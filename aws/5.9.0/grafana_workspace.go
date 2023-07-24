// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	grafanaworkspace "github.com/golingon/terraproviders/aws/5.9.0/grafanaworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGrafanaWorkspace creates a new instance of [GrafanaWorkspace].
func NewGrafanaWorkspace(name string, args GrafanaWorkspaceArgs) *GrafanaWorkspace {
	return &GrafanaWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GrafanaWorkspace)(nil)

// GrafanaWorkspace represents the Terraform resource aws_grafana_workspace.
type GrafanaWorkspace struct {
	Name      string
	Args      GrafanaWorkspaceArgs
	state     *grafanaWorkspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GrafanaWorkspace].
func (gw *GrafanaWorkspace) Type() string {
	return "aws_grafana_workspace"
}

// LocalName returns the local name for [GrafanaWorkspace].
func (gw *GrafanaWorkspace) LocalName() string {
	return gw.Name
}

// Configuration returns the configuration (args) for [GrafanaWorkspace].
func (gw *GrafanaWorkspace) Configuration() interface{} {
	return gw.Args
}

// DependOn is used for other resources to depend on [GrafanaWorkspace].
func (gw *GrafanaWorkspace) DependOn() terra.Reference {
	return terra.ReferenceResource(gw)
}

// Dependencies returns the list of resources [GrafanaWorkspace] depends_on.
func (gw *GrafanaWorkspace) Dependencies() terra.Dependencies {
	return gw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GrafanaWorkspace].
func (gw *GrafanaWorkspace) LifecycleManagement() *terra.Lifecycle {
	return gw.Lifecycle
}

// Attributes returns the attributes for [GrafanaWorkspace].
func (gw *GrafanaWorkspace) Attributes() grafanaWorkspaceAttributes {
	return grafanaWorkspaceAttributes{ref: terra.ReferenceResource(gw)}
}

// ImportState imports the given attribute values into [GrafanaWorkspace]'s state.
func (gw *GrafanaWorkspace) ImportState(av io.Reader) error {
	gw.state = &grafanaWorkspaceState{}
	if err := json.NewDecoder(av).Decode(gw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gw.Type(), gw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GrafanaWorkspace] has state.
func (gw *GrafanaWorkspace) State() (*grafanaWorkspaceState, bool) {
	return gw.state, gw.state != nil
}

// StateMust returns the state for [GrafanaWorkspace]. Panics if the state is nil.
func (gw *GrafanaWorkspace) StateMust() *grafanaWorkspaceState {
	if gw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gw.Type(), gw.LocalName()))
	}
	return gw.state
}

