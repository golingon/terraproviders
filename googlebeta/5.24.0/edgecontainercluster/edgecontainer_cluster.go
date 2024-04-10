// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package edgecontainercluster

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type MaintenanceEvents struct{}

type Authorization struct {
	// AdminUsers: required
	AdminUsers *AdminUsers `hcl:"admin_users,block" validate:"required"`
}

type AdminUsers struct {
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type ControlPlane struct {
	// Local: optional
	Local *Local `hcl:"local,block"`
	// Remote: optional
	Remote *Remote `hcl:"remote,block"`
}

type Local struct {
	// MachineFilter: string, optional
	MachineFilter terra.StringValue `hcl:"machine_filter,attr"`
	// NodeCount: number, optional
	NodeCount terra.NumberValue `hcl:"node_count,attr"`
	// NodeLocation: string, optional
	NodeLocation terra.StringValue `hcl:"node_location,attr"`
	// SharedDeploymentPolicy: string, optional
	SharedDeploymentPolicy terra.StringValue `hcl:"shared_deployment_policy,attr"`
}

type Remote struct {
	// NodeLocation: string, optional
	NodeLocation terra.StringValue `hcl:"node_location,attr"`
}

type ControlPlaneEncryption struct {
	// KmsKey: string, optional
	KmsKey terra.StringValue `hcl:"kms_key,attr"`
	// KmsStatus: min=0
	KmsStatus []KmsStatus `hcl:"kms_status,block" validate:"min=0"`
}

type KmsStatus struct{}

type Fleet struct {
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
}

type MaintenancePolicy struct {
	// MaintenancePolicyWindow: required
	Window *MaintenancePolicyWindow `hcl:"window,block" validate:"required"`
}

type MaintenancePolicyWindow struct {
	// RecurringWindow: required
	RecurringWindow *RecurringWindow `hcl:"recurring_window,block" validate:"required"`
}

type RecurringWindow struct {
	// Recurrence: string, optional
	Recurrence terra.StringValue `hcl:"recurrence,attr"`
	// RecurringWindowWindow: optional
	Window *RecurringWindowWindow `hcl:"window,block"`
}

type RecurringWindowWindow struct {
	// EndTime: string, optional
	EndTime terra.StringValue `hcl:"end_time,attr"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
}

type Networking struct {
	// ClusterIpv4CidrBlocks: list of string, required
	ClusterIpv4CidrBlocks terra.ListValue[terra.StringValue] `hcl:"cluster_ipv4_cidr_blocks,attr" validate:"required"`
	// ClusterIpv6CidrBlocks: list of string, optional
	ClusterIpv6CidrBlocks terra.ListValue[terra.StringValue] `hcl:"cluster_ipv6_cidr_blocks,attr"`
	// ServicesIpv4CidrBlocks: list of string, required
	ServicesIpv4CidrBlocks terra.ListValue[terra.StringValue] `hcl:"services_ipv4_cidr_blocks,attr" validate:"required"`
	// ServicesIpv6CidrBlocks: list of string, optional
	ServicesIpv6CidrBlocks terra.ListValue[terra.StringValue] `hcl:"services_ipv6_cidr_blocks,attr"`
}

type SystemAddonsConfig struct {
	// Ingress: optional
	Ingress *Ingress `hcl:"ingress,block"`
}

type Ingress struct {
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// Ipv4Vip: string, optional
	Ipv4Vip terra.StringValue `hcl:"ipv4_vip,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type MaintenanceEventsAttributes struct {
	ref terra.Reference
}

func (me MaintenanceEventsAttributes) InternalRef() (terra.Reference, error) {
	return me.ref, nil
}

func (me MaintenanceEventsAttributes) InternalWithRef(ref terra.Reference) MaintenanceEventsAttributes {
	return MaintenanceEventsAttributes{ref: ref}
}

func (me MaintenanceEventsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return me.ref.InternalTokens()
}

func (me MaintenanceEventsAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("create_time"))
}

func (me MaintenanceEventsAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("end_time"))
}

func (me MaintenanceEventsAttributes) Operation() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("operation"))
}

func (me MaintenanceEventsAttributes) Schedule() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("schedule"))
}

func (me MaintenanceEventsAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("start_time"))
}

