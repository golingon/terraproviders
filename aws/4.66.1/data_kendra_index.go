// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datakendraindex "github.com/golingon/terraproviders/aws/4.66.1/datakendraindex"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKendraIndex creates a new instance of [DataKendraIndex].
func NewDataKendraIndex(name string, args DataKendraIndexArgs) *DataKendraIndex {
	return &DataKendraIndex{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKendraIndex)(nil)

// DataKendraIndex represents the Terraform data resource aws_kendra_index.
type DataKendraIndex struct {
	Name string
	Args DataKendraIndexArgs
}

// DataSource returns the Terraform object type for [DataKendraIndex].
func (ki *DataKendraIndex) DataSource() string {
	return "aws_kendra_index"
}

// LocalName returns the local name for [DataKendraIndex].
func (ki *DataKendraIndex) LocalName() string {
	return ki.Name
}

// Configuration returns the configuration (args) for [DataKendraIndex].
func (ki *DataKendraIndex) Configuration() interface{} {
	return ki.Args
}

// Attributes returns the attributes for [DataKendraIndex].
func (ki *DataKendraIndex) Attributes() dataKendraIndexAttributes {
	return dataKendraIndexAttributes{ref: terra.ReferenceDataResource(ki)}
}

// DataKendraIndexArgs contains the configurations for aws_kendra_index.
type DataKendraIndexArgs struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// CapacityUnits: min=0
	CapacityUnits []datakendraindex.CapacityUnits `hcl:"capacity_units,block" validate:"min=0"`
	// DocumentMetadataConfigurationUpdates: min=0
	DocumentMetadataConfigurationUpdates []datakendraindex.DocumentMetadataConfigurationUpdates `hcl:"document_metadata_configuration_updates,block" validate:"min=0"`
	// IndexStatistics: min=0
	IndexStatistics []datakendraindex.IndexStatistics `hcl:"index_statistics,block" validate:"min=0"`
	// ServerSideEncryptionConfiguration: min=0
	ServerSideEncryptionConfiguration []datakendraindex.ServerSideEncryptionConfiguration `hcl:"server_side_encryption_configuration,block" validate:"min=0"`
	// UserGroupResolutionConfiguration: min=0
	UserGroupResolutionConfiguration []datakendraindex.UserGroupResolutionConfiguration `hcl:"user_group_resolution_configuration,block" validate:"min=0"`
	// UserTokenConfigurations: min=0
	UserTokenConfigurations []datakendraindex.UserTokenConfigurations `hcl:"user_token_configurations,block" validate:"min=0"`
}
type dataKendraIndexAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kendra_index.
func (ki dataKendraIndexAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_kendra_index.
func (ki dataKendraIndexAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("created_at"))
}

// Description returns a reference to field description of aws_kendra_index.
func (ki dataKendraIndexAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("description"))
}

// Edition returns a reference to field edition of aws_kendra_index.
func (ki dataKendraIndexAttributes) Edition() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("edition"))
}

// ErrorMessage returns a reference to field error_message of aws_kendra_index.
func (ki dataKendraIndexAttributes) ErrorMessage() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("error_message"))
}

// Id returns a reference to field id of aws_kendra_index.
func (ki dataKendraIndexAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("id"))
}

// Name returns a reference to field name of aws_kendra_index.
func (ki dataKendraIndexAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_kendra_index.
func (ki dataKendraIndexAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("role_arn"))
}

// Status returns a reference to field status of aws_kendra_index.
func (ki dataKendraIndexAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_kendra_index.
func (ki dataKendraIndexAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ki.ref.Append("tags"))
}

// UpdatedAt returns a reference to field updated_at of aws_kendra_index.
func (ki dataKendraIndexAttributes) UpdatedAt() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("updated_at"))
}

// UserContextPolicy returns a reference to field user_context_policy of aws_kendra_index.
func (ki dataKendraIndexAttributes) UserContextPolicy() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("user_context_policy"))
}

func (ki dataKendraIndexAttributes) CapacityUnits() terra.ListValue[datakendraindex.CapacityUnitsAttributes] {
	return terra.ReferenceAsList[datakendraindex.CapacityUnitsAttributes](ki.ref.Append("capacity_units"))
}

func (ki dataKendraIndexAttributes) DocumentMetadataConfigurationUpdates() terra.SetValue[datakendraindex.DocumentMetadataConfigurationUpdatesAttributes] {
	return terra.ReferenceAsSet[datakendraindex.DocumentMetadataConfigurationUpdatesAttributes](ki.ref.Append("document_metadata_configuration_updates"))
}

func (ki dataKendraIndexAttributes) IndexStatistics() terra.ListValue[datakendraindex.IndexStatisticsAttributes] {
	return terra.ReferenceAsList[datakendraindex.IndexStatisticsAttributes](ki.ref.Append("index_statistics"))
}

func (ki dataKendraIndexAttributes) ServerSideEncryptionConfiguration() terra.ListValue[datakendraindex.ServerSideEncryptionConfigurationAttributes] {
	return terra.ReferenceAsList[datakendraindex.ServerSideEncryptionConfigurationAttributes](ki.ref.Append("server_side_encryption_configuration"))
}

func (ki dataKendraIndexAttributes) UserGroupResolutionConfiguration() terra.ListValue[datakendraindex.UserGroupResolutionConfigurationAttributes] {
	return terra.ReferenceAsList[datakendraindex.UserGroupResolutionConfigurationAttributes](ki.ref.Append("user_group_resolution_configuration"))
}

func (ki dataKendraIndexAttributes) UserTokenConfigurations() terra.ListValue[datakendraindex.UserTokenConfigurationsAttributes] {
	return terra.ReferenceAsList[datakendraindex.UserTokenConfigurationsAttributes](ki.ref.Append("user_token_configurations"))
}
