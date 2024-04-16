// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_redis_instance

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type MaintenancePolicy struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// MaintenancePolicyWeeklyMaintenanceWindow: min=0
	WeeklyMaintenanceWindow []MaintenancePolicyWeeklyMaintenanceWindow `hcl:"weekly_maintenance_window,block" validate:"min=0"`
}

type MaintenancePolicyWeeklyMaintenanceWindow struct {
	// Day: string, required
	Day terra.StringValue `hcl:"day,attr" validate:"required"`
	// MaintenancePolicyWeeklyMaintenanceWindowStartTime: required
	StartTime *MaintenancePolicyWeeklyMaintenanceWindowStartTime `hcl:"start_time,block" validate:"required"`
}

type MaintenancePolicyWeeklyMaintenanceWindowStartTime struct {
	// Hours: number, optional
	Hours terra.NumberValue `hcl:"hours,attr"`
	// Minutes: number, optional
	Minutes terra.NumberValue `hcl:"minutes,attr"`
	// Nanos: number, optional
	Nanos terra.NumberValue `hcl:"nanos,attr"`
	// Seconds: number, optional
	Seconds terra.NumberValue `hcl:"seconds,attr"`
}

type PersistenceConfig struct {
	// PersistenceMode: string, optional
	PersistenceMode terra.StringValue `hcl:"persistence_mode,attr"`
	// RdbSnapshotPeriod: string, optional
	RdbSnapshotPeriod terra.StringValue `hcl:"rdb_snapshot_period,attr"`
	// RdbSnapshotStartTime: string, optional
	RdbSnapshotStartTime terra.StringValue `hcl:"rdb_snapshot_start_time,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type MaintenanceScheduleAttributes struct {
	ref terra.Reference
}

func (ms MaintenanceScheduleAttributes) InternalRef() (terra.Reference, error) {
	return ms.ref, nil
}

func (ms MaintenanceScheduleAttributes) InternalWithRef(ref terra.Reference) MaintenanceScheduleAttributes {
	return MaintenanceScheduleAttributes{ref: ref}
}

func (ms MaintenanceScheduleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ms.ref.InternalTokens()
}

func (ms MaintenanceScheduleAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("end_time"))
}

func (ms MaintenanceScheduleAttributes) ScheduleDeadlineTime() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("schedule_deadline_time"))
}

func (ms MaintenanceScheduleAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("start_time"))
}

type NodesAttributes struct {
	ref terra.Reference
}

func (n NodesAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n NodesAttributes) InternalWithRef(ref terra.Reference) NodesAttributes {
	return NodesAttributes{ref: ref}
}

func (n NodesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n NodesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("id"))
}

func (n NodesAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("zone"))
}

type ServerCaCertsAttributes struct {
	ref terra.Reference
}

func (scc ServerCaCertsAttributes) InternalRef() (terra.Reference, error) {
	return scc.ref, nil
}

func (scc ServerCaCertsAttributes) InternalWithRef(ref terra.Reference) ServerCaCertsAttributes {
	return ServerCaCertsAttributes{ref: ref}
}

func (scc ServerCaCertsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return scc.ref.InternalTokens()
}

func (scc ServerCaCertsAttributes) Cert() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("cert"))
}

func (scc ServerCaCertsAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("create_time"))
}

func (scc ServerCaCertsAttributes) ExpireTime() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("expire_time"))
}

func (scc ServerCaCertsAttributes) SerialNumber() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("serial_number"))
}

func (scc ServerCaCertsAttributes) Sha1Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("sha1_fingerprint"))
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

func (mp MaintenancePolicyAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(mp.ref.Append("create_time"))
}

func (mp MaintenancePolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mp.ref.Append("description"))
}

func (mp MaintenancePolicyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(mp.ref.Append("update_time"))
}

func (mp MaintenancePolicyAttributes) WeeklyMaintenanceWindow() terra.ListValue[MaintenancePolicyWeeklyMaintenanceWindowAttributes] {
	return terra.ReferenceAsList[MaintenancePolicyWeeklyMaintenanceWindowAttributes](mp.ref.Append("weekly_maintenance_window"))
}

type MaintenancePolicyWeeklyMaintenanceWindowAttributes struct {
	ref terra.Reference
}

func (wmw MaintenancePolicyWeeklyMaintenanceWindowAttributes) InternalRef() (terra.Reference, error) {
	return wmw.ref, nil
}

func (wmw MaintenancePolicyWeeklyMaintenanceWindowAttributes) InternalWithRef(ref terra.Reference) MaintenancePolicyWeeklyMaintenanceWindowAttributes {
	return MaintenancePolicyWeeklyMaintenanceWindowAttributes{ref: ref}
}

func (wmw MaintenancePolicyWeeklyMaintenanceWindowAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wmw.ref.InternalTokens()
}

func (wmw MaintenancePolicyWeeklyMaintenanceWindowAttributes) Day() terra.StringValue {
	return terra.ReferenceAsString(wmw.ref.Append("day"))
}

func (wmw MaintenancePolicyWeeklyMaintenanceWindowAttributes) Duration() terra.StringValue {
	return terra.ReferenceAsString(wmw.ref.Append("duration"))
}

func (wmw MaintenancePolicyWeeklyMaintenanceWindowAttributes) StartTime() terra.ListValue[MaintenancePolicyWeeklyMaintenanceWindowStartTimeAttributes] {
	return terra.ReferenceAsList[MaintenancePolicyWeeklyMaintenanceWindowStartTimeAttributes](wmw.ref.Append("start_time"))
}

type MaintenancePolicyWeeklyMaintenanceWindowStartTimeAttributes struct {
	ref terra.Reference
}

func (st MaintenancePolicyWeeklyMaintenanceWindowStartTimeAttributes) InternalRef() (terra.Reference, error) {
	return st.ref, nil
}

func (st MaintenancePolicyWeeklyMaintenanceWindowStartTimeAttributes) InternalWithRef(ref terra.Reference) MaintenancePolicyWeeklyMaintenanceWindowStartTimeAttributes {
	return MaintenancePolicyWeeklyMaintenanceWindowStartTimeAttributes{ref: ref}
}

func (st MaintenancePolicyWeeklyMaintenanceWindowStartTimeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return st.ref.InternalTokens()
}

func (st MaintenancePolicyWeeklyMaintenanceWindowStartTimeAttributes) Hours() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("hours"))
}

func (st MaintenancePolicyWeeklyMaintenanceWindowStartTimeAttributes) Minutes() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("minutes"))
}

func (st MaintenancePolicyWeeklyMaintenanceWindowStartTimeAttributes) Nanos() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("nanos"))
}

func (st MaintenancePolicyWeeklyMaintenanceWindowStartTimeAttributes) Seconds() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("seconds"))
}

type PersistenceConfigAttributes struct {
	ref terra.Reference
}

func (pc PersistenceConfigAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PersistenceConfigAttributes) InternalWithRef(ref terra.Reference) PersistenceConfigAttributes {
	return PersistenceConfigAttributes{ref: ref}
}

func (pc PersistenceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PersistenceConfigAttributes) PersistenceMode() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("persistence_mode"))
}

func (pc PersistenceConfigAttributes) RdbNextSnapshotTime() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("rdb_next_snapshot_time"))
}

func (pc PersistenceConfigAttributes) RdbSnapshotPeriod() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("rdb_snapshot_period"))
}

func (pc PersistenceConfigAttributes) RdbSnapshotStartTime() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("rdb_snapshot_start_time"))
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

type MaintenanceScheduleState struct {
	EndTime              string `json:"end_time"`
	ScheduleDeadlineTime string `json:"schedule_deadline_time"`
	StartTime            string `json:"start_time"`
}

type NodesState struct {
	Id   string `json:"id"`
	Zone string `json:"zone"`
}

type ServerCaCertsState struct {
	Cert            string `json:"cert"`
	CreateTime      string `json:"create_time"`
	ExpireTime      string `json:"expire_time"`
	SerialNumber    string `json:"serial_number"`
	Sha1Fingerprint string `json:"sha1_fingerprint"`
}

type MaintenancePolicyState struct {
	CreateTime              string                                          `json:"create_time"`
	Description             string                                          `json:"description"`
	UpdateTime              string                                          `json:"update_time"`
	WeeklyMaintenanceWindow []MaintenancePolicyWeeklyMaintenanceWindowState `json:"weekly_maintenance_window"`
}

type MaintenancePolicyWeeklyMaintenanceWindowState struct {
	Day       string                                                   `json:"day"`
	Duration  string                                                   `json:"duration"`
	StartTime []MaintenancePolicyWeeklyMaintenanceWindowStartTimeState `json:"start_time"`
}

type MaintenancePolicyWeeklyMaintenanceWindowStartTimeState struct {
	Hours   float64 `json:"hours"`
	Minutes float64 `json:"minutes"`
	Nanos   float64 `json:"nanos"`
	Seconds float64 `json:"seconds"`
}

type PersistenceConfigState struct {
	PersistenceMode      string `json:"persistence_mode"`
	RdbNextSnapshotTime  string `json:"rdb_next_snapshot_time"`
	RdbSnapshotPeriod    string `json:"rdb_snapshot_period"`
	RdbSnapshotStartTime string `json:"rdb_snapshot_start_time"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
