// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datablueprintdefinition "github.com/golingon/terraproviders/azurerm/3.65.0/datablueprintdefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBlueprintDefinition creates a new instance of [DataBlueprintDefinition].
func NewDataBlueprintDefinition(name string, args DataBlueprintDefinitionArgs) *DataBlueprintDefinition {
	return &DataBlueprintDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBlueprintDefinition)(nil)

// DataBlueprintDefinition represents the Terraform data resource azurerm_blueprint_definition.
type DataBlueprintDefinition struct {
	Name string
	Args DataBlueprintDefinitionArgs
}

// DataSource returns the Terraform object type for [DataBlueprintDefinition].
func (bd *DataBlueprintDefinition) DataSource() string {
	return "azurerm_blueprint_definition"
}

// LocalName returns the local name for [DataBlueprintDefinition].
func (bd *DataBlueprintDefinition) LocalName() string {
	return bd.Name
}

// Configuration returns the configuration (args) for [DataBlueprintDefinition].
func (bd *DataBlueprintDefinition) Configuration() interface{} {
	return bd.Args
}

// Attributes returns the attributes for [DataBlueprintDefinition].
func (bd *DataBlueprintDefinition) Attributes() dataBlueprintDefinitionAttributes {
	return dataBlueprintDefinitionAttributes{ref: terra.ReferenceDataResource(bd)}
}

// DataBlueprintDefinitionArgs contains the configurations for azurerm_blueprint_definition.
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

// Description returns a reference to field description of azurerm_blueprint_definition.
func (bd dataBlueprintDefinitionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_blueprint_definition.
func (bd dataBlueprintDefinitionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_blueprint_definition.
func (bd dataBlueprintDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("id"))
}

// LastModified returns a reference to field last_modified of azurerm_blueprint_definition.
func (bd dataBlueprintDefinitionAttributes) LastModified() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("last_modified"))
}

// Name returns a reference to field name of azurerm_blueprint_definition.
func (bd dataBlueprintDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("name"))
}

// ScopeId returns a reference to field scope_id of azurerm_blueprint_definition.
func (bd dataBlueprintDefinitionAttributes) ScopeId() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("scope_id"))
}

// TargetScope returns a reference to field target_scope of azurerm_blueprint_definition.
func (bd dataBlueprintDefinitionAttributes) TargetScope() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("target_scope"))
}

// TimeCreated returns a reference to field time_created of azurerm_blueprint_definition.
func (bd dataBlueprintDefinitionAttributes) TimeCreated() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("time_created"))
}

// Versions returns a reference to field versions of azurerm_blueprint_definition.
func (bd dataBlueprintDefinitionAttributes) Versions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bd.ref.Append("versions"))
}

func (bd dataBlueprintDefinitionAttributes) Timeouts() datablueprintdefinition.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datablueprintdefinition.TimeoutsAttributes](bd.ref.Append("timeouts"))
}
