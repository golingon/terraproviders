// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedservicesnowflake "github.com/golingon/terraproviders/azurerm/3.68.0/datafactorylinkedservicesnowflake"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceSnowflake creates a new instance of [DataFactoryLinkedServiceSnowflake].
func NewDataFactoryLinkedServiceSnowflake(name string, args DataFactoryLinkedServiceSnowflakeArgs) *DataFactoryLinkedServiceSnowflake {
	return &DataFactoryLinkedServiceSnowflake{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceSnowflake)(nil)

// DataFactoryLinkedServiceSnowflake represents the Terraform resource azurerm_data_factory_linked_service_snowflake.
type DataFactoryLinkedServiceSnowflake struct {
	Name      string
	Args      DataFactoryLinkedServiceSnowflakeArgs
	state     *dataFactoryLinkedServiceSnowflakeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceSnowflake].
func (dflss *DataFactoryLinkedServiceSnowflake) Type() string {
	return "azurerm_data_factory_linked_service_snowflake"
}

// LocalName returns the local name for [DataFactoryLinkedServiceSnowflake].
func (dflss *DataFactoryLinkedServiceSnowflake) LocalName() string {
	return dflss.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceSnowflake].
func (dflss *DataFactoryLinkedServiceSnowflake) Configuration() interface{} {
	return dflss.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceSnowflake].
func (dflss *DataFactoryLinkedServiceSnowflake) DependOn() terra.Reference {
	return terra.ReferenceResource(dflss)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceSnowflake] depends_on.
func (dflss *DataFactoryLinkedServiceSnowflake) Dependencies() terra.Dependencies {
	return dflss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceSnowflake].
func (dflss *DataFactoryLinkedServiceSnowflake) LifecycleManagement() *terra.Lifecycle {
	return dflss.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceSnowflake].
func (dflss *DataFactoryLinkedServiceSnowflake) Attributes() dataFactoryLinkedServiceSnowflakeAttributes {
	return dataFactoryLinkedServiceSnowflakeAttributes{ref: terra.ReferenceResource(dflss)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceSnowflake]'s state.
func (dflss *DataFactoryLinkedServiceSnowflake) ImportState(av io.Reader) error {
	dflss.state = &dataFactoryLinkedServiceSnowflakeState{}
	if err := json.NewDecoder(av).Decode(dflss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflss.Type(), dflss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceSnowflake] has state.
func (dflss *DataFactoryLinkedServiceSnowflake) State() (*dataFactoryLinkedServiceSnowflakeState, bool) {
	return dflss.state, dflss.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceSnowflake]. Panics if the state is nil.
func (dflss *DataFactoryLinkedServiceSnowflake) StateMust() *dataFactoryLinkedServiceSnowflakeState {
	if dflss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflss.Type(), dflss.LocalName()))
	}
	return dflss.state
}

// DataFactoryLinkedServiceSnowflakeArgs contains the configurations for azurerm_data_factory_linked_service_snowflake.
type DataFactoryLinkedServiceSnowflakeArgs struct {
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
	// KeyVaultPassword: optional
	KeyVaultPassword *datafactorylinkedservicesnowflake.KeyVaultPassword `hcl:"key_vault_password,block"`
	// Timeouts: optional
	Timeouts *datafactorylinkedservicesnowflake.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceSnowflakeAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_snowflake.
func (dflss dataFactoryLinkedServiceSnowflakeAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflss.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_snowflake.
func (dflss dataFactoryLinkedServiceSnowflakeAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflss.ref.Append("annotations"))
}

// ConnectionString returns a reference to field connection_string of azurerm_data_factory_linked_service_snowflake.
func (dflss dataFactoryLinkedServiceSnowflakeAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("connection_string"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_snowflake.
func (dflss dataFactoryLinkedServiceSnowflakeAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_snowflake.
func (dflss dataFactoryLinkedServiceSnowflakeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_snowflake.
func (dflss dataFactoryLinkedServiceSnowflakeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_snowflake.
func (dflss dataFactoryLinkedServiceSnowflakeAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_snowflake.
func (dflss dataFactoryLinkedServiceSnowflakeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_snowflake.
func (dflss dataFactoryLinkedServiceSnowflakeAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflss.ref.Append("parameters"))
}

func (dflss dataFactoryLinkedServiceSnowflakeAttributes) KeyVaultPassword() terra.ListValue[datafactorylinkedservicesnowflake.KeyVaultPasswordAttributes] {
	return terra.ReferenceAsList[datafactorylinkedservicesnowflake.KeyVaultPasswordAttributes](dflss.ref.Append("key_vault_password"))
}

func (dflss dataFactoryLinkedServiceSnowflakeAttributes) Timeouts() datafactorylinkedservicesnowflake.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedservicesnowflake.TimeoutsAttributes](dflss.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceSnowflakeState struct {
	AdditionalProperties   map[string]string                                         `json:"additional_properties"`
	Annotations            []string                                                  `json:"annotations"`
	ConnectionString       string                                                    `json:"connection_string"`
	DataFactoryId          string                                                    `json:"data_factory_id"`
	Description            string                                                    `json:"description"`
	Id                     string                                                    `json:"id"`
	IntegrationRuntimeName string                                                    `json:"integration_runtime_name"`
	Name                   string                                                    `json:"name"`
	Parameters             map[string]string                                         `json:"parameters"`
	KeyVaultPassword       []datafactorylinkedservicesnowflake.KeyVaultPasswordState `json:"key_vault_password"`
	Timeouts               *datafactorylinkedservicesnowflake.TimeoutsState          `json:"timeouts"`
}
