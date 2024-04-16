// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_vmwareengine_nsx_credentials

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_vmwareengine_nsx_credentials.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gvnc *DataSource) DataSource() string {
	return "google_vmwareengine_nsx_credentials"
}

// LocalName returns the local name for [DataSource].
func (gvnc *DataSource) LocalName() string {
	return gvnc.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gvnc *DataSource) Configuration() interface{} {
	return gvnc.Args
}

// Attributes returns the attributes for [DataSource].
func (gvnc *DataSource) Attributes() dataGoogleVmwareengineNsxCredentialsAttributes {
	return dataGoogleVmwareengineNsxCredentialsAttributes{ref: terra.ReferenceDataSource(gvnc)}
}

// DataArgs contains the configurations for google_vmwareengine_nsx_credentials.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
}

type dataGoogleVmwareengineNsxCredentialsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_vmwareengine_nsx_credentials.
func (gvnc dataGoogleVmwareengineNsxCredentialsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gvnc.ref.Append("id"))
}

// Parent returns a reference to field parent of google_vmwareengine_nsx_credentials.
func (gvnc dataGoogleVmwareengineNsxCredentialsAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(gvnc.ref.Append("parent"))
}

// Password returns a reference to field password of google_vmwareengine_nsx_credentials.
func (gvnc dataGoogleVmwareengineNsxCredentialsAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(gvnc.ref.Append("password"))
}

// Username returns a reference to field username of google_vmwareengine_nsx_credentials.
func (gvnc dataGoogleVmwareengineNsxCredentialsAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(gvnc.ref.Append("username"))
}
