// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	vmwareengineprivatecloud "github.com/golingon/terraproviders/googlebeta/4.76.0/vmwareengineprivatecloud"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVmwareenginePrivateCloud creates a new instance of [VmwareenginePrivateCloud].
func NewVmwareenginePrivateCloud(name string, args VmwareenginePrivateCloudArgs) *VmwareenginePrivateCloud {
	return &VmwareenginePrivateCloud{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VmwareenginePrivateCloud)(nil)

// VmwareenginePrivateCloud represents the Terraform resource google_vmwareengine_private_cloud.
type VmwareenginePrivateCloud struct {
	Name      string
	Args      VmwareenginePrivateCloudArgs
	state     *vmwareenginePrivateCloudState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VmwareenginePrivateCloud].
func (vpc *VmwareenginePrivateCloud) Type() string {
	return "google_vmwareengine_private_cloud"
}

// LocalName returns the local name for [VmwareenginePrivateCloud].
func (vpc *VmwareenginePrivateCloud) LocalName() string {
	return vpc.Name
}

// Configuration returns the configuration (args) for [VmwareenginePrivateCloud].
func (vpc *VmwareenginePrivateCloud) Configuration() interface{} {
	return vpc.Args
}

// DependOn is used for other resources to depend on [VmwareenginePrivateCloud].
func (vpc *VmwareenginePrivateCloud) DependOn() terra.Reference {
	return terra.ReferenceResource(vpc)
}

// Dependencies returns the list of resources [VmwareenginePrivateCloud] depends_on.
func (vpc *VmwareenginePrivateCloud) Dependencies() terra.Dependencies {
	return vpc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VmwareenginePrivateCloud].
func (vpc *VmwareenginePrivateCloud) LifecycleManagement() *terra.Lifecycle {
	return vpc.Lifecycle
}

// Attributes returns the attributes for [VmwareenginePrivateCloud].
func (vpc *VmwareenginePrivateCloud) Attributes() vmwareenginePrivateCloudAttributes {
	return vmwareenginePrivateCloudAttributes{ref: terra.ReferenceResource(vpc)}
}

// ImportState imports the given attribute values into [VmwareenginePrivateCloud]'s state.
func (vpc *VmwareenginePrivateCloud) ImportState(av io.Reader) error {
	vpc.state = &vmwareenginePrivateCloudState{}
	if err := json.NewDecoder(av).Decode(vpc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vpc.Type(), vpc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VmwareenginePrivateCloud] has state.
func (vpc *VmwareenginePrivateCloud) State() (*vmwareenginePrivateCloudState, bool) {
	return vpc.state, vpc.state != nil
}

// StateMust returns the state for [VmwareenginePrivateCloud]. Panics if the state is nil.
func (vpc *VmwareenginePrivateCloud) StateMust() *vmwareenginePrivateCloudState {
	if vpc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vpc.Type(), vpc.LocalName()))
	}
	return vpc.state
}

// VmwareenginePrivateCloudArgs contains the configurations for google_vmwareengine_private_cloud.
type VmwareenginePrivateCloudArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Hcx: min=0
	Hcx []vmwareengineprivatecloud.Hcx `hcl:"hcx,block" validate:"min=0"`
	// Nsx: min=0
	Nsx []vmwareengineprivatecloud.Nsx `hcl:"nsx,block" validate:"min=0"`
	// Vcenter: min=0
	Vcenter []vmwareengineprivatecloud.Vcenter `hcl:"vcenter,block" validate:"min=0"`
	// ManagementCluster: required
	ManagementCluster *vmwareengineprivatecloud.ManagementCluster `hcl:"management_cluster,block" validate:"required"`
	// NetworkConfig: required
	NetworkConfig *vmwareengineprivatecloud.NetworkConfig `hcl:"network_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *vmwareengineprivatecloud.Timeouts `hcl:"timeouts,block"`
}
type vmwareenginePrivateCloudAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_vmwareengine_private_cloud.
func (vpc vmwareenginePrivateCloudAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("description"))
}

// Id returns a reference to field id of google_vmwareengine_private_cloud.
func (vpc vmwareenginePrivateCloudAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("id"))
}

// Location returns a reference to field location of google_vmwareengine_private_cloud.
func (vpc vmwareenginePrivateCloudAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("location"))
}

// Name returns a reference to field name of google_vmwareengine_private_cloud.
func (vpc vmwareenginePrivateCloudAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("name"))
}

// Project returns a reference to field project of google_vmwareengine_private_cloud.
func (vpc vmwareenginePrivateCloudAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("project"))
}

// State returns a reference to field state of google_vmwareengine_private_cloud.
func (vpc vmwareenginePrivateCloudAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("state"))
}

// Uid returns a reference to field uid of google_vmwareengine_private_cloud.
func (vpc vmwareenginePrivateCloudAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("uid"))
}

func (vpc vmwareenginePrivateCloudAttributes) Hcx() terra.ListValue[vmwareengineprivatecloud.HcxAttributes] {
	return terra.ReferenceAsList[vmwareengineprivatecloud.HcxAttributes](vpc.ref.Append("hcx"))
}

func (vpc vmwareenginePrivateCloudAttributes) Nsx() terra.ListValue[vmwareengineprivatecloud.NsxAttributes] {
	return terra.ReferenceAsList[vmwareengineprivatecloud.NsxAttributes](vpc.ref.Append("nsx"))
}

func (vpc vmwareenginePrivateCloudAttributes) Vcenter() terra.ListValue[vmwareengineprivatecloud.VcenterAttributes] {
	return terra.ReferenceAsList[vmwareengineprivatecloud.VcenterAttributes](vpc.ref.Append("vcenter"))
}

func (vpc vmwareenginePrivateCloudAttributes) ManagementCluster() terra.ListValue[vmwareengineprivatecloud.ManagementClusterAttributes] {
	return terra.ReferenceAsList[vmwareengineprivatecloud.ManagementClusterAttributes](vpc.ref.Append("management_cluster"))
}

func (vpc vmwareenginePrivateCloudAttributes) NetworkConfig() terra.ListValue[vmwareengineprivatecloud.NetworkConfigAttributes] {
	return terra.ReferenceAsList[vmwareengineprivatecloud.NetworkConfigAttributes](vpc.ref.Append("network_config"))
}

func (vpc vmwareenginePrivateCloudAttributes) Timeouts() vmwareengineprivatecloud.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vmwareengineprivatecloud.TimeoutsAttributes](vpc.ref.Append("timeouts"))
}

type vmwareenginePrivateCloudState struct {
	Description       string                                            `json:"description"`
	Id                string                                            `json:"id"`
	Location          string                                            `json:"location"`
	Name              string                                            `json:"name"`
	Project           string                                            `json:"project"`
	State             string                                            `json:"state"`
	Uid               string                                            `json:"uid"`
	Hcx               []vmwareengineprivatecloud.HcxState               `json:"hcx"`
	Nsx               []vmwareengineprivatecloud.NsxState               `json:"nsx"`
	Vcenter           []vmwareengineprivatecloud.VcenterState           `json:"vcenter"`
	ManagementCluster []vmwareengineprivatecloud.ManagementClusterState `json:"management_cluster"`
	NetworkConfig     []vmwareengineprivatecloud.NetworkConfigState     `json:"network_config"`
	Timeouts          *vmwareengineprivatecloud.TimeoutsState           `json:"timeouts"`
}