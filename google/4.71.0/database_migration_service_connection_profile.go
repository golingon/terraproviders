// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	databasemigrationserviceconnectionprofile "github.com/golingon/terraproviders/google/4.71.0/databasemigrationserviceconnectionprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatabaseMigrationServiceConnectionProfile creates a new instance of [DatabaseMigrationServiceConnectionProfile].
func NewDatabaseMigrationServiceConnectionProfile(name string, args DatabaseMigrationServiceConnectionProfileArgs) *DatabaseMigrationServiceConnectionProfile {
	return &DatabaseMigrationServiceConnectionProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatabaseMigrationServiceConnectionProfile)(nil)

// DatabaseMigrationServiceConnectionProfile represents the Terraform resource google_database_migration_service_connection_profile.
type DatabaseMigrationServiceConnectionProfile struct {
	Name      string
	Args      DatabaseMigrationServiceConnectionProfileArgs
	state     *databaseMigrationServiceConnectionProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatabaseMigrationServiceConnectionProfile].
func (dmscp *DatabaseMigrationServiceConnectionProfile) Type() string {
	return "google_database_migration_service_connection_profile"
}

// LocalName returns the local name for [DatabaseMigrationServiceConnectionProfile].
func (dmscp *DatabaseMigrationServiceConnectionProfile) LocalName() string {
	return dmscp.Name
}

// Configuration returns the configuration (args) for [DatabaseMigrationServiceConnectionProfile].
func (dmscp *DatabaseMigrationServiceConnectionProfile) Configuration() interface{} {
	return dmscp.Args
}

// DependOn is used for other resources to depend on [DatabaseMigrationServiceConnectionProfile].
func (dmscp *DatabaseMigrationServiceConnectionProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(dmscp)
}

// Dependencies returns the list of resources [DatabaseMigrationServiceConnectionProfile] depends_on.
func (dmscp *DatabaseMigrationServiceConnectionProfile) Dependencies() terra.Dependencies {
	return dmscp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatabaseMigrationServiceConnectionProfile].
func (dmscp *DatabaseMigrationServiceConnectionProfile) LifecycleManagement() *terra.Lifecycle {
	return dmscp.Lifecycle
}

// Attributes returns the attributes for [DatabaseMigrationServiceConnectionProfile].
func (dmscp *DatabaseMigrationServiceConnectionProfile) Attributes() databaseMigrationServiceConnectionProfileAttributes {
	return databaseMigrationServiceConnectionProfileAttributes{ref: terra.ReferenceResource(dmscp)}
}

// ImportState imports the given attribute values into [DatabaseMigrationServiceConnectionProfile]'s state.
func (dmscp *DatabaseMigrationServiceConnectionProfile) ImportState(av io.Reader) error {
	dmscp.state = &databaseMigrationServiceConnectionProfileState{}
	if err := json.NewDecoder(av).Decode(dmscp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmscp.Type(), dmscp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatabaseMigrationServiceConnectionProfile] has state.
func (dmscp *DatabaseMigrationServiceConnectionProfile) State() (*databaseMigrationServiceConnectionProfileState, bool) {
	return dmscp.state, dmscp.state != nil
}

// StateMust returns the state for [DatabaseMigrationServiceConnectionProfile]. Panics if the state is nil.
func (dmscp *DatabaseMigrationServiceConnectionProfile) StateMust() *databaseMigrationServiceConnectionProfileState {
	if dmscp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmscp.Type(), dmscp.LocalName()))
	}
	return dmscp.state
}

