// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datakubernetescluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AciConnectorLinux struct{}

type AgentPoolProfile struct {
	// UpgradeSettings: min=0
	UpgradeSettings []UpgradeSettings `hcl:"upgrade_settings,block" validate:"min=0"`
}

type UpgradeSettings struct{}

type AzureActiveDirectoryRoleBasedAccessControl struct{}

type Identity struct{}

type IngressApplicationGateway struct {
	// IngressApplicationGatewayIdentity: min=0
	IngressApplicationGatewayIdentity []IngressApplicationGatewayIdentity `hcl:"ingress_application_gateway_identity,block" validate:"min=0"`
}

type IngressApplicationGatewayIdentity struct{}

type KeyManagementService struct{}

type KeyVaultSecretsProvider struct {
	// SecretIdentity: min=0
	SecretIdentity []SecretIdentity `hcl:"secret_identity,block" validate:"min=0"`
}

type SecretIdentity struct{}

type KubeAdminConfig struct{}

type KubeConfig struct{}

type KubeletIdentity struct{}

type LinuxProfile struct {
	// SshKey: min=0
	SshKey []SshKey `hcl:"ssh_key,block" validate:"min=0"`
}

type SshKey struct{}

type MicrosoftDefender struct{}

type NetworkProfile struct{}

type OmsAgent struct {
	// OmsAgentIdentity: min=0
	OmsAgentIdentity []OmsAgentIdentity `hcl:"oms_agent_identity,block" validate:"min=0"`
}

type OmsAgentIdentity struct{}

type ServicePrincipal struct{}

type StorageProfile struct{}

type WindowsProfile struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type AciConnectorLinuxAttributes struct {
	ref terra.Reference
}

func (acl AciConnectorLinuxAttributes) InternalRef() (terra.Reference, error) {
	return acl.ref, nil
}

func (acl AciConnectorLinuxAttributes) InternalWithRef(ref terra.Reference) AciConnectorLinuxAttributes {
	return AciConnectorLinuxAttributes{ref: ref}
}

func (acl AciConnectorLinuxAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return acl.ref.InternalTokens()
}

func (acl AciConnectorLinuxAttributes) SubnetName() terra.StringValue {
	return terra.ReferenceAsString(acl.ref.Append("subnet_name"))
}

type AgentPoolProfileAttributes struct {
	ref terra.Reference
}

func (app AgentPoolProfileAttributes) InternalRef() (terra.Reference, error) {
	return app.ref, nil
}

func (app AgentPoolProfileAttributes) InternalWithRef(ref terra.Reference) AgentPoolProfileAttributes {
	return AgentPoolProfileAttributes{ref: ref}
}

func (app AgentPoolProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return app.ref.InternalTokens()
}

func (app AgentPoolProfileAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(app.ref.Append("count"))
}

func (app AgentPoolProfileAttributes) EnableAutoScaling() terra.BoolValue {
	return terra.ReferenceAsBool(app.ref.Append("enable_auto_scaling"))
}

func (app AgentPoolProfileAttributes) EnableNodePublicIp() terra.BoolValue {
	return terra.ReferenceAsBool(app.ref.Append("enable_node_public_ip"))
}

func (app AgentPoolProfileAttributes) MaxCount() terra.NumberValue {
	return terra.ReferenceAsNumber(app.ref.Append("max_count"))
}

func (app AgentPoolProfileAttributes) MaxPods() terra.NumberValue {
	return terra.ReferenceAsNumber(app.ref.Append("max_pods"))
}

func (app AgentPoolProfileAttributes) MinCount() terra.NumberValue {
	return terra.ReferenceAsNumber(app.ref.Append("min_count"))
}

func (app AgentPoolProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("name"))
}

func (app AgentPoolProfileAttributes) NodeLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](app.ref.Append("node_labels"))
}

func (app AgentPoolProfileAttributes) NodePublicIpPrefixId() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("node_public_ip_prefix_id"))
}

