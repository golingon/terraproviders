// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datacomputemachinetypes

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type MachineTypes struct {
	// Accelerators: min=0
	Accelerators []Accelerators `hcl:"accelerators,block" validate:"min=0"`
	// BundledLocalSsds: min=0
	BundledLocalSsds []BundledLocalSsds `hcl:"bundled_local_ssds,block" validate:"min=0"`
	// Deprecated: min=0
	Deprecated []Deprecated `hcl:"deprecated,block" validate:"min=0"`
}

type Accelerators struct{}

type BundledLocalSsds struct{}

type Deprecated struct{}

type MachineTypesAttributes struct {
	ref terra.Reference
}

func (mt MachineTypesAttributes) InternalRef() (terra.Reference, error) {
	return mt.ref, nil
}

func (mt MachineTypesAttributes) InternalWithRef(ref terra.Reference) MachineTypesAttributes {
	return MachineTypesAttributes{ref: ref}
}

func (mt MachineTypesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mt.ref.InternalTokens()
}

func (mt MachineTypesAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("description"))
}

func (mt MachineTypesAttributes) GuestCpus() terra.NumberValue {
	return terra.ReferenceAsNumber(mt.ref.Append("guest_cpus"))
}

func (mt MachineTypesAttributes) IsSharedCpus() terra.BoolValue {
	return terra.ReferenceAsBool(mt.ref.Append("is_shared_cpus"))
}

func (mt MachineTypesAttributes) MaximumPersistentDisks() terra.NumberValue {
	return terra.ReferenceAsNumber(mt.ref.Append("maximum_persistent_disks"))
}

func (mt MachineTypesAttributes) MaximumPersistentDisksSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(mt.ref.Append("maximum_persistent_disks_size_gb"))
}

func (mt MachineTypesAttributes) MemoryMb() terra.NumberValue {
	return terra.ReferenceAsNumber(mt.ref.Append("memory_mb"))
}

func (mt MachineTypesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("name"))
}

func (mt MachineTypesAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("self_link"))
}

func (mt MachineTypesAttributes) Accelerators() terra.ListValue[AcceleratorsAttributes] {
	return terra.ReferenceAsList[AcceleratorsAttributes](mt.ref.Append("accelerators"))
}

func (mt MachineTypesAttributes) BundledLocalSsds() terra.SetValue[BundledLocalSsdsAttributes] {
	return terra.ReferenceAsSet[BundledLocalSsdsAttributes](mt.ref.Append("bundled_local_ssds"))
}

func (mt MachineTypesAttributes) Deprecated() terra.SetValue[DeprecatedAttributes] {
	return terra.ReferenceAsSet[DeprecatedAttributes](mt.ref.Append("deprecated"))
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

func (a AcceleratorsAttributes) GuestAcceleratorCount() terra.NumberValue {
	return terra.ReferenceAsNumber(a.ref.Append("guest_accelerator_count"))
}

func (a AcceleratorsAttributes) GuestAcceleratorType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("guest_accelerator_type"))
}

type BundledLocalSsdsAttributes struct {
	ref terra.Reference
}

func (bls BundledLocalSsdsAttributes) InternalRef() (terra.Reference, error) {
	return bls.ref, nil
}

func (bls BundledLocalSsdsAttributes) InternalWithRef(ref terra.Reference) BundledLocalSsdsAttributes {
	return BundledLocalSsdsAttributes{ref: ref}
}

func (bls BundledLocalSsdsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bls.ref.InternalTokens()
}

func (bls BundledLocalSsdsAttributes) DefaultInterface() terra.StringValue {
	return terra.ReferenceAsString(bls.ref.Append("default_interface"))
}

func (bls BundledLocalSsdsAttributes) PartitionCount() terra.NumberValue {
	return terra.ReferenceAsNumber(bls.ref.Append("partition_count"))
}

type DeprecatedAttributes struct {
	ref terra.Reference
}

func (d DeprecatedAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DeprecatedAttributes) InternalWithRef(ref terra.Reference) DeprecatedAttributes {
	return DeprecatedAttributes{ref: ref}
}

func (d DeprecatedAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DeprecatedAttributes) Replacement() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("replacement"))
}

func (d DeprecatedAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("state"))
}

type MachineTypesState struct {
	Description                  string                  `json:"description"`
	GuestCpus                    float64                 `json:"guest_cpus"`
	IsSharedCpus                 bool                    `json:"is_shared_cpus"`
	MaximumPersistentDisks       float64                 `json:"maximum_persistent_disks"`
	MaximumPersistentDisksSizeGb float64                 `json:"maximum_persistent_disks_size_gb"`
	MemoryMb                     float64                 `json:"memory_mb"`
	Name                         string                  `json:"name"`
	SelfLink                     string                  `json:"self_link"`
	Accelerators                 []AcceleratorsState     `json:"accelerators"`
	BundledLocalSsds             []BundledLocalSsdsState `json:"bundled_local_ssds"`
	Deprecated                   []DeprecatedState       `json:"deprecated"`
}

type AcceleratorsState struct {
	GuestAcceleratorCount float64 `json:"guest_accelerator_count"`
	GuestAcceleratorType  string  `json:"guest_accelerator_type"`
}

type BundledLocalSsdsState struct {
	DefaultInterface string  `json:"default_interface"`
	PartitionCount   float64 `json:"partition_count"`
}

type DeprecatedState struct {
	Replacement string `json:"replacement"`
	State       string `json:"state"`
}
