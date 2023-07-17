// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementproduct "github.com/golingon/terraproviders/azurerm/3.65.0/apimanagementproduct"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementProduct creates a new instance of [ApiManagementProduct].
func NewApiManagementProduct(name string, args ApiManagementProductArgs) *ApiManagementProduct {
	return &ApiManagementProduct{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementProduct)(nil)

// ApiManagementProduct represents the Terraform resource azurerm_api_management_product.
type ApiManagementProduct struct {
	Name      string
	Args      ApiManagementProductArgs
	state     *apiManagementProductState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementProduct].
func (amp *ApiManagementProduct) Type() string {
	return "azurerm_api_management_product"
}

// LocalName returns the local name for [ApiManagementProduct].
func (amp *ApiManagementProduct) LocalName() string {
	return amp.Name
}

// Configuration returns the configuration (args) for [ApiManagementProduct].
func (amp *ApiManagementProduct) Configuration() interface{} {
	return amp.Args
}

// DependOn is used for other resources to depend on [ApiManagementProduct].
func (amp *ApiManagementProduct) DependOn() terra.Reference {
	return terra.ReferenceResource(amp)
}

// Dependencies returns the list of resources [ApiManagementProduct] depends_on.
func (amp *ApiManagementProduct) Dependencies() terra.Dependencies {
	return amp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementProduct].
func (amp *ApiManagementProduct) LifecycleManagement() *terra.Lifecycle {
	return amp.Lifecycle
}

// Attributes returns the attributes for [ApiManagementProduct].
func (amp *ApiManagementProduct) Attributes() apiManagementProductAttributes {
	return apiManagementProductAttributes{ref: terra.ReferenceResource(amp)}
}

// ImportState imports the given attribute values into [ApiManagementProduct]'s state.
func (amp *ApiManagementProduct) ImportState(av io.Reader) error {
	amp.state = &apiManagementProductState{}
	if err := json.NewDecoder(av).Decode(amp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amp.Type(), amp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementProduct] has state.
func (amp *ApiManagementProduct) State() (*apiManagementProductState, bool) {
	return amp.state, amp.state != nil
}

// StateMust returns the state for [ApiManagementProduct]. Panics if the state is nil.
func (amp *ApiManagementProduct) StateMust() *apiManagementProductState {
	if amp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amp.Type(), amp.LocalName()))
	}
	return amp.state
}

// ApiManagementProductArgs contains the configurations for azurerm_api_management_product.
type ApiManagementProductArgs struct {
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
	Timeouts *apimanagementproduct.Timeouts `hcl:"timeouts,block"`
}
type apiManagementProductAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_product.
func (amp apiManagementProductAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("api_management_name"))
}

// ApprovalRequired returns a reference to field approval_required of azurerm_api_management_product.
func (amp apiManagementProductAttributes) ApprovalRequired() terra.BoolValue {
	return terra.ReferenceAsBool(amp.ref.Append("approval_required"))
}

// Description returns a reference to field description of azurerm_api_management_product.
func (amp apiManagementProductAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_product.
func (amp apiManagementProductAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_product.
func (amp apiManagementProductAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("id"))
}

// ProductId returns a reference to field product_id of azurerm_api_management_product.
func (amp apiManagementProductAttributes) ProductId() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("product_id"))
}

// Published returns a reference to field published of azurerm_api_management_product.
func (amp apiManagementProductAttributes) Published() terra.BoolValue {
	return terra.ReferenceAsBool(amp.ref.Append("published"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_product.
func (amp apiManagementProductAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("resource_group_name"))
}

// SubscriptionRequired returns a reference to field subscription_required of azurerm_api_management_product.
func (amp apiManagementProductAttributes) SubscriptionRequired() terra.BoolValue {
	return terra.ReferenceAsBool(amp.ref.Append("subscription_required"))
}

// SubscriptionsLimit returns a reference to field subscriptions_limit of azurerm_api_management_product.
func (amp apiManagementProductAttributes) SubscriptionsLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(amp.ref.Append("subscriptions_limit"))
}

// Terms returns a reference to field terms of azurerm_api_management_product.
func (amp apiManagementProductAttributes) Terms() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("terms"))
}

func (amp apiManagementProductAttributes) Timeouts() apimanagementproduct.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementproduct.TimeoutsAttributes](amp.ref.Append("timeouts"))
}

type apiManagementProductState struct {
	ApiManagementName    string                              `json:"api_management_name"`
	ApprovalRequired     bool                                `json:"approval_required"`
	Description          string                              `json:"description"`
	DisplayName          string                              `json:"display_name"`
	Id                   string                              `json:"id"`
	ProductId            string                              `json:"product_id"`
	Published            bool                                `json:"published"`
	ResourceGroupName    string                              `json:"resource_group_name"`
	SubscriptionRequired bool                                `json:"subscription_required"`
	SubscriptionsLimit   float64                             `json:"subscriptions_limit"`
	Terms                string                              `json:"terms"`
	Timeouts             *apimanagementproduct.TimeoutsState `json:"timeouts"`
}
