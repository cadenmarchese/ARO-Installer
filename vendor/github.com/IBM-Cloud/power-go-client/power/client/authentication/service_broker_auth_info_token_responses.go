// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBrokerAuthInfoTokenReader is a Reader for the ServiceBrokerAuthInfoToken structure.
type ServiceBrokerAuthInfoTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerAuthInfoTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerAuthInfoTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerAuthInfoTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceBrokerAuthInfoTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServiceBrokerAuthInfoTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceBrokerAuthInfoTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerAuthInfoTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /auth/v1/info/token] serviceBroker.auth.info.token", response, response.Code())
	}
}

// NewServiceBrokerAuthInfoTokenOK creates a ServiceBrokerAuthInfoTokenOK with default headers values
func NewServiceBrokerAuthInfoTokenOK() *ServiceBrokerAuthInfoTokenOK {
	return &ServiceBrokerAuthInfoTokenOK{}
}

/*
ServiceBrokerAuthInfoTokenOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerAuthInfoTokenOK struct {
	Payload *models.TokenExtra
}

// IsSuccess returns true when this service broker auth info token o k response has a 2xx status code
func (o *ServiceBrokerAuthInfoTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker auth info token o k response has a 3xx status code
func (o *ServiceBrokerAuthInfoTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth info token o k response has a 4xx status code
func (o *ServiceBrokerAuthInfoTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker auth info token o k response has a 5xx status code
func (o *ServiceBrokerAuthInfoTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth info token o k response a status code equal to that given
func (o *ServiceBrokerAuthInfoTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker auth info token o k response
func (o *ServiceBrokerAuthInfoTokenOK) Code() int {
	return 200
}

func (o *ServiceBrokerAuthInfoTokenOK) Error() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenOK) String() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenOK) GetPayload() *models.TokenExtra {
	return o.Payload
}

func (o *ServiceBrokerAuthInfoTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenExtra)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthInfoTokenBadRequest creates a ServiceBrokerAuthInfoTokenBadRequest with default headers values
func NewServiceBrokerAuthInfoTokenBadRequest() *ServiceBrokerAuthInfoTokenBadRequest {
	return &ServiceBrokerAuthInfoTokenBadRequest{}
}

/*
ServiceBrokerAuthInfoTokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBrokerAuthInfoTokenBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth info token bad request response has a 2xx status code
func (o *ServiceBrokerAuthInfoTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth info token bad request response has a 3xx status code
func (o *ServiceBrokerAuthInfoTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth info token bad request response has a 4xx status code
func (o *ServiceBrokerAuthInfoTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker auth info token bad request response has a 5xx status code
func (o *ServiceBrokerAuthInfoTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth info token bad request response a status code equal to that given
func (o *ServiceBrokerAuthInfoTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service broker auth info token bad request response
func (o *ServiceBrokerAuthInfoTokenBadRequest) Code() int {
	return 400
}

func (o *ServiceBrokerAuthInfoTokenBadRequest) Error() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenBadRequest) String() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthInfoTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthInfoTokenUnauthorized creates a ServiceBrokerAuthInfoTokenUnauthorized with default headers values
func NewServiceBrokerAuthInfoTokenUnauthorized() *ServiceBrokerAuthInfoTokenUnauthorized {
	return &ServiceBrokerAuthInfoTokenUnauthorized{}
}

/*
ServiceBrokerAuthInfoTokenUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBrokerAuthInfoTokenUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth info token unauthorized response has a 2xx status code
func (o *ServiceBrokerAuthInfoTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth info token unauthorized response has a 3xx status code
func (o *ServiceBrokerAuthInfoTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth info token unauthorized response has a 4xx status code
func (o *ServiceBrokerAuthInfoTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker auth info token unauthorized response has a 5xx status code
func (o *ServiceBrokerAuthInfoTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth info token unauthorized response a status code equal to that given
func (o *ServiceBrokerAuthInfoTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the service broker auth info token unauthorized response
func (o *ServiceBrokerAuthInfoTokenUnauthorized) Code() int {
	return 401
}

func (o *ServiceBrokerAuthInfoTokenUnauthorized) Error() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenUnauthorized) String() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthInfoTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthInfoTokenForbidden creates a ServiceBrokerAuthInfoTokenForbidden with default headers values
func NewServiceBrokerAuthInfoTokenForbidden() *ServiceBrokerAuthInfoTokenForbidden {
	return &ServiceBrokerAuthInfoTokenForbidden{}
}

/*
ServiceBrokerAuthInfoTokenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServiceBrokerAuthInfoTokenForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth info token forbidden response has a 2xx status code
func (o *ServiceBrokerAuthInfoTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth info token forbidden response has a 3xx status code
func (o *ServiceBrokerAuthInfoTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth info token forbidden response has a 4xx status code
func (o *ServiceBrokerAuthInfoTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker auth info token forbidden response has a 5xx status code
func (o *ServiceBrokerAuthInfoTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth info token forbidden response a status code equal to that given
func (o *ServiceBrokerAuthInfoTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the service broker auth info token forbidden response
func (o *ServiceBrokerAuthInfoTokenForbidden) Code() int {
	return 403
}

func (o *ServiceBrokerAuthInfoTokenForbidden) Error() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenForbidden) String() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthInfoTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthInfoTokenNotFound creates a ServiceBrokerAuthInfoTokenNotFound with default headers values
func NewServiceBrokerAuthInfoTokenNotFound() *ServiceBrokerAuthInfoTokenNotFound {
	return &ServiceBrokerAuthInfoTokenNotFound{}
}

/*
ServiceBrokerAuthInfoTokenNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServiceBrokerAuthInfoTokenNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth info token not found response has a 2xx status code
func (o *ServiceBrokerAuthInfoTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth info token not found response has a 3xx status code
func (o *ServiceBrokerAuthInfoTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth info token not found response has a 4xx status code
func (o *ServiceBrokerAuthInfoTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker auth info token not found response has a 5xx status code
func (o *ServiceBrokerAuthInfoTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth info token not found response a status code equal to that given
func (o *ServiceBrokerAuthInfoTokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the service broker auth info token not found response
func (o *ServiceBrokerAuthInfoTokenNotFound) Code() int {
	return 404
}

func (o *ServiceBrokerAuthInfoTokenNotFound) Error() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenNotFound) String() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthInfoTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthInfoTokenInternalServerError creates a ServiceBrokerAuthInfoTokenInternalServerError with default headers values
func NewServiceBrokerAuthInfoTokenInternalServerError() *ServiceBrokerAuthInfoTokenInternalServerError {
	return &ServiceBrokerAuthInfoTokenInternalServerError{}
}

/*
ServiceBrokerAuthInfoTokenInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerAuthInfoTokenInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth info token internal server error response has a 2xx status code
func (o *ServiceBrokerAuthInfoTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth info token internal server error response has a 3xx status code
func (o *ServiceBrokerAuthInfoTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth info token internal server error response has a 4xx status code
func (o *ServiceBrokerAuthInfoTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker auth info token internal server error response has a 5xx status code
func (o *ServiceBrokerAuthInfoTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker auth info token internal server error response a status code equal to that given
func (o *ServiceBrokerAuthInfoTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the service broker auth info token internal server error response
func (o *ServiceBrokerAuthInfoTokenInternalServerError) Code() int {
	return 500
}

func (o *ServiceBrokerAuthInfoTokenInternalServerError) Error() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenInternalServerError) String() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthInfoTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}