// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lambdafunctionurl "github.com/golingon/terraproviders/aws/5.0.1/lambdafunctionurl"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLambdaFunctionUrl creates a new instance of [LambdaFunctionUrl].
func NewLambdaFunctionUrl(name string, args LambdaFunctionUrlArgs) *LambdaFunctionUrl {
	return &LambdaFunctionUrl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LambdaFunctionUrl)(nil)

// LambdaFunctionUrl represents the Terraform resource aws_lambda_function_url.
type LambdaFunctionUrl struct {
	Name      string
	Args      LambdaFunctionUrlArgs
	state     *lambdaFunctionUrlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LambdaFunctionUrl].
func (lfu *LambdaFunctionUrl) Type() string {
	return "aws_lambda_function_url"
}

// LocalName returns the local name for [LambdaFunctionUrl].
func (lfu *LambdaFunctionUrl) LocalName() string {
	return lfu.Name
}

// Configuration returns the configuration (args) for [LambdaFunctionUrl].
func (lfu *LambdaFunctionUrl) Configuration() interface{} {
	return lfu.Args
}

// DependOn is used for other resources to depend on [LambdaFunctionUrl].
func (lfu *LambdaFunctionUrl) DependOn() terra.Reference {
	return terra.ReferenceResource(lfu)
}

// Dependencies returns the list of resources [LambdaFunctionUrl] depends_on.
func (lfu *LambdaFunctionUrl) Dependencies() terra.Dependencies {
	return lfu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LambdaFunctionUrl].
func (lfu *LambdaFunctionUrl) LifecycleManagement() *terra.Lifecycle {
	return lfu.Lifecycle
}

// Attributes returns the attributes for [LambdaFunctionUrl].
func (lfu *LambdaFunctionUrl) Attributes() lambdaFunctionUrlAttributes {
	return lambdaFunctionUrlAttributes{ref: terra.ReferenceResource(lfu)}
}

// ImportState imports the given attribute values into [LambdaFunctionUrl]'s state.
func (lfu *LambdaFunctionUrl) ImportState(av io.Reader) error {
	lfu.state = &lambdaFunctionUrlState{}
	if err := json.NewDecoder(av).Decode(lfu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lfu.Type(), lfu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LambdaFunctionUrl] has state.
func (lfu *LambdaFunctionUrl) State() (*lambdaFunctionUrlState, bool) {
	return lfu.state, lfu.state != nil
}

// StateMust returns the state for [LambdaFunctionUrl]. Panics if the state is nil.
func (lfu *LambdaFunctionUrl) StateMust() *lambdaFunctionUrlState {
	if lfu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lfu.Type(), lfu.LocalName()))
	}
	return lfu.state
}

// LambdaFunctionUrlArgs contains the configurations for aws_lambda_function_url.
type LambdaFunctionUrlArgs struct {
	// AuthorizationType: string, required
	AuthorizationType terra.StringValue `hcl:"authorization_type,attr" validate:"required"`
	// FunctionName: string, required
	FunctionName terra.StringValue `hcl:"function_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InvokeMode: string, optional
	InvokeMode terra.StringValue `hcl:"invoke_mode,attr"`
	// Qualifier: string, optional
	Qualifier terra.StringValue `hcl:"qualifier,attr"`
	// Cors: optional
	Cors *lambdafunctionurl.Cors `hcl:"cors,block"`
	// Timeouts: optional
	Timeouts *lambdafunctionurl.Timeouts `hcl:"timeouts,block"`
}
type lambdaFunctionUrlAttributes struct {
	ref terra.Reference
}

// AuthorizationType returns a reference to field authorization_type of aws_lambda_function_url.
func (lfu lambdaFunctionUrlAttributes) AuthorizationType() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("authorization_type"))
}

// FunctionArn returns a reference to field function_arn of aws_lambda_function_url.
func (lfu lambdaFunctionUrlAttributes) FunctionArn() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("function_arn"))
}

// FunctionName returns a reference to field function_name of aws_lambda_function_url.
func (lfu lambdaFunctionUrlAttributes) FunctionName() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("function_name"))
}

// FunctionUrl returns a reference to field function_url of aws_lambda_function_url.
func (lfu lambdaFunctionUrlAttributes) FunctionUrl() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("function_url"))
}

// Id returns a reference to field id of aws_lambda_function_url.
func (lfu lambdaFunctionUrlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("id"))
}

// InvokeMode returns a reference to field invoke_mode of aws_lambda_function_url.
func (lfu lambdaFunctionUrlAttributes) InvokeMode() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("invoke_mode"))
}

// Qualifier returns a reference to field qualifier of aws_lambda_function_url.
func (lfu lambdaFunctionUrlAttributes) Qualifier() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("qualifier"))
}

// UrlId returns a reference to field url_id of aws_lambda_function_url.
func (lfu lambdaFunctionUrlAttributes) UrlId() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("url_id"))
}

func (lfu lambdaFunctionUrlAttributes) Cors() terra.ListValue[lambdafunctionurl.CorsAttributes] {
	return terra.ReferenceAsList[lambdafunctionurl.CorsAttributes](lfu.ref.Append("cors"))
}

func (lfu lambdaFunctionUrlAttributes) Timeouts() lambdafunctionurl.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lambdafunctionurl.TimeoutsAttributes](lfu.ref.Append("timeouts"))
}

type lambdaFunctionUrlState struct {
	AuthorizationType string                           `json:"authorization_type"`
	FunctionArn       string                           `json:"function_arn"`
	FunctionName      string                           `json:"function_name"`
	FunctionUrl       string                           `json:"function_url"`
	Id                string                           `json:"id"`
	InvokeMode        string                           `json:"invoke_mode"`
	Qualifier         string                           `json:"qualifier"`
	UrlId             string                           `json:"url_id"`
	Cors              []lambdafunctionurl.CorsState    `json:"cors"`
	Timeouts          *lambdafunctionurl.TimeoutsState `json:"timeouts"`
}
