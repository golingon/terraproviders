// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_codebuild_resource_policy

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

// Resource represents the Terraform resource aws_codebuild_resource_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *awsCodebuildResourcePolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acrp *Resource) Type() string {
	return "aws_codebuild_resource_policy"
}

// LocalName returns the local name for [Resource].
func (acrp *Resource) LocalName() string {
	return acrp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acrp *Resource) Configuration() interface{} {
	return acrp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acrp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acrp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acrp *Resource) Dependencies() terra.Dependencies {
	return acrp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acrp *Resource) LifecycleManagement() *terra.Lifecycle {
	return acrp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acrp *Resource) Attributes() awsCodebuildResourcePolicyAttributes {
	return awsCodebuildResourcePolicyAttributes{ref: terra.ReferenceResource(acrp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acrp *Resource) ImportState(state io.Reader) error {
	acrp.state = &awsCodebuildResourcePolicyState{}
	if err := json.NewDecoder(state).Decode(acrp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acrp.Type(), acrp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acrp *Resource) State() (*awsCodebuildResourcePolicyState, bool) {
	return acrp.state, acrp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acrp *Resource) StateMust() *awsCodebuildResourcePolicyState {
	if acrp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acrp.Type(), acrp.LocalName()))
	}
	return acrp.state
}

// Args contains the configurations for aws_codebuild_resource_policy.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
}

type awsCodebuildResourcePolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_codebuild_resource_policy.
func (acrp awsCodebuildResourcePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acrp.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_codebuild_resource_policy.
func (acrp awsCodebuildResourcePolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(acrp.ref.Append("policy"))
}

// ResourceArn returns a reference to field resource_arn of aws_codebuild_resource_policy.
func (acrp awsCodebuildResourcePolicyAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(acrp.ref.Append("resource_arn"))
}

type awsCodebuildResourcePolicyState struct {
	Id          string `json:"id"`
	Policy      string `json:"policy"`
	ResourceArn string `json:"resource_arn"`
}
