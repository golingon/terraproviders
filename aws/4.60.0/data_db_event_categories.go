// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDbEventCategories creates a new instance of [DataDbEventCategories].
func NewDataDbEventCategories(name string, args DataDbEventCategoriesArgs) *DataDbEventCategories {
	return &DataDbEventCategories{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDbEventCategories)(nil)

// DataDbEventCategories represents the Terraform data resource aws_db_event_categories.
type DataDbEventCategories struct {
	Name string
	Args DataDbEventCategoriesArgs
}

// DataSource returns the Terraform object type for [DataDbEventCategories].
func (dec *DataDbEventCategories) DataSource() string {
	return "aws_db_event_categories"
}

// LocalName returns the local name for [DataDbEventCategories].
func (dec *DataDbEventCategories) LocalName() string {
	return dec.Name
}

// Configuration returns the configuration (args) for [DataDbEventCategories].
func (dec *DataDbEventCategories) Configuration() interface{} {
	return dec.Args
}

// Attributes returns the attributes for [DataDbEventCategories].
func (dec *DataDbEventCategories) Attributes() dataDbEventCategoriesAttributes {
	return dataDbEventCategoriesAttributes{ref: terra.ReferenceDataResource(dec)}
}

// DataDbEventCategoriesArgs contains the configurations for aws_db_event_categories.
type DataDbEventCategoriesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SourceType: string, optional
	SourceType terra.StringValue `hcl:"source_type,attr"`
}
type dataDbEventCategoriesAttributes struct {
	ref terra.Reference
}

// EventCategories returns a reference to field event_categories of aws_db_event_categories.
func (dec dataDbEventCategoriesAttributes) EventCategories() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dec.ref.Append("event_categories"))
}

// Id returns a reference to field id of aws_db_event_categories.
func (dec dataDbEventCategoriesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dec.ref.Append("id"))
}

// SourceType returns a reference to field source_type of aws_db_event_categories.
func (dec dataDbEventCategoriesAttributes) SourceType() terra.StringValue {
	return terra.ReferenceAsString(dec.ref.Append("source_type"))
}
