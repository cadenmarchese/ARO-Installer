// Package manifests deals with creating manifests for all manifests to be installed for the cluster
package manifests

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/vincent-petithory/dataurl"
	"sigs.k8s.io/yaml"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/templates/content/bootkube"
	"github.com/openshift/installer/pkg/asset/tls"
	"github.com/openshift/installer/pkg/types"
	azuretypes "github.com/openshift/installer/pkg/types/azure"
	"github.com/openshift/installer/pkg/types/vsphere"
)

const (
	manifestDir = "manifests"
)

var (
	kubeSysConfigPath = filepath.Join(manifestDir, "cluster-config.yaml")

	_ asset.WritableAsset = (*Manifests)(nil)

	customTmplFuncs = template.FuncMap{
		"indent": indent,
		"add": func(i, j int) int {
			return i + j
		},
	}
)

// Manifests generates the dependent operator config.yaml files
type Manifests struct {
	KubeSysConfig *configurationObject
	FileList      []*asset.File
}

type genericData map[string]string

// Name returns a human friendly name for the operator
func (m *Manifests) Name() string {
	return "Common Manifests"
}

// Dependencies returns all of the dependencies directly needed by a
// Manifests asset.
func (m *Manifests) Dependencies() []asset.Asset {
	return []asset.Asset{
		&installconfig.ClusterID{},
		&installconfig.InstallConfig{},
		&Ingress{},
		&DNS{},
		&Infrastructure{},
		&Networking{},
		&Proxy{},
		&Scheduler{},
		&ImageContentSourcePolicy{},
		&ClusterCSIDriverConfig{},
		&ImageDigestMirrorSet{},
		&tls.RootCA{},
		&tls.MCSCertKey{},

		&bootkube.CVOOverrides{},
		&bootkube.KubeCloudConfig{},
		&bootkube.KubeSystemConfigmapRootCA{},
		&bootkube.MachineConfigServerTLSSecret{},
		&bootkube.OpenshiftConfigSecretPullSecret{},
		&bootkube.AROWorkerRegistries{},
		&bootkube.AROIngressService{},
		&bootkube.ARODNSConfig{},
		&bootkube.AROImageRegistry{},
		&bootkube.AROImageRegistryConfig{},
	}
}

// Generate generates the respective operator config.yml files
func (m *Manifests) Generate(dependencies asset.Parents) error {
	ingress := &Ingress{}
	dns := &DNS{}
	network := &Networking{}
	infra := &Infrastructure{}
	installConfig := &installconfig.InstallConfig{}
	proxy := &Proxy{}
	scheduler := &Scheduler{}
	imageContentSourcePolicy := &ImageContentSourcePolicy{}
	clusterCSIDriverConfig := &ClusterCSIDriverConfig{}
	imageDigestMirrorSet := &ImageDigestMirrorSet{}

	dependencies.Get(installConfig, ingress, dns, network, infra, proxy, scheduler, imageContentSourcePolicy, imageDigestMirrorSet, clusterCSIDriverConfig)

	redactedConfig, err := redactedInstallConfig(*installConfig.Config)
	if err != nil {
		return errors.Wrap(err, "failed to redact install-config")
	}
	// mao go to kube-system config map
	m.KubeSysConfig = configMap("kube-system", "cluster-config-v1", genericData{
		"install-config": string(redactedConfig),
	})
	if m.KubeSysConfig.Metadata.Annotations == nil {
		m.KubeSysConfig.Metadata.Annotations = make(map[string]string, 1)
	}
	m.KubeSysConfig.Metadata.Annotations["kubernetes.io/description"] = "The install-config content used to create the cluster.  The cluster configuration may have evolved since installation, so check cluster configuration resources directly if you are interested in the current cluster state."

	kubeSysConfigData, err := yaml.Marshal(m.KubeSysConfig)
	if err != nil {
		return errors.Wrap(err, "failed to create kube-system/cluster-config-v1 configmap")
	}

	m.FileList = []*asset.File{
		{
			Filename: kubeSysConfigPath,
			Data:     kubeSysConfigData,
		},
	}
	m.FileList = append(m.FileList, m.generateBootKubeManifests(dependencies)...)

	m.FileList = append(m.FileList, ingress.Files()...)
	m.FileList = append(m.FileList, dns.Files()...)
	m.FileList = append(m.FileList, network.Files()...)
	m.FileList = append(m.FileList, infra.Files()...)
	m.FileList = append(m.FileList, proxy.Files()...)
	m.FileList = append(m.FileList, scheduler.Files()...)
	m.FileList = append(m.FileList, imageContentSourcePolicy.Files()...)
	m.FileList = append(m.FileList, clusterCSIDriverConfig.Files()...)
	m.FileList = append(m.FileList, imageDigestMirrorSet.Files()...)

	asset.SortFiles(m.FileList)

	return nil
}