func (app AgentPoolProfileAttributes) NodeTaints() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](app.ref.Append("node_taints"))
}

func (app AgentPoolProfileAttributes) OrchestratorVersion() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("orchestrator_version"))
}

func (app AgentPoolProfileAttributes) OsDiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(app.ref.Append("os_disk_size_gb"))
}

func (app AgentPoolProfileAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("os_type"))
}

func (app AgentPoolProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](app.ref.Append("tags"))
}

func (app AgentPoolProfileAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("type"))
}

func (app AgentPoolProfileAttributes) VmSize() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("vm_size"))
}

func (app AgentPoolProfileAttributes) VnetSubnetId() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("vnet_subnet_id"))
}

func (app AgentPoolProfileAttributes) Zones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](app.ref.Append("zones"))
}

func (app AgentPoolProfileAttributes) UpgradeSettings() terra.ListValue[UpgradeSettingsAttributes] {
	return terra.ReferenceAsList[UpgradeSettingsAttributes](app.ref.Append("upgrade_settings"))
}

type UpgradeSettingsAttributes struct {
	ref terra.Reference
}

func (us UpgradeSettingsAttributes) InternalRef() (terra.Reference, error) {
	return us.ref, nil
}

func (us UpgradeSettingsAttributes) InternalWithRef(ref terra.Reference) UpgradeSettingsAttributes {
	return UpgradeSettingsAttributes{ref: ref}
}

func (us UpgradeSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return us.ref.InternalTokens()
}

func (us UpgradeSettingsAttributes) MaxSurge() terra.StringValue {
	return terra.ReferenceAsString(us.ref.Append("max_surge"))
}

type AzureActiveDirectoryRoleBasedAccessControlAttributes struct {
	ref terra.Reference
}

func (aadrbac AzureActiveDirectoryRoleBasedAccessControlAttributes) InternalRef() (terra.Reference, error) {
	return aadrbac.ref, nil
}

func (aadrbac AzureActiveDirectoryRoleBasedAccessControlAttributes) InternalWithRef(ref terra.Reference) AzureActiveDirectoryRoleBasedAccessControlAttributes {
	return AzureActiveDirectoryRoleBasedAccessControlAttributes{ref: ref}
}

func (aadrbac AzureActiveDirectoryRoleBasedAccessControlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aadrbac.ref.InternalTokens()
}

func (aadrbac AzureActiveDirectoryRoleBasedAccessControlAttributes) AdminGroupObjectIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aadrbac.ref.Append("admin_group_object_ids"))
}

func (aadrbac AzureActiveDirectoryRoleBasedAccessControlAttributes) AzureRbacEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aadrbac.ref.Append("azure_rbac_enabled"))
}

func (aadrbac AzureActiveDirectoryRoleBasedAccessControlAttributes) ClientAppId() terra.StringValue {
	return terra.ReferenceAsString(aadrbac.ref.Append("client_app_id"))
}

func (aadrbac AzureActiveDirectoryRoleBasedAccessControlAttributes) Managed() terra.BoolValue {
	return terra.ReferenceAsBool(aadrbac.ref.Append("managed"))
}

func (aadrbac AzureActiveDirectoryRoleBasedAccessControlAttributes) ServerAppId() terra.StringValue {
	return terra.ReferenceAsString(aadrbac.ref.Append("server_app_id"))
}

