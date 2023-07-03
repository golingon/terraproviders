// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataresourcegroupstaggingapiresources "github.com/golingon/terraproviders/aws/5.6.2/dataresourcegroupstaggingapiresources"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataResourcegroupstaggingapiResources creates a new instance of [DataResourcegroupstaggingapiResources].
func NewDataResourcegroupstaggingapiResources(name string, args DataResourcegroupstaggingapiResourcesArgs) *DataResourcegroupstaggingapiResources {
	return &DataResourcegroupstaggingapiResources{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataResourcegroupstaggingapiResources)(nil)

// DataResourcegroupstaggingapiResources represents the Terraform data resource aws_resourcegroupstaggingapi_resources.
type DataResourcegroupstaggingapiResources struct {
	Name string
	Args DataResourcegroupstaggingapiResourcesArgs
}

// DataSource returns the Terraform object type for [DataResourcegroupstaggingapiResources].
func (rr *DataResourcegroupstaggingapiResources) DataSource() string {
	return "aws_resourcegroupstaggingapi_resources"
}

// LocalName returns the local name for [DataResourcegroupstaggingapiResources].
func (rr *DataResourcegroupstaggingapiResources) LocalName() string {
	return rr.Name
}

// Configuration returns the configuration (args) for [DataResourcegroupstaggingapiResources].
func (rr *DataResourcegroupstaggingapiResources) Configuration() interface{} {
	return rr.Args
}

// Attributes returns the attributes for [DataResourcegroupstaggingapiResources].
func (rr *DataResourcegroupstaggingapiResources) Attributes() dataResourcegroupstaggingapiResourcesAttributes {
	return dataResourcegroupstaggingapiResourcesAttributes{ref: terra.ReferenceDataResource(rr)}
}

// DataResourcegroupstaggingapiResourcesArgs contains the configurations for aws_resourcegroupstaggingapi_resources.
type DataResourcegroupstaggingapiResourcesArgs struct {
	// ExcludeCompliantResources: bool, optional
	ExcludeCompliantResources terra.BoolValue `hcl:"exclude_compliant_resources,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludeComplianceDetails: bool, optional
	IncludeComplianceDetails terra.BoolValue `hcl:"include_compliance_details,attr"`
	// ResourceArnList: set of string, optional
	ResourceArnList terra.SetValue[terra.StringValue] `hcl:"resource_arn_list,attr"`
	// ResourceTypeFilters: set of string, optional
	ResourceTypeFilters terra.SetValue[terra.StringValue] `hcl:"resource_type_filters,attr"`
	// ResourceTagMappingList: min=0
	ResourceTagMappingList []dataresourcegroupstaggingapiresources.ResourceTagMappingList `hcl:"resource_tag_mapping_list,block" validate:"min=0"`
	// TagFilter: min=0,max=50
	TagFilter []dataresourcegroupstaggingapiresources.TagFilter `hcl:"tag_filter,block" validate:"min=0,max=50"`
}
type dataResourcegroupstaggingapiResourcesAttributes struct {
	ref terra.Reference
}

// ExcludeCompliantResources returns a reference to field exclude_compliant_resources of aws_resourcegroupstaggingapi_resources.
func (rr dataResourcegroupstaggingapiResourcesAttributes) ExcludeCompliantResources() terra.BoolValue {
	return terra.ReferenceAsBool(rr.ref.Append("exclude_compliant_resources"))
}

// Id returns a reference to field id of aws_resourcegroupstaggingapi_resources.
func (rr dataResourcegroupstaggingapiResourcesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("id"))
}

// IncludeComplianceDetails returns a reference to field include_compliance_details of aws_resourcegroupstaggingapi_resources.
func (rr dataResourcegroupstaggingapiResourcesAttributes) IncludeComplianceDetails() terra.BoolValue {
	return terra.ReferenceAsBool(rr.ref.Append("include_compliance_details"))
}

// ResourceArnList returns a reference to field resource_arn_list of aws_resourcegroupstaggingapi_resources.
func (rr dataResourcegroupstaggingapiResourcesAttributes) ResourceArnList() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rr.ref.Append("resource_arn_list"))
}

// ResourceTypeFilters returns a reference to field resource_type_filters of aws_resourcegroupstaggingapi_resources.
func (rr dataResourcegroupstaggingapiResourcesAttributes) ResourceTypeFilters() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rr.ref.Append("resource_type_filters"))
}

func (rr dataResourcegroupstaggingapiResourcesAttributes) ResourceTagMappingList() terra.ListValue[dataresourcegroupstaggingapiresources.ResourceTagMappingListAttributes] {
	return terra.ReferenceAsList[dataresourcegroupstaggingapiresources.ResourceTagMappingListAttributes](rr.ref.Append("resource_tag_mapping_list"))
}

func (rr dataResourcegroupstaggingapiResourcesAttributes) TagFilter() terra.ListValue[dataresourcegroupstaggingapiresources.TagFilterAttributes] {
	return terra.ReferenceAsList[dataresourcegroupstaggingapiresources.TagFilterAttributes](rr.ref.Append("tag_filter"))
}