// Files returns the files generated by the asset.
func (m *Manifests) Files() []*asset.File {
	return m.FileList
}

func (m *Manifests) generateBootKubeManifests(dependencies asset.Parents) []*asset.File {
	clusterID := &installconfig.ClusterID{}
	installConfig := &installconfig.InstallConfig{}
	mcsCertKey := &tls.MCSCertKey{}
	rootCA := &tls.RootCA{}
	aroDNSConfig := &bootkube.ARODNSConfig{}
	aroImageRegistryConfig := &bootkube.AROImageRegistryConfig{}
	dependencies.Get(
		clusterID,
		installConfig,
		mcsCertKey,
		rootCA,
		aroDNSConfig,
		aroImageRegistryConfig,
	)

	templateData := &bootkubeTemplateData{
		CVOCapabilities:               installConfig.Config.Capabilities,
		CVOClusterID:                  clusterID.UUID,
		McsTLSCert:                    base64.StdEncoding.EncodeToString(mcsCertKey.Cert()),
		McsTLSKey:                     base64.StdEncoding.EncodeToString(mcsCertKey.Key()),
		PullSecretBase64:              base64.StdEncoding.EncodeToString([]byte(installConfig.Config.PullSecret)),
		RootCaCert:                    string(rootCA.Cert()),
		IsFCOS:                        installConfig.Config.IsFCOS(),
		IsSCOS:                        installConfig.Config.IsSCOS(),
		IsOKD:                         installConfig.Config.IsOKD(),
		AROWorkerRegistries:           aroWorkerRegistries(installConfig.Config.ImageDigestSources),
		AROIngressIP:                  aroDNSConfig.IngressIP,
		AROIngressInternal:            installConfig.Config.Publish == types.InternalPublishingStrategy,
		AROImageRegistryHTTPSecret:    aroImageRegistryConfig.HTTPSecret,
		AROImageRegistryAccountName:   aroImageRegistryConfig.AccountName,
		AROImageRegistryContainerName: aroImageRegistryConfig.ContainerName,
	}

	switch installConfig.Config.Platform.Name() {
	case azuretypes.Name:
		templateData.AROCloudName = installConfig.Azure.CloudName.Name()
	}

	files := []*asset.File{}
	for _, a := range []asset.WritableAsset{
		&bootkube.CVOOverrides{},
		&bootkube.KubeCloudConfig{},
		&bootkube.KubeSystemConfigmapRootCA{},
		&bootkube.MachineConfigServerTLSSecret{},
		&bootkube.OpenshiftConfigSecretPullSecret{},
		&bootkube.AROWorkerRegistries{},
		&bootkube.AROIngressService{},
		&bootkube.AROImageRegistry{},
	} {
		dependencies.Get(a)
		for _, f := range a.Files() {
			files = append(files, &asset.File{
				Filename: filepath.Join(manifestDir, strings.TrimSuffix(filepath.Base(f.Filename), ".template")),
				Data:     applyTemplateData(f.Data, templateData),
			})
		}
	}
	return files
}

