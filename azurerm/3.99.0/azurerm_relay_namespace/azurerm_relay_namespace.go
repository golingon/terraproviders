// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_relay_namespace

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

// Resource represents the Terraform resource azurerm_relay_namespace.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermRelayNamespaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (arn *Resource) Type() string {
	return "azurerm_relay_namespace"
}

// LocalName returns the local name for [Resource].
func (arn *Resource) LocalName() string {
	return arn.Name
}

// Configuration returns the configuration (args) for [Resource].
func (arn *Resource) Configuration() interface{} {
	return arn.Args
}

// DependOn is used for other resources to depend on [Resource].
func (arn *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(arn)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (arn *Resource) Dependencies() terra.Dependencies {
	return arn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (arn *Resource) LifecycleManagement() *terra.Lifecycle {
	return arn.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (arn *Resource) Attributes() azurermRelayNamespaceAttributes {
	return azurermRelayNamespaceAttributes{ref: terra.ReferenceResource(arn)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (arn *Resource) ImportState(state io.Reader) error {
	arn.state = &azurermRelayNamespaceState{}
	if err := json.NewDecoder(state).Decode(arn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", arn.Type(), arn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (arn *Resource) State() (*azurermRelayNamespaceState, bool) {
	return arn.state, arn.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (arn *Resource) StateMust() *azurermRelayNamespaceState {
	if arn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", arn.Type(), arn.LocalName()))
	}
	return arn.state
}

// Args contains the configurations for azurerm_relay_namespace.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermRelayNamespaceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_relay_namespace.
func (arn azurermRelayNamespaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_relay_namespace.
func (arn azurermRelayNamespaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("location"))
}

// MetricId returns a reference to field metric_id of azurerm_relay_namespace.
func (arn azurermRelayNamespaceAttributes) MetricId() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("metric_id"))
}

// Name returns a reference to field name of azurerm_relay_namespace.
func (arn azurermRelayNamespaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_relay_namespace.
func (arn azurermRelayNamespaceAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("primary_connection_string"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_relay_namespace.
func (arn azurermRelayNamespaceAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("primary_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_relay_namespace.
func (arn azurermRelayNamespaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("resource_group_name"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_relay_namespace.
func (arn azurermRelayNamespaceAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("secondary_connection_string"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_relay_namespace.
func (arn azurermRelayNamespaceAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("secondary_key"))
}

// SkuName returns a reference to field sku_name of azurerm_relay_namespace.
func (arn azurermRelayNamespaceAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_relay_namespace.
func (arn azurermRelayNamespaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](arn.ref.Append("tags"))
}

func (arn azurermRelayNamespaceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](arn.ref.Append("timeouts"))
}

type azurermRelayNamespaceState struct {
	Id                        string            `json:"id"`
	Location                  string            `json:"location"`
	MetricId                  string            `json:"metric_id"`
	Name                      string            `json:"name"`
	PrimaryConnectionString   string            `json:"primary_connection_string"`
	PrimaryKey                string            `json:"primary_key"`
	ResourceGroupName         string            `json:"resource_group_name"`
	SecondaryConnectionString string            `json:"secondary_connection_string"`
	SecondaryKey              string            `json:"secondary_key"`
	SkuName                   string            `json:"sku_name"`
	Tags                      map[string]string `json:"tags"`
	Timeouts                  *TimeoutsState    `json:"timeouts"`
}
