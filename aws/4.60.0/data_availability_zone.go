// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataavailabilityzone "github.com/golingon/terraproviders/aws/4.60.0/dataavailabilityzone"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataAvailabilityZone(name string, args DataAvailabilityZoneArgs) *DataAvailabilityZone {
	return &DataAvailabilityZone{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAvailabilityZone)(nil)

type DataAvailabilityZone struct {
	Name string
	Args DataAvailabilityZoneArgs
}

func (az *DataAvailabilityZone) DataSource() string {
	return "aws_availability_zone"
}

func (az *DataAvailabilityZone) LocalName() string {
	return az.Name
}

func (az *DataAvailabilityZone) Configuration() interface{} {
	return az.Args
}

func (az *DataAvailabilityZone) Attributes() dataAvailabilityZoneAttributes {
	return dataAvailabilityZoneAttributes{ref: terra.ReferenceDataResource(az)}
}

type DataAvailabilityZoneArgs struct {
	// AllAvailabilityZones: bool, optional
	AllAvailabilityZones terra.BoolValue `hcl:"all_availability_zones,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// ZoneId: string, optional
	ZoneId terra.StringValue `hcl:"zone_id,attr"`
	// Filter: min=0
	Filter []dataavailabilityzone.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataavailabilityzone.Timeouts `hcl:"timeouts,block"`
}
type dataAvailabilityZoneAttributes struct {
	ref terra.Reference
}

func (az dataAvailabilityZoneAttributes) AllAvailabilityZones() terra.BoolValue {
	return terra.ReferenceBool(az.ref.Append("all_availability_zones"))
}

func (az dataAvailabilityZoneAttributes) GroupName() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("group_name"))
}

func (az dataAvailabilityZoneAttributes) Id() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("id"))
}

func (az dataAvailabilityZoneAttributes) Name() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("name"))
}

func (az dataAvailabilityZoneAttributes) NameSuffix() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("name_suffix"))
}

func (az dataAvailabilityZoneAttributes) NetworkBorderGroup() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("network_border_group"))
}

func (az dataAvailabilityZoneAttributes) OptInStatus() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("opt_in_status"))
}

func (az dataAvailabilityZoneAttributes) ParentZoneId() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("parent_zone_id"))
}

func (az dataAvailabilityZoneAttributes) ParentZoneName() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("parent_zone_name"))
}

func (az dataAvailabilityZoneAttributes) Region() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("region"))
}

func (az dataAvailabilityZoneAttributes) State() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("state"))
}

func (az dataAvailabilityZoneAttributes) ZoneId() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("zone_id"))
}

func (az dataAvailabilityZoneAttributes) ZoneType() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("zone_type"))
}

func (az dataAvailabilityZoneAttributes) Filter() terra.SetValue[dataavailabilityzone.FilterAttributes] {
	return terra.ReferenceSet[dataavailabilityzone.FilterAttributes](az.ref.Append("filter"))
}

func (az dataAvailabilityZoneAttributes) Timeouts() dataavailabilityzone.TimeoutsAttributes {
	return terra.ReferenceSingle[dataavailabilityzone.TimeoutsAttributes](az.ref.Append("timeouts"))
}
