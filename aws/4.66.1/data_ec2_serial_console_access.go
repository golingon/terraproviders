// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2serialconsoleaccess "github.com/golingon/terraproviders/aws/4.66.1/dataec2serialconsoleaccess"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2SerialConsoleAccess creates a new instance of [DataEc2SerialConsoleAccess].
func NewDataEc2SerialConsoleAccess(name string, args DataEc2SerialConsoleAccessArgs) *DataEc2SerialConsoleAccess {
	return &DataEc2SerialConsoleAccess{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2SerialConsoleAccess)(nil)

// DataEc2SerialConsoleAccess represents the Terraform data resource aws_ec2_serial_console_access.
type DataEc2SerialConsoleAccess struct {
	Name string
	Args DataEc2SerialConsoleAccessArgs
}

// DataSource returns the Terraform object type for [DataEc2SerialConsoleAccess].
func (esca *DataEc2SerialConsoleAccess) DataSource() string {
	return "aws_ec2_serial_console_access"
}

// LocalName returns the local name for [DataEc2SerialConsoleAccess].
func (esca *DataEc2SerialConsoleAccess) LocalName() string {
	return esca.Name
}

// Configuration returns the configuration (args) for [DataEc2SerialConsoleAccess].
func (esca *DataEc2SerialConsoleAccess) Configuration() interface{} {
	return esca.Args
}

// Attributes returns the attributes for [DataEc2SerialConsoleAccess].
func (esca *DataEc2SerialConsoleAccess) Attributes() dataEc2SerialConsoleAccessAttributes {
	return dataEc2SerialConsoleAccessAttributes{ref: terra.ReferenceDataResource(esca)}
}

// DataEc2SerialConsoleAccessArgs contains the configurations for aws_ec2_serial_console_access.
type DataEc2SerialConsoleAccessArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *dataec2serialconsoleaccess.Timeouts `hcl:"timeouts,block"`
}
type dataEc2SerialConsoleAccessAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of aws_ec2_serial_console_access.
func (esca dataEc2SerialConsoleAccessAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(esca.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_ec2_serial_console_access.
func (esca dataEc2SerialConsoleAccessAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(esca.ref.Append("id"))
}

func (esca dataEc2SerialConsoleAccessAttributes) Timeouts() dataec2serialconsoleaccess.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2serialconsoleaccess.TimeoutsAttributes](esca.ref.Append("timeouts"))
}
