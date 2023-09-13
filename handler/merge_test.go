package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTemplates(t *testing.T) {
	mergedData, err := mergeTemplates("testdata/tmpl.tpl", "testdata/tmpl_1.tpl", "testdata/tmpl_2.tpl")
	assert.NoError(t, err)
	assert.NotEmpty(t, mergedData)
	assert.Equal(t, "Template 01\nTemplate 02", string(mergedData))
}
