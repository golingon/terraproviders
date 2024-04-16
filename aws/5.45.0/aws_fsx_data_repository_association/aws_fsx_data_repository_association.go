// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_fsx_data_repository_association

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_fsx_data_repository_association.
type Resource struct {
	Name      string
	Args      Args
	state     *awsFsxDataRepositoryAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (afdra *Resource) Type() string {
	return "aws_fsx_data_repository_association"
}

// LocalName returns the local name for [Resource].
func (afdra *Resource) LocalName() string {
	return afdra.Name
}

// Configuration returns the configuration (args) for [Resource].
func (afdra *Resource) Configuration() interface{} {
	return afdra.Args
}

// DependOn is used for other resources to depend on [Resource].
func (afdra *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(afdra)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (afdra *Resource) Dependencies() terra.Dependencies {
	return afdra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (afdra *Resource) LifecycleManagement() *terra.Lifecycle {
	return afdra.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (afdra *Resource) Attributes() awsFsxDataRepositoryAssociationAttributes {
	return awsFsxDataRepositoryAssociationAttributes{ref: terra.ReferenceResource(afdra)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (afdra *Resource) ImportState(state io.Reader) error {
	afdra.state = &awsFsxDataRepositoryAssociationState{}
	if err := json.NewDecoder(state).Decode(afdra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", afdra.Type(), afdra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (afdra *Resource) State() (*awsFsxDataRepositoryAssociationState, bool) {
	return afdra.state, afdra.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (afdra *Resource) StateMust() *awsFsxDataRepositoryAssociationState {
	if afdra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", afdra.Type(), afdra.LocalName()))
	}
	return afdra.state
}

// Args contains the configurations for aws_fsx_data_repository_association.
type Args struct {
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
	S3 *S3 `hcl:"s3,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsFsxDataRepositoryAssociationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_fsx_data_repository_association.
func (afdra awsFsxDataRepositoryAssociationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(afdra.ref.Append("arn"))
}

// AssociationId returns a reference to field association_id of aws_fsx_data_repository_association.
func (afdra awsFsxDataRepositoryAssociationAttributes) AssociationId() terra.StringValue {
	return terra.ReferenceAsString(afdra.ref.Append("association_id"))
}

// BatchImportMetaDataOnCreate returns a reference to field batch_import_meta_data_on_create of aws_fsx_data_repository_association.
func (afdra awsFsxDataRepositoryAssociationAttributes) BatchImportMetaDataOnCreate() terra.BoolValue {
	return terra.ReferenceAsBool(afdra.ref.Append("batch_import_meta_data_on_create"))
}

// DataRepositoryPath returns a reference to field data_repository_path of aws_fsx_data_repository_association.
func (afdra awsFsxDataRepositoryAssociationAttributes) DataRepositoryPath() terra.StringValue {
	return terra.ReferenceAsString(afdra.ref.Append("data_repository_path"))
}

// DeleteDataInFilesystem returns a reference to field delete_data_in_filesystem of aws_fsx_data_repository_association.
func (afdra awsFsxDataRepositoryAssociationAttributes) DeleteDataInFilesystem() terra.BoolValue {
	return terra.ReferenceAsBool(afdra.ref.Append("delete_data_in_filesystem"))
}

// FileSystemId returns a reference to field file_system_id of aws_fsx_data_repository_association.
func (afdra awsFsxDataRepositoryAssociationAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(afdra.ref.Append("file_system_id"))
}

// FileSystemPath returns a reference to field file_system_path of aws_fsx_data_repository_association.
func (afdra awsFsxDataRepositoryAssociationAttributes) FileSystemPath() terra.StringValue {
	return terra.ReferenceAsString(afdra.ref.Append("file_system_path"))
}

// Id returns a reference to field id of aws_fsx_data_repository_association.
func (afdra awsFsxDataRepositoryAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(afdra.ref.Append("id"))
}

// ImportedFileChunkSize returns a reference to field imported_file_chunk_size of aws_fsx_data_repository_association.
func (afdra awsFsxDataRepositoryAssociationAttributes) ImportedFileChunkSize() terra.NumberValue {
	return terra.ReferenceAsNumber(afdra.ref.Append("imported_file_chunk_size"))
}

// Tags returns a reference to field tags of aws_fsx_data_repository_association.
func (afdra awsFsxDataRepositoryAssociationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](afdra.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_fsx_data_repository_association.
func (afdra awsFsxDataRepositoryAssociationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](afdra.ref.Append("tags_all"))
}

func (afdra awsFsxDataRepositoryAssociationAttributes) S3() terra.ListValue[S3Attributes] {
	return terra.ReferenceAsList[S3Attributes](afdra.ref.Append("s3"))
}

func (afdra awsFsxDataRepositoryAssociationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](afdra.ref.Append("timeouts"))
}

type awsFsxDataRepositoryAssociationState struct {
	Arn                         string            `json:"arn"`
	AssociationId               string            `json:"association_id"`
	BatchImportMetaDataOnCreate bool              `json:"batch_import_meta_data_on_create"`
	DataRepositoryPath          string            `json:"data_repository_path"`
	DeleteDataInFilesystem      bool              `json:"delete_data_in_filesystem"`
	FileSystemId                string            `json:"file_system_id"`
	FileSystemPath              string            `json:"file_system_path"`
	Id                          string            `json:"id"`
	ImportedFileChunkSize       float64           `json:"imported_file_chunk_size"`
	Tags                        map[string]string `json:"tags"`
	TagsAll                     map[string]string `json:"tags_all"`
	S3                          []S3State         `json:"s3"`
	Timeouts                    *TimeoutsState    `json:"timeouts"`
}
