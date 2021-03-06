package codegen

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mlukasik-dev/grpcerr"
	"gopkg.in/yaml.v2"
)

// Config allows to tweak the behavior of Generate function.
type Config struct {
	GoFilename, JsFilename string
}

// Generate generates Go and JavaScript code from given config file
// into outdir directory.
func Generate(filename, outdir string, config ...Config) error {
	// read config file into memory.
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}
	// Unmarshal read file content into struct.
	var c grpcErrors
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config data: %v", err)
	}

	// Merge definitions with uses.
	for i, rpc := range c.Uses {
		for j, use := range rpc {
			def := findDefinition(use.Name, c.Definitions)
			if def == nil {
				return fmt.Errorf("definition for %s was not found", use.Name)
			}
			if use.Title == nil {
				use.Title = &def.Title
			}
			if use.Message == nil {
				use.Message = &def.Message
			}
			if use.Options.Template == nil {
				use.Options.Template = &def.Options.Template
			}
			rpc[j] = use
		}
		c.Uses[i] = rpc
	}

	// Create output directory (if not exists).
	if err = os.MkdirAll(outdir, 0700); err != nil {
		return err
	}

	// Generating Go file based on configfile and template.
	var goFilename = filepath.Join(outdir, "gen.go")
	if len(config) > 0 && config[0].GoFilename != "" {
		goFilename = filepath.Join(outdir, config[0].GoFilename)
	}
	f, err := os.Create(goFilename)
	if err != nil {
		return fmt.Errorf("failed to create %s file: %v", goFilename, err)
	}
	err = GoTmpl.Execute(f, c)
	if err != nil {
		return fmt.Errorf("failed to generate Go code: %v", err)
	}
	return nil
}

func findDefinition(name string, defs []grpcerr.GrpcErr) *grpcerr.GrpcErr {
	for _, def := range defs {
		if def.Name == name {
			return &def
		}
	}
	return nil
}

type grpcErrors struct {
	Definitions []grpcerr.GrpcErr `yaml:"definitions"`
	Uses        map[string][]use  `yaml:"uses"`
}

type optionsUse struct {
	Template *bool `yaml:"template"`
}

type use struct {
	Name    string     `yaml:"name"`
	Title   *string    `yaml:"title"`
	Message *string    `yaml:"message"`
	Options optionsUse `yaml:"options"`
}
