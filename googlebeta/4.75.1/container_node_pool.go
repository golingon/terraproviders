// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	containernodepool "github.com/golingon/terraproviders/googlebeta/4.75.1/containernodepool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerNodePool creates a new instance of [ContainerNodePool].
func NewContainerNodePool(name string, args ContainerNodePoolArgs) *ContainerNodePool {
	return &ContainerNodePool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerNodePool)(nil)

// ContainerNodePool represents the Terraform resource google_container_node_pool.
type ContainerNodePool struct {
	Name      string
	Args      ContainerNodePoolArgs
	state     *containerNodePoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerNodePool].
func (cnp *ContainerNodePool) Type() string {
	return "google_container_node_pool"
}

// LocalName returns the local name for [ContainerNodePool].
func (cnp *ContainerNodePool) LocalName() string {
	return cnp.Name
}

// Configuration returns the configuration (args) for [ContainerNodePool].
func (cnp *ContainerNodePool) Configuration() interface{} {
	return cnp.Args
}

// DependOn is used for other resources to depend on [ContainerNodePool].
func (cnp *ContainerNodePool) DependOn() terra.Reference {
	return terra.ReferenceResource(cnp)
}

// Dependencies returns the list of resources [ContainerNodePool] depends_on.
func (cnp *ContainerNodePool) Dependencies() terra.Dependencies {
	return cnp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerNodePool].
func (cnp *ContainerNodePool) LifecycleManagement() *terra.Lifecycle {
	return cnp.Lifecycle
}

// Attributes returns the attributes for [ContainerNodePool].
func (cnp *ContainerNodePool) Attributes() containerNodePoolAttributes {
	return containerNodePoolAttributes{ref: terra.ReferenceResource(cnp)}
}

// ImportState imports the given attribute values into [ContainerNodePool]'s state.
func (cnp *ContainerNodePool) ImportState(av io.Reader) error {
	cnp.state = &containerNodePoolState{}
	if err := json.NewDecoder(av).Decode(cnp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cnp.Type(), cnp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerNodePool] has state.
func (cnp *ContainerNodePool) State() (*containerNodePoolState, bool) {
	return cnp.state, cnp.state != nil
}

// StateMust returns the state for [ContainerNodePool]. Panics if the state is nil.
func (cnp *ContainerNodePool) StateMust() *containerNodePoolState {
	if cnp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cnp.Type(), cnp.LocalName()))
	}
	return cnp.state
}

// ContainerNodePoolArgs contains the configurations for google_container_node_pool.
type ContainerNodePoolArgs struct {
	// Cluster: string, required
	Cluster terra.StringValue `hcl:"cluster,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InitialNodeCount: number, optional
	InitialNodeCount terra.NumberValue `hcl:"initial_node_count,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// MaxPodsPerNode: number, optional
	MaxPodsPerNode terra.NumberValue `hcl:"max_pods_per_node,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// NodeCount: number, optional
	NodeCount terra.NumberValue `hcl:"node_count,attr"`
	// NodeLocations: set of string, optional
	NodeLocations terra.SetValue[terra.StringValue] `hcl:"node_locations,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// Autoscaling: optional
	Autoscaling *containernodepool.Autoscaling `hcl:"autoscaling,block"`
	// Management: optional
	Management *containernodepool.Management `hcl:"management,block"`
	// NetworkConfig: optional
	NetworkConfig *containernodepool.NetworkConfig `hcl:"network_config,block"`
	// NodeConfig: optional
	NodeConfig *containernodepool.NodeConfig `hcl:"node_config,block"`
	// PlacementPolicy: optional
	PlacementPolicy *containernodepool.PlacementPolicy `hcl:"placement_policy,block"`
	// Timeouts: optional
	Timeouts *containernodepool.Timeouts `hcl:"timeouts,block"`
	// UpgradeSettings: optional
	UpgradeSettings *containernodepool.UpgradeSettings `hcl:"upgrade_settings,block"`
}
type containerNodePoolAttributes struct {
	ref terra.Reference
}

// Cluster returns a reference to field cluster of google_container_node_pool.
func (cnp containerNodePoolAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("cluster"))
}

// Id returns a reference to field id of google_container_node_pool.
func (cnp containerNodePoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("id"))
}

// InitialNodeCount returns a reference to field initial_node_count of google_container_node_pool.
func (cnp containerNodePoolAttributes) InitialNodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(cnp.ref.Append("initial_node_count"))
}

// InstanceGroupUrls returns a reference to field instance_group_urls of google_container_node_pool.
func (cnp containerNodePoolAttributes) InstanceGroupUrls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cnp.ref.Append("instance_group_urls"))
}

