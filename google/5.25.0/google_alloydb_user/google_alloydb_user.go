// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_alloydb_user

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

// Resource represents the Terraform resource google_alloydb_user.
type Resource struct {
	Name      string
	Args      Args
	state     *googleAlloydbUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gau *Resource) Type() string {
	return "google_alloydb_user"
}

// LocalName returns the local name for [Resource].
func (gau *Resource) LocalName() string {
	return gau.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gau *Resource) Configuration() interface{} {
	return gau.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gau *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gau)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gau *Resource) Dependencies() terra.Dependencies {
	return gau.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gau *Resource) LifecycleManagement() *terra.Lifecycle {
	return gau.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gau *Resource) Attributes() googleAlloydbUserAttributes {
	return googleAlloydbUserAttributes{ref: terra.ReferenceResource(gau)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gau *Resource) ImportState(state io.Reader) error {
	gau.state = &googleAlloydbUserState{}
	if err := json.NewDecoder(state).Decode(gau.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gau.Type(), gau.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gau *Resource) State() (*googleAlloydbUserState, bool) {
	return gau.state, gau.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gau *Resource) StateMust() *googleAlloydbUserState {
	if gau.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gau.Type(), gau.LocalName()))
	}
	return gau.state
}

// Args contains the configurations for google_alloydb_user.
type Args struct {
	// Cluster: string, required
	Cluster terra.StringValue `hcl:"cluster,attr" validate:"required"`
	// DatabaseRoles: list of string, optional
	DatabaseRoles terra.ListValue[terra.StringValue] `hcl:"database_roles,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// UserId: string, required
	UserId terra.StringValue `hcl:"user_id,attr" validate:"required"`
	// UserType: string, required
	UserType terra.StringValue `hcl:"user_type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleAlloydbUserAttributes struct {
	ref terra.Reference
}

// Cluster returns a reference to field cluster of google_alloydb_user.
func (gau googleAlloydbUserAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(gau.ref.Append("cluster"))
}

// DatabaseRoles returns a reference to field database_roles of google_alloydb_user.
func (gau googleAlloydbUserAttributes) DatabaseRoles() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gau.ref.Append("database_roles"))
}

// Id returns a reference to field id of google_alloydb_user.
func (gau googleAlloydbUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gau.ref.Append("id"))
}

// Name returns a reference to field name of google_alloydb_user.
func (gau googleAlloydbUserAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gau.ref.Append("name"))
}

// Password returns a reference to field password of google_alloydb_user.
func (gau googleAlloydbUserAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(gau.ref.Append("password"))
}

// UserId returns a reference to field user_id of google_alloydb_user.
func (gau googleAlloydbUserAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(gau.ref.Append("user_id"))
}

// UserType returns a reference to field user_type of google_alloydb_user.
func (gau googleAlloydbUserAttributes) UserType() terra.StringValue {
	return terra.ReferenceAsString(gau.ref.Append("user_type"))
}

func (gau googleAlloydbUserAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gau.ref.Append("timeouts"))
}

type googleAlloydbUserState struct {
	Cluster       string         `json:"cluster"`
	DatabaseRoles []string       `json:"database_roles"`
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Password      string         `json:"password"`
	UserId        string         `json:"user_id"`
	UserType      string         `json:"user_type"`
	Timeouts      *TimeoutsState `json:"timeouts"`
}
