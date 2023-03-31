// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementproduct "github.com/golingon/terraproviders/azurerm/3.49.0/apimanagementproduct"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApiManagementProduct(name string, args ApiManagementProductArgs) *ApiManagementProduct {
	return &ApiManagementProduct{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementProduct)(nil)

type ApiManagementProduct struct {
	Name  string
	Args  ApiManagementProductArgs
	state *apiManagementProductState
}

func (amp *ApiManagementProduct) Type() string {
	return "azurerm_api_management_product"
}

func (amp *ApiManagementProduct) LocalName() string {
	return amp.Name
}

func (amp *ApiManagementProduct) Configuration() interface{} {
	return amp.Args
}

func (amp *ApiManagementProduct) Attributes() apiManagementProductAttributes {
	return apiManagementProductAttributes{ref: terra.ReferenceResource(amp)}
}

func (amp *ApiManagementProduct) ImportState(av io.Reader) error {
	amp.state = &apiManagementProductState{}
	if err := json.NewDecoder(av).Decode(amp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amp.Type(), amp.LocalName(), err)
	}
	return nil
}

func (amp *ApiManagementProduct) State() (*apiManagementProductState, bool) {
	return amp.state, amp.state != nil
}

func (amp *ApiManagementProduct) StateMust() *apiManagementProductState {
	if amp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amp.Type(), amp.LocalName()))
	}
	return amp.state
}

func (amp *ApiManagementProduct) DependOn() terra.Reference {
	return terra.ReferenceResource(amp)
}

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
	// DependsOn contains resources that ApiManagementProduct depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apiManagementProductAttributes struct {
	ref terra.Reference
}

func (amp apiManagementProductAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceString(amp.ref.Append("api_management_name"))
}

func (amp apiManagementProductAttributes) ApprovalRequired() terra.BoolValue {
	return terra.ReferenceBool(amp.ref.Append("approval_required"))
}

func (amp apiManagementProductAttributes) Description() terra.StringValue {
	return terra.ReferenceString(amp.ref.Append("description"))
}

func (amp apiManagementProductAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(amp.ref.Append("display_name"))
}

func (amp apiManagementProductAttributes) Id() terra.StringValue {
	return terra.ReferenceString(amp.ref.Append("id"))
}

func (amp apiManagementProductAttributes) ProductId() terra.StringValue {
	return terra.ReferenceString(amp.ref.Append("product_id"))
}

func (amp apiManagementProductAttributes) Published() terra.BoolValue {
	return terra.ReferenceBool(amp.ref.Append("published"))
}

func (amp apiManagementProductAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(amp.ref.Append("resource_group_name"))
}

func (amp apiManagementProductAttributes) SubscriptionRequired() terra.BoolValue {
	return terra.ReferenceBool(amp.ref.Append("subscription_required"))
}

func (amp apiManagementProductAttributes) SubscriptionsLimit() terra.NumberValue {
	return terra.ReferenceNumber(amp.ref.Append("subscriptions_limit"))
}

func (amp apiManagementProductAttributes) Terms() terra.StringValue {
	return terra.ReferenceString(amp.ref.Append("terms"))
}

func (amp apiManagementProductAttributes) Timeouts() apimanagementproduct.TimeoutsAttributes {
	return terra.ReferenceSingle[apimanagementproduct.TimeoutsAttributes](amp.ref.Append("timeouts"))
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
