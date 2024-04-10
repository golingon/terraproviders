// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package eksnodegroup

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
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

func (r ResourcesAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ResourcesAttributes) InternalWithRef(ref terra.Reference) ResourcesAttributes {
	return ResourcesAttributes{ref: ref}
}

func (r ResourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ResourcesAttributes) RemoteAccessSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("remote_access_security_group_id"))
}

func (r ResourcesAttributes) AutoscalingGroups() terra.ListValue[AutoscalingGroupsAttributes] {
	return terra.ReferenceAsList[AutoscalingGroupsAttributes](r.ref.Append("autoscaling_groups"))
}

type AutoscalingGroupsAttributes struct {
	ref terra.Reference
}

func (ag AutoscalingGroupsAttributes) InternalRef() (terra.Reference, error) {
	return ag.ref, nil
}

func (ag AutoscalingGroupsAttributes) InternalWithRef(ref terra.Reference) AutoscalingGroupsAttributes {
	return AutoscalingGroupsAttributes{ref: ref}
}

func (ag AutoscalingGroupsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ag.ref.InternalTokens()
}

func (ag AutoscalingGroupsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("name"))
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

type RemoteAccessAttributes struct {
	ref terra.Reference
}

func (ra RemoteAccessAttributes) InternalRef() (terra.Reference, error) {
	return ra.ref, nil
}

func (ra RemoteAccessAttributes) InternalWithRef(ref terra.Reference) RemoteAccessAttributes {
	return RemoteAccessAttributes{ref: ref}
}

func (ra RemoteAccessAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ra.ref.InternalTokens()
}

func (ra RemoteAccessAttributes) Ec2SshKey() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("ec2_ssh_key"))
}

func (ra RemoteAccessAttributes) SourceSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ra.ref.Append("source_security_group_ids"))
}

type ScalingConfigAttributes struct {
	ref terra.Reference
}

func (sc ScalingConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc ScalingConfigAttributes) InternalWithRef(ref terra.Reference) ScalingConfigAttributes {
	return ScalingConfigAttributes{ref: ref}
}

func (sc ScalingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc ScalingConfigAttributes) DesiredSize() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("desired_size"))
}

func (sc ScalingConfigAttributes) MaxSize() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("max_size"))
}

func (sc ScalingConfigAttributes) MinSize() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("min_size"))
}

type TaintAttributes struct {
	ref terra.Reference
}

func (t TaintAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TaintAttributes) InternalWithRef(ref terra.Reference) TaintAttributes {
	return TaintAttributes{ref: ref}
}

func (t TaintAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TaintAttributes) Effect() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("effect"))
}

func (t TaintAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("key"))
}

func (t TaintAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("value"))
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

type UpdateConfigAttributes struct {
	ref terra.Reference
}

func (uc UpdateConfigAttributes) InternalRef() (terra.Reference, error) {
	return uc.ref, nil
}

func (uc UpdateConfigAttributes) InternalWithRef(ref terra.Reference) UpdateConfigAttributes {
	return UpdateConfigAttributes{ref: ref}
}

func (uc UpdateConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return uc.ref.InternalTokens()
}

func (uc UpdateConfigAttributes) MaxUnavailable() terra.NumberValue {
	return terra.ReferenceAsNumber(uc.ref.Append("max_unavailable"))
}

func (uc UpdateConfigAttributes) MaxUnavailablePercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(uc.ref.Append("max_unavailable_percentage"))
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
