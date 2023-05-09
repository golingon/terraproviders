// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	prometheusworkspace "github.com/golingon/terraproviders/aws/4.66.1/prometheusworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrometheusWorkspace creates a new instance of [PrometheusWorkspace].
func NewPrometheusWorkspace(name string, args PrometheusWorkspaceArgs) *PrometheusWorkspace {
	return &PrometheusWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrometheusWorkspace)(nil)

// PrometheusWorkspace represents the Terraform resource aws_prometheus_workspace.
type PrometheusWorkspace struct {
	Name      string
	Args      PrometheusWorkspaceArgs
	state     *prometheusWorkspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrometheusWorkspace].
func (pw *PrometheusWorkspace) Type() string {
	return "aws_prometheus_workspace"
}

// LocalName returns the local name for [PrometheusWorkspace].
func (pw *PrometheusWorkspace) LocalName() string {
	return pw.Name
}

// Configuration returns the configuration (args) for [PrometheusWorkspace].
func (pw *PrometheusWorkspace) Configuration() interface{} {
	return pw.Args
}

// DependOn is used for other resources to depend on [PrometheusWorkspace].
func (pw *PrometheusWorkspace) DependOn() terra.Reference {
	return terra.ReferenceResource(pw)
}

// Dependencies returns the list of resources [PrometheusWorkspace] depends_on.
func (pw *PrometheusWorkspace) Dependencies() terra.Dependencies {
	return pw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrometheusWorkspace].
func (pw *PrometheusWorkspace) LifecycleManagement() *terra.Lifecycle {
	return pw.Lifecycle
}

// Attributes returns the attributes for [PrometheusWorkspace].
func (pw *PrometheusWorkspace) Attributes() prometheusWorkspaceAttributes {
	return prometheusWorkspaceAttributes{ref: terra.ReferenceResource(pw)}
}

// ImportState imports the given attribute values into [PrometheusWorkspace]'s state.
func (pw *PrometheusWorkspace) ImportState(av io.Reader) error {
	pw.state = &prometheusWorkspaceState{}
	if err := json.NewDecoder(av).Decode(pw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pw.Type(), pw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrometheusWorkspace] has state.
func (pw *PrometheusWorkspace) State() (*prometheusWorkspaceState, bool) {
	return pw.state, pw.state != nil
}

// StateMust returns the state for [PrometheusWorkspace]. Panics if the state is nil.
func (pw *PrometheusWorkspace) StateMust() *prometheusWorkspaceState {
	if pw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pw.Type(), pw.LocalName()))
	}
	return pw.state
}

// PrometheusWorkspaceArgs contains the configurations for aws_prometheus_workspace.
type PrometheusWorkspaceArgs struct {
	// Alias: string, optional
	Alias terra.StringValue `hcl:"alias,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// LoggingConfiguration: optional
	LoggingConfiguration *prometheusworkspace.LoggingConfiguration `hcl:"logging_configuration,block"`
}
type prometheusWorkspaceAttributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of aws_prometheus_workspace.
func (pw prometheusWorkspaceAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(pw.ref.Append("alias"))
}

// Arn returns a reference to field arn of aws_prometheus_workspace.
func (pw prometheusWorkspaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(pw.ref.Append("arn"))
}

// Id returns a reference to field id of aws_prometheus_workspace.
func (pw prometheusWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pw.ref.Append("id"))
}

// PrometheusEndpoint returns a reference to field prometheus_endpoint of aws_prometheus_workspace.
func (pw prometheusWorkspaceAttributes) PrometheusEndpoint() terra.StringValue {
	return terra.ReferenceAsString(pw.ref.Append("prometheus_endpoint"))
}

// Tags returns a reference to field tags of aws_prometheus_workspace.
func (pw prometheusWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pw.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_prometheus_workspace.
func (pw prometheusWorkspaceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pw.ref.Append("tags_all"))
}

func (pw prometheusWorkspaceAttributes) LoggingConfiguration() terra.ListValue[prometheusworkspace.LoggingConfigurationAttributes] {
	return terra.ReferenceAsList[prometheusworkspace.LoggingConfigurationAttributes](pw.ref.Append("logging_configuration"))
}

type prometheusWorkspaceState struct {
	Alias                string                                          `json:"alias"`
	Arn                  string                                          `json:"arn"`
	Id                   string                                          `json:"id"`
	PrometheusEndpoint   string                                          `json:"prometheus_endpoint"`
	Tags                 map[string]string                               `json:"tags"`
	TagsAll              map[string]string                               `json:"tags_all"`
	LoggingConfiguration []prometheusworkspace.LoggingConfigurationState `json:"logging_configuration"`
}
