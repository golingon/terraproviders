// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedservicepostgresql "github.com/golingon/terraproviders/azurerm/3.49.0/datafactorylinkedservicepostgresql"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDataFactoryLinkedServicePostgresql(name string, args DataFactoryLinkedServicePostgresqlArgs) *DataFactoryLinkedServicePostgresql {
	return &DataFactoryLinkedServicePostgresql{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServicePostgresql)(nil)

type DataFactoryLinkedServicePostgresql struct {
	Name  string
	Args  DataFactoryLinkedServicePostgresqlArgs
	state *dataFactoryLinkedServicePostgresqlState
}

func (dflsp *DataFactoryLinkedServicePostgresql) Type() string {
	return "azurerm_data_factory_linked_service_postgresql"
}

func (dflsp *DataFactoryLinkedServicePostgresql) LocalName() string {
	return dflsp.Name
}

func (dflsp *DataFactoryLinkedServicePostgresql) Configuration() interface{} {
	return dflsp.Args
}

func (dflsp *DataFactoryLinkedServicePostgresql) Attributes() dataFactoryLinkedServicePostgresqlAttributes {
	return dataFactoryLinkedServicePostgresqlAttributes{ref: terra.ReferenceResource(dflsp)}
}

func (dflsp *DataFactoryLinkedServicePostgresql) ImportState(av io.Reader) error {
	dflsp.state = &dataFactoryLinkedServicePostgresqlState{}
	if err := json.NewDecoder(av).Decode(dflsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflsp.Type(), dflsp.LocalName(), err)
	}
	return nil
}

func (dflsp *DataFactoryLinkedServicePostgresql) State() (*dataFactoryLinkedServicePostgresqlState, bool) {
	return dflsp.state, dflsp.state != nil
}

func (dflsp *DataFactoryLinkedServicePostgresql) StateMust() *dataFactoryLinkedServicePostgresqlState {
	if dflsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflsp.Type(), dflsp.LocalName()))
	}
	return dflsp.state
}

func (dflsp *DataFactoryLinkedServicePostgresql) DependOn() terra.Reference {
	return terra.ReferenceResource(dflsp)
}

type DataFactoryLinkedServicePostgresqlArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// ConnectionString: string, required
	ConnectionString terra.StringValue `hcl:"connection_string,attr" validate:"required"`
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
	// Timeouts: optional
	Timeouts *datafactorylinkedservicepostgresql.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that DataFactoryLinkedServicePostgresql depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dataFactoryLinkedServicePostgresqlAttributes struct {
	ref terra.Reference
}

func (dflsp dataFactoryLinkedServicePostgresqlAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dflsp.ref.Append("additional_properties"))
}

func (dflsp dataFactoryLinkedServicePostgresqlAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dflsp.ref.Append("annotations"))
}

func (dflsp dataFactoryLinkedServicePostgresqlAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceString(dflsp.ref.Append("connection_string"))
}

func (dflsp dataFactoryLinkedServicePostgresqlAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceString(dflsp.ref.Append("data_factory_id"))
}

func (dflsp dataFactoryLinkedServicePostgresqlAttributes) Description() terra.StringValue {
	return terra.ReferenceString(dflsp.ref.Append("description"))
}

func (dflsp dataFactoryLinkedServicePostgresqlAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dflsp.ref.Append("id"))
}

func (dflsp dataFactoryLinkedServicePostgresqlAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceString(dflsp.ref.Append("integration_runtime_name"))
}

func (dflsp dataFactoryLinkedServicePostgresqlAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dflsp.ref.Append("name"))
}

func (dflsp dataFactoryLinkedServicePostgresqlAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dflsp.ref.Append("parameters"))
}

func (dflsp dataFactoryLinkedServicePostgresqlAttributes) Timeouts() datafactorylinkedservicepostgresql.TimeoutsAttributes {
	return terra.ReferenceSingle[datafactorylinkedservicepostgresql.TimeoutsAttributes](dflsp.ref.Append("timeouts"))
}

type dataFactoryLinkedServicePostgresqlState struct {
	AdditionalProperties   map[string]string                                 `json:"additional_properties"`
	Annotations            []string                                          `json:"annotations"`
	ConnectionString       string                                            `json:"connection_string"`
	DataFactoryId          string                                            `json:"data_factory_id"`
	Description            string                                            `json:"description"`
	Id                     string                                            `json:"id"`
	IntegrationRuntimeName string                                            `json:"integration_runtime_name"`
	Name                   string                                            `json:"name"`
	Parameters             map[string]string                                 `json:"parameters"`
	Timeouts               *datafactorylinkedservicepostgresql.TimeoutsState `json:"timeouts"`
}
