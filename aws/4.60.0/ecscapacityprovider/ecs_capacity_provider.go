// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ecscapacityprovider

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AutoScalingGroupProvider struct {
	// AutoScalingGroupArn: string, required
	AutoScalingGroupArn terra.StringValue `hcl:"auto_scaling_group_arn,attr" validate:"required"`
	// ManagedTerminationProtection: string, optional
	ManagedTerminationProtection terra.StringValue `hcl:"managed_termination_protection,attr"`
	// ManagedScaling: optional
	ManagedScaling *ManagedScaling `hcl:"managed_scaling,block"`
}

type ManagedScaling struct {
	// InstanceWarmupPeriod: number, optional
	InstanceWarmupPeriod terra.NumberValue `hcl:"instance_warmup_period,attr"`
	// MaximumScalingStepSize: number, optional
	MaximumScalingStepSize terra.NumberValue `hcl:"maximum_scaling_step_size,attr"`
	// MinimumScalingStepSize: number, optional
	MinimumScalingStepSize terra.NumberValue `hcl:"minimum_scaling_step_size,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// TargetCapacity: number, optional
	TargetCapacity terra.NumberValue `hcl:"target_capacity,attr"`
}

type AutoScalingGroupProviderAttributes struct {
	ref terra.Reference
}

func (asgp AutoScalingGroupProviderAttributes) InternalRef() terra.Reference {
	return asgp.ref
}

func (asgp AutoScalingGroupProviderAttributes) InternalWithRef(ref terra.Reference) AutoScalingGroupProviderAttributes {
	return AutoScalingGroupProviderAttributes{ref: ref}
}

func (asgp AutoScalingGroupProviderAttributes) InternalTokens() hclwrite.Tokens {
	return asgp.ref.InternalTokens()
}

func (asgp AutoScalingGroupProviderAttributes) AutoScalingGroupArn() terra.StringValue {
	return terra.ReferenceAsString(asgp.ref.Append("auto_scaling_group_arn"))
}

func (asgp AutoScalingGroupProviderAttributes) ManagedTerminationProtection() terra.StringValue {
	return terra.ReferenceAsString(asgp.ref.Append("managed_termination_protection"))
}

func (asgp AutoScalingGroupProviderAttributes) ManagedScaling() terra.ListValue[ManagedScalingAttributes] {
	return terra.ReferenceAsList[ManagedScalingAttributes](asgp.ref.Append("managed_scaling"))
}

type ManagedScalingAttributes struct {
	ref terra.Reference
}

func (ms ManagedScalingAttributes) InternalRef() terra.Reference {
	return ms.ref
}

func (ms ManagedScalingAttributes) InternalWithRef(ref terra.Reference) ManagedScalingAttributes {
	return ManagedScalingAttributes{ref: ref}
}

func (ms ManagedScalingAttributes) InternalTokens() hclwrite.Tokens {
	return ms.ref.InternalTokens()
}

func (ms ManagedScalingAttributes) InstanceWarmupPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("instance_warmup_period"))
}

func (ms ManagedScalingAttributes) MaximumScalingStepSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("maximum_scaling_step_size"))
}

func (ms ManagedScalingAttributes) MinimumScalingStepSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("minimum_scaling_step_size"))
}

func (ms ManagedScalingAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("status"))
}

func (ms ManagedScalingAttributes) TargetCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("target_capacity"))
}

type AutoScalingGroupProviderState struct {
	AutoScalingGroupArn          string                `json:"auto_scaling_group_arn"`
	ManagedTerminationProtection string                `json:"managed_termination_protection"`
	ManagedScaling               []ManagedScalingState `json:"managed_scaling"`
}

type ManagedScalingState struct {
	InstanceWarmupPeriod   float64 `json:"instance_warmup_period"`
	MaximumScalingStepSize float64 `json:"maximum_scaling_step_size"`
	MinimumScalingStepSize float64 `json:"minimum_scaling_step_size"`
	Status                 string  `json:"status"`
	TargetCapacity         float64 `json:"target_capacity"`
}