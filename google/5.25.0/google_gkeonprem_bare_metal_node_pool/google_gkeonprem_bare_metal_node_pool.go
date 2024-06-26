// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_gkeonprem_bare_metal_node_pool

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

// Resource represents the Terraform resource google_gkeonprem_bare_metal_node_pool.
type Resource struct {
	Name      string
	Args      Args
	state     *googleGkeonpremBareMetalNodePoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ggbmnp *Resource) Type() string {
	return "google_gkeonprem_bare_metal_node_pool"
}

// LocalName returns the local name for [Resource].
func (ggbmnp *Resource) LocalName() string {
	return ggbmnp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ggbmnp *Resource) Configuration() interface{} {
	return ggbmnp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ggbmnp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ggbmnp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ggbmnp *Resource) Dependencies() terra.Dependencies {
	return ggbmnp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ggbmnp *Resource) LifecycleManagement() *terra.Lifecycle {
	return ggbmnp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ggbmnp *Resource) Attributes() googleGkeonpremBareMetalNodePoolAttributes {
	return googleGkeonpremBareMetalNodePoolAttributes{ref: terra.ReferenceResource(ggbmnp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ggbmnp *Resource) ImportState(state io.Reader) error {
	ggbmnp.state = &googleGkeonpremBareMetalNodePoolState{}
	if err := json.NewDecoder(state).Decode(ggbmnp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ggbmnp.Type(), ggbmnp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ggbmnp *Resource) State() (*googleGkeonpremBareMetalNodePoolState, bool) {
	return ggbmnp.state, ggbmnp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ggbmnp *Resource) StateMust() *googleGkeonpremBareMetalNodePoolState {
	if ggbmnp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ggbmnp.Type(), ggbmnp.LocalName()))
	}
	return ggbmnp.state
}

// Args contains the configurations for google_gkeonprem_bare_metal_node_pool.
type Args struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// BareMetalCluster: string, required
	BareMetalCluster terra.StringValue `hcl:"bare_metal_cluster,attr" validate:"required"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// NodePoolConfig: required
	NodePoolConfig *NodePoolConfig `hcl:"node_pool_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleGkeonpremBareMetalNodePoolAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ggbmnp.ref.Append("annotations"))
}

// BareMetalCluster returns a reference to field bare_metal_cluster of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) BareMetalCluster() terra.StringValue {
	return terra.ReferenceAsString(ggbmnp.ref.Append("bare_metal_cluster"))
}

// CreateTime returns a reference to field create_time of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ggbmnp.ref.Append("create_time"))
}

// DeleteTime returns a reference to field delete_time of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) DeleteTime() terra.StringValue {
	return terra.ReferenceAsString(ggbmnp.ref.Append("delete_time"))
}

// DisplayName returns a reference to field display_name of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ggbmnp.ref.Append("display_name"))
}

// EffectiveAnnotations returns a reference to field effective_annotations of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) EffectiveAnnotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ggbmnp.ref.Append("effective_annotations"))
}

// Etag returns a reference to field etag of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ggbmnp.ref.Append("etag"))
}

// Id returns a reference to field id of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ggbmnp.ref.Append("id"))
}

// Location returns a reference to field location of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ggbmnp.ref.Append("location"))
}

// Name returns a reference to field name of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ggbmnp.ref.Append("name"))
}

// Project returns a reference to field project of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ggbmnp.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(ggbmnp.ref.Append("reconciling"))
}

// State returns a reference to field state of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ggbmnp.ref.Append("state"))
}

// Uid returns a reference to field uid of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(ggbmnp.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_gkeonprem_bare_metal_node_pool.
func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ggbmnp.ref.Append("update_time"))
}

func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) Status() terra.ListValue[StatusAttributes] {
	return terra.ReferenceAsList[StatusAttributes](ggbmnp.ref.Append("status"))
}

func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) NodePoolConfig() terra.ListValue[NodePoolConfigAttributes] {
	return terra.ReferenceAsList[NodePoolConfigAttributes](ggbmnp.ref.Append("node_pool_config"))
}

func (ggbmnp googleGkeonpremBareMetalNodePoolAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ggbmnp.ref.Append("timeouts"))
}

type googleGkeonpremBareMetalNodePoolState struct {
	Annotations          map[string]string     `json:"annotations"`
	BareMetalCluster     string                `json:"bare_metal_cluster"`
	CreateTime           string                `json:"create_time"`
	DeleteTime           string                `json:"delete_time"`
	DisplayName          string                `json:"display_name"`
	EffectiveAnnotations map[string]string     `json:"effective_annotations"`
	Etag                 string                `json:"etag"`
	Id                   string                `json:"id"`
	Location             string                `json:"location"`
	Name                 string                `json:"name"`
	Project              string                `json:"project"`
	Reconciling          bool                  `json:"reconciling"`
	State                string                `json:"state"`
	Uid                  string                `json:"uid"`
	UpdateTime           string                `json:"update_time"`
	Status               []StatusState         `json:"status"`
	NodePoolConfig       []NodePoolConfigState `json:"node_pool_config"`
	Timeouts             *TimeoutsState        `json:"timeouts"`
}
