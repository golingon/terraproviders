// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataCloudfrontDistribution creates a new instance of [DataCloudfrontDistribution].
func NewDataCloudfrontDistribution(name string, args DataCloudfrontDistributionArgs) *DataCloudfrontDistribution {
	return &DataCloudfrontDistribution{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudfrontDistribution)(nil)

// DataCloudfrontDistribution represents the Terraform data resource aws_cloudfront_distribution.
type DataCloudfrontDistribution struct {
	Name string
	Args DataCloudfrontDistributionArgs
}

// DataSource returns the Terraform object type for [DataCloudfrontDistribution].
func (cd *DataCloudfrontDistribution) DataSource() string {
	return "aws_cloudfront_distribution"
}

// LocalName returns the local name for [DataCloudfrontDistribution].
func (cd *DataCloudfrontDistribution) LocalName() string {
	return cd.Name
}

// Configuration returns the configuration (args) for [DataCloudfrontDistribution].
func (cd *DataCloudfrontDistribution) Configuration() interface{} {
	return cd.Args
}

// Attributes returns the attributes for [DataCloudfrontDistribution].
func (cd *DataCloudfrontDistribution) Attributes() dataCloudfrontDistributionAttributes {
	return dataCloudfrontDistributionAttributes{ref: terra.ReferenceDataResource(cd)}
}

// DataCloudfrontDistributionArgs contains the configurations for aws_cloudfront_distribution.
type DataCloudfrontDistributionArgs struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataCloudfrontDistributionAttributes struct {
	ref terra.Reference
}

// Aliases returns a reference to field aliases of aws_cloudfront_distribution.
func (cd dataCloudfrontDistributionAttributes) Aliases() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cd.ref.Append("aliases"))
}

// Arn returns a reference to field arn of aws_cloudfront_distribution.
func (cd dataCloudfrontDistributionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("arn"))
}

// DomainName returns a reference to field domain_name of aws_cloudfront_distribution.
func (cd dataCloudfrontDistributionAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("domain_name"))
}

// Enabled returns a reference to field enabled of aws_cloudfront_distribution.
func (cd dataCloudfrontDistributionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cd.ref.Append("enabled"))
}

// Etag returns a reference to field etag of aws_cloudfront_distribution.
func (cd dataCloudfrontDistributionAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("etag"))
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_cloudfront_distribution.
func (cd dataCloudfrontDistributionAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("hosted_zone_id"))
}

// Id returns a reference to field id of aws_cloudfront_distribution.
func (cd dataCloudfrontDistributionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("id"))
}

// InProgressValidationBatches returns a reference to field in_progress_validation_batches of aws_cloudfront_distribution.
func (cd dataCloudfrontDistributionAttributes) InProgressValidationBatches() terra.NumberValue {
	return terra.ReferenceAsNumber(cd.ref.Append("in_progress_validation_batches"))
}

// LastModifiedTime returns a reference to field last_modified_time of aws_cloudfront_distribution.
func (cd dataCloudfrontDistributionAttributes) LastModifiedTime() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("last_modified_time"))
}

// Status returns a reference to field status of aws_cloudfront_distribution.
func (cd dataCloudfrontDistributionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_cloudfront_distribution.
func (cd dataCloudfrontDistributionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cd.ref.Append("tags"))
}

// WebAclId returns a reference to field web_acl_id of aws_cloudfront_distribution.
func (cd dataCloudfrontDistributionAttributes) WebAclId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("web_acl_id"))
}
