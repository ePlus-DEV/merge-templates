package handler

import (
	"bytes"
	"html/template"
	"os"
)

func MergeTemplates(outPath string, mainTemplatePath string, templatePaths ...string) error {
	mergedTemplateData, err := mergeTemplates(mainTemplatePath, templatePaths...)
	if err != nil {
		return err
	}
	return os.WriteFile(outPath, mergedTemplateData, 0644)
}

func mergeTemplates(mainTemplatePath string, templatePaths ...string) ([]byte, error) {
	tplPaths := []string{mainTemplatePath}
	tplPaths = append(tplPaths, templatePaths...)
	mergedTemplate, err := template.ParseFiles(tplPaths...)
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer
	err = mergedTemplate.Execute(&result, nil)
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
