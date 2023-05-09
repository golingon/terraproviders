// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataamiids "github.com/golingon/terraproviders/aws/4.66.1/dataamiids"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAmiIds creates a new instance of [DataAmiIds].
func NewDataAmiIds(name string, args DataAmiIdsArgs) *DataAmiIds {
	return &DataAmiIds{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAmiIds)(nil)

// DataAmiIds represents the Terraform data resource aws_ami_ids.
type DataAmiIds struct {
	Name string
	Args DataAmiIdsArgs
}

// DataSource returns the Terraform object type for [DataAmiIds].
func (ai *DataAmiIds) DataSource() string {
	return "aws_ami_ids"
}

// LocalName returns the local name for [DataAmiIds].
func (ai *DataAmiIds) LocalName() string {
	return ai.Name
}

// Configuration returns the configuration (args) for [DataAmiIds].
func (ai *DataAmiIds) Configuration() interface{} {
	return ai.Args
}

// Attributes returns the attributes for [DataAmiIds].
func (ai *DataAmiIds) Attributes() dataAmiIdsAttributes {
	return dataAmiIdsAttributes{ref: terra.ReferenceDataResource(ai)}
}

// DataAmiIdsArgs contains the configurations for aws_ami_ids.
type DataAmiIdsArgs struct {
	// ExecutableUsers: list of string, optional
	ExecutableUsers terra.ListValue[terra.StringValue] `hcl:"executable_users,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludeDeprecated: bool, optional
	IncludeDeprecated terra.BoolValue `hcl:"include_deprecated,attr"`
	// NameRegex: string, optional
	NameRegex terra.StringValue `hcl:"name_regex,attr"`
	// Owners: list of string, required
	Owners terra.ListValue[terra.StringValue] `hcl:"owners,attr" validate:"required"`
	// SortAscending: bool, optional
	SortAscending terra.BoolValue `hcl:"sort_ascending,attr"`
	// Filter: min=0
	Filter []dataamiids.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataamiids.Timeouts `hcl:"timeouts,block"`
}
type dataAmiIdsAttributes struct {
	ref terra.Reference
}

// ExecutableUsers returns a reference to field executable_users of aws_ami_ids.
func (ai dataAmiIdsAttributes) ExecutableUsers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ai.ref.Append("executable_users"))
}

// Id returns a reference to field id of aws_ami_ids.
func (ai dataAmiIdsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_ami_ids.
func (ai dataAmiIdsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ai.ref.Append("ids"))
}

// IncludeDeprecated returns a reference to field include_deprecated of aws_ami_ids.
func (ai dataAmiIdsAttributes) IncludeDeprecated() terra.BoolValue {
	return terra.ReferenceAsBool(ai.ref.Append("include_deprecated"))
}

// NameRegex returns a reference to field name_regex of aws_ami_ids.
func (ai dataAmiIdsAttributes) NameRegex() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("name_regex"))
}

// Owners returns a reference to field owners of aws_ami_ids.
func (ai dataAmiIdsAttributes) Owners() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ai.ref.Append("owners"))
}

// SortAscending returns a reference to field sort_ascending of aws_ami_ids.
func (ai dataAmiIdsAttributes) SortAscending() terra.BoolValue {
	return terra.ReferenceAsBool(ai.ref.Append("sort_ascending"))
}

func (ai dataAmiIdsAttributes) Filter() terra.SetValue[dataamiids.FilterAttributes] {
	return terra.ReferenceAsSet[dataamiids.FilterAttributes](ai.ref.Append("filter"))
}

func (ai dataAmiIdsAttributes) Timeouts() dataamiids.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataamiids.TimeoutsAttributes](ai.ref.Append("timeouts"))
}
