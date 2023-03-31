// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataTagsTagKey(name string, args DataTagsTagKeyArgs) *DataTagsTagKey {
	return &DataTagsTagKey{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataTagsTagKey)(nil)

type DataTagsTagKey struct {
	Name string
	Args DataTagsTagKeyArgs
}

func (ttk *DataTagsTagKey) DataSource() string {
	return "google_tags_tag_key"
}

func (ttk *DataTagsTagKey) LocalName() string {
	return ttk.Name
}

func (ttk *DataTagsTagKey) Configuration() interface{} {
	return ttk.Args
}

func (ttk *DataTagsTagKey) Attributes() dataTagsTagKeyAttributes {
	return dataTagsTagKeyAttributes{ref: terra.ReferenceDataResource(ttk)}
}

type DataTagsTagKeyArgs struct {
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// ShortName: string, required
	ShortName terra.StringValue `hcl:"short_name,attr" validate:"required"`
}
type dataTagsTagKeyAttributes struct {
	ref terra.Reference
}

func (ttk dataTagsTagKeyAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceString(ttk.ref.Append("create_time"))
}

func (ttk dataTagsTagKeyAttributes) Description() terra.StringValue {
	return terra.ReferenceString(ttk.ref.Append("description"))
}

func (ttk dataTagsTagKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ttk.ref.Append("id"))
}

func (ttk dataTagsTagKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ttk.ref.Append("name"))
}

func (ttk dataTagsTagKeyAttributes) NamespacedName() terra.StringValue {
	return terra.ReferenceString(ttk.ref.Append("namespaced_name"))
}

func (ttk dataTagsTagKeyAttributes) Parent() terra.StringValue {
	return terra.ReferenceString(ttk.ref.Append("parent"))
}

func (ttk dataTagsTagKeyAttributes) ShortName() terra.StringValue {
	return terra.ReferenceString(ttk.ref.Append("short_name"))
}

func (ttk dataTagsTagKeyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceString(ttk.ref.Append("update_time"))
}
