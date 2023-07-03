// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	hpccache "github.com/golingon/terraproviders/azurerm/3.63.0/hpccache"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHpcCache creates a new instance of [HpcCache].
func NewHpcCache(name string, args HpcCacheArgs) *HpcCache {
	return &HpcCache{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HpcCache)(nil)

// HpcCache represents the Terraform resource azurerm_hpc_cache.
type HpcCache struct {
	Name      string
	Args      HpcCacheArgs
	state     *hpcCacheState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HpcCache].
func (hc *HpcCache) Type() string {
	return "azurerm_hpc_cache"
}

// LocalName returns the local name for [HpcCache].
func (hc *HpcCache) LocalName() string {
	return hc.Name
}

// Configuration returns the configuration (args) for [HpcCache].
func (hc *HpcCache) Configuration() interface{} {
	return hc.Args
}

// DependOn is used for other resources to depend on [HpcCache].
func (hc *HpcCache) DependOn() terra.Reference {
	return terra.ReferenceResource(hc)
}

// Dependencies returns the list of resources [HpcCache] depends_on.
func (hc *HpcCache) Dependencies() terra.Dependencies {
	return hc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HpcCache].
func (hc *HpcCache) LifecycleManagement() *terra.Lifecycle {
	return hc.Lifecycle
}

// Attributes returns the attributes for [HpcCache].
func (hc *HpcCache) Attributes() hpcCacheAttributes {
	return hpcCacheAttributes{ref: terra.ReferenceResource(hc)}
}

// ImportState imports the given attribute values into [HpcCache]'s state.
func (hc *HpcCache) ImportState(av io.Reader) error {
	hc.state = &hpcCacheState{}
	if err := json.NewDecoder(av).Decode(hc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hc.Type(), hc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HpcCache] has state.
func (hc *HpcCache) State() (*hpcCacheState, bool) {
	return hc.state, hc.state != nil
}

// StateMust returns the state for [HpcCache]. Panics if the state is nil.
func (hc *HpcCache) StateMust() *hpcCacheState {
	if hc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hc.Type(), hc.LocalName()))
	}
	return hc.state
}

