// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	identitystoregroup "github.com/golingon/terraproviders/aws/5.7.0/identitystoregroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentitystoreGroup creates a new instance of [IdentitystoreGroup].
func NewIdentitystoreGroup(name string, args IdentitystoreGroupArgs) *IdentitystoreGroup {
	return &IdentitystoreGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentitystoreGroup)(nil)

// IdentitystoreGroup represents the Terraform resource aws_identitystore_group.
type IdentitystoreGroup struct {
	Name      string
	Args      IdentitystoreGroupArgs
	state     *identitystoreGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentitystoreGroup].
func (ig *IdentitystoreGroup) Type() string {
	return "aws_identitystore_group"
}

// LocalName returns the local name for [IdentitystoreGroup].
func (ig *IdentitystoreGroup) LocalName() string {
	return ig.Name
}

// Configuration returns the configuration (args) for [IdentitystoreGroup].
func (ig *IdentitystoreGroup) Configuration() interface{} {
	return ig.Args
}

// DependOn is used for other resources to depend on [IdentitystoreGroup].
func (ig *IdentitystoreGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ig)
}

// Dependencies returns the list of resources [IdentitystoreGroup] depends_on.
func (ig *IdentitystoreGroup) Dependencies() terra.Dependencies {
	return ig.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentitystoreGroup].
func (ig *IdentitystoreGroup) LifecycleManagement() *terra.Lifecycle {
	return ig.Lifecycle
}

// Attributes returns the attributes for [IdentitystoreGroup].
func (ig *IdentitystoreGroup) Attributes() identitystoreGroupAttributes {
	return identitystoreGroupAttributes{ref: terra.ReferenceResource(ig)}
}

// ImportState imports the given attribute values into [IdentitystoreGroup]'s state.
func (ig *IdentitystoreGroup) ImportState(av io.Reader) error {
	ig.state = &identitystoreGroupState{}
	if err := json.NewDecoder(av).Decode(ig.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ig.Type(), ig.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentitystoreGroup] has state.
func (ig *IdentitystoreGroup) State() (*identitystoreGroupState, bool) {
	return ig.state, ig.state != nil
}

// StateMust returns the state for [IdentitystoreGroup]. Panics if the state is nil.
func (ig *IdentitystoreGroup) StateMust() *identitystoreGroupState {
	if ig.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ig.Type(), ig.LocalName()))
	}
	return ig.state
}

// IdentitystoreGroupArgs contains the configurations for aws_identitystore_group.
type IdentitystoreGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityStoreId: string, required
	IdentityStoreId terra.StringValue `hcl:"identity_store_id,attr" validate:"required"`
	// ExternalIds: min=0
	ExternalIds []identitystoregroup.ExternalIds `hcl:"external_ids,block" validate:"min=0"`
}
type identitystoreGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_identitystore_group.
func (ig identitystoreGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of aws_identitystore_group.
func (ig identitystoreGroupAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("display_name"))
}

// GroupId returns a reference to field group_id of aws_identitystore_group.
func (ig identitystoreGroupAttributes) GroupId() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("group_id"))
}

// Id returns a reference to field id of aws_identitystore_group.
func (ig identitystoreGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("id"))
}

// IdentityStoreId returns a reference to field identity_store_id of aws_identitystore_group.
func (ig identitystoreGroupAttributes) IdentityStoreId() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("identity_store_id"))
}

func (ig identitystoreGroupAttributes) ExternalIds() terra.ListValue[identitystoregroup.ExternalIdsAttributes] {
	return terra.ReferenceAsList[identitystoregroup.ExternalIdsAttributes](ig.ref.Append("external_ids"))
}

type identitystoreGroupState struct {
	Description     string                                `json:"description"`
	DisplayName     string                                `json:"display_name"`
	GroupId         string                                `json:"group_id"`
	Id              string                                `json:"id"`
	IdentityStoreId string                                `json:"identity_store_id"`
	ExternalIds     []identitystoregroup.ExternalIdsState `json:"external_ids"`
}
