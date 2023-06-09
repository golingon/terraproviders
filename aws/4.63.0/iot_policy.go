// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotPolicy creates a new instance of [IotPolicy].
func NewIotPolicy(name string, args IotPolicyArgs) *IotPolicy {
	return &IotPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotPolicy)(nil)

// IotPolicy represents the Terraform resource aws_iot_policy.
type IotPolicy struct {
	Name      string
	Args      IotPolicyArgs
	state     *iotPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotPolicy].
func (ip *IotPolicy) Type() string {
	return "aws_iot_policy"
}

// LocalName returns the local name for [IotPolicy].
func (ip *IotPolicy) LocalName() string {
	return ip.Name
}

// Configuration returns the configuration (args) for [IotPolicy].
func (ip *IotPolicy) Configuration() interface{} {
	return ip.Args
}

// DependOn is used for other resources to depend on [IotPolicy].
func (ip *IotPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ip)
}

// Dependencies returns the list of resources [IotPolicy] depends_on.
func (ip *IotPolicy) Dependencies() terra.Dependencies {
	return ip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotPolicy].
func (ip *IotPolicy) LifecycleManagement() *terra.Lifecycle {
	return ip.Lifecycle
}

// Attributes returns the attributes for [IotPolicy].
func (ip *IotPolicy) Attributes() iotPolicyAttributes {
	return iotPolicyAttributes{ref: terra.ReferenceResource(ip)}
}

// ImportState imports the given attribute values into [IotPolicy]'s state.
func (ip *IotPolicy) ImportState(av io.Reader) error {
	ip.state = &iotPolicyState{}
	if err := json.NewDecoder(av).Decode(ip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ip.Type(), ip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotPolicy] has state.
func (ip *IotPolicy) State() (*iotPolicyState, bool) {
	return ip.state, ip.state != nil
}

// StateMust returns the state for [IotPolicy]. Panics if the state is nil.
func (ip *IotPolicy) StateMust() *iotPolicyState {
	if ip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ip.Type(), ip.LocalName()))
	}
	return ip.state
}

// IotPolicyArgs contains the configurations for aws_iot_policy.
type IotPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
}
type iotPolicyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iot_policy.
func (ip iotPolicyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("arn"))
}

// DefaultVersionId returns a reference to field default_version_id of aws_iot_policy.
func (ip iotPolicyAttributes) DefaultVersionId() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("default_version_id"))
}

// Id returns a reference to field id of aws_iot_policy.
func (ip iotPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("id"))
}

// Name returns a reference to field name of aws_iot_policy.
func (ip iotPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("name"))
}

// Policy returns a reference to field policy of aws_iot_policy.
func (ip iotPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("policy"))
}

type iotPolicyState struct {
	Arn              string `json:"arn"`
	DefaultVersionId string `json:"default_version_id"`
	Id               string `json:"id"`
	Name             string `json:"name"`
	Policy           string `json:"policy"`
}
