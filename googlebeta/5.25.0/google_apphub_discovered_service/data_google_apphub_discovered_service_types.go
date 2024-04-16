// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_apphub_discovered_service

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataServicePropertiesAttributes struct {
	ref terra.Reference
}

func (sp DataServicePropertiesAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp DataServicePropertiesAttributes) InternalWithRef(ref terra.Reference) DataServicePropertiesAttributes {
	return DataServicePropertiesAttributes{ref: ref}
}

func (sp DataServicePropertiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp DataServicePropertiesAttributes) GcpProject() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("gcp_project"))
}

func (sp DataServicePropertiesAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("location"))
}

func (sp DataServicePropertiesAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("zone"))
}

type DataServiceReferenceAttributes struct {
	ref terra.Reference
}

func (sr DataServiceReferenceAttributes) InternalRef() (terra.Reference, error) {
	return sr.ref, nil
}

func (sr DataServiceReferenceAttributes) InternalWithRef(ref terra.Reference) DataServiceReferenceAttributes {
	return DataServiceReferenceAttributes{ref: ref}
}

func (sr DataServiceReferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sr.ref.InternalTokens()
}

func (sr DataServiceReferenceAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("path"))
}

func (sr DataServiceReferenceAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("uri"))
}

type DataServicePropertiesState struct {
	GcpProject string `json:"gcp_project"`
	Location   string `json:"location"`
	Zone       string `json:"zone"`
}

type DataServiceReferenceState struct {
	Path string `json:"path"`
	Uri  string `json:"uri"`
}
