// Code generated by go-swagger; DO NOT EDIT.

package vcenters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/vcds-go-client/vcds/models"
)

// ControllersPublicV1ProxyAddVcenterNfsStorageReader is a Reader for the ControllersPublicV1ProxyAddVcenterNfsStorage structure.
type ControllersPublicV1ProxyAddVcenterNfsStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControllersPublicV1ProxyAddVcenterNfsStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewControllersPublicV1ProxyAddVcenterNfsStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewControllersPublicV1ProxyAddVcenterNfsStorageAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewControllersPublicV1ProxyAddVcenterNfsStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewControllersPublicV1ProxyAddVcenterNfsStorageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewControllersPublicV1ProxyAddVcenterNfsStorageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewControllersPublicV1ProxyAddVcenterNfsStorageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewControllersPublicV1ProxyAddVcenterNfsStorageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewControllersPublicV1ProxyAddVcenterNfsStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewControllersPublicV1ProxyAddVcenterNfsStorageOK creates a ControllersPublicV1ProxyAddVcenterNfsStorageOK with default headers values
func NewControllersPublicV1ProxyAddVcenterNfsStorageOK() *ControllersPublicV1ProxyAddVcenterNfsStorageOK {
	return &ControllersPublicV1ProxyAddVcenterNfsStorageOK{}
}

/*ControllersPublicV1ProxyAddVcenterNfsStorageOK handles this case with default header values.

Success. The request has been successfully verified.
*/
type ControllersPublicV1ProxyAddVcenterNfsStorageOK struct {
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageOK) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/clusters/{cluster_id}/shared_storages][%d] controllersPublicV1ProxyAddVcenterNfsStorageOK ", 200)
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewControllersPublicV1ProxyAddVcenterNfsStorageAccepted creates a ControllersPublicV1ProxyAddVcenterNfsStorageAccepted with default headers values
func NewControllersPublicV1ProxyAddVcenterNfsStorageAccepted() *ControllersPublicV1ProxyAddVcenterNfsStorageAccepted {
	return &ControllersPublicV1ProxyAddVcenterNfsStorageAccepted{}
}

/*ControllersPublicV1ProxyAddVcenterNfsStorageAccepted handles this case with default header values.

Success. The request for adding new NFS storage has been accepted.
*/
type ControllersPublicV1ProxyAddVcenterNfsStorageAccepted struct {
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/clusters/{cluster_id}/shared_storages][%d] controllersPublicV1ProxyAddVcenterNfsStorageAccepted ", 202)
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewControllersPublicV1ProxyAddVcenterNfsStorageBadRequest creates a ControllersPublicV1ProxyAddVcenterNfsStorageBadRequest with default headers values
func NewControllersPublicV1ProxyAddVcenterNfsStorageBadRequest() *ControllersPublicV1ProxyAddVcenterNfsStorageBadRequest {
	return &ControllersPublicV1ProxyAddVcenterNfsStorageBadRequest{}
}

/*ControllersPublicV1ProxyAddVcenterNfsStorageBadRequest handles this case with default header values.

Bad request. Check your request parameters.
*/
type ControllersPublicV1ProxyAddVcenterNfsStorageBadRequest struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/clusters/{cluster_id}/shared_storages][%d] controllersPublicV1ProxyAddVcenterNfsStorageBadRequest  %+v", 400, o.Payload)
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyAddVcenterNfsStorageUnauthorized creates a ControllersPublicV1ProxyAddVcenterNfsStorageUnauthorized with default headers values
func NewControllersPublicV1ProxyAddVcenterNfsStorageUnauthorized() *ControllersPublicV1ProxyAddVcenterNfsStorageUnauthorized {
	return &ControllersPublicV1ProxyAddVcenterNfsStorageUnauthorized{}
}

/*ControllersPublicV1ProxyAddVcenterNfsStorageUnauthorized handles this case with default header values.

Unauthorized. The IAM token is invalid or expired. To retrieve your IAM token, run `ibmcloud login` and then `ibmcloud iam oauth-tokens`.
*/
type ControllersPublicV1ProxyAddVcenterNfsStorageUnauthorized struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/clusters/{cluster_id}/shared_storages][%d] controllersPublicV1ProxyAddVcenterNfsStorageUnauthorized  %+v", 401, o.Payload)
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyAddVcenterNfsStorageForbidden creates a ControllersPublicV1ProxyAddVcenterNfsStorageForbidden with default headers values
func NewControllersPublicV1ProxyAddVcenterNfsStorageForbidden() *ControllersPublicV1ProxyAddVcenterNfsStorageForbidden {
	return &ControllersPublicV1ProxyAddVcenterNfsStorageForbidden{}
}

