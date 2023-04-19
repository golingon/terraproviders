// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	containerazurenodepool "github.com/golingon/terraproviders/googlebeta/4.62.0/containerazurenodepool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerAzureNodePool creates a new instance of [ContainerAzureNodePool].
func NewContainerAzureNodePool(name string, args ContainerAzureNodePoolArgs) *ContainerAzureNodePool {
	return &ContainerAzureNodePool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAzureNodePool)(nil)

// ContainerAzureNodePool represents the Terraform resource google_container_azure_node_pool.
type ContainerAzureNodePool struct {
	Name      string
	Args      ContainerAzureNodePoolArgs
	state     *containerAzureNodePoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerAzureNodePool].
func (canp *ContainerAzureNodePool) Type() string {
	return "google_container_azure_node_pool"
}

// LocalName returns the local name for [ContainerAzureNodePool].
func (canp *ContainerAzureNodePool) LocalName() string {
	return canp.Name
}

// Configuration returns the configuration (args) for [ContainerAzureNodePool].
func (canp *ContainerAzureNodePool) Configuration() interface{} {
	return canp.Args
}

// DependOn is used for other resources to depend on [ContainerAzureNodePool].
func (canp *ContainerAzureNodePool) DependOn() terra.Reference {
	return terra.ReferenceResource(canp)
}

// Dependencies returns the list of resources [ContainerAzureNodePool] depends_on.
func (canp *ContainerAzureNodePool) Dependencies() terra.Dependencies {
	return canp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerAzureNodePool].
func (canp *ContainerAzureNodePool) LifecycleManagement() *terra.Lifecycle {
	return canp.Lifecycle
}

// Attributes returns the attributes for [ContainerAzureNodePool].
func (canp *ContainerAzureNodePool) Attributes() containerAzureNodePoolAttributes {
	return containerAzureNodePoolAttributes{ref: terra.ReferenceResource(canp)}
}

// ImportState imports the given attribute values into [ContainerAzureNodePool]'s state.
func (canp *ContainerAzureNodePool) ImportState(av io.Reader) error {
	canp.state = &containerAzureNodePoolState{}
	if err := json.NewDecoder(av).Decode(canp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", canp.Type(), canp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerAzureNodePool] has state.
func (canp *ContainerAzureNodePool) State() (*containerAzureNodePoolState, bool) {
	return canp.state, canp.state != nil
}

// StateMust returns the state for [ContainerAzureNodePool]. Panics if the state is nil.
func (canp *ContainerAzureNodePool) StateMust() *containerAzureNodePoolState {
	if canp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", canp.Type(), canp.LocalName()))
	}
	return canp.state
}

// ContainerAzureNodePoolArgs contains the configurations for google_container_azure_node_pool.
type ContainerAzureNodePoolArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// AzureAvailabilityZone: string, optional
	AzureAvailabilityZone terra.StringValue `hcl:"azure_availability_zone,attr"`
	// Cluster: string, required
	Cluster terra.StringValue `hcl:"cluster,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// Autoscaling: required
	Autoscaling *containerazurenodepool.Autoscaling `hcl:"autoscaling,block" validate:"required"`
	// Config: required
	Config *containerazurenodepool.Config `hcl:"config,block" validate:"required"`
	// MaxPodsConstraint: required
	MaxPodsConstraint *containerazurenodepool.MaxPodsConstraint `hcl:"max_pods_constraint,block" validate:"required"`
	// Timeouts: optional
	Timeouts *containerazurenodepool.Timeouts `hcl:"timeouts,block"`
}
type containerAzureNodePoolAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](canp.ref.Append("annotations"))
}

// AzureAvailabilityZone returns a reference to field azure_availability_zone of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) AzureAvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("azure_availability_zone"))
}

// Cluster returns a reference to field cluster of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("cluster"))
}

// CreateTime returns a reference to field create_time of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("create_time"))
}

// Etag returns a reference to field etag of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("etag"))
}

// Id returns a reference to field id of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("id"))
}

// Location returns a reference to field location of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("location"))
}

// Name returns a reference to field name of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("name"))
}

// Project returns a reference to field project of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(canp.ref.Append("reconciling"))
}

// State returns a reference to field state of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("state"))
}

// SubnetId returns a reference to field subnet_id of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("subnet_id"))
}

// Uid returns a reference to field uid of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("update_time"))
}

// Version returns a reference to field version of google_container_azure_node_pool.
func (canp containerAzureNodePoolAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("version"))
}

func (canp containerAzureNodePoolAttributes) Autoscaling() terra.ListValue[containerazurenodepool.AutoscalingAttributes] {
	return terra.ReferenceAsList[containerazurenodepool.AutoscalingAttributes](canp.ref.Append("autoscaling"))
}

func (canp containerAzureNodePoolAttributes) Config() terra.ListValue[containerazurenodepool.ConfigAttributes] {
	return terra.ReferenceAsList[containerazurenodepool.ConfigAttributes](canp.ref.Append("config"))
}

func (canp containerAzureNodePoolAttributes) MaxPodsConstraint() terra.ListValue[containerazurenodepool.MaxPodsConstraintAttributes] {
	return terra.ReferenceAsList[containerazurenodepool.MaxPodsConstraintAttributes](canp.ref.Append("max_pods_constraint"))
}

func (canp containerAzureNodePoolAttributes) Timeouts() containerazurenodepool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerazurenodepool.TimeoutsAttributes](canp.ref.Append("timeouts"))
}

type containerAzureNodePoolState struct {
	Annotations           map[string]string                               `json:"annotations"`
	AzureAvailabilityZone string                                          `json:"azure_availability_zone"`
	Cluster               string                                          `json:"cluster"`
	CreateTime            string                                          `json:"create_time"`
	Etag                  string                                          `json:"etag"`
	Id                    string                                          `json:"id"`
	Location              string                                          `json:"location"`
	Name                  string                                          `json:"name"`
	Project               string                                          `json:"project"`
	Reconciling           bool                                            `json:"reconciling"`
	State                 string                                          `json:"state"`
	SubnetId              string                                          `json:"subnet_id"`
	Uid                   string                                          `json:"uid"`
	UpdateTime            string                                          `json:"update_time"`
	Version               string                                          `json:"version"`
	Autoscaling           []containerazurenodepool.AutoscalingState       `json:"autoscaling"`
	Config                []containerazurenodepool.ConfigState            `json:"config"`
	MaxPodsConstraint     []containerazurenodepool.MaxPodsConstraintState `json:"max_pods_constraint"`
	Timeouts              *containerazurenodepool.TimeoutsState           `json:"timeouts"`
}
