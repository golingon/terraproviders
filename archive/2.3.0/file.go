// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package archive

import (
	"encoding/json"
	"fmt"
	file "github.com/golingon/terraproviders/archive/2.3.0/file"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFile creates a new instance of [File].
func NewFile(name string, args FileArgs) *File {
	return &File{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*File)(nil)

// File represents the Terraform resource archive_file.
type File struct {
	Name      string
	Args      FileArgs
	state     *fileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [File].
func (f *File) Type() string {
	return "archive_file"
}

// LocalName returns the local name for [File].
func (f *File) LocalName() string {
	return f.Name
}

// Configuration returns the configuration (args) for [File].
func (f *File) Configuration() interface{} {
	return f.Args
}

// DependOn is used for other resources to depend on [File].
func (f *File) DependOn() terra.Reference {
	return terra.ReferenceResource(f)
}

// Dependencies returns the list of resources [File] depends_on.
func (f *File) Dependencies() terra.Dependencies {
	return f.DependsOn
}

// LifecycleManagement returns the lifecycle block for [File].
func (f *File) LifecycleManagement() *terra.Lifecycle {
	return f.Lifecycle
}

// Attributes returns the attributes for [File].
func (f *File) Attributes() fileAttributes {
	return fileAttributes{ref: terra.ReferenceResource(f)}
}

// ImportState imports the given attribute values into [File]'s state.
func (f *File) ImportState(av io.Reader) error {
	f.state = &fileState{}
	if err := json.NewDecoder(av).Decode(f.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", f.Type(), f.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [File] has state.
func (f *File) State() (*fileState, bool) {
	return f.state, f.state != nil
}

// StateMust returns the state for [File]. Panics if the state is nil.
func (f *File) StateMust() *fileState {
	if f.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", f.Type(), f.LocalName()))
	}
	return f.state
}

// FileArgs contains the configurations for archive_file.
type FileArgs struct {
	// Excludes: set of string, optional
	Excludes terra.SetValue[terra.StringValue] `hcl:"excludes,attr"`
	// OutputFileMode: string, optional
	OutputFileMode terra.StringValue `hcl:"output_file_mode,attr"`
	// OutputPath: string, required
	OutputPath terra.StringValue `hcl:"output_path,attr" validate:"required"`
	// SourceContent: string, optional
	SourceContent terra.StringValue `hcl:"source_content,attr"`
	// SourceContentFilename: string, optional
	SourceContentFilename terra.StringValue `hcl:"source_content_filename,attr"`
	// SourceDir: string, optional
	SourceDir terra.StringValue `hcl:"source_dir,attr"`
	// SourceFile: string, optional
	SourceFile terra.StringValue `hcl:"source_file,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Source: min=0
	Source []file.Source `hcl:"source,block" validate:"min=0"`
}
type fileAttributes struct {
	ref terra.Reference
}

// Excludes returns a reference to field excludes of archive_file.
func (f fileAttributes) Excludes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](f.ref.Append("excludes"))
}

// Id returns a reference to field id of archive_file.
func (f fileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("id"))
}

// OutputBase64Sha256 returns a reference to field output_base64sha256 of archive_file.
func (f fileAttributes) OutputBase64Sha256() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("output_base64sha256"))
}

// OutputFileMode returns a reference to field output_file_mode of archive_file.
func (f fileAttributes) OutputFileMode() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("output_file_mode"))
}

// OutputMd5 returns a reference to field output_md5 of archive_file.
func (f fileAttributes) OutputMd5() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("output_md5"))
}

// OutputPath returns a reference to field output_path of archive_file.
func (f fileAttributes) OutputPath() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("output_path"))
}

// OutputSha returns a reference to field output_sha of archive_file.
func (f fileAttributes) OutputSha() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("output_sha"))
}

// OutputSize returns a reference to field output_size of archive_file.
func (f fileAttributes) OutputSize() terra.NumberValue {
	return terra.ReferenceAsNumber(f.ref.Append("output_size"))
}

// SourceContent returns a reference to field source_content of archive_file.
func (f fileAttributes) SourceContent() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("source_content"))
}

// SourceContentFilename returns a reference to field source_content_filename of archive_file.
func (f fileAttributes) SourceContentFilename() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("source_content_filename"))
}

// SourceDir returns a reference to field source_dir of archive_file.
func (f fileAttributes) SourceDir() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("source_dir"))
}

// SourceFile returns a reference to field source_file of archive_file.
func (f fileAttributes) SourceFile() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("source_file"))
}

// Type returns a reference to field type of archive_file.
func (f fileAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("type"))
}

func (f fileAttributes) Source() terra.SetValue[file.SourceAttributes] {
	return terra.ReferenceAsSet[file.SourceAttributes](f.ref.Append("source"))
}

type fileState struct {
	Excludes              []string           `json:"excludes"`
	Id                    string             `json:"id"`
	OutputBase64Sha256    string             `json:"output_base64sha256"`
	OutputFileMode        string             `json:"output_file_mode"`
	OutputMd5             string             `json:"output_md5"`
	OutputPath            string             `json:"output_path"`
	OutputSha             string             `json:"output_sha"`
	OutputSize            float64            `json:"output_size"`
	SourceContent         string             `json:"source_content"`
	SourceContentFilename string             `json:"source_content_filename"`
	SourceDir             string             `json:"source_dir"`
	SourceFile            string             `json:"source_file"`
	Type                  string             `json:"type"`
	Source                []file.SourceState `json:"source"`
}
