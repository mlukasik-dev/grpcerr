package grpcerr

import (
	"bytes"
	"fmt"
	"text/template"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Options struct {
	Template bool `yaml:"template"`
}

type GrpcErr struct {
	Name        string `yaml:"name"`
	Title       string `yaml:"title"`
	Message     string `yaml:"message"`
	MessageTmpl *template.Template
	Options     `yaml:"options"`
}

func (e *GrpcErr) Parse(data interface{}) *GrpcErr {
	if e.Options.Template {
		var tpl bytes.Buffer
		err := e.MessageTmpl.Execute(&tpl, data)
		if err != nil {
		}
		e.Message = tpl.String()
	}
	return nil
}

func (e *GrpcErr) New(code codes.Code, debug string) error {
	st, err := status.New(code, debug).WithDetails(&errdetails.ErrorInfo{Reason: e.Name, Metadata: map[string]string{
		"title":   e.Title,
		"message": e.Message,
	}})
	if err != nil {
	}
	return st.Err()
}

func (e *GrpcErr) Newf(code codes.Code, debugFmt string, a ...interface{}) error {
	return e.New(code, fmt.Sprintf(debugFmt, a...))
}
