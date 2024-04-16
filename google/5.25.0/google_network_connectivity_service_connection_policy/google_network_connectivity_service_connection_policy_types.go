// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_network_connectivity_service_connection_policy

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type PscConfig struct {
	// Limit: string, optional
	Limit terra.StringValue `hcl:"limit,attr"`
	// Subnetworks: list of string, required
	Subnetworks terra.ListValue[terra.StringValue] `hcl:"subnetworks,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type PscConnectionsAttributes struct {
	ref terra.Reference
}

func (pc PscConnectionsAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PscConnectionsAttributes) InternalWithRef(ref terra.Reference) PscConnectionsAttributes {
	return PscConnectionsAttributes{ref: ref}
}

func (pc PscConnectionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PscConnectionsAttributes) ConsumerAddress() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("consumer_address"))
}

func (pc PscConnectionsAttributes) ConsumerForwardingRule() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("consumer_forwarding_rule"))
}

func (pc PscConnectionsAttributes) ConsumerTargetProject() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("consumer_target_project"))
}

func (pc PscConnectionsAttributes) ErrorType() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("error_type"))
}

func (pc PscConnectionsAttributes) GceOperation() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("gce_operation"))
}

func (pc PscConnectionsAttributes) PscConnectionId() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("psc_connection_id"))
}

func (pc PscConnectionsAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("state"))
}

func (pc PscConnectionsAttributes) Error() terra.ListValue[PscConnectionsErrorAttributes] {
	return terra.ReferenceAsList[PscConnectionsErrorAttributes](pc.ref.Append("error"))
}

func (pc PscConnectionsAttributes) ErrorInfo() terra.ListValue[PscConnectionsErrorInfoAttributes] {
	return terra.ReferenceAsList[PscConnectionsErrorInfoAttributes](pc.ref.Append("error_info"))
}

type PscConnectionsErrorAttributes struct {
	ref terra.Reference
}

func (e PscConnectionsErrorAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e PscConnectionsErrorAttributes) InternalWithRef(ref terra.Reference) PscConnectionsErrorAttributes {
	return PscConnectionsErrorAttributes{ref: ref}
}

func (e PscConnectionsErrorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e PscConnectionsErrorAttributes) Code() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("code"))
}

func (e PscConnectionsErrorAttributes) Details() terra.ListValue[terra.MapValue[terra.StringValue]] {
	return terra.ReferenceAsList[terra.MapValue[terra.StringValue]](e.ref.Append("details"))
}

func (e PscConnectionsErrorAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("message"))
}

type PscConnectionsErrorInfoAttributes struct {
	ref terra.Reference
}

func (ei PscConnectionsErrorInfoAttributes) InternalRef() (terra.Reference, error) {
	return ei.ref, nil
}

func (ei PscConnectionsErrorInfoAttributes) InternalWithRef(ref terra.Reference) PscConnectionsErrorInfoAttributes {
	return PscConnectionsErrorInfoAttributes{ref: ref}
}

func (ei PscConnectionsErrorInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ei.ref.InternalTokens()
}

func (ei PscConnectionsErrorInfoAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("domain"))
}

func (ei PscConnectionsErrorInfoAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ei.ref.Append("metadata"))
}

func (ei PscConnectionsErrorInfoAttributes) Reason() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("reason"))
}

type PscConfigAttributes struct {
	ref terra.Reference
}

func (pc PscConfigAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PscConfigAttributes) InternalWithRef(ref terra.Reference) PscConfigAttributes {
	return PscConfigAttributes{ref: ref}
}

func (pc PscConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PscConfigAttributes) Limit() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("limit"))
}

func (pc PscConfigAttributes) Subnetworks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("subnetworks"))
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

type PscConnectionsState struct {
	ConsumerAddress        string                         `json:"consumer_address"`
	ConsumerForwardingRule string                         `json:"consumer_forwarding_rule"`
	ConsumerTargetProject  string                         `json:"consumer_target_project"`
	ErrorType              string                         `json:"error_type"`
	GceOperation           string                         `json:"gce_operation"`
	PscConnectionId        string                         `json:"psc_connection_id"`
	State                  string                         `json:"state"`
	Error                  []PscConnectionsErrorState     `json:"error"`
	ErrorInfo              []PscConnectionsErrorInfoState `json:"error_info"`
}

type PscConnectionsErrorState struct {
	Code    float64             `json:"code"`
	Details []map[string]string `json:"details"`
	Message string              `json:"message"`
}

type PscConnectionsErrorInfoState struct {
	Domain   string            `json:"domain"`
	Metadata map[string]string `json:"metadata"`
	Reason   string            `json:"reason"`
}

type PscConfigState struct {
	Limit       string   `json:"limit"`
	Subnetworks []string `json:"subnetworks"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
