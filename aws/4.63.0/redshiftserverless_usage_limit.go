// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedshiftserverlessUsageLimit creates a new instance of [RedshiftserverlessUsageLimit].
func NewRedshiftserverlessUsageLimit(name string, args RedshiftserverlessUsageLimitArgs) *RedshiftserverlessUsageLimit {
	return &RedshiftserverlessUsageLimit{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftserverlessUsageLimit)(nil)

// RedshiftserverlessUsageLimit represents the Terraform resource aws_redshiftserverless_usage_limit.
type RedshiftserverlessUsageLimit struct {
	Name      string
	Args      RedshiftserverlessUsageLimitArgs
	state     *redshiftserverlessUsageLimitState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedshiftserverlessUsageLimit].
func (rul *RedshiftserverlessUsageLimit) Type() string {
	return "aws_redshiftserverless_usage_limit"
}

// LocalName returns the local name for [RedshiftserverlessUsageLimit].
func (rul *RedshiftserverlessUsageLimit) LocalName() string {
	return rul.Name
}

// Configuration returns the configuration (args) for [RedshiftserverlessUsageLimit].
func (rul *RedshiftserverlessUsageLimit) Configuration() interface{} {
	return rul.Args
}

// DependOn is used for other resources to depend on [RedshiftserverlessUsageLimit].
func (rul *RedshiftserverlessUsageLimit) DependOn() terra.Reference {
	return terra.ReferenceResource(rul)
}

// Dependencies returns the list of resources [RedshiftserverlessUsageLimit] depends_on.
func (rul *RedshiftserverlessUsageLimit) Dependencies() terra.Dependencies {
	return rul.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedshiftserverlessUsageLimit].
func (rul *RedshiftserverlessUsageLimit) LifecycleManagement() *terra.Lifecycle {
	return rul.Lifecycle
}

// Attributes returns the attributes for [RedshiftserverlessUsageLimit].
func (rul *RedshiftserverlessUsageLimit) Attributes() redshiftserverlessUsageLimitAttributes {
	return redshiftserverlessUsageLimitAttributes{ref: terra.ReferenceResource(rul)}
}

// ImportState imports the given attribute values into [RedshiftserverlessUsageLimit]'s state.
func (rul *RedshiftserverlessUsageLimit) ImportState(av io.Reader) error {
	rul.state = &redshiftserverlessUsageLimitState{}
	if err := json.NewDecoder(av).Decode(rul.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rul.Type(), rul.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedshiftserverlessUsageLimit] has state.
func (rul *RedshiftserverlessUsageLimit) State() (*redshiftserverlessUsageLimitState, bool) {
	return rul.state, rul.state != nil
}

// StateMust returns the state for [RedshiftserverlessUsageLimit]. Panics if the state is nil.
func (rul *RedshiftserverlessUsageLimit) StateMust() *redshiftserverlessUsageLimitState {
	if rul.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rul.Type(), rul.LocalName()))
	}
	return rul.state
}

// RedshiftserverlessUsageLimitArgs contains the configurations for aws_redshiftserverless_usage_limit.
type RedshiftserverlessUsageLimitArgs struct {
	// Amount: number, required
	Amount terra.NumberValue `hcl:"amount,attr" validate:"required"`
	// BreachAction: string, optional
	BreachAction terra.StringValue `hcl:"breach_action,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Period: string, optional
	Period terra.StringValue `hcl:"period,attr"`
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
	// UsageType: string, required
	UsageType terra.StringValue `hcl:"usage_type,attr" validate:"required"`
}
type redshiftserverlessUsageLimitAttributes struct {
	ref terra.Reference
}

// Amount returns a reference to field amount of aws_redshiftserverless_usage_limit.
func (rul redshiftserverlessUsageLimitAttributes) Amount() terra.NumberValue {
	return terra.ReferenceAsNumber(rul.ref.Append("amount"))
}

// Arn returns a reference to field arn of aws_redshiftserverless_usage_limit.
func (rul redshiftserverlessUsageLimitAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rul.ref.Append("arn"))
}

// BreachAction returns a reference to field breach_action of aws_redshiftserverless_usage_limit.
func (rul redshiftserverlessUsageLimitAttributes) BreachAction() terra.StringValue {
	return terra.ReferenceAsString(rul.ref.Append("breach_action"))
}

// Id returns a reference to field id of aws_redshiftserverless_usage_limit.
func (rul redshiftserverlessUsageLimitAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rul.ref.Append("id"))
}

// Period returns a reference to field period of aws_redshiftserverless_usage_limit.
func (rul redshiftserverlessUsageLimitAttributes) Period() terra.StringValue {
	return terra.ReferenceAsString(rul.ref.Append("period"))
}

// ResourceArn returns a reference to field resource_arn of aws_redshiftserverless_usage_limit.
func (rul redshiftserverlessUsageLimitAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(rul.ref.Append("resource_arn"))
}

// UsageType returns a reference to field usage_type of aws_redshiftserverless_usage_limit.
func (rul redshiftserverlessUsageLimitAttributes) UsageType() terra.StringValue {
	return terra.ReferenceAsString(rul.ref.Append("usage_type"))
}

type redshiftserverlessUsageLimitState struct {
	Amount       float64 `json:"amount"`
	Arn          string  `json:"arn"`
	BreachAction string  `json:"breach_action"`
	Id           string  `json:"id"`
	Period       string  `json:"period"`
	ResourceArn  string  `json:"resource_arn"`
	UsageType    string  `json:"usage_type"`
}
