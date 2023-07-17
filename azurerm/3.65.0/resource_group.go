// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	resourcegroup "github.com/golingon/terraproviders/azurerm/3.65.0/resourcegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewResourceGroup creates a new instance of [ResourceGroup].
func NewResourceGroup(name string, args ResourceGroupArgs) *ResourceGroup {
	return &ResourceGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ResourceGroup)(nil)

// ResourceGroup represents the Terraform resource azurerm_resource_group.
type ResourceGroup struct {
	Name      string
	Args      ResourceGroupArgs
	state     *resourceGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ResourceGroup].
func (rg *ResourceGroup) Type() string {
	return "azurerm_resource_group"
}

// LocalName returns the local name for [ResourceGroup].
func (rg *ResourceGroup) LocalName() string {
	return rg.Name
}

// Configuration returns the configuration (args) for [ResourceGroup].
func (rg *ResourceGroup) Configuration() interface{} {
	return rg.Args
}

// DependOn is used for other resources to depend on [ResourceGroup].
func (rg *ResourceGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(rg)
}

// Dependencies returns the list of resources [ResourceGroup] depends_on.
func (rg *ResourceGroup) Dependencies() terra.Dependencies {
	return rg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ResourceGroup].
func (rg *ResourceGroup) LifecycleManagement() *terra.Lifecycle {
	return rg.Lifecycle
}

// Attributes returns the attributes for [ResourceGroup].
func (rg *ResourceGroup) Attributes() resourceGroupAttributes {
	return resourceGroupAttributes{ref: terra.ReferenceResource(rg)}
}

// ImportState imports the given attribute values into [ResourceGroup]'s state.
func (rg *ResourceGroup) ImportState(av io.Reader) error {
	rg.state = &resourceGroupState{}
	if err := json.NewDecoder(av).Decode(rg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rg.Type(), rg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ResourceGroup] has state.
func (rg *ResourceGroup) State() (*resourceGroupState, bool) {
	return rg.state, rg.state != nil
}

// StateMust returns the state for [ResourceGroup]. Panics if the state is nil.
func (rg *ResourceGroup) StateMust() *resourceGroupState {
	if rg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rg.Type(), rg.LocalName()))
	}
	return rg.state
}

// ResourceGroupArgs contains the configurations for azurerm_resource_group.
type ResourceGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagedBy: string, optional
	ManagedBy terra.StringValue `hcl:"managed_by,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *resourcegroup.Timeouts `hcl:"timeouts,block"`
}
type resourceGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_resource_group.
func (rg resourceGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_resource_group.
func (rg resourceGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("location"))
}

// ManagedBy returns a reference to field managed_by of azurerm_resource_group.
func (rg resourceGroupAttributes) ManagedBy() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("managed_by"))
}

// Name returns a reference to field name of azurerm_resource_group.
func (rg resourceGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_resource_group.
func (rg resourceGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rg.ref.Append("tags"))
}

func (rg resourceGroupAttributes) Timeouts() resourcegroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[resourcegroup.TimeoutsAttributes](rg.ref.Append("timeouts"))
}

type resourceGroupState struct {
	Id        string                       `json:"id"`
	Location  string                       `json:"location"`
	ManagedBy string                       `json:"managed_by"`
	Name      string                       `json:"name"`
	Tags      map[string]string            `json:"tags"`
	Timeouts  *resourcegroup.TimeoutsState `json:"timeouts"`
}
