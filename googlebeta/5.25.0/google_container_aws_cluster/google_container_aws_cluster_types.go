// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_container_aws_cluster

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Authorization struct {
	// AuthorizationAdminGroups: min=0
	AdminGroups []AuthorizationAdminGroups `hcl:"admin_groups,block" validate:"min=0"`
	// AuthorizationAdminUsers: min=1
	AdminUsers []AuthorizationAdminUsers `hcl:"admin_users,block" validate:"min=1"`
}

type AuthorizationAdminGroups struct {
	// Group: string, required
	Group terra.StringValue `hcl:"group,attr" validate:"required"`
}

type AuthorizationAdminUsers struct {
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type BinaryAuthorization struct {
	// EvaluationMode: string, optional
	EvaluationMode terra.StringValue `hcl:"evaluation_mode,attr"`
}

type ControlPlane struct {
	// IamInstanceProfile: string, required
	IamInstanceProfile terra.StringValue `hcl:"iam_instance_profile,attr" validate:"required"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// SecurityGroupIds: list of string, optional
	SecurityGroupIds terra.ListValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// SubnetIds: list of string, required
	SubnetIds terra.ListValue[terra.StringValue] `hcl:"subnet_ids,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// ControlPlaneAwsServicesAuthentication: required
	AwsServicesAuthentication *ControlPlaneAwsServicesAuthentication `hcl:"aws_services_authentication,block" validate:"required"`
	// ControlPlaneConfigEncryption: required
	ConfigEncryption *ControlPlaneConfigEncryption `hcl:"config_encryption,block" validate:"required"`
	// ControlPlaneDatabaseEncryption: required
	DatabaseEncryption *ControlPlaneDatabaseEncryption `hcl:"database_encryption,block" validate:"required"`
	// ControlPlaneInstancePlacement: optional
	InstancePlacement *ControlPlaneInstancePlacement `hcl:"instance_placement,block"`
	// ControlPlaneMainVolume: optional
	MainVolume *ControlPlaneMainVolume `hcl:"main_volume,block"`
	// ControlPlaneProxyConfig: optional
	ProxyConfig *ControlPlaneProxyConfig `hcl:"proxy_config,block"`
	// ControlPlaneRootVolume: optional
	RootVolume *ControlPlaneRootVolume `hcl:"root_volume,block"`
	// ControlPlaneSshConfig: optional
	SshConfig *ControlPlaneSshConfig `hcl:"ssh_config,block"`
}

type ControlPlaneAwsServicesAuthentication struct {
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// RoleSessionName: string, optional
	RoleSessionName terra.StringValue `hcl:"role_session_name,attr"`
}

type ControlPlaneConfigEncryption struct {
	// KmsKeyArn: string, required
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr" validate:"required"`
}

type ControlPlaneDatabaseEncryption struct {
	// KmsKeyArn: string, required
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr" validate:"required"`
}

type ControlPlaneInstancePlacement struct {
	// Tenancy: string, optional
	Tenancy terra.StringValue `hcl:"tenancy,attr"`
}

type ControlPlaneMainVolume struct {
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// SizeGib: number, optional
	SizeGib terra.NumberValue `hcl:"size_gib,attr"`
	// Throughput: number, optional
	Throughput terra.NumberValue `hcl:"throughput,attr"`
	// VolumeType: string, optional
	VolumeType terra.StringValue `hcl:"volume_type,attr"`
}

type ControlPlaneProxyConfig struct {
	// SecretArn: string, required
	SecretArn terra.StringValue `hcl:"secret_arn,attr" validate:"required"`
	// SecretVersion: string, required
	SecretVersion terra.StringValue `hcl:"secret_version,attr" validate:"required"`
}

type ControlPlaneRootVolume struct {
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// SizeGib: number, optional
	SizeGib terra.NumberValue `hcl:"size_gib,attr"`
	// Throughput: number, optional
	Throughput terra.NumberValue `hcl:"throughput,attr"`
	// VolumeType: string, optional
	VolumeType terra.StringValue `hcl:"volume_type,attr"`
}

type ControlPlaneSshConfig struct {
	// Ec2KeyPair: string, required
	Ec2KeyPair terra.StringValue `hcl:"ec2_key_pair,attr" validate:"required"`
}

type Fleet struct {
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}

type LoggingConfig struct {
	// LoggingConfigComponentConfig: optional
	ComponentConfig *LoggingConfigComponentConfig `hcl:"component_config,block"`
}

type LoggingConfigComponentConfig struct {
	// EnableComponents: list of string, optional
	EnableComponents terra.ListValue[terra.StringValue] `hcl:"enable_components,attr"`
}

type Networking struct {
	// PerNodePoolSgRulesDisabled: bool, optional
	PerNodePoolSgRulesDisabled terra.BoolValue `hcl:"per_node_pool_sg_rules_disabled,attr"`
	// PodAddressCidrBlocks: list of string, required
	PodAddressCidrBlocks terra.ListValue[terra.StringValue] `hcl:"pod_address_cidr_blocks,attr" validate:"required"`
	// ServiceAddressCidrBlocks: list of string, required
	ServiceAddressCidrBlocks terra.ListValue[terra.StringValue] `hcl:"service_address_cidr_blocks,attr" validate:"required"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type WorkloadIdentityConfigAttributes struct {
	ref terra.Reference
}

func (wic WorkloadIdentityConfigAttributes) InternalRef() (terra.Reference, error) {
	return wic.ref, nil
}

func (wic WorkloadIdentityConfigAttributes) InternalWithRef(ref terra.Reference) WorkloadIdentityConfigAttributes {
	return WorkloadIdentityConfigAttributes{ref: ref}
}

func (wic WorkloadIdentityConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wic.ref.InternalTokens()
}

func (wic WorkloadIdentityConfigAttributes) IdentityProvider() terra.StringValue {
	return terra.ReferenceAsString(wic.ref.Append("identity_provider"))
}

func (wic WorkloadIdentityConfigAttributes) IssuerUri() terra.StringValue {
	return terra.ReferenceAsString(wic.ref.Append("issuer_uri"))
}

func (wic WorkloadIdentityConfigAttributes) WorkloadPool() terra.StringValue {
	return terra.ReferenceAsString(wic.ref.Append("workload_pool"))
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

func (a AuthorizationAttributes) AdminGroups() terra.ListValue[AuthorizationAdminGroupsAttributes] {
	return terra.ReferenceAsList[AuthorizationAdminGroupsAttributes](a.ref.Append("admin_groups"))
}

func (a AuthorizationAttributes) AdminUsers() terra.ListValue[AuthorizationAdminUsersAttributes] {
	return terra.ReferenceAsList[AuthorizationAdminUsersAttributes](a.ref.Append("admin_users"))
}

type AuthorizationAdminGroupsAttributes struct {
	ref terra.Reference
}

func (ag AuthorizationAdminGroupsAttributes) InternalRef() (terra.Reference, error) {
	return ag.ref, nil
}

func (ag AuthorizationAdminGroupsAttributes) InternalWithRef(ref terra.Reference) AuthorizationAdminGroupsAttributes {
	return AuthorizationAdminGroupsAttributes{ref: ref}
}

func (ag AuthorizationAdminGroupsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ag.ref.InternalTokens()
}

func (ag AuthorizationAdminGroupsAttributes) Group() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("group"))
}

type AuthorizationAdminUsersAttributes struct {
	ref terra.Reference
}

func (au AuthorizationAdminUsersAttributes) InternalRef() (terra.Reference, error) {
	return au.ref, nil
}

func (au AuthorizationAdminUsersAttributes) InternalWithRef(ref terra.Reference) AuthorizationAdminUsersAttributes {
	return AuthorizationAdminUsersAttributes{ref: ref}
}

func (au AuthorizationAdminUsersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return au.ref.InternalTokens()
}

func (au AuthorizationAdminUsersAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("username"))
}

type BinaryAuthorizationAttributes struct {
	ref terra.Reference
}

func (ba BinaryAuthorizationAttributes) InternalRef() (terra.Reference, error) {
	return ba.ref, nil
}

func (ba BinaryAuthorizationAttributes) InternalWithRef(ref terra.Reference) BinaryAuthorizationAttributes {
	return BinaryAuthorizationAttributes{ref: ref}
}

func (ba BinaryAuthorizationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ba.ref.InternalTokens()
}

func (ba BinaryAuthorizationAttributes) EvaluationMode() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("evaluation_mode"))
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

func (cp ControlPlaneAttributes) IamInstanceProfile() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("iam_instance_profile"))
}

func (cp ControlPlaneAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("instance_type"))
}

func (cp ControlPlaneAttributes) SecurityGroupIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cp.ref.Append("security_group_ids"))
}

func (cp ControlPlaneAttributes) SubnetIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cp.ref.Append("subnet_ids"))
}

func (cp ControlPlaneAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cp.ref.Append("tags"))
}

func (cp ControlPlaneAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("version"))
}

func (cp ControlPlaneAttributes) AwsServicesAuthentication() terra.ListValue[ControlPlaneAwsServicesAuthenticationAttributes] {
	return terra.ReferenceAsList[ControlPlaneAwsServicesAuthenticationAttributes](cp.ref.Append("aws_services_authentication"))
}

func (cp ControlPlaneAttributes) ConfigEncryption() terra.ListValue[ControlPlaneConfigEncryptionAttributes] {
	return terra.ReferenceAsList[ControlPlaneConfigEncryptionAttributes](cp.ref.Append("config_encryption"))
}

func (cp ControlPlaneAttributes) DatabaseEncryption() terra.ListValue[ControlPlaneDatabaseEncryptionAttributes] {
	return terra.ReferenceAsList[ControlPlaneDatabaseEncryptionAttributes](cp.ref.Append("database_encryption"))
}

func (cp ControlPlaneAttributes) InstancePlacement() terra.ListValue[ControlPlaneInstancePlacementAttributes] {
	return terra.ReferenceAsList[ControlPlaneInstancePlacementAttributes](cp.ref.Append("instance_placement"))
}

func (cp ControlPlaneAttributes) MainVolume() terra.ListValue[ControlPlaneMainVolumeAttributes] {
	return terra.ReferenceAsList[ControlPlaneMainVolumeAttributes](cp.ref.Append("main_volume"))
}

func (cp ControlPlaneAttributes) ProxyConfig() terra.ListValue[ControlPlaneProxyConfigAttributes] {
	return terra.ReferenceAsList[ControlPlaneProxyConfigAttributes](cp.ref.Append("proxy_config"))
}

func (cp ControlPlaneAttributes) RootVolume() terra.ListValue[ControlPlaneRootVolumeAttributes] {
	return terra.ReferenceAsList[ControlPlaneRootVolumeAttributes](cp.ref.Append("root_volume"))
}

func (cp ControlPlaneAttributes) SshConfig() terra.ListValue[ControlPlaneSshConfigAttributes] {
	return terra.ReferenceAsList[ControlPlaneSshConfigAttributes](cp.ref.Append("ssh_config"))
}

type ControlPlaneAwsServicesAuthenticationAttributes struct {
	ref terra.Reference
}

func (asa ControlPlaneAwsServicesAuthenticationAttributes) InternalRef() (terra.Reference, error) {
	return asa.ref, nil
}

func (asa ControlPlaneAwsServicesAuthenticationAttributes) InternalWithRef(ref terra.Reference) ControlPlaneAwsServicesAuthenticationAttributes {
	return ControlPlaneAwsServicesAuthenticationAttributes{ref: ref}
}

func (asa ControlPlaneAwsServicesAuthenticationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return asa.ref.InternalTokens()
}

func (asa ControlPlaneAwsServicesAuthenticationAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("role_arn"))
}

func (asa ControlPlaneAwsServicesAuthenticationAttributes) RoleSessionName() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("role_session_name"))
}

type ControlPlaneConfigEncryptionAttributes struct {
	ref terra.Reference
}

func (ce ControlPlaneConfigEncryptionAttributes) InternalRef() (terra.Reference, error) {
	return ce.ref, nil
}

func (ce ControlPlaneConfigEncryptionAttributes) InternalWithRef(ref terra.Reference) ControlPlaneConfigEncryptionAttributes {
	return ControlPlaneConfigEncryptionAttributes{ref: ref}
}

func (ce ControlPlaneConfigEncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ce.ref.InternalTokens()
}

func (ce ControlPlaneConfigEncryptionAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("kms_key_arn"))
}

type ControlPlaneDatabaseEncryptionAttributes struct {
	ref terra.Reference
}

func (de ControlPlaneDatabaseEncryptionAttributes) InternalRef() (terra.Reference, error) {
	return de.ref, nil
}

func (de ControlPlaneDatabaseEncryptionAttributes) InternalWithRef(ref terra.Reference) ControlPlaneDatabaseEncryptionAttributes {
	return ControlPlaneDatabaseEncryptionAttributes{ref: ref}
}

func (de ControlPlaneDatabaseEncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return de.ref.InternalTokens()
}

func (de ControlPlaneDatabaseEncryptionAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("kms_key_arn"))
}

type ControlPlaneInstancePlacementAttributes struct {
	ref terra.Reference
}

func (ip ControlPlaneInstancePlacementAttributes) InternalRef() (terra.Reference, error) {
	return ip.ref, nil
}

func (ip ControlPlaneInstancePlacementAttributes) InternalWithRef(ref terra.Reference) ControlPlaneInstancePlacementAttributes {
	return ControlPlaneInstancePlacementAttributes{ref: ref}
}

func (ip ControlPlaneInstancePlacementAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ip.ref.InternalTokens()
}

func (ip ControlPlaneInstancePlacementAttributes) Tenancy() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("tenancy"))
}

type ControlPlaneMainVolumeAttributes struct {
	ref terra.Reference
}

func (mv ControlPlaneMainVolumeAttributes) InternalRef() (terra.Reference, error) {
	return mv.ref, nil
}

func (mv ControlPlaneMainVolumeAttributes) InternalWithRef(ref terra.Reference) ControlPlaneMainVolumeAttributes {
	return ControlPlaneMainVolumeAttributes{ref: ref}
}

func (mv ControlPlaneMainVolumeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mv.ref.InternalTokens()
}

func (mv ControlPlaneMainVolumeAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(mv.ref.Append("iops"))
}

func (mv ControlPlaneMainVolumeAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(mv.ref.Append("kms_key_arn"))
}

func (mv ControlPlaneMainVolumeAttributes) SizeGib() terra.NumberValue {
	return terra.ReferenceAsNumber(mv.ref.Append("size_gib"))
}

func (mv ControlPlaneMainVolumeAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(mv.ref.Append("throughput"))
}

func (mv ControlPlaneMainVolumeAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(mv.ref.Append("volume_type"))
}

type ControlPlaneProxyConfigAttributes struct {
	ref terra.Reference
}

func (pc ControlPlaneProxyConfigAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc ControlPlaneProxyConfigAttributes) InternalWithRef(ref terra.Reference) ControlPlaneProxyConfigAttributes {
	return ControlPlaneProxyConfigAttributes{ref: ref}
}

func (pc ControlPlaneProxyConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc ControlPlaneProxyConfigAttributes) SecretArn() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("secret_arn"))
}

func (pc ControlPlaneProxyConfigAttributes) SecretVersion() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("secret_version"))
}

type ControlPlaneRootVolumeAttributes struct {
	ref terra.Reference
}

func (rv ControlPlaneRootVolumeAttributes) InternalRef() (terra.Reference, error) {
	return rv.ref, nil
}

func (rv ControlPlaneRootVolumeAttributes) InternalWithRef(ref terra.Reference) ControlPlaneRootVolumeAttributes {
	return ControlPlaneRootVolumeAttributes{ref: ref}
}

func (rv ControlPlaneRootVolumeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rv.ref.InternalTokens()
}

func (rv ControlPlaneRootVolumeAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(rv.ref.Append("iops"))
}

func (rv ControlPlaneRootVolumeAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(rv.ref.Append("kms_key_arn"))
}

func (rv ControlPlaneRootVolumeAttributes) SizeGib() terra.NumberValue {
	return terra.ReferenceAsNumber(rv.ref.Append("size_gib"))
}

func (rv ControlPlaneRootVolumeAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(rv.ref.Append("throughput"))
}

func (rv ControlPlaneRootVolumeAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(rv.ref.Append("volume_type"))
}

type ControlPlaneSshConfigAttributes struct {
	ref terra.Reference
}

func (sc ControlPlaneSshConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc ControlPlaneSshConfigAttributes) InternalWithRef(ref terra.Reference) ControlPlaneSshConfigAttributes {
	return ControlPlaneSshConfigAttributes{ref: ref}
}

func (sc ControlPlaneSshConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc ControlPlaneSshConfigAttributes) Ec2KeyPair() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("ec2_key_pair"))
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

type LoggingConfigAttributes struct {
	ref terra.Reference
}

func (lc LoggingConfigAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc LoggingConfigAttributes) InternalWithRef(ref terra.Reference) LoggingConfigAttributes {
	return LoggingConfigAttributes{ref: ref}
}

func (lc LoggingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc LoggingConfigAttributes) ComponentConfig() terra.ListValue[LoggingConfigComponentConfigAttributes] {
	return terra.ReferenceAsList[LoggingConfigComponentConfigAttributes](lc.ref.Append("component_config"))
}

type LoggingConfigComponentConfigAttributes struct {
	ref terra.Reference
}

func (cc LoggingConfigComponentConfigAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc LoggingConfigComponentConfigAttributes) InternalWithRef(ref terra.Reference) LoggingConfigComponentConfigAttributes {
	return LoggingConfigComponentConfigAttributes{ref: ref}
}

func (cc LoggingConfigComponentConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc LoggingConfigComponentConfigAttributes) EnableComponents() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cc.ref.Append("enable_components"))
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

func (n NetworkingAttributes) PerNodePoolSgRulesDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(n.ref.Append("per_node_pool_sg_rules_disabled"))
}

func (n NetworkingAttributes) PodAddressCidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](n.ref.Append("pod_address_cidr_blocks"))
}

func (n NetworkingAttributes) ServiceAddressCidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](n.ref.Append("service_address_cidr_blocks"))
}

func (n NetworkingAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("vpc_id"))
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

type WorkloadIdentityConfigState struct {
	IdentityProvider string `json:"identity_provider"`
	IssuerUri        string `json:"issuer_uri"`
	WorkloadPool     string `json:"workload_pool"`
}

type AuthorizationState struct {
	AdminGroups []AuthorizationAdminGroupsState `json:"admin_groups"`
	AdminUsers  []AuthorizationAdminUsersState  `json:"admin_users"`
}

type AuthorizationAdminGroupsState struct {
	Group string `json:"group"`
}

type AuthorizationAdminUsersState struct {
	Username string `json:"username"`
}

type BinaryAuthorizationState struct {
	EvaluationMode string `json:"evaluation_mode"`
}

type ControlPlaneState struct {
	IamInstanceProfile        string                                       `json:"iam_instance_profile"`
	InstanceType              string                                       `json:"instance_type"`
	SecurityGroupIds          []string                                     `json:"security_group_ids"`
	SubnetIds                 []string                                     `json:"subnet_ids"`
	Tags                      map[string]string                            `json:"tags"`
	Version                   string                                       `json:"version"`
	AwsServicesAuthentication []ControlPlaneAwsServicesAuthenticationState `json:"aws_services_authentication"`
	ConfigEncryption          []ControlPlaneConfigEncryptionState          `json:"config_encryption"`
	DatabaseEncryption        []ControlPlaneDatabaseEncryptionState        `json:"database_encryption"`
	InstancePlacement         []ControlPlaneInstancePlacementState         `json:"instance_placement"`
	MainVolume                []ControlPlaneMainVolumeState                `json:"main_volume"`
	ProxyConfig               []ControlPlaneProxyConfigState               `json:"proxy_config"`
	RootVolume                []ControlPlaneRootVolumeState                `json:"root_volume"`
	SshConfig                 []ControlPlaneSshConfigState                 `json:"ssh_config"`
}

type ControlPlaneAwsServicesAuthenticationState struct {
	RoleArn         string `json:"role_arn"`
	RoleSessionName string `json:"role_session_name"`
}

type ControlPlaneConfigEncryptionState struct {
	KmsKeyArn string `json:"kms_key_arn"`
}

type ControlPlaneDatabaseEncryptionState struct {
	KmsKeyArn string `json:"kms_key_arn"`
}

type ControlPlaneInstancePlacementState struct {
	Tenancy string `json:"tenancy"`
}

type ControlPlaneMainVolumeState struct {
	Iops       float64 `json:"iops"`
	KmsKeyArn  string  `json:"kms_key_arn"`
	SizeGib    float64 `json:"size_gib"`
	Throughput float64 `json:"throughput"`
	VolumeType string  `json:"volume_type"`
}

type ControlPlaneProxyConfigState struct {
	SecretArn     string `json:"secret_arn"`
	SecretVersion string `json:"secret_version"`
}

type ControlPlaneRootVolumeState struct {
	Iops       float64 `json:"iops"`
	KmsKeyArn  string  `json:"kms_key_arn"`
	SizeGib    float64 `json:"size_gib"`
	Throughput float64 `json:"throughput"`
	VolumeType string  `json:"volume_type"`
}

type ControlPlaneSshConfigState struct {
	Ec2KeyPair string `json:"ec2_key_pair"`
}

type FleetState struct {
	Membership string `json:"membership"`
	Project    string `json:"project"`
}

type LoggingConfigState struct {
	ComponentConfig []LoggingConfigComponentConfigState `json:"component_config"`
}

type LoggingConfigComponentConfigState struct {
	EnableComponents []string `json:"enable_components"`
}

type NetworkingState struct {
	PerNodePoolSgRulesDisabled bool     `json:"per_node_pool_sg_rules_disabled"`
	PodAddressCidrBlocks       []string `json:"pod_address_cidr_blocks"`
	ServiceAddressCidrBlocks   []string `json:"service_address_cidr_blocks"`
	VpcId                      string   `json:"vpc_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
