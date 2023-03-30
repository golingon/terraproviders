// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataS3BucketObjects(name string, args DataS3BucketObjectsArgs) *DataS3BucketObjects {
	return &DataS3BucketObjects{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataS3BucketObjects)(nil)

type DataS3BucketObjects struct {
	Name string
	Args DataS3BucketObjectsArgs
}

func (sbo *DataS3BucketObjects) DataSource() string {
	return "aws_s3_bucket_objects"
}

func (sbo *DataS3BucketObjects) LocalName() string {
	return sbo.Name
}

func (sbo *DataS3BucketObjects) Configuration() interface{} {
	return sbo.Args
}

func (sbo *DataS3BucketObjects) Attributes() dataS3BucketObjectsAttributes {
	return dataS3BucketObjectsAttributes{ref: terra.ReferenceDataResource(sbo)}
}

type DataS3BucketObjectsArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Delimiter: string, optional
	Delimiter terra.StringValue `hcl:"delimiter,attr"`
	// EncodingType: string, optional
	EncodingType terra.StringValue `hcl:"encoding_type,attr"`
	// FetchOwner: bool, optional
	FetchOwner terra.BoolValue `hcl:"fetch_owner,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxKeys: number, optional
	MaxKeys terra.NumberValue `hcl:"max_keys,attr"`
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
	// StartAfter: string, optional
	StartAfter terra.StringValue `hcl:"start_after,attr"`
}
type dataS3BucketObjectsAttributes struct {
	ref terra.Reference
}

func (sbo dataS3BucketObjectsAttributes) Bucket() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("bucket"))
}

func (sbo dataS3BucketObjectsAttributes) CommonPrefixes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](sbo.ref.Append("common_prefixes"))
}

func (sbo dataS3BucketObjectsAttributes) Delimiter() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("delimiter"))
}

func (sbo dataS3BucketObjectsAttributes) EncodingType() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("encoding_type"))
}

func (sbo dataS3BucketObjectsAttributes) FetchOwner() terra.BoolValue {
	return terra.ReferenceBool(sbo.ref.Append("fetch_owner"))
}

func (sbo dataS3BucketObjectsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("id"))
}

func (sbo dataS3BucketObjectsAttributes) Keys() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](sbo.ref.Append("keys"))
}

func (sbo dataS3BucketObjectsAttributes) MaxKeys() terra.NumberValue {
	return terra.ReferenceNumber(sbo.ref.Append("max_keys"))
}

func (sbo dataS3BucketObjectsAttributes) Owners() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](sbo.ref.Append("owners"))
}

func (sbo dataS3BucketObjectsAttributes) Prefix() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("prefix"))
}

func (sbo dataS3BucketObjectsAttributes) StartAfter() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("start_after"))
}
