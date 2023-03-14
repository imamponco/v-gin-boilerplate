package vhttp

import "net/http"

// Common error
var ResourceNotFoundError = NewError("E_COMM_1", "Resource Not Found", http.StatusOK)
var InvalidResourceVersionError = NewError("E_COMM_2", "Invalid Resource Version", http.StatusBadRequest)
var ListQueryParamError = NewError("E_COMM_2", "Query param cannot be null", http.StatusBadRequest)
var UnauthorizedError = NewError("401", "Unauthorized", http.StatusUnauthorized)
var ForbiddenError = NewError("403", "Forbidden", http.StatusForbidden)
var TooManyRequestError = NewError("429", "Too Many Request", http.StatusTooManyRequests)

// Base Errors
var BadRequestError = NewError("400", "Bad Request", http.StatusBadRequest)
var InternalError = NewError("500", "Internal Error", http.StatusInternalServerError)

// Authorization Errors
var EmptyAuthorizationError = NewError("E_AUTH_1", "Authorization value is empty",
	http.StatusBadRequest)
var MalformedTokenError = NewError("E_AUTH_2", "Malformed token",
	http.StatusBadRequest)
var InvalidJWTFormatError = NewError("E_AUTH_3", "Invalid JWT Format",
	http.StatusBadRequest)
var InvalidJWTIssuerError = NewError("E_AUTH_4", "Invalid JWT Issuer",
	http.StatusUnauthorized)
var ExpiredJWTError = NewError("E_AUTH_5", "Expired JWT",
	http.StatusUnauthorized)
