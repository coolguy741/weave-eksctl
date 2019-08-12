package fileprocessor

import (
	"bytes"
	"github.com/pkg/errors"
	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	templateExtension = ".tmpl"
)

// TemplateParameters represents the API variables that can be used to template a profile. This set of variables will
// be applied to the go template files found. Templates filenames must end in .templ
type TemplateParameters struct {
	ClusterName string
}

// NewTemplateParameters creates a set of variables for templating given a ClusterConfig object
func NewTemplateParameters(clusterConfig *api.ClusterConfig) TemplateParameters {
	return TemplateParameters{
		ClusterName: clusterConfig.Metadata.Name,
	}
}

// GoTemplateProcessor is a FileProcessor that executes Go Templates
type GoTemplateProcessor struct {
	Params TemplateParameters
}

// ProcessFile takes a template file and executes the template applying the TemplateParameters
func (p *GoTemplateProcessor) ProcessFile(file File, baseDir string) (*File, error) {
	parsedTemplate, err := template.New(file.Name).Parse(string(file.Data))
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse manifest template file %q", file.Name)
	}

	out := bytes.NewBuffer(nil)
	if err = parsedTemplate.Execute(out, p.Params); err != nil {
		return nil, errors.Wrapf(err, "cannot execute template for file %q", file.Name)
	}

	relPath, err := filepath.Rel(baseDir, file.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get relative path for file %q", file.Name)
	}
	newFileName := strings.TrimSuffix(relPath, templateExtension)
	return &File{
		Data: out.Bytes(),
		Name: newFileName,
	}, nil
}
