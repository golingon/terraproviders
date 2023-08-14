// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudtasksqueue

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AppEngineRoutingOverride struct {
	// Instance: string, optional
	Instance terra.StringValue `hcl:"instance,attr"`
	// Service: string, optional
	Service terra.StringValue `hcl:"service,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
}

type RateLimits struct {
	// MaxConcurrentDispatches: number, optional
	MaxConcurrentDispatches terra.NumberValue `hcl:"max_concurrent_dispatches,attr"`
	// MaxDispatchesPerSecond: number, optional
	MaxDispatchesPerSecond terra.NumberValue `hcl:"max_dispatches_per_second,attr"`
}

type RetryConfig struct {
	// MaxAttempts: number, optional
	MaxAttempts terra.NumberValue `hcl:"max_attempts,attr"`
	// MaxBackoff: string, optional
	MaxBackoff terra.StringValue `hcl:"max_backoff,attr"`
	// MaxDoublings: number, optional
	MaxDoublings terra.NumberValue `hcl:"max_doublings,attr"`
	// MaxRetryDuration: string, optional
	MaxRetryDuration terra.StringValue `hcl:"max_retry_duration,attr"`
	// MinBackoff: string, optional
	MinBackoff terra.StringValue `hcl:"min_backoff,attr"`
}

type StackdriverLoggingConfig struct {
	// SamplingRatio: number, required
	SamplingRatio terra.NumberValue `hcl:"sampling_ratio,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AppEngineRoutingOverrideAttributes struct {
	ref terra.Reference
}

func (aero AppEngineRoutingOverrideAttributes) InternalRef() (terra.Reference, error) {
	return aero.ref, nil
}

func (aero AppEngineRoutingOverrideAttributes) InternalWithRef(ref terra.Reference) AppEngineRoutingOverrideAttributes {
	return AppEngineRoutingOverrideAttributes{ref: ref}
}

func (aero AppEngineRoutingOverrideAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aero.ref.InternalTokens()
}

func (aero AppEngineRoutingOverrideAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(aero.ref.Append("host"))
}

func (aero AppEngineRoutingOverrideAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(aero.ref.Append("instance"))
}

func (aero AppEngineRoutingOverrideAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(aero.ref.Append("service"))
}

func (aero AppEngineRoutingOverrideAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(aero.ref.Append("version"))
}

type RateLimitsAttributes struct {
	ref terra.Reference
}

func (rl RateLimitsAttributes) InternalRef() (terra.Reference, error) {
	return rl.ref, nil
}

func (rl RateLimitsAttributes) InternalWithRef(ref terra.Reference) RateLimitsAttributes {
	return RateLimitsAttributes{ref: ref}
}

func (rl RateLimitsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rl.ref.InternalTokens()
}

func (rl RateLimitsAttributes) MaxBurstSize() terra.NumberValue {
	return terra.ReferenceAsNumber(rl.ref.Append("max_burst_size"))
}

func (rl RateLimitsAttributes) MaxConcurrentDispatches() terra.NumberValue {
	return terra.ReferenceAsNumber(rl.ref.Append("max_concurrent_dispatches"))
}

func (rl RateLimitsAttributes) MaxDispatchesPerSecond() terra.NumberValue {
	return terra.ReferenceAsNumber(rl.ref.Append("max_dispatches_per_second"))
}

type RetryConfigAttributes struct {
	ref terra.Reference
}

func (rc RetryConfigAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc RetryConfigAttributes) InternalWithRef(ref terra.Reference) RetryConfigAttributes {
	return RetryConfigAttributes{ref: ref}
}

func (rc RetryConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc RetryConfigAttributes) MaxAttempts() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("max_attempts"))
}

func (rc RetryConfigAttributes) MaxBackoff() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("max_backoff"))
}

func (rc RetryConfigAttributes) MaxDoublings() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("max_doublings"))
}

func (rc RetryConfigAttributes) MaxRetryDuration() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("max_retry_duration"))
}

func (rc RetryConfigAttributes) MinBackoff() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("min_backoff"))
}

type StackdriverLoggingConfigAttributes struct {
	ref terra.Reference
}

func (slc StackdriverLoggingConfigAttributes) InternalRef() (terra.Reference, error) {
	return slc.ref, nil
}

func (slc StackdriverLoggingConfigAttributes) InternalWithRef(ref terra.Reference) StackdriverLoggingConfigAttributes {
	return StackdriverLoggingConfigAttributes{ref: ref}
}

func (slc StackdriverLoggingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return slc.ref.InternalTokens()
}

func (slc StackdriverLoggingConfigAttributes) SamplingRatio() terra.NumberValue {
	return terra.ReferenceAsNumber(slc.ref.Append("sampling_ratio"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AppEngineRoutingOverrideState struct {
	Host     string `json:"host"`
	Instance string `json:"instance"`
	Service  string `json:"service"`
	Version  string `json:"version"`
}

type RateLimitsState struct {
	MaxBurstSize            float64 `json:"max_burst_size"`
	MaxConcurrentDispatches float64 `json:"max_concurrent_dispatches"`
	MaxDispatchesPerSecond  float64 `json:"max_dispatches_per_second"`
}

type RetryConfigState struct {
	MaxAttempts      float64 `json:"max_attempts"`
	MaxBackoff       string  `json:"max_backoff"`
	MaxDoublings     float64 `json:"max_doublings"`
	MaxRetryDuration string  `json:"max_retry_duration"`
	MinBackoff       string  `json:"min_backoff"`
}

type StackdriverLoggingConfigState struct {
	SamplingRatio float64 `json:"sampling_ratio"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
