// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/powerman/go-service-example/api/openapi/model"
	"github.com/powerman/go-service-example/internal/app"
)

// ListContactsHandlerFunc turns a function with the right signature into a list contacts handler
type ListContactsHandlerFunc func(ListContactsParams, *app.Auth) ListContactsResponder

// Handle executing the request and returning a response
func (fn ListContactsHandlerFunc) Handle(params ListContactsParams, principal *app.Auth) ListContactsResponder {
	return fn(params, principal)
}

// ListContactsHandler interface for that can handle valid list contacts params
type ListContactsHandler interface {
	Handle(ListContactsParams, *app.Auth) ListContactsResponder
}

// NewListContacts creates a new http.Handler for the list contacts operation
func NewListContacts(ctx *middleware.Context, handler ListContactsHandler) *ListContacts {
	return &ListContacts{Context: ctx, Handler: handler}
}

/*ListContacts swagger:route GET /contacts listContacts

Return all contacts ordered by ID ASC using pagination.

*/
type ListContacts struct {
	Context *middleware.Context
	Handler ListContactsHandler
}

func (o *ListContacts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListContactsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *app.Auth
	if uprinc != nil {
		principal = uprinc.(*app.Auth) // this is really a app.Auth, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ListContactsBody list contacts body
//
// swagger:model ListContactsBody
type ListContactsBody struct {
	model.SeekPagination
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ListContactsBody) UnmarshalJSON(raw []byte) error {
	// ListContactsParamsBodyAO0
	var listContactsParamsBodyAO0 model.SeekPagination
	if err := swag.ReadJSON(raw, &listContactsParamsBodyAO0); err != nil {
		return err
	}
	o.SeekPagination = listContactsParamsBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ListContactsBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	listContactsParamsBodyAO0, err := swag.WriteJSON(o.SeekPagination)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, listContactsParamsBodyAO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this list contacts body
func (o *ListContactsBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with model.SeekPagination
	if err := o.SeekPagination.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *ListContactsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListContactsBody) UnmarshalBinary(b []byte) error {
	var res ListContactsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
