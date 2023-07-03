// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	sqluser "github.com/golingon/terraproviders/googlebeta/4.71.0/sqluser"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlUser creates a new instance of [SqlUser].
func NewSqlUser(name string, args SqlUserArgs) *SqlUser {
	return &SqlUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlUser)(nil)

// SqlUser represents the Terraform resource google_sql_user.
type SqlUser struct {
	Name      string
	Args      SqlUserArgs
	state     *sqlUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlUser].
func (su *SqlUser) Type() string {
	return "google_sql_user"
}

// LocalName returns the local name for [SqlUser].
func (su *SqlUser) LocalName() string {
	return su.Name
}

// Configuration returns the configuration (args) for [SqlUser].
func (su *SqlUser) Configuration() interface{} {
	return su.Args
}

// DependOn is used for other resources to depend on [SqlUser].
func (su *SqlUser) DependOn() terra.Reference {
	return terra.ReferenceResource(su)
}

// Dependencies returns the list of resources [SqlUser] depends_on.
func (su *SqlUser) Dependencies() terra.Dependencies {
	return su.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlUser].
func (su *SqlUser) LifecycleManagement() *terra.Lifecycle {
	return su.Lifecycle
}

// Attributes returns the attributes for [SqlUser].
func (su *SqlUser) Attributes() sqlUserAttributes {
	return sqlUserAttributes{ref: terra.ReferenceResource(su)}
}

// ImportState imports the given attribute values into [SqlUser]'s state.
func (su *SqlUser) ImportState(av io.Reader) error {
	su.state = &sqlUserState{}
	if err := json.NewDecoder(av).Decode(su.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", su.Type(), su.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlUser] has state.
func (su *SqlUser) State() (*sqlUserState, bool) {
	return su.state, su.state != nil
}

// StateMust returns the state for [SqlUser]. Panics if the state is nil.
func (su *SqlUser) StateMust() *sqlUserState {
	if su.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", su.Type(), su.LocalName()))
	}
	return su.state
}

// SqlUserArgs contains the configurations for google_sql_user.
type SqlUserArgs struct {
	// DeletionPolicy: string, optional
	DeletionPolicy terra.StringValue `hcl:"deletion_policy,attr"`
	// Host: string, optional
	Host terra.StringValue `hcl:"host,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// SqlServerUserDetails: min=0
	SqlServerUserDetails []sqluser.SqlServerUserDetails `hcl:"sql_server_user_details,block" validate:"min=0"`
	// PasswordPolicy: optional
	PasswordPolicy *sqluser.PasswordPolicy `hcl:"password_policy,block"`
	// Timeouts: optional
	Timeouts *sqluser.Timeouts `hcl:"timeouts,block"`
}
type sqlUserAttributes struct {
	ref terra.Reference
}

// DeletionPolicy returns a reference to field deletion_policy of google_sql_user.
func (su sqlUserAttributes) DeletionPolicy() terra.StringValue {
	return terra.ReferenceAsString(su.ref.Append("deletion_policy"))
}

// Host returns a reference to field host of google_sql_user.
func (su sqlUserAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(su.ref.Append("host"))
}

// Id returns a reference to field id of google_sql_user.
func (su sqlUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(su.ref.Append("id"))
}

// Instance returns a reference to field instance of google_sql_user.
func (su sqlUserAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(su.ref.Append("instance"))
}

// Name returns a reference to field name of google_sql_user.
func (su sqlUserAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(su.ref.Append("name"))
}

// Password returns a reference to field password of google_sql_user.
func (su sqlUserAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(su.ref.Append("password"))
}

// Project returns a reference to field project of google_sql_user.
func (su sqlUserAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(su.ref.Append("project"))
}

// Type returns a reference to field type of google_sql_user.
func (su sqlUserAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(su.ref.Append("type"))
}

func (su sqlUserAttributes) SqlServerUserDetails() terra.ListValue[sqluser.SqlServerUserDetailsAttributes] {
	return terra.ReferenceAsList[sqluser.SqlServerUserDetailsAttributes](su.ref.Append("sql_server_user_details"))
}

func (su sqlUserAttributes) PasswordPolicy() terra.ListValue[sqluser.PasswordPolicyAttributes] {
	return terra.ReferenceAsList[sqluser.PasswordPolicyAttributes](su.ref.Append("password_policy"))
}

func (su sqlUserAttributes) Timeouts() sqluser.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqluser.TimeoutsAttributes](su.ref.Append("timeouts"))
}

type sqlUserState struct {
	DeletionPolicy       string                              `json:"deletion_policy"`
	Host                 string                              `json:"host"`
	Id                   string                              `json:"id"`
	Instance             string                              `json:"instance"`
	Name                 string                              `json:"name"`
	Password             string                              `json:"password"`
	Project              string                              `json:"project"`
	Type                 string                              `json:"type"`
	SqlServerUserDetails []sqluser.SqlServerUserDetailsState `json:"sql_server_user_details"`
	PasswordPolicy       []sqluser.PasswordPolicyState       `json:"password_policy"`
	Timeouts             *sqluser.TimeoutsState              `json:"timeouts"`
}
