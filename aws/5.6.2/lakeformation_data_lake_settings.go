// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lakeformationdatalakesettings "github.com/golingon/terraproviders/aws/5.6.2/lakeformationdatalakesettings"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLakeformationDataLakeSettings creates a new instance of [LakeformationDataLakeSettings].
func NewLakeformationDataLakeSettings(name string, args LakeformationDataLakeSettingsArgs) *LakeformationDataLakeSettings {
	return &LakeformationDataLakeSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LakeformationDataLakeSettings)(nil)

// LakeformationDataLakeSettings represents the Terraform resource aws_lakeformation_data_lake_settings.
type LakeformationDataLakeSettings struct {
	Name      string
	Args      LakeformationDataLakeSettingsArgs
	state     *lakeformationDataLakeSettingsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LakeformationDataLakeSettings].
func (ldls *LakeformationDataLakeSettings) Type() string {
	return "aws_lakeformation_data_lake_settings"
}

// LocalName returns the local name for [LakeformationDataLakeSettings].
func (ldls *LakeformationDataLakeSettings) LocalName() string {
	return ldls.Name
}

// Configuration returns the configuration (args) for [LakeformationDataLakeSettings].
func (ldls *LakeformationDataLakeSettings) Configuration() interface{} {
	return ldls.Args
}

// DependOn is used for other resources to depend on [LakeformationDataLakeSettings].
func (ldls *LakeformationDataLakeSettings) DependOn() terra.Reference {
	return terra.ReferenceResource(ldls)
}

// Dependencies returns the list of resources [LakeformationDataLakeSettings] depends_on.
func (ldls *LakeformationDataLakeSettings) Dependencies() terra.Dependencies {
	return ldls.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LakeformationDataLakeSettings].
func (ldls *LakeformationDataLakeSettings) LifecycleManagement() *terra.Lifecycle {
	return ldls.Lifecycle
}

// Attributes returns the attributes for [LakeformationDataLakeSettings].
func (ldls *LakeformationDataLakeSettings) Attributes() lakeformationDataLakeSettingsAttributes {
	return lakeformationDataLakeSettingsAttributes{ref: terra.ReferenceResource(ldls)}
}

// ImportState imports the given attribute values into [LakeformationDataLakeSettings]'s state.
func (ldls *LakeformationDataLakeSettings) ImportState(av io.Reader) error {
	ldls.state = &lakeformationDataLakeSettingsState{}
	if err := json.NewDecoder(av).Decode(ldls.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ldls.Type(), ldls.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LakeformationDataLakeSettings] has state.
func (ldls *LakeformationDataLakeSettings) State() (*lakeformationDataLakeSettingsState, bool) {
	return ldls.state, ldls.state != nil
}

// StateMust returns the state for [LakeformationDataLakeSettings]. Panics if the state is nil.
func (ldls *LakeformationDataLakeSettings) StateMust() *lakeformationDataLakeSettingsState {
	if ldls.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ldls.Type(), ldls.LocalName()))
	}
	return ldls.state
}

