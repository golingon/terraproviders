// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataImagebuilderComponent creates a new instance of [DataImagebuilderComponent].
func NewDataImagebuilderComponent(name string, args DataImagebuilderComponentArgs) *DataImagebuilderComponent {
	return &DataImagebuilderComponent{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataImagebuilderComponent)(nil)

// DataImagebuilderComponent represents the Terraform data resource aws_imagebuilder_component.
type DataImagebuilderComponent struct {
	Name string
	Args DataImagebuilderComponentArgs
}

// DataSource returns the Terraform object type for [DataImagebuilderComponent].
func (ic *DataImagebuilderComponent) DataSource() string {
	return "aws_imagebuilder_component"
}

// LocalName returns the local name for [DataImagebuilderComponent].
func (ic *DataImagebuilderComponent) LocalName() string {
	return ic.Name
}

// Configuration returns the configuration (args) for [DataImagebuilderComponent].
func (ic *DataImagebuilderComponent) Configuration() interface{} {
	return ic.Args
}

// Attributes returns the attributes for [DataImagebuilderComponent].
func (ic *DataImagebuilderComponent) Attributes() dataImagebuilderComponentAttributes {
	return dataImagebuilderComponentAttributes{ref: terra.ReferenceDataResource(ic)}
}

// DataImagebuilderComponentArgs contains the configurations for aws_imagebuilder_component.
type DataImagebuilderComponentArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataImagebuilderComponentAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("arn"))
}

// ChangeDescription returns a reference to field change_description of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) ChangeDescription() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("change_description"))
}

// Data returns a reference to field data of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) Data() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("data"))
}

// DateCreated returns a reference to field date_created of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) DateCreated() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("date_created"))
}

// Description returns a reference to field description of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("description"))
}

// Encrypted returns a reference to field encrypted of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("encrypted"))
}

// Id returns a reference to field id of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("kms_key_id"))
}

// Name returns a reference to field name of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("owner"))
}

// Platform returns a reference to field platform of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) Platform() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("platform"))
}

// SupportedOsVersions returns a reference to field supported_os_versions of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) SupportedOsVersions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ic.ref.Append("supported_os_versions"))
}

// Tags returns a reference to field tags of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ic.ref.Append("tags"))
}

// Type returns a reference to field type of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("type"))
}

// Version returns a reference to field version of aws_imagebuilder_component.
func (ic dataImagebuilderComponentAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("version"))
}
