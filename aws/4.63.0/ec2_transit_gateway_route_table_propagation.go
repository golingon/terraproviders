// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2TransitGatewayRouteTablePropagation creates a new instance of [Ec2TransitGatewayRouteTablePropagation].
func NewEc2TransitGatewayRouteTablePropagation(name string, args Ec2TransitGatewayRouteTablePropagationArgs) *Ec2TransitGatewayRouteTablePropagation {
	return &Ec2TransitGatewayRouteTablePropagation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2TransitGatewayRouteTablePropagation)(nil)

// Ec2TransitGatewayRouteTablePropagation represents the Terraform resource aws_ec2_transit_gateway_route_table_propagation.
type Ec2TransitGatewayRouteTablePropagation struct {
	Name      string
	Args      Ec2TransitGatewayRouteTablePropagationArgs
	state     *ec2TransitGatewayRouteTablePropagationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2TransitGatewayRouteTablePropagation].
func (etgrtp *Ec2TransitGatewayRouteTablePropagation) Type() string {
	return "aws_ec2_transit_gateway_route_table_propagation"
}

// LocalName returns the local name for [Ec2TransitGatewayRouteTablePropagation].
func (etgrtp *Ec2TransitGatewayRouteTablePropagation) LocalName() string {
	return etgrtp.Name
}

// Configuration returns the configuration (args) for [Ec2TransitGatewayRouteTablePropagation].
func (etgrtp *Ec2TransitGatewayRouteTablePropagation) Configuration() interface{} {
	return etgrtp.Args
}

// DependOn is used for other resources to depend on [Ec2TransitGatewayRouteTablePropagation].
func (etgrtp *Ec2TransitGatewayRouteTablePropagation) DependOn() terra.Reference {
	return terra.ReferenceResource(etgrtp)
}

// Dependencies returns the list of resources [Ec2TransitGatewayRouteTablePropagation] depends_on.
func (etgrtp *Ec2TransitGatewayRouteTablePropagation) Dependencies() terra.Dependencies {
	return etgrtp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2TransitGatewayRouteTablePropagation].
func (etgrtp *Ec2TransitGatewayRouteTablePropagation) LifecycleManagement() *terra.Lifecycle {
	return etgrtp.Lifecycle
}

// Attributes returns the attributes for [Ec2TransitGatewayRouteTablePropagation].
func (etgrtp *Ec2TransitGatewayRouteTablePropagation) Attributes() ec2TransitGatewayRouteTablePropagationAttributes {
	return ec2TransitGatewayRouteTablePropagationAttributes{ref: terra.ReferenceResource(etgrtp)}
}

// ImportState imports the given attribute values into [Ec2TransitGatewayRouteTablePropagation]'s state.
func (etgrtp *Ec2TransitGatewayRouteTablePropagation) ImportState(av io.Reader) error {
	etgrtp.state = &ec2TransitGatewayRouteTablePropagationState{}
	if err := json.NewDecoder(av).Decode(etgrtp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", etgrtp.Type(), etgrtp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2TransitGatewayRouteTablePropagation] has state.
func (etgrtp *Ec2TransitGatewayRouteTablePropagation) State() (*ec2TransitGatewayRouteTablePropagationState, bool) {
	return etgrtp.state, etgrtp.state != nil
}

// StateMust returns the state for [Ec2TransitGatewayRouteTablePropagation]. Panics if the state is nil.
func (etgrtp *Ec2TransitGatewayRouteTablePropagation) StateMust() *ec2TransitGatewayRouteTablePropagationState {
	if etgrtp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", etgrtp.Type(), etgrtp.LocalName()))
	}
	return etgrtp.state
}

// Ec2TransitGatewayRouteTablePropagationArgs contains the configurations for aws_ec2_transit_gateway_route_table_propagation.
type Ec2TransitGatewayRouteTablePropagationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TransitGatewayAttachmentId: string, required
	TransitGatewayAttachmentId terra.StringValue `hcl:"transit_gateway_attachment_id,attr" validate:"required"`
	// TransitGatewayRouteTableId: string, required
	TransitGatewayRouteTableId terra.StringValue `hcl:"transit_gateway_route_table_id,attr" validate:"required"`
}
type ec2TransitGatewayRouteTablePropagationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_transit_gateway_route_table_propagation.
func (etgrtp ec2TransitGatewayRouteTablePropagationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgrtp.ref.Append("id"))
}

// ResourceId returns a reference to field resource_id of aws_ec2_transit_gateway_route_table_propagation.
func (etgrtp ec2TransitGatewayRouteTablePropagationAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(etgrtp.ref.Append("resource_id"))
}

// ResourceType returns a reference to field resource_type of aws_ec2_transit_gateway_route_table_propagation.
func (etgrtp ec2TransitGatewayRouteTablePropagationAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(etgrtp.ref.Append("resource_type"))
}

// TransitGatewayAttachmentId returns a reference to field transit_gateway_attachment_id of aws_ec2_transit_gateway_route_table_propagation.
func (etgrtp ec2TransitGatewayRouteTablePropagationAttributes) TransitGatewayAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(etgrtp.ref.Append("transit_gateway_attachment_id"))
}

// TransitGatewayRouteTableId returns a reference to field transit_gateway_route_table_id of aws_ec2_transit_gateway_route_table_propagation.
func (etgrtp ec2TransitGatewayRouteTablePropagationAttributes) TransitGatewayRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(etgrtp.ref.Append("transit_gateway_route_table_id"))
}

type ec2TransitGatewayRouteTablePropagationState struct {
	Id                         string `json:"id"`
	ResourceId                 string `json:"resource_id"`
	ResourceType               string `json:"resource_type"`
	TransitGatewayAttachmentId string `json:"transit_gateway_attachment_id"`
	TransitGatewayRouteTableId string `json:"transit_gateway_route_table_id"`
}
