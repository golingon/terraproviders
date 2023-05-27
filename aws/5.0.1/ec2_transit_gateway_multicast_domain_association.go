// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ec2transitgatewaymulticastdomainassociation "github.com/golingon/terraproviders/aws/5.0.1/ec2transitgatewaymulticastdomainassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2TransitGatewayMulticastDomainAssociation creates a new instance of [Ec2TransitGatewayMulticastDomainAssociation].
func NewEc2TransitGatewayMulticastDomainAssociation(name string, args Ec2TransitGatewayMulticastDomainAssociationArgs) *Ec2TransitGatewayMulticastDomainAssociation {
	return &Ec2TransitGatewayMulticastDomainAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2TransitGatewayMulticastDomainAssociation)(nil)

// Ec2TransitGatewayMulticastDomainAssociation represents the Terraform resource aws_ec2_transit_gateway_multicast_domain_association.
type Ec2TransitGatewayMulticastDomainAssociation struct {
	Name      string
	Args      Ec2TransitGatewayMulticastDomainAssociationArgs
	state     *ec2TransitGatewayMulticastDomainAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2TransitGatewayMulticastDomainAssociation].
func (etgmda *Ec2TransitGatewayMulticastDomainAssociation) Type() string {
	return "aws_ec2_transit_gateway_multicast_domain_association"
}

// LocalName returns the local name for [Ec2TransitGatewayMulticastDomainAssociation].
func (etgmda *Ec2TransitGatewayMulticastDomainAssociation) LocalName() string {
	return etgmda.Name
}

// Configuration returns the configuration (args) for [Ec2TransitGatewayMulticastDomainAssociation].
func (etgmda *Ec2TransitGatewayMulticastDomainAssociation) Configuration() interface{} {
	return etgmda.Args
}

// DependOn is used for other resources to depend on [Ec2TransitGatewayMulticastDomainAssociation].
func (etgmda *Ec2TransitGatewayMulticastDomainAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(etgmda)
}

// Dependencies returns the list of resources [Ec2TransitGatewayMulticastDomainAssociation] depends_on.
func (etgmda *Ec2TransitGatewayMulticastDomainAssociation) Dependencies() terra.Dependencies {
	return etgmda.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2TransitGatewayMulticastDomainAssociation].
func (etgmda *Ec2TransitGatewayMulticastDomainAssociation) LifecycleManagement() *terra.Lifecycle {
	return etgmda.Lifecycle
}

// Attributes returns the attributes for [Ec2TransitGatewayMulticastDomainAssociation].
func (etgmda *Ec2TransitGatewayMulticastDomainAssociation) Attributes() ec2TransitGatewayMulticastDomainAssociationAttributes {
	return ec2TransitGatewayMulticastDomainAssociationAttributes{ref: terra.ReferenceResource(etgmda)}
}

// ImportState imports the given attribute values into [Ec2TransitGatewayMulticastDomainAssociation]'s state.
func (etgmda *Ec2TransitGatewayMulticastDomainAssociation) ImportState(av io.Reader) error {
	etgmda.state = &ec2TransitGatewayMulticastDomainAssociationState{}
	if err := json.NewDecoder(av).Decode(etgmda.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", etgmda.Type(), etgmda.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2TransitGatewayMulticastDomainAssociation] has state.
func (etgmda *Ec2TransitGatewayMulticastDomainAssociation) State() (*ec2TransitGatewayMulticastDomainAssociationState, bool) {
	return etgmda.state, etgmda.state != nil
}

// StateMust returns the state for [Ec2TransitGatewayMulticastDomainAssociation]. Panics if the state is nil.
func (etgmda *Ec2TransitGatewayMulticastDomainAssociation) StateMust() *ec2TransitGatewayMulticastDomainAssociationState {
	if etgmda.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", etgmda.Type(), etgmda.LocalName()))
	}
	return etgmda.state
}

// Ec2TransitGatewayMulticastDomainAssociationArgs contains the configurations for aws_ec2_transit_gateway_multicast_domain_association.
type Ec2TransitGatewayMulticastDomainAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// TransitGatewayAttachmentId: string, required
	TransitGatewayAttachmentId terra.StringValue `hcl:"transit_gateway_attachment_id,attr" validate:"required"`
	// TransitGatewayMulticastDomainId: string, required
	TransitGatewayMulticastDomainId terra.StringValue `hcl:"transit_gateway_multicast_domain_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *ec2transitgatewaymulticastdomainassociation.Timeouts `hcl:"timeouts,block"`
}
type ec2TransitGatewayMulticastDomainAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_transit_gateway_multicast_domain_association.
func (etgmda ec2TransitGatewayMulticastDomainAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgmda.ref.Append("id"))
}

// SubnetId returns a reference to field subnet_id of aws_ec2_transit_gateway_multicast_domain_association.
func (etgmda ec2TransitGatewayMulticastDomainAssociationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(etgmda.ref.Append("subnet_id"))
}

// TransitGatewayAttachmentId returns a reference to field transit_gateway_attachment_id of aws_ec2_transit_gateway_multicast_domain_association.
func (etgmda ec2TransitGatewayMulticastDomainAssociationAttributes) TransitGatewayAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(etgmda.ref.Append("transit_gateway_attachment_id"))
}

// TransitGatewayMulticastDomainId returns a reference to field transit_gateway_multicast_domain_id of aws_ec2_transit_gateway_multicast_domain_association.
func (etgmda ec2TransitGatewayMulticastDomainAssociationAttributes) TransitGatewayMulticastDomainId() terra.StringValue {
	return terra.ReferenceAsString(etgmda.ref.Append("transit_gateway_multicast_domain_id"))
}

func (etgmda ec2TransitGatewayMulticastDomainAssociationAttributes) Timeouts() ec2transitgatewaymulticastdomainassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ec2transitgatewaymulticastdomainassociation.TimeoutsAttributes](etgmda.ref.Append("timeouts"))
}

type ec2TransitGatewayMulticastDomainAssociationState struct {
	Id                              string                                                     `json:"id"`
	SubnetId                        string                                                     `json:"subnet_id"`
	TransitGatewayAttachmentId      string                                                     `json:"transit_gateway_attachment_id"`
	TransitGatewayMulticastDomainId string                                                     `json:"transit_gateway_multicast_domain_id"`
	Timeouts                        *ec2transitgatewaymulticastdomainassociation.TimeoutsState `json:"timeouts"`
}
