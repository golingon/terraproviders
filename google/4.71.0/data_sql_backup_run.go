// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataSqlBackupRun creates a new instance of [DataSqlBackupRun].
func NewDataSqlBackupRun(name string, args DataSqlBackupRunArgs) *DataSqlBackupRun {
	return &DataSqlBackupRun{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSqlBackupRun)(nil)

// DataSqlBackupRun represents the Terraform data resource google_sql_backup_run.
type DataSqlBackupRun struct {
	Name string
	Args DataSqlBackupRunArgs
}

// DataSource returns the Terraform object type for [DataSqlBackupRun].
func (sbr *DataSqlBackupRun) DataSource() string {
	return "google_sql_backup_run"
}

// LocalName returns the local name for [DataSqlBackupRun].
func (sbr *DataSqlBackupRun) LocalName() string {
	return sbr.Name
}

// Configuration returns the configuration (args) for [DataSqlBackupRun].
func (sbr *DataSqlBackupRun) Configuration() interface{} {
	return sbr.Args
}

// Attributes returns the attributes for [DataSqlBackupRun].
func (sbr *DataSqlBackupRun) Attributes() dataSqlBackupRunAttributes {
	return dataSqlBackupRunAttributes{ref: terra.ReferenceDataResource(sbr)}
}

// DataSqlBackupRunArgs contains the configurations for google_sql_backup_run.
type DataSqlBackupRunArgs struct {
	// BackupId: number, optional
	BackupId terra.NumberValue `hcl:"backup_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// MostRecent: bool, optional
	MostRecent terra.BoolValue `hcl:"most_recent,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataSqlBackupRunAttributes struct {
	ref terra.Reference
}

// BackupId returns a reference to field backup_id of google_sql_backup_run.
func (sbr dataSqlBackupRunAttributes) BackupId() terra.NumberValue {
	return terra.ReferenceAsNumber(sbr.ref.Append("backup_id"))
}

// Id returns a reference to field id of google_sql_backup_run.
func (sbr dataSqlBackupRunAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbr.ref.Append("id"))
}

// Instance returns a reference to field instance of google_sql_backup_run.
func (sbr dataSqlBackupRunAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(sbr.ref.Append("instance"))
}

// Location returns a reference to field location of google_sql_backup_run.
func (sbr dataSqlBackupRunAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sbr.ref.Append("location"))
}

// MostRecent returns a reference to field most_recent of google_sql_backup_run.
func (sbr dataSqlBackupRunAttributes) MostRecent() terra.BoolValue {
	return terra.ReferenceAsBool(sbr.ref.Append("most_recent"))
}

// Project returns a reference to field project of google_sql_backup_run.
func (sbr dataSqlBackupRunAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sbr.ref.Append("project"))
}

// StartTime returns a reference to field start_time of google_sql_backup_run.
func (sbr dataSqlBackupRunAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(sbr.ref.Append("start_time"))
}

// Status returns a reference to field status of google_sql_backup_run.
func (sbr dataSqlBackupRunAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(sbr.ref.Append("status"))
}
