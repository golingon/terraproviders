// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datarediscache "github.com/golingon/terraproviders/azurerm/3.55.0/datarediscache"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRedisCache creates a new instance of [DataRedisCache].
func NewDataRedisCache(name string, args DataRedisCacheArgs) *DataRedisCache {
	return &DataRedisCache{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRedisCache)(nil)

// DataRedisCache represents the Terraform data resource azurerm_redis_cache.
type DataRedisCache struct {
	Name string
	Args DataRedisCacheArgs
}

// DataSource returns the Terraform object type for [DataRedisCache].
func (rc *DataRedisCache) DataSource() string {
	return "azurerm_redis_cache"
}

// LocalName returns the local name for [DataRedisCache].
func (rc *DataRedisCache) LocalName() string {
	return rc.Name
}

// Configuration returns the configuration (args) for [DataRedisCache].
func (rc *DataRedisCache) Configuration() interface{} {
	return rc.Args
}

// Attributes returns the attributes for [DataRedisCache].
func (rc *DataRedisCache) Attributes() dataRedisCacheAttributes {
	return dataRedisCacheAttributes{ref: terra.ReferenceDataResource(rc)}
}

// DataRedisCacheArgs contains the configurations for azurerm_redis_cache.
type DataRedisCacheArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// PatchSchedule: min=0
	PatchSchedule []datarediscache.PatchSchedule `hcl:"patch_schedule,block" validate:"min=0"`
	// RedisConfiguration: min=0
	RedisConfiguration []datarediscache.RedisConfiguration `hcl:"redis_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datarediscache.Timeouts `hcl:"timeouts,block"`
}
type dataRedisCacheAttributes struct {
	ref terra.Reference
}

// Capacity returns a reference to field capacity of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("capacity"))
}

// EnableNonSslPort returns a reference to field enable_non_ssl_port of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) EnableNonSslPort() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("enable_non_ssl_port"))
}

// Family returns a reference to field family of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) Family() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("family"))
}

// Hostname returns a reference to field hostname of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("hostname"))
}

// Id returns a reference to field id of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("location"))
}

// MinimumTlsVersion returns a reference to field minimum_tls_version of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) MinimumTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("minimum_tls_version"))
}

// Name returns a reference to field name of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("name"))
}

// Port returns a reference to field port of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("port"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("primary_access_key"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("primary_connection_string"))
}

// PrivateStaticIpAddress returns a reference to field private_static_ip_address of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) PrivateStaticIpAddress() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("private_static_ip_address"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("secondary_access_key"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("secondary_connection_string"))
}

// ShardCount returns a reference to field shard_count of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) ShardCount() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("shard_count"))
}

// SkuName returns a reference to field sku_name of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("sku_name"))
}

// SslPort returns a reference to field ssl_port of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) SslPort() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("ssl_port"))
}

// SubnetId returns a reference to field subnet_id of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rc.ref.Append("tags"))
}

// Zones returns a reference to field zones of azurerm_redis_cache.
func (rc dataRedisCacheAttributes) Zones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rc.ref.Append("zones"))
}

func (rc dataRedisCacheAttributes) PatchSchedule() terra.ListValue[datarediscache.PatchScheduleAttributes] {
	return terra.ReferenceAsList[datarediscache.PatchScheduleAttributes](rc.ref.Append("patch_schedule"))
}

func (rc dataRedisCacheAttributes) RedisConfiguration() terra.ListValue[datarediscache.RedisConfigurationAttributes] {
	return terra.ReferenceAsList[datarediscache.RedisConfigurationAttributes](rc.ref.Append("redis_configuration"))
}

func (rc dataRedisCacheAttributes) Timeouts() datarediscache.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datarediscache.TimeoutsAttributes](rc.ref.Append("timeouts"))
}
