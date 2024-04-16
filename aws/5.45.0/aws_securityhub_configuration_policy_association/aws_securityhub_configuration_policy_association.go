// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_securityhub_configuration_policy_association

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

// Resource represents the Terraform resource aws_securityhub_configuration_policy_association.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSecurityhubConfigurationPolicyAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ascpa *Resource) Type() string {
	return "aws_securityhub_configuration_policy_association"
}

// LocalName returns the local name for [Resource].
func (ascpa *Resource) LocalName() string {
	return ascpa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ascpa *Resource) Configuration() interface{} {
	return ascpa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ascpa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ascpa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ascpa *Resource) Dependencies() terra.Dependencies {
	return ascpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ascpa *Resource) LifecycleManagement() *terra.Lifecycle {
	return ascpa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ascpa *Resource) Attributes() awsSecurityhubConfigurationPolicyAssociationAttributes {
	return awsSecurityhubConfigurationPolicyAssociationAttributes{ref: terra.ReferenceResource(ascpa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ascpa *Resource) ImportState(state io.Reader) error {
	ascpa.state = &awsSecurityhubConfigurationPolicyAssociationState{}
	if err := json.NewDecoder(state).Decode(ascpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ascpa.Type(), ascpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ascpa *Resource) State() (*awsSecurityhubConfigurationPolicyAssociationState, bool) {
	return ascpa.state, ascpa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ascpa *Resource) StateMust() *awsSecurityhubConfigurationPolicyAssociationState {
	if ascpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ascpa.Type(), ascpa.LocalName()))
	}
	return ascpa.state
}

// Args contains the configurations for aws_securityhub_configuration_policy_association.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyId: string, required
	PolicyId terra.StringValue `hcl:"policy_id,attr" validate:"required"`
	// TargetId: string, required
	TargetId terra.StringValue `hcl:"target_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsSecurityhubConfigurationPolicyAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_securityhub_configuration_policy_association.
func (ascpa awsSecurityhubConfigurationPolicyAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ascpa.ref.Append("id"))
}

// PolicyId returns a reference to field policy_id of aws_securityhub_configuration_policy_association.
func (ascpa awsSecurityhubConfigurationPolicyAssociationAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(ascpa.ref.Append("policy_id"))
}

// TargetId returns a reference to field target_id of aws_securityhub_configuration_policy_association.
func (ascpa awsSecurityhubConfigurationPolicyAssociationAttributes) TargetId() terra.StringValue {
	return terra.ReferenceAsString(ascpa.ref.Append("target_id"))
}

func (ascpa awsSecurityhubConfigurationPolicyAssociationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ascpa.ref.Append("timeouts"))
}

type awsSecurityhubConfigurationPolicyAssociationState struct {
	Id       string         `json:"id"`
	PolicyId string         `json:"policy_id"`
	TargetId string         `json:"target_id"`
	Timeouts *TimeoutsState `json:"timeouts"`
}
