// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	fluidrelayserver "github.com/golingon/terraproviders/azurerm/3.49.0/fluidrelayserver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFluidRelayServer creates a new instance of [FluidRelayServer].
func NewFluidRelayServer(name string, args FluidRelayServerArgs) *FluidRelayServer {
	return &FluidRelayServer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FluidRelayServer)(nil)

// FluidRelayServer represents the Terraform resource azurerm_fluid_relay_server.
type FluidRelayServer struct {
	Name      string
	Args      FluidRelayServerArgs
	state     *fluidRelayServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FluidRelayServer].
func (frs *FluidRelayServer) Type() string {
	return "azurerm_fluid_relay_server"
}

// LocalName returns the local name for [FluidRelayServer].
func (frs *FluidRelayServer) LocalName() string {
	return frs.Name
}

// Configuration returns the configuration (args) for [FluidRelayServer].
func (frs *FluidRelayServer) Configuration() interface{} {
	return frs.Args
}

// DependOn is used for other resources to depend on [FluidRelayServer].
func (frs *FluidRelayServer) DependOn() terra.Reference {
	return terra.ReferenceResource(frs)
}

// Dependencies returns the list of resources [FluidRelayServer] depends_on.
func (frs *FluidRelayServer) Dependencies() terra.Dependencies {
	return frs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FluidRelayServer].
func (frs *FluidRelayServer) LifecycleManagement() *terra.Lifecycle {
	return frs.Lifecycle
}

// Attributes returns the attributes for [FluidRelayServer].
func (frs *FluidRelayServer) Attributes() fluidRelayServerAttributes {
	return fluidRelayServerAttributes{ref: terra.ReferenceResource(frs)}
}

// ImportState imports the given attribute values into [FluidRelayServer]'s state.
func (frs *FluidRelayServer) ImportState(av io.Reader) error {
	frs.state = &fluidRelayServerState{}
	if err := json.NewDecoder(av).Decode(frs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", frs.Type(), frs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FluidRelayServer] has state.
func (frs *FluidRelayServer) State() (*fluidRelayServerState, bool) {
	return frs.state, frs.state != nil
}

// StateMust returns the state for [FluidRelayServer]. Panics if the state is nil.
func (frs *FluidRelayServer) StateMust() *fluidRelayServerState {
	if frs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", frs.Type(), frs.LocalName()))
	}
	return frs.state
}

// FluidRelayServerArgs contains the configurations for azurerm_fluid_relay_server.
type FluidRelayServerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageSku: string, optional
	StorageSku terra.StringValue `hcl:"storage_sku,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: optional
	Identity *fluidrelayserver.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *fluidrelayserver.Timeouts `hcl:"timeouts,block"`
}
type fluidRelayServerAttributes struct {
	ref terra.Reference
}

// FrsTenantId returns a reference to field frs_tenant_id of azurerm_fluid_relay_server.
func (frs fluidRelayServerAttributes) FrsTenantId() terra.StringValue {
	return terra.ReferenceAsString(frs.ref.Append("frs_tenant_id"))
}

// Id returns a reference to field id of azurerm_fluid_relay_server.
func (frs fluidRelayServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(frs.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_fluid_relay_server.
func (frs fluidRelayServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(frs.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_fluid_relay_server.
func (frs fluidRelayServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(frs.ref.Append("name"))
}

// OrdererEndpoints returns a reference to field orderer_endpoints of azurerm_fluid_relay_server.
func (frs fluidRelayServerAttributes) OrdererEndpoints() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](frs.ref.Append("orderer_endpoints"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_fluid_relay_server.
func (frs fluidRelayServerAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(frs.ref.Append("primary_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_fluid_relay_server.
func (frs fluidRelayServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(frs.ref.Append("resource_group_name"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_fluid_relay_server.
func (frs fluidRelayServerAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(frs.ref.Append("secondary_key"))
}

// ServiceEndpoints returns a reference to field service_endpoints of azurerm_fluid_relay_server.
func (frs fluidRelayServerAttributes) ServiceEndpoints() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](frs.ref.Append("service_endpoints"))
}

// StorageEndpoints returns a reference to field storage_endpoints of azurerm_fluid_relay_server.
func (frs fluidRelayServerAttributes) StorageEndpoints() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](frs.ref.Append("storage_endpoints"))
}

// StorageSku returns a reference to field storage_sku of azurerm_fluid_relay_server.
func (frs fluidRelayServerAttributes) StorageSku() terra.StringValue {
	return terra.ReferenceAsString(frs.ref.Append("storage_sku"))
}

// Tags returns a reference to field tags of azurerm_fluid_relay_server.
func (frs fluidRelayServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](frs.ref.Append("tags"))
}

func (frs fluidRelayServerAttributes) Identity() terra.ListValue[fluidrelayserver.IdentityAttributes] {
	return terra.ReferenceAsList[fluidrelayserver.IdentityAttributes](frs.ref.Append("identity"))
}

func (frs fluidRelayServerAttributes) Timeouts() fluidrelayserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[fluidrelayserver.TimeoutsAttributes](frs.ref.Append("timeouts"))
}

type fluidRelayServerState struct {
	FrsTenantId       string                           `json:"frs_tenant_id"`
	Id                string                           `json:"id"`
	Location          string                           `json:"location"`
	Name              string                           `json:"name"`
	OrdererEndpoints  []string                         `json:"orderer_endpoints"`
	PrimaryKey        string                           `json:"primary_key"`
	ResourceGroupName string                           `json:"resource_group_name"`
	SecondaryKey      string                           `json:"secondary_key"`
	ServiceEndpoints  []string                         `json:"service_endpoints"`
	StorageEndpoints  []string                         `json:"storage_endpoints"`
	StorageSku        string                           `json:"storage_sku"`
	Tags              map[string]string                `json:"tags"`
	Identity          []fluidrelayserver.IdentityState `json:"identity"`
	Timeouts          *fluidrelayserver.TimeoutsState  `json:"timeouts"`
}
