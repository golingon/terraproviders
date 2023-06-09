// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataapimanagementproduct "github.com/golingon/terraproviders/azurerm/3.64.0/dataapimanagementproduct"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataApiManagementProduct creates a new instance of [DataApiManagementProduct].
func NewDataApiManagementProduct(name string, args DataApiManagementProductArgs) *DataApiManagementProduct {
	return &DataApiManagementProduct{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApiManagementProduct)(nil)

// DataApiManagementProduct represents the Terraform data resource azurerm_api_management_product.
type DataApiManagementProduct struct {
	Name string
	Args DataApiManagementProductArgs
}

// DataSource returns the Terraform object type for [DataApiManagementProduct].
func (amp *DataApiManagementProduct) DataSource() string {
	return "azurerm_api_management_product"
}

// LocalName returns the local name for [DataApiManagementProduct].
func (amp *DataApiManagementProduct) LocalName() string {
	return amp.Name
}

// Configuration returns the configuration (args) for [DataApiManagementProduct].
func (amp *DataApiManagementProduct) Configuration() interface{} {
	return amp.Args
}

// Attributes returns the attributes for [DataApiManagementProduct].
func (amp *DataApiManagementProduct) Attributes() dataApiManagementProductAttributes {
	return dataApiManagementProductAttributes{ref: terra.ReferenceDataResource(amp)}
}

// DataApiManagementProductArgs contains the configurations for azurerm_api_management_product.
type DataApiManagementProductArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProductId: string, required
	ProductId terra.StringValue `hcl:"product_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataapimanagementproduct.Timeouts `hcl:"timeouts,block"`
}
type dataApiManagementProductAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_product.
func (amp dataApiManagementProductAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("api_management_name"))
}

// ApprovalRequired returns a reference to field approval_required of azurerm_api_management_product.
func (amp dataApiManagementProductAttributes) ApprovalRequired() terra.BoolValue {
	return terra.ReferenceAsBool(amp.ref.Append("approval_required"))
}

// Description returns a reference to field description of azurerm_api_management_product.
func (amp dataApiManagementProductAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_product.
func (amp dataApiManagementProductAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_product.
func (amp dataApiManagementProductAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("id"))
}

// ProductId returns a reference to field product_id of azurerm_api_management_product.
func (amp dataApiManagementProductAttributes) ProductId() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("product_id"))
}

// Published returns a reference to field published of azurerm_api_management_product.
func (amp dataApiManagementProductAttributes) Published() terra.BoolValue {
	return terra.ReferenceAsBool(amp.ref.Append("published"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_product.
func (amp dataApiManagementProductAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("resource_group_name"))
}

// SubscriptionRequired returns a reference to field subscription_required of azurerm_api_management_product.
func (amp dataApiManagementProductAttributes) SubscriptionRequired() terra.BoolValue {
	return terra.ReferenceAsBool(amp.ref.Append("subscription_required"))
}

// SubscriptionsLimit returns a reference to field subscriptions_limit of azurerm_api_management_product.
func (amp dataApiManagementProductAttributes) SubscriptionsLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(amp.ref.Append("subscriptions_limit"))
}

// Terms returns a reference to field terms of azurerm_api_management_product.
func (amp dataApiManagementProductAttributes) Terms() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("terms"))
}

func (amp dataApiManagementProductAttributes) Timeouts() dataapimanagementproduct.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataapimanagementproduct.TimeoutsAttributes](amp.ref.Append("timeouts"))
}
