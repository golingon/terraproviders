// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDatastreamStaticIps creates a new instance of [DataDatastreamStaticIps].
func NewDataDatastreamStaticIps(name string, args DataDatastreamStaticIpsArgs) *DataDatastreamStaticIps {
	return &DataDatastreamStaticIps{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDatastreamStaticIps)(nil)

// DataDatastreamStaticIps represents the Terraform data resource google_datastream_static_ips.
type DataDatastreamStaticIps struct {
	Name string
	Args DataDatastreamStaticIpsArgs
}

// DataSource returns the Terraform object type for [DataDatastreamStaticIps].
func (dsi *DataDatastreamStaticIps) DataSource() string {
	return "google_datastream_static_ips"
}

// LocalName returns the local name for [DataDatastreamStaticIps].
func (dsi *DataDatastreamStaticIps) LocalName() string {
	return dsi.Name
}

// Configuration returns the configuration (args) for [DataDatastreamStaticIps].
func (dsi *DataDatastreamStaticIps) Configuration() interface{} {
	return dsi.Args
}

// Attributes returns the attributes for [DataDatastreamStaticIps].
func (dsi *DataDatastreamStaticIps) Attributes() dataDatastreamStaticIpsAttributes {
	return dataDatastreamStaticIpsAttributes{ref: terra.ReferenceDataResource(dsi)}
}

// DataDatastreamStaticIpsArgs contains the configurations for google_datastream_static_ips.
type DataDatastreamStaticIpsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataDatastreamStaticIpsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_datastream_static_ips.
func (dsi dataDatastreamStaticIpsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsi.ref.Append("id"))
}

// Location returns a reference to field location of google_datastream_static_ips.
func (dsi dataDatastreamStaticIpsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dsi.ref.Append("location"))
}

// Project returns a reference to field project of google_datastream_static_ips.
func (dsi dataDatastreamStaticIpsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dsi.ref.Append("project"))
}

// StaticIps returns a reference to field static_ips of google_datastream_static_ips.
func (dsi dataDatastreamStaticIpsAttributes) StaticIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dsi.ref.Append("static_ips"))
}