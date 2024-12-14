package parser

import (
	"testing"
)

func shouldNotErr(t *testing.T, err error) {
	if err != nil {
		t.Fail()
		t.Logf("fail: expected no err, got %s", err.Error())
	}
}

func TestParseFiletypeYaml(t *testing.T) {
	path := "test_files/test_1.yaml"
	fileOpenAPI, err := Parse(path)
	shouldNotErr(t, err)
	if fileOpenAPI.Filetype != "yaml" {
		t.Fail()
		t.Logf("fail: expected yaml filetype, got %s", fileOpenAPI.Filetype)
	}
}

func TestParseFiletypeJson(t *testing.T) {
	path := "test_files/test_1.yaml"
	fileOpenAPI, err := Parse(path)
	shouldNotErr(t, err)
	if fileOpenAPI.Filetype != "yaml" {
		t.Fail()
		t.Logf("fail: expected yaml filetype, got %s", fileOpenAPI.Filetype)
	}
}

func TestParseYaml(t *testing.T) {
	path := "test_files/test_1.yaml"
	fileOpenAPI, err := Parse(path)
	shouldNotErr(t, err)
	if fileOpenAPI.Spec.Ver != "3.1.0" {
		t.Fail()
		t.Logf("expected: openapi ver 3.1.0, got: %s", fileOpenAPI.Spec.Ver)
	}
	if fileOpenAPI.Spec.Info.Version != "3.1.0" {
		t.Fail()
		t.Logf("expected: openapi ver 3.1.0, got: %s", fileOpenAPI.Spec.Ver)
	}
}

func TestParseJson(t *testing.T) {
	path := "test_files/test_1.json"
	fileOpenAPI, err := Parse(path)
	shouldNotErr(t, err)
	if fileOpenAPI.Spec.Ver != "3.1.0" {
		t.Fail()
		t.Logf("expected: openapi ver 3.1.0, got: %s", fileOpenAPI.Spec.Ver)
	}
}
