// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	folderaccessapprovalsettings "github.com/golingon/terraproviders/google/4.72.1/folderaccessapprovalsettings"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFolderAccessApprovalSettings creates a new instance of [FolderAccessApprovalSettings].
func NewFolderAccessApprovalSettings(name string, args FolderAccessApprovalSettingsArgs) *FolderAccessApprovalSettings {
	return &FolderAccessApprovalSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FolderAccessApprovalSettings)(nil)

// FolderAccessApprovalSettings represents the Terraform resource google_folder_access_approval_settings.
type FolderAccessApprovalSettings struct {
	Name      string
	Args      FolderAccessApprovalSettingsArgs
	state     *folderAccessApprovalSettingsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FolderAccessApprovalSettings].
func (faas *FolderAccessApprovalSettings) Type() string {
	return "google_folder_access_approval_settings"
}

// LocalName returns the local name for [FolderAccessApprovalSettings].
func (faas *FolderAccessApprovalSettings) LocalName() string {
	return faas.Name
}

// Configuration returns the configuration (args) for [FolderAccessApprovalSettings].
func (faas *FolderAccessApprovalSettings) Configuration() interface{} {
	return faas.Args
}

// DependOn is used for other resources to depend on [FolderAccessApprovalSettings].
func (faas *FolderAccessApprovalSettings) DependOn() terra.Reference {
	return terra.ReferenceResource(faas)
}

// Dependencies returns the list of resources [FolderAccessApprovalSettings] depends_on.
func (faas *FolderAccessApprovalSettings) Dependencies() terra.Dependencies {
	return faas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FolderAccessApprovalSettings].
func (faas *FolderAccessApprovalSettings) LifecycleManagement() *terra.Lifecycle {
	return faas.Lifecycle
}

// Attributes returns the attributes for [FolderAccessApprovalSettings].
func (faas *FolderAccessApprovalSettings) Attributes() folderAccessApprovalSettingsAttributes {
	return folderAccessApprovalSettingsAttributes{ref: terra.ReferenceResource(faas)}
}

// ImportState imports the given attribute values into [FolderAccessApprovalSettings]'s state.
func (faas *FolderAccessApprovalSettings) ImportState(av io.Reader) error {
	faas.state = &folderAccessApprovalSettingsState{}
	if err := json.NewDecoder(av).Decode(faas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", faas.Type(), faas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FolderAccessApprovalSettings] has state.
func (faas *FolderAccessApprovalSettings) State() (*folderAccessApprovalSettingsState, bool) {
	return faas.state, faas.state != nil
}

// StateMust returns the state for [FolderAccessApprovalSettings]. Panics if the state is nil.
func (faas *FolderAccessApprovalSettings) StateMust() *folderAccessApprovalSettingsState {
	if faas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", faas.Type(), faas.LocalName()))
	}
	return faas.state
}

// FolderAccessApprovalSettingsArgs contains the configurations for google_folder_access_approval_settings.
type FolderAccessApprovalSettingsArgs struct {
	// ActiveKeyVersion: string, optional
	ActiveKeyVersion terra.StringValue `hcl:"active_key_version,attr"`
	// FolderId: string, required
	FolderId terra.StringValue `hcl:"folder_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NotificationEmails: set of string, optional
	NotificationEmails terra.SetValue[terra.StringValue] `hcl:"notification_emails,attr"`
	// EnrolledServices: min=1
	EnrolledServices []folderaccessapprovalsettings.EnrolledServices `hcl:"enrolled_services,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *folderaccessapprovalsettings.Timeouts `hcl:"timeouts,block"`
}
type folderAccessApprovalSettingsAttributes struct {
	ref terra.Reference
}

// ActiveKeyVersion returns a reference to field active_key_version of google_folder_access_approval_settings.
func (faas folderAccessApprovalSettingsAttributes) ActiveKeyVersion() terra.StringValue {
	return terra.ReferenceAsString(faas.ref.Append("active_key_version"))
}

// AncestorHasActiveKeyVersion returns a reference to field ancestor_has_active_key_version of google_folder_access_approval_settings.
func (faas folderAccessApprovalSettingsAttributes) AncestorHasActiveKeyVersion() terra.BoolValue {
	return terra.ReferenceAsBool(faas.ref.Append("ancestor_has_active_key_version"))
}

// EnrolledAncestor returns a reference to field enrolled_ancestor of google_folder_access_approval_settings.
func (faas folderAccessApprovalSettingsAttributes) EnrolledAncestor() terra.BoolValue {
	return terra.ReferenceAsBool(faas.ref.Append("enrolled_ancestor"))
}

// FolderId returns a reference to field folder_id of google_folder_access_approval_settings.
func (faas folderAccessApprovalSettingsAttributes) FolderId() terra.StringValue {
	return terra.ReferenceAsString(faas.ref.Append("folder_id"))
}

// Id returns a reference to field id of google_folder_access_approval_settings.
func (faas folderAccessApprovalSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(faas.ref.Append("id"))
}

// InvalidKeyVersion returns a reference to field invalid_key_version of google_folder_access_approval_settings.
func (faas folderAccessApprovalSettingsAttributes) InvalidKeyVersion() terra.BoolValue {
	return terra.ReferenceAsBool(faas.ref.Append("invalid_key_version"))
}

// Name returns a reference to field name of google_folder_access_approval_settings.
func (faas folderAccessApprovalSettingsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(faas.ref.Append("name"))
}

// NotificationEmails returns a reference to field notification_emails of google_folder_access_approval_settings.
func (faas folderAccessApprovalSettingsAttributes) NotificationEmails() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](faas.ref.Append("notification_emails"))
}

func (faas folderAccessApprovalSettingsAttributes) EnrolledServices() terra.SetValue[folderaccessapprovalsettings.EnrolledServicesAttributes] {
	return terra.ReferenceAsSet[folderaccessapprovalsettings.EnrolledServicesAttributes](faas.ref.Append("enrolled_services"))
}

func (faas folderAccessApprovalSettingsAttributes) Timeouts() folderaccessapprovalsettings.TimeoutsAttributes {
	return terra.ReferenceAsSingle[folderaccessapprovalsettings.TimeoutsAttributes](faas.ref.Append("timeouts"))
}

type folderAccessApprovalSettingsState struct {
	ActiveKeyVersion            string                                               `json:"active_key_version"`
	AncestorHasActiveKeyVersion bool                                                 `json:"ancestor_has_active_key_version"`
	EnrolledAncestor            bool                                                 `json:"enrolled_ancestor"`
	FolderId                    string                                               `json:"folder_id"`
	Id                          string                                               `json:"id"`
	InvalidKeyVersion           bool                                                 `json:"invalid_key_version"`
	Name                        string                                               `json:"name"`
	NotificationEmails          []string                                             `json:"notification_emails"`
	EnrolledServices            []folderaccessapprovalsettings.EnrolledServicesState `json:"enrolled_services"`
	Timeouts                    *folderaccessapprovalsettings.TimeoutsState          `json:"timeouts"`
}
