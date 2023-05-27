// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataavailabilityzone "github.com/golingon/terraproviders/aws/5.0.1/dataavailabilityzone"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAvailabilityZone creates a new instance of [DataAvailabilityZone].
func NewDataAvailabilityZone(name string, args DataAvailabilityZoneArgs) *DataAvailabilityZone {
	return &DataAvailabilityZone{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAvailabilityZone)(nil)

// DataAvailabilityZone represents the Terraform data resource aws_availability_zone.
type DataAvailabilityZone struct {
	Name string
	Args DataAvailabilityZoneArgs
}

// DataSource returns the Terraform object type for [DataAvailabilityZone].
func (az *DataAvailabilityZone) DataSource() string {
	return "aws_availability_zone"
}

// LocalName returns the local name for [DataAvailabilityZone].
func (az *DataAvailabilityZone) LocalName() string {
	return az.Name
}

// Configuration returns the configuration (args) for [DataAvailabilityZone].
func (az *DataAvailabilityZone) Configuration() interface{} {
	return az.Args
}

// Attributes returns the attributes for [DataAvailabilityZone].
func (az *DataAvailabilityZone) Attributes() dataAvailabilityZoneAttributes {
	return dataAvailabilityZoneAttributes{ref: terra.ReferenceDataResource(az)}
}

// DataAvailabilityZoneArgs contains the configurations for aws_availability_zone.
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

// AllAvailabilityZones returns a reference to field all_availability_zones of aws_availability_zone.
func (az dataAvailabilityZoneAttributes) AllAvailabilityZones() terra.BoolValue {
	return terra.ReferenceAsBool(az.ref.Append("all_availability_zones"))
}

// GroupName returns a reference to field group_name of aws_availability_zone.
func (az dataAvailabilityZoneAttributes) GroupName() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("group_name"))
}

// Id returns a reference to field id of aws_availability_zone.
func (az dataAvailabilityZoneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("id"))
}

// Name returns a reference to field name of aws_availability_zone.
func (az dataAvailabilityZoneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("name"))
}

// NameSuffix returns a reference to field name_suffix of aws_availability_zone.
func (az dataAvailabilityZoneAttributes) NameSuffix() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("name_suffix"))
}

// NetworkBorderGroup returns a reference to field network_border_group of aws_availability_zone.
func (az dataAvailabilityZoneAttributes) NetworkBorderGroup() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("network_border_group"))
}

// OptInStatus returns a reference to field opt_in_status of aws_availability_zone.
func (az dataAvailabilityZoneAttributes) OptInStatus() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("opt_in_status"))
}

// ParentZoneId returns a reference to field parent_zone_id of aws_availability_zone.
func (az dataAvailabilityZoneAttributes) ParentZoneId() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("parent_zone_id"))
}

// ParentZoneName returns a reference to field parent_zone_name of aws_availability_zone.
func (az dataAvailabilityZoneAttributes) ParentZoneName() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("parent_zone_name"))
}

// Region returns a reference to field region of aws_availability_zone.
func (az dataAvailabilityZoneAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("region"))
}

// State returns a reference to field state of aws_availability_zone.
func (az dataAvailabilityZoneAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("state"))
}

// ZoneId returns a reference to field zone_id of aws_availability_zone.
func (az dataAvailabilityZoneAttributes) ZoneId() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("zone_id"))
}

// ZoneType returns a reference to field zone_type of aws_availability_zone.
func (az dataAvailabilityZoneAttributes) ZoneType() terra.StringValue {
	return terra.ReferenceAsString(az.ref.Append("zone_type"))
}

func (az dataAvailabilityZoneAttributes) Filter() terra.SetValue[dataavailabilityzone.FilterAttributes] {
	return terra.ReferenceAsSet[dataavailabilityzone.FilterAttributes](az.ref.Append("filter"))
}

func (az dataAvailabilityZoneAttributes) Timeouts() dataavailabilityzone.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataavailabilityzone.TimeoutsAttributes](az.ref.Append("timeouts"))
}
