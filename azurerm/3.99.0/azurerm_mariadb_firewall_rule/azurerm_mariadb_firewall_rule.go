// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_mariadb_firewall_rule

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_mariadb_firewall_rule.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermMariadbFirewallRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (amfr *Resource) Type() string {
	return "azurerm_mariadb_firewall_rule"
}

// LocalName returns the local name for [Resource].
func (amfr *Resource) LocalName() string {
	return amfr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (amfr *Resource) Configuration() interface{} {
	return amfr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (amfr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(amfr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (amfr *Resource) Dependencies() terra.Dependencies {
	return amfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (amfr *Resource) LifecycleManagement() *terra.Lifecycle {
	return amfr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (amfr *Resource) Attributes() azurermMariadbFirewallRuleAttributes {
	return azurermMariadbFirewallRuleAttributes{ref: terra.ReferenceResource(amfr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (amfr *Resource) ImportState(state io.Reader) error {
	amfr.state = &azurermMariadbFirewallRuleState{}
	if err := json.NewDecoder(state).Decode(amfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amfr.Type(), amfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (amfr *Resource) State() (*azurermMariadbFirewallRuleState, bool) {
	return amfr.state, amfr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (amfr *Resource) StateMust() *azurermMariadbFirewallRuleState {
	if amfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amfr.Type(), amfr.LocalName()))
	}
	return amfr.state
}

// Args contains the configurations for azurerm_mariadb_firewall_rule.
type Args struct {
	// EndIpAddress: string, required
	EndIpAddress terra.StringValue `hcl:"end_ip_address,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// StartIpAddress: string, required
	StartIpAddress terra.StringValue `hcl:"start_ip_address,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermMariadbFirewallRuleAttributes struct {
	ref terra.Reference
}

// EndIpAddress returns a reference to field end_ip_address of azurerm_mariadb_firewall_rule.
func (amfr azurermMariadbFirewallRuleAttributes) EndIpAddress() terra.StringValue {
	return terra.ReferenceAsString(amfr.ref.Append("end_ip_address"))
}

// Id returns a reference to field id of azurerm_mariadb_firewall_rule.
func (amfr azurermMariadbFirewallRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amfr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mariadb_firewall_rule.
func (amfr azurermMariadbFirewallRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amfr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mariadb_firewall_rule.
func (amfr azurermMariadbFirewallRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amfr.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mariadb_firewall_rule.
func (amfr azurermMariadbFirewallRuleAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(amfr.ref.Append("server_name"))
}

// StartIpAddress returns a reference to field start_ip_address of azurerm_mariadb_firewall_rule.
func (amfr azurermMariadbFirewallRuleAttributes) StartIpAddress() terra.StringValue {
	return terra.ReferenceAsString(amfr.ref.Append("start_ip_address"))
}

func (amfr azurermMariadbFirewallRuleAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](amfr.ref.Append("timeouts"))
}

type azurermMariadbFirewallRuleState struct {
	EndIpAddress      string         `json:"end_ip_address"`
	Id                string         `json:"id"`
	Name              string         `json:"name"`
	ResourceGroupName string         `json:"resource_group_name"`
	ServerName        string         `json:"server_name"`
	StartIpAddress    string         `json:"start_ip_address"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