func (aadrbac AzureActiveDirectoryRoleBasedAccessControlAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(aadrbac.ref.Append("tenant_id"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type IngressApplicationGatewayAttributes struct {
	ref terra.Reference
}

func (iag IngressApplicationGatewayAttributes) InternalRef() (terra.Reference, error) {
	return iag.ref, nil
}

func (iag IngressApplicationGatewayAttributes) InternalWithRef(ref terra.Reference) IngressApplicationGatewayAttributes {
	return IngressApplicationGatewayAttributes{ref: ref}
}

func (iag IngressApplicationGatewayAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return iag.ref.InternalTokens()
}

func (iag IngressApplicationGatewayAttributes) EffectiveGatewayId() terra.StringValue {
	return terra.ReferenceAsString(iag.ref.Append("effective_gateway_id"))
}

func (iag IngressApplicationGatewayAttributes) GatewayId() terra.StringValue {
	return terra.ReferenceAsString(iag.ref.Append("gateway_id"))
}

func (iag IngressApplicationGatewayAttributes) GatewayName() terra.StringValue {
	return terra.ReferenceAsString(iag.ref.Append("gateway_name"))
}

func (iag IngressApplicationGatewayAttributes) SubnetCidr() terra.StringValue {
	return terra.ReferenceAsString(iag.ref.Append("subnet_cidr"))
}

func (iag IngressApplicationGatewayAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(iag.ref.Append("subnet_id"))
}

func (iag IngressApplicationGatewayAttributes) IngressApplicationGatewayIdentity() terra.ListValue[IngressApplicationGatewayIdentityAttributes] {
	return terra.ReferenceAsList[IngressApplicationGatewayIdentityAttributes](iag.ref.Append("ingress_application_gateway_identity"))
}

type IngressApplicationGatewayIdentityAttributes struct {
	ref terra.Reference
}

func (iagi IngressApplicationGatewayIdentityAttributes) InternalRef() (terra.Reference, error) {
	return iagi.ref, nil
}

func (iagi IngressApplicationGatewayIdentityAttributes) InternalWithRef(ref terra.Reference) IngressApplicationGatewayIdentityAttributes {
	return IngressApplicationGatewayIdentityAttributes{ref: ref}
}

func (iagi IngressApplicationGatewayIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return iagi.ref.InternalTokens()
}

func (iagi IngressApplicationGatewayIdentityAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(iagi.ref.Append("client_id"))
}

func (iagi IngressApplicationGatewayIdentityAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(iagi.ref.Append("object_id"))
}

func (iagi IngressApplicationGatewayIdentityAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(iagi.ref.Append("user_assigned_identity_id"))
}

type KeyManagementServiceAttributes struct {
	ref terra.Reference
}

func (kms KeyManagementServiceAttributes) InternalRef() (terra.Reference, error) {
	return kms.ref, nil
}

func (kms KeyManagementServiceAttributes) InternalWithRef(ref terra.Reference) KeyManagementServiceAttributes {
	return KeyManagementServiceAttributes{ref: ref}
}

func (kms KeyManagementServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kms.ref.InternalTokens()
}

func (kms KeyManagementServiceAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(kms.ref.Append("key_vault_key_id"))
}

func (kms KeyManagementServiceAttributes) KeyVaultNetworkAccess() terra.StringValue {
	return terra.ReferenceAsString(kms.ref.Append("key_vault_network_access"))
}

type KeyVaultSecretsProviderAttributes struct {
	ref terra.Reference
}

func (kvsp KeyVaultSecretsProviderAttributes) InternalRef() (terra.Reference, error) {
	return kvsp.ref, nil
}

func (kvsp KeyVaultSecretsProviderAttributes) InternalWithRef(ref terra.Reference) KeyVaultSecretsProviderAttributes {
	return KeyVaultSecretsProviderAttributes{ref: ref}
}

func (kvsp KeyVaultSecretsProviderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kvsp.ref.InternalTokens()
}

func (kvsp KeyVaultSecretsProviderAttributes) SecretRotationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kvsp.ref.Append("secret_rotation_enabled"))
}

func (kvsp KeyVaultSecretsProviderAttributes) SecretRotationInterval() terra.StringValue {
	return terra.ReferenceAsString(kvsp.ref.Append("secret_rotation_interval"))
}

func (kvsp KeyVaultSecretsProviderAttributes) SecretIdentity() terra.ListValue[SecretIdentityAttributes] {
	return terra.ReferenceAsList[SecretIdentityAttributes](kvsp.ref.Append("secret_identity"))
}

type SecretIdentityAttributes struct {
	ref terra.Reference
}

func (si SecretIdentityAttributes) InternalRef() (terra.Reference, error) {
	return si.ref, nil
}

func (si SecretIdentityAttributes) InternalWithRef(ref terra.Reference) SecretIdentityAttributes {
	return SecretIdentityAttributes{ref: ref}
}

func (si SecretIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return si.ref.InternalTokens()
}

func (si SecretIdentityAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("client_id"))
}

