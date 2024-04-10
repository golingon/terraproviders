// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	apimanagementapipolicy "github.com/golingon/terraproviders/azurerm/3.98.0/apimanagementapipolicy"
	"io"
)

// NewApiManagementApiPolicy creates a new instance of [ApiManagementApiPolicy].
func NewApiManagementApiPolicy(name string, args ApiManagementApiPolicyArgs) *ApiManagementApiPolicy {
	return &ApiManagementApiPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementApiPolicy)(nil)

// ApiManagementApiPolicy represents the Terraform resource azurerm_api_management_api_policy.
type ApiManagementApiPolicy struct {
	Name      string
	Args      ApiManagementApiPolicyArgs
	state     *apiManagementApiPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementApiPolicy].
func (amap *ApiManagementApiPolicy) Type() string {
	return "azurerm_api_management_api_policy"
}

// LocalName returns the local name for [ApiManagementApiPolicy].
func (amap *ApiManagementApiPolicy) LocalName() string {
	return amap.Name
}

// Configuration returns the configuration (args) for [ApiManagementApiPolicy].
func (amap *ApiManagementApiPolicy) Configuration() interface{} {
	return amap.Args
}

// DependOn is used for other resources to depend on [ApiManagementApiPolicy].
func (amap *ApiManagementApiPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(amap)
}

// Dependencies returns the list of resources [ApiManagementApiPolicy] depends_on.
func (amap *ApiManagementApiPolicy) Dependencies() terra.Dependencies {
	return amap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementApiPolicy].
func (amap *ApiManagementApiPolicy) LifecycleManagement() *terra.Lifecycle {
	return amap.Lifecycle
}

// Attributes returns the attributes for [ApiManagementApiPolicy].
func (amap *ApiManagementApiPolicy) Attributes() apiManagementApiPolicyAttributes {
	return apiManagementApiPolicyAttributes{ref: terra.ReferenceResource(amap)}
}

// ImportState imports the given attribute values into [ApiManagementApiPolicy]'s state.
func (amap *ApiManagementApiPolicy) ImportState(av io.Reader) error {
	amap.state = &apiManagementApiPolicyState{}
	if err := json.NewDecoder(av).Decode(amap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amap.Type(), amap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementApiPolicy] has state.
func (amap *ApiManagementApiPolicy) State() (*apiManagementApiPolicyState, bool) {
	return amap.state, amap.state != nil
}

// StateMust returns the state for [ApiManagementApiPolicy]. Panics if the state is nil.
func (amap *ApiManagementApiPolicy) StateMust() *apiManagementApiPolicyState {
	if amap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amap.Type(), amap.LocalName()))
	}
	return amap.state
}

// ApiManagementApiPolicyArgs contains the configurations for azurerm_api_management_api_policy.
type ApiManagementApiPolicyArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// ApiName: string, required
	ApiName terra.StringValue `hcl:"api_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// XmlContent: string, optional
	XmlContent terra.StringValue `hcl:"xml_content,attr"`
	// XmlLink: string, optional
	XmlLink terra.StringValue `hcl:"xml_link,attr"`
	// Timeouts: optional
	Timeouts *apimanagementapipolicy.Timeouts `hcl:"timeouts,block"`
}
type apiManagementApiPolicyAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_api_policy.
func (amap apiManagementApiPolicyAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amap.ref.Append("api_management_name"))
}

// ApiName returns a reference to field api_name of azurerm_api_management_api_policy.
func (amap apiManagementApiPolicyAttributes) ApiName() terra.StringValue {
	return terra.ReferenceAsString(amap.ref.Append("api_name"))
}

// Id returns a reference to field id of azurerm_api_management_api_policy.
func (amap apiManagementApiPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amap.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_api_policy.
func (amap apiManagementApiPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amap.ref.Append("resource_group_name"))
}

// XmlContent returns a reference to field xml_content of azurerm_api_management_api_policy.
func (amap apiManagementApiPolicyAttributes) XmlContent() terra.StringValue {
	return terra.ReferenceAsString(amap.ref.Append("xml_content"))
}

// XmlLink returns a reference to field xml_link of azurerm_api_management_api_policy.
func (amap apiManagementApiPolicyAttributes) XmlLink() terra.StringValue {
	return terra.ReferenceAsString(amap.ref.Append("xml_link"))
}

func (amap apiManagementApiPolicyAttributes) Timeouts() apimanagementapipolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementapipolicy.TimeoutsAttributes](amap.ref.Append("timeouts"))
}

type apiManagementApiPolicyState struct {
	ApiManagementName string                                `json:"api_management_name"`
	ApiName           string                                `json:"api_name"`
	Id                string                                `json:"id"`
	ResourceGroupName string                                `json:"resource_group_name"`
	XmlContent        string                                `json:"xml_content"`
	XmlLink           string                                `json:"xml_link"`
	Timeouts          *apimanagementapipolicy.TimeoutsState `json:"timeouts"`
}
