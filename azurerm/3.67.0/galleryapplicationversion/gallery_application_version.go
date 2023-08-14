// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package galleryapplicationversion

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ManageAction struct {
	// Install: string, required
	Install terra.StringValue `hcl:"install,attr" validate:"required"`
	// Remove: string, required
	Remove terra.StringValue `hcl:"remove,attr" validate:"required"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type Source struct {
	// DefaultConfigurationLink: string, optional
	DefaultConfigurationLink terra.StringValue `hcl:"default_configuration_link,attr"`
	// MediaLink: string, required
	MediaLink terra.StringValue `hcl:"media_link,attr" validate:"required"`
}

type TargetRegion struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RegionalReplicaCount: number, required
	RegionalReplicaCount terra.NumberValue `hcl:"regional_replica_count,attr" validate:"required"`
	// StorageAccountType: string, optional
	StorageAccountType terra.StringValue `hcl:"storage_account_type,attr"`
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

type ManageActionAttributes struct {
	ref terra.Reference
}

func (ma ManageActionAttributes) InternalRef() (terra.Reference, error) {
	return ma.ref, nil
}

func (ma ManageActionAttributes) InternalWithRef(ref terra.Reference) ManageActionAttributes {
	return ManageActionAttributes{ref: ref}
}

func (ma ManageActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ma.ref.InternalTokens()
}

func (ma ManageActionAttributes) Install() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("install"))
}

func (ma ManageActionAttributes) Remove() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("remove"))
}

func (ma ManageActionAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("update"))
}

type SourceAttributes struct {
	ref terra.Reference
}

func (s SourceAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SourceAttributes) InternalWithRef(ref terra.Reference) SourceAttributes {
	return SourceAttributes{ref: ref}
}

func (s SourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SourceAttributes) DefaultConfigurationLink() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("default_configuration_link"))
}

func (s SourceAttributes) MediaLink() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("media_link"))
}

type TargetRegionAttributes struct {
	ref terra.Reference
}

func (tr TargetRegionAttributes) InternalRef() (terra.Reference, error) {
	return tr.ref, nil
}

func (tr TargetRegionAttributes) InternalWithRef(ref terra.Reference) TargetRegionAttributes {
	return TargetRegionAttributes{ref: ref}
}

func (tr TargetRegionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tr.ref.InternalTokens()
}

func (tr TargetRegionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("name"))
}

func (tr TargetRegionAttributes) RegionalReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(tr.ref.Append("regional_replica_count"))
}

func (tr TargetRegionAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("storage_account_type"))
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

type ManageActionState struct {
	Install string `json:"install"`
	Remove  string `json:"remove"`
	Update  string `json:"update"`
}

type SourceState struct {
	DefaultConfigurationLink string `json:"default_configuration_link"`
	MediaLink                string `json:"media_link"`
}

type TargetRegionState struct {
	Name                 string  `json:"name"`
	RegionalReplicaCount float64 `json:"regional_replica_count"`
	StorageAccountType   string  `json:"storage_account_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}