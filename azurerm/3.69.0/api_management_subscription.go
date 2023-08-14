// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementsubscription "github.com/golingon/terraproviders/azurerm/3.69.0/apimanagementsubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementSubscription creates a new instance of [ApiManagementSubscription].
func NewApiManagementSubscription(name string, args ApiManagementSubscriptionArgs) *ApiManagementSubscription {
	return &ApiManagementSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementSubscription)(nil)

// ApiManagementSubscription represents the Terraform resource azurerm_api_management_subscription.
type ApiManagementSubscription struct {
	Name      string
	Args      ApiManagementSubscriptionArgs
	state     *apiManagementSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementSubscription].
func (ams *ApiManagementSubscription) Type() string {
	return "azurerm_api_management_subscription"
}

// LocalName returns the local name for [ApiManagementSubscription].
func (ams *ApiManagementSubscription) LocalName() string {
	return ams.Name
}

// Configuration returns the configuration (args) for [ApiManagementSubscription].
func (ams *ApiManagementSubscription) Configuration() interface{} {
	return ams.Args
}

// DependOn is used for other resources to depend on [ApiManagementSubscription].
func (ams *ApiManagementSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(ams)
}

// Dependencies returns the list of resources [ApiManagementSubscription] depends_on.
func (ams *ApiManagementSubscription) Dependencies() terra.Dependencies {
	return ams.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementSubscription].
func (ams *ApiManagementSubscription) LifecycleManagement() *terra.Lifecycle {
	return ams.Lifecycle
}

// Attributes returns the attributes for [ApiManagementSubscription].
func (ams *ApiManagementSubscription) Attributes() apiManagementSubscriptionAttributes {
	return apiManagementSubscriptionAttributes{ref: terra.ReferenceResource(ams)}
}

// ImportState imports the given attribute values into [ApiManagementSubscription]'s state.
func (ams *ApiManagementSubscription) ImportState(av io.Reader) error {
	ams.state = &apiManagementSubscriptionState{}
	if err := json.NewDecoder(av).Decode(ams.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ams.Type(), ams.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementSubscription] has state.
func (ams *ApiManagementSubscription) State() (*apiManagementSubscriptionState, bool) {
	return ams.state, ams.state != nil
}

// StateMust returns the state for [ApiManagementSubscription]. Panics if the state is nil.
func (ams *ApiManagementSubscription) StateMust() *apiManagementSubscriptionState {
	if ams.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ams.Type(), ams.LocalName()))
	}
	return ams.state
}

// ApiManagementSubscriptionArgs contains the configurations for azurerm_api_management_subscription.
type ApiManagementSubscriptionArgs struct {
	// AllowTracing: bool, optional
	AllowTracing terra.BoolValue `hcl:"allow_tracing,attr"`
	// ApiId: string, optional
	ApiId terra.StringValue `hcl:"api_id,attr"`
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrimaryKey: string, optional
	PrimaryKey terra.StringValue `hcl:"primary_key,attr"`
	// ProductId: string, optional
	ProductId terra.StringValue `hcl:"product_id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SecondaryKey: string, optional
	SecondaryKey terra.StringValue `hcl:"secondary_key,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// SubscriptionId: string, optional
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr"`
	// UserId: string, optional
	UserId terra.StringValue `hcl:"user_id,attr"`
	// Timeouts: optional
	Timeouts *apimanagementsubscription.Timeouts `hcl:"timeouts,block"`
}
type apiManagementSubscriptionAttributes struct {
	ref terra.Reference
}

// AllowTracing returns a reference to field allow_tracing of azurerm_api_management_subscription.
func (ams apiManagementSubscriptionAttributes) AllowTracing() terra.BoolValue {
	return terra.ReferenceAsBool(ams.ref.Append("allow_tracing"))
}

// ApiId returns a reference to field api_id of azurerm_api_management_subscription.
func (ams apiManagementSubscriptionAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("api_id"))
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_subscription.
func (ams apiManagementSubscriptionAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("api_management_name"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_subscription.
func (ams apiManagementSubscriptionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_subscription.
func (ams apiManagementSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("id"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_api_management_subscription.
func (ams apiManagementSubscriptionAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("primary_key"))
}

// ProductId returns a reference to field product_id of azurerm_api_management_subscription.
func (ams apiManagementSubscriptionAttributes) ProductId() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("product_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_subscription.
func (ams apiManagementSubscriptionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("resource_group_name"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_api_management_subscription.
func (ams apiManagementSubscriptionAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("secondary_key"))
}

// State returns a reference to field state of azurerm_api_management_subscription.
func (ams apiManagementSubscriptionAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("state"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_api_management_subscription.
func (ams apiManagementSubscriptionAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("subscription_id"))
}

// UserId returns a reference to field user_id of azurerm_api_management_subscription.
func (ams apiManagementSubscriptionAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("user_id"))
}

func (ams apiManagementSubscriptionAttributes) Timeouts() apimanagementsubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementsubscription.TimeoutsAttributes](ams.ref.Append("timeouts"))
}

type apiManagementSubscriptionState struct {
	AllowTracing      bool                                     `json:"allow_tracing"`
	ApiId             string                                   `json:"api_id"`
	ApiManagementName string                                   `json:"api_management_name"`
	DisplayName       string                                   `json:"display_name"`
	Id                string                                   `json:"id"`
	PrimaryKey        string                                   `json:"primary_key"`
	ProductId         string                                   `json:"product_id"`
	ResourceGroupName string                                   `json:"resource_group_name"`
	SecondaryKey      string                                   `json:"secondary_key"`
	State             string                                   `json:"state"`
	SubscriptionId    string                                   `json:"subscription_id"`
	UserId            string                                   `json:"user_id"`
	Timeouts          *apimanagementsubscription.TimeoutsState `json:"timeouts"`
}
