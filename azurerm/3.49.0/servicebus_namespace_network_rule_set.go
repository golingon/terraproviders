// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	servicebusnamespacenetworkruleset "github.com/golingon/terraproviders/azurerm/3.49.0/servicebusnamespacenetworkruleset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicebusNamespaceNetworkRuleSet creates a new instance of [ServicebusNamespaceNetworkRuleSet].
func NewServicebusNamespaceNetworkRuleSet(name string, args ServicebusNamespaceNetworkRuleSetArgs) *ServicebusNamespaceNetworkRuleSet {
	return &ServicebusNamespaceNetworkRuleSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicebusNamespaceNetworkRuleSet)(nil)

// ServicebusNamespaceNetworkRuleSet represents the Terraform resource azurerm_servicebus_namespace_network_rule_set.
type ServicebusNamespaceNetworkRuleSet struct {
	Name      string
	Args      ServicebusNamespaceNetworkRuleSetArgs
	state     *servicebusNamespaceNetworkRuleSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicebusNamespaceNetworkRuleSet].
func (snnrs *ServicebusNamespaceNetworkRuleSet) Type() string {
	return "azurerm_servicebus_namespace_network_rule_set"
}

// LocalName returns the local name for [ServicebusNamespaceNetworkRuleSet].
func (snnrs *ServicebusNamespaceNetworkRuleSet) LocalName() string {
	return snnrs.Name
}

// Configuration returns the configuration (args) for [ServicebusNamespaceNetworkRuleSet].
func (snnrs *ServicebusNamespaceNetworkRuleSet) Configuration() interface{} {
	return snnrs.Args
}

// DependOn is used for other resources to depend on [ServicebusNamespaceNetworkRuleSet].
func (snnrs *ServicebusNamespaceNetworkRuleSet) DependOn() terra.Reference {
	return terra.ReferenceResource(snnrs)
}

// Dependencies returns the list of resources [ServicebusNamespaceNetworkRuleSet] depends_on.
func (snnrs *ServicebusNamespaceNetworkRuleSet) Dependencies() terra.Dependencies {
	return snnrs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicebusNamespaceNetworkRuleSet].
func (snnrs *ServicebusNamespaceNetworkRuleSet) LifecycleManagement() *terra.Lifecycle {
	return snnrs.Lifecycle
}

// Attributes returns the attributes for [ServicebusNamespaceNetworkRuleSet].
func (snnrs *ServicebusNamespaceNetworkRuleSet) Attributes() servicebusNamespaceNetworkRuleSetAttributes {
	return servicebusNamespaceNetworkRuleSetAttributes{ref: terra.ReferenceResource(snnrs)}
}

// ImportState imports the given attribute values into [ServicebusNamespaceNetworkRuleSet]'s state.
func (snnrs *ServicebusNamespaceNetworkRuleSet) ImportState(av io.Reader) error {
	snnrs.state = &servicebusNamespaceNetworkRuleSetState{}
	if err := json.NewDecoder(av).Decode(snnrs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", snnrs.Type(), snnrs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicebusNamespaceNetworkRuleSet] has state.
func (snnrs *ServicebusNamespaceNetworkRuleSet) State() (*servicebusNamespaceNetworkRuleSetState, bool) {
	return snnrs.state, snnrs.state != nil
}

// StateMust returns the state for [ServicebusNamespaceNetworkRuleSet]. Panics if the state is nil.
func (snnrs *ServicebusNamespaceNetworkRuleSet) StateMust() *servicebusNamespaceNetworkRuleSetState {
	if snnrs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", snnrs.Type(), snnrs.LocalName()))
	}
	return snnrs.state
}

// ServicebusNamespaceNetworkRuleSetArgs contains the configurations for azurerm_servicebus_namespace_network_rule_set.
type ServicebusNamespaceNetworkRuleSetArgs struct {
	// DefaultAction: string, optional
	DefaultAction terra.StringValue `hcl:"default_action,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpRules: set of string, optional
	IpRules terra.SetValue[terra.StringValue] `hcl:"ip_rules,attr"`
	// NamespaceId: string, required
	NamespaceId terra.StringValue `hcl:"namespace_id,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// TrustedServicesAllowed: bool, optional
	TrustedServicesAllowed terra.BoolValue `hcl:"trusted_services_allowed,attr"`
	// NetworkRules: min=0
	NetworkRules []servicebusnamespacenetworkruleset.NetworkRules `hcl:"network_rules,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *servicebusnamespacenetworkruleset.Timeouts `hcl:"timeouts,block"`
}
type servicebusNamespaceNetworkRuleSetAttributes struct {
	ref terra.Reference
}

// DefaultAction returns a reference to field default_action of azurerm_servicebus_namespace_network_rule_set.
func (snnrs servicebusNamespaceNetworkRuleSetAttributes) DefaultAction() terra.StringValue {
	return terra.ReferenceAsString(snnrs.ref.Append("default_action"))
}

// Id returns a reference to field id of azurerm_servicebus_namespace_network_rule_set.
func (snnrs servicebusNamespaceNetworkRuleSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(snnrs.ref.Append("id"))
}

// IpRules returns a reference to field ip_rules of azurerm_servicebus_namespace_network_rule_set.
func (snnrs servicebusNamespaceNetworkRuleSetAttributes) IpRules() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](snnrs.ref.Append("ip_rules"))
}

// NamespaceId returns a reference to field namespace_id of azurerm_servicebus_namespace_network_rule_set.
func (snnrs servicebusNamespaceNetworkRuleSetAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(snnrs.ref.Append("namespace_id"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_servicebus_namespace_network_rule_set.
func (snnrs servicebusNamespaceNetworkRuleSetAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(snnrs.ref.Append("public_network_access_enabled"))
}

// TrustedServicesAllowed returns a reference to field trusted_services_allowed of azurerm_servicebus_namespace_network_rule_set.
func (snnrs servicebusNamespaceNetworkRuleSetAttributes) TrustedServicesAllowed() terra.BoolValue {
	return terra.ReferenceAsBool(snnrs.ref.Append("trusted_services_allowed"))
}

func (snnrs servicebusNamespaceNetworkRuleSetAttributes) NetworkRules() terra.SetValue[servicebusnamespacenetworkruleset.NetworkRulesAttributes] {
	return terra.ReferenceAsSet[servicebusnamespacenetworkruleset.NetworkRulesAttributes](snnrs.ref.Append("network_rules"))
}

func (snnrs servicebusNamespaceNetworkRuleSetAttributes) Timeouts() servicebusnamespacenetworkruleset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicebusnamespacenetworkruleset.TimeoutsAttributes](snnrs.ref.Append("timeouts"))
}

type servicebusNamespaceNetworkRuleSetState struct {
	DefaultAction              string                                                `json:"default_action"`
	Id                         string                                                `json:"id"`
	IpRules                    []string                                              `json:"ip_rules"`
	NamespaceId                string                                                `json:"namespace_id"`
	PublicNetworkAccessEnabled bool                                                  `json:"public_network_access_enabled"`
	TrustedServicesAllowed     bool                                                  `json:"trusted_services_allowed"`
	NetworkRules               []servicebusnamespacenetworkruleset.NetworkRulesState `json:"network_rules"`
	Timeouts                   *servicebusnamespacenetworkruleset.TimeoutsState      `json:"timeouts"`
}