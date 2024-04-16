// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_data_factory_linked_service_web

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

// Resource represents the Terraform resource azurerm_data_factory_linked_service_web.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermDataFactoryLinkedServiceWebState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adflsw *Resource) Type() string {
	return "azurerm_data_factory_linked_service_web"
}

// LocalName returns the local name for [Resource].
func (adflsw *Resource) LocalName() string {
	return adflsw.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adflsw *Resource) Configuration() interface{} {
	return adflsw.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adflsw *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adflsw)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adflsw *Resource) Dependencies() terra.Dependencies {
	return adflsw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adflsw *Resource) LifecycleManagement() *terra.Lifecycle {
	return adflsw.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adflsw *Resource) Attributes() azurermDataFactoryLinkedServiceWebAttributes {
	return azurermDataFactoryLinkedServiceWebAttributes{ref: terra.ReferenceResource(adflsw)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adflsw *Resource) ImportState(state io.Reader) error {
	adflsw.state = &azurermDataFactoryLinkedServiceWebState{}
	if err := json.NewDecoder(state).Decode(adflsw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adflsw.Type(), adflsw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adflsw *Resource) State() (*azurermDataFactoryLinkedServiceWebState, bool) {
	return adflsw.state, adflsw.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adflsw *Resource) StateMust() *azurermDataFactoryLinkedServiceWebState {
	if adflsw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adflsw.Type(), adflsw.LocalName()))
	}
	return adflsw.state
}

// Args contains the configurations for azurerm_data_factory_linked_service_web.
type Args struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// AuthenticationType: string, required
	AuthenticationType terra.StringValue `hcl:"authentication_type,attr" validate:"required"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationRuntimeName: string, optional
	IntegrationRuntimeName terra.StringValue `hcl:"integration_runtime_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
	// Username: string, optional
	Username terra.StringValue `hcl:"username,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermDataFactoryLinkedServiceWebAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_web.
func (adflsw azurermDataFactoryLinkedServiceWebAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adflsw.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_web.
func (adflsw azurermDataFactoryLinkedServiceWebAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](adflsw.ref.Append("annotations"))
}

// AuthenticationType returns a reference to field authentication_type of azurerm_data_factory_linked_service_web.
func (adflsw azurermDataFactoryLinkedServiceWebAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(adflsw.ref.Append("authentication_type"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_web.
func (adflsw azurermDataFactoryLinkedServiceWebAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(adflsw.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_web.
func (adflsw azurermDataFactoryLinkedServiceWebAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(adflsw.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_web.
func (adflsw azurermDataFactoryLinkedServiceWebAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adflsw.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_web.
func (adflsw azurermDataFactoryLinkedServiceWebAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(adflsw.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_web.
func (adflsw azurermDataFactoryLinkedServiceWebAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adflsw.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_web.
func (adflsw azurermDataFactoryLinkedServiceWebAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adflsw.ref.Append("parameters"))
}

// Password returns a reference to field password of azurerm_data_factory_linked_service_web.
func (adflsw azurermDataFactoryLinkedServiceWebAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(adflsw.ref.Append("password"))
}

// Url returns a reference to field url of azurerm_data_factory_linked_service_web.
func (adflsw azurermDataFactoryLinkedServiceWebAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(adflsw.ref.Append("url"))
}

// Username returns a reference to field username of azurerm_data_factory_linked_service_web.
func (adflsw azurermDataFactoryLinkedServiceWebAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(adflsw.ref.Append("username"))
}

func (adflsw azurermDataFactoryLinkedServiceWebAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adflsw.ref.Append("timeouts"))
}

type azurermDataFactoryLinkedServiceWebState struct {
	AdditionalProperties   map[string]string `json:"additional_properties"`
	Annotations            []string          `json:"annotations"`
	AuthenticationType     string            `json:"authentication_type"`
	DataFactoryId          string            `json:"data_factory_id"`
	Description            string            `json:"description"`
	Id                     string            `json:"id"`
	IntegrationRuntimeName string            `json:"integration_runtime_name"`
	Name                   string            `json:"name"`
	Parameters             map[string]string `json:"parameters"`
	Password               string            `json:"password"`
	Url                    string            `json:"url"`
	Username               string            `json:"username"`
	Timeouts               *TimeoutsState    `json:"timeouts"`
}
