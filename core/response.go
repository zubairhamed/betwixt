package core

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	. "github.com/zubairhamed/goap"
)

func NewMethodNotAllowedResponse() Response {
	return &MethodNotAllowedResponse{}
}

type MethodNotAllowedResponse struct {
}

func (r *MethodNotAllowedResponse) GetResponseCode() CoapCode {
	return COAPCODE_405_METHOD_NOT_ALLOWED
}

func (r *MethodNotAllowedResponse) GetResponseValue() ResponseValue {
	return NewEmptyValue()
}

func NewConflictResponse() Response {
	return &ConflictResponse{}
}

type ConflictResponse struct {
}

func (r *ConflictResponse) GetResponseCode() CoapCode {
	return COAPCODE_409_CONFLICT
}

func (r *ConflictResponse) GetResponseValue() ResponseValue {
	return NewEmptyValue()
}

func NewBadRequestResponse() Response {
	return &BadRequestResponse{}
}

type BadRequestResponse struct {
}

func (r *BadRequestResponse) GetResponseCode() CoapCode {
	return COAPCODE_400_BAD_REQUEST
}

func (r *BadRequestResponse) GetResponseValue() ResponseValue {
	return NewEmptyValue()
}

func NewNotFoundResponse() Response {
	return &NotFoundResponse{}
}

type NotFoundResponse struct {
}

func (r *NotFoundResponse) GetResponseCode() CoapCode {
	return COAPCODE_404_NOT_FOUND
}

func (r *NotFoundResponse) GetResponseValue() ResponseValue {
	return NewEmptyValue()
}

func NewUnauthorizedResponse() Response {
	return &UnauthorizedResponse{}
}

type UnauthorizedResponse struct {
}

func (r *UnauthorizedResponse) GetResponseCode() CoapCode {
	return COAPCODE_401_UNAUTHORIZED
}

func (r *UnauthorizedResponse) GetResponseValue() ResponseValue {
	return NewEmptyValue()
}

func NewChangedResponse() Response {
	return &ChangedResponse{}
}

type ChangedResponse struct {
}

func (r *ChangedResponse) GetResponseCode() CoapCode {
	return COAPCODE_204_CHANGED
}

func (r *ChangedResponse) GetResponseValue() ResponseValue {
	return NewEmptyValue()
}

func NewCreatedResponse() Response {
	return &CreatedResponse{}
}

type CreatedResponse struct {
}

func (r *CreatedResponse) GetResponseCode() CoapCode {
	return COAPCODE_201_CREATED
}

func (r *CreatedResponse) GetResponseValue() ResponseValue {
	return NewEmptyValue()
}

func NewDeletedResponse() Response {
	return &DeletedResponse{}
}

type DeletedResponse struct {
}

func (r *DeletedResponse) GetResponseCode() CoapCode {
	return COAPCODE_202_DELETED
}

func (r *DeletedResponse) GetResponseValue() ResponseValue {
	return NewEmptyValue()
}

func NewContentResponse(val ResponseValue) Response {
	return &ContentResponse{
		val: val,
	}
}

type ContentResponse struct {
	val ResponseValue
}

func (r *ContentResponse) GetResponseCode() CoapCode {
	return COAPCODE_205_CONTENT
}

func (r *ContentResponse) GetResponseValue() ResponseValue {
	return r.val
}
