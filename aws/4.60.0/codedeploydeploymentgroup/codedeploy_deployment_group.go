// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package codedeploydeploymentgroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AlarmConfiguration struct {
	// Alarms: set of string, optional
	Alarms terra.SetValue[terra.StringValue] `hcl:"alarms,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// IgnorePollAlarmFailure: bool, optional
	IgnorePollAlarmFailure terra.BoolValue `hcl:"ignore_poll_alarm_failure,attr"`
}

type AutoRollbackConfiguration struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Events: set of string, optional
	Events terra.SetValue[terra.StringValue] `hcl:"events,attr"`
}

type BlueGreenDeploymentConfig struct {
	// DeploymentReadyOption: optional
	DeploymentReadyOption *DeploymentReadyOption `hcl:"deployment_ready_option,block"`
	// GreenFleetProvisioningOption: optional
	GreenFleetProvisioningOption *GreenFleetProvisioningOption `hcl:"green_fleet_provisioning_option,block"`
	// TerminateBlueInstancesOnDeploymentSuccess: optional
	TerminateBlueInstancesOnDeploymentSuccess *TerminateBlueInstancesOnDeploymentSuccess `hcl:"terminate_blue_instances_on_deployment_success,block"`
}

type DeploymentReadyOption struct {
	// ActionOnTimeout: string, optional
	ActionOnTimeout terra.StringValue `hcl:"action_on_timeout,attr"`
	// WaitTimeInMinutes: number, optional
	WaitTimeInMinutes terra.NumberValue `hcl:"wait_time_in_minutes,attr"`
}

type GreenFleetProvisioningOption struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
}

type TerminateBlueInstancesOnDeploymentSuccess struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
	// TerminationWaitTimeInMinutes: number, optional
	TerminationWaitTimeInMinutes terra.NumberValue `hcl:"termination_wait_time_in_minutes,attr"`
}

type DeploymentStyle struct {
	// DeploymentOption: string, optional
	DeploymentOption terra.StringValue `hcl:"deployment_option,attr"`
	// DeploymentType: string, optional
	DeploymentType terra.StringValue `hcl:"deployment_type,attr"`
}

type Ec2TagFilter struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type Ec2TagSet struct {
	// Ec2TagSetEc2TagFilter: min=0
	Ec2TagFilter []Ec2TagSetEc2TagFilter `hcl:"ec2_tag_filter,block" validate:"min=0"`
}

type Ec2TagSetEc2TagFilter struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type EcsService struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
}

type LoadBalancerInfo struct {
	// ElbInfo: min=0
	ElbInfo []ElbInfo `hcl:"elb_info,block" validate:"min=0"`
	// TargetGroupInfo: min=0
	TargetGroupInfo []TargetGroupInfo `hcl:"target_group_info,block" validate:"min=0"`
	// TargetGroupPairInfo: optional
	TargetGroupPairInfo *TargetGroupPairInfo `hcl:"target_group_pair_info,block"`
}

type ElbInfo struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}

type TargetGroupInfo struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}

type TargetGroupPairInfo struct {
	// ProdTrafficRoute: required
	ProdTrafficRoute *ProdTrafficRoute `hcl:"prod_traffic_route,block" validate:"required"`
	// TargetGroup: min=1,max=2
	TargetGroup []TargetGroup `hcl:"target_group,block" validate:"min=1,max=2"`
	// TestTrafficRoute: optional
	TestTrafficRoute *TestTrafficRoute `hcl:"test_traffic_route,block"`
}

type ProdTrafficRoute struct {
	// ListenerArns: set of string, required
	ListenerArns terra.SetValue[terra.StringValue] `hcl:"listener_arns,attr" validate:"required"`
}

type TargetGroup struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type TestTrafficRoute struct {
	// ListenerArns: set of string, required
	ListenerArns terra.SetValue[terra.StringValue] `hcl:"listener_arns,attr" validate:"required"`
}

type OnPremisesInstanceTagFilter struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type TriggerConfiguration struct {
	// TriggerEvents: set of string, required
	TriggerEvents terra.SetValue[terra.StringValue] `hcl:"trigger_events,attr" validate:"required"`
	// TriggerName: string, required
	TriggerName terra.StringValue `hcl:"trigger_name,attr" validate:"required"`
	// TriggerTargetArn: string, required
	TriggerTargetArn terra.StringValue `hcl:"trigger_target_arn,attr" validate:"required"`
}

type AlarmConfigurationAttributes struct {
	ref terra.Reference
}

func (ac AlarmConfigurationAttributes) InternalRef() terra.Reference {
	return ac.ref
}

func (ac AlarmConfigurationAttributes) InternalWithRef(ref terra.Reference) AlarmConfigurationAttributes {
	return AlarmConfigurationAttributes{ref: ref}
}

func (ac AlarmConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return ac.ref.InternalTokens()
}

func (ac AlarmConfigurationAttributes) Alarms() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](ac.ref.Append("alarms"))
}

func (ac AlarmConfigurationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(ac.ref.Append("enabled"))
}

func (ac AlarmConfigurationAttributes) IgnorePollAlarmFailure() terra.BoolValue {
	return terra.ReferenceBool(ac.ref.Append("ignore_poll_alarm_failure"))
}

type AutoRollbackConfigurationAttributes struct {
	ref terra.Reference
}

func (arc AutoRollbackConfigurationAttributes) InternalRef() terra.Reference {
	return arc.ref
}

func (arc AutoRollbackConfigurationAttributes) InternalWithRef(ref terra.Reference) AutoRollbackConfigurationAttributes {
	return AutoRollbackConfigurationAttributes{ref: ref}
}

func (arc AutoRollbackConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return arc.ref.InternalTokens()
}

func (arc AutoRollbackConfigurationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(arc.ref.Append("enabled"))
}

func (arc AutoRollbackConfigurationAttributes) Events() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](arc.ref.Append("events"))
}

type BlueGreenDeploymentConfigAttributes struct {
	ref terra.Reference
}

func (bgdc BlueGreenDeploymentConfigAttributes) InternalRef() terra.Reference {
	return bgdc.ref
}

func (bgdc BlueGreenDeploymentConfigAttributes) InternalWithRef(ref terra.Reference) BlueGreenDeploymentConfigAttributes {
	return BlueGreenDeploymentConfigAttributes{ref: ref}
}

func (bgdc BlueGreenDeploymentConfigAttributes) InternalTokens() hclwrite.Tokens {
	return bgdc.ref.InternalTokens()
}

func (bgdc BlueGreenDeploymentConfigAttributes) DeploymentReadyOption() terra.ListValue[DeploymentReadyOptionAttributes] {
	return terra.ReferenceList[DeploymentReadyOptionAttributes](bgdc.ref.Append("deployment_ready_option"))
}

func (bgdc BlueGreenDeploymentConfigAttributes) GreenFleetProvisioningOption() terra.ListValue[GreenFleetProvisioningOptionAttributes] {
	return terra.ReferenceList[GreenFleetProvisioningOptionAttributes](bgdc.ref.Append("green_fleet_provisioning_option"))
}

func (bgdc BlueGreenDeploymentConfigAttributes) TerminateBlueInstancesOnDeploymentSuccess() terra.ListValue[TerminateBlueInstancesOnDeploymentSuccessAttributes] {
	return terra.ReferenceList[TerminateBlueInstancesOnDeploymentSuccessAttributes](bgdc.ref.Append("terminate_blue_instances_on_deployment_success"))
}

type DeploymentReadyOptionAttributes struct {
	ref terra.Reference
}

func (dro DeploymentReadyOptionAttributes) InternalRef() terra.Reference {
	return dro.ref
}

func (dro DeploymentReadyOptionAttributes) InternalWithRef(ref terra.Reference) DeploymentReadyOptionAttributes {
	return DeploymentReadyOptionAttributes{ref: ref}
}

func (dro DeploymentReadyOptionAttributes) InternalTokens() hclwrite.Tokens {
	return dro.ref.InternalTokens()
}

func (dro DeploymentReadyOptionAttributes) ActionOnTimeout() terra.StringValue {
	return terra.ReferenceString(dro.ref.Append("action_on_timeout"))
}

func (dro DeploymentReadyOptionAttributes) WaitTimeInMinutes() terra.NumberValue {
	return terra.ReferenceNumber(dro.ref.Append("wait_time_in_minutes"))
}

type GreenFleetProvisioningOptionAttributes struct {
	ref terra.Reference
}

func (gfpo GreenFleetProvisioningOptionAttributes) InternalRef() terra.Reference {
	return gfpo.ref
}

func (gfpo GreenFleetProvisioningOptionAttributes) InternalWithRef(ref terra.Reference) GreenFleetProvisioningOptionAttributes {
	return GreenFleetProvisioningOptionAttributes{ref: ref}
}

func (gfpo GreenFleetProvisioningOptionAttributes) InternalTokens() hclwrite.Tokens {
	return gfpo.ref.InternalTokens()
}

func (gfpo GreenFleetProvisioningOptionAttributes) Action() terra.StringValue {
	return terra.ReferenceString(gfpo.ref.Append("action"))
}

type TerminateBlueInstancesOnDeploymentSuccessAttributes struct {
	ref terra.Reference
}

func (tbiods TerminateBlueInstancesOnDeploymentSuccessAttributes) InternalRef() terra.Reference {
	return tbiods.ref
}

func (tbiods TerminateBlueInstancesOnDeploymentSuccessAttributes) InternalWithRef(ref terra.Reference) TerminateBlueInstancesOnDeploymentSuccessAttributes {
	return TerminateBlueInstancesOnDeploymentSuccessAttributes{ref: ref}
}

func (tbiods TerminateBlueInstancesOnDeploymentSuccessAttributes) InternalTokens() hclwrite.Tokens {
	return tbiods.ref.InternalTokens()
}

func (tbiods TerminateBlueInstancesOnDeploymentSuccessAttributes) Action() terra.StringValue {
	return terra.ReferenceString(tbiods.ref.Append("action"))
}

func (tbiods TerminateBlueInstancesOnDeploymentSuccessAttributes) TerminationWaitTimeInMinutes() terra.NumberValue {
	return terra.ReferenceNumber(tbiods.ref.Append("termination_wait_time_in_minutes"))
}

type DeploymentStyleAttributes struct {
	ref terra.Reference
}

func (ds DeploymentStyleAttributes) InternalRef() terra.Reference {
	return ds.ref
}

func (ds DeploymentStyleAttributes) InternalWithRef(ref terra.Reference) DeploymentStyleAttributes {
	return DeploymentStyleAttributes{ref: ref}
}

func (ds DeploymentStyleAttributes) InternalTokens() hclwrite.Tokens {
	return ds.ref.InternalTokens()
}

func (ds DeploymentStyleAttributes) DeploymentOption() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("deployment_option"))
}

func (ds DeploymentStyleAttributes) DeploymentType() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("deployment_type"))
}

type Ec2TagFilterAttributes struct {
	ref terra.Reference
}

func (etf Ec2TagFilterAttributes) InternalRef() terra.Reference {
	return etf.ref
}

func (etf Ec2TagFilterAttributes) InternalWithRef(ref terra.Reference) Ec2TagFilterAttributes {
	return Ec2TagFilterAttributes{ref: ref}
}

func (etf Ec2TagFilterAttributes) InternalTokens() hclwrite.Tokens {
	return etf.ref.InternalTokens()
}

func (etf Ec2TagFilterAttributes) Key() terra.StringValue {
	return terra.ReferenceString(etf.ref.Append("key"))
}

func (etf Ec2TagFilterAttributes) Type() terra.StringValue {
	return terra.ReferenceString(etf.ref.Append("type"))
}

func (etf Ec2TagFilterAttributes) Value() terra.StringValue {
	return terra.ReferenceString(etf.ref.Append("value"))
}

type Ec2TagSetAttributes struct {
	ref terra.Reference
}

func (ets Ec2TagSetAttributes) InternalRef() terra.Reference {
	return ets.ref
}

func (ets Ec2TagSetAttributes) InternalWithRef(ref terra.Reference) Ec2TagSetAttributes {
	return Ec2TagSetAttributes{ref: ref}
}

func (ets Ec2TagSetAttributes) InternalTokens() hclwrite.Tokens {
	return ets.ref.InternalTokens()
}

func (ets Ec2TagSetAttributes) Ec2TagFilter() terra.SetValue[Ec2TagSetEc2TagFilterAttributes] {
	return terra.ReferenceSet[Ec2TagSetEc2TagFilterAttributes](ets.ref.Append("ec2_tag_filter"))
}

type Ec2TagSetEc2TagFilterAttributes struct {
	ref terra.Reference
}

func (etf Ec2TagSetEc2TagFilterAttributes) InternalRef() terra.Reference {
	return etf.ref
}

func (etf Ec2TagSetEc2TagFilterAttributes) InternalWithRef(ref terra.Reference) Ec2TagSetEc2TagFilterAttributes {
	return Ec2TagSetEc2TagFilterAttributes{ref: ref}
}

func (etf Ec2TagSetEc2TagFilterAttributes) InternalTokens() hclwrite.Tokens {
	return etf.ref.InternalTokens()
}

func (etf Ec2TagSetEc2TagFilterAttributes) Key() terra.StringValue {
	return terra.ReferenceString(etf.ref.Append("key"))
}

func (etf Ec2TagSetEc2TagFilterAttributes) Type() terra.StringValue {
	return terra.ReferenceString(etf.ref.Append("type"))
}

func (etf Ec2TagSetEc2TagFilterAttributes) Value() terra.StringValue {
	return terra.ReferenceString(etf.ref.Append("value"))
}

type EcsServiceAttributes struct {
	ref terra.Reference
}

func (es EcsServiceAttributes) InternalRef() terra.Reference {
	return es.ref
}

func (es EcsServiceAttributes) InternalWithRef(ref terra.Reference) EcsServiceAttributes {
	return EcsServiceAttributes{ref: ref}
}

func (es EcsServiceAttributes) InternalTokens() hclwrite.Tokens {
	return es.ref.InternalTokens()
}

func (es EcsServiceAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceString(es.ref.Append("cluster_name"))
}

func (es EcsServiceAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceString(es.ref.Append("service_name"))
}

type LoadBalancerInfoAttributes struct {
	ref terra.Reference
}

func (lbi LoadBalancerInfoAttributes) InternalRef() terra.Reference {
	return lbi.ref
}

func (lbi LoadBalancerInfoAttributes) InternalWithRef(ref terra.Reference) LoadBalancerInfoAttributes {
	return LoadBalancerInfoAttributes{ref: ref}
}

func (lbi LoadBalancerInfoAttributes) InternalTokens() hclwrite.Tokens {
	return lbi.ref.InternalTokens()
}

func (lbi LoadBalancerInfoAttributes) ElbInfo() terra.SetValue[ElbInfoAttributes] {
	return terra.ReferenceSet[ElbInfoAttributes](lbi.ref.Append("elb_info"))
}

func (lbi LoadBalancerInfoAttributes) TargetGroupInfo() terra.SetValue[TargetGroupInfoAttributes] {
	return terra.ReferenceSet[TargetGroupInfoAttributes](lbi.ref.Append("target_group_info"))
}

func (lbi LoadBalancerInfoAttributes) TargetGroupPairInfo() terra.ListValue[TargetGroupPairInfoAttributes] {
	return terra.ReferenceList[TargetGroupPairInfoAttributes](lbi.ref.Append("target_group_pair_info"))
}

type ElbInfoAttributes struct {
	ref terra.Reference
}

func (ei ElbInfoAttributes) InternalRef() terra.Reference {
	return ei.ref
}

func (ei ElbInfoAttributes) InternalWithRef(ref terra.Reference) ElbInfoAttributes {
	return ElbInfoAttributes{ref: ref}
}

func (ei ElbInfoAttributes) InternalTokens() hclwrite.Tokens {
	return ei.ref.InternalTokens()
}

func (ei ElbInfoAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ei.ref.Append("name"))
}

type TargetGroupInfoAttributes struct {
	ref terra.Reference
}

func (tgi TargetGroupInfoAttributes) InternalRef() terra.Reference {
	return tgi.ref
}

func (tgi TargetGroupInfoAttributes) InternalWithRef(ref terra.Reference) TargetGroupInfoAttributes {
	return TargetGroupInfoAttributes{ref: ref}
}

func (tgi TargetGroupInfoAttributes) InternalTokens() hclwrite.Tokens {
	return tgi.ref.InternalTokens()
}

func (tgi TargetGroupInfoAttributes) Name() terra.StringValue {
	return terra.ReferenceString(tgi.ref.Append("name"))
}

type TargetGroupPairInfoAttributes struct {
	ref terra.Reference
}

func (tgpi TargetGroupPairInfoAttributes) InternalRef() terra.Reference {
	return tgpi.ref
}

func (tgpi TargetGroupPairInfoAttributes) InternalWithRef(ref terra.Reference) TargetGroupPairInfoAttributes {
	return TargetGroupPairInfoAttributes{ref: ref}
}

func (tgpi TargetGroupPairInfoAttributes) InternalTokens() hclwrite.Tokens {
	return tgpi.ref.InternalTokens()
}

func (tgpi TargetGroupPairInfoAttributes) ProdTrafficRoute() terra.ListValue[ProdTrafficRouteAttributes] {
	return terra.ReferenceList[ProdTrafficRouteAttributes](tgpi.ref.Append("prod_traffic_route"))
}

func (tgpi TargetGroupPairInfoAttributes) TargetGroup() terra.ListValue[TargetGroupAttributes] {
	return terra.ReferenceList[TargetGroupAttributes](tgpi.ref.Append("target_group"))
}

func (tgpi TargetGroupPairInfoAttributes) TestTrafficRoute() terra.ListValue[TestTrafficRouteAttributes] {
	return terra.ReferenceList[TestTrafficRouteAttributes](tgpi.ref.Append("test_traffic_route"))
}

type ProdTrafficRouteAttributes struct {
	ref terra.Reference
}

func (ptr ProdTrafficRouteAttributes) InternalRef() terra.Reference {
	return ptr.ref
}

func (ptr ProdTrafficRouteAttributes) InternalWithRef(ref terra.Reference) ProdTrafficRouteAttributes {
	return ProdTrafficRouteAttributes{ref: ref}
}

func (ptr ProdTrafficRouteAttributes) InternalTokens() hclwrite.Tokens {
	return ptr.ref.InternalTokens()
}

func (ptr ProdTrafficRouteAttributes) ListenerArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](ptr.ref.Append("listener_arns"))
}

type TargetGroupAttributes struct {
	ref terra.Reference
}

func (tg TargetGroupAttributes) InternalRef() terra.Reference {
	return tg.ref
}

func (tg TargetGroupAttributes) InternalWithRef(ref terra.Reference) TargetGroupAttributes {
	return TargetGroupAttributes{ref: ref}
}

func (tg TargetGroupAttributes) InternalTokens() hclwrite.Tokens {
	return tg.ref.InternalTokens()
}

func (tg TargetGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(tg.ref.Append("name"))
}

type TestTrafficRouteAttributes struct {
	ref terra.Reference
}

func (ttr TestTrafficRouteAttributes) InternalRef() terra.Reference {
	return ttr.ref
}

func (ttr TestTrafficRouteAttributes) InternalWithRef(ref terra.Reference) TestTrafficRouteAttributes {
	return TestTrafficRouteAttributes{ref: ref}
}

func (ttr TestTrafficRouteAttributes) InternalTokens() hclwrite.Tokens {
	return ttr.ref.InternalTokens()
}

func (ttr TestTrafficRouteAttributes) ListenerArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](ttr.ref.Append("listener_arns"))
}

type OnPremisesInstanceTagFilterAttributes struct {
	ref terra.Reference
}

func (opitf OnPremisesInstanceTagFilterAttributes) InternalRef() terra.Reference {
	return opitf.ref
}

func (opitf OnPremisesInstanceTagFilterAttributes) InternalWithRef(ref terra.Reference) OnPremisesInstanceTagFilterAttributes {
	return OnPremisesInstanceTagFilterAttributes{ref: ref}
}

func (opitf OnPremisesInstanceTagFilterAttributes) InternalTokens() hclwrite.Tokens {
	return opitf.ref.InternalTokens()
}

func (opitf OnPremisesInstanceTagFilterAttributes) Key() terra.StringValue {
	return terra.ReferenceString(opitf.ref.Append("key"))
}

func (opitf OnPremisesInstanceTagFilterAttributes) Type() terra.StringValue {
	return terra.ReferenceString(opitf.ref.Append("type"))
}

func (opitf OnPremisesInstanceTagFilterAttributes) Value() terra.StringValue {
	return terra.ReferenceString(opitf.ref.Append("value"))
}

type TriggerConfigurationAttributes struct {
	ref terra.Reference
}

func (tc TriggerConfigurationAttributes) InternalRef() terra.Reference {
	return tc.ref
}

func (tc TriggerConfigurationAttributes) InternalWithRef(ref terra.Reference) TriggerConfigurationAttributes {
	return TriggerConfigurationAttributes{ref: ref}
}

func (tc TriggerConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return tc.ref.InternalTokens()
}

func (tc TriggerConfigurationAttributes) TriggerEvents() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](tc.ref.Append("trigger_events"))
}

func (tc TriggerConfigurationAttributes) TriggerName() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("trigger_name"))
}

func (tc TriggerConfigurationAttributes) TriggerTargetArn() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("trigger_target_arn"))
}

type AlarmConfigurationState struct {
	Alarms                 []string `json:"alarms"`
	Enabled                bool     `json:"enabled"`
	IgnorePollAlarmFailure bool     `json:"ignore_poll_alarm_failure"`
}

type AutoRollbackConfigurationState struct {
	Enabled bool     `json:"enabled"`
	Events  []string `json:"events"`
}

type BlueGreenDeploymentConfigState struct {
	DeploymentReadyOption                     []DeploymentReadyOptionState                     `json:"deployment_ready_option"`
	GreenFleetProvisioningOption              []GreenFleetProvisioningOptionState              `json:"green_fleet_provisioning_option"`
	TerminateBlueInstancesOnDeploymentSuccess []TerminateBlueInstancesOnDeploymentSuccessState `json:"terminate_blue_instances_on_deployment_success"`
}

type DeploymentReadyOptionState struct {
	ActionOnTimeout   string  `json:"action_on_timeout"`
	WaitTimeInMinutes float64 `json:"wait_time_in_minutes"`
}

type GreenFleetProvisioningOptionState struct {
	Action string `json:"action"`
}

type TerminateBlueInstancesOnDeploymentSuccessState struct {
	Action                       string  `json:"action"`
	TerminationWaitTimeInMinutes float64 `json:"termination_wait_time_in_minutes"`
}

type DeploymentStyleState struct {
	DeploymentOption string `json:"deployment_option"`
	DeploymentType   string `json:"deployment_type"`
}

type Ec2TagFilterState struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Ec2TagSetState struct {
	Ec2TagFilter []Ec2TagSetEc2TagFilterState `json:"ec2_tag_filter"`
}

type Ec2TagSetEc2TagFilterState struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type EcsServiceState struct {
	ClusterName string `json:"cluster_name"`
	ServiceName string `json:"service_name"`
}

type LoadBalancerInfoState struct {
	ElbInfo             []ElbInfoState             `json:"elb_info"`
	TargetGroupInfo     []TargetGroupInfoState     `json:"target_group_info"`
	TargetGroupPairInfo []TargetGroupPairInfoState `json:"target_group_pair_info"`
}

type ElbInfoState struct {
	Name string `json:"name"`
}

type TargetGroupInfoState struct {
	Name string `json:"name"`
}

type TargetGroupPairInfoState struct {
	ProdTrafficRoute []ProdTrafficRouteState `json:"prod_traffic_route"`
	TargetGroup      []TargetGroupState      `json:"target_group"`
	TestTrafficRoute []TestTrafficRouteState `json:"test_traffic_route"`
}

type ProdTrafficRouteState struct {
	ListenerArns []string `json:"listener_arns"`
}

type TargetGroupState struct {
	Name string `json:"name"`
}

type TestTrafficRouteState struct {
	ListenerArns []string `json:"listener_arns"`
}

type OnPremisesInstanceTagFilterState struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type TriggerConfigurationState struct {
	TriggerEvents    []string `json:"trigger_events"`
	TriggerName      string   `json:"trigger_name"`
	TriggerTargetArn string   `json:"trigger_target_arn"`
}
