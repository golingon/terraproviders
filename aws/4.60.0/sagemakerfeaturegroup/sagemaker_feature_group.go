// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sagemakerfeaturegroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type FeatureDefinition struct {
	// FeatureName: string, optional
	FeatureName terra.StringValue `hcl:"feature_name,attr"`
	// FeatureType: string, optional
	FeatureType terra.StringValue `hcl:"feature_type,attr"`
}

type OfflineStoreConfig struct {
	// DisableGlueTableCreation: bool, optional
	DisableGlueTableCreation terra.BoolValue `hcl:"disable_glue_table_creation,attr"`
	// DataCatalogConfig: optional
	DataCatalogConfig *DataCatalogConfig `hcl:"data_catalog_config,block"`
	// S3StorageConfig: required
	S3StorageConfig *S3StorageConfig `hcl:"s3_storage_config,block" validate:"required"`
}

type DataCatalogConfig struct {
	// Catalog: string, optional
	Catalog terra.StringValue `hcl:"catalog,attr"`
	// Database: string, optional
	Database terra.StringValue `hcl:"database,attr"`
	// TableName: string, optional
	TableName terra.StringValue `hcl:"table_name,attr"`
}

type S3StorageConfig struct {
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// S3Uri: string, required
	S3Uri terra.StringValue `hcl:"s3_uri,attr" validate:"required"`
}

type OnlineStoreConfig struct {
	// EnableOnlineStore: bool, optional
	EnableOnlineStore terra.BoolValue `hcl:"enable_online_store,attr"`
	// SecurityConfig: optional
	SecurityConfig *SecurityConfig `hcl:"security_config,block"`
}

type SecurityConfig struct {
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
}

type FeatureDefinitionAttributes struct {
	ref terra.Reference
}

func (fd FeatureDefinitionAttributes) InternalRef() (terra.Reference, error) {
	return fd.ref, nil
}

func (fd FeatureDefinitionAttributes) InternalWithRef(ref terra.Reference) FeatureDefinitionAttributes {
	return FeatureDefinitionAttributes{ref: ref}
}

func (fd FeatureDefinitionAttributes) InternalTokens() hclwrite.Tokens {
	return fd.ref.InternalTokens()
}

func (fd FeatureDefinitionAttributes) FeatureName() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("feature_name"))
}

func (fd FeatureDefinitionAttributes) FeatureType() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("feature_type"))
}

type OfflineStoreConfigAttributes struct {
	ref terra.Reference
}

func (osc OfflineStoreConfigAttributes) InternalRef() (terra.Reference, error) {
	return osc.ref, nil
}

func (osc OfflineStoreConfigAttributes) InternalWithRef(ref terra.Reference) OfflineStoreConfigAttributes {
	return OfflineStoreConfigAttributes{ref: ref}
}

func (osc OfflineStoreConfigAttributes) InternalTokens() hclwrite.Tokens {
	return osc.ref.InternalTokens()
}

func (osc OfflineStoreConfigAttributes) DisableGlueTableCreation() terra.BoolValue {
	return terra.ReferenceAsBool(osc.ref.Append("disable_glue_table_creation"))
}

func (osc OfflineStoreConfigAttributes) DataCatalogConfig() terra.ListValue[DataCatalogConfigAttributes] {
	return terra.ReferenceAsList[DataCatalogConfigAttributes](osc.ref.Append("data_catalog_config"))
}

func (osc OfflineStoreConfigAttributes) S3StorageConfig() terra.ListValue[S3StorageConfigAttributes] {
	return terra.ReferenceAsList[S3StorageConfigAttributes](osc.ref.Append("s3_storage_config"))
}

type DataCatalogConfigAttributes struct {
	ref terra.Reference
}

func (dcc DataCatalogConfigAttributes) InternalRef() (terra.Reference, error) {
	return dcc.ref, nil
}

func (dcc DataCatalogConfigAttributes) InternalWithRef(ref terra.Reference) DataCatalogConfigAttributes {
	return DataCatalogConfigAttributes{ref: ref}
}

func (dcc DataCatalogConfigAttributes) InternalTokens() hclwrite.Tokens {
	return dcc.ref.InternalTokens()
}

func (dcc DataCatalogConfigAttributes) Catalog() terra.StringValue {
	return terra.ReferenceAsString(dcc.ref.Append("catalog"))
}

func (dcc DataCatalogConfigAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(dcc.ref.Append("database"))
}

func (dcc DataCatalogConfigAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(dcc.ref.Append("table_name"))
}

type S3StorageConfigAttributes struct {
	ref terra.Reference
}

func (ssc S3StorageConfigAttributes) InternalRef() (terra.Reference, error) {
	return ssc.ref, nil
}

func (ssc S3StorageConfigAttributes) InternalWithRef(ref terra.Reference) S3StorageConfigAttributes {
	return S3StorageConfigAttributes{ref: ref}
}

func (ssc S3StorageConfigAttributes) InternalTokens() hclwrite.Tokens {
	return ssc.ref.InternalTokens()
}

func (ssc S3StorageConfigAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("kms_key_id"))
}

func (ssc S3StorageConfigAttributes) S3Uri() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("s3_uri"))
}

type OnlineStoreConfigAttributes struct {
	ref terra.Reference
}

func (osc OnlineStoreConfigAttributes) InternalRef() (terra.Reference, error) {
	return osc.ref, nil
}

func (osc OnlineStoreConfigAttributes) InternalWithRef(ref terra.Reference) OnlineStoreConfigAttributes {
	return OnlineStoreConfigAttributes{ref: ref}
}

func (osc OnlineStoreConfigAttributes) InternalTokens() hclwrite.Tokens {
	return osc.ref.InternalTokens()
}

func (osc OnlineStoreConfigAttributes) EnableOnlineStore() terra.BoolValue {
	return terra.ReferenceAsBool(osc.ref.Append("enable_online_store"))
}

func (osc OnlineStoreConfigAttributes) SecurityConfig() terra.ListValue[SecurityConfigAttributes] {
	return terra.ReferenceAsList[SecurityConfigAttributes](osc.ref.Append("security_config"))
}

type SecurityConfigAttributes struct {
	ref terra.Reference
}

func (sc SecurityConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SecurityConfigAttributes) InternalWithRef(ref terra.Reference) SecurityConfigAttributes {
	return SecurityConfigAttributes{ref: ref}
}

func (sc SecurityConfigAttributes) InternalTokens() hclwrite.Tokens {
	return sc.ref.InternalTokens()
}

func (sc SecurityConfigAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("kms_key_id"))
}

type FeatureDefinitionState struct {
	FeatureName string `json:"feature_name"`
	FeatureType string `json:"feature_type"`
}

type OfflineStoreConfigState struct {
	DisableGlueTableCreation bool                     `json:"disable_glue_table_creation"`
	DataCatalogConfig        []DataCatalogConfigState `json:"data_catalog_config"`
	S3StorageConfig          []S3StorageConfigState   `json:"s3_storage_config"`
}

type DataCatalogConfigState struct {
	Catalog   string `json:"catalog"`
	Database  string `json:"database"`
	TableName string `json:"table_name"`
}

type S3StorageConfigState struct {
	KmsKeyId string `json:"kms_key_id"`
	S3Uri    string `json:"s3_uri"`
}

type OnlineStoreConfigState struct {
	EnableOnlineStore bool                  `json:"enable_online_store"`
	SecurityConfig    []SecurityConfigState `json:"security_config"`
}

type SecurityConfigState struct {
	KmsKeyId string `json:"kms_key_id"`
}
