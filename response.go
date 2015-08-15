// Various responses to a LWM2M Server request
// Typically responses represent a CoAP Request code
package betwixt

import (
	"github.com/zubairhamed/canopus"
)

// Created creates a LWM2M Response (CreatedResponse) with CoAP code 201 - Created
func Created() Lwm2mResponse {
	return &CreatedResponse{}
}

type CreatedResponse struct {
}

func (r *CreatedResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_201_CREATED
}

func (r *CreatedResponse) GetResponseValue() Value {
	return Empty()
}

// Deleted creates a LWM2M Response (DeletedResponse) with CoAP code 202 - Deleted
func Deleted() Lwm2mResponse {
	return &DeletedResponse{}
}

type DeletedResponse struct {
}

func (r *DeletedResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_202_DELETED
}

func (r *DeletedResponse) GetResponseValue() Value {
	return Empty()
}

// Changed creates a LWM2M Response (ChangedResponse) with CoAP code 204 - Changed
func Changed() Lwm2mResponse {
	return &ChangedResponse{}
}

type ChangedResponse struct {
}

func (r *ChangedResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_204_CHANGED
}

func (r *ChangedResponse) GetResponseValue() Value {
	return Empty()
}

// Content creates a LWM2M Response (ContentResponse) with CoAP code 205 - Content
func Content(val Value) Lwm2mResponse {
	return &ContentResponse{
		val: val,
	}
}

type ContentResponse struct {
	val Value
}

func (r *ContentResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_205_CONTENT
}

func (r *ContentResponse) GetResponseValue() Value {
	return r.val
}

// BadRequest creates a LWM2M Response (BadRequestResponse) with CoAP code 400 - Bad Request
func BadRequest() Lwm2mResponse {
	return &BadRequestResponse{}
}

type BadRequestResponse struct {
}

func (r *BadRequestResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_400_BAD_REQUEST
}

func (r *BadRequestResponse) GetResponseValue() Value {
	return Empty()
}

// Unauthorized creates a LWM2M Response (UnauthorizedResponse) with CoAP code 401 - Unauthorized
func Unauthorized() Lwm2mResponse {
	return &UnauthorizedResponse{}
}

type UnauthorizedResponse struct {
}

func (r *UnauthorizedResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_401_UNAUTHORIZED
}

func (r *UnauthorizedResponse) GetResponseValue() Value {
	return Empty()
}

// NotFound creates a LWM2M Response (NotFoundResponse) with CoAP code 404 - Not Found
func NotFound() Lwm2mResponse {
	return &NotFoundResponse{}
}

type NotFoundResponse struct {
}

func (r *NotFoundResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_404_NOT_FOUND
}

func (r *NotFoundResponse) GetResponseValue() Value {
	return Empty()
}

/// MethodNotAllowed creates a LWM2M Response (MethodNotAllowedResponse) with CoAP code 405 - Method Not Allowed
func MethodNotAllowed() Lwm2mResponse {
	return &MethodNotAllowedResponse{}
}

type MethodNotAllowedResponse struct {
}

func (r *MethodNotAllowedResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (r *MethodNotAllowedResponse) GetResponseValue() Value {
	return Empty()
}

// Conflict creates a LWM2M Response (ConflictResponse) with CoAP code 409 - Conflict
func Conflict() Lwm2mResponse {
	return &ConflictResponse{}
}

type ConflictResponse struct {
}

func (r *ConflictResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_409_CONFLICT
}

func (r *ConflictResponse) GetResponseValue() Value {
	return Empty()
}
