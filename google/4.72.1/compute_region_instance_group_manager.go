// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeregioninstancegroupmanager "github.com/golingon/terraproviders/google/4.72.1/computeregioninstancegroupmanager"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionInstanceGroupManager creates a new instance of [ComputeRegionInstanceGroupManager].
func NewComputeRegionInstanceGroupManager(name string, args ComputeRegionInstanceGroupManagerArgs) *ComputeRegionInstanceGroupManager {
	return &ComputeRegionInstanceGroupManager{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionInstanceGroupManager)(nil)

// ComputeRegionInstanceGroupManager represents the Terraform resource google_compute_region_instance_group_manager.
type ComputeRegionInstanceGroupManager struct {
	Name      string
	Args      ComputeRegionInstanceGroupManagerArgs
	state     *computeRegionInstanceGroupManagerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionInstanceGroupManager].
func (crigm *ComputeRegionInstanceGroupManager) Type() string {
	return "google_compute_region_instance_group_manager"
}

// LocalName returns the local name for [ComputeRegionInstanceGroupManager].
func (crigm *ComputeRegionInstanceGroupManager) LocalName() string {
	return crigm.Name
}

// Configuration returns the configuration (args) for [ComputeRegionInstanceGroupManager].
func (crigm *ComputeRegionInstanceGroupManager) Configuration() interface{} {
	return crigm.Args
}

// DependOn is used for other resources to depend on [ComputeRegionInstanceGroupManager].
func (crigm *ComputeRegionInstanceGroupManager) DependOn() terra.Reference {
	return terra.ReferenceResource(crigm)
}

// Dependencies returns the list of resources [ComputeRegionInstanceGroupManager] depends_on.
func (crigm *ComputeRegionInstanceGroupManager) Dependencies() terra.Dependencies {
	return crigm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionInstanceGroupManager].
func (crigm *ComputeRegionInstanceGroupManager) LifecycleManagement() *terra.Lifecycle {
	return crigm.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionInstanceGroupManager].
func (crigm *ComputeRegionInstanceGroupManager) Attributes() computeRegionInstanceGroupManagerAttributes {
	return computeRegionInstanceGroupManagerAttributes{ref: terra.ReferenceResource(crigm)}
}

// ImportState imports the given attribute values into [ComputeRegionInstanceGroupManager]'s state.
func (crigm *ComputeRegionInstanceGroupManager) ImportState(av io.Reader) error {
	crigm.state = &computeRegionInstanceGroupManagerState{}
	if err := json.NewDecoder(av).Decode(crigm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crigm.Type(), crigm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionInstanceGroupManager] has state.
func (crigm *ComputeRegionInstanceGroupManager) State() (*computeRegionInstanceGroupManagerState, bool) {
	return crigm.state, crigm.state != nil
}

// StateMust returns the state for [ComputeRegionInstanceGroupManager]. Panics if the state is nil.
func (crigm *ComputeRegionInstanceGroupManager) StateMust() *computeRegionInstanceGroupManagerState {
	if crigm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crigm.Type(), crigm.LocalName()))
	}
	return crigm.state
}

// ComputeRegionInstanceGroupManagerArgs contains the configurations for google_compute_region_instance_group_manager.
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
}
type computeRegionInstanceGroupManagerAttributes struct {
	ref terra.Reference
}

// BaseInstanceName returns a reference to field base_instance_name of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) BaseInstanceName() terra.StringValue {
	return terra.ReferenceAsString(crigm.ref.Append("base_instance_name"))
}

// Description returns a reference to field description of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crigm.ref.Append("description"))
}

// DistributionPolicyTargetShape returns a reference to field distribution_policy_target_shape of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) DistributionPolicyTargetShape() terra.StringValue {
	return terra.ReferenceAsString(crigm.ref.Append("distribution_policy_target_shape"))
}

// DistributionPolicyZones returns a reference to field distribution_policy_zones of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) DistributionPolicyZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crigm.ref.Append("distribution_policy_zones"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(crigm.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crigm.ref.Append("id"))
}

// InstanceGroup returns a reference to field instance_group of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) InstanceGroup() terra.StringValue {
	return terra.ReferenceAsString(crigm.ref.Append("instance_group"))
}

// ListManagedInstancesResults returns a reference to field list_managed_instances_results of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) ListManagedInstancesResults() terra.StringValue {
	return terra.ReferenceAsString(crigm.ref.Append("list_managed_instances_results"))
}

// Name returns a reference to field name of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crigm.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crigm.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crigm.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crigm.ref.Append("self_link"))
}

// TargetPools returns a reference to field target_pools of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) TargetPools() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crigm.ref.Append("target_pools"))
}

// TargetSize returns a reference to field target_size of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) TargetSize() terra.NumberValue {
	return terra.ReferenceAsNumber(crigm.ref.Append("target_size"))
}

// WaitForInstances returns a reference to field wait_for_instances of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) WaitForInstances() terra.BoolValue {
	return terra.ReferenceAsBool(crigm.ref.Append("wait_for_instances"))
}

// WaitForInstancesStatus returns a reference to field wait_for_instances_status of google_compute_region_instance_group_manager.
func (crigm computeRegionInstanceGroupManagerAttributes) WaitForInstancesStatus() terra.StringValue {
	return terra.ReferenceAsString(crigm.ref.Append("wait_for_instances_status"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) Status() terra.ListValue[computeregioninstancegroupmanager.StatusAttributes] {
	return terra.ReferenceAsList[computeregioninstancegroupmanager.StatusAttributes](crigm.ref.Append("status"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) AutoHealingPolicies() terra.ListValue[computeregioninstancegroupmanager.AutoHealingPoliciesAttributes] {
	return terra.ReferenceAsList[computeregioninstancegroupmanager.AutoHealingPoliciesAttributes](crigm.ref.Append("auto_healing_policies"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) NamedPort() terra.SetValue[computeregioninstancegroupmanager.NamedPortAttributes] {
	return terra.ReferenceAsSet[computeregioninstancegroupmanager.NamedPortAttributes](crigm.ref.Append("named_port"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) StatefulDisk() terra.SetValue[computeregioninstancegroupmanager.StatefulDiskAttributes] {
	return terra.ReferenceAsSet[computeregioninstancegroupmanager.StatefulDiskAttributes](crigm.ref.Append("stateful_disk"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) Timeouts() computeregioninstancegroupmanager.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregioninstancegroupmanager.TimeoutsAttributes](crigm.ref.Append("timeouts"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) UpdatePolicy() terra.ListValue[computeregioninstancegroupmanager.UpdatePolicyAttributes] {
	return terra.ReferenceAsList[computeregioninstancegroupmanager.UpdatePolicyAttributes](crigm.ref.Append("update_policy"))
}

func (crigm computeRegionInstanceGroupManagerAttributes) Version() terra.ListValue[computeregioninstancegroupmanager.VersionAttributes] {
	return terra.ReferenceAsList[computeregioninstancegroupmanager.VersionAttributes](crigm.ref.Append("version"))
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
