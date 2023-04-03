// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacosmosdbsqlroledefinition "github.com/golingon/terraproviders/azurerm/3.49.0/datacosmosdbsqlroledefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCosmosdbSqlRoleDefinition creates a new instance of [DataCosmosdbSqlRoleDefinition].
func NewDataCosmosdbSqlRoleDefinition(name string, args DataCosmosdbSqlRoleDefinitionArgs) *DataCosmosdbSqlRoleDefinition {
	return &DataCosmosdbSqlRoleDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCosmosdbSqlRoleDefinition)(nil)

// DataCosmosdbSqlRoleDefinition represents the Terraform data resource azurerm_cosmosdb_sql_role_definition.
type DataCosmosdbSqlRoleDefinition struct {
	Name string
	Args DataCosmosdbSqlRoleDefinitionArgs
}

// DataSource returns the Terraform object type for [DataCosmosdbSqlRoleDefinition].
func (csrd *DataCosmosdbSqlRoleDefinition) DataSource() string {
	return "azurerm_cosmosdb_sql_role_definition"
}

// LocalName returns the local name for [DataCosmosdbSqlRoleDefinition].
func (csrd *DataCosmosdbSqlRoleDefinition) LocalName() string {
	return csrd.Name
}

// Configuration returns the configuration (args) for [DataCosmosdbSqlRoleDefinition].
func (csrd *DataCosmosdbSqlRoleDefinition) Configuration() interface{} {
	return csrd.Args
}

// Attributes returns the attributes for [DataCosmosdbSqlRoleDefinition].
func (csrd *DataCosmosdbSqlRoleDefinition) Attributes() dataCosmosdbSqlRoleDefinitionAttributes {
	return dataCosmosdbSqlRoleDefinitionAttributes{ref: terra.ReferenceDataResource(csrd)}
}

// DataCosmosdbSqlRoleDefinitionArgs contains the configurations for azurerm_cosmosdb_sql_role_definition.
type DataCosmosdbSqlRoleDefinitionArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RoleDefinitionId: string, required
	RoleDefinitionId terra.StringValue `hcl:"role_definition_id,attr" validate:"required"`
	// Permissions: min=0
	Permissions []datacosmosdbsqlroledefinition.Permissions `hcl:"permissions,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datacosmosdbsqlroledefinition.Timeouts `hcl:"timeouts,block"`
}
type dataCosmosdbSqlRoleDefinitionAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_sql_role_definition.
func (csrd dataCosmosdbSqlRoleDefinitionAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(csrd.ref.Append("account_name"))
}

// AssignableScopes returns a reference to field assignable_scopes of azurerm_cosmosdb_sql_role_definition.
func (csrd dataCosmosdbSqlRoleDefinitionAttributes) AssignableScopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](csrd.ref.Append("assignable_scopes"))
}

// Id returns a reference to field id of azurerm_cosmosdb_sql_role_definition.
func (csrd dataCosmosdbSqlRoleDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csrd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_sql_role_definition.
func (csrd dataCosmosdbSqlRoleDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csrd.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_sql_role_definition.
func (csrd dataCosmosdbSqlRoleDefinitionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(csrd.ref.Append("resource_group_name"))
}

// RoleDefinitionId returns a reference to field role_definition_id of azurerm_cosmosdb_sql_role_definition.
func (csrd dataCosmosdbSqlRoleDefinitionAttributes) RoleDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(csrd.ref.Append("role_definition_id"))
}

// Type returns a reference to field type of azurerm_cosmosdb_sql_role_definition.
func (csrd dataCosmosdbSqlRoleDefinitionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(csrd.ref.Append("type"))
}

func (csrd dataCosmosdbSqlRoleDefinitionAttributes) Permissions() terra.SetValue[datacosmosdbsqlroledefinition.PermissionsAttributes] {
	return terra.ReferenceAsSet[datacosmosdbsqlroledefinition.PermissionsAttributes](csrd.ref.Append("permissions"))
}

func (csrd dataCosmosdbSqlRoleDefinitionAttributes) Timeouts() datacosmosdbsqlroledefinition.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacosmosdbsqlroledefinition.TimeoutsAttributes](csrd.ref.Append("timeouts"))
}
