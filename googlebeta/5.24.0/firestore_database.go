// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	firestoredatabase "github.com/golingon/terraproviders/googlebeta/5.24.0/firestoredatabase"
	"io"
)

// NewFirestoreDatabase creates a new instance of [FirestoreDatabase].
func NewFirestoreDatabase(name string, args FirestoreDatabaseArgs) *FirestoreDatabase {
	return &FirestoreDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirestoreDatabase)(nil)

// FirestoreDatabase represents the Terraform resource google_firestore_database.
type FirestoreDatabase struct {
	Name      string
	Args      FirestoreDatabaseArgs
	state     *firestoreDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirestoreDatabase].
func (fd *FirestoreDatabase) Type() string {
	return "google_firestore_database"
}

// LocalName returns the local name for [FirestoreDatabase].
func (fd *FirestoreDatabase) LocalName() string {
	return fd.Name
}

// Configuration returns the configuration (args) for [FirestoreDatabase].
func (fd *FirestoreDatabase) Configuration() interface{} {
	return fd.Args
}

// DependOn is used for other resources to depend on [FirestoreDatabase].
func (fd *FirestoreDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(fd)
}

// Dependencies returns the list of resources [FirestoreDatabase] depends_on.
func (fd *FirestoreDatabase) Dependencies() terra.Dependencies {
	return fd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirestoreDatabase].
func (fd *FirestoreDatabase) LifecycleManagement() *terra.Lifecycle {
	return fd.Lifecycle
}

// Attributes returns the attributes for [FirestoreDatabase].
func (fd *FirestoreDatabase) Attributes() firestoreDatabaseAttributes {
	return firestoreDatabaseAttributes{ref: terra.ReferenceResource(fd)}
}

// ImportState imports the given attribute values into [FirestoreDatabase]'s state.
func (fd *FirestoreDatabase) ImportState(av io.Reader) error {
	fd.state = &firestoreDatabaseState{}
	if err := json.NewDecoder(av).Decode(fd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fd.Type(), fd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirestoreDatabase] has state.
func (fd *FirestoreDatabase) State() (*firestoreDatabaseState, bool) {
	return fd.state, fd.state != nil
}

// StateMust returns the state for [FirestoreDatabase]. Panics if the state is nil.
func (fd *FirestoreDatabase) StateMust() *firestoreDatabaseState {
	if fd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fd.Type(), fd.LocalName()))
	}
	return fd.state
}

