// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataproc_autoscaling_policy

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type BasicAlgorithm struct {
	// CooldownPeriod: string, optional
	CooldownPeriod terra.StringValue `hcl:"cooldown_period,attr"`
	// BasicAlgorithmYarnConfig: required
	YarnConfig *BasicAlgorithmYarnConfig `hcl:"yarn_config,block" validate:"required"`
}

type BasicAlgorithmYarnConfig struct {
	// GracefulDecommissionTimeout: string, required
	GracefulDecommissionTimeout terra.StringValue `hcl:"graceful_decommission_timeout,attr" validate:"required"`
	// ScaleDownFactor: number, required
	ScaleDownFactor terra.NumberValue `hcl:"scale_down_factor,attr" validate:"required"`
	// ScaleDownMinWorkerFraction: number, optional
	ScaleDownMinWorkerFraction terra.NumberValue `hcl:"scale_down_min_worker_fraction,attr"`
	// ScaleUpFactor: number, required
	ScaleUpFactor terra.NumberValue `hcl:"scale_up_factor,attr" validate:"required"`
	// ScaleUpMinWorkerFraction: number, optional
	ScaleUpMinWorkerFraction terra.NumberValue `hcl:"scale_up_min_worker_fraction,attr"`
}

type SecondaryWorkerConfig struct {
	// MaxInstances: number, optional
	MaxInstances terra.NumberValue `hcl:"max_instances,attr"`
	// MinInstances: number, optional
	MinInstances terra.NumberValue `hcl:"min_instances,attr"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type WorkerConfig struct {
	// MaxInstances: number, required
	MaxInstances terra.NumberValue `hcl:"max_instances,attr" validate:"required"`
	// MinInstances: number, optional
	MinInstances terra.NumberValue `hcl:"min_instances,attr"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
}

type BasicAlgorithmAttributes struct {
	ref terra.Reference
}

func (ba BasicAlgorithmAttributes) InternalRef() (terra.Reference, error) {
	return ba.ref, nil
}

func (ba BasicAlgorithmAttributes) InternalWithRef(ref terra.Reference) BasicAlgorithmAttributes {
	return BasicAlgorithmAttributes{ref: ref}
}

func (ba BasicAlgorithmAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ba.ref.InternalTokens()
}

func (ba BasicAlgorithmAttributes) CooldownPeriod() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("cooldown_period"))
}

func (ba BasicAlgorithmAttributes) YarnConfig() terra.ListValue[BasicAlgorithmYarnConfigAttributes] {
	return terra.ReferenceAsList[BasicAlgorithmYarnConfigAttributes](ba.ref.Append("yarn_config"))
}

type BasicAlgorithmYarnConfigAttributes struct {
	ref terra.Reference
}

func (yc BasicAlgorithmYarnConfigAttributes) InternalRef() (terra.Reference, error) {
	return yc.ref, nil
}

func (yc BasicAlgorithmYarnConfigAttributes) InternalWithRef(ref terra.Reference) BasicAlgorithmYarnConfigAttributes {
	return BasicAlgorithmYarnConfigAttributes{ref: ref}
}

func (yc BasicAlgorithmYarnConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return yc.ref.InternalTokens()
}

func (yc BasicAlgorithmYarnConfigAttributes) GracefulDecommissionTimeout() terra.StringValue {
	return terra.ReferenceAsString(yc.ref.Append("graceful_decommission_timeout"))
}

func (yc BasicAlgorithmYarnConfigAttributes) ScaleDownFactor() terra.NumberValue {
	return terra.ReferenceAsNumber(yc.ref.Append("scale_down_factor"))
}

func (yc BasicAlgorithmYarnConfigAttributes) ScaleDownMinWorkerFraction() terra.NumberValue {
	return terra.ReferenceAsNumber(yc.ref.Append("scale_down_min_worker_fraction"))
}

func (yc BasicAlgorithmYarnConfigAttributes) ScaleUpFactor() terra.NumberValue {
	return terra.ReferenceAsNumber(yc.ref.Append("scale_up_factor"))
}

func (yc BasicAlgorithmYarnConfigAttributes) ScaleUpMinWorkerFraction() terra.NumberValue {
	return terra.ReferenceAsNumber(yc.ref.Append("scale_up_min_worker_fraction"))
}

type SecondaryWorkerConfigAttributes struct {
	ref terra.Reference
}

func (swc SecondaryWorkerConfigAttributes) InternalRef() (terra.Reference, error) {
	return swc.ref, nil
}

func (swc SecondaryWorkerConfigAttributes) InternalWithRef(ref terra.Reference) SecondaryWorkerConfigAttributes {
	return SecondaryWorkerConfigAttributes{ref: ref}
}

func (swc SecondaryWorkerConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return swc.ref.InternalTokens()
}

func (swc SecondaryWorkerConfigAttributes) MaxInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(swc.ref.Append("max_instances"))
}

func (swc SecondaryWorkerConfigAttributes) MinInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(swc.ref.Append("min_instances"))
}

func (swc SecondaryWorkerConfigAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(swc.ref.Append("weight"))
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

type WorkerConfigAttributes struct {
	ref terra.Reference
}

func (wc WorkerConfigAttributes) InternalRef() (terra.Reference, error) {
	return wc.ref, nil
}

func (wc WorkerConfigAttributes) InternalWithRef(ref terra.Reference) WorkerConfigAttributes {
	return WorkerConfigAttributes{ref: ref}
}

func (wc WorkerConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wc.ref.InternalTokens()
}

func (wc WorkerConfigAttributes) MaxInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(wc.ref.Append("max_instances"))
}

func (wc WorkerConfigAttributes) MinInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(wc.ref.Append("min_instances"))
}

func (wc WorkerConfigAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(wc.ref.Append("weight"))
}

type BasicAlgorithmState struct {
	CooldownPeriod string                          `json:"cooldown_period"`
	YarnConfig     []BasicAlgorithmYarnConfigState `json:"yarn_config"`
}

type BasicAlgorithmYarnConfigState struct {
	GracefulDecommissionTimeout string  `json:"graceful_decommission_timeout"`
	ScaleDownFactor             float64 `json:"scale_down_factor"`
	ScaleDownMinWorkerFraction  float64 `json:"scale_down_min_worker_fraction"`
	ScaleUpFactor               float64 `json:"scale_up_factor"`
	ScaleUpMinWorkerFraction    float64 `json:"scale_up_min_worker_fraction"`
}

type SecondaryWorkerConfigState struct {
	MaxInstances float64 `json:"max_instances"`
	MinInstances float64 `json:"min_instances"`
	Weight       float64 `json:"weight"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type WorkerConfigState struct {
	MaxInstances float64 `json:"max_instances"`
	MinInstances float64 `json:"min_instances"`
	Weight       float64 `json:"weight"`
}
