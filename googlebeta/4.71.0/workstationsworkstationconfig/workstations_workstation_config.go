// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package workstationsworkstationconfig

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Conditions struct{}

type Container struct {
	// Args: list of string, optional
	Args terra.ListValue[terra.StringValue] `hcl:"args,attr"`
	// Command: list of string, optional
	Command terra.ListValue[terra.StringValue] `hcl:"command,attr"`
	// Env: map of string, optional
	Env terra.MapValue[terra.StringValue] `hcl:"env,attr"`
	// Image: string, optional
	Image terra.StringValue `hcl:"image,attr"`
	// RunAsUser: number, optional
	RunAsUser terra.NumberValue `hcl:"run_as_user,attr"`
	// WorkingDir: string, optional
	WorkingDir terra.StringValue `hcl:"working_dir,attr"`
}

type EncryptionKey struct {
	// KmsKey: string, required
	KmsKey terra.StringValue `hcl:"kms_key,attr" validate:"required"`
	// KmsKeyServiceAccount: string, required
	KmsKeyServiceAccount terra.StringValue `hcl:"kms_key_service_account,attr" validate:"required"`
}

type Host struct {
	// GceInstance: optional
	GceInstance *GceInstance `hcl:"gce_instance,block"`
}

type GceInstance struct {
	// BootDiskSizeGb: number, optional
	BootDiskSizeGb terra.NumberValue `hcl:"boot_disk_size_gb,attr"`
	// DisablePublicIpAddresses: bool, optional
	DisablePublicIpAddresses terra.BoolValue `hcl:"disable_public_ip_addresses,attr"`
	// MachineType: string, optional
	MachineType terra.StringValue `hcl:"machine_type,attr"`
	// PoolSize: number, optional
	PoolSize terra.NumberValue `hcl:"pool_size,attr"`
	// ServiceAccount: string, optional
	ServiceAccount terra.StringValue `hcl:"service_account,attr"`
	// Tags: list of string, optional
	Tags terra.ListValue[terra.StringValue] `hcl:"tags,attr"`
	// ConfidentialInstanceConfig: optional
	ConfidentialInstanceConfig *ConfidentialInstanceConfig `hcl:"confidential_instance_config,block"`
	// ShieldedInstanceConfig: optional
	ShieldedInstanceConfig *ShieldedInstanceConfig `hcl:"shielded_instance_config,block"`
}

type ConfidentialInstanceConfig struct {
	// EnableConfidentialCompute: bool, optional
	EnableConfidentialCompute terra.BoolValue `hcl:"enable_confidential_compute,attr"`
}

type ShieldedInstanceConfig struct {
	// EnableIntegrityMonitoring: bool, optional
	EnableIntegrityMonitoring terra.BoolValue `hcl:"enable_integrity_monitoring,attr"`
	// EnableSecureBoot: bool, optional
	EnableSecureBoot terra.BoolValue `hcl:"enable_secure_boot,attr"`
	// EnableVtpm: bool, optional
	EnableVtpm terra.BoolValue `hcl:"enable_vtpm,attr"`
}

type PersistentDirectories struct {
	// MountPath: string, optional
	MountPath terra.StringValue `hcl:"mount_path,attr"`
	// GcePd: optional
	GcePd *GcePd `hcl:"gce_pd,block"`
}

type GcePd struct {
	// DiskType: string, optional
	DiskType terra.StringValue `hcl:"disk_type,attr"`
	// FsType: string, optional
	FsType terra.StringValue `hcl:"fs_type,attr"`
	// ReclaimPolicy: string, optional
	ReclaimPolicy terra.StringValue `hcl:"reclaim_policy,attr"`
	// SizeGb: number, optional
	SizeGb terra.NumberValue `hcl:"size_gb,attr"`
	// SourceSnapshot: string, optional
	SourceSnapshot terra.StringValue `hcl:"source_snapshot,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ConditionsAttributes struct {
	ref terra.Reference
}

func (c ConditionsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConditionsAttributes) InternalWithRef(ref terra.Reference) ConditionsAttributes {
	return ConditionsAttributes{ref: ref}
}

func (c ConditionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConditionsAttributes) Code() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("code"))
}

func (c ConditionsAttributes) Details() terra.ListValue[terra.MapValue[terra.StringValue]] {
	return terra.ReferenceAsList[terra.MapValue[terra.StringValue]](c.ref.Append("details"))
}

func (c ConditionsAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("message"))
}

type ContainerAttributes struct {
	ref terra.Reference
}

func (c ContainerAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ContainerAttributes) InternalWithRef(ref terra.Reference) ContainerAttributes {
	return ContainerAttributes{ref: ref}
}

func (c ContainerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ContainerAttributes) Args() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("args"))
}

func (c ContainerAttributes) Command() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("command"))
}

func (c ContainerAttributes) Env() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("env"))
}

func (c ContainerAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("image"))
}

func (c ContainerAttributes) RunAsUser() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("run_as_user"))
}

func (c ContainerAttributes) WorkingDir() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("working_dir"))
}

type EncryptionKeyAttributes struct {
	ref terra.Reference
}

func (ek EncryptionKeyAttributes) InternalRef() (terra.Reference, error) {
	return ek.ref, nil
}

func (ek EncryptionKeyAttributes) InternalWithRef(ref terra.Reference) EncryptionKeyAttributes {
	return EncryptionKeyAttributes{ref: ref}
}

func (ek EncryptionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ek.ref.InternalTokens()
}

func (ek EncryptionKeyAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(ek.ref.Append("kms_key"))
}

func (ek EncryptionKeyAttributes) KmsKeyServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(ek.ref.Append("kms_key_service_account"))
}

type HostAttributes struct {
	ref terra.Reference
}

func (h HostAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h HostAttributes) InternalWithRef(ref terra.Reference) HostAttributes {
	return HostAttributes{ref: ref}
}

func (h HostAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h HostAttributes) GceInstance() terra.ListValue[GceInstanceAttributes] {
	return terra.ReferenceAsList[GceInstanceAttributes](h.ref.Append("gce_instance"))
}

type GceInstanceAttributes struct {
	ref terra.Reference
}

func (gi GceInstanceAttributes) InternalRef() (terra.Reference, error) {
	return gi.ref, nil
}

func (gi GceInstanceAttributes) InternalWithRef(ref terra.Reference) GceInstanceAttributes {
	return GceInstanceAttributes{ref: ref}
}

func (gi GceInstanceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gi.ref.InternalTokens()
}

func (gi GceInstanceAttributes) BootDiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(gi.ref.Append("boot_disk_size_gb"))
}

func (gi GceInstanceAttributes) DisablePublicIpAddresses() terra.BoolValue {
	return terra.ReferenceAsBool(gi.ref.Append("disable_public_ip_addresses"))
}

func (gi GceInstanceAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(gi.ref.Append("machine_type"))
}

func (gi GceInstanceAttributes) PoolSize() terra.NumberValue {
	return terra.ReferenceAsNumber(gi.ref.Append("pool_size"))
}

func (gi GceInstanceAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(gi.ref.Append("service_account"))
}

func (gi GceInstanceAttributes) Tags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gi.ref.Append("tags"))
}

func (gi GceInstanceAttributes) ConfidentialInstanceConfig() terra.ListValue[ConfidentialInstanceConfigAttributes] {
	return terra.ReferenceAsList[ConfidentialInstanceConfigAttributes](gi.ref.Append("confidential_instance_config"))
}

func (gi GceInstanceAttributes) ShieldedInstanceConfig() terra.ListValue[ShieldedInstanceConfigAttributes] {
	return terra.ReferenceAsList[ShieldedInstanceConfigAttributes](gi.ref.Append("shielded_instance_config"))
}

type ConfidentialInstanceConfigAttributes struct {
	ref terra.Reference
}

func (cic ConfidentialInstanceConfigAttributes) InternalRef() (terra.Reference, error) {
	return cic.ref, nil
}

func (cic ConfidentialInstanceConfigAttributes) InternalWithRef(ref terra.Reference) ConfidentialInstanceConfigAttributes {
	return ConfidentialInstanceConfigAttributes{ref: ref}
}

func (cic ConfidentialInstanceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cic.ref.InternalTokens()
}

func (cic ConfidentialInstanceConfigAttributes) EnableConfidentialCompute() terra.BoolValue {
	return terra.ReferenceAsBool(cic.ref.Append("enable_confidential_compute"))
}

type ShieldedInstanceConfigAttributes struct {
	ref terra.Reference
}

func (sic ShieldedInstanceConfigAttributes) InternalRef() (terra.Reference, error) {
	return sic.ref, nil
}

func (sic ShieldedInstanceConfigAttributes) InternalWithRef(ref terra.Reference) ShieldedInstanceConfigAttributes {
	return ShieldedInstanceConfigAttributes{ref: ref}
}

func (sic ShieldedInstanceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sic.ref.InternalTokens()
}

func (sic ShieldedInstanceConfigAttributes) EnableIntegrityMonitoring() terra.BoolValue {
	return terra.ReferenceAsBool(sic.ref.Append("enable_integrity_monitoring"))
}

func (sic ShieldedInstanceConfigAttributes) EnableSecureBoot() terra.BoolValue {
	return terra.ReferenceAsBool(sic.ref.Append("enable_secure_boot"))
}

func (sic ShieldedInstanceConfigAttributes) EnableVtpm() terra.BoolValue {
	return terra.ReferenceAsBool(sic.ref.Append("enable_vtpm"))
}

type PersistentDirectoriesAttributes struct {
	ref terra.Reference
}

func (pd PersistentDirectoriesAttributes) InternalRef() (terra.Reference, error) {
	return pd.ref, nil
}

func (pd PersistentDirectoriesAttributes) InternalWithRef(ref terra.Reference) PersistentDirectoriesAttributes {
	return PersistentDirectoriesAttributes{ref: ref}
}

func (pd PersistentDirectoriesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pd.ref.InternalTokens()
}

func (pd PersistentDirectoriesAttributes) MountPath() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("mount_path"))
}

func (pd PersistentDirectoriesAttributes) GcePd() terra.ListValue[GcePdAttributes] {
	return terra.ReferenceAsList[GcePdAttributes](pd.ref.Append("gce_pd"))
}

type GcePdAttributes struct {
	ref terra.Reference
}

func (gp GcePdAttributes) InternalRef() (terra.Reference, error) {
	return gp.ref, nil
}

func (gp GcePdAttributes) InternalWithRef(ref terra.Reference) GcePdAttributes {
	return GcePdAttributes{ref: ref}
}

func (gp GcePdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gp.ref.InternalTokens()
}

func (gp GcePdAttributes) DiskType() terra.StringValue {
	return terra.ReferenceAsString(gp.ref.Append("disk_type"))
}

func (gp GcePdAttributes) FsType() terra.StringValue {
	return terra.ReferenceAsString(gp.ref.Append("fs_type"))
}

func (gp GcePdAttributes) ReclaimPolicy() terra.StringValue {
	return terra.ReferenceAsString(gp.ref.Append("reclaim_policy"))
}

func (gp GcePdAttributes) SizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(gp.ref.Append("size_gb"))
}

func (gp GcePdAttributes) SourceSnapshot() terra.StringValue {
	return terra.ReferenceAsString(gp.ref.Append("source_snapshot"))
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

type ConditionsState struct {
	Code    float64             `json:"code"`
	Details []map[string]string `json:"details"`
	Message string              `json:"message"`
}

type ContainerState struct {
	Args       []string          `json:"args"`
	Command    []string          `json:"command"`
	Env        map[string]string `json:"env"`
	Image      string            `json:"image"`
	RunAsUser  float64           `json:"run_as_user"`
	WorkingDir string            `json:"working_dir"`
}

type EncryptionKeyState struct {
	KmsKey               string `json:"kms_key"`
	KmsKeyServiceAccount string `json:"kms_key_service_account"`
}

type HostState struct {
	GceInstance []GceInstanceState `json:"gce_instance"`
}

type GceInstanceState struct {
	BootDiskSizeGb             float64                           `json:"boot_disk_size_gb"`
	DisablePublicIpAddresses   bool                              `json:"disable_public_ip_addresses"`
	MachineType                string                            `json:"machine_type"`
	PoolSize                   float64                           `json:"pool_size"`
	ServiceAccount             string                            `json:"service_account"`
	Tags                       []string                          `json:"tags"`
	ConfidentialInstanceConfig []ConfidentialInstanceConfigState `json:"confidential_instance_config"`
	ShieldedInstanceConfig     []ShieldedInstanceConfigState     `json:"shielded_instance_config"`
}

type ConfidentialInstanceConfigState struct {
	EnableConfidentialCompute bool `json:"enable_confidential_compute"`
}

type ShieldedInstanceConfigState struct {
	EnableIntegrityMonitoring bool `json:"enable_integrity_monitoring"`
	EnableSecureBoot          bool `json:"enable_secure_boot"`
	EnableVtpm                bool `json:"enable_vtpm"`
}

type PersistentDirectoriesState struct {
	MountPath string       `json:"mount_path"`
	GcePd     []GcePdState `json:"gce_pd"`
}

type GcePdState struct {
	DiskType       string  `json:"disk_type"`
	FsType         string  `json:"fs_type"`
	ReclaimPolicy  string  `json:"reclaim_policy"`
	SizeGb         float64 `json:"size_gb"`
	SourceSnapshot string  `json:"source_snapshot"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
