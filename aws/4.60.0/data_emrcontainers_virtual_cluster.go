// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataemrcontainersvirtualcluster "github.com/golingon/terraproviders/aws/4.60.0/dataemrcontainersvirtualcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEmrcontainersVirtualCluster creates a new instance of [DataEmrcontainersVirtualCluster].
func NewDataEmrcontainersVirtualCluster(name string, args DataEmrcontainersVirtualClusterArgs) *DataEmrcontainersVirtualCluster {
	return &DataEmrcontainersVirtualCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEmrcontainersVirtualCluster)(nil)

// DataEmrcontainersVirtualCluster represents the Terraform data resource aws_emrcontainers_virtual_cluster.
type DataEmrcontainersVirtualCluster struct {
	Name string
	Args DataEmrcontainersVirtualClusterArgs
}

// DataSource returns the Terraform object type for [DataEmrcontainersVirtualCluster].
func (evc *DataEmrcontainersVirtualCluster) DataSource() string {
	return "aws_emrcontainers_virtual_cluster"
}

// LocalName returns the local name for [DataEmrcontainersVirtualCluster].
func (evc *DataEmrcontainersVirtualCluster) LocalName() string {
	return evc.Name
}

// Configuration returns the configuration (args) for [DataEmrcontainersVirtualCluster].
func (evc *DataEmrcontainersVirtualCluster) Configuration() interface{} {
	return evc.Args
}

// Attributes returns the attributes for [DataEmrcontainersVirtualCluster].
func (evc *DataEmrcontainersVirtualCluster) Attributes() dataEmrcontainersVirtualClusterAttributes {
	return dataEmrcontainersVirtualClusterAttributes{ref: terra.ReferenceDataResource(evc)}
}

// DataEmrcontainersVirtualClusterArgs contains the configurations for aws_emrcontainers_virtual_cluster.
type DataEmrcontainersVirtualClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualClusterId: string, required
	VirtualClusterId terra.StringValue `hcl:"virtual_cluster_id,attr" validate:"required"`
	// ContainerProvider: min=0
	ContainerProvider []dataemrcontainersvirtualcluster.ContainerProvider `hcl:"container_provider,block" validate:"min=0"`
}
type dataEmrcontainersVirtualClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_emrcontainers_virtual_cluster.
func (evc dataEmrcontainersVirtualClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(evc.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_emrcontainers_virtual_cluster.
func (evc dataEmrcontainersVirtualClusterAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(evc.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_emrcontainers_virtual_cluster.
func (evc dataEmrcontainersVirtualClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(evc.ref.Append("id"))
}

// Name returns a reference to field name of aws_emrcontainers_virtual_cluster.
func (evc dataEmrcontainersVirtualClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(evc.ref.Append("name"))
}

// State returns a reference to field state of aws_emrcontainers_virtual_cluster.
func (evc dataEmrcontainersVirtualClusterAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(evc.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_emrcontainers_virtual_cluster.
func (evc dataEmrcontainersVirtualClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](evc.ref.Append("tags"))
}

// VirtualClusterId returns a reference to field virtual_cluster_id of aws_emrcontainers_virtual_cluster.
func (evc dataEmrcontainersVirtualClusterAttributes) VirtualClusterId() terra.StringValue {
	return terra.ReferenceAsString(evc.ref.Append("virtual_cluster_id"))
}

func (evc dataEmrcontainersVirtualClusterAttributes) ContainerProvider() terra.ListValue[dataemrcontainersvirtualcluster.ContainerProviderAttributes] {
	return terra.ReferenceAsList[dataemrcontainersvirtualcluster.ContainerProviderAttributes](evc.ref.Append("container_provider"))
}
