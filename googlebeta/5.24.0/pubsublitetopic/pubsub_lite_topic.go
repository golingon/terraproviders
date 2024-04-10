// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package pubsublitetopic

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type PartitionConfig struct {
	// Count: number, required
	Count terra.NumberValue `hcl:"count,attr" validate:"required"`
	// Capacity: optional
	Capacity *Capacity `hcl:"capacity,block"`
}

type Capacity struct {
	// PublishMibPerSec: number, required
	PublishMibPerSec terra.NumberValue `hcl:"publish_mib_per_sec,attr" validate:"required"`
	// SubscribeMibPerSec: number, required
	SubscribeMibPerSec terra.NumberValue `hcl:"subscribe_mib_per_sec,attr" validate:"required"`
}

type ReservationConfig struct {
	// ThroughputReservation: string, optional
	ThroughputReservation terra.StringValue `hcl:"throughput_reservation,attr"`
}

type RetentionConfig struct {
	// PerPartitionBytes: string, required
	PerPartitionBytes terra.StringValue `hcl:"per_partition_bytes,attr" validate:"required"`
	// Period: string, optional
	Period terra.StringValue `hcl:"period,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type PartitionConfigAttributes struct {
	ref terra.Reference
}

func (pc PartitionConfigAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PartitionConfigAttributes) InternalWithRef(ref terra.Reference) PartitionConfigAttributes {
	return PartitionConfigAttributes{ref: ref}
}

func (pc PartitionConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PartitionConfigAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(pc.ref.Append("count"))
}

func (pc PartitionConfigAttributes) Capacity() terra.ListValue[CapacityAttributes] {
	return terra.ReferenceAsList[CapacityAttributes](pc.ref.Append("capacity"))
}

type CapacityAttributes struct {
	ref terra.Reference
}

func (c CapacityAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CapacityAttributes) InternalWithRef(ref terra.Reference) CapacityAttributes {
	return CapacityAttributes{ref: ref}
}

func (c CapacityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CapacityAttributes) PublishMibPerSec() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("publish_mib_per_sec"))
}

func (c CapacityAttributes) SubscribeMibPerSec() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("subscribe_mib_per_sec"))
}

type ReservationConfigAttributes struct {
	ref terra.Reference
}

func (rc ReservationConfigAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc ReservationConfigAttributes) InternalWithRef(ref terra.Reference) ReservationConfigAttributes {
	return ReservationConfigAttributes{ref: ref}
}

func (rc ReservationConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc ReservationConfigAttributes) ThroughputReservation() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("throughput_reservation"))
}

type RetentionConfigAttributes struct {
	ref terra.Reference
}

func (rc RetentionConfigAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc RetentionConfigAttributes) InternalWithRef(ref terra.Reference) RetentionConfigAttributes {
	return RetentionConfigAttributes{ref: ref}
}

func (rc RetentionConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc RetentionConfigAttributes) PerPartitionBytes() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("per_partition_bytes"))
}

func (rc RetentionConfigAttributes) Period() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("period"))
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

type PartitionConfigState struct {
	Count    float64         `json:"count"`
	Capacity []CapacityState `json:"capacity"`
}

type CapacityState struct {
	PublishMibPerSec   float64 `json:"publish_mib_per_sec"`
	SubscribeMibPerSec float64 `json:"subscribe_mib_per_sec"`
}

type ReservationConfigState struct {
	ThroughputReservation string `json:"throughput_reservation"`
}

type RetentionConfigState struct {
	PerPartitionBytes string `json:"per_partition_bytes"`
	Period            string `json:"period"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
