// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datas3controlmultiregionaccesspoint "github.com/golingon/terraproviders/aws/5.11.0/datas3controlmultiregionaccesspoint"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataS3ControlMultiRegionAccessPoint creates a new instance of [DataS3ControlMultiRegionAccessPoint].
func NewDataS3ControlMultiRegionAccessPoint(name string, args DataS3ControlMultiRegionAccessPointArgs) *DataS3ControlMultiRegionAccessPoint {
	return &DataS3ControlMultiRegionAccessPoint{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataS3ControlMultiRegionAccessPoint)(nil)

// DataS3ControlMultiRegionAccessPoint represents the Terraform data resource aws_s3control_multi_region_access_point.
type DataS3ControlMultiRegionAccessPoint struct {
	Name string
	Args DataS3ControlMultiRegionAccessPointArgs
}

// DataSource returns the Terraform object type for [DataS3ControlMultiRegionAccessPoint].
func (smrap *DataS3ControlMultiRegionAccessPoint) DataSource() string {
	return "aws_s3control_multi_region_access_point"
}

// LocalName returns the local name for [DataS3ControlMultiRegionAccessPoint].
func (smrap *DataS3ControlMultiRegionAccessPoint) LocalName() string {
	return smrap.Name
}

// Configuration returns the configuration (args) for [DataS3ControlMultiRegionAccessPoint].
func (smrap *DataS3ControlMultiRegionAccessPoint) Configuration() interface{} {
	return smrap.Args
}

// Attributes returns the attributes for [DataS3ControlMultiRegionAccessPoint].
func (smrap *DataS3ControlMultiRegionAccessPoint) Attributes() dataS3ControlMultiRegionAccessPointAttributes {
	return dataS3ControlMultiRegionAccessPointAttributes{ref: terra.ReferenceDataResource(smrap)}
}

// DataS3ControlMultiRegionAccessPointArgs contains the configurations for aws_s3control_multi_region_access_point.
type DataS3ControlMultiRegionAccessPointArgs struct {
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicAccessBlock: min=0
	PublicAccessBlock []datas3controlmultiregionaccesspoint.PublicAccessBlock `hcl:"public_access_block,block" validate:"min=0"`
	// Regions: min=0
	Regions []datas3controlmultiregionaccesspoint.Regions `hcl:"regions,block" validate:"min=0"`
}
type dataS3ControlMultiRegionAccessPointAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_s3control_multi_region_access_point.
func (smrap dataS3ControlMultiRegionAccessPointAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("account_id"))
}

// Alias returns a reference to field alias of aws_s3control_multi_region_access_point.
func (smrap dataS3ControlMultiRegionAccessPointAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("alias"))
}

// Arn returns a reference to field arn of aws_s3control_multi_region_access_point.
func (smrap dataS3ControlMultiRegionAccessPointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_s3control_multi_region_access_point.
func (smrap dataS3ControlMultiRegionAccessPointAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("created_at"))
}

// DomainName returns a reference to field domain_name of aws_s3control_multi_region_access_point.
func (smrap dataS3ControlMultiRegionAccessPointAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("domain_name"))
}

// Id returns a reference to field id of aws_s3control_multi_region_access_point.
func (smrap dataS3ControlMultiRegionAccessPointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("id"))
}

// Name returns a reference to field name of aws_s3control_multi_region_access_point.
func (smrap dataS3ControlMultiRegionAccessPointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("name"))
}

// Status returns a reference to field status of aws_s3control_multi_region_access_point.
func (smrap dataS3ControlMultiRegionAccessPointAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("status"))
}

func (smrap dataS3ControlMultiRegionAccessPointAttributes) PublicAccessBlock() terra.ListValue[datas3controlmultiregionaccesspoint.PublicAccessBlockAttributes] {
	return terra.ReferenceAsList[datas3controlmultiregionaccesspoint.PublicAccessBlockAttributes](smrap.ref.Append("public_access_block"))
}

func (smrap dataS3ControlMultiRegionAccessPointAttributes) Regions() terra.ListValue[datas3controlmultiregionaccesspoint.RegionsAttributes] {
	return terra.ReferenceAsList[datas3controlmultiregionaccesspoint.RegionsAttributes](smrap.ref.Append("regions"))
}
