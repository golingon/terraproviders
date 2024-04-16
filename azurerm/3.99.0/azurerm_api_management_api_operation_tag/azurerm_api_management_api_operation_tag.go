// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_api_management_api_operation_tag

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

// Resource represents the Terraform resource azurerm_api_management_api_operation_tag.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermApiManagementApiOperationTagState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aamaot *Resource) Type() string {
	return "azurerm_api_management_api_operation_tag"
}

// LocalName returns the local name for [Resource].
func (aamaot *Resource) LocalName() string {
	return aamaot.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aamaot *Resource) Configuration() interface{} {
	return aamaot.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aamaot *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aamaot)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aamaot *Resource) Dependencies() terra.Dependencies {
	return aamaot.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aamaot *Resource) LifecycleManagement() *terra.Lifecycle {
	return aamaot.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aamaot *Resource) Attributes() azurermApiManagementApiOperationTagAttributes {
	return azurermApiManagementApiOperationTagAttributes{ref: terra.ReferenceResource(aamaot)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aamaot *Resource) ImportState(state io.Reader) error {
	aamaot.state = &azurermApiManagementApiOperationTagState{}
	if err := json.NewDecoder(state).Decode(aamaot.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aamaot.Type(), aamaot.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aamaot *Resource) State() (*azurermApiManagementApiOperationTagState, bool) {
	return aamaot.state, aamaot.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aamaot *Resource) StateMust() *azurermApiManagementApiOperationTagState {
	if aamaot.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aamaot.Type(), aamaot.LocalName()))
	}
	return aamaot.state
}

// Args contains the configurations for azurerm_api_management_api_operation_tag.
type Args struct {
	// ApiOperationId: string, required
	ApiOperationId terra.StringValue `hcl:"api_operation_id,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermApiManagementApiOperationTagAttributes struct {
	ref terra.Reference
}

// ApiOperationId returns a reference to field api_operation_id of azurerm_api_management_api_operation_tag.
func (aamaot azurermApiManagementApiOperationTagAttributes) ApiOperationId() terra.StringValue {
	return terra.ReferenceAsString(aamaot.ref.Append("api_operation_id"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_api_operation_tag.
func (aamaot azurermApiManagementApiOperationTagAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aamaot.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_api_operation_tag.
func (aamaot azurermApiManagementApiOperationTagAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aamaot.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_api_operation_tag.
func (aamaot azurermApiManagementApiOperationTagAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aamaot.ref.Append("name"))
}

func (aamaot azurermApiManagementApiOperationTagAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aamaot.ref.Append("timeouts"))
}

type azurermApiManagementApiOperationTagState struct {
	ApiOperationId string         `json:"api_operation_id"`
	DisplayName    string         `json:"display_name"`
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Timeouts       *TimeoutsState `json:"timeouts"`
}
