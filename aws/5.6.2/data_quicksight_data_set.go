// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataquicksightdataset "github.com/golingon/terraproviders/aws/5.6.2/dataquicksightdataset"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataQuicksightDataSet creates a new instance of [DataQuicksightDataSet].
func NewDataQuicksightDataSet(name string, args DataQuicksightDataSetArgs) *DataQuicksightDataSet {
	return &DataQuicksightDataSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataQuicksightDataSet)(nil)

// DataQuicksightDataSet represents the Terraform data resource aws_quicksight_data_set.
type DataQuicksightDataSet struct {
	Name string
	Args DataQuicksightDataSetArgs
}

// DataSource returns the Terraform object type for [DataQuicksightDataSet].
func (qds *DataQuicksightDataSet) DataSource() string {
	return "aws_quicksight_data_set"
}

// LocalName returns the local name for [DataQuicksightDataSet].
func (qds *DataQuicksightDataSet) LocalName() string {
	return qds.Name
}

// Configuration returns the configuration (args) for [DataQuicksightDataSet].
func (qds *DataQuicksightDataSet) Configuration() interface{} {
	return qds.Args
}

// Attributes returns the attributes for [DataQuicksightDataSet].
func (qds *DataQuicksightDataSet) Attributes() dataQuicksightDataSetAttributes {
	return dataQuicksightDataSetAttributes{ref: terra.ReferenceDataResource(qds)}
}

// DataQuicksightDataSetArgs contains the configurations for aws_quicksight_data_set.
type DataQuicksightDataSetArgs struct {
	// AwsAccountId: string, optional
	AwsAccountId terra.StringValue `hcl:"aws_account_id,attr"`
	// DataSetId: string, required
	DataSetId terra.StringValue `hcl:"data_set_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ColumnGroups: min=0
	ColumnGroups []dataquicksightdataset.ColumnGroups `hcl:"column_groups,block" validate:"min=0"`
	// DataSetUsageConfiguration: min=0
	DataSetUsageConfiguration []dataquicksightdataset.DataSetUsageConfiguration `hcl:"data_set_usage_configuration,block" validate:"min=0"`
	// FieldFolders: min=0
	FieldFolders []dataquicksightdataset.FieldFolders `hcl:"field_folders,block" validate:"min=0"`
	// LogicalTableMap: min=0
	LogicalTableMap []dataquicksightdataset.LogicalTableMap `hcl:"logical_table_map,block" validate:"min=0"`
	// Permissions: min=0
	Permissions []dataquicksightdataset.Permissions `hcl:"permissions,block" validate:"min=0"`
	// PhysicalTableMap: min=0
	PhysicalTableMap []dataquicksightdataset.PhysicalTableMap `hcl:"physical_table_map,block" validate:"min=0"`
	// RowLevelPermissionDataSet: min=0
	RowLevelPermissionDataSet []dataquicksightdataset.RowLevelPermissionDataSet `hcl:"row_level_permission_data_set,block" validate:"min=0"`
	// RowLevelPermissionTagConfiguration: min=0
	RowLevelPermissionTagConfiguration []dataquicksightdataset.RowLevelPermissionTagConfiguration `hcl:"row_level_permission_tag_configuration,block" validate:"min=0"`
	// ColumnLevelPermissionRules: min=0
	ColumnLevelPermissionRules []dataquicksightdataset.ColumnLevelPermissionRules `hcl:"column_level_permission_rules,block" validate:"min=0"`
}
type dataQuicksightDataSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_quicksight_data_set.
func (qds dataQuicksightDataSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("arn"))
}

// AwsAccountId returns a reference to field aws_account_id of aws_quicksight_data_set.
func (qds dataQuicksightDataSetAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("aws_account_id"))
}

// DataSetId returns a reference to field data_set_id of aws_quicksight_data_set.
func (qds dataQuicksightDataSetAttributes) DataSetId() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("data_set_id"))
}

// Id returns a reference to field id of aws_quicksight_data_set.
func (qds dataQuicksightDataSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("id"))
}

// ImportMode returns a reference to field import_mode of aws_quicksight_data_set.
func (qds dataQuicksightDataSetAttributes) ImportMode() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("import_mode"))
}

// Name returns a reference to field name of aws_quicksight_data_set.
func (qds dataQuicksightDataSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_quicksight_data_set.
func (qds dataQuicksightDataSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](qds.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_quicksight_data_set.
func (qds dataQuicksightDataSetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](qds.ref.Append("tags_all"))
}

func (qds dataQuicksightDataSetAttributes) ColumnGroups() terra.ListValue[dataquicksightdataset.ColumnGroupsAttributes] {
	return terra.ReferenceAsList[dataquicksightdataset.ColumnGroupsAttributes](qds.ref.Append("column_groups"))
}

func (qds dataQuicksightDataSetAttributes) DataSetUsageConfiguration() terra.ListValue[dataquicksightdataset.DataSetUsageConfigurationAttributes] {
	return terra.ReferenceAsList[dataquicksightdataset.DataSetUsageConfigurationAttributes](qds.ref.Append("data_set_usage_configuration"))
}

func (qds dataQuicksightDataSetAttributes) FieldFolders() terra.SetValue[dataquicksightdataset.FieldFoldersAttributes] {
	return terra.ReferenceAsSet[dataquicksightdataset.FieldFoldersAttributes](qds.ref.Append("field_folders"))
}

func (qds dataQuicksightDataSetAttributes) LogicalTableMap() terra.SetValue[dataquicksightdataset.LogicalTableMapAttributes] {
	return terra.ReferenceAsSet[dataquicksightdataset.LogicalTableMapAttributes](qds.ref.Append("logical_table_map"))
}

func (qds dataQuicksightDataSetAttributes) Permissions() terra.ListValue[dataquicksightdataset.PermissionsAttributes] {
	return terra.ReferenceAsList[dataquicksightdataset.PermissionsAttributes](qds.ref.Append("permissions"))
}

func (qds dataQuicksightDataSetAttributes) PhysicalTableMap() terra.SetValue[dataquicksightdataset.PhysicalTableMapAttributes] {
	return terra.ReferenceAsSet[dataquicksightdataset.PhysicalTableMapAttributes](qds.ref.Append("physical_table_map"))
}

func (qds dataQuicksightDataSetAttributes) RowLevelPermissionDataSet() terra.ListValue[dataquicksightdataset.RowLevelPermissionDataSetAttributes] {
	return terra.ReferenceAsList[dataquicksightdataset.RowLevelPermissionDataSetAttributes](qds.ref.Append("row_level_permission_data_set"))
}

func (qds dataQuicksightDataSetAttributes) RowLevelPermissionTagConfiguration() terra.ListValue[dataquicksightdataset.RowLevelPermissionTagConfigurationAttributes] {
	return terra.ReferenceAsList[dataquicksightdataset.RowLevelPermissionTagConfigurationAttributes](qds.ref.Append("row_level_permission_tag_configuration"))
}

func (qds dataQuicksightDataSetAttributes) ColumnLevelPermissionRules() terra.ListValue[dataquicksightdataset.ColumnLevelPermissionRulesAttributes] {
	return terra.ReferenceAsList[dataquicksightdataset.ColumnLevelPermissionRulesAttributes](qds.ref.Append("column_level_permission_rules"))
}