func (me MaintenanceEventsAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("state"))
}

func (me MaintenanceEventsAttributes) TargetVersion() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("target_version"))
}

func (me MaintenanceEventsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("type"))
}

func (me MaintenanceEventsAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("update_time"))
}

func (me MaintenanceEventsAttributes) Uuid() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("uuid"))
}

type AuthorizationAttributes struct {
	ref terra.Reference
}

func (a AuthorizationAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AuthorizationAttributes) InternalWithRef(ref terra.Reference) AuthorizationAttributes {
	return AuthorizationAttributes{ref: ref}
}

func (a AuthorizationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AuthorizationAttributes) AdminUsers() terra.ListValue[AdminUsersAttributes] {
	return terra.ReferenceAsList[AdminUsersAttributes](a.ref.Append("admin_users"))
}

type AdminUsersAttributes struct {
	ref terra.Reference
}

func (au AdminUsersAttributes) InternalRef() (terra.Reference, error) {
	return au.ref, nil
}

func (au AdminUsersAttributes) InternalWithRef(ref terra.Reference) AdminUsersAttributes {
	return AdminUsersAttributes{ref: ref}
}

func (au AdminUsersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return au.ref.InternalTokens()
}

func (au AdminUsersAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("username"))
}

type ControlPlaneAttributes struct {
	ref terra.Reference
}

func (cp ControlPlaneAttributes) InternalRef() (terra.Reference, error) {
	return cp.ref, nil
}

func (cp ControlPlaneAttributes) InternalWithRef(ref terra.Reference) ControlPlaneAttributes {
	return ControlPlaneAttributes{ref: ref}
}

func (cp ControlPlaneAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cp.ref.InternalTokens()
}

func (cp ControlPlaneAttributes) Local() terra.ListValue[LocalAttributes] {
	return terra.ReferenceAsList[LocalAttributes](cp.ref.Append("local"))
}

func (cp ControlPlaneAttributes) Remote() terra.ListValue[RemoteAttributes] {
	return terra.ReferenceAsList[RemoteAttributes](cp.ref.Append("remote"))
}

type LocalAttributes struct {
	ref terra.Reference
}

func (l LocalAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LocalAttributes) InternalWithRef(ref terra.Reference) LocalAttributes {
	return LocalAttributes{ref: ref}
}

func (l LocalAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l LocalAttributes) MachineFilter() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("machine_filter"))
}

func (l LocalAttributes) NodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("node_count"))
}

func (l LocalAttributes) NodeLocation() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("node_location"))
}

func (l LocalAttributes) SharedDeploymentPolicy() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("shared_deployment_policy"))
}

type RemoteAttributes struct {
	ref terra.Reference
}

func (r RemoteAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RemoteAttributes) InternalWithRef(ref terra.Reference) RemoteAttributes {
	return RemoteAttributes{ref: ref}
}

func (r RemoteAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RemoteAttributes) NodeLocation() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("node_location"))
}

type ControlPlaneEncryptionAttributes struct {
	ref terra.Reference
}

func (cpe ControlPlaneEncryptionAttributes) InternalRef() (terra.Reference, error) {
	return cpe.ref, nil
}

func (cpe ControlPlaneEncryptionAttributes) InternalWithRef(ref terra.Reference) ControlPlaneEncryptionAttributes {
	return ControlPlaneEncryptionAttributes{ref: ref}
}

func (cpe ControlPlaneEncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cpe.ref.InternalTokens()
}

func (cpe ControlPlaneEncryptionAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(cpe.ref.Append("kms_key"))
}

func (cpe ControlPlaneEncryptionAttributes) KmsKeyActiveVersion() terra.StringValue {
	return terra.ReferenceAsString(cpe.ref.Append("kms_key_active_version"))
}

func (cpe ControlPlaneEncryptionAttributes) KmsKeyState() terra.StringValue {
	return terra.ReferenceAsString(cpe.ref.Append("kms_key_state"))
}

func (cpe ControlPlaneEncryptionAttributes) KmsStatus() terra.ListValue[KmsStatusAttributes] {
	return terra.ReferenceAsList[KmsStatusAttributes](cpe.ref.Append("kms_status"))
}

type KmsStatusAttributes struct {
	ref terra.Reference
}

