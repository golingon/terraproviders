// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package apigatewaydocumentationpart

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Location struct {
	// Method: string, optional
	Method terra.StringValue `hcl:"method,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// StatusCode: string, optional
	StatusCode terra.StringValue `hcl:"status_code,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type LocationAttributes struct {
	ref terra.Reference
}

func (l LocationAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LocationAttributes) InternalWithRef(ref terra.Reference) LocationAttributes {
	return LocationAttributes{ref: ref}
}

func (l LocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l LocationAttributes) Method() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("method"))
}

func (l LocationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("name"))
}

func (l LocationAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("path"))
}

func (l LocationAttributes) StatusCode() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("status_code"))
}

func (l LocationAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("type"))
}

type LocationState struct {
	Method     string `json:"method"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	StatusCode string `json:"status_code"`
	Type       string `json:"type"`
}