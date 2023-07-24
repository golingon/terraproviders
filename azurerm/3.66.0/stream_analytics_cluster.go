// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticscluster "github.com/golingon/terraproviders/azurerm/3.66.0/streamanalyticscluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsCluster creates a new instance of [StreamAnalyticsCluster].
func NewStreamAnalyticsCluster(name string, args StreamAnalyticsClusterArgs) *StreamAnalyticsCluster {
	return &StreamAnalyticsCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsCluster)(nil)

// StreamAnalyticsCluster represents the Terraform resource azurerm_stream_analytics_cluster.
type StreamAnalyticsCluster struct {
	Name      string
	Args      StreamAnalyticsClusterArgs
	state     *streamAnalyticsClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsCluster].
func (sac *StreamAnalyticsCluster) Type() string {
	return "azurerm_stream_analytics_cluster"
}

// LocalName returns the local name for [StreamAnalyticsCluster].
func (sac *StreamAnalyticsCluster) LocalName() string {
	return sac.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsCluster].
func (sac *StreamAnalyticsCluster) Configuration() interface{} {
	return sac.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsCluster].
func (sac *StreamAnalyticsCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(sac)
}

// Dependencies returns the list of resources [StreamAnalyticsCluster] depends_on.
func (sac *StreamAnalyticsCluster) Dependencies() terra.Dependencies {
	return sac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsCluster].
func (sac *StreamAnalyticsCluster) LifecycleManagement() *terra.Lifecycle {
	return sac.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsCluster].
func (sac *StreamAnalyticsCluster) Attributes() streamAnalyticsClusterAttributes {
	return streamAnalyticsClusterAttributes{ref: terra.ReferenceResource(sac)}
}

// ImportState imports the given attribute values into [StreamAnalyticsCluster]'s state.
func (sac *StreamAnalyticsCluster) ImportState(av io.Reader) error {
	sac.state = &streamAnalyticsClusterState{}
	if err := json.NewDecoder(av).Decode(sac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sac.Type(), sac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsCluster] has state.
func (sac *StreamAnalyticsCluster) State() (*streamAnalyticsClusterState, bool) {
	return sac.state, sac.state != nil
}

// StateMust returns the state for [StreamAnalyticsCluster]. Panics if the state is nil.
func (sac *StreamAnalyticsCluster) StateMust() *streamAnalyticsClusterState {
	if sac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sac.Type(), sac.LocalName()))
	}
	return sac.state
}

// StreamAnalyticsClusterArgs contains the configurations for azurerm_stream_analytics_cluster.
type StreamAnalyticsClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StreamingCapacity: number, required
	StreamingCapacity terra.NumberValue `hcl:"streaming_capacity,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *streamanalyticscluster.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsClusterAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_stream_analytics_cluster.
func (sac streamAnalyticsClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sac.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_stream_analytics_cluster.
func (sac streamAnalyticsClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sac.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_stream_analytics_cluster.
func (sac streamAnalyticsClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sac.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_cluster.
func (sac streamAnalyticsClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sac.ref.Append("resource_group_name"))
}

// StreamingCapacity returns a reference to field streaming_capacity of azurerm_stream_analytics_cluster.
func (sac streamAnalyticsClusterAttributes) StreamingCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(sac.ref.Append("streaming_capacity"))
}

// Tags returns a reference to field tags of azurerm_stream_analytics_cluster.
func (sac streamAnalyticsClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sac.ref.Append("tags"))
}

func (sac streamAnalyticsClusterAttributes) Timeouts() streamanalyticscluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticscluster.TimeoutsAttributes](sac.ref.Append("timeouts"))
}

type streamAnalyticsClusterState struct {
	Id                string                                `json:"id"`
	Location          string                                `json:"location"`
	Name              string                                `json:"name"`
	ResourceGroupName string                                `json:"resource_group_name"`
	StreamingCapacity float64                               `json:"streaming_capacity"`
	Tags              map[string]string                     `json:"tags"`
	Timeouts          *streamanalyticscluster.TimeoutsState `json:"timeouts"`
}
