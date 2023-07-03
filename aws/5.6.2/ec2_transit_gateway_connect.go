// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ec2transitgatewayconnect "github.com/golingon/terraproviders/aws/5.6.2/ec2transitgatewayconnect"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2TransitGatewayConnect creates a new instance of [Ec2TransitGatewayConnect].
func NewEc2TransitGatewayConnect(name string, args Ec2TransitGatewayConnectArgs) *Ec2TransitGatewayConnect {
	return &Ec2TransitGatewayConnect{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2TransitGatewayConnect)(nil)

// Ec2TransitGatewayConnect represents the Terraform resource aws_ec2_transit_gateway_connect.
type Ec2TransitGatewayConnect struct {
	Name      string
	Args      Ec2TransitGatewayConnectArgs
	state     *ec2TransitGatewayConnectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2TransitGatewayConnect].
func (etgc *Ec2TransitGatewayConnect) Type() string {
	return "aws_ec2_transit_gateway_connect"
}

// LocalName returns the local name for [Ec2TransitGatewayConnect].
func (etgc *Ec2TransitGatewayConnect) LocalName() string {
	return etgc.Name
}

// Configuration returns the configuration (args) for [Ec2TransitGatewayConnect].
func (etgc *Ec2TransitGatewayConnect) Configuration() interface{} {
	return etgc.Args
}

// DependOn is used for other resources to depend on [Ec2TransitGatewayConnect].
func (etgc *Ec2TransitGatewayConnect) DependOn() terra.Reference {
	return terra.ReferenceResource(etgc)
}

// Dependencies returns the list of resources [Ec2TransitGatewayConnect] depends_on.
func (etgc *Ec2TransitGatewayConnect) Dependencies() terra.Dependencies {
	return etgc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2TransitGatewayConnect].
func (etgc *Ec2TransitGatewayConnect) LifecycleManagement() *terra.Lifecycle {
	return etgc.Lifecycle
}

// Attributes returns the attributes for [Ec2TransitGatewayConnect].
func (etgc *Ec2TransitGatewayConnect) Attributes() ec2TransitGatewayConnectAttributes {
	return ec2TransitGatewayConnectAttributes{ref: terra.ReferenceResource(etgc)}
}

// ImportState imports the given attribute values into [Ec2TransitGatewayConnect]'s state.
func (etgc *Ec2TransitGatewayConnect) ImportState(av io.Reader) error {
	etgc.state = &ec2TransitGatewayConnectState{}
	if err := json.NewDecoder(av).Decode(etgc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", etgc.Type(), etgc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2TransitGatewayConnect] has state.
func (etgc *Ec2TransitGatewayConnect) State() (*ec2TransitGatewayConnectState, bool) {
	return etgc.state, etgc.state != nil
}

// StateMust returns the state for [Ec2TransitGatewayConnect]. Panics if the state is nil.
func (etgc *Ec2TransitGatewayConnect) StateMust() *ec2TransitGatewayConnectState {
	if etgc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", etgc.Type(), etgc.LocalName()))
	}
	return etgc.state
}

// Ec2TransitGatewayConnectArgs contains the configurations for aws_ec2_transit_gateway_connect.
type Ec2TransitGatewayConnectArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TransitGatewayDefaultRouteTableAssociation: bool, optional
	TransitGatewayDefaultRouteTableAssociation terra.BoolValue `hcl:"transit_gateway_default_route_table_association,attr"`
	// TransitGatewayDefaultRouteTablePropagation: bool, optional
	TransitGatewayDefaultRouteTablePropagation terra.BoolValue `hcl:"transit_gateway_default_route_table_propagation,attr"`
	// TransitGatewayId: string, required
	TransitGatewayId terra.StringValue `hcl:"transit_gateway_id,attr" validate:"required"`
	// TransportAttachmentId: string, required
	TransportAttachmentId terra.StringValue `hcl:"transport_attachment_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *ec2transitgatewayconnect.Timeouts `hcl:"timeouts,block"`
}
type ec2TransitGatewayConnectAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_transit_gateway_connect.
func (etgc ec2TransitGatewayConnectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgc.ref.Append("id"))
}

// Protocol returns a reference to field protocol of aws_ec2_transit_gateway_connect.
func (etgc ec2TransitGatewayConnectAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(etgc.ref.Append("protocol"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_connect.
func (etgc ec2TransitGatewayConnectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_transit_gateway_connect.
func (etgc ec2TransitGatewayConnectAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgc.ref.Append("tags_all"))
}

// TransitGatewayDefaultRouteTableAssociation returns a reference to field transit_gateway_default_route_table_association of aws_ec2_transit_gateway_connect.
func (etgc ec2TransitGatewayConnectAttributes) TransitGatewayDefaultRouteTableAssociation() terra.BoolValue {
	return terra.ReferenceAsBool(etgc.ref.Append("transit_gateway_default_route_table_association"))
}

// TransitGatewayDefaultRouteTablePropagation returns a reference to field transit_gateway_default_route_table_propagation of aws_ec2_transit_gateway_connect.
func (etgc ec2TransitGatewayConnectAttributes) TransitGatewayDefaultRouteTablePropagation() terra.BoolValue {
	return terra.ReferenceAsBool(etgc.ref.Append("transit_gateway_default_route_table_propagation"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_ec2_transit_gateway_connect.
func (etgc ec2TransitGatewayConnectAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(etgc.ref.Append("transit_gateway_id"))
}

// TransportAttachmentId returns a reference to field transport_attachment_id of aws_ec2_transit_gateway_connect.
func (etgc ec2TransitGatewayConnectAttributes) TransportAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(etgc.ref.Append("transport_attachment_id"))
}

func (etgc ec2TransitGatewayConnectAttributes) Timeouts() ec2transitgatewayconnect.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ec2transitgatewayconnect.TimeoutsAttributes](etgc.ref.Append("timeouts"))
}

type ec2TransitGatewayConnectState struct {
	Id                                         string                                  `json:"id"`
	Protocol                                   string                                  `json:"protocol"`
	Tags                                       map[string]string                       `json:"tags"`
	TagsAll                                    map[string]string                       `json:"tags_all"`
	TransitGatewayDefaultRouteTableAssociation bool                                    `json:"transit_gateway_default_route_table_association"`
	TransitGatewayDefaultRouteTablePropagation bool                                    `json:"transit_gateway_default_route_table_propagation"`
	TransitGatewayId                           string                                  `json:"transit_gateway_id"`
	TransportAttachmentId                      string                                  `json:"transport_attachment_id"`
	Timeouts                                   *ec2transitgatewayconnect.TimeoutsState `json:"timeouts"`
}
