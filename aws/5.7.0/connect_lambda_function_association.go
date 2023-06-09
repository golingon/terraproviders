// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConnectLambdaFunctionAssociation creates a new instance of [ConnectLambdaFunctionAssociation].
func NewConnectLambdaFunctionAssociation(name string, args ConnectLambdaFunctionAssociationArgs) *ConnectLambdaFunctionAssociation {
	return &ConnectLambdaFunctionAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConnectLambdaFunctionAssociation)(nil)

// ConnectLambdaFunctionAssociation represents the Terraform resource aws_connect_lambda_function_association.
type ConnectLambdaFunctionAssociation struct {
	Name      string
	Args      ConnectLambdaFunctionAssociationArgs
	state     *connectLambdaFunctionAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConnectLambdaFunctionAssociation].
func (clfa *ConnectLambdaFunctionAssociation) Type() string {
	return "aws_connect_lambda_function_association"
}

// LocalName returns the local name for [ConnectLambdaFunctionAssociation].
func (clfa *ConnectLambdaFunctionAssociation) LocalName() string {
	return clfa.Name
}

// Configuration returns the configuration (args) for [ConnectLambdaFunctionAssociation].
func (clfa *ConnectLambdaFunctionAssociation) Configuration() interface{} {
	return clfa.Args
}

// DependOn is used for other resources to depend on [ConnectLambdaFunctionAssociation].
func (clfa *ConnectLambdaFunctionAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(clfa)
}

// Dependencies returns the list of resources [ConnectLambdaFunctionAssociation] depends_on.
func (clfa *ConnectLambdaFunctionAssociation) Dependencies() terra.Dependencies {
	return clfa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConnectLambdaFunctionAssociation].
func (clfa *ConnectLambdaFunctionAssociation) LifecycleManagement() *terra.Lifecycle {
	return clfa.Lifecycle
}

// Attributes returns the attributes for [ConnectLambdaFunctionAssociation].
func (clfa *ConnectLambdaFunctionAssociation) Attributes() connectLambdaFunctionAssociationAttributes {
	return connectLambdaFunctionAssociationAttributes{ref: terra.ReferenceResource(clfa)}
}

// ImportState imports the given attribute values into [ConnectLambdaFunctionAssociation]'s state.
func (clfa *ConnectLambdaFunctionAssociation) ImportState(av io.Reader) error {
	clfa.state = &connectLambdaFunctionAssociationState{}
	if err := json.NewDecoder(av).Decode(clfa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", clfa.Type(), clfa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConnectLambdaFunctionAssociation] has state.
func (clfa *ConnectLambdaFunctionAssociation) State() (*connectLambdaFunctionAssociationState, bool) {
	return clfa.state, clfa.state != nil
}

// StateMust returns the state for [ConnectLambdaFunctionAssociation]. Panics if the state is nil.
func (clfa *ConnectLambdaFunctionAssociation) StateMust() *connectLambdaFunctionAssociationState {
	if clfa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", clfa.Type(), clfa.LocalName()))
	}
	return clfa.state
}

// ConnectLambdaFunctionAssociationArgs contains the configurations for aws_connect_lambda_function_association.
type ConnectLambdaFunctionAssociationArgs struct {
	// FunctionArn: string, required
	FunctionArn terra.StringValue `hcl:"function_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
}
type connectLambdaFunctionAssociationAttributes struct {
	ref terra.Reference
}

// FunctionArn returns a reference to field function_arn of aws_connect_lambda_function_association.
func (clfa connectLambdaFunctionAssociationAttributes) FunctionArn() terra.StringValue {
	return terra.ReferenceAsString(clfa.ref.Append("function_arn"))
}

// Id returns a reference to field id of aws_connect_lambda_function_association.
func (clfa connectLambdaFunctionAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(clfa.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_lambda_function_association.
func (clfa connectLambdaFunctionAssociationAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(clfa.ref.Append("instance_id"))
}

type connectLambdaFunctionAssociationState struct {
	FunctionArn string `json:"function_arn"`
	Id          string `json:"id"`
	InstanceId  string `json:"instance_id"`
}