func (ks KmsStatusAttributes) InternalRef() (terra.Reference, error) {
	return ks.ref, nil
}

func (ks KmsStatusAttributes) InternalWithRef(ref terra.Reference) KmsStatusAttributes {
	return KmsStatusAttributes{ref: ref}
}

func (ks KmsStatusAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ks.ref.InternalTokens()
}

func (ks KmsStatusAttributes) Code() terra.NumberValue {
	return terra.ReferenceAsNumber(ks.ref.Append("code"))
}

func (ks KmsStatusAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("message"))
}

type FleetAttributes struct {
	ref terra.Reference
}

func (f FleetAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FleetAttributes) InternalWithRef(ref terra.Reference) FleetAttributes {
	return FleetAttributes{ref: ref}
}

func (f FleetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FleetAttributes) Membership() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("membership"))
}

func (f FleetAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("project"))
}

type MaintenancePolicyAttributes struct {
	ref terra.Reference
}

func (mp MaintenancePolicyAttributes) InternalRef() (terra.Reference, error) {
	return mp.ref, nil
}

func (mp MaintenancePolicyAttributes) InternalWithRef(ref terra.Reference) MaintenancePolicyAttributes {
	return MaintenancePolicyAttributes{ref: ref}
}

func (mp MaintenancePolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mp.ref.InternalTokens()
}

func (mp MaintenancePolicyAttributes) Window() terra.ListValue[MaintenancePolicyWindowAttributes] {
	return terra.ReferenceAsList[MaintenancePolicyWindowAttributes](mp.ref.Append("window"))
}

type MaintenancePolicyWindowAttributes struct {
	ref terra.Reference
}

func (w MaintenancePolicyWindowAttributes) InternalRef() (terra.Reference, error) {
	return w.ref, nil
}

func (w MaintenancePolicyWindowAttributes) InternalWithRef(ref terra.Reference) MaintenancePolicyWindowAttributes {
	return MaintenancePolicyWindowAttributes{ref: ref}
}

func (w MaintenancePolicyWindowAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return w.ref.InternalTokens()
}

func (w MaintenancePolicyWindowAttributes) RecurringWindow() terra.ListValue[RecurringWindowAttributes] {
	return terra.ReferenceAsList[RecurringWindowAttributes](w.ref.Append("recurring_window"))
}

type RecurringWindowAttributes struct {
	ref terra.Reference
}

func (rw RecurringWindowAttributes) InternalRef() (terra.Reference, error) {
	return rw.ref, nil
}

func (rw RecurringWindowAttributes) InternalWithRef(ref terra.Reference) RecurringWindowAttributes {
	return RecurringWindowAttributes{ref: ref}
}

func (rw RecurringWindowAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rw.ref.InternalTokens()
}

func (rw RecurringWindowAttributes) Recurrence() terra.StringValue {
	return terra.ReferenceAsString(rw.ref.Append("recurrence"))
}

func (rw RecurringWindowAttributes) Window() terra.ListValue[RecurringWindowWindowAttributes] {
	return terra.ReferenceAsList[RecurringWindowWindowAttributes](rw.ref.Append("window"))
}

type RecurringWindowWindowAttributes struct {
	ref terra.Reference
}

func (w RecurringWindowWindowAttributes) InternalRef() (terra.Reference, error) {
	return w.ref, nil
}

func (w RecurringWindowWindowAttributes) InternalWithRef(ref terra.Reference) RecurringWindowWindowAttributes {
	return RecurringWindowWindowAttributes{ref: ref}
}

func (w RecurringWindowWindowAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return w.ref.InternalTokens()
}

func (w RecurringWindowWindowAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("end_time"))
}

func (w RecurringWindowWindowAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("start_time"))
}

type NetworkingAttributes struct {
	ref terra.Reference
}

func (n NetworkingAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n NetworkingAttributes) InternalWithRef(ref terra.Reference) NetworkingAttributes {
	return NetworkingAttributes{ref: ref}
}

func (n NetworkingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n NetworkingAttributes) ClusterIpv4CidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](n.ref.Append("cluster_ipv4_cidr_blocks"))
}

func (n NetworkingAttributes) ClusterIpv6CidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](n.ref.Append("cluster_ipv6_cidr_blocks"))
}

