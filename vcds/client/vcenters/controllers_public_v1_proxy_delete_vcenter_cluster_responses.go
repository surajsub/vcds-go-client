// Code generated by go-swagger; DO NOT EDIT.

package vcenters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/vcds-go-client/vcds/models"
)

// ControllersPublicV1ProxyDeleteVcenterClusterReader is a Reader for the ControllersPublicV1ProxyDeleteVcenterCluster structure.
type ControllersPublicV1ProxyDeleteVcenterClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControllersPublicV1ProxyDeleteVcenterClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewControllersPublicV1ProxyDeleteVcenterClusterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewControllersPublicV1ProxyDeleteVcenterClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewControllersPublicV1ProxyDeleteVcenterClusterAccepted creates a ControllersPublicV1ProxyDeleteVcenterClusterAccepted with default headers values
func NewControllersPublicV1ProxyDeleteVcenterClusterAccepted() *ControllersPublicV1ProxyDeleteVcenterClusterAccepted {
	return &ControllersPublicV1ProxyDeleteVcenterClusterAccepted{}
}

/*ControllersPublicV1ProxyDeleteVcenterClusterAccepted handles this case with default header values.

Success. The request for deleting the cluster has been accepted.
*/
type ControllersPublicV1ProxyDeleteVcenterClusterAccepted struct {
}

func (o *ControllersPublicV1ProxyDeleteVcenterClusterAccepted) Error() string {
	return fmt.Sprintf("[DELETE /v1/vcenters/{instance_id}/clusters/{cluster_id}][%d] controllersPublicV1ProxyDeleteVcenterClusterAccepted ", 202)
}

func (o *ControllersPublicV1ProxyDeleteVcenterClusterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewControllersPublicV1ProxyDeleteVcenterClusterUnauthorized creates a ControllersPublicV1ProxyDeleteVcenterClusterUnauthorized with default headers values
func NewControllersPublicV1ProxyDeleteVcenterClusterUnauthorized() *ControllersPublicV1ProxyDeleteVcenterClusterUnauthorized {
	return &ControllersPublicV1ProxyDeleteVcenterClusterUnauthorized{}
}

/*ControllersPublicV1ProxyDeleteVcenterClusterUnauthorized handles this case with default header values.

Unauthorized. The IAM token is invalid or expired. To retrieve your IAM token, run `ibmcloud login` and then `ibmcloud iam oauth-tokens`.
*/
type ControllersPublicV1ProxyDeleteVcenterClusterUnauthorized struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyDeleteVcenterClusterUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/vcenters/{instance_id}/clusters/{cluster_id}][%d] controllersPublicV1ProxyDeleteVcenterClusterUnauthorized  %+v", 401, o.Payload)
}

func (o *ControllersPublicV1ProxyDeleteVcenterClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
