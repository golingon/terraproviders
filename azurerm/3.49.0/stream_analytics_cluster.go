// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticscluster "github.com/golingon/terraproviders/azurerm/3.49.0/streamanalyticscluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewStreamAnalyticsCluster(name string, args StreamAnalyticsClusterArgs) *StreamAnalyticsCluster {
	return &StreamAnalyticsCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsCluster)(nil)

type StreamAnalyticsCluster struct {
	Name  string
	Args  StreamAnalyticsClusterArgs
	state *streamAnalyticsClusterState
}

func (sac *StreamAnalyticsCluster) Type() string {
	return "azurerm_stream_analytics_cluster"
}

func (sac *StreamAnalyticsCluster) LocalName() string {
	return sac.Name
}

func (sac *StreamAnalyticsCluster) Configuration() interface{} {
	return sac.Args
}

func (sac *StreamAnalyticsCluster) Attributes() streamAnalyticsClusterAttributes {
	return streamAnalyticsClusterAttributes{ref: terra.ReferenceResource(sac)}
}

func (sac *StreamAnalyticsCluster) ImportState(av io.Reader) error {
	sac.state = &streamAnalyticsClusterState{}
	if err := json.NewDecoder(av).Decode(sac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sac.Type(), sac.LocalName(), err)
	}
	return nil
}

func (sac *StreamAnalyticsCluster) State() (*streamAnalyticsClusterState, bool) {
	return sac.state, sac.state != nil
}

func (sac *StreamAnalyticsCluster) StateMust() *streamAnalyticsClusterState {
	if sac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sac.Type(), sac.LocalName()))
	}
	return sac.state
}

func (sac *StreamAnalyticsCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(sac)
}

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
	// DependsOn contains resources that StreamAnalyticsCluster depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type streamAnalyticsClusterAttributes struct {
	ref terra.Reference
}

func (sac streamAnalyticsClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sac.ref.Append("id"))
}

func (sac streamAnalyticsClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceString(sac.ref.Append("location"))
}

func (sac streamAnalyticsClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sac.ref.Append("name"))
}

func (sac streamAnalyticsClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(sac.ref.Append("resource_group_name"))
}

func (sac streamAnalyticsClusterAttributes) StreamingCapacity() terra.NumberValue {
	return terra.ReferenceNumber(sac.ref.Append("streaming_capacity"))
}

func (sac streamAnalyticsClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sac.ref.Append("tags"))
}

func (sac streamAnalyticsClusterAttributes) Timeouts() streamanalyticscluster.TimeoutsAttributes {
	return terra.ReferenceSingle[streamanalyticscluster.TimeoutsAttributes](sac.ref.Append("timeouts"))
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
