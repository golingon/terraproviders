// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_container_aws_node_pool

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_container_aws_node_pool.
type Resource struct {
	Name      string
	Args      Args
	state     *googleContainerAwsNodePoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcanp *Resource) Type() string {
	return "google_container_aws_node_pool"
}

// LocalName returns the local name for [Resource].
func (gcanp *Resource) LocalName() string {
	return gcanp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcanp *Resource) Configuration() interface{} {
	return gcanp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcanp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcanp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcanp *Resource) Dependencies() terra.Dependencies {
	return gcanp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcanp *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcanp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcanp *Resource) Attributes() googleContainerAwsNodePoolAttributes {
	return googleContainerAwsNodePoolAttributes{ref: terra.ReferenceResource(gcanp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcanp *Resource) ImportState(state io.Reader) error {
	gcanp.state = &googleContainerAwsNodePoolState{}
	if err := json.NewDecoder(state).Decode(gcanp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcanp.Type(), gcanp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcanp *Resource) State() (*googleContainerAwsNodePoolState, bool) {
	return gcanp.state, gcanp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcanp *Resource) StateMust() *googleContainerAwsNodePoolState {
	if gcanp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcanp.Type(), gcanp.LocalName()))
	}
	return gcanp.state
}

// Args contains the configurations for google_container_aws_node_pool.
type Args struct {
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
	Autoscaling *Autoscaling `hcl:"autoscaling,block" validate:"required"`
	// Config: required
	Config *Config `hcl:"config,block" validate:"required"`
	// Management: optional
	Management *Management `hcl:"management,block"`
	// MaxPodsConstraint: required
	MaxPodsConstraint *MaxPodsConstraint `hcl:"max_pods_constraint,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// UpdateSettings: optional
	UpdateSettings *UpdateSettings `hcl:"update_settings,block"`
}

type googleContainerAwsNodePoolAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcanp.ref.Append("annotations"))
}

// Cluster returns a reference to field cluster of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(gcanp.ref.Append("cluster"))
}

// CreateTime returns a reference to field create_time of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gcanp.ref.Append("create_time"))
}

// EffectiveAnnotations returns a reference to field effective_annotations of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) EffectiveAnnotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcanp.ref.Append("effective_annotations"))
}

// Etag returns a reference to field etag of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gcanp.ref.Append("etag"))
}

// Id returns a reference to field id of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcanp.ref.Append("id"))
}

// Location returns a reference to field location of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gcanp.ref.Append("location"))
}

// Name returns a reference to field name of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcanp.ref.Append("name"))
}

// Project returns a reference to field project of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcanp.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(gcanp.ref.Append("reconciling"))
}

// State returns a reference to field state of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gcanp.ref.Append("state"))
}

// SubnetId returns a reference to field subnet_id of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(gcanp.ref.Append("subnet_id"))
}

// Uid returns a reference to field uid of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(gcanp.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gcanp.ref.Append("update_time"))
}

// Version returns a reference to field version of google_container_aws_node_pool.
func (gcanp googleContainerAwsNodePoolAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(gcanp.ref.Append("version"))
}

func (gcanp googleContainerAwsNodePoolAttributes) Autoscaling() terra.ListValue[AutoscalingAttributes] {
	return terra.ReferenceAsList[AutoscalingAttributes](gcanp.ref.Append("autoscaling"))
}

func (gcanp googleContainerAwsNodePoolAttributes) Config() terra.ListValue[ConfigAttributes] {
	return terra.ReferenceAsList[ConfigAttributes](gcanp.ref.Append("config"))
}

func (gcanp googleContainerAwsNodePoolAttributes) Management() terra.ListValue[ManagementAttributes] {
	return terra.ReferenceAsList[ManagementAttributes](gcanp.ref.Append("management"))
}

func (gcanp googleContainerAwsNodePoolAttributes) MaxPodsConstraint() terra.ListValue[MaxPodsConstraintAttributes] {
	return terra.ReferenceAsList[MaxPodsConstraintAttributes](gcanp.ref.Append("max_pods_constraint"))
}

func (gcanp googleContainerAwsNodePoolAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcanp.ref.Append("timeouts"))
}

func (gcanp googleContainerAwsNodePoolAttributes) UpdateSettings() terra.ListValue[UpdateSettingsAttributes] {
	return terra.ReferenceAsList[UpdateSettingsAttributes](gcanp.ref.Append("update_settings"))
}

type googleContainerAwsNodePoolState struct {
	Annotations          map[string]string        `json:"annotations"`
	Cluster              string                   `json:"cluster"`
	CreateTime           string                   `json:"create_time"`
	EffectiveAnnotations map[string]string        `json:"effective_annotations"`
	Etag                 string                   `json:"etag"`
	Id                   string                   `json:"id"`
	Location             string                   `json:"location"`
	Name                 string                   `json:"name"`
	Project              string                   `json:"project"`
	Reconciling          bool                     `json:"reconciling"`
	State                string                   `json:"state"`
	SubnetId             string                   `json:"subnet_id"`
	Uid                  string                   `json:"uid"`
	UpdateTime           string                   `json:"update_time"`
	Version              string                   `json:"version"`
	Autoscaling          []AutoscalingState       `json:"autoscaling"`
	Config               []ConfigState            `json:"config"`
	Management           []ManagementState        `json:"management"`
	MaxPodsConstraint    []MaxPodsConstraintState `json:"max_pods_constraint"`
	Timeouts             *TimeoutsState           `json:"timeouts"`
	UpdateSettings       []UpdateSettingsState    `json:"update_settings"`
}
