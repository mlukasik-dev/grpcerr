package grpcerr

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

// AppErr represents structured error response.
type AppErr struct {

	// Application specific code.
	Code string `json:"code"`

	// User readable title (summary) of the error.
	Title string `json:"title"`

	// User readable message.
	Message string `json:"message"`

	// Optional debug info.
	Debug string `json:"debug,omitempty"`
}

func Convert(e error) (appErr AppErr) {
	st := status.Convert(e)
	if st != nil && len(st.Details()) > 0 {
		d, ok := st.Details()[0].(*errdetails.ErrorInfo)
		if !ok || d == nil {
			appErr = AppErr{Code: "UNKNOWN_INTERNAL", Title: "", Message: ""}
			return appErr
		}
		appErr := AppErr{
			Code:    d.Reason,
			Title:   d.Metadata["title"],
			Message: d.Metadata["message"],
			Debug:   st.Message(),
		}
		return appErr
	}
	appErr = AppErr{Code: "UNKNOWN_INTERNAL", Title: "", Message: ""}
	return appErr
}
