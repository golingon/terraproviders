// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	folderiamauditconfig "github.com/golingon/terraproviders/googlebeta/4.62.0/folderiamauditconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFolderIamAuditConfig creates a new instance of [FolderIamAuditConfig].
func NewFolderIamAuditConfig(name string, args FolderIamAuditConfigArgs) *FolderIamAuditConfig {
	return &FolderIamAuditConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FolderIamAuditConfig)(nil)

// FolderIamAuditConfig represents the Terraform resource google_folder_iam_audit_config.
type FolderIamAuditConfig struct {
	Name      string
	Args      FolderIamAuditConfigArgs
	state     *folderIamAuditConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FolderIamAuditConfig].
func (fiac *FolderIamAuditConfig) Type() string {
	return "google_folder_iam_audit_config"
}

// LocalName returns the local name for [FolderIamAuditConfig].
func (fiac *FolderIamAuditConfig) LocalName() string {
	return fiac.Name
}

// Configuration returns the configuration (args) for [FolderIamAuditConfig].
func (fiac *FolderIamAuditConfig) Configuration() interface{} {
	return fiac.Args
}

// DependOn is used for other resources to depend on [FolderIamAuditConfig].
func (fiac *FolderIamAuditConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(fiac)
}

// Dependencies returns the list of resources [FolderIamAuditConfig] depends_on.
func (fiac *FolderIamAuditConfig) Dependencies() terra.Dependencies {
	return fiac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FolderIamAuditConfig].
func (fiac *FolderIamAuditConfig) LifecycleManagement() *terra.Lifecycle {
	return fiac.Lifecycle
}

// Attributes returns the attributes for [FolderIamAuditConfig].
func (fiac *FolderIamAuditConfig) Attributes() folderIamAuditConfigAttributes {
	return folderIamAuditConfigAttributes{ref: terra.ReferenceResource(fiac)}
}

// ImportState imports the given attribute values into [FolderIamAuditConfig]'s state.
func (fiac *FolderIamAuditConfig) ImportState(av io.Reader) error {
	fiac.state = &folderIamAuditConfigState{}
	if err := json.NewDecoder(av).Decode(fiac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fiac.Type(), fiac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FolderIamAuditConfig] has state.
func (fiac *FolderIamAuditConfig) State() (*folderIamAuditConfigState, bool) {
	return fiac.state, fiac.state != nil
}

// StateMust returns the state for [FolderIamAuditConfig]. Panics if the state is nil.
func (fiac *FolderIamAuditConfig) StateMust() *folderIamAuditConfigState {
	if fiac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fiac.Type(), fiac.LocalName()))
	}
	return fiac.state
}

// FolderIamAuditConfigArgs contains the configurations for google_folder_iam_audit_config.
type FolderIamAuditConfigArgs struct {
	// Folder: string, required
	Folder terra.StringValue `hcl:"folder,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// AuditLogConfig: min=1
	AuditLogConfig []folderiamauditconfig.AuditLogConfig `hcl:"audit_log_config,block" validate:"min=1"`
}
type folderIamAuditConfigAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_folder_iam_audit_config.
func (fiac folderIamAuditConfigAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(fiac.ref.Append("etag"))
}

// Folder returns a reference to field folder of google_folder_iam_audit_config.
func (fiac folderIamAuditConfigAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(fiac.ref.Append("folder"))
}

// Id returns a reference to field id of google_folder_iam_audit_config.
func (fiac folderIamAuditConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fiac.ref.Append("id"))
}

// Service returns a reference to field service of google_folder_iam_audit_config.
func (fiac folderIamAuditConfigAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(fiac.ref.Append("service"))
}

func (fiac folderIamAuditConfigAttributes) AuditLogConfig() terra.SetValue[folderiamauditconfig.AuditLogConfigAttributes] {
	return terra.ReferenceAsSet[folderiamauditconfig.AuditLogConfigAttributes](fiac.ref.Append("audit_log_config"))
}

type folderIamAuditConfigState struct {
	Etag           string                                     `json:"etag"`
	Folder         string                                     `json:"folder"`
	Id             string                                     `json:"id"`
	Service        string                                     `json:"service"`
	AuditLogConfig []folderiamauditconfig.AuditLogConfigState `json:"audit_log_config"`
}
