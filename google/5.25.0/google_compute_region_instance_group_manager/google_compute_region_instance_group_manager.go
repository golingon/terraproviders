// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_region_instance_group_manager

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

// Resource represents the Terraform resource google_compute_region_instance_group_manager.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeRegionInstanceGroupManagerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcrigm *Resource) Type() string {
	return "google_compute_region_instance_group_manager"
}

// LocalName returns the local name for [Resource].
func (gcrigm *Resource) LocalName() string {
	return gcrigm.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcrigm *Resource) Configuration() interface{} {
	return gcrigm.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcrigm *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcrigm)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcrigm *Resource) Dependencies() terra.Dependencies {
	return gcrigm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcrigm *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcrigm.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcrigm *Resource) Attributes() googleComputeRegionInstanceGroupManagerAttributes {
	return googleComputeRegionInstanceGroupManagerAttributes{ref: terra.ReferenceResource(gcrigm)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcrigm *Resource) ImportState(state io.Reader) error {
	gcrigm.state = &googleComputeRegionInstanceGroupManagerState{}
	if err := json.NewDecoder(state).Decode(gcrigm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcrigm.Type(), gcrigm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcrigm *Resource) State() (*googleComputeRegionInstanceGroupManagerState, bool) {
	return gcrigm.state, gcrigm.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcrigm *Resource) StateMust() *googleComputeRegionInstanceGroupManagerState {
	if gcrigm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcrigm.Type(), gcrigm.LocalName()))
	}
	return gcrigm.state
}

