// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dataappserviceenvironment

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ClusterSetting struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type ClusterSettingAttributes struct {
	ref terra.Reference
}

func (cs ClusterSettingAttributes) InternalRef() (terra.Reference, error) {
	return cs.ref, nil
}

func (cs ClusterSettingAttributes) InternalWithRef(ref terra.Reference) ClusterSettingAttributes {
	return ClusterSettingAttributes{ref: ref}
}

func (cs ClusterSettingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cs.ref.InternalTokens()
}

func (cs ClusterSettingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("name"))
}

func (cs ClusterSettingAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("value"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type ClusterSettingState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
