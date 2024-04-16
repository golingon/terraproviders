// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_machine_types

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_compute_machine_types.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gcmt *DataSource) DataSource() string {
	return "google_compute_machine_types"
}

// LocalName returns the local name for [DataSource].
func (gcmt *DataSource) LocalName() string {
	return gcmt.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gcmt *DataSource) Configuration() interface{} {
	return gcmt.Args
}

// Attributes returns the attributes for [DataSource].
func (gcmt *DataSource) Attributes() dataGoogleComputeMachineTypesAttributes {
	return dataGoogleComputeMachineTypesAttributes{ref: terra.ReferenceDataSource(gcmt)}
}

// DataArgs contains the configurations for google_compute_machine_types.
type DataArgs struct {
	// Filter: string, optional
	Filter terra.StringValue `hcl:"filter,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
}

type dataGoogleComputeMachineTypesAttributes struct {
	ref terra.Reference
}

// Filter returns a reference to field filter of google_compute_machine_types.
func (gcmt dataGoogleComputeMachineTypesAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(gcmt.ref.Append("filter"))
}

// Id returns a reference to field id of google_compute_machine_types.
func (gcmt dataGoogleComputeMachineTypesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcmt.ref.Append("id"))
}

// Project returns a reference to field project of google_compute_machine_types.
func (gcmt dataGoogleComputeMachineTypesAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcmt.ref.Append("project"))
}

// Zone returns a reference to field zone of google_compute_machine_types.
func (gcmt dataGoogleComputeMachineTypesAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(gcmt.ref.Append("zone"))
}

func (gcmt dataGoogleComputeMachineTypesAttributes) MachineTypes() terra.ListValue[DataMachineTypesAttributes] {
	return terra.ReferenceAsList[DataMachineTypesAttributes](gcmt.ref.Append("machine_types"))
}