func (si SecretIdentityAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("object_id"))
}

func (si SecretIdentityAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("user_assigned_identity_id"))
}

type KubeAdminConfigAttributes struct {
	ref terra.Reference
}

func (kac KubeAdminConfigAttributes) InternalRef() (terra.Reference, error) {
	return kac.ref, nil
}

func (kac KubeAdminConfigAttributes) InternalWithRef(ref terra.Reference) KubeAdminConfigAttributes {
	return KubeAdminConfigAttributes{ref: ref}
}

func (kac KubeAdminConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kac.ref.InternalTokens()
}

func (kac KubeAdminConfigAttributes) ClientCertificate() terra.StringValue {
	return terra.ReferenceAsString(kac.ref.Append("client_certificate"))
}

func (kac KubeAdminConfigAttributes) ClientKey() terra.StringValue {
	return terra.ReferenceAsString(kac.ref.Append("client_key"))
}

func (kac KubeAdminConfigAttributes) ClusterCaCertificate() terra.StringValue {
	return terra.ReferenceAsString(kac.ref.Append("cluster_ca_certificate"))
}

func (kac KubeAdminConfigAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(kac.ref.Append("host"))
}

func (kac KubeAdminConfigAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(kac.ref.Append("password"))
}

func (kac KubeAdminConfigAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(kac.ref.Append("username"))
}

type KubeConfigAttributes struct {
	ref terra.Reference
}

func (kc KubeConfigAttributes) InternalRef() (terra.Reference, error) {
	return kc.ref, nil
}

func (kc KubeConfigAttributes) InternalWithRef(ref terra.Reference) KubeConfigAttributes {
	return KubeConfigAttributes{ref: ref}
}

func (kc KubeConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kc.ref.InternalTokens()
}

func (kc KubeConfigAttributes) ClientCertificate() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("client_certificate"))
}

func (kc KubeConfigAttributes) ClientKey() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("client_key"))
}

func (kc KubeConfigAttributes) ClusterCaCertificate() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("cluster_ca_certificate"))
}

func (kc KubeConfigAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("host"))
}

func (kc KubeConfigAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("password"))
}

func (kc KubeConfigAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("username"))
}

type KubeletIdentityAttributes struct {
	ref terra.Reference
}

func (ki KubeletIdentityAttributes) InternalRef() (terra.Reference, error) {
	return ki.ref, nil
}

func (ki KubeletIdentityAttributes) InternalWithRef(ref terra.Reference) KubeletIdentityAttributes {
	return KubeletIdentityAttributes{ref: ref}
}

func (ki KubeletIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ki.ref.InternalTokens()
}

func (ki KubeletIdentityAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("client_id"))
}

func (ki KubeletIdentityAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("object_id"))
}

func (ki KubeletIdentityAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("user_assigned_identity_id"))
}

type LinuxProfileAttributes struct {
	ref terra.Reference
}

func (lp LinuxProfileAttributes) InternalRef() (terra.Reference, error) {
	return lp.ref, nil
}

func (lp LinuxProfileAttributes) InternalWithRef(ref terra.Reference) LinuxProfileAttributes {
	return LinuxProfileAttributes{ref: ref}
}

func (lp LinuxProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lp.ref.InternalTokens()
}

func (lp LinuxProfileAttributes) AdminUsername() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("admin_username"))
}

