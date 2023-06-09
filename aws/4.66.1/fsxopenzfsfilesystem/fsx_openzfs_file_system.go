// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package fsxopenzfsfilesystem

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DiskIopsConfiguration struct {
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
}

type RootVolumeConfiguration struct {
	// CopyTagsToSnapshots: bool, optional
	CopyTagsToSnapshots terra.BoolValue `hcl:"copy_tags_to_snapshots,attr"`
	// DataCompressionType: string, optional
	DataCompressionType terra.StringValue `hcl:"data_compression_type,attr"`
	// ReadOnly: bool, optional
	ReadOnly terra.BoolValue `hcl:"read_only,attr"`
	// RecordSizeKib: number, optional
	RecordSizeKib terra.NumberValue `hcl:"record_size_kib,attr"`
	// NfsExports: optional
	NfsExports *NfsExports `hcl:"nfs_exports,block"`
	// UserAndGroupQuotas: min=0,max=100
	UserAndGroupQuotas []UserAndGroupQuotas `hcl:"user_and_group_quotas,block" validate:"min=0,max=100"`
}

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

type UserAndGroupQuotas struct {
	// Id: number, required
	Id terra.NumberValue `hcl:"id,attr" validate:"required"`
	// StorageCapacityQuotaGib: number, required
	StorageCapacityQuotaGib terra.NumberValue `hcl:"storage_capacity_quota_gib,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DiskIopsConfigurationAttributes struct {
	ref terra.Reference
}

func (dic DiskIopsConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return dic.ref, nil
}

func (dic DiskIopsConfigurationAttributes) InternalWithRef(ref terra.Reference) DiskIopsConfigurationAttributes {
	return DiskIopsConfigurationAttributes{ref: ref}
}

func (dic DiskIopsConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dic.ref.InternalTokens()
}

func (dic DiskIopsConfigurationAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(dic.ref.Append("iops"))
}

func (dic DiskIopsConfigurationAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(dic.ref.Append("mode"))
}

type RootVolumeConfigurationAttributes struct {
	ref terra.Reference
}

func (rvc RootVolumeConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return rvc.ref, nil
}

func (rvc RootVolumeConfigurationAttributes) InternalWithRef(ref terra.Reference) RootVolumeConfigurationAttributes {
	return RootVolumeConfigurationAttributes{ref: ref}
}

func (rvc RootVolumeConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rvc.ref.InternalTokens()
}

func (rvc RootVolumeConfigurationAttributes) CopyTagsToSnapshots() terra.BoolValue {
	return terra.ReferenceAsBool(rvc.ref.Append("copy_tags_to_snapshots"))
}

func (rvc RootVolumeConfigurationAttributes) DataCompressionType() terra.StringValue {
	return terra.ReferenceAsString(rvc.ref.Append("data_compression_type"))
}

func (rvc RootVolumeConfigurationAttributes) ReadOnly() terra.BoolValue {
	return terra.ReferenceAsBool(rvc.ref.Append("read_only"))
}

func (rvc RootVolumeConfigurationAttributes) RecordSizeKib() terra.NumberValue {
	return terra.ReferenceAsNumber(rvc.ref.Append("record_size_kib"))
}

func (rvc RootVolumeConfigurationAttributes) NfsExports() terra.ListValue[NfsExportsAttributes] {
	return terra.ReferenceAsList[NfsExportsAttributes](rvc.ref.Append("nfs_exports"))
}

func (rvc RootVolumeConfigurationAttributes) UserAndGroupQuotas() terra.SetValue[UserAndGroupQuotasAttributes] {
	return terra.ReferenceAsSet[UserAndGroupQuotasAttributes](rvc.ref.Append("user_and_group_quotas"))
}

type NfsExportsAttributes struct {
	ref terra.Reference
}

func (ne NfsExportsAttributes) InternalRef() (terra.Reference, error) {
	return ne.ref, nil
}

func (ne NfsExportsAttributes) InternalWithRef(ref terra.Reference) NfsExportsAttributes {
	return NfsExportsAttributes{ref: ref}
}

func (ne NfsExportsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ne.ref.InternalTokens()
}

func (ne NfsExportsAttributes) ClientConfigurations() terra.SetValue[ClientConfigurationsAttributes] {
	return terra.ReferenceAsSet[ClientConfigurationsAttributes](ne.ref.Append("client_configurations"))
}

type ClientConfigurationsAttributes struct {
	ref terra.Reference
}

func (cc ClientConfigurationsAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc ClientConfigurationsAttributes) InternalWithRef(ref terra.Reference) ClientConfigurationsAttributes {
	return ClientConfigurationsAttributes{ref: ref}
}

func (cc ClientConfigurationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc ClientConfigurationsAttributes) Clients() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("clients"))
}

func (cc ClientConfigurationsAttributes) Options() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cc.ref.Append("options"))
}

type UserAndGroupQuotasAttributes struct {
	ref terra.Reference
}

func (uagq UserAndGroupQuotasAttributes) InternalRef() (terra.Reference, error) {
	return uagq.ref, nil
}

func (uagq UserAndGroupQuotasAttributes) InternalWithRef(ref terra.Reference) UserAndGroupQuotasAttributes {
	return UserAndGroupQuotasAttributes{ref: ref}
}

func (uagq UserAndGroupQuotasAttributes) InternalTokens() (hclwrite.Tokens, error) {
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

type DiskIopsConfigurationState struct {
	Iops float64 `json:"iops"`
	Mode string  `json:"mode"`
}

type RootVolumeConfigurationState struct {
	CopyTagsToSnapshots bool                      `json:"copy_tags_to_snapshots"`
	DataCompressionType string                    `json:"data_compression_type"`
	ReadOnly            bool                      `json:"read_only"`
	RecordSizeKib       float64                   `json:"record_size_kib"`
	NfsExports          []NfsExportsState         `json:"nfs_exports"`
	UserAndGroupQuotas  []UserAndGroupQuotasState `json:"user_and_group_quotas"`
}

type NfsExportsState struct {
	ClientConfigurations []ClientConfigurationsState `json:"client_configurations"`
}

type ClientConfigurationsState struct {
	Clients string   `json:"clients"`
	Options []string `json:"options"`
}

type UserAndGroupQuotasState struct {
	Id                      float64 `json:"id"`
	StorageCapacityQuotaGib float64 `json:"storage_capacity_quota_gib"`
	Type                    string  `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
