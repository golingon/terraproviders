// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_api_management_api_operation_policy

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

// Resource represents the Terraform resource azurerm_api_management_api_operation_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermApiManagementApiOperationPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aamaop *Resource) Type() string {
	return "azurerm_api_management_api_operation_policy"
}

// LocalName returns the local name for [Resource].
func (aamaop *Resource) LocalName() string {
	return aamaop.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aamaop *Resource) Configuration() interface{} {
	return aamaop.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aamaop *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aamaop)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aamaop *Resource) Dependencies() terra.Dependencies {
	return aamaop.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aamaop *Resource) LifecycleManagement() *terra.Lifecycle {
	return aamaop.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aamaop *Resource) Attributes() azurermApiManagementApiOperationPolicyAttributes {
	return azurermApiManagementApiOperationPolicyAttributes{ref: terra.ReferenceResource(aamaop)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aamaop *Resource) ImportState(state io.Reader) error {
	aamaop.state = &azurermApiManagementApiOperationPolicyState{}
	if err := json.NewDecoder(state).Decode(aamaop.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aamaop.Type(), aamaop.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aamaop *Resource) State() (*azurermApiManagementApiOperationPolicyState, bool) {
	return aamaop.state, aamaop.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aamaop *Resource) StateMust() *azurermApiManagementApiOperationPolicyState {
	if aamaop.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aamaop.Type(), aamaop.LocalName()))
	}
	return aamaop.state
}

// Args contains the configurations for azurerm_api_management_api_operation_policy.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermApiManagementApiOperationPolicyAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_api_operation_policy.
func (aamaop azurermApiManagementApiOperationPolicyAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(aamaop.ref.Append("api_management_name"))
}

// ApiName returns a reference to field api_name of azurerm_api_management_api_operation_policy.
func (aamaop azurermApiManagementApiOperationPolicyAttributes) ApiName() terra.StringValue {
	return terra.ReferenceAsString(aamaop.ref.Append("api_name"))
}

// Id returns a reference to field id of azurerm_api_management_api_operation_policy.
func (aamaop azurermApiManagementApiOperationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aamaop.ref.Append("id"))
}

// OperationId returns a reference to field operation_id of azurerm_api_management_api_operation_policy.
func (aamaop azurermApiManagementApiOperationPolicyAttributes) OperationId() terra.StringValue {
	return terra.ReferenceAsString(aamaop.ref.Append("operation_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_api_operation_policy.
func (aamaop azurermApiManagementApiOperationPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aamaop.ref.Append("resource_group_name"))
}

// XmlContent returns a reference to field xml_content of azurerm_api_management_api_operation_policy.
func (aamaop azurermApiManagementApiOperationPolicyAttributes) XmlContent() terra.StringValue {
	return terra.ReferenceAsString(aamaop.ref.Append("xml_content"))
}

// XmlLink returns a reference to field xml_link of azurerm_api_management_api_operation_policy.
func (aamaop azurermApiManagementApiOperationPolicyAttributes) XmlLink() terra.StringValue {
	return terra.ReferenceAsString(aamaop.ref.Append("xml_link"))
}

func (aamaop azurermApiManagementApiOperationPolicyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aamaop.ref.Append("timeouts"))
}

type azurermApiManagementApiOperationPolicyState struct {
	ApiManagementName string         `json:"api_management_name"`
	ApiName           string         `json:"api_name"`
	Id                string         `json:"id"`
	OperationId       string         `json:"operation_id"`
	ResourceGroupName string         `json:"resource_group_name"`
	XmlContent        string         `json:"xml_content"`
	XmlLink           string         `json:"xml_link"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
