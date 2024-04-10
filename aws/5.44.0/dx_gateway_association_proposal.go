// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewDxGatewayAssociationProposal creates a new instance of [DxGatewayAssociationProposal].
func NewDxGatewayAssociationProposal(name string, args DxGatewayAssociationProposalArgs) *DxGatewayAssociationProposal {
	return &DxGatewayAssociationProposal{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DxGatewayAssociationProposal)(nil)

// DxGatewayAssociationProposal represents the Terraform resource aws_dx_gateway_association_proposal.
type DxGatewayAssociationProposal struct {
	Name      string
	Args      DxGatewayAssociationProposalArgs
	state     *dxGatewayAssociationProposalState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DxGatewayAssociationProposal].
func (dgap *DxGatewayAssociationProposal) Type() string {
	return "aws_dx_gateway_association_proposal"
}

// LocalName returns the local name for [DxGatewayAssociationProposal].
func (dgap *DxGatewayAssociationProposal) LocalName() string {
	return dgap.Name
}

// Configuration returns the configuration (args) for [DxGatewayAssociationProposal].
func (dgap *DxGatewayAssociationProposal) Configuration() interface{} {
	return dgap.Args
}

// DependOn is used for other resources to depend on [DxGatewayAssociationProposal].
func (dgap *DxGatewayAssociationProposal) DependOn() terra.Reference {
	return terra.ReferenceResource(dgap)
}

// Dependencies returns the list of resources [DxGatewayAssociationProposal] depends_on.
func (dgap *DxGatewayAssociationProposal) Dependencies() terra.Dependencies {
	return dgap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DxGatewayAssociationProposal].
func (dgap *DxGatewayAssociationProposal) LifecycleManagement() *terra.Lifecycle {
	return dgap.Lifecycle
}

// Attributes returns the attributes for [DxGatewayAssociationProposal].
func (dgap *DxGatewayAssociationProposal) Attributes() dxGatewayAssociationProposalAttributes {
	return dxGatewayAssociationProposalAttributes{ref: terra.ReferenceResource(dgap)}
}

// ImportState imports the given attribute values into [DxGatewayAssociationProposal]'s state.
func (dgap *DxGatewayAssociationProposal) ImportState(av io.Reader) error {
	dgap.state = &dxGatewayAssociationProposalState{}
	if err := json.NewDecoder(av).Decode(dgap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dgap.Type(), dgap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DxGatewayAssociationProposal] has state.
func (dgap *DxGatewayAssociationProposal) State() (*dxGatewayAssociationProposalState, bool) {
	return dgap.state, dgap.state != nil
}

// StateMust returns the state for [DxGatewayAssociationProposal]. Panics if the state is nil.
func (dgap *DxGatewayAssociationProposal) StateMust() *dxGatewayAssociationProposalState {
	if dgap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dgap.Type(), dgap.LocalName()))
	}
	return dgap.state
}

// DxGatewayAssociationProposalArgs contains the configurations for aws_dx_gateway_association_proposal.
type DxGatewayAssociationProposalArgs struct {
	// AllowedPrefixes: set of string, optional
	AllowedPrefixes terra.SetValue[terra.StringValue] `hcl:"allowed_prefixes,attr"`
	// AssociatedGatewayId: string, required
	AssociatedGatewayId terra.StringValue `hcl:"associated_gateway_id,attr" validate:"required"`
	// DxGatewayId: string, required
	DxGatewayId terra.StringValue `hcl:"dx_gateway_id,attr" validate:"required"`
	// DxGatewayOwnerAccountId: string, required
	DxGatewayOwnerAccountId terra.StringValue `hcl:"dx_gateway_owner_account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dxGatewayAssociationProposalAttributes struct {
	ref terra.Reference
}

// AllowedPrefixes returns a reference to field allowed_prefixes of aws_dx_gateway_association_proposal.
func (dgap dxGatewayAssociationProposalAttributes) AllowedPrefixes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dgap.ref.Append("allowed_prefixes"))
}

// AssociatedGatewayId returns a reference to field associated_gateway_id of aws_dx_gateway_association_proposal.
func (dgap dxGatewayAssociationProposalAttributes) AssociatedGatewayId() terra.StringValue {
	return terra.ReferenceAsString(dgap.ref.Append("associated_gateway_id"))
}

// AssociatedGatewayOwnerAccountId returns a reference to field associated_gateway_owner_account_id of aws_dx_gateway_association_proposal.
func (dgap dxGatewayAssociationProposalAttributes) AssociatedGatewayOwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(dgap.ref.Append("associated_gateway_owner_account_id"))
}

// AssociatedGatewayType returns a reference to field associated_gateway_type of aws_dx_gateway_association_proposal.
func (dgap dxGatewayAssociationProposalAttributes) AssociatedGatewayType() terra.StringValue {
	return terra.ReferenceAsString(dgap.ref.Append("associated_gateway_type"))
}

// DxGatewayId returns a reference to field dx_gateway_id of aws_dx_gateway_association_proposal.
func (dgap dxGatewayAssociationProposalAttributes) DxGatewayId() terra.StringValue {
	return terra.ReferenceAsString(dgap.ref.Append("dx_gateway_id"))
}

// DxGatewayOwnerAccountId returns a reference to field dx_gateway_owner_account_id of aws_dx_gateway_association_proposal.
func (dgap dxGatewayAssociationProposalAttributes) DxGatewayOwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(dgap.ref.Append("dx_gateway_owner_account_id"))
}

// Id returns a reference to field id of aws_dx_gateway_association_proposal.
func (dgap dxGatewayAssociationProposalAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dgap.ref.Append("id"))
}

type dxGatewayAssociationProposalState struct {
	AllowedPrefixes                 []string `json:"allowed_prefixes"`
	AssociatedGatewayId             string   `json:"associated_gateway_id"`
	AssociatedGatewayOwnerAccountId string   `json:"associated_gateway_owner_account_id"`
	AssociatedGatewayType           string   `json:"associated_gateway_type"`
	DxGatewayId                     string   `json:"dx_gateway_id"`
	DxGatewayOwnerAccountId         string   `json:"dx_gateway_owner_account_id"`
	Id                              string   `json:"id"`
}
