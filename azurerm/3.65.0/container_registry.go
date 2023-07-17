// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	containerregistry "github.com/golingon/terraproviders/azurerm/3.65.0/containerregistry"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerRegistry creates a new instance of [ContainerRegistry].
func NewContainerRegistry(name string, args ContainerRegistryArgs) *ContainerRegistry {
	return &ContainerRegistry{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerRegistry)(nil)

// ContainerRegistry represents the Terraform resource azurerm_container_registry.
type ContainerRegistry struct {
	Name      string
	Args      ContainerRegistryArgs
	state     *containerRegistryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerRegistry].
func (cr *ContainerRegistry) Type() string {
	return "azurerm_container_registry"
}

// LocalName returns the local name for [ContainerRegistry].
func (cr *ContainerRegistry) LocalName() string {
	return cr.Name
}

// Configuration returns the configuration (args) for [ContainerRegistry].
func (cr *ContainerRegistry) Configuration() interface{} {
	return cr.Args
}

// DependOn is used for other resources to depend on [ContainerRegistry].
func (cr *ContainerRegistry) DependOn() terra.Reference {
	return terra.ReferenceResource(cr)
}

// Dependencies returns the list of resources [ContainerRegistry] depends_on.
func (cr *ContainerRegistry) Dependencies() terra.Dependencies {
	return cr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerRegistry].
func (cr *ContainerRegistry) LifecycleManagement() *terra.Lifecycle {
	return cr.Lifecycle
}

// Attributes returns the attributes for [ContainerRegistry].
func (cr *ContainerRegistry) Attributes() containerRegistryAttributes {
	return containerRegistryAttributes{ref: terra.ReferenceResource(cr)}
}

// ImportState imports the given attribute values into [ContainerRegistry]'s state.
func (cr *ContainerRegistry) ImportState(av io.Reader) error {
	cr.state = &containerRegistryState{}
	if err := json.NewDecoder(av).Decode(cr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cr.Type(), cr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerRegistry] has state.
func (cr *ContainerRegistry) State() (*containerRegistryState, bool) {
	return cr.state, cr.state != nil
}

// StateMust returns the state for [ContainerRegistry]. Panics if the state is nil.
func (cr *ContainerRegistry) StateMust() *containerRegistryState {
	if cr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cr.Type(), cr.LocalName()))
	}
	return cr.state
}

