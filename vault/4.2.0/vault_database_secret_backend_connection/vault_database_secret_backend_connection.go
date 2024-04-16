// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_database_secret_backend_connection

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource vault_database_secret_backend_connection.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultDatabaseSecretBackendConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vdsbc *Resource) Type() string {
	return "vault_database_secret_backend_connection"
}

// LocalName returns the local name for [Resource].
func (vdsbc *Resource) LocalName() string {
	return vdsbc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vdsbc *Resource) Configuration() interface{} {
	return vdsbc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vdsbc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vdsbc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vdsbc *Resource) Dependencies() terra.Dependencies {
	return vdsbc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vdsbc *Resource) LifecycleManagement() *terra.Lifecycle {
	return vdsbc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vdsbc *Resource) Attributes() vaultDatabaseSecretBackendConnectionAttributes {
	return vaultDatabaseSecretBackendConnectionAttributes{ref: terra.ReferenceResource(vdsbc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vdsbc *Resource) ImportState(state io.Reader) error {
	vdsbc.state = &vaultDatabaseSecretBackendConnectionState{}
	if err := json.NewDecoder(state).Decode(vdsbc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vdsbc.Type(), vdsbc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vdsbc *Resource) State() (*vaultDatabaseSecretBackendConnectionState, bool) {
	return vdsbc.state, vdsbc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vdsbc *Resource) StateMust() *vaultDatabaseSecretBackendConnectionState {
	if vdsbc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vdsbc.Type(), vdsbc.LocalName()))
	}
	return vdsbc.state
}

