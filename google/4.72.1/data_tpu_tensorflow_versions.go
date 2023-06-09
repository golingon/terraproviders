// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataTpuTensorflowVersions creates a new instance of [DataTpuTensorflowVersions].
func NewDataTpuTensorflowVersions(name string, args DataTpuTensorflowVersionsArgs) *DataTpuTensorflowVersions {
	return &DataTpuTensorflowVersions{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataTpuTensorflowVersions)(nil)

// DataTpuTensorflowVersions represents the Terraform data resource google_tpu_tensorflow_versions.
type DataTpuTensorflowVersions struct {
	Name string
	Args DataTpuTensorflowVersionsArgs
}

// DataSource returns the Terraform object type for [DataTpuTensorflowVersions].
func (ttv *DataTpuTensorflowVersions) DataSource() string {
	return "google_tpu_tensorflow_versions"
}

// LocalName returns the local name for [DataTpuTensorflowVersions].
func (ttv *DataTpuTensorflowVersions) LocalName() string {
	return ttv.Name
}

// Configuration returns the configuration (args) for [DataTpuTensorflowVersions].
func (ttv *DataTpuTensorflowVersions) Configuration() interface{} {
	return ttv.Args
}

// Attributes returns the attributes for [DataTpuTensorflowVersions].
func (ttv *DataTpuTensorflowVersions) Attributes() dataTpuTensorflowVersionsAttributes {
	return dataTpuTensorflowVersionsAttributes{ref: terra.ReferenceDataResource(ttv)}
}

// DataTpuTensorflowVersionsArgs contains the configurations for google_tpu_tensorflow_versions.
type DataTpuTensorflowVersionsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
}
type dataTpuTensorflowVersionsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_tpu_tensorflow_versions.
func (ttv dataTpuTensorflowVersionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("id"))
}

// Project returns a reference to field project of google_tpu_tensorflow_versions.
func (ttv dataTpuTensorflowVersionsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("project"))
}

// Versions returns a reference to field versions of google_tpu_tensorflow_versions.
func (ttv dataTpuTensorflowVersionsAttributes) Versions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ttv.ref.Append("versions"))
}

// Zone returns a reference to field zone of google_tpu_tensorflow_versions.
func (ttv dataTpuTensorflowVersionsAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(ttv.ref.Append("zone"))
}
