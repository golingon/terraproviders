// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dxgatewayassociation "github.com/golingon/terraproviders/aws/5.8.0/dxgatewayassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDxGatewayAssociation creates a new instance of [DxGatewayAssociation].
func NewDxGatewayAssociation(name string, args DxGatewayAssociationArgs) *DxGatewayAssociation {
	return &DxGatewayAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DxGatewayAssociation)(nil)

// DxGatewayAssociation represents the Terraform resource aws_dx_gateway_association.
type DxGatewayAssociation struct {
	Name      string
	Args      DxGatewayAssociationArgs
	state     *dxGatewayAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DxGatewayAssociation].
func (dga *DxGatewayAssociation) Type() string {
	return "aws_dx_gateway_association"
}

// LocalName returns the local name for [DxGatewayAssociation].
func (dga *DxGatewayAssociation) LocalName() string {
	return dga.Name
}

// Configuration returns the configuration (args) for [DxGatewayAssociation].
func (dga *DxGatewayAssociation) Configuration() interface{} {
	return dga.Args
}

// DependOn is used for other resources to depend on [DxGatewayAssociation].
func (dga *DxGatewayAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(dga)
}

// Dependencies returns the list of resources [DxGatewayAssociation] depends_on.
func (dga *DxGatewayAssociation) Dependencies() terra.Dependencies {
	return dga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DxGatewayAssociation].
func (dga *DxGatewayAssociation) LifecycleManagement() *terra.Lifecycle {
	return dga.Lifecycle
}

// Attributes returns the attributes for [DxGatewayAssociation].
func (dga *DxGatewayAssociation) Attributes() dxGatewayAssociationAttributes {
	return dxGatewayAssociationAttributes{ref: terra.ReferenceResource(dga)}
}

// ImportState imports the given attribute values into [DxGatewayAssociation]'s state.
func (dga *DxGatewayAssociation) ImportState(av io.Reader) error {
	dga.state = &dxGatewayAssociationState{}
	if err := json.NewDecoder(av).Decode(dga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dga.Type(), dga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DxGatewayAssociation] has state.
func (dga *DxGatewayAssociation) State() (*dxGatewayAssociationState, bool) {
	return dga.state, dga.state != nil
}

// StateMust returns the state for [DxGatewayAssociation]. Panics if the state is nil.
func (dga *DxGatewayAssociation) StateMust() *dxGatewayAssociationState {
	if dga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dga.Type(), dga.LocalName()))
	}
	return dga.state
}

// DxGatewayAssociationArgs contains the configurations for aws_dx_gateway_association.
type DxGatewayAssociationArgs struct {
	// AllowedPrefixes: set of string, optional
	AllowedPrefixes terra.SetValue[terra.StringValue] `hcl:"allowed_prefixes,attr"`
	// AssociatedGatewayId: string, optional
	AssociatedGatewayId terra.StringValue `hcl:"associated_gateway_id,attr"`
	// AssociatedGatewayOwnerAccountId: string, optional
	AssociatedGatewayOwnerAccountId terra.StringValue `hcl:"associated_gateway_owner_account_id,attr"`
	// DxGatewayId: string, required
	DxGatewayId terra.StringValue `hcl:"dx_gateway_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProposalId: string, optional
	ProposalId terra.StringValue `hcl:"proposal_id,attr"`
	// VpnGatewayId: string, optional
	VpnGatewayId terra.StringValue `hcl:"vpn_gateway_id,attr"`
	// Timeouts: optional
	Timeouts *dxgatewayassociation.Timeouts `hcl:"timeouts,block"`
}
type dxGatewayAssociationAttributes struct {
	ref terra.Reference
}

// AllowedPrefixes returns a reference to field allowed_prefixes of aws_dx_gateway_association.
func (dga dxGatewayAssociationAttributes) AllowedPrefixes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dga.ref.Append("allowed_prefixes"))
}

// AssociatedGatewayId returns a reference to field associated_gateway_id of aws_dx_gateway_association.
func (dga dxGatewayAssociationAttributes) AssociatedGatewayId() terra.StringValue {
	return terra.ReferenceAsString(dga.ref.Append("associated_gateway_id"))
}

// AssociatedGatewayOwnerAccountId returns a reference to field associated_gateway_owner_account_id of aws_dx_gateway_association.
func (dga dxGatewayAssociationAttributes) AssociatedGatewayOwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(dga.ref.Append("associated_gateway_owner_account_id"))
}

// AssociatedGatewayType returns a reference to field associated_gateway_type of aws_dx_gateway_association.
func (dga dxGatewayAssociationAttributes) AssociatedGatewayType() terra.StringValue {
	return terra.ReferenceAsString(dga.ref.Append("associated_gateway_type"))
}

// DxGatewayAssociationId returns a reference to field dx_gateway_association_id of aws_dx_gateway_association.
func (dga dxGatewayAssociationAttributes) DxGatewayAssociationId() terra.StringValue {
	return terra.ReferenceAsString(dga.ref.Append("dx_gateway_association_id"))
}

// DxGatewayId returns a reference to field dx_gateway_id of aws_dx_gateway_association.
func (dga dxGatewayAssociationAttributes) DxGatewayId() terra.StringValue {
	return terra.ReferenceAsString(dga.ref.Append("dx_gateway_id"))
}

// DxGatewayOwnerAccountId returns a reference to field dx_gateway_owner_account_id of aws_dx_gateway_association.
func (dga dxGatewayAssociationAttributes) DxGatewayOwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(dga.ref.Append("dx_gateway_owner_account_id"))
}

// Id returns a reference to field id of aws_dx_gateway_association.
func (dga dxGatewayAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dga.ref.Append("id"))
}

// ProposalId returns a reference to field proposal_id of aws_dx_gateway_association.
func (dga dxGatewayAssociationAttributes) ProposalId() terra.StringValue {
	return terra.ReferenceAsString(dga.ref.Append("proposal_id"))
}

// VpnGatewayId returns a reference to field vpn_gateway_id of aws_dx_gateway_association.
func (dga dxGatewayAssociationAttributes) VpnGatewayId() terra.StringValue {
	return terra.ReferenceAsString(dga.ref.Append("vpn_gateway_id"))
}

func (dga dxGatewayAssociationAttributes) Timeouts() dxgatewayassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dxgatewayassociation.TimeoutsAttributes](dga.ref.Append("timeouts"))
}

type dxGatewayAssociationState struct {
	AllowedPrefixes                 []string                            `json:"allowed_prefixes"`
	AssociatedGatewayId             string                              `json:"associated_gateway_id"`
	AssociatedGatewayOwnerAccountId string                              `json:"associated_gateway_owner_account_id"`
	AssociatedGatewayType           string                              `json:"associated_gateway_type"`
	DxGatewayAssociationId          string                              `json:"dx_gateway_association_id"`
	DxGatewayId                     string                              `json:"dx_gateway_id"`
	DxGatewayOwnerAccountId         string                              `json:"dx_gateway_owner_account_id"`
	Id                              string                              `json:"id"`
	ProposalId                      string                              `json:"proposal_id"`
	VpnGatewayId                    string                              `json:"vpn_gateway_id"`
	Timeouts                        *dxgatewayassociation.TimeoutsState `json:"timeouts"`
}
