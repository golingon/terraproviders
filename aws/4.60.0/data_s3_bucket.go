// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataS3Bucket(name string, args DataS3BucketArgs) *DataS3Bucket {
	return &DataS3Bucket{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataS3Bucket)(nil)

type DataS3Bucket struct {
	Name string
	Args DataS3BucketArgs
}

func (sb *DataS3Bucket) DataSource() string {
	return "aws_s3_bucket"
}

func (sb *DataS3Bucket) LocalName() string {
	return sb.Name
}

func (sb *DataS3Bucket) Configuration() interface{} {
	return sb.Args
}

func (sb *DataS3Bucket) Attributes() dataS3BucketAttributes {
	return dataS3BucketAttributes{ref: terra.ReferenceDataResource(sb)}
}

type DataS3BucketArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataS3BucketAttributes struct {
	ref terra.Reference
}

func (sb dataS3BucketAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(sb.ref.Append("arn"))
}

func (sb dataS3BucketAttributes) Bucket() terra.StringValue {
	return terra.ReferenceString(sb.ref.Append("bucket"))
}

func (sb dataS3BucketAttributes) BucketDomainName() terra.StringValue {
	return terra.ReferenceString(sb.ref.Append("bucket_domain_name"))
}

func (sb dataS3BucketAttributes) BucketRegionalDomainName() terra.StringValue {
	return terra.ReferenceString(sb.ref.Append("bucket_regional_domain_name"))
}

func (sb dataS3BucketAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceString(sb.ref.Append("hosted_zone_id"))
}

func (sb dataS3BucketAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sb.ref.Append("id"))
}

func (sb dataS3BucketAttributes) Region() terra.StringValue {
	return terra.ReferenceString(sb.ref.Append("region"))
}

func (sb dataS3BucketAttributes) WebsiteDomain() terra.StringValue {
	return terra.ReferenceString(sb.ref.Append("website_domain"))
}

func (sb dataS3BucketAttributes) WebsiteEndpoint() terra.StringValue {
	return terra.ReferenceString(sb.ref.Append("website_endpoint"))
}
