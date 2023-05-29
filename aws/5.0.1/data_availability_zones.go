// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataavailabilityzones "github.com/golingon/terraproviders/aws/5.0.1/dataavailabilityzones"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAvailabilityZones creates a new instance of [DataAvailabilityZones].
func NewDataAvailabilityZones(name string, args DataAvailabilityZonesArgs) *DataAvailabilityZones {
	return &DataAvailabilityZones{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAvailabilityZones)(nil)

// DataAvailabilityZones represents the Terraform data resource aws_availability_zones.
type DataAvailabilityZones struct {
	Name string
	Args DataAvailabilityZonesArgs
}

// DataSource returns the Terraform object type for [DataAvailabilityZones].
func (az *DataAvailabilityZones) DataSource() string {
	return "aws_availability_zones"
}

// LocalName returns the local name for [DataAvailabilityZones].
func (az *DataAvailabilityZones) LocalName() string {
	return az.Name
}

// Configuration returns the configuration (args) for [DataAvailabilityZones].
func (az *DataAvailabilityZones) Configuration() interface{} {
	return az.Args
}

// Attributes returns the attributes for [DataAvailabilityZones].
func (az *DataAvailabilityZones) Attributes() dataAvailabilityZonesAttributes {
	return dataAvailabilityZonesAttributes{ref: terra.ReferenceDataResource(az)}
}

// DataAvailabilityZonesArgs contains the configurations for aws_availability_zones.
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

// AllAvailabilityZones returns a reference to field all_availability_zones of aws_availability_zones.
func (az dataAvailabilityZonesAttributes) AllAvailabilityZones() terra.BoolValue {
	return terra.ReferenceAsBool(az.ref.Append("all_availability_zones"))
}

// ExcludeNames returns a reference to field exclude_names of aws_availability_zones.
func (az dataAvailabilityZonesAttributes) ExcludeNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](az.ref.Append("exclude_names"))
}

// ExcludeZoneIds returns a reference to field exclude_zone_ids of aws_availability_zones.
func (az dataAvailabilityZonesAttributes) ExcludeZoneIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](az.ref.Append("exclude_zone_ids"))
}

// GroupNames returns a reference to field group_names of aws_availability_zones.
func (az dataAvailabilityZonesAttributes) GroupNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](az.ref.Append("group_names"))
}

// Id returns a reference to field id of aws_availability_zones.
func (az dataAvailabilityZonesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("id"))
}

// Names returns a reference to field names of aws_availability_zones.
func (az dataAvailabilityZonesAttributes) Names() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](az.ref.Append("names"))
}

// State returns a reference to field state of aws_availability_zones.
func (az dataAvailabilityZonesAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("state"))
}

// ZoneIds returns a reference to field zone_ids of aws_availability_zones.
func (az dataAvailabilityZonesAttributes) ZoneIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](az.ref.Append("zone_ids"))
}

func (az dataAvailabilityZonesAttributes) Filter() terra.SetValue[dataavailabilityzones.FilterAttributes] {
	return terra.ReferenceAsSet[dataavailabilityzones.FilterAttributes](az.ref.Append("filter"))
}

func (az dataAvailabilityZonesAttributes) Timeouts() dataavailabilityzones.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataavailabilityzones.TimeoutsAttributes](az.ref.Append("timeouts"))
}