// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datagluedatacatalogencryptionsettings "github.com/golingon/terraproviders/aws/5.0.1/datagluedatacatalogencryptionsettings"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataGlueDataCatalogEncryptionSettings creates a new instance of [DataGlueDataCatalogEncryptionSettings].
func NewDataGlueDataCatalogEncryptionSettings(name string, args DataGlueDataCatalogEncryptionSettingsArgs) *DataGlueDataCatalogEncryptionSettings {
	return &DataGlueDataCatalogEncryptionSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGlueDataCatalogEncryptionSettings)(nil)

// DataGlueDataCatalogEncryptionSettings represents the Terraform data resource aws_glue_data_catalog_encryption_settings.
type DataGlueDataCatalogEncryptionSettings struct {
	Name string
	Args DataGlueDataCatalogEncryptionSettingsArgs
}

// DataSource returns the Terraform object type for [DataGlueDataCatalogEncryptionSettings].
func (gdces *DataGlueDataCatalogEncryptionSettings) DataSource() string {
	return "aws_glue_data_catalog_encryption_settings"
}

// LocalName returns the local name for [DataGlueDataCatalogEncryptionSettings].
func (gdces *DataGlueDataCatalogEncryptionSettings) LocalName() string {
	return gdces.Name
}

// Configuration returns the configuration (args) for [DataGlueDataCatalogEncryptionSettings].
func (gdces *DataGlueDataCatalogEncryptionSettings) Configuration() interface{} {
	return gdces.Args
}

// Attributes returns the attributes for [DataGlueDataCatalogEncryptionSettings].
func (gdces *DataGlueDataCatalogEncryptionSettings) Attributes() dataGlueDataCatalogEncryptionSettingsAttributes {
	return dataGlueDataCatalogEncryptionSettingsAttributes{ref: terra.ReferenceDataResource(gdces)}
}

// DataGlueDataCatalogEncryptionSettingsArgs contains the configurations for aws_glue_data_catalog_encryption_settings.
type DataGlueDataCatalogEncryptionSettingsArgs struct {
	// CatalogId: string, required
	CatalogId terra.StringValue `hcl:"catalog_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// DataCatalogEncryptionSettings: min=0
	DataCatalogEncryptionSettings []datagluedatacatalogencryptionsettings.DataCatalogEncryptionSettings `hcl:"data_catalog_encryption_settings,block" validate:"min=0"`
}
type dataGlueDataCatalogEncryptionSettingsAttributes struct {
	ref terra.Reference
}

// CatalogId returns a reference to field catalog_id of aws_glue_data_catalog_encryption_settings.
func (gdces dataGlueDataCatalogEncryptionSettingsAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(gdces.ref.Append("catalog_id"))
}

// Id returns a reference to field id of aws_glue_data_catalog_encryption_settings.
func (gdces dataGlueDataCatalogEncryptionSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdces.ref.Append("id"))
}

func (gdces dataGlueDataCatalogEncryptionSettingsAttributes) DataCatalogEncryptionSettings() terra.ListValue[datagluedatacatalogencryptionsettings.DataCatalogEncryptionSettingsAttributes] {
	return terra.ReferenceAsList[datagluedatacatalogencryptionsettings.DataCatalogEncryptionSettingsAttributes](gdces.ref.Append("data_catalog_encryption_settings"))
}
