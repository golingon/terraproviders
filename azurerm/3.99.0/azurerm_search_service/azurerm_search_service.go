// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_search_service

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

// Resource represents the Terraform resource azurerm_search_service.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSearchServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ass *Resource) Type() string {
	return "azurerm_search_service"
}

// LocalName returns the local name for [Resource].
func (ass *Resource) LocalName() string {
	return ass.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ass *Resource) Configuration() interface{} {
	return ass.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ass *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ass)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ass *Resource) Dependencies() terra.Dependencies {
	return ass.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ass *Resource) LifecycleManagement() *terra.Lifecycle {
	return ass.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ass *Resource) Attributes() azurermSearchServiceAttributes {
	return azurermSearchServiceAttributes{ref: terra.ReferenceResource(ass)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ass *Resource) ImportState(state io.Reader) error {
	ass.state = &azurermSearchServiceState{}
	if err := json.NewDecoder(state).Decode(ass.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ass.Type(), ass.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ass *Resource) State() (*azurermSearchServiceState, bool) {
	return ass.state, ass.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ass *Resource) StateMust() *azurermSearchServiceState {
	if ass.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ass.Type(), ass.LocalName()))
	}
	return ass.state
}

// Args contains the configurations for azurerm_search_service.
type Args struct {
	// AllowedIps: set of string, optional
	AllowedIps terra.SetValue[terra.StringValue] `hcl:"allowed_ips,attr"`
	// AuthenticationFailureMode: string, optional
	AuthenticationFailureMode terra.StringValue `hcl:"authentication_failure_mode,attr"`
	// CustomerManagedKeyEnforcementEnabled: bool, optional
	CustomerManagedKeyEnforcementEnabled terra.BoolValue `hcl:"customer_managed_key_enforcement_enabled,attr"`
	// HostingMode: string, optional
	HostingMode terra.StringValue `hcl:"hosting_mode,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalAuthenticationEnabled: bool, optional
	LocalAuthenticationEnabled terra.BoolValue `hcl:"local_authentication_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PartitionCount: number, optional
	PartitionCount terra.NumberValue `hcl:"partition_count,attr"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ReplicaCount: number, optional
	ReplicaCount terra.NumberValue `hcl:"replica_count,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SemanticSearchSku: string, optional
	SemanticSearchSku terra.StringValue `hcl:"semantic_search_sku,attr"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: optional
	Identity *Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSearchServiceAttributes struct {
	ref terra.Reference
}

// AllowedIps returns a reference to field allowed_ips of azurerm_search_service.
func (ass azurermSearchServiceAttributes) AllowedIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ass.ref.Append("allowed_ips"))
}

// AuthenticationFailureMode returns a reference to field authentication_failure_mode of azurerm_search_service.
func (ass azurermSearchServiceAttributes) AuthenticationFailureMode() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("authentication_failure_mode"))
}

// CustomerManagedKeyEnforcementEnabled returns a reference to field customer_managed_key_enforcement_enabled of azurerm_search_service.
func (ass azurermSearchServiceAttributes) CustomerManagedKeyEnforcementEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ass.ref.Append("customer_managed_key_enforcement_enabled"))
}

// HostingMode returns a reference to field hosting_mode of azurerm_search_service.
func (ass azurermSearchServiceAttributes) HostingMode() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("hosting_mode"))
}

// Id returns a reference to field id of azurerm_search_service.
func (ass azurermSearchServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("id"))
}

// LocalAuthenticationEnabled returns a reference to field local_authentication_enabled of azurerm_search_service.
func (ass azurermSearchServiceAttributes) LocalAuthenticationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ass.ref.Append("local_authentication_enabled"))
}

// Location returns a reference to field location of azurerm_search_service.
func (ass azurermSearchServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_search_service.
func (ass azurermSearchServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("name"))
}

// PartitionCount returns a reference to field partition_count of azurerm_search_service.
func (ass azurermSearchServiceAttributes) PartitionCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ass.ref.Append("partition_count"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_search_service.
func (ass azurermSearchServiceAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("primary_key"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_search_service.
func (ass azurermSearchServiceAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ass.ref.Append("public_network_access_enabled"))
}

// ReplicaCount returns a reference to field replica_count of azurerm_search_service.
func (ass azurermSearchServiceAttributes) ReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ass.ref.Append("replica_count"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_search_service.
func (ass azurermSearchServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("resource_group_name"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_search_service.
func (ass azurermSearchServiceAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("secondary_key"))
}

// SemanticSearchSku returns a reference to field semantic_search_sku of azurerm_search_service.
func (ass azurermSearchServiceAttributes) SemanticSearchSku() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("semantic_search_sku"))
}

// Sku returns a reference to field sku of azurerm_search_service.
func (ass azurermSearchServiceAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_search_service.
func (ass azurermSearchServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ass.ref.Append("tags"))
}

func (ass azurermSearchServiceAttributes) QueryKeys() terra.ListValue[QueryKeysAttributes] {
	return terra.ReferenceAsList[QueryKeysAttributes](ass.ref.Append("query_keys"))
}

func (ass azurermSearchServiceAttributes) Identity() terra.ListValue[IdentityAttributes] {
	return terra.ReferenceAsList[IdentityAttributes](ass.ref.Append("identity"))
}

func (ass azurermSearchServiceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ass.ref.Append("timeouts"))
}

type azurermSearchServiceState struct {
	AllowedIps                           []string          `json:"allowed_ips"`
	AuthenticationFailureMode            string            `json:"authentication_failure_mode"`
	CustomerManagedKeyEnforcementEnabled bool              `json:"customer_managed_key_enforcement_enabled"`
	HostingMode                          string            `json:"hosting_mode"`
	Id                                   string            `json:"id"`
	LocalAuthenticationEnabled           bool              `json:"local_authentication_enabled"`
	Location                             string            `json:"location"`
	Name                                 string            `json:"name"`
	PartitionCount                       float64           `json:"partition_count"`
	PrimaryKey                           string            `json:"primary_key"`
	PublicNetworkAccessEnabled           bool              `json:"public_network_access_enabled"`
	ReplicaCount                         float64           `json:"replica_count"`
	ResourceGroupName                    string            `json:"resource_group_name"`
	SecondaryKey                         string            `json:"secondary_key"`
	SemanticSearchSku                    string            `json:"semantic_search_sku"`
	Sku                                  string            `json:"sku"`
	Tags                                 map[string]string `json:"tags"`
	QueryKeys                            []QueryKeysState  `json:"query_keys"`
	Identity                             []IdentityState   `json:"identity"`
	Timeouts                             *TimeoutsState    `json:"timeouts"`
}
