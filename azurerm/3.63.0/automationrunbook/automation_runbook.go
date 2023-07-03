// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package automationrunbook

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type JobSchedule struct {
	// JobScheduleId: string, optional
	JobScheduleId terra.StringValue `hcl:"job_schedule_id,attr"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// RunOn: string, optional
	RunOn terra.StringValue `hcl:"run_on,attr"`
	// ScheduleName: string, optional
	ScheduleName terra.StringValue `hcl:"schedule_name,attr"`
}

type Draft struct {
	// EditModeEnabled: bool, optional
	EditModeEnabled terra.BoolValue `hcl:"edit_mode_enabled,attr"`
	// OutputTypes: list of string, optional
	OutputTypes terra.ListValue[terra.StringValue] `hcl:"output_types,attr"`
	// ContentLink: optional
	ContentLink *ContentLink `hcl:"content_link,block"`
	// Parameters: min=0
	Parameters []Parameters `hcl:"parameters,block" validate:"min=0"`
}

type ContentLink struct {
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// ContentLinkHash: optional
	Hash *ContentLinkHash `hcl:"hash,block"`
}

type ContentLinkHash struct {
	// Algorithm: string, required
	Algorithm terra.StringValue `hcl:"algorithm,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type Parameters struct {
	// DefaultValue: string, optional
	DefaultValue terra.StringValue `hcl:"default_value,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Mandatory: bool, optional
	Mandatory terra.BoolValue `hcl:"mandatory,attr"`
	// Position: number, optional
	Position terra.NumberValue `hcl:"position,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type PublishContentLink struct {
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// PublishContentLinkHash: optional
	Hash *PublishContentLinkHash `hcl:"hash,block"`
}

