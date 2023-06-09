// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEcsAccountSettingDefault creates a new instance of [EcsAccountSettingDefault].
func NewEcsAccountSettingDefault(name string, args EcsAccountSettingDefaultArgs) *EcsAccountSettingDefault {
	return &EcsAccountSettingDefault{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EcsAccountSettingDefault)(nil)

// EcsAccountSettingDefault represents the Terraform resource aws_ecs_account_setting_default.
type EcsAccountSettingDefault struct {
	Name      string
	Args      EcsAccountSettingDefaultArgs
	state     *ecsAccountSettingDefaultState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EcsAccountSettingDefault].
func (easd *EcsAccountSettingDefault) Type() string {
	return "aws_ecs_account_setting_default"
}

// LocalName returns the local name for [EcsAccountSettingDefault].
func (easd *EcsAccountSettingDefault) LocalName() string {
	return easd.Name
}

// Configuration returns the configuration (args) for [EcsAccountSettingDefault].
func (easd *EcsAccountSettingDefault) Configuration() interface{} {
	return easd.Args
}

// DependOn is used for other resources to depend on [EcsAccountSettingDefault].
func (easd *EcsAccountSettingDefault) DependOn() terra.Reference {
	return terra.ReferenceResource(easd)
}

// Dependencies returns the list of resources [EcsAccountSettingDefault] depends_on.
func (easd *EcsAccountSettingDefault) Dependencies() terra.Dependencies {
	return easd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EcsAccountSettingDefault].
func (easd *EcsAccountSettingDefault) LifecycleManagement() *terra.Lifecycle {
	return easd.Lifecycle
}

// Attributes returns the attributes for [EcsAccountSettingDefault].
func (easd *EcsAccountSettingDefault) Attributes() ecsAccountSettingDefaultAttributes {
	return ecsAccountSettingDefaultAttributes{ref: terra.ReferenceResource(easd)}
}

// ImportState imports the given attribute values into [EcsAccountSettingDefault]'s state.
func (easd *EcsAccountSettingDefault) ImportState(av io.Reader) error {
	easd.state = &ecsAccountSettingDefaultState{}
	if err := json.NewDecoder(av).Decode(easd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", easd.Type(), easd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EcsAccountSettingDefault] has state.
func (easd *EcsAccountSettingDefault) State() (*ecsAccountSettingDefaultState, bool) {
	return easd.state, easd.state != nil
}

// StateMust returns the state for [EcsAccountSettingDefault]. Panics if the state is nil.
func (easd *EcsAccountSettingDefault) StateMust() *ecsAccountSettingDefaultState {
	if easd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", easd.Type(), easd.LocalName()))
	}
	return easd.state
}

// EcsAccountSettingDefaultArgs contains the configurations for aws_ecs_account_setting_default.
type EcsAccountSettingDefaultArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}
type ecsAccountSettingDefaultAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ecs_account_setting_default.
func (easd ecsAccountSettingDefaultAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(easd.ref.Append("id"))
}

// Name returns a reference to field name of aws_ecs_account_setting_default.
func (easd ecsAccountSettingDefaultAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(easd.ref.Append("name"))
}

// PrincipalArn returns a reference to field principal_arn of aws_ecs_account_setting_default.
func (easd ecsAccountSettingDefaultAttributes) PrincipalArn() terra.StringValue {
	return terra.ReferenceAsString(easd.ref.Append("principal_arn"))
}

// Value returns a reference to field value of aws_ecs_account_setting_default.
func (easd ecsAccountSettingDefaultAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(easd.ref.Append("value"))
}

type ecsAccountSettingDefaultState struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	PrincipalArn string `json:"principal_arn"`
	Value        string `json:"value"`
}