func applyTemplateData(data []byte, templateData interface{}) []byte {
	template := template.Must(template.New("template").Funcs(customTmplFuncs).Parse(string(data)))
	buf := &bytes.Buffer{}
	if err := template.Execute(buf, templateData); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// Load returns the manifests asset from disk.
func (m *Manifests) Load(f asset.FileFetcher) (bool, error) {
	yamlFileList, err := f.FetchByPattern(filepath.Join(manifestDir, "*.yaml"))
	if err != nil {
		return false, errors.Wrap(err, "failed to load *.yaml files")
	}
	ymlFileList, err := f.FetchByPattern(filepath.Join(manifestDir, "*.yml"))
	if err != nil {
		return false, errors.Wrap(err, "failed to load *.yml files")
	}
	jsonFileList, err := f.FetchByPattern(filepath.Join(manifestDir, "*.json"))
	if err != nil {
		return false, errors.Wrap(err, "failed to load *.json files")
	}
	fileList := append(yamlFileList, ymlFileList...)
	fileList = append(fileList, jsonFileList...)

	if len(fileList) == 0 {
		return false, nil
	}

	kubeSysConfig := &configurationObject{}
	var found bool
	for _, file := range fileList {
		if file.Filename == kubeSysConfigPath {
			if err := yaml.Unmarshal(file.Data, kubeSysConfig); err != nil {
				return false, errors.Wrapf(err, "failed to unmarshal %s", kubeSysConfigPath)
			}
			found = true
		}
	}

	if !found {
		return false, nil

	}

	m.FileList, m.KubeSysConfig = fileList, kubeSysConfig

	asset.SortFiles(m.FileList)

	return true, nil
}

func redactedInstallConfig(config types.InstallConfig) ([]byte, error) {
	newConfig := config

	newConfig.PullSecret = ""
	if newConfig.Platform.VSphere != nil {
		p := config.VSphere
		newVCenters := make([]vsphere.VCenter, len(p.VCenters))
		for i, v := range p.VCenters {
			newVCenters[i].Server = v.Server
			newVCenters[i].Datacenters = v.Datacenters
		}
		newVSpherePlatform := vsphere.Platform{
			DeprecatedVCenter:          p.DeprecatedVCenter,
			DeprecatedUsername:         "",
			DeprecatedPassword:         "",
			DeprecatedDatacenter:       p.DeprecatedDatacenter,
			DeprecatedDefaultDatastore: p.DeprecatedDefaultDatastore,
			DeprecatedFolder:           p.DeprecatedFolder,
			DeprecatedCluster:          p.DeprecatedCluster,
			DeprecatedResourcePool:     p.DeprecatedResourcePool,
			ClusterOSImage:             p.ClusterOSImage,
			DeprecatedAPIVIP:           p.DeprecatedAPIVIP,
			APIVIPs:                    p.APIVIPs,
			DeprecatedIngressVIP:       p.DeprecatedIngressVIP,
			IngressVIPs:                p.IngressVIPs,
			DefaultMachinePlatform:     p.DefaultMachinePlatform,
			DeprecatedNetwork:          p.DeprecatedNetwork,
			DiskType:                   p.DiskType,
			VCenters:                   newVCenters,
			FailureDomains:             p.FailureDomains,
		}
		newConfig.Platform.VSphere = &newVSpherePlatform
	}

	return yaml.Marshal(newConfig)
}

func indent(indention int, v string) string {
	newline := "\n" + strings.Repeat(" ", indention)
	return strings.Replace(v, "\n", newline, -1)
}

func aroWorkerRegistries(idss []types.ImageDigestSource) string {
	b := &bytes.Buffer{}

	fmt.Fprintf(b, "unqualified-search-registries = [\"registry.access.redhat.com\", \"docker.io\"]\n")

	for _, ids := range idss {
		fmt.Fprintf(b, "\n[[registry]]\n  prefix = \"\"\n  location = \"%s\"\n  mirror-by-digest-only = true\n", ids.Source)

		for _, mirror := range ids.Mirrors {
			fmt.Fprintf(b, "\n  [[registry.mirror]]\n    location = \"%s\"\n", mirror)
		}
	}

	du := dataurl.New(b.Bytes(), "text/plain")
	du.Encoding = dataurl.EncodingASCII

	return du.String()
}