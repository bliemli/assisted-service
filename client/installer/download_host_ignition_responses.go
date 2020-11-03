// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// DownloadHostIgnitionReader is a Reader for the DownloadHostIgnition structure.
type DownloadHostIgnitionReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *DownloadHostIgnitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadHostIgnitionOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDownloadHostIgnitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDownloadHostIgnitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadHostIgnitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDownloadHostIgnitionMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDownloadHostIgnitionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDownloadHostIgnitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDownloadHostIgnitionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDownloadHostIgnitionOK creates a DownloadHostIgnitionOK with default headers values
func NewDownloadHostIgnitionOK(writer io.Writer) *DownloadHostIgnitionOK {
	return &DownloadHostIgnitionOK{
		Payload: writer,
	}
}

/*DownloadHostIgnitionOK handles this case with default header values.

Success.
*/
type DownloadHostIgnitionOK struct {
	Payload io.Writer
}

func (o *DownloadHostIgnitionOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/downloads/ignition][%d] downloadHostIgnitionOK  %+v", 200, o.Payload)
}

func (o *DownloadHostIgnitionOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *DownloadHostIgnitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadHostIgnitionUnauthorized creates a DownloadHostIgnitionUnauthorized with default headers values
func NewDownloadHostIgnitionUnauthorized() *DownloadHostIgnitionUnauthorized {
	return &DownloadHostIgnitionUnauthorized{}
}

/*DownloadHostIgnitionUnauthorized handles this case with default header values.

Unauthorized.
*/
type DownloadHostIgnitionUnauthorized struct {
	Payload *models.InfraError
}

func (o *DownloadHostIgnitionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/downloads/ignition][%d] downloadHostIgnitionUnauthorized  %+v", 401, o.Payload)
}

func (o *DownloadHostIgnitionUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *DownloadHostIgnitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadHostIgnitionForbidden creates a DownloadHostIgnitionForbidden with default headers values
func NewDownloadHostIgnitionForbidden() *DownloadHostIgnitionForbidden {
	return &DownloadHostIgnitionForbidden{}
}

/*DownloadHostIgnitionForbidden handles this case with default header values.

Forbidden.
*/
type DownloadHostIgnitionForbidden struct {
	Payload *models.InfraError
}

func (o *DownloadHostIgnitionForbidden) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/downloads/ignition][%d] downloadHostIgnitionForbidden  %+v", 403, o.Payload)
}

func (o *DownloadHostIgnitionForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *DownloadHostIgnitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadHostIgnitionNotFound creates a DownloadHostIgnitionNotFound with default headers values
func NewDownloadHostIgnitionNotFound() *DownloadHostIgnitionNotFound {
	return &DownloadHostIgnitionNotFound{}
}

/*DownloadHostIgnitionNotFound handles this case with default header values.

Error.
*/
type DownloadHostIgnitionNotFound struct {
	Payload *models.Error
}

func (o *DownloadHostIgnitionNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/downloads/ignition][%d] downloadHostIgnitionNotFound  %+v", 404, o.Payload)
}

func (o *DownloadHostIgnitionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadHostIgnitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadHostIgnitionMethodNotAllowed creates a DownloadHostIgnitionMethodNotAllowed with default headers values
func NewDownloadHostIgnitionMethodNotAllowed() *DownloadHostIgnitionMethodNotAllowed {
	return &DownloadHostIgnitionMethodNotAllowed{}
}

/*DownloadHostIgnitionMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type DownloadHostIgnitionMethodNotAllowed struct {
	Payload *models.Error
}

func (o *DownloadHostIgnitionMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/downloads/ignition][%d] downloadHostIgnitionMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *DownloadHostIgnitionMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadHostIgnitionMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadHostIgnitionConflict creates a DownloadHostIgnitionConflict with default headers values
func NewDownloadHostIgnitionConflict() *DownloadHostIgnitionConflict {
	return &DownloadHostIgnitionConflict{}
}

/*DownloadHostIgnitionConflict handles this case with default header values.

Error.
*/
type DownloadHostIgnitionConflict struct {
	Payload *models.Error
}

func (o *DownloadHostIgnitionConflict) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/downloads/ignition][%d] downloadHostIgnitionConflict  %+v", 409, o.Payload)
}

func (o *DownloadHostIgnitionConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadHostIgnitionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadHostIgnitionInternalServerError creates a DownloadHostIgnitionInternalServerError with default headers values
func NewDownloadHostIgnitionInternalServerError() *DownloadHostIgnitionInternalServerError {
	return &DownloadHostIgnitionInternalServerError{}
}

/*DownloadHostIgnitionInternalServerError handles this case with default header values.

Error.
*/
type DownloadHostIgnitionInternalServerError struct {
	Payload *models.Error
}

func (o *DownloadHostIgnitionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/downloads/ignition][%d] downloadHostIgnitionInternalServerError  %+v", 500, o.Payload)
}

func (o *DownloadHostIgnitionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadHostIgnitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadHostIgnitionServiceUnavailable creates a DownloadHostIgnitionServiceUnavailable with default headers values
func NewDownloadHostIgnitionServiceUnavailable() *DownloadHostIgnitionServiceUnavailable {
	return &DownloadHostIgnitionServiceUnavailable{}
}

/*DownloadHostIgnitionServiceUnavailable handles this case with default header values.

Unavailable.
*/
type DownloadHostIgnitionServiceUnavailable struct {
	Payload *models.Error
}

func (o *DownloadHostIgnitionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/hosts/{host_id}/downloads/ignition][%d] downloadHostIgnitionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DownloadHostIgnitionServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadHostIgnitionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