type PublishContentLinkHash struct {
	// Algorithm: string, required
	Algorithm terra.StringValue `hcl:"algorithm,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type JobScheduleAttributes struct {
	ref terra.Reference
}

func (js JobScheduleAttributes) InternalRef() (terra.Reference, error) {
	return js.ref, nil
}

func (js JobScheduleAttributes) InternalWithRef(ref terra.Reference) JobScheduleAttributes {
	return JobScheduleAttributes{ref: ref}
}

func (js JobScheduleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return js.ref.InternalTokens()
}

func (js JobScheduleAttributes) JobScheduleId() terra.StringValue {
	return terra.ReferenceAsString(js.ref.Append("job_schedule_id"))
}

func (js JobScheduleAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](js.ref.Append("parameters"))
}

func (js JobScheduleAttributes) RunOn() terra.StringValue {
	return terra.ReferenceAsString(js.ref.Append("run_on"))
}

func (js JobScheduleAttributes) ScheduleName() terra.StringValue {
	return terra.ReferenceAsString(js.ref.Append("schedule_name"))
}

type DraftAttributes struct {
	ref terra.Reference
}

func (d DraftAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DraftAttributes) InternalWithRef(ref terra.Reference) DraftAttributes {
	return DraftAttributes{ref: ref}
}

func (d DraftAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DraftAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("creation_time"))
}

func (d DraftAttributes) EditModeEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("edit_mode_enabled"))
}

func (d DraftAttributes) LastModifiedTime() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("last_modified_time"))
}

func (d DraftAttributes) OutputTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("output_types"))
}

func (d DraftAttributes) ContentLink() terra.ListValue[ContentLinkAttributes] {
	return terra.ReferenceAsList[ContentLinkAttributes](d.ref.Append("content_link"))
}

func (d DraftAttributes) Parameters() terra.ListValue[ParametersAttributes] {
	return terra.ReferenceAsList[ParametersAttributes](d.ref.Append("parameters"))
}

type ContentLinkAttributes struct {
	ref terra.Reference
}

func (cl ContentLinkAttributes) InternalRef() (terra.Reference, error) {
	return cl.ref, nil
}

func (cl ContentLinkAttributes) InternalWithRef(ref terra.Reference) ContentLinkAttributes {
	return ContentLinkAttributes{ref: ref}
}

func (cl ContentLinkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cl.ref.InternalTokens()
}

func (cl ContentLinkAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("uri"))
}

func (cl ContentLinkAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("version"))
}

func (cl ContentLinkAttributes) Hash() terra.ListValue[ContentLinkHashAttributes] {
	return terra.ReferenceAsList[ContentLinkHashAttributes](cl.ref.Append("hash"))
}

type ContentLinkHashAttributes struct {
	ref terra.Reference
}

func (h ContentLinkHashAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h ContentLinkHashAttributes) InternalWithRef(ref terra.Reference) ContentLinkHashAttributes {
	return ContentLinkHashAttributes{ref: ref}
}

func (h ContentLinkHashAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h ContentLinkHashAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("algorithm"))
}

func (h ContentLinkHashAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("value"))
}

type ParametersAttributes struct {
	ref terra.Reference
}

func (p ParametersAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ParametersAttributes) InternalWithRef(ref terra.Reference) ParametersAttributes {
	return ParametersAttributes{ref: ref}
}

func (p ParametersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ParametersAttributes) DefaultValue() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("default_value"))
}

func (p ParametersAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("key"))
}

func (p ParametersAttributes) Mandatory() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("mandatory"))
}

func (p ParametersAttributes) Position() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("position"))
}

func (p ParametersAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("type"))
}

type PublishContentLinkAttributes struct {
	ref terra.Reference
}

func (pcl PublishContentLinkAttributes) InternalRef() (terra.Reference, error) {
	return pcl.ref, nil
}

func (pcl PublishContentLinkAttributes) InternalWithRef(ref terra.Reference) PublishContentLinkAttributes {
	return PublishContentLinkAttributes{ref: ref}
}

func (pcl PublishContentLinkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pcl.ref.InternalTokens()
}

func (pcl PublishContentLinkAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(pcl.ref.Append("uri"))
}

func (pcl PublishContentLinkAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(pcl.ref.Append("version"))
}

func (pcl PublishContentLinkAttributes) Hash() terra.ListValue[PublishContentLinkHashAttributes] {
	return terra.ReferenceAsList[PublishContentLinkHashAttributes](pcl.ref.Append("hash"))
}

type PublishContentLinkHashAttributes struct {
	ref terra.Reference
}

func (h PublishContentLinkHashAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h PublishContentLinkHashAttributes) InternalWithRef(ref terra.Reference) PublishContentLinkHashAttributes {
	return PublishContentLinkHashAttributes{ref: ref}
}

func (h PublishContentLinkHashAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h PublishContentLinkHashAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("algorithm"))
}

func (h PublishContentLinkHashAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("value"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type JobScheduleState struct {
	JobScheduleId string            `json:"job_schedule_id"`
	Parameters    map[string]string `json:"parameters"`
	RunOn         string            `json:"run_on"`
	ScheduleName  string            `json:"schedule_name"`
}

type DraftState struct {
	CreationTime     string             `json:"creation_time"`
	EditModeEnabled  bool               `json:"edit_mode_enabled"`
	LastModifiedTime string             `json:"last_modified_time"`
	OutputTypes      []string           `json:"output_types"`
	ContentLink      []ContentLinkState `json:"content_link"`
	Parameters       []ParametersState  `json:"parameters"`
}

type ContentLinkState struct {
	Uri     string                 `json:"uri"`
	Version string                 `json:"version"`
	Hash    []ContentLinkHashState `json:"hash"`
}

type ContentLinkHashState struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
}

type ParametersState struct {
	DefaultValue string  `json:"default_value"`
	Key          string  `json:"key"`
	Mandatory    bool    `json:"mandatory"`
	Position     float64 `json:"position"`
	Type         string  `json:"type"`
}

type PublishContentLinkState struct {
	Uri     string                        `json:"uri"`
	Version string                        `json:"version"`
	Hash    []PublishContentLinkHashState `json:"hash"`
}

type PublishContentLinkHashState struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
