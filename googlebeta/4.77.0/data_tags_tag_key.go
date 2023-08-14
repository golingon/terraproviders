// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataTagsTagKey creates a new instance of [DataTagsTagKey].
func NewDataTagsTagKey(name string, args DataTagsTagKeyArgs) *DataTagsTagKey {
	return &DataTagsTagKey{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataTagsTagKey)(nil)

// DataTagsTagKey represents the Terraform data resource google_tags_tag_key.
type DataTagsTagKey struct {
	Name string
	Args DataTagsTagKeyArgs
}

// DataSource returns the Terraform object type for [DataTagsTagKey].
func (ttk *DataTagsTagKey) DataSource() string {
	return "google_tags_tag_key"
}

// LocalName returns the local name for [DataTagsTagKey].
func (ttk *DataTagsTagKey) LocalName() string {
	return ttk.Name
}

// Configuration returns the configuration (args) for [DataTagsTagKey].
func (ttk *DataTagsTagKey) Configuration() interface{} {
	return ttk.Args
}

// Attributes returns the attributes for [DataTagsTagKey].
func (ttk *DataTagsTagKey) Attributes() dataTagsTagKeyAttributes {
	return dataTagsTagKeyAttributes{ref: terra.ReferenceDataResource(ttk)}
}

// DataTagsTagKeyArgs contains the configurations for google_tags_tag_key.
type DataTagsTagKeyArgs struct {
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// ShortName: string, required
	ShortName terra.StringValue `hcl:"short_name,attr" validate:"required"`
}
type dataTagsTagKeyAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_tags_tag_key.
func (ttk dataTagsTagKeyAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("create_time"))
}

// Description returns a reference to field description of google_tags_tag_key.
func (ttk dataTagsTagKeyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("description"))
}

// Id returns a reference to field id of google_tags_tag_key.
func (ttk dataTagsTagKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("id"))
}

// Name returns a reference to field name of google_tags_tag_key.
func (ttk dataTagsTagKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("name"))
}

// NamespacedName returns a reference to field namespaced_name of google_tags_tag_key.
func (ttk dataTagsTagKeyAttributes) NamespacedName() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("namespaced_name"))
}

// Parent returns a reference to field parent of google_tags_tag_key.
func (ttk dataTagsTagKeyAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("parent"))
}

// ShortName returns a reference to field short_name of google_tags_tag_key.
func (ttk dataTagsTagKeyAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("short_name"))
}

// UpdateTime returns a reference to field update_time of google_tags_tag_key.
func (ttk dataTagsTagKeyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ttk.ref.Append("update_time"))
}