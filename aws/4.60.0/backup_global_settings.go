// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewBackupGlobalSettings(name string, args BackupGlobalSettingsArgs) *BackupGlobalSettings {
	return &BackupGlobalSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BackupGlobalSettings)(nil)

type BackupGlobalSettings struct {
	Name  string
	Args  BackupGlobalSettingsArgs
	state *backupGlobalSettingsState
}

func (bgs *BackupGlobalSettings) Type() string {
	return "aws_backup_global_settings"
}

func (bgs *BackupGlobalSettings) LocalName() string {
	return bgs.Name
}

func (bgs *BackupGlobalSettings) Configuration() interface{} {
	return bgs.Args
}

func (bgs *BackupGlobalSettings) Attributes() backupGlobalSettingsAttributes {
	return backupGlobalSettingsAttributes{ref: terra.ReferenceResource(bgs)}
}

func (bgs *BackupGlobalSettings) ImportState(av io.Reader) error {
	bgs.state = &backupGlobalSettingsState{}
	if err := json.NewDecoder(av).Decode(bgs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bgs.Type(), bgs.LocalName(), err)
	}
	return nil
}

func (bgs *BackupGlobalSettings) State() (*backupGlobalSettingsState, bool) {
	return bgs.state, bgs.state != nil
}

func (bgs *BackupGlobalSettings) StateMust() *backupGlobalSettingsState {
	if bgs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bgs.Type(), bgs.LocalName()))
	}
	return bgs.state
}

func (bgs *BackupGlobalSettings) DependOn() terra.Reference {
	return terra.ReferenceResource(bgs)
}

type BackupGlobalSettingsArgs struct {
	// GlobalSettings: map of string, required
	GlobalSettings terra.MapValue[terra.StringValue] `hcl:"global_settings,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// DependsOn contains resources that BackupGlobalSettings depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type backupGlobalSettingsAttributes struct {
	ref terra.Reference
}

func (bgs backupGlobalSettingsAttributes) GlobalSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](bgs.ref.Append("global_settings"))
}

func (bgs backupGlobalSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(bgs.ref.Append("id"))
}

type backupGlobalSettingsState struct {
	GlobalSettings map[string]string `json:"global_settings"`
	Id             string            `json:"id"`
}
