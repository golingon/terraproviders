// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ssoadmin_permission_set_inline_policy

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

// Resource represents the Terraform resource aws_ssoadmin_permission_set_inline_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSsoadminPermissionSetInlinePolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aspsip *Resource) Type() string {
	return "aws_ssoadmin_permission_set_inline_policy"
}

// LocalName returns the local name for [Resource].
func (aspsip *Resource) LocalName() string {
	return aspsip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aspsip *Resource) Configuration() interface{} {
	return aspsip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aspsip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aspsip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aspsip *Resource) Dependencies() terra.Dependencies {
	return aspsip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aspsip *Resource) LifecycleManagement() *terra.Lifecycle {
	return aspsip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aspsip *Resource) Attributes() awsSsoadminPermissionSetInlinePolicyAttributes {
	return awsSsoadminPermissionSetInlinePolicyAttributes{ref: terra.ReferenceResource(aspsip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aspsip *Resource) ImportState(state io.Reader) error {
	aspsip.state = &awsSsoadminPermissionSetInlinePolicyState{}
	if err := json.NewDecoder(state).Decode(aspsip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aspsip.Type(), aspsip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aspsip *Resource) State() (*awsSsoadminPermissionSetInlinePolicyState, bool) {
	return aspsip.state, aspsip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aspsip *Resource) StateMust() *awsSsoadminPermissionSetInlinePolicyState {
	if aspsip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aspsip.Type(), aspsip.LocalName()))
	}
	return aspsip.state
}

// Args contains the configurations for aws_ssoadmin_permission_set_inline_policy.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InlinePolicy: string, required
	InlinePolicy terra.StringValue `hcl:"inline_policy,attr" validate:"required"`
	// InstanceArn: string, required
	InstanceArn terra.StringValue `hcl:"instance_arn,attr" validate:"required"`
	// PermissionSetArn: string, required
	PermissionSetArn terra.StringValue `hcl:"permission_set_arn,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsSsoadminPermissionSetInlinePolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ssoadmin_permission_set_inline_policy.
func (aspsip awsSsoadminPermissionSetInlinePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aspsip.ref.Append("id"))
}

// InlinePolicy returns a reference to field inline_policy of aws_ssoadmin_permission_set_inline_policy.
func (aspsip awsSsoadminPermissionSetInlinePolicyAttributes) InlinePolicy() terra.StringValue {
	return terra.ReferenceAsString(aspsip.ref.Append("inline_policy"))
}

// InstanceArn returns a reference to field instance_arn of aws_ssoadmin_permission_set_inline_policy.
func (aspsip awsSsoadminPermissionSetInlinePolicyAttributes) InstanceArn() terra.StringValue {
	return terra.ReferenceAsString(aspsip.ref.Append("instance_arn"))
}

// PermissionSetArn returns a reference to field permission_set_arn of aws_ssoadmin_permission_set_inline_policy.
func (aspsip awsSsoadminPermissionSetInlinePolicyAttributes) PermissionSetArn() terra.StringValue {
	return terra.ReferenceAsString(aspsip.ref.Append("permission_set_arn"))
}

func (aspsip awsSsoadminPermissionSetInlinePolicyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aspsip.ref.Append("timeouts"))
}

type awsSsoadminPermissionSetInlinePolicyState struct {
	Id               string         `json:"id"`
	InlinePolicy     string         `json:"inline_policy"`
	InstanceArn      string         `json:"instance_arn"`
	PermissionSetArn string         `json:"permission_set_arn"`
	Timeouts         *TimeoutsState `json:"timeouts"`
}
