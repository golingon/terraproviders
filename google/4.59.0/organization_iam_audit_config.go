// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	organizationiamauditconfig "github.com/golingon/terraproviders/google/4.59.0/organizationiamauditconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewOrganizationIamAuditConfig(name string, args OrganizationIamAuditConfigArgs) *OrganizationIamAuditConfig {
	return &OrganizationIamAuditConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrganizationIamAuditConfig)(nil)

type OrganizationIamAuditConfig struct {
	Name  string
	Args  OrganizationIamAuditConfigArgs
	state *organizationIamAuditConfigState
}

func (oiac *OrganizationIamAuditConfig) Type() string {
	return "google_organization_iam_audit_config"
}

func (oiac *OrganizationIamAuditConfig) LocalName() string {
	return oiac.Name
}

func (oiac *OrganizationIamAuditConfig) Configuration() interface{} {
	return oiac.Args
}

func (oiac *OrganizationIamAuditConfig) Attributes() organizationIamAuditConfigAttributes {
	return organizationIamAuditConfigAttributes{ref: terra.ReferenceResource(oiac)}
}

func (oiac *OrganizationIamAuditConfig) ImportState(av io.Reader) error {
	oiac.state = &organizationIamAuditConfigState{}
	if err := json.NewDecoder(av).Decode(oiac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oiac.Type(), oiac.LocalName(), err)
	}
	return nil
}

func (oiac *OrganizationIamAuditConfig) State() (*organizationIamAuditConfigState, bool) {
	return oiac.state, oiac.state != nil
}

func (oiac *OrganizationIamAuditConfig) StateMust() *organizationIamAuditConfigState {
	if oiac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oiac.Type(), oiac.LocalName()))
	}
	return oiac.state
}

func (oiac *OrganizationIamAuditConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(oiac)
}

type OrganizationIamAuditConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// AuditLogConfig: min=1
	AuditLogConfig []organizationiamauditconfig.AuditLogConfig `hcl:"audit_log_config,block" validate:"min=1"`
	// DependsOn contains resources that OrganizationIamAuditConfig depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type organizationIamAuditConfigAttributes struct {
	ref terra.Reference
}

func (oiac organizationIamAuditConfigAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(oiac.ref.Append("etag"))
}

func (oiac organizationIamAuditConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceString(oiac.ref.Append("id"))
}

func (oiac organizationIamAuditConfigAttributes) OrgId() terra.StringValue {
	return terra.ReferenceString(oiac.ref.Append("org_id"))
}

func (oiac organizationIamAuditConfigAttributes) Service() terra.StringValue {
	return terra.ReferenceString(oiac.ref.Append("service"))
}

func (oiac organizationIamAuditConfigAttributes) AuditLogConfig() terra.SetValue[organizationiamauditconfig.AuditLogConfigAttributes] {
	return terra.ReferenceSet[organizationiamauditconfig.AuditLogConfigAttributes](oiac.ref.Append("audit_log_config"))
}

type organizationIamAuditConfigState struct {
	Etag           string                                           `json:"etag"`
	Id             string                                           `json:"id"`
	OrgId          string                                           `json:"org_id"`
	Service        string                                           `json:"service"`
	AuditLogConfig []organizationiamauditconfig.AuditLogConfigState `json:"audit_log_config"`
}
