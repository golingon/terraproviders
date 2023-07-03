// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	inspector2memberassociation "github.com/golingon/terraproviders/aws/5.6.2/inspector2memberassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewInspector2MemberAssociation creates a new instance of [Inspector2MemberAssociation].
func NewInspector2MemberAssociation(name string, args Inspector2MemberAssociationArgs) *Inspector2MemberAssociation {
	return &Inspector2MemberAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Inspector2MemberAssociation)(nil)

// Inspector2MemberAssociation represents the Terraform resource aws_inspector2_member_association.
type Inspector2MemberAssociation struct {
	Name      string
	Args      Inspector2MemberAssociationArgs
	state     *inspector2MemberAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Inspector2MemberAssociation].
func (ima *Inspector2MemberAssociation) Type() string {
	return "aws_inspector2_member_association"
}

// LocalName returns the local name for [Inspector2MemberAssociation].
func (ima *Inspector2MemberAssociation) LocalName() string {
	return ima.Name
}

// Configuration returns the configuration (args) for [Inspector2MemberAssociation].
func (ima *Inspector2MemberAssociation) Configuration() interface{} {
	return ima.Args
}

// DependOn is used for other resources to depend on [Inspector2MemberAssociation].
func (ima *Inspector2MemberAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(ima)
}

// Dependencies returns the list of resources [Inspector2MemberAssociation] depends_on.
func (ima *Inspector2MemberAssociation) Dependencies() terra.Dependencies {
	return ima.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Inspector2MemberAssociation].
func (ima *Inspector2MemberAssociation) LifecycleManagement() *terra.Lifecycle {
	return ima.Lifecycle
}

// Attributes returns the attributes for [Inspector2MemberAssociation].
func (ima *Inspector2MemberAssociation) Attributes() inspector2MemberAssociationAttributes {
	return inspector2MemberAssociationAttributes{ref: terra.ReferenceResource(ima)}
}

// ImportState imports the given attribute values into [Inspector2MemberAssociation]'s state.
func (ima *Inspector2MemberAssociation) ImportState(av io.Reader) error {
	ima.state = &inspector2MemberAssociationState{}
	if err := json.NewDecoder(av).Decode(ima.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ima.Type(), ima.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Inspector2MemberAssociation] has state.
func (ima *Inspector2MemberAssociation) State() (*inspector2MemberAssociationState, bool) {
	return ima.state, ima.state != nil
}

// StateMust returns the state for [Inspector2MemberAssociation]. Panics if the state is nil.
func (ima *Inspector2MemberAssociation) StateMust() *inspector2MemberAssociationState {
	if ima.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ima.Type(), ima.LocalName()))
	}
	return ima.state
}

// Inspector2MemberAssociationArgs contains the configurations for aws_inspector2_member_association.
type Inspector2MemberAssociationArgs struct {
	// AccountId: string, required
	AccountId terra.StringValue `hcl:"account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *inspector2memberassociation.Timeouts `hcl:"timeouts,block"`
}
type inspector2MemberAssociationAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_inspector2_member_association.
func (ima inspector2MemberAssociationAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(ima.ref.Append("account_id"))
}

// DelegatedAdminAccountId returns a reference to field delegated_admin_account_id of aws_inspector2_member_association.
func (ima inspector2MemberAssociationAttributes) DelegatedAdminAccountId() terra.StringValue {
	return terra.ReferenceAsString(ima.ref.Append("delegated_admin_account_id"))
}

// Id returns a reference to field id of aws_inspector2_member_association.
func (ima inspector2MemberAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ima.ref.Append("id"))
}

// RelationshipStatus returns a reference to field relationship_status of aws_inspector2_member_association.
func (ima inspector2MemberAssociationAttributes) RelationshipStatus() terra.StringValue {
	return terra.ReferenceAsString(ima.ref.Append("relationship_status"))
}

// UpdatedAt returns a reference to field updated_at of aws_inspector2_member_association.
func (ima inspector2MemberAssociationAttributes) UpdatedAt() terra.StringValue {
	return terra.ReferenceAsString(ima.ref.Append("updated_at"))
}

func (ima inspector2MemberAssociationAttributes) Timeouts() inspector2memberassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[inspector2memberassociation.TimeoutsAttributes](ima.ref.Append("timeouts"))
}

type inspector2MemberAssociationState struct {
	AccountId               string                                     `json:"account_id"`
	DelegatedAdminAccountId string                                     `json:"delegated_admin_account_id"`
	Id                      string                                     `json:"id"`
	RelationshipStatus      string                                     `json:"relationship_status"`
	UpdatedAt               string                                     `json:"updated_at"`
	Timeouts                *inspector2memberassociation.TimeoutsState `json:"timeouts"`
}
