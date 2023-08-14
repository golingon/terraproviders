// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2TransitGatewayPolicyTable creates a new instance of [Ec2TransitGatewayPolicyTable].
func NewEc2TransitGatewayPolicyTable(name string, args Ec2TransitGatewayPolicyTableArgs) *Ec2TransitGatewayPolicyTable {
	return &Ec2TransitGatewayPolicyTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2TransitGatewayPolicyTable)(nil)

// Ec2TransitGatewayPolicyTable represents the Terraform resource aws_ec2_transit_gateway_policy_table.
type Ec2TransitGatewayPolicyTable struct {
	Name      string
	Args      Ec2TransitGatewayPolicyTableArgs
	state     *ec2TransitGatewayPolicyTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2TransitGatewayPolicyTable].
func (etgpt *Ec2TransitGatewayPolicyTable) Type() string {
	return "aws_ec2_transit_gateway_policy_table"
}

// LocalName returns the local name for [Ec2TransitGatewayPolicyTable].
func (etgpt *Ec2TransitGatewayPolicyTable) LocalName() string {
	return etgpt.Name
}

// Configuration returns the configuration (args) for [Ec2TransitGatewayPolicyTable].
func (etgpt *Ec2TransitGatewayPolicyTable) Configuration() interface{} {
	return etgpt.Args
}

// DependOn is used for other resources to depend on [Ec2TransitGatewayPolicyTable].
func (etgpt *Ec2TransitGatewayPolicyTable) DependOn() terra.Reference {
	return terra.ReferenceResource(etgpt)
}

// Dependencies returns the list of resources [Ec2TransitGatewayPolicyTable] depends_on.
func (etgpt *Ec2TransitGatewayPolicyTable) Dependencies() terra.Dependencies {
	return etgpt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2TransitGatewayPolicyTable].
func (etgpt *Ec2TransitGatewayPolicyTable) LifecycleManagement() *terra.Lifecycle {
	return etgpt.Lifecycle
}

// Attributes returns the attributes for [Ec2TransitGatewayPolicyTable].
func (etgpt *Ec2TransitGatewayPolicyTable) Attributes() ec2TransitGatewayPolicyTableAttributes {
	return ec2TransitGatewayPolicyTableAttributes{ref: terra.ReferenceResource(etgpt)}
}

// ImportState imports the given attribute values into [Ec2TransitGatewayPolicyTable]'s state.
func (etgpt *Ec2TransitGatewayPolicyTable) ImportState(av io.Reader) error {
	etgpt.state = &ec2TransitGatewayPolicyTableState{}
	if err := json.NewDecoder(av).Decode(etgpt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", etgpt.Type(), etgpt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2TransitGatewayPolicyTable] has state.
func (etgpt *Ec2TransitGatewayPolicyTable) State() (*ec2TransitGatewayPolicyTableState, bool) {
	return etgpt.state, etgpt.state != nil
}

// StateMust returns the state for [Ec2TransitGatewayPolicyTable]. Panics if the state is nil.
func (etgpt *Ec2TransitGatewayPolicyTable) StateMust() *ec2TransitGatewayPolicyTableState {
	if etgpt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", etgpt.Type(), etgpt.LocalName()))
	}
	return etgpt.state
}

// Ec2TransitGatewayPolicyTableArgs contains the configurations for aws_ec2_transit_gateway_policy_table.
type Ec2TransitGatewayPolicyTableArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TransitGatewayId: string, required
	TransitGatewayId terra.StringValue `hcl:"transit_gateway_id,attr" validate:"required"`
}
type ec2TransitGatewayPolicyTableAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_transit_gateway_policy_table.
func (etgpt ec2TransitGatewayPolicyTableAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(etgpt.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ec2_transit_gateway_policy_table.
func (etgpt ec2TransitGatewayPolicyTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgpt.ref.Append("id"))
}

// State returns a reference to field state of aws_ec2_transit_gateway_policy_table.
func (etgpt ec2TransitGatewayPolicyTableAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(etgpt.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_policy_table.
func (etgpt ec2TransitGatewayPolicyTableAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgpt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_transit_gateway_policy_table.
func (etgpt ec2TransitGatewayPolicyTableAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgpt.ref.Append("tags_all"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_ec2_transit_gateway_policy_table.
func (etgpt ec2TransitGatewayPolicyTableAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(etgpt.ref.Append("transit_gateway_id"))
}

type ec2TransitGatewayPolicyTableState struct {
	Arn              string            `json:"arn"`
	Id               string            `json:"id"`
	State            string            `json:"state"`
	Tags             map[string]string `json:"tags"`
	TagsAll          map[string]string `json:"tags_all"`
	TransitGatewayId string            `json:"transit_gateway_id"`
}