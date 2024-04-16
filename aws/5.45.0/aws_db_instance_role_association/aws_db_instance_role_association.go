// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_db_instance_role_association

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_db_instance_role_association.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDbInstanceRoleAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adira *Resource) Type() string {
	return "aws_db_instance_role_association"
}

// LocalName returns the local name for [Resource].
func (adira *Resource) LocalName() string {
	return adira.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adira *Resource) Configuration() interface{} {
	return adira.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adira *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adira)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adira *Resource) Dependencies() terra.Dependencies {
	return adira.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adira *Resource) LifecycleManagement() *terra.Lifecycle {
	return adira.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adira *Resource) Attributes() awsDbInstanceRoleAssociationAttributes {
	return awsDbInstanceRoleAssociationAttributes{ref: terra.ReferenceResource(adira)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adira *Resource) ImportState(state io.Reader) error {
	adira.state = &awsDbInstanceRoleAssociationState{}
	if err := json.NewDecoder(state).Decode(adira.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adira.Type(), adira.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adira *Resource) State() (*awsDbInstanceRoleAssociationState, bool) {
	return adira.state, adira.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adira *Resource) StateMust() *awsDbInstanceRoleAssociationState {
	if adira.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adira.Type(), adira.LocalName()))
	}
	return adira.state
}

// Args contains the configurations for aws_db_instance_role_association.
type Args struct {
	// DbInstanceIdentifier: string, required
	DbInstanceIdentifier terra.StringValue `hcl:"db_instance_identifier,attr" validate:"required"`
	// FeatureName: string, required
	FeatureName terra.StringValue `hcl:"feature_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}

type awsDbInstanceRoleAssociationAttributes struct {
	ref terra.Reference
}

// DbInstanceIdentifier returns a reference to field db_instance_identifier of aws_db_instance_role_association.
func (adira awsDbInstanceRoleAssociationAttributes) DbInstanceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(adira.ref.Append("db_instance_identifier"))
}

// FeatureName returns a reference to field feature_name of aws_db_instance_role_association.
func (adira awsDbInstanceRoleAssociationAttributes) FeatureName() terra.StringValue {
	return terra.ReferenceAsString(adira.ref.Append("feature_name"))
}

// Id returns a reference to field id of aws_db_instance_role_association.
func (adira awsDbInstanceRoleAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adira.ref.Append("id"))
}

// RoleArn returns a reference to field role_arn of aws_db_instance_role_association.
func (adira awsDbInstanceRoleAssociationAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(adira.ref.Append("role_arn"))
}

type awsDbInstanceRoleAssociationState struct {
	DbInstanceIdentifier string `json:"db_instance_identifier"`
	FeatureName          string `json:"feature_name"`
	Id                   string `json:"id"`
	RoleArn              string `json:"role_arn"`
}
