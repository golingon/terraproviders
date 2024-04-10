// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataDnsManagedZone creates a new instance of [DataDnsManagedZone].
func NewDataDnsManagedZone(name string, args DataDnsManagedZoneArgs) *DataDnsManagedZone {
	return &DataDnsManagedZone{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDnsManagedZone)(nil)

// DataDnsManagedZone represents the Terraform data resource google_dns_managed_zone.
type DataDnsManagedZone struct {
	Name string
	Args DataDnsManagedZoneArgs
}

// DataSource returns the Terraform object type for [DataDnsManagedZone].
func (dmz *DataDnsManagedZone) DataSource() string {
	return "google_dns_managed_zone"
}

// LocalName returns the local name for [DataDnsManagedZone].
func (dmz *DataDnsManagedZone) LocalName() string {
	return dmz.Name
}

// Configuration returns the configuration (args) for [DataDnsManagedZone].
func (dmz *DataDnsManagedZone) Configuration() interface{} {
	return dmz.Args
}

// Attributes returns the attributes for [DataDnsManagedZone].
func (dmz *DataDnsManagedZone) Attributes() dataDnsManagedZoneAttributes {
	return dataDnsManagedZoneAttributes{ref: terra.ReferenceDataResource(dmz)}
}

// DataDnsManagedZoneArgs contains the configurations for google_dns_managed_zone.
type DataDnsManagedZoneArgs struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataDnsManagedZoneAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_dns_managed_zone.
func (dmz dataDnsManagedZoneAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dmz.ref.Append("description"))
}

// DnsName returns a reference to field dns_name of google_dns_managed_zone.
func (dmz dataDnsManagedZoneAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(dmz.ref.Append("dns_name"))
}

// Id returns a reference to field id of google_dns_managed_zone.
func (dmz dataDnsManagedZoneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmz.ref.Append("id"))
}

// ManagedZoneId returns a reference to field managed_zone_id of google_dns_managed_zone.
func (dmz dataDnsManagedZoneAttributes) ManagedZoneId() terra.NumberValue {
	return terra.ReferenceAsNumber(dmz.ref.Append("managed_zone_id"))
}

// Name returns a reference to field name of google_dns_managed_zone.
func (dmz dataDnsManagedZoneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dmz.ref.Append("name"))
}

// NameServers returns a reference to field name_servers of google_dns_managed_zone.
func (dmz dataDnsManagedZoneAttributes) NameServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dmz.ref.Append("name_servers"))
}

// Project returns a reference to field project of google_dns_managed_zone.
func (dmz dataDnsManagedZoneAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmz.ref.Append("project"))
}

// Visibility returns a reference to field visibility of google_dns_managed_zone.
func (dmz dataDnsManagedZoneAttributes) Visibility() terra.StringValue {
	return terra.ReferenceAsString(dmz.ref.Append("visibility"))
}