// LakeformationDataLakeSettingsArgs contains the configurations for aws_lakeformation_data_lake_settings.
type LakeformationDataLakeSettingsArgs struct {
	// Admins: set of string, optional
	Admins terra.SetValue[terra.StringValue] `hcl:"admins,attr"`
	// AllowExternalDataFiltering: bool, optional
	AllowExternalDataFiltering terra.BoolValue `hcl:"allow_external_data_filtering,attr"`
	// AuthorizedSessionTagValueList: list of string, optional
	AuthorizedSessionTagValueList terra.ListValue[terra.StringValue] `hcl:"authorized_session_tag_value_list,attr"`
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// ExternalDataFilteringAllowList: set of string, optional
	ExternalDataFilteringAllowList terra.SetValue[terra.StringValue] `hcl:"external_data_filtering_allow_list,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TrustedResourceOwners: list of string, optional
	TrustedResourceOwners terra.ListValue[terra.StringValue] `hcl:"trusted_resource_owners,attr"`
	// CreateDatabaseDefaultPermissions: min=0,max=3
	CreateDatabaseDefaultPermissions []lakeformationdatalakesettings.CreateDatabaseDefaultPermissions `hcl:"create_database_default_permissions,block" validate:"min=0,max=3"`
	// CreateTableDefaultPermissions: min=0,max=3
	CreateTableDefaultPermissions []lakeformationdatalakesettings.CreateTableDefaultPermissions `hcl:"create_table_default_permissions,block" validate:"min=0,max=3"`
}
type lakeformationDataLakeSettingsAttributes struct {
	ref terra.Reference
}

// Admins returns a reference to field admins of aws_lakeformation_data_lake_settings.
func (ldls lakeformationDataLakeSettingsAttributes) Admins() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ldls.ref.Append("admins"))
}

// AllowExternalDataFiltering returns a reference to field allow_external_data_filtering of aws_lakeformation_data_lake_settings.
func (ldls lakeformationDataLakeSettingsAttributes) AllowExternalDataFiltering() terra.BoolValue {
	return terra.ReferenceAsBool(ldls.ref.Append("allow_external_data_filtering"))
}

// AuthorizedSessionTagValueList returns a reference to field authorized_session_tag_value_list of aws_lakeformation_data_lake_settings.
func (ldls lakeformationDataLakeSettingsAttributes) AuthorizedSessionTagValueList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ldls.ref.Append("authorized_session_tag_value_list"))
}

// CatalogId returns a reference to field catalog_id of aws_lakeformation_data_lake_settings.
func (ldls lakeformationDataLakeSettingsAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(ldls.ref.Append("catalog_id"))
}

// ExternalDataFilteringAllowList returns a reference to field external_data_filtering_allow_list of aws_lakeformation_data_lake_settings.
func (ldls lakeformationDataLakeSettingsAttributes) ExternalDataFilteringAllowList() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ldls.ref.Append("external_data_filtering_allow_list"))
}

// Id returns a reference to field id of aws_lakeformation_data_lake_settings.
func (ldls lakeformationDataLakeSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ldls.ref.Append("id"))
}

// TrustedResourceOwners returns a reference to field trusted_resource_owners of aws_lakeformation_data_lake_settings.
func (ldls lakeformationDataLakeSettingsAttributes) TrustedResourceOwners() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ldls.ref.Append("trusted_resource_owners"))
}

func (ldls lakeformationDataLakeSettingsAttributes) CreateDatabaseDefaultPermissions() terra.ListValue[lakeformationdatalakesettings.CreateDatabaseDefaultPermissionsAttributes] {
	return terra.ReferenceAsList[lakeformationdatalakesettings.CreateDatabaseDefaultPermissionsAttributes](ldls.ref.Append("create_database_default_permissions"))
}

func (ldls lakeformationDataLakeSettingsAttributes) CreateTableDefaultPermissions() terra.ListValue[lakeformationdatalakesettings.CreateTableDefaultPermissionsAttributes] {
	return terra.ReferenceAsList[lakeformationdatalakesettings.CreateTableDefaultPermissionsAttributes](ldls.ref.Append("create_table_default_permissions"))
}

type lakeformationDataLakeSettingsState struct {
	Admins                           []string                                                              `json:"admins"`
	AllowExternalDataFiltering       bool                                                                  `json:"allow_external_data_filtering"`
	AuthorizedSessionTagValueList    []string                                                              `json:"authorized_session_tag_value_list"`
	CatalogId                        string                                                                `json:"catalog_id"`
	ExternalDataFilteringAllowList   []string                                                              `json:"external_data_filtering_allow_list"`
	Id                               string                                                                `json:"id"`
	TrustedResourceOwners            []string                                                              `json:"trusted_resource_owners"`
	CreateDatabaseDefaultPermissions []lakeformationdatalakesettings.CreateDatabaseDefaultPermissionsState `json:"create_database_default_permissions"`
	CreateTableDefaultPermissions    []lakeformationdatalakesettings.CreateTableDefaultPermissionsState    `json:"create_table_default_permissions"`
}