func (lp LinuxProfileAttributes) SshKey() terra.ListValue[SshKeyAttributes] {
	return terra.ReferenceAsList[SshKeyAttributes](lp.ref.Append("ssh_key"))
}

type SshKeyAttributes struct {
	ref terra.Reference
}

func (sk SshKeyAttributes) InternalRef() (terra.Reference, error) {
	return sk.ref, nil
}

func (sk SshKeyAttributes) InternalWithRef(ref terra.Reference) SshKeyAttributes {
	return SshKeyAttributes{ref: ref}
}

func (sk SshKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sk.ref.InternalTokens()
}

func (sk SshKeyAttributes) KeyData() terra.StringValue {
	return terra.ReferenceAsString(sk.ref.Append("key_data"))
}

type MicrosoftDefenderAttributes struct {
	ref terra.Reference
}

func (md MicrosoftDefenderAttributes) InternalRef() (terra.Reference, error) {
	return md.ref, nil
}

func (md MicrosoftDefenderAttributes) InternalWithRef(ref terra.Reference) MicrosoftDefenderAttributes {
	return MicrosoftDefenderAttributes{ref: ref}
}

func (md MicrosoftDefenderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return md.ref.InternalTokens()
}

func (md MicrosoftDefenderAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("log_analytics_workspace_id"))
}

type NetworkProfileAttributes struct {
	ref terra.Reference
}

func (np NetworkProfileAttributes) InternalRef() (terra.Reference, error) {
	return np.ref, nil
}

func (np NetworkProfileAttributes) InternalWithRef(ref terra.Reference) NetworkProfileAttributes {
	return NetworkProfileAttributes{ref: ref}
}

func (np NetworkProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return np.ref.InternalTokens()
}

func (np NetworkProfileAttributes) DnsServiceIp() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("dns_service_ip"))
}

func (np NetworkProfileAttributes) DockerBridgeCidr() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("docker_bridge_cidr"))
}

func (np NetworkProfileAttributes) LoadBalancerSku() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("load_balancer_sku"))
}

func (np NetworkProfileAttributes) NetworkPlugin() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("network_plugin"))
}

func (np NetworkProfileAttributes) NetworkPolicy() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("network_policy"))
}

func (np NetworkProfileAttributes) PodCidr() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("pod_cidr"))
}

func (np NetworkProfileAttributes) ServiceCidr() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("service_cidr"))
}

type OmsAgentAttributes struct {
	ref terra.Reference
}

func (oa OmsAgentAttributes) InternalRef() (terra.Reference, error) {
	return oa.ref, nil
}

func (oa OmsAgentAttributes) InternalWithRef(ref terra.Reference) OmsAgentAttributes {
	return OmsAgentAttributes{ref: ref}
}

func (oa OmsAgentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oa.ref.InternalTokens()
}

func (oa OmsAgentAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("log_analytics_workspace_id"))
}

func (oa OmsAgentAttributes) MsiAuthForMonitoringEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(oa.ref.Append("msi_auth_for_monitoring_enabled"))
}

func (oa OmsAgentAttributes) OmsAgentIdentity() terra.ListValue[OmsAgentIdentityAttributes] {
	return terra.ReferenceAsList[OmsAgentIdentityAttributes](oa.ref.Append("oms_agent_identity"))
}

type OmsAgentIdentityAttributes struct {
	ref terra.Reference
}

func (oai OmsAgentIdentityAttributes) InternalRef() (terra.Reference, error) {
	return oai.ref, nil
}

func (oai OmsAgentIdentityAttributes) InternalWithRef(ref terra.Reference) OmsAgentIdentityAttributes {
	return OmsAgentIdentityAttributes{ref: ref}
}

func (oai OmsAgentIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oai.ref.InternalTokens()
}

func (oai OmsAgentIdentityAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(oai.ref.Append("client_id"))
}

