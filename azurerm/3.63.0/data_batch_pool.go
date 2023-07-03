// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	databatchpool "github.com/golingon/terraproviders/azurerm/3.63.0/databatchpool"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBatchPool creates a new instance of [DataBatchPool].
func NewDataBatchPool(name string, args DataBatchPoolArgs) *DataBatchPool {
	return &DataBatchPool{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBatchPool)(nil)

// DataBatchPool represents the Terraform data resource azurerm_batch_pool.
type DataBatchPool struct {
	Name string
	Args DataBatchPoolArgs
}

// DataSource returns the Terraform object type for [DataBatchPool].
func (bp *DataBatchPool) DataSource() string {
	return "azurerm_batch_pool"
}

// LocalName returns the local name for [DataBatchPool].
func (bp *DataBatchPool) LocalName() string {
	return bp.Name
}

// Configuration returns the configuration (args) for [DataBatchPool].
func (bp *DataBatchPool) Configuration() interface{} {
	return bp.Args
}

// Attributes returns the attributes for [DataBatchPool].
func (bp *DataBatchPool) Attributes() dataBatchPoolAttributes {
	return dataBatchPoolAttributes{ref: terra.ReferenceDataResource(bp)}
}

// DataBatchPoolArgs contains the configurations for azurerm_batch_pool.
type DataBatchPoolArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AutoScale: min=0
	AutoScale []databatchpool.AutoScale `hcl:"auto_scale,block" validate:"min=0"`
	// Certificate: min=0
	Certificate []databatchpool.Certificate `hcl:"certificate,block" validate:"min=0"`
	// ContainerConfiguration: min=0
	ContainerConfiguration []databatchpool.ContainerConfiguration `hcl:"container_configuration,block" validate:"min=0"`
	// DataDisks: min=0
	DataDisks []databatchpool.DataDisks `hcl:"data_disks,block" validate:"min=0"`
	// DiskEncryption: min=0
	DiskEncryption []databatchpool.DiskEncryption `hcl:"disk_encryption,block" validate:"min=0"`
	// Extensions: min=0
	Extensions []databatchpool.Extensions `hcl:"extensions,block" validate:"min=0"`
	// FixedScale: min=0
	FixedScale []databatchpool.FixedScale `hcl:"fixed_scale,block" validate:"min=0"`
	// Mount: min=0
	Mount []databatchpool.Mount `hcl:"mount,block" validate:"min=0"`
	// NetworkConfiguration: min=0
	NetworkConfiguration []databatchpool.NetworkConfiguration `hcl:"network_configuration,block" validate:"min=0"`
	// NodePlacement: min=0
	NodePlacement []databatchpool.NodePlacement `hcl:"node_placement,block" validate:"min=0"`
	// StartTask: min=0
	StartTask []databatchpool.StartTask `hcl:"start_task,block" validate:"min=0"`
	// StorageImageReference: min=0
	StorageImageReference []databatchpool.StorageImageReference `hcl:"storage_image_reference,block" validate:"min=0"`
	// TaskSchedulingPolicy: min=0
	TaskSchedulingPolicy []databatchpool.TaskSchedulingPolicy `hcl:"task_scheduling_policy,block" validate:"min=0"`
	// UserAccounts: min=0
	UserAccounts []databatchpool.UserAccounts `hcl:"user_accounts,block" validate:"min=0"`
	// Windows: min=0
	Windows []databatchpool.Windows `hcl:"windows,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *databatchpool.Timeouts `hcl:"timeouts,block"`
}
type dataBatchPoolAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_batch_pool.
func (bp dataBatchPoolAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("account_name"))
}

// DisplayName returns a reference to field display_name of azurerm_batch_pool.
func (bp dataBatchPoolAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_batch_pool.
func (bp dataBatchPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("id"))
}

// InterNodeCommunication returns a reference to field inter_node_communication of azurerm_batch_pool.
func (bp dataBatchPoolAttributes) InterNodeCommunication() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("inter_node_communication"))
}

// LicenseType returns a reference to field license_type of azurerm_batch_pool.
func (bp dataBatchPoolAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("license_type"))
}

// MaxTasksPerNode returns a reference to field max_tasks_per_node of azurerm_batch_pool.
func (bp dataBatchPoolAttributes) MaxTasksPerNode() terra.NumberValue {
	return terra.ReferenceAsNumber(bp.ref.Append("max_tasks_per_node"))
}

// Metadata returns a reference to field metadata of azurerm_batch_pool.
func (bp dataBatchPoolAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bp.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_batch_pool.
func (bp dataBatchPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("name"))
}

// NodeAgentSkuId returns a reference to field node_agent_sku_id of azurerm_batch_pool.
func (bp dataBatchPoolAttributes) NodeAgentSkuId() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("node_agent_sku_id"))
}

// OsDiskPlacement returns a reference to field os_disk_placement of azurerm_batch_pool.
func (bp dataBatchPoolAttributes) OsDiskPlacement() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("os_disk_placement"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_batch_pool.
func (bp dataBatchPoolAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("resource_group_name"))
}

// VmSize returns a reference to field vm_size of azurerm_batch_pool.
func (bp dataBatchPoolAttributes) VmSize() terra.StringValue {
	return terra.ReferenceAsString(bp.ref.Append("vm_size"))
}

func (bp dataBatchPoolAttributes) AutoScale() terra.ListValue[databatchpool.AutoScaleAttributes] {
	return terra.ReferenceAsList[databatchpool.AutoScaleAttributes](bp.ref.Append("auto_scale"))
}

func (bp dataBatchPoolAttributes) Certificate() terra.ListValue[databatchpool.CertificateAttributes] {
	return terra.ReferenceAsList[databatchpool.CertificateAttributes](bp.ref.Append("certificate"))
}

func (bp dataBatchPoolAttributes) ContainerConfiguration() terra.ListValue[databatchpool.ContainerConfigurationAttributes] {
	return terra.ReferenceAsList[databatchpool.ContainerConfigurationAttributes](bp.ref.Append("container_configuration"))
}

func (bp dataBatchPoolAttributes) DataDisks() terra.ListValue[databatchpool.DataDisksAttributes] {
	return terra.ReferenceAsList[databatchpool.DataDisksAttributes](bp.ref.Append("data_disks"))
}

func (bp dataBatchPoolAttributes) DiskEncryption() terra.ListValue[databatchpool.DiskEncryptionAttributes] {
	return terra.ReferenceAsList[databatchpool.DiskEncryptionAttributes](bp.ref.Append("disk_encryption"))
}

func (bp dataBatchPoolAttributes) Extensions() terra.ListValue[databatchpool.ExtensionsAttributes] {
	return terra.ReferenceAsList[databatchpool.ExtensionsAttributes](bp.ref.Append("extensions"))
}

func (bp dataBatchPoolAttributes) FixedScale() terra.ListValue[databatchpool.FixedScaleAttributes] {
	return terra.ReferenceAsList[databatchpool.FixedScaleAttributes](bp.ref.Append("fixed_scale"))
}

func (bp dataBatchPoolAttributes) Mount() terra.ListValue[databatchpool.MountAttributes] {
	return terra.ReferenceAsList[databatchpool.MountAttributes](bp.ref.Append("mount"))
}

func (bp dataBatchPoolAttributes) NetworkConfiguration() terra.ListValue[databatchpool.NetworkConfigurationAttributes] {
	return terra.ReferenceAsList[databatchpool.NetworkConfigurationAttributes](bp.ref.Append("network_configuration"))
}

func (bp dataBatchPoolAttributes) NodePlacement() terra.ListValue[databatchpool.NodePlacementAttributes] {
	return terra.ReferenceAsList[databatchpool.NodePlacementAttributes](bp.ref.Append("node_placement"))
}

func (bp dataBatchPoolAttributes) StartTask() terra.ListValue[databatchpool.StartTaskAttributes] {
	return terra.ReferenceAsList[databatchpool.StartTaskAttributes](bp.ref.Append("start_task"))
}

func (bp dataBatchPoolAttributes) StorageImageReference() terra.ListValue[databatchpool.StorageImageReferenceAttributes] {
	return terra.ReferenceAsList[databatchpool.StorageImageReferenceAttributes](bp.ref.Append("storage_image_reference"))
}

func (bp dataBatchPoolAttributes) TaskSchedulingPolicy() terra.ListValue[databatchpool.TaskSchedulingPolicyAttributes] {
	return terra.ReferenceAsList[databatchpool.TaskSchedulingPolicyAttributes](bp.ref.Append("task_scheduling_policy"))
}

func (bp dataBatchPoolAttributes) UserAccounts() terra.ListValue[databatchpool.UserAccountsAttributes] {
	return terra.ReferenceAsList[databatchpool.UserAccountsAttributes](bp.ref.Append("user_accounts"))
}

func (bp dataBatchPoolAttributes) Windows() terra.ListValue[databatchpool.WindowsAttributes] {
	return terra.ReferenceAsList[databatchpool.WindowsAttributes](bp.ref.Append("windows"))
}

func (bp dataBatchPoolAttributes) Timeouts() databatchpool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databatchpool.TimeoutsAttributes](bp.ref.Append("timeouts"))
}
