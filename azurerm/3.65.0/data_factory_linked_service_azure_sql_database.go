// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedserviceazuresqldatabase "github.com/golingon/terraproviders/azurerm/3.65.0/datafactorylinkedserviceazuresqldatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceAzureSqlDatabase creates a new instance of [DataFactoryLinkedServiceAzureSqlDatabase].
func NewDataFactoryLinkedServiceAzureSqlDatabase(name string, args DataFactoryLinkedServiceAzureSqlDatabaseArgs) *DataFactoryLinkedServiceAzureSqlDatabase {
	return &DataFactoryLinkedServiceAzureSqlDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceAzureSqlDatabase)(nil)

// DataFactoryLinkedServiceAzureSqlDatabase represents the Terraform resource azurerm_data_factory_linked_service_azure_sql_database.
type DataFactoryLinkedServiceAzureSqlDatabase struct {
	Name      string
	Args      DataFactoryLinkedServiceAzureSqlDatabaseArgs
	state     *dataFactoryLinkedServiceAzureSqlDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceAzureSqlDatabase].
func (dflsasd *DataFactoryLinkedServiceAzureSqlDatabase) Type() string {
	return "azurerm_data_factory_linked_service_azure_sql_database"
}

// LocalName returns the local name for [DataFactoryLinkedServiceAzureSqlDatabase].
func (dflsasd *DataFactoryLinkedServiceAzureSqlDatabase) LocalName() string {
	return dflsasd.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceAzureSqlDatabase].
func (dflsasd *DataFactoryLinkedServiceAzureSqlDatabase) Configuration() interface{} {
	return dflsasd.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceAzureSqlDatabase].
func (dflsasd *DataFactoryLinkedServiceAzureSqlDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(dflsasd)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceAzureSqlDatabase] depends_on.
func (dflsasd *DataFactoryLinkedServiceAzureSqlDatabase) Dependencies() terra.Dependencies {
	return dflsasd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceAzureSqlDatabase].
func (dflsasd *DataFactoryLinkedServiceAzureSqlDatabase) LifecycleManagement() *terra.Lifecycle {
	return dflsasd.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceAzureSqlDatabase].
func (dflsasd *DataFactoryLinkedServiceAzureSqlDatabase) Attributes() dataFactoryLinkedServiceAzureSqlDatabaseAttributes {
	return dataFactoryLinkedServiceAzureSqlDatabaseAttributes{ref: terra.ReferenceResource(dflsasd)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceAzureSqlDatabase]'s state.
func (dflsasd *DataFactoryLinkedServiceAzureSqlDatabase) ImportState(av io.Reader) error {
	dflsasd.state = &dataFactoryLinkedServiceAzureSqlDatabaseState{}
	if err := json.NewDecoder(av).Decode(dflsasd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflsasd.Type(), dflsasd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceAzureSqlDatabase] has state.
func (dflsasd *DataFactoryLinkedServiceAzureSqlDatabase) State() (*dataFactoryLinkedServiceAzureSqlDatabaseState, bool) {
	return dflsasd.state, dflsasd.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceAzureSqlDatabase]. Panics if the state is nil.
func (dflsasd *DataFactoryLinkedServiceAzureSqlDatabase) StateMust() *dataFactoryLinkedServiceAzureSqlDatabaseState {
	if dflsasd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflsasd.Type(), dflsasd.LocalName()))
	}
	return dflsasd.state
}

