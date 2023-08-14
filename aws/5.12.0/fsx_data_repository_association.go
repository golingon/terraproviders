// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	fsxdatarepositoryassociation "github.com/golingon/terraproviders/aws/5.12.0/fsxdatarepositoryassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFsxDataRepositoryAssociation creates a new instance of [FsxDataRepositoryAssociation].
func NewFsxDataRepositoryAssociation(name string, args FsxDataRepositoryAssociationArgs) *FsxDataRepositoryAssociation {
	return &FsxDataRepositoryAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FsxDataRepositoryAssociation)(nil)

// FsxDataRepositoryAssociation represents the Terraform resource aws_fsx_data_repository_association.
type FsxDataRepositoryAssociation struct {
	Name      string
	Args      FsxDataRepositoryAssociationArgs
	state     *fsxDataRepositoryAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FsxDataRepositoryAssociation].
func (fdra *FsxDataRepositoryAssociation) Type() string {
	return "aws_fsx_data_repository_association"
}

// LocalName returns the local name for [FsxDataRepositoryAssociation].
func (fdra *FsxDataRepositoryAssociation) LocalName() string {
	return fdra.Name
}

// Configuration returns the configuration (args) for [FsxDataRepositoryAssociation].
func (fdra *FsxDataRepositoryAssociation) Configuration() interface{} {
	return fdra.Args
}

// DependOn is used for other resources to depend on [FsxDataRepositoryAssociation].
func (fdra *FsxDataRepositoryAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(fdra)
}

// Dependencies returns the list of resources [FsxDataRepositoryAssociation] depends_on.
func (fdra *FsxDataRepositoryAssociation) Dependencies() terra.Dependencies {
	return fdra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FsxDataRepositoryAssociation].
func (fdra *FsxDataRepositoryAssociation) LifecycleManagement() *terra.Lifecycle {
	return fdra.Lifecycle
}

// Attributes returns the attributes for [FsxDataRepositoryAssociation].
func (fdra *FsxDataRepositoryAssociation) Attributes() fsxDataRepositoryAssociationAttributes {
	return fsxDataRepositoryAssociationAttributes{ref: terra.ReferenceResource(fdra)}
}

// ImportState imports the given attribute values into [FsxDataRepositoryAssociation]'s state.
func (fdra *FsxDataRepositoryAssociation) ImportState(av io.Reader) error {
	fdra.state = &fsxDataRepositoryAssociationState{}
	if err := json.NewDecoder(av).Decode(fdra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fdra.Type(), fdra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FsxDataRepositoryAssociation] has state.
func (fdra *FsxDataRepositoryAssociation) State() (*fsxDataRepositoryAssociationState, bool) {
	return fdra.state, fdra.state != nil
}

// StateMust returns the state for [FsxDataRepositoryAssociation]. Panics if the state is nil.
func (fdra *FsxDataRepositoryAssociation) StateMust() *fsxDataRepositoryAssociationState {
	if fdra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fdra.Type(), fdra.LocalName()))
	}
	return fdra.state
}

// FsxDataRepositoryAssociationArgs contains the configurations for aws_fsx_data_repository_association.
type FsxDataRepositoryAssociationArgs struct {
	// BatchImportMetaDataOnCreate: bool, optional
	BatchImportMetaDataOnCreate terra.BoolValue `hcl:"batch_import_meta_data_on_create,attr"`
	// DataRepositoryPath: string, required
	DataRepositoryPath terra.StringValue `hcl:"data_repository_path,attr" validate:"required"`
	// DeleteDataInFilesystem: bool, optional
	DeleteDataInFilesystem terra.BoolValue `hcl:"delete_data_in_filesystem,attr"`
	// FileSystemId: string, required
	FileSystemId terra.StringValue `hcl:"file_system_id,attr" validate:"required"`
	// FileSystemPath: string, required
	FileSystemPath terra.StringValue `hcl:"file_system_path,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImportedFileChunkSize: number, optional
	ImportedFileChunkSize terra.NumberValue `hcl:"imported_file_chunk_size,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// S3: optional
	S3 *fsxdatarepositoryassociation.S3 `hcl:"s3,block"`
	// Timeouts: optional
	Timeouts *fsxdatarepositoryassociation.Timeouts `hcl:"timeouts,block"`
}
type fsxDataRepositoryAssociationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_fsx_data_repository_association.
func (fdra fsxDataRepositoryAssociationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(fdra.ref.Append("arn"))
}

