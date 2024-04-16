// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_resourceexplorer2_search

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_resourceexplorer2_search.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (ars *DataSource) DataSource() string {
	return "aws_resourceexplorer2_search"
}

// LocalName returns the local name for [DataSource].
func (ars *DataSource) LocalName() string {
	return ars.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (ars *DataSource) Configuration() interface{} {
	return ars.Args
}

// Attributes returns the attributes for [DataSource].
func (ars *DataSource) Attributes() dataAwsResourceexplorer2SearchAttributes {
	return dataAwsResourceexplorer2SearchAttributes{ref: terra.ReferenceDataSource(ars)}
}

// DataArgs contains the configurations for aws_resourceexplorer2_search.
type DataArgs struct {
	// QueryString: string, required
	QueryString terra.StringValue `hcl:"query_string,attr" validate:"required"`
	// ViewArn: string, optional
	ViewArn terra.StringValue `hcl:"view_arn,attr"`
	// ResourceCount: min=0
	ResourceCount []DataResourceCount `hcl:"resource_count,block" validate:"min=0"`
	// Resources: min=0
	Resources []DataResources `hcl:"resources,block" validate:"min=0"`
}

type dataAwsResourceexplorer2SearchAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_resourceexplorer2_search.
func (ars dataAwsResourceexplorer2SearchAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ars.ref.Append("id"))
}

// QueryString returns a reference to field query_string of aws_resourceexplorer2_search.
func (ars dataAwsResourceexplorer2SearchAttributes) QueryString() terra.StringValue {
	return terra.ReferenceAsString(ars.ref.Append("query_string"))
}

// ViewArn returns a reference to field view_arn of aws_resourceexplorer2_search.
func (ars dataAwsResourceexplorer2SearchAttributes) ViewArn() terra.StringValue {
	return terra.ReferenceAsString(ars.ref.Append("view_arn"))
}

func (ars dataAwsResourceexplorer2SearchAttributes) ResourceCount() terra.ListValue[DataResourceCountAttributes] {
	return terra.ReferenceAsList[DataResourceCountAttributes](ars.ref.Append("resource_count"))
}

func (ars dataAwsResourceexplorer2SearchAttributes) Resources() terra.ListValue[DataResourcesAttributes] {
	return terra.ReferenceAsList[DataResourcesAttributes](ars.ref.Append("resources"))
}