func (oai OmsAgentIdentityAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(oai.ref.Append("object_id"))
}

func (oai OmsAgentIdentityAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(oai.ref.Append("user_assigned_identity_id"))
}

type ServicePrincipalAttributes struct {
	ref terra.Reference
}

func (sp ServicePrincipalAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp ServicePrincipalAttributes) InternalWithRef(ref terra.Reference) ServicePrincipalAttributes {
	return ServicePrincipalAttributes{ref: ref}
}

func (sp ServicePrincipalAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp ServicePrincipalAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("client_id"))
}

type StorageProfileAttributes struct {
	ref terra.Reference
}

func (sp StorageProfileAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp StorageProfileAttributes) InternalWithRef(ref terra.Reference) StorageProfileAttributes {
	return StorageProfileAttributes{ref: ref}
}

func (sp StorageProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp StorageProfileAttributes) BlobDriverEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("blob_driver_enabled"))
}

func (sp StorageProfileAttributes) DiskDriverEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("disk_driver_enabled"))
}

func (sp StorageProfileAttributes) DiskDriverVersion() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("disk_driver_version"))
}

func (sp StorageProfileAttributes) FileDriverEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("file_driver_enabled"))
}

func (sp StorageProfileAttributes) SnapshotControllerEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("snapshot_controller_enabled"))
}

type WindowsProfileAttributes struct {
	ref terra.Reference
}

func (wp WindowsProfileAttributes) InternalRef() (terra.Reference, error) {
	return wp.ref, nil
}

func (wp WindowsProfileAttributes) InternalWithRef(ref terra.Reference) WindowsProfileAttributes {
	return WindowsProfileAttributes{ref: ref}
}

func (wp WindowsProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wp.ref.InternalTokens()
}

