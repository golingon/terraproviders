// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	alloydbbackup "github.com/golingon/terraproviders/google/5.24.0/alloydbbackup"
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
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// BackupId: string, required
	BackupId terra.StringValue `hcl:"backup_id,attr" validate:"required"`
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// EncryptionInfo: min=0
	EncryptionInfo []alloydbbackup.EncryptionInfo `hcl:"encryption_info,block" validate:"min=0"`
	// ExpiryQuantity: min=0
	ExpiryQuantity []alloydbbackup.ExpiryQuantity `hcl:"expiry_quantity,block" validate:"min=0"`
	// EncryptionConfig: optional
	EncryptionConfig *alloydbbackup.EncryptionConfig `hcl:"encryption_config,block"`
	// Timeouts: optional
	Timeouts *alloydbbackup.Timeouts `hcl:"timeouts,block"`
}
type alloydbBackupAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_alloydb_backup.
func (ab alloydbBackupAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ab.ref.Append("annotations"))
}

// BackupId returns a reference to field backup_id of google_alloydb_backup.
func (ab alloydbBackupAttributes) BackupId() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("backup_id"))
}

// ClusterName returns a reference to field cluster_name of google_alloydb_backup.
func (ab alloydbBackupAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("cluster_name"))
}

// ClusterUid returns a reference to field cluster_uid of google_alloydb_backup.
func (ab alloydbBackupAttributes) ClusterUid() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("cluster_uid"))
}

// CreateTime returns a reference to field create_time of google_alloydb_backup.
func (ab alloydbBackupAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("create_time"))
}

// DeleteTime returns a reference to field delete_time of google_alloydb_backup.
func (ab alloydbBackupAttributes) DeleteTime() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("delete_time"))
}

// Description returns a reference to field description of google_alloydb_backup.
func (ab alloydbBackupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_alloydb_backup.
func (ab alloydbBackupAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("display_name"))
}

// EffectiveAnnotations returns a reference to field effective_annotations of google_alloydb_backup.
func (ab alloydbBackupAttributes) EffectiveAnnotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ab.ref.Append("effective_annotations"))
}

// EffectiveLabels returns a reference to field effective_labels of google_alloydb_backup.
func (ab alloydbBackupAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ab.ref.Append("effective_labels"))
}

// Etag returns a reference to field etag of google_alloydb_backup.
func (ab alloydbBackupAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("etag"))
}

// ExpiryTime returns a reference to field expiry_time of google_alloydb_backup.
func (ab alloydbBackupAttributes) ExpiryTime() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("expiry_time"))
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

// SizeBytes returns a reference to field size_bytes of google_alloydb_backup.
func (ab alloydbBackupAttributes) SizeBytes() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("size_bytes"))
}

// State returns a reference to field state of google_alloydb_backup.
func (ab alloydbBackupAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("state"))
}

// TerraformLabels returns a reference to field terraform_labels of google_alloydb_backup.
func (ab alloydbBackupAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ab.ref.Append("terraform_labels"))
}

// Type returns a reference to field type of google_alloydb_backup.
func (ab alloydbBackupAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("type"))
}

// Uid returns a reference to field uid of google_alloydb_backup.
func (ab alloydbBackupAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_alloydb_backup.
func (ab alloydbBackupAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("update_time"))
}

func (ab alloydbBackupAttributes) EncryptionInfo() terra.ListValue[alloydbbackup.EncryptionInfoAttributes] {
	return terra.ReferenceAsList[alloydbbackup.EncryptionInfoAttributes](ab.ref.Append("encryption_info"))
}

func (ab alloydbBackupAttributes) ExpiryQuantity() terra.ListValue[alloydbbackup.ExpiryQuantityAttributes] {
	return terra.ReferenceAsList[alloydbbackup.ExpiryQuantityAttributes](ab.ref.Append("expiry_quantity"))
}

func (ab alloydbBackupAttributes) EncryptionConfig() terra.ListValue[alloydbbackup.EncryptionConfigAttributes] {
	return terra.ReferenceAsList[alloydbbackup.EncryptionConfigAttributes](ab.ref.Append("encryption_config"))
}

func (ab alloydbBackupAttributes) Timeouts() alloydbbackup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[alloydbbackup.TimeoutsAttributes](ab.ref.Append("timeouts"))
}

type alloydbBackupState struct {
	Annotations          map[string]string                     `json:"annotations"`
	BackupId             string                                `json:"backup_id"`
	ClusterName          string                                `json:"cluster_name"`
	ClusterUid           string                                `json:"cluster_uid"`
	CreateTime           string                                `json:"create_time"`
	DeleteTime           string                                `json:"delete_time"`
	Description          string                                `json:"description"`
	DisplayName          string                                `json:"display_name"`
	EffectiveAnnotations map[string]string                     `json:"effective_annotations"`
	EffectiveLabels      map[string]string                     `json:"effective_labels"`
	Etag                 string                                `json:"etag"`
	ExpiryTime           string                                `json:"expiry_time"`
	Id                   string                                `json:"id"`
	Labels               map[string]string                     `json:"labels"`
	Location             string                                `json:"location"`
	Name                 string                                `json:"name"`
	Project              string                                `json:"project"`
	Reconciling          bool                                  `json:"reconciling"`
	SizeBytes            string                                `json:"size_bytes"`
	State                string                                `json:"state"`
	TerraformLabels      map[string]string                     `json:"terraform_labels"`
	Type                 string                                `json:"type"`
	Uid                  string                                `json:"uid"`
	UpdateTime           string                                `json:"update_time"`
	EncryptionInfo       []alloydbbackup.EncryptionInfoState   `json:"encryption_info"`
	ExpiryQuantity       []alloydbbackup.ExpiryQuantityState   `json:"expiry_quantity"`
	EncryptionConfig     []alloydbbackup.EncryptionConfigState `json:"encryption_config"`
	Timeouts             *alloydbbackup.TimeoutsState          `json:"timeouts"`
}
