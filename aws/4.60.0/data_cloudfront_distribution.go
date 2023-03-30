// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataCloudfrontDistribution(name string, args DataCloudfrontDistributionArgs) *DataCloudfrontDistribution {
	return &DataCloudfrontDistribution{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudfrontDistribution)(nil)

type DataCloudfrontDistribution struct {
	Name string
	Args DataCloudfrontDistributionArgs
}

func (cd *DataCloudfrontDistribution) DataSource() string {
	return "aws_cloudfront_distribution"
}

func (cd *DataCloudfrontDistribution) LocalName() string {
	return cd.Name
}

func (cd *DataCloudfrontDistribution) Configuration() interface{} {
	return cd.Args
}

func (cd *DataCloudfrontDistribution) Attributes() dataCloudfrontDistributionAttributes {
	return dataCloudfrontDistributionAttributes{ref: terra.ReferenceDataResource(cd)}
}

type DataCloudfrontDistributionArgs struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataCloudfrontDistributionAttributes struct {
	ref terra.Reference
}

func (cd dataCloudfrontDistributionAttributes) Aliases() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cd.ref.Append("aliases"))
}

func (cd dataCloudfrontDistributionAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("arn"))
}

func (cd dataCloudfrontDistributionAttributes) DomainName() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("domain_name"))
}

func (cd dataCloudfrontDistributionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(cd.ref.Append("enabled"))
}

func (cd dataCloudfrontDistributionAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("etag"))
}

func (cd dataCloudfrontDistributionAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("hosted_zone_id"))
}

func (cd dataCloudfrontDistributionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("id"))
}

func (cd dataCloudfrontDistributionAttributes) InProgressValidationBatches() terra.NumberValue {
	return terra.ReferenceNumber(cd.ref.Append("in_progress_validation_batches"))
}

func (cd dataCloudfrontDistributionAttributes) LastModifiedTime() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("last_modified_time"))
}

func (cd dataCloudfrontDistributionAttributes) Status() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("status"))
}

func (cd dataCloudfrontDistributionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cd.ref.Append("tags"))
}
