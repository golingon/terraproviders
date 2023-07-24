// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsmanagedprivateendpoint "github.com/golingon/terraproviders/azurerm/3.66.0/streamanalyticsmanagedprivateendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsManagedPrivateEndpoint creates a new instance of [StreamAnalyticsManagedPrivateEndpoint].
func NewStreamAnalyticsManagedPrivateEndpoint(name string, args StreamAnalyticsManagedPrivateEndpointArgs) *StreamAnalyticsManagedPrivateEndpoint {
	return &StreamAnalyticsManagedPrivateEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsManagedPrivateEndpoint)(nil)

// StreamAnalyticsManagedPrivateEndpoint represents the Terraform resource azurerm_stream_analytics_managed_private_endpoint.
type StreamAnalyticsManagedPrivateEndpoint struct {
	Name      string
	Args      StreamAnalyticsManagedPrivateEndpointArgs
	state     *streamAnalyticsManagedPrivateEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsManagedPrivateEndpoint].
func (sampe *StreamAnalyticsManagedPrivateEndpoint) Type() string {
	return "azurerm_stream_analytics_managed_private_endpoint"
}

// LocalName returns the local name for [StreamAnalyticsManagedPrivateEndpoint].
func (sampe *StreamAnalyticsManagedPrivateEndpoint) LocalName() string {
	return sampe.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsManagedPrivateEndpoint].
func (sampe *StreamAnalyticsManagedPrivateEndpoint) Configuration() interface{} {
	return sampe.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsManagedPrivateEndpoint].
func (sampe *StreamAnalyticsManagedPrivateEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(sampe)
}

// Dependencies returns the list of resources [StreamAnalyticsManagedPrivateEndpoint] depends_on.
func (sampe *StreamAnalyticsManagedPrivateEndpoint) Dependencies() terra.Dependencies {
	return sampe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsManagedPrivateEndpoint].
func (sampe *StreamAnalyticsManagedPrivateEndpoint) LifecycleManagement() *terra.Lifecycle {
	return sampe.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsManagedPrivateEndpoint].
func (sampe *StreamAnalyticsManagedPrivateEndpoint) Attributes() streamAnalyticsManagedPrivateEndpointAttributes {
	return streamAnalyticsManagedPrivateEndpointAttributes{ref: terra.ReferenceResource(sampe)}
}

// ImportState imports the given attribute values into [StreamAnalyticsManagedPrivateEndpoint]'s state.
func (sampe *StreamAnalyticsManagedPrivateEndpoint) ImportState(av io.Reader) error {
	sampe.state = &streamAnalyticsManagedPrivateEndpointState{}
	if err := json.NewDecoder(av).Decode(sampe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sampe.Type(), sampe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsManagedPrivateEndpoint] has state.
func (sampe *StreamAnalyticsManagedPrivateEndpoint) State() (*streamAnalyticsManagedPrivateEndpointState, bool) {
	return sampe.state, sampe.state != nil
}

// StateMust returns the state for [StreamAnalyticsManagedPrivateEndpoint]. Panics if the state is nil.
func (sampe *StreamAnalyticsManagedPrivateEndpoint) StateMust() *streamAnalyticsManagedPrivateEndpointState {
	if sampe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sampe.Type(), sampe.LocalName()))
	}
	return sampe.state
}

// StreamAnalyticsManagedPrivateEndpointArgs contains the configurations for azurerm_stream_analytics_managed_private_endpoint.
type StreamAnalyticsManagedPrivateEndpointArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StreamAnalyticsClusterName: string, required
	StreamAnalyticsClusterName terra.StringValue `hcl:"stream_analytics_cluster_name,attr" validate:"required"`
	// SubresourceName: string, required
	SubresourceName terra.StringValue `hcl:"subresource_name,attr" validate:"required"`
	// TargetResourceId: string, required
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsmanagedprivateendpoint.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsManagedPrivateEndpointAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_stream_analytics_managed_private_endpoint.
func (sampe streamAnalyticsManagedPrivateEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sampe.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_managed_private_endpoint.
func (sampe streamAnalyticsManagedPrivateEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sampe.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_managed_private_endpoint.
func (sampe streamAnalyticsManagedPrivateEndpointAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sampe.ref.Append("resource_group_name"))
}

// StreamAnalyticsClusterName returns a reference to field stream_analytics_cluster_name of azurerm_stream_analytics_managed_private_endpoint.
func (sampe streamAnalyticsManagedPrivateEndpointAttributes) StreamAnalyticsClusterName() terra.StringValue {
	return terra.ReferenceAsString(sampe.ref.Append("stream_analytics_cluster_name"))
}

// SubresourceName returns a reference to field subresource_name of azurerm_stream_analytics_managed_private_endpoint.
func (sampe streamAnalyticsManagedPrivateEndpointAttributes) SubresourceName() terra.StringValue {
	return terra.ReferenceAsString(sampe.ref.Append("subresource_name"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_stream_analytics_managed_private_endpoint.
func (sampe streamAnalyticsManagedPrivateEndpointAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(sampe.ref.Append("target_resource_id"))
}

func (sampe streamAnalyticsManagedPrivateEndpointAttributes) Timeouts() streamanalyticsmanagedprivateendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsmanagedprivateendpoint.TimeoutsAttributes](sampe.ref.Append("timeouts"))
}

type streamAnalyticsManagedPrivateEndpointState struct {
	Id                         string                                               `json:"id"`
	Name                       string                                               `json:"name"`
	ResourceGroupName          string                                               `json:"resource_group_name"`
	StreamAnalyticsClusterName string                                               `json:"stream_analytics_cluster_name"`
	SubresourceName            string                                               `json:"subresource_name"`
	TargetResourceId           string                                               `json:"target_resource_id"`
	Timeouts                   *streamanalyticsmanagedprivateendpoint.TimeoutsState `json:"timeouts"`
}
