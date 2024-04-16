// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_connect_lambda_function_association

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

// Resource represents the Terraform resource aws_connect_lambda_function_association.
type Resource struct {
	Name      string
	Args      Args
	state     *awsConnectLambdaFunctionAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aclfa *Resource) Type() string {
	return "aws_connect_lambda_function_association"
}

// LocalName returns the local name for [Resource].
func (aclfa *Resource) LocalName() string {
	return aclfa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aclfa *Resource) Configuration() interface{} {
	return aclfa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aclfa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aclfa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aclfa *Resource) Dependencies() terra.Dependencies {
	return aclfa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aclfa *Resource) LifecycleManagement() *terra.Lifecycle {
	return aclfa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aclfa *Resource) Attributes() awsConnectLambdaFunctionAssociationAttributes {
	return awsConnectLambdaFunctionAssociationAttributes{ref: terra.ReferenceResource(aclfa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aclfa *Resource) ImportState(state io.Reader) error {
	aclfa.state = &awsConnectLambdaFunctionAssociationState{}
	if err := json.NewDecoder(state).Decode(aclfa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aclfa.Type(), aclfa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aclfa *Resource) State() (*awsConnectLambdaFunctionAssociationState, bool) {
	return aclfa.state, aclfa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aclfa *Resource) StateMust() *awsConnectLambdaFunctionAssociationState {
	if aclfa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aclfa.Type(), aclfa.LocalName()))
	}
	return aclfa.state
}

// Args contains the configurations for aws_connect_lambda_function_association.
type Args struct {
	// FunctionArn: string, required
	FunctionArn terra.StringValue `hcl:"function_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
}

type awsConnectLambdaFunctionAssociationAttributes struct {
	ref terra.Reference
}

// FunctionArn returns a reference to field function_arn of aws_connect_lambda_function_association.
func (aclfa awsConnectLambdaFunctionAssociationAttributes) FunctionArn() terra.StringValue {
	return terra.ReferenceAsString(aclfa.ref.Append("function_arn"))
}

// Id returns a reference to field id of aws_connect_lambda_function_association.
func (aclfa awsConnectLambdaFunctionAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aclfa.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_lambda_function_association.
func (aclfa awsConnectLambdaFunctionAssociationAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(aclfa.ref.Append("instance_id"))
}

type awsConnectLambdaFunctionAssociationState struct {
	FunctionArn string `json:"function_arn"`
	Id          string `json:"id"`
	InstanceId  string `json:"instance_id"`
}
