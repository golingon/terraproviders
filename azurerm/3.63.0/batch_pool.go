// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	batchpool "github.com/golingon/terraproviders/azurerm/3.63.0/batchpool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBatchPool creates a new instance of [BatchPool].
func NewBatchPool(name string, args BatchPoolArgs) *BatchPool {
	return &BatchPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BatchPool)(nil)

// BatchPool represents the Terraform resource azurerm_batch_pool.
type BatchPool struct {
	Name      string
	Args      BatchPoolArgs
	state     *batchPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BatchPool].
func (bp *BatchPool) Type() string {
	return "azurerm_batch_pool"
}

// LocalName returns the local name for [BatchPool].
func (bp *BatchPool) LocalName() string {
	return bp.Name
}

// Configuration returns the configuration (args) for [BatchPool].
func (bp *BatchPool) Configuration() interface{} {
	return bp.Args
}

// DependOn is used for other resources to depend on [BatchPool].
func (bp *BatchPool) DependOn() terra.Reference {
	return terra.ReferenceResource(bp)
}

// Dependencies returns the list of resources [BatchPool] depends_on.
func (bp *BatchPool) Dependencies() terra.Dependencies {
	return bp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BatchPool].
func (bp *BatchPool) LifecycleManagement() *terra.Lifecycle {
	return bp.Lifecycle
}

// Attributes returns the attributes for [BatchPool].
func (bp *BatchPool) Attributes() batchPoolAttributes {
	return batchPoolAttributes{ref: terra.ReferenceResource(bp)}
}

// ImportState imports the given attribute values into [BatchPool]'s state.
func (bp *BatchPool) ImportState(av io.Reader) error {
	bp.state = &batchPoolState{}
	if err := json.NewDecoder(av).Decode(bp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bp.Type(), bp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BatchPool] has state.
func (bp *BatchPool) State() (*batchPoolState, bool) {
	return bp.state, bp.state != nil
}

// StateMust returns the state for [BatchPool]. Panics if the state is nil.
func (bp *BatchPool) StateMust() *batchPoolState {
	if bp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bp.Type(), bp.LocalName()))
	}
	return bp.state
}