func (wp WindowsProfileAttributes) AdminUsername() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("admin_username"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type AciConnectorLinuxState struct {
	SubnetName string `json:"subnet_name"`
}

type AgentPoolProfileState struct {
	Count                float64                `json:"count"`
	EnableAutoScaling    bool                   `json:"enable_auto_scaling"`
	EnableNodePublicIp   bool                   `json:"enable_node_public_ip"`
	MaxCount             float64                `json:"max_count"`
	MaxPods              float64                `json:"max_pods"`
	MinCount             float64                `json:"min_count"`
	Name                 string                 `json:"name"`
	NodeLabels           map[string]string      `json:"node_labels"`
	NodePublicIpPrefixId string                 `json:"node_public_ip_prefix_id"`
	NodeTaints           []string               `json:"node_taints"`
	OrchestratorVersion  string                 `json:"orchestrator_version"`
	OsDiskSizeGb         float64                `json:"os_disk_size_gb"`
	OsType               string                 `json:"os_type"`
	Tags                 map[string]string      `json:"tags"`
	Type                 string                 `json:"type"`
	VmSize               string                 `json:"vm_size"`
	VnetSubnetId         string                 `json:"vnet_subnet_id"`
	Zones                []string               `json:"zones"`
	UpgradeSettings      []UpgradeSettingsState `json:"upgrade_settings"`
}

type UpgradeSettingsState struct {
	MaxSurge string `json:"max_surge"`
}

type AzureActiveDirectoryRoleBasedAccessControlState struct {
	AdminGroupObjectIds []string `json:"admin_group_object_ids"`
	AzureRbacEnabled    bool     `json:"azure_rbac_enabled"`
	ClientAppId         string   `json:"client_app_id"`
	Managed             bool     `json:"managed"`
	ServerAppId         string   `json:"server_app_id"`
	TenantId            string   `json:"tenant_id"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type IngressApplicationGatewayState struct {
	EffectiveGatewayId                string                                   `json:"effective_gateway_id"`
	GatewayId                         string                                   `json:"gateway_id"`
	GatewayName                       string                                   `json:"gateway_name"`
	SubnetCidr                        string                                   `json:"subnet_cidr"`
	SubnetId                          string                                   `json:"subnet_id"`
	IngressApplicationGatewayIdentity []IngressApplicationGatewayIdentityState `json:"ingress_application_gateway_identity"`
}

type IngressApplicationGatewayIdentityState struct {
	ClientId               string `json:"client_id"`
	ObjectId               string `json:"object_id"`
	UserAssignedIdentityId string `json:"user_assigned_identity_id"`
}

type KeyManagementServiceState struct {
	KeyVaultKeyId         string `json:"key_vault_key_id"`
	KeyVaultNetworkAccess string `json:"key_vault_network_access"`
}

type KeyVaultSecretsProviderState struct {
	SecretRotationEnabled  bool                  `json:"secret_rotation_enabled"`
	SecretRotationInterval string                `json:"secret_rotation_interval"`
	SecretIdentity         []SecretIdentityState `json:"secret_identity"`
}

type SecretIdentityState struct {
	ClientId               string `json:"client_id"`
	ObjectId               string `json:"object_id"`
	UserAssignedIdentityId string `json:"user_assigned_identity_id"`
}

type KubeAdminConfigState struct {
	ClientCertificate    string `json:"client_certificate"`
	ClientKey            string `json:"client_key"`
	ClusterCaCertificate string `json:"cluster_ca_certificate"`
	Host                 string `json:"host"`
	Password             string `json:"password"`
	Username             string `json:"username"`
}

type KubeConfigState struct {
	ClientCertificate    string `json:"client_certificate"`
	ClientKey            string `json:"client_key"`
	ClusterCaCertificate string `json:"cluster_ca_certificate"`
	Host                 string `json:"host"`
	Password             string `json:"password"`
	Username             string `json:"username"`
}

type KubeletIdentityState struct {
	ClientId               string `json:"client_id"`
	ObjectId               string `json:"object_id"`
	UserAssignedIdentityId string `json:"user_assigned_identity_id"`
}

type LinuxProfileState struct {
	AdminUsername string        `json:"admin_username"`
	SshKey        []SshKeyState `json:"ssh_key"`
}

type SshKeyState struct {
	KeyData string `json:"key_data"`
}

type MicrosoftDefenderState struct {
	LogAnalyticsWorkspaceId string `json:"log_analytics_workspace_id"`
}

type NetworkProfileState struct {
	DnsServiceIp     string `json:"dns_service_ip"`
	DockerBridgeCidr string `json:"docker_bridge_cidr"`
	LoadBalancerSku  string `json:"load_balancer_sku"`
	NetworkPlugin    string `json:"network_plugin"`
	NetworkPolicy    string `json:"network_policy"`
	PodCidr          string `json:"pod_cidr"`
	ServiceCidr      string `json:"service_cidr"`
}

type OmsAgentState struct {
	LogAnalyticsWorkspaceId     string                  `json:"log_analytics_workspace_id"`
	MsiAuthForMonitoringEnabled bool                    `json:"msi_auth_for_monitoring_enabled"`
	OmsAgentIdentity            []OmsAgentIdentityState `json:"oms_agent_identity"`
}

type OmsAgentIdentityState struct {
	ClientId               string `json:"client_id"`
	ObjectId               string `json:"object_id"`
	UserAssignedIdentityId string `json:"user_assigned_identity_id"`
}

type ServicePrincipalState struct {
	ClientId string `json:"client_id"`
}

type StorageProfileState struct {
	BlobDriverEnabled         bool   `json:"blob_driver_enabled"`
	DiskDriverEnabled         bool   `json:"disk_driver_enabled"`
	DiskDriverVersion         string `json:"disk_driver_version"`
	FileDriverEnabled         bool   `json:"file_driver_enabled"`
	SnapshotControllerEnabled bool   `json:"snapshot_controller_enabled"`
}

type WindowsProfileState struct {
	AdminUsername string `json:"admin_username"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
