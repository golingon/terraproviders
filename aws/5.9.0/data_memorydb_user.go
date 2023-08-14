// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datamemorydbuser "github.com/golingon/terraproviders/aws/5.9.0/datamemorydbuser"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMemorydbUser creates a new instance of [DataMemorydbUser].
func NewDataMemorydbUser(name string, args DataMemorydbUserArgs) *DataMemorydbUser {
	return &DataMemorydbUser{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMemorydbUser)(nil)

// DataMemorydbUser represents the Terraform data resource aws_memorydb_user.
type DataMemorydbUser struct {
	Name string
	Args DataMemorydbUserArgs
}

// DataSource returns the Terraform object type for [DataMemorydbUser].
func (mu *DataMemorydbUser) DataSource() string {
	return "aws_memorydb_user"
}

// LocalName returns the local name for [DataMemorydbUser].
func (mu *DataMemorydbUser) LocalName() string {
	return mu.Name
}

// Configuration returns the configuration (args) for [DataMemorydbUser].
func (mu *DataMemorydbUser) Configuration() interface{} {
	return mu.Args
}

// Attributes returns the attributes for [DataMemorydbUser].
func (mu *DataMemorydbUser) Attributes() dataMemorydbUserAttributes {
	return dataMemorydbUserAttributes{ref: terra.ReferenceDataResource(mu)}
}

// DataMemorydbUserArgs contains the configurations for aws_memorydb_user.
type DataMemorydbUserArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// UserName: string, required
	UserName terra.StringValue `hcl:"user_name,attr" validate:"required"`
	// AuthenticationMode: min=0
	AuthenticationMode []datamemorydbuser.AuthenticationMode `hcl:"authentication_mode,block" validate:"min=0"`
}
type dataMemorydbUserAttributes struct {
	ref terra.Reference
}

// AccessString returns a reference to field access_string of aws_memorydb_user.
func (mu dataMemorydbUserAttributes) AccessString() terra.StringValue {
	return terra.ReferenceAsString(mu.ref.Append("access_string"))
}

// Arn returns a reference to field arn of aws_memorydb_user.
func (mu dataMemorydbUserAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mu.ref.Append("arn"))
}

// Id returns a reference to field id of aws_memorydb_user.
func (mu dataMemorydbUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mu.ref.Append("id"))
}

// MinimumEngineVersion returns a reference to field minimum_engine_version of aws_memorydb_user.
func (mu dataMemorydbUserAttributes) MinimumEngineVersion() terra.StringValue {
	return terra.ReferenceAsString(mu.ref.Append("minimum_engine_version"))
}

// Tags returns a reference to field tags of aws_memorydb_user.
func (mu dataMemorydbUserAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mu.ref.Append("tags"))
}

// UserName returns a reference to field user_name of aws_memorydb_user.
func (mu dataMemorydbUserAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(mu.ref.Append("user_name"))
}

func (mu dataMemorydbUserAttributes) AuthenticationMode() terra.ListValue[datamemorydbuser.AuthenticationModeAttributes] {
	return terra.ReferenceAsList[datamemorydbuser.AuthenticationModeAttributes](mu.ref.Append("authentication_mode"))
}