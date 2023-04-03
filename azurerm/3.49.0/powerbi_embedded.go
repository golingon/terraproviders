// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	powerbiembedded "github.com/golingon/terraproviders/azurerm/3.49.0/powerbiembedded"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPowerbiEmbedded creates a new instance of [PowerbiEmbedded].
func NewPowerbiEmbedded(name string, args PowerbiEmbeddedArgs) *PowerbiEmbedded {
	return &PowerbiEmbedded{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PowerbiEmbedded)(nil)

// PowerbiEmbedded represents the Terraform resource azurerm_powerbi_embedded.
type PowerbiEmbedded struct {
	Name      string
	Args      PowerbiEmbeddedArgs
	state     *powerbiEmbeddedState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PowerbiEmbedded].
func (pe *PowerbiEmbedded) Type() string {
	return "azurerm_powerbi_embedded"
}

// LocalName returns the local name for [PowerbiEmbedded].
func (pe *PowerbiEmbedded) LocalName() string {
	return pe.Name
}

// Configuration returns the configuration (args) for [PowerbiEmbedded].
func (pe *PowerbiEmbedded) Configuration() interface{} {
	return pe.Args
}

// DependOn is used for other resources to depend on [PowerbiEmbedded].
func (pe *PowerbiEmbedded) DependOn() terra.Reference {
	return terra.ReferenceResource(pe)
}

// Dependencies returns the list of resources [PowerbiEmbedded] depends_on.
func (pe *PowerbiEmbedded) Dependencies() terra.Dependencies {
	return pe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PowerbiEmbedded].
func (pe *PowerbiEmbedded) LifecycleManagement() *terra.Lifecycle {
	return pe.Lifecycle
}

// Attributes returns the attributes for [PowerbiEmbedded].
func (pe *PowerbiEmbedded) Attributes() powerbiEmbeddedAttributes {
	return powerbiEmbeddedAttributes{ref: terra.ReferenceResource(pe)}
}

// ImportState imports the given attribute values into [PowerbiEmbedded]'s state.
func (pe *PowerbiEmbedded) ImportState(av io.Reader) error {
	pe.state = &powerbiEmbeddedState{}
	if err := json.NewDecoder(av).Decode(pe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pe.Type(), pe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PowerbiEmbedded] has state.
func (pe *PowerbiEmbedded) State() (*powerbiEmbeddedState, bool) {
	return pe.state, pe.state != nil
}

// StateMust returns the state for [PowerbiEmbedded]. Panics if the state is nil.
func (pe *PowerbiEmbedded) StateMust() *powerbiEmbeddedState {
	if pe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pe.Type(), pe.LocalName()))
	}
	return pe.state
}

// PowerbiEmbeddedArgs contains the configurations for azurerm_powerbi_embedded.
type PowerbiEmbeddedArgs struct {
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
	Timeouts *powerbiembedded.Timeouts `hcl:"timeouts,block"`
}
type powerbiEmbeddedAttributes struct {
	ref terra.Reference
}

// Administrators returns a reference to field administrators of azurerm_powerbi_embedded.
func (pe powerbiEmbeddedAttributes) Administrators() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pe.ref.Append("administrators"))
}

// Id returns a reference to field id of azurerm_powerbi_embedded.
func (pe powerbiEmbeddedAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_powerbi_embedded.
func (pe powerbiEmbeddedAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("location"))
}

// Mode returns a reference to field mode of azurerm_powerbi_embedded.
func (pe powerbiEmbeddedAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_powerbi_embedded.
func (pe powerbiEmbeddedAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_powerbi_embedded.
func (pe powerbiEmbeddedAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_powerbi_embedded.
func (pe powerbiEmbeddedAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_powerbi_embedded.
func (pe powerbiEmbeddedAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pe.ref.Append("tags"))
}

func (pe powerbiEmbeddedAttributes) Timeouts() powerbiembedded.TimeoutsAttributes {
	return terra.ReferenceAsSingle[powerbiembedded.TimeoutsAttributes](pe.ref.Append("timeouts"))
}

type powerbiEmbeddedState struct {
	Administrators    []string                       `json:"administrators"`
	Id                string                         `json:"id"`
	Location          string                         `json:"location"`
	Mode              string                         `json:"mode"`
	Name              string                         `json:"name"`
	ResourceGroupName string                         `json:"resource_group_name"`
	SkuName           string                         `json:"sku_name"`
	Tags              map[string]string              `json:"tags"`
	Timeouts          *powerbiembedded.TimeoutsState `json:"timeouts"`
}
