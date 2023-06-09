// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datatemplatespecversion "github.com/golingon/terraproviders/azurerm/3.52.0/datatemplatespecversion"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataTemplateSpecVersion creates a new instance of [DataTemplateSpecVersion].
func NewDataTemplateSpecVersion(name string, args DataTemplateSpecVersionArgs) *DataTemplateSpecVersion {
	return &DataTemplateSpecVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataTemplateSpecVersion)(nil)

// DataTemplateSpecVersion represents the Terraform data resource azurerm_template_spec_version.
type DataTemplateSpecVersion struct {
	Name string
	Args DataTemplateSpecVersionArgs
}

// DataSource returns the Terraform object type for [DataTemplateSpecVersion].
func (tsv *DataTemplateSpecVersion) DataSource() string {
	return "azurerm_template_spec_version"
}

// LocalName returns the local name for [DataTemplateSpecVersion].
func (tsv *DataTemplateSpecVersion) LocalName() string {
	return tsv.Name
}

// Configuration returns the configuration (args) for [DataTemplateSpecVersion].
func (tsv *DataTemplateSpecVersion) Configuration() interface{} {
	return tsv.Args
}

// Attributes returns the attributes for [DataTemplateSpecVersion].
func (tsv *DataTemplateSpecVersion) Attributes() dataTemplateSpecVersionAttributes {
	return dataTemplateSpecVersionAttributes{ref: terra.ReferenceDataResource(tsv)}
}

// DataTemplateSpecVersionArgs contains the configurations for azurerm_template_spec_version.
type DataTemplateSpecVersionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datatemplatespecversion.Timeouts `hcl:"timeouts,block"`
}
type dataTemplateSpecVersionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_template_spec_version.
func (tsv dataTemplateSpecVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tsv.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_template_spec_version.
func (tsv dataTemplateSpecVersionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tsv.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_template_spec_version.
func (tsv dataTemplateSpecVersionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(tsv.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_template_spec_version.
func (tsv dataTemplateSpecVersionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tsv.ref.Append("tags"))
}

// TemplateBody returns a reference to field template_body of azurerm_template_spec_version.
func (tsv dataTemplateSpecVersionAttributes) TemplateBody() terra.StringValue {
	return terra.ReferenceAsString(tsv.ref.Append("template_body"))
}

// Version returns a reference to field version of azurerm_template_spec_version.
func (tsv dataTemplateSpecVersionAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(tsv.ref.Append("version"))
}

func (tsv dataTemplateSpecVersionAttributes) Timeouts() datatemplatespecversion.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datatemplatespecversion.TimeoutsAttributes](tsv.ref.Append("timeouts"))
}
