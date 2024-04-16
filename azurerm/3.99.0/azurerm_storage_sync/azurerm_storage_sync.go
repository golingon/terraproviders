// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_storage_sync

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

// Resource represents the Terraform resource azurerm_storage_sync.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermStorageSyncState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ass *Resource) Type() string {
	return "azurerm_storage_sync"
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
func (ass *Resource) Attributes() azurermStorageSyncAttributes {
	return azurermStorageSyncAttributes{ref: terra.ReferenceResource(ass)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ass *Resource) ImportState(state io.Reader) error {
	ass.state = &azurermStorageSyncState{}
	if err := json.NewDecoder(state).Decode(ass.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ass.Type(), ass.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ass *Resource) State() (*azurermStorageSyncState, bool) {
	return ass.state, ass.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ass *Resource) StateMust() *azurermStorageSyncState {
	if ass.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ass.Type(), ass.LocalName()))
	}
	return ass.state
}

// Args contains the configurations for azurerm_storage_sync.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncomingTrafficPolicy: string, optional
	IncomingTrafficPolicy terra.StringValue `hcl:"incoming_traffic_policy,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermStorageSyncAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_storage_sync.
func (ass azurermStorageSyncAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("id"))
}

// IncomingTrafficPolicy returns a reference to field incoming_traffic_policy of azurerm_storage_sync.
func (ass azurermStorageSyncAttributes) IncomingTrafficPolicy() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("incoming_traffic_policy"))
}

// Location returns a reference to field location of azurerm_storage_sync.
func (ass azurermStorageSyncAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_storage_sync.
func (ass azurermStorageSyncAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_storage_sync.
func (ass azurermStorageSyncAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_storage_sync.
func (ass azurermStorageSyncAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ass.ref.Append("tags"))
}

func (ass azurermStorageSyncAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ass.ref.Append("timeouts"))
}

type azurermStorageSyncState struct {
	Id                    string            `json:"id"`
	IncomingTrafficPolicy string            `json:"incoming_traffic_policy"`
	Location              string            `json:"location"`
	Name                  string            `json:"name"`
	ResourceGroupName     string            `json:"resource_group_name"`
	Tags                  map[string]string `json:"tags"`
	Timeouts              *TimeoutsState    `json:"timeouts"`
}
