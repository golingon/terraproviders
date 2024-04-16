// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_data_fusion_instance

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Accelerators struct {
	// AcceleratorType: string, required
	AcceleratorType terra.StringValue `hcl:"accelerator_type,attr" validate:"required"`
	// State: string, required
	State terra.StringValue `hcl:"state,attr" validate:"required"`
}

type CryptoKeyConfig struct {
	// KeyReference: string, required
	KeyReference terra.StringValue `hcl:"key_reference,attr" validate:"required"`
}

type EventPublishConfig struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Topic: string, required
	Topic terra.StringValue `hcl:"topic,attr" validate:"required"`
}

type NetworkConfig struct {
	// IpAllocation: string, required
	IpAllocation terra.StringValue `hcl:"ip_allocation,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AcceleratorsAttributes struct {
	ref terra.Reference
}

func (a AcceleratorsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AcceleratorsAttributes) InternalWithRef(ref terra.Reference) AcceleratorsAttributes {
	return AcceleratorsAttributes{ref: ref}
}

func (a AcceleratorsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AcceleratorsAttributes) AcceleratorType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("accelerator_type"))
}

func (a AcceleratorsAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("state"))
}

type CryptoKeyConfigAttributes struct {
	ref terra.Reference
}

func (ckc CryptoKeyConfigAttributes) InternalRef() (terra.Reference, error) {
	return ckc.ref, nil
}

func (ckc CryptoKeyConfigAttributes) InternalWithRef(ref terra.Reference) CryptoKeyConfigAttributes {
	return CryptoKeyConfigAttributes{ref: ref}
}

func (ckc CryptoKeyConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ckc.ref.InternalTokens()
}

func (ckc CryptoKeyConfigAttributes) KeyReference() terra.StringValue {
	return terra.ReferenceAsString(ckc.ref.Append("key_reference"))
}

type EventPublishConfigAttributes struct {
	ref terra.Reference
}

func (epc EventPublishConfigAttributes) InternalRef() (terra.Reference, error) {
	return epc.ref, nil
}

func (epc EventPublishConfigAttributes) InternalWithRef(ref terra.Reference) EventPublishConfigAttributes {
	return EventPublishConfigAttributes{ref: ref}
}

func (epc EventPublishConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return epc.ref.InternalTokens()
}

func (epc EventPublishConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(epc.ref.Append("enabled"))
}

func (epc EventPublishConfigAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(epc.ref.Append("topic"))
}

type NetworkConfigAttributes struct {
	ref terra.Reference
}

func (nc NetworkConfigAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc NetworkConfigAttributes) InternalWithRef(ref terra.Reference) NetworkConfigAttributes {
	return NetworkConfigAttributes{ref: ref}
}

func (nc NetworkConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc NetworkConfigAttributes) IpAllocation() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("ip_allocation"))
}

func (nc NetworkConfigAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("network"))
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

type AcceleratorsState struct {
	AcceleratorType string `json:"accelerator_type"`
	State           string `json:"state"`
}

type CryptoKeyConfigState struct {
	KeyReference string `json:"key_reference"`
}

type EventPublishConfigState struct {
	Enabled bool   `json:"enabled"`
	Topic   string `json:"topic"`
}

type NetworkConfigState struct {
	IpAllocation string `json:"ip_allocation"`
	Network      string `json:"network"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