// HpcCacheArgs contains the configurations for azurerm_hpc_cache.
type HpcCacheArgs struct {
	// AutomaticallyRotateKeyToLatestEnabled: bool, optional
	AutomaticallyRotateKeyToLatestEnabled terra.BoolValue `hcl:"automatically_rotate_key_to_latest_enabled,attr"`
	// CacheSizeInGb: number, required
	CacheSizeInGb terra.NumberValue `hcl:"cache_size_in_gb,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultKeyId: string, optional
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Mtu: number, optional
	Mtu terra.NumberValue `hcl:"mtu,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NtpServer: string, optional
	NtpServer terra.StringValue `hcl:"ntp_server,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// DefaultAccessPolicy: optional
	DefaultAccessPolicy *hpccache.DefaultAccessPolicy `hcl:"default_access_policy,block"`
	// DirectoryActiveDirectory: optional
	DirectoryActiveDirectory *hpccache.DirectoryActiveDirectory `hcl:"directory_active_directory,block"`
	// DirectoryFlatFile: optional
	DirectoryFlatFile *hpccache.DirectoryFlatFile `hcl:"directory_flat_file,block"`
	// DirectoryLdap: optional
	DirectoryLdap *hpccache.DirectoryLdap `hcl:"directory_ldap,block"`
	// Dns: optional
	Dns *hpccache.Dns `hcl:"dns,block"`
	// Identity: optional
	Identity *hpccache.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *hpccache.Timeouts `hcl:"timeouts,block"`
}
type hpcCacheAttributes struct {
	ref terra.Reference
}

// AutomaticallyRotateKeyToLatestEnabled returns a reference to field automatically_rotate_key_to_latest_enabled of azurerm_hpc_cache.
func (hc hpcCacheAttributes) AutomaticallyRotateKeyToLatestEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(hc.ref.Append("automatically_rotate_key_to_latest_enabled"))
}

// CacheSizeInGb returns a reference to field cache_size_in_gb of azurerm_hpc_cache.
func (hc hpcCacheAttributes) CacheSizeInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("cache_size_in_gb"))
}

// Id returns a reference to field id of azurerm_hpc_cache.
func (hc hpcCacheAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("id"))
}

// KeyVaultKeyId returns a reference to field key_vault_key_id of azurerm_hpc_cache.
func (hc hpcCacheAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("key_vault_key_id"))
}

// Location returns a reference to field location of azurerm_hpc_cache.
func (hc hpcCacheAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("location"))
}

// MountAddresses returns a reference to field mount_addresses of azurerm_hpc_cache.
func (hc hpcCacheAttributes) MountAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](hc.ref.Append("mount_addresses"))
}

// Mtu returns a reference to field mtu of azurerm_hpc_cache.
func (hc hpcCacheAttributes) Mtu() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("mtu"))
}

// Name returns a reference to field name of azurerm_hpc_cache.
func (hc hpcCacheAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("name"))
}

// NtpServer returns a reference to field ntp_server of azurerm_hpc_cache.
func (hc hpcCacheAttributes) NtpServer() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("ntp_server"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_hpc_cache.
func (hc hpcCacheAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_hpc_cache.
func (hc hpcCacheAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("sku_name"))
}

// SubnetId returns a reference to field subnet_id of azurerm_hpc_cache.
func (hc hpcCacheAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_hpc_cache.
func (hc hpcCacheAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hc.ref.Append("tags"))
}

func (hc hpcCacheAttributes) DefaultAccessPolicy() terra.ListValue[hpccache.DefaultAccessPolicyAttributes] {
	return terra.ReferenceAsList[hpccache.DefaultAccessPolicyAttributes](hc.ref.Append("default_access_policy"))
}

func (hc hpcCacheAttributes) DirectoryActiveDirectory() terra.ListValue[hpccache.DirectoryActiveDirectoryAttributes] {
	return terra.ReferenceAsList[hpccache.DirectoryActiveDirectoryAttributes](hc.ref.Append("directory_active_directory"))
}

func (hc hpcCacheAttributes) DirectoryFlatFile() terra.ListValue[hpccache.DirectoryFlatFileAttributes] {
	return terra.ReferenceAsList[hpccache.DirectoryFlatFileAttributes](hc.ref.Append("directory_flat_file"))
}

func (hc hpcCacheAttributes) DirectoryLdap() terra.ListValue[hpccache.DirectoryLdapAttributes] {
	return terra.ReferenceAsList[hpccache.DirectoryLdapAttributes](hc.ref.Append("directory_ldap"))
}

func (hc hpcCacheAttributes) Dns() terra.ListValue[hpccache.DnsAttributes] {
	return terra.ReferenceAsList[hpccache.DnsAttributes](hc.ref.Append("dns"))
}

func (hc hpcCacheAttributes) Identity() terra.ListValue[hpccache.IdentityAttributes] {
	return terra.ReferenceAsList[hpccache.IdentityAttributes](hc.ref.Append("identity"))
}

func (hc hpcCacheAttributes) Timeouts() hpccache.TimeoutsAttributes {
	return terra.ReferenceAsSingle[hpccache.TimeoutsAttributes](hc.ref.Append("timeouts"))
}

type hpcCacheState struct {
	AutomaticallyRotateKeyToLatestEnabled bool                                     `json:"automatically_rotate_key_to_latest_enabled"`
	CacheSizeInGb                         float64                                  `json:"cache_size_in_gb"`
	Id                                    string                                   `json:"id"`
	KeyVaultKeyId                         string                                   `json:"key_vault_key_id"`
	Location                              string                                   `json:"location"`
	MountAddresses                        []string                                 `json:"mount_addresses"`
	Mtu                                   float64                                  `json:"mtu"`
	Name                                  string                                   `json:"name"`
	NtpServer                             string                                   `json:"ntp_server"`
	ResourceGroupName                     string                                   `json:"resource_group_name"`
	SkuName                               string                                   `json:"sku_name"`
	SubnetId                              string                                   `json:"subnet_id"`
	Tags                                  map[string]string                        `json:"tags"`
	DefaultAccessPolicy                   []hpccache.DefaultAccessPolicyState      `json:"default_access_policy"`
	DirectoryActiveDirectory              []hpccache.DirectoryActiveDirectoryState `json:"directory_active_directory"`
	DirectoryFlatFile                     []hpccache.DirectoryFlatFileState        `json:"directory_flat_file"`
	DirectoryLdap                         []hpccache.DirectoryLdapState            `json:"directory_ldap"`
	Dns                                   []hpccache.DnsState                      `json:"dns"`
	Identity                              []hpccache.IdentityState                 `json:"identity"`
	Timeouts                              *hpccache.TimeoutsState                  `json:"timeouts"`
}
