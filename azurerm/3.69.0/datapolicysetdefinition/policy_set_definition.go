// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datapolicysetdefinition

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PolicyDefinitionGroup struct{}

type PolicyDefinitionReference struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type PolicyDefinitionGroupAttributes struct {
	ref terra.Reference
}

func (pdg PolicyDefinitionGroupAttributes) InternalRef() (terra.Reference, error) {
	return pdg.ref, nil
}

func (pdg PolicyDefinitionGroupAttributes) InternalWithRef(ref terra.Reference) PolicyDefinitionGroupAttributes {
	return PolicyDefinitionGroupAttributes{ref: ref}
}

func (pdg PolicyDefinitionGroupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pdg.ref.InternalTokens()
}

func (pdg PolicyDefinitionGroupAttributes) AdditionalMetadataResourceId() terra.StringValue {
	return terra.ReferenceAsString(pdg.ref.Append("additional_metadata_resource_id"))
}

func (pdg PolicyDefinitionGroupAttributes) Category() terra.StringValue {
	return terra.ReferenceAsString(pdg.ref.Append("category"))
}

func (pdg PolicyDefinitionGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(pdg.ref.Append("description"))
}

func (pdg PolicyDefinitionGroupAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(pdg.ref.Append("display_name"))
}

func (pdg PolicyDefinitionGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdg.ref.Append("name"))
}

type PolicyDefinitionReferenceAttributes struct {
	ref terra.Reference
}

func (pdr PolicyDefinitionReferenceAttributes) InternalRef() (terra.Reference, error) {
	return pdr.ref, nil
}

func (pdr PolicyDefinitionReferenceAttributes) InternalWithRef(ref terra.Reference) PolicyDefinitionReferenceAttributes {
	return PolicyDefinitionReferenceAttributes{ref: ref}
}

func (pdr PolicyDefinitionReferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pdr.ref.InternalTokens()
}

func (pdr PolicyDefinitionReferenceAttributes) ParameterValues() terra.StringValue {
	return terra.ReferenceAsString(pdr.ref.Append("parameter_values"))
}

func (pdr PolicyDefinitionReferenceAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdr.ref.Append("parameters"))
}

func (pdr PolicyDefinitionReferenceAttributes) PolicyDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(pdr.ref.Append("policy_definition_id"))
}

func (pdr PolicyDefinitionReferenceAttributes) PolicyGroupNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pdr.ref.Append("policy_group_names"))
}

func (pdr PolicyDefinitionReferenceAttributes) ReferenceId() terra.StringValue {
	return terra.ReferenceAsString(pdr.ref.Append("reference_id"))
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

type PolicyDefinitionGroupState struct {
	AdditionalMetadataResourceId string `json:"additional_metadata_resource_id"`
	Category                     string `json:"category"`
	Description                  string `json:"description"`
	DisplayName                  string `json:"display_name"`
	Name                         string `json:"name"`
}

type PolicyDefinitionReferenceState struct {
	ParameterValues    string            `json:"parameter_values"`
	Parameters         map[string]string `json:"parameters"`
	PolicyDefinitionId string            `json:"policy_definition_id"`
	PolicyGroupNames   []string          `json:"policy_group_names"`
	ReferenceId        string            `json:"reference_id"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}