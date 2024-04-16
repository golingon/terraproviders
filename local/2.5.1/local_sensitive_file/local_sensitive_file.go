// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package local_sensitive_file

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

// Resource represents the Terraform resource local_sensitive_file.
type Resource struct {
	Name      string
	Args      Args
	state     *localSensitiveFileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (lsf *Resource) Type() string {
	return "local_sensitive_file"
}

// LocalName returns the local name for [Resource].
func (lsf *Resource) LocalName() string {
	return lsf.Name
}

// Configuration returns the configuration (args) for [Resource].
func (lsf *Resource) Configuration() interface{} {
	return lsf.Args
}

// DependOn is used for other resources to depend on [Resource].
func (lsf *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(lsf)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (lsf *Resource) Dependencies() terra.Dependencies {
	return lsf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (lsf *Resource) LifecycleManagement() *terra.Lifecycle {
	return lsf.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (lsf *Resource) Attributes() localSensitiveFileAttributes {
	return localSensitiveFileAttributes{ref: terra.ReferenceResource(lsf)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (lsf *Resource) ImportState(state io.Reader) error {
	lsf.state = &localSensitiveFileState{}
	if err := json.NewDecoder(state).Decode(lsf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lsf.Type(), lsf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (lsf *Resource) State() (*localSensitiveFileState, bool) {
	return lsf.state, lsf.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (lsf *Resource) StateMust() *localSensitiveFileState {
	if lsf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lsf.Type(), lsf.LocalName()))
	}
	return lsf.state
}

// Args contains the configurations for local_sensitive_file.
type Args struct {
	// Content: string, optional
	Content terra.StringValue `hcl:"content,attr"`
	// ContentBase64: string, optional
	ContentBase64 terra.StringValue `hcl:"content_base64,attr"`
	// DirectoryPermission: string, optional
	DirectoryPermission terra.StringValue `hcl:"directory_permission,attr"`
	// FilePermission: string, optional
	FilePermission terra.StringValue `hcl:"file_permission,attr"`
	// Filename: string, required
	Filename terra.StringValue `hcl:"filename,attr" validate:"required"`
	// Source: string, optional
	Source terra.StringValue `hcl:"source,attr"`
}

type localSensitiveFileAttributes struct {
	ref terra.Reference
}

// Content returns a reference to field content of local_sensitive_file.
func (lsf localSensitiveFileAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(lsf.ref.Append("content"))
}

// ContentBase64 returns a reference to field content_base64 of local_sensitive_file.
func (lsf localSensitiveFileAttributes) ContentBase64() terra.StringValue {
	return terra.ReferenceAsString(lsf.ref.Append("content_base64"))
}

// ContentBase64Sha256 returns a reference to field content_base64sha256 of local_sensitive_file.
func (lsf localSensitiveFileAttributes) ContentBase64Sha256() terra.StringValue {
	return terra.ReferenceAsString(lsf.ref.Append("content_base64sha256"))
}

// ContentBase64Sha512 returns a reference to field content_base64sha512 of local_sensitive_file.
func (lsf localSensitiveFileAttributes) ContentBase64Sha512() terra.StringValue {
	return terra.ReferenceAsString(lsf.ref.Append("content_base64sha512"))
}

// ContentMd5 returns a reference to field content_md5 of local_sensitive_file.
func (lsf localSensitiveFileAttributes) ContentMd5() terra.StringValue {
	return terra.ReferenceAsString(lsf.ref.Append("content_md5"))
}

// ContentSha1 returns a reference to field content_sha1 of local_sensitive_file.
func (lsf localSensitiveFileAttributes) ContentSha1() terra.StringValue {
	return terra.ReferenceAsString(lsf.ref.Append("content_sha1"))
}

// ContentSha256 returns a reference to field content_sha256 of local_sensitive_file.
func (lsf localSensitiveFileAttributes) ContentSha256() terra.StringValue {
	return terra.ReferenceAsString(lsf.ref.Append("content_sha256"))
}

// ContentSha512 returns a reference to field content_sha512 of local_sensitive_file.
func (lsf localSensitiveFileAttributes) ContentSha512() terra.StringValue {
	return terra.ReferenceAsString(lsf.ref.Append("content_sha512"))
}

// DirectoryPermission returns a reference to field directory_permission of local_sensitive_file.
func (lsf localSensitiveFileAttributes) DirectoryPermission() terra.StringValue {
	return terra.ReferenceAsString(lsf.ref.Append("directory_permission"))
}

// FilePermission returns a reference to field file_permission of local_sensitive_file.
func (lsf localSensitiveFileAttributes) FilePermission() terra.StringValue {
	return terra.ReferenceAsString(lsf.ref.Append("file_permission"))
}

// Filename returns a reference to field filename of local_sensitive_file.
func (lsf localSensitiveFileAttributes) Filename() terra.StringValue {
	return terra.ReferenceAsString(lsf.ref.Append("filename"))
}

// Id returns a reference to field id of local_sensitive_file.
func (lsf localSensitiveFileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lsf.ref.Append("id"))
}

// Source returns a reference to field source of local_sensitive_file.
func (lsf localSensitiveFileAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(lsf.ref.Append("source"))
}

type localSensitiveFileState struct {
	Content             string `json:"content"`
	ContentBase64       string `json:"content_base64"`
	ContentBase64Sha256 string `json:"content_base64sha256"`
	ContentBase64Sha512 string `json:"content_base64sha512"`
	ContentMd5          string `json:"content_md5"`
	ContentSha1         string `json:"content_sha1"`
	ContentSha256       string `json:"content_sha256"`
	ContentSha512       string `json:"content_sha512"`
	DirectoryPermission string `json:"directory_permission"`
	FilePermission      string `json:"file_permission"`
	Filename            string `json:"filename"`
	Id                  string `json:"id"`
	Source              string `json:"source"`
}
