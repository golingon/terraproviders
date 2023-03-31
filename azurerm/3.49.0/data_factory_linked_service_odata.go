// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedserviceodata "github.com/golingon/terraproviders/azurerm/3.49.0/datafactorylinkedserviceodata"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDataFactoryLinkedServiceOdata(name string, args DataFactoryLinkedServiceOdataArgs) *DataFactoryLinkedServiceOdata {
	return &DataFactoryLinkedServiceOdata{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceOdata)(nil)

type DataFactoryLinkedServiceOdata struct {
	Name  string
	Args  DataFactoryLinkedServiceOdataArgs
	state *dataFactoryLinkedServiceOdataState
}

func (dflso *DataFactoryLinkedServiceOdata) Type() string {
	return "azurerm_data_factory_linked_service_odata"
}

func (dflso *DataFactoryLinkedServiceOdata) LocalName() string {
	return dflso.Name
}

func (dflso *DataFactoryLinkedServiceOdata) Configuration() interface{} {
	return dflso.Args
}

func (dflso *DataFactoryLinkedServiceOdata) Attributes() dataFactoryLinkedServiceOdataAttributes {
	return dataFactoryLinkedServiceOdataAttributes{ref: terra.ReferenceResource(dflso)}
}

func (dflso *DataFactoryLinkedServiceOdata) ImportState(av io.Reader) error {
	dflso.state = &dataFactoryLinkedServiceOdataState{}
	if err := json.NewDecoder(av).Decode(dflso.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflso.Type(), dflso.LocalName(), err)
	}
	return nil
}

func (dflso *DataFactoryLinkedServiceOdata) State() (*dataFactoryLinkedServiceOdataState, bool) {
	return dflso.state, dflso.state != nil
}

func (dflso *DataFactoryLinkedServiceOdata) StateMust() *dataFactoryLinkedServiceOdataState {
	if dflso.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflso.Type(), dflso.LocalName()))
	}
	return dflso.state
}

func (dflso *DataFactoryLinkedServiceOdata) DependOn() terra.Reference {
	return terra.ReferenceResource(dflso)
}

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
	// DependsOn contains resources that DataFactoryLinkedServiceOdata depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dataFactoryLinkedServiceOdataAttributes struct {
	ref terra.Reference
}

func (dflso dataFactoryLinkedServiceOdataAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dflso.ref.Append("additional_properties"))
}

func (dflso dataFactoryLinkedServiceOdataAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dflso.ref.Append("annotations"))
}

func (dflso dataFactoryLinkedServiceOdataAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceString(dflso.ref.Append("data_factory_id"))
}

func (dflso dataFactoryLinkedServiceOdataAttributes) Description() terra.StringValue {
	return terra.ReferenceString(dflso.ref.Append("description"))
}

func (dflso dataFactoryLinkedServiceOdataAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dflso.ref.Append("id"))
}

func (dflso dataFactoryLinkedServiceOdataAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceString(dflso.ref.Append("integration_runtime_name"))
}

func (dflso dataFactoryLinkedServiceOdataAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dflso.ref.Append("name"))
}

func (dflso dataFactoryLinkedServiceOdataAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dflso.ref.Append("parameters"))
}

func (dflso dataFactoryLinkedServiceOdataAttributes) Url() terra.StringValue {
	return terra.ReferenceString(dflso.ref.Append("url"))
}

func (dflso dataFactoryLinkedServiceOdataAttributes) BasicAuthentication() terra.ListValue[datafactorylinkedserviceodata.BasicAuthenticationAttributes] {
	return terra.ReferenceList[datafactorylinkedserviceodata.BasicAuthenticationAttributes](dflso.ref.Append("basic_authentication"))
}

func (dflso dataFactoryLinkedServiceOdataAttributes) Timeouts() datafactorylinkedserviceodata.TimeoutsAttributes {
	return terra.ReferenceSingle[datafactorylinkedserviceodata.TimeoutsAttributes](dflso.ref.Append("timeouts"))
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
