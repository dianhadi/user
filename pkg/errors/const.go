package errors

import "net/http"

const (
	msgUnclassified = "Unclassified Error"

	StatusOK StatusCode = 200
	msgOK               = "Success"

	StatusRequestBodyInvalid   StatusCode = 400
	msgRequestBodyInvalid                 = "Request Body is not valid"
	StatusAuthorizationInvalid StatusCode = 401
	msgAuthorizationInvalid               = "Authorization is not valid"
	StatusDataNotFound         StatusCode = 404
	msgDataNotFound                       = "Data not found"

	StatusInternalError      StatusCode = 500
	msgInternalError                    = "Internal Server Error"
	StatusServiceUnavailable StatusCode = 503
	msgServiceUnavailable               = "Service Unavailable"
)

var (
	errHttpStatus = map[StatusCode]int{
		StatusOK:                   http.StatusOK,
		StatusRequestBodyInvalid:   http.StatusBadRequest,
		StatusAuthorizationInvalid: http.StatusUnauthorized,
		StatusInternalError:        http.StatusInternalServerError,
		StatusServiceUnavailable:   http.StatusServiceUnavailable,
		StatusDataNotFound:         http.StatusNotFound,
	}
	errMsg = map[StatusCode]string{
		StatusOK:                   msgOK,
		StatusRequestBodyInvalid:   msgRequestBodyInvalid,
		StatusAuthorizationInvalid: msgAuthorizationInvalid,
		StatusInternalError:        msgInternalError,
		StatusDataNotFound:         msgDataNotFound,
	}
)

const rootPath = "github.com/dianhadi/user/"
