// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataeksnodegroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type LaunchTemplate struct{}

type RemoteAccess struct{}

type Resources struct {
	// AutoscalingGroups: min=0
	AutoscalingGroups []AutoscalingGroups `hcl:"autoscaling_groups,block" validate:"min=0"`
}

type AutoscalingGroups struct{}

type ScalingConfig struct{}

type Taints struct{}

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

type TaintsAttributes struct {
	ref terra.Reference
}

func (t TaintsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TaintsAttributes) InternalWithRef(ref terra.Reference) TaintsAttributes {
	return TaintsAttributes{ref: ref}
}

func (t TaintsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TaintsAttributes) Effect() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("effect"))
}

func (t TaintsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("key"))
}

func (t TaintsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("value"))
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

type ResourcesState struct {
	RemoteAccessSecurityGroupId string                   `json:"remote_access_security_group_id"`
	AutoscalingGroups           []AutoscalingGroupsState `json:"autoscaling_groups"`
}

type AutoscalingGroupsState struct {
	Name string `json:"name"`
}

type ScalingConfigState struct {
	DesiredSize float64 `json:"desired_size"`
	MaxSize     float64 `json:"max_size"`
	MinSize     float64 `json:"min_size"`
}

type TaintsState struct {
	Effect string `json:"effect"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}
