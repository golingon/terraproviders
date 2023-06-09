// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataCloudfrontLogDeliveryCanonicalUserId creates a new instance of [DataCloudfrontLogDeliveryCanonicalUserId].
func NewDataCloudfrontLogDeliveryCanonicalUserId(name string, args DataCloudfrontLogDeliveryCanonicalUserIdArgs) *DataCloudfrontLogDeliveryCanonicalUserId {
	return &DataCloudfrontLogDeliveryCanonicalUserId{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudfrontLogDeliveryCanonicalUserId)(nil)

// DataCloudfrontLogDeliveryCanonicalUserId represents the Terraform data resource aws_cloudfront_log_delivery_canonical_user_id.
type DataCloudfrontLogDeliveryCanonicalUserId struct {
	Name string
	Args DataCloudfrontLogDeliveryCanonicalUserIdArgs
}

// DataSource returns the Terraform object type for [DataCloudfrontLogDeliveryCanonicalUserId].
func (cldcui *DataCloudfrontLogDeliveryCanonicalUserId) DataSource() string {
	return "aws_cloudfront_log_delivery_canonical_user_id"
}

// LocalName returns the local name for [DataCloudfrontLogDeliveryCanonicalUserId].
func (cldcui *DataCloudfrontLogDeliveryCanonicalUserId) LocalName() string {
	return cldcui.Name
}

// Configuration returns the configuration (args) for [DataCloudfrontLogDeliveryCanonicalUserId].
func (cldcui *DataCloudfrontLogDeliveryCanonicalUserId) Configuration() interface{} {
	return cldcui.Args
}

// Attributes returns the attributes for [DataCloudfrontLogDeliveryCanonicalUserId].
func (cldcui *DataCloudfrontLogDeliveryCanonicalUserId) Attributes() dataCloudfrontLogDeliveryCanonicalUserIdAttributes {
	return dataCloudfrontLogDeliveryCanonicalUserIdAttributes{ref: terra.ReferenceDataResource(cldcui)}
}

// DataCloudfrontLogDeliveryCanonicalUserIdArgs contains the configurations for aws_cloudfront_log_delivery_canonical_user_id.
type DataCloudfrontLogDeliveryCanonicalUserIdArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type dataCloudfrontLogDeliveryCanonicalUserIdAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_cloudfront_log_delivery_canonical_user_id.
func (cldcui dataCloudfrontLogDeliveryCanonicalUserIdAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cldcui.ref.Append("id"))
}

// Region returns a reference to field region of aws_cloudfront_log_delivery_canonical_user_id.
func (cldcui dataCloudfrontLogDeliveryCanonicalUserIdAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cldcui.ref.Append("region"))
}
