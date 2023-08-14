// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datamemorydbsnapshot "github.com/golingon/terraproviders/aws/5.12.0/datamemorydbsnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMemorydbSnapshot creates a new instance of [DataMemorydbSnapshot].
func NewDataMemorydbSnapshot(name string, args DataMemorydbSnapshotArgs) *DataMemorydbSnapshot {
	return &DataMemorydbSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMemorydbSnapshot)(nil)

// DataMemorydbSnapshot represents the Terraform data resource aws_memorydb_snapshot.
type DataMemorydbSnapshot struct {
	Name string
	Args DataMemorydbSnapshotArgs
}

// DataSource returns the Terraform object type for [DataMemorydbSnapshot].
func (ms *DataMemorydbSnapshot) DataSource() string {
	return "aws_memorydb_snapshot"
}

// LocalName returns the local name for [DataMemorydbSnapshot].
func (ms *DataMemorydbSnapshot) LocalName() string {
	return ms.Name
}

// Configuration returns the configuration (args) for [DataMemorydbSnapshot].
func (ms *DataMemorydbSnapshot) Configuration() interface{} {
	return ms.Args
}

// Attributes returns the attributes for [DataMemorydbSnapshot].
func (ms *DataMemorydbSnapshot) Attributes() dataMemorydbSnapshotAttributes {
	return dataMemorydbSnapshotAttributes{ref: terra.ReferenceDataResource(ms)}
}

// DataMemorydbSnapshotArgs contains the configurations for aws_memorydb_snapshot.
type DataMemorydbSnapshotArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ClusterConfiguration: min=0
	ClusterConfiguration []datamemorydbsnapshot.ClusterConfiguration `hcl:"cluster_configuration,block" validate:"min=0"`
}
type dataMemorydbSnapshotAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_memorydb_snapshot.
func (ms dataMemorydbSnapshotAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("arn"))
}

// ClusterName returns a reference to field cluster_name of aws_memorydb_snapshot.
func (ms dataMemorydbSnapshotAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("cluster_name"))
}

// Id returns a reference to field id of aws_memorydb_snapshot.
func (ms dataMemorydbSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_memorydb_snapshot.
func (ms dataMemorydbSnapshotAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("kms_key_arn"))
}

// Name returns a reference to field name of aws_memorydb_snapshot.
func (ms dataMemorydbSnapshotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("name"))
}

// Source returns a reference to field source of aws_memorydb_snapshot.
func (ms dataMemorydbSnapshotAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("source"))
}

// Tags returns a reference to field tags of aws_memorydb_snapshot.
func (ms dataMemorydbSnapshotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ms.ref.Append("tags"))
}

func (ms dataMemorydbSnapshotAttributes) ClusterConfiguration() terra.ListValue[datamemorydbsnapshot.ClusterConfigurationAttributes] {
	return terra.ReferenceAsList[datamemorydbsnapshot.ClusterConfigurationAttributes](ms.ref.Append("cluster_configuration"))
}
