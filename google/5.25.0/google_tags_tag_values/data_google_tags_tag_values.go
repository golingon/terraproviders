// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_tags_tag_values

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_tags_tag_values.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gttv *DataSource) DataSource() string {
	return "google_tags_tag_values"
}

// LocalName returns the local name for [DataSource].
func (gttv *DataSource) LocalName() string {
	return gttv.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gttv *DataSource) Configuration() interface{} {
	return gttv.Args
}

// Attributes returns the attributes for [DataSource].
func (gttv *DataSource) Attributes() dataGoogleTagsTagValuesAttributes {
	return dataGoogleTagsTagValuesAttributes{ref: terra.ReferenceDataSource(gttv)}
}

// DataArgs contains the configurations for google_tags_tag_values.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
}

type dataGoogleTagsTagValuesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_tags_tag_values.
func (gttv dataGoogleTagsTagValuesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gttv.ref.Append("id"))
}

// Parent returns a reference to field parent of google_tags_tag_values.
func (gttv dataGoogleTagsTagValuesAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(gttv.ref.Append("parent"))
}

func (gttv dataGoogleTagsTagValuesAttributes) Values() terra.ListValue[DataValuesAttributes] {
	return terra.ReferenceAsList[DataValuesAttributes](gttv.ref.Append("values"))
}
