package generic_error

import "net/http"

const (
	ErrorCodeUnknown                    string = "unknown_error"
	ErrorCodeInternalServerError        string = "internal_server_error"
	ErrorCodeFormat                     string = "invalid_format"
	ErrorCodeFieldValue                 string = "invalid_field_value"
	ErrorCodeValidation                 string = "validation_failed"
	ErrorCodeForbidden                  string = "forbidden"
	ErrorCodeNotFound                   string = "resource_not_found"
	ErrorCodeExternalServiceUnavailable string = "external_service_unavailable"
	ErrorCodeExternalServiceError       string = "external_service_error"
	ErrorCodeUnsupported                string = "operation_unsupported"
	ErrorCodeExpired                    string = "operation_expired"
)

var CommonErrorDescriptions = map[string]string{
	ErrorCodeUnknown:                    "Unknown error.",
	ErrorCodeInternalServerError:        "Internal server error.",
	ErrorCodeFormat:                     "Invalid format of request.",
	ErrorCodeFieldValue:                 "Invalid value of request field.",
	ErrorCodeValidation:                 "Validation failed.",
	ErrorCodeForbidden:                  "Access denied.",
	ErrorCodeNotFound:                   "Resource not found.",
	ErrorCodeExternalServiceUnavailable: "External service is temporarily unavailable.",
	ErrorCodeExternalServiceError:       "External service reported error.",
	ErrorCodeUnsupported:                "Operation unsupported.",
	ErrorCodeExpired:                    "Operation expired.",
}

var CommonErrorHttpCodes = map[string]int{
	ErrorCodeUnknown:                    http.StatusInternalServerError,
	ErrorCodeInternalServerError:        http.StatusInternalServerError,
	ErrorCodeForbidden:                  http.StatusForbidden,
	ErrorCodeNotFound:                   http.StatusNotFound,
	ErrorCodeExternalServiceUnavailable: http.StatusInternalServerError,
	ErrorCodeExternalServiceError:       http.StatusInternalServerError,
}
