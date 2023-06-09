// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedservicemysql "github.com/golingon/terraproviders/azurerm/3.52.0/datafactorylinkedservicemysql"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceMysql creates a new instance of [DataFactoryLinkedServiceMysql].
func NewDataFactoryLinkedServiceMysql(name string, args DataFactoryLinkedServiceMysqlArgs) *DataFactoryLinkedServiceMysql {
	return &DataFactoryLinkedServiceMysql{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceMysql)(nil)

// DataFactoryLinkedServiceMysql represents the Terraform resource azurerm_data_factory_linked_service_mysql.
type DataFactoryLinkedServiceMysql struct {
	Name      string
	Args      DataFactoryLinkedServiceMysqlArgs
	state     *dataFactoryLinkedServiceMysqlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceMysql].
func (dflsm *DataFactoryLinkedServiceMysql) Type() string {
	return "azurerm_data_factory_linked_service_mysql"
}

// LocalName returns the local name for [DataFactoryLinkedServiceMysql].
func (dflsm *DataFactoryLinkedServiceMysql) LocalName() string {
	return dflsm.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceMysql].
func (dflsm *DataFactoryLinkedServiceMysql) Configuration() interface{} {
	return dflsm.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceMysql].
func (dflsm *DataFactoryLinkedServiceMysql) DependOn() terra.Reference {
	return terra.ReferenceResource(dflsm)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceMysql] depends_on.
func (dflsm *DataFactoryLinkedServiceMysql) Dependencies() terra.Dependencies {
	return dflsm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceMysql].
func (dflsm *DataFactoryLinkedServiceMysql) LifecycleManagement() *terra.Lifecycle {
	return dflsm.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceMysql].
func (dflsm *DataFactoryLinkedServiceMysql) Attributes() dataFactoryLinkedServiceMysqlAttributes {
	return dataFactoryLinkedServiceMysqlAttributes{ref: terra.ReferenceResource(dflsm)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceMysql]'s state.
func (dflsm *DataFactoryLinkedServiceMysql) ImportState(av io.Reader) error {
	dflsm.state = &dataFactoryLinkedServiceMysqlState{}
	if err := json.NewDecoder(av).Decode(dflsm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflsm.Type(), dflsm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceMysql] has state.
func (dflsm *DataFactoryLinkedServiceMysql) State() (*dataFactoryLinkedServiceMysqlState, bool) {
	return dflsm.state, dflsm.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceMysql]. Panics if the state is nil.
func (dflsm *DataFactoryLinkedServiceMysql) StateMust() *dataFactoryLinkedServiceMysqlState {
	if dflsm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflsm.Type(), dflsm.LocalName()))
	}
	return dflsm.state
}

// DataFactoryLinkedServiceMysqlArgs contains the configurations for azurerm_data_factory_linked_service_mysql.
type DataFactoryLinkedServiceMysqlArgs struct {
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
	Timeouts *datafactorylinkedservicemysql.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceMysqlAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_mysql.
func (dflsm dataFactoryLinkedServiceMysqlAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsm.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_mysql.
func (dflsm dataFactoryLinkedServiceMysqlAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflsm.ref.Append("annotations"))
}

// ConnectionString returns a reference to field connection_string of azurerm_data_factory_linked_service_mysql.
func (dflsm dataFactoryLinkedServiceMysqlAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(dflsm.ref.Append("connection_string"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_mysql.
func (dflsm dataFactoryLinkedServiceMysqlAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflsm.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_mysql.
func (dflsm dataFactoryLinkedServiceMysqlAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflsm.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_mysql.
func (dflsm dataFactoryLinkedServiceMysqlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflsm.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_mysql.
func (dflsm dataFactoryLinkedServiceMysqlAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflsm.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_mysql.
func (dflsm dataFactoryLinkedServiceMysqlAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflsm.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_mysql.
func (dflsm dataFactoryLinkedServiceMysqlAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsm.ref.Append("parameters"))
}

func (dflsm dataFactoryLinkedServiceMysqlAttributes) Timeouts() datafactorylinkedservicemysql.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedservicemysql.TimeoutsAttributes](dflsm.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceMysqlState struct {
	AdditionalProperties   map[string]string                            `json:"additional_properties"`
	Annotations            []string                                     `json:"annotations"`
	ConnectionString       string                                       `json:"connection_string"`
	DataFactoryId          string                                       `json:"data_factory_id"`
	Description            string                                       `json:"description"`
	Id                     string                                       `json:"id"`
	IntegrationRuntimeName string                                       `json:"integration_runtime_name"`
	Name                   string                                       `json:"name"`
	Parameters             map[string]string                            `json:"parameters"`
	Timeouts               *datafactorylinkedservicemysql.TimeoutsState `json:"timeouts"`
}
