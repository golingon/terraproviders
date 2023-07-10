// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementproductapi "github.com/golingon/terraproviders/azurerm/3.64.0/apimanagementproductapi"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementProductApi creates a new instance of [ApiManagementProductApi].
func NewApiManagementProductApi(name string, args ApiManagementProductApiArgs) *ApiManagementProductApi {
	return &ApiManagementProductApi{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementProductApi)(nil)

// ApiManagementProductApi represents the Terraform resource azurerm_api_management_product_api.
type ApiManagementProductApi struct {
	Name      string
	Args      ApiManagementProductApiArgs
	state     *apiManagementProductApiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementProductApi].
func (ampa *ApiManagementProductApi) Type() string {
	return "azurerm_api_management_product_api"
}

// LocalName returns the local name for [ApiManagementProductApi].
func (ampa *ApiManagementProductApi) LocalName() string {
	return ampa.Name
}

// Configuration returns the configuration (args) for [ApiManagementProductApi].
func (ampa *ApiManagementProductApi) Configuration() interface{} {
	return ampa.Args
}

// DependOn is used for other resources to depend on [ApiManagementProductApi].
func (ampa *ApiManagementProductApi) DependOn() terra.Reference {
	return terra.ReferenceResource(ampa)
}

// Dependencies returns the list of resources [ApiManagementProductApi] depends_on.
func (ampa *ApiManagementProductApi) Dependencies() terra.Dependencies {
	return ampa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementProductApi].
func (ampa *ApiManagementProductApi) LifecycleManagement() *terra.Lifecycle {
	return ampa.Lifecycle
}

// Attributes returns the attributes for [ApiManagementProductApi].
func (ampa *ApiManagementProductApi) Attributes() apiManagementProductApiAttributes {
	return apiManagementProductApiAttributes{ref: terra.ReferenceResource(ampa)}
}

// ImportState imports the given attribute values into [ApiManagementProductApi]'s state.
func (ampa *ApiManagementProductApi) ImportState(av io.Reader) error {
	ampa.state = &apiManagementProductApiState{}
	if err := json.NewDecoder(av).Decode(ampa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ampa.Type(), ampa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementProductApi] has state.
func (ampa *ApiManagementProductApi) State() (*apiManagementProductApiState, bool) {
	return ampa.state, ampa.state != nil
}

// StateMust returns the state for [ApiManagementProductApi]. Panics if the state is nil.
func (ampa *ApiManagementProductApi) StateMust() *apiManagementProductApiState {
	if ampa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ampa.Type(), ampa.LocalName()))
	}
	return ampa.state
}

// ApiManagementProductApiArgs contains the configurations for azurerm_api_management_product_api.
type ApiManagementProductApiArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// ApiName: string, required
	ApiName terra.StringValue `hcl:"api_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProductId: string, required
	ProductId terra.StringValue `hcl:"product_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementproductapi.Timeouts `hcl:"timeouts,block"`
}
type apiManagementProductApiAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_product_api.
func (ampa apiManagementProductApiAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(ampa.ref.Append("api_management_name"))
}

// ApiName returns a reference to field api_name of azurerm_api_management_product_api.
func (ampa apiManagementProductApiAttributes) ApiName() terra.StringValue {
	return terra.ReferenceAsString(ampa.ref.Append("api_name"))
}

// Id returns a reference to field id of azurerm_api_management_product_api.
func (ampa apiManagementProductApiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ampa.ref.Append("id"))
}

// ProductId returns a reference to field product_id of azurerm_api_management_product_api.
func (ampa apiManagementProductApiAttributes) ProductId() terra.StringValue {
	return terra.ReferenceAsString(ampa.ref.Append("product_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_product_api.
func (ampa apiManagementProductApiAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ampa.ref.Append("resource_group_name"))
}

func (ampa apiManagementProductApiAttributes) Timeouts() apimanagementproductapi.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementproductapi.TimeoutsAttributes](ampa.ref.Append("timeouts"))
}

type apiManagementProductApiState struct {
	ApiManagementName string                                 `json:"api_management_name"`
	ApiName           string                                 `json:"api_name"`
	Id                string                                 `json:"id"`
	ProductId         string                                 `json:"product_id"`
	ResourceGroupName string                                 `json:"resource_group_name"`
	Timeouts          *apimanagementproductapi.TimeoutsState `json:"timeouts"`
}
