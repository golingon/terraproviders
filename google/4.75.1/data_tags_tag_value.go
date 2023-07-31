// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataTagsTagValue creates a new instance of [DataTagsTagValue].
func NewDataTagsTagValue(name string, args DataTagsTagValueArgs) *DataTagsTagValue {
	return &DataTagsTagValue{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataTagsTagValue)(nil)

// DataTagsTagValue represents the Terraform data resource google_tags_tag_value.
type DataTagsTagValue struct {
	Name string
	Args DataTagsTagValueArgs
}

// DataSource returns the Terraform object type for [DataTagsTagValue].
func (ttv *DataTagsTagValue) DataSource() string {
	return "google_tags_tag_value"
}

// LocalName returns the local name for [DataTagsTagValue].
func (ttv *DataTagsTagValue) LocalName() string {
	return ttv.Name
}

// Configuration returns the configuration (args) for [DataTagsTagValue].
func (ttv *DataTagsTagValue) Configuration() interface{} {
	return ttv.Args
}

// Attributes returns the attributes for [DataTagsTagValue].
func (ttv *DataTagsTagValue) Attributes() dataTagsTagValueAttributes {
	return dataTagsTagValueAttributes{ref: terra.ReferenceDataResource(ttv)}
}

// DataTagsTagValueArgs contains the configurations for google_tags_tag_value.
type DataTagsTagValueArgs struct {
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// ShortName: string, required
	ShortName terra.StringValue `hcl:"short_name,attr" validate:"required"`
}
type dataTagsTagValueAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_tags_tag_value.
func (ttv dataTagsTagValueAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("create_time"))
}

// Description returns a reference to field description of google_tags_tag_value.
func (ttv dataTagsTagValueAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("description"))
}

// Id returns a reference to field id of google_tags_tag_value.
func (ttv dataTagsTagValueAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("id"))
}

// Name returns a reference to field name of google_tags_tag_value.
func (ttv dataTagsTagValueAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("name"))
}

// NamespacedName returns a reference to field namespaced_name of google_tags_tag_value.
func (ttv dataTagsTagValueAttributes) NamespacedName() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("namespaced_name"))
}

// Parent returns a reference to field parent of google_tags_tag_value.
func (ttv dataTagsTagValueAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("parent"))
}

// ShortName returns a reference to field short_name of google_tags_tag_value.
func (ttv dataTagsTagValueAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("short_name"))
}

// UpdateTime returns a reference to field update_time of google_tags_tag_value.
func (ttv dataTagsTagValueAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("update_time"))
}
