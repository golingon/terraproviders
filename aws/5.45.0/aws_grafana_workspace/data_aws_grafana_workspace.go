// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_grafana_workspace

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_grafana_workspace.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (agw *DataSource) DataSource() string {
	return "aws_grafana_workspace"
}

// LocalName returns the local name for [DataSource].
func (agw *DataSource) LocalName() string {
	return agw.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (agw *DataSource) Configuration() interface{} {
	return agw.Args
}

// Attributes returns the attributes for [DataSource].
func (agw *DataSource) Attributes() dataAwsGrafanaWorkspaceAttributes {
	return dataAwsGrafanaWorkspaceAttributes{ref: terra.ReferenceDataSource(agw)}
}

// DataArgs contains the configurations for aws_grafana_workspace.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
}

type dataAwsGrafanaWorkspaceAttributes struct {
	ref terra.Reference
}

// AccountAccessType returns a reference to field account_access_type of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) AccountAccessType() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("account_access_type"))
}

// Arn returns a reference to field arn of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("arn"))
}

// AuthenticationProviders returns a reference to field authentication_providers of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) AuthenticationProviders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](agw.ref.Append("authentication_providers"))
}

// CreatedDate returns a reference to field created_date of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("created_date"))
}

// DataSources returns a reference to field data_sources of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) DataSources() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](agw.ref.Append("data_sources"))
}

// Description returns a reference to field description of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("description"))
}

// Endpoint returns a reference to field endpoint of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("endpoint"))
}

// GrafanaVersion returns a reference to field grafana_version of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) GrafanaVersion() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("grafana_version"))
}

// Id returns a reference to field id of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("last_updated_date"))
}

// Name returns a reference to field name of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("name"))
}

// NotificationDestinations returns a reference to field notification_destinations of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) NotificationDestinations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](agw.ref.Append("notification_destinations"))
}

// OrganizationRoleName returns a reference to field organization_role_name of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) OrganizationRoleName() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("organization_role_name"))
}

// OrganizationalUnits returns a reference to field organizational_units of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) OrganizationalUnits() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](agw.ref.Append("organizational_units"))
}

// PermissionType returns a reference to field permission_type of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) PermissionType() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("permission_type"))
}

// RoleArn returns a reference to field role_arn of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("role_arn"))
}

// SamlConfigurationStatus returns a reference to field saml_configuration_status of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) SamlConfigurationStatus() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("saml_configuration_status"))
}

// StackSetName returns a reference to field stack_set_name of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) StackSetName() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("stack_set_name"))
}

// Status returns a reference to field status of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agw.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of aws_grafana_workspace.
func (agw dataAwsGrafanaWorkspaceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(agw.ref.Append("workspace_id"))
}
