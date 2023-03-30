// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package iotthinggroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Metadata struct {
	// RootToParentGroups: min=0
	RootToParentGroups []RootToParentGroups `hcl:"root_to_parent_groups,block" validate:"min=0"`
}

type RootToParentGroups struct{}

type Properties struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// AttributePayload: optional
	AttributePayload *AttributePayload `hcl:"attribute_payload,block"`
}

type AttributePayload struct {
	// Attributes: map of string, optional
	Attributes terra.MapValue[terra.StringValue] `hcl:"attributes,attr"`
}

type MetadataAttributes struct {
	ref terra.Reference
}

func (m MetadataAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m MetadataAttributes) InternalWithRef(ref terra.Reference) MetadataAttributes {
	return MetadataAttributes{ref: ref}
}

func (m MetadataAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m MetadataAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("creation_date"))
}

func (m MetadataAttributes) ParentGroupName() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("parent_group_name"))
}

func (m MetadataAttributes) RootToParentGroups() terra.ListValue[RootToParentGroupsAttributes] {
	return terra.ReferenceList[RootToParentGroupsAttributes](m.ref.Append("root_to_parent_groups"))
}

type RootToParentGroupsAttributes struct {
	ref terra.Reference
}

func (rtpg RootToParentGroupsAttributes) InternalRef() terra.Reference {
	return rtpg.ref
}

func (rtpg RootToParentGroupsAttributes) InternalWithRef(ref terra.Reference) RootToParentGroupsAttributes {
	return RootToParentGroupsAttributes{ref: ref}
}

func (rtpg RootToParentGroupsAttributes) InternalTokens() hclwrite.Tokens {
	return rtpg.ref.InternalTokens()
}

func (rtpg RootToParentGroupsAttributes) GroupArn() terra.StringValue {
	return terra.ReferenceString(rtpg.ref.Append("group_arn"))
}

func (rtpg RootToParentGroupsAttributes) GroupName() terra.StringValue {
	return terra.ReferenceString(rtpg.ref.Append("group_name"))
}

type PropertiesAttributes struct {
	ref terra.Reference
}

func (p PropertiesAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p PropertiesAttributes) InternalWithRef(ref terra.Reference) PropertiesAttributes {
	return PropertiesAttributes{ref: ref}
}

func (p PropertiesAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p PropertiesAttributes) Description() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("description"))
}

func (p PropertiesAttributes) AttributePayload() terra.ListValue[AttributePayloadAttributes] {
	return terra.ReferenceList[AttributePayloadAttributes](p.ref.Append("attribute_payload"))
}

type AttributePayloadAttributes struct {
	ref terra.Reference
}

func (ap AttributePayloadAttributes) InternalRef() terra.Reference {
	return ap.ref
}

func (ap AttributePayloadAttributes) InternalWithRef(ref terra.Reference) AttributePayloadAttributes {
	return AttributePayloadAttributes{ref: ref}
}

func (ap AttributePayloadAttributes) InternalTokens() hclwrite.Tokens {
	return ap.ref.InternalTokens()
}

func (ap AttributePayloadAttributes) Attributes() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ap.ref.Append("attributes"))
}

type MetadataState struct {
	CreationDate       string                    `json:"creation_date"`
	ParentGroupName    string                    `json:"parent_group_name"`
	RootToParentGroups []RootToParentGroupsState `json:"root_to_parent_groups"`
}

type RootToParentGroupsState struct {
	GroupArn  string `json:"group_arn"`
	GroupName string `json:"group_name"`
}

type PropertiesState struct {
	Description      string                  `json:"description"`
	AttributePayload []AttributePayloadState `json:"attribute_payload"`
}

type AttributePayloadState struct {
	Attributes map[string]string `json:"attributes"`
}
