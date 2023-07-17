// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datakubernetesnodepoolsnapshot "github.com/golingon/terraproviders/azurerm/3.65.0/datakubernetesnodepoolsnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKubernetesNodePoolSnapshot creates a new instance of [DataKubernetesNodePoolSnapshot].
func NewDataKubernetesNodePoolSnapshot(name string, args DataKubernetesNodePoolSnapshotArgs) *DataKubernetesNodePoolSnapshot {
	return &DataKubernetesNodePoolSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKubernetesNodePoolSnapshot)(nil)

// DataKubernetesNodePoolSnapshot represents the Terraform data resource azurerm_kubernetes_node_pool_snapshot.
type DataKubernetesNodePoolSnapshot struct {
	Name string
	Args DataKubernetesNodePoolSnapshotArgs
}

// DataSource returns the Terraform object type for [DataKubernetesNodePoolSnapshot].
func (knps *DataKubernetesNodePoolSnapshot) DataSource() string {
	return "azurerm_kubernetes_node_pool_snapshot"
}

// LocalName returns the local name for [DataKubernetesNodePoolSnapshot].
func (knps *DataKubernetesNodePoolSnapshot) LocalName() string {
	return knps.Name
}

// Configuration returns the configuration (args) for [DataKubernetesNodePoolSnapshot].
func (knps *DataKubernetesNodePoolSnapshot) Configuration() interface{} {
	return knps.Args
}

// Attributes returns the attributes for [DataKubernetesNodePoolSnapshot].
func (knps *DataKubernetesNodePoolSnapshot) Attributes() dataKubernetesNodePoolSnapshotAttributes {
	return dataKubernetesNodePoolSnapshotAttributes{ref: terra.ReferenceDataResource(knps)}
}

// DataKubernetesNodePoolSnapshotArgs contains the configurations for azurerm_kubernetes_node_pool_snapshot.
type DataKubernetesNodePoolSnapshotArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datakubernetesnodepoolsnapshot.Timeouts `hcl:"timeouts,block"`
}
type dataKubernetesNodePoolSnapshotAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_kubernetes_node_pool_snapshot.
func (knps dataKubernetesNodePoolSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(knps.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_kubernetes_node_pool_snapshot.
func (knps dataKubernetesNodePoolSnapshotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(knps.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kubernetes_node_pool_snapshot.
func (knps dataKubernetesNodePoolSnapshotAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(knps.ref.Append("resource_group_name"))
}

// SourceNodePoolId returns a reference to field source_node_pool_id of azurerm_kubernetes_node_pool_snapshot.
func (knps dataKubernetesNodePoolSnapshotAttributes) SourceNodePoolId() terra.StringValue {
	return terra.ReferenceAsString(knps.ref.Append("source_node_pool_id"))
}

// Tags returns a reference to field tags of azurerm_kubernetes_node_pool_snapshot.
func (knps dataKubernetesNodePoolSnapshotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](knps.ref.Append("tags"))
}

func (knps dataKubernetesNodePoolSnapshotAttributes) Timeouts() datakubernetesnodepoolsnapshot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datakubernetesnodepoolsnapshot.TimeoutsAttributes](knps.ref.Append("timeouts"))
}