// AssociationId returns a reference to field association_id of aws_fsx_data_repository_association.
func (fdra fsxDataRepositoryAssociationAttributes) AssociationId() terra.StringValue {
	return terra.ReferenceAsString(fdra.ref.Append("association_id"))
}

// BatchImportMetaDataOnCreate returns a reference to field batch_import_meta_data_on_create of aws_fsx_data_repository_association.
func (fdra fsxDataRepositoryAssociationAttributes) BatchImportMetaDataOnCreate() terra.BoolValue {
	return terra.ReferenceAsBool(fdra.ref.Append("batch_import_meta_data_on_create"))
}

// DataRepositoryPath returns a reference to field data_repository_path of aws_fsx_data_repository_association.
func (fdra fsxDataRepositoryAssociationAttributes) DataRepositoryPath() terra.StringValue {
	return terra.ReferenceAsString(fdra.ref.Append("data_repository_path"))
}

// DeleteDataInFilesystem returns a reference to field delete_data_in_filesystem of aws_fsx_data_repository_association.
func (fdra fsxDataRepositoryAssociationAttributes) DeleteDataInFilesystem() terra.BoolValue {
	return terra.ReferenceAsBool(fdra.ref.Append("delete_data_in_filesystem"))
}

// FileSystemId returns a reference to field file_system_id of aws_fsx_data_repository_association.
func (fdra fsxDataRepositoryAssociationAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(fdra.ref.Append("file_system_id"))
}

// FileSystemPath returns a reference to field file_system_path of aws_fsx_data_repository_association.
func (fdra fsxDataRepositoryAssociationAttributes) FileSystemPath() terra.StringValue {
	return terra.ReferenceAsString(fdra.ref.Append("file_system_path"))
}

// Id returns a reference to field id of aws_fsx_data_repository_association.
func (fdra fsxDataRepositoryAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fdra.ref.Append("id"))
}

// ImportedFileChunkSize returns a reference to field imported_file_chunk_size of aws_fsx_data_repository_association.
func (fdra fsxDataRepositoryAssociationAttributes) ImportedFileChunkSize() terra.NumberValue {
	return terra.ReferenceAsNumber(fdra.ref.Append("imported_file_chunk_size"))
}

// Tags returns a reference to field tags of aws_fsx_data_repository_association.
func (fdra fsxDataRepositoryAssociationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fdra.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_fsx_data_repository_association.
func (fdra fsxDataRepositoryAssociationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fdra.ref.Append("tags_all"))
}

func (fdra fsxDataRepositoryAssociationAttributes) S3() terra.ListValue[fsxdatarepositoryassociation.S3Attributes] {
	return terra.ReferenceAsList[fsxdatarepositoryassociation.S3Attributes](fdra.ref.Append("s3"))
}

func (fdra fsxDataRepositoryAssociationAttributes) Timeouts() fsxdatarepositoryassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[fsxdatarepositoryassociation.TimeoutsAttributes](fdra.ref.Append("timeouts"))
}

type fsxDataRepositoryAssociationState struct {
	Arn                         string                                      `json:"arn"`
	AssociationId               string                                      `json:"association_id"`
	BatchImportMetaDataOnCreate bool                                        `json:"batch_import_meta_data_on_create"`
	DataRepositoryPath          string                                      `json:"data_repository_path"`
	DeleteDataInFilesystem      bool                                        `json:"delete_data_in_filesystem"`
	FileSystemId                string                                      `json:"file_system_id"`
	FileSystemPath              string                                      `json:"file_system_path"`
	Id                          string                                      `json:"id"`
	ImportedFileChunkSize       float64                                     `json:"imported_file_chunk_size"`
	Tags                        map[string]string                           `json:"tags"`
	TagsAll                     map[string]string                           `json:"tags_all"`
	S3                          []fsxdatarepositoryassociation.S3State      `json:"s3"`
	Timeouts                    *fsxdatarepositoryassociation.TimeoutsState `json:"timeouts"`
}
