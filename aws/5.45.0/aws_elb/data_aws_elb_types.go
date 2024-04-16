// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_elb

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataAccessLogsAttributes struct {
	ref terra.Reference
}

func (al DataAccessLogsAttributes) InternalRef() (terra.Reference, error) {
	return al.ref, nil
}

func (al DataAccessLogsAttributes) InternalWithRef(ref terra.Reference) DataAccessLogsAttributes {
	return DataAccessLogsAttributes{ref: ref}
}

func (al DataAccessLogsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return al.ref.InternalTokens()
}

func (al DataAccessLogsAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("bucket"))
}

func (al DataAccessLogsAttributes) BucketPrefix() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("bucket_prefix"))
}

func (al DataAccessLogsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(al.ref.Append("enabled"))
}

func (al DataAccessLogsAttributes) Interval() terra.NumberValue {
	return terra.ReferenceAsNumber(al.ref.Append("interval"))
}

type DataHealthCheckAttributes struct {
	ref terra.Reference
}

func (hc DataHealthCheckAttributes) InternalRef() (terra.Reference, error) {
	return hc.ref, nil
}

func (hc DataHealthCheckAttributes) InternalWithRef(ref terra.Reference) DataHealthCheckAttributes {
	return DataHealthCheckAttributes{ref: ref}
}

func (hc DataHealthCheckAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hc.ref.InternalTokens()
}

func (hc DataHealthCheckAttributes) HealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("healthy_threshold"))
}

func (hc DataHealthCheckAttributes) Interval() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("interval"))
}

func (hc DataHealthCheckAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("target"))
}

func (hc DataHealthCheckAttributes) Timeout() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("timeout"))
}

func (hc DataHealthCheckAttributes) UnhealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("unhealthy_threshold"))
}

type DataListenerAttributes struct {
	ref terra.Reference
}

func (l DataListenerAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l DataListenerAttributes) InternalWithRef(ref terra.Reference) DataListenerAttributes {
	return DataListenerAttributes{ref: ref}
}

func (l DataListenerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l DataListenerAttributes) InstancePort() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("instance_port"))
}

func (l DataListenerAttributes) InstanceProtocol() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("instance_protocol"))
}

func (l DataListenerAttributes) LbPort() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("lb_port"))
}

func (l DataListenerAttributes) LbProtocol() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("lb_protocol"))
}

func (l DataListenerAttributes) SslCertificateId() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("ssl_certificate_id"))
}

type DataAccessLogsState struct {
	Bucket       string  `json:"bucket"`
	BucketPrefix string  `json:"bucket_prefix"`
	Enabled      bool    `json:"enabled"`
	Interval     float64 `json:"interval"`
}

type DataHealthCheckState struct {
	HealthyThreshold   float64 `json:"healthy_threshold"`
	Interval           float64 `json:"interval"`
	Target             string  `json:"target"`
	Timeout            float64 `json:"timeout"`
	UnhealthyThreshold float64 `json:"unhealthy_threshold"`
}

type DataListenerState struct {
	InstancePort     float64 `json:"instance_port"`
	InstanceProtocol string  `json:"instance_protocol"`
	LbPort           float64 `json:"lb_port"`
	LbProtocol       string  `json:"lb_protocol"`
	SslCertificateId string  `json:"ssl_certificate_id"`
}
