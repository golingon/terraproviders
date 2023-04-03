// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementnamedvalue "github.com/golingon/terraproviders/azurerm/3.49.0/apimanagementnamedvalue"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementNamedValue creates a new instance of [ApiManagementNamedValue].
func NewApiManagementNamedValue(name string, args ApiManagementNamedValueArgs) *ApiManagementNamedValue {
	return &ApiManagementNamedValue{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementNamedValue)(nil)

// ApiManagementNamedValue represents the Terraform resource azurerm_api_management_named_value.
type ApiManagementNamedValue struct {
	Name      string
	Args      ApiManagementNamedValueArgs
	state     *apiManagementNamedValueState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementNamedValue].
func (amnv *ApiManagementNamedValue) Type() string {
	return "azurerm_api_management_named_value"
}

// LocalName returns the local name for [ApiManagementNamedValue].
func (amnv *ApiManagementNamedValue) LocalName() string {
	return amnv.Name
}

// Configuration returns the configuration (args) for [ApiManagementNamedValue].
func (amnv *ApiManagementNamedValue) Configuration() interface{} {
	return amnv.Args
}

// DependOn is used for other resources to depend on [ApiManagementNamedValue].
func (amnv *ApiManagementNamedValue) DependOn() terra.Reference {
	return terra.ReferenceResource(amnv)
}

// Dependencies returns the list of resources [ApiManagementNamedValue] depends_on.
func (amnv *ApiManagementNamedValue) Dependencies() terra.Dependencies {
	return amnv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementNamedValue].
func (amnv *ApiManagementNamedValue) LifecycleManagement() *terra.Lifecycle {
	return amnv.Lifecycle
}

// Attributes returns the attributes for [ApiManagementNamedValue].
func (amnv *ApiManagementNamedValue) Attributes() apiManagementNamedValueAttributes {
	return apiManagementNamedValueAttributes{ref: terra.ReferenceResource(amnv)}
}

// ImportState imports the given attribute values into [ApiManagementNamedValue]'s state.
func (amnv *ApiManagementNamedValue) ImportState(av io.Reader) error {
	amnv.state = &apiManagementNamedValueState{}
	if err := json.NewDecoder(av).Decode(amnv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amnv.Type(), amnv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementNamedValue] has state.
func (amnv *ApiManagementNamedValue) State() (*apiManagementNamedValueState, bool) {
	return amnv.state, amnv.state != nil
}

// StateMust returns the state for [ApiManagementNamedValue]. Panics if the state is nil.
func (amnv *ApiManagementNamedValue) StateMust() *apiManagementNamedValueState {
	if amnv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amnv.Type(), amnv.LocalName()))
	}
	return amnv.state
}

// ApiManagementNamedValueArgs contains the configurations for azurerm_api_management_named_value.
type ApiManagementNamedValueArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Secret: bool, optional
	Secret terra.BoolValue `hcl:"secret,attr"`
	// Tags: list of string, optional
	Tags terra.ListValue[terra.StringValue] `hcl:"tags,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
	// Timeouts: optional
	Timeouts *apimanagementnamedvalue.Timeouts `hcl:"timeouts,block"`
	// ValueFromKeyVault: optional
	ValueFromKeyVault *apimanagementnamedvalue.ValueFromKeyVault `hcl:"value_from_key_vault,block"`
}
type apiManagementNamedValueAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_named_value.
func (amnv apiManagementNamedValueAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amnv.ref.Append("api_management_name"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_named_value.
func (amnv apiManagementNamedValueAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(amnv.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_named_value.
func (amnv apiManagementNamedValueAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amnv.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_named_value.
func (amnv apiManagementNamedValueAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amnv.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_named_value.
func (amnv apiManagementNamedValueAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amnv.ref.Append("resource_group_name"))
}

// Secret returns a reference to field secret of azurerm_api_management_named_value.
func (amnv apiManagementNamedValueAttributes) Secret() terra.BoolValue {
	return terra.ReferenceAsBool(amnv.ref.Append("secret"))
}

// Tags returns a reference to field tags of azurerm_api_management_named_value.
func (amnv apiManagementNamedValueAttributes) Tags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](amnv.ref.Append("tags"))
}

// Value returns a reference to field value of azurerm_api_management_named_value.
func (amnv apiManagementNamedValueAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(amnv.ref.Append("value"))
}

func (amnv apiManagementNamedValueAttributes) Timeouts() apimanagementnamedvalue.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementnamedvalue.TimeoutsAttributes](amnv.ref.Append("timeouts"))
}

func (amnv apiManagementNamedValueAttributes) ValueFromKeyVault() terra.ListValue[apimanagementnamedvalue.ValueFromKeyVaultAttributes] {
	return terra.ReferenceAsList[apimanagementnamedvalue.ValueFromKeyVaultAttributes](amnv.ref.Append("value_from_key_vault"))
}

type apiManagementNamedValueState struct {
	ApiManagementName string                                           `json:"api_management_name"`
	DisplayName       string                                           `json:"display_name"`
	Id                string                                           `json:"id"`
	Name              string                                           `json:"name"`
	ResourceGroupName string                                           `json:"resource_group_name"`
	Secret            bool                                             `json:"secret"`
	Tags              []string                                         `json:"tags"`
	Value             string                                           `json:"value"`
	Timeouts          *apimanagementnamedvalue.TimeoutsState           `json:"timeouts"`
	ValueFromKeyVault []apimanagementnamedvalue.ValueFromKeyVaultState `json:"value_from_key_vault"`
}
