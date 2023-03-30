// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package eksnodegroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Resources struct {
	// AutoscalingGroups: min=0
	AutoscalingGroups []AutoscalingGroups `hcl:"autoscaling_groups,block" validate:"min=0"`
}

type AutoscalingGroups struct{}

type LaunchTemplate struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
}

type RemoteAccess struct {
	// Ec2SshKey: string, optional
	Ec2SshKey terra.StringValue `hcl:"ec2_ssh_key,attr"`
	// SourceSecurityGroupIds: set of string, optional
	SourceSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"source_security_group_ids,attr"`
}

type ScalingConfig struct {
	// DesiredSize: number, required
	DesiredSize terra.NumberValue `hcl:"desired_size,attr" validate:"required"`
	// MaxSize: number, required
	MaxSize terra.NumberValue `hcl:"max_size,attr" validate:"required"`
	// MinSize: number, required
	MinSize terra.NumberValue `hcl:"min_size,attr" validate:"required"`
}

type Taint struct {
	// Effect: string, required
	Effect terra.StringValue `hcl:"effect,attr" validate:"required"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type UpdateConfig struct {
	// MaxUnavailable: number, optional
	MaxUnavailable terra.NumberValue `hcl:"max_unavailable,attr"`
	// MaxUnavailablePercentage: number, optional
	MaxUnavailablePercentage terra.NumberValue `hcl:"max_unavailable_percentage,attr"`
}

type ResourcesAttributes struct {
	ref terra.Reference
}

func (r ResourcesAttributes) InternalRef() terra.Reference {
	return r.ref
}

func (r ResourcesAttributes) InternalWithRef(ref terra.Reference) ResourcesAttributes {
	return ResourcesAttributes{ref: ref}
}

func (r ResourcesAttributes) InternalTokens() hclwrite.Tokens {
	return r.ref.InternalTokens()
}

func (r ResourcesAttributes) RemoteAccessSecurityGroupId() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("remote_access_security_group_id"))
}

func (r ResourcesAttributes) AutoscalingGroups() terra.ListValue[AutoscalingGroupsAttributes] {
	return terra.ReferenceList[AutoscalingGroupsAttributes](r.ref.Append("autoscaling_groups"))
}

type AutoscalingGroupsAttributes struct {
	ref terra.Reference
}

func (ag AutoscalingGroupsAttributes) InternalRef() terra.Reference {
	return ag.ref
}

func (ag AutoscalingGroupsAttributes) InternalWithRef(ref terra.Reference) AutoscalingGroupsAttributes {
	return AutoscalingGroupsAttributes{ref: ref}
}

func (ag AutoscalingGroupsAttributes) InternalTokens() hclwrite.Tokens {
	return ag.ref.InternalTokens()
}

func (ag AutoscalingGroupsAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ag.ref.Append("name"))
}

type LaunchTemplateAttributes struct {
	ref terra.Reference
}

func (lt LaunchTemplateAttributes) InternalRef() terra.Reference {
	return lt.ref
}

func (lt LaunchTemplateAttributes) InternalWithRef(ref terra.Reference) LaunchTemplateAttributes {
	return LaunchTemplateAttributes{ref: ref}
}

func (lt LaunchTemplateAttributes) InternalTokens() hclwrite.Tokens {
	return lt.ref.InternalTokens()
}

func (lt LaunchTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceString(lt.ref.Append("id"))
}

func (lt LaunchTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceString(lt.ref.Append("name"))
}

func (lt LaunchTemplateAttributes) Version() terra.StringValue {
	return terra.ReferenceString(lt.ref.Append("version"))
}

type RemoteAccessAttributes struct {
	ref terra.Reference
}

func (ra RemoteAccessAttributes) InternalRef() terra.Reference {
	return ra.ref
}

func (ra RemoteAccessAttributes) InternalWithRef(ref terra.Reference) RemoteAccessAttributes {
	return RemoteAccessAttributes{ref: ref}
}

func (ra RemoteAccessAttributes) InternalTokens() hclwrite.Tokens {
	return ra.ref.InternalTokens()
}

func (ra RemoteAccessAttributes) Ec2SshKey() terra.StringValue {
	return terra.ReferenceString(ra.ref.Append("ec2_ssh_key"))
}

func (ra RemoteAccessAttributes) SourceSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](ra.ref.Append("source_security_group_ids"))
}

type ScalingConfigAttributes struct {
	ref terra.Reference
}

func (sc ScalingConfigAttributes) InternalRef() terra.Reference {
	return sc.ref
}

func (sc ScalingConfigAttributes) InternalWithRef(ref terra.Reference) ScalingConfigAttributes {
	return ScalingConfigAttributes{ref: ref}
}

func (sc ScalingConfigAttributes) InternalTokens() hclwrite.Tokens {
	return sc.ref.InternalTokens()
}

func (sc ScalingConfigAttributes) DesiredSize() terra.NumberValue {
	return terra.ReferenceNumber(sc.ref.Append("desired_size"))
}

func (sc ScalingConfigAttributes) MaxSize() terra.NumberValue {
	return terra.ReferenceNumber(sc.ref.Append("max_size"))
}

func (sc ScalingConfigAttributes) MinSize() terra.NumberValue {
	return terra.ReferenceNumber(sc.ref.Append("min_size"))
}

type TaintAttributes struct {
	ref terra.Reference
}

func (t TaintAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TaintAttributes) InternalWithRef(ref terra.Reference) TaintAttributes {
	return TaintAttributes{ref: ref}
}

func (t TaintAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TaintAttributes) Effect() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("effect"))
}

func (t TaintAttributes) Key() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("key"))
}

func (t TaintAttributes) Value() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("value"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type UpdateConfigAttributes struct {
	ref terra.Reference
}

func (uc UpdateConfigAttributes) InternalRef() terra.Reference {
	return uc.ref
}

func (uc UpdateConfigAttributes) InternalWithRef(ref terra.Reference) UpdateConfigAttributes {
	return UpdateConfigAttributes{ref: ref}
}

func (uc UpdateConfigAttributes) InternalTokens() hclwrite.Tokens {
	return uc.ref.InternalTokens()
}

func (uc UpdateConfigAttributes) MaxUnavailable() terra.NumberValue {
	return terra.ReferenceNumber(uc.ref.Append("max_unavailable"))
}

func (uc UpdateConfigAttributes) MaxUnavailablePercentage() terra.NumberValue {
	return terra.ReferenceNumber(uc.ref.Append("max_unavailable_percentage"))
}

type ResourcesState struct {
	RemoteAccessSecurityGroupId string                   `json:"remote_access_security_group_id"`
	AutoscalingGroups           []AutoscalingGroupsState `json:"autoscaling_groups"`
}

type AutoscalingGroupsState struct {
	Name string `json:"name"`
}

type LaunchTemplateState struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type RemoteAccessState struct {
	Ec2SshKey              string   `json:"ec2_ssh_key"`
	SourceSecurityGroupIds []string `json:"source_security_group_ids"`
}

type ScalingConfigState struct {
	DesiredSize float64 `json:"desired_size"`
	MaxSize     float64 `json:"max_size"`
	MinSize     float64 `json:"min_size"`
}

type TaintState struct {
	Effect string `json:"effect"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type UpdateConfigState struct {
	MaxUnavailable           float64 `json:"max_unavailable"`
	MaxUnavailablePercentage float64 `json:"max_unavailable_percentage"`
}
