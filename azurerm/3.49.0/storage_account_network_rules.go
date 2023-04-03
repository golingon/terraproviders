// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storageaccountnetworkrules "github.com/golingon/terraproviders/azurerm/3.49.0/storageaccountnetworkrules"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageAccountNetworkRules creates a new instance of [StorageAccountNetworkRules].
func NewStorageAccountNetworkRules(name string, args StorageAccountNetworkRulesArgs) *StorageAccountNetworkRules {
	return &StorageAccountNetworkRules{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageAccountNetworkRules)(nil)

// StorageAccountNetworkRules represents the Terraform resource azurerm_storage_account_network_rules.
type StorageAccountNetworkRules struct {
	Name      string
	Args      StorageAccountNetworkRulesArgs
	state     *storageAccountNetworkRulesState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageAccountNetworkRules].
func (sanr *StorageAccountNetworkRules) Type() string {
	return "azurerm_storage_account_network_rules"
}

// LocalName returns the local name for [StorageAccountNetworkRules].
func (sanr *StorageAccountNetworkRules) LocalName() string {
	return sanr.Name
}

// Configuration returns the configuration (args) for [StorageAccountNetworkRules].
func (sanr *StorageAccountNetworkRules) Configuration() interface{} {
	return sanr.Args
}

// DependOn is used for other resources to depend on [StorageAccountNetworkRules].
func (sanr *StorageAccountNetworkRules) DependOn() terra.Reference {
	return terra.ReferenceResource(sanr)
}

// Dependencies returns the list of resources [StorageAccountNetworkRules] depends_on.
func (sanr *StorageAccountNetworkRules) Dependencies() terra.Dependencies {
	return sanr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageAccountNetworkRules].
func (sanr *StorageAccountNetworkRules) LifecycleManagement() *terra.Lifecycle {
	return sanr.Lifecycle
}

// Attributes returns the attributes for [StorageAccountNetworkRules].
func (sanr *StorageAccountNetworkRules) Attributes() storageAccountNetworkRulesAttributes {
	return storageAccountNetworkRulesAttributes{ref: terra.ReferenceResource(sanr)}
}

// ImportState imports the given attribute values into [StorageAccountNetworkRules]'s state.
func (sanr *StorageAccountNetworkRules) ImportState(av io.Reader) error {
	sanr.state = &storageAccountNetworkRulesState{}
	if err := json.NewDecoder(av).Decode(sanr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sanr.Type(), sanr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageAccountNetworkRules] has state.
func (sanr *StorageAccountNetworkRules) State() (*storageAccountNetworkRulesState, bool) {
	return sanr.state, sanr.state != nil
}

// StateMust returns the state for [StorageAccountNetworkRules]. Panics if the state is nil.
func (sanr *StorageAccountNetworkRules) StateMust() *storageAccountNetworkRulesState {
	if sanr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sanr.Type(), sanr.LocalName()))
	}
	return sanr.state
}

// StorageAccountNetworkRulesArgs contains the configurations for azurerm_storage_account_network_rules.
type StorageAccountNetworkRulesArgs struct {
	// Bypass: set of string, optional
	Bypass terra.SetValue[terra.StringValue] `hcl:"bypass,attr"`
	// DefaultAction: string, required
	DefaultAction terra.StringValue `hcl:"default_action,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpRules: set of string, optional
	IpRules terra.SetValue[terra.StringValue] `hcl:"ip_rules,attr"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// VirtualNetworkSubnetIds: set of string, optional
	VirtualNetworkSubnetIds terra.SetValue[terra.StringValue] `hcl:"virtual_network_subnet_ids,attr"`
	// PrivateLinkAccess: min=0
	PrivateLinkAccess []storageaccountnetworkrules.PrivateLinkAccess `hcl:"private_link_access,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *storageaccountnetworkrules.Timeouts `hcl:"timeouts,block"`
}
type storageAccountNetworkRulesAttributes struct {
	ref terra.Reference
}

// Bypass returns a reference to field bypass of azurerm_storage_account_network_rules.
func (sanr storageAccountNetworkRulesAttributes) Bypass() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sanr.ref.Append("bypass"))
}

// DefaultAction returns a reference to field default_action of azurerm_storage_account_network_rules.
func (sanr storageAccountNetworkRulesAttributes) DefaultAction() terra.StringValue {
	return terra.ReferenceAsString(sanr.ref.Append("default_action"))
}

// Id returns a reference to field id of azurerm_storage_account_network_rules.
func (sanr storageAccountNetworkRulesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sanr.ref.Append("id"))
}

// IpRules returns a reference to field ip_rules of azurerm_storage_account_network_rules.
func (sanr storageAccountNetworkRulesAttributes) IpRules() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sanr.ref.Append("ip_rules"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_storage_account_network_rules.
func (sanr storageAccountNetworkRulesAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(sanr.ref.Append("storage_account_id"))
}

// VirtualNetworkSubnetIds returns a reference to field virtual_network_subnet_ids of azurerm_storage_account_network_rules.
func (sanr storageAccountNetworkRulesAttributes) VirtualNetworkSubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sanr.ref.Append("virtual_network_subnet_ids"))
}

func (sanr storageAccountNetworkRulesAttributes) PrivateLinkAccess() terra.ListValue[storageaccountnetworkrules.PrivateLinkAccessAttributes] {
	return terra.ReferenceAsList[storageaccountnetworkrules.PrivateLinkAccessAttributes](sanr.ref.Append("private_link_access"))
}

func (sanr storageAccountNetworkRulesAttributes) Timeouts() storageaccountnetworkrules.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storageaccountnetworkrules.TimeoutsAttributes](sanr.ref.Append("timeouts"))
}

type storageAccountNetworkRulesState struct {
	Bypass                  []string                                            `json:"bypass"`
	DefaultAction           string                                              `json:"default_action"`
	Id                      string                                              `json:"id"`
	IpRules                 []string                                            `json:"ip_rules"`
	StorageAccountId        string                                              `json:"storage_account_id"`
	VirtualNetworkSubnetIds []string                                            `json:"virtual_network_subnet_ids"`
	PrivateLinkAccess       []storageaccountnetworkrules.PrivateLinkAccessState `json:"private_link_access"`
	Timeouts                *storageaccountnetworkrules.TimeoutsState           `json:"timeouts"`
}
