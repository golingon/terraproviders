// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudwatchLogDataProtectionPolicy creates a new instance of [CloudwatchLogDataProtectionPolicy].
func NewCloudwatchLogDataProtectionPolicy(name string, args CloudwatchLogDataProtectionPolicyArgs) *CloudwatchLogDataProtectionPolicy {
	return &CloudwatchLogDataProtectionPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudwatchLogDataProtectionPolicy)(nil)

// CloudwatchLogDataProtectionPolicy represents the Terraform resource aws_cloudwatch_log_data_protection_policy.
type CloudwatchLogDataProtectionPolicy struct {
	Name      string
	Args      CloudwatchLogDataProtectionPolicyArgs
	state     *cloudwatchLogDataProtectionPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudwatchLogDataProtectionPolicy].
func (cldpp *CloudwatchLogDataProtectionPolicy) Type() string {
	return "aws_cloudwatch_log_data_protection_policy"
}

// LocalName returns the local name for [CloudwatchLogDataProtectionPolicy].
func (cldpp *CloudwatchLogDataProtectionPolicy) LocalName() string {
	return cldpp.Name
}

// Configuration returns the configuration (args) for [CloudwatchLogDataProtectionPolicy].
func (cldpp *CloudwatchLogDataProtectionPolicy) Configuration() interface{} {
	return cldpp.Args
}

// DependOn is used for other resources to depend on [CloudwatchLogDataProtectionPolicy].
func (cldpp *CloudwatchLogDataProtectionPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(cldpp)
}

// Dependencies returns the list of resources [CloudwatchLogDataProtectionPolicy] depends_on.
func (cldpp *CloudwatchLogDataProtectionPolicy) Dependencies() terra.Dependencies {
	return cldpp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudwatchLogDataProtectionPolicy].
func (cldpp *CloudwatchLogDataProtectionPolicy) LifecycleManagement() *terra.Lifecycle {
	return cldpp.Lifecycle
}

// Attributes returns the attributes for [CloudwatchLogDataProtectionPolicy].
func (cldpp *CloudwatchLogDataProtectionPolicy) Attributes() cloudwatchLogDataProtectionPolicyAttributes {
	return cloudwatchLogDataProtectionPolicyAttributes{ref: terra.ReferenceResource(cldpp)}
}

// ImportState imports the given attribute values into [CloudwatchLogDataProtectionPolicy]'s state.
func (cldpp *CloudwatchLogDataProtectionPolicy) ImportState(av io.Reader) error {
	cldpp.state = &cloudwatchLogDataProtectionPolicyState{}
	if err := json.NewDecoder(av).Decode(cldpp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cldpp.Type(), cldpp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudwatchLogDataProtectionPolicy] has state.
func (cldpp *CloudwatchLogDataProtectionPolicy) State() (*cloudwatchLogDataProtectionPolicyState, bool) {
	return cldpp.state, cldpp.state != nil
}

// StateMust returns the state for [CloudwatchLogDataProtectionPolicy]. Panics if the state is nil.
func (cldpp *CloudwatchLogDataProtectionPolicy) StateMust() *cloudwatchLogDataProtectionPolicyState {
	if cldpp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cldpp.Type(), cldpp.LocalName()))
	}
	return cldpp.state
}

// CloudwatchLogDataProtectionPolicyArgs contains the configurations for aws_cloudwatch_log_data_protection_policy.
type CloudwatchLogDataProtectionPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogGroupName: string, required
	LogGroupName terra.StringValue `hcl:"log_group_name,attr" validate:"required"`
	// PolicyDocument: string, required
	PolicyDocument terra.StringValue `hcl:"policy_document,attr" validate:"required"`
}
type cloudwatchLogDataProtectionPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_cloudwatch_log_data_protection_policy.
func (cldpp cloudwatchLogDataProtectionPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cldpp.ref.Append("id"))
}

// LogGroupName returns a reference to field log_group_name of aws_cloudwatch_log_data_protection_policy.
func (cldpp cloudwatchLogDataProtectionPolicyAttributes) LogGroupName() terra.StringValue {
	return terra.ReferenceAsString(cldpp.ref.Append("log_group_name"))
}

// PolicyDocument returns a reference to field policy_document of aws_cloudwatch_log_data_protection_policy.
func (cldpp cloudwatchLogDataProtectionPolicyAttributes) PolicyDocument() terra.StringValue {
	return terra.ReferenceAsString(cldpp.ref.Append("policy_document"))
}

type cloudwatchLogDataProtectionPolicyState struct {
	Id             string `json:"id"`
	LogGroupName   string `json:"log_group_name"`
	PolicyDocument string `json:"policy_document"`
}
