// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataComputeInstanceSerialPort creates a new instance of [DataComputeInstanceSerialPort].
func NewDataComputeInstanceSerialPort(name string, args DataComputeInstanceSerialPortArgs) *DataComputeInstanceSerialPort {
	return &DataComputeInstanceSerialPort{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeInstanceSerialPort)(nil)

// DataComputeInstanceSerialPort represents the Terraform data resource google_compute_instance_serial_port.
type DataComputeInstanceSerialPort struct {
	Name string
	Args DataComputeInstanceSerialPortArgs
}

// DataSource returns the Terraform object type for [DataComputeInstanceSerialPort].
func (cisp *DataComputeInstanceSerialPort) DataSource() string {
	return "google_compute_instance_serial_port"
}

// LocalName returns the local name for [DataComputeInstanceSerialPort].
func (cisp *DataComputeInstanceSerialPort) LocalName() string {
	return cisp.Name
}

// Configuration returns the configuration (args) for [DataComputeInstanceSerialPort].
func (cisp *DataComputeInstanceSerialPort) Configuration() interface{} {
	return cisp.Args
}

// Attributes returns the attributes for [DataComputeInstanceSerialPort].
func (cisp *DataComputeInstanceSerialPort) Attributes() dataComputeInstanceSerialPortAttributes {
	return dataComputeInstanceSerialPortAttributes{ref: terra.ReferenceDataResource(cisp)}
}

// DataComputeInstanceSerialPortArgs contains the configurations for google_compute_instance_serial_port.
type DataComputeInstanceSerialPortArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
}
type dataComputeInstanceSerialPortAttributes struct {
	ref terra.Reference
}

// Contents returns a reference to field contents of google_compute_instance_serial_port.
func (cisp dataComputeInstanceSerialPortAttributes) Contents() terra.StringValue {
	return terra.ReferenceAsString(cisp.ref.Append("contents"))
}

// Id returns a reference to field id of google_compute_instance_serial_port.
func (cisp dataComputeInstanceSerialPortAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cisp.ref.Append("id"))
}

// Instance returns a reference to field instance of google_compute_instance_serial_port.
func (cisp dataComputeInstanceSerialPortAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(cisp.ref.Append("instance"))
}

// Port returns a reference to field port of google_compute_instance_serial_port.
func (cisp dataComputeInstanceSerialPortAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(cisp.ref.Append("port"))
}

// Project returns a reference to field project of google_compute_instance_serial_port.
func (cisp dataComputeInstanceSerialPortAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cisp.ref.Append("project"))
}

// Zone returns a reference to field zone of google_compute_instance_serial_port.
func (cisp dataComputeInstanceSerialPortAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cisp.ref.Append("zone"))
}
