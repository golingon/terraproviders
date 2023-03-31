// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lambdaalias "github.com/golingon/terraproviders/aws/4.60.0/lambdaalias"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLambdaAlias creates a new instance of [LambdaAlias].
func NewLambdaAlias(name string, args LambdaAliasArgs) *LambdaAlias {
	return &LambdaAlias{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LambdaAlias)(nil)

// LambdaAlias represents the Terraform resource aws_lambda_alias.
type LambdaAlias struct {
	Name      string
	Args      LambdaAliasArgs
	state     *lambdaAliasState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LambdaAlias].
func (la *LambdaAlias) Type() string {
	return "aws_lambda_alias"
}

// LocalName returns the local name for [LambdaAlias].
func (la *LambdaAlias) LocalName() string {
	return la.Name
}

// Configuration returns the configuration (args) for [LambdaAlias].
func (la *LambdaAlias) Configuration() interface{} {
	return la.Args
}

// DependOn is used for other resources to depend on [LambdaAlias].
func (la *LambdaAlias) DependOn() terra.Reference {
	return terra.ReferenceResource(la)
}

// Dependencies returns the list of resources [LambdaAlias] depends_on.
func (la *LambdaAlias) Dependencies() terra.Dependencies {
	return la.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LambdaAlias].
func (la *LambdaAlias) LifecycleManagement() *terra.Lifecycle {
	return la.Lifecycle
}

// Attributes returns the attributes for [LambdaAlias].
func (la *LambdaAlias) Attributes() lambdaAliasAttributes {
	return lambdaAliasAttributes{ref: terra.ReferenceResource(la)}
}

// ImportState imports the given attribute values into [LambdaAlias]'s state.
func (la *LambdaAlias) ImportState(av io.Reader) error {
	la.state = &lambdaAliasState{}
	if err := json.NewDecoder(av).Decode(la.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", la.Type(), la.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LambdaAlias] has state.
func (la *LambdaAlias) State() (*lambdaAliasState, bool) {
	return la.state, la.state != nil
}

// StateMust returns the state for [LambdaAlias]. Panics if the state is nil.
func (la *LambdaAlias) StateMust() *lambdaAliasState {
	if la.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", la.Type(), la.LocalName()))
	}
	return la.state
}

// LambdaAliasArgs contains the configurations for aws_lambda_alias.
type LambdaAliasArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FunctionName: string, required
	FunctionName terra.StringValue `hcl:"function_name,attr" validate:"required"`
	// FunctionVersion: string, required
	FunctionVersion terra.StringValue `hcl:"function_version,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoutingConfig: optional
	RoutingConfig *lambdaalias.RoutingConfig `hcl:"routing_config,block"`
}
type lambdaAliasAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lambda_alias.
func (la lambdaAliasAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("arn"))
}

// Description returns a reference to field description of aws_lambda_alias.
func (la lambdaAliasAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("description"))
}

// FunctionName returns a reference to field function_name of aws_lambda_alias.
func (la lambdaAliasAttributes) FunctionName() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("function_name"))
}

// FunctionVersion returns a reference to field function_version of aws_lambda_alias.
func (la lambdaAliasAttributes) FunctionVersion() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("function_version"))
}

// Id returns a reference to field id of aws_lambda_alias.
func (la lambdaAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("id"))
}

// InvokeArn returns a reference to field invoke_arn of aws_lambda_alias.
func (la lambdaAliasAttributes) InvokeArn() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("invoke_arn"))
}

// Name returns a reference to field name of aws_lambda_alias.
func (la lambdaAliasAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("name"))
}

func (la lambdaAliasAttributes) RoutingConfig() terra.ListValue[lambdaalias.RoutingConfigAttributes] {
	return terra.ReferenceAsList[lambdaalias.RoutingConfigAttributes](la.ref.Append("routing_config"))
}

type lambdaAliasState struct {
	Arn             string                           `json:"arn"`
	Description     string                           `json:"description"`
	FunctionName    string                           `json:"function_name"`
	FunctionVersion string                           `json:"function_version"`
	Id              string                           `json:"id"`
	InvokeArn       string                           `json:"invoke_arn"`
	Name            string                           `json:"name"`
	RoutingConfig   []lambdaalias.RoutingConfigState `json:"routing_config"`
}
