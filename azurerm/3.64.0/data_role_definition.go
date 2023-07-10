// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataroledefinition "github.com/golingon/terraproviders/azurerm/3.64.0/dataroledefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRoleDefinition creates a new instance of [DataRoleDefinition].
func NewDataRoleDefinition(name string, args DataRoleDefinitionArgs) *DataRoleDefinition {
	return &DataRoleDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRoleDefinition)(nil)

// DataRoleDefinition represents the Terraform data resource azurerm_role_definition.
type DataRoleDefinition struct {
	Name string
	Args DataRoleDefinitionArgs
}

// DataSource returns the Terraform object type for [DataRoleDefinition].
func (rd *DataRoleDefinition) DataSource() string {
	return "azurerm_role_definition"
}

// LocalName returns the local name for [DataRoleDefinition].
func (rd *DataRoleDefinition) LocalName() string {
	return rd.Name
}

// Configuration returns the configuration (args) for [DataRoleDefinition].
func (rd *DataRoleDefinition) Configuration() interface{} {
	return rd.Args
}

// Attributes returns the attributes for [DataRoleDefinition].
func (rd *DataRoleDefinition) Attributes() dataRoleDefinitionAttributes {
	return dataRoleDefinitionAttributes{ref: terra.ReferenceDataResource(rd)}
}

// DataRoleDefinitionArgs contains the configurations for azurerm_role_definition.
type DataRoleDefinitionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// RoleDefinitionId: string, optional
	RoleDefinitionId terra.StringValue `hcl:"role_definition_id,attr"`
	// Scope: string, optional
	Scope terra.StringValue `hcl:"scope,attr"`
	// Permissions: min=0
	Permissions []dataroledefinition.Permissions `hcl:"permissions,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataroledefinition.Timeouts `hcl:"timeouts,block"`
}
type dataRoleDefinitionAttributes struct {
	ref terra.Reference
}

// AssignableScopes returns a reference to field assignable_scopes of azurerm_role_definition.
func (rd dataRoleDefinitionAttributes) AssignableScopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rd.ref.Append("assignable_scopes"))
}

// Description returns a reference to field description of azurerm_role_definition.
func (rd dataRoleDefinitionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_role_definition.
func (rd dataRoleDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_role_definition.
func (rd dataRoleDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("name"))
}

// RoleDefinitionId returns a reference to field role_definition_id of azurerm_role_definition.
func (rd dataRoleDefinitionAttributes) RoleDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("role_definition_id"))
}

// Scope returns a reference to field scope of azurerm_role_definition.
func (rd dataRoleDefinitionAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("scope"))
}

// Type returns a reference to field type of azurerm_role_definition.
func (rd dataRoleDefinitionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(rd.ref.Append("type"))
}

func (rd dataRoleDefinitionAttributes) Permissions() terra.ListValue[dataroledefinition.PermissionsAttributes] {
	return terra.ReferenceAsList[dataroledefinition.PermissionsAttributes](rd.ref.Append("permissions"))
}

func (rd dataRoleDefinitionAttributes) Timeouts() dataroledefinition.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataroledefinition.TimeoutsAttributes](rd.ref.Append("timeouts"))
}
