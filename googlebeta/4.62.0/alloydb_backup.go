// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	alloydbbackup "github.com/golingon/terraproviders/googlebeta/4.62.0/alloydbbackup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAlloydbBackup creates a new instance of [AlloydbBackup].
func NewAlloydbBackup(name string, args AlloydbBackupArgs) *AlloydbBackup {
	return &AlloydbBackup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AlloydbBackup)(nil)

// AlloydbBackup represents the Terraform resource google_alloydb_backup.
type AlloydbBackup struct {
	Name      string
	Args      AlloydbBackupArgs
	state     *alloydbBackupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AlloydbBackup].
func (ab *AlloydbBackup) Type() string {
	return "google_alloydb_backup"
}

// LocalName returns the local name for [AlloydbBackup].
func (ab *AlloydbBackup) LocalName() string {
	return ab.Name
}

// Configuration returns the configuration (args) for [AlloydbBackup].
func (ab *AlloydbBackup) Configuration() interface{} {
	return ab.Args
}

// DependOn is used for other resources to depend on [AlloydbBackup].
func (ab *AlloydbBackup) DependOn() terra.Reference {
	return terra.ReferenceResource(ab)
}

// Dependencies returns the list of resources [AlloydbBackup] depends_on.
func (ab *AlloydbBackup) Dependencies() terra.Dependencies {
	return ab.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AlloydbBackup].
func (ab *AlloydbBackup) LifecycleManagement() *terra.Lifecycle {
	return ab.Lifecycle
}

// Attributes returns the attributes for [AlloydbBackup].
func (ab *AlloydbBackup) Attributes() alloydbBackupAttributes {
	return alloydbBackupAttributes{ref: terra.ReferenceResource(ab)}
}

// ImportState imports the given attribute values into [AlloydbBackup]'s state.
func (ab *AlloydbBackup) ImportState(av io.Reader) error {
	ab.state = &alloydbBackupState{}
	if err := json.NewDecoder(av).Decode(ab.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ab.Type(), ab.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AlloydbBackup] has state.
func (ab *AlloydbBackup) State() (*alloydbBackupState, bool) {
	return ab.state, ab.state != nil
}

// StateMust returns the state for [AlloydbBackup]. Panics if the state is nil.
func (ab *AlloydbBackup) StateMust() *alloydbBackupState {
	if ab.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ab.Type(), ab.LocalName()))
	}
	return ab.state
}

// AlloydbBackupArgs contains the configurations for google_alloydb_backup.
type AlloydbBackupArgs struct {
	// BackupId: string, required
	BackupId terra.StringValue `hcl:"backup_id,attr" validate:"required"`
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *alloydbbackup.Timeouts `hcl:"timeouts,block"`
}
type alloydbBackupAttributes struct {
	ref terra.Reference
}

// BackupId returns a reference to field backup_id of google_alloydb_backup.
func (ab alloydbBackupAttributes) BackupId() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("backup_id"))
}

// ClusterName returns a reference to field cluster_name of google_alloydb_backup.
func (ab alloydbBackupAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("cluster_name"))
}

// CreateTime returns a reference to field create_time of google_alloydb_backup.
func (ab alloydbBackupAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("create_time"))
}

// Description returns a reference to field description of google_alloydb_backup.
func (ab alloydbBackupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("description"))
}

// Etag returns a reference to field etag of google_alloydb_backup.
func (ab alloydbBackupAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("etag"))
}

// Id returns a reference to field id of google_alloydb_backup.
func (ab alloydbBackupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("id"))
}

// Labels returns a reference to field labels of google_alloydb_backup.
func (ab alloydbBackupAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ab.ref.Append("labels"))
}

// Location returns a reference to field location of google_alloydb_backup.
func (ab alloydbBackupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("location"))
}

// Name returns a reference to field name of google_alloydb_backup.
func (ab alloydbBackupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("name"))
}

// Project returns a reference to field project of google_alloydb_backup.
func (ab alloydbBackupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_alloydb_backup.
func (ab alloydbBackupAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(ab.ref.Append("reconciling"))
}

// State returns a reference to field state of google_alloydb_backup.
func (ab alloydbBackupAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("state"))
}

// Uid returns a reference to field uid of google_alloydb_backup.
func (ab alloydbBackupAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_alloydb_backup.
func (ab alloydbBackupAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("update_time"))
}

func (ab alloydbBackupAttributes) Timeouts() alloydbbackup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[alloydbbackup.TimeoutsAttributes](ab.ref.Append("timeouts"))
}

type alloydbBackupState struct {
	BackupId    string                       `json:"backup_id"`
	ClusterName string                       `json:"cluster_name"`
	CreateTime  string                       `json:"create_time"`
	Description string                       `json:"description"`
	Etag        string                       `json:"etag"`
	Id          string                       `json:"id"`
	Labels      map[string]string            `json:"labels"`
	Location    string                       `json:"location"`
	Name        string                       `json:"name"`
	Project     string                       `json:"project"`
	Reconciling bool                         `json:"reconciling"`
	State       string                       `json:"state"`
	Uid         string                       `json:"uid"`
	UpdateTime  string                       `json:"update_time"`
	Timeouts    *alloydbbackup.TimeoutsState `json:"timeouts"`
}
