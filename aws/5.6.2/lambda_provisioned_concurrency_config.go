// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lambdaprovisionedconcurrencyconfig "github.com/golingon/terraproviders/aws/5.6.2/lambdaprovisionedconcurrencyconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLambdaProvisionedConcurrencyConfig creates a new instance of [LambdaProvisionedConcurrencyConfig].
func NewLambdaProvisionedConcurrencyConfig(name string, args LambdaProvisionedConcurrencyConfigArgs) *LambdaProvisionedConcurrencyConfig {
	return &LambdaProvisionedConcurrencyConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LambdaProvisionedConcurrencyConfig)(nil)

// LambdaProvisionedConcurrencyConfig represents the Terraform resource aws_lambda_provisioned_concurrency_config.
type LambdaProvisionedConcurrencyConfig struct {
	Name      string
	Args      LambdaProvisionedConcurrencyConfigArgs
	state     *lambdaProvisionedConcurrencyConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LambdaProvisionedConcurrencyConfig].
func (lpcc *LambdaProvisionedConcurrencyConfig) Type() string {
	return "aws_lambda_provisioned_concurrency_config"
}

// LocalName returns the local name for [LambdaProvisionedConcurrencyConfig].
func (lpcc *LambdaProvisionedConcurrencyConfig) LocalName() string {
	return lpcc.Name
}

// Configuration returns the configuration (args) for [LambdaProvisionedConcurrencyConfig].
func (lpcc *LambdaProvisionedConcurrencyConfig) Configuration() interface{} {
	return lpcc.Args
}

// DependOn is used for other resources to depend on [LambdaProvisionedConcurrencyConfig].
func (lpcc *LambdaProvisionedConcurrencyConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(lpcc)
}

// Dependencies returns the list of resources [LambdaProvisionedConcurrencyConfig] depends_on.
func (lpcc *LambdaProvisionedConcurrencyConfig) Dependencies() terra.Dependencies {
	return lpcc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LambdaProvisionedConcurrencyConfig].
func (lpcc *LambdaProvisionedConcurrencyConfig) LifecycleManagement() *terra.Lifecycle {
	return lpcc.Lifecycle
}

// Attributes returns the attributes for [LambdaProvisionedConcurrencyConfig].
func (lpcc *LambdaProvisionedConcurrencyConfig) Attributes() lambdaProvisionedConcurrencyConfigAttributes {
	return lambdaProvisionedConcurrencyConfigAttributes{ref: terra.ReferenceResource(lpcc)}
}

// ImportState imports the given attribute values into [LambdaProvisionedConcurrencyConfig]'s state.
func (lpcc *LambdaProvisionedConcurrencyConfig) ImportState(av io.Reader) error {
	lpcc.state = &lambdaProvisionedConcurrencyConfigState{}
	if err := json.NewDecoder(av).Decode(lpcc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lpcc.Type(), lpcc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LambdaProvisionedConcurrencyConfig] has state.
func (lpcc *LambdaProvisionedConcurrencyConfig) State() (*lambdaProvisionedConcurrencyConfigState, bool) {
	return lpcc.state, lpcc.state != nil
}

// StateMust returns the state for [LambdaProvisionedConcurrencyConfig]. Panics if the state is nil.
func (lpcc *LambdaProvisionedConcurrencyConfig) StateMust() *lambdaProvisionedConcurrencyConfigState {
	if lpcc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lpcc.Type(), lpcc.LocalName()))
	}
	return lpcc.state
}

// LambdaProvisionedConcurrencyConfigArgs contains the configurations for aws_lambda_provisioned_concurrency_config.
type LambdaProvisionedConcurrencyConfigArgs struct {
	// FunctionName: string, required
	FunctionName terra.StringValue `hcl:"function_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProvisionedConcurrentExecutions: number, required
	ProvisionedConcurrentExecutions terra.NumberValue `hcl:"provisioned_concurrent_executions,attr" validate:"required"`
	// Qualifier: string, required
	Qualifier terra.StringValue `hcl:"qualifier,attr" validate:"required"`
	// SkipDestroy: bool, optional
	SkipDestroy terra.BoolValue `hcl:"skip_destroy,attr"`
	// Timeouts: optional
	Timeouts *lambdaprovisionedconcurrencyconfig.Timeouts `hcl:"timeouts,block"`
}
type lambdaProvisionedConcurrencyConfigAttributes struct {
	ref terra.Reference
}

// FunctionName returns a reference to field function_name of aws_lambda_provisioned_concurrency_config.
func (lpcc lambdaProvisionedConcurrencyConfigAttributes) FunctionName() terra.StringValue {
	return terra.ReferenceAsString(lpcc.ref.Append("function_name"))
}

// Id returns a reference to field id of aws_lambda_provisioned_concurrency_config.
func (lpcc lambdaProvisionedConcurrencyConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lpcc.ref.Append("id"))
}

// ProvisionedConcurrentExecutions returns a reference to field provisioned_concurrent_executions of aws_lambda_provisioned_concurrency_config.
func (lpcc lambdaProvisionedConcurrencyConfigAttributes) ProvisionedConcurrentExecutions() terra.NumberValue {
	return terra.ReferenceAsNumber(lpcc.ref.Append("provisioned_concurrent_executions"))
}

// Qualifier returns a reference to field qualifier of aws_lambda_provisioned_concurrency_config.
func (lpcc lambdaProvisionedConcurrencyConfigAttributes) Qualifier() terra.StringValue {
	return terra.ReferenceAsString(lpcc.ref.Append("qualifier"))
}

// SkipDestroy returns a reference to field skip_destroy of aws_lambda_provisioned_concurrency_config.
func (lpcc lambdaProvisionedConcurrencyConfigAttributes) SkipDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(lpcc.ref.Append("skip_destroy"))
}

func (lpcc lambdaProvisionedConcurrencyConfigAttributes) Timeouts() lambdaprovisionedconcurrencyconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lambdaprovisionedconcurrencyconfig.TimeoutsAttributes](lpcc.ref.Append("timeouts"))
}

type lambdaProvisionedConcurrencyConfigState struct {
	FunctionName                    string                                            `json:"function_name"`
	Id                              string                                            `json:"id"`
	ProvisionedConcurrentExecutions float64                                           `json:"provisioned_concurrent_executions"`
	Qualifier                       string                                            `json:"qualifier"`
	SkipDestroy                     bool                                              `json:"skip_destroy"`
	Timeouts                        *lambdaprovisionedconcurrencyconfig.TimeoutsState `json:"timeouts"`
}
