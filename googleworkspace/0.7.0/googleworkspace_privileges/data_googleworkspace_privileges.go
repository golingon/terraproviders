// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googleworkspace_privileges

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource googleworkspace_privileges.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gp *DataSource) DataSource() string {
	return "googleworkspace_privileges"
}

// LocalName returns the local name for [DataSource].
func (gp *DataSource) LocalName() string {
	return gp.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gp *DataSource) Configuration() interface{} {
	return gp.Args
}

// Attributes returns the attributes for [DataSource].
func (gp *DataSource) Attributes() dataGoogleworkspacePrivilegesAttributes {
	return dataGoogleworkspacePrivilegesAttributes{ref: terra.ReferenceDataSource(gp)}
}

// DataArgs contains the configurations for googleworkspace_privileges.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}

type dataGoogleworkspacePrivilegesAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of googleworkspace_privileges.
func (gp dataGoogleworkspacePrivilegesAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gp.ref.Append("etag"))
}

// Id returns a reference to field id of googleworkspace_privileges.
func (gp dataGoogleworkspacePrivilegesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gp.ref.Append("id"))
}

func (gp dataGoogleworkspacePrivilegesAttributes) Items() terra.ListValue[DataItemsAttributes] {
	return terra.ReferenceAsList[DataItemsAttributes](gp.ref.Append("items"))
}
