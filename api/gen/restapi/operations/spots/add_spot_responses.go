// Code generated by go-swagger; DO NOT EDIT.

package spots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/suujia/flow/api/gen/models"
)

// AddSpotCreatedCode is the HTTP code returned for type AddSpotCreated
const AddSpotCreatedCode int = 201

/*AddSpotCreated Created spot

swagger:response addSpotCreated
*/
type AddSpotCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Spots `json:"body,omitempty"`
}

// NewAddSpotCreated creates AddSpotCreated with default headers values
func NewAddSpotCreated() *AddSpotCreated {

	return &AddSpotCreated{}
}

// WithPayload adds the payload to the add spot created response
func (o *AddSpotCreated) WithPayload(payload *models.Spots) *AddSpotCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add spot created response
func (o *AddSpotCreated) SetPayload(payload *models.Spots) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSpotCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddSpotDefault error response

swagger:response addSpotDefault
*/
type AddSpotDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddSpotDefault creates AddSpotDefault with default headers values
func NewAddSpotDefault(code int) *AddSpotDefault {
	if code <= 0 {
		code = 500
	}

	return &AddSpotDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add spot default response
func (o *AddSpotDefault) WithStatusCode(code int) *AddSpotDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add spot default response
func (o *AddSpotDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add spot default response
func (o *AddSpotDefault) WithPayload(payload *models.Error) *AddSpotDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add spot default response
func (o *AddSpotDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSpotDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
