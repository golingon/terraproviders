// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_prometheus_rule_group_namespace

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

// Resource represents the Terraform resource aws_prometheus_rule_group_namespace.
type Resource struct {
	Name      string
	Args      Args
	state     *awsPrometheusRuleGroupNamespaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aprgn *Resource) Type() string {
	return "aws_prometheus_rule_group_namespace"
}

// LocalName returns the local name for [Resource].
func (aprgn *Resource) LocalName() string {
	return aprgn.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aprgn *Resource) Configuration() interface{} {
	return aprgn.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aprgn *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aprgn)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aprgn *Resource) Dependencies() terra.Dependencies {
	return aprgn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aprgn *Resource) LifecycleManagement() *terra.Lifecycle {
	return aprgn.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aprgn *Resource) Attributes() awsPrometheusRuleGroupNamespaceAttributes {
	return awsPrometheusRuleGroupNamespaceAttributes{ref: terra.ReferenceResource(aprgn)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aprgn *Resource) ImportState(state io.Reader) error {
	aprgn.state = &awsPrometheusRuleGroupNamespaceState{}
	if err := json.NewDecoder(state).Decode(aprgn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aprgn.Type(), aprgn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aprgn *Resource) State() (*awsPrometheusRuleGroupNamespaceState, bool) {
	return aprgn.state, aprgn.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aprgn *Resource) StateMust() *awsPrometheusRuleGroupNamespaceState {
	if aprgn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aprgn.Type(), aprgn.LocalName()))
	}
	return aprgn.state
}

// Args contains the configurations for aws_prometheus_rule_group_namespace.
type Args struct {
	// Data: string, required
	Data terra.StringValue `hcl:"data,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
}

type awsPrometheusRuleGroupNamespaceAttributes struct {
	ref terra.Reference
}

// Data returns a reference to field data of aws_prometheus_rule_group_namespace.
func (aprgn awsPrometheusRuleGroupNamespaceAttributes) Data() terra.StringValue {
	return terra.ReferenceAsString(aprgn.ref.Append("data"))
}

// Id returns a reference to field id of aws_prometheus_rule_group_namespace.
func (aprgn awsPrometheusRuleGroupNamespaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aprgn.ref.Append("id"))
}

// Name returns a reference to field name of aws_prometheus_rule_group_namespace.
func (aprgn awsPrometheusRuleGroupNamespaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aprgn.ref.Append("name"))
}

// WorkspaceId returns a reference to field workspace_id of aws_prometheus_rule_group_namespace.
func (aprgn awsPrometheusRuleGroupNamespaceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(aprgn.ref.Append("workspace_id"))
}

type awsPrometheusRuleGroupNamespaceState struct {
	Data        string `json:"data"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	WorkspaceId string `json:"workspace_id"`
}
