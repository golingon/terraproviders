// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementapioperationpolicy "github.com/golingon/terraproviders/azurerm/3.49.0/apimanagementapioperationpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementApiOperationPolicy creates a new instance of [ApiManagementApiOperationPolicy].
func NewApiManagementApiOperationPolicy(name string, args ApiManagementApiOperationPolicyArgs) *ApiManagementApiOperationPolicy {
	return &ApiManagementApiOperationPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementApiOperationPolicy)(nil)

// ApiManagementApiOperationPolicy represents the Terraform resource azurerm_api_management_api_operation_policy.
type ApiManagementApiOperationPolicy struct {
	Name      string
	Args      ApiManagementApiOperationPolicyArgs
	state     *apiManagementApiOperationPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementApiOperationPolicy].
func (amaop *ApiManagementApiOperationPolicy) Type() string {
	return "azurerm_api_management_api_operation_policy"
}

// LocalName returns the local name for [ApiManagementApiOperationPolicy].
func (amaop *ApiManagementApiOperationPolicy) LocalName() string {
	return amaop.Name
}

// Configuration returns the configuration (args) for [ApiManagementApiOperationPolicy].
func (amaop *ApiManagementApiOperationPolicy) Configuration() interface{} {
	return amaop.Args
}

// DependOn is used for other resources to depend on [ApiManagementApiOperationPolicy].
func (amaop *ApiManagementApiOperationPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(amaop)
}

// Dependencies returns the list of resources [ApiManagementApiOperationPolicy] depends_on.
func (amaop *ApiManagementApiOperationPolicy) Dependencies() terra.Dependencies {
	return amaop.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementApiOperationPolicy].
func (amaop *ApiManagementApiOperationPolicy) LifecycleManagement() *terra.Lifecycle {
	return amaop.Lifecycle
}

// Attributes returns the attributes for [ApiManagementApiOperationPolicy].
func (amaop *ApiManagementApiOperationPolicy) Attributes() apiManagementApiOperationPolicyAttributes {
	return apiManagementApiOperationPolicyAttributes{ref: terra.ReferenceResource(amaop)}
}

// ImportState imports the given attribute values into [ApiManagementApiOperationPolicy]'s state.
func (amaop *ApiManagementApiOperationPolicy) ImportState(av io.Reader) error {
	amaop.state = &apiManagementApiOperationPolicyState{}
	if err := json.NewDecoder(av).Decode(amaop.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amaop.Type(), amaop.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementApiOperationPolicy] has state.
func (amaop *ApiManagementApiOperationPolicy) State() (*apiManagementApiOperationPolicyState, bool) {
	return amaop.state, amaop.state != nil
}

// StateMust returns the state for [ApiManagementApiOperationPolicy]. Panics if the state is nil.
func (amaop *ApiManagementApiOperationPolicy) StateMust() *apiManagementApiOperationPolicyState {
	if amaop.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amaop.Type(), amaop.LocalName()))
	}
	return amaop.state
}

// ApiManagementApiOperationPolicyArgs contains the configurations for azurerm_api_management_api_operation_policy.
type ApiManagementApiOperationPolicyArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// ApiName: string, required
	ApiName terra.StringValue `hcl:"api_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OperationId: string, required
	OperationId terra.StringValue `hcl:"operation_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// XmlContent: string, optional
	XmlContent terra.StringValue `hcl:"xml_content,attr"`
	// XmlLink: string, optional
	XmlLink terra.StringValue `hcl:"xml_link,attr"`
	// Timeouts: optional
	Timeouts *apimanagementapioperationpolicy.Timeouts `hcl:"timeouts,block"`
}
type apiManagementApiOperationPolicyAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_api_operation_policy.
func (amaop apiManagementApiOperationPolicyAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amaop.ref.Append("api_management_name"))
}

// ApiName returns a reference to field api_name of azurerm_api_management_api_operation_policy.
func (amaop apiManagementApiOperationPolicyAttributes) ApiName() terra.StringValue {
	return terra.ReferenceAsString(amaop.ref.Append("api_name"))
}

// Id returns a reference to field id of azurerm_api_management_api_operation_policy.
func (amaop apiManagementApiOperationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amaop.ref.Append("id"))
}

// OperationId returns a reference to field operation_id of azurerm_api_management_api_operation_policy.
func (amaop apiManagementApiOperationPolicyAttributes) OperationId() terra.StringValue {
	return terra.ReferenceAsString(amaop.ref.Append("operation_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_api_operation_policy.
func (amaop apiManagementApiOperationPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amaop.ref.Append("resource_group_name"))
}

// XmlContent returns a reference to field xml_content of azurerm_api_management_api_operation_policy.
func (amaop apiManagementApiOperationPolicyAttributes) XmlContent() terra.StringValue {
	return terra.ReferenceAsString(amaop.ref.Append("xml_content"))
}

// XmlLink returns a reference to field xml_link of azurerm_api_management_api_operation_policy.
func (amaop apiManagementApiOperationPolicyAttributes) XmlLink() terra.StringValue {
	return terra.ReferenceAsString(amaop.ref.Append("xml_link"))
}

func (amaop apiManagementApiOperationPolicyAttributes) Timeouts() apimanagementapioperationpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementapioperationpolicy.TimeoutsAttributes](amaop.ref.Append("timeouts"))
}

type apiManagementApiOperationPolicyState struct {
	ApiManagementName string                                         `json:"api_management_name"`
	ApiName           string                                         `json:"api_name"`
	Id                string                                         `json:"id"`
	OperationId       string                                         `json:"operation_id"`
	ResourceGroupName string                                         `json:"resource_group_name"`
	XmlContent        string                                         `json:"xml_content"`
	XmlLink           string                                         `json:"xml_link"`
	Timeouts          *apimanagementapioperationpolicy.TimeoutsState `json:"timeouts"`
}
