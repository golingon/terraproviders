// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagertransitgatewayregistration "github.com/golingon/terraproviders/aws/5.7.0/networkmanagertransitgatewayregistration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerTransitGatewayRegistration creates a new instance of [NetworkmanagerTransitGatewayRegistration].
func NewNetworkmanagerTransitGatewayRegistration(name string, args NetworkmanagerTransitGatewayRegistrationArgs) *NetworkmanagerTransitGatewayRegistration {
	return &NetworkmanagerTransitGatewayRegistration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerTransitGatewayRegistration)(nil)

// NetworkmanagerTransitGatewayRegistration represents the Terraform resource aws_networkmanager_transit_gateway_registration.
type NetworkmanagerTransitGatewayRegistration struct {
	Name      string
	Args      NetworkmanagerTransitGatewayRegistrationArgs
	state     *networkmanagerTransitGatewayRegistrationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerTransitGatewayRegistration].
func (ntgr *NetworkmanagerTransitGatewayRegistration) Type() string {
	return "aws_networkmanager_transit_gateway_registration"
}

// LocalName returns the local name for [NetworkmanagerTransitGatewayRegistration].
func (ntgr *NetworkmanagerTransitGatewayRegistration) LocalName() string {
	return ntgr.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerTransitGatewayRegistration].
func (ntgr *NetworkmanagerTransitGatewayRegistration) Configuration() interface{} {
	return ntgr.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerTransitGatewayRegistration].
func (ntgr *NetworkmanagerTransitGatewayRegistration) DependOn() terra.Reference {
	return terra.ReferenceResource(ntgr)
}

// Dependencies returns the list of resources [NetworkmanagerTransitGatewayRegistration] depends_on.
func (ntgr *NetworkmanagerTransitGatewayRegistration) Dependencies() terra.Dependencies {
	return ntgr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerTransitGatewayRegistration].
func (ntgr *NetworkmanagerTransitGatewayRegistration) LifecycleManagement() *terra.Lifecycle {
	return ntgr.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerTransitGatewayRegistration].
func (ntgr *NetworkmanagerTransitGatewayRegistration) Attributes() networkmanagerTransitGatewayRegistrationAttributes {
	return networkmanagerTransitGatewayRegistrationAttributes{ref: terra.ReferenceResource(ntgr)}
}

// ImportState imports the given attribute values into [NetworkmanagerTransitGatewayRegistration]'s state.
func (ntgr *NetworkmanagerTransitGatewayRegistration) ImportState(av io.Reader) error {
	ntgr.state = &networkmanagerTransitGatewayRegistrationState{}
	if err := json.NewDecoder(av).Decode(ntgr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ntgr.Type(), ntgr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerTransitGatewayRegistration] has state.
func (ntgr *NetworkmanagerTransitGatewayRegistration) State() (*networkmanagerTransitGatewayRegistrationState, bool) {
	return ntgr.state, ntgr.state != nil
}

// StateMust returns the state for [NetworkmanagerTransitGatewayRegistration]. Panics if the state is nil.
func (ntgr *NetworkmanagerTransitGatewayRegistration) StateMust() *networkmanagerTransitGatewayRegistrationState {
	if ntgr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ntgr.Type(), ntgr.LocalName()))
	}
	return ntgr.state
}

// NetworkmanagerTransitGatewayRegistrationArgs contains the configurations for aws_networkmanager_transit_gateway_registration.
type NetworkmanagerTransitGatewayRegistrationArgs struct {
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TransitGatewayArn: string, required
	TransitGatewayArn terra.StringValue `hcl:"transit_gateway_arn,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagertransitgatewayregistration.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerTransitGatewayRegistrationAttributes struct {
	ref terra.Reference
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_transit_gateway_registration.
func (ntgr networkmanagerTransitGatewayRegistrationAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(ntgr.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_transit_gateway_registration.
func (ntgr networkmanagerTransitGatewayRegistrationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ntgr.ref.Append("id"))
}

// TransitGatewayArn returns a reference to field transit_gateway_arn of aws_networkmanager_transit_gateway_registration.
func (ntgr networkmanagerTransitGatewayRegistrationAttributes) TransitGatewayArn() terra.StringValue {
	return terra.ReferenceAsString(ntgr.ref.Append("transit_gateway_arn"))
}

func (ntgr networkmanagerTransitGatewayRegistrationAttributes) Timeouts() networkmanagertransitgatewayregistration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagertransitgatewayregistration.TimeoutsAttributes](ntgr.ref.Append("timeouts"))
}

type networkmanagerTransitGatewayRegistrationState struct {
	GlobalNetworkId   string                                                  `json:"global_network_id"`
	Id                string                                                  `json:"id"`
	TransitGatewayArn string                                                  `json:"transit_gateway_arn"`
	Timeouts          *networkmanagertransitgatewayregistration.TimeoutsState `json:"timeouts"`
}