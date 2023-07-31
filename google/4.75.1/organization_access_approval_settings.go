// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	organizationaccessapprovalsettings "github.com/golingon/terraproviders/google/4.75.1/organizationaccessapprovalsettings"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOrganizationAccessApprovalSettings creates a new instance of [OrganizationAccessApprovalSettings].
func NewOrganizationAccessApprovalSettings(name string, args OrganizationAccessApprovalSettingsArgs) *OrganizationAccessApprovalSettings {
	return &OrganizationAccessApprovalSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrganizationAccessApprovalSettings)(nil)

// OrganizationAccessApprovalSettings represents the Terraform resource google_organization_access_approval_settings.
type OrganizationAccessApprovalSettings struct {
	Name      string
	Args      OrganizationAccessApprovalSettingsArgs
	state     *organizationAccessApprovalSettingsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OrganizationAccessApprovalSettings].
func (oaas *OrganizationAccessApprovalSettings) Type() string {
	return "google_organization_access_approval_settings"
}

// LocalName returns the local name for [OrganizationAccessApprovalSettings].
func (oaas *OrganizationAccessApprovalSettings) LocalName() string {
	return oaas.Name
}

// Configuration returns the configuration (args) for [OrganizationAccessApprovalSettings].
func (oaas *OrganizationAccessApprovalSettings) Configuration() interface{} {
	return oaas.Args
}

// DependOn is used for other resources to depend on [OrganizationAccessApprovalSettings].
func (oaas *OrganizationAccessApprovalSettings) DependOn() terra.Reference {
	return terra.ReferenceResource(oaas)
}

// Dependencies returns the list of resources [OrganizationAccessApprovalSettings] depends_on.
func (oaas *OrganizationAccessApprovalSettings) Dependencies() terra.Dependencies {
	return oaas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OrganizationAccessApprovalSettings].
func (oaas *OrganizationAccessApprovalSettings) LifecycleManagement() *terra.Lifecycle {
	return oaas.Lifecycle
}

// Attributes returns the attributes for [OrganizationAccessApprovalSettings].
func (oaas *OrganizationAccessApprovalSettings) Attributes() organizationAccessApprovalSettingsAttributes {
	return organizationAccessApprovalSettingsAttributes{ref: terra.ReferenceResource(oaas)}
}

// ImportState imports the given attribute values into [OrganizationAccessApprovalSettings]'s state.
func (oaas *OrganizationAccessApprovalSettings) ImportState(av io.Reader) error {
	oaas.state = &organizationAccessApprovalSettingsState{}
	if err := json.NewDecoder(av).Decode(oaas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oaas.Type(), oaas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OrganizationAccessApprovalSettings] has state.
func (oaas *OrganizationAccessApprovalSettings) State() (*organizationAccessApprovalSettingsState, bool) {
	return oaas.state, oaas.state != nil
}

// StateMust returns the state for [OrganizationAccessApprovalSettings]. Panics if the state is nil.
func (oaas *OrganizationAccessApprovalSettings) StateMust() *organizationAccessApprovalSettingsState {
	if oaas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oaas.Type(), oaas.LocalName()))
	}
	return oaas.state
}

// OrganizationAccessApprovalSettingsArgs contains the configurations for google_organization_access_approval_settings.
type OrganizationAccessApprovalSettingsArgs struct {
	// ActiveKeyVersion: string, optional
	ActiveKeyVersion terra.StringValue `hcl:"active_key_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NotificationEmails: set of string, optional
	NotificationEmails terra.SetValue[terra.StringValue] `hcl:"notification_emails,attr"`
	// OrganizationId: string, required
	OrganizationId terra.StringValue `hcl:"organization_id,attr" validate:"required"`
	// EnrolledServices: min=1
	EnrolledServices []organizationaccessapprovalsettings.EnrolledServices `hcl:"enrolled_services,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *organizationaccessapprovalsettings.Timeouts `hcl:"timeouts,block"`
}
type organizationAccessApprovalSettingsAttributes struct {
	ref terra.Reference
}

// ActiveKeyVersion returns a reference to field active_key_version of google_organization_access_approval_settings.
func (oaas organizationAccessApprovalSettingsAttributes) ActiveKeyVersion() terra.StringValue {
	return terra.ReferenceAsString(oaas.ref.Append("active_key_version"))
}

// AncestorHasActiveKeyVersion returns a reference to field ancestor_has_active_key_version of google_organization_access_approval_settings.
func (oaas organizationAccessApprovalSettingsAttributes) AncestorHasActiveKeyVersion() terra.BoolValue {
	return terra.ReferenceAsBool(oaas.ref.Append("ancestor_has_active_key_version"))
}

// EnrolledAncestor returns a reference to field enrolled_ancestor of google_organization_access_approval_settings.
func (oaas organizationAccessApprovalSettingsAttributes) EnrolledAncestor() terra.BoolValue {
	return terra.ReferenceAsBool(oaas.ref.Append("enrolled_ancestor"))
}

// Id returns a reference to field id of google_organization_access_approval_settings.
func (oaas organizationAccessApprovalSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oaas.ref.Append("id"))
}

// InvalidKeyVersion returns a reference to field invalid_key_version of google_organization_access_approval_settings.
func (oaas organizationAccessApprovalSettingsAttributes) InvalidKeyVersion() terra.BoolValue {
	return terra.ReferenceAsBool(oaas.ref.Append("invalid_key_version"))
}

// Name returns a reference to field name of google_organization_access_approval_settings.
func (oaas organizationAccessApprovalSettingsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(oaas.ref.Append("name"))
}

// NotificationEmails returns a reference to field notification_emails of google_organization_access_approval_settings.
func (oaas organizationAccessApprovalSettingsAttributes) NotificationEmails() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oaas.ref.Append("notification_emails"))
}

// OrganizationId returns a reference to field organization_id of google_organization_access_approval_settings.
func (oaas organizationAccessApprovalSettingsAttributes) OrganizationId() terra.StringValue {
	return terra.ReferenceAsString(oaas.ref.Append("organization_id"))
}

func (oaas organizationAccessApprovalSettingsAttributes) EnrolledServices() terra.SetValue[organizationaccessapprovalsettings.EnrolledServicesAttributes] {
	return terra.ReferenceAsSet[organizationaccessapprovalsettings.EnrolledServicesAttributes](oaas.ref.Append("enrolled_services"))
}

func (oaas organizationAccessApprovalSettingsAttributes) Timeouts() organizationaccessapprovalsettings.TimeoutsAttributes {
	return terra.ReferenceAsSingle[organizationaccessapprovalsettings.TimeoutsAttributes](oaas.ref.Append("timeouts"))
}

type organizationAccessApprovalSettingsState struct {
	ActiveKeyVersion            string                                                     `json:"active_key_version"`
	AncestorHasActiveKeyVersion bool                                                       `json:"ancestor_has_active_key_version"`
	EnrolledAncestor            bool                                                       `json:"enrolled_ancestor"`
	Id                          string                                                     `json:"id"`
	InvalidKeyVersion           bool                                                       `json:"invalid_key_version"`
	Name                        string                                                     `json:"name"`
	NotificationEmails          []string                                                   `json:"notification_emails"`
	OrganizationId              string                                                     `json:"organization_id"`
	EnrolledServices            []organizationaccessapprovalsettings.EnrolledServicesState `json:"enrolled_services"`
	Timeouts                    *organizationaccessapprovalsettings.TimeoutsState          `json:"timeouts"`
}