// Args contains the configurations for google_compute_region_instance_group_manager.
type Args struct {
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
	// AllInstancesConfig: optional
	AllInstancesConfig *AllInstancesConfig `hcl:"all_instances_config,block"`
	// AutoHealingPolicies: optional
	AutoHealingPolicies *AutoHealingPolicies `hcl:"auto_healing_policies,block"`
	// InstanceLifecyclePolicy: optional
	InstanceLifecyclePolicy *InstanceLifecyclePolicy `hcl:"instance_lifecycle_policy,block"`
	// NamedPort: min=0
	NamedPort []NamedPort `hcl:"named_port,block" validate:"min=0"`
	// StatefulDisk: min=0
	StatefulDisk []StatefulDisk `hcl:"stateful_disk,block" validate:"min=0"`
	// StatefulExternalIp: min=0
	StatefulExternalIp []StatefulExternalIp `hcl:"stateful_external_ip,block" validate:"min=0"`
	// StatefulInternalIp: min=0
	StatefulInternalIp []StatefulInternalIp `hcl:"stateful_internal_ip,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// UpdatePolicy: optional
	UpdatePolicy *UpdatePolicy `hcl:"update_policy,block"`
	// Version: min=1
	Version []Version `hcl:"version,block" validate:"min=1"`
}

type googleComputeRegionInstanceGroupManagerAttributes struct {
	ref terra.Reference
}

// BaseInstanceName returns a reference to field base_instance_name of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) BaseInstanceName() terra.StringValue {
	return terra.ReferenceAsString(gcrigm.ref.Append("base_instance_name"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gcrigm.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcrigm.ref.Append("description"))
}

// DistributionPolicyTargetShape returns a reference to field distribution_policy_target_shape of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) DistributionPolicyTargetShape() terra.StringValue {
	return terra.ReferenceAsString(gcrigm.ref.Append("distribution_policy_target_shape"))
}

// DistributionPolicyZones returns a reference to field distribution_policy_zones of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) DistributionPolicyZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gcrigm.ref.Append("distribution_policy_zones"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(gcrigm.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcrigm.ref.Append("id"))
}

// InstanceGroup returns a reference to field instance_group of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) InstanceGroup() terra.StringValue {
	return terra.ReferenceAsString(gcrigm.ref.Append("instance_group"))
}

// ListManagedInstancesResults returns a reference to field list_managed_instances_results of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) ListManagedInstancesResults() terra.StringValue {
	return terra.ReferenceAsString(gcrigm.ref.Append("list_managed_instances_results"))
}

// Name returns a reference to field name of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcrigm.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcrigm.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gcrigm.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gcrigm.ref.Append("self_link"))
}

// TargetPools returns a reference to field target_pools of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) TargetPools() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gcrigm.ref.Append("target_pools"))
}

// TargetSize returns a reference to field target_size of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) TargetSize() terra.NumberValue {
	return terra.ReferenceAsNumber(gcrigm.ref.Append("target_size"))
}

// WaitForInstances returns a reference to field wait_for_instances of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) WaitForInstances() terra.BoolValue {
	return terra.ReferenceAsBool(gcrigm.ref.Append("wait_for_instances"))
}

// WaitForInstancesStatus returns a reference to field wait_for_instances_status of google_compute_region_instance_group_manager.
func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) WaitForInstancesStatus() terra.StringValue {
	return terra.ReferenceAsString(gcrigm.ref.Append("wait_for_instances_status"))
}

func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) Status() terra.ListValue[StatusAttributes] {
	return terra.ReferenceAsList[StatusAttributes](gcrigm.ref.Append("status"))
}

func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) AllInstancesConfig() terra.ListValue[AllInstancesConfigAttributes] {
	return terra.ReferenceAsList[AllInstancesConfigAttributes](gcrigm.ref.Append("all_instances_config"))
}

func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) AutoHealingPolicies() terra.ListValue[AutoHealingPoliciesAttributes] {
	return terra.ReferenceAsList[AutoHealingPoliciesAttributes](gcrigm.ref.Append("auto_healing_policies"))
}

func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) InstanceLifecyclePolicy() terra.ListValue[InstanceLifecyclePolicyAttributes] {
	return terra.ReferenceAsList[InstanceLifecyclePolicyAttributes](gcrigm.ref.Append("instance_lifecycle_policy"))
}

func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) NamedPort() terra.SetValue[NamedPortAttributes] {
	return terra.ReferenceAsSet[NamedPortAttributes](gcrigm.ref.Append("named_port"))
}

func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) StatefulDisk() terra.SetValue[StatefulDiskAttributes] {
	return terra.ReferenceAsSet[StatefulDiskAttributes](gcrigm.ref.Append("stateful_disk"))
}

func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) StatefulExternalIp() terra.ListValue[StatefulExternalIpAttributes] {
	return terra.ReferenceAsList[StatefulExternalIpAttributes](gcrigm.ref.Append("stateful_external_ip"))
}

func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) StatefulInternalIp() terra.ListValue[StatefulInternalIpAttributes] {
	return terra.ReferenceAsList[StatefulInternalIpAttributes](gcrigm.ref.Append("stateful_internal_ip"))
}

func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcrigm.ref.Append("timeouts"))
}

func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) UpdatePolicy() terra.ListValue[UpdatePolicyAttributes] {
	return terra.ReferenceAsList[UpdatePolicyAttributes](gcrigm.ref.Append("update_policy"))
}

func (gcrigm googleComputeRegionInstanceGroupManagerAttributes) Version() terra.ListValue[VersionAttributes] {
	return terra.ReferenceAsList[VersionAttributes](gcrigm.ref.Append("version"))
}

type googleComputeRegionInstanceGroupManagerState struct {
	BaseInstanceName              string                         `json:"base_instance_name"`
	CreationTimestamp             string                         `json:"creation_timestamp"`
	Description                   string                         `json:"description"`
	DistributionPolicyTargetShape string                         `json:"distribution_policy_target_shape"`
	DistributionPolicyZones       []string                       `json:"distribution_policy_zones"`
	Fingerprint                   string                         `json:"fingerprint"`
	Id                            string                         `json:"id"`
	InstanceGroup                 string                         `json:"instance_group"`
	ListManagedInstancesResults   string                         `json:"list_managed_instances_results"`
	Name                          string                         `json:"name"`
	Project                       string                         `json:"project"`
	Region                        string                         `json:"region"`
	SelfLink                      string                         `json:"self_link"`
	TargetPools                   []string                       `json:"target_pools"`
	TargetSize                    float64                        `json:"target_size"`
	WaitForInstances              bool                           `json:"wait_for_instances"`
	WaitForInstancesStatus        string                         `json:"wait_for_instances_status"`
	Status                        []StatusState                  `json:"status"`
	AllInstancesConfig            []AllInstancesConfigState      `json:"all_instances_config"`
	AutoHealingPolicies           []AutoHealingPoliciesState     `json:"auto_healing_policies"`
	InstanceLifecyclePolicy       []InstanceLifecyclePolicyState `json:"instance_lifecycle_policy"`
	NamedPort                     []NamedPortState               `json:"named_port"`
	StatefulDisk                  []StatefulDiskState            `json:"stateful_disk"`
	StatefulExternalIp            []StatefulExternalIpState      `json:"stateful_external_ip"`
	StatefulInternalIp            []StatefulInternalIpState      `json:"stateful_internal_ip"`
	Timeouts                      *TimeoutsState                 `json:"timeouts"`
	UpdatePolicy                  []UpdatePolicyState            `json:"update_policy"`
	Version                       []VersionState                 `json:"version"`
}
