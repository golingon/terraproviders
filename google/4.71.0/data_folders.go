// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datafolders "github.com/golingon/terraproviders/google/4.71.0/datafolders"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataFolders creates a new instance of [DataFolders].
func NewDataFolders(name string, args DataFoldersArgs) *DataFolders {
	return &DataFolders{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataFolders)(nil)

// DataFolders represents the Terraform data resource google_folders.
type DataFolders struct {
	Name string
	Args DataFoldersArgs
}

// DataSource returns the Terraform object type for [DataFolders].
func (f *DataFolders) DataSource() string {
	return "google_folders"
}

// LocalName returns the local name for [DataFolders].
func (f *DataFolders) LocalName() string {
	return f.Name
}

// Configuration returns the configuration (args) for [DataFolders].
func (f *DataFolders) Configuration() interface{} {
	return f.Args
}

// Attributes returns the attributes for [DataFolders].
func (f *DataFolders) Attributes() dataFoldersAttributes {
	return dataFoldersAttributes{ref: terra.ReferenceDataResource(f)}
}

// DataFoldersArgs contains the configurations for google_folders.
type DataFoldersArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ParentId: string, required
	ParentId terra.StringValue `hcl:"parent_id,attr" validate:"required"`
	// Folders: min=0
	Folders []datafolders.Folders `hcl:"folders,block" validate:"min=0"`
}
type dataFoldersAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_folders.
func (f dataFoldersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("id"))
}

// ParentId returns a reference to field parent_id of google_folders.
func (f dataFoldersAttributes) ParentId() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("parent_id"))
}

func (f dataFoldersAttributes) Folders() terra.ListValue[datafolders.FoldersAttributes] {
	return terra.ReferenceAsList[datafolders.FoldersAttributes](f.ref.Append("folders"))
}
