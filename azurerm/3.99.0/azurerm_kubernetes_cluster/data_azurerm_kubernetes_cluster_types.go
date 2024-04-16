// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_kubernetes_cluster

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataAciConnectorLinuxAttributes struct {
	ref terra.Reference
}

func (acl DataAciConnectorLinuxAttributes) InternalRef() (terra.Reference, error) {
	return acl.ref, nil
}

func (acl DataAciConnectorLinuxAttributes) InternalWithRef(ref terra.Reference) DataAciConnectorLinuxAttributes {
	return DataAciConnectorLinuxAttributes{ref: ref}
}

func (acl DataAciConnectorLinuxAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return acl.ref.InternalTokens()
}

func (acl DataAciConnectorLinuxAttributes) SubnetName() terra.StringValue {
	return terra.ReferenceAsString(acl.ref.Append("subnet_name"))
}

type DataAgentPoolProfileAttributes struct {
	ref terra.Reference
}

func (app DataAgentPoolProfileAttributes) InternalRef() (terra.Reference, error) {
	return app.ref, nil
}

func (app DataAgentPoolProfileAttributes) InternalWithRef(ref terra.Reference) DataAgentPoolProfileAttributes {
	return DataAgentPoolProfileAttributes{ref: ref}
}

func (app DataAgentPoolProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return app.ref.InternalTokens()
}

func (app DataAgentPoolProfileAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(app.ref.Append("count"))
}

func (app DataAgentPoolProfileAttributes) EnableAutoScaling() terra.BoolValue {
	return terra.ReferenceAsBool(app.ref.Append("enable_auto_scaling"))
}

func (app DataAgentPoolProfileAttributes) EnableNodePublicIp() terra.BoolValue {
	return terra.ReferenceAsBool(app.ref.Append("enable_node_public_ip"))
}

func (app DataAgentPoolProfileAttributes) MaxCount() terra.NumberValue {
	return terra.ReferenceAsNumber(app.ref.Append("max_count"))
}

func (app DataAgentPoolProfileAttributes) MaxPods() terra.NumberValue {
	return terra.ReferenceAsNumber(app.ref.Append("max_pods"))
}

func (app DataAgentPoolProfileAttributes) MinCount() terra.NumberValue {
	return terra.ReferenceAsNumber(app.ref.Append("min_count"))
}

func (app DataAgentPoolProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("name"))
}

func (app DataAgentPoolProfileAttributes) NodeLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](app.ref.Append("node_labels"))
}

func (app DataAgentPoolProfileAttributes) NodePublicIpPrefixId() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("node_public_ip_prefix_id"))
}

func (app DataAgentPoolProfileAttributes) NodeTaints() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](app.ref.Append("node_taints"))
}

func (app DataAgentPoolProfileAttributes) OrchestratorVersion() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("orchestrator_version"))
}

func (app DataAgentPoolProfileAttributes) OsDiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(app.ref.Append("os_disk_size_gb"))
}

func (app DataAgentPoolProfileAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("os_type"))
}

func (app DataAgentPoolProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](app.ref.Append("tags"))
}

func (app DataAgentPoolProfileAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("type"))
}

func (app DataAgentPoolProfileAttributes) VmSize() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("vm_size"))
}

func (app DataAgentPoolProfileAttributes) VnetSubnetId() terra.StringValue {
	return terra.ReferenceAsString(app.ref.Append("vnet_subnet_id"))
}

func (app DataAgentPoolProfileAttributes) Zones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](app.ref.Append("zones"))
}

func (app DataAgentPoolProfileAttributes) UpgradeSettings() terra.ListValue[DataAgentPoolProfileUpgradeSettingsAttributes] {
	return terra.ReferenceAsList[DataAgentPoolProfileUpgradeSettingsAttributes](app.ref.Append("upgrade_settings"))
}

type DataAgentPoolProfileUpgradeSettingsAttributes struct {
	ref terra.Reference
}

func (us DataAgentPoolProfileUpgradeSettingsAttributes) InternalRef() (terra.Reference, error) {
	return us.ref, nil
}

func (us DataAgentPoolProfileUpgradeSettingsAttributes) InternalWithRef(ref terra.Reference) DataAgentPoolProfileUpgradeSettingsAttributes {
	return DataAgentPoolProfileUpgradeSettingsAttributes{ref: ref}
}

func (us DataAgentPoolProfileUpgradeSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return us.ref.InternalTokens()
}

func (us DataAgentPoolProfileUpgradeSettingsAttributes) MaxSurge() terra.StringValue {
	return terra.ReferenceAsString(us.ref.Append("max_surge"))
}

type DataAzureActiveDirectoryRoleBasedAccessControlAttributes struct {
	ref terra.Reference
}

func (aadrbac DataAzureActiveDirectoryRoleBasedAccessControlAttributes) InternalRef() (terra.Reference, error) {
	return aadrbac.ref, nil
}

func (aadrbac DataAzureActiveDirectoryRoleBasedAccessControlAttributes) InternalWithRef(ref terra.Reference) DataAzureActiveDirectoryRoleBasedAccessControlAttributes {
	return DataAzureActiveDirectoryRoleBasedAccessControlAttributes{ref: ref}
}

func (aadrbac DataAzureActiveDirectoryRoleBasedAccessControlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aadrbac.ref.InternalTokens()
}

func (aadrbac DataAzureActiveDirectoryRoleBasedAccessControlAttributes) AdminGroupObjectIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aadrbac.ref.Append("admin_group_object_ids"))
}

func (aadrbac DataAzureActiveDirectoryRoleBasedAccessControlAttributes) AzureRbacEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aadrbac.ref.Append("azure_rbac_enabled"))
}

func (aadrbac DataAzureActiveDirectoryRoleBasedAccessControlAttributes) ClientAppId() terra.StringValue {
	return terra.ReferenceAsString(aadrbac.ref.Append("client_app_id"))
}

func (aadrbac DataAzureActiveDirectoryRoleBasedAccessControlAttributes) Managed() terra.BoolValue {
	return terra.ReferenceAsBool(aadrbac.ref.Append("managed"))
}

func (aadrbac DataAzureActiveDirectoryRoleBasedAccessControlAttributes) ServerAppId() terra.StringValue {
	return terra.ReferenceAsString(aadrbac.ref.Append("server_app_id"))
}

func (aadrbac DataAzureActiveDirectoryRoleBasedAccessControlAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(aadrbac.ref.Append("tenant_id"))
}

type DataIdentityAttributes struct {
	ref terra.Reference
}

func (i DataIdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i DataIdentityAttributes) InternalWithRef(ref terra.Reference) DataIdentityAttributes {
	return DataIdentityAttributes{ref: ref}
}

func (i DataIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i DataIdentityAttributes) IdentityIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i DataIdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i DataIdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i DataIdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type DataIngressApplicationGatewayAttributes struct {
	ref terra.Reference
}

func (iag DataIngressApplicationGatewayAttributes) InternalRef() (terra.Reference, error) {
	return iag.ref, nil
}

func (iag DataIngressApplicationGatewayAttributes) InternalWithRef(ref terra.Reference) DataIngressApplicationGatewayAttributes {
	return DataIngressApplicationGatewayAttributes{ref: ref}
}

func (iag DataIngressApplicationGatewayAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return iag.ref.InternalTokens()
}

func (iag DataIngressApplicationGatewayAttributes) EffectiveGatewayId() terra.StringValue {
	return terra.ReferenceAsString(iag.ref.Append("effective_gateway_id"))
}

func (iag DataIngressApplicationGatewayAttributes) GatewayId() terra.StringValue {
	return terra.ReferenceAsString(iag.ref.Append("gateway_id"))
}

func (iag DataIngressApplicationGatewayAttributes) GatewayName() terra.StringValue {
	return terra.ReferenceAsString(iag.ref.Append("gateway_name"))
}

func (iag DataIngressApplicationGatewayAttributes) SubnetCidr() terra.StringValue {
	return terra.ReferenceAsString(iag.ref.Append("subnet_cidr"))
}

func (iag DataIngressApplicationGatewayAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(iag.ref.Append("subnet_id"))
}

func (iag DataIngressApplicationGatewayAttributes) IngressApplicationGatewayIdentity() terra.ListValue[DataIngressApplicationGatewayIngressApplicationGatewayIdentityAttributes] {
	return terra.ReferenceAsList[DataIngressApplicationGatewayIngressApplicationGatewayIdentityAttributes](iag.ref.Append("ingress_application_gateway_identity"))
}