func (n NetworkingAttributes) NetworkType() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("network_type"))
}

func (n NetworkingAttributes) ServicesIpv4CidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](n.ref.Append("services_ipv4_cidr_blocks"))
}

func (n NetworkingAttributes) ServicesIpv6CidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](n.ref.Append("services_ipv6_cidr_blocks"))
}

type SystemAddonsConfigAttributes struct {
	ref terra.Reference
}

func (sac SystemAddonsConfigAttributes) InternalRef() (terra.Reference, error) {
	return sac.ref, nil
}

func (sac SystemAddonsConfigAttributes) InternalWithRef(ref terra.Reference) SystemAddonsConfigAttributes {
	return SystemAddonsConfigAttributes{ref: ref}
}

func (sac SystemAddonsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sac.ref.InternalTokens()
}

func (sac SystemAddonsConfigAttributes) Ingress() terra.ListValue[IngressAttributes] {
	return terra.ReferenceAsList[IngressAttributes](sac.ref.Append("ingress"))
}

type IngressAttributes struct {
	ref terra.Reference
}

func (i IngressAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IngressAttributes) InternalWithRef(ref terra.Reference) IngressAttributes {
	return IngressAttributes{ref: ref}
}

func (i IngressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IngressAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("disabled"))
}

func (i IngressAttributes) Ipv4Vip() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("ipv4_vip"))
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

type MaintenanceEventsState struct {
	CreateTime    string `json:"create_time"`
	EndTime       string `json:"end_time"`
	Operation     string `json:"operation"`
	Schedule      string `json:"schedule"`
	StartTime     string `json:"start_time"`
	State         string `json:"state"`
	TargetVersion string `json:"target_version"`
	Type          string `json:"type"`
	UpdateTime    string `json:"update_time"`
	Uuid          string `json:"uuid"`
}

type AuthorizationState struct {
	AdminUsers []AdminUsersState `json:"admin_users"`
}

type AdminUsersState struct {
	Username string `json:"username"`
}

type ControlPlaneState struct {
	Local  []LocalState  `json:"local"`
	Remote []RemoteState `json:"remote"`
}

type LocalState struct {
	MachineFilter          string  `json:"machine_filter"`
	NodeCount              float64 `json:"node_count"`
	NodeLocation           string  `json:"node_location"`
	SharedDeploymentPolicy string  `json:"shared_deployment_policy"`
}

type RemoteState struct {
	NodeLocation string `json:"node_location"`
}

type ControlPlaneEncryptionState struct {
	KmsKey              string           `json:"kms_key"`
	KmsKeyActiveVersion string           `json:"kms_key_active_version"`
	KmsKeyState         string           `json:"kms_key_state"`
	KmsStatus           []KmsStatusState `json:"kms_status"`
}

type KmsStatusState struct {
	Code    float64 `json:"code"`
	Message string  `json:"message"`
}

type FleetState struct {
	Membership string `json:"membership"`
	Project    string `json:"project"`
}

type MaintenancePolicyState struct {
	Window []MaintenancePolicyWindowState `json:"window"`
}

type MaintenancePolicyWindowState struct {
	RecurringWindow []RecurringWindowState `json:"recurring_window"`
}

type RecurringWindowState struct {
	Recurrence string                       `json:"recurrence"`
	Window     []RecurringWindowWindowState `json:"window"`
}

type RecurringWindowWindowState struct {
	EndTime   string `json:"end_time"`
	StartTime string `json:"start_time"`
}

type NetworkingState struct {
	ClusterIpv4CidrBlocks  []string `json:"cluster_ipv4_cidr_blocks"`
	ClusterIpv6CidrBlocks  []string `json:"cluster_ipv6_cidr_blocks"`
	NetworkType            string   `json:"network_type"`
	ServicesIpv4CidrBlocks []string `json:"services_ipv4_cidr_blocks"`
	ServicesIpv6CidrBlocks []string `json:"services_ipv6_cidr_blocks"`
}

type SystemAddonsConfigState struct {
	Ingress []IngressState `json:"ingress"`
}

type IngressState struct {
	Disabled bool   `json:"disabled"`
	Ipv4Vip  string `json:"ipv4_vip"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