// Location returns a reference to field location of google_container_node_pool.
func (cnp containerNodePoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("location"))
}

// ManagedInstanceGroupUrls returns a reference to field managed_instance_group_urls of google_container_node_pool.
func (cnp containerNodePoolAttributes) ManagedInstanceGroupUrls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cnp.ref.Append("managed_instance_group_urls"))
}

// MaxPodsPerNode returns a reference to field max_pods_per_node of google_container_node_pool.
func (cnp containerNodePoolAttributes) MaxPodsPerNode() terra.NumberValue {
	return terra.ReferenceAsNumber(cnp.ref.Append("max_pods_per_node"))
}

// Name returns a reference to field name of google_container_node_pool.
func (cnp containerNodePoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of google_container_node_pool.
func (cnp containerNodePoolAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("name_prefix"))
}

// NodeCount returns a reference to field node_count of google_container_node_pool.
func (cnp containerNodePoolAttributes) NodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(cnp.ref.Append("node_count"))
}

// NodeLocations returns a reference to field node_locations of google_container_node_pool.
func (cnp containerNodePoolAttributes) NodeLocations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cnp.ref.Append("node_locations"))
}

// Operation returns a reference to field operation of google_container_node_pool.
func (cnp containerNodePoolAttributes) Operation() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("operation"))
}

// Project returns a reference to field project of google_container_node_pool.
func (cnp containerNodePoolAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("project"))
}

// Version returns a reference to field version of google_container_node_pool.
func (cnp containerNodePoolAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("version"))
}

func (cnp containerNodePoolAttributes) Autoscaling() terra.ListValue[containernodepool.AutoscalingAttributes] {
	return terra.ReferenceAsList[containernodepool.AutoscalingAttributes](cnp.ref.Append("autoscaling"))
}

func (cnp containerNodePoolAttributes) Management() terra.ListValue[containernodepool.ManagementAttributes] {
	return terra.ReferenceAsList[containernodepool.ManagementAttributes](cnp.ref.Append("management"))
}

func (cnp containerNodePoolAttributes) NetworkConfig() terra.ListValue[containernodepool.NetworkConfigAttributes] {
	return terra.ReferenceAsList[containernodepool.NetworkConfigAttributes](cnp.ref.Append("network_config"))
}

func (cnp containerNodePoolAttributes) NodeConfig() terra.ListValue[containernodepool.NodeConfigAttributes] {
	return terra.ReferenceAsList[containernodepool.NodeConfigAttributes](cnp.ref.Append("node_config"))
}

func (cnp containerNodePoolAttributes) PlacementPolicy() terra.ListValue[containernodepool.PlacementPolicyAttributes] {
	return terra.ReferenceAsList[containernodepool.PlacementPolicyAttributes](cnp.ref.Append("placement_policy"))
}

func (cnp containerNodePoolAttributes) Timeouts() containernodepool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containernodepool.TimeoutsAttributes](cnp.ref.Append("timeouts"))
}

func (cnp containerNodePoolAttributes) UpgradeSettings() terra.ListValue[containernodepool.UpgradeSettingsAttributes] {
	return terra.ReferenceAsList[containernodepool.UpgradeSettingsAttributes](cnp.ref.Append("upgrade_settings"))
}

type containerNodePoolState struct {
	Cluster                  string                                   `json:"cluster"`
	Id                       string                                   `json:"id"`
	InitialNodeCount         float64                                  `json:"initial_node_count"`
	InstanceGroupUrls        []string                                 `json:"instance_group_urls"`
	Location                 string                                   `json:"location"`
	ManagedInstanceGroupUrls []string                                 `json:"managed_instance_group_urls"`
	MaxPodsPerNode           float64                                  `json:"max_pods_per_node"`
	Name                     string                                   `json:"name"`
	NamePrefix               string                                   `json:"name_prefix"`
	NodeCount                float64                                  `json:"node_count"`
	NodeLocations            []string                                 `json:"node_locations"`
	Operation                string                                   `json:"operation"`
	Project                  string                                   `json:"project"`
	Version                  string                                   `json:"version"`
	Autoscaling              []containernodepool.AutoscalingState     `json:"autoscaling"`
	Management               []containernodepool.ManagementState      `json:"management"`
	NetworkConfig            []containernodepool.NetworkConfigState   `json:"network_config"`
	NodeConfig               []containernodepool.NodeConfigState      `json:"node_config"`
	PlacementPolicy          []containernodepool.PlacementPolicyState `json:"placement_policy"`
	Timeouts                 *containernodepool.TimeoutsState         `json:"timeouts"`
	UpgradeSettings          []containernodepool.UpgradeSettingsState `json:"upgrade_settings"`
}