type DataIngressApplicationGatewayIngressApplicationGatewayIdentityAttributes struct {
	ref terra.Reference
}

func (iagi DataIngressApplicationGatewayIngressApplicationGatewayIdentityAttributes) InternalRef() (terra.Reference, error) {
	return iagi.ref, nil
}

func (iagi DataIngressApplicationGatewayIngressApplicationGatewayIdentityAttributes) InternalWithRef(ref terra.Reference) DataIngressApplicationGatewayIngressApplicationGatewayIdentityAttributes {
	return DataIngressApplicationGatewayIngressApplicationGatewayIdentityAttributes{ref: ref}
}

func (iagi DataIngressApplicationGatewayIngressApplicationGatewayIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return iagi.ref.InternalTokens()
}

func (iagi DataIngressApplicationGatewayIngressApplicationGatewayIdentityAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(iagi.ref.Append("client_id"))
}

func (iagi DataIngressApplicationGatewayIngressApplicationGatewayIdentityAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(iagi.ref.Append("object_id"))
}

func (iagi DataIngressApplicationGatewayIngressApplicationGatewayIdentityAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(iagi.ref.Append("user_assigned_identity_id"))
}

type DataKeyManagementServiceAttributes struct {
	ref terra.Reference
}

func (kms DataKeyManagementServiceAttributes) InternalRef() (terra.Reference, error) {
	return kms.ref, nil
}

func (kms DataKeyManagementServiceAttributes) InternalWithRef(ref terra.Reference) DataKeyManagementServiceAttributes {
	return DataKeyManagementServiceAttributes{ref: ref}
}

func (kms DataKeyManagementServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kms.ref.InternalTokens()
}

func (kms DataKeyManagementServiceAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(kms.ref.Append("key_vault_key_id"))
}

func (kms DataKeyManagementServiceAttributes) KeyVaultNetworkAccess() terra.StringValue {
	return terra.ReferenceAsString(kms.ref.Append("key_vault_network_access"))
}

type DataKeyVaultSecretsProviderAttributes struct {
	ref terra.Reference
}

func (kvsp DataKeyVaultSecretsProviderAttributes) InternalRef() (terra.Reference, error) {
	return kvsp.ref, nil
}

func (kvsp DataKeyVaultSecretsProviderAttributes) InternalWithRef(ref terra.Reference) DataKeyVaultSecretsProviderAttributes {
	return DataKeyVaultSecretsProviderAttributes{ref: ref}
}

func (kvsp DataKeyVaultSecretsProviderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kvsp.ref.InternalTokens()
}

func (kvsp DataKeyVaultSecretsProviderAttributes) SecretRotationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kvsp.ref.Append("secret_rotation_enabled"))
}

func (kvsp DataKeyVaultSecretsProviderAttributes) SecretRotationInterval() terra.StringValue {
	return terra.ReferenceAsString(kvsp.ref.Append("secret_rotation_interval"))
}

func (kvsp DataKeyVaultSecretsProviderAttributes) SecretIdentity() terra.ListValue[DataKeyVaultSecretsProviderSecretIdentityAttributes] {
	return terra.ReferenceAsList[DataKeyVaultSecretsProviderSecretIdentityAttributes](kvsp.ref.Append("secret_identity"))
}

type DataKeyVaultSecretsProviderSecretIdentityAttributes struct {
	ref terra.Reference
}

func (si DataKeyVaultSecretsProviderSecretIdentityAttributes) InternalRef() (terra.Reference, error) {
	return si.ref, nil
}

func (si DataKeyVaultSecretsProviderSecretIdentityAttributes) InternalWithRef(ref terra.Reference) DataKeyVaultSecretsProviderSecretIdentityAttributes {
	return DataKeyVaultSecretsProviderSecretIdentityAttributes{ref: ref}
}

func (si DataKeyVaultSecretsProviderSecretIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return si.ref.InternalTokens()
}

func (si DataKeyVaultSecretsProviderSecretIdentityAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("client_id"))
}

func (si DataKeyVaultSecretsProviderSecretIdentityAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("object_id"))
}

func (si DataKeyVaultSecretsProviderSecretIdentityAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("user_assigned_identity_id"))
}

type DataKubeAdminConfigAttributes struct {
	ref terra.Reference
}