// DatabaseMigrationServiceConnectionProfileArgs contains the configurations for google_database_migration_service_connection_profile.
type DatabaseMigrationServiceConnectionProfileArgs struct {
	// ConnectionProfileId: string, required
	ConnectionProfileId terra.StringValue `hcl:"connection_profile_id,attr" validate:"required"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Error: min=0
	Error []databasemigrationserviceconnectionprofile.Error `hcl:"error,block" validate:"min=0"`
	// Alloydb: optional
	Alloydb *databasemigrationserviceconnectionprofile.Alloydb `hcl:"alloydb,block"`
	// Cloudsql: optional
	Cloudsql *databasemigrationserviceconnectionprofile.Cloudsql `hcl:"cloudsql,block"`
	// Mysql: optional
	Mysql *databasemigrationserviceconnectionprofile.Mysql `hcl:"mysql,block"`
	// Postgresql: optional
	Postgresql *databasemigrationserviceconnectionprofile.Postgresql `hcl:"postgresql,block"`
	// Timeouts: optional
	Timeouts *databasemigrationserviceconnectionprofile.Timeouts `hcl:"timeouts,block"`
}
type databaseMigrationServiceConnectionProfileAttributes struct {
	ref terra.Reference
}

// ConnectionProfileId returns a reference to field connection_profile_id of google_database_migration_service_connection_profile.
func (dmscp databaseMigrationServiceConnectionProfileAttributes) ConnectionProfileId() terra.StringValue {
	return terra.ReferenceAsString(dmscp.ref.Append("connection_profile_id"))
}

// CreateTime returns a reference to field create_time of google_database_migration_service_connection_profile.
func (dmscp databaseMigrationServiceConnectionProfileAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(dmscp.ref.Append("create_time"))
}

// Dbprovider returns a reference to field dbprovider of google_database_migration_service_connection_profile.
func (dmscp databaseMigrationServiceConnectionProfileAttributes) Dbprovider() terra.StringValue {
	return terra.ReferenceAsString(dmscp.ref.Append("dbprovider"))
}

// DisplayName returns a reference to field display_name of google_database_migration_service_connection_profile.
func (dmscp databaseMigrationServiceConnectionProfileAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dmscp.ref.Append("display_name"))
}

// Id returns a reference to field id of google_database_migration_service_connection_profile.
func (dmscp databaseMigrationServiceConnectionProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmscp.ref.Append("id"))
}

// Labels returns a reference to field labels of google_database_migration_service_connection_profile.
func (dmscp databaseMigrationServiceConnectionProfileAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dmscp.ref.Append("labels"))
}

// Location returns a reference to field location of google_database_migration_service_connection_profile.
func (dmscp databaseMigrationServiceConnectionProfileAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dmscp.ref.Append("location"))
}

// Name returns a reference to field name of google_database_migration_service_connection_profile.
func (dmscp databaseMigrationServiceConnectionProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dmscp.ref.Append("name"))
}

// Project returns a reference to field project of google_database_migration_service_connection_profile.
func (dmscp databaseMigrationServiceConnectionProfileAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmscp.ref.Append("project"))
}

// State returns a reference to field state of google_database_migration_service_connection_profile.
func (dmscp databaseMigrationServiceConnectionProfileAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(dmscp.ref.Append("state"))
}

func (dmscp databaseMigrationServiceConnectionProfileAttributes) Error() terra.ListValue[databasemigrationserviceconnectionprofile.ErrorAttributes] {
	return terra.ReferenceAsList[databasemigrationserviceconnectionprofile.ErrorAttributes](dmscp.ref.Append("error"))
}

func (dmscp databaseMigrationServiceConnectionProfileAttributes) Alloydb() terra.ListValue[databasemigrationserviceconnectionprofile.AlloydbAttributes] {
	return terra.ReferenceAsList[databasemigrationserviceconnectionprofile.AlloydbAttributes](dmscp.ref.Append("alloydb"))
}

func (dmscp databaseMigrationServiceConnectionProfileAttributes) Cloudsql() terra.ListValue[databasemigrationserviceconnectionprofile.CloudsqlAttributes] {
	return terra.ReferenceAsList[databasemigrationserviceconnectionprofile.CloudsqlAttributes](dmscp.ref.Append("cloudsql"))
}

func (dmscp databaseMigrationServiceConnectionProfileAttributes) Mysql() terra.ListValue[databasemigrationserviceconnectionprofile.MysqlAttributes] {
	return terra.ReferenceAsList[databasemigrationserviceconnectionprofile.MysqlAttributes](dmscp.ref.Append("mysql"))
}

func (dmscp databaseMigrationServiceConnectionProfileAttributes) Postgresql() terra.ListValue[databasemigrationserviceconnectionprofile.PostgresqlAttributes] {
	return terra.ReferenceAsList[databasemigrationserviceconnectionprofile.PostgresqlAttributes](dmscp.ref.Append("postgresql"))
}

func (dmscp databaseMigrationServiceConnectionProfileAttributes) Timeouts() databasemigrationserviceconnectionprofile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databasemigrationserviceconnectionprofile.TimeoutsAttributes](dmscp.ref.Append("timeouts"))
}

type databaseMigrationServiceConnectionProfileState struct {
	ConnectionProfileId string                                                      `json:"connection_profile_id"`
	CreateTime          string                                                      `json:"create_time"`
	Dbprovider          string                                                      `json:"dbprovider"`
	DisplayName         string                                                      `json:"display_name"`
	Id                  string                                                      `json:"id"`
	Labels              map[string]string                                           `json:"labels"`
	Location            string                                                      `json:"location"`
	Name                string                                                      `json:"name"`
	Project             string                                                      `json:"project"`
	State               string                                                      `json:"state"`
	Error               []databasemigrationserviceconnectionprofile.ErrorState      `json:"error"`
	Alloydb             []databasemigrationserviceconnectionprofile.AlloydbState    `json:"alloydb"`
	Cloudsql            []databasemigrationserviceconnectionprofile.CloudsqlState   `json:"cloudsql"`
	Mysql               []databasemigrationserviceconnectionprofile.MysqlState      `json:"mysql"`
	Postgresql          []databasemigrationserviceconnectionprofile.PostgresqlState `json:"postgresql"`
	Timeouts            *databasemigrationserviceconnectionprofile.TimeoutsState    `json:"timeouts"`
}