/*ControllersPublicV1ProxyAddVcenterNfsStorageForbidden handles this case with default header values.

Forbidden. Access to the specified resource is not authorized. Check the IAM access policy for the `VMware Solutions` service.
*/
type ControllersPublicV1ProxyAddVcenterNfsStorageForbidden struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/clusters/{cluster_id}/shared_storages][%d] controllersPublicV1ProxyAddVcenterNfsStorageForbidden  %+v", 403, o.Payload)
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyAddVcenterNfsStorageNotFound creates a ControllersPublicV1ProxyAddVcenterNfsStorageNotFound with default headers values
func NewControllersPublicV1ProxyAddVcenterNfsStorageNotFound() *ControllersPublicV1ProxyAddVcenterNfsStorageNotFound {
	return &ControllersPublicV1ProxyAddVcenterNfsStorageNotFound{}
}

/*ControllersPublicV1ProxyAddVcenterNfsStorageNotFound handles this case with default header values.

Not found. The resource cannot be found.
*/
type ControllersPublicV1ProxyAddVcenterNfsStorageNotFound struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/clusters/{cluster_id}/shared_storages][%d] controllersPublicV1ProxyAddVcenterNfsStorageNotFound  %+v", 404, o.Payload)
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyAddVcenterNfsStorageConflict creates a ControllersPublicV1ProxyAddVcenterNfsStorageConflict with default headers values
func NewControllersPublicV1ProxyAddVcenterNfsStorageConflict() *ControllersPublicV1ProxyAddVcenterNfsStorageConflict {
	return &ControllersPublicV1ProxyAddVcenterNfsStorageConflict{}
}

/*ControllersPublicV1ProxyAddVcenterNfsStorageConflict handles this case with default header values.

Conflict. The request cannot be completed because of a conflict with the current state of the target resource.
*/
type ControllersPublicV1ProxyAddVcenterNfsStorageConflict struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageConflict) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/clusters/{cluster_id}/shared_storages][%d] controllersPublicV1ProxyAddVcenterNfsStorageConflict  %+v", 409, o.Payload)
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyAddVcenterNfsStorageInternalServerError creates a ControllersPublicV1ProxyAddVcenterNfsStorageInternalServerError with default headers values
func NewControllersPublicV1ProxyAddVcenterNfsStorageInternalServerError() *ControllersPublicV1ProxyAddVcenterNfsStorageInternalServerError {
	return &ControllersPublicV1ProxyAddVcenterNfsStorageInternalServerError{}
}

/*ControllersPublicV1ProxyAddVcenterNfsStorageInternalServerError handles this case with default header values.

Internal server error. Your request cannot be processed. Please wait a few minutes and try again.
*/
type ControllersPublicV1ProxyAddVcenterNfsStorageInternalServerError struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/clusters/{cluster_id}/shared_storages][%d] controllersPublicV1ProxyAddVcenterNfsStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ControllersPublicV1ProxyAddVcenterNfsStorageBody controllers public v1 proxy add vcenter nfs storage body
swagger:model ControllersPublicV1ProxyAddVcenterNfsStorageBody
*/
type ControllersPublicV1ProxyAddVcenterNfsStorageBody struct {

	// shared storages
	// Required: true
	SharedStorages []*models.SharedStorageConfig `json:"shared_storages"`
}

// Validate validates this controllers public v1 proxy add vcenter nfs storage body
func (o *ControllersPublicV1ProxyAddVcenterNfsStorageBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSharedStorages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ControllersPublicV1ProxyAddVcenterNfsStorageBody) validateSharedStorages(formats strfmt.Registry) error {

	if err := validate.Required("nfs_order_data"+"."+"shared_storages", "body", o.SharedStorages); err != nil {
		return err
	}

	for i := 0; i < len(o.SharedStorages); i++ {
		if swag.IsZero(o.SharedStorages[i]) { // not required
			continue
		}

		if o.SharedStorages[i] != nil {
			if err := o.SharedStorages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nfs_order_data" + "." + "shared_storages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ControllersPublicV1ProxyAddVcenterNfsStorageBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ControllersPublicV1ProxyAddVcenterNfsStorageBody) UnmarshalBinary(b []byte) error {
	var res ControllersPublicV1ProxyAddVcenterNfsStorageBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
