// Code generated by apperr. DO NOT EDIT.

package gen

import (
	"text/template"

	"github.com/mlukasik-dev/grpcerr"
)

var GetUser = struct {
	UnknownInternal  grpcerr.GrpcErr
	ResourceNotFound grpcerr.GrpcErr
}{
	grpcerr.GrpcErr{
		Name:        "UNKNOWN_INTERNAL",
		Title:       "Ups. Sorry for the inconvenience",
		Message:     "Some error happened to our services, we're trying to fix it. (reload page and so on...)",
		MessageTmpl: nil,
		Options: grpcerr.Options{
			Template: false,
		}},
	grpcerr.GrpcErr{
		Name:        "RESOURCE_NOT_FOUND",
		Title:       "Not found",
		Message:     "",
		MessageTmpl: template.Must(template.New("").Parse("User with {{ .Email }} was not found")),
		Options: grpcerr.Options{
			Template: true,
		}},
}
