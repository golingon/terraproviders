// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package healthcarehl7v2store

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type NotificationConfig struct {
	// PubsubTopic: string, required
	PubsubTopic terra.StringValue `hcl:"pubsub_topic,attr" validate:"required"`
}

type NotificationConfigs struct {
	// Filter: string, optional
	Filter terra.StringValue `hcl:"filter,attr"`
	// PubsubTopic: string, required
	PubsubTopic terra.StringValue `hcl:"pubsub_topic,attr" validate:"required"`
}

type ParserConfig struct {
	// AllowNullHeader: bool, optional
	AllowNullHeader terra.BoolValue `hcl:"allow_null_header,attr"`
	// Schema: string, optional
	Schema terra.StringValue `hcl:"schema,attr"`
	// SegmentTerminator: string, optional
	SegmentTerminator terra.StringValue `hcl:"segment_terminator,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type NotificationConfigAttributes struct {
	ref terra.Reference
}

func (nc NotificationConfigAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc NotificationConfigAttributes) InternalWithRef(ref terra.Reference) NotificationConfigAttributes {
	return NotificationConfigAttributes{ref: ref}
}

func (nc NotificationConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc NotificationConfigAttributes) PubsubTopic() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("pubsub_topic"))
}

type NotificationConfigsAttributes struct {
	ref terra.Reference
}

func (nc NotificationConfigsAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc NotificationConfigsAttributes) InternalWithRef(ref terra.Reference) NotificationConfigsAttributes {
	return NotificationConfigsAttributes{ref: ref}
}

func (nc NotificationConfigsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc NotificationConfigsAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("filter"))
}

func (nc NotificationConfigsAttributes) PubsubTopic() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("pubsub_topic"))
}

type ParserConfigAttributes struct {
	ref terra.Reference
}

func (pc ParserConfigAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc ParserConfigAttributes) InternalWithRef(ref terra.Reference) ParserConfigAttributes {
	return ParserConfigAttributes{ref: ref}
}

func (pc ParserConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc ParserConfigAttributes) AllowNullHeader() terra.BoolValue {
	return terra.ReferenceAsBool(pc.ref.Append("allow_null_header"))
}

func (pc ParserConfigAttributes) Schema() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("schema"))
}

func (pc ParserConfigAttributes) SegmentTerminator() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("segment_terminator"))
}

func (pc ParserConfigAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("version"))
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

type NotificationConfigState struct {
	PubsubTopic string `json:"pubsub_topic"`
}

type NotificationConfigsState struct {
	Filter      string `json:"filter"`
	PubsubTopic string `json:"pubsub_topic"`
}

type ParserConfigState struct {
	AllowNullHeader   bool   `json:"allow_null_header"`
	Schema            string `json:"schema"`
	SegmentTerminator string `json:"segment_terminator"`
	Version           string `json:"version"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
