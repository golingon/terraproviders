// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataroledefinition "github.com/golingon/terraproviders/azurerm/3.49.0/dataroledefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataRoleDefinition(name string, args DataRoleDefinitionArgs) *DataRoleDefinition {
	return &DataRoleDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRoleDefinition)(nil)

type DataRoleDefinition struct {
	Name string
	Args DataRoleDefinitionArgs
}

func (rd *DataRoleDefinition) DataSource() string {
	return "azurerm_role_definition"
}

func (rd *DataRoleDefinition) LocalName() string {
	return rd.Name
}

func (rd *DataRoleDefinition) Configuration() interface{} {
	return rd.Args
}

func (rd *DataRoleDefinition) Attributes() dataRoleDefinitionAttributes {
	return dataRoleDefinitionAttributes{ref: terra.ReferenceDataResource(rd)}
}

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

func (rd dataRoleDefinitionAttributes) AssignableScopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](rd.ref.Append("assignable_scopes"))
}

func (rd dataRoleDefinitionAttributes) Description() terra.StringValue {
	return terra.ReferenceString(rd.ref.Append("description"))
}

func (rd dataRoleDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rd.ref.Append("id"))
}

func (rd dataRoleDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(rd.ref.Append("name"))
}

func (rd dataRoleDefinitionAttributes) RoleDefinitionId() terra.StringValue {
	return terra.ReferenceString(rd.ref.Append("role_definition_id"))
}

func (rd dataRoleDefinitionAttributes) Scope() terra.StringValue {
	return terra.ReferenceString(rd.ref.Append("scope"))
}

func (rd dataRoleDefinitionAttributes) Type() terra.StringValue {
	return terra.ReferenceString(rd.ref.Append("type"))
}

func (rd dataRoleDefinitionAttributes) Permissions() terra.ListValue[dataroledefinition.PermissionsAttributes] {
	return terra.ReferenceList[dataroledefinition.PermissionsAttributes](rd.ref.Append("permissions"))
}

func (rd dataRoleDefinitionAttributes) Timeouts() dataroledefinition.TimeoutsAttributes {
	return terra.ReferenceSingle[dataroledefinition.TimeoutsAttributes](rd.ref.Append("timeouts"))
}
