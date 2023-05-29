// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMacieMemberAccountAssociation creates a new instance of [MacieMemberAccountAssociation].
func NewMacieMemberAccountAssociation(name string, args MacieMemberAccountAssociationArgs) *MacieMemberAccountAssociation {
	return &MacieMemberAccountAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MacieMemberAccountAssociation)(nil)

// MacieMemberAccountAssociation represents the Terraform resource aws_macie_member_account_association.
type MacieMemberAccountAssociation struct {
	Name      string
	Args      MacieMemberAccountAssociationArgs
	state     *macieMemberAccountAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MacieMemberAccountAssociation].
func (mmaa *MacieMemberAccountAssociation) Type() string {
	return "aws_macie_member_account_association"
}

// LocalName returns the local name for [MacieMemberAccountAssociation].
func (mmaa *MacieMemberAccountAssociation) LocalName() string {
	return mmaa.Name
}

// Configuration returns the configuration (args) for [MacieMemberAccountAssociation].
func (mmaa *MacieMemberAccountAssociation) Configuration() interface{} {
	return mmaa.Args
}

// DependOn is used for other resources to depend on [MacieMemberAccountAssociation].
func (mmaa *MacieMemberAccountAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(mmaa)
}

// Dependencies returns the list of resources [MacieMemberAccountAssociation] depends_on.
func (mmaa *MacieMemberAccountAssociation) Dependencies() terra.Dependencies {
	return mmaa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MacieMemberAccountAssociation].
func (mmaa *MacieMemberAccountAssociation) LifecycleManagement() *terra.Lifecycle {
	return mmaa.Lifecycle
}

// Attributes returns the attributes for [MacieMemberAccountAssociation].
func (mmaa *MacieMemberAccountAssociation) Attributes() macieMemberAccountAssociationAttributes {
	return macieMemberAccountAssociationAttributes{ref: terra.ReferenceResource(mmaa)}
}

// ImportState imports the given attribute values into [MacieMemberAccountAssociation]'s state.
func (mmaa *MacieMemberAccountAssociation) ImportState(av io.Reader) error {
	mmaa.state = &macieMemberAccountAssociationState{}
	if err := json.NewDecoder(av).Decode(mmaa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mmaa.Type(), mmaa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MacieMemberAccountAssociation] has state.
func (mmaa *MacieMemberAccountAssociation) State() (*macieMemberAccountAssociationState, bool) {
	return mmaa.state, mmaa.state != nil
}

// StateMust returns the state for [MacieMemberAccountAssociation]. Panics if the state is nil.
func (mmaa *MacieMemberAccountAssociation) StateMust() *macieMemberAccountAssociationState {
	if mmaa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mmaa.Type(), mmaa.LocalName()))
	}
	return mmaa.state
}

// MacieMemberAccountAssociationArgs contains the configurations for aws_macie_member_account_association.
type MacieMemberAccountAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MemberAccountId: string, required
	MemberAccountId terra.StringValue `hcl:"member_account_id,attr" validate:"required"`
}
type macieMemberAccountAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_macie_member_account_association.
func (mmaa macieMemberAccountAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mmaa.ref.Append("id"))
}

// MemberAccountId returns a reference to field member_account_id of aws_macie_member_account_association.
func (mmaa macieMemberAccountAssociationAttributes) MemberAccountId() terra.StringValue {
	return terra.ReferenceAsString(mmaa.ref.Append("member_account_id"))
}

type macieMemberAccountAssociationState struct {
	Id              string `json:"id"`
	MemberAccountId string `json:"member_account_id"`
}