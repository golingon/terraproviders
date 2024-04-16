// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_folder_access_approval_settings

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

// Resource represents the Terraform resource google_folder_access_approval_settings.
type Resource struct {
	Name      string
	Args      Args
	state     *googleFolderAccessApprovalSettingsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gfaas *Resource) Type() string {
	return "google_folder_access_approval_settings"
}

// LocalName returns the local name for [Resource].
func (gfaas *Resource) LocalName() string {
	return gfaas.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gfaas *Resource) Configuration() interface{} {
	return gfaas.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gfaas *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gfaas)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gfaas *Resource) Dependencies() terra.Dependencies {
	return gfaas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gfaas *Resource) LifecycleManagement() *terra.Lifecycle {
	return gfaas.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gfaas *Resource) Attributes() googleFolderAccessApprovalSettingsAttributes {
	return googleFolderAccessApprovalSettingsAttributes{ref: terra.ReferenceResource(gfaas)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gfaas *Resource) ImportState(state io.Reader) error {
	gfaas.state = &googleFolderAccessApprovalSettingsState{}
	if err := json.NewDecoder(state).Decode(gfaas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gfaas.Type(), gfaas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gfaas *Resource) State() (*googleFolderAccessApprovalSettingsState, bool) {
	return gfaas.state, gfaas.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gfaas *Resource) StateMust() *googleFolderAccessApprovalSettingsState {
	if gfaas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gfaas.Type(), gfaas.LocalName()))
	}
	return gfaas.state
}

// Args contains the configurations for google_folder_access_approval_settings.
type Args struct {
	// ActiveKeyVersion: string, optional
	ActiveKeyVersion terra.StringValue `hcl:"active_key_version,attr"`
	// FolderId: string, required
	FolderId terra.StringValue `hcl:"folder_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NotificationEmails: set of string, optional
	NotificationEmails terra.SetValue[terra.StringValue] `hcl:"notification_emails,attr"`
	// EnrolledServices: min=1
	EnrolledServices []EnrolledServices `hcl:"enrolled_services,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleFolderAccessApprovalSettingsAttributes struct {
	ref terra.Reference
}

// ActiveKeyVersion returns a reference to field active_key_version of google_folder_access_approval_settings.
func (gfaas googleFolderAccessApprovalSettingsAttributes) ActiveKeyVersion() terra.StringValue {
	return terra.ReferenceAsString(gfaas.ref.Append("active_key_version"))
}

// AncestorHasActiveKeyVersion returns a reference to field ancestor_has_active_key_version of google_folder_access_approval_settings.
func (gfaas googleFolderAccessApprovalSettingsAttributes) AncestorHasActiveKeyVersion() terra.BoolValue {
	return terra.ReferenceAsBool(gfaas.ref.Append("ancestor_has_active_key_version"))
}

// EnrolledAncestor returns a reference to field enrolled_ancestor of google_folder_access_approval_settings.
func (gfaas googleFolderAccessApprovalSettingsAttributes) EnrolledAncestor() terra.BoolValue {
	return terra.ReferenceAsBool(gfaas.ref.Append("enrolled_ancestor"))
}

// FolderId returns a reference to field folder_id of google_folder_access_approval_settings.
func (gfaas googleFolderAccessApprovalSettingsAttributes) FolderId() terra.StringValue {
	return terra.ReferenceAsString(gfaas.ref.Append("folder_id"))
}

// Id returns a reference to field id of google_folder_access_approval_settings.
func (gfaas googleFolderAccessApprovalSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gfaas.ref.Append("id"))
}

// InvalidKeyVersion returns a reference to field invalid_key_version of google_folder_access_approval_settings.
func (gfaas googleFolderAccessApprovalSettingsAttributes) InvalidKeyVersion() terra.BoolValue {
	return terra.ReferenceAsBool(gfaas.ref.Append("invalid_key_version"))
}

// Name returns a reference to field name of google_folder_access_approval_settings.
func (gfaas googleFolderAccessApprovalSettingsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gfaas.ref.Append("name"))
}

// NotificationEmails returns a reference to field notification_emails of google_folder_access_approval_settings.
func (gfaas googleFolderAccessApprovalSettingsAttributes) NotificationEmails() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gfaas.ref.Append("notification_emails"))
}

func (gfaas googleFolderAccessApprovalSettingsAttributes) EnrolledServices() terra.SetValue[EnrolledServicesAttributes] {
	return terra.ReferenceAsSet[EnrolledServicesAttributes](gfaas.ref.Append("enrolled_services"))
}

func (gfaas googleFolderAccessApprovalSettingsAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gfaas.ref.Append("timeouts"))
}

type googleFolderAccessApprovalSettingsState struct {
	ActiveKeyVersion            string                  `json:"active_key_version"`
	AncestorHasActiveKeyVersion bool                    `json:"ancestor_has_active_key_version"`
	EnrolledAncestor            bool                    `json:"enrolled_ancestor"`
	FolderId                    string                  `json:"folder_id"`
	Id                          string                  `json:"id"`
	InvalidKeyVersion           bool                    `json:"invalid_key_version"`
	Name                        string                  `json:"name"`
	NotificationEmails          []string                `json:"notification_emails"`
	EnrolledServices            []EnrolledServicesState `json:"enrolled_services"`
	Timeouts                    *TimeoutsState          `json:"timeouts"`
}
