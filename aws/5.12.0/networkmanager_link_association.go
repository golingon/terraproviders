// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagerlinkassociation "github.com/golingon/terraproviders/aws/5.12.0/networkmanagerlinkassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerLinkAssociation creates a new instance of [NetworkmanagerLinkAssociation].
func NewNetworkmanagerLinkAssociation(name string, args NetworkmanagerLinkAssociationArgs) *NetworkmanagerLinkAssociation {
	return &NetworkmanagerLinkAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerLinkAssociation)(nil)

// NetworkmanagerLinkAssociation represents the Terraform resource aws_networkmanager_link_association.
type NetworkmanagerLinkAssociation struct {
	Name      string
	Args      NetworkmanagerLinkAssociationArgs
	state     *networkmanagerLinkAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerLinkAssociation].
func (nla *NetworkmanagerLinkAssociation) Type() string {
	return "aws_networkmanager_link_association"
}

// LocalName returns the local name for [NetworkmanagerLinkAssociation].
func (nla *NetworkmanagerLinkAssociation) LocalName() string {
	return nla.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerLinkAssociation].
func (nla *NetworkmanagerLinkAssociation) Configuration() interface{} {
	return nla.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerLinkAssociation].
func (nla *NetworkmanagerLinkAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(nla)
}

// Dependencies returns the list of resources [NetworkmanagerLinkAssociation] depends_on.
func (nla *NetworkmanagerLinkAssociation) Dependencies() terra.Dependencies {
	return nla.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerLinkAssociation].
func (nla *NetworkmanagerLinkAssociation) LifecycleManagement() *terra.Lifecycle {
	return nla.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerLinkAssociation].
func (nla *NetworkmanagerLinkAssociation) Attributes() networkmanagerLinkAssociationAttributes {
	return networkmanagerLinkAssociationAttributes{ref: terra.ReferenceResource(nla)}
}

// ImportState imports the given attribute values into [NetworkmanagerLinkAssociation]'s state.
func (nla *NetworkmanagerLinkAssociation) ImportState(av io.Reader) error {
	nla.state = &networkmanagerLinkAssociationState{}
	if err := json.NewDecoder(av).Decode(nla.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nla.Type(), nla.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerLinkAssociation] has state.
func (nla *NetworkmanagerLinkAssociation) State() (*networkmanagerLinkAssociationState, bool) {
	return nla.state, nla.state != nil
}

// StateMust returns the state for [NetworkmanagerLinkAssociation]. Panics if the state is nil.
func (nla *NetworkmanagerLinkAssociation) StateMust() *networkmanagerLinkAssociationState {
	if nla.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nla.Type(), nla.LocalName()))
	}
	return nla.state
}

// NetworkmanagerLinkAssociationArgs contains the configurations for aws_networkmanager_link_association.
type NetworkmanagerLinkAssociationArgs struct {
	// DeviceId: string, required
	DeviceId terra.StringValue `hcl:"device_id,attr" validate:"required"`
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkId: string, required
	LinkId terra.StringValue `hcl:"link_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagerlinkassociation.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerLinkAssociationAttributes struct {
	ref terra.Reference
}

// DeviceId returns a reference to field device_id of aws_networkmanager_link_association.
func (nla networkmanagerLinkAssociationAttributes) DeviceId() terra.StringValue {
	return terra.ReferenceAsString(nla.ref.Append("device_id"))
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_link_association.
func (nla networkmanagerLinkAssociationAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(nla.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_link_association.
func (nla networkmanagerLinkAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nla.ref.Append("id"))
}

// LinkId returns a reference to field link_id of aws_networkmanager_link_association.
func (nla networkmanagerLinkAssociationAttributes) LinkId() terra.StringValue {
	return terra.ReferenceAsString(nla.ref.Append("link_id"))
}

func (nla networkmanagerLinkAssociationAttributes) Timeouts() networkmanagerlinkassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagerlinkassociation.TimeoutsAttributes](nla.ref.Append("timeouts"))
}

type networkmanagerLinkAssociationState struct {
	DeviceId        string                                       `json:"device_id"`
	GlobalNetworkId string                                       `json:"global_network_id"`
	Id              string                                       `json:"id"`
	LinkId          string                                       `json:"link_id"`
	Timeouts        *networkmanagerlinkassociation.TimeoutsState `json:"timeouts"`
}
