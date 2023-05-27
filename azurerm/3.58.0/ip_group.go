// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	ipgroup "github.com/golingon/terraproviders/azurerm/3.58.0/ipgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIpGroup creates a new instance of [IpGroup].
func NewIpGroup(name string, args IpGroupArgs) *IpGroup {
	return &IpGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IpGroup)(nil)

// IpGroup represents the Terraform resource azurerm_ip_group.
type IpGroup struct {
	Name      string
	Args      IpGroupArgs
	state     *ipGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IpGroup].
func (ig *IpGroup) Type() string {
	return "azurerm_ip_group"
}

// LocalName returns the local name for [IpGroup].
func (ig *IpGroup) LocalName() string {
	return ig.Name
}

// Configuration returns the configuration (args) for [IpGroup].
func (ig *IpGroup) Configuration() interface{} {
	return ig.Args
}

// DependOn is used for other resources to depend on [IpGroup].
func (ig *IpGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ig)
}

// Dependencies returns the list of resources [IpGroup] depends_on.
func (ig *IpGroup) Dependencies() terra.Dependencies {
	return ig.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IpGroup].
func (ig *IpGroup) LifecycleManagement() *terra.Lifecycle {
	return ig.Lifecycle
}

// Attributes returns the attributes for [IpGroup].
func (ig *IpGroup) Attributes() ipGroupAttributes {
	return ipGroupAttributes{ref: terra.ReferenceResource(ig)}
}

// ImportState imports the given attribute values into [IpGroup]'s state.
func (ig *IpGroup) ImportState(av io.Reader) error {
	ig.state = &ipGroupState{}
	if err := json.NewDecoder(av).Decode(ig.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ig.Type(), ig.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IpGroup] has state.
func (ig *IpGroup) State() (*ipGroupState, bool) {
	return ig.state, ig.state != nil
}

// StateMust returns the state for [IpGroup]. Panics if the state is nil.
func (ig *IpGroup) StateMust() *ipGroupState {
	if ig.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ig.Type(), ig.LocalName()))
	}
	return ig.state
}

// IpGroupArgs contains the configurations for azurerm_ip_group.
type IpGroupArgs struct {
	// Cidrs: set of string, optional
	Cidrs terra.SetValue[terra.StringValue] `hcl:"cidrs,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *ipgroup.Timeouts `hcl:"timeouts,block"`
}
type ipGroupAttributes struct {
	ref terra.Reference
}

// Cidrs returns a reference to field cidrs of azurerm_ip_group.
func (ig ipGroupAttributes) Cidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ig.ref.Append("cidrs"))
}

// FirewallIds returns a reference to field firewall_ids of azurerm_ip_group.
func (ig ipGroupAttributes) FirewallIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ig.ref.Append("firewall_ids"))
}

// FirewallPolicyIds returns a reference to field firewall_policy_ids of azurerm_ip_group.
func (ig ipGroupAttributes) FirewallPolicyIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ig.ref.Append("firewall_policy_ids"))
}

// Id returns a reference to field id of azurerm_ip_group.
func (ig ipGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_ip_group.
func (ig ipGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_ip_group.
func (ig ipGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_ip_group.
func (ig ipGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_ip_group.
func (ig ipGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ig.ref.Append("tags"))
}

func (ig ipGroupAttributes) Timeouts() ipgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ipgroup.TimeoutsAttributes](ig.ref.Append("timeouts"))
}

type ipGroupState struct {
	Cidrs             []string               `json:"cidrs"`
	FirewallIds       []string               `json:"firewall_ids"`
	FirewallPolicyIds []string               `json:"firewall_policy_ids"`
	Id                string                 `json:"id"`
	Location          string                 `json:"location"`
	Name              string                 `json:"name"`
	ResourceGroupName string                 `json:"resource_group_name"`
	Tags              map[string]string      `json:"tags"`
	Timeouts          *ipgroup.TimeoutsState `json:"timeouts"`
}
