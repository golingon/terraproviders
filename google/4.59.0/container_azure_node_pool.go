// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	containerazurenodepool "github.com/golingon/terraproviders/google/4.59.0/containerazurenodepool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewContainerAzureNodePool(name string, args ContainerAzureNodePoolArgs) *ContainerAzureNodePool {
	return &ContainerAzureNodePool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAzureNodePool)(nil)

type ContainerAzureNodePool struct {
	Name  string
	Args  ContainerAzureNodePoolArgs
	state *containerAzureNodePoolState
}

func (canp *ContainerAzureNodePool) Type() string {
	return "google_container_azure_node_pool"
}

func (canp *ContainerAzureNodePool) LocalName() string {
	return canp.Name
}

func (canp *ContainerAzureNodePool) Configuration() interface{} {
	return canp.Args
}

func (canp *ContainerAzureNodePool) Attributes() containerAzureNodePoolAttributes {
	return containerAzureNodePoolAttributes{ref: terra.ReferenceResource(canp)}
}

func (canp *ContainerAzureNodePool) ImportState(av io.Reader) error {
	canp.state = &containerAzureNodePoolState{}
	if err := json.NewDecoder(av).Decode(canp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", canp.Type(), canp.LocalName(), err)
	}
	return nil
}

func (canp *ContainerAzureNodePool) State() (*containerAzureNodePoolState, bool) {
	return canp.state, canp.state != nil
}

func (canp *ContainerAzureNodePool) StateMust() *containerAzureNodePoolState {
	if canp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", canp.Type(), canp.LocalName()))
	}
	return canp.state
}

func (canp *ContainerAzureNodePool) DependOn() terra.Reference {
	return terra.ReferenceResource(canp)
}

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
	// DependsOn contains resources that ContainerAzureNodePool depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type containerAzureNodePoolAttributes struct {
	ref terra.Reference
}

func (canp containerAzureNodePoolAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](canp.ref.Append("annotations"))
}

func (canp containerAzureNodePoolAttributes) AzureAvailabilityZone() terra.StringValue {
	return terra.ReferenceString(canp.ref.Append("azure_availability_zone"))
}

func (canp containerAzureNodePoolAttributes) Cluster() terra.StringValue {
	return terra.ReferenceString(canp.ref.Append("cluster"))
}

func (canp containerAzureNodePoolAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceString(canp.ref.Append("create_time"))
}

func (canp containerAzureNodePoolAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(canp.ref.Append("etag"))
}

func (canp containerAzureNodePoolAttributes) Id() terra.StringValue {
	return terra.ReferenceString(canp.ref.Append("id"))
}

func (canp containerAzureNodePoolAttributes) Location() terra.StringValue {
	return terra.ReferenceString(canp.ref.Append("location"))
}

func (canp containerAzureNodePoolAttributes) Name() terra.StringValue {
	return terra.ReferenceString(canp.ref.Append("name"))
}

func (canp containerAzureNodePoolAttributes) Project() terra.StringValue {
	return terra.ReferenceString(canp.ref.Append("project"))
}

func (canp containerAzureNodePoolAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceBool(canp.ref.Append("reconciling"))
}

func (canp containerAzureNodePoolAttributes) State() terra.StringValue {
	return terra.ReferenceString(canp.ref.Append("state"))
}

func (canp containerAzureNodePoolAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceString(canp.ref.Append("subnet_id"))
}

func (canp containerAzureNodePoolAttributes) Uid() terra.StringValue {
	return terra.ReferenceString(canp.ref.Append("uid"))
}

func (canp containerAzureNodePoolAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceString(canp.ref.Append("update_time"))
}

func (canp containerAzureNodePoolAttributes) Version() terra.StringValue {
	return terra.ReferenceString(canp.ref.Append("version"))
}

func (canp containerAzureNodePoolAttributes) Autoscaling() terra.ListValue[containerazurenodepool.AutoscalingAttributes] {
	return terra.ReferenceList[containerazurenodepool.AutoscalingAttributes](canp.ref.Append("autoscaling"))
}

func (canp containerAzureNodePoolAttributes) Config() terra.ListValue[containerazurenodepool.ConfigAttributes] {
	return terra.ReferenceList[containerazurenodepool.ConfigAttributes](canp.ref.Append("config"))
}

func (canp containerAzureNodePoolAttributes) MaxPodsConstraint() terra.ListValue[containerazurenodepool.MaxPodsConstraintAttributes] {
	return terra.ReferenceList[containerazurenodepool.MaxPodsConstraintAttributes](canp.ref.Append("max_pods_constraint"))
}

func (canp containerAzureNodePoolAttributes) Timeouts() containerazurenodepool.TimeoutsAttributes {
	return terra.ReferenceSingle[containerazurenodepool.TimeoutsAttributes](canp.ref.Append("timeouts"))
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
