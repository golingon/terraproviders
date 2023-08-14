// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dialogflowcxversion

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type NluSettings struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type NluSettingsAttributes struct {
	ref terra.Reference
}

func (ns NluSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ns.ref, nil
}

func (ns NluSettingsAttributes) InternalWithRef(ref terra.Reference) NluSettingsAttributes {
	return NluSettingsAttributes{ref: ref}
}

func (ns NluSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ns.ref.InternalTokens()
}

func (ns NluSettingsAttributes) ClassificationThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(ns.ref.Append("classification_threshold"))
}

func (ns NluSettingsAttributes) ModelTrainingMode() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("model_training_mode"))
}

func (ns NluSettingsAttributes) ModelType() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("model_type"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type NluSettingsState struct {
	ClassificationThreshold float64 `json:"classification_threshold"`
	ModelTrainingMode       string  `json:"model_training_mode"`
	ModelType               string  `json:"model_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}