// GrafanaWorkspaceArgs contains the configurations for aws_grafana_workspace.
type GrafanaWorkspaceArgs struct {
	// AccountAccessType: string, required
	AccountAccessType terra.StringValue `hcl:"account_access_type,attr" validate:"required"`
	// AuthenticationProviders: list of string, required
	AuthenticationProviders terra.ListValue[terra.StringValue] `hcl:"authentication_providers,attr" validate:"required"`
	// Configuration: string, optional
	Configuration terra.StringValue `hcl:"configuration,attr"`
	// DataSources: list of string, optional
	DataSources terra.ListValue[terra.StringValue] `hcl:"data_sources,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// GrafanaVersion: string, optional
	GrafanaVersion terra.StringValue `hcl:"grafana_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NotificationDestinations: list of string, optional
	NotificationDestinations terra.ListValue[terra.StringValue] `hcl:"notification_destinations,attr"`
	// OrganizationRoleName: string, optional
	OrganizationRoleName terra.StringValue `hcl:"organization_role_name,attr"`
	// OrganizationalUnits: list of string, optional
	OrganizationalUnits terra.ListValue[terra.StringValue] `hcl:"organizational_units,attr"`
	// PermissionType: string, required
	PermissionType terra.StringValue `hcl:"permission_type,attr" validate:"required"`
	// RoleArn: string, optional
	RoleArn terra.StringValue `hcl:"role_arn,attr"`
	// StackSetName: string, optional
	StackSetName terra.StringValue `hcl:"stack_set_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// NetworkAccessControl: optional
	NetworkAccessControl *grafanaworkspace.NetworkAccessControl `hcl:"network_access_control,block"`
	// Timeouts: optional
	Timeouts *grafanaworkspace.Timeouts `hcl:"timeouts,block"`
	// VpcConfiguration: optional
	VpcConfiguration *grafanaworkspace.VpcConfiguration `hcl:"vpc_configuration,block"`
}
type grafanaWorkspaceAttributes struct {
	ref terra.Reference
}

// AccountAccessType returns a reference to field account_access_type of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) AccountAccessType() terra.StringValue {
	return terra.ReferenceAsString(gw.ref.Append("account_access_type"))
}

// Arn returns a reference to field arn of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(gw.ref.Append("arn"))
}

// AuthenticationProviders returns a reference to field authentication_providers of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) AuthenticationProviders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gw.ref.Append("authentication_providers"))
}

// Configuration returns a reference to field configuration of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) Configuration() terra.StringValue {
	return terra.ReferenceAsString(gw.ref.Append("configuration"))
}

// DataSources returns a reference to field data_sources of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) DataSources() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gw.ref.Append("data_sources"))
}

// Description returns a reference to field description of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gw.ref.Append("description"))
}

// Endpoint returns a reference to field endpoint of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(gw.ref.Append("endpoint"))
}

// GrafanaVersion returns a reference to field grafana_version of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) GrafanaVersion() terra.StringValue {
	return terra.ReferenceAsString(gw.ref.Append("grafana_version"))
}

// Id returns a reference to field id of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gw.ref.Append("id"))
}

// Name returns a reference to field name of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gw.ref.Append("name"))
}

// NotificationDestinations returns a reference to field notification_destinations of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) NotificationDestinations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gw.ref.Append("notification_destinations"))
}

// OrganizationRoleName returns a reference to field organization_role_name of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) OrganizationRoleName() terra.StringValue {
	return terra.ReferenceAsString(gw.ref.Append("organization_role_name"))
}

// OrganizationalUnits returns a reference to field organizational_units of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) OrganizationalUnits() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gw.ref.Append("organizational_units"))
}

// PermissionType returns a reference to field permission_type of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) PermissionType() terra.StringValue {
	return terra.ReferenceAsString(gw.ref.Append("permission_type"))
}

// RoleArn returns a reference to field role_arn of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(gw.ref.Append("role_arn"))
}

// SamlConfigurationStatus returns a reference to field saml_configuration_status of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) SamlConfigurationStatus() terra.StringValue {
	return terra.ReferenceAsString(gw.ref.Append("saml_configuration_status"))
}

// StackSetName returns a reference to field stack_set_name of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) StackSetName() terra.StringValue {
	return terra.ReferenceAsString(gw.ref.Append("stack_set_name"))
}

// Tags returns a reference to field tags of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gw.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_grafana_workspace.
func (gw grafanaWorkspaceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gw.ref.Append("tags_all"))
}

func (gw grafanaWorkspaceAttributes) NetworkAccessControl() terra.ListValue[grafanaworkspace.NetworkAccessControlAttributes] {
	return terra.ReferenceAsList[grafanaworkspace.NetworkAccessControlAttributes](gw.ref.Append("network_access_control"))
}

func (gw grafanaWorkspaceAttributes) Timeouts() grafanaworkspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[grafanaworkspace.TimeoutsAttributes](gw.ref.Append("timeouts"))
}

func (gw grafanaWorkspaceAttributes) VpcConfiguration() terra.ListValue[grafanaworkspace.VpcConfigurationAttributes] {
	return terra.ReferenceAsList[grafanaworkspace.VpcConfigurationAttributes](gw.ref.Append("vpc_configuration"))
}

type grafanaWorkspaceState struct {
	AccountAccessType        string                                       `json:"account_access_type"`
	Arn                      string                                       `json:"arn"`
	AuthenticationProviders  []string                                     `json:"authentication_providers"`
	Configuration            string                                       `json:"configuration"`
	DataSources              []string                                     `json:"data_sources"`
	Description              string                                       `json:"description"`
	Endpoint                 string                                       `json:"endpoint"`
	GrafanaVersion           string                                       `json:"grafana_version"`
	Id                       string                                       `json:"id"`
	Name                     string                                       `json:"name"`
	NotificationDestinations []string                                     `json:"notification_destinations"`
	OrganizationRoleName     string                                       `json:"organization_role_name"`
	OrganizationalUnits      []string                                     `json:"organizational_units"`
	PermissionType           string                                       `json:"permission_type"`
	RoleArn                  string                                       `json:"role_arn"`
	SamlConfigurationStatus  string                                       `json:"saml_configuration_status"`
	StackSetName             string                                       `json:"stack_set_name"`
	Tags                     map[string]string                            `json:"tags"`
	TagsAll                  map[string]string                            `json:"tags_all"`
	NetworkAccessControl     []grafanaworkspace.NetworkAccessControlState `json:"network_access_control"`
	Timeouts                 *grafanaworkspace.TimeoutsState              `json:"timeouts"`
	VpcConfiguration         []grafanaworkspace.VpcConfigurationState     `json:"vpc_configuration"`
}
