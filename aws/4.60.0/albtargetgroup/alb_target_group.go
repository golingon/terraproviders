// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package albtargetgroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type HealthCheck struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// HealthyThreshold: number, optional
	HealthyThreshold terra.NumberValue `hcl:"healthy_threshold,attr"`
	// Interval: number, optional
	Interval terra.NumberValue `hcl:"interval,attr"`
	// Matcher: string, optional
	Matcher terra.StringValue `hcl:"matcher,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Port: string, optional
	Port terra.StringValue `hcl:"port,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// Timeout: number, optional
	Timeout terra.NumberValue `hcl:"timeout,attr"`
	// UnhealthyThreshold: number, optional
	UnhealthyThreshold terra.NumberValue `hcl:"unhealthy_threshold,attr"`
}

type Stickiness struct {
	// CookieDuration: number, optional
	CookieDuration terra.NumberValue `hcl:"cookie_duration,attr"`
	// CookieName: string, optional
	CookieName terra.StringValue `hcl:"cookie_name,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type TargetFailover struct {
	// OnDeregistration: string, required
	OnDeregistration terra.StringValue `hcl:"on_deregistration,attr" validate:"required"`
	// OnUnhealthy: string, required
	OnUnhealthy terra.StringValue `hcl:"on_unhealthy,attr" validate:"required"`
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

func (hc HealthCheckAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(hc.ref.Append("enabled"))
}

func (hc HealthCheckAttributes) HealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("healthy_threshold"))
}

func (hc HealthCheckAttributes) Interval() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("interval"))
}

func (hc HealthCheckAttributes) Matcher() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("matcher"))
}

func (hc HealthCheckAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("path"))
}

func (hc HealthCheckAttributes) Port() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("port"))
}

func (hc HealthCheckAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("protocol"))
}

func (hc HealthCheckAttributes) Timeout() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("timeout"))
}

func (hc HealthCheckAttributes) UnhealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("unhealthy_threshold"))
}

type StickinessAttributes struct {
	ref terra.Reference
}

func (s StickinessAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s StickinessAttributes) InternalWithRef(ref terra.Reference) StickinessAttributes {
	return StickinessAttributes{ref: ref}
}

func (s StickinessAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s StickinessAttributes) CookieDuration() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("cookie_duration"))
}

func (s StickinessAttributes) CookieName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("cookie_name"))
}

func (s StickinessAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enabled"))
}

func (s StickinessAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("type"))
}

type TargetFailoverAttributes struct {
	ref terra.Reference
}

func (tf TargetFailoverAttributes) InternalRef() terra.Reference {
	return tf.ref
}

func (tf TargetFailoverAttributes) InternalWithRef(ref terra.Reference) TargetFailoverAttributes {
	return TargetFailoverAttributes{ref: ref}
}

func (tf TargetFailoverAttributes) InternalTokens() hclwrite.Tokens {
	return tf.ref.InternalTokens()
}

func (tf TargetFailoverAttributes) OnDeregistration() terra.StringValue {
	return terra.ReferenceAsString(tf.ref.Append("on_deregistration"))
}

func (tf TargetFailoverAttributes) OnUnhealthy() terra.StringValue {
	return terra.ReferenceAsString(tf.ref.Append("on_unhealthy"))
}

type HealthCheckState struct {
	Enabled            bool    `json:"enabled"`
	HealthyThreshold   float64 `json:"healthy_threshold"`
	Interval           float64 `json:"interval"`
	Matcher            string  `json:"matcher"`
	Path               string  `json:"path"`
	Port               string  `json:"port"`
	Protocol           string  `json:"protocol"`
	Timeout            float64 `json:"timeout"`
	UnhealthyThreshold float64 `json:"unhealthy_threshold"`
}

type StickinessState struct {
	CookieDuration float64 `json:"cookie_duration"`
	CookieName     string  `json:"cookie_name"`
	Enabled        bool    `json:"enabled"`
	Type           string  `json:"type"`
}

type TargetFailoverState struct {
	OnDeregistration string `json:"on_deregistration"`
	OnUnhealthy      string `json:"on_unhealthy"`
}
