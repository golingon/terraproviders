// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package gameliftgameservergroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AutoScalingPolicy struct {
	// EstimatedInstanceWarmup: number, optional
	EstimatedInstanceWarmup terra.NumberValue `hcl:"estimated_instance_warmup,attr"`
	// TargetTrackingConfiguration: required
	TargetTrackingConfiguration *TargetTrackingConfiguration `hcl:"target_tracking_configuration,block" validate:"required"`
}

type TargetTrackingConfiguration struct {
	// TargetValue: number, required
	TargetValue terra.NumberValue `hcl:"target_value,attr" validate:"required"`
}

type InstanceDefinition struct {
	// InstanceType: string, required
	InstanceType terra.StringValue `hcl:"instance_type,attr" validate:"required"`
	// WeightedCapacity: string, optional
	WeightedCapacity terra.StringValue `hcl:"weighted_capacity,attr"`
}

type LaunchTemplate struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type AutoScalingPolicyAttributes struct {
	ref terra.Reference
}

func (asp AutoScalingPolicyAttributes) InternalRef() (terra.Reference, error) {
	return asp.ref, nil
}

func (asp AutoScalingPolicyAttributes) InternalWithRef(ref terra.Reference) AutoScalingPolicyAttributes {
	return AutoScalingPolicyAttributes{ref: ref}
}

func (asp AutoScalingPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return asp.ref.InternalTokens()
}

func (asp AutoScalingPolicyAttributes) EstimatedInstanceWarmup() terra.NumberValue {
	return terra.ReferenceAsNumber(asp.ref.Append("estimated_instance_warmup"))
}

func (asp AutoScalingPolicyAttributes) TargetTrackingConfiguration() terra.ListValue[TargetTrackingConfigurationAttributes] {
	return terra.ReferenceAsList[TargetTrackingConfigurationAttributes](asp.ref.Append("target_tracking_configuration"))
}

type TargetTrackingConfigurationAttributes struct {
	ref terra.Reference
}

func (ttc TargetTrackingConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ttc.ref, nil
}

func (ttc TargetTrackingConfigurationAttributes) InternalWithRef(ref terra.Reference) TargetTrackingConfigurationAttributes {
	return TargetTrackingConfigurationAttributes{ref: ref}
}

func (ttc TargetTrackingConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ttc.ref.InternalTokens()
}

func (ttc TargetTrackingConfigurationAttributes) TargetValue() terra.NumberValue {
	return terra.ReferenceAsNumber(ttc.ref.Append("target_value"))
}

type InstanceDefinitionAttributes struct {
	ref terra.Reference
}

func (id InstanceDefinitionAttributes) InternalRef() (terra.Reference, error) {
	return id.ref, nil
}

func (id InstanceDefinitionAttributes) InternalWithRef(ref terra.Reference) InstanceDefinitionAttributes {
	return InstanceDefinitionAttributes{ref: ref}
}

func (id InstanceDefinitionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return id.ref.InternalTokens()
}

func (id InstanceDefinitionAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(id.ref.Append("instance_type"))
}

func (id InstanceDefinitionAttributes) WeightedCapacity() terra.StringValue {
	return terra.ReferenceAsString(id.ref.Append("weighted_capacity"))
}

type LaunchTemplateAttributes struct {
	ref terra.Reference
}

func (lt LaunchTemplateAttributes) InternalRef() (terra.Reference, error) {
	return lt.ref, nil
}

func (lt LaunchTemplateAttributes) InternalWithRef(ref terra.Reference) LaunchTemplateAttributes {
	return LaunchTemplateAttributes{ref: ref}
}

func (lt LaunchTemplateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lt.ref.InternalTokens()
}

func (lt LaunchTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("id"))
}

func (lt LaunchTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("name"))
}

func (lt LaunchTemplateAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("version"))
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

type AutoScalingPolicyState struct {
	EstimatedInstanceWarmup     float64                            `json:"estimated_instance_warmup"`
	TargetTrackingConfiguration []TargetTrackingConfigurationState `json:"target_tracking_configuration"`
}

type TargetTrackingConfigurationState struct {
	TargetValue float64 `json:"target_value"`
}

type InstanceDefinitionState struct {
	InstanceType     string `json:"instance_type"`
	WeightedCapacity string `json:"weighted_capacity"`
}

type LaunchTemplateState struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
