// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataWafIpset creates a new instance of [DataWafIpset].
func NewDataWafIpset(name string, args DataWafIpsetArgs) *DataWafIpset {
	return &DataWafIpset{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWafIpset)(nil)

// DataWafIpset represents the Terraform data resource aws_waf_ipset.
type DataWafIpset struct {
	Name string
	Args DataWafIpsetArgs
}

// DataSource returns the Terraform object type for [DataWafIpset].
func (wi *DataWafIpset) DataSource() string {
	return "aws_waf_ipset"
}

// LocalName returns the local name for [DataWafIpset].
func (wi *DataWafIpset) LocalName() string {
	return wi.Name
}

// Configuration returns the configuration (args) for [DataWafIpset].
func (wi *DataWafIpset) Configuration() interface{} {
	return wi.Args
}

// Attributes returns the attributes for [DataWafIpset].
func (wi *DataWafIpset) Attributes() dataWafIpsetAttributes {
	return dataWafIpsetAttributes{ref: terra.ReferenceDataResource(wi)}
}

// DataWafIpsetArgs contains the configurations for aws_waf_ipset.
type DataWafIpsetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type dataWafIpsetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_waf_ipset.
func (wi dataWafIpsetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wi.ref.Append("id"))
}

// Name returns a reference to field name of aws_waf_ipset.
func (wi dataWafIpsetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wi.ref.Append("name"))
}
