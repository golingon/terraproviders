// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedserviceazuredatabricks "github.com/golingon/terraproviders/azurerm/3.55.0/datafactorylinkedserviceazuredatabricks"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceAzureDatabricks creates a new instance of [DataFactoryLinkedServiceAzureDatabricks].
func NewDataFactoryLinkedServiceAzureDatabricks(name string, args DataFactoryLinkedServiceAzureDatabricksArgs) *DataFactoryLinkedServiceAzureDatabricks {
	return &DataFactoryLinkedServiceAzureDatabricks{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceAzureDatabricks)(nil)

// DataFactoryLinkedServiceAzureDatabricks represents the Terraform resource azurerm_data_factory_linked_service_azure_databricks.
type DataFactoryLinkedServiceAzureDatabricks struct {
	Name      string
	Args      DataFactoryLinkedServiceAzureDatabricksArgs
	state     *dataFactoryLinkedServiceAzureDatabricksState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceAzureDatabricks].
func (dflsad *DataFactoryLinkedServiceAzureDatabricks) Type() string {
	return "azurerm_data_factory_linked_service_azure_databricks"
}

// LocalName returns the local name for [DataFactoryLinkedServiceAzureDatabricks].
func (dflsad *DataFactoryLinkedServiceAzureDatabricks) LocalName() string {
	return dflsad.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceAzureDatabricks].
func (dflsad *DataFactoryLinkedServiceAzureDatabricks) Configuration() interface{} {
	return dflsad.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceAzureDatabricks].
func (dflsad *DataFactoryLinkedServiceAzureDatabricks) DependOn() terra.Reference {
	return terra.ReferenceResource(dflsad)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceAzureDatabricks] depends_on.
func (dflsad *DataFactoryLinkedServiceAzureDatabricks) Dependencies() terra.Dependencies {
	return dflsad.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceAzureDatabricks].
func (dflsad *DataFactoryLinkedServiceAzureDatabricks) LifecycleManagement() *terra.Lifecycle {
	return dflsad.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceAzureDatabricks].
func (dflsad *DataFactoryLinkedServiceAzureDatabricks) Attributes() dataFactoryLinkedServiceAzureDatabricksAttributes {
	return dataFactoryLinkedServiceAzureDatabricksAttributes{ref: terra.ReferenceResource(dflsad)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceAzureDatabricks]'s state.
func (dflsad *DataFactoryLinkedServiceAzureDatabricks) ImportState(av io.Reader) error {
	dflsad.state = &dataFactoryLinkedServiceAzureDatabricksState{}
	if err := json.NewDecoder(av).Decode(dflsad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflsad.Type(), dflsad.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceAzureDatabricks] has state.
func (dflsad *DataFactoryLinkedServiceAzureDatabricks) State() (*dataFactoryLinkedServiceAzureDatabricksState, bool) {
	return dflsad.state, dflsad.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceAzureDatabricks]. Panics if the state is nil.
func (dflsad *DataFactoryLinkedServiceAzureDatabricks) StateMust() *dataFactoryLinkedServiceAzureDatabricksState {
	if dflsad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflsad.Type(), dflsad.LocalName()))
	}
	return dflsad.state
}

// DataFactoryLinkedServiceAzureDatabricksArgs contains the configurations for azurerm_data_factory_linked_service_azure_databricks.
type DataFactoryLinkedServiceAzureDatabricksArgs struct {
	// AccessToken: string, optional
	AccessToken terra.StringValue `hcl:"access_token,attr"`
	// AdbDomain: string, required
	AdbDomain terra.StringValue `hcl:"adb_domain,attr" validate:"required"`
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ExistingClusterId: string, optional
	ExistingClusterId terra.StringValue `hcl:"existing_cluster_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationRuntimeName: string, optional
	IntegrationRuntimeName terra.StringValue `hcl:"integration_runtime_name,attr"`
	// MsiWorkSpaceResourceId: string, optional
	MsiWorkSpaceResourceId terra.StringValue `hcl:"msi_work_space_resource_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// InstancePool: optional
	InstancePool *datafactorylinkedserviceazuredatabricks.InstancePool `hcl:"instance_pool,block"`
	// KeyVaultPassword: optional
	KeyVaultPassword *datafactorylinkedserviceazuredatabricks.KeyVaultPassword `hcl:"key_vault_password,block"`
	// NewClusterConfig: optional
	NewClusterConfig *datafactorylinkedserviceazuredatabricks.NewClusterConfig `hcl:"new_cluster_config,block"`
	// Timeouts: optional
	Timeouts *datafactorylinkedserviceazuredatabricks.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceAzureDatabricksAttributes struct {
	ref terra.Reference
}

// AccessToken returns a reference to field access_token of azurerm_data_factory_linked_service_azure_databricks.
func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) AccessToken() terra.StringValue {
	return terra.ReferenceAsString(dflsad.ref.Append("access_token"))
}