// BatchPoolArgs contains the configurations for azurerm_batch_pool.
type BatchPoolArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InterNodeCommunication: string, optional
	InterNodeCommunication terra.StringValue `hcl:"inter_node_communication,attr"`
	// LicenseType: string, optional
	LicenseType terra.StringValue `hcl:"license_type,attr"`
	// MaxTasksPerNode: number, optional
	MaxTasksPerNode terra.NumberValue `hcl:"max_tasks_per_node,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NodeAgentSkuId: string, required
	NodeAgentSkuId terra.StringValue `hcl:"node_agent_sku_id,attr" validate:"required"`
	// OsDiskPlacement: string, optional
	OsDiskPlacement terra.StringValue `hcl:"os_disk_placement,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StopPendingResizeOperation: bool, optional
	StopPendingResizeOperation terra.BoolValue `hcl:"stop_pending_resize_operation,attr"`
	// TargetNodeCommunicationMode: string, optional
	TargetNodeCommunicationMode terra.StringValue `hcl:"target_node_communication_mode,attr"`
	// VmSize: string, required
	VmSize terra.StringValue `hcl:"vm_size,attr" validate:"required"`
	// AutoScale: optional
	AutoScale *batchpool.AutoScale `hcl:"auto_scale,block"`
	// Certificate: min=0
	Certificate []batchpool.Certificate `hcl:"certificate,block" validate:"min=0"`
	// ContainerConfiguration: optional
	ContainerConfiguration *batchpool.ContainerConfiguration `hcl:"container_configuration,block"`
	// DataDisks: min=0
	DataDisks []batchpool.DataDisks `hcl:"data_disks,block" validate:"min=0"`
	// DiskEncryption: min=0
	DiskEncryption []batchpool.DiskEncryption `hcl:"disk_encryption,block" validate:"min=0"`
	// Extensions: min=0
	Extensions []batchpool.Extensions `hcl:"extensions,block" validate:"min=0"`
	// FixedScale: optional
	FixedScale *batchpool.FixedScale `hcl:"fixed_scale,block"`
	// Identity: optional
	Identity *batchpool.Identity `hcl:"identity,block"`
	// Mount: min=0
	Mount []batchpool.Mount `hcl:"mount,block" validate:"min=0"`
	// NetworkConfiguration: optional
	NetworkConfiguration *batchpool.NetworkConfiguration `hcl:"network_configuration,block"`
	// NodePlacement: min=0
	NodePlacement []batchpool.NodePlacement `hcl:"node_placement,block" validate:"min=0"`
	// StartTask: optional
	StartTask *batchpool.StartTask `hcl:"start_task,block"`
	// StorageImageReference: required
	StorageImageReference *batchpool.StorageImageReference `hcl:"storage_image_reference,block" validate:"required"`
	// TaskSchedulingPolicy: min=0
	TaskSchedulingPolicy []batchpool.TaskSchedulingPolicy `hcl:"task_scheduling_policy,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *batchpool.Timeouts `hcl:"timeouts,block"`
	// UserAccounts: min=0
	UserAccounts []batchpool.UserAccounts `hcl:"user_accounts,block" validate:"min=0"`
	// Windows: min=0
	Windows []batchpool.Windows `hcl:"windows,block" validate:"min=0"`
}
type batchPoolAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_batch_pool.
func (bp batchPoolAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("account_name"))
}

// DisplayName returns a reference to field display_name of azurerm_batch_pool.
func (bp batchPoolAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_batch_pool.
func (bp batchPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("id"))
}

// InterNodeCommunication returns a reference to field inter_node_communication of azurerm_batch_pool.
func (bp batchPoolAttributes) InterNodeCommunication() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("inter_node_communication"))
}

// LicenseType returns a reference to field license_type of azurerm_batch_pool.
func (bp batchPoolAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("license_type"))
}

// MaxTasksPerNode returns a reference to field max_tasks_per_node of azurerm_batch_pool.
func (bp batchPoolAttributes) MaxTasksPerNode() terra.NumberValue {
	return terra.ReferenceAsNumber(bp.ref.Append("max_tasks_per_node"))
}

// Metadata returns a reference to field metadata of azurerm_batch_pool.
func (bp batchPoolAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bp.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_batch_pool.
func (bp batchPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("name"))
}

// NodeAgentSkuId returns a reference to field node_agent_sku_id of azurerm_batch_pool.
func (bp batchPoolAttributes) NodeAgentSkuId() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("node_agent_sku_id"))
}

// OsDiskPlacement returns a reference to field os_disk_placement of azurerm_batch_pool.
func (bp batchPoolAttributes) OsDiskPlacement() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("os_disk_placement"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_batch_pool.
func (bp batchPoolAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("resource_group_name"))
}

// StopPendingResizeOperation returns a reference to field stop_pending_resize_operation of azurerm_batch_pool.
func (bp batchPoolAttributes) StopPendingResizeOperation() terra.BoolValue {
	return terra.ReferenceAsBool(bp.ref.Append("stop_pending_resize_operation"))
}

// TargetNodeCommunicationMode returns a reference to field target_node_communication_mode of azurerm_batch_pool.
func (bp batchPoolAttributes) TargetNodeCommunicationMode() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("target_node_communication_mode"))
}

// VmSize returns a reference to field vm_size of azurerm_batch_pool.
func (bp batchPoolAttributes) VmSize() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("vm_size"))
}

func (bp batchPoolAttributes) AutoScale() terra.ListValue[batchpool.AutoScaleAttributes] {
	return terra.ReferenceAsList[batchpool.AutoScaleAttributes](bp.ref.Append("auto_scale"))
}

func (bp batchPoolAttributes) Certificate() terra.ListValue[batchpool.CertificateAttributes] {
	return terra.ReferenceAsList[batchpool.CertificateAttributes](bp.ref.Append("certificate"))
}

func (bp batchPoolAttributes) ContainerConfiguration() terra.ListValue[batchpool.ContainerConfigurationAttributes] {
	return terra.ReferenceAsList[batchpool.ContainerConfigurationAttributes](bp.ref.Append("container_configuration"))
}

func (bp batchPoolAttributes) DataDisks() terra.ListValue[batchpool.DataDisksAttributes] {
	return terra.ReferenceAsList[batchpool.DataDisksAttributes](bp.ref.Append("data_disks"))
}

func (bp batchPoolAttributes) DiskEncryption() terra.ListValue[batchpool.DiskEncryptionAttributes] {
	return terra.ReferenceAsList[batchpool.DiskEncryptionAttributes](bp.ref.Append("disk_encryption"))
}

func (bp batchPoolAttributes) Extensions() terra.ListValue[batchpool.ExtensionsAttributes] {
	return terra.ReferenceAsList[batchpool.ExtensionsAttributes](bp.ref.Append("extensions"))
}

func (bp batchPoolAttributes) FixedScale() terra.ListValue[batchpool.FixedScaleAttributes] {
	return terra.ReferenceAsList[batchpool.FixedScaleAttributes](bp.ref.Append("fixed_scale"))
}

func (bp batchPoolAttributes) Identity() terra.ListValue[batchpool.IdentityAttributes] {
	return terra.ReferenceAsList[batchpool.IdentityAttributes](bp.ref.Append("identity"))
}

func (bp batchPoolAttributes) Mount() terra.ListValue[batchpool.MountAttributes] {
	return terra.ReferenceAsList[batchpool.MountAttributes](bp.ref.Append("mount"))
}

func (bp batchPoolAttributes) NetworkConfiguration() terra.ListValue[batchpool.NetworkConfigurationAttributes] {
	return terra.ReferenceAsList[batchpool.NetworkConfigurationAttributes](bp.ref.Append("network_configuration"))
}

func (bp batchPoolAttributes) NodePlacement() terra.ListValue[batchpool.NodePlacementAttributes] {
	return terra.ReferenceAsList[batchpool.NodePlacementAttributes](bp.ref.Append("node_placement"))
}

func (bp batchPoolAttributes) StartTask() terra.ListValue[batchpool.StartTaskAttributes] {
	return terra.ReferenceAsList[batchpool.StartTaskAttributes](bp.ref.Append("start_task"))
}

func (bp batchPoolAttributes) StorageImageReference() terra.ListValue[batchpool.StorageImageReferenceAttributes] {
	return terra.ReferenceAsList[batchpool.StorageImageReferenceAttributes](bp.ref.Append("storage_image_reference"))
}

func (bp batchPoolAttributes) TaskSchedulingPolicy() terra.ListValue[batchpool.TaskSchedulingPolicyAttributes] {
	return terra.ReferenceAsList[batchpool.TaskSchedulingPolicyAttributes](bp.ref.Append("task_scheduling_policy"))
}

func (bp batchPoolAttributes) Timeouts() batchpool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[batchpool.TimeoutsAttributes](bp.ref.Append("timeouts"))
}

func (bp batchPoolAttributes) UserAccounts() terra.ListValue[batchpool.UserAccountsAttributes] {
	return terra.ReferenceAsList[batchpool.UserAccountsAttributes](bp.ref.Append("user_accounts"))
}

func (bp batchPoolAttributes) Windows() terra.ListValue[batchpool.WindowsAttributes] {
	return terra.ReferenceAsList[batchpool.WindowsAttributes](bp.ref.Append("windows"))
}

type batchPoolState struct {
	AccountName                 string                                  `json:"account_name"`
	DisplayName                 string                                  `json:"display_name"`
	Id                          string                                  `json:"id"`
	InterNodeCommunication      string                                  `json:"inter_node_communication"`
	LicenseType                 string                                  `json:"license_type"`
	MaxTasksPerNode             float64                                 `json:"max_tasks_per_node"`
	Metadata                    map[string]string                       `json:"metadata"`
	Name                        string                                  `json:"name"`
	NodeAgentSkuId              string                                  `json:"node_agent_sku_id"`
	OsDiskPlacement             string                                  `json:"os_disk_placement"`
	ResourceGroupName           string                                  `json:"resource_group_name"`
	StopPendingResizeOperation  bool                                    `json:"stop_pending_resize_operation"`
	TargetNodeCommunicationMode string                                  `json:"target_node_communication_mode"`
	VmSize                      string                                  `json:"vm_size"`
	AutoScale                   []batchpool.AutoScaleState              `json:"auto_scale"`
	Certificate                 []batchpool.CertificateState            `json:"certificate"`
	ContainerConfiguration      []batchpool.ContainerConfigurationState `json:"container_configuration"`
	DataDisks                   []batchpool.DataDisksState              `json:"data_disks"`
	DiskEncryption              []batchpool.DiskEncryptionState         `json:"disk_encryption"`
	Extensions                  []batchpool.ExtensionsState             `json:"extensions"`
	FixedScale                  []batchpool.FixedScaleState             `json:"fixed_scale"`
	Identity                    []batchpool.IdentityState               `json:"identity"`
	Mount                       []batchpool.MountState                  `json:"mount"`
	NetworkConfiguration        []batchpool.NetworkConfigurationState   `json:"network_configuration"`
	NodePlacement               []batchpool.NodePlacementState          `json:"node_placement"`
	StartTask                   []batchpool.StartTaskState              `json:"start_task"`
	StorageImageReference       []batchpool.StorageImageReferenceState  `json:"storage_image_reference"`
	TaskSchedulingPolicy        []batchpool.TaskSchedulingPolicyState   `json:"task_scheduling_policy"`
	Timeouts                    *batchpool.TimeoutsState                `json:"timeouts"`
	UserAccounts                []batchpool.UserAccountsState           `json:"user_accounts"`
	Windows                     []batchpool.WindowsState                `json:"windows"`
}
