// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package fsxopenzfsvolume

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type NfsExports struct {
	// ClientConfigurations: min=1,max=25
	ClientConfigurations []ClientConfigurations `hcl:"client_configurations,block" validate:"min=1,max=25"`
}

type ClientConfigurations struct {
	// Clients: string, required
	Clients terra.StringValue `hcl:"clients,attr" validate:"required"`
	// Options: list of string, required
	Options terra.ListValue[terra.StringValue] `hcl:"options,attr" validate:"required"`
}

type OriginSnapshot struct {
	// CopyStrategy: string, required
	CopyStrategy terra.StringValue `hcl:"copy_strategy,attr" validate:"required"`
	// SnapshotArn: string, required
	SnapshotArn terra.StringValue `hcl:"snapshot_arn,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type UserAndGroupQuotas struct {
	// Id: number, required
	Id terra.NumberValue `hcl:"id,attr" validate:"required"`
	// StorageCapacityQuotaGib: number, required
	StorageCapacityQuotaGib terra.NumberValue `hcl:"storage_capacity_quota_gib,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type NfsExportsAttributes struct {
	ref terra.Reference
}

func (ne NfsExportsAttributes) InternalRef() terra.Reference {
	return ne.ref
}

func (ne NfsExportsAttributes) InternalWithRef(ref terra.Reference) NfsExportsAttributes {
	return NfsExportsAttributes{ref: ref}
}

func (ne NfsExportsAttributes) InternalTokens() hclwrite.Tokens {
	return ne.ref.InternalTokens()
}

func (ne NfsExportsAttributes) ClientConfigurations() terra.SetValue[ClientConfigurationsAttributes] {
	return terra.ReferenceAsSet[ClientConfigurationsAttributes](ne.ref.Append("client_configurations"))
}

type ClientConfigurationsAttributes struct {
	ref terra.Reference
}

func (cc ClientConfigurationsAttributes) InternalRef() terra.Reference {
	return cc.ref
}

func (cc ClientConfigurationsAttributes) InternalWithRef(ref terra.Reference) ClientConfigurationsAttributes {
	return ClientConfigurationsAttributes{ref: ref}
}

func (cc ClientConfigurationsAttributes) InternalTokens() hclwrite.Tokens {
	return cc.ref.InternalTokens()
}

func (cc ClientConfigurationsAttributes) Clients() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("clients"))
}

func (cc ClientConfigurationsAttributes) Options() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cc.ref.Append("options"))
}

type OriginSnapshotAttributes struct {
	ref terra.Reference
}

func (os OriginSnapshotAttributes) InternalRef() terra.Reference {
	return os.ref
}

func (os OriginSnapshotAttributes) InternalWithRef(ref terra.Reference) OriginSnapshotAttributes {
	return OriginSnapshotAttributes{ref: ref}
}

func (os OriginSnapshotAttributes) InternalTokens() hclwrite.Tokens {
	return os.ref.InternalTokens()
}

func (os OriginSnapshotAttributes) CopyStrategy() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("copy_strategy"))
}

func (os OriginSnapshotAttributes) SnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("snapshot_arn"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
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

type UserAndGroupQuotasAttributes struct {
	ref terra.Reference
}

func (uagq UserAndGroupQuotasAttributes) InternalRef() terra.Reference {
	return uagq.ref
}

func (uagq UserAndGroupQuotasAttributes) InternalWithRef(ref terra.Reference) UserAndGroupQuotasAttributes {
	return UserAndGroupQuotasAttributes{ref: ref}
}

func (uagq UserAndGroupQuotasAttributes) InternalTokens() hclwrite.Tokens {
	return uagq.ref.InternalTokens()
}

func (uagq UserAndGroupQuotasAttributes) Id() terra.NumberValue {
	return terra.ReferenceAsNumber(uagq.ref.Append("id"))
}

func (uagq UserAndGroupQuotasAttributes) StorageCapacityQuotaGib() terra.NumberValue {
	return terra.ReferenceAsNumber(uagq.ref.Append("storage_capacity_quota_gib"))
}

func (uagq UserAndGroupQuotasAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(uagq.ref.Append("type"))
}

type NfsExportsState struct {
	ClientConfigurations []ClientConfigurationsState `json:"client_configurations"`
}

type ClientConfigurationsState struct {
	Clients string   `json:"clients"`
	Options []string `json:"options"`
}

type OriginSnapshotState struct {
	CopyStrategy string `json:"copy_strategy"`
	SnapshotArn  string `json:"snapshot_arn"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type UserAndGroupQuotasState struct {
	Id                      float64 `json:"id"`
	StorageCapacityQuotaGib float64 `json:"storage_capacity_quota_gib"`
	Type                    string  `json:"type"`
}