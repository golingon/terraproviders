// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	projectaccessapprovalsettings "github.com/golingon/terraproviders/googlebeta/4.73.1/projectaccessapprovalsettings"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewProjectAccessApprovalSettings creates a new instance of [ProjectAccessApprovalSettings].
func NewProjectAccessApprovalSettings(name string, args ProjectAccessApprovalSettingsArgs) *ProjectAccessApprovalSettings {
	return &ProjectAccessApprovalSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ProjectAccessApprovalSettings)(nil)

// ProjectAccessApprovalSettings represents the Terraform resource google_project_access_approval_settings.
type ProjectAccessApprovalSettings struct {
	Name      string
	Args      ProjectAccessApprovalSettingsArgs
	state     *projectAccessApprovalSettingsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ProjectAccessApprovalSettings].
func (paas *ProjectAccessApprovalSettings) Type() string {
	return "google_project_access_approval_settings"
}

// LocalName returns the local name for [ProjectAccessApprovalSettings].
func (paas *ProjectAccessApprovalSettings) LocalName() string {
	return paas.Name
}

// Configuration returns the configuration (args) for [ProjectAccessApprovalSettings].
func (paas *ProjectAccessApprovalSettings) Configuration() interface{} {
	return paas.Args
}

// DependOn is used for other resources to depend on [ProjectAccessApprovalSettings].
func (paas *ProjectAccessApprovalSettings) DependOn() terra.Reference {
	return terra.ReferenceResource(paas)
}

// Dependencies returns the list of resources [ProjectAccessApprovalSettings] depends_on.
func (paas *ProjectAccessApprovalSettings) Dependencies() terra.Dependencies {
	return paas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ProjectAccessApprovalSettings].
func (paas *ProjectAccessApprovalSettings) LifecycleManagement() *terra.Lifecycle {
	return paas.Lifecycle
}

// Attributes returns the attributes for [ProjectAccessApprovalSettings].
func (paas *ProjectAccessApprovalSettings) Attributes() projectAccessApprovalSettingsAttributes {
	return projectAccessApprovalSettingsAttributes{ref: terra.ReferenceResource(paas)}
}

// ImportState imports the given attribute values into [ProjectAccessApprovalSettings]'s state.
func (paas *ProjectAccessApprovalSettings) ImportState(av io.Reader) error {
	paas.state = &projectAccessApprovalSettingsState{}
	if err := json.NewDecoder(av).Decode(paas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", paas.Type(), paas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ProjectAccessApprovalSettings] has state.
func (paas *ProjectAccessApprovalSettings) State() (*projectAccessApprovalSettingsState, bool) {
	return paas.state, paas.state != nil
}

// StateMust returns the state for [ProjectAccessApprovalSettings]. Panics if the state is nil.
func (paas *ProjectAccessApprovalSettings) StateMust() *projectAccessApprovalSettingsState {
	if paas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", paas.Type(), paas.LocalName()))
	}
	return paas.state
}

// ProjectAccessApprovalSettingsArgs contains the configurations for google_project_access_approval_settings.
type ProjectAccessApprovalSettingsArgs struct {
	// ActiveKeyVersion: string, optional
	ActiveKeyVersion terra.StringValue `hcl:"active_key_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NotificationEmails: set of string, optional
	NotificationEmails terra.SetValue[terra.StringValue] `hcl:"notification_emails,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ProjectId: string, required
	ProjectId terra.StringValue `hcl:"project_id,attr" validate:"required"`
	// EnrolledServices: min=1
	EnrolledServices []projectaccessapprovalsettings.EnrolledServices `hcl:"enrolled_services,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *projectaccessapprovalsettings.Timeouts `hcl:"timeouts,block"`
}
type projectAccessApprovalSettingsAttributes struct {
	ref terra.Reference
}

// ActiveKeyVersion returns a reference to field active_key_version of google_project_access_approval_settings.
func (paas projectAccessApprovalSettingsAttributes) ActiveKeyVersion() terra.StringValue {
	return terra.ReferenceAsString(paas.ref.Append("active_key_version"))
}

// AncestorHasActiveKeyVersion returns a reference to field ancestor_has_active_key_version of google_project_access_approval_settings.
func (paas projectAccessApprovalSettingsAttributes) AncestorHasActiveKeyVersion() terra.BoolValue {
	return terra.ReferenceAsBool(paas.ref.Append("ancestor_has_active_key_version"))
}

// EnrolledAncestor returns a reference to field enrolled_ancestor of google_project_access_approval_settings.
func (paas projectAccessApprovalSettingsAttributes) EnrolledAncestor() terra.BoolValue {
	return terra.ReferenceAsBool(paas.ref.Append("enrolled_ancestor"))
}

// Id returns a reference to field id of google_project_access_approval_settings.
func (paas projectAccessApprovalSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(paas.ref.Append("id"))
}

// InvalidKeyVersion returns a reference to field invalid_key_version of google_project_access_approval_settings.
func (paas projectAccessApprovalSettingsAttributes) InvalidKeyVersion() terra.BoolValue {
	return terra.ReferenceAsBool(paas.ref.Append("invalid_key_version"))
}

// Name returns a reference to field name of google_project_access_approval_settings.
func (paas projectAccessApprovalSettingsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(paas.ref.Append("name"))
}

// NotificationEmails returns a reference to field notification_emails of google_project_access_approval_settings.
func (paas projectAccessApprovalSettingsAttributes) NotificationEmails() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](paas.ref.Append("notification_emails"))
}

// Project returns a reference to field project of google_project_access_approval_settings.
func (paas projectAccessApprovalSettingsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(paas.ref.Append("project"))
}

// ProjectId returns a reference to field project_id of google_project_access_approval_settings.
func (paas projectAccessApprovalSettingsAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(paas.ref.Append("project_id"))
}

func (paas projectAccessApprovalSettingsAttributes) EnrolledServices() terra.SetValue[projectaccessapprovalsettings.EnrolledServicesAttributes] {
	return terra.ReferenceAsSet[projectaccessapprovalsettings.EnrolledServicesAttributes](paas.ref.Append("enrolled_services"))
}

func (paas projectAccessApprovalSettingsAttributes) Timeouts() projectaccessapprovalsettings.TimeoutsAttributes {
	return terra.ReferenceAsSingle[projectaccessapprovalsettings.TimeoutsAttributes](paas.ref.Append("timeouts"))
}

type projectAccessApprovalSettingsState struct {
	ActiveKeyVersion            string                                                `json:"active_key_version"`
	AncestorHasActiveKeyVersion bool                                                  `json:"ancestor_has_active_key_version"`
	EnrolledAncestor            bool                                                  `json:"enrolled_ancestor"`
	Id                          string                                                `json:"id"`
	InvalidKeyVersion           bool                                                  `json:"invalid_key_version"`
	Name                        string                                                `json:"name"`
	NotificationEmails          []string                                              `json:"notification_emails"`
	Project                     string                                                `json:"project"`
	ProjectId                   string                                                `json:"project_id"`
	EnrolledServices            []projectaccessapprovalsettings.EnrolledServicesState `json:"enrolled_services"`
	Timeouts                    *projectaccessapprovalsettings.TimeoutsState          `json:"timeouts"`
}
