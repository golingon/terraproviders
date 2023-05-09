// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpngatewayroutepropagation "github.com/golingon/terraproviders/aws/4.66.1/vpngatewayroutepropagation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpnGatewayRoutePropagation creates a new instance of [VpnGatewayRoutePropagation].
func NewVpnGatewayRoutePropagation(name string, args VpnGatewayRoutePropagationArgs) *VpnGatewayRoutePropagation {
	return &VpnGatewayRoutePropagation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpnGatewayRoutePropagation)(nil)

// VpnGatewayRoutePropagation represents the Terraform resource aws_vpn_gateway_route_propagation.
type VpnGatewayRoutePropagation struct {
	Name      string
	Args      VpnGatewayRoutePropagationArgs
	state     *vpnGatewayRoutePropagationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpnGatewayRoutePropagation].
func (vgrp *VpnGatewayRoutePropagation) Type() string {
	return "aws_vpn_gateway_route_propagation"
}

// LocalName returns the local name for [VpnGatewayRoutePropagation].
func (vgrp *VpnGatewayRoutePropagation) LocalName() string {
	return vgrp.Name
}

// Configuration returns the configuration (args) for [VpnGatewayRoutePropagation].
func (vgrp *VpnGatewayRoutePropagation) Configuration() interface{} {
	return vgrp.Args
}

// DependOn is used for other resources to depend on [VpnGatewayRoutePropagation].
func (vgrp *VpnGatewayRoutePropagation) DependOn() terra.Reference {
	return terra.ReferenceResource(vgrp)
}

// Dependencies returns the list of resources [VpnGatewayRoutePropagation] depends_on.
func (vgrp *VpnGatewayRoutePropagation) Dependencies() terra.Dependencies {
	return vgrp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpnGatewayRoutePropagation].
func (vgrp *VpnGatewayRoutePropagation) LifecycleManagement() *terra.Lifecycle {
	return vgrp.Lifecycle
}

// Attributes returns the attributes for [VpnGatewayRoutePropagation].
func (vgrp *VpnGatewayRoutePropagation) Attributes() vpnGatewayRoutePropagationAttributes {
	return vpnGatewayRoutePropagationAttributes{ref: terra.ReferenceResource(vgrp)}
}

// ImportState imports the given attribute values into [VpnGatewayRoutePropagation]'s state.
func (vgrp *VpnGatewayRoutePropagation) ImportState(av io.Reader) error {
	vgrp.state = &vpnGatewayRoutePropagationState{}
	if err := json.NewDecoder(av).Decode(vgrp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vgrp.Type(), vgrp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpnGatewayRoutePropagation] has state.
func (vgrp *VpnGatewayRoutePropagation) State() (*vpnGatewayRoutePropagationState, bool) {
	return vgrp.state, vgrp.state != nil
}

// StateMust returns the state for [VpnGatewayRoutePropagation]. Panics if the state is nil.
func (vgrp *VpnGatewayRoutePropagation) StateMust() *vpnGatewayRoutePropagationState {
	if vgrp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vgrp.Type(), vgrp.LocalName()))
	}
	return vgrp.state
}

// VpnGatewayRoutePropagationArgs contains the configurations for aws_vpn_gateway_route_propagation.
type VpnGatewayRoutePropagationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RouteTableId: string, required
	RouteTableId terra.StringValue `hcl:"route_table_id,attr" validate:"required"`
	// VpnGatewayId: string, required
	VpnGatewayId terra.StringValue `hcl:"vpn_gateway_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *vpngatewayroutepropagation.Timeouts `hcl:"timeouts,block"`
}
type vpnGatewayRoutePropagationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpn_gateway_route_propagation.
func (vgrp vpnGatewayRoutePropagationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vgrp.ref.Append("id"))
}

// RouteTableId returns a reference to field route_table_id of aws_vpn_gateway_route_propagation.
func (vgrp vpnGatewayRoutePropagationAttributes) RouteTableId() terra.StringValue {
	return terra.ReferenceAsString(vgrp.ref.Append("route_table_id"))
}

// VpnGatewayId returns a reference to field vpn_gateway_id of aws_vpn_gateway_route_propagation.
func (vgrp vpnGatewayRoutePropagationAttributes) VpnGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vgrp.ref.Append("vpn_gateway_id"))
}

func (vgrp vpnGatewayRoutePropagationAttributes) Timeouts() vpngatewayroutepropagation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpngatewayroutepropagation.TimeoutsAttributes](vgrp.ref.Append("timeouts"))
}

type vpnGatewayRoutePropagationState struct {
	Id           string                                    `json:"id"`
	RouteTableId string                                    `json:"route_table_id"`
	VpnGatewayId string                                    `json:"vpn_gateway_id"`
	Timeouts     *vpngatewayroutepropagation.TimeoutsState `json:"timeouts"`
}
