// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package archive

import (
	datafile "github.com/golingon/terraproviders/archive/2.4.0/datafile"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataFile creates a new instance of [DataFile].
func NewDataFile(name string, args DataFileArgs) *DataFile {
	return &DataFile{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataFile)(nil)

// DataFile represents the Terraform data resource archive_file.
type DataFile struct {
	Name string
	Args DataFileArgs
}

// DataSource returns the Terraform object type for [DataFile].
func (f *DataFile) DataSource() string {
	return "archive_file"
}

// LocalName returns the local name for [DataFile].
func (f *DataFile) LocalName() string {
	return f.Name
}

// Configuration returns the configuration (args) for [DataFile].
func (f *DataFile) Configuration() interface{} {
	return f.Args
}

// Attributes returns the attributes for [DataFile].
func (f *DataFile) Attributes() dataFileAttributes {
	return dataFileAttributes{ref: terra.ReferenceDataResource(f)}
}

// DataFileArgs contains the configurations for archive_file.
type DataFileArgs struct {
	// ExcludeSymlinkDirectories: bool, optional
	ExcludeSymlinkDirectories terra.BoolValue `hcl:"exclude_symlink_directories,attr"`
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
	Source []datafile.Source `hcl:"source,block" validate:"min=0"`
}
type dataFileAttributes struct {
	ref terra.Reference
}

// ExcludeSymlinkDirectories returns a reference to field exclude_symlink_directories of archive_file.
func (f dataFileAttributes) ExcludeSymlinkDirectories() terra.BoolValue {
	return terra.ReferenceAsBool(f.ref.Append("exclude_symlink_directories"))
}

// Excludes returns a reference to field excludes of archive_file.
func (f dataFileAttributes) Excludes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](f.ref.Append("excludes"))
}

// Id returns a reference to field id of archive_file.
func (f dataFileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("id"))
}

// OutputBase64Sha256 returns a reference to field output_base64sha256 of archive_file.
func (f dataFileAttributes) OutputBase64Sha256() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("output_base64sha256"))
}

// OutputBase64Sha512 returns a reference to field output_base64sha512 of archive_file.
func (f dataFileAttributes) OutputBase64Sha512() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("output_base64sha512"))
}

// OutputFileMode returns a reference to field output_file_mode of archive_file.
func (f dataFileAttributes) OutputFileMode() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("output_file_mode"))
}

// OutputMd5 returns a reference to field output_md5 of archive_file.
func (f dataFileAttributes) OutputMd5() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("output_md5"))
}

// OutputPath returns a reference to field output_path of archive_file.
func (f dataFileAttributes) OutputPath() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("output_path"))
}

// OutputSha returns a reference to field output_sha of archive_file.
func (f dataFileAttributes) OutputSha() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("output_sha"))
}

// OutputSha256 returns a reference to field output_sha256 of archive_file.
func (f dataFileAttributes) OutputSha256() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("output_sha256"))
}

// OutputSha512 returns a reference to field output_sha512 of archive_file.
func (f dataFileAttributes) OutputSha512() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("output_sha512"))
}

// OutputSize returns a reference to field output_size of archive_file.
func (f dataFileAttributes) OutputSize() terra.NumberValue {
	return terra.ReferenceAsNumber(f.ref.Append("output_size"))
}

// SourceContent returns a reference to field source_content of archive_file.
func (f dataFileAttributes) SourceContent() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("source_content"))
}

// SourceContentFilename returns a reference to field source_content_filename of archive_file.
func (f dataFileAttributes) SourceContentFilename() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("source_content_filename"))
}

// SourceDir returns a reference to field source_dir of archive_file.
func (f dataFileAttributes) SourceDir() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("source_dir"))
}

// SourceFile returns a reference to field source_file of archive_file.
func (f dataFileAttributes) SourceFile() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("source_file"))
}

// Type returns a reference to field type of archive_file.
func (f dataFileAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("type"))
}

func (f dataFileAttributes) Source() terra.SetValue[datafile.SourceAttributes] {
	return terra.ReferenceAsSet[datafile.SourceAttributes](f.ref.Append("source"))
}
