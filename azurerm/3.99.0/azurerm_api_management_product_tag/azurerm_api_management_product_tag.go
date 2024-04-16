// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_api_management_product_tag

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

// Resource represents the Terraform resource azurerm_api_management_product_tag.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermApiManagementProductTagState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aampt *Resource) Type() string {
	return "azurerm_api_management_product_tag"
}

// LocalName returns the local name for [Resource].
func (aampt *Resource) LocalName() string {
	return aampt.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aampt *Resource) Configuration() interface{} {
	return aampt.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aampt *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aampt)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aampt *Resource) Dependencies() terra.Dependencies {
	return aampt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aampt *Resource) LifecycleManagement() *terra.Lifecycle {
	return aampt.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aampt *Resource) Attributes() azurermApiManagementProductTagAttributes {
	return azurermApiManagementProductTagAttributes{ref: terra.ReferenceResource(aampt)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aampt *Resource) ImportState(state io.Reader) error {
	aampt.state = &azurermApiManagementProductTagState{}
	if err := json.NewDecoder(state).Decode(aampt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aampt.Type(), aampt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aampt *Resource) State() (*azurermApiManagementProductTagState, bool) {
	return aampt.state, aampt.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aampt *Resource) StateMust() *azurermApiManagementProductTagState {
	if aampt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aampt.Type(), aampt.LocalName()))
	}
	return aampt.state
}

// Args contains the configurations for azurerm_api_management_product_tag.
type Args struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// ApiManagementProductId: string, required
	ApiManagementProductId terra.StringValue `hcl:"api_management_product_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermApiManagementProductTagAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_product_tag.
func (aampt azurermApiManagementProductTagAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(aampt.ref.Append("api_management_name"))
}

// ApiManagementProductId returns a reference to field api_management_product_id of azurerm_api_management_product_tag.
func (aampt azurermApiManagementProductTagAttributes) ApiManagementProductId() terra.StringValue {
	return terra.ReferenceAsString(aampt.ref.Append("api_management_product_id"))
}

// Id returns a reference to field id of azurerm_api_management_product_tag.
func (aampt azurermApiManagementProductTagAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aampt.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_product_tag.
func (aampt azurermApiManagementProductTagAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aampt.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_product_tag.
func (aampt azurermApiManagementProductTagAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aampt.ref.Append("resource_group_name"))
}

func (aampt azurermApiManagementProductTagAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aampt.ref.Append("timeouts"))
}

type azurermApiManagementProductTagState struct {
	ApiManagementName      string         `json:"api_management_name"`
	ApiManagementProductId string         `json:"api_management_product_id"`
	Id                     string         `json:"id"`
	Name                   string         `json:"name"`
	ResourceGroupName      string         `json:"resource_group_name"`
	Timeouts               *TimeoutsState `json:"timeouts"`
}
