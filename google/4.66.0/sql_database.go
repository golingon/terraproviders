// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	sqldatabase "github.com/golingon/terraproviders/google/4.66.0/sqldatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlDatabase creates a new instance of [SqlDatabase].
func NewSqlDatabase(name string, args SqlDatabaseArgs) *SqlDatabase {
	return &SqlDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlDatabase)(nil)

// SqlDatabase represents the Terraform resource google_sql_database.
type SqlDatabase struct {
	Name      string
	Args      SqlDatabaseArgs
	state     *sqlDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlDatabase].
func (sd *SqlDatabase) Type() string {
	return "google_sql_database"
}

// LocalName returns the local name for [SqlDatabase].
func (sd *SqlDatabase) LocalName() string {
	return sd.Name
}

// Configuration returns the configuration (args) for [SqlDatabase].
func (sd *SqlDatabase) Configuration() interface{} {
	return sd.Args
}

// DependOn is used for other resources to depend on [SqlDatabase].
func (sd *SqlDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(sd)
}

// Dependencies returns the list of resources [SqlDatabase] depends_on.
func (sd *SqlDatabase) Dependencies() terra.Dependencies {
	return sd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlDatabase].
func (sd *SqlDatabase) LifecycleManagement() *terra.Lifecycle {
	return sd.Lifecycle
}

// Attributes returns the attributes for [SqlDatabase].
func (sd *SqlDatabase) Attributes() sqlDatabaseAttributes {
	return sqlDatabaseAttributes{ref: terra.ReferenceResource(sd)}
}

// ImportState imports the given attribute values into [SqlDatabase]'s state.
func (sd *SqlDatabase) ImportState(av io.Reader) error {
	sd.state = &sqlDatabaseState{}
	if err := json.NewDecoder(av).Decode(sd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sd.Type(), sd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlDatabase] has state.
func (sd *SqlDatabase) State() (*sqlDatabaseState, bool) {
	return sd.state, sd.state != nil
}

// StateMust returns the state for [SqlDatabase]. Panics if the state is nil.
func (sd *SqlDatabase) StateMust() *sqlDatabaseState {
	if sd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sd.Type(), sd.LocalName()))
	}
	return sd.state
}

// SqlDatabaseArgs contains the configurations for google_sql_database.
type SqlDatabaseArgs struct {
	// Charset: string, optional
	Charset terra.StringValue `hcl:"charset,attr"`
	// Collation: string, optional
	Collation terra.StringValue `hcl:"collation,attr"`
	// DeletionPolicy: string, optional
	DeletionPolicy terra.StringValue `hcl:"deletion_policy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *sqldatabase.Timeouts `hcl:"timeouts,block"`
}
type sqlDatabaseAttributes struct {
	ref terra.Reference
}

// Charset returns a reference to field charset of google_sql_database.
func (sd sqlDatabaseAttributes) Charset() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("charset"))
}

// Collation returns a reference to field collation of google_sql_database.
func (sd sqlDatabaseAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("collation"))
}

// DeletionPolicy returns a reference to field deletion_policy of google_sql_database.
func (sd sqlDatabaseAttributes) DeletionPolicy() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("deletion_policy"))
}

// Id returns a reference to field id of google_sql_database.
func (sd sqlDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("id"))
}

// Instance returns a reference to field instance of google_sql_database.
func (sd sqlDatabaseAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("instance"))
}

// Name returns a reference to field name of google_sql_database.
func (sd sqlDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("name"))
}

// Project returns a reference to field project of google_sql_database.
func (sd sqlDatabaseAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_sql_database.
func (sd sqlDatabaseAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("self_link"))
}

func (sd sqlDatabaseAttributes) Timeouts() sqldatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqldatabase.TimeoutsAttributes](sd.ref.Append("timeouts"))
}

type sqlDatabaseState struct {
	Charset        string                     `json:"charset"`
	Collation      string                     `json:"collation"`
	DeletionPolicy string                     `json:"deletion_policy"`
	Id             string                     `json:"id"`
	Instance       string                     `json:"instance"`
	Name           string                     `json:"name"`
	Project        string                     `json:"project"`
	SelfLink       string                     `json:"self_link"`
	Timeouts       *sqldatabase.TimeoutsState `json:"timeouts"`
}
