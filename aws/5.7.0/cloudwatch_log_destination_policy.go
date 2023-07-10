// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudwatchLogDestinationPolicy creates a new instance of [CloudwatchLogDestinationPolicy].
func NewCloudwatchLogDestinationPolicy(name string, args CloudwatchLogDestinationPolicyArgs) *CloudwatchLogDestinationPolicy {
	return &CloudwatchLogDestinationPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudwatchLogDestinationPolicy)(nil)

// CloudwatchLogDestinationPolicy represents the Terraform resource aws_cloudwatch_log_destination_policy.
type CloudwatchLogDestinationPolicy struct {
	Name      string
	Args      CloudwatchLogDestinationPolicyArgs
	state     *cloudwatchLogDestinationPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudwatchLogDestinationPolicy].
func (cldp *CloudwatchLogDestinationPolicy) Type() string {
	return "aws_cloudwatch_log_destination_policy"
}

// LocalName returns the local name for [CloudwatchLogDestinationPolicy].
func (cldp *CloudwatchLogDestinationPolicy) LocalName() string {
	return cldp.Name
}

// Configuration returns the configuration (args) for [CloudwatchLogDestinationPolicy].
func (cldp *CloudwatchLogDestinationPolicy) Configuration() interface{} {
	return cldp.Args
}

// DependOn is used for other resources to depend on [CloudwatchLogDestinationPolicy].
func (cldp *CloudwatchLogDestinationPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(cldp)
}

// Dependencies returns the list of resources [CloudwatchLogDestinationPolicy] depends_on.
func (cldp *CloudwatchLogDestinationPolicy) Dependencies() terra.Dependencies {
	return cldp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudwatchLogDestinationPolicy].
func (cldp *CloudwatchLogDestinationPolicy) LifecycleManagement() *terra.Lifecycle {
	return cldp.Lifecycle
}

// Attributes returns the attributes for [CloudwatchLogDestinationPolicy].
func (cldp *CloudwatchLogDestinationPolicy) Attributes() cloudwatchLogDestinationPolicyAttributes {
	return cloudwatchLogDestinationPolicyAttributes{ref: terra.ReferenceResource(cldp)}
}

// ImportState imports the given attribute values into [CloudwatchLogDestinationPolicy]'s state.
func (cldp *CloudwatchLogDestinationPolicy) ImportState(av io.Reader) error {
	cldp.state = &cloudwatchLogDestinationPolicyState{}
	if err := json.NewDecoder(av).Decode(cldp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cldp.Type(), cldp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudwatchLogDestinationPolicy] has state.
func (cldp *CloudwatchLogDestinationPolicy) State() (*cloudwatchLogDestinationPolicyState, bool) {
	return cldp.state, cldp.state != nil
}

// StateMust returns the state for [CloudwatchLogDestinationPolicy]. Panics if the state is nil.
func (cldp *CloudwatchLogDestinationPolicy) StateMust() *cloudwatchLogDestinationPolicyState {
	if cldp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cldp.Type(), cldp.LocalName()))
	}
	return cldp.state
}

// CloudwatchLogDestinationPolicyArgs contains the configurations for aws_cloudwatch_log_destination_policy.
type CloudwatchLogDestinationPolicyArgs struct {
	// AccessPolicy: string, required
	AccessPolicy terra.StringValue `hcl:"access_policy,attr" validate:"required"`
	// DestinationName: string, required
	DestinationName terra.StringValue `hcl:"destination_name,attr" validate:"required"`
	// ForceUpdate: bool, optional
	ForceUpdate terra.BoolValue `hcl:"force_update,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type cloudwatchLogDestinationPolicyAttributes struct {
	ref terra.Reference
}

// AccessPolicy returns a reference to field access_policy of aws_cloudwatch_log_destination_policy.
func (cldp cloudwatchLogDestinationPolicyAttributes) AccessPolicy() terra.StringValue {
	return terra.ReferenceAsString(cldp.ref.Append("access_policy"))
}

// DestinationName returns a reference to field destination_name of aws_cloudwatch_log_destination_policy.
func (cldp cloudwatchLogDestinationPolicyAttributes) DestinationName() terra.StringValue {
	return terra.ReferenceAsString(cldp.ref.Append("destination_name"))
}

// ForceUpdate returns a reference to field force_update of aws_cloudwatch_log_destination_policy.
func (cldp cloudwatchLogDestinationPolicyAttributes) ForceUpdate() terra.BoolValue {
	return terra.ReferenceAsBool(cldp.ref.Append("force_update"))
}

// Id returns a reference to field id of aws_cloudwatch_log_destination_policy.
func (cldp cloudwatchLogDestinationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cldp.ref.Append("id"))
}

type cloudwatchLogDestinationPolicyState struct {
	AccessPolicy    string `json:"access_policy"`
	DestinationName string `json:"destination_name"`
	ForceUpdate     bool   `json:"force_update"`
	Id              string `json:"id"`
}
