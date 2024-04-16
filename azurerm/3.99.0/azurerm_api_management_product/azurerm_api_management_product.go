// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_api_management_product

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

// Resource represents the Terraform resource azurerm_api_management_product.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermApiManagementProductState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aamp *Resource) Type() string {
	return "azurerm_api_management_product"
}

// LocalName returns the local name for [Resource].
func (aamp *Resource) LocalName() string {
	return aamp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aamp *Resource) Configuration() interface{} {
	return aamp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aamp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aamp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aamp *Resource) Dependencies() terra.Dependencies {
	return aamp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aamp *Resource) LifecycleManagement() *terra.Lifecycle {
	return aamp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aamp *Resource) Attributes() azurermApiManagementProductAttributes {
	return azurermApiManagementProductAttributes{ref: terra.ReferenceResource(aamp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aamp *Resource) ImportState(state io.Reader) error {
	aamp.state = &azurermApiManagementProductState{}
	if err := json.NewDecoder(state).Decode(aamp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aamp.Type(), aamp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aamp *Resource) State() (*azurermApiManagementProductState, bool) {
	return aamp.state, aamp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aamp *Resource) StateMust() *azurermApiManagementProductState {
	if aamp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aamp.Type(), aamp.LocalName()))
	}
	return aamp.state
}

// Args contains the configurations for azurerm_api_management_product.
type Args struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// ApprovalRequired: bool, optional
	ApprovalRequired terra.BoolValue `hcl:"approval_required,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProductId: string, required
	ProductId terra.StringValue `hcl:"product_id,attr" validate:"required"`
	// Published: bool, required
	Published terra.BoolValue `hcl:"published,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SubscriptionRequired: bool, optional
	SubscriptionRequired terra.BoolValue `hcl:"subscription_required,attr"`
	// SubscriptionsLimit: number, optional
	SubscriptionsLimit terra.NumberValue `hcl:"subscriptions_limit,attr"`
	// Terms: string, optional
	Terms terra.StringValue `hcl:"terms,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermApiManagementProductAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_product.
func (aamp azurermApiManagementProductAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(aamp.ref.Append("api_management_name"))
}

// ApprovalRequired returns a reference to field approval_required of azurerm_api_management_product.
func (aamp azurermApiManagementProductAttributes) ApprovalRequired() terra.BoolValue {
	return terra.ReferenceAsBool(aamp.ref.Append("approval_required"))
}

// Description returns a reference to field description of azurerm_api_management_product.
func (aamp azurermApiManagementProductAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aamp.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_product.
func (aamp azurermApiManagementProductAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aamp.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_product.
func (aamp azurermApiManagementProductAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aamp.ref.Append("id"))
}

// ProductId returns a reference to field product_id of azurerm_api_management_product.
func (aamp azurermApiManagementProductAttributes) ProductId() terra.StringValue {
	return terra.ReferenceAsString(aamp.ref.Append("product_id"))
}

// Published returns a reference to field published of azurerm_api_management_product.
func (aamp azurermApiManagementProductAttributes) Published() terra.BoolValue {
	return terra.ReferenceAsBool(aamp.ref.Append("published"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_product.
func (aamp azurermApiManagementProductAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aamp.ref.Append("resource_group_name"))
}

// SubscriptionRequired returns a reference to field subscription_required of azurerm_api_management_product.
func (aamp azurermApiManagementProductAttributes) SubscriptionRequired() terra.BoolValue {
	return terra.ReferenceAsBool(aamp.ref.Append("subscription_required"))
}

// SubscriptionsLimit returns a reference to field subscriptions_limit of azurerm_api_management_product.
func (aamp azurermApiManagementProductAttributes) SubscriptionsLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(aamp.ref.Append("subscriptions_limit"))
}

// Terms returns a reference to field terms of azurerm_api_management_product.
func (aamp azurermApiManagementProductAttributes) Terms() terra.StringValue {
	return terra.ReferenceAsString(aamp.ref.Append("terms"))
}

func (aamp azurermApiManagementProductAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aamp.ref.Append("timeouts"))
}

type azurermApiManagementProductState struct {
	ApiManagementName    string         `json:"api_management_name"`
	ApprovalRequired     bool           `json:"approval_required"`
	Description          string         `json:"description"`
	DisplayName          string         `json:"display_name"`
	Id                   string         `json:"id"`
	ProductId            string         `json:"product_id"`
	Published            bool           `json:"published"`
	ResourceGroupName    string         `json:"resource_group_name"`
	SubscriptionRequired bool           `json:"subscription_required"`
	SubscriptionsLimit   float64        `json:"subscriptions_limit"`
	Terms                string         `json:"terms"`
	Timeouts             *TimeoutsState `json:"timeouts"`
}
