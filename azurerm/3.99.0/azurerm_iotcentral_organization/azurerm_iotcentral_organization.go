// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_iotcentral_organization

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

// Resource represents the Terraform resource azurerm_iotcentral_organization.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermIotcentralOrganizationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aio *Resource) Type() string {
	return "azurerm_iotcentral_organization"
}

// LocalName returns the local name for [Resource].
func (aio *Resource) LocalName() string {
	return aio.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aio *Resource) Configuration() interface{} {
	return aio.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aio *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aio)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aio *Resource) Dependencies() terra.Dependencies {
	return aio.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aio *Resource) LifecycleManagement() *terra.Lifecycle {
	return aio.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aio *Resource) Attributes() azurermIotcentralOrganizationAttributes {
	return azurermIotcentralOrganizationAttributes{ref: terra.ReferenceResource(aio)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aio *Resource) ImportState(state io.Reader) error {
	aio.state = &azurermIotcentralOrganizationState{}
	if err := json.NewDecoder(state).Decode(aio.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aio.Type(), aio.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aio *Resource) State() (*azurermIotcentralOrganizationState, bool) {
	return aio.state, aio.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aio *Resource) StateMust() *azurermIotcentralOrganizationState {
	if aio.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aio.Type(), aio.LocalName()))
	}
	return aio.state
}

// Args contains the configurations for azurerm_iotcentral_organization.
type Args struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IotcentralApplicationId: string, required
	IotcentralApplicationId terra.StringValue `hcl:"iotcentral_application_id,attr" validate:"required"`
	// OrganizationId: string, required
	OrganizationId terra.StringValue `hcl:"organization_id,attr" validate:"required"`
	// ParentOrganizationId: string, optional
	ParentOrganizationId terra.StringValue `hcl:"parent_organization_id,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermIotcentralOrganizationAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azurerm_iotcentral_organization.
func (aio azurermIotcentralOrganizationAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aio.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_iotcentral_organization.
func (aio azurermIotcentralOrganizationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aio.ref.Append("id"))
}

// IotcentralApplicationId returns a reference to field iotcentral_application_id of azurerm_iotcentral_organization.
func (aio azurermIotcentralOrganizationAttributes) IotcentralApplicationId() terra.StringValue {
	return terra.ReferenceAsString(aio.ref.Append("iotcentral_application_id"))
}

// OrganizationId returns a reference to field organization_id of azurerm_iotcentral_organization.
func (aio azurermIotcentralOrganizationAttributes) OrganizationId() terra.StringValue {
	return terra.ReferenceAsString(aio.ref.Append("organization_id"))
}

// ParentOrganizationId returns a reference to field parent_organization_id of azurerm_iotcentral_organization.
func (aio azurermIotcentralOrganizationAttributes) ParentOrganizationId() terra.StringValue {
	return terra.ReferenceAsString(aio.ref.Append("parent_organization_id"))
}

func (aio azurermIotcentralOrganizationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aio.ref.Append("timeouts"))
}

type azurermIotcentralOrganizationState struct {
	DisplayName             string         `json:"display_name"`
	Id                      string         `json:"id"`
	IotcentralApplicationId string         `json:"iotcentral_application_id"`
	OrganizationId          string         `json:"organization_id"`
	ParentOrganizationId    string         `json:"parent_organization_id"`
	Timeouts                *TimeoutsState `json:"timeouts"`
}