func (kac DataKubeAdminConfigAttributes) InternalRef() (terra.Reference, error) {
	return kac.ref, nil
}

func (kac DataKubeAdminConfigAttributes) InternalWithRef(ref terra.Reference) DataKubeAdminConfigAttributes {
	return DataKubeAdminConfigAttributes{ref: ref}
}

func (kac DataKubeAdminConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kac.ref.InternalTokens()
}

func (kac DataKubeAdminConfigAttributes) ClientCertificate() terra.StringValue {
	return terra.ReferenceAsString(kac.ref.Append("client_certificate"))
}

func (kac DataKubeAdminConfigAttributes) ClientKey() terra.StringValue {
	return terra.ReferenceAsString(kac.ref.Append("client_key"))
}

func (kac DataKubeAdminConfigAttributes) ClusterCaCertificate() terra.StringValue {
	return terra.ReferenceAsString(kac.ref.Append("cluster_ca_certificate"))
}

func (kac DataKubeAdminConfigAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(kac.ref.Append("host"))
}

func (kac DataKubeAdminConfigAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(kac.ref.Append("password"))
}

func (kac DataKubeAdminConfigAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(kac.ref.Append("username"))
}

type DataKubeConfigAttributes struct {
	ref terra.Reference
}

func (kc DataKubeConfigAttributes) InternalRef() (terra.Reference, error) {
	return kc.ref, nil
}

func (kc DataKubeConfigAttributes) InternalWithRef(ref terra.Reference) DataKubeConfigAttributes {
	return DataKubeConfigAttributes{ref: ref}
}

func (kc DataKubeConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kc.ref.InternalTokens()
}

func (kc DataKubeConfigAttributes) ClientCertificate() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("client_certificate"))
}

func (kc DataKubeConfigAttributes) ClientKey() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("client_key"))
}

func (kc DataKubeConfigAttributes) ClusterCaCertificate() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("cluster_ca_certificate"))
}

func (kc DataKubeConfigAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("host"))
}

func (kc DataKubeConfigAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("password"))
}

func (kc DataKubeConfigAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("username"))
}

type DataKubeletIdentityAttributes struct {
	ref terra.Reference
}

func (ki DataKubeletIdentityAttributes) InternalRef() (terra.Reference, error) {
	return ki.ref, nil
}

func (ki DataKubeletIdentityAttributes) InternalWithRef(ref terra.Reference) DataKubeletIdentityAttributes {
	return DataKubeletIdentityAttributes{ref: ref}
}

func (ki DataKubeletIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ki.ref.InternalTokens()
}

func (ki DataKubeletIdentityAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("client_id"))
}

func (ki DataKubeletIdentityAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("object_id"))
}

func (ki DataKubeletIdentityAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("user_assigned_identity_id"))
}

type DataLinuxProfileAttributes struct {
	ref terra.Reference
}

func (lp DataLinuxProfileAttributes) InternalRef() (terra.Reference, error) {
	return lp.ref, nil
}

func (lp DataLinuxProfileAttributes) InternalWithRef(ref terra.Reference) DataLinuxProfileAttributes {
	return DataLinuxProfileAttributes{ref: ref}
}

func (lp DataLinuxProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lp.ref.InternalTokens()
}

func (lp DataLinuxProfileAttributes) AdminUsername() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("admin_username"))
}

func (lp DataLinuxProfileAttributes) SshKey() terra.ListValue[DataLinuxProfileSshKeyAttributes] {
	return terra.ReferenceAsList[DataLinuxProfileSshKeyAttributes](lp.ref.Append("ssh_key"))
}

type DataLinuxProfileSshKeyAttributes struct {
	ref terra.Reference
}

func (sk DataLinuxProfileSshKeyAttributes) InternalRef() (terra.Reference, error) {
	return sk.ref, nil
}

func (sk DataLinuxProfileSshKeyAttributes) InternalWithRef(ref terra.Reference) DataLinuxProfileSshKeyAttributes {
	return DataLinuxProfileSshKeyAttributes{ref: ref}
}

func (sk DataLinuxProfileSshKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sk.ref.InternalTokens()
}

func (sk DataLinuxProfileSshKeyAttributes) KeyData() terra.StringValue {
	return terra.ReferenceAsString(sk.ref.Append("key_data"))
}

type DataMicrosoftDefenderAttributes struct {
	ref terra.Reference
}

