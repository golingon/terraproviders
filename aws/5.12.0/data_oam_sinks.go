// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataOamSinks creates a new instance of [DataOamSinks].
func NewDataOamSinks(name string, args DataOamSinksArgs) *DataOamSinks {
	return &DataOamSinks{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOamSinks)(nil)

// DataOamSinks represents the Terraform data resource aws_oam_sinks.
type DataOamSinks struct {
	Name string
	Args DataOamSinksArgs
}

// DataSource returns the Terraform object type for [DataOamSinks].
func (os *DataOamSinks) DataSource() string {
	return "aws_oam_sinks"
}

// LocalName returns the local name for [DataOamSinks].
func (os *DataOamSinks) LocalName() string {
	return os.Name
}

// Configuration returns the configuration (args) for [DataOamSinks].
func (os *DataOamSinks) Configuration() interface{} {
	return os.Args
}

// Attributes returns the attributes for [DataOamSinks].
func (os *DataOamSinks) Attributes() dataOamSinksAttributes {
	return dataOamSinksAttributes{ref: terra.ReferenceDataResource(os)}
}

// DataOamSinksArgs contains the configurations for aws_oam_sinks.
type DataOamSinksArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataOamSinksAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_oam_sinks.
func (os dataOamSinksAttributes) Arns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](os.ref.Append("arns"))
}

// Id returns a reference to field id of aws_oam_sinks.
func (os dataOamSinksAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("id"))
}
