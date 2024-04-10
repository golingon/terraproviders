// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataS3Bucket creates a new instance of [DataS3Bucket].
func NewDataS3Bucket(name string, args DataS3BucketArgs) *DataS3Bucket {
	return &DataS3Bucket{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataS3Bucket)(nil)

// DataS3Bucket represents the Terraform data resource aws_s3_bucket.
type DataS3Bucket struct {
	Name string
	Args DataS3BucketArgs
}

// DataSource returns the Terraform object type for [DataS3Bucket].
func (sb *DataS3Bucket) DataSource() string {
	return "aws_s3_bucket"
}

// LocalName returns the local name for [DataS3Bucket].
func (sb *DataS3Bucket) LocalName() string {
	return sb.Name
}

// Configuration returns the configuration (args) for [DataS3Bucket].
func (sb *DataS3Bucket) Configuration() interface{} {
	return sb.Args
}

// Attributes returns the attributes for [DataS3Bucket].
func (sb *DataS3Bucket) Attributes() dataS3BucketAttributes {
	return dataS3BucketAttributes{ref: terra.ReferenceDataResource(sb)}
}

// DataS3BucketArgs contains the configurations for aws_s3_bucket.
type DataS3BucketArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataS3BucketAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_s3_bucket.
func (sb dataS3BucketAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("arn"))
}

// Bucket returns a reference to field bucket of aws_s3_bucket.
func (sb dataS3BucketAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("bucket"))
}

// BucketDomainName returns a reference to field bucket_domain_name of aws_s3_bucket.
func (sb dataS3BucketAttributes) BucketDomainName() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("bucket_domain_name"))
}

// BucketRegionalDomainName returns a reference to field bucket_regional_domain_name of aws_s3_bucket.
func (sb dataS3BucketAttributes) BucketRegionalDomainName() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("bucket_regional_domain_name"))
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_s3_bucket.
func (sb dataS3BucketAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("hosted_zone_id"))
}

// Id returns a reference to field id of aws_s3_bucket.
func (sb dataS3BucketAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("id"))
}

// Region returns a reference to field region of aws_s3_bucket.
func (sb dataS3BucketAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("region"))
}

// WebsiteDomain returns a reference to field website_domain of aws_s3_bucket.
func (sb dataS3BucketAttributes) WebsiteDomain() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("website_domain"))
}

// WebsiteEndpoint returns a reference to field website_endpoint of aws_s3_bucket.
func (sb dataS3BucketAttributes) WebsiteEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("website_endpoint"))
}