func (md DataMicrosoftDefenderAttributes) InternalRef() (terra.Reference, error) {
	return md.ref, nil
}

func (md DataMicrosoftDefenderAttributes) InternalWithRef(ref terra.Reference) DataMicrosoftDefenderAttributes {
	return DataMicrosoftDefenderAttributes{ref: ref}
}

func (md DataMicrosoftDefenderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return md.ref.InternalTokens()
}

func (md DataMicrosoftDefenderAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("log_analytics_workspace_id"))
}

type DataNetworkProfileAttributes struct {
	ref terra.Reference
}

func (np DataNetworkProfileAttributes) InternalRef() (terra.Reference, error) {
	return np.ref, nil
}

func (np DataNetworkProfileAttributes) InternalWithRef(ref terra.Reference) DataNetworkProfileAttributes {
	return DataNetworkProfileAttributes{ref: ref}
}

func (np DataNetworkProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return np.ref.InternalTokens()
}

func (np DataNetworkProfileAttributes) DnsServiceIp() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("dns_service_ip"))
}

func (np DataNetworkProfileAttributes) DockerBridgeCidr() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("docker_bridge_cidr"))
}

func (np DataNetworkProfileAttributes) LoadBalancerSku() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("load_balancer_sku"))
}

func (np DataNetworkProfileAttributes) NetworkPlugin() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("network_plugin"))
}

func (np DataNetworkProfileAttributes) NetworkPolicy() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("network_policy"))
}

func (np DataNetworkProfileAttributes) PodCidr() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("pod_cidr"))
}

func (np DataNetworkProfileAttributes) ServiceCidr() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("service_cidr"))
}

type DataOmsAgentAttributes struct {
	ref terra.Reference
}

func (oa DataOmsAgentAttributes) InternalRef() (terra.Reference, error) {
	return oa.ref, nil
}

func (oa DataOmsAgentAttributes) InternalWithRef(ref terra.Reference) DataOmsAgentAttributes {
	return DataOmsAgentAttributes{ref: ref}
}

func (oa DataOmsAgentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oa.ref.InternalTokens()
}

func (oa DataOmsAgentAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("log_analytics_workspace_id"))
}

func (oa DataOmsAgentAttributes) MsiAuthForMonitoringEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(oa.ref.Append("msi_auth_for_monitoring_enabled"))
}

func (oa DataOmsAgentAttributes) OmsAgentIdentity() terra.ListValue[DataOmsAgentOmsAgentIdentityAttributes] {
	return terra.ReferenceAsList[DataOmsAgentOmsAgentIdentityAttributes](oa.ref.Append("oms_agent_identity"))
}

type DataOmsAgentOmsAgentIdentityAttributes struct {
	ref terra.Reference
}

func (oai DataOmsAgentOmsAgentIdentityAttributes) InternalRef() (terra.Reference, error) {
	return oai.ref, nil
}

func (oai DataOmsAgentOmsAgentIdentityAttributes) InternalWithRef(ref terra.Reference) DataOmsAgentOmsAgentIdentityAttributes {
	return DataOmsAgentOmsAgentIdentityAttributes{ref: ref}
}

func (oai DataOmsAgentOmsAgentIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oai.ref.InternalTokens()
}

func (oai DataOmsAgentOmsAgentIdentityAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(oai.ref.Append("client_id"))
}

func (oai DataOmsAgentOmsAgentIdentityAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(oai.ref.Append("object_id"))
}

func (oai DataOmsAgentOmsAgentIdentityAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(oai.ref.Append("user_assigned_identity_id"))
}

type DataServiceMeshProfileAttributes struct {
	ref terra.Reference
}

func (smp DataServiceMeshProfileAttributes) InternalRef() (terra.Reference, error) {
	return smp.ref, nil
}

func (smp DataServiceMeshProfileAttributes) InternalWithRef(ref terra.Reference) DataServiceMeshProfileAttributes {
	return DataServiceMeshProfileAttributes{ref: ref}
}

func (smp DataServiceMeshProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return smp.ref.InternalTokens()
}

func (smp DataServiceMeshProfileAttributes) ExternalIngressGatewayEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(smp.ref.Append("external_ingress_gateway_enabled"))
}

func (smp DataServiceMeshProfileAttributes) InternalIngressGatewayEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(smp.ref.Append("internal_ingress_gateway_enabled"))
}

