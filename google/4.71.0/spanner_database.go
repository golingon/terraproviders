// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	spannerdatabase "github.com/golingon/terraproviders/google/4.71.0/spannerdatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpannerDatabase creates a new instance of [SpannerDatabase].
func NewSpannerDatabase(name string, args SpannerDatabaseArgs) *SpannerDatabase {
	return &SpannerDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpannerDatabase)(nil)

// SpannerDatabase represents the Terraform resource google_spanner_database.
type SpannerDatabase struct {
	Name      string
	Args      SpannerDatabaseArgs
	state     *spannerDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpannerDatabase].
func (sd *SpannerDatabase) Type() string {
	return "google_spanner_database"
}

// LocalName returns the local name for [SpannerDatabase].
func (sd *SpannerDatabase) LocalName() string {
	return sd.Name
}

// Configuration returns the configuration (args) for [SpannerDatabase].
func (sd *SpannerDatabase) Configuration() interface{} {
	return sd.Args
}

// DependOn is used for other resources to depend on [SpannerDatabase].
func (sd *SpannerDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(sd)
}

// Dependencies returns the list of resources [SpannerDatabase] depends_on.
func (sd *SpannerDatabase) Dependencies() terra.Dependencies {
	return sd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpannerDatabase].
func (sd *SpannerDatabase) LifecycleManagement() *terra.Lifecycle {
	return sd.Lifecycle
}

// Attributes returns the attributes for [SpannerDatabase].
func (sd *SpannerDatabase) Attributes() spannerDatabaseAttributes {
	return spannerDatabaseAttributes{ref: terra.ReferenceResource(sd)}
}

// ImportState imports the given attribute values into [SpannerDatabase]'s state.
func (sd *SpannerDatabase) ImportState(av io.Reader) error {
	sd.state = &spannerDatabaseState{}
	if err := json.NewDecoder(av).Decode(sd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sd.Type(), sd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpannerDatabase] has state.
func (sd *SpannerDatabase) State() (*spannerDatabaseState, bool) {
	return sd.state, sd.state != nil
}

// StateMust returns the state for [SpannerDatabase]. Panics if the state is nil.
func (sd *SpannerDatabase) StateMust() *spannerDatabaseState {
	if sd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sd.Type(), sd.LocalName()))
	}
	return sd.state
}

// SpannerDatabaseArgs contains the configurations for google_spanner_database.
type SpannerDatabaseArgs struct {
	// DatabaseDialect: string, optional
	DatabaseDialect terra.StringValue `hcl:"database_dialect,attr"`
	// Ddl: list of string, optional
	Ddl terra.ListValue[terra.StringValue] `hcl:"ddl,attr"`
	// DeletionProtection: bool, optional
	DeletionProtection terra.BoolValue `hcl:"deletion_protection,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// VersionRetentionPeriod: string, optional
	VersionRetentionPeriod terra.StringValue `hcl:"version_retention_period,attr"`
	// EncryptionConfig: optional
	EncryptionConfig *spannerdatabase.EncryptionConfig `hcl:"encryption_config,block"`
	// Timeouts: optional
	Timeouts *spannerdatabase.Timeouts `hcl:"timeouts,block"`
}
type spannerDatabaseAttributes struct {
	ref terra.Reference
}

// DatabaseDialect returns a reference to field database_dialect of google_spanner_database.
func (sd spannerDatabaseAttributes) DatabaseDialect() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("database_dialect"))
}

// Ddl returns a reference to field ddl of google_spanner_database.
func (sd spannerDatabaseAttributes) Ddl() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sd.ref.Append("ddl"))
}

// DeletionProtection returns a reference to field deletion_protection of google_spanner_database.
func (sd spannerDatabaseAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(sd.ref.Append("deletion_protection"))
}

// Id returns a reference to field id of google_spanner_database.
func (sd spannerDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("id"))
}

// Instance returns a reference to field instance of google_spanner_database.
func (sd spannerDatabaseAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("instance"))
}

// Name returns a reference to field name of google_spanner_database.
func (sd spannerDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("name"))
}

// Project returns a reference to field project of google_spanner_database.
func (sd spannerDatabaseAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("project"))
}

// State returns a reference to field state of google_spanner_database.
func (sd spannerDatabaseAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("state"))
}

// VersionRetentionPeriod returns a reference to field version_retention_period of google_spanner_database.
func (sd spannerDatabaseAttributes) VersionRetentionPeriod() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("version_retention_period"))
}

func (sd spannerDatabaseAttributes) EncryptionConfig() terra.ListValue[spannerdatabase.EncryptionConfigAttributes] {
	return terra.ReferenceAsList[spannerdatabase.EncryptionConfigAttributes](sd.ref.Append("encryption_config"))
}

func (sd spannerDatabaseAttributes) Timeouts() spannerdatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[spannerdatabase.TimeoutsAttributes](sd.ref.Append("timeouts"))
}

type spannerDatabaseState struct {
	DatabaseDialect        string                                  `json:"database_dialect"`
	Ddl                    []string                                `json:"ddl"`
	DeletionProtection     bool                                    `json:"deletion_protection"`
	Id                     string                                  `json:"id"`
	Instance               string                                  `json:"instance"`
	Name                   string                                  `json:"name"`
	Project                string                                  `json:"project"`
	State                  string                                  `json:"state"`
	VersionRetentionPeriod string                                  `json:"version_retention_period"`
	EncryptionConfig       []spannerdatabase.EncryptionConfigState `json:"encryption_config"`
	Timeouts               *spannerdatabase.TimeoutsState          `json:"timeouts"`
}
