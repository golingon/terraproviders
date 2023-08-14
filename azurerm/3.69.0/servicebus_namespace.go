// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	servicebusnamespace "github.com/golingon/terraproviders/azurerm/3.69.0/servicebusnamespace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicebusNamespace creates a new instance of [ServicebusNamespace].
func NewServicebusNamespace(name string, args ServicebusNamespaceArgs) *ServicebusNamespace {
	return &ServicebusNamespace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicebusNamespace)(nil)

// ServicebusNamespace represents the Terraform resource azurerm_servicebus_namespace.
type ServicebusNamespace struct {
	Name      string
	Args      ServicebusNamespaceArgs
	state     *servicebusNamespaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicebusNamespace].
func (sn *ServicebusNamespace) Type() string {
	return "azurerm_servicebus_namespace"
}

// LocalName returns the local name for [ServicebusNamespace].
func (sn *ServicebusNamespace) LocalName() string {
	return sn.Name
}

// Configuration returns the configuration (args) for [ServicebusNamespace].
func (sn *ServicebusNamespace) Configuration() interface{} {
	return sn.Args
}

// DependOn is used for other resources to depend on [ServicebusNamespace].
func (sn *ServicebusNamespace) DependOn() terra.Reference {
	return terra.ReferenceResource(sn)
}

// Dependencies returns the list of resources [ServicebusNamespace] depends_on.
func (sn *ServicebusNamespace) Dependencies() terra.Dependencies {
	return sn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicebusNamespace].
func (sn *ServicebusNamespace) LifecycleManagement() *terra.Lifecycle {
	return sn.Lifecycle
}

// Attributes returns the attributes for [ServicebusNamespace].
func (sn *ServicebusNamespace) Attributes() servicebusNamespaceAttributes {
	return servicebusNamespaceAttributes{ref: terra.ReferenceResource(sn)}
}

// ImportState imports the given attribute values into [ServicebusNamespace]'s state.
func (sn *ServicebusNamespace) ImportState(av io.Reader) error {
	sn.state = &servicebusNamespaceState{}
	if err := json.NewDecoder(av).Decode(sn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sn.Type(), sn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicebusNamespace] has state.
func (sn *ServicebusNamespace) State() (*servicebusNamespaceState, bool) {
	return sn.state, sn.state != nil
}

// StateMust returns the state for [ServicebusNamespace]. Panics if the state is nil.
func (sn *ServicebusNamespace) StateMust() *servicebusNamespaceState {
	if sn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sn.Type(), sn.LocalName()))
	}
	return sn.state
}

// ServicebusNamespaceArgs contains the configurations for azurerm_servicebus_namespace.
type ServicebusNamespaceArgs struct {
	// Capacity: number, optional
	Capacity terra.NumberValue `hcl:"capacity,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalAuthEnabled: bool, optional
	LocalAuthEnabled terra.BoolValue `hcl:"local_auth_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
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
	// CustomerManagedKey: optional
	CustomerManagedKey *servicebusnamespace.CustomerManagedKey `hcl:"customer_managed_key,block"`
	// Identity: optional
	Identity *servicebusnamespace.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *servicebusnamespace.Timeouts `hcl:"timeouts,block"`
}
type servicebusNamespaceAttributes struct {
	ref terra.Reference
}

// Capacity returns a reference to field capacity of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(sn.ref.Append("capacity"))
}

// DefaultPrimaryConnectionString returns a reference to field default_primary_connection_string of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) DefaultPrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("default_primary_connection_string"))
}

// DefaultPrimaryKey returns a reference to field default_primary_key of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) DefaultPrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("default_primary_key"))
}

// DefaultSecondaryConnectionString returns a reference to field default_secondary_connection_string of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) DefaultSecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("default_secondary_connection_string"))
}

// DefaultSecondaryKey returns a reference to field default_secondary_key of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) DefaultSecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("default_secondary_key"))
}

// Endpoint returns a reference to field endpoint of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("endpoint"))
}

// Id returns a reference to field id of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("id"))
}

// LocalAuthEnabled returns a reference to field local_auth_enabled of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) LocalAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sn.ref.Append("local_auth_enabled"))
}

// Location returns a reference to field location of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("location"))
}

// MinimumTlsVersion returns a reference to field minimum_tls_version of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) MinimumTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("minimum_tls_version"))
}

// Name returns a reference to field name of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sn.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sn.ref.Append("tags"))
}

// ZoneRedundant returns a reference to field zone_redundant of azurerm_servicebus_namespace.
func (sn servicebusNamespaceAttributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(sn.ref.Append("zone_redundant"))
}

func (sn servicebusNamespaceAttributes) CustomerManagedKey() terra.ListValue[servicebusnamespace.CustomerManagedKeyAttributes] {
	return terra.ReferenceAsList[servicebusnamespace.CustomerManagedKeyAttributes](sn.ref.Append("customer_managed_key"))
}

func (sn servicebusNamespaceAttributes) Identity() terra.ListValue[servicebusnamespace.IdentityAttributes] {
	return terra.ReferenceAsList[servicebusnamespace.IdentityAttributes](sn.ref.Append("identity"))
}

func (sn servicebusNamespaceAttributes) Timeouts() servicebusnamespace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicebusnamespace.TimeoutsAttributes](sn.ref.Append("timeouts"))
}

type servicebusNamespaceState struct {
	Capacity                         float64                                       `json:"capacity"`
	DefaultPrimaryConnectionString   string                                        `json:"default_primary_connection_string"`
	DefaultPrimaryKey                string                                        `json:"default_primary_key"`
	DefaultSecondaryConnectionString string                                        `json:"default_secondary_connection_string"`
	DefaultSecondaryKey              string                                        `json:"default_secondary_key"`
	Endpoint                         string                                        `json:"endpoint"`
	Id                               string                                        `json:"id"`
	LocalAuthEnabled                 bool                                          `json:"local_auth_enabled"`
	Location                         string                                        `json:"location"`
	MinimumTlsVersion                string                                        `json:"minimum_tls_version"`
	Name                             string                                        `json:"name"`
	PublicNetworkAccessEnabled       bool                                          `json:"public_network_access_enabled"`
	ResourceGroupName                string                                        `json:"resource_group_name"`
	Sku                              string                                        `json:"sku"`
	Tags                             map[string]string                             `json:"tags"`
	ZoneRedundant                    bool                                          `json:"zone_redundant"`
	CustomerManagedKey               []servicebusnamespace.CustomerManagedKeyState `json:"customer_managed_key"`
	Identity                         []servicebusnamespace.IdentityState           `json:"identity"`
	Timeouts                         *servicebusnamespace.TimeoutsState            `json:"timeouts"`
}
