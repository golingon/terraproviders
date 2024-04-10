// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	managementgroupsubscriptionassociation "github.com/golingon/terraproviders/azurerm/3.98.0/managementgroupsubscriptionassociation"
	"io"
)

// NewManagementGroupSubscriptionAssociation creates a new instance of [ManagementGroupSubscriptionAssociation].
func NewManagementGroupSubscriptionAssociation(name string, args ManagementGroupSubscriptionAssociationArgs) *ManagementGroupSubscriptionAssociation {
	return &ManagementGroupSubscriptionAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ManagementGroupSubscriptionAssociation)(nil)

// ManagementGroupSubscriptionAssociation represents the Terraform resource azurerm_management_group_subscription_association.
type ManagementGroupSubscriptionAssociation struct {
	Name      string
	Args      ManagementGroupSubscriptionAssociationArgs
	state     *managementGroupSubscriptionAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ManagementGroupSubscriptionAssociation].
func (mgsa *ManagementGroupSubscriptionAssociation) Type() string {
	return "azurerm_management_group_subscription_association"
}

// LocalName returns the local name for [ManagementGroupSubscriptionAssociation].
func (mgsa *ManagementGroupSubscriptionAssociation) LocalName() string {
	return mgsa.Name
}

// Configuration returns the configuration (args) for [ManagementGroupSubscriptionAssociation].
func (mgsa *ManagementGroupSubscriptionAssociation) Configuration() interface{} {
	return mgsa.Args
}

// DependOn is used for other resources to depend on [ManagementGroupSubscriptionAssociation].
func (mgsa *ManagementGroupSubscriptionAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(mgsa)
}

// Dependencies returns the list of resources [ManagementGroupSubscriptionAssociation] depends_on.
func (mgsa *ManagementGroupSubscriptionAssociation) Dependencies() terra.Dependencies {
	return mgsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ManagementGroupSubscriptionAssociation].
func (mgsa *ManagementGroupSubscriptionAssociation) LifecycleManagement() *terra.Lifecycle {
	return mgsa.Lifecycle
}

// Attributes returns the attributes for [ManagementGroupSubscriptionAssociation].
func (mgsa *ManagementGroupSubscriptionAssociation) Attributes() managementGroupSubscriptionAssociationAttributes {
	return managementGroupSubscriptionAssociationAttributes{ref: terra.ReferenceResource(mgsa)}
}

// ImportState imports the given attribute values into [ManagementGroupSubscriptionAssociation]'s state.
func (mgsa *ManagementGroupSubscriptionAssociation) ImportState(av io.Reader) error {
	mgsa.state = &managementGroupSubscriptionAssociationState{}
	if err := json.NewDecoder(av).Decode(mgsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mgsa.Type(), mgsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ManagementGroupSubscriptionAssociation] has state.
func (mgsa *ManagementGroupSubscriptionAssociation) State() (*managementGroupSubscriptionAssociationState, bool) {
	return mgsa.state, mgsa.state != nil
}

// StateMust returns the state for [ManagementGroupSubscriptionAssociation]. Panics if the state is nil.
func (mgsa *ManagementGroupSubscriptionAssociation) StateMust() *managementGroupSubscriptionAssociationState {
	if mgsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mgsa.Type(), mgsa.LocalName()))
	}
	return mgsa.state
}

// ManagementGroupSubscriptionAssociationArgs contains the configurations for azurerm_management_group_subscription_association.
type ManagementGroupSubscriptionAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagementGroupId: string, required
	ManagementGroupId terra.StringValue `hcl:"management_group_id,attr" validate:"required"`
	// SubscriptionId: string, required
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *managementgroupsubscriptionassociation.Timeouts `hcl:"timeouts,block"`
}
type managementGroupSubscriptionAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_management_group_subscription_association.
func (mgsa managementGroupSubscriptionAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mgsa.ref.Append("id"))
}

// ManagementGroupId returns a reference to field management_group_id of azurerm_management_group_subscription_association.
func (mgsa managementGroupSubscriptionAssociationAttributes) ManagementGroupId() terra.StringValue {
	return terra.ReferenceAsString(mgsa.ref.Append("management_group_id"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_management_group_subscription_association.
func (mgsa managementGroupSubscriptionAssociationAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(mgsa.ref.Append("subscription_id"))
}

func (mgsa managementGroupSubscriptionAssociationAttributes) Timeouts() managementgroupsubscriptionassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[managementgroupsubscriptionassociation.TimeoutsAttributes](mgsa.ref.Append("timeouts"))
}

type managementGroupSubscriptionAssociationState struct {
	Id                string                                                `json:"id"`
	ManagementGroupId string                                                `json:"management_group_id"`
	SubscriptionId    string                                                `json:"subscription_id"`
	Timeouts          *managementgroupsubscriptionassociation.TimeoutsState `json:"timeouts"`
}
