// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeregioninstancegroupmanager "github.com/golingon/terraproviders/google/4.59.0/computeregioninstancegroupmanager"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeRegionInstanceGroupManager(name string, args ComputeRegionInstanceGroupManagerArgs) *ComputeRegionInstanceGroupManager {
	return &ComputeRegionInstanceGroupManager{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionInstanceGroupManager)(nil)

type ComputeRegionInstanceGroupManager struct {
	Name  string
	Args  ComputeRegionInstanceGroupManagerArgs
	state *computeRegionInstanceGroupManagerState
}

func (crigm *ComputeRegionInstanceGroupManager) Type() string {
	return "google_compute_region_instance_group_manager"
}

func (crigm *ComputeRegionInstanceGroupManager) LocalName() string {
	return crigm.Name
}

func (crigm *ComputeRegionInstanceGroupManager) Configuration() interface{} {
	return crigm.Args
}

func (crigm *ComputeRegionInstanceGroupManager) Attributes() computeRegionInstanceGroupManagerAttributes {
	return computeRegionInstanceGroupManagerAttributes{ref: terra.ReferenceResource(crigm)}
}

func (crigm *ComputeRegionInstanceGroupManager) ImportState(av io.Reader) error {
	crigm.state = &computeRegionInstanceGroupManagerState{}
	if err := json.NewDecoder(av).Decode(crigm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crigm.Type(), crigm.LocalName(), err)
	}
	return nil
}

func (crigm *ComputeRegionInstanceGroupManager) State() (*computeRegionInstanceGroupManagerState, bool) {
	return crigm.state, crigm.state != nil
}

func (crigm *ComputeRegionInstanceGroupManager) StateMust() *computeRegionInstanceGroupManagerState {
	if crigm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crigm.Type(), crigm.LocalName()))
	}
	return crigm.state
}

func (crigm *ComputeRegionInstanceGroupManager) DependOn() terra.Reference {
	return terra.ReferenceResource(crigm)
}

type ComputeRegionInstanceGroupManagerArgs struct {
	// BaseInstanceName: string, required
	BaseInstanceName terra.StringValue `hcl:"base_instance_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DistributionPolicyTargetShape: string, optional
	DistributionPolicyTargetShape terra.StringValue `hcl:"distribution_policy_target_shape,attr"`
	// DistributionPolicyZones: set of string, optional
	DistributionPolicyZones terra.SetValue[terra.StringValue] `hcl:"distribution_policy_zones,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ListManagedInstancesResults: string, optional
	ListManagedInstancesResults terra.StringValue `hcl:"list_managed_instances_results,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// TargetPools: set of string, optional
	TargetPools terra.SetValue[terra.StringValue] `hcl:"target_pools,attr"`
	// TargetSize: number, optional
	TargetSize terra.NumberValue `hcl:"target_size,attr"`
	// WaitForInstances: bool, optional
	WaitForInstances terra.BoolValue `hcl:"wait_for_instances,attr"`
	// WaitForInstancesStatus: string, optional
	WaitForInstancesStatus terra.StringValue `hcl:"wait_for_instances_status,attr"`
	// Status: min=0
	Status []computeregioninstancegroupmanager.Status `hcl:"status,block" validate:"min=0"`
	// AutoHealingPolicies: optional
	AutoHealingPolicies *computeregioninstancegroupmanager.AutoHealingPolicies `hcl:"auto_healing_policies,block"`
	// NamedPort: min=0
	NamedPort []computeregioninstancegroupmanager.NamedPort `hcl:"named_port,block" validate:"min=0"`
	// StatefulDisk: min=0
	StatefulDisk []computeregioninstancegroupmanager.StatefulDisk `hcl:"stateful_disk,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computeregioninstancegroupmanager.Timeouts `hcl:"timeouts,block"`
	// UpdatePolicy: optional
	UpdatePolicy *computeregioninstancegroupmanager.UpdatePolicy `hcl:"update_policy,block"`
	// Version: min=1
	Version []computeregioninstancegroupmanager.Version `hcl:"version,block" validate:"min=1"`
	// DependsOn contains resources that ComputeRegionInstanceGroupManager depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeRegionInstanceGroupManagerAttributes struct {
	ref terra.Reference
}