// AdbDomain returns a reference to field adb_domain of azurerm_data_factory_linked_service_azure_databricks.
func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) AdbDomain() terra.StringValue {
	return terra.ReferenceAsString(dflsad.ref.Append("adb_domain"))
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_azure_databricks.
func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsad.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_azure_databricks.
func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflsad.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_azure_databricks.
func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflsad.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_azure_databricks.
func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflsad.ref.Append("description"))
}

// ExistingClusterId returns a reference to field existing_cluster_id of azurerm_data_factory_linked_service_azure_databricks.
func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) ExistingClusterId() terra.StringValue {
	return terra.ReferenceAsString(dflsad.ref.Append("existing_cluster_id"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_azure_databricks.
func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflsad.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_azure_databricks.
func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflsad.ref.Append("integration_runtime_name"))
}

// MsiWorkSpaceResourceId returns a reference to field msi_work_space_resource_id of azurerm_data_factory_linked_service_azure_databricks.
func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) MsiWorkSpaceResourceId() terra.StringValue {
	return terra.ReferenceAsString(dflsad.ref.Append("msi_work_space_resource_id"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_azure_databricks.
func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflsad.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_azure_databricks.
func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsad.ref.Append("parameters"))
}

func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) InstancePool() terra.ListValue[datafactorylinkedserviceazuredatabricks.InstancePoolAttributes] {
	return terra.ReferenceAsList[datafactorylinkedserviceazuredatabricks.InstancePoolAttributes](dflsad.ref.Append("instance_pool"))
}

func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) KeyVaultPassword() terra.ListValue[datafactorylinkedserviceazuredatabricks.KeyVaultPasswordAttributes] {
	return terra.ReferenceAsList[datafactorylinkedserviceazuredatabricks.KeyVaultPasswordAttributes](dflsad.ref.Append("key_vault_password"))
}

func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) NewClusterConfig() terra.ListValue[datafactorylinkedserviceazuredatabricks.NewClusterConfigAttributes] {
	return terra.ReferenceAsList[datafactorylinkedserviceazuredatabricks.NewClusterConfigAttributes](dflsad.ref.Append("new_cluster_config"))
}

func (dflsad dataFactoryLinkedServiceAzureDatabricksAttributes) Timeouts() datafactorylinkedserviceazuredatabricks.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedserviceazuredatabricks.TimeoutsAttributes](dflsad.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceAzureDatabricksState struct {
	AccessToken            string                                                          `json:"access_token"`
	AdbDomain              string                                                          `json:"adb_domain"`
	AdditionalProperties   map[string]string                                               `json:"additional_properties"`
	Annotations            []string                                                        `json:"annotations"`
	DataFactoryId          string                                                          `json:"data_factory_id"`
	Description            string                                                          `json:"description"`
	ExistingClusterId      string                                                          `json:"existing_cluster_id"`
	Id                     string                                                          `json:"id"`
	IntegrationRuntimeName string                                                          `json:"integration_runtime_name"`
	MsiWorkSpaceResourceId string                                                          `json:"msi_work_space_resource_id"`
	Name                   string                                                          `json:"name"`
	Parameters             map[string]string                                               `json:"parameters"`
	InstancePool           []datafactorylinkedserviceazuredatabricks.InstancePoolState     `json:"instance_pool"`
	KeyVaultPassword       []datafactorylinkedserviceazuredatabricks.KeyVaultPasswordState `json:"key_vault_password"`
	NewClusterConfig       []datafactorylinkedserviceazuredatabricks.NewClusterConfigState `json:"new_cluster_config"`
	Timeouts               *datafactorylinkedserviceazuredatabricks.TimeoutsState          `json:"timeouts"`
}
