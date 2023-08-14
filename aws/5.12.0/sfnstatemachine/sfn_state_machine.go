// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sfnstatemachine

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type LoggingConfiguration struct {
	// IncludeExecutionData: bool, optional
	IncludeExecutionData terra.BoolValue `hcl:"include_execution_data,attr"`
	// Level: string, optional
	Level terra.StringValue `hcl:"level,attr"`
	// LogDestination: string, optional
	LogDestination terra.StringValue `hcl:"log_destination,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type TracingConfiguration struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type LoggingConfigurationAttributes struct {
	ref terra.Reference
}

func (lc LoggingConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc LoggingConfigurationAttributes) InternalWithRef(ref terra.Reference) LoggingConfigurationAttributes {
	return LoggingConfigurationAttributes{ref: ref}
}

func (lc LoggingConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc LoggingConfigurationAttributes) IncludeExecutionData() terra.BoolValue {
	return terra.ReferenceAsBool(lc.ref.Append("include_execution_data"))
}

func (lc LoggingConfigurationAttributes) Level() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("level"))
}

func (lc LoggingConfigurationAttributes) LogDestination() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("log_destination"))
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

type TracingConfigurationAttributes struct {
	ref terra.Reference
}

func (tc TracingConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return tc.ref, nil
}

func (tc TracingConfigurationAttributes) InternalWithRef(ref terra.Reference) TracingConfigurationAttributes {
	return TracingConfigurationAttributes{ref: ref}
}

func (tc TracingConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tc.ref.InternalTokens()
}

func (tc TracingConfigurationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(tc.ref.Append("enabled"))
}

type LoggingConfigurationState struct {
	IncludeExecutionData bool   `json:"include_execution_data"`
	Level                string `json:"level"`
	LogDestination       string `json:"log_destination"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type TracingConfigurationState struct {
	Enabled bool `json:"enabled"`
}