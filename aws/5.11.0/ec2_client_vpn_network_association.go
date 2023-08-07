// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ec2clientvpnnetworkassociation "github.com/golingon/terraproviders/aws/5.11.0/ec2clientvpnnetworkassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2ClientVpnNetworkAssociation creates a new instance of [Ec2ClientVpnNetworkAssociation].
func NewEc2ClientVpnNetworkAssociation(name string, args Ec2ClientVpnNetworkAssociationArgs) *Ec2ClientVpnNetworkAssociation {
	return &Ec2ClientVpnNetworkAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2ClientVpnNetworkAssociation)(nil)

// Ec2ClientVpnNetworkAssociation represents the Terraform resource aws_ec2_client_vpn_network_association.
type Ec2ClientVpnNetworkAssociation struct {
	Name      string
	Args      Ec2ClientVpnNetworkAssociationArgs
	state     *ec2ClientVpnNetworkAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2ClientVpnNetworkAssociation].
func (ecvna *Ec2ClientVpnNetworkAssociation) Type() string {
	return "aws_ec2_client_vpn_network_association"
}

// LocalName returns the local name for [Ec2ClientVpnNetworkAssociation].
func (ecvna *Ec2ClientVpnNetworkAssociation) LocalName() string {
	return ecvna.Name
}

// Configuration returns the configuration (args) for [Ec2ClientVpnNetworkAssociation].
func (ecvna *Ec2ClientVpnNetworkAssociation) Configuration() interface{} {
	return ecvna.Args
}

// DependOn is used for other resources to depend on [Ec2ClientVpnNetworkAssociation].
func (ecvna *Ec2ClientVpnNetworkAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(ecvna)
}

// Dependencies returns the list of resources [Ec2ClientVpnNetworkAssociation] depends_on.
func (ecvna *Ec2ClientVpnNetworkAssociation) Dependencies() terra.Dependencies {
	return ecvna.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2ClientVpnNetworkAssociation].
func (ecvna *Ec2ClientVpnNetworkAssociation) LifecycleManagement() *terra.Lifecycle {
	return ecvna.Lifecycle
}

// Attributes returns the attributes for [Ec2ClientVpnNetworkAssociation].
func (ecvna *Ec2ClientVpnNetworkAssociation) Attributes() ec2ClientVpnNetworkAssociationAttributes {
	return ec2ClientVpnNetworkAssociationAttributes{ref: terra.ReferenceResource(ecvna)}
}

// ImportState imports the given attribute values into [Ec2ClientVpnNetworkAssociation]'s state.
func (ecvna *Ec2ClientVpnNetworkAssociation) ImportState(av io.Reader) error {
	ecvna.state = &ec2ClientVpnNetworkAssociationState{}
	if err := json.NewDecoder(av).Decode(ecvna.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ecvna.Type(), ecvna.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2ClientVpnNetworkAssociation] has state.
func (ecvna *Ec2ClientVpnNetworkAssociation) State() (*ec2ClientVpnNetworkAssociationState, bool) {
	return ecvna.state, ecvna.state != nil
}

// StateMust returns the state for [Ec2ClientVpnNetworkAssociation]. Panics if the state is nil.
func (ecvna *Ec2ClientVpnNetworkAssociation) StateMust() *ec2ClientVpnNetworkAssociationState {
	if ecvna.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ecvna.Type(), ecvna.LocalName()))
	}
	return ecvna.state
}

// Ec2ClientVpnNetworkAssociationArgs contains the configurations for aws_ec2_client_vpn_network_association.
type Ec2ClientVpnNetworkAssociationArgs struct {
	// ClientVpnEndpointId: string, required
	ClientVpnEndpointId terra.StringValue `hcl:"client_vpn_endpoint_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *ec2clientvpnnetworkassociation.Timeouts `hcl:"timeouts,block"`
}
type ec2ClientVpnNetworkAssociationAttributes struct {
	ref terra.Reference
}

// AssociationId returns a reference to field association_id of aws_ec2_client_vpn_network_association.
func (ecvna ec2ClientVpnNetworkAssociationAttributes) AssociationId() terra.StringValue {
	return terra.ReferenceAsString(ecvna.ref.Append("association_id"))
}

// ClientVpnEndpointId returns a reference to field client_vpn_endpoint_id of aws_ec2_client_vpn_network_association.
func (ecvna ec2ClientVpnNetworkAssociationAttributes) ClientVpnEndpointId() terra.StringValue {
	return terra.ReferenceAsString(ecvna.ref.Append("client_vpn_endpoint_id"))
}

// Id returns a reference to field id of aws_ec2_client_vpn_network_association.
func (ecvna ec2ClientVpnNetworkAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ecvna.ref.Append("id"))
}

// SubnetId returns a reference to field subnet_id of aws_ec2_client_vpn_network_association.
func (ecvna ec2ClientVpnNetworkAssociationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(ecvna.ref.Append("subnet_id"))
}

// VpcId returns a reference to field vpc_id of aws_ec2_client_vpn_network_association.
func (ecvna ec2ClientVpnNetworkAssociationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ecvna.ref.Append("vpc_id"))
}

func (ecvna ec2ClientVpnNetworkAssociationAttributes) Timeouts() ec2clientvpnnetworkassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ec2clientvpnnetworkassociation.TimeoutsAttributes](ecvna.ref.Append("timeouts"))
}

type ec2ClientVpnNetworkAssociationState struct {
	AssociationId       string                                        `json:"association_id"`
	ClientVpnEndpointId string                                        `json:"client_vpn_endpoint_id"`
	Id                  string                                        `json:"id"`
	SubnetId            string                                        `json:"subnet_id"`
	VpcId               string                                        `json:"vpc_id"`
	Timeouts            *ec2clientvpnnetworkassociation.TimeoutsState `json:"timeouts"`
}