// Args contains the configurations for vault_database_secret_backend_connection.
type Args struct {
	// AllowedRoles: list of string, optional
	AllowedRoles terra.ListValue[terra.StringValue] `hcl:"allowed_roles,attr"`
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// Data: map of string, optional
	Data terra.MapValue[terra.StringValue] `hcl:"data,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// PluginName: string, optional
	PluginName terra.StringValue `hcl:"plugin_name,attr"`
	// RootRotationStatements: list of string, optional
	RootRotationStatements terra.ListValue[terra.StringValue] `hcl:"root_rotation_statements,attr"`
	// VerifyConnection: bool, optional
	VerifyConnection terra.BoolValue `hcl:"verify_connection,attr"`
	// Cassandra: optional
	Cassandra *Cassandra `hcl:"cassandra,block"`
	// Couchbase: optional
	Couchbase *Couchbase `hcl:"couchbase,block"`
	// Elasticsearch: optional
	Elasticsearch *Elasticsearch `hcl:"elasticsearch,block"`
	// Hana: optional
	Hana *Hana `hcl:"hana,block"`
	// Influxdb: optional
	Influxdb *Influxdb `hcl:"influxdb,block"`
	// Mongodb: optional
	Mongodb *Mongodb `hcl:"mongodb,block"`
	// Mongodbatlas: optional
	Mongodbatlas *Mongodbatlas `hcl:"mongodbatlas,block"`
	// Mssql: optional
	Mssql *Mssql `hcl:"mssql,block"`
	// Mysql: optional
	Mysql *Mysql `hcl:"mysql,block"`
	// MysqlAurora: optional
	MysqlAurora *MysqlAurora `hcl:"mysql_aurora,block"`
	// MysqlLegacy: optional
	MysqlLegacy *MysqlLegacy `hcl:"mysql_legacy,block"`
	// MysqlRds: optional
	MysqlRds *MysqlRds `hcl:"mysql_rds,block"`
	// Oracle: optional
	Oracle *Oracle `hcl:"oracle,block"`
	// Postgresql: optional
	Postgresql *Postgresql `hcl:"postgresql,block"`
	// Redis: optional
	Redis *Redis `hcl:"redis,block"`
	// RedisElasticache: optional
	RedisElasticache *RedisElasticache `hcl:"redis_elasticache,block"`
	// Redshift: optional
	Redshift *Redshift `hcl:"redshift,block"`
	// Snowflake: optional
	Snowflake *Snowflake `hcl:"snowflake,block"`
}

type vaultDatabaseSecretBackendConnectionAttributes struct {
	ref terra.Reference
}

// AllowedRoles returns a reference to field allowed_roles of vault_database_secret_backend_connection.
func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) AllowedRoles() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vdsbc.ref.Append("allowed_roles"))
}

// Backend returns a reference to field backend of vault_database_secret_backend_connection.
func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(vdsbc.ref.Append("backend"))
}

// Data returns a reference to field data of vault_database_secret_backend_connection.
func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Data() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vdsbc.ref.Append("data"))
}

// Id returns a reference to field id of vault_database_secret_backend_connection.
func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vdsbc.ref.Append("id"))
}

// Name returns a reference to field name of vault_database_secret_backend_connection.
func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vdsbc.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_database_secret_backend_connection.
func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vdsbc.ref.Append("namespace"))
}

// PluginName returns a reference to field plugin_name of vault_database_secret_backend_connection.
func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) PluginName() terra.StringValue {
	return terra.ReferenceAsString(vdsbc.ref.Append("plugin_name"))
}

// RootRotationStatements returns a reference to field root_rotation_statements of vault_database_secret_backend_connection.
func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) RootRotationStatements() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vdsbc.ref.Append("root_rotation_statements"))
}

// VerifyConnection returns a reference to field verify_connection of vault_database_secret_backend_connection.
func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) VerifyConnection() terra.BoolValue {
	return terra.ReferenceAsBool(vdsbc.ref.Append("verify_connection"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Cassandra() terra.ListValue[CassandraAttributes] {
	return terra.ReferenceAsList[CassandraAttributes](vdsbc.ref.Append("cassandra"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Couchbase() terra.ListValue[CouchbaseAttributes] {
	return terra.ReferenceAsList[CouchbaseAttributes](vdsbc.ref.Append("couchbase"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Elasticsearch() terra.ListValue[ElasticsearchAttributes] {
	return terra.ReferenceAsList[ElasticsearchAttributes](vdsbc.ref.Append("elasticsearch"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Hana() terra.ListValue[HanaAttributes] {
	return terra.ReferenceAsList[HanaAttributes](vdsbc.ref.Append("hana"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Influxdb() terra.ListValue[InfluxdbAttributes] {
	return terra.ReferenceAsList[InfluxdbAttributes](vdsbc.ref.Append("influxdb"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Mongodb() terra.ListValue[MongodbAttributes] {
	return terra.ReferenceAsList[MongodbAttributes](vdsbc.ref.Append("mongodb"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Mongodbatlas() terra.ListValue[MongodbatlasAttributes] {
	return terra.ReferenceAsList[MongodbatlasAttributes](vdsbc.ref.Append("mongodbatlas"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Mssql() terra.ListValue[MssqlAttributes] {
	return terra.ReferenceAsList[MssqlAttributes](vdsbc.ref.Append("mssql"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Mysql() terra.ListValue[MysqlAttributes] {
	return terra.ReferenceAsList[MysqlAttributes](vdsbc.ref.Append("mysql"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) MysqlAurora() terra.ListValue[MysqlAuroraAttributes] {
	return terra.ReferenceAsList[MysqlAuroraAttributes](vdsbc.ref.Append("mysql_aurora"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) MysqlLegacy() terra.ListValue[MysqlLegacyAttributes] {
	return terra.ReferenceAsList[MysqlLegacyAttributes](vdsbc.ref.Append("mysql_legacy"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) MysqlRds() terra.ListValue[MysqlRdsAttributes] {
	return terra.ReferenceAsList[MysqlRdsAttributes](vdsbc.ref.Append("mysql_rds"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Oracle() terra.ListValue[OracleAttributes] {
	return terra.ReferenceAsList[OracleAttributes](vdsbc.ref.Append("oracle"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Postgresql() terra.ListValue[PostgresqlAttributes] {
	return terra.ReferenceAsList[PostgresqlAttributes](vdsbc.ref.Append("postgresql"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Redis() terra.ListValue[RedisAttributes] {
	return terra.ReferenceAsList[RedisAttributes](vdsbc.ref.Append("redis"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) RedisElasticache() terra.ListValue[RedisElasticacheAttributes] {
	return terra.ReferenceAsList[RedisElasticacheAttributes](vdsbc.ref.Append("redis_elasticache"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Redshift() terra.ListValue[RedshiftAttributes] {
	return terra.ReferenceAsList[RedshiftAttributes](vdsbc.ref.Append("redshift"))
}

func (vdsbc vaultDatabaseSecretBackendConnectionAttributes) Snowflake() terra.ListValue[SnowflakeAttributes] {
	return terra.ReferenceAsList[SnowflakeAttributes](vdsbc.ref.Append("snowflake"))
}

type vaultDatabaseSecretBackendConnectionState struct {
	AllowedRoles           []string                `json:"allowed_roles"`
	Backend                string                  `json:"backend"`
	Data                   map[string]string       `json:"data"`
	Id                     string                  `json:"id"`
	Name                   string                  `json:"name"`
	Namespace              string                  `json:"namespace"`
	PluginName             string                  `json:"plugin_name"`
	RootRotationStatements []string                `json:"root_rotation_statements"`
	VerifyConnection       bool                    `json:"verify_connection"`
	Cassandra              []CassandraState        `json:"cassandra"`
	Couchbase              []CouchbaseState        `json:"couchbase"`
	Elasticsearch          []ElasticsearchState    `json:"elasticsearch"`
	Hana                   []HanaState             `json:"hana"`
	Influxdb               []InfluxdbState         `json:"influxdb"`
	Mongodb                []MongodbState          `json:"mongodb"`
	Mongodbatlas           []MongodbatlasState     `json:"mongodbatlas"`
	Mssql                  []MssqlState            `json:"mssql"`
	Mysql                  []MysqlState            `json:"mysql"`
	MysqlAurora            []MysqlAuroraState      `json:"mysql_aurora"`
	MysqlLegacy            []MysqlLegacyState      `json:"mysql_legacy"`
	MysqlRds               []MysqlRdsState         `json:"mysql_rds"`
	Oracle                 []OracleState           `json:"oracle"`
	Postgresql             []PostgresqlState       `json:"postgresql"`
	Redis                  []RedisState            `json:"redis"`
	RedisElasticache       []RedisElasticacheState `json:"redis_elasticache"`
	Redshift               []RedshiftState         `json:"redshift"`
	Snowflake              []SnowflakeState        `json:"snowflake"`
}
