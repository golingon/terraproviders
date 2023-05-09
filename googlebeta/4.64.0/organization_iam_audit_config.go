// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	organizationiamauditconfig "github.com/golingon/terraproviders/googlebeta/4.64.0/organizationiamauditconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOrganizationIamAuditConfig creates a new instance of [OrganizationIamAuditConfig].
func NewOrganizationIamAuditConfig(name string, args OrganizationIamAuditConfigArgs) *OrganizationIamAuditConfig {
	return &OrganizationIamAuditConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrganizationIamAuditConfig)(nil)

// OrganizationIamAuditConfig represents the Terraform resource google_organization_iam_audit_config.
type OrganizationIamAuditConfig struct {
	Name      string
	Args      OrganizationIamAuditConfigArgs
	state     *organizationIamAuditConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OrganizationIamAuditConfig].
func (oiac *OrganizationIamAuditConfig) Type() string {
	return "google_organization_iam_audit_config"
}

// LocalName returns the local name for [OrganizationIamAuditConfig].
func (oiac *OrganizationIamAuditConfig) LocalName() string {
	return oiac.Name
}

// Configuration returns the configuration (args) for [OrganizationIamAuditConfig].
func (oiac *OrganizationIamAuditConfig) Configuration() interface{} {
	return oiac.Args
}

// DependOn is used for other resources to depend on [OrganizationIamAuditConfig].
func (oiac *OrganizationIamAuditConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(oiac)
}

// Dependencies returns the list of resources [OrganizationIamAuditConfig] depends_on.
func (oiac *OrganizationIamAuditConfig) Dependencies() terra.Dependencies {
	return oiac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OrganizationIamAuditConfig].
func (oiac *OrganizationIamAuditConfig) LifecycleManagement() *terra.Lifecycle {
	return oiac.Lifecycle
}

// Attributes returns the attributes for [OrganizationIamAuditConfig].
func (oiac *OrganizationIamAuditConfig) Attributes() organizationIamAuditConfigAttributes {
	return organizationIamAuditConfigAttributes{ref: terra.ReferenceResource(oiac)}
}

// ImportState imports the given attribute values into [OrganizationIamAuditConfig]'s state.
func (oiac *OrganizationIamAuditConfig) ImportState(av io.Reader) error {
	oiac.state = &organizationIamAuditConfigState{}
	if err := json.NewDecoder(av).Decode(oiac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oiac.Type(), oiac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OrganizationIamAuditConfig] has state.
func (oiac *OrganizationIamAuditConfig) State() (*organizationIamAuditConfigState, bool) {
	return oiac.state, oiac.state != nil
}

// StateMust returns the state for [OrganizationIamAuditConfig]. Panics if the state is nil.
func (oiac *OrganizationIamAuditConfig) StateMust() *organizationIamAuditConfigState {
	if oiac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oiac.Type(), oiac.LocalName()))
	}
	return oiac.state
}

// OrganizationIamAuditConfigArgs contains the configurations for google_organization_iam_audit_config.
type OrganizationIamAuditConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// AuditLogConfig: min=1
	AuditLogConfig []organizationiamauditconfig.AuditLogConfig `hcl:"audit_log_config,block" validate:"min=1"`
}
type organizationIamAuditConfigAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_organization_iam_audit_config.
func (oiac organizationIamAuditConfigAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(oiac.ref.Append("etag"))
}

// Id returns a reference to field id of google_organization_iam_audit_config.
func (oiac organizationIamAuditConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oiac.ref.Append("id"))
}

// OrgId returns a reference to field org_id of google_organization_iam_audit_config.
func (oiac organizationIamAuditConfigAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(oiac.ref.Append("org_id"))
}

// Service returns a reference to field service of google_organization_iam_audit_config.
func (oiac organizationIamAuditConfigAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(oiac.ref.Append("service"))
}

func (oiac organizationIamAuditConfigAttributes) AuditLogConfig() terra.SetValue[organizationiamauditconfig.AuditLogConfigAttributes] {
	return terra.ReferenceAsSet[organizationiamauditconfig.AuditLogConfigAttributes](oiac.ref.Append("audit_log_config"))
}

type organizationIamAuditConfigState struct {
	Etag           string                                           `json:"etag"`
	Id             string                                           `json:"id"`
	OrgId          string                                           `json:"org_id"`
	Service        string                                           `json:"service"`
	AuditLogConfig []organizationiamauditconfig.AuditLogConfigState `json:"audit_log_config"`
}
