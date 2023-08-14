// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotRoleAlias creates a new instance of [IotRoleAlias].
func NewIotRoleAlias(name string, args IotRoleAliasArgs) *IotRoleAlias {
	return &IotRoleAlias{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotRoleAlias)(nil)

// IotRoleAlias represents the Terraform resource aws_iot_role_alias.
type IotRoleAlias struct {
	Name      string
	Args      IotRoleAliasArgs
	state     *iotRoleAliasState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotRoleAlias].
func (ira *IotRoleAlias) Type() string {
	return "aws_iot_role_alias"
}

// LocalName returns the local name for [IotRoleAlias].
func (ira *IotRoleAlias) LocalName() string {
	return ira.Name
}

// Configuration returns the configuration (args) for [IotRoleAlias].
func (ira *IotRoleAlias) Configuration() interface{} {
	return ira.Args
}

// DependOn is used for other resources to depend on [IotRoleAlias].
func (ira *IotRoleAlias) DependOn() terra.Reference {
	return terra.ReferenceResource(ira)
}

// Dependencies returns the list of resources [IotRoleAlias] depends_on.
func (ira *IotRoleAlias) Dependencies() terra.Dependencies {
	return ira.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotRoleAlias].
func (ira *IotRoleAlias) LifecycleManagement() *terra.Lifecycle {
	return ira.Lifecycle
}

// Attributes returns the attributes for [IotRoleAlias].
func (ira *IotRoleAlias) Attributes() iotRoleAliasAttributes {
	return iotRoleAliasAttributes{ref: terra.ReferenceResource(ira)}
}

// ImportState imports the given attribute values into [IotRoleAlias]'s state.
func (ira *IotRoleAlias) ImportState(av io.Reader) error {
	ira.state = &iotRoleAliasState{}
	if err := json.NewDecoder(av).Decode(ira.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ira.Type(), ira.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotRoleAlias] has state.
func (ira *IotRoleAlias) State() (*iotRoleAliasState, bool) {
	return ira.state, ira.state != nil
}

// StateMust returns the state for [IotRoleAlias]. Panics if the state is nil.
func (ira *IotRoleAlias) StateMust() *iotRoleAliasState {
	if ira.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ira.Type(), ira.LocalName()))
	}
	return ira.state
}

// IotRoleAliasArgs contains the configurations for aws_iot_role_alias.
type IotRoleAliasArgs struct {
	// Alias: string, required
	Alias terra.StringValue `hcl:"alias,attr" validate:"required"`
	// CredentialDuration: number, optional
	CredentialDuration terra.NumberValue `hcl:"credential_duration,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}
type iotRoleAliasAttributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of aws_iot_role_alias.
func (ira iotRoleAliasAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(ira.ref.Append("alias"))
}

// Arn returns a reference to field arn of aws_iot_role_alias.
func (ira iotRoleAliasAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ira.ref.Append("arn"))
}

// CredentialDuration returns a reference to field credential_duration of aws_iot_role_alias.
func (ira iotRoleAliasAttributes) CredentialDuration() terra.NumberValue {
	return terra.ReferenceAsNumber(ira.ref.Append("credential_duration"))
}

// Id returns a reference to field id of aws_iot_role_alias.
func (ira iotRoleAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ira.ref.Append("id"))
}

// RoleArn returns a reference to field role_arn of aws_iot_role_alias.
func (ira iotRoleAliasAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ira.ref.Append("role_arn"))
}

type iotRoleAliasState struct {
	Alias              string  `json:"alias"`
	Arn                string  `json:"arn"`
	CredentialDuration float64 `json:"credential_duration"`
	Id                 string  `json:"id"`
	RoleArn            string  `json:"role_arn"`
}