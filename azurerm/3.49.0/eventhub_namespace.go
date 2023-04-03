// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	eventhubnamespace "github.com/golingon/terraproviders/azurerm/3.49.0/eventhubnamespace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventhubNamespace creates a new instance of [EventhubNamespace].
func NewEventhubNamespace(name string, args EventhubNamespaceArgs) *EventhubNamespace {
	return &EventhubNamespace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventhubNamespace)(nil)

// EventhubNamespace represents the Terraform resource azurerm_eventhub_namespace.
type EventhubNamespace struct {
	Name      string
	Args      EventhubNamespaceArgs
	state     *eventhubNamespaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventhubNamespace].
func (en *EventhubNamespace) Type() string {
	return "azurerm_eventhub_namespace"
}

// LocalName returns the local name for [EventhubNamespace].
func (en *EventhubNamespace) LocalName() string {
	return en.Name
}

// Configuration returns the configuration (args) for [EventhubNamespace].
func (en *EventhubNamespace) Configuration() interface{} {
	return en.Args
}

// DependOn is used for other resources to depend on [EventhubNamespace].
func (en *EventhubNamespace) DependOn() terra.Reference {
	return terra.ReferenceResource(en)
}

// Dependencies returns the list of resources [EventhubNamespace] depends_on.
func (en *EventhubNamespace) Dependencies() terra.Dependencies {
	return en.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventhubNamespace].
func (en *EventhubNamespace) LifecycleManagement() *terra.Lifecycle {
	return en.Lifecycle
}

// Attributes returns the attributes for [EventhubNamespace].
func (en *EventhubNamespace) Attributes() eventhubNamespaceAttributes {
	return eventhubNamespaceAttributes{ref: terra.ReferenceResource(en)}
}

// ImportState imports the given attribute values into [EventhubNamespace]'s state.
func (en *EventhubNamespace) ImportState(av io.Reader) error {
	en.state = &eventhubNamespaceState{}
	if err := json.NewDecoder(av).Decode(en.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", en.Type(), en.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventhubNamespace] has state.
func (en *EventhubNamespace) State() (*eventhubNamespaceState, bool) {
	return en.state, en.state != nil
}

// StateMust returns the state for [EventhubNamespace]. Panics if the state is nil.
func (en *EventhubNamespace) StateMust() *eventhubNamespaceState {
	if en.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", en.Type(), en.LocalName()))
	}
	return en.state
}

// EventhubNamespaceArgs contains the configurations for azurerm_eventhub_namespace.
type EventhubNamespaceArgs struct {
	// AutoInflateEnabled: bool, optional
	AutoInflateEnabled terra.BoolValue `hcl:"auto_inflate_enabled,attr"`
	// Capacity: number, optional
	Capacity terra.NumberValue `hcl:"capacity,attr"`
	// DedicatedClusterId: string, optional
	DedicatedClusterId terra.StringValue `hcl:"dedicated_cluster_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalAuthenticationEnabled: bool, optional
	LocalAuthenticationEnabled terra.BoolValue `hcl:"local_authentication_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MaximumThroughputUnits: number, optional
	MaximumThroughputUnits terra.NumberValue `hcl:"maximum_throughput_units,attr"`
	// MinimumTlsVersion: string, optional
	MinimumTlsVersion terra.StringValue `hcl:"minimum_tls_version,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ZoneRedundant: bool, optional
	ZoneRedundant terra.BoolValue `hcl:"zone_redundant,attr"`
	// NetworkRulesets: min=0
	NetworkRulesets []eventhubnamespace.NetworkRulesets `hcl:"network_rulesets,block" validate:"min=0"`
	// Identity: optional
	Identity *eventhubnamespace.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *eventhubnamespace.Timeouts `hcl:"timeouts,block"`
}
type eventhubNamespaceAttributes struct {
	ref terra.Reference
}

// AutoInflateEnabled returns a reference to field auto_inflate_enabled of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) AutoInflateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(en.ref.Append("auto_inflate_enabled"))
}

// Capacity returns a reference to field capacity of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(en.ref.Append("capacity"))
}

// DedicatedClusterId returns a reference to field dedicated_cluster_id of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) DedicatedClusterId() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("dedicated_cluster_id"))
}

// DefaultPrimaryConnectionString returns a reference to field default_primary_connection_string of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) DefaultPrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("default_primary_connection_string"))
}

// DefaultPrimaryConnectionStringAlias returns a reference to field default_primary_connection_string_alias of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) DefaultPrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("default_primary_connection_string_alias"))
}

// DefaultPrimaryKey returns a reference to field default_primary_key of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) DefaultPrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("default_primary_key"))
}

// DefaultSecondaryConnectionString returns a reference to field default_secondary_connection_string of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) DefaultSecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("default_secondary_connection_string"))
}

// DefaultSecondaryConnectionStringAlias returns a reference to field default_secondary_connection_string_alias of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) DefaultSecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("default_secondary_connection_string_alias"))
}

// DefaultSecondaryKey returns a reference to field default_secondary_key of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) DefaultSecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("default_secondary_key"))
}

// Id returns a reference to field id of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("id"))
}

// LocalAuthenticationEnabled returns a reference to field local_authentication_enabled of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) LocalAuthenticationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(en.ref.Append("local_authentication_enabled"))
}

// Location returns a reference to field location of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("location"))
}

// MaximumThroughputUnits returns a reference to field maximum_throughput_units of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) MaximumThroughputUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(en.ref.Append("maximum_throughput_units"))
}

// MinimumTlsVersion returns a reference to field minimum_tls_version of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) MinimumTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("minimum_tls_version"))
}

// Name returns a reference to field name of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(en.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](en.ref.Append("tags"))
}

// ZoneRedundant returns a reference to field zone_redundant of azurerm_eventhub_namespace.
func (en eventhubNamespaceAttributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(en.ref.Append("zone_redundant"))
}

func (en eventhubNamespaceAttributes) NetworkRulesets() terra.ListValue[eventhubnamespace.NetworkRulesetsAttributes] {
	return terra.ReferenceAsList[eventhubnamespace.NetworkRulesetsAttributes](en.ref.Append("network_rulesets"))
}

func (en eventhubNamespaceAttributes) Identity() terra.ListValue[eventhubnamespace.IdentityAttributes] {
	return terra.ReferenceAsList[eventhubnamespace.IdentityAttributes](en.ref.Append("identity"))
}

func (en eventhubNamespaceAttributes) Timeouts() eventhubnamespace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventhubnamespace.TimeoutsAttributes](en.ref.Append("timeouts"))
}

type eventhubNamespaceState struct {
	AutoInflateEnabled                    bool                                     `json:"auto_inflate_enabled"`
	Capacity                              float64                                  `json:"capacity"`
	DedicatedClusterId                    string                                   `json:"dedicated_cluster_id"`
	DefaultPrimaryConnectionString        string                                   `json:"default_primary_connection_string"`
	DefaultPrimaryConnectionStringAlias   string                                   `json:"default_primary_connection_string_alias"`
	DefaultPrimaryKey                     string                                   `json:"default_primary_key"`
	DefaultSecondaryConnectionString      string                                   `json:"default_secondary_connection_string"`
	DefaultSecondaryConnectionStringAlias string                                   `json:"default_secondary_connection_string_alias"`
	DefaultSecondaryKey                   string                                   `json:"default_secondary_key"`
	Id                                    string                                   `json:"id"`
	LocalAuthenticationEnabled            bool                                     `json:"local_authentication_enabled"`
	Location                              string                                   `json:"location"`
	MaximumThroughputUnits                float64                                  `json:"maximum_throughput_units"`
	MinimumTlsVersion                     string                                   `json:"minimum_tls_version"`
	Name                                  string                                   `json:"name"`
	PublicNetworkAccessEnabled            bool                                     `json:"public_network_access_enabled"`
	ResourceGroupName                     string                                   `json:"resource_group_name"`
	Sku                                   string                                   `json:"sku"`
	Tags                                  map[string]string                        `json:"tags"`
	ZoneRedundant                         bool                                     `json:"zone_redundant"`
	NetworkRulesets                       []eventhubnamespace.NetworkRulesetsState `json:"network_rulesets"`
	Identity                              []eventhubnamespace.IdentityState        `json:"identity"`
	Timeouts                              *eventhubnamespace.TimeoutsState         `json:"timeouts"`
}
