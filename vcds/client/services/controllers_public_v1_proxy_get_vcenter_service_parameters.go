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

// NewControllersPublicV1ProxyGetVcenterServiceParams creates a new ControllersPublicV1ProxyGetVcenterServiceParams object
// with the default values initialized.
func NewControllersPublicV1ProxyGetVcenterServiceParams() *ControllersPublicV1ProxyGetVcenterServiceParams {
	var ()
	return &ControllersPublicV1ProxyGetVcenterServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewControllersPublicV1ProxyGetVcenterServiceParamsWithTimeout creates a new ControllersPublicV1ProxyGetVcenterServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewControllersPublicV1ProxyGetVcenterServiceParamsWithTimeout(timeout time.Duration) *ControllersPublicV1ProxyGetVcenterServiceParams {
	var ()
	return &ControllersPublicV1ProxyGetVcenterServiceParams{

		timeout: timeout,
	}
}

// NewControllersPublicV1ProxyGetVcenterServiceParamsWithContext creates a new ControllersPublicV1ProxyGetVcenterServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewControllersPublicV1ProxyGetVcenterServiceParamsWithContext(ctx context.Context) *ControllersPublicV1ProxyGetVcenterServiceParams {
	var ()
	return &ControllersPublicV1ProxyGetVcenterServiceParams{

		Context: ctx,
	}
}

// NewControllersPublicV1ProxyGetVcenterServiceParamsWithHTTPClient creates a new ControllersPublicV1ProxyGetVcenterServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewControllersPublicV1ProxyGetVcenterServiceParamsWithHTTPClient(client *http.Client) *ControllersPublicV1ProxyGetVcenterServiceParams {
	var ()
	return &ControllersPublicV1ProxyGetVcenterServiceParams{
		HTTPClient: client,
	}
}

/*ControllersPublicV1ProxyGetVcenterServiceParams contains all the parameters to send to the API endpoint
for the controllers public v1 proxy get vcenter service operation typically these are written to a http.Request
*/
type ControllersPublicV1ProxyGetVcenterServiceParams struct {

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

// WithTimeout adds the timeout to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) WithTimeout(timeout time.Duration) *ControllersPublicV1ProxyGetVcenterServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) WithContext(ctx context.Context) *ControllersPublicV1ProxyGetVcenterServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) WithHTTPClient(client *http.Client) *ControllersPublicV1ProxyGetVcenterServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) WithAuthorization(authorization string) *ControllersPublicV1ProxyGetVcenterServiceParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithInstanceID adds the instanceID to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) WithInstanceID(instanceID string) *ControllersPublicV1ProxyGetVcenterServiceParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithServiceInstanceID adds the serviceInstanceID to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) WithServiceInstanceID(serviceInstanceID string) *ControllersPublicV1ProxyGetVcenterServiceParams {
	o.SetServiceInstanceID(serviceInstanceID)
	return o
}

// SetServiceInstanceID adds the serviceInstanceId to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) SetServiceInstanceID(serviceInstanceID string) {
	o.ServiceInstanceID = serviceInstanceID
}

// WithXGlobalTransactionID adds the xGlobalTransactionID to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) WithXGlobalTransactionID(xGlobalTransactionID *string) *ControllersPublicV1ProxyGetVcenterServiceParams {
	o.SetXGlobalTransactionID(xGlobalTransactionID)
	return o
}

// SetXGlobalTransactionID adds the xGlobalTransactionId to the controllers public v1 proxy get vcenter service params
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) SetXGlobalTransactionID(xGlobalTransactionID *string) {
	o.XGlobalTransactionID = xGlobalTransactionID
}

// WriteToRequest writes these params to a swagger request
func (o *ControllersPublicV1ProxyGetVcenterServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