func (smp DataServiceMeshProfileAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(smp.ref.Append("mode"))
}

type DataServicePrincipalAttributes struct {
	ref terra.Reference
}

func (sp DataServicePrincipalAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp DataServicePrincipalAttributes) InternalWithRef(ref terra.Reference) DataServicePrincipalAttributes {
	return DataServicePrincipalAttributes{ref: ref}
}

func (sp DataServicePrincipalAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp DataServicePrincipalAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("client_id"))
}

type DataStorageProfileAttributes struct {
	ref terra.Reference
}

func (sp DataStorageProfileAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp DataStorageProfileAttributes) InternalWithRef(ref terra.Reference) DataStorageProfileAttributes {
	return DataStorageProfileAttributes{ref: ref}
}

func (sp DataStorageProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp DataStorageProfileAttributes) BlobDriverEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("blob_driver_enabled"))
}

func (sp DataStorageProfileAttributes) DiskDriverEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("disk_driver_enabled"))
}

func (sp DataStorageProfileAttributes) DiskDriverVersion() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("disk_driver_version"))
}

func (sp DataStorageProfileAttributes) FileDriverEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("file_driver_enabled"))
}

func (sp DataStorageProfileAttributes) SnapshotControllerEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("snapshot_controller_enabled"))
}

type DataWindowsProfileAttributes struct {
	ref terra.Reference
}

func (wp DataWindowsProfileAttributes) InternalRef() (terra.Reference, error) {
	return wp.ref, nil
}

func (wp DataWindowsProfileAttributes) InternalWithRef(ref terra.Reference) DataWindowsProfileAttributes {
	return DataWindowsProfileAttributes{ref: ref}
}

func (wp DataWindowsProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wp.ref.InternalTokens()
}

