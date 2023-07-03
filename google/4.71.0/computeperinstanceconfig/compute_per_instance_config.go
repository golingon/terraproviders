// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package computeperinstanceconfig

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PreservedState struct {
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Disk: min=0
	Disk []Disk `hcl:"disk,block" validate:"min=0"`
}

type Disk struct {
	// DeleteRule: string, optional
	DeleteRule terra.StringValue `hcl:"delete_rule,attr"`
	// DeviceName: string, required
	DeviceName terra.StringValue `hcl:"device_name,attr" validate:"required"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type PreservedStateAttributes struct {
	ref terra.Reference
}

func (ps PreservedStateAttributes) InternalRef() (terra.Reference, error) {
	return ps.ref, nil
}

func (ps PreservedStateAttributes) InternalWithRef(ref terra.Reference) PreservedStateAttributes {
	return PreservedStateAttributes{ref: ref}
}

func (ps PreservedStateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ps.ref.InternalTokens()
}

func (ps PreservedStateAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ps.ref.Append("metadata"))
}

func (ps PreservedStateAttributes) Disk() terra.SetValue[DiskAttributes] {
	return terra.ReferenceAsSet[DiskAttributes](ps.ref.Append("disk"))
}

type DiskAttributes struct {
	ref terra.Reference
}

func (d DiskAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DiskAttributes) InternalWithRef(ref terra.Reference) DiskAttributes {
	return DiskAttributes{ref: ref}
}

func (d DiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DiskAttributes) DeleteRule() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("delete_rule"))
}

func (d DiskAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("device_name"))
}

func (d DiskAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("mode"))
}

func (d DiskAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("source"))
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

type PreservedStateState struct {
	Metadata map[string]string `json:"metadata"`
	Disk     []DiskState       `json:"disk"`
}

type DiskState struct {
	DeleteRule string `json:"delete_rule"`
	DeviceName string `json:"device_name"`
	Mode       string `json:"mode"`
	Source     string `json:"source"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
