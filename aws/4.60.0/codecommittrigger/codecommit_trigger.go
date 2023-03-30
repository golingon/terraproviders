// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package codecommittrigger

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Trigger struct {
	// Branches: list of string, optional
	Branches terra.ListValue[terra.StringValue] `hcl:"branches,attr"`
	// CustomData: string, optional
	CustomData terra.StringValue `hcl:"custom_data,attr"`
	// DestinationArn: string, required
	DestinationArn terra.StringValue `hcl:"destination_arn,attr" validate:"required"`
	// Events: list of string, required
	Events terra.ListValue[terra.StringValue] `hcl:"events,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type TriggerAttributes struct {
	ref terra.Reference
}

func (t TriggerAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TriggerAttributes) InternalWithRef(ref terra.Reference) TriggerAttributes {
	return TriggerAttributes{ref: ref}
}

func (t TriggerAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TriggerAttributes) Branches() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](t.ref.Append("branches"))
}

func (t TriggerAttributes) CustomData() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("custom_data"))
}

func (t TriggerAttributes) DestinationArn() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("destination_arn"))
}

func (t TriggerAttributes) Events() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](t.ref.Append("events"))
}

func (t TriggerAttributes) Name() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("name"))
}

type TriggerState struct {
	Branches       []string `json:"branches"`
	CustomData     string   `json:"custom_data"`
	DestinationArn string   `json:"destination_arn"`
	Events         []string `json:"events"`
	Name           string   `json:"name"`
}
