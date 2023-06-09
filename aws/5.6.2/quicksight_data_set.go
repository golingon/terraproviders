// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	quicksightdataset "github.com/golingon/terraproviders/aws/5.6.2/quicksightdataset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewQuicksightDataSet creates a new instance of [QuicksightDataSet].
func NewQuicksightDataSet(name string, args QuicksightDataSetArgs) *QuicksightDataSet {
	return &QuicksightDataSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*QuicksightDataSet)(nil)

// QuicksightDataSet represents the Terraform resource aws_quicksight_data_set.
type QuicksightDataSet struct {
	Name      string
	Args      QuicksightDataSetArgs
	state     *quicksightDataSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [QuicksightDataSet].
func (qds *QuicksightDataSet) Type() string {
	return "aws_quicksight_data_set"
}

// LocalName returns the local name for [QuicksightDataSet].
func (qds *QuicksightDataSet) LocalName() string {
	return qds.Name
}

// Configuration returns the configuration (args) for [QuicksightDataSet].
func (qds *QuicksightDataSet) Configuration() interface{} {
	return qds.Args
}

// DependOn is used for other resources to depend on [QuicksightDataSet].
func (qds *QuicksightDataSet) DependOn() terra.Reference {
	return terra.ReferenceResource(qds)
}

// Dependencies returns the list of resources [QuicksightDataSet] depends_on.
func (qds *QuicksightDataSet) Dependencies() terra.Dependencies {
	return qds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [QuicksightDataSet].
func (qds *QuicksightDataSet) LifecycleManagement() *terra.Lifecycle {
	return qds.Lifecycle
}

// Attributes returns the attributes for [QuicksightDataSet].
func (qds *QuicksightDataSet) Attributes() quicksightDataSetAttributes {
	return quicksightDataSetAttributes{ref: terra.ReferenceResource(qds)}
}

// ImportState imports the given attribute values into [QuicksightDataSet]'s state.
func (qds *QuicksightDataSet) ImportState(av io.Reader) error {
	qds.state = &quicksightDataSetState{}
	if err := json.NewDecoder(av).Decode(qds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", qds.Type(), qds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [QuicksightDataSet] has state.
func (qds *QuicksightDataSet) State() (*quicksightDataSetState, bool) {
	return qds.state, qds.state != nil
}

// StateMust returns the state for [QuicksightDataSet]. Panics if the state is nil.
func (qds *QuicksightDataSet) StateMust() *quicksightDataSetState {
	if qds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", qds.Type(), qds.LocalName()))
	}
	return qds.state
}

// QuicksightDataSetArgs contains the configurations for aws_quicksight_data_set.
type QuicksightDataSetArgs struct {
	// AwsAccountId: string, optional
	AwsAccountId terra.StringValue `hcl:"aws_account_id,attr"`
	// DataSetId: string, required
	DataSetId terra.StringValue `hcl:"data_set_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImportMode: string, required
	ImportMode terra.StringValue `hcl:"import_mode,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// OutputColumns: min=0
	OutputColumns []quicksightdataset.OutputColumns `hcl:"output_columns,block" validate:"min=0"`
	// ColumnGroups: min=0,max=8
	ColumnGroups []quicksightdataset.ColumnGroups `hcl:"column_groups,block" validate:"min=0,max=8"`
	// ColumnLevelPermissionRules: min=0
	ColumnLevelPermissionRules []quicksightdataset.ColumnLevelPermissionRules `hcl:"column_level_permission_rules,block" validate:"min=0"`
	// DataSetUsageConfiguration: optional
	DataSetUsageConfiguration *quicksightdataset.DataSetUsageConfiguration `hcl:"data_set_usage_configuration,block"`
	// FieldFolders: min=0,max=1000
	FieldFolders []quicksightdataset.FieldFolders `hcl:"field_folders,block" validate:"min=0,max=1000"`
	// LogicalTableMap: min=0,max=64
	LogicalTableMap []quicksightdataset.LogicalTableMap `hcl:"logical_table_map,block" validate:"min=0,max=64"`
	// Permissions: min=0,max=64
	Permissions []quicksightdataset.Permissions `hcl:"permissions,block" validate:"min=0,max=64"`
	// PhysicalTableMap: min=0,max=32
	PhysicalTableMap []quicksightdataset.PhysicalTableMap `hcl:"physical_table_map,block" validate:"min=0,max=32"`
	// RefreshProperties: optional
	RefreshProperties *quicksightdataset.RefreshProperties `hcl:"refresh_properties,block"`
	// RowLevelPermissionDataSet: optional
	RowLevelPermissionDataSet *quicksightdataset.RowLevelPermissionDataSet `hcl:"row_level_permission_data_set,block"`
	// RowLevelPermissionTagConfiguration: optional
	RowLevelPermissionTagConfiguration *quicksightdataset.RowLevelPermissionTagConfiguration `hcl:"row_level_permission_tag_configuration,block"`
}
type quicksightDataSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_quicksight_data_set.
func (qds quicksightDataSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("arn"))
}

// AwsAccountId returns a reference to field aws_account_id of aws_quicksight_data_set.
func (qds quicksightDataSetAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("aws_account_id"))
}

// DataSetId returns a reference to field data_set_id of aws_quicksight_data_set.
func (qds quicksightDataSetAttributes) DataSetId() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("data_set_id"))
}

// Id returns a reference to field id of aws_quicksight_data_set.
func (qds quicksightDataSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("id"))
}

