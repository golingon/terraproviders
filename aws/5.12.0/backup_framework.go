// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	backupframework "github.com/golingon/terraproviders/aws/5.12.0/backupframework"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBackupFramework creates a new instance of [BackupFramework].
func NewBackupFramework(name string, args BackupFrameworkArgs) *BackupFramework {
	return &BackupFramework{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BackupFramework)(nil)

// BackupFramework represents the Terraform resource aws_backup_framework.
type BackupFramework struct {
	Name      string
	Args      BackupFrameworkArgs
	state     *backupFrameworkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BackupFramework].
func (bf *BackupFramework) Type() string {
	return "aws_backup_framework"
}

// LocalName returns the local name for [BackupFramework].
func (bf *BackupFramework) LocalName() string {
	return bf.Name
}

// Configuration returns the configuration (args) for [BackupFramework].
func (bf *BackupFramework) Configuration() interface{} {
	return bf.Args
}

// DependOn is used for other resources to depend on [BackupFramework].
func (bf *BackupFramework) DependOn() terra.Reference {
	return terra.ReferenceResource(bf)
}

// Dependencies returns the list of resources [BackupFramework] depends_on.
func (bf *BackupFramework) Dependencies() terra.Dependencies {
	return bf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BackupFramework].
func (bf *BackupFramework) LifecycleManagement() *terra.Lifecycle {
	return bf.Lifecycle
}

// Attributes returns the attributes for [BackupFramework].
func (bf *BackupFramework) Attributes() backupFrameworkAttributes {
	return backupFrameworkAttributes{ref: terra.ReferenceResource(bf)}
}

// ImportState imports the given attribute values into [BackupFramework]'s state.
func (bf *BackupFramework) ImportState(av io.Reader) error {
	bf.state = &backupFrameworkState{}
	if err := json.NewDecoder(av).Decode(bf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bf.Type(), bf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BackupFramework] has state.
func (bf *BackupFramework) State() (*backupFrameworkState, bool) {
	return bf.state, bf.state != nil
}

// StateMust returns the state for [BackupFramework]. Panics if the state is nil.
func (bf *BackupFramework) StateMust() *backupFrameworkState {
	if bf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bf.Type(), bf.LocalName()))
	}
	return bf.state
}

// BackupFrameworkArgs contains the configurations for aws_backup_framework.
type BackupFrameworkArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Control: min=1
	Control []backupframework.Control `hcl:"control,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *backupframework.Timeouts `hcl:"timeouts,block"`
}
type backupFrameworkAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_backup_framework.
func (bf backupFrameworkAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(bf.ref.Append("arn"))
}

// CreationTime returns a reference to field creation_time of aws_backup_framework.
func (bf backupFrameworkAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(bf.ref.Append("creation_time"))
}

// DeploymentStatus returns a reference to field deployment_status of aws_backup_framework.
func (bf backupFrameworkAttributes) DeploymentStatus() terra.StringValue {
	return terra.ReferenceAsString(bf.ref.Append("deployment_status"))
}

// Description returns a reference to field description of aws_backup_framework.
func (bf backupFrameworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(bf.ref.Append("description"))
}

// Id returns a reference to field id of aws_backup_framework.
func (bf backupFrameworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bf.ref.Append("id"))
}

// Name returns a reference to field name of aws_backup_framework.
func (bf backupFrameworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bf.ref.Append("name"))
}

// Status returns a reference to field status of aws_backup_framework.
func (bf backupFrameworkAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(bf.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_backup_framework.
func (bf backupFrameworkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bf.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_backup_framework.
func (bf backupFrameworkAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bf.ref.Append("tags_all"))
}

func (bf backupFrameworkAttributes) Control() terra.SetValue[backupframework.ControlAttributes] {
	return terra.ReferenceAsSet[backupframework.ControlAttributes](bf.ref.Append("control"))
}

func (bf backupFrameworkAttributes) Timeouts() backupframework.TimeoutsAttributes {
	return terra.ReferenceAsSingle[backupframework.TimeoutsAttributes](bf.ref.Append("timeouts"))
}

type backupFrameworkState struct {
	Arn              string                         `json:"arn"`
	CreationTime     string                         `json:"creation_time"`
	DeploymentStatus string                         `json:"deployment_status"`
	Description      string                         `json:"description"`
	Id               string                         `json:"id"`
	Name             string                         `json:"name"`
	Status           string                         `json:"status"`
	Tags             map[string]string              `json:"tags"`
	TagsAll          map[string]string              `json:"tags_all"`
	Control          []backupframework.ControlState `json:"control"`
	Timeouts         *backupframework.TimeoutsState `json:"timeouts"`
}
