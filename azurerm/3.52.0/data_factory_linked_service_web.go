// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedserviceweb "github.com/golingon/terraproviders/azurerm/3.52.0/datafactorylinkedserviceweb"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceWeb creates a new instance of [DataFactoryLinkedServiceWeb].
func NewDataFactoryLinkedServiceWeb(name string, args DataFactoryLinkedServiceWebArgs) *DataFactoryLinkedServiceWeb {
	return &DataFactoryLinkedServiceWeb{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceWeb)(nil)

// DataFactoryLinkedServiceWeb represents the Terraform resource azurerm_data_factory_linked_service_web.
type DataFactoryLinkedServiceWeb struct {
	Name      string
	Args      DataFactoryLinkedServiceWebArgs
	state     *dataFactoryLinkedServiceWebState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceWeb].
func (dflsw *DataFactoryLinkedServiceWeb) Type() string {
	return "azurerm_data_factory_linked_service_web"
}

// LocalName returns the local name for [DataFactoryLinkedServiceWeb].
func (dflsw *DataFactoryLinkedServiceWeb) LocalName() string {
	return dflsw.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceWeb].
func (dflsw *DataFactoryLinkedServiceWeb) Configuration() interface{} {
	return dflsw.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceWeb].
func (dflsw *DataFactoryLinkedServiceWeb) DependOn() terra.Reference {
	return terra.ReferenceResource(dflsw)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceWeb] depends_on.
func (dflsw *DataFactoryLinkedServiceWeb) Dependencies() terra.Dependencies {
	return dflsw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceWeb].
func (dflsw *DataFactoryLinkedServiceWeb) LifecycleManagement() *terra.Lifecycle {
	return dflsw.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceWeb].
func (dflsw *DataFactoryLinkedServiceWeb) Attributes() dataFactoryLinkedServiceWebAttributes {
	return dataFactoryLinkedServiceWebAttributes{ref: terra.ReferenceResource(dflsw)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceWeb]'s state.
func (dflsw *DataFactoryLinkedServiceWeb) ImportState(av io.Reader) error {
	dflsw.state = &dataFactoryLinkedServiceWebState{}
	if err := json.NewDecoder(av).Decode(dflsw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflsw.Type(), dflsw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceWeb] has state.
func (dflsw *DataFactoryLinkedServiceWeb) State() (*dataFactoryLinkedServiceWebState, bool) {
	return dflsw.state, dflsw.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceWeb]. Panics if the state is nil.
func (dflsw *DataFactoryLinkedServiceWeb) StateMust() *dataFactoryLinkedServiceWebState {
	if dflsw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflsw.Type(), dflsw.LocalName()))
	}
	return dflsw.state
}

// DataFactoryLinkedServiceWebArgs contains the configurations for azurerm_data_factory_linked_service_web.
type DataFactoryLinkedServiceWebArgs struct {
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
	Timeouts *datafactorylinkedserviceweb.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceWebAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_web.
func (dflsw dataFactoryLinkedServiceWebAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsw.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_web.
func (dflsw dataFactoryLinkedServiceWebAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflsw.ref.Append("annotations"))
}

// AuthenticationType returns a reference to field authentication_type of azurerm_data_factory_linked_service_web.
func (dflsw dataFactoryLinkedServiceWebAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(dflsw.ref.Append("authentication_type"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_web.
func (dflsw dataFactoryLinkedServiceWebAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflsw.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_web.
func (dflsw dataFactoryLinkedServiceWebAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflsw.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_web.
func (dflsw dataFactoryLinkedServiceWebAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflsw.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_web.
func (dflsw dataFactoryLinkedServiceWebAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflsw.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_web.
func (dflsw dataFactoryLinkedServiceWebAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflsw.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_web.
func (dflsw dataFactoryLinkedServiceWebAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsw.ref.Append("parameters"))
}

// Password returns a reference to field password of azurerm_data_factory_linked_service_web.
func (dflsw dataFactoryLinkedServiceWebAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(dflsw.ref.Append("password"))
}

// Url returns a reference to field url of azurerm_data_factory_linked_service_web.
func (dflsw dataFactoryLinkedServiceWebAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(dflsw.ref.Append("url"))
}

// Username returns a reference to field username of azurerm_data_factory_linked_service_web.
func (dflsw dataFactoryLinkedServiceWebAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(dflsw.ref.Append("username"))
}

func (dflsw dataFactoryLinkedServiceWebAttributes) Timeouts() datafactorylinkedserviceweb.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedserviceweb.TimeoutsAttributes](dflsw.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceWebState struct {
	AdditionalProperties   map[string]string                          `json:"additional_properties"`
	Annotations            []string                                   `json:"annotations"`
	AuthenticationType     string                                     `json:"authentication_type"`
	DataFactoryId          string                                     `json:"data_factory_id"`
	Description            string                                     `json:"description"`
	Id                     string                                     `json:"id"`
	IntegrationRuntimeName string                                     `json:"integration_runtime_name"`
	Name                   string                                     `json:"name"`
	Parameters             map[string]string                          `json:"parameters"`
	Password               string                                     `json:"password"`
	Url                    string                                     `json:"url"`
	Username               string                                     `json:"username"`
	Timeouts               *datafactorylinkedserviceweb.TimeoutsState `json:"timeouts"`
}
