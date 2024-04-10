// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package computediskasyncreplication

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type SecondaryDisk struct {
	// Disk: string, required
	Disk terra.StringValue `hcl:"disk,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type SecondaryDiskAttributes struct {
	ref terra.Reference
}

func (sd SecondaryDiskAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd SecondaryDiskAttributes) InternalWithRef(ref terra.Reference) SecondaryDiskAttributes {
	return SecondaryDiskAttributes{ref: ref}
}

func (sd SecondaryDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sd.ref.InternalTokens()
}

func (sd SecondaryDiskAttributes) Disk() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("disk"))
}

func (sd SecondaryDiskAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("state"))
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

type SecondaryDiskState struct {
	Disk  string `json:"disk"`
	State string `json:"state"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
