// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package firebaseextensionsinstance

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ErrorStatus struct{}

type RuntimeData struct {
	// FatalError: min=0
	FatalError []FatalError `hcl:"fatal_error,block" validate:"min=0"`
	// ProcessingState: min=0
	ProcessingState []ProcessingState `hcl:"processing_state,block" validate:"min=0"`
}

type FatalError struct{}

type ProcessingState struct{}

type Config struct {
	// AllowedEventTypes: list of string, optional
	AllowedEventTypes terra.ListValue[terra.StringValue] `hcl:"allowed_event_types,attr"`
	// EventarcChannel: string, optional
	EventarcChannel terra.StringValue `hcl:"eventarc_channel,attr"`
	// ExtensionRef: string, required
	ExtensionRef terra.StringValue `hcl:"extension_ref,attr" validate:"required"`
	// ExtensionVersion: string, optional
	ExtensionVersion terra.StringValue `hcl:"extension_version,attr"`
	// Params: map of string, required
	Params terra.MapValue[terra.StringValue] `hcl:"params,attr" validate:"required"`
	// SystemParams: map of string, optional
	SystemParams terra.MapValue[terra.StringValue] `hcl:"system_params,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ErrorStatusAttributes struct {
	ref terra.Reference
}

func (es ErrorStatusAttributes) InternalRef() (terra.Reference, error) {
	return es.ref, nil
}

func (es ErrorStatusAttributes) InternalWithRef(ref terra.Reference) ErrorStatusAttributes {
	return ErrorStatusAttributes{ref: ref}
}

func (es ErrorStatusAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return es.ref.InternalTokens()
}

func (es ErrorStatusAttributes) Code() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("code"))
}

func (es ErrorStatusAttributes) Details() terra.ListValue[terra.MapValue[terra.StringValue]] {
	return terra.ReferenceAsList[terra.MapValue[terra.StringValue]](es.ref.Append("details"))
}

func (es ErrorStatusAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("message"))
}

type RuntimeDataAttributes struct {
	ref terra.Reference
}

func (rd RuntimeDataAttributes) InternalRef() (terra.Reference, error) {
	return rd.ref, nil
}

func (rd RuntimeDataAttributes) InternalWithRef(ref terra.Reference) RuntimeDataAttributes {
	return RuntimeDataAttributes{ref: ref}
}

func (rd RuntimeDataAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rd.ref.InternalTokens()
}

func (rd RuntimeDataAttributes) StateUpdateTime() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("state_update_time"))
}

func (rd RuntimeDataAttributes) FatalError() terra.ListValue[FatalErrorAttributes] {
	return terra.ReferenceAsList[FatalErrorAttributes](rd.ref.Append("fatal_error"))
}

func (rd RuntimeDataAttributes) ProcessingState() terra.ListValue[ProcessingStateAttributes] {
	return terra.ReferenceAsList[ProcessingStateAttributes](rd.ref.Append("processing_state"))
}

type FatalErrorAttributes struct {
	ref terra.Reference
}

func (fe FatalErrorAttributes) InternalRef() (terra.Reference, error) {
	return fe.ref, nil
}

func (fe FatalErrorAttributes) InternalWithRef(ref terra.Reference) FatalErrorAttributes {
	return FatalErrorAttributes{ref: ref}
}

func (fe FatalErrorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fe.ref.InternalTokens()
}

func (fe FatalErrorAttributes) ErrorMessage() terra.StringValue {
	return terra.ReferenceAsString(fe.ref.Append("error_message"))
}

type ProcessingStateAttributes struct {
	ref terra.Reference
}

func (ps ProcessingStateAttributes) InternalRef() (terra.Reference, error) {
	return ps.ref, nil
}

func (ps ProcessingStateAttributes) InternalWithRef(ref terra.Reference) ProcessingStateAttributes {
	return ProcessingStateAttributes{ref: ref}
}

func (ps ProcessingStateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ps.ref.InternalTokens()
}

func (ps ProcessingStateAttributes) DetailMessage() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("detail_message"))
}

func (ps ProcessingStateAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("state"))
}

type ConfigAttributes struct {
	ref terra.Reference
}

func (c ConfigAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigAttributes) InternalWithRef(ref terra.Reference) ConfigAttributes {
	return ConfigAttributes{ref: ref}
}

func (c ConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigAttributes) AllowedEventTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("allowed_event_types"))
}

func (c ConfigAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("create_time"))
}

func (c ConfigAttributes) EventarcChannel() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("eventarc_channel"))
}

func (c ConfigAttributes) ExtensionRef() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("extension_ref"))
}

func (c ConfigAttributes) ExtensionVersion() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("extension_version"))
}

func (c ConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c ConfigAttributes) Params() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("params"))
}

func (c ConfigAttributes) PopulatedPostinstallContent() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("populated_postinstall_content"))
}

func (c ConfigAttributes) SystemParams() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("system_params"))
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

type ErrorStatusState struct {
	Code    float64             `json:"code"`
	Details []map[string]string `json:"details"`
	Message string              `json:"message"`
}

type RuntimeDataState struct {
	StateUpdateTime string                 `json:"state_update_time"`
	FatalError      []FatalErrorState      `json:"fatal_error"`
	ProcessingState []ProcessingStateState `json:"processing_state"`
}

type FatalErrorState struct {
	ErrorMessage string `json:"error_message"`
}

type ProcessingStateState struct {
	DetailMessage string `json:"detail_message"`
	State         string `json:"state"`
}

type ConfigState struct {
	AllowedEventTypes           []string          `json:"allowed_event_types"`
	CreateTime                  string            `json:"create_time"`
	EventarcChannel             string            `json:"eventarc_channel"`
	ExtensionRef                string            `json:"extension_ref"`
	ExtensionVersion            string            `json:"extension_version"`
	Name                        string            `json:"name"`
	Params                      map[string]string `json:"params"`
	PopulatedPostinstallContent string            `json:"populated_postinstall_content"`
	SystemParams                map[string]string `json:"system_params"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
