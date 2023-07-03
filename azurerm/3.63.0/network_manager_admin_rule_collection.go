// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkmanageradminrulecollection "github.com/golingon/terraproviders/azurerm/3.63.0/networkmanageradminrulecollection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkManagerAdminRuleCollection creates a new instance of [NetworkManagerAdminRuleCollection].
func NewNetworkManagerAdminRuleCollection(name string, args NetworkManagerAdminRuleCollectionArgs) *NetworkManagerAdminRuleCollection {
	return &NetworkManagerAdminRuleCollection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManagerAdminRuleCollection)(nil)

// NetworkManagerAdminRuleCollection represents the Terraform resource azurerm_network_manager_admin_rule_collection.
type NetworkManagerAdminRuleCollection struct {
	Name      string
	Args      NetworkManagerAdminRuleCollectionArgs
	state     *networkManagerAdminRuleCollectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkManagerAdminRuleCollection].
func (nmarc *NetworkManagerAdminRuleCollection) Type() string {
	return "azurerm_network_manager_admin_rule_collection"
}

// LocalName returns the local name for [NetworkManagerAdminRuleCollection].
func (nmarc *NetworkManagerAdminRuleCollection) LocalName() string {
	return nmarc.Name
}

// Configuration returns the configuration (args) for [NetworkManagerAdminRuleCollection].
func (nmarc *NetworkManagerAdminRuleCollection) Configuration() interface{} {
	return nmarc.Args
}

// DependOn is used for other resources to depend on [NetworkManagerAdminRuleCollection].
func (nmarc *NetworkManagerAdminRuleCollection) DependOn() terra.Reference {
	return terra.ReferenceResource(nmarc)
}

// Dependencies returns the list of resources [NetworkManagerAdminRuleCollection] depends_on.
func (nmarc *NetworkManagerAdminRuleCollection) Dependencies() terra.Dependencies {
	return nmarc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkManagerAdminRuleCollection].
func (nmarc *NetworkManagerAdminRuleCollection) LifecycleManagement() *terra.Lifecycle {
	return nmarc.Lifecycle
}

// Attributes returns the attributes for [NetworkManagerAdminRuleCollection].
func (nmarc *NetworkManagerAdminRuleCollection) Attributes() networkManagerAdminRuleCollectionAttributes {
	return networkManagerAdminRuleCollectionAttributes{ref: terra.ReferenceResource(nmarc)}
}

// ImportState imports the given attribute values into [NetworkManagerAdminRuleCollection]'s state.
func (nmarc *NetworkManagerAdminRuleCollection) ImportState(av io.Reader) error {
	nmarc.state = &networkManagerAdminRuleCollectionState{}
	if err := json.NewDecoder(av).Decode(nmarc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nmarc.Type(), nmarc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkManagerAdminRuleCollection] has state.
func (nmarc *NetworkManagerAdminRuleCollection) State() (*networkManagerAdminRuleCollectionState, bool) {
	return nmarc.state, nmarc.state != nil
}

// StateMust returns the state for [NetworkManagerAdminRuleCollection]. Panics if the state is nil.
func (nmarc *NetworkManagerAdminRuleCollection) StateMust() *networkManagerAdminRuleCollectionState {
	if nmarc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nmarc.Type(), nmarc.LocalName()))
	}
	return nmarc.state
}

// NetworkManagerAdminRuleCollectionArgs contains the configurations for azurerm_network_manager_admin_rule_collection.
type NetworkManagerAdminRuleCollectionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkGroupIds: list of string, required
	NetworkGroupIds terra.ListValue[terra.StringValue] `hcl:"network_group_ids,attr" validate:"required"`
	// SecurityAdminConfigurationId: string, required
	SecurityAdminConfigurationId terra.StringValue `hcl:"security_admin_configuration_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanageradminrulecollection.Timeouts `hcl:"timeouts,block"`
}
type networkManagerAdminRuleCollectionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_network_manager_admin_rule_collection.
func (nmarc networkManagerAdminRuleCollectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nmarc.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_network_manager_admin_rule_collection.
func (nmarc networkManagerAdminRuleCollectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nmarc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_network_manager_admin_rule_collection.
func (nmarc networkManagerAdminRuleCollectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nmarc.ref.Append("name"))
}

// NetworkGroupIds returns a reference to field network_group_ids of azurerm_network_manager_admin_rule_collection.
func (nmarc networkManagerAdminRuleCollectionAttributes) NetworkGroupIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nmarc.ref.Append("network_group_ids"))
}

// SecurityAdminConfigurationId returns a reference to field security_admin_configuration_id of azurerm_network_manager_admin_rule_collection.
func (nmarc networkManagerAdminRuleCollectionAttributes) SecurityAdminConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(nmarc.ref.Append("security_admin_configuration_id"))
}

func (nmarc networkManagerAdminRuleCollectionAttributes) Timeouts() networkmanageradminrulecollection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanageradminrulecollection.TimeoutsAttributes](nmarc.ref.Append("timeouts"))
}

type networkManagerAdminRuleCollectionState struct {
	Description                  string                                           `json:"description"`
	Id                           string                                           `json:"id"`
	Name                         string                                           `json:"name"`
	NetworkGroupIds              []string                                         `json:"network_group_ids"`
	SecurityAdminConfigurationId string                                           `json:"security_admin_configuration_id"`
	Timeouts                     *networkmanageradminrulecollection.TimeoutsState `json:"timeouts"`
}
