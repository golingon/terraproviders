// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementproducttag "github.com/golingon/terraproviders/azurerm/3.65.0/apimanagementproducttag"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementProductTag creates a new instance of [ApiManagementProductTag].
func NewApiManagementProductTag(name string, args ApiManagementProductTagArgs) *ApiManagementProductTag {
	return &ApiManagementProductTag{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementProductTag)(nil)

// ApiManagementProductTag represents the Terraform resource azurerm_api_management_product_tag.
type ApiManagementProductTag struct {
	Name      string
	Args      ApiManagementProductTagArgs
	state     *apiManagementProductTagState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementProductTag].
func (ampt *ApiManagementProductTag) Type() string {
	return "azurerm_api_management_product_tag"
}

// LocalName returns the local name for [ApiManagementProductTag].
func (ampt *ApiManagementProductTag) LocalName() string {
	return ampt.Name
}

// Configuration returns the configuration (args) for [ApiManagementProductTag].
func (ampt *ApiManagementProductTag) Configuration() interface{} {
	return ampt.Args
}

// DependOn is used for other resources to depend on [ApiManagementProductTag].
func (ampt *ApiManagementProductTag) DependOn() terra.Reference {
	return terra.ReferenceResource(ampt)
}

// Dependencies returns the list of resources [ApiManagementProductTag] depends_on.
func (ampt *ApiManagementProductTag) Dependencies() terra.Dependencies {
	return ampt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementProductTag].
func (ampt *ApiManagementProductTag) LifecycleManagement() *terra.Lifecycle {
	return ampt.Lifecycle
}

// Attributes returns the attributes for [ApiManagementProductTag].
func (ampt *ApiManagementProductTag) Attributes() apiManagementProductTagAttributes {
	return apiManagementProductTagAttributes{ref: terra.ReferenceResource(ampt)}
}

// ImportState imports the given attribute values into [ApiManagementProductTag]'s state.
func (ampt *ApiManagementProductTag) ImportState(av io.Reader) error {
	ampt.state = &apiManagementProductTagState{}
	if err := json.NewDecoder(av).Decode(ampt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ampt.Type(), ampt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementProductTag] has state.
func (ampt *ApiManagementProductTag) State() (*apiManagementProductTagState, bool) {
	return ampt.state, ampt.state != nil
}

// StateMust returns the state for [ApiManagementProductTag]. Panics if the state is nil.
func (ampt *ApiManagementProductTag) StateMust() *apiManagementProductTagState {
	if ampt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ampt.Type(), ampt.LocalName()))
	}
	return ampt.state
}

// ApiManagementProductTagArgs contains the configurations for azurerm_api_management_product_tag.
type ApiManagementProductTagArgs struct {
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
	Timeouts *apimanagementproducttag.Timeouts `hcl:"timeouts,block"`
}
type apiManagementProductTagAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_product_tag.
func (ampt apiManagementProductTagAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(ampt.ref.Append("api_management_name"))
}

// ApiManagementProductId returns a reference to field api_management_product_id of azurerm_api_management_product_tag.
func (ampt apiManagementProductTagAttributes) ApiManagementProductId() terra.StringValue {
	return terra.ReferenceAsString(ampt.ref.Append("api_management_product_id"))
}

// Id returns a reference to field id of azurerm_api_management_product_tag.
func (ampt apiManagementProductTagAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ampt.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_product_tag.
func (ampt apiManagementProductTagAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ampt.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_product_tag.
func (ampt apiManagementProductTagAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ampt.ref.Append("resource_group_name"))
}

func (ampt apiManagementProductTagAttributes) Timeouts() apimanagementproducttag.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementproducttag.TimeoutsAttributes](ampt.ref.Append("timeouts"))
}

type apiManagementProductTagState struct {
	ApiManagementName      string                                 `json:"api_management_name"`
	ApiManagementProductId string                                 `json:"api_management_product_id"`
	Id                     string                                 `json:"id"`
	Name                   string                                 `json:"name"`
	ResourceGroupName      string                                 `json:"resource_group_name"`
	Timeouts               *apimanagementproducttag.TimeoutsState `json:"timeouts"`
}