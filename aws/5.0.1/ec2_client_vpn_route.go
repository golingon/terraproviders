// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ec2clientvpnroute "github.com/golingon/terraproviders/aws/5.0.1/ec2clientvpnroute"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2ClientVpnRoute creates a new instance of [Ec2ClientVpnRoute].
func NewEc2ClientVpnRoute(name string, args Ec2ClientVpnRouteArgs) *Ec2ClientVpnRoute {
	return &Ec2ClientVpnRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2ClientVpnRoute)(nil)

// Ec2ClientVpnRoute represents the Terraform resource aws_ec2_client_vpn_route.
type Ec2ClientVpnRoute struct {
	Name      string
	Args      Ec2ClientVpnRouteArgs
	state     *ec2ClientVpnRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2ClientVpnRoute].
func (ecvr *Ec2ClientVpnRoute) Type() string {
	return "aws_ec2_client_vpn_route"
}

// LocalName returns the local name for [Ec2ClientVpnRoute].
func (ecvr *Ec2ClientVpnRoute) LocalName() string {
	return ecvr.Name
}

// Configuration returns the configuration (args) for [Ec2ClientVpnRoute].
func (ecvr *Ec2ClientVpnRoute) Configuration() interface{} {
	return ecvr.Args
}

// DependOn is used for other resources to depend on [Ec2ClientVpnRoute].
func (ecvr *Ec2ClientVpnRoute) DependOn() terra.Reference {
	return terra.ReferenceResource(ecvr)
}

// Dependencies returns the list of resources [Ec2ClientVpnRoute] depends_on.
func (ecvr *Ec2ClientVpnRoute) Dependencies() terra.Dependencies {
	return ecvr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2ClientVpnRoute].
func (ecvr *Ec2ClientVpnRoute) LifecycleManagement() *terra.Lifecycle {
	return ecvr.Lifecycle
}

// Attributes returns the attributes for [Ec2ClientVpnRoute].
func (ecvr *Ec2ClientVpnRoute) Attributes() ec2ClientVpnRouteAttributes {
	return ec2ClientVpnRouteAttributes{ref: terra.ReferenceResource(ecvr)}
}

// ImportState imports the given attribute values into [Ec2ClientVpnRoute]'s state.
func (ecvr *Ec2ClientVpnRoute) ImportState(av io.Reader) error {
	ecvr.state = &ec2ClientVpnRouteState{}
	if err := json.NewDecoder(av).Decode(ecvr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ecvr.Type(), ecvr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2ClientVpnRoute] has state.
func (ecvr *Ec2ClientVpnRoute) State() (*ec2ClientVpnRouteState, bool) {
	return ecvr.state, ecvr.state != nil
}

// StateMust returns the state for [Ec2ClientVpnRoute]. Panics if the state is nil.
func (ecvr *Ec2ClientVpnRoute) StateMust() *ec2ClientVpnRouteState {
	if ecvr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ecvr.Type(), ecvr.LocalName()))
	}
	return ecvr.state
}

// Ec2ClientVpnRouteArgs contains the configurations for aws_ec2_client_vpn_route.
type Ec2ClientVpnRouteArgs struct {
	// ClientVpnEndpointId: string, required
	ClientVpnEndpointId terra.StringValue `hcl:"client_vpn_endpoint_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DestinationCidrBlock: string, required
	DestinationCidrBlock terra.StringValue `hcl:"destination_cidr_block,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TargetVpcSubnetId: string, required
	TargetVpcSubnetId terra.StringValue `hcl:"target_vpc_subnet_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *ec2clientvpnroute.Timeouts `hcl:"timeouts,block"`
}
type ec2ClientVpnRouteAttributes struct {
	ref terra.Reference
}

// ClientVpnEndpointId returns a reference to field client_vpn_endpoint_id of aws_ec2_client_vpn_route.
func (ecvr ec2ClientVpnRouteAttributes) ClientVpnEndpointId() terra.StringValue {
	return terra.ReferenceAsString(ecvr.ref.Append("client_vpn_endpoint_id"))
}

// Description returns a reference to field description of aws_ec2_client_vpn_route.
func (ecvr ec2ClientVpnRouteAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ecvr.ref.Append("description"))
}

// DestinationCidrBlock returns a reference to field destination_cidr_block of aws_ec2_client_vpn_route.
func (ecvr ec2ClientVpnRouteAttributes) DestinationCidrBlock() terra.StringValue {
	return terra.ReferenceAsString(ecvr.ref.Append("destination_cidr_block"))
}

// Id returns a reference to field id of aws_ec2_client_vpn_route.
func (ecvr ec2ClientVpnRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ecvr.ref.Append("id"))
}

// Origin returns a reference to field origin of aws_ec2_client_vpn_route.
func (ecvr ec2ClientVpnRouteAttributes) Origin() terra.StringValue {
	return terra.ReferenceAsString(ecvr.ref.Append("origin"))
}

// TargetVpcSubnetId returns a reference to field target_vpc_subnet_id of aws_ec2_client_vpn_route.
func (ecvr ec2ClientVpnRouteAttributes) TargetVpcSubnetId() terra.StringValue {
	return terra.ReferenceAsString(ecvr.ref.Append("target_vpc_subnet_id"))
}

// Type returns a reference to field type of aws_ec2_client_vpn_route.
func (ecvr ec2ClientVpnRouteAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ecvr.ref.Append("type"))
}

func (ecvr ec2ClientVpnRouteAttributes) Timeouts() ec2clientvpnroute.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ec2clientvpnroute.TimeoutsAttributes](ecvr.ref.Append("timeouts"))
}

type ec2ClientVpnRouteState struct {
	ClientVpnEndpointId  string                           `json:"client_vpn_endpoint_id"`
	Description          string                           `json:"description"`
	DestinationCidrBlock string                           `json:"destination_cidr_block"`
	Id                   string                           `json:"id"`
	Origin               string                           `json:"origin"`
	TargetVpcSubnetId    string                           `json:"target_vpc_subnet_id"`
	Type                 string                           `json:"type"`
	Timeouts             *ec2clientvpnroute.TimeoutsState `json:"timeouts"`
}
