// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	grafanaworkspacesamlconfiguration "github.com/golingon/terraproviders/aws/5.0.1/grafanaworkspacesamlconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGrafanaWorkspaceSamlConfiguration creates a new instance of [GrafanaWorkspaceSamlConfiguration].
func NewGrafanaWorkspaceSamlConfiguration(name string, args GrafanaWorkspaceSamlConfigurationArgs) *GrafanaWorkspaceSamlConfiguration {
	return &GrafanaWorkspaceSamlConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GrafanaWorkspaceSamlConfiguration)(nil)

// GrafanaWorkspaceSamlConfiguration represents the Terraform resource aws_grafana_workspace_saml_configuration.
type GrafanaWorkspaceSamlConfiguration struct {
	Name      string
	Args      GrafanaWorkspaceSamlConfigurationArgs
	state     *grafanaWorkspaceSamlConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GrafanaWorkspaceSamlConfiguration].
func (gwsc *GrafanaWorkspaceSamlConfiguration) Type() string {
	return "aws_grafana_workspace_saml_configuration"
}

// LocalName returns the local name for [GrafanaWorkspaceSamlConfiguration].
func (gwsc *GrafanaWorkspaceSamlConfiguration) LocalName() string {
	return gwsc.Name
}

// Configuration returns the configuration (args) for [GrafanaWorkspaceSamlConfiguration].
func (gwsc *GrafanaWorkspaceSamlConfiguration) Configuration() interface{} {
	return gwsc.Args
}

// DependOn is used for other resources to depend on [GrafanaWorkspaceSamlConfiguration].
func (gwsc *GrafanaWorkspaceSamlConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(gwsc)
}

// Dependencies returns the list of resources [GrafanaWorkspaceSamlConfiguration] depends_on.
func (gwsc *GrafanaWorkspaceSamlConfiguration) Dependencies() terra.Dependencies {
	return gwsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GrafanaWorkspaceSamlConfiguration].
func (gwsc *GrafanaWorkspaceSamlConfiguration) LifecycleManagement() *terra.Lifecycle {
	return gwsc.Lifecycle
}

// Attributes returns the attributes for [GrafanaWorkspaceSamlConfiguration].
func (gwsc *GrafanaWorkspaceSamlConfiguration) Attributes() grafanaWorkspaceSamlConfigurationAttributes {
	return grafanaWorkspaceSamlConfigurationAttributes{ref: terra.ReferenceResource(gwsc)}
}

// ImportState imports the given attribute values into [GrafanaWorkspaceSamlConfiguration]'s state.
func (gwsc *GrafanaWorkspaceSamlConfiguration) ImportState(av io.Reader) error {
	gwsc.state = &grafanaWorkspaceSamlConfigurationState{}
	if err := json.NewDecoder(av).Decode(gwsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gwsc.Type(), gwsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GrafanaWorkspaceSamlConfiguration] has state.
func (gwsc *GrafanaWorkspaceSamlConfiguration) State() (*grafanaWorkspaceSamlConfigurationState, bool) {
	return gwsc.state, gwsc.state != nil
}

// StateMust returns the state for [GrafanaWorkspaceSamlConfiguration]. Panics if the state is nil.
func (gwsc *GrafanaWorkspaceSamlConfiguration) StateMust() *grafanaWorkspaceSamlConfigurationState {
	if gwsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gwsc.Type(), gwsc.LocalName()))
	}
	return gwsc.state
}

// GrafanaWorkspaceSamlConfigurationArgs contains the configurations for aws_grafana_workspace_saml_configuration.
type GrafanaWorkspaceSamlConfigurationArgs struct {
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
	Timeouts *grafanaworkspacesamlconfiguration.Timeouts `hcl:"timeouts,block"`
}
type grafanaWorkspaceSamlConfigurationAttributes struct {
	ref terra.Reference
}

// AdminRoleValues returns a reference to field admin_role_values of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) AdminRoleValues() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gwsc.ref.Append("admin_role_values"))
}

