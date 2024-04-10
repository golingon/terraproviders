// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	containerawsnodepool "github.com/golingon/terraproviders/google/5.24.0/containerawsnodepool"
	"io"
)

// NewContainerAwsNodePool creates a new instance of [ContainerAwsNodePool].
func NewContainerAwsNodePool(name string, args ContainerAwsNodePoolArgs) *ContainerAwsNodePool {
	return &ContainerAwsNodePool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAwsNodePool)(nil)

// ContainerAwsNodePool represents the Terraform resource google_container_aws_node_pool.
type ContainerAwsNodePool struct {
	Name      string
	Args      ContainerAwsNodePoolArgs
	state     *containerAwsNodePoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerAwsNodePool].
func (canp *ContainerAwsNodePool) Type() string {
	return "google_container_aws_node_pool"
}

// LocalName returns the local name for [ContainerAwsNodePool].
func (canp *ContainerAwsNodePool) LocalName() string {
	return canp.Name
}

// Configuration returns the configuration (args) for [ContainerAwsNodePool].
func (canp *ContainerAwsNodePool) Configuration() interface{} {
	return canp.Args
}

// DependOn is used for other resources to depend on [ContainerAwsNodePool].
func (canp *ContainerAwsNodePool) DependOn() terra.Reference {
	return terra.ReferenceResource(canp)
}

// Dependencies returns the list of resources [ContainerAwsNodePool] depends_on.
func (canp *ContainerAwsNodePool) Dependencies() terra.Dependencies {
	return canp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerAwsNodePool].
func (canp *ContainerAwsNodePool) LifecycleManagement() *terra.Lifecycle {
	return canp.Lifecycle
}

// Attributes returns the attributes for [ContainerAwsNodePool].
func (canp *ContainerAwsNodePool) Attributes() containerAwsNodePoolAttributes {
	return containerAwsNodePoolAttributes{ref: terra.ReferenceResource(canp)}
}

// ImportState imports the given attribute values into [ContainerAwsNodePool]'s state.
func (canp *ContainerAwsNodePool) ImportState(av io.Reader) error {
	canp.state = &containerAwsNodePoolState{}
	if err := json.NewDecoder(av).Decode(canp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", canp.Type(), canp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerAwsNodePool] has state.
func (canp *ContainerAwsNodePool) State() (*containerAwsNodePoolState, bool) {
	return canp.state, canp.state != nil
}

// StateMust returns the state for [ContainerAwsNodePool]. Panics if the state is nil.
func (canp *ContainerAwsNodePool) StateMust() *containerAwsNodePoolState {
	if canp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", canp.Type(), canp.LocalName()))
	}
	return canp.state
}

// ContainerAwsNodePoolArgs contains the configurations for google_container_aws_node_pool.
type ContainerAwsNodePoolArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
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
	Autoscaling *containerawsnodepool.Autoscaling `hcl:"autoscaling,block" validate:"required"`
	// Config: required
	Config *containerawsnodepool.Config `hcl:"config,block" validate:"required"`
	// Management: optional
	Management *containerawsnodepool.Management `hcl:"management,block"`
	// MaxPodsConstraint: required
	MaxPodsConstraint *containerawsnodepool.MaxPodsConstraint `hcl:"max_pods_constraint,block" validate:"required"`
	// Timeouts: optional
	Timeouts *containerawsnodepool.Timeouts `hcl:"timeouts,block"`
	// UpdateSettings: optional
	UpdateSettings *containerawsnodepool.UpdateSettings `hcl:"update_settings,block"`
}
type containerAwsNodePoolAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](canp.ref.Append("annotations"))
}

// Cluster returns a reference to field cluster of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("cluster"))
}

// CreateTime returns a reference to field create_time of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("create_time"))
}

// EffectiveAnnotations returns a reference to field effective_annotations of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) EffectiveAnnotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](canp.ref.Append("effective_annotations"))
}

// Etag returns a reference to field etag of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("etag"))
}

// Id returns a reference to field id of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("id"))
}

// Location returns a reference to field location of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("location"))
}

// Name returns a reference to field name of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("name"))
}

// Project returns a reference to field project of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(canp.ref.Append("reconciling"))
}

// State returns a reference to field state of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("state"))
}

// SubnetId returns a reference to field subnet_id of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("subnet_id"))
}

// Uid returns a reference to field uid of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("update_time"))
}

// Version returns a reference to field version of google_container_aws_node_pool.
func (canp containerAwsNodePoolAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(canp.ref.Append("version"))
}

func (canp containerAwsNodePoolAttributes) Autoscaling() terra.ListValue[containerawsnodepool.AutoscalingAttributes] {
	return terra.ReferenceAsList[containerawsnodepool.AutoscalingAttributes](canp.ref.Append("autoscaling"))
}

func (canp containerAwsNodePoolAttributes) Config() terra.ListValue[containerawsnodepool.ConfigAttributes] {
	return terra.ReferenceAsList[containerawsnodepool.ConfigAttributes](canp.ref.Append("config"))
}

func (canp containerAwsNodePoolAttributes) Management() terra.ListValue[containerawsnodepool.ManagementAttributes] {
	return terra.ReferenceAsList[containerawsnodepool.ManagementAttributes](canp.ref.Append("management"))
}

func (canp containerAwsNodePoolAttributes) MaxPodsConstraint() terra.ListValue[containerawsnodepool.MaxPodsConstraintAttributes] {
	return terra.ReferenceAsList[containerawsnodepool.MaxPodsConstraintAttributes](canp.ref.Append("max_pods_constraint"))
}

func (canp containerAwsNodePoolAttributes) Timeouts() containerawsnodepool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerawsnodepool.TimeoutsAttributes](canp.ref.Append("timeouts"))
}

func (canp containerAwsNodePoolAttributes) UpdateSettings() terra.ListValue[containerawsnodepool.UpdateSettingsAttributes] {
	return terra.ReferenceAsList[containerawsnodepool.UpdateSettingsAttributes](canp.ref.Append("update_settings"))
}

type containerAwsNodePoolState struct {
	Annotations          map[string]string                             `json:"annotations"`
	Cluster              string                                        `json:"cluster"`
	CreateTime           string                                        `json:"create_time"`
	EffectiveAnnotations map[string]string                             `json:"effective_annotations"`
	Etag                 string                                        `json:"etag"`
	Id                   string                                        `json:"id"`
	Location             string                                        `json:"location"`
	Name                 string                                        `json:"name"`
	Project              string                                        `json:"project"`
	Reconciling          bool                                          `json:"reconciling"`
	State                string                                        `json:"state"`
	SubnetId             string                                        `json:"subnet_id"`
	Uid                  string                                        `json:"uid"`
	UpdateTime           string                                        `json:"update_time"`
	Version              string                                        `json:"version"`
	Autoscaling          []containerawsnodepool.AutoscalingState       `json:"autoscaling"`
	Config               []containerawsnodepool.ConfigState            `json:"config"`
	Management           []containerawsnodepool.ManagementState        `json:"management"`
	MaxPodsConstraint    []containerawsnodepool.MaxPodsConstraintState `json:"max_pods_constraint"`
	Timeouts             *containerawsnodepool.TimeoutsState           `json:"timeouts"`
	UpdateSettings       []containerawsnodepool.UpdateSettingsState    `json:"update_settings"`
}
