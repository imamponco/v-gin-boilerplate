package vhttp

import (
	"github.com/gin-gonic/gin"
)

const (
	SuccessCode     = "OK"
	SuccessMessage  = "Success"
	SuccessCodeHTTP = 200
)

const (
	BadRequestCode     = "400"
	BadRequestCodeHTTP = 400
	BadRequestMessage  = "Bad Request"
)

const (
	UnauthorizedCode     = "401"
	UnauthorizedCodeHTTP = 401
	UnauthorizedMessage  = "Unauthorized"
)

const (
	ForbiddenCode     = "403"
	ForbiddenCodeHTTP = 403
	ForbiddenMessage  = "Forbidden"
)

const (
	InternalErrorCode     = "500"
	InternalErrorCodeHTTP = 500
	InternalErrorMessage  = "Internal Error"
)

// Data response bad request
type DataResponseBadRequest struct {
	Debug map[string]interface{} `json:"__debug"`
}

// Response use standard for response
type Response struct {
	Success bool              `json:"success" example:"true"`
	Code    string            `json:"code" example:"OK"`
	Message string            `json:"message" example:"Success"`
	Data    interface{}       `json:"data" swaggerignore:"true"`
	Header  map[string]string `json:"-"`
}

func (h *ResponseHandler) Success(c *gin.Context, data interface{}) *gin.Context {

	// Set response based on data
	resp := &Response{
		Success: true,
		Code:    SuccessCode,
		Message: SuccessMessage,
		Header:  make(map[string]string),
		Data:    data,
	}

	// Set response JSON
	c.JSON(SuccessCodeHTTP, resp)

	return c
}

func (h *ResponseHandler) BadRequest(c *gin.Context, Err *Error) *gin.Context {

	// Set message based on params
	resp := &Response{
		Success: false,
		Code:    Err.code,
		Message: Err.message,
		Header:  make(map[string]string),
		Data:    nil,
	}

	// Return error if debug is on
	if h.debug {
		resp.Data = &DataResponseBadRequest{
			Debug: map[string]interface{}{
				"metadata": Err.Error(),
			},
		}
	}

	// Set response JSON
	c.JSON(BadRequestCodeHTTP, resp)

	return c
}

func (h *ResponseHandler) Unauthorized(c *gin.Context, Err *Error) *gin.Context {

	// Set message based on params
	resp := &Response{
		Success: false,
		Code:    Err.code,
		Message: Err.message,
		Header:  make(map[string]string),
		Data:    nil,
	}

	// Return error if debug is on
	if h.debug {
		resp.Data = &DataResponseBadRequest{
			Debug: map[string]interface{}{
				"metadata": Err.Error(),
			},
		}
	}

	// Set response JSON
	c.JSON(UnauthorizedCodeHTTP, resp)

	return c
}

func (h *ResponseHandler) CustomizeError(c *gin.Context, Err *Error) *gin.Context {

	// Set message based on params
	resp := &Response{
		Success: Err.success,
		Code:    Err.code,
		Message: Err.message,
		Header:  make(map[string]string),
		Data:    nil,
	}

	// Return error if debug is on
	if h.debug {
		resp.Data = &DataResponseBadRequest{
			Debug: map[string]interface{}{
				"metadata": Err.Error(),
			},
		}
	}

	// Set response JSON
	c.JSON(Err.httpStatus, resp)

	return c
}

func (h *ResponseHandler) Forbidden(c *gin.Context, Err *Error) *gin.Context {

	// Set message based on params
	resp := &Response{
		Success: false,
		Code:    Err.code,
		Message: Err.message,
		Header:  make(map[string]string),
		Data:    nil,
	}

	// Return error if debug is on
	if h.debug {
		resp.Data = &DataResponseBadRequest{
			Debug: map[string]interface{}{
				"metadata": Err.Error(),
			},
		}
	}

	// Set response JSON
	c.JSON(ForbiddenCodeHTTP, resp)

	return c
}

func (h *ResponseHandler) OK(c *gin.Context) *gin.Context {
	// Set response OK
	resp := &Response{
		Success: true,
		Code:    SuccessCode,
		Message: SuccessMessage,
		Header:  make(map[string]string),
		Data:    nil,
	}

	// Set response JSON
	c.JSON(SuccessCodeHTTP, resp)

	return c
}

func (h *ResponseHandler) InternalErrorResponse(c *gin.Context) *gin.Context {
	// Set message based on params
	resp := &Response{
		Success: false,
		Code:    InternalErrorCode,
		Message: InternalErrorMessage,
		Header:  make(map[string]string),
		Data:    nil,
	}

	// Set response JSON
	c.JSON(InternalErrorCodeHTTP, resp)

	return c
}

type ResponseSuccess struct {
	Success bool              `json:"success" example:"true"`
	Code    string            `json:"code" example:"OK"`
	Message string            `json:"message" example:"Success"`
	Data    interface{}       `json:"data"`
	Header  map[string]string `json:"-"`
}

type ResponseBadRequest struct {
	Success bool              `json:"success" example:"false"`
	Code    string            `json:"code" example:"400"`
	Message string            `json:"message" example:"Bad Request"`
	Data    interface{}       `json:"data"`
	Header  map[string]string `json:"-"`
}

type ResponseInternalError struct {
	Success bool              `json:"success" example:"false"`
	Code    string            `json:"code" example:"500"`
	Message string            `json:"message" example:"Internal Error"`
	Data    interface{}       `json:"data"`
	Header  map[string]string `json:"-"`
}
