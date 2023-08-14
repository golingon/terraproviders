// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	eventhubcluster "github.com/golingon/terraproviders/azurerm/3.69.0/eventhubcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventhubCluster creates a new instance of [EventhubCluster].
func NewEventhubCluster(name string, args EventhubClusterArgs) *EventhubCluster {
	return &EventhubCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventhubCluster)(nil)

// EventhubCluster represents the Terraform resource azurerm_eventhub_cluster.
type EventhubCluster struct {
	Name      string
	Args      EventhubClusterArgs
	state     *eventhubClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventhubCluster].
func (ec *EventhubCluster) Type() string {
	return "azurerm_eventhub_cluster"
}

// LocalName returns the local name for [EventhubCluster].
func (ec *EventhubCluster) LocalName() string {
	return ec.Name
}

// Configuration returns the configuration (args) for [EventhubCluster].
func (ec *EventhubCluster) Configuration() interface{} {
	return ec.Args
}

// DependOn is used for other resources to depend on [EventhubCluster].
func (ec *EventhubCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(ec)
}

// Dependencies returns the list of resources [EventhubCluster] depends_on.
func (ec *EventhubCluster) Dependencies() terra.Dependencies {
	return ec.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventhubCluster].
func (ec *EventhubCluster) LifecycleManagement() *terra.Lifecycle {
	return ec.Lifecycle
}

// Attributes returns the attributes for [EventhubCluster].
func (ec *EventhubCluster) Attributes() eventhubClusterAttributes {
	return eventhubClusterAttributes{ref: terra.ReferenceResource(ec)}
}

// ImportState imports the given attribute values into [EventhubCluster]'s state.
func (ec *EventhubCluster) ImportState(av io.Reader) error {
	ec.state = &eventhubClusterState{}
	if err := json.NewDecoder(av).Decode(ec.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ec.Type(), ec.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventhubCluster] has state.
func (ec *EventhubCluster) State() (*eventhubClusterState, bool) {
	return ec.state, ec.state != nil
}

// StateMust returns the state for [EventhubCluster]. Panics if the state is nil.
func (ec *EventhubCluster) StateMust() *eventhubClusterState {
	if ec.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ec.Type(), ec.LocalName()))
	}
	return ec.state
}

// EventhubClusterArgs contains the configurations for azurerm_eventhub_cluster.
type EventhubClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *eventhubcluster.Timeouts `hcl:"timeouts,block"`
}
type eventhubClusterAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_eventhub_cluster.
func (ec eventhubClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_eventhub_cluster.
func (ec eventhubClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_eventhub_cluster.
func (ec eventhubClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventhub_cluster.
func (ec eventhubClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_eventhub_cluster.
func (ec eventhubClusterAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_eventhub_cluster.
func (ec eventhubClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ec.ref.Append("tags"))
}

func (ec eventhubClusterAttributes) Timeouts() eventhubcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventhubcluster.TimeoutsAttributes](ec.ref.Append("timeouts"))
}

type eventhubClusterState struct {
	Id                string                         `json:"id"`
	Location          string                         `json:"location"`
	Name              string                         `json:"name"`
	ResourceGroupName string                         `json:"resource_group_name"`
	SkuName           string                         `json:"sku_name"`
	Tags              map[string]string              `json:"tags"`
	Timeouts          *eventhubcluster.TimeoutsState `json:"timeouts"`
}