// FirestoreDatabaseArgs contains the configurations for google_firestore_database.
type FirestoreDatabaseArgs struct {
	// AppEngineIntegrationMode: string, optional
	AppEngineIntegrationMode terra.StringValue `hcl:"app_engine_integration_mode,attr"`
	// ConcurrencyMode: string, optional
	ConcurrencyMode terra.StringValue `hcl:"concurrency_mode,attr"`
	// DeleteProtectionState: string, optional
	DeleteProtectionState terra.StringValue `hcl:"delete_protection_state,attr"`
	// DeletionPolicy: string, optional
	DeletionPolicy terra.StringValue `hcl:"deletion_policy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocationId: string, required
	LocationId terra.StringValue `hcl:"location_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PointInTimeRecoveryEnablement: string, optional
	PointInTimeRecoveryEnablement terra.StringValue `hcl:"point_in_time_recovery_enablement,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// CmekConfig: optional
	CmekConfig *firestoredatabase.CmekConfig `hcl:"cmek_config,block"`
	// Timeouts: optional
	Timeouts *firestoredatabase.Timeouts `hcl:"timeouts,block"`
}
type firestoreDatabaseAttributes struct {
	ref terra.Reference
}

// AppEngineIntegrationMode returns a reference to field app_engine_integration_mode of google_firestore_database.
func (fd firestoreDatabaseAttributes) AppEngineIntegrationMode() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("app_engine_integration_mode"))
}

// ConcurrencyMode returns a reference to field concurrency_mode of google_firestore_database.
func (fd firestoreDatabaseAttributes) ConcurrencyMode() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("concurrency_mode"))
}

// CreateTime returns a reference to field create_time of google_firestore_database.
func (fd firestoreDatabaseAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("create_time"))
}

// DeleteProtectionState returns a reference to field delete_protection_state of google_firestore_database.
func (fd firestoreDatabaseAttributes) DeleteProtectionState() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("delete_protection_state"))
}

// DeletionPolicy returns a reference to field deletion_policy of google_firestore_database.
func (fd firestoreDatabaseAttributes) DeletionPolicy() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("deletion_policy"))
}

// EarliestVersionTime returns a reference to field earliest_version_time of google_firestore_database.
func (fd firestoreDatabaseAttributes) EarliestVersionTime() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("earliest_version_time"))
}

// Etag returns a reference to field etag of google_firestore_database.
func (fd firestoreDatabaseAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("etag"))
}

// Id returns a reference to field id of google_firestore_database.
func (fd firestoreDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("id"))
}

// KeyPrefix returns a reference to field key_prefix of google_firestore_database.
func (fd firestoreDatabaseAttributes) KeyPrefix() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("key_prefix"))
}

// LocationId returns a reference to field location_id of google_firestore_database.
func (fd firestoreDatabaseAttributes) LocationId() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("location_id"))
}

// Name returns a reference to field name of google_firestore_database.
func (fd firestoreDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("name"))
}

// PointInTimeRecoveryEnablement returns a reference to field point_in_time_recovery_enablement of google_firestore_database.
func (fd firestoreDatabaseAttributes) PointInTimeRecoveryEnablement() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("point_in_time_recovery_enablement"))
}

// Project returns a reference to field project of google_firestore_database.
func (fd firestoreDatabaseAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("project"))
}

// Type returns a reference to field type of google_firestore_database.
func (fd firestoreDatabaseAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("type"))
}

// Uid returns a reference to field uid of google_firestore_database.
func (fd firestoreDatabaseAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_firestore_database.
func (fd firestoreDatabaseAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("update_time"))
}

// VersionRetentionPeriod returns a reference to field version_retention_period of google_firestore_database.
func (fd firestoreDatabaseAttributes) VersionRetentionPeriod() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("version_retention_period"))
}

func (fd firestoreDatabaseAttributes) CmekConfig() terra.ListValue[firestoredatabase.CmekConfigAttributes] {
	return terra.ReferenceAsList[firestoredatabase.CmekConfigAttributes](fd.ref.Append("cmek_config"))
}

func (fd firestoreDatabaseAttributes) Timeouts() firestoredatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firestoredatabase.TimeoutsAttributes](fd.ref.Append("timeouts"))
}

type firestoreDatabaseState struct {
	AppEngineIntegrationMode      string                              `json:"app_engine_integration_mode"`
	ConcurrencyMode               string                              `json:"concurrency_mode"`
	CreateTime                    string                              `json:"create_time"`
	DeleteProtectionState         string                              `json:"delete_protection_state"`
	DeletionPolicy                string                              `json:"deletion_policy"`
	EarliestVersionTime           string                              `json:"earliest_version_time"`
	Etag                          string                              `json:"etag"`
	Id                            string                              `json:"id"`
	KeyPrefix                     string                              `json:"key_prefix"`
	LocationId                    string                              `json:"location_id"`
	Name                          string                              `json:"name"`
	PointInTimeRecoveryEnablement string                              `json:"point_in_time_recovery_enablement"`
	Project                       string                              `json:"project"`
	Type                          string                              `json:"type"`
	Uid                           string                              `json:"uid"`
	UpdateTime                    string                              `json:"update_time"`
	VersionRetentionPeriod        string                              `json:"version_retention_period"`
	CmekConfig                    []firestoredatabase.CmekConfigState `json:"cmek_config"`
	Timeouts                      *firestoredatabase.TimeoutsState    `json:"timeouts"`
}
