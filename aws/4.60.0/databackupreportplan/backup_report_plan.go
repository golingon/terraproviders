// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package databackupreportplan

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ReportDeliveryChannel struct{}

type ReportSetting struct{}

type ReportDeliveryChannelAttributes struct {
	ref terra.Reference
}

func (rdc ReportDeliveryChannelAttributes) InternalRef() terra.Reference {
	return rdc.ref
}

func (rdc ReportDeliveryChannelAttributes) InternalWithRef(ref terra.Reference) ReportDeliveryChannelAttributes {
	return ReportDeliveryChannelAttributes{ref: ref}
}

func (rdc ReportDeliveryChannelAttributes) InternalTokens() hclwrite.Tokens {
	return rdc.ref.InternalTokens()
}

func (rdc ReportDeliveryChannelAttributes) Formats() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](rdc.ref.Append("formats"))
}

func (rdc ReportDeliveryChannelAttributes) S3BucketName() terra.StringValue {
	return terra.ReferenceString(rdc.ref.Append("s3_bucket_name"))
}

func (rdc ReportDeliveryChannelAttributes) S3KeyPrefix() terra.StringValue {
	return terra.ReferenceString(rdc.ref.Append("s3_key_prefix"))
}

type ReportSettingAttributes struct {
	ref terra.Reference
}

func (rs ReportSettingAttributes) InternalRef() terra.Reference {
	return rs.ref
}

func (rs ReportSettingAttributes) InternalWithRef(ref terra.Reference) ReportSettingAttributes {
	return ReportSettingAttributes{ref: ref}
}

func (rs ReportSettingAttributes) InternalTokens() hclwrite.Tokens {
	return rs.ref.InternalTokens()
}

func (rs ReportSettingAttributes) FrameworkArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](rs.ref.Append("framework_arns"))
}

func (rs ReportSettingAttributes) NumberOfFrameworks() terra.NumberValue {
	return terra.ReferenceNumber(rs.ref.Append("number_of_frameworks"))
}

func (rs ReportSettingAttributes) ReportTemplate() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("report_template"))
}

type ReportDeliveryChannelState struct {
	Formats      []string `json:"formats"`
	S3BucketName string   `json:"s3_bucket_name"`
	S3KeyPrefix  string   `json:"s3_key_prefix"`
}

type ReportSettingState struct {
	FrameworkArns      []string `json:"framework_arns"`
	NumberOfFrameworks float64  `json:"number_of_frameworks"`
	ReportTemplate     string   `json:"report_template"`
}
