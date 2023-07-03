// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	projectiamauditconfig "github.com/golingon/terraproviders/google/4.71.0/projectiamauditconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewProjectIamAuditConfig creates a new instance of [ProjectIamAuditConfig].
func NewProjectIamAuditConfig(name string, args ProjectIamAuditConfigArgs) *ProjectIamAuditConfig {
	return &ProjectIamAuditConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ProjectIamAuditConfig)(nil)

// ProjectIamAuditConfig represents the Terraform resource google_project_iam_audit_config.
type ProjectIamAuditConfig struct {
	Name      string
	Args      ProjectIamAuditConfigArgs
	state     *projectIamAuditConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ProjectIamAuditConfig].
func (piac *ProjectIamAuditConfig) Type() string {
	return "google_project_iam_audit_config"
}

// LocalName returns the local name for [ProjectIamAuditConfig].
func (piac *ProjectIamAuditConfig) LocalName() string {
	return piac.Name
}

// Configuration returns the configuration (args) for [ProjectIamAuditConfig].
func (piac *ProjectIamAuditConfig) Configuration() interface{} {
	return piac.Args
}

// DependOn is used for other resources to depend on [ProjectIamAuditConfig].
func (piac *ProjectIamAuditConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(piac)
}

// Dependencies returns the list of resources [ProjectIamAuditConfig] depends_on.
func (piac *ProjectIamAuditConfig) Dependencies() terra.Dependencies {
	return piac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ProjectIamAuditConfig].
func (piac *ProjectIamAuditConfig) LifecycleManagement() *terra.Lifecycle {
	return piac.Lifecycle
}

// Attributes returns the attributes for [ProjectIamAuditConfig].
func (piac *ProjectIamAuditConfig) Attributes() projectIamAuditConfigAttributes {
	return projectIamAuditConfigAttributes{ref: terra.ReferenceResource(piac)}
}

// ImportState imports the given attribute values into [ProjectIamAuditConfig]'s state.
func (piac *ProjectIamAuditConfig) ImportState(av io.Reader) error {
	piac.state = &projectIamAuditConfigState{}
	if err := json.NewDecoder(av).Decode(piac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", piac.Type(), piac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ProjectIamAuditConfig] has state.
func (piac *ProjectIamAuditConfig) State() (*projectIamAuditConfigState, bool) {
	return piac.state, piac.state != nil
}

// StateMust returns the state for [ProjectIamAuditConfig]. Panics if the state is nil.
func (piac *ProjectIamAuditConfig) StateMust() *projectIamAuditConfigState {
	if piac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", piac.Type(), piac.LocalName()))
	}
	return piac.state
}

// ProjectIamAuditConfigArgs contains the configurations for google_project_iam_audit_config.
type ProjectIamAuditConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// AuditLogConfig: min=1
	AuditLogConfig []projectiamauditconfig.AuditLogConfig `hcl:"audit_log_config,block" validate:"min=1"`
}
type projectIamAuditConfigAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_project_iam_audit_config.
func (piac projectIamAuditConfigAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(piac.ref.Append("etag"))
}

// Id returns a reference to field id of google_project_iam_audit_config.
func (piac projectIamAuditConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(piac.ref.Append("id"))
}

// Project returns a reference to field project of google_project_iam_audit_config.
func (piac projectIamAuditConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(piac.ref.Append("project"))
}

// Service returns a reference to field service of google_project_iam_audit_config.
func (piac projectIamAuditConfigAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(piac.ref.Append("service"))
}

func (piac projectIamAuditConfigAttributes) AuditLogConfig() terra.SetValue[projectiamauditconfig.AuditLogConfigAttributes] {
	return terra.ReferenceAsSet[projectiamauditconfig.AuditLogConfigAttributes](piac.ref.Append("audit_log_config"))
}

type projectIamAuditConfigState struct {
	Etag           string                                      `json:"etag"`
	Id             string                                      `json:"id"`
	Project        string                                      `json:"project"`
	Service        string                                      `json:"service"`
	AuditLogConfig []projectiamauditconfig.AuditLogConfigState `json:"audit_log_config"`
}
