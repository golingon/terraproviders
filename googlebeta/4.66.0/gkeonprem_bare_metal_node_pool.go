// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	gkeonprembaremetalnodepool "github.com/golingon/terraproviders/googlebeta/4.66.0/gkeonprembaremetalnodepool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGkeonpremBareMetalNodePool creates a new instance of [GkeonpremBareMetalNodePool].
func NewGkeonpremBareMetalNodePool(name string, args GkeonpremBareMetalNodePoolArgs) *GkeonpremBareMetalNodePool {
	return &GkeonpremBareMetalNodePool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeonpremBareMetalNodePool)(nil)

// GkeonpremBareMetalNodePool represents the Terraform resource google_gkeonprem_bare_metal_node_pool.
type GkeonpremBareMetalNodePool struct {
	Name      string
	Args      GkeonpremBareMetalNodePoolArgs
	state     *gkeonpremBareMetalNodePoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeonpremBareMetalNodePool].
func (gbmnp *GkeonpremBareMetalNodePool) Type() string {
	return "google_gkeonprem_bare_metal_node_pool"
}

// LocalName returns the local name for [GkeonpremBareMetalNodePool].
func (gbmnp *GkeonpremBareMetalNodePool) LocalName() string {
	return gbmnp.Name
}

// Configuration returns the configuration (args) for [GkeonpremBareMetalNodePool].
func (gbmnp *GkeonpremBareMetalNodePool) Configuration() interface{} {
	return gbmnp.Args
}

// DependOn is used for other resources to depend on [GkeonpremBareMetalNodePool].
func (gbmnp *GkeonpremBareMetalNodePool) DependOn() terra.Reference {
	return terra.ReferenceResource(gbmnp)
}

// Dependencies returns the list of resources [GkeonpremBareMetalNodePool] depends_on.
func (gbmnp *GkeonpremBareMetalNodePool) Dependencies() terra.Dependencies {
	return gbmnp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeonpremBareMetalNodePool].
func (gbmnp *GkeonpremBareMetalNodePool) LifecycleManagement() *terra.Lifecycle {
	return gbmnp.Lifecycle
}

// Attributes returns the attributes for [GkeonpremBareMetalNodePool].
func (gbmnp *GkeonpremBareMetalNodePool) Attributes() gkeonpremBareMetalNodePoolAttributes {
	return gkeonpremBareMetalNodePoolAttributes{ref: terra.ReferenceResource(gbmnp)}
}

// ImportState imports the given attribute values into [GkeonpremBareMetalNodePool]'s state.
func (gbmnp *GkeonpremBareMetalNodePool) ImportState(av io.Reader) error {
	gbmnp.state = &gkeonpremBareMetalNodePoolState{}
	if err := json.NewDecoder(av).Decode(gbmnp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gbmnp.Type(), gbmnp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeonpremBareMetalNodePool] has state.
func (gbmnp *GkeonpremBareMetalNodePool) State() (*gkeonpremBareMetalNodePoolState, bool) {
	return gbmnp.state, gbmnp.state != nil
}

// StateMust returns the state for [GkeonpremBareMetalNodePool]. Panics if the state is nil.
func (gbmnp *GkeonpremBareMetalNodePool) StateMust() *gkeonpremBareMetalNodePoolState {
	if gbmnp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gbmnp.Type(), gbmnp.LocalName()))
	}
	return gbmnp.state
}

// GkeonpremBareMetalNodePoolArgs contains the configurations for google_gkeonprem_bare_metal_node_pool.
type GkeonpremBareMetalNodePoolArgs struct {
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
	// Status: min=0
	Status []gkeonprembaremetalnodepool.Status `hcl:"status,block" validate:"min=0"`
	// NodePoolConfig: required
	NodePoolConfig *gkeonprembaremetalnodepool.NodePoolConfig `hcl:"node_pool_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *gkeonprembaremetalnodepool.Timeouts `hcl:"timeouts,block"`
}
type gkeonpremBareMetalNodePoolAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gbmnp.ref.Append("annotations"))
}

// BareMetalCluster returns a reference to field bare_metal_cluster of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) BareMetalCluster() terra.StringValue {
	return terra.ReferenceAsString(gbmnp.ref.Append("bare_metal_cluster"))
}

// CreateTime returns a reference to field create_time of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gbmnp.ref.Append("create_time"))
}

// DeleteTime returns a reference to field delete_time of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) DeleteTime() terra.StringValue {
	return terra.ReferenceAsString(gbmnp.ref.Append("delete_time"))
}

// DisplayName returns a reference to field display_name of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gbmnp.ref.Append("display_name"))
}

// Etag returns a reference to field etag of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gbmnp.ref.Append("etag"))
}

// Id returns a reference to field id of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbmnp.ref.Append("id"))
}

// Location returns a reference to field location of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gbmnp.ref.Append("location"))
}

// Name returns a reference to field name of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gbmnp.ref.Append("name"))
}

// Project returns a reference to field project of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbmnp.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(gbmnp.ref.Append("reconciling"))
}

// State returns a reference to field state of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gbmnp.ref.Append("state"))
}

// Uid returns a reference to field uid of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(gbmnp.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_gkeonprem_bare_metal_node_pool.
func (gbmnp gkeonpremBareMetalNodePoolAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gbmnp.ref.Append("update_time"))
}

func (gbmnp gkeonpremBareMetalNodePoolAttributes) Status() terra.ListValue[gkeonprembaremetalnodepool.StatusAttributes] {
	return terra.ReferenceAsList[gkeonprembaremetalnodepool.StatusAttributes](gbmnp.ref.Append("status"))
}

func (gbmnp gkeonpremBareMetalNodePoolAttributes) NodePoolConfig() terra.ListValue[gkeonprembaremetalnodepool.NodePoolConfigAttributes] {
	return terra.ReferenceAsList[gkeonprembaremetalnodepool.NodePoolConfigAttributes](gbmnp.ref.Append("node_pool_config"))
}

func (gbmnp gkeonpremBareMetalNodePoolAttributes) Timeouts() gkeonprembaremetalnodepool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gkeonprembaremetalnodepool.TimeoutsAttributes](gbmnp.ref.Append("timeouts"))
}

type gkeonpremBareMetalNodePoolState struct {
	Annotations      map[string]string                                `json:"annotations"`
	BareMetalCluster string                                           `json:"bare_metal_cluster"`
	CreateTime       string                                           `json:"create_time"`
	DeleteTime       string                                           `json:"delete_time"`
	DisplayName      string                                           `json:"display_name"`
	Etag             string                                           `json:"etag"`
	Id               string                                           `json:"id"`
	Location         string                                           `json:"location"`
	Name             string                                           `json:"name"`
	Project          string                                           `json:"project"`
	Reconciling      bool                                             `json:"reconciling"`
	State            string                                           `json:"state"`
	Uid              string                                           `json:"uid"`
	UpdateTime       string                                           `json:"update_time"`
	Status           []gkeonprembaremetalnodepool.StatusState         `json:"status"`
	NodePoolConfig   []gkeonprembaremetalnodepool.NodePoolConfigState `json:"node_pool_config"`
	Timeouts         *gkeonprembaremetalnodepool.TimeoutsState        `json:"timeouts"`
}