// AllowedOrganizations returns a reference to field allowed_organizations of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) AllowedOrganizations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gwsc.ref.Append("allowed_organizations"))
}

// EditorRoleValues returns a reference to field editor_role_values of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) EditorRoleValues() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gwsc.ref.Append("editor_role_values"))
}

// EmailAssertion returns a reference to field email_assertion of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) EmailAssertion() terra.StringValue {
	return terra.ReferenceAsString(gwsc.ref.Append("email_assertion"))
}

// GroupsAssertion returns a reference to field groups_assertion of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) GroupsAssertion() terra.StringValue {
	return terra.ReferenceAsString(gwsc.ref.Append("groups_assertion"))
}

// Id returns a reference to field id of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gwsc.ref.Append("id"))
}

// IdpMetadataUrl returns a reference to field idp_metadata_url of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) IdpMetadataUrl() terra.StringValue {
	return terra.ReferenceAsString(gwsc.ref.Append("idp_metadata_url"))
}

// IdpMetadataXml returns a reference to field idp_metadata_xml of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) IdpMetadataXml() terra.StringValue {
	return terra.ReferenceAsString(gwsc.ref.Append("idp_metadata_xml"))
}

// LoginAssertion returns a reference to field login_assertion of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) LoginAssertion() terra.StringValue {
	return terra.ReferenceAsString(gwsc.ref.Append("login_assertion"))
}

// LoginValidityDuration returns a reference to field login_validity_duration of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) LoginValidityDuration() terra.NumberValue {
	return terra.ReferenceAsNumber(gwsc.ref.Append("login_validity_duration"))
}

// NameAssertion returns a reference to field name_assertion of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) NameAssertion() terra.StringValue {
	return terra.ReferenceAsString(gwsc.ref.Append("name_assertion"))
}

// OrgAssertion returns a reference to field org_assertion of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) OrgAssertion() terra.StringValue {
	return terra.ReferenceAsString(gwsc.ref.Append("org_assertion"))
}

// RoleAssertion returns a reference to field role_assertion of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) RoleAssertion() terra.StringValue {
	return terra.ReferenceAsString(gwsc.ref.Append("role_assertion"))
}

// Status returns a reference to field status of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(gwsc.ref.Append("status"))
}

// WorkspaceId returns a reference to field workspace_id of aws_grafana_workspace_saml_configuration.
func (gwsc grafanaWorkspaceSamlConfigurationAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(gwsc.ref.Append("workspace_id"))
}

func (gwsc grafanaWorkspaceSamlConfigurationAttributes) Timeouts() grafanaworkspacesamlconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[grafanaworkspacesamlconfiguration.TimeoutsAttributes](gwsc.ref.Append("timeouts"))
}

type grafanaWorkspaceSamlConfigurationState struct {
	AdminRoleValues       []string                                         `json:"admin_role_values"`
	AllowedOrganizations  []string                                         `json:"allowed_organizations"`
	EditorRoleValues      []string                                         `json:"editor_role_values"`
	EmailAssertion        string                                           `json:"email_assertion"`
	GroupsAssertion       string                                           `json:"groups_assertion"`
	Id                    string                                           `json:"id"`
	IdpMetadataUrl        string                                           `json:"idp_metadata_url"`
	IdpMetadataXml        string                                           `json:"idp_metadata_xml"`
	LoginAssertion        string                                           `json:"login_assertion"`
	LoginValidityDuration float64                                          `json:"login_validity_duration"`
	NameAssertion         string                                           `json:"name_assertion"`
	OrgAssertion          string                                           `json:"org_assertion"`
	RoleAssertion         string                                           `json:"role_assertion"`
	Status                string                                           `json:"status"`
	WorkspaceId           string                                           `json:"workspace_id"`
	Timeouts              *grafanaworkspacesamlconfiguration.TimeoutsState `json:"timeouts"`
}
