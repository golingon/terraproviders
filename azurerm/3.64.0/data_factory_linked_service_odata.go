// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedserviceodata "github.com/golingon/terraproviders/azurerm/3.64.0/datafactorylinkedserviceodata"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceOdata creates a new instance of [DataFactoryLinkedServiceOdata].
func NewDataFactoryLinkedServiceOdata(name string, args DataFactoryLinkedServiceOdataArgs) *DataFactoryLinkedServiceOdata {
	return &DataFactoryLinkedServiceOdata{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceOdata)(nil)

// DataFactoryLinkedServiceOdata represents the Terraform resource azurerm_data_factory_linked_service_odata.
type DataFactoryLinkedServiceOdata struct {
	Name      string
	Args      DataFactoryLinkedServiceOdataArgs
	state     *dataFactoryLinkedServiceOdataState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceOdata].
func (dflso *DataFactoryLinkedServiceOdata) Type() string {
	return "azurerm_data_factory_linked_service_odata"
}

// LocalName returns the local name for [DataFactoryLinkedServiceOdata].
func (dflso *DataFactoryLinkedServiceOdata) LocalName() string {
	return dflso.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceOdata].
func (dflso *DataFactoryLinkedServiceOdata) Configuration() interface{} {
	return dflso.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceOdata].
func (dflso *DataFactoryLinkedServiceOdata) DependOn() terra.Reference {
	return terra.ReferenceResource(dflso)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceOdata] depends_on.
func (dflso *DataFactoryLinkedServiceOdata) Dependencies() terra.Dependencies {
	return dflso.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceOdata].
func (dflso *DataFactoryLinkedServiceOdata) LifecycleManagement() *terra.Lifecycle {
	return dflso.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceOdata].
func (dflso *DataFactoryLinkedServiceOdata) Attributes() dataFactoryLinkedServiceOdataAttributes {
	return dataFactoryLinkedServiceOdataAttributes{ref: terra.ReferenceResource(dflso)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceOdata]'s state.
func (dflso *DataFactoryLinkedServiceOdata) ImportState(av io.Reader) error {
	dflso.state = &dataFactoryLinkedServiceOdataState{}
	if err := json.NewDecoder(av).Decode(dflso.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflso.Type(), dflso.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceOdata] has state.
func (dflso *DataFactoryLinkedServiceOdata) State() (*dataFactoryLinkedServiceOdataState, bool) {
	return dflso.state, dflso.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceOdata]. Panics if the state is nil.
func (dflso *DataFactoryLinkedServiceOdata) StateMust() *dataFactoryLinkedServiceOdataState {
	if dflso.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflso.Type(), dflso.LocalName()))
	}
	return dflso.state
}

// DataFactoryLinkedServiceOdataArgs contains the configurations for azurerm_data_factory_linked_service_odata.
type DataFactoryLinkedServiceOdataArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
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
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
	// BasicAuthentication: optional
	BasicAuthentication *datafactorylinkedserviceodata.BasicAuthentication `hcl:"basic_authentication,block"`
	// Timeouts: optional
	Timeouts *datafactorylinkedserviceodata.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceOdataAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_odata.
func (dflso dataFactoryLinkedServiceOdataAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflso.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_odata.
func (dflso dataFactoryLinkedServiceOdataAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflso.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_odata.
func (dflso dataFactoryLinkedServiceOdataAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflso.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_odata.
func (dflso dataFactoryLinkedServiceOdataAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflso.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_odata.
func (dflso dataFactoryLinkedServiceOdataAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflso.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_odata.
func (dflso dataFactoryLinkedServiceOdataAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflso.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_odata.
func (dflso dataFactoryLinkedServiceOdataAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflso.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_odata.
func (dflso dataFactoryLinkedServiceOdataAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflso.ref.Append("parameters"))
}

// Url returns a reference to field url of azurerm_data_factory_linked_service_odata.
func (dflso dataFactoryLinkedServiceOdataAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(dflso.ref.Append("url"))
}

func (dflso dataFactoryLinkedServiceOdataAttributes) BasicAuthentication() terra.ListValue[datafactorylinkedserviceodata.BasicAuthenticationAttributes] {
	return terra.ReferenceAsList[datafactorylinkedserviceodata.BasicAuthenticationAttributes](dflso.ref.Append("basic_authentication"))
}

func (dflso dataFactoryLinkedServiceOdataAttributes) Timeouts() datafactorylinkedserviceodata.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedserviceodata.TimeoutsAttributes](dflso.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceOdataState struct {
	AdditionalProperties   map[string]string                                        `json:"additional_properties"`
	Annotations            []string                                                 `json:"annotations"`
	DataFactoryId          string                                                   `json:"data_factory_id"`
	Description            string                                                   `json:"description"`
	Id                     string                                                   `json:"id"`
	IntegrationRuntimeName string                                                   `json:"integration_runtime_name"`
	Name                   string                                                   `json:"name"`
	Parameters             map[string]string                                        `json:"parameters"`
	Url                    string                                                   `json:"url"`
	BasicAuthentication    []datafactorylinkedserviceodata.BasicAuthenticationState `json:"basic_authentication"`
	Timeouts               *datafactorylinkedserviceodata.TimeoutsState             `json:"timeouts"`
}
