// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_msk_cluster_policy

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

// Resource represents the Terraform resource aws_msk_cluster_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *awsMskClusterPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (amcp *Resource) Type() string {
	return "aws_msk_cluster_policy"
}

// LocalName returns the local name for [Resource].
func (amcp *Resource) LocalName() string {
	return amcp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (amcp *Resource) Configuration() interface{} {
	return amcp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (amcp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(amcp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (amcp *Resource) Dependencies() terra.Dependencies {
	return amcp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (amcp *Resource) LifecycleManagement() *terra.Lifecycle {
	return amcp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (amcp *Resource) Attributes() awsMskClusterPolicyAttributes {
	return awsMskClusterPolicyAttributes{ref: terra.ReferenceResource(amcp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (amcp *Resource) ImportState(state io.Reader) error {
	amcp.state = &awsMskClusterPolicyState{}
	if err := json.NewDecoder(state).Decode(amcp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amcp.Type(), amcp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (amcp *Resource) State() (*awsMskClusterPolicyState, bool) {
	return amcp.state, amcp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (amcp *Resource) StateMust() *awsMskClusterPolicyState {
	if amcp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amcp.Type(), amcp.LocalName()))
	}
	return amcp.state
}

// Args contains the configurations for aws_msk_cluster_policy.
type Args struct {
	// ClusterArn: string, required
	ClusterArn terra.StringValue `hcl:"cluster_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
}

type awsMskClusterPolicyAttributes struct {
	ref terra.Reference
}

// ClusterArn returns a reference to field cluster_arn of aws_msk_cluster_policy.
func (amcp awsMskClusterPolicyAttributes) ClusterArn() terra.StringValue {
	return terra.ReferenceAsString(amcp.ref.Append("cluster_arn"))
}

// CurrentVersion returns a reference to field current_version of aws_msk_cluster_policy.
func (amcp awsMskClusterPolicyAttributes) CurrentVersion() terra.StringValue {
	return terra.ReferenceAsString(amcp.ref.Append("current_version"))
}

// Id returns a reference to field id of aws_msk_cluster_policy.
func (amcp awsMskClusterPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amcp.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_msk_cluster_policy.
func (amcp awsMskClusterPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(amcp.ref.Append("policy"))
}

type awsMskClusterPolicyState struct {
	ClusterArn     string `json:"cluster_arn"`
	CurrentVersion string `json:"current_version"`
	Id             string `json:"id"`
	Policy         string `json:"policy"`
}