// DataFactoryLinkedServiceAzureSqlDatabaseArgs contains the configurations for azurerm_data_factory_linked_service_azure_sql_database.
type DataFactoryLinkedServiceAzureSqlDatabaseArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// ConnectionString: string, optional
	ConnectionString terra.StringValue `hcl:"connection_string,attr"`
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
	// ServicePrincipalId: string, optional
	ServicePrincipalId terra.StringValue `hcl:"service_principal_id,attr"`
	// ServicePrincipalKey: string, optional
	ServicePrincipalKey terra.StringValue `hcl:"service_principal_key,attr"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// UseManagedIdentity: bool, optional
	UseManagedIdentity terra.BoolValue `hcl:"use_managed_identity,attr"`
	// KeyVaultConnectionString: optional
	KeyVaultConnectionString *datafactorylinkedserviceazuresqldatabase.KeyVaultConnectionString `hcl:"key_vault_connection_string,block"`
	// KeyVaultPassword: optional
	KeyVaultPassword *datafactorylinkedserviceazuresqldatabase.KeyVaultPassword `hcl:"key_vault_password,block"`
	// Timeouts: optional
	Timeouts *datafactorylinkedserviceazuresqldatabase.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceAzureSqlDatabaseAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_azure_sql_database.
func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsasd.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_azure_sql_database.
func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflsasd.ref.Append("annotations"))
}

// ConnectionString returns a reference to field connection_string of azurerm_data_factory_linked_service_azure_sql_database.
func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(dflsasd.ref.Append("connection_string"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_azure_sql_database.
func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflsasd.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_azure_sql_database.
func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflsasd.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_azure_sql_database.
func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflsasd.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_azure_sql_database.
func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflsasd.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_azure_sql_database.
func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflsasd.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_azure_sql_database.
func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsasd.ref.Append("parameters"))
}

// ServicePrincipalId returns a reference to field service_principal_id of azurerm_data_factory_linked_service_azure_sql_database.
func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) ServicePrincipalId() terra.StringValue {
	return terra.ReferenceAsString(dflsasd.ref.Append("service_principal_id"))
}

// ServicePrincipalKey returns a reference to field service_principal_key of azurerm_data_factory_linked_service_azure_sql_database.
func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) ServicePrincipalKey() terra.StringValue {
	return terra.ReferenceAsString(dflsasd.ref.Append("service_principal_key"))
}

// TenantId returns a reference to field tenant_id of azurerm_data_factory_linked_service_azure_sql_database.
func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(dflsasd.ref.Append("tenant_id"))
}

// UseManagedIdentity returns a reference to field use_managed_identity of azurerm_data_factory_linked_service_azure_sql_database.
func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) UseManagedIdentity() terra.BoolValue {
	return terra.ReferenceAsBool(dflsasd.ref.Append("use_managed_identity"))
}

func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) KeyVaultConnectionString() terra.ListValue[datafactorylinkedserviceazuresqldatabase.KeyVaultConnectionStringAttributes] {
	return terra.ReferenceAsList[datafactorylinkedserviceazuresqldatabase.KeyVaultConnectionStringAttributes](dflsasd.ref.Append("key_vault_connection_string"))
}

func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) KeyVaultPassword() terra.ListValue[datafactorylinkedserviceazuresqldatabase.KeyVaultPasswordAttributes] {
	return terra.ReferenceAsList[datafactorylinkedserviceazuresqldatabase.KeyVaultPasswordAttributes](dflsasd.ref.Append("key_vault_password"))
}

func (dflsasd dataFactoryLinkedServiceAzureSqlDatabaseAttributes) Timeouts() datafactorylinkedserviceazuresqldatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedserviceazuresqldatabase.TimeoutsAttributes](dflsasd.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceAzureSqlDatabaseState struct {
	AdditionalProperties     map[string]string                                                        `json:"additional_properties"`
	Annotations              []string                                                                 `json:"annotations"`
	ConnectionString         string                                                                   `json:"connection_string"`
	DataFactoryId            string                                                                   `json:"data_factory_id"`
	Description              string                                                                   `json:"description"`
	Id                       string                                                                   `json:"id"`
	IntegrationRuntimeName   string                                                                   `json:"integration_runtime_name"`
	Name                     string                                                                   `json:"name"`
	Parameters               map[string]string                                                        `json:"parameters"`
	ServicePrincipalId       string                                                                   `json:"service_principal_id"`
	ServicePrincipalKey      string                                                                   `json:"service_principal_key"`
	TenantId                 string                                                                   `json:"tenant_id"`
	UseManagedIdentity       bool                                                                     `json:"use_managed_identity"`
	KeyVaultConnectionString []datafactorylinkedserviceazuresqldatabase.KeyVaultConnectionStringState `json:"key_vault_connection_string"`
	KeyVaultPassword         []datafactorylinkedserviceazuresqldatabase.KeyVaultPasswordState         `json:"key_vault_password"`
	Timeouts                 *datafactorylinkedserviceazuresqldatabase.TimeoutsState                  `json:"timeouts"`
}
