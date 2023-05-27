// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lambdafunctioneventinvokeconfig "github.com/golingon/terraproviders/aws/5.0.1/lambdafunctioneventinvokeconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLambdaFunctionEventInvokeConfig creates a new instance of [LambdaFunctionEventInvokeConfig].
func NewLambdaFunctionEventInvokeConfig(name string, args LambdaFunctionEventInvokeConfigArgs) *LambdaFunctionEventInvokeConfig {
	return &LambdaFunctionEventInvokeConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LambdaFunctionEventInvokeConfig)(nil)

// LambdaFunctionEventInvokeConfig represents the Terraform resource aws_lambda_function_event_invoke_config.
type LambdaFunctionEventInvokeConfig struct {
	Name      string
	Args      LambdaFunctionEventInvokeConfigArgs
	state     *lambdaFunctionEventInvokeConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LambdaFunctionEventInvokeConfig].
func (lfeic *LambdaFunctionEventInvokeConfig) Type() string {
	return "aws_lambda_function_event_invoke_config"
}

// LocalName returns the local name for [LambdaFunctionEventInvokeConfig].
func (lfeic *LambdaFunctionEventInvokeConfig) LocalName() string {
	return lfeic.Name
}

// Configuration returns the configuration (args) for [LambdaFunctionEventInvokeConfig].
func (lfeic *LambdaFunctionEventInvokeConfig) Configuration() interface{} {
	return lfeic.Args
}

// DependOn is used for other resources to depend on [LambdaFunctionEventInvokeConfig].
func (lfeic *LambdaFunctionEventInvokeConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(lfeic)
}

// Dependencies returns the list of resources [LambdaFunctionEventInvokeConfig] depends_on.
func (lfeic *LambdaFunctionEventInvokeConfig) Dependencies() terra.Dependencies {
	return lfeic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LambdaFunctionEventInvokeConfig].
func (lfeic *LambdaFunctionEventInvokeConfig) LifecycleManagement() *terra.Lifecycle {
	return lfeic.Lifecycle
}

// Attributes returns the attributes for [LambdaFunctionEventInvokeConfig].
func (lfeic *LambdaFunctionEventInvokeConfig) Attributes() lambdaFunctionEventInvokeConfigAttributes {
	return lambdaFunctionEventInvokeConfigAttributes{ref: terra.ReferenceResource(lfeic)}
}

// ImportState imports the given attribute values into [LambdaFunctionEventInvokeConfig]'s state.
func (lfeic *LambdaFunctionEventInvokeConfig) ImportState(av io.Reader) error {
	lfeic.state = &lambdaFunctionEventInvokeConfigState{}
	if err := json.NewDecoder(av).Decode(lfeic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lfeic.Type(), lfeic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LambdaFunctionEventInvokeConfig] has state.
func (lfeic *LambdaFunctionEventInvokeConfig) State() (*lambdaFunctionEventInvokeConfigState, bool) {
	return lfeic.state, lfeic.state != nil
}

// StateMust returns the state for [LambdaFunctionEventInvokeConfig]. Panics if the state is nil.
func (lfeic *LambdaFunctionEventInvokeConfig) StateMust() *lambdaFunctionEventInvokeConfigState {
	if lfeic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lfeic.Type(), lfeic.LocalName()))
	}
	return lfeic.state
}

// LambdaFunctionEventInvokeConfigArgs contains the configurations for aws_lambda_function_event_invoke_config.
type LambdaFunctionEventInvokeConfigArgs struct {
	// FunctionName: string, required
	FunctionName terra.StringValue `hcl:"function_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaximumEventAgeInSeconds: number, optional
	MaximumEventAgeInSeconds terra.NumberValue `hcl:"maximum_event_age_in_seconds,attr"`
	// MaximumRetryAttempts: number, optional
	MaximumRetryAttempts terra.NumberValue `hcl:"maximum_retry_attempts,attr"`
	// Qualifier: string, optional
	Qualifier terra.StringValue `hcl:"qualifier,attr"`
	// DestinationConfig: optional
	DestinationConfig *lambdafunctioneventinvokeconfig.DestinationConfig `hcl:"destination_config,block"`
}
type lambdaFunctionEventInvokeConfigAttributes struct {
	ref terra.Reference
}

// FunctionName returns a reference to field function_name of aws_lambda_function_event_invoke_config.
func (lfeic lambdaFunctionEventInvokeConfigAttributes) FunctionName() terra.StringValue {
	return terra.ReferenceAsString(lfeic.ref.Append("function_name"))
}

// Id returns a reference to field id of aws_lambda_function_event_invoke_config.
func (lfeic lambdaFunctionEventInvokeConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lfeic.ref.Append("id"))
}

// MaximumEventAgeInSeconds returns a reference to field maximum_event_age_in_seconds of aws_lambda_function_event_invoke_config.
func (lfeic lambdaFunctionEventInvokeConfigAttributes) MaximumEventAgeInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(lfeic.ref.Append("maximum_event_age_in_seconds"))
}

// MaximumRetryAttempts returns a reference to field maximum_retry_attempts of aws_lambda_function_event_invoke_config.
func (lfeic lambdaFunctionEventInvokeConfigAttributes) MaximumRetryAttempts() terra.NumberValue {
	return terra.ReferenceAsNumber(lfeic.ref.Append("maximum_retry_attempts"))
}

// Qualifier returns a reference to field qualifier of aws_lambda_function_event_invoke_config.
func (lfeic lambdaFunctionEventInvokeConfigAttributes) Qualifier() terra.StringValue {
	return terra.ReferenceAsString(lfeic.ref.Append("qualifier"))
}

func (lfeic lambdaFunctionEventInvokeConfigAttributes) DestinationConfig() terra.ListValue[lambdafunctioneventinvokeconfig.DestinationConfigAttributes] {
	return terra.ReferenceAsList[lambdafunctioneventinvokeconfig.DestinationConfigAttributes](lfeic.ref.Append("destination_config"))
}

type lambdaFunctionEventInvokeConfigState struct {
	FunctionName             string                                                   `json:"function_name"`
	Id                       string                                                   `json:"id"`
	MaximumEventAgeInSeconds float64                                                  `json:"maximum_event_age_in_seconds"`
	MaximumRetryAttempts     float64                                                  `json:"maximum_retry_attempts"`
	Qualifier                string                                                   `json:"qualifier"`
	DestinationConfig        []lambdafunctioneventinvokeconfig.DestinationConfigState `json:"destination_config"`
}
