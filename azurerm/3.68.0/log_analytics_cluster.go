// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	loganalyticscluster "github.com/golingon/terraproviders/azurerm/3.68.0/loganalyticscluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogAnalyticsCluster creates a new instance of [LogAnalyticsCluster].
func NewLogAnalyticsCluster(name string, args LogAnalyticsClusterArgs) *LogAnalyticsCluster {
	return &LogAnalyticsCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogAnalyticsCluster)(nil)

// LogAnalyticsCluster represents the Terraform resource azurerm_log_analytics_cluster.
type LogAnalyticsCluster struct {
	Name      string
	Args      LogAnalyticsClusterArgs
	state     *logAnalyticsClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogAnalyticsCluster].
func (lac *LogAnalyticsCluster) Type() string {
	return "azurerm_log_analytics_cluster"
}

// LocalName returns the local name for [LogAnalyticsCluster].
func (lac *LogAnalyticsCluster) LocalName() string {
	return lac.Name
}

// Configuration returns the configuration (args) for [LogAnalyticsCluster].
func (lac *LogAnalyticsCluster) Configuration() interface{} {
	return lac.Args
}

// DependOn is used for other resources to depend on [LogAnalyticsCluster].
func (lac *LogAnalyticsCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(lac)
}

// Dependencies returns the list of resources [LogAnalyticsCluster] depends_on.
func (lac *LogAnalyticsCluster) Dependencies() terra.Dependencies {
	return lac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogAnalyticsCluster].
func (lac *LogAnalyticsCluster) LifecycleManagement() *terra.Lifecycle {
	return lac.Lifecycle
}

// Attributes returns the attributes for [LogAnalyticsCluster].
func (lac *LogAnalyticsCluster) Attributes() logAnalyticsClusterAttributes {
	return logAnalyticsClusterAttributes{ref: terra.ReferenceResource(lac)}
}

// ImportState imports the given attribute values into [LogAnalyticsCluster]'s state.
func (lac *LogAnalyticsCluster) ImportState(av io.Reader) error {
	lac.state = &logAnalyticsClusterState{}
	if err := json.NewDecoder(av).Decode(lac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lac.Type(), lac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogAnalyticsCluster] has state.
func (lac *LogAnalyticsCluster) State() (*logAnalyticsClusterState, bool) {
	return lac.state, lac.state != nil
}

// StateMust returns the state for [LogAnalyticsCluster]. Panics if the state is nil.
func (lac *LogAnalyticsCluster) StateMust() *logAnalyticsClusterState {
	if lac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lac.Type(), lac.LocalName()))
	}
	return lac.state
}

// LogAnalyticsClusterArgs contains the configurations for azurerm_log_analytics_cluster.
type LogAnalyticsClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SizeGb: number, optional
	SizeGb terra.NumberValue `hcl:"size_gb,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: required
	Identity *loganalyticscluster.Identity `hcl:"identity,block" validate:"required"`
	// Timeouts: optional
	Timeouts *loganalyticscluster.Timeouts `hcl:"timeouts,block"`
}
type logAnalyticsClusterAttributes struct {
	ref terra.Reference
}

// ClusterId returns a reference to field cluster_id of azurerm_log_analytics_cluster.
func (lac logAnalyticsClusterAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(lac.ref.Append("cluster_id"))
}

// Id returns a reference to field id of azurerm_log_analytics_cluster.
func (lac logAnalyticsClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lac.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_log_analytics_cluster.
func (lac logAnalyticsClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lac.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_log_analytics_cluster.
func (lac logAnalyticsClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lac.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_log_analytics_cluster.
func (lac logAnalyticsClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lac.ref.Append("resource_group_name"))
}

// SizeGb returns a reference to field size_gb of azurerm_log_analytics_cluster.
func (lac logAnalyticsClusterAttributes) SizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(lac.ref.Append("size_gb"))
}

// Tags returns a reference to field tags of azurerm_log_analytics_cluster.
func (lac logAnalyticsClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lac.ref.Append("tags"))
}

func (lac logAnalyticsClusterAttributes) Identity() terra.ListValue[loganalyticscluster.IdentityAttributes] {
	return terra.ReferenceAsList[loganalyticscluster.IdentityAttributes](lac.ref.Append("identity"))
}

func (lac logAnalyticsClusterAttributes) Timeouts() loganalyticscluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[loganalyticscluster.TimeoutsAttributes](lac.ref.Append("timeouts"))
}

type logAnalyticsClusterState struct {
	ClusterId         string                              `json:"cluster_id"`
	Id                string                              `json:"id"`
	Location          string                              `json:"location"`
	Name              string                              `json:"name"`
	ResourceGroupName string                              `json:"resource_group_name"`
	SizeGb            float64                             `json:"size_gb"`
	Tags              map[string]string                   `json:"tags"`
	Identity          []loganalyticscluster.IdentityState `json:"identity"`
	Timeouts          *loganalyticscluster.TimeoutsState  `json:"timeouts"`
}
