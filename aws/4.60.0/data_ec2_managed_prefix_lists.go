// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2managedprefixlists "github.com/golingon/terraproviders/aws/4.60.0/dataec2managedprefixlists"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2ManagedPrefixLists creates a new instance of [DataEc2ManagedPrefixLists].
func NewDataEc2ManagedPrefixLists(name string, args DataEc2ManagedPrefixListsArgs) *DataEc2ManagedPrefixLists {
	return &DataEc2ManagedPrefixLists{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2ManagedPrefixLists)(nil)

// DataEc2ManagedPrefixLists represents the Terraform data resource aws_ec2_managed_prefix_lists.
type DataEc2ManagedPrefixLists struct {
	Name string
	Args DataEc2ManagedPrefixListsArgs
}

// DataSource returns the Terraform object type for [DataEc2ManagedPrefixLists].
func (empl *DataEc2ManagedPrefixLists) DataSource() string {
	return "aws_ec2_managed_prefix_lists"
}

// LocalName returns the local name for [DataEc2ManagedPrefixLists].
func (empl *DataEc2ManagedPrefixLists) LocalName() string {
	return empl.Name
}

// Configuration returns the configuration (args) for [DataEc2ManagedPrefixLists].
func (empl *DataEc2ManagedPrefixLists) Configuration() interface{} {
	return empl.Args
}

// Attributes returns the attributes for [DataEc2ManagedPrefixLists].
func (empl *DataEc2ManagedPrefixLists) Attributes() dataEc2ManagedPrefixListsAttributes {
	return dataEc2ManagedPrefixListsAttributes{ref: terra.ReferenceDataResource(empl)}
}

// DataEc2ManagedPrefixListsArgs contains the configurations for aws_ec2_managed_prefix_lists.
type DataEc2ManagedPrefixListsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2managedprefixlists.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataEc2ManagedPrefixListsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_managed_prefix_lists.
func (empl dataEc2ManagedPrefixListsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(empl.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_ec2_managed_prefix_lists.
func (empl dataEc2ManagedPrefixListsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](empl.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_ec2_managed_prefix_lists.
func (empl dataEc2ManagedPrefixListsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](empl.ref.Append("tags"))
}

func (empl dataEc2ManagedPrefixListsAttributes) Filter() terra.SetValue[dataec2managedprefixlists.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2managedprefixlists.FilterAttributes](empl.ref.Append("filter"))
}