func (crigm computeRegionInstanceGroupManagerAttributes) BaseInstanceName() terra.StringValue {
	return terra.ReferenceString(crigm.ref.Append("base_instance_name"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) Description() terra.StringValue {
	return terra.ReferenceString(crigm.ref.Append("description"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) DistributionPolicyTargetShape() terra.StringValue {
	return terra.ReferenceString(crigm.ref.Append("distribution_policy_target_shape"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) DistributionPolicyZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](crigm.ref.Append("distribution_policy_zones"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceString(crigm.ref.Append("fingerprint"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) Id() terra.StringValue {
	return terra.ReferenceString(crigm.ref.Append("id"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) InstanceGroup() terra.StringValue {
	return terra.ReferenceString(crigm.ref.Append("instance_group"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) ListManagedInstancesResults() terra.StringValue {
	return terra.ReferenceString(crigm.ref.Append("list_managed_instances_results"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) Name() terra.StringValue {
	return terra.ReferenceString(crigm.ref.Append("name"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) Project() terra.StringValue {
	return terra.ReferenceString(crigm.ref.Append("project"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) Region() terra.StringValue {
	return terra.ReferenceString(crigm.ref.Append("region"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(crigm.ref.Append("self_link"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) TargetPools() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](crigm.ref.Append("target_pools"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) TargetSize() terra.NumberValue {
	return terra.ReferenceNumber(crigm.ref.Append("target_size"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) WaitForInstances() terra.BoolValue {
	return terra.ReferenceBool(crigm.ref.Append("wait_for_instances"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) WaitForInstancesStatus() terra.StringValue {
	return terra.ReferenceString(crigm.ref.Append("wait_for_instances_status"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) Status() terra.ListValue[computeregioninstancegroupmanager.StatusAttributes] {
	return terra.ReferenceList[computeregioninstancegroupmanager.StatusAttributes](crigm.ref.Append("status"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) AutoHealingPolicies() terra.ListValue[computeregioninstancegroupmanager.AutoHealingPoliciesAttributes] {
	return terra.ReferenceList[computeregioninstancegroupmanager.AutoHealingPoliciesAttributes](crigm.ref.Append("auto_healing_policies"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) NamedPort() terra.SetValue[computeregioninstancegroupmanager.NamedPortAttributes] {
	return terra.ReferenceSet[computeregioninstancegroupmanager.NamedPortAttributes](crigm.ref.Append("named_port"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) StatefulDisk() terra.SetValue[computeregioninstancegroupmanager.StatefulDiskAttributes] {
	return terra.ReferenceSet[computeregioninstancegroupmanager.StatefulDiskAttributes](crigm.ref.Append("stateful_disk"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) Timeouts() computeregioninstancegroupmanager.TimeoutsAttributes {
	return terra.ReferenceSingle[computeregioninstancegroupmanager.TimeoutsAttributes](crigm.ref.Append("timeouts"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) UpdatePolicy() terra.ListValue[computeregioninstancegroupmanager.UpdatePolicyAttributes] {
	return terra.ReferenceList[computeregioninstancegroupmanager.UpdatePolicyAttributes](crigm.ref.Append("update_policy"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) Version() terra.ListValue[computeregioninstancegroupmanager.VersionAttributes] {
	return terra.ReferenceList[computeregioninstancegroupmanager.VersionAttributes](crigm.ref.Append("version"))
}

type computeRegionInstanceGroupManagerState struct {
	BaseInstanceName              string                                                       `json:"base_instance_name"`
	Description                   string                                                       `json:"description"`
	DistributionPolicyTargetShape string                                                       `json:"distribution_policy_target_shape"`
	DistributionPolicyZones       []string                                                     `json:"distribution_policy_zones"`
	Fingerprint                   string                                                       `json:"fingerprint"`
	Id                            string                                                       `json:"id"`
	InstanceGroup                 string                                                       `json:"instance_group"`
	ListManagedInstancesResults   string                                                       `json:"list_managed_instances_results"`
	Name                          string                                                       `json:"name"`
	Project                       string                                                       `json:"project"`
	Region                        string                                                       `json:"region"`
	SelfLink                      string                                                       `json:"self_link"`
	TargetPools                   []string                                                     `json:"target_pools"`
	TargetSize                    float64                                                      `json:"target_size"`
	WaitForInstances              bool                                                         `json:"wait_for_instances"`
	WaitForInstancesStatus        string                                                       `json:"wait_for_instances_status"`
	Status                        []computeregioninstancegroupmanager.StatusState              `json:"status"`
	AutoHealingPolicies           []computeregioninstancegroupmanager.AutoHealingPoliciesState `json:"auto_healing_policies"`
	NamedPort                     []computeregioninstancegroupmanager.NamedPortState           `json:"named_port"`
	StatefulDisk                  []computeregioninstancegroupmanager.StatefulDiskState        `json:"stateful_disk"`
	Timeouts                      *computeregioninstancegroupmanager.TimeoutsState             `json:"timeouts"`
	UpdatePolicy                  []computeregioninstancegroupmanager.UpdatePolicyState        `json:"update_policy"`
	Version                       []computeregioninstancegroupmanager.VersionState             `json:"version"`
}