// ContainerRegistryArgs contains the configurations for azurerm_container_registry.
type ContainerRegistryArgs struct {
	// AdminEnabled: bool, optional
	AdminEnabled terra.BoolValue `hcl:"admin_enabled,attr"`
	// AnonymousPullEnabled: bool, optional
	AnonymousPullEnabled terra.BoolValue `hcl:"anonymous_pull_enabled,attr"`
	// DataEndpointEnabled: bool, optional
	DataEndpointEnabled terra.BoolValue `hcl:"data_endpoint_enabled,attr"`
	// ExportPolicyEnabled: bool, optional
	ExportPolicyEnabled terra.BoolValue `hcl:"export_policy_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkRuleBypassOption: string, optional
	NetworkRuleBypassOption terra.StringValue `hcl:"network_rule_bypass_option,attr"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// QuarantinePolicyEnabled: bool, optional
	QuarantinePolicyEnabled terra.BoolValue `hcl:"quarantine_policy_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ZoneRedundancyEnabled: bool, optional
	ZoneRedundancyEnabled terra.BoolValue `hcl:"zone_redundancy_enabled,attr"`
	// Encryption: min=0
	Encryption []containerregistry.Encryption `hcl:"encryption,block" validate:"min=0"`
	// NetworkRuleSet: min=0
	NetworkRuleSet []containerregistry.NetworkRuleSet `hcl:"network_rule_set,block" validate:"min=0"`
	// RetentionPolicy: min=0
	RetentionPolicy []containerregistry.RetentionPolicy `hcl:"retention_policy,block" validate:"min=0"`
	// TrustPolicy: min=0
	TrustPolicy []containerregistry.TrustPolicy `hcl:"trust_policy,block" validate:"min=0"`
	// Georeplications: min=0
	Georeplications []containerregistry.Georeplications `hcl:"georeplications,block" validate:"min=0"`
	// Identity: optional
	Identity *containerregistry.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *containerregistry.Timeouts `hcl:"timeouts,block"`
}
type containerRegistryAttributes struct {
	ref terra.Reference
}

// AdminEnabled returns a reference to field admin_enabled of azurerm_container_registry.
func (cr containerRegistryAttributes) AdminEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cr.ref.Append("admin_enabled"))
}

// AdminPassword returns a reference to field admin_password of azurerm_container_registry.
func (cr containerRegistryAttributes) AdminPassword() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("admin_password"))
}

// AdminUsername returns a reference to field admin_username of azurerm_container_registry.
func (cr containerRegistryAttributes) AdminUsername() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("admin_username"))
}

// AnonymousPullEnabled returns a reference to field anonymous_pull_enabled of azurerm_container_registry.
func (cr containerRegistryAttributes) AnonymousPullEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cr.ref.Append("anonymous_pull_enabled"))
}

// DataEndpointEnabled returns a reference to field data_endpoint_enabled of azurerm_container_registry.
func (cr containerRegistryAttributes) DataEndpointEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cr.ref.Append("data_endpoint_enabled"))
}

// ExportPolicyEnabled returns a reference to field export_policy_enabled of azurerm_container_registry.
func (cr containerRegistryAttributes) ExportPolicyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cr.ref.Append("export_policy_enabled"))
}

// Id returns a reference to field id of azurerm_container_registry.
func (cr containerRegistryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_container_registry.
func (cr containerRegistryAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("location"))
}

// LoginServer returns a reference to field login_server of azurerm_container_registry.
func (cr containerRegistryAttributes) LoginServer() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("login_server"))
}

// Name returns a reference to field name of azurerm_container_registry.
func (cr containerRegistryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("name"))
}

// NetworkRuleBypassOption returns a reference to field network_rule_bypass_option of azurerm_container_registry.
func (cr containerRegistryAttributes) NetworkRuleBypassOption() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("network_rule_bypass_option"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_container_registry.
func (cr containerRegistryAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cr.ref.Append("public_network_access_enabled"))
}

// QuarantinePolicyEnabled returns a reference to field quarantine_policy_enabled of azurerm_container_registry.
func (cr containerRegistryAttributes) QuarantinePolicyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cr.ref.Append("quarantine_policy_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_container_registry.
func (cr containerRegistryAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_container_registry.
func (cr containerRegistryAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_container_registry.
func (cr containerRegistryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cr.ref.Append("tags"))
}

// ZoneRedundancyEnabled returns a reference to field zone_redundancy_enabled of azurerm_container_registry.
func (cr containerRegistryAttributes) ZoneRedundancyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cr.ref.Append("zone_redundancy_enabled"))
}

func (cr containerRegistryAttributes) Encryption() terra.ListValue[containerregistry.EncryptionAttributes] {
	return terra.ReferenceAsList[containerregistry.EncryptionAttributes](cr.ref.Append("encryption"))
}

func (cr containerRegistryAttributes) NetworkRuleSet() terra.ListValue[containerregistry.NetworkRuleSetAttributes] {
	return terra.ReferenceAsList[containerregistry.NetworkRuleSetAttributes](cr.ref.Append("network_rule_set"))
}

func (cr containerRegistryAttributes) RetentionPolicy() terra.ListValue[containerregistry.RetentionPolicyAttributes] {
	return terra.ReferenceAsList[containerregistry.RetentionPolicyAttributes](cr.ref.Append("retention_policy"))
}

func (cr containerRegistryAttributes) TrustPolicy() terra.ListValue[containerregistry.TrustPolicyAttributes] {
	return terra.ReferenceAsList[containerregistry.TrustPolicyAttributes](cr.ref.Append("trust_policy"))
}

func (cr containerRegistryAttributes) Georeplications() terra.ListValue[containerregistry.GeoreplicationsAttributes] {
	return terra.ReferenceAsList[containerregistry.GeoreplicationsAttributes](cr.ref.Append("georeplications"))
}

func (cr containerRegistryAttributes) Identity() terra.ListValue[containerregistry.IdentityAttributes] {
	return terra.ReferenceAsList[containerregistry.IdentityAttributes](cr.ref.Append("identity"))
}

func (cr containerRegistryAttributes) Timeouts() containerregistry.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerregistry.TimeoutsAttributes](cr.ref.Append("timeouts"))
}

type containerRegistryState struct {
	AdminEnabled               bool                                     `json:"admin_enabled"`
	AdminPassword              string                                   `json:"admin_password"`
	AdminUsername              string                                   `json:"admin_username"`
	AnonymousPullEnabled       bool                                     `json:"anonymous_pull_enabled"`
	DataEndpointEnabled        bool                                     `json:"data_endpoint_enabled"`
	ExportPolicyEnabled        bool                                     `json:"export_policy_enabled"`
	Id                         string                                   `json:"id"`
	Location                   string                                   `json:"location"`
	LoginServer                string                                   `json:"login_server"`
	Name                       string                                   `json:"name"`
	NetworkRuleBypassOption    string                                   `json:"network_rule_bypass_option"`
	PublicNetworkAccessEnabled bool                                     `json:"public_network_access_enabled"`
	QuarantinePolicyEnabled    bool                                     `json:"quarantine_policy_enabled"`
	ResourceGroupName          string                                   `json:"resource_group_name"`
	Sku                        string                                   `json:"sku"`
	Tags                       map[string]string                        `json:"tags"`
	ZoneRedundancyEnabled      bool                                     `json:"zone_redundancy_enabled"`
	Encryption                 []containerregistry.EncryptionState      `json:"encryption"`
	NetworkRuleSet             []containerregistry.NetworkRuleSetState  `json:"network_rule_set"`
	RetentionPolicy            []containerregistry.RetentionPolicyState `json:"retention_policy"`
	TrustPolicy                []containerregistry.TrustPolicyState     `json:"trust_policy"`
	Georeplications            []containerregistry.GeoreplicationsState `json:"georeplications"`
	Identity                   []containerregistry.IdentityState        `json:"identity"`
	Timeouts                   *containerregistry.TimeoutsState         `json:"timeouts"`
}