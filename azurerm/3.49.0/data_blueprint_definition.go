// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datablueprintdefinition "github.com/golingon/terraproviders/azurerm/3.49.0/datablueprintdefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataBlueprintDefinition(name string, args DataBlueprintDefinitionArgs) *DataBlueprintDefinition {
	return &DataBlueprintDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBlueprintDefinition)(nil)

type DataBlueprintDefinition struct {
	Name string
	Args DataBlueprintDefinitionArgs
}

func (bd *DataBlueprintDefinition) DataSource() string {
	return "azurerm_blueprint_definition"
}

func (bd *DataBlueprintDefinition) LocalName() string {
	return bd.Name
}

func (bd *DataBlueprintDefinition) Configuration() interface{} {
	return bd.Args
}

func (bd *DataBlueprintDefinition) Attributes() dataBlueprintDefinitionAttributes {
	return dataBlueprintDefinitionAttributes{ref: terra.ReferenceDataResource(bd)}
}

type DataBlueprintDefinitionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ScopeId: string, required
	ScopeId terra.StringValue `hcl:"scope_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datablueprintdefinition.Timeouts `hcl:"timeouts,block"`
}
type dataBlueprintDefinitionAttributes struct {
	ref terra.Reference
}

func (bd dataBlueprintDefinitionAttributes) Description() terra.StringValue {
	return terra.ReferenceString(bd.ref.Append("description"))
}

func (bd dataBlueprintDefinitionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(bd.ref.Append("display_name"))
}

func (bd dataBlueprintDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(bd.ref.Append("id"))
}

func (bd dataBlueprintDefinitionAttributes) LastModified() terra.StringValue {
	return terra.ReferenceString(bd.ref.Append("last_modified"))
}

func (bd dataBlueprintDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(bd.ref.Append("name"))
}

func (bd dataBlueprintDefinitionAttributes) ScopeId() terra.StringValue {
	return terra.ReferenceString(bd.ref.Append("scope_id"))
}

func (bd dataBlueprintDefinitionAttributes) TargetScope() terra.StringValue {
	return terra.ReferenceString(bd.ref.Append("target_scope"))
}

func (bd dataBlueprintDefinitionAttributes) TimeCreated() terra.StringValue {
	return terra.ReferenceString(bd.ref.Append("time_created"))
}

func (bd dataBlueprintDefinitionAttributes) Versions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](bd.ref.Append("versions"))
}

func (bd dataBlueprintDefinitionAttributes) Timeouts() datablueprintdefinition.TimeoutsAttributes {
	return terra.ReferenceSingle[datablueprintdefinition.TimeoutsAttributes](bd.ref.Append("timeouts"))
}
