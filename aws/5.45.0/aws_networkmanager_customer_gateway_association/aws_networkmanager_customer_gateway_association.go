// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_networkmanager_customer_gateway_association

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_networkmanager_customer_gateway_association.
type Resource struct {
	Name      string
	Args      Args
	state     *awsNetworkmanagerCustomerGatewayAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ancga *Resource) Type() string {
	return "aws_networkmanager_customer_gateway_association"
}

// LocalName returns the local name for [Resource].
func (ancga *Resource) LocalName() string {
	return ancga.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ancga *Resource) Configuration() interface{} {
	return ancga.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ancga *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ancga)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ancga *Resource) Dependencies() terra.Dependencies {
	return ancga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ancga *Resource) LifecycleManagement() *terra.Lifecycle {
	return ancga.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ancga *Resource) Attributes() awsNetworkmanagerCustomerGatewayAssociationAttributes {
	return awsNetworkmanagerCustomerGatewayAssociationAttributes{ref: terra.ReferenceResource(ancga)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ancga *Resource) ImportState(state io.Reader) error {
	ancga.state = &awsNetworkmanagerCustomerGatewayAssociationState{}
	if err := json.NewDecoder(state).Decode(ancga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ancga.Type(), ancga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ancga *Resource) State() (*awsNetworkmanagerCustomerGatewayAssociationState, bool) {
	return ancga.state, ancga.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ancga *Resource) StateMust() *awsNetworkmanagerCustomerGatewayAssociationState {
	if ancga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ancga.Type(), ancga.LocalName()))
	}
	return ancga.state
}

// Args contains the configurations for aws_networkmanager_customer_gateway_association.
type Args struct {
	// CustomerGatewayArn: string, required
	CustomerGatewayArn terra.StringValue `hcl:"customer_gateway_arn,attr" validate:"required"`
	// DeviceId: string, required
	DeviceId terra.StringValue `hcl:"device_id,attr" validate:"required"`
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkId: string, optional
	LinkId terra.StringValue `hcl:"link_id,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsNetworkmanagerCustomerGatewayAssociationAttributes struct {
	ref terra.Reference
}

// CustomerGatewayArn returns a reference to field customer_gateway_arn of aws_networkmanager_customer_gateway_association.
func (ancga awsNetworkmanagerCustomerGatewayAssociationAttributes) CustomerGatewayArn() terra.StringValue {
	return terra.ReferenceAsString(ancga.ref.Append("customer_gateway_arn"))
}

// DeviceId returns a reference to field device_id of aws_networkmanager_customer_gateway_association.
func (ancga awsNetworkmanagerCustomerGatewayAssociationAttributes) DeviceId() terra.StringValue {
	return terra.ReferenceAsString(ancga.ref.Append("device_id"))
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_customer_gateway_association.
func (ancga awsNetworkmanagerCustomerGatewayAssociationAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(ancga.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_customer_gateway_association.
func (ancga awsNetworkmanagerCustomerGatewayAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ancga.ref.Append("id"))
}

// LinkId returns a reference to field link_id of aws_networkmanager_customer_gateway_association.
func (ancga awsNetworkmanagerCustomerGatewayAssociationAttributes) LinkId() terra.StringValue {
	return terra.ReferenceAsString(ancga.ref.Append("link_id"))
}

func (ancga awsNetworkmanagerCustomerGatewayAssociationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ancga.ref.Append("timeouts"))
}

type awsNetworkmanagerCustomerGatewayAssociationState struct {
	CustomerGatewayArn string         `json:"customer_gateway_arn"`
	DeviceId           string         `json:"device_id"`
	GlobalNetworkId    string         `json:"global_network_id"`
	Id                 string         `json:"id"`
	LinkId             string         `json:"link_id"`
	Timeouts           *TimeoutsState `json:"timeouts"`
}
