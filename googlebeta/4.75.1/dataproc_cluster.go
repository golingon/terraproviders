// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataproccluster "github.com/golingon/terraproviders/googlebeta/4.75.1/dataproccluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataprocCluster creates a new instance of [DataprocCluster].
func NewDataprocCluster(name string, args DataprocClusterArgs) *DataprocCluster {
	return &DataprocCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocCluster)(nil)

// DataprocCluster represents the Terraform resource google_dataproc_cluster.
type DataprocCluster struct {
	Name      string
	Args      DataprocClusterArgs
	state     *dataprocClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocCluster].
func (dc *DataprocCluster) Type() string {
	return "google_dataproc_cluster"
}

// LocalName returns the local name for [DataprocCluster].
func (dc *DataprocCluster) LocalName() string {
	return dc.Name
}

// Configuration returns the configuration (args) for [DataprocCluster].
func (dc *DataprocCluster) Configuration() interface{} {
	return dc.Args
}

// DependOn is used for other resources to depend on [DataprocCluster].
func (dc *DataprocCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(dc)
}

// Dependencies returns the list of resources [DataprocCluster] depends_on.
func (dc *DataprocCluster) Dependencies() terra.Dependencies {
	return dc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocCluster].
func (dc *DataprocCluster) LifecycleManagement() *terra.Lifecycle {
	return dc.Lifecycle
}

// Attributes returns the attributes for [DataprocCluster].
func (dc *DataprocCluster) Attributes() dataprocClusterAttributes {
	return dataprocClusterAttributes{ref: terra.ReferenceResource(dc)}
}

// ImportState imports the given attribute values into [DataprocCluster]'s state.
func (dc *DataprocCluster) ImportState(av io.Reader) error {
	dc.state = &dataprocClusterState{}
	if err := json.NewDecoder(av).Decode(dc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dc.Type(), dc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocCluster] has state.
func (dc *DataprocCluster) State() (*dataprocClusterState, bool) {
	return dc.state, dc.state != nil
}

// StateMust returns the state for [DataprocCluster]. Panics if the state is nil.
func (dc *DataprocCluster) StateMust() *dataprocClusterState {
	if dc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dc.Type(), dc.LocalName()))
	}
	return dc.state
}

// DataprocClusterArgs contains the configurations for google_dataproc_cluster.
type DataprocClusterArgs struct {
	// GracefulDecommissionTimeout: string, optional
	GracefulDecommissionTimeout terra.StringValue `hcl:"graceful_decommission_timeout,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ClusterConfig: optional
	ClusterConfig *dataproccluster.ClusterConfig `hcl:"cluster_config,block"`
	// Timeouts: optional
	Timeouts *dataproccluster.Timeouts `hcl:"timeouts,block"`
	// VirtualClusterConfig: optional
	VirtualClusterConfig *dataproccluster.VirtualClusterConfig `hcl:"virtual_cluster_config,block"`
}
type dataprocClusterAttributes struct {
	ref terra.Reference
}

// GracefulDecommissionTimeout returns a reference to field graceful_decommission_timeout of google_dataproc_cluster.
func (dc dataprocClusterAttributes) GracefulDecommissionTimeout() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("graceful_decommission_timeout"))
}

// Id returns a reference to field id of google_dataproc_cluster.
func (dc dataprocClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("id"))
}

// Labels returns a reference to field labels of google_dataproc_cluster.
func (dc dataprocClusterAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dc.ref.Append("labels"))
}

// Name returns a reference to field name of google_dataproc_cluster.
func (dc dataprocClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("name"))
}

// Project returns a reference to field project of google_dataproc_cluster.
func (dc dataprocClusterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("project"))
}

// Region returns a reference to field region of google_dataproc_cluster.
func (dc dataprocClusterAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("region"))
}

func (dc dataprocClusterAttributes) ClusterConfig() terra.ListValue[dataproccluster.ClusterConfigAttributes] {
	return terra.ReferenceAsList[dataproccluster.ClusterConfigAttributes](dc.ref.Append("cluster_config"))
}

func (dc dataprocClusterAttributes) Timeouts() dataproccluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataproccluster.TimeoutsAttributes](dc.ref.Append("timeouts"))
}

func (dc dataprocClusterAttributes) VirtualClusterConfig() terra.ListValue[dataproccluster.VirtualClusterConfigAttributes] {
	return terra.ReferenceAsList[dataproccluster.VirtualClusterConfigAttributes](dc.ref.Append("virtual_cluster_config"))
}

type dataprocClusterState struct {
	GracefulDecommissionTimeout string                                      `json:"graceful_decommission_timeout"`
	Id                          string                                      `json:"id"`
	Labels                      map[string]string                           `json:"labels"`
	Name                        string                                      `json:"name"`
	Project                     string                                      `json:"project"`
	Region                      string                                      `json:"region"`
	ClusterConfig               []dataproccluster.ClusterConfigState        `json:"cluster_config"`
	Timeouts                    *dataproccluster.TimeoutsState              `json:"timeouts"`
	VirtualClusterConfig        []dataproccluster.VirtualClusterConfigState `json:"virtual_cluster_config"`
}
