// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkmanagersecurityadminconfiguration "github.com/golingon/terraproviders/azurerm/3.69.0/networkmanagersecurityadminconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkManagerSecurityAdminConfiguration creates a new instance of [NetworkManagerSecurityAdminConfiguration].
func NewNetworkManagerSecurityAdminConfiguration(name string, args NetworkManagerSecurityAdminConfigurationArgs) *NetworkManagerSecurityAdminConfiguration {
	return &NetworkManagerSecurityAdminConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManagerSecurityAdminConfiguration)(nil)

// NetworkManagerSecurityAdminConfiguration represents the Terraform resource azurerm_network_manager_security_admin_configuration.
type NetworkManagerSecurityAdminConfiguration struct {
	Name      string
	Args      NetworkManagerSecurityAdminConfigurationArgs
	state     *networkManagerSecurityAdminConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkManagerSecurityAdminConfiguration].
func (nmsac *NetworkManagerSecurityAdminConfiguration) Type() string {
	return "azurerm_network_manager_security_admin_configuration"
}

// LocalName returns the local name for [NetworkManagerSecurityAdminConfiguration].
func (nmsac *NetworkManagerSecurityAdminConfiguration) LocalName() string {
	return nmsac.Name
}

// Configuration returns the configuration (args) for [NetworkManagerSecurityAdminConfiguration].
func (nmsac *NetworkManagerSecurityAdminConfiguration) Configuration() interface{} {
	return nmsac.Args
}

// DependOn is used for other resources to depend on [NetworkManagerSecurityAdminConfiguration].
func (nmsac *NetworkManagerSecurityAdminConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(nmsac)
}

// Dependencies returns the list of resources [NetworkManagerSecurityAdminConfiguration] depends_on.
func (nmsac *NetworkManagerSecurityAdminConfiguration) Dependencies() terra.Dependencies {
	return nmsac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkManagerSecurityAdminConfiguration].
func (nmsac *NetworkManagerSecurityAdminConfiguration) LifecycleManagement() *terra.Lifecycle {
	return nmsac.Lifecycle
}

// Attributes returns the attributes for [NetworkManagerSecurityAdminConfiguration].
func (nmsac *NetworkManagerSecurityAdminConfiguration) Attributes() networkManagerSecurityAdminConfigurationAttributes {
	return networkManagerSecurityAdminConfigurationAttributes{ref: terra.ReferenceResource(nmsac)}
}

// ImportState imports the given attribute values into [NetworkManagerSecurityAdminConfiguration]'s state.
func (nmsac *NetworkManagerSecurityAdminConfiguration) ImportState(av io.Reader) error {
	nmsac.state = &networkManagerSecurityAdminConfigurationState{}
	if err := json.NewDecoder(av).Decode(nmsac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nmsac.Type(), nmsac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkManagerSecurityAdminConfiguration] has state.
func (nmsac *NetworkManagerSecurityAdminConfiguration) State() (*networkManagerSecurityAdminConfigurationState, bool) {
	return nmsac.state, nmsac.state != nil
}

// StateMust returns the state for [NetworkManagerSecurityAdminConfiguration]. Panics if the state is nil.
func (nmsac *NetworkManagerSecurityAdminConfiguration) StateMust() *networkManagerSecurityAdminConfigurationState {
	if nmsac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nmsac.Type(), nmsac.LocalName()))
	}
	return nmsac.state
}

// NetworkManagerSecurityAdminConfigurationArgs contains the configurations for azurerm_network_manager_security_admin_configuration.
type NetworkManagerSecurityAdminConfigurationArgs struct {
	// ApplyOnNetworkIntentPolicyBasedServices: list of string, optional
	ApplyOnNetworkIntentPolicyBasedServices terra.ListValue[terra.StringValue] `hcl:"apply_on_network_intent_policy_based_services,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkManagerId: string, required
	NetworkManagerId terra.StringValue `hcl:"network_manager_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagersecurityadminconfiguration.Timeouts `hcl:"timeouts,block"`
}
type networkManagerSecurityAdminConfigurationAttributes struct {
	ref terra.Reference
}

// ApplyOnNetworkIntentPolicyBasedServices returns a reference to field apply_on_network_intent_policy_based_services of azurerm_network_manager_security_admin_configuration.
func (nmsac networkManagerSecurityAdminConfigurationAttributes) ApplyOnNetworkIntentPolicyBasedServices() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nmsac.ref.Append("apply_on_network_intent_policy_based_services"))
}

// Description returns a reference to field description of azurerm_network_manager_security_admin_configuration.
func (nmsac networkManagerSecurityAdminConfigurationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nmsac.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_network_manager_security_admin_configuration.
func (nmsac networkManagerSecurityAdminConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nmsac.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_network_manager_security_admin_configuration.
func (nmsac networkManagerSecurityAdminConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nmsac.ref.Append("name"))
}

// NetworkManagerId returns a reference to field network_manager_id of azurerm_network_manager_security_admin_configuration.
func (nmsac networkManagerSecurityAdminConfigurationAttributes) NetworkManagerId() terra.StringValue {
	return terra.ReferenceAsString(nmsac.ref.Append("network_manager_id"))
}

func (nmsac networkManagerSecurityAdminConfigurationAttributes) Timeouts() networkmanagersecurityadminconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagersecurityadminconfiguration.TimeoutsAttributes](nmsac.ref.Append("timeouts"))
}

type networkManagerSecurityAdminConfigurationState struct {
	ApplyOnNetworkIntentPolicyBasedServices []string                                                `json:"apply_on_network_intent_policy_based_services"`
	Description                             string                                                  `json:"description"`
	Id                                      string                                                  `json:"id"`
	Name                                    string                                                  `json:"name"`
	NetworkManagerId                        string                                                  `json:"network_manager_id"`
	Timeouts                                *networkmanagersecurityadminconfiguration.TimeoutsState `json:"timeouts"`
}
