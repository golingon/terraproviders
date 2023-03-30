// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataebssnapshotids "github.com/golingon/terraproviders/aws/4.60.0/dataebssnapshotids"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataEbsSnapshotIds(name string, args DataEbsSnapshotIdsArgs) *DataEbsSnapshotIds {
	return &DataEbsSnapshotIds{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEbsSnapshotIds)(nil)

type DataEbsSnapshotIds struct {
	Name string
	Args DataEbsSnapshotIdsArgs
}

func (esi *DataEbsSnapshotIds) DataSource() string {
	return "aws_ebs_snapshot_ids"
}

func (esi *DataEbsSnapshotIds) LocalName() string {
	return esi.Name
}

func (esi *DataEbsSnapshotIds) Configuration() interface{} {
	return esi.Args
}

func (esi *DataEbsSnapshotIds) Attributes() dataEbsSnapshotIdsAttributes {
	return dataEbsSnapshotIdsAttributes{ref: terra.ReferenceDataResource(esi)}
}

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

func (esi dataEbsSnapshotIdsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(esi.ref.Append("id"))
}

func (esi dataEbsSnapshotIdsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](esi.ref.Append("ids"))
}

func (esi dataEbsSnapshotIdsAttributes) Owners() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](esi.ref.Append("owners"))
}

func (esi dataEbsSnapshotIdsAttributes) RestorableByUserIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](esi.ref.Append("restorable_by_user_ids"))
}

func (esi dataEbsSnapshotIdsAttributes) Filter() terra.SetValue[dataebssnapshotids.FilterAttributes] {
	return terra.ReferenceSet[dataebssnapshotids.FilterAttributes](esi.ref.Append("filter"))
}

func (esi dataEbsSnapshotIdsAttributes) Timeouts() dataebssnapshotids.TimeoutsAttributes {
	return terra.ReferenceSingle[dataebssnapshotids.TimeoutsAttributes](esi.ref.Append("timeouts"))
}
