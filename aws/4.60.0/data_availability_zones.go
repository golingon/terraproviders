// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataavailabilityzones "github.com/golingon/terraproviders/aws/4.60.0/dataavailabilityzones"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataAvailabilityZones(name string, args DataAvailabilityZonesArgs) *DataAvailabilityZones {
	return &DataAvailabilityZones{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAvailabilityZones)(nil)

type DataAvailabilityZones struct {
	Name string
	Args DataAvailabilityZonesArgs
}

func (az *DataAvailabilityZones) DataSource() string {
	return "aws_availability_zones"
}

func (az *DataAvailabilityZones) LocalName() string {
	return az.Name
}

func (az *DataAvailabilityZones) Configuration() interface{} {
	return az.Args
}

func (az *DataAvailabilityZones) Attributes() dataAvailabilityZonesAttributes {
	return dataAvailabilityZonesAttributes{ref: terra.ReferenceDataResource(az)}
}

type DataAvailabilityZonesArgs struct {
	// AllAvailabilityZones: bool, optional
	AllAvailabilityZones terra.BoolValue `hcl:"all_availability_zones,attr"`
	// ExcludeNames: set of string, optional
	ExcludeNames terra.SetValue[terra.StringValue] `hcl:"exclude_names,attr"`
	// ExcludeZoneIds: set of string, optional
	ExcludeZoneIds terra.SetValue[terra.StringValue] `hcl:"exclude_zone_ids,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Filter: min=0
	Filter []dataavailabilityzones.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataavailabilityzones.Timeouts `hcl:"timeouts,block"`
}
type dataAvailabilityZonesAttributes struct {
	ref terra.Reference
}

func (az dataAvailabilityZonesAttributes) AllAvailabilityZones() terra.BoolValue {
	return terra.ReferenceBool(az.ref.Append("all_availability_zones"))
}

func (az dataAvailabilityZonesAttributes) ExcludeNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](az.ref.Append("exclude_names"))
}

func (az dataAvailabilityZonesAttributes) ExcludeZoneIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](az.ref.Append("exclude_zone_ids"))
}

func (az dataAvailabilityZonesAttributes) GroupNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](az.ref.Append("group_names"))
}

func (az dataAvailabilityZonesAttributes) Id() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("id"))
}

func (az dataAvailabilityZonesAttributes) Names() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](az.ref.Append("names"))
}

func (az dataAvailabilityZonesAttributes) State() terra.StringValue {
	return terra.ReferenceString(az.ref.Append("state"))
}

func (az dataAvailabilityZonesAttributes) ZoneIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](az.ref.Append("zone_ids"))
}

func (az dataAvailabilityZonesAttributes) Filter() terra.SetValue[dataavailabilityzones.FilterAttributes] {
	return terra.ReferenceSet[dataavailabilityzones.FilterAttributes](az.ref.Append("filter"))
}

func (az dataAvailabilityZonesAttributes) Timeouts() dataavailabilityzones.TimeoutsAttributes {
	return terra.ReferenceSingle[dataavailabilityzones.TimeoutsAttributes](az.ref.Append("timeouts"))
}
