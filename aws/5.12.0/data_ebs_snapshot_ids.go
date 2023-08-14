// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataebssnapshotids "github.com/golingon/terraproviders/aws/5.12.0/dataebssnapshotids"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEbsSnapshotIds creates a new instance of [DataEbsSnapshotIds].
func NewDataEbsSnapshotIds(name string, args DataEbsSnapshotIdsArgs) *DataEbsSnapshotIds {
	return &DataEbsSnapshotIds{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEbsSnapshotIds)(nil)

// DataEbsSnapshotIds represents the Terraform data resource aws_ebs_snapshot_ids.
type DataEbsSnapshotIds struct {
	Name string
	Args DataEbsSnapshotIdsArgs
}

// DataSource returns the Terraform object type for [DataEbsSnapshotIds].
func (esi *DataEbsSnapshotIds) DataSource() string {
	return "aws_ebs_snapshot_ids"
}

// LocalName returns the local name for [DataEbsSnapshotIds].
func (esi *DataEbsSnapshotIds) LocalName() string {
	return esi.Name
}

// Configuration returns the configuration (args) for [DataEbsSnapshotIds].
func (esi *DataEbsSnapshotIds) Configuration() interface{} {
	return esi.Args
}

// Attributes returns the attributes for [DataEbsSnapshotIds].
func (esi *DataEbsSnapshotIds) Attributes() dataEbsSnapshotIdsAttributes {
	return dataEbsSnapshotIdsAttributes{ref: terra.ReferenceDataResource(esi)}
}

// DataEbsSnapshotIdsArgs contains the configurations for aws_ebs_snapshot_ids.
type DataEbsSnapshotIdsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Owners: list of string, optional
	Owners terra.ListValue[terra.StringValue] `hcl:"owners,attr"`
	// RestorableByUserIds: list of string, optional
	RestorableByUserIds terra.ListValue[terra.StringValue] `hcl:"restorable_by_user_ids,attr"`
	// Filter: min=0
	Filter []dataebssnapshotids.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataebssnapshotids.Timeouts `hcl:"timeouts,block"`
}
type dataEbsSnapshotIdsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ebs_snapshot_ids.
func (esi dataEbsSnapshotIdsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(esi.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_ebs_snapshot_ids.
func (esi dataEbsSnapshotIdsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](esi.ref.Append("ids"))
}

// Owners returns a reference to field owners of aws_ebs_snapshot_ids.
func (esi dataEbsSnapshotIdsAttributes) Owners() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](esi.ref.Append("owners"))
}

// RestorableByUserIds returns a reference to field restorable_by_user_ids of aws_ebs_snapshot_ids.
func (esi dataEbsSnapshotIdsAttributes) RestorableByUserIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](esi.ref.Append("restorable_by_user_ids"))
}

func (esi dataEbsSnapshotIdsAttributes) Filter() terra.SetValue[dataebssnapshotids.FilterAttributes] {
	return terra.ReferenceAsSet[dataebssnapshotids.FilterAttributes](esi.ref.Append("filter"))
}

func (esi dataEbsSnapshotIdsAttributes) Timeouts() dataebssnapshotids.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataebssnapshotids.TimeoutsAttributes](esi.ref.Append("timeouts"))
}
