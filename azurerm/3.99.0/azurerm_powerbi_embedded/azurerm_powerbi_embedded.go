// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_powerbi_embedded

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

// Resource represents the Terraform resource azurerm_powerbi_embedded.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermPowerbiEmbeddedState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ape *Resource) Type() string {
	return "azurerm_powerbi_embedded"
}

// LocalName returns the local name for [Resource].
func (ape *Resource) LocalName() string {
	return ape.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ape *Resource) Configuration() interface{} {
	return ape.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ape *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ape)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ape *Resource) Dependencies() terra.Dependencies {
	return ape.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ape *Resource) LifecycleManagement() *terra.Lifecycle {
	return ape.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ape *Resource) Attributes() azurermPowerbiEmbeddedAttributes {
	return azurermPowerbiEmbeddedAttributes{ref: terra.ReferenceResource(ape)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ape *Resource) ImportState(state io.Reader) error {
	ape.state = &azurermPowerbiEmbeddedState{}
	if err := json.NewDecoder(state).Decode(ape.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ape.Type(), ape.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ape *Resource) State() (*azurermPowerbiEmbeddedState, bool) {
	return ape.state, ape.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ape *Resource) StateMust() *azurermPowerbiEmbeddedState {
	if ape.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ape.Type(), ape.LocalName()))
	}
	return ape.state
}

// Args contains the configurations for azurerm_powerbi_embedded.
type Args struct {
	// Administrators: set of string, required
	Administrators terra.SetValue[terra.StringValue] `hcl:"administrators,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
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

type azurermPowerbiEmbeddedAttributes struct {
	ref terra.Reference
}

// Administrators returns a reference to field administrators of azurerm_powerbi_embedded.
func (ape azurermPowerbiEmbeddedAttributes) Administrators() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ape.ref.Append("administrators"))
}

// Id returns a reference to field id of azurerm_powerbi_embedded.
func (ape azurermPowerbiEmbeddedAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ape.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_powerbi_embedded.
func (ape azurermPowerbiEmbeddedAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ape.ref.Append("location"))
}

// Mode returns a reference to field mode of azurerm_powerbi_embedded.
func (ape azurermPowerbiEmbeddedAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(ape.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_powerbi_embedded.
func (ape azurermPowerbiEmbeddedAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ape.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_powerbi_embedded.
func (ape azurermPowerbiEmbeddedAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ape.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_powerbi_embedded.
func (ape azurermPowerbiEmbeddedAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ape.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_powerbi_embedded.
func (ape azurermPowerbiEmbeddedAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ape.ref.Append("tags"))
}

func (ape azurermPowerbiEmbeddedAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ape.ref.Append("timeouts"))
}

type azurermPowerbiEmbeddedState struct {
	Administrators    []string          `json:"administrators"`
	Id                string            `json:"id"`
	Location          string            `json:"location"`
	Mode              string            `json:"mode"`
	Name              string            `json:"name"`
	ResourceGroupName string            `json:"resource_group_name"`
	SkuName           string            `json:"sku_name"`
	Tags              map[string]string `json:"tags"`
	Timeouts          *TimeoutsState    `json:"timeouts"`
}
