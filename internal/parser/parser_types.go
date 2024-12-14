package parser

import "errors"

var ErrWrongFileFormat = errors.New("wrong file format")

/*
An OAS document consists of several sections, including:

Info: General information about the API, such as title, version, description, and contact details.
Paths: Definitions of API endpoints (URL paths) and the operations (HTTP methods) supported by each endpoint.
Components: Reusable components such as schemas (data models), parameters, responses, request bodies, and security schemes.
Security: Specification of security requirements and mechanisms, including authentication methods and authorization scopes.
Tags: Grouping of API operations into categories or tags for organizational purposes.
External Documentation: Links to external documentation or resources related to the API.
Servers: Base URLs and server-specific information for API endpoints.
Examples: Examples of API requests and responses to illustrate usage.
*/

type Info struct {
	Version string `json:"version" yaml:"version"`
}

type Path struct {
	URL    string
	Method string
}

type Component struct{}

type Tag struct {
	Name string
}

type ExternalDoc struct {
	Link string
}

type Server struct {
	BaseURL string
}

type Example struct{}

// version: 1.0.0
// title: Simple Petstore
// description: This is a slimmed down single path version of the Petstore definition.

type OpenAPI struct {
	Ver  string `json:"openapi" yaml:"openapi"`
	Info Info   `json:"info" yaml:"info"`
}

type FileOpenAPI struct {
	Filetype string
	Spec     OpenAPI
}
