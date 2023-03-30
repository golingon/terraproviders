// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewBackupRegionSettings(name string, args BackupRegionSettingsArgs) *BackupRegionSettings {
	return &BackupRegionSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BackupRegionSettings)(nil)

type BackupRegionSettings struct {
	Name  string
	Args  BackupRegionSettingsArgs
	state *backupRegionSettingsState
}

func (brs *BackupRegionSettings) Type() string {
	return "aws_backup_region_settings"
}

func (brs *BackupRegionSettings) LocalName() string {
	return brs.Name
}

func (brs *BackupRegionSettings) Configuration() interface{} {
	return brs.Args
}

func (brs *BackupRegionSettings) Attributes() backupRegionSettingsAttributes {
	return backupRegionSettingsAttributes{ref: terra.ReferenceResource(brs)}
}

func (brs *BackupRegionSettings) ImportState(av io.Reader) error {
	brs.state = &backupRegionSettingsState{}
	if err := json.NewDecoder(av).Decode(brs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", brs.Type(), brs.LocalName(), err)
	}
	return nil
}

func (brs *BackupRegionSettings) State() (*backupRegionSettingsState, bool) {
	return brs.state, brs.state != nil
}

func (brs *BackupRegionSettings) StateMust() *backupRegionSettingsState {
	if brs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", brs.Type(), brs.LocalName()))
	}
	return brs.state
}

func (brs *BackupRegionSettings) DependOn() terra.Reference {
	return terra.ReferenceResource(brs)
}

type BackupRegionSettingsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceTypeManagementPreference: map of bool, optional
	ResourceTypeManagementPreference terra.MapValue[terra.BoolValue] `hcl:"resource_type_management_preference,attr"`
	// ResourceTypeOptInPreference: map of bool, required
	ResourceTypeOptInPreference terra.MapValue[terra.BoolValue] `hcl:"resource_type_opt_in_preference,attr" validate:"required"`
	// DependsOn contains resources that BackupRegionSettings depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type backupRegionSettingsAttributes struct {
	ref terra.Reference
}

func (brs backupRegionSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(brs.ref.Append("id"))
}

func (brs backupRegionSettingsAttributes) ResourceTypeManagementPreference() terra.MapValue[terra.BoolValue] {
	return terra.ReferenceMap[terra.BoolValue](brs.ref.Append("resource_type_management_preference"))
}

func (brs backupRegionSettingsAttributes) ResourceTypeOptInPreference() terra.MapValue[terra.BoolValue] {
	return terra.ReferenceMap[terra.BoolValue](brs.ref.Append("resource_type_opt_in_preference"))
}

type backupRegionSettingsState struct {
	Id                               string          `json:"id"`
	ResourceTypeManagementPreference map[string]bool `json:"resource_type_management_preference"`
	ResourceTypeOptInPreference      map[string]bool `json:"resource_type_opt_in_preference"`
}
