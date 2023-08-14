// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	route53recoverycontrolconfigcluster "github.com/golingon/terraproviders/aws/5.12.0/route53recoverycontrolconfigcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoute53RecoverycontrolconfigCluster creates a new instance of [Route53RecoverycontrolconfigCluster].
func NewRoute53RecoverycontrolconfigCluster(name string, args Route53RecoverycontrolconfigClusterArgs) *Route53RecoverycontrolconfigCluster {
	return &Route53RecoverycontrolconfigCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53RecoverycontrolconfigCluster)(nil)

// Route53RecoverycontrolconfigCluster represents the Terraform resource aws_route53recoverycontrolconfig_cluster.
type Route53RecoverycontrolconfigCluster struct {
	Name      string
	Args      Route53RecoverycontrolconfigClusterArgs
	state     *route53RecoverycontrolconfigClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53RecoverycontrolconfigCluster].
func (rc *Route53RecoverycontrolconfigCluster) Type() string {
	return "aws_route53recoverycontrolconfig_cluster"
}

// LocalName returns the local name for [Route53RecoverycontrolconfigCluster].
func (rc *Route53RecoverycontrolconfigCluster) LocalName() string {
	return rc.Name
}

// Configuration returns the configuration (args) for [Route53RecoverycontrolconfigCluster].
func (rc *Route53RecoverycontrolconfigCluster) Configuration() interface{} {
	return rc.Args
}

// DependOn is used for other resources to depend on [Route53RecoverycontrolconfigCluster].
func (rc *Route53RecoverycontrolconfigCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(rc)
}

// Dependencies returns the list of resources [Route53RecoverycontrolconfigCluster] depends_on.
func (rc *Route53RecoverycontrolconfigCluster) Dependencies() terra.Dependencies {
	return rc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53RecoverycontrolconfigCluster].
func (rc *Route53RecoverycontrolconfigCluster) LifecycleManagement() *terra.Lifecycle {
	return rc.Lifecycle
}

// Attributes returns the attributes for [Route53RecoverycontrolconfigCluster].
func (rc *Route53RecoverycontrolconfigCluster) Attributes() route53RecoverycontrolconfigClusterAttributes {
	return route53RecoverycontrolconfigClusterAttributes{ref: terra.ReferenceResource(rc)}
}

// ImportState imports the given attribute values into [Route53RecoverycontrolconfigCluster]'s state.
func (rc *Route53RecoverycontrolconfigCluster) ImportState(av io.Reader) error {
	rc.state = &route53RecoverycontrolconfigClusterState{}
	if err := json.NewDecoder(av).Decode(rc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rc.Type(), rc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53RecoverycontrolconfigCluster] has state.
func (rc *Route53RecoverycontrolconfigCluster) State() (*route53RecoverycontrolconfigClusterState, bool) {
	return rc.state, rc.state != nil
}

// StateMust returns the state for [Route53RecoverycontrolconfigCluster]. Panics if the state is nil.
func (rc *Route53RecoverycontrolconfigCluster) StateMust() *route53RecoverycontrolconfigClusterState {
	if rc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rc.Type(), rc.LocalName()))
	}
	return rc.state
}

// Route53RecoverycontrolconfigClusterArgs contains the configurations for aws_route53recoverycontrolconfig_cluster.
type Route53RecoverycontrolconfigClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ClusterEndpoints: min=0
	ClusterEndpoints []route53recoverycontrolconfigcluster.ClusterEndpoints `hcl:"cluster_endpoints,block" validate:"min=0"`
}
type route53RecoverycontrolconfigClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53recoverycontrolconfig_cluster.
func (rc route53RecoverycontrolconfigClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("arn"))
}

// Id returns a reference to field id of aws_route53recoverycontrolconfig_cluster.
func (rc route53RecoverycontrolconfigClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("id"))
}

// Name returns a reference to field name of aws_route53recoverycontrolconfig_cluster.
func (rc route53RecoverycontrolconfigClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("name"))
}

// Status returns a reference to field status of aws_route53recoverycontrolconfig_cluster.
func (rc route53RecoverycontrolconfigClusterAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("status"))
}

func (rc route53RecoverycontrolconfigClusterAttributes) ClusterEndpoints() terra.ListValue[route53recoverycontrolconfigcluster.ClusterEndpointsAttributes] {
	return terra.ReferenceAsList[route53recoverycontrolconfigcluster.ClusterEndpointsAttributes](rc.ref.Append("cluster_endpoints"))
}

type route53RecoverycontrolconfigClusterState struct {
	Arn              string                                                      `json:"arn"`
	Id               string                                                      `json:"id"`
	Name             string                                                      `json:"name"`
	Status           string                                                      `json:"status"`
	ClusterEndpoints []route53recoverycontrolconfigcluster.ClusterEndpointsState `json:"cluster_endpoints"`
}