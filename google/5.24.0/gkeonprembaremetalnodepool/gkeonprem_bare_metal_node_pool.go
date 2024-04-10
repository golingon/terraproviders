// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package gkeonprembaremetalnodepool

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Status struct {
	// Conditions: min=0
	Conditions []Conditions `hcl:"conditions,block" validate:"min=0"`
}

type Conditions struct{}

type NodePoolConfig struct {
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// OperatingSystem: string, optional
	OperatingSystem terra.StringValue `hcl:"operating_system,attr"`
	// NodeConfigs: min=1
	NodeConfigs []NodeConfigs `hcl:"node_configs,block" validate:"min=1"`
	// Taints: min=0
	Taints []Taints `hcl:"taints,block" validate:"min=0"`
}

type NodeConfigs struct {
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// NodeIp: string, optional
	NodeIp terra.StringValue `hcl:"node_ip,attr"`
}

type Taints struct {
	// Effect: string, optional
	Effect terra.StringValue `hcl:"effect,attr"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type StatusAttributes struct {
	ref terra.Reference
}

func (s StatusAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StatusAttributes) InternalWithRef(ref terra.Reference) StatusAttributes {
	return StatusAttributes{ref: ref}
}

func (s StatusAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StatusAttributes) ErrorMessage() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("error_message"))
}

func (s StatusAttributes) Conditions() terra.ListValue[ConditionsAttributes] {
	return terra.ReferenceAsList[ConditionsAttributes](s.ref.Append("conditions"))
}

type ConditionsAttributes struct {
	ref terra.Reference
}

func (c ConditionsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConditionsAttributes) InternalWithRef(ref terra.Reference) ConditionsAttributes {
	return ConditionsAttributes{ref: ref}
}

func (c ConditionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConditionsAttributes) LastTransitionTime() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("last_transition_time"))
}

func (c ConditionsAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("message"))
}

func (c ConditionsAttributes) Reason() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("reason"))
}

func (c ConditionsAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("state"))
}

func (c ConditionsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("type"))
}

type NodePoolConfigAttributes struct {
	ref terra.Reference
}

func (npc NodePoolConfigAttributes) InternalRef() (terra.Reference, error) {
	return npc.ref, nil
}

func (npc NodePoolConfigAttributes) InternalWithRef(ref terra.Reference) NodePoolConfigAttributes {
	return NodePoolConfigAttributes{ref: ref}
}

func (npc NodePoolConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return npc.ref.InternalTokens()
}

func (npc NodePoolConfigAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](npc.ref.Append("labels"))
}

func (npc NodePoolConfigAttributes) OperatingSystem() terra.StringValue {
	return terra.ReferenceAsString(npc.ref.Append("operating_system"))
}

func (npc NodePoolConfigAttributes) NodeConfigs() terra.ListValue[NodeConfigsAttributes] {
	return terra.ReferenceAsList[NodeConfigsAttributes](npc.ref.Append("node_configs"))
}

func (npc NodePoolConfigAttributes) Taints() terra.ListValue[TaintsAttributes] {
	return terra.ReferenceAsList[TaintsAttributes](npc.ref.Append("taints"))
}

type NodeConfigsAttributes struct {
	ref terra.Reference
}

func (nc NodeConfigsAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc NodeConfigsAttributes) InternalWithRef(ref terra.Reference) NodeConfigsAttributes {
	return NodeConfigsAttributes{ref: ref}
}

func (nc NodeConfigsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc NodeConfigsAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nc.ref.Append("labels"))
}

func (nc NodeConfigsAttributes) NodeIp() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("node_ip"))
}

type TaintsAttributes struct {
	ref terra.Reference
}

func (t TaintsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TaintsAttributes) InternalWithRef(ref terra.Reference) TaintsAttributes {
	return TaintsAttributes{ref: ref}
}

func (t TaintsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TaintsAttributes) Effect() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("effect"))
}

func (t TaintsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("key"))
}

func (t TaintsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("value"))
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

type StatusState struct {
	ErrorMessage string            `json:"error_message"`
	Conditions   []ConditionsState `json:"conditions"`
}

type ConditionsState struct {
	LastTransitionTime string `json:"last_transition_time"`
	Message            string `json:"message"`
	Reason             string `json:"reason"`
	State              string `json:"state"`
	Type               string `json:"type"`
}

type NodePoolConfigState struct {
	Labels          map[string]string  `json:"labels"`
	OperatingSystem string             `json:"operating_system"`
	NodeConfigs     []NodeConfigsState `json:"node_configs"`
	Taints          []TaintsState      `json:"taints"`
}

type NodeConfigsState struct {
	Labels map[string]string `json:"labels"`
	NodeIp string            `json:"node_ip"`
}

type TaintsState struct {
	Effect string `json:"effect"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
