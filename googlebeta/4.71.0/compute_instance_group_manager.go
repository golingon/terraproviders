// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeinstancegroupmanager "github.com/golingon/terraproviders/googlebeta/4.71.0/computeinstancegroupmanager"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeInstanceGroupManager creates a new instance of [ComputeInstanceGroupManager].
func NewComputeInstanceGroupManager(name string, args ComputeInstanceGroupManagerArgs) *ComputeInstanceGroupManager {
	return &ComputeInstanceGroupManager{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeInstanceGroupManager)(nil)

// ComputeInstanceGroupManager represents the Terraform resource google_compute_instance_group_manager.
type ComputeInstanceGroupManager struct {
	Name      string
	Args      ComputeInstanceGroupManagerArgs
	state     *computeInstanceGroupManagerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeInstanceGroupManager].
func (cigm *ComputeInstanceGroupManager) Type() string {
	return "google_compute_instance_group_manager"
}

// LocalName returns the local name for [ComputeInstanceGroupManager].
func (cigm *ComputeInstanceGroupManager) LocalName() string {
	return cigm.Name
}

// Configuration returns the configuration (args) for [ComputeInstanceGroupManager].
func (cigm *ComputeInstanceGroupManager) Configuration() interface{} {
	return cigm.Args
}

// DependOn is used for other resources to depend on [ComputeInstanceGroupManager].
func (cigm *ComputeInstanceGroupManager) DependOn() terra.Reference {
	return terra.ReferenceResource(cigm)
}

// Dependencies returns the list of resources [ComputeInstanceGroupManager] depends_on.
func (cigm *ComputeInstanceGroupManager) Dependencies() terra.Dependencies {
	return cigm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeInstanceGroupManager].
func (cigm *ComputeInstanceGroupManager) LifecycleManagement() *terra.Lifecycle {
	return cigm.Lifecycle
}

// Attributes returns the attributes for [ComputeInstanceGroupManager].
func (cigm *ComputeInstanceGroupManager) Attributes() computeInstanceGroupManagerAttributes {
	return computeInstanceGroupManagerAttributes{ref: terra.ReferenceResource(cigm)}
}

// ImportState imports the given attribute values into [ComputeInstanceGroupManager]'s state.
func (cigm *ComputeInstanceGroupManager) ImportState(av io.Reader) error {
	cigm.state = &computeInstanceGroupManagerState{}
	if err := json.NewDecoder(av).Decode(cigm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cigm.Type(), cigm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeInstanceGroupManager] has state.
func (cigm *ComputeInstanceGroupManager) State() (*computeInstanceGroupManagerState, bool) {
	return cigm.state, cigm.state != nil
}

// StateMust returns the state for [ComputeInstanceGroupManager]. Panics if the state is nil.
func (cigm *ComputeInstanceGroupManager) StateMust() *computeInstanceGroupManagerState {
	if cigm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cigm.Type(), cigm.LocalName()))
	}
	return cigm.state
}

