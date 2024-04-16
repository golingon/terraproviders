// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_grafana_workspace_saml_configuration

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

// Resource represents the Terraform resource aws_grafana_workspace_saml_configuration.
type Resource struct {
	Name      string
	Args      Args
	state     *awsGrafanaWorkspaceSamlConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (agwsc *Resource) Type() string {
	return "aws_grafana_workspace_saml_configuration"
}

// LocalName returns the local name for [Resource].
func (agwsc *Resource) LocalName() string {
	return agwsc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (agwsc *Resource) Configuration() interface{} {
	return agwsc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (agwsc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(agwsc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (agwsc *Resource) Dependencies() terra.Dependencies {
	return agwsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (agwsc *Resource) LifecycleManagement() *terra.Lifecycle {
	return agwsc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (agwsc *Resource) Attributes() awsGrafanaWorkspaceSamlConfigurationAttributes {
	return awsGrafanaWorkspaceSamlConfigurationAttributes{ref: terra.ReferenceResource(agwsc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (agwsc *Resource) ImportState(state io.Reader) error {
	agwsc.state = &awsGrafanaWorkspaceSamlConfigurationState{}
	if err := json.NewDecoder(state).Decode(agwsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agwsc.Type(), agwsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (agwsc *Resource) State() (*awsGrafanaWorkspaceSamlConfigurationState, bool) {
	return agwsc.state, agwsc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (agwsc *Resource) StateMust() *awsGrafanaWorkspaceSamlConfigurationState {
	if agwsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agwsc.Type(), agwsc.LocalName()))
	}
	return agwsc.state
}

// Args contains the configurations for aws_grafana_workspace_saml_configuration.
type Args struct {
	// AdminRoleValues: list of string, optional
	AdminRoleValues terra.ListValue[terra.StringValue] `hcl:"admin_role_values,attr"`
	// AllowedOrganizations: list of string, optional
	AllowedOrganizations terra.ListValue[terra.StringValue] `hcl:"allowed_organizations,attr"`
	// EditorRoleValues: list of string, required
	EditorRoleValues terra.ListValue[terra.StringValue] `hcl:"editor_role_values,attr" validate:"required"`
	// EmailAssertion: string, optional
	EmailAssertion terra.StringValue `hcl:"email_assertion,attr"`
	// GroupsAssertion: string, optional
	GroupsAssertion terra.StringValue `hcl:"groups_assertion,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdpMetadataUrl: string, optional
	IdpMetadataUrl terra.StringValue `hcl:"idp_metadata_url,attr"`
	// IdpMetadataXml: string, optional
	IdpMetadataXml terra.StringValue `hcl:"idp_metadata_xml,attr"`
	// LoginAssertion: string, optional
	LoginAssertion terra.StringValue `hcl:"login_assertion,attr"`
	// LoginValidityDuration: number, optional
	LoginValidityDuration terra.NumberValue `hcl:"login_validity_duration,attr"`
	// NameAssertion: string, optional
	NameAssertion terra.StringValue `hcl:"name_assertion,attr"`
	// OrgAssertion: string, optional
	OrgAssertion terra.StringValue `hcl:"org_assertion,attr"`
	// RoleAssertion: string, optional
	RoleAssertion terra.StringValue `hcl:"role_assertion,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsGrafanaWorkspaceSamlConfigurationAttributes struct {
	ref terra.Reference
}

// AdminRoleValues returns a reference to field admin_role_values of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) AdminRoleValues() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](agwsc.ref.Append("admin_role_values"))
}

// AllowedOrganizations returns a reference to field allowed_organizations of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) AllowedOrganizations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](agwsc.ref.Append("allowed_organizations"))
}

// EditorRoleValues returns a reference to field editor_role_values of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) EditorRoleValues() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](agwsc.ref.Append("editor_role_values"))
}

// EmailAssertion returns a reference to field email_assertion of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) EmailAssertion() terra.StringValue {
	return terra.ReferenceAsString(agwsc.ref.Append("email_assertion"))
}

// GroupsAssertion returns a reference to field groups_assertion of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) GroupsAssertion() terra.StringValue {
	return terra.ReferenceAsString(agwsc.ref.Append("groups_assertion"))
}

// Id returns a reference to field id of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agwsc.ref.Append("id"))
}

// IdpMetadataUrl returns a reference to field idp_metadata_url of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) IdpMetadataUrl() terra.StringValue {
	return terra.ReferenceAsString(agwsc.ref.Append("idp_metadata_url"))
}

// IdpMetadataXml returns a reference to field idp_metadata_xml of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) IdpMetadataXml() terra.StringValue {
	return terra.ReferenceAsString(agwsc.ref.Append("idp_metadata_xml"))
}

// LoginAssertion returns a reference to field login_assertion of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) LoginAssertion() terra.StringValue {
	return terra.ReferenceAsString(agwsc.ref.Append("login_assertion"))
}

// LoginValidityDuration returns a reference to field login_validity_duration of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) LoginValidityDuration() terra.NumberValue {
	return terra.ReferenceAsNumber(agwsc.ref.Append("login_validity_duration"))
}

// NameAssertion returns a reference to field name_assertion of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) NameAssertion() terra.StringValue {
	return terra.ReferenceAsString(agwsc.ref.Append("name_assertion"))
}

// OrgAssertion returns a reference to field org_assertion of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) OrgAssertion() terra.StringValue {
	return terra.ReferenceAsString(agwsc.ref.Append("org_assertion"))
}

// RoleAssertion returns a reference to field role_assertion of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) RoleAssertion() terra.StringValue {
	return terra.ReferenceAsString(agwsc.ref.Append("role_assertion"))
}

// Status returns a reference to field status of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(agwsc.ref.Append("status"))
}

// WorkspaceId returns a reference to field workspace_id of aws_grafana_workspace_saml_configuration.
func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(agwsc.ref.Append("workspace_id"))
}

func (agwsc awsGrafanaWorkspaceSamlConfigurationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](agwsc.ref.Append("timeouts"))
}

type awsGrafanaWorkspaceSamlConfigurationState struct {
	AdminRoleValues       []string       `json:"admin_role_values"`
	AllowedOrganizations  []string       `json:"allowed_organizations"`
	EditorRoleValues      []string       `json:"editor_role_values"`
	EmailAssertion        string         `json:"email_assertion"`
	GroupsAssertion       string         `json:"groups_assertion"`
	Id                    string         `json:"id"`
	IdpMetadataUrl        string         `json:"idp_metadata_url"`
	IdpMetadataXml        string         `json:"idp_metadata_xml"`
	LoginAssertion        string         `json:"login_assertion"`
	LoginValidityDuration float64        `json:"login_validity_duration"`
	NameAssertion         string         `json:"name_assertion"`
	OrgAssertion          string         `json:"org_assertion"`
	RoleAssertion         string         `json:"role_assertion"`
	Status                string         `json:"status"`
	WorkspaceId           string         `json:"workspace_id"`
	Timeouts              *TimeoutsState `json:"timeouts"`
}
