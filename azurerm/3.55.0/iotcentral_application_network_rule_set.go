// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iotcentralapplicationnetworkruleset "github.com/golingon/terraproviders/azurerm/3.55.0/iotcentralapplicationnetworkruleset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotcentralApplicationNetworkRuleSet creates a new instance of [IotcentralApplicationNetworkRuleSet].
func NewIotcentralApplicationNetworkRuleSet(name string, args IotcentralApplicationNetworkRuleSetArgs) *IotcentralApplicationNetworkRuleSet {
	return &IotcentralApplicationNetworkRuleSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotcentralApplicationNetworkRuleSet)(nil)

// IotcentralApplicationNetworkRuleSet represents the Terraform resource azurerm_iotcentral_application_network_rule_set.
type IotcentralApplicationNetworkRuleSet struct {
	Name      string
	Args      IotcentralApplicationNetworkRuleSetArgs
	state     *iotcentralApplicationNetworkRuleSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotcentralApplicationNetworkRuleSet].
func (ianrs *IotcentralApplicationNetworkRuleSet) Type() string {
	return "azurerm_iotcentral_application_network_rule_set"
}

// LocalName returns the local name for [IotcentralApplicationNetworkRuleSet].
func (ianrs *IotcentralApplicationNetworkRuleSet) LocalName() string {
	return ianrs.Name
}

// Configuration returns the configuration (args) for [IotcentralApplicationNetworkRuleSet].
func (ianrs *IotcentralApplicationNetworkRuleSet) Configuration() interface{} {
	return ianrs.Args
}

// DependOn is used for other resources to depend on [IotcentralApplicationNetworkRuleSet].
func (ianrs *IotcentralApplicationNetworkRuleSet) DependOn() terra.Reference {
	return terra.ReferenceResource(ianrs)
}

// Dependencies returns the list of resources [IotcentralApplicationNetworkRuleSet] depends_on.
func (ianrs *IotcentralApplicationNetworkRuleSet) Dependencies() terra.Dependencies {
	return ianrs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotcentralApplicationNetworkRuleSet].
func (ianrs *IotcentralApplicationNetworkRuleSet) LifecycleManagement() *terra.Lifecycle {
	return ianrs.Lifecycle
}

// Attributes returns the attributes for [IotcentralApplicationNetworkRuleSet].
func (ianrs *IotcentralApplicationNetworkRuleSet) Attributes() iotcentralApplicationNetworkRuleSetAttributes {
	return iotcentralApplicationNetworkRuleSetAttributes{ref: terra.ReferenceResource(ianrs)}
}

// ImportState imports the given attribute values into [IotcentralApplicationNetworkRuleSet]'s state.
func (ianrs *IotcentralApplicationNetworkRuleSet) ImportState(av io.Reader) error {
	ianrs.state = &iotcentralApplicationNetworkRuleSetState{}
	if err := json.NewDecoder(av).Decode(ianrs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ianrs.Type(), ianrs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotcentralApplicationNetworkRuleSet] has state.
func (ianrs *IotcentralApplicationNetworkRuleSet) State() (*iotcentralApplicationNetworkRuleSetState, bool) {
	return ianrs.state, ianrs.state != nil
}

// StateMust returns the state for [IotcentralApplicationNetworkRuleSet]. Panics if the state is nil.
func (ianrs *IotcentralApplicationNetworkRuleSet) StateMust() *iotcentralApplicationNetworkRuleSetState {
	if ianrs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ianrs.Type(), ianrs.LocalName()))
	}
	return ianrs.state
}

// IotcentralApplicationNetworkRuleSetArgs contains the configurations for azurerm_iotcentral_application_network_rule_set.
type IotcentralApplicationNetworkRuleSetArgs struct {
	// ApplyToDevice: bool, optional
	ApplyToDevice terra.BoolValue `hcl:"apply_to_device,attr"`
	// DefaultAction: string, optional
	DefaultAction terra.StringValue `hcl:"default_action,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IotcentralApplicationId: string, required
	IotcentralApplicationId terra.StringValue `hcl:"iotcentral_application_id,attr" validate:"required"`
	// IpRule: min=0
	IpRule []iotcentralapplicationnetworkruleset.IpRule `hcl:"ip_rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *iotcentralapplicationnetworkruleset.Timeouts `hcl:"timeouts,block"`
}
type iotcentralApplicationNetworkRuleSetAttributes struct {
	ref terra.Reference
}

// ApplyToDevice returns a reference to field apply_to_device of azurerm_iotcentral_application_network_rule_set.
func (ianrs iotcentralApplicationNetworkRuleSetAttributes) ApplyToDevice() terra.BoolValue {
	return terra.ReferenceAsBool(ianrs.ref.Append("apply_to_device"))
}

// DefaultAction returns a reference to field default_action of azurerm_iotcentral_application_network_rule_set.
func (ianrs iotcentralApplicationNetworkRuleSetAttributes) DefaultAction() terra.StringValue {
	return terra.ReferenceAsString(ianrs.ref.Append("default_action"))
}

// Id returns a reference to field id of azurerm_iotcentral_application_network_rule_set.
func (ianrs iotcentralApplicationNetworkRuleSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ianrs.ref.Append("id"))
}

// IotcentralApplicationId returns a reference to field iotcentral_application_id of azurerm_iotcentral_application_network_rule_set.
func (ianrs iotcentralApplicationNetworkRuleSetAttributes) IotcentralApplicationId() terra.StringValue {
	return terra.ReferenceAsString(ianrs.ref.Append("iotcentral_application_id"))
}

func (ianrs iotcentralApplicationNetworkRuleSetAttributes) IpRule() terra.ListValue[iotcentralapplicationnetworkruleset.IpRuleAttributes] {
	return terra.ReferenceAsList[iotcentralapplicationnetworkruleset.IpRuleAttributes](ianrs.ref.Append("ip_rule"))
}

func (ianrs iotcentralApplicationNetworkRuleSetAttributes) Timeouts() iotcentralapplicationnetworkruleset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iotcentralapplicationnetworkruleset.TimeoutsAttributes](ianrs.ref.Append("timeouts"))
}

type iotcentralApplicationNetworkRuleSetState struct {
	ApplyToDevice           bool                                               `json:"apply_to_device"`
	DefaultAction           string                                             `json:"default_action"`
	Id                      string                                             `json:"id"`
	IotcentralApplicationId string                                             `json:"iotcentral_application_id"`
	IpRule                  []iotcentralapplicationnetworkruleset.IpRuleState  `json:"ip_rule"`
	Timeouts                *iotcentralapplicationnetworkruleset.TimeoutsState `json:"timeouts"`
}
