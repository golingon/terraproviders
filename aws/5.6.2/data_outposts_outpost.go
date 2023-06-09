// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataOutpostsOutpost creates a new instance of [DataOutpostsOutpost].
func NewDataOutpostsOutpost(name string, args DataOutpostsOutpostArgs) *DataOutpostsOutpost {
	return &DataOutpostsOutpost{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOutpostsOutpost)(nil)

// DataOutpostsOutpost represents the Terraform data resource aws_outposts_outpost.
type DataOutpostsOutpost struct {
	Name string
	Args DataOutpostsOutpostArgs
}

// DataSource returns the Terraform object type for [DataOutpostsOutpost].
func (oo *DataOutpostsOutpost) DataSource() string {
	return "aws_outposts_outpost"
}

// LocalName returns the local name for [DataOutpostsOutpost].
func (oo *DataOutpostsOutpost) LocalName() string {
	return oo.Name
}

// Configuration returns the configuration (args) for [DataOutpostsOutpost].
func (oo *DataOutpostsOutpost) Configuration() interface{} {
	return oo.Args
}

// Attributes returns the attributes for [DataOutpostsOutpost].
func (oo *DataOutpostsOutpost) Attributes() dataOutpostsOutpostAttributes {
	return dataOutpostsOutpostAttributes{ref: terra.ReferenceDataResource(oo)}
}

// DataOutpostsOutpostArgs contains the configurations for aws_outposts_outpost.
type DataOutpostsOutpostArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// OwnerId: string, optional
	OwnerId terra.StringValue `hcl:"owner_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataOutpostsOutpostAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_outposts_outpost.
func (oo dataOutpostsOutpostAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("arn"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_outposts_outpost.
func (oo dataOutpostsOutpostAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("availability_zone"))
}

// AvailabilityZoneId returns a reference to field availability_zone_id of aws_outposts_outpost.
func (oo dataOutpostsOutpostAttributes) AvailabilityZoneId() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("availability_zone_id"))
}

// Description returns a reference to field description of aws_outposts_outpost.
func (oo dataOutpostsOutpostAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("description"))
}

// Id returns a reference to field id of aws_outposts_outpost.
func (oo dataOutpostsOutpostAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("id"))
}

// LifecycleStatus returns a reference to field lifecycle_status of aws_outposts_outpost.
func (oo dataOutpostsOutpostAttributes) LifecycleStatus() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("lifecycle_status"))
}

// Name returns a reference to field name of aws_outposts_outpost.
func (oo dataOutpostsOutpostAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("name"))
}

// OwnerId returns a reference to field owner_id of aws_outposts_outpost.
func (oo dataOutpostsOutpostAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("owner_id"))
}

// SiteArn returns a reference to field site_arn of aws_outposts_outpost.
func (oo dataOutpostsOutpostAttributes) SiteArn() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("site_arn"))
}

// SiteId returns a reference to field site_id of aws_outposts_outpost.
func (oo dataOutpostsOutpostAttributes) SiteId() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("site_id"))
}

// SupportedHardwareType returns a reference to field supported_hardware_type of aws_outposts_outpost.
func (oo dataOutpostsOutpostAttributes) SupportedHardwareType() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("supported_hardware_type"))
}

// Tags returns a reference to field tags of aws_outposts_outpost.
func (oo dataOutpostsOutpostAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oo.ref.Append("tags"))
}
