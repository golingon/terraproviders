// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dataapphubdiscoveredworkload

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type WorkloadProperties struct{}

type WorkloadReference struct{}

type WorkloadPropertiesAttributes struct {
	ref terra.Reference
}

func (wp WorkloadPropertiesAttributes) InternalRef() (terra.Reference, error) {
	return wp.ref, nil
}

func (wp WorkloadPropertiesAttributes) InternalWithRef(ref terra.Reference) WorkloadPropertiesAttributes {
	return WorkloadPropertiesAttributes{ref: ref}
}

func (wp WorkloadPropertiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wp.ref.InternalTokens()
}

func (wp WorkloadPropertiesAttributes) GcpProject() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("gcp_project"))
}

func (wp WorkloadPropertiesAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("location"))
}

func (wp WorkloadPropertiesAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("zone"))
}

type WorkloadReferenceAttributes struct {
	ref terra.Reference
}

func (wr WorkloadReferenceAttributes) InternalRef() (terra.Reference, error) {
	return wr.ref, nil
}

func (wr WorkloadReferenceAttributes) InternalWithRef(ref terra.Reference) WorkloadReferenceAttributes {
	return WorkloadReferenceAttributes{ref: ref}
}

func (wr WorkloadReferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wr.ref.InternalTokens()
}

func (wr WorkloadReferenceAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(wr.ref.Append("uri"))
}

type WorkloadPropertiesState struct {
	GcpProject string `json:"gcp_project"`
	Location   string `json:"location"`
	Zone       string `json:"zone"`
}

type WorkloadReferenceState struct {
	Uri string `json:"uri"`
}
