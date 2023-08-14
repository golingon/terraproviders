// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package notebooksinstance

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AcceleratorConfig struct {
	// CoreCount: number, required
	CoreCount terra.NumberValue `hcl:"core_count,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type ContainerImage struct {
	// Repository: string, required
	Repository terra.StringValue `hcl:"repository,attr" validate:"required"`
	// Tag: string, optional
	Tag terra.StringValue `hcl:"tag,attr"`
}

type ReservationAffinity struct {
	// ConsumeReservationType: string, required
	ConsumeReservationType terra.StringValue `hcl:"consume_reservation_type,attr" validate:"required"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Values: list of string, optional
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr"`
}

type ShieldedInstanceConfig struct {
	// EnableIntegrityMonitoring: bool, optional
	EnableIntegrityMonitoring terra.BoolValue `hcl:"enable_integrity_monitoring,attr"`
	// EnableSecureBoot: bool, optional
	EnableSecureBoot terra.BoolValue `hcl:"enable_secure_boot,attr"`
	// EnableVtpm: bool, optional
	EnableVtpm terra.BoolValue `hcl:"enable_vtpm,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type VmImage struct {
	// ImageFamily: string, optional
	ImageFamily terra.StringValue `hcl:"image_family,attr"`
	// ImageName: string, optional
	ImageName terra.StringValue `hcl:"image_name,attr"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
}

type AcceleratorConfigAttributes struct {
	ref terra.Reference
}

func (ac AcceleratorConfigAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AcceleratorConfigAttributes) InternalWithRef(ref terra.Reference) AcceleratorConfigAttributes {
	return AcceleratorConfigAttributes{ref: ref}
}

func (ac AcceleratorConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AcceleratorConfigAttributes) CoreCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ac.ref.Append("core_count"))
}

func (ac AcceleratorConfigAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("type"))
}

type ContainerImageAttributes struct {
	ref terra.Reference
}

func (ci ContainerImageAttributes) InternalRef() (terra.Reference, error) {
	return ci.ref, nil
}

func (ci ContainerImageAttributes) InternalWithRef(ref terra.Reference) ContainerImageAttributes {
	return ContainerImageAttributes{ref: ref}
}

func (ci ContainerImageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ci.ref.InternalTokens()
}

func (ci ContainerImageAttributes) Repository() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("repository"))
}

func (ci ContainerImageAttributes) Tag() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("tag"))
}

type ReservationAffinityAttributes struct {
	ref terra.Reference
}

func (ra ReservationAffinityAttributes) InternalRef() (terra.Reference, error) {
	return ra.ref, nil
}

func (ra ReservationAffinityAttributes) InternalWithRef(ref terra.Reference) ReservationAffinityAttributes {
	return ReservationAffinityAttributes{ref: ref}
}

func (ra ReservationAffinityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ra.ref.InternalTokens()
}

func (ra ReservationAffinityAttributes) ConsumeReservationType() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("consume_reservation_type"))
}

func (ra ReservationAffinityAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("key"))
}

func (ra ReservationAffinityAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ra.ref.Append("values"))
}

type ShieldedInstanceConfigAttributes struct {
	ref terra.Reference
}

func (sic ShieldedInstanceConfigAttributes) InternalRef() (terra.Reference, error) {
	return sic.ref, nil
}

func (sic ShieldedInstanceConfigAttributes) InternalWithRef(ref terra.Reference) ShieldedInstanceConfigAttributes {
	return ShieldedInstanceConfigAttributes{ref: ref}
}

func (sic ShieldedInstanceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sic.ref.InternalTokens()
}

func (sic ShieldedInstanceConfigAttributes) EnableIntegrityMonitoring() terra.BoolValue {
	return terra.ReferenceAsBool(sic.ref.Append("enable_integrity_monitoring"))
}

func (sic ShieldedInstanceConfigAttributes) EnableSecureBoot() terra.BoolValue {
	return terra.ReferenceAsBool(sic.ref.Append("enable_secure_boot"))
}

func (sic ShieldedInstanceConfigAttributes) EnableVtpm() terra.BoolValue {
	return terra.ReferenceAsBool(sic.ref.Append("enable_vtpm"))
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

type VmImageAttributes struct {
	ref terra.Reference
}

func (vi VmImageAttributes) InternalRef() (terra.Reference, error) {
	return vi.ref, nil
}

func (vi VmImageAttributes) InternalWithRef(ref terra.Reference) VmImageAttributes {
	return VmImageAttributes{ref: ref}
}

func (vi VmImageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vi.ref.InternalTokens()
}

func (vi VmImageAttributes) ImageFamily() terra.StringValue {
	return terra.ReferenceAsString(vi.ref.Append("image_family"))
}

func (vi VmImageAttributes) ImageName() terra.StringValue {
	return terra.ReferenceAsString(vi.ref.Append("image_name"))
}

func (vi VmImageAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vi.ref.Append("project"))
}

type AcceleratorConfigState struct {
	CoreCount float64 `json:"core_count"`
	Type      string  `json:"type"`
}

type ContainerImageState struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type ReservationAffinityState struct {
	ConsumeReservationType string   `json:"consume_reservation_type"`
	Key                    string   `json:"key"`
	Values                 []string `json:"values"`
}

type ShieldedInstanceConfigState struct {
	EnableIntegrityMonitoring bool `json:"enable_integrity_monitoring"`
	EnableSecureBoot          bool `json:"enable_secure_boot"`
	EnableVtpm                bool `json:"enable_vtpm"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type VmImageState struct {
	ImageFamily string `json:"image_family"`
	ImageName   string `json:"image_name"`
	Project     string `json:"project"`
}
