// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataComputeLbIpRanges creates a new instance of [DataComputeLbIpRanges].
func NewDataComputeLbIpRanges(name string, args DataComputeLbIpRangesArgs) *DataComputeLbIpRanges {
	return &DataComputeLbIpRanges{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeLbIpRanges)(nil)

// DataComputeLbIpRanges represents the Terraform data resource google_compute_lb_ip_ranges.
type DataComputeLbIpRanges struct {
	Name string
	Args DataComputeLbIpRangesArgs
}

// DataSource returns the Terraform object type for [DataComputeLbIpRanges].
func (clir *DataComputeLbIpRanges) DataSource() string {
	return "google_compute_lb_ip_ranges"
}

// LocalName returns the local name for [DataComputeLbIpRanges].
func (clir *DataComputeLbIpRanges) LocalName() string {
	return clir.Name
}

// Configuration returns the configuration (args) for [DataComputeLbIpRanges].
func (clir *DataComputeLbIpRanges) Configuration() interface{} {
	return clir.Args
}

// Attributes returns the attributes for [DataComputeLbIpRanges].
func (clir *DataComputeLbIpRanges) Attributes() dataComputeLbIpRangesAttributes {
	return dataComputeLbIpRangesAttributes{ref: terra.ReferenceDataResource(clir)}
}

// DataComputeLbIpRangesArgs contains the configurations for google_compute_lb_ip_ranges.
type DataComputeLbIpRangesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataComputeLbIpRangesAttributes struct {
	ref terra.Reference
}

// HttpSslTcpInternal returns a reference to field http_ssl_tcp_internal of google_compute_lb_ip_ranges.
func (clir dataComputeLbIpRangesAttributes) HttpSslTcpInternal() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](clir.ref.Append("http_ssl_tcp_internal"))
}

// Id returns a reference to field id of google_compute_lb_ip_ranges.
func (clir dataComputeLbIpRangesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(clir.ref.Append("id"))
}

// Network returns a reference to field network of google_compute_lb_ip_ranges.
func (clir dataComputeLbIpRangesAttributes) Network() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](clir.ref.Append("network"))
}