// ImportMode returns a reference to field import_mode of aws_quicksight_data_set.
func (qds quicksightDataSetAttributes) ImportMode() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("import_mode"))
}

// Name returns a reference to field name of aws_quicksight_data_set.
func (qds quicksightDataSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_quicksight_data_set.
func (qds quicksightDataSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](qds.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_quicksight_data_set.
func (qds quicksightDataSetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](qds.ref.Append("tags_all"))
}

func (qds quicksightDataSetAttributes) OutputColumns() terra.ListValue[quicksightdataset.OutputColumnsAttributes] {
	return terra.ReferenceAsList[quicksightdataset.OutputColumnsAttributes](qds.ref.Append("output_columns"))
}

func (qds quicksightDataSetAttributes) ColumnGroups() terra.ListValue[quicksightdataset.ColumnGroupsAttributes] {
	return terra.ReferenceAsList[quicksightdataset.ColumnGroupsAttributes](qds.ref.Append("column_groups"))
}

func (qds quicksightDataSetAttributes) ColumnLevelPermissionRules() terra.ListValue[quicksightdataset.ColumnLevelPermissionRulesAttributes] {
	return terra.ReferenceAsList[quicksightdataset.ColumnLevelPermissionRulesAttributes](qds.ref.Append("column_level_permission_rules"))
}

func (qds quicksightDataSetAttributes) DataSetUsageConfiguration() terra.ListValue[quicksightdataset.DataSetUsageConfigurationAttributes] {
	return terra.ReferenceAsList[quicksightdataset.DataSetUsageConfigurationAttributes](qds.ref.Append("data_set_usage_configuration"))
}

func (qds quicksightDataSetAttributes) FieldFolders() terra.SetValue[quicksightdataset.FieldFoldersAttributes] {
	return terra.ReferenceAsSet[quicksightdataset.FieldFoldersAttributes](qds.ref.Append("field_folders"))
}

func (qds quicksightDataSetAttributes) LogicalTableMap() terra.SetValue[quicksightdataset.LogicalTableMapAttributes] {
	return terra.ReferenceAsSet[quicksightdataset.LogicalTableMapAttributes](qds.ref.Append("logical_table_map"))
}

func (qds quicksightDataSetAttributes) Permissions() terra.ListValue[quicksightdataset.PermissionsAttributes] {
	return terra.ReferenceAsList[quicksightdataset.PermissionsAttributes](qds.ref.Append("permissions"))
}

func (qds quicksightDataSetAttributes) PhysicalTableMap() terra.SetValue[quicksightdataset.PhysicalTableMapAttributes] {
	return terra.ReferenceAsSet[quicksightdataset.PhysicalTableMapAttributes](qds.ref.Append("physical_table_map"))
}

func (qds quicksightDataSetAttributes) RefreshProperties() terra.ListValue[quicksightdataset.RefreshPropertiesAttributes] {
	return terra.ReferenceAsList[quicksightdataset.RefreshPropertiesAttributes](qds.ref.Append("refresh_properties"))
}

func (qds quicksightDataSetAttributes) RowLevelPermissionDataSet() terra.ListValue[quicksightdataset.RowLevelPermissionDataSetAttributes] {
	return terra.ReferenceAsList[quicksightdataset.RowLevelPermissionDataSetAttributes](qds.ref.Append("row_level_permission_data_set"))
}

func (qds quicksightDataSetAttributes) RowLevelPermissionTagConfiguration() terra.ListValue[quicksightdataset.RowLevelPermissionTagConfigurationAttributes] {
	return terra.ReferenceAsList[quicksightdataset.RowLevelPermissionTagConfigurationAttributes](qds.ref.Append("row_level_permission_tag_configuration"))
}

type quicksightDataSetState struct {
	Arn                                string                                                      `json:"arn"`
	AwsAccountId                       string                                                      `json:"aws_account_id"`
	DataSetId                          string                                                      `json:"data_set_id"`
	Id                                 string                                                      `json:"id"`
	ImportMode                         string                                                      `json:"import_mode"`
	Name                               string                                                      `json:"name"`
	Tags                               map[string]string                                           `json:"tags"`
	TagsAll                            map[string]string                                           `json:"tags_all"`
	OutputColumns                      []quicksightdataset.OutputColumnsState                      `json:"output_columns"`
	ColumnGroups                       []quicksightdataset.ColumnGroupsState                       `json:"column_groups"`
	ColumnLevelPermissionRules         []quicksightdataset.ColumnLevelPermissionRulesState         `json:"column_level_permission_rules"`
	DataSetUsageConfiguration          []quicksightdataset.DataSetUsageConfigurationState          `json:"data_set_usage_configuration"`
	FieldFolders                       []quicksightdataset.FieldFoldersState                       `json:"field_folders"`
	LogicalTableMap                    []quicksightdataset.LogicalTableMapState                    `json:"logical_table_map"`
	Permissions                        []quicksightdataset.PermissionsState                        `json:"permissions"`
	PhysicalTableMap                   []quicksightdataset.PhysicalTableMapState                   `json:"physical_table_map"`
	RefreshProperties                  []quicksightdataset.RefreshPropertiesState                  `json:"refresh_properties"`
	RowLevelPermissionDataSet          []quicksightdataset.RowLevelPermissionDataSetState          `json:"row_level_permission_data_set"`
	RowLevelPermissionTagConfiguration []quicksightdataset.RowLevelPermissionTagConfigurationState `json:"row_level_permission_tag_configuration"`
}
