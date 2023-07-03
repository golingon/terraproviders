// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAuditmanagerOrganizationAdminAccountRegistration creates a new instance of [AuditmanagerOrganizationAdminAccountRegistration].
func NewAuditmanagerOrganizationAdminAccountRegistration(name string, args AuditmanagerOrganizationAdminAccountRegistrationArgs) *AuditmanagerOrganizationAdminAccountRegistration {
	return &AuditmanagerOrganizationAdminAccountRegistration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AuditmanagerOrganizationAdminAccountRegistration)(nil)

// AuditmanagerOrganizationAdminAccountRegistration represents the Terraform resource aws_auditmanager_organization_admin_account_registration.
type AuditmanagerOrganizationAdminAccountRegistration struct {
	Name      string
	Args      AuditmanagerOrganizationAdminAccountRegistrationArgs
	state     *auditmanagerOrganizationAdminAccountRegistrationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AuditmanagerOrganizationAdminAccountRegistration].
func (aoaar *AuditmanagerOrganizationAdminAccountRegistration) Type() string {
	return "aws_auditmanager_organization_admin_account_registration"
}

// LocalName returns the local name for [AuditmanagerOrganizationAdminAccountRegistration].
func (aoaar *AuditmanagerOrganizationAdminAccountRegistration) LocalName() string {
	return aoaar.Name
}

// Configuration returns the configuration (args) for [AuditmanagerOrganizationAdminAccountRegistration].
func (aoaar *AuditmanagerOrganizationAdminAccountRegistration) Configuration() interface{} {
	return aoaar.Args
}

// DependOn is used for other resources to depend on [AuditmanagerOrganizationAdminAccountRegistration].
func (aoaar *AuditmanagerOrganizationAdminAccountRegistration) DependOn() terra.Reference {
	return terra.ReferenceResource(aoaar)
}

// Dependencies returns the list of resources [AuditmanagerOrganizationAdminAccountRegistration] depends_on.
func (aoaar *AuditmanagerOrganizationAdminAccountRegistration) Dependencies() terra.Dependencies {
	return aoaar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AuditmanagerOrganizationAdminAccountRegistration].
func (aoaar *AuditmanagerOrganizationAdminAccountRegistration) LifecycleManagement() *terra.Lifecycle {
	return aoaar.Lifecycle
}

// Attributes returns the attributes for [AuditmanagerOrganizationAdminAccountRegistration].
func (aoaar *AuditmanagerOrganizationAdminAccountRegistration) Attributes() auditmanagerOrganizationAdminAccountRegistrationAttributes {
	return auditmanagerOrganizationAdminAccountRegistrationAttributes{ref: terra.ReferenceResource(aoaar)}
}

// ImportState imports the given attribute values into [AuditmanagerOrganizationAdminAccountRegistration]'s state.
func (aoaar *AuditmanagerOrganizationAdminAccountRegistration) ImportState(av io.Reader) error {
	aoaar.state = &auditmanagerOrganizationAdminAccountRegistrationState{}
	if err := json.NewDecoder(av).Decode(aoaar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aoaar.Type(), aoaar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AuditmanagerOrganizationAdminAccountRegistration] has state.
func (aoaar *AuditmanagerOrganizationAdminAccountRegistration) State() (*auditmanagerOrganizationAdminAccountRegistrationState, bool) {
	return aoaar.state, aoaar.state != nil
}

// StateMust returns the state for [AuditmanagerOrganizationAdminAccountRegistration]. Panics if the state is nil.
func (aoaar *AuditmanagerOrganizationAdminAccountRegistration) StateMust() *auditmanagerOrganizationAdminAccountRegistrationState {
	if aoaar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aoaar.Type(), aoaar.LocalName()))
	}
	return aoaar.state
}

// AuditmanagerOrganizationAdminAccountRegistrationArgs contains the configurations for aws_auditmanager_organization_admin_account_registration.
type AuditmanagerOrganizationAdminAccountRegistrationArgs struct {
	// AdminAccountId: string, required
	AdminAccountId terra.StringValue `hcl:"admin_account_id,attr" validate:"required"`
}
type auditmanagerOrganizationAdminAccountRegistrationAttributes struct {
	ref terra.Reference
}

// AdminAccountId returns a reference to field admin_account_id of aws_auditmanager_organization_admin_account_registration.
func (aoaar auditmanagerOrganizationAdminAccountRegistrationAttributes) AdminAccountId() terra.StringValue {
	return terra.ReferenceAsString(aoaar.ref.Append("admin_account_id"))
}

// Id returns a reference to field id of aws_auditmanager_organization_admin_account_registration.
func (aoaar auditmanagerOrganizationAdminAccountRegistrationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aoaar.ref.Append("id"))
}

// OrganizationId returns a reference to field organization_id of aws_auditmanager_organization_admin_account_registration.
func (aoaar auditmanagerOrganizationAdminAccountRegistrationAttributes) OrganizationId() terra.StringValue {
	return terra.ReferenceAsString(aoaar.ref.Append("organization_id"))
}

type auditmanagerOrganizationAdminAccountRegistrationState struct {
	AdminAccountId string `json:"admin_account_id"`
	Id             string `json:"id"`
	OrganizationId string `json:"organization_id"`
}
