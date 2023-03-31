// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datamariadbserver

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type StorageProfile struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type StorageProfileAttributes struct {
	ref terra.Reference
}

func (sp StorageProfileAttributes) InternalRef() terra.Reference {
	return sp.ref
}

func (sp StorageProfileAttributes) InternalWithRef(ref terra.Reference) StorageProfileAttributes {
	return StorageProfileAttributes{ref: ref}
}

func (sp StorageProfileAttributes) InternalTokens() hclwrite.Tokens {
	return sp.ref.InternalTokens()
}

func (sp StorageProfileAttributes) AutoGrow() terra.StringValue {
	return terra.ReferenceString(sp.ref.Append("auto_grow"))
}

func (sp StorageProfileAttributes) BackupRetentionDays() terra.NumberValue {
	return terra.ReferenceNumber(sp.ref.Append("backup_retention_days"))
}

func (sp StorageProfileAttributes) GeoRedundantBackup() terra.StringValue {
	return terra.ReferenceString(sp.ref.Append("geo_redundant_backup"))
}

func (sp StorageProfileAttributes) StorageMb() terra.NumberValue {
	return terra.ReferenceNumber(sp.ref.Append("storage_mb"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

type StorageProfileState struct {
	AutoGrow            string  `json:"auto_grow"`
	BackupRetentionDays float64 `json:"backup_retention_days"`
	GeoRedundantBackup  string  `json:"geo_redundant_backup"`
	StorageMb           float64 `json:"storage_mb"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
