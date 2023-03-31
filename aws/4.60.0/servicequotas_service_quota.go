// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicequotasServiceQuota creates a new instance of [ServicequotasServiceQuota].
func NewServicequotasServiceQuota(name string, args ServicequotasServiceQuotaArgs) *ServicequotasServiceQuota {
	return &ServicequotasServiceQuota{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicequotasServiceQuota)(nil)

// ServicequotasServiceQuota represents the Terraform resource aws_servicequotas_service_quota.
type ServicequotasServiceQuota struct {
	Name      string
	Args      ServicequotasServiceQuotaArgs
	state     *servicequotasServiceQuotaState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicequotasServiceQuota].
func (ssq *ServicequotasServiceQuota) Type() string {
	return "aws_servicequotas_service_quota"
}

// LocalName returns the local name for [ServicequotasServiceQuota].
func (ssq *ServicequotasServiceQuota) LocalName() string {
	return ssq.Name
}

// Configuration returns the configuration (args) for [ServicequotasServiceQuota].
func (ssq *ServicequotasServiceQuota) Configuration() interface{} {
	return ssq.Args
}

// DependOn is used for other resources to depend on [ServicequotasServiceQuota].
func (ssq *ServicequotasServiceQuota) DependOn() terra.Reference {
	return terra.ReferenceResource(ssq)
}

// Dependencies returns the list of resources [ServicequotasServiceQuota] depends_on.
func (ssq *ServicequotasServiceQuota) Dependencies() terra.Dependencies {
	return ssq.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicequotasServiceQuota].
func (ssq *ServicequotasServiceQuota) LifecycleManagement() *terra.Lifecycle {
	return ssq.Lifecycle
}

// Attributes returns the attributes for [ServicequotasServiceQuota].
func (ssq *ServicequotasServiceQuota) Attributes() servicequotasServiceQuotaAttributes {
	return servicequotasServiceQuotaAttributes{ref: terra.ReferenceResource(ssq)}
}

// ImportState imports the given attribute values into [ServicequotasServiceQuota]'s state.
func (ssq *ServicequotasServiceQuota) ImportState(av io.Reader) error {
	ssq.state = &servicequotasServiceQuotaState{}
	if err := json.NewDecoder(av).Decode(ssq.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssq.Type(), ssq.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicequotasServiceQuota] has state.
func (ssq *ServicequotasServiceQuota) State() (*servicequotasServiceQuotaState, bool) {
	return ssq.state, ssq.state != nil
}

// StateMust returns the state for [ServicequotasServiceQuota]. Panics if the state is nil.
func (ssq *ServicequotasServiceQuota) StateMust() *servicequotasServiceQuotaState {
	if ssq.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssq.Type(), ssq.LocalName()))
	}
	return ssq.state
}

// ServicequotasServiceQuotaArgs contains the configurations for aws_servicequotas_service_quota.
type ServicequotasServiceQuotaArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// QuotaCode: string, required
	QuotaCode terra.StringValue `hcl:"quota_code,attr" validate:"required"`
	// ServiceCode: string, required
	ServiceCode terra.StringValue `hcl:"service_code,attr" validate:"required"`
	// Value: number, required
	Value terra.NumberValue `hcl:"value,attr" validate:"required"`
}
type servicequotasServiceQuotaAttributes struct {
	ref terra.Reference
}

// Adjustable returns a reference to field adjustable of aws_servicequotas_service_quota.
func (ssq servicequotasServiceQuotaAttributes) Adjustable() terra.BoolValue {
	return terra.ReferenceAsBool(ssq.ref.Append("adjustable"))
}

// Arn returns a reference to field arn of aws_servicequotas_service_quota.
func (ssq servicequotasServiceQuotaAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("arn"))
}

// DefaultValue returns a reference to field default_value of aws_servicequotas_service_quota.
func (ssq servicequotasServiceQuotaAttributes) DefaultValue() terra.NumberValue {
	return terra.ReferenceAsNumber(ssq.ref.Append("default_value"))
}

// Id returns a reference to field id of aws_servicequotas_service_quota.
func (ssq servicequotasServiceQuotaAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("id"))
}

// QuotaCode returns a reference to field quota_code of aws_servicequotas_service_quota.
func (ssq servicequotasServiceQuotaAttributes) QuotaCode() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("quota_code"))
}

// QuotaName returns a reference to field quota_name of aws_servicequotas_service_quota.
func (ssq servicequotasServiceQuotaAttributes) QuotaName() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("quota_name"))
}

// RequestId returns a reference to field request_id of aws_servicequotas_service_quota.
func (ssq servicequotasServiceQuotaAttributes) RequestId() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("request_id"))
}

// RequestStatus returns a reference to field request_status of aws_servicequotas_service_quota.
func (ssq servicequotasServiceQuotaAttributes) RequestStatus() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("request_status"))
}

// ServiceCode returns a reference to field service_code of aws_servicequotas_service_quota.
func (ssq servicequotasServiceQuotaAttributes) ServiceCode() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("service_code"))
}

// ServiceName returns a reference to field service_name of aws_servicequotas_service_quota.
func (ssq servicequotasServiceQuotaAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(ssq.ref.Append("service_name"))
}

// Value returns a reference to field value of aws_servicequotas_service_quota.
func (ssq servicequotasServiceQuotaAttributes) Value() terra.NumberValue {
	return terra.ReferenceAsNumber(ssq.ref.Append("value"))
}

type servicequotasServiceQuotaState struct {
	Adjustable    bool    `json:"adjustable"`
	Arn           string  `json:"arn"`
	DefaultValue  float64 `json:"default_value"`
	Id            string  `json:"id"`
	QuotaCode     string  `json:"quota_code"`
	QuotaName     string  `json:"quota_name"`
	RequestId     string  `json:"request_id"`
	RequestStatus string  `json:"request_status"`
	ServiceCode   string  `json:"service_code"`
	ServiceName   string  `json:"service_name"`
	Value         float64 `json:"value"`
}
