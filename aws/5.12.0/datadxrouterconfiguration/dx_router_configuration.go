// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datadxrouterconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Router struct{}

type RouterAttributes struct {
	ref terra.Reference
}

func (r RouterAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RouterAttributes) InternalWithRef(ref terra.Reference) RouterAttributes {
	return RouterAttributes{ref: ref}
}

func (r RouterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RouterAttributes) Platform() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("platform"))
}

func (r RouterAttributes) RouterTypeIdentifier() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("router_type_identifier"))
}

func (r RouterAttributes) Software() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("software"))
}

func (r RouterAttributes) Vendor() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("vendor"))
}

func (r RouterAttributes) XsltTemplateName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("xslt_template_name"))
}

func (r RouterAttributes) XsltTemplateNameForMacSec() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("xslt_template_name_for_mac_sec"))
}

type RouterState struct {
	Platform                  string `json:"platform"`
	RouterTypeIdentifier      string `json:"router_type_identifier"`
	Software                  string `json:"software"`
	Vendor                    string `json:"vendor"`
	XsltTemplateName          string `json:"xslt_template_name"`
	XsltTemplateNameForMacSec string `json:"xslt_template_name_for_mac_sec"`
}
