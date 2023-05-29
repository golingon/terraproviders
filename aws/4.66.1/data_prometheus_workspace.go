// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataPrometheusWorkspace creates a new instance of [DataPrometheusWorkspace].
func NewDataPrometheusWorkspace(name string, args DataPrometheusWorkspaceArgs) *DataPrometheusWorkspace {
	return &DataPrometheusWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrometheusWorkspace)(nil)

// DataPrometheusWorkspace represents the Terraform data resource aws_prometheus_workspace.
type DataPrometheusWorkspace struct {
	Name string
	Args DataPrometheusWorkspaceArgs
}

// DataSource returns the Terraform object type for [DataPrometheusWorkspace].
func (pw *DataPrometheusWorkspace) DataSource() string {
	return "aws_prometheus_workspace"
}

// LocalName returns the local name for [DataPrometheusWorkspace].
func (pw *DataPrometheusWorkspace) LocalName() string {
	return pw.Name
}

// Configuration returns the configuration (args) for [DataPrometheusWorkspace].
func (pw *DataPrometheusWorkspace) Configuration() interface{} {
	return pw.Args
}

// Attributes returns the attributes for [DataPrometheusWorkspace].
func (pw *DataPrometheusWorkspace) Attributes() dataPrometheusWorkspaceAttributes {
	return dataPrometheusWorkspaceAttributes{ref: terra.ReferenceDataResource(pw)}
}

// DataPrometheusWorkspaceArgs contains the configurations for aws_prometheus_workspace.
type DataPrometheusWorkspaceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
}
type dataPrometheusWorkspaceAttributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of aws_prometheus_workspace.
func (pw dataPrometheusWorkspaceAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(pw.ref.Append("alias"))
}

// Arn returns a reference to field arn of aws_prometheus_workspace.
func (pw dataPrometheusWorkspaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(pw.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_prometheus_workspace.
func (pw dataPrometheusWorkspaceAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(pw.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_prometheus_workspace.
func (pw dataPrometheusWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pw.ref.Append("id"))
}

// PrometheusEndpoint returns a reference to field prometheus_endpoint of aws_prometheus_workspace.
func (pw dataPrometheusWorkspaceAttributes) PrometheusEndpoint() terra.StringValue {
	return terra.ReferenceAsString(pw.ref.Append("prometheus_endpoint"))
}

// Status returns a reference to field status of aws_prometheus_workspace.
func (pw dataPrometheusWorkspaceAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(pw.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_prometheus_workspace.
func (pw dataPrometheusWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pw.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of aws_prometheus_workspace.
func (pw dataPrometheusWorkspaceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(pw.ref.Append("workspace_id"))
}