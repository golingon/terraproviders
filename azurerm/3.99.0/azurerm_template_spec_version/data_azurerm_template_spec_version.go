// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_template_spec_version

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_template_spec_version.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (atsv *DataSource) DataSource() string {
	return "azurerm_template_spec_version"
}

// LocalName returns the local name for [DataSource].
func (atsv *DataSource) LocalName() string {
	return atsv.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (atsv *DataSource) Configuration() interface{} {
	return atsv.Args
}

// Attributes returns the attributes for [DataSource].
func (atsv *DataSource) Attributes() dataAzurermTemplateSpecVersionAttributes {
	return dataAzurermTemplateSpecVersionAttributes{ref: terra.ReferenceDataSource(atsv)}
}

// DataArgs contains the configurations for azurerm_template_spec_version.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermTemplateSpecVersionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_template_spec_version.
func (atsv dataAzurermTemplateSpecVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(atsv.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_template_spec_version.
func (atsv dataAzurermTemplateSpecVersionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(atsv.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_template_spec_version.
func (atsv dataAzurermTemplateSpecVersionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(atsv.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_template_spec_version.
func (atsv dataAzurermTemplateSpecVersionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](atsv.ref.Append("tags"))
}

// TemplateBody returns a reference to field template_body of azurerm_template_spec_version.
func (atsv dataAzurermTemplateSpecVersionAttributes) TemplateBody() terra.StringValue {
	return terra.ReferenceAsString(atsv.ref.Append("template_body"))
}

// Version returns a reference to field version of azurerm_template_spec_version.
func (atsv dataAzurermTemplateSpecVersionAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(atsv.ref.Append("version"))
}

func (atsv dataAzurermTemplateSpecVersionAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](atsv.ref.Append("timeouts"))
}
