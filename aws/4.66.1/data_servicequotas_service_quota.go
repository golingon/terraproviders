// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataservicequotasservicequota "github.com/golingon/terraproviders/aws/4.66.1/dataservicequotasservicequota"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServicequotasServiceQuota creates a new instance of [DataServicequotasServiceQuota].
func NewDataServicequotasServiceQuota(name string, args DataServicequotasServiceQuotaArgs) *DataServicequotasServiceQuota {
	return &DataServicequotasServiceQuota{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicequotasServiceQuota)(nil)

// DataServicequotasServiceQuota represents the Terraform data resource aws_servicequotas_service_quota.
type DataServicequotasServiceQuota struct {
	Name string
	Args DataServicequotasServiceQuotaArgs
}

// DataSource returns the Terraform object type for [DataServicequotasServiceQuota].
func (ssq *DataServicequotasServiceQuota) DataSource() string {
	return "aws_servicequotas_service_quota"
}

// LocalName returns the local name for [DataServicequotasServiceQuota].
func (ssq *DataServicequotasServiceQuota) LocalName() string {
	return ssq.Name
}

// Configuration returns the configuration (args) for [DataServicequotasServiceQuota].
func (ssq *DataServicequotasServiceQuota) Configuration() interface{} {
	return ssq.Args
}

// Attributes returns the attributes for [DataServicequotasServiceQuota].
func (ssq *DataServicequotasServiceQuota) Attributes() dataServicequotasServiceQuotaAttributes {
	return dataServicequotasServiceQuotaAttributes{ref: terra.ReferenceDataResource(ssq)}
}

// DataServicequotasServiceQuotaArgs contains the configurations for aws_servicequotas_service_quota.
type DataServicequotasServiceQuotaArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// QuotaCode: string, optional
	QuotaCode terra.StringValue `hcl:"quota_code,attr"`
	// QuotaName: string, optional
	QuotaName terra.StringValue `hcl:"quota_name,attr"`
	// ServiceCode: string, required
	ServiceCode terra.StringValue `hcl:"service_code,attr" validate:"required"`
	// UsageMetric: min=0
	UsageMetric []dataservicequotasservicequota.UsageMetric `hcl:"usage_metric,block" validate:"min=0"`
}
type dataServicequotasServiceQuotaAttributes struct {
	ref terra.Reference
}

// Adjustable returns a reference to field adjustable of aws_servicequotas_service_quota.
func (ssq dataServicequotasServiceQuotaAttributes) Adjustable() terra.BoolValue {
	return terra.ReferenceAsBool(ssq.ref.Append("adjustable"))
}

// Arn returns a reference to field arn of aws_servicequotas_service_quota.
func (ssq dataServicequotasServiceQuotaAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("arn"))
}

// DefaultValue returns a reference to field default_value of aws_servicequotas_service_quota.
func (ssq dataServicequotasServiceQuotaAttributes) DefaultValue() terra.NumberValue {
	return terra.ReferenceAsNumber(ssq.ref.Append("default_value"))
}

// GlobalQuota returns a reference to field global_quota of aws_servicequotas_service_quota.
func (ssq dataServicequotasServiceQuotaAttributes) GlobalQuota() terra.BoolValue {
	return terra.ReferenceAsBool(ssq.ref.Append("global_quota"))
}

// Id returns a reference to field id of aws_servicequotas_service_quota.
func (ssq dataServicequotasServiceQuotaAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("id"))
}

// QuotaCode returns a reference to field quota_code of aws_servicequotas_service_quota.
func (ssq dataServicequotasServiceQuotaAttributes) QuotaCode() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("quota_code"))
}

// QuotaName returns a reference to field quota_name of aws_servicequotas_service_quota.
func (ssq dataServicequotasServiceQuotaAttributes) QuotaName() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("quota_name"))
}

// ServiceCode returns a reference to field service_code of aws_servicequotas_service_quota.
func (ssq dataServicequotasServiceQuotaAttributes) ServiceCode() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("service_code"))
}

// ServiceName returns a reference to field service_name of aws_servicequotas_service_quota.
func (ssq dataServicequotasServiceQuotaAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("service_name"))
}

// Value returns a reference to field value of aws_servicequotas_service_quota.
func (ssq dataServicequotasServiceQuotaAttributes) Value() terra.NumberValue {
	return terra.ReferenceAsNumber(ssq.ref.Append("value"))
}

func (ssq dataServicequotasServiceQuotaAttributes) UsageMetric() terra.ListValue[dataservicequotasservicequota.UsageMetricAttributes] {
	return terra.ReferenceAsList[dataservicequotasservicequota.UsageMetricAttributes](ssq.ref.Append("usage_metric"))
}
