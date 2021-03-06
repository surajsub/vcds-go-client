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

// NewControllersPublicV1ProxyGetServiceCatalogListParams creates a new ControllersPublicV1ProxyGetServiceCatalogListParams object
// with the default values initialized.
func NewControllersPublicV1ProxyGetServiceCatalogListParams() *ControllersPublicV1ProxyGetServiceCatalogListParams {
	var ()
	return &ControllersPublicV1ProxyGetServiceCatalogListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewControllersPublicV1ProxyGetServiceCatalogListParamsWithTimeout creates a new ControllersPublicV1ProxyGetServiceCatalogListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewControllersPublicV1ProxyGetServiceCatalogListParamsWithTimeout(timeout time.Duration) *ControllersPublicV1ProxyGetServiceCatalogListParams {
	var ()
	return &ControllersPublicV1ProxyGetServiceCatalogListParams{

		timeout: timeout,
	}
}

// NewControllersPublicV1ProxyGetServiceCatalogListParamsWithContext creates a new ControllersPublicV1ProxyGetServiceCatalogListParams object
// with the default values initialized, and the ability to set a context for a request
func NewControllersPublicV1ProxyGetServiceCatalogListParamsWithContext(ctx context.Context) *ControllersPublicV1ProxyGetServiceCatalogListParams {
	var ()
	return &ControllersPublicV1ProxyGetServiceCatalogListParams{

		Context: ctx,
	}
}

// NewControllersPublicV1ProxyGetServiceCatalogListParamsWithHTTPClient creates a new ControllersPublicV1ProxyGetServiceCatalogListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewControllersPublicV1ProxyGetServiceCatalogListParamsWithHTTPClient(client *http.Client) *ControllersPublicV1ProxyGetServiceCatalogListParams {
	var ()
	return &ControllersPublicV1ProxyGetServiceCatalogListParams{
		HTTPClient: client,
	}
}

/*ControllersPublicV1ProxyGetServiceCatalogListParams contains all the parameters to send to the API endpoint
for the controllers public v1 proxy get service catalog list operation typically these are written to a http.Request
*/
type ControllersPublicV1ProxyGetServiceCatalogListParams struct {

	/*XGlobalTransactionID
	  Global transaction ID for request correlation.

	*/
	XGlobalTransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the controllers public v1 proxy get service catalog list params
func (o *ControllersPublicV1ProxyGetServiceCatalogListParams) WithTimeout(timeout time.Duration) *ControllersPublicV1ProxyGetServiceCatalogListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the controllers public v1 proxy get service catalog list params
func (o *ControllersPublicV1ProxyGetServiceCatalogListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the controllers public v1 proxy get service catalog list params
func (o *ControllersPublicV1ProxyGetServiceCatalogListParams) WithContext(ctx context.Context) *ControllersPublicV1ProxyGetServiceCatalogListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the controllers public v1 proxy get service catalog list params
func (o *ControllersPublicV1ProxyGetServiceCatalogListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the controllers public v1 proxy get service catalog list params
func (o *ControllersPublicV1ProxyGetServiceCatalogListParams) WithHTTPClient(client *http.Client) *ControllersPublicV1ProxyGetServiceCatalogListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the controllers public v1 proxy get service catalog list params
func (o *ControllersPublicV1ProxyGetServiceCatalogListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXGlobalTransactionID adds the xGlobalTransactionID to the controllers public v1 proxy get service catalog list params
func (o *ControllersPublicV1ProxyGetServiceCatalogListParams) WithXGlobalTransactionID(xGlobalTransactionID *string) *ControllersPublicV1ProxyGetServiceCatalogListParams {
	o.SetXGlobalTransactionID(xGlobalTransactionID)
	return o
}

// SetXGlobalTransactionID adds the xGlobalTransactionId to the controllers public v1 proxy get service catalog list params
func (o *ControllersPublicV1ProxyGetServiceCatalogListParams) SetXGlobalTransactionID(xGlobalTransactionID *string) {
	o.XGlobalTransactionID = xGlobalTransactionID
}

// WriteToRequest writes these params to a swagger request
func (o *ControllersPublicV1ProxyGetServiceCatalogListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
