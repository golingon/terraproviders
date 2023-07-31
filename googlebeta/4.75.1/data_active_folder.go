// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataActiveFolder creates a new instance of [DataActiveFolder].
func NewDataActiveFolder(name string, args DataActiveFolderArgs) *DataActiveFolder {
	return &DataActiveFolder{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataActiveFolder)(nil)

// DataActiveFolder represents the Terraform data resource google_active_folder.
type DataActiveFolder struct {
	Name string
	Args DataActiveFolderArgs
}

// DataSource returns the Terraform object type for [DataActiveFolder].
func (af *DataActiveFolder) DataSource() string {
	return "google_active_folder"
}

// LocalName returns the local name for [DataActiveFolder].
func (af *DataActiveFolder) LocalName() string {
	return af.Name
}

// Configuration returns the configuration (args) for [DataActiveFolder].
func (af *DataActiveFolder) Configuration() interface{} {
	return af.Args
}

// Attributes returns the attributes for [DataActiveFolder].
func (af *DataActiveFolder) Attributes() dataActiveFolderAttributes {
	return dataActiveFolderAttributes{ref: terra.ReferenceDataResource(af)}
}

// DataActiveFolderArgs contains the configurations for google_active_folder.
type DataActiveFolderArgs struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
}
type dataActiveFolderAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_active_folder.
func (af dataActiveFolderAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("display_name"))
}

// Id returns a reference to field id of google_active_folder.
func (af dataActiveFolderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("id"))
}

// Name returns a reference to field name of google_active_folder.
func (af dataActiveFolderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("name"))
}

// Parent returns a reference to field parent of google_active_folder.
func (af dataActiveFolderAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("parent"))
}
