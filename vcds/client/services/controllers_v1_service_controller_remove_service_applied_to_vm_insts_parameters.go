// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams creates a new ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams object
// with the default values initialized.
func NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams() *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams {
	var ()
	return &ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParamsWithTimeout creates a new ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParamsWithTimeout(timeout time.Duration) *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams {
	var ()
	return &ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams{

		timeout: timeout,
	}
}

// NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParamsWithContext creates a new ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams object
// with the default values initialized, and the ability to set a context for a request
func NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParamsWithContext(ctx context.Context) *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams {
	var ()
	return &ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams{

		Context: ctx,
	}
}

// NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParamsWithHTTPClient creates a new ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParamsWithHTTPClient(client *http.Client) *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams {
	var ()
	return &ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams{
		HTTPClient: client,
	}
}

/*ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams contains all the parameters to send to the API endpoint
for the controllers v1 service controller remove service applied to vm insts operation typically these are written to a http.Request
*/
type ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams struct {

	/*Authorization
	  Your IBM Cloud Identity and Access Management (IAM) token. To retrieve your IAM token, run `ibmcloud iam oauth-tokens`.

	*/
	Authorization string
	/*InstanceID
	  Instance ID.

	*/
	InstanceID string
	/*ServiceInstanceID
	  Service Instance ID.

	*/
	ServiceInstanceID string
	/*XGlobalTransactionID
	  Global transaction ID for request correlation.

	*/
	XGlobalTransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) WithTimeout(timeout time.Duration) *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) WithContext(ctx context.Context) *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) WithHTTPClient(client *http.Client) *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) WithAuthorization(authorization string) *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithInstanceID adds the instanceID to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) WithInstanceID(instanceID string) *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithServiceInstanceID adds the serviceInstanceID to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) WithServiceInstanceID(serviceInstanceID string) *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams {
	o.SetServiceInstanceID(serviceInstanceID)
	return o
}

// SetServiceInstanceID adds the serviceInstanceId to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) SetServiceInstanceID(serviceInstanceID string) {
	o.ServiceInstanceID = serviceInstanceID
}

// WithXGlobalTransactionID adds the xGlobalTransactionID to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) WithXGlobalTransactionID(xGlobalTransactionID *string) *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams {
	o.SetXGlobalTransactionID(xGlobalTransactionID)
	return o
}

// SetXGlobalTransactionID adds the xGlobalTransactionId to the controllers v1 service controller remove service applied to vm insts params
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) SetXGlobalTransactionID(xGlobalTransactionID *string) {
	o.XGlobalTransactionID = xGlobalTransactionID
}

// WriteToRequest writes these params to a swagger request
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.InstanceID); err != nil {
		return err
	}

	// path param service_instance_id
	if err := r.SetPathParam("service_instance_id", o.ServiceInstanceID); err != nil {
		return err
	}

	if o.XGlobalTransactionID != nil {

		// header param x-global-transaction-id
		if err := r.SetHeaderParam("x-global-transaction-id", *o.XGlobalTransactionID); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
