// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	managementgroup "github.com/golingon/terraproviders/azurerm/3.49.0/managementgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewManagementGroup(name string, args ManagementGroupArgs) *ManagementGroup {
	return &ManagementGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ManagementGroup)(nil)

type ManagementGroup struct {
	Name  string
	Args  ManagementGroupArgs
	state *managementGroupState
}

func (mg *ManagementGroup) Type() string {
	return "azurerm_management_group"
}

func (mg *ManagementGroup) LocalName() string {
	return mg.Name
}

func (mg *ManagementGroup) Configuration() interface{} {
	return mg.Args
}

func (mg *ManagementGroup) Attributes() managementGroupAttributes {
	return managementGroupAttributes{ref: terra.ReferenceResource(mg)}
}

func (mg *ManagementGroup) ImportState(av io.Reader) error {
	mg.state = &managementGroupState{}
	if err := json.NewDecoder(av).Decode(mg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mg.Type(), mg.LocalName(), err)
	}
	return nil
}

func (mg *ManagementGroup) State() (*managementGroupState, bool) {
	return mg.state, mg.state != nil
}

func (mg *ManagementGroup) StateMust() *managementGroupState {
	if mg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mg.Type(), mg.LocalName()))
	}
	return mg.state
}

func (mg *ManagementGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(mg)
}

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
	// DependsOn contains resources that ManagementGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type managementGroupAttributes struct {
	ref terra.Reference
}

func (mg managementGroupAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(mg.ref.Append("display_name"))
}

func (mg managementGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mg.ref.Append("id"))
}

func (mg managementGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(mg.ref.Append("name"))
}

func (mg managementGroupAttributes) ParentManagementGroupId() terra.StringValue {
	return terra.ReferenceString(mg.ref.Append("parent_management_group_id"))
}

func (mg managementGroupAttributes) SubscriptionIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](mg.ref.Append("subscription_ids"))
}

func (mg managementGroupAttributes) Timeouts() managementgroup.TimeoutsAttributes {
	return terra.ReferenceSingle[managementgroup.TimeoutsAttributes](mg.ref.Append("timeouts"))
}

type managementGroupState struct {
	DisplayName             string                         `json:"display_name"`
	Id                      string                         `json:"id"`
	Name                    string                         `json:"name"`
	ParentManagementGroupId string                         `json:"parent_management_group_id"`
	SubscriptionIds         []string                       `json:"subscription_ids"`
	Timeouts                *managementgroup.TimeoutsState `json:"timeouts"`
}
