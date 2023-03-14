package vhttp

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nbs-go/nlogger"
	"net/http"
)

type ResponseHandler struct {
	debug  bool
	logger nlogger.Logger
}
type ResponseHandlerOptions struct {
	Debug  bool
	Logger nlogger.Logger
}

func NewResponseHandler(args ...ResponseHandlerOptions) ResponseHandler {
	// Get options
	var options ResponseHandlerOptions
	if len(args) > 0 {
		options = args[0]
	} else {
		options = ResponseHandlerOptions{
			Debug: false,
		}
	}

	return ResponseHandler{
		debug:  options.Debug,
		logger: options.Logger,
	}
}

type Error struct {
	success    bool
	code       string
	message    string
	httpStatus int
	err        error
}

type options struct {
	setSuccess bool
}

type SetOptionFn = func(*options)

func SetSuccess(success bool) SetOptionFn {
	return func(o *options) {
		o.setSuccess = success
	}
}

// Error implement standard go error interface. If source error is exists then it will print error cause
func (e *Error) Error() string {
	errMsg := e.err.Error()
	if errMsg == "" {
		errMsg = e.message
	}
	return errMsg
}

func NewError(code string, message string, httpStatus int, args ...SetOptionFn) *Error {

	// Evaluate options
	o := evaluateOptions(args)

	err := &Error{
		success:    false,
		code:       code,
		message:    message,
		httpStatus: httpStatus,
		err:        errors.New(""),
	}

	// Set success
	if o.setSuccess {
		err.success = o.setSuccess
	}

	return err
}

func defaultOptions() *options {
	return &options{
		setSuccess: false,
	}
}

func evaluateOptions(args []SetOptionFn) *options {
	optCopy := defaultOptions()
	for _, fn := range args {
		fn(optCopy)
	}
	return optCopy
}

func (e *Error) WithStatus(httpStatus int) *Error {
	e.httpStatus = httpStatus
	return e
}

func (e *Error) Code(code string) *Error {
	e.code = code
	return e
}

func (e *Error) Wrap(err error) *Error {
	e.err = err
	return e
}

func (h *ResponseHandler) ErrorHandler(c *gin.Context, err error) *gin.Context {
	// Init result
	result := c

	// Assert error
	var hErr *Error
	ok := errors.As(err, &hErr)
	if ok {
		switch hErr.httpStatus {
		case http.StatusTooManyRequests:
			return h.CustomizeError(c, hErr)
		case http.StatusOK:
			return h.CustomizeError(c, hErr)
		case http.StatusBadRequest:
			return h.BadRequest(c, hErr)
		case http.StatusUnauthorized:
			return h.Unauthorized(c, hErr)
		case http.StatusForbidden:
			return h.Forbidden(c, hErr)
		default:
			h.logger.Errorf("%v", err)
			return h.InternalErrorResponse(c)
		}
	}

	h.logger.Errorf("%v", err)

	result = h.InternalErrorResponse(c)

	return result
}
