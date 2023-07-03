// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBackupRegionSettings creates a new instance of [BackupRegionSettings].
func NewBackupRegionSettings(name string, args BackupRegionSettingsArgs) *BackupRegionSettings {
	return &BackupRegionSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BackupRegionSettings)(nil)

// BackupRegionSettings represents the Terraform resource aws_backup_region_settings.
type BackupRegionSettings struct {
	Name      string
	Args      BackupRegionSettingsArgs
	state     *backupRegionSettingsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BackupRegionSettings].
func (brs *BackupRegionSettings) Type() string {
	return "aws_backup_region_settings"
}

// LocalName returns the local name for [BackupRegionSettings].
func (brs *BackupRegionSettings) LocalName() string {
	return brs.Name
}

// Configuration returns the configuration (args) for [BackupRegionSettings].
func (brs *BackupRegionSettings) Configuration() interface{} {
	return brs.Args
}

// DependOn is used for other resources to depend on [BackupRegionSettings].
func (brs *BackupRegionSettings) DependOn() terra.Reference {
	return terra.ReferenceResource(brs)
}

// Dependencies returns the list of resources [BackupRegionSettings] depends_on.
func (brs *BackupRegionSettings) Dependencies() terra.Dependencies {
	return brs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BackupRegionSettings].
func (brs *BackupRegionSettings) LifecycleManagement() *terra.Lifecycle {
	return brs.Lifecycle
}

// Attributes returns the attributes for [BackupRegionSettings].
func (brs *BackupRegionSettings) Attributes() backupRegionSettingsAttributes {
	return backupRegionSettingsAttributes{ref: terra.ReferenceResource(brs)}
}

// ImportState imports the given attribute values into [BackupRegionSettings]'s state.
func (brs *BackupRegionSettings) ImportState(av io.Reader) error {
	brs.state = &backupRegionSettingsState{}
	if err := json.NewDecoder(av).Decode(brs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", brs.Type(), brs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BackupRegionSettings] has state.
func (brs *BackupRegionSettings) State() (*backupRegionSettingsState, bool) {
	return brs.state, brs.state != nil
}

// StateMust returns the state for [BackupRegionSettings]. Panics if the state is nil.
func (brs *BackupRegionSettings) StateMust() *backupRegionSettingsState {
	if brs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", brs.Type(), brs.LocalName()))
	}
	return brs.state
}

// BackupRegionSettingsArgs contains the configurations for aws_backup_region_settings.
type BackupRegionSettingsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceTypeManagementPreference: map of bool, optional
	ResourceTypeManagementPreference terra.MapValue[terra.BoolValue] `hcl:"resource_type_management_preference,attr"`
	// ResourceTypeOptInPreference: map of bool, required
	ResourceTypeOptInPreference terra.MapValue[terra.BoolValue] `hcl:"resource_type_opt_in_preference,attr" validate:"required"`
}
type backupRegionSettingsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_backup_region_settings.
func (brs backupRegionSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(brs.ref.Append("id"))
}

// ResourceTypeManagementPreference returns a reference to field resource_type_management_preference of aws_backup_region_settings.
func (brs backupRegionSettingsAttributes) ResourceTypeManagementPreference() terra.MapValue[terra.BoolValue] {
	return terra.ReferenceAsMap[terra.BoolValue](brs.ref.Append("resource_type_management_preference"))
}

// ResourceTypeOptInPreference returns a reference to field resource_type_opt_in_preference of aws_backup_region_settings.
func (brs backupRegionSettingsAttributes) ResourceTypeOptInPreference() terra.MapValue[terra.BoolValue] {
	return terra.ReferenceAsMap[terra.BoolValue](brs.ref.Append("resource_type_opt_in_preference"))
}

type backupRegionSettingsState struct {
	Id                               string          `json:"id"`
	ResourceTypeManagementPreference map[string]bool `json:"resource_type_management_preference"`
	ResourceTypeOptInPreference      map[string]bool `json:"resource_type_opt_in_preference"`
}
