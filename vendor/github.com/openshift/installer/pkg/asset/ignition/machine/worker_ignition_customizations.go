package machine

import (
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sigs.k8s.io/yaml"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/ignition"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/templates/content/bootkube"
	"github.com/openshift/installer/pkg/asset/tls"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
)

var (
	workerMachineConfigFileName = filepath.Join(directory, "99_openshift-installer-ignition_worker.yaml")
)

// WorkerIgnitionCustomizations is an asset that checks for any customizations a user might
// have made to the pointer ignition configs before creating the cluster. If customizations
// are made, then the updates are reconciled as a MachineConfig file
type WorkerIgnitionCustomizations struct {
	File *asset.File
}

var _ asset.WritableAsset = (*WorkerIgnitionCustomizations)(nil)

// Dependencies returns the dependencies for WorkerIgnitionCustomizations
func (a *WorkerIgnitionCustomizations) Dependencies() []asset.Asset {
	return []asset.Asset{
		&installconfig.InstallConfig{},
		&tls.RootCA{},
		&Worker{},
		&bootkube.ARODNSConfig{},
	}
}

// Generate queries for input from the user.
func (a *WorkerIgnitionCustomizations) Generate(dependencies asset.Parents) error {
	installConfig := &installconfig.InstallConfig{}
	rootCA := &tls.RootCA{}
	worker := &Worker{}
	aroDNSConfig := &bootkube.ARODNSConfig{}
	dependencies.Get(installConfig, rootCA, worker, aroDNSConfig)

	defaultPointerIgnition := pointerIgnitionConfig(installConfig.Config, aroDNSConfig, rootCA.Cert(), "worker")
	savedPointerIgnition := worker.Config

	// Create a machineconfig if the ignition has been modified
	savedPointerIgnitionJSON, err := ignition.Marshal(savedPointerIgnition)
	if err != nil {
		return errors.Wrap(err, "failed Marshal savedPointerIgnition")
	}
	defaultPointerIgnitionJSON, err := ignition.Marshal(defaultPointerIgnition)
	if err != nil {
		return errors.Wrap(err, "failed Marshal defaultPointerIgnition")
	}
	if string(savedPointerIgnitionJSON) != string(defaultPointerIgnitionJSON) {
		logrus.Infof("Worker pointer ignition was modified. Saving contents to a machineconfig")
		mc := &mcfgv1.MachineConfig{}
		mc, err = generatePointerMachineConfig(*savedPointerIgnition, "worker")
		if err != nil {
			return errors.Wrap(err, "failed to generate worker installer machineconfig")
		}
		configData, err := yaml.Marshal(mc)
		if err != nil {
			return errors.Wrap(err, "failed to marshal worker installer machineconfig")
		}
		a.File = &asset.File{
			Filename: workerMachineConfigFileName,
			Data:     configData,
		}
	}

	return nil
}

// Name returns the human-friendly name of the asset.
func (a *WorkerIgnitionCustomizations) Name() string {
	return "Worker Ignition Customization Check"
}

// Files returns the files generated by the asset.
func (a *WorkerIgnitionCustomizations) Files() []*asset.File {
	if a.File != nil {
		return []*asset.File{a.File}
	}
	return []*asset.File{}
}

// Load does nothing, since we consume the ignition-configs
func (a *WorkerIgnitionCustomizations) Load(f asset.FileFetcher) (found bool, err error) {
	return false, nil
}