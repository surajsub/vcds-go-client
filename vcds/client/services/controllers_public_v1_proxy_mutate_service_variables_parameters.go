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

// NewControllersPublicV1ProxyMutateServiceVariablesParams creates a new ControllersPublicV1ProxyMutateServiceVariablesParams object
// with the default values initialized.
func NewControllersPublicV1ProxyMutateServiceVariablesParams() *ControllersPublicV1ProxyMutateServiceVariablesParams {
	var ()
	return &ControllersPublicV1ProxyMutateServiceVariablesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewControllersPublicV1ProxyMutateServiceVariablesParamsWithTimeout creates a new ControllersPublicV1ProxyMutateServiceVariablesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewControllersPublicV1ProxyMutateServiceVariablesParamsWithTimeout(timeout time.Duration) *ControllersPublicV1ProxyMutateServiceVariablesParams {
	var ()
	return &ControllersPublicV1ProxyMutateServiceVariablesParams{

		timeout: timeout,
	}
}

// NewControllersPublicV1ProxyMutateServiceVariablesParamsWithContext creates a new ControllersPublicV1ProxyMutateServiceVariablesParams object
// with the default values initialized, and the ability to set a context for a request
func NewControllersPublicV1ProxyMutateServiceVariablesParamsWithContext(ctx context.Context) *ControllersPublicV1ProxyMutateServiceVariablesParams {
	var ()
	return &ControllersPublicV1ProxyMutateServiceVariablesParams{

		Context: ctx,
	}
}

// NewControllersPublicV1ProxyMutateServiceVariablesParamsWithHTTPClient creates a new ControllersPublicV1ProxyMutateServiceVariablesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewControllersPublicV1ProxyMutateServiceVariablesParamsWithHTTPClient(client *http.Client) *ControllersPublicV1ProxyMutateServiceVariablesParams {
	var ()
	return &ControllersPublicV1ProxyMutateServiceVariablesParams{
		HTTPClient: client,
	}
}

/*ControllersPublicV1ProxyMutateServiceVariablesParams contains all the parameters to send to the API endpoint
for the controllers public v1 proxy mutate service variables operation typically these are written to a http.Request
*/
type ControllersPublicV1ProxyMutateServiceVariablesParams struct {

	/*Authorization
	  Your IBM Cloud Identity and Access Management (IAM) token. To retrieve your IAM token, run `ibmcloud iam oauth-tokens`.

	*/
	Authorization string
	/*ServiceInstanceID
	  Service Instance ID.

	*/
	ServiceInstanceID string
	/*ServiceVariablePatchData
	  The list of variables to be updated or deleted for a service instance.

	*/
	ServiceVariablePatchData ControllersPublicV1ProxyMutateServiceVariablesBody
	/*XGlobalTransactionID
	  Global transaction ID for request correlation.

	*/
	XGlobalTransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) WithTimeout(timeout time.Duration) *ControllersPublicV1ProxyMutateServiceVariablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) WithContext(ctx context.Context) *ControllersPublicV1ProxyMutateServiceVariablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) WithHTTPClient(client *http.Client) *ControllersPublicV1ProxyMutateServiceVariablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) WithAuthorization(authorization string) *ControllersPublicV1ProxyMutateServiceVariablesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithServiceInstanceID adds the serviceInstanceID to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) WithServiceInstanceID(serviceInstanceID string) *ControllersPublicV1ProxyMutateServiceVariablesParams {
	o.SetServiceInstanceID(serviceInstanceID)
	return o
}

// SetServiceInstanceID adds the serviceInstanceId to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) SetServiceInstanceID(serviceInstanceID string) {
	o.ServiceInstanceID = serviceInstanceID
}

// WithServiceVariablePatchData adds the serviceVariablePatchData to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) WithServiceVariablePatchData(serviceVariablePatchData ControllersPublicV1ProxyMutateServiceVariablesBody) *ControllersPublicV1ProxyMutateServiceVariablesParams {
	o.SetServiceVariablePatchData(serviceVariablePatchData)
	return o
}

// SetServiceVariablePatchData adds the serviceVariablePatchData to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) SetServiceVariablePatchData(serviceVariablePatchData ControllersPublicV1ProxyMutateServiceVariablesBody) {
	o.ServiceVariablePatchData = serviceVariablePatchData
}

// WithXGlobalTransactionID adds the xGlobalTransactionID to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) WithXGlobalTransactionID(xGlobalTransactionID *string) *ControllersPublicV1ProxyMutateServiceVariablesParams {
	o.SetXGlobalTransactionID(xGlobalTransactionID)
	return o
}

// SetXGlobalTransactionID adds the xGlobalTransactionId to the controllers public v1 proxy mutate service variables params
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) SetXGlobalTransactionID(xGlobalTransactionID *string) {
	o.XGlobalTransactionID = xGlobalTransactionID
}

// WriteToRequest writes these params to a swagger request
func (o *ControllersPublicV1ProxyMutateServiceVariablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param service_instance_id
	if err := r.SetPathParam("service_instance_id", o.ServiceInstanceID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.ServiceVariablePatchData); err != nil {
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