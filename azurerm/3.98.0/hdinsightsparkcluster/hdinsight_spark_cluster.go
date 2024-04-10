// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package hdinsightsparkcluster

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ComponentVersion struct {
	// Spark: string, required
	Spark terra.StringValue `hcl:"spark,attr" validate:"required"`
}

type ComputeIsolation struct {
	// ComputeIsolationEnabled: bool, optional
	ComputeIsolationEnabled terra.BoolValue `hcl:"compute_isolation_enabled,attr"`
	// HostSku: string, optional
	HostSku terra.StringValue `hcl:"host_sku,attr"`
}

type DiskEncryption struct {
	// EncryptionAlgorithm: string, optional
	EncryptionAlgorithm terra.StringValue `hcl:"encryption_algorithm,attr"`
	// EncryptionAtHostEnabled: bool, optional
	EncryptionAtHostEnabled terra.BoolValue `hcl:"encryption_at_host_enabled,attr"`
	// KeyVaultKeyId: string, optional
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr"`
	// KeyVaultManagedIdentityId: string, optional
	KeyVaultManagedIdentityId terra.StringValue `hcl:"key_vault_managed_identity_id,attr"`
}

type Extension struct {
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// PrimaryKey: string, required
	PrimaryKey terra.StringValue `hcl:"primary_key,attr" validate:"required"`
}

type Gateway struct {
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type Metastores struct {
	// Ambari: optional
	Ambari *Ambari `hcl:"ambari,block"`
	// Hive: optional
	Hive *Hive `hcl:"hive,block"`
	// Oozie: optional
	Oozie *Oozie `hcl:"oozie,block"`
}

type Ambari struct {
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Server: string, required
	Server terra.StringValue `hcl:"server,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type Hive struct {
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Server: string, required
	Server terra.StringValue `hcl:"server,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type Oozie struct {
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Server: string, required
	Server terra.StringValue `hcl:"server,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type Monitor struct {
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// PrimaryKey: string, required
	PrimaryKey terra.StringValue `hcl:"primary_key,attr" validate:"required"`
}

type Network struct {
	// ConnectionDirection: string, optional
	ConnectionDirection terra.StringValue `hcl:"connection_direction,attr"`
	// PrivateLinkEnabled: bool, optional
	PrivateLinkEnabled terra.BoolValue `hcl:"private_link_enabled,attr"`
}

type Roles struct {
	// HeadNode: required
	HeadNode *HeadNode `hcl:"head_node,block" validate:"required"`
	// WorkerNode: required
	WorkerNode *WorkerNode `hcl:"worker_node,block" validate:"required"`
	// ZookeeperNode: required
	ZookeeperNode *ZookeeperNode `hcl:"zookeeper_node,block" validate:"required"`
}

type HeadNode struct {
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// SshKeys: set of string, optional
	SshKeys terra.SetValue[terra.StringValue] `hcl:"ssh_keys,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
	// VirtualNetworkId: string, optional
	VirtualNetworkId terra.StringValue `hcl:"virtual_network_id,attr"`
	// VmSize: string, required
	VmSize terra.StringValue `hcl:"vm_size,attr" validate:"required"`
	// HeadNodeScriptActions: min=0
	ScriptActions []HeadNodeScriptActions `hcl:"script_actions,block" validate:"min=0"`
}

type HeadNodeScriptActions struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: string, optional
	Parameters terra.StringValue `hcl:"parameters,attr"`
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
}

type WorkerNode struct {
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// SshKeys: set of string, optional
	SshKeys terra.SetValue[terra.StringValue] `hcl:"ssh_keys,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
	// TargetInstanceCount: number, required
	TargetInstanceCount terra.NumberValue `hcl:"target_instance_count,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
	// VirtualNetworkId: string, optional
	VirtualNetworkId terra.StringValue `hcl:"virtual_network_id,attr"`
	// VmSize: string, required
	VmSize terra.StringValue `hcl:"vm_size,attr" validate:"required"`
	// Autoscale: optional
	Autoscale *Autoscale `hcl:"autoscale,block"`
	// WorkerNodeScriptActions: min=0
	ScriptActions []WorkerNodeScriptActions `hcl:"script_actions,block" validate:"min=0"`
}

type Autoscale struct {
	// Capacity: optional
	Capacity *Capacity `hcl:"capacity,block"`
	// Recurrence: optional
	Recurrence *Recurrence `hcl:"recurrence,block"`
}

type Capacity struct {
	// MaxInstanceCount: number, required
	MaxInstanceCount terra.NumberValue `hcl:"max_instance_count,attr" validate:"required"`
	// MinInstanceCount: number, required
	MinInstanceCount terra.NumberValue `hcl:"min_instance_count,attr" validate:"required"`
}

type Recurrence struct {
	// Timezone: string, required
	Timezone terra.StringValue `hcl:"timezone,attr" validate:"required"`
	// Schedule: min=1
	Schedule []Schedule `hcl:"schedule,block" validate:"min=1"`
}

type Schedule struct {
	// Days: list of string, required
	Days terra.ListValue[terra.StringValue] `hcl:"days,attr" validate:"required"`
	// TargetInstanceCount: number, required
	TargetInstanceCount terra.NumberValue `hcl:"target_instance_count,attr" validate:"required"`
	// Time: string, required
	Time terra.StringValue `hcl:"time,attr" validate:"required"`
}

type WorkerNodeScriptActions struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: string, optional
	Parameters terra.StringValue `hcl:"parameters,attr"`
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
}

type ZookeeperNode struct {
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// SshKeys: set of string, optional
	SshKeys terra.SetValue[terra.StringValue] `hcl:"ssh_keys,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
	// VirtualNetworkId: string, optional
	VirtualNetworkId terra.StringValue `hcl:"virtual_network_id,attr"`
	// VmSize: string, required
	VmSize terra.StringValue `hcl:"vm_size,attr" validate:"required"`
	// ZookeeperNodeScriptActions: min=0
	ScriptActions []ZookeeperNodeScriptActions `hcl:"script_actions,block" validate:"min=0"`
}

type ZookeeperNodeScriptActions struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: string, optional
	Parameters terra.StringValue `hcl:"parameters,attr"`
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
}

type SecurityProfile struct {
	// AaddsResourceId: string, required
	AaddsResourceId terra.StringValue `hcl:"aadds_resource_id,attr" validate:"required"`
	// ClusterUsersGroupDns: set of string, optional
	ClusterUsersGroupDns terra.SetValue[terra.StringValue] `hcl:"cluster_users_group_dns,attr"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// DomainUserPassword: string, required
	DomainUserPassword terra.StringValue `hcl:"domain_user_password,attr" validate:"required"`
	// DomainUsername: string, required
	DomainUsername terra.StringValue `hcl:"domain_username,attr" validate:"required"`
	// LdapsUrls: set of string, required
	LdapsUrls terra.SetValue[terra.StringValue] `hcl:"ldaps_urls,attr" validate:"required"`
	// MsiResourceId: string, required
	MsiResourceId terra.StringValue `hcl:"msi_resource_id,attr" validate:"required"`
}

type StorageAccount struct {
	// IsDefault: bool, required
	IsDefault terra.BoolValue `hcl:"is_default,attr" validate:"required"`
	// StorageAccountKey: string, required
	StorageAccountKey terra.StringValue `hcl:"storage_account_key,attr" validate:"required"`
	// StorageContainerId: string, required
	StorageContainerId terra.StringValue `hcl:"storage_container_id,attr" validate:"required"`
	// StorageResourceId: string, optional
	StorageResourceId terra.StringValue `hcl:"storage_resource_id,attr"`
}

type StorageAccountGen2 struct {
	// FilesystemId: string, required
	FilesystemId terra.StringValue `hcl:"filesystem_id,attr" validate:"required"`
	// IsDefault: bool, required
	IsDefault terra.BoolValue `hcl:"is_default,attr" validate:"required"`
	// ManagedIdentityResourceId: string, required
	ManagedIdentityResourceId terra.StringValue `hcl:"managed_identity_resource_id,attr" validate:"required"`
	// StorageResourceId: string, required
	StorageResourceId terra.StringValue `hcl:"storage_resource_id,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ComponentVersionAttributes struct {
	ref terra.Reference
}

func (cv ComponentVersionAttributes) InternalRef() (terra.Reference, error) {
	return cv.ref, nil
}

func (cv ComponentVersionAttributes) InternalWithRef(ref terra.Reference) ComponentVersionAttributes {
	return ComponentVersionAttributes{ref: ref}
}

func (cv ComponentVersionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cv.ref.InternalTokens()
}

func (cv ComponentVersionAttributes) Spark() terra.StringValue {
	return terra.ReferenceAsString(cv.ref.Append("spark"))
}

type ComputeIsolationAttributes struct {
	ref terra.Reference
}

func (ci ComputeIsolationAttributes) InternalRef() (terra.Reference, error) {
	return ci.ref, nil
}

func (ci ComputeIsolationAttributes) InternalWithRef(ref terra.Reference) ComputeIsolationAttributes {
	return ComputeIsolationAttributes{ref: ref}
}

func (ci ComputeIsolationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ci.ref.InternalTokens()
}

func (ci ComputeIsolationAttributes) ComputeIsolationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("compute_isolation_enabled"))
}

func (ci ComputeIsolationAttributes) HostSku() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("host_sku"))
}

type DiskEncryptionAttributes struct {
	ref terra.Reference
}

func (de DiskEncryptionAttributes) InternalRef() (terra.Reference, error) {
	return de.ref, nil
}

func (de DiskEncryptionAttributes) InternalWithRef(ref terra.Reference) DiskEncryptionAttributes {
	return DiskEncryptionAttributes{ref: ref}
}

func (de DiskEncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return de.ref.InternalTokens()
}

func (de DiskEncryptionAttributes) EncryptionAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("encryption_algorithm"))
}

func (de DiskEncryptionAttributes) EncryptionAtHostEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(de.ref.Append("encryption_at_host_enabled"))
}

func (de DiskEncryptionAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("key_vault_key_id"))
}

func (de DiskEncryptionAttributes) KeyVaultManagedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("key_vault_managed_identity_id"))
}

type ExtensionAttributes struct {
	ref terra.Reference
}

func (e ExtensionAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ExtensionAttributes) InternalWithRef(ref terra.Reference) ExtensionAttributes {
	return ExtensionAttributes{ref: ref}
}

func (e ExtensionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e ExtensionAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("log_analytics_workspace_id"))
}

func (e ExtensionAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("primary_key"))
}

type GatewayAttributes struct {
	ref terra.Reference
}

func (g GatewayAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g GatewayAttributes) InternalWithRef(ref terra.Reference) GatewayAttributes {
	return GatewayAttributes{ref: ref}
}

func (g GatewayAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return g.ref.InternalTokens()
}

func (g GatewayAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("password"))
}

func (g GatewayAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("username"))
}

type MetastoresAttributes struct {
	ref terra.Reference
}

func (m MetastoresAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MetastoresAttributes) InternalWithRef(ref terra.Reference) MetastoresAttributes {
	return MetastoresAttributes{ref: ref}
}

func (m MetastoresAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MetastoresAttributes) Ambari() terra.ListValue[AmbariAttributes] {
	return terra.ReferenceAsList[AmbariAttributes](m.ref.Append("ambari"))
}

func (m MetastoresAttributes) Hive() terra.ListValue[HiveAttributes] {
	return terra.ReferenceAsList[HiveAttributes](m.ref.Append("hive"))
}

func (m MetastoresAttributes) Oozie() terra.ListValue[OozieAttributes] {
	return terra.ReferenceAsList[OozieAttributes](m.ref.Append("oozie"))
}

type AmbariAttributes struct {
	ref terra.Reference
}

func (a AmbariAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AmbariAttributes) InternalWithRef(ref terra.Reference) AmbariAttributes {
	return AmbariAttributes{ref: ref}
}

func (a AmbariAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AmbariAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("database_name"))
}

func (a AmbariAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("password"))
}

func (a AmbariAttributes) Server() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("server"))
}

func (a AmbariAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("username"))
}

type HiveAttributes struct {
	ref terra.Reference
}

func (h HiveAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h HiveAttributes) InternalWithRef(ref terra.Reference) HiveAttributes {
	return HiveAttributes{ref: ref}
}

func (h HiveAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h HiveAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("database_name"))
}

func (h HiveAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("password"))
}

func (h HiveAttributes) Server() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("server"))
}

func (h HiveAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("username"))
}

type OozieAttributes struct {
	ref terra.Reference
}

func (o OozieAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o OozieAttributes) InternalWithRef(ref terra.Reference) OozieAttributes {
	return OozieAttributes{ref: ref}
}

func (o OozieAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o OozieAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("database_name"))
}

func (o OozieAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("password"))
}

func (o OozieAttributes) Server() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("server"))
}

func (o OozieAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("username"))
}

type MonitorAttributes struct {
	ref terra.Reference
}

func (m MonitorAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MonitorAttributes) InternalWithRef(ref terra.Reference) MonitorAttributes {
	return MonitorAttributes{ref: ref}
}

func (m MonitorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MonitorAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("log_analytics_workspace_id"))
}

func (m MonitorAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("primary_key"))
}

type NetworkAttributes struct {
	ref terra.Reference
}

func (n NetworkAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n NetworkAttributes) InternalWithRef(ref terra.Reference) NetworkAttributes {
	return NetworkAttributes{ref: ref}
}

func (n NetworkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n NetworkAttributes) ConnectionDirection() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("connection_direction"))
}

func (n NetworkAttributes) PrivateLinkEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(n.ref.Append("private_link_enabled"))
}

type RolesAttributes struct {
	ref terra.Reference
}

func (r RolesAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RolesAttributes) InternalWithRef(ref terra.Reference) RolesAttributes {
	return RolesAttributes{ref: ref}
}

func (r RolesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RolesAttributes) HeadNode() terra.ListValue[HeadNodeAttributes] {
	return terra.ReferenceAsList[HeadNodeAttributes](r.ref.Append("head_node"))
}

func (r RolesAttributes) WorkerNode() terra.ListValue[WorkerNodeAttributes] {
	return terra.ReferenceAsList[WorkerNodeAttributes](r.ref.Append("worker_node"))
}

func (r RolesAttributes) ZookeeperNode() terra.ListValue[ZookeeperNodeAttributes] {
	return terra.ReferenceAsList[ZookeeperNodeAttributes](r.ref.Append("zookeeper_node"))
}

type HeadNodeAttributes struct {
	ref terra.Reference
}

func (hn HeadNodeAttributes) InternalRef() (terra.Reference, error) {
	return hn.ref, nil
}

func (hn HeadNodeAttributes) InternalWithRef(ref terra.Reference) HeadNodeAttributes {
	return HeadNodeAttributes{ref: ref}
}

func (hn HeadNodeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hn.ref.InternalTokens()
}

func (hn HeadNodeAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(hn.ref.Append("password"))
}

func (hn HeadNodeAttributes) SshKeys() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hn.ref.Append("ssh_keys"))
}

func (hn HeadNodeAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(hn.ref.Append("subnet_id"))
}

func (hn HeadNodeAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(hn.ref.Append("username"))
}

func (hn HeadNodeAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(hn.ref.Append("virtual_network_id"))
}

func (hn HeadNodeAttributes) VmSize() terra.StringValue {
	return terra.ReferenceAsString(hn.ref.Append("vm_size"))
}

func (hn HeadNodeAttributes) ScriptActions() terra.ListValue[HeadNodeScriptActionsAttributes] {
	return terra.ReferenceAsList[HeadNodeScriptActionsAttributes](hn.ref.Append("script_actions"))
}

type HeadNodeScriptActionsAttributes struct {
	ref terra.Reference
}

func (sa HeadNodeScriptActionsAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa HeadNodeScriptActionsAttributes) InternalWithRef(ref terra.Reference) HeadNodeScriptActionsAttributes {
	return HeadNodeScriptActionsAttributes{ref: ref}
}

func (sa HeadNodeScriptActionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa HeadNodeScriptActionsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("name"))
}

func (sa HeadNodeScriptActionsAttributes) Parameters() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("parameters"))
}

func (sa HeadNodeScriptActionsAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("uri"))
}

type WorkerNodeAttributes struct {
	ref terra.Reference
}

func (wn WorkerNodeAttributes) InternalRef() (terra.Reference, error) {
	return wn.ref, nil
}

func (wn WorkerNodeAttributes) InternalWithRef(ref terra.Reference) WorkerNodeAttributes {
	return WorkerNodeAttributes{ref: ref}
}

func (wn WorkerNodeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wn.ref.InternalTokens()
}

func (wn WorkerNodeAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(wn.ref.Append("password"))
}

func (wn WorkerNodeAttributes) SshKeys() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wn.ref.Append("ssh_keys"))
}

func (wn WorkerNodeAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(wn.ref.Append("subnet_id"))
}

func (wn WorkerNodeAttributes) TargetInstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(wn.ref.Append("target_instance_count"))
}

func (wn WorkerNodeAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(wn.ref.Append("username"))
}

func (wn WorkerNodeAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(wn.ref.Append("virtual_network_id"))
}

func (wn WorkerNodeAttributes) VmSize() terra.StringValue {
	return terra.ReferenceAsString(wn.ref.Append("vm_size"))
}

func (wn WorkerNodeAttributes) Autoscale() terra.ListValue[AutoscaleAttributes] {
	return terra.ReferenceAsList[AutoscaleAttributes](wn.ref.Append("autoscale"))
}

func (wn WorkerNodeAttributes) ScriptActions() terra.ListValue[WorkerNodeScriptActionsAttributes] {
	return terra.ReferenceAsList[WorkerNodeScriptActionsAttributes](wn.ref.Append("script_actions"))
}

type AutoscaleAttributes struct {
	ref terra.Reference
}

func (a AutoscaleAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AutoscaleAttributes) InternalWithRef(ref terra.Reference) AutoscaleAttributes {
	return AutoscaleAttributes{ref: ref}
}

func (a AutoscaleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AutoscaleAttributes) Capacity() terra.ListValue[CapacityAttributes] {
	return terra.ReferenceAsList[CapacityAttributes](a.ref.Append("capacity"))
}

func (a AutoscaleAttributes) Recurrence() terra.ListValue[RecurrenceAttributes] {
	return terra.ReferenceAsList[RecurrenceAttributes](a.ref.Append("recurrence"))
}

type CapacityAttributes struct {
	ref terra.Reference
}

func (c CapacityAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CapacityAttributes) InternalWithRef(ref terra.Reference) CapacityAttributes {
	return CapacityAttributes{ref: ref}
}

func (c CapacityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CapacityAttributes) MaxInstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("max_instance_count"))
}

func (c CapacityAttributes) MinInstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("min_instance_count"))
}

type RecurrenceAttributes struct {
	ref terra.Reference
}

func (r RecurrenceAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RecurrenceAttributes) InternalWithRef(ref terra.Reference) RecurrenceAttributes {
	return RecurrenceAttributes{ref: ref}
}

func (r RecurrenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RecurrenceAttributes) Timezone() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("timezone"))
}

func (r RecurrenceAttributes) Schedule() terra.ListValue[ScheduleAttributes] {
	return terra.ReferenceAsList[ScheduleAttributes](r.ref.Append("schedule"))
}

type ScheduleAttributes struct {
	ref terra.Reference
}

func (s ScheduleAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ScheduleAttributes) InternalWithRef(ref terra.Reference) ScheduleAttributes {
	return ScheduleAttributes{ref: ref}
}

func (s ScheduleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ScheduleAttributes) Days() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("days"))
}

func (s ScheduleAttributes) TargetInstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("target_instance_count"))
}

func (s ScheduleAttributes) Time() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("time"))
}

type WorkerNodeScriptActionsAttributes struct {
	ref terra.Reference
}

func (sa WorkerNodeScriptActionsAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa WorkerNodeScriptActionsAttributes) InternalWithRef(ref terra.Reference) WorkerNodeScriptActionsAttributes {
	return WorkerNodeScriptActionsAttributes{ref: ref}
}

func (sa WorkerNodeScriptActionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa WorkerNodeScriptActionsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("name"))
}

func (sa WorkerNodeScriptActionsAttributes) Parameters() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("parameters"))
}

func (sa WorkerNodeScriptActionsAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("uri"))
}

type ZookeeperNodeAttributes struct {
	ref terra.Reference
}

func (zn ZookeeperNodeAttributes) InternalRef() (terra.Reference, error) {
	return zn.ref, nil
}

func (zn ZookeeperNodeAttributes) InternalWithRef(ref terra.Reference) ZookeeperNodeAttributes {
	return ZookeeperNodeAttributes{ref: ref}
}

func (zn ZookeeperNodeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return zn.ref.InternalTokens()
}

func (zn ZookeeperNodeAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(zn.ref.Append("password"))
}

func (zn ZookeeperNodeAttributes) SshKeys() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](zn.ref.Append("ssh_keys"))
}

func (zn ZookeeperNodeAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(zn.ref.Append("subnet_id"))
}

func (zn ZookeeperNodeAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(zn.ref.Append("username"))
}

func (zn ZookeeperNodeAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(zn.ref.Append("virtual_network_id"))
}

func (zn ZookeeperNodeAttributes) VmSize() terra.StringValue {
	return terra.ReferenceAsString(zn.ref.Append("vm_size"))
}

func (zn ZookeeperNodeAttributes) ScriptActions() terra.ListValue[ZookeeperNodeScriptActionsAttributes] {
	return terra.ReferenceAsList[ZookeeperNodeScriptActionsAttributes](zn.ref.Append("script_actions"))
}

type ZookeeperNodeScriptActionsAttributes struct {
	ref terra.Reference
}

func (sa ZookeeperNodeScriptActionsAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa ZookeeperNodeScriptActionsAttributes) InternalWithRef(ref terra.Reference) ZookeeperNodeScriptActionsAttributes {
	return ZookeeperNodeScriptActionsAttributes{ref: ref}
}

func (sa ZookeeperNodeScriptActionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa ZookeeperNodeScriptActionsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("name"))
}

func (sa ZookeeperNodeScriptActionsAttributes) Parameters() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("parameters"))
}

func (sa ZookeeperNodeScriptActionsAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("uri"))
}

type SecurityProfileAttributes struct {
	ref terra.Reference
}

func (sp SecurityProfileAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp SecurityProfileAttributes) InternalWithRef(ref terra.Reference) SecurityProfileAttributes {
	return SecurityProfileAttributes{ref: ref}
}

func (sp SecurityProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp SecurityProfileAttributes) AaddsResourceId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("aadds_resource_id"))
}

func (sp SecurityProfileAttributes) ClusterUsersGroupDns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sp.ref.Append("cluster_users_group_dns"))
}

func (sp SecurityProfileAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("domain_name"))
}

func (sp SecurityProfileAttributes) DomainUserPassword() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("domain_user_password"))
}

func (sp SecurityProfileAttributes) DomainUsername() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("domain_username"))
}

func (sp SecurityProfileAttributes) LdapsUrls() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sp.ref.Append("ldaps_urls"))
}

func (sp SecurityProfileAttributes) MsiResourceId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("msi_resource_id"))
}

type StorageAccountAttributes struct {
	ref terra.Reference
}

func (sa StorageAccountAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa StorageAccountAttributes) InternalWithRef(ref terra.Reference) StorageAccountAttributes {
	return StorageAccountAttributes{ref: ref}
}

func (sa StorageAccountAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa StorageAccountAttributes) IsDefault() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("is_default"))
}

func (sa StorageAccountAttributes) StorageAccountKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("storage_account_key"))
}

func (sa StorageAccountAttributes) StorageContainerId() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("storage_container_id"))
}

func (sa StorageAccountAttributes) StorageResourceId() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("storage_resource_id"))
}

type StorageAccountGen2Attributes struct {
	ref terra.Reference
}

func (sag StorageAccountGen2Attributes) InternalRef() (terra.Reference, error) {
	return sag.ref, nil
}

func (sag StorageAccountGen2Attributes) InternalWithRef(ref terra.Reference) StorageAccountGen2Attributes {
	return StorageAccountGen2Attributes{ref: ref}
}

func (sag StorageAccountGen2Attributes) InternalTokens() (hclwrite.Tokens, error) {
	return sag.ref.InternalTokens()
}

func (sag StorageAccountGen2Attributes) FilesystemId() terra.StringValue {
	return terra.ReferenceAsString(sag.ref.Append("filesystem_id"))
}

func (sag StorageAccountGen2Attributes) IsDefault() terra.BoolValue {
	return terra.ReferenceAsBool(sag.ref.Append("is_default"))
}

func (sag StorageAccountGen2Attributes) ManagedIdentityResourceId() terra.StringValue {
	return terra.ReferenceAsString(sag.ref.Append("managed_identity_resource_id"))
}

func (sag StorageAccountGen2Attributes) StorageResourceId() terra.StringValue {
	return terra.ReferenceAsString(sag.ref.Append("storage_resource_id"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ComponentVersionState struct {
	Spark string `json:"spark"`
}

type ComputeIsolationState struct {
	ComputeIsolationEnabled bool   `json:"compute_isolation_enabled"`
	HostSku                 string `json:"host_sku"`
}

type DiskEncryptionState struct {
	EncryptionAlgorithm       string `json:"encryption_algorithm"`
	EncryptionAtHostEnabled   bool   `json:"encryption_at_host_enabled"`
	KeyVaultKeyId             string `json:"key_vault_key_id"`
	KeyVaultManagedIdentityId string `json:"key_vault_managed_identity_id"`
}

type ExtensionState struct {
	LogAnalyticsWorkspaceId string `json:"log_analytics_workspace_id"`
	PrimaryKey              string `json:"primary_key"`
}

type GatewayState struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type MetastoresState struct {
	Ambari []AmbariState `json:"ambari"`
	Hive   []HiveState   `json:"hive"`
	Oozie  []OozieState  `json:"oozie"`
}

type AmbariState struct {
	DatabaseName string `json:"database_name"`
	Password     string `json:"password"`
	Server       string `json:"server"`
	Username     string `json:"username"`
}

type HiveState struct {
	DatabaseName string `json:"database_name"`
	Password     string `json:"password"`
	Server       string `json:"server"`
	Username     string `json:"username"`
}

type OozieState struct {
	DatabaseName string `json:"database_name"`
	Password     string `json:"password"`
	Server       string `json:"server"`
	Username     string `json:"username"`
}

type MonitorState struct {
	LogAnalyticsWorkspaceId string `json:"log_analytics_workspace_id"`
	PrimaryKey              string `json:"primary_key"`
}

type NetworkState struct {
	ConnectionDirection string `json:"connection_direction"`
	PrivateLinkEnabled  bool   `json:"private_link_enabled"`
}

type RolesState struct {
	HeadNode      []HeadNodeState      `json:"head_node"`
	WorkerNode    []WorkerNodeState    `json:"worker_node"`
	ZookeeperNode []ZookeeperNodeState `json:"zookeeper_node"`
}

type HeadNodeState struct {
	Password         string                       `json:"password"`
	SshKeys          []string                     `json:"ssh_keys"`
	SubnetId         string                       `json:"subnet_id"`
	Username         string                       `json:"username"`
	VirtualNetworkId string                       `json:"virtual_network_id"`
	VmSize           string                       `json:"vm_size"`
	ScriptActions    []HeadNodeScriptActionsState `json:"script_actions"`
}

type HeadNodeScriptActionsState struct {
	Name       string `json:"name"`
	Parameters string `json:"parameters"`
	Uri        string `json:"uri"`
}

type WorkerNodeState struct {
	Password            string                         `json:"password"`
	SshKeys             []string                       `json:"ssh_keys"`
	SubnetId            string                         `json:"subnet_id"`
	TargetInstanceCount float64                        `json:"target_instance_count"`
	Username            string                         `json:"username"`
	VirtualNetworkId    string                         `json:"virtual_network_id"`
	VmSize              string                         `json:"vm_size"`
	Autoscale           []AutoscaleState               `json:"autoscale"`
	ScriptActions       []WorkerNodeScriptActionsState `json:"script_actions"`
}

type AutoscaleState struct {
	Capacity   []CapacityState   `json:"capacity"`
	Recurrence []RecurrenceState `json:"recurrence"`
}

type CapacityState struct {
	MaxInstanceCount float64 `json:"max_instance_count"`
	MinInstanceCount float64 `json:"min_instance_count"`
}

type RecurrenceState struct {
	Timezone string          `json:"timezone"`
	Schedule []ScheduleState `json:"schedule"`
}

type ScheduleState struct {
	Days                []string `json:"days"`
	TargetInstanceCount float64  `json:"target_instance_count"`
	Time                string   `json:"time"`
}

type WorkerNodeScriptActionsState struct {
	Name       string `json:"name"`
	Parameters string `json:"parameters"`
	Uri        string `json:"uri"`
}

type ZookeeperNodeState struct {
	Password         string                            `json:"password"`
	SshKeys          []string                          `json:"ssh_keys"`
	SubnetId         string                            `json:"subnet_id"`
	Username         string                            `json:"username"`
	VirtualNetworkId string                            `json:"virtual_network_id"`
	VmSize           string                            `json:"vm_size"`
	ScriptActions    []ZookeeperNodeScriptActionsState `json:"script_actions"`
}

type ZookeeperNodeScriptActionsState struct {
	Name       string `json:"name"`
	Parameters string `json:"parameters"`
	Uri        string `json:"uri"`
}

type SecurityProfileState struct {
	AaddsResourceId      string   `json:"aadds_resource_id"`
	ClusterUsersGroupDns []string `json:"cluster_users_group_dns"`
	DomainName           string   `json:"domain_name"`
	DomainUserPassword   string   `json:"domain_user_password"`
	DomainUsername       string   `json:"domain_username"`
	LdapsUrls            []string `json:"ldaps_urls"`
	MsiResourceId        string   `json:"msi_resource_id"`
}

type StorageAccountState struct {
	IsDefault          bool   `json:"is_default"`
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
	StorageResourceId  string `json:"storage_resource_id"`
}

type StorageAccountGen2State struct {
	FilesystemId              string `json:"filesystem_id"`
	IsDefault                 bool   `json:"is_default"`
	ManagedIdentityResourceId string `json:"managed_identity_resource_id"`
	StorageResourceId         string `json:"storage_resource_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
