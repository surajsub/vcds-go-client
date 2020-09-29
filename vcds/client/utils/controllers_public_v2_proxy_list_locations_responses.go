// Code generated by go-swagger; DO NOT EDIT.

package utils

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/vcds-go-client/vcds/models"
)

// ControllersPublicV2ProxyListLocationsReader is a Reader for the ControllersPublicV2ProxyListLocations structure.
type ControllersPublicV2ProxyListLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControllersPublicV2ProxyListLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewControllersPublicV2ProxyListLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewControllersPublicV2ProxyListLocationsOK creates a ControllersPublicV2ProxyListLocationsOK with default headers values
func NewControllersPublicV2ProxyListLocationsOK() *ControllersPublicV2ProxyListLocationsOK {
	return &ControllersPublicV2ProxyListLocationsOK{}
}

/*ControllersPublicV2ProxyListLocationsOK handles this case with default header values.

Success.
*/
type ControllersPublicV2ProxyListLocationsOK struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.DefLocations
}

func (o *ControllersPublicV2ProxyListLocationsOK) Error() string {
	return fmt.Sprintf("[GET /v2/locations][%d] controllersPublicV2ProxyListLocationsOK  %+v", 200, o.Payload)
}

func (o *ControllersPublicV2ProxyListLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.DefLocations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}