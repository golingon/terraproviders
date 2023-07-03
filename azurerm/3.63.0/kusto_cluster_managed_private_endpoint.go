// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	kustoclustermanagedprivateendpoint "github.com/golingon/terraproviders/azurerm/3.63.0/kustoclustermanagedprivateendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKustoClusterManagedPrivateEndpoint creates a new instance of [KustoClusterManagedPrivateEndpoint].
func NewKustoClusterManagedPrivateEndpoint(name string, args KustoClusterManagedPrivateEndpointArgs) *KustoClusterManagedPrivateEndpoint {
	return &KustoClusterManagedPrivateEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KustoClusterManagedPrivateEndpoint)(nil)

// KustoClusterManagedPrivateEndpoint represents the Terraform resource azurerm_kusto_cluster_managed_private_endpoint.
type KustoClusterManagedPrivateEndpoint struct {
	Name      string
	Args      KustoClusterManagedPrivateEndpointArgs
	state     *kustoClusterManagedPrivateEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KustoClusterManagedPrivateEndpoint].
func (kcmpe *KustoClusterManagedPrivateEndpoint) Type() string {
	return "azurerm_kusto_cluster_managed_private_endpoint"
}

// LocalName returns the local name for [KustoClusterManagedPrivateEndpoint].
func (kcmpe *KustoClusterManagedPrivateEndpoint) LocalName() string {
	return kcmpe.Name
}

// Configuration returns the configuration (args) for [KustoClusterManagedPrivateEndpoint].
func (kcmpe *KustoClusterManagedPrivateEndpoint) Configuration() interface{} {
	return kcmpe.Args
}

// DependOn is used for other resources to depend on [KustoClusterManagedPrivateEndpoint].
func (kcmpe *KustoClusterManagedPrivateEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(kcmpe)
}

// Dependencies returns the list of resources [KustoClusterManagedPrivateEndpoint] depends_on.
func (kcmpe *KustoClusterManagedPrivateEndpoint) Dependencies() terra.Dependencies {
	return kcmpe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KustoClusterManagedPrivateEndpoint].
func (kcmpe *KustoClusterManagedPrivateEndpoint) LifecycleManagement() *terra.Lifecycle {
	return kcmpe.Lifecycle
}

// Attributes returns the attributes for [KustoClusterManagedPrivateEndpoint].
func (kcmpe *KustoClusterManagedPrivateEndpoint) Attributes() kustoClusterManagedPrivateEndpointAttributes {
	return kustoClusterManagedPrivateEndpointAttributes{ref: terra.ReferenceResource(kcmpe)}
}

// ImportState imports the given attribute values into [KustoClusterManagedPrivateEndpoint]'s state.
func (kcmpe *KustoClusterManagedPrivateEndpoint) ImportState(av io.Reader) error {
	kcmpe.state = &kustoClusterManagedPrivateEndpointState{}
	if err := json.NewDecoder(av).Decode(kcmpe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kcmpe.Type(), kcmpe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KustoClusterManagedPrivateEndpoint] has state.
func (kcmpe *KustoClusterManagedPrivateEndpoint) State() (*kustoClusterManagedPrivateEndpointState, bool) {
	return kcmpe.state, kcmpe.state != nil
}

// StateMust returns the state for [KustoClusterManagedPrivateEndpoint]. Panics if the state is nil.
func (kcmpe *KustoClusterManagedPrivateEndpoint) StateMust() *kustoClusterManagedPrivateEndpointState {
	if kcmpe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kcmpe.Type(), kcmpe.LocalName()))
	}
	return kcmpe.state
}

// KustoClusterManagedPrivateEndpointArgs contains the configurations for azurerm_kusto_cluster_managed_private_endpoint.
type KustoClusterManagedPrivateEndpointArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// GroupId: string, required
	GroupId terra.StringValue `hcl:"group_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateLinkResourceId: string, required
	PrivateLinkResourceId terra.StringValue `hcl:"private_link_resource_id,attr" validate:"required"`
	// PrivateLinkResourceRegion: string, optional
	PrivateLinkResourceRegion terra.StringValue `hcl:"private_link_resource_region,attr"`
	// RequestMessage: string, optional
	RequestMessage terra.StringValue `hcl:"request_message,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *kustoclustermanagedprivateendpoint.Timeouts `hcl:"timeouts,block"`
}
type kustoClusterManagedPrivateEndpointAttributes struct {
	ref terra.Reference
}

// ClusterName returns a reference to field cluster_name of azurerm_kusto_cluster_managed_private_endpoint.
func (kcmpe kustoClusterManagedPrivateEndpointAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(kcmpe.ref.Append("cluster_name"))
}

// GroupId returns a reference to field group_id of azurerm_kusto_cluster_managed_private_endpoint.
func (kcmpe kustoClusterManagedPrivateEndpointAttributes) GroupId() terra.StringValue {
	return terra.ReferenceAsString(kcmpe.ref.Append("group_id"))
}

// Id returns a reference to field id of azurerm_kusto_cluster_managed_private_endpoint.
func (kcmpe kustoClusterManagedPrivateEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kcmpe.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_kusto_cluster_managed_private_endpoint.
func (kcmpe kustoClusterManagedPrivateEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kcmpe.ref.Append("name"))
}

// PrivateLinkResourceId returns a reference to field private_link_resource_id of azurerm_kusto_cluster_managed_private_endpoint.
func (kcmpe kustoClusterManagedPrivateEndpointAttributes) PrivateLinkResourceId() terra.StringValue {
	return terra.ReferenceAsString(kcmpe.ref.Append("private_link_resource_id"))
}

// PrivateLinkResourceRegion returns a reference to field private_link_resource_region of azurerm_kusto_cluster_managed_private_endpoint.
func (kcmpe kustoClusterManagedPrivateEndpointAttributes) PrivateLinkResourceRegion() terra.StringValue {
	return terra.ReferenceAsString(kcmpe.ref.Append("private_link_resource_region"))
}

// RequestMessage returns a reference to field request_message of azurerm_kusto_cluster_managed_private_endpoint.
func (kcmpe kustoClusterManagedPrivateEndpointAttributes) RequestMessage() terra.StringValue {
	return terra.ReferenceAsString(kcmpe.ref.Append("request_message"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kusto_cluster_managed_private_endpoint.
func (kcmpe kustoClusterManagedPrivateEndpointAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kcmpe.ref.Append("resource_group_name"))
}

func (kcmpe kustoClusterManagedPrivateEndpointAttributes) Timeouts() kustoclustermanagedprivateendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kustoclustermanagedprivateendpoint.TimeoutsAttributes](kcmpe.ref.Append("timeouts"))
}

type kustoClusterManagedPrivateEndpointState struct {
	ClusterName               string                                            `json:"cluster_name"`
	GroupId                   string                                            `json:"group_id"`
	Id                        string                                            `json:"id"`
	Name                      string                                            `json:"name"`
	PrivateLinkResourceId     string                                            `json:"private_link_resource_id"`
	PrivateLinkResourceRegion string                                            `json:"private_link_resource_region"`
	RequestMessage            string                                            `json:"request_message"`
	ResourceGroupName         string                                            `json:"resource_group_name"`
	Timeouts                  *kustoclustermanagedprivateendpoint.TimeoutsState `json:"timeouts"`
}
