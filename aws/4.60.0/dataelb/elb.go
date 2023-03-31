// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataelb

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AccessLogs struct{}

type HealthCheck struct{}

type Listener struct{}

type AccessLogsAttributes struct {
	ref terra.Reference
}

func (al AccessLogsAttributes) InternalRef() terra.Reference {
	return al.ref
}

func (al AccessLogsAttributes) InternalWithRef(ref terra.Reference) AccessLogsAttributes {
	return AccessLogsAttributes{ref: ref}
}

func (al AccessLogsAttributes) InternalTokens() hclwrite.Tokens {
	return al.ref.InternalTokens()
}

func (al AccessLogsAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("bucket"))
}

func (al AccessLogsAttributes) BucketPrefix() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("bucket_prefix"))
}

func (al AccessLogsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(al.ref.Append("enabled"))
}

func (al AccessLogsAttributes) Interval() terra.NumberValue {
	return terra.ReferenceAsNumber(al.ref.Append("interval"))
}

type HealthCheckAttributes struct {
	ref terra.Reference
}

func (hc HealthCheckAttributes) InternalRef() terra.Reference {
	return hc.ref
}

func (hc HealthCheckAttributes) InternalWithRef(ref terra.Reference) HealthCheckAttributes {
	return HealthCheckAttributes{ref: ref}
}

func (hc HealthCheckAttributes) InternalTokens() hclwrite.Tokens {
	return hc.ref.InternalTokens()
}

func (hc HealthCheckAttributes) HealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("healthy_threshold"))
}

func (hc HealthCheckAttributes) Interval() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("interval"))
}

func (hc HealthCheckAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("target"))
}

func (hc HealthCheckAttributes) Timeout() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("timeout"))
}

func (hc HealthCheckAttributes) UnhealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("unhealthy_threshold"))
}

type ListenerAttributes struct {
	ref terra.Reference
}

func (l ListenerAttributes) InternalRef() terra.Reference {
	return l.ref
}

func (l ListenerAttributes) InternalWithRef(ref terra.Reference) ListenerAttributes {
	return ListenerAttributes{ref: ref}
}

func (l ListenerAttributes) InternalTokens() hclwrite.Tokens {
	return l.ref.InternalTokens()
}

func (l ListenerAttributes) InstancePort() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("instance_port"))
}

func (l ListenerAttributes) InstanceProtocol() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("instance_protocol"))
}

func (l ListenerAttributes) LbPort() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("lb_port"))
}

func (l ListenerAttributes) LbProtocol() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("lb_protocol"))
}

func (l ListenerAttributes) SslCertificateId() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("ssl_certificate_id"))
}

type AccessLogsState struct {
	Bucket       string  `json:"bucket"`
	BucketPrefix string  `json:"bucket_prefix"`
	Enabled      bool    `json:"enabled"`
	Interval     float64 `json:"interval"`
}

type HealthCheckState struct {
	HealthyThreshold   float64 `json:"healthy_threshold"`
	Interval           float64 `json:"interval"`
	Target             string  `json:"target"`
	Timeout            float64 `json:"timeout"`
	UnhealthyThreshold float64 `json:"unhealthy_threshold"`
}

type ListenerState struct {
	InstancePort     float64 `json:"instance_port"`
	InstanceProtocol string  `json:"instance_protocol"`
	LbPort           float64 `json:"lb_port"`
	LbProtocol       string  `json:"lb_protocol"`
	SslCertificateId string  `json:"ssl_certificate_id"`
}
