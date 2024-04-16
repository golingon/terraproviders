// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ecrpublic_repository

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CatalogData struct {
	// AboutText: string, optional
	AboutText terra.StringValue `hcl:"about_text,attr"`
	// Architectures: set of string, optional
	Architectures terra.SetValue[terra.StringValue] `hcl:"architectures,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// LogoImageBlob: string, optional
	LogoImageBlob terra.StringValue `hcl:"logo_image_blob,attr"`
	// OperatingSystems: set of string, optional
	OperatingSystems terra.SetValue[terra.StringValue] `hcl:"operating_systems,attr"`
	// UsageText: string, optional
	UsageText terra.StringValue `hcl:"usage_text,attr"`
}

type Timeouts struct {
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type CatalogDataAttributes struct {
	ref terra.Reference
}

func (cd CatalogDataAttributes) InternalRef() (terra.Reference, error) {
	return cd.ref, nil
}

func (cd CatalogDataAttributes) InternalWithRef(ref terra.Reference) CatalogDataAttributes {
	return CatalogDataAttributes{ref: ref}
}

func (cd CatalogDataAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cd.ref.InternalTokens()
}

func (cd CatalogDataAttributes) AboutText() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("about_text"))
}

func (cd CatalogDataAttributes) Architectures() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cd.ref.Append("architectures"))
}

func (cd CatalogDataAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("description"))
}

func (cd CatalogDataAttributes) LogoImageBlob() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("logo_image_blob"))
}

func (cd CatalogDataAttributes) OperatingSystems() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cd.ref.Append("operating_systems"))
}

func (cd CatalogDataAttributes) UsageText() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("usage_text"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

type CatalogDataState struct {
	AboutText        string   `json:"about_text"`
	Architectures    []string `json:"architectures"`
	Description      string   `json:"description"`
	LogoImageBlob    string   `json:"logo_image_blob"`
	OperatingSystems []string `json:"operating_systems"`
	UsageText        string   `json:"usage_text"`
}

type TimeoutsState struct {
	Delete string `json:"delete"`
}
