package compat_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/mongodb/mongodb-atlas-kubernetes/pkg/util/compat"
)

func TestJSONSliceMerge(t *testing.T) {
	assert := assert.New(t)

	type Item struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}

	dst := []Item{
		{"00001", "dst1"},
		{"00002", "dst2"},
		{"00003", "dst3"},
	}

	src := []Item{
		{"99999", "src1"},  // different key, different value
		{"", "src2"},       // no key, different value
		{"", ""},           // no key, no value
		{"12345", "extra"}, // extra value
	}

	expected := []Item{
		{"99999", "src1"},
		{"00002", "src2"},
		{"00003", "dst3"},
		{"12345", "extra"},
	}

	err := JSONSliceMerge(&dst, src)
	assert.NoError(err)
	assert.Equal(expected, dst)
}