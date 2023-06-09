// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	gkeonpremvmwarenodepool "github.com/golingon/terraproviders/googlebeta/4.64.0/gkeonpremvmwarenodepool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGkeonpremVmwareNodePool creates a new instance of [GkeonpremVmwareNodePool].
func NewGkeonpremVmwareNodePool(name string, args GkeonpremVmwareNodePoolArgs) *GkeonpremVmwareNodePool {
	return &GkeonpremVmwareNodePool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeonpremVmwareNodePool)(nil)

// GkeonpremVmwareNodePool represents the Terraform resource google_gkeonprem_vmware_node_pool.
type GkeonpremVmwareNodePool struct {
	Name      string
	Args      GkeonpremVmwareNodePoolArgs
	state     *gkeonpremVmwareNodePoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeonpremVmwareNodePool].
func (gvnp *GkeonpremVmwareNodePool) Type() string {
	return "google_gkeonprem_vmware_node_pool"
}

// LocalName returns the local name for [GkeonpremVmwareNodePool].
func (gvnp *GkeonpremVmwareNodePool) LocalName() string {
	return gvnp.Name
}

// Configuration returns the configuration (args) for [GkeonpremVmwareNodePool].
func (gvnp *GkeonpremVmwareNodePool) Configuration() interface{} {
	return gvnp.Args
}

// DependOn is used for other resources to depend on [GkeonpremVmwareNodePool].
func (gvnp *GkeonpremVmwareNodePool) DependOn() terra.Reference {
	return terra.ReferenceResource(gvnp)
}

// Dependencies returns the list of resources [GkeonpremVmwareNodePool] depends_on.
func (gvnp *GkeonpremVmwareNodePool) Dependencies() terra.Dependencies {
	return gvnp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeonpremVmwareNodePool].
func (gvnp *GkeonpremVmwareNodePool) LifecycleManagement() *terra.Lifecycle {
	return gvnp.Lifecycle
}

// Attributes returns the attributes for [GkeonpremVmwareNodePool].
func (gvnp *GkeonpremVmwareNodePool) Attributes() gkeonpremVmwareNodePoolAttributes {
	return gkeonpremVmwareNodePoolAttributes{ref: terra.ReferenceResource(gvnp)}
}

// ImportState imports the given attribute values into [GkeonpremVmwareNodePool]'s state.
func (gvnp *GkeonpremVmwareNodePool) ImportState(av io.Reader) error {
	gvnp.state = &gkeonpremVmwareNodePoolState{}
	if err := json.NewDecoder(av).Decode(gvnp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gvnp.Type(), gvnp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeonpremVmwareNodePool] has state.
func (gvnp *GkeonpremVmwareNodePool) State() (*gkeonpremVmwareNodePoolState, bool) {
	return gvnp.state, gvnp.state != nil
}

// StateMust returns the state for [GkeonpremVmwareNodePool]. Panics if the state is nil.
func (gvnp *GkeonpremVmwareNodePool) StateMust() *gkeonpremVmwareNodePoolState {
	if gvnp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gvnp.Type(), gvnp.LocalName()))
	}
	return gvnp.state
}

// GkeonpremVmwareNodePoolArgs contains the configurations for google_gkeonprem_vmware_node_pool.
type GkeonpremVmwareNodePoolArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
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
	// VmwareCluster: string, required
	VmwareCluster terra.StringValue `hcl:"vmware_cluster,attr" validate:"required"`
	// Status: min=0
	Status []gkeonpremvmwarenodepool.Status `hcl:"status,block" validate:"min=0"`
	// Config: required
	Config *gkeonpremvmwarenodepool.Config `hcl:"config,block" validate:"required"`
	// NodePoolAutoscaling: optional
	NodePoolAutoscaling *gkeonpremvmwarenodepool.NodePoolAutoscaling `hcl:"node_pool_autoscaling,block"`
	// Timeouts: optional
	Timeouts *gkeonpremvmwarenodepool.Timeouts `hcl:"timeouts,block"`
}
type gkeonpremVmwareNodePoolAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gvnp.ref.Append("annotations"))
}

// CreateTime returns a reference to field create_time of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gvnp.ref.Append("create_time"))
}

// DeleteTime returns a reference to field delete_time of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) DeleteTime() terra.StringValue {
	return terra.ReferenceAsString(gvnp.ref.Append("delete_time"))
}

// DisplayName returns a reference to field display_name of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gvnp.ref.Append("display_name"))
}

// Etag returns a reference to field etag of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gvnp.ref.Append("etag"))
}

// Id returns a reference to field id of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gvnp.ref.Append("id"))
}

// Location returns a reference to field location of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gvnp.ref.Append("location"))
}

// Name returns a reference to field name of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gvnp.ref.Append("name"))
}

// OnPremVersion returns a reference to field on_prem_version of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) OnPremVersion() terra.StringValue {
	return terra.ReferenceAsString(gvnp.ref.Append("on_prem_version"))
}

// Project returns a reference to field project of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gvnp.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(gvnp.ref.Append("reconciling"))
}

// State returns a reference to field state of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gvnp.ref.Append("state"))
}

// Uid returns a reference to field uid of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(gvnp.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gvnp.ref.Append("update_time"))
}

// VmwareCluster returns a reference to field vmware_cluster of google_gkeonprem_vmware_node_pool.
func (gvnp gkeonpremVmwareNodePoolAttributes) VmwareCluster() terra.StringValue {
	return terra.ReferenceAsString(gvnp.ref.Append("vmware_cluster"))
}

func (gvnp gkeonpremVmwareNodePoolAttributes) Status() terra.ListValue[gkeonpremvmwarenodepool.StatusAttributes] {
	return terra.ReferenceAsList[gkeonpremvmwarenodepool.StatusAttributes](gvnp.ref.Append("status"))
}

func (gvnp gkeonpremVmwareNodePoolAttributes) Config() terra.ListValue[gkeonpremvmwarenodepool.ConfigAttributes] {
	return terra.ReferenceAsList[gkeonpremvmwarenodepool.ConfigAttributes](gvnp.ref.Append("config"))
}

func (gvnp gkeonpremVmwareNodePoolAttributes) NodePoolAutoscaling() terra.ListValue[gkeonpremvmwarenodepool.NodePoolAutoscalingAttributes] {
	return terra.ReferenceAsList[gkeonpremvmwarenodepool.NodePoolAutoscalingAttributes](gvnp.ref.Append("node_pool_autoscaling"))
}

func (gvnp gkeonpremVmwareNodePoolAttributes) Timeouts() gkeonpremvmwarenodepool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gkeonpremvmwarenodepool.TimeoutsAttributes](gvnp.ref.Append("timeouts"))
}

type gkeonpremVmwareNodePoolState struct {
	Annotations         map[string]string                                  `json:"annotations"`
	CreateTime          string                                             `json:"create_time"`
	DeleteTime          string                                             `json:"delete_time"`
	DisplayName         string                                             `json:"display_name"`
	Etag                string                                             `json:"etag"`
	Id                  string                                             `json:"id"`
	Location            string                                             `json:"location"`
	Name                string                                             `json:"name"`
	OnPremVersion       string                                             `json:"on_prem_version"`
	Project             string                                             `json:"project"`
	Reconciling         bool                                               `json:"reconciling"`
	State               string                                             `json:"state"`
	Uid                 string                                             `json:"uid"`
	UpdateTime          string                                             `json:"update_time"`
	VmwareCluster       string                                             `json:"vmware_cluster"`
	Status              []gkeonpremvmwarenodepool.StatusState              `json:"status"`
	Config              []gkeonpremvmwarenodepool.ConfigState              `json:"config"`
	NodePoolAutoscaling []gkeonpremvmwarenodepool.NodePoolAutoscalingState `json:"node_pool_autoscaling"`
	Timeouts            *gkeonpremvmwarenodepool.TimeoutsState             `json:"timeouts"`
}