// ComputeInstanceGroupManagerArgs contains the configurations for google_compute_instance_group_manager.
type ComputeInstanceGroupManagerArgs struct {
	// BaseInstanceName: string, required
	BaseInstanceName terra.StringValue `hcl:"base_instance_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ListManagedInstancesResults: string, optional
	ListManagedInstancesResults terra.StringValue `hcl:"list_managed_instances_results,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// TargetPools: set of string, optional
	TargetPools terra.SetValue[terra.StringValue] `hcl:"target_pools,attr"`
	// TargetSize: number, optional
	TargetSize terra.NumberValue `hcl:"target_size,attr"`
	// WaitForInstances: bool, optional
	WaitForInstances terra.BoolValue `hcl:"wait_for_instances,attr"`
	// WaitForInstancesStatus: string, optional
	WaitForInstancesStatus terra.StringValue `hcl:"wait_for_instances_status,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Status: min=0
	Status []computeinstancegroupmanager.Status `hcl:"status,block" validate:"min=0"`
	// AllInstancesConfig: optional
	AllInstancesConfig *computeinstancegroupmanager.AllInstancesConfig `hcl:"all_instances_config,block"`
	// AutoHealingPolicies: optional
	AutoHealingPolicies *computeinstancegroupmanager.AutoHealingPolicies `hcl:"auto_healing_policies,block"`
	// InstanceLifecyclePolicy: optional
	InstanceLifecyclePolicy *computeinstancegroupmanager.InstanceLifecyclePolicy `hcl:"instance_lifecycle_policy,block"`
	// NamedPort: min=0
	NamedPort []computeinstancegroupmanager.NamedPort `hcl:"named_port,block" validate:"min=0"`
	// StatefulDisk: min=0
	StatefulDisk []computeinstancegroupmanager.StatefulDisk `hcl:"stateful_disk,block" validate:"min=0"`
	// StatefulExternalIp: min=0
	StatefulExternalIp []computeinstancegroupmanager.StatefulExternalIp `hcl:"stateful_external_ip,block" validate:"min=0"`
	// StatefulInternalIp: min=0
	StatefulInternalIp []computeinstancegroupmanager.StatefulInternalIp `hcl:"stateful_internal_ip,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computeinstancegroupmanager.Timeouts `hcl:"timeouts,block"`
	// UpdatePolicy: optional
	UpdatePolicy *computeinstancegroupmanager.UpdatePolicy `hcl:"update_policy,block"`
	// Version: min=1
	Version []computeinstancegroupmanager.Version `hcl:"version,block" validate:"min=1"`
}
type computeInstanceGroupManagerAttributes struct {
	ref terra.Reference
}

// BaseInstanceName returns a reference to field base_instance_name of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) BaseInstanceName() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("base_instance_name"))
}

// Description returns a reference to field description of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("description"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("id"))
}

// InstanceGroup returns a reference to field instance_group of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) InstanceGroup() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("instance_group"))
}

// ListManagedInstancesResults returns a reference to field list_managed_instances_results of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) ListManagedInstancesResults() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("list_managed_instances_results"))
}

// Name returns a reference to field name of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("name"))
}

// Operation returns a reference to field operation of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) Operation() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("operation"))
}

// Project returns a reference to field project of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("self_link"))
}

// TargetPools returns a reference to field target_pools of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) TargetPools() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cigm.ref.Append("target_pools"))
}

// TargetSize returns a reference to field target_size of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) TargetSize() terra.NumberValue {
	return terra.ReferenceAsNumber(cigm.ref.Append("target_size"))
}

// WaitForInstances returns a reference to field wait_for_instances of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) WaitForInstances() terra.BoolValue {
	return terra.ReferenceAsBool(cigm.ref.Append("wait_for_instances"))
}

// WaitForInstancesStatus returns a reference to field wait_for_instances_status of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) WaitForInstancesStatus() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("wait_for_instances_status"))
}

// Zone returns a reference to field zone of google_compute_instance_group_manager.
func (cigm computeInstanceGroupManagerAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("zone"))
}

func (cigm computeInstanceGroupManagerAttributes) Status() terra.ListValue[computeinstancegroupmanager.StatusAttributes] {
	return terra.ReferenceAsList[computeinstancegroupmanager.StatusAttributes](cigm.ref.Append("status"))
}

func (cigm computeInstanceGroupManagerAttributes) AllInstancesConfig() terra.ListValue[computeinstancegroupmanager.AllInstancesConfigAttributes] {
	return terra.ReferenceAsList[computeinstancegroupmanager.AllInstancesConfigAttributes](cigm.ref.Append("all_instances_config"))
}

func (cigm computeInstanceGroupManagerAttributes) AutoHealingPolicies() terra.ListValue[computeinstancegroupmanager.AutoHealingPoliciesAttributes] {
	return terra.ReferenceAsList[computeinstancegroupmanager.AutoHealingPoliciesAttributes](cigm.ref.Append("auto_healing_policies"))
}

func (cigm computeInstanceGroupManagerAttributes) InstanceLifecyclePolicy() terra.ListValue[computeinstancegroupmanager.InstanceLifecyclePolicyAttributes] {
	return terra.ReferenceAsList[computeinstancegroupmanager.InstanceLifecyclePolicyAttributes](cigm.ref.Append("instance_lifecycle_policy"))
}

