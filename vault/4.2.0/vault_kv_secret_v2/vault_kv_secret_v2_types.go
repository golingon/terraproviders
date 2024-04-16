// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_kv_secret_v2

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CustomMetadata struct {
	// CasRequired: bool, optional
	CasRequired terra.BoolValue `hcl:"cas_required,attr"`
	// Data: map of string, optional
	Data terra.MapValue[terra.StringValue] `hcl:"data,attr"`
	// DeleteVersionAfter: number, optional
	DeleteVersionAfter terra.NumberValue `hcl:"delete_version_after,attr"`
	// MaxVersions: number, optional
	MaxVersions terra.NumberValue `hcl:"max_versions,attr"`
}

type CustomMetadataAttributes struct {
	ref terra.Reference
}

func (cm CustomMetadataAttributes) InternalRef() (terra.Reference, error) {
	return cm.ref, nil
}

func (cm CustomMetadataAttributes) InternalWithRef(ref terra.Reference) CustomMetadataAttributes {
	return CustomMetadataAttributes{ref: ref}
}

func (cm CustomMetadataAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cm.ref.InternalTokens()
}

func (cm CustomMetadataAttributes) CasRequired() terra.BoolValue {
	return terra.ReferenceAsBool(cm.ref.Append("cas_required"))
}

func (cm CustomMetadataAttributes) Data() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cm.ref.Append("data"))
}

func (cm CustomMetadataAttributes) DeleteVersionAfter() terra.NumberValue {
	return terra.ReferenceAsNumber(cm.ref.Append("delete_version_after"))
}

func (cm CustomMetadataAttributes) MaxVersions() terra.NumberValue {
	return terra.ReferenceAsNumber(cm.ref.Append("max_versions"))
}

type CustomMetadataState struct {
	CasRequired        bool              `json:"cas_required"`
	Data               map[string]string `json:"data"`
	DeleteVersionAfter float64           `json:"delete_version_after"`
	MaxVersions        float64           `json:"max_versions"`
}
