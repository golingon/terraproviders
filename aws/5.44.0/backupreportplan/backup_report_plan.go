// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package backupreportplan

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ReportDeliveryChannel struct {
	// Formats: set of string, optional
	Formats terra.SetValue[terra.StringValue] `hcl:"formats,attr"`
	// S3BucketName: string, required
	S3BucketName terra.StringValue `hcl:"s3_bucket_name,attr" validate:"required"`
	// S3KeyPrefix: string, optional
	S3KeyPrefix terra.StringValue `hcl:"s3_key_prefix,attr"`
}

type ReportSetting struct {
	// Accounts: set of string, optional
	Accounts terra.SetValue[terra.StringValue] `hcl:"accounts,attr"`
	// FrameworkArns: set of string, optional
	FrameworkArns terra.SetValue[terra.StringValue] `hcl:"framework_arns,attr"`
	// NumberOfFrameworks: number, optional
	NumberOfFrameworks terra.NumberValue `hcl:"number_of_frameworks,attr"`
	// OrganizationUnits: set of string, optional
	OrganizationUnits terra.SetValue[terra.StringValue] `hcl:"organization_units,attr"`
	// Regions: set of string, optional
	Regions terra.SetValue[terra.StringValue] `hcl:"regions,attr"`
	// ReportTemplate: string, required
	ReportTemplate terra.StringValue `hcl:"report_template,attr" validate:"required"`
}

type ReportDeliveryChannelAttributes struct {
	ref terra.Reference
}

func (rdc ReportDeliveryChannelAttributes) InternalRef() (terra.Reference, error) {
	return rdc.ref, nil
}

func (rdc ReportDeliveryChannelAttributes) InternalWithRef(ref terra.Reference) ReportDeliveryChannelAttributes {
	return ReportDeliveryChannelAttributes{ref: ref}
}

func (rdc ReportDeliveryChannelAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rdc.ref.InternalTokens()
}

func (rdc ReportDeliveryChannelAttributes) Formats() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rdc.ref.Append("formats"))
}

func (rdc ReportDeliveryChannelAttributes) S3BucketName() terra.StringValue {
	return terra.ReferenceAsString(rdc.ref.Append("s3_bucket_name"))
}

func (rdc ReportDeliveryChannelAttributes) S3KeyPrefix() terra.StringValue {
	return terra.ReferenceAsString(rdc.ref.Append("s3_key_prefix"))
}

type ReportSettingAttributes struct {
	ref terra.Reference
}

func (rs ReportSettingAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs ReportSettingAttributes) InternalWithRef(ref terra.Reference) ReportSettingAttributes {
	return ReportSettingAttributes{ref: ref}
}

func (rs ReportSettingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs ReportSettingAttributes) Accounts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rs.ref.Append("accounts"))
}

func (rs ReportSettingAttributes) FrameworkArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rs.ref.Append("framework_arns"))
}

func (rs ReportSettingAttributes) NumberOfFrameworks() terra.NumberValue {
	return terra.ReferenceAsNumber(rs.ref.Append("number_of_frameworks"))
}

func (rs ReportSettingAttributes) OrganizationUnits() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rs.ref.Append("organization_units"))
}

func (rs ReportSettingAttributes) Regions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rs.ref.Append("regions"))
}

func (rs ReportSettingAttributes) ReportTemplate() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("report_template"))
}

type ReportDeliveryChannelState struct {
	Formats      []string `json:"formats"`
	S3BucketName string   `json:"s3_bucket_name"`
	S3KeyPrefix  string   `json:"s3_key_prefix"`
}

type ReportSettingState struct {
	Accounts           []string `json:"accounts"`
	FrameworkArns      []string `json:"framework_arns"`
	NumberOfFrameworks float64  `json:"number_of_frameworks"`
	OrganizationUnits  []string `json:"organization_units"`
	Regions            []string `json:"regions"`
	ReportTemplate     string   `json:"report_template"`
}