func (cigm computeInstanceGroupManagerAttributes) NamedPort() terra.SetValue[computeinstancegroupmanager.NamedPortAttributes] {
	return terra.ReferenceAsSet[computeinstancegroupmanager.NamedPortAttributes](cigm.ref.Append("named_port"))
}

func (cigm computeInstanceGroupManagerAttributes) StatefulDisk() terra.SetValue[computeinstancegroupmanager.StatefulDiskAttributes] {
	return terra.ReferenceAsSet[computeinstancegroupmanager.StatefulDiskAttributes](cigm.ref.Append("stateful_disk"))
}

func (cigm computeInstanceGroupManagerAttributes) StatefulExternalIp() terra.ListValue[computeinstancegroupmanager.StatefulExternalIpAttributes] {
	return terra.ReferenceAsList[computeinstancegroupmanager.StatefulExternalIpAttributes](cigm.ref.Append("stateful_external_ip"))
}

func (cigm computeInstanceGroupManagerAttributes) StatefulInternalIp() terra.ListValue[computeinstancegroupmanager.StatefulInternalIpAttributes] {
	return terra.ReferenceAsList[computeinstancegroupmanager.StatefulInternalIpAttributes](cigm.ref.Append("stateful_internal_ip"))
}

func (cigm computeInstanceGroupManagerAttributes) Timeouts() computeinstancegroupmanager.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeinstancegroupmanager.TimeoutsAttributes](cigm.ref.Append("timeouts"))
}

func (cigm computeInstanceGroupManagerAttributes) UpdatePolicy() terra.ListValue[computeinstancegroupmanager.UpdatePolicyAttributes] {
	return terra.ReferenceAsList[computeinstancegroupmanager.UpdatePolicyAttributes](cigm.ref.Append("update_policy"))
}

func (cigm computeInstanceGroupManagerAttributes) Version() terra.ListValue[computeinstancegroupmanager.VersionAttributes] {
	return terra.ReferenceAsList[computeinstancegroupmanager.VersionAttributes](cigm.ref.Append("version"))
}

type computeInstanceGroupManagerState struct {
	BaseInstanceName            string                                                     `json:"base_instance_name"`
	Description                 string                                                     `json:"description"`
	Fingerprint                 string                                                     `json:"fingerprint"`
	Id                          string                                                     `json:"id"`
	InstanceGroup               string                                                     `json:"instance_group"`
	ListManagedInstancesResults string                                                     `json:"list_managed_instances_results"`
	Name                        string                                                     `json:"name"`
	Operation                   string                                                     `json:"operation"`
	Project                     string                                                     `json:"project"`
	SelfLink                    string                                                     `json:"self_link"`
	TargetPools                 []string                                                   `json:"target_pools"`
	TargetSize                  float64                                                    `json:"target_size"`
	WaitForInstances            bool                                                       `json:"wait_for_instances"`
	WaitForInstancesStatus      string                                                     `json:"wait_for_instances_status"`
	Zone                        string                                                     `json:"zone"`
	Status                      []computeinstancegroupmanager.StatusState                  `json:"status"`
	AllInstancesConfig          []computeinstancegroupmanager.AllInstancesConfigState      `json:"all_instances_config"`
	AutoHealingPolicies         []computeinstancegroupmanager.AutoHealingPoliciesState     `json:"auto_healing_policies"`
	InstanceLifecyclePolicy     []computeinstancegroupmanager.InstanceLifecyclePolicyState `json:"instance_lifecycle_policy"`
	NamedPort                   []computeinstancegroupmanager.NamedPortState               `json:"named_port"`
	StatefulDisk                []computeinstancegroupmanager.StatefulDiskState            `json:"stateful_disk"`
	StatefulExternalIp          []computeinstancegroupmanager.StatefulExternalIpState      `json:"stateful_external_ip"`
	StatefulInternalIp          []computeinstancegroupmanager.StatefulInternalIpState      `json:"stateful_internal_ip"`
	Timeouts                    *computeinstancegroupmanager.TimeoutsState                 `json:"timeouts"`
	UpdatePolicy                []computeinstancegroupmanager.UpdatePolicyState            `json:"update_policy"`
	Version                     []computeinstancegroupmanager.VersionState                 `json:"version"`
}