func (wp DataWindowsProfileAttributes) AdminUsername() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("admin_username"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataAciConnectorLinuxState struct {
	SubnetName string `json:"subnet_name"`
}

type DataAgentPoolProfileState struct {
	Count                float64                                    `json:"count"`
	EnableAutoScaling    bool                                       `json:"enable_auto_scaling"`
	EnableNodePublicIp   bool                                       `json:"enable_node_public_ip"`
	MaxCount             float64                                    `json:"max_count"`
	MaxPods              float64                                    `json:"max_pods"`
	MinCount             float64                                    `json:"min_count"`
	Name                 string                                     `json:"name"`
	NodeLabels           map[string]string                          `json:"node_labels"`
	NodePublicIpPrefixId string                                     `json:"node_public_ip_prefix_id"`
	NodeTaints           []string                                   `json:"node_taints"`
	OrchestratorVersion  string                                     `json:"orchestrator_version"`
	OsDiskSizeGb         float64                                    `json:"os_disk_size_gb"`
	OsType               string                                     `json:"os_type"`
	Tags                 map[string]string                          `json:"tags"`
	Type                 string                                     `json:"type"`
	VmSize               string                                     `json:"vm_size"`
	VnetSubnetId         string                                     `json:"vnet_subnet_id"`
	Zones                []string                                   `json:"zones"`
	UpgradeSettings      []DataAgentPoolProfileUpgradeSettingsState `json:"upgrade_settings"`
}

type DataAgentPoolProfileUpgradeSettingsState struct {
	MaxSurge string `json:"max_surge"`
}

type DataAzureActiveDirectoryRoleBasedAccessControlState struct {
	AdminGroupObjectIds []string `json:"admin_group_object_ids"`
	AzureRbacEnabled    bool     `json:"azure_rbac_enabled"`
	ClientAppId         string   `json:"client_app_id"`
	Managed             bool     `json:"managed"`
	ServerAppId         string   `json:"server_app_id"`
	TenantId            string   `json:"tenant_id"`
}

type DataIdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type DataIngressApplicationGatewayState struct {
	EffectiveGatewayId                string                                                                `json:"effective_gateway_id"`
	GatewayId                         string                                                                `json:"gateway_id"`
	GatewayName                       string                                                                `json:"gateway_name"`
	SubnetCidr                        string                                                                `json:"subnet_cidr"`
	SubnetId                          string                                                                `json:"subnet_id"`
	IngressApplicationGatewayIdentity []DataIngressApplicationGatewayIngressApplicationGatewayIdentityState `json:"ingress_application_gateway_identity"`
}

type DataIngressApplicationGatewayIngressApplicationGatewayIdentityState struct {
	ClientId               string `json:"client_id"`
	ObjectId               string `json:"object_id"`
	UserAssignedIdentityId string `json:"user_assigned_identity_id"`
}

type DataKeyManagementServiceState struct {
	KeyVaultKeyId         string `json:"key_vault_key_id"`
	KeyVaultNetworkAccess string `json:"key_vault_network_access"`
}

type DataKeyVaultSecretsProviderState struct {
	SecretRotationEnabled  bool                                             `json:"secret_rotation_enabled"`
	SecretRotationInterval string                                           `json:"secret_rotation_interval"`
	SecretIdentity         []DataKeyVaultSecretsProviderSecretIdentityState `json:"secret_identity"`
}

type DataKeyVaultSecretsProviderSecretIdentityState struct {
	ClientId               string `json:"client_id"`
	ObjectId               string `json:"object_id"`
	UserAssignedIdentityId string `json:"user_assigned_identity_id"`
}

type DataKubeAdminConfigState struct {
	ClientCertificate    string `json:"client_certificate"`
	ClientKey            string `json:"client_key"`
	ClusterCaCertificate string `json:"cluster_ca_certificate"`
	Host                 string `json:"host"`
	Password             string `json:"password"`
	Username             string `json:"username"`
}

type DataKubeConfigState struct {
	ClientCertificate    string `json:"client_certificate"`
	ClientKey            string `json:"client_key"`
	ClusterCaCertificate string `json:"cluster_ca_certificate"`
	Host                 string `json:"host"`
	Password             string `json:"password"`
	Username             string `json:"username"`
}

type DataKubeletIdentityState struct {
	ClientId               string `json:"client_id"`
	ObjectId               string `json:"object_id"`
	UserAssignedIdentityId string `json:"user_assigned_identity_id"`
}

type DataLinuxProfileState struct {
	AdminUsername string                        `json:"admin_username"`
	SshKey        []DataLinuxProfileSshKeyState `json:"ssh_key"`
}

type DataLinuxProfileSshKeyState struct {
	KeyData string `json:"key_data"`
}

type DataMicrosoftDefenderState struct {
	LogAnalyticsWorkspaceId string `json:"log_analytics_workspace_id"`
}

type DataNetworkProfileState struct {
	DnsServiceIp     string `json:"dns_service_ip"`
	DockerBridgeCidr string `json:"docker_bridge_cidr"`
	LoadBalancerSku  string `json:"load_balancer_sku"`
	NetworkPlugin    string `json:"network_plugin"`
	NetworkPolicy    string `json:"network_policy"`
	PodCidr          string `json:"pod_cidr"`
	ServiceCidr      string `json:"service_cidr"`
}

type DataOmsAgentState struct {
	LogAnalyticsWorkspaceId     string                              `json:"log_analytics_workspace_id"`
	MsiAuthForMonitoringEnabled bool                                `json:"msi_auth_for_monitoring_enabled"`
	OmsAgentIdentity            []DataOmsAgentOmsAgentIdentityState `json:"oms_agent_identity"`
}

type DataOmsAgentOmsAgentIdentityState struct {
	ClientId               string `json:"client_id"`
	ObjectId               string `json:"object_id"`
	UserAssignedIdentityId string `json:"user_assigned_identity_id"`
}

type DataServiceMeshProfileState struct {
	ExternalIngressGatewayEnabled bool   `json:"external_ingress_gateway_enabled"`
	InternalIngressGatewayEnabled bool   `json:"internal_ingress_gateway_enabled"`
	Mode                          string `json:"mode"`
}

type DataServicePrincipalState struct {
	ClientId string `json:"client_id"`
}

type DataStorageProfileState struct {
	BlobDriverEnabled         bool   `json:"blob_driver_enabled"`
	DiskDriverEnabled         bool   `json:"disk_driver_enabled"`
	DiskDriverVersion         string `json:"disk_driver_version"`
	FileDriverEnabled         bool   `json:"file_driver_enabled"`
	SnapshotControllerEnabled bool   `json:"snapshot_controller_enabled"`
}

type DataWindowsProfileState struct {
	AdminUsername string `json:"admin_username"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
