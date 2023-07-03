// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	managementgroup "github.com/golingon/terraproviders/azurerm/3.63.0/managementgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewManagementGroup creates a new instance of [ManagementGroup].
func NewManagementGroup(name string, args ManagementGroupArgs) *ManagementGroup {
	return &ManagementGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ManagementGroup)(nil)

// ManagementGroup represents the Terraform resource azurerm_management_group.
type ManagementGroup struct {
	Name      string
	Args      ManagementGroupArgs
	state     *managementGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ManagementGroup].
func (mg *ManagementGroup) Type() string {
	return "azurerm_management_group"
}

// LocalName returns the local name for [ManagementGroup].
func (mg *ManagementGroup) LocalName() string {
	return mg.Name
}

// Configuration returns the configuration (args) for [ManagementGroup].
func (mg *ManagementGroup) Configuration() interface{} {
	return mg.Args
}

// DependOn is used for other resources to depend on [ManagementGroup].
func (mg *ManagementGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(mg)
}

// Dependencies returns the list of resources [ManagementGroup] depends_on.
func (mg *ManagementGroup) Dependencies() terra.Dependencies {
	return mg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ManagementGroup].
func (mg *ManagementGroup) LifecycleManagement() *terra.Lifecycle {
	return mg.Lifecycle
}

// Attributes returns the attributes for [ManagementGroup].
func (mg *ManagementGroup) Attributes() managementGroupAttributes {
	return managementGroupAttributes{ref: terra.ReferenceResource(mg)}
}

// ImportState imports the given attribute values into [ManagementGroup]'s state.
func (mg *ManagementGroup) ImportState(av io.Reader) error {
	mg.state = &managementGroupState{}
	if err := json.NewDecoder(av).Decode(mg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mg.Type(), mg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ManagementGroup] has state.
func (mg *ManagementGroup) State() (*managementGroupState, bool) {
	return mg.state, mg.state != nil
}

// StateMust returns the state for [ManagementGroup]. Panics if the state is nil.
func (mg *ManagementGroup) StateMust() *managementGroupState {
	if mg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mg.Type(), mg.LocalName()))
	}
	return mg.state
}

// ManagementGroupArgs contains the configurations for azurerm_management_group.
type ManagementGroupArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ParentManagementGroupId: string, optional
	ParentManagementGroupId terra.StringValue `hcl:"parent_management_group_id,attr"`
	// SubscriptionIds: set of string, optional
	SubscriptionIds terra.SetValue[terra.StringValue] `hcl:"subscription_ids,attr"`
	// Timeouts: optional
	Timeouts *managementgroup.Timeouts `hcl:"timeouts,block"`
}
type managementGroupAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azurerm_management_group.
func (mg managementGroupAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_management_group.
func (mg managementGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_management_group.
func (mg managementGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("name"))
}

// ParentManagementGroupId returns a reference to field parent_management_group_id of azurerm_management_group.
func (mg managementGroupAttributes) ParentManagementGroupId() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("parent_management_group_id"))
}

// SubscriptionIds returns a reference to field subscription_ids of azurerm_management_group.
func (mg managementGroupAttributes) SubscriptionIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mg.ref.Append("subscription_ids"))
}

func (mg managementGroupAttributes) Timeouts() managementgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[managementgroup.TimeoutsAttributes](mg.ref.Append("timeouts"))
}

type managementGroupState struct {
	DisplayName             string                         `json:"display_name"`
	Id                      string                         `json:"id"`
	Name                    string                         `json:"name"`
	ParentManagementGroupId string                         `json:"parent_management_group_id"`
	SubscriptionIds         []string                       `json:"subscription_ids"`
	Timeouts                *managementgroup.TimeoutsState `json:"timeouts"`
}
