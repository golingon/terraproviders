// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_api_management_product_policy

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

// Resource represents the Terraform resource azurerm_api_management_product_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermApiManagementProductPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aampp *Resource) Type() string {
	return "azurerm_api_management_product_policy"
}

// LocalName returns the local name for [Resource].
func (aampp *Resource) LocalName() string {
	return aampp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aampp *Resource) Configuration() interface{} {
	return aampp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aampp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aampp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aampp *Resource) Dependencies() terra.Dependencies {
	return aampp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aampp *Resource) LifecycleManagement() *terra.Lifecycle {
	return aampp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aampp *Resource) Attributes() azurermApiManagementProductPolicyAttributes {
	return azurermApiManagementProductPolicyAttributes{ref: terra.ReferenceResource(aampp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aampp *Resource) ImportState(state io.Reader) error {
	aampp.state = &azurermApiManagementProductPolicyState{}
	if err := json.NewDecoder(state).Decode(aampp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aampp.Type(), aampp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aampp *Resource) State() (*azurermApiManagementProductPolicyState, bool) {
	return aampp.state, aampp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aampp *Resource) StateMust() *azurermApiManagementProductPolicyState {
	if aampp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aampp.Type(), aampp.LocalName()))
	}
	return aampp.state
}

// Args contains the configurations for azurerm_api_management_product_policy.
type Args struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProductId: string, required
	ProductId terra.StringValue `hcl:"product_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// XmlContent: string, optional
	XmlContent terra.StringValue `hcl:"xml_content,attr"`
	// XmlLink: string, optional
	XmlLink terra.StringValue `hcl:"xml_link,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermApiManagementProductPolicyAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_product_policy.
func (aampp azurermApiManagementProductPolicyAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(aampp.ref.Append("api_management_name"))
}

// Id returns a reference to field id of azurerm_api_management_product_policy.
func (aampp azurermApiManagementProductPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aampp.ref.Append("id"))
}

// ProductId returns a reference to field product_id of azurerm_api_management_product_policy.
func (aampp azurermApiManagementProductPolicyAttributes) ProductId() terra.StringValue {
	return terra.ReferenceAsString(aampp.ref.Append("product_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_product_policy.
func (aampp azurermApiManagementProductPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aampp.ref.Append("resource_group_name"))
}

// XmlContent returns a reference to field xml_content of azurerm_api_management_product_policy.
func (aampp azurermApiManagementProductPolicyAttributes) XmlContent() terra.StringValue {
	return terra.ReferenceAsString(aampp.ref.Append("xml_content"))
}

// XmlLink returns a reference to field xml_link of azurerm_api_management_product_policy.
func (aampp azurermApiManagementProductPolicyAttributes) XmlLink() terra.StringValue {
	return terra.ReferenceAsString(aampp.ref.Append("xml_link"))
}

func (aampp azurermApiManagementProductPolicyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aampp.ref.Append("timeouts"))
}

type azurermApiManagementProductPolicyState struct {
	ApiManagementName string         `json:"api_management_name"`
	Id                string         `json:"id"`
	ProductId         string         `json:"product_id"`
	ResourceGroupName string         `json:"resource_group_name"`
	XmlContent        string         `json:"xml_content"`
	XmlLink           string         `json:"xml_link"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
