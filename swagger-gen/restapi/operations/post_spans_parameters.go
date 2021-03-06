// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/jaegertracing/jaeger/swagger-gen/models"
)

// NewPostSpansParams creates a new PostSpansParams object
// with the default values initialized.
func NewPostSpansParams() PostSpansParams {
	var ()
	return PostSpansParams{}
}

// PostSpansParams contains all the bound params for the post spans operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostSpans
type PostSpansParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*A list of spans that belong to any trace.
	  Required: true
	  In: body
	*/
	Spans models.ListOfSpans
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PostSpansParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ListOfSpans
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("spans", "body"))
			} else {
				res = append(res, errors.NewParseError("spans", "body", "", err))
			}

		} else {

			if len(res) == 0 {
				o.Spans = body
			}
		}

	} else {
		res = append(res, errors.Required("spans", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
