// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	alloydbcluster "github.com/golingon/terraproviders/google/4.63.1/alloydbcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAlloydbCluster creates a new instance of [AlloydbCluster].
func NewAlloydbCluster(name string, args AlloydbClusterArgs) *AlloydbCluster {
	return &AlloydbCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AlloydbCluster)(nil)

// AlloydbCluster represents the Terraform resource google_alloydb_cluster.
type AlloydbCluster struct {
	Name      string
	Args      AlloydbClusterArgs
	state     *alloydbClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AlloydbCluster].
func (ac *AlloydbCluster) Type() string {
	return "google_alloydb_cluster"
}

// LocalName returns the local name for [AlloydbCluster].
func (ac *AlloydbCluster) LocalName() string {
	return ac.Name
}

// Configuration returns the configuration (args) for [AlloydbCluster].
func (ac *AlloydbCluster) Configuration() interface{} {
	return ac.Args
}

// DependOn is used for other resources to depend on [AlloydbCluster].
func (ac *AlloydbCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(ac)
}

// Dependencies returns the list of resources [AlloydbCluster] depends_on.
func (ac *AlloydbCluster) Dependencies() terra.Dependencies {
	return ac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AlloydbCluster].
func (ac *AlloydbCluster) LifecycleManagement() *terra.Lifecycle {
	return ac.Lifecycle
}

// Attributes returns the attributes for [AlloydbCluster].
func (ac *AlloydbCluster) Attributes() alloydbClusterAttributes {
	return alloydbClusterAttributes{ref: terra.ReferenceResource(ac)}
}

// ImportState imports the given attribute values into [AlloydbCluster]'s state.
func (ac *AlloydbCluster) ImportState(av io.Reader) error {
	ac.state = &alloydbClusterState{}
	if err := json.NewDecoder(av).Decode(ac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ac.Type(), ac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AlloydbCluster] has state.
func (ac *AlloydbCluster) State() (*alloydbClusterState, bool) {
	return ac.state, ac.state != nil
}

// StateMust returns the state for [AlloydbCluster]. Panics if the state is nil.
func (ac *AlloydbCluster) StateMust() *alloydbClusterState {
	if ac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ac.Type(), ac.LocalName()))
	}
	return ac.state
}

// AlloydbClusterArgs contains the configurations for google_alloydb_cluster.
type AlloydbClusterArgs struct {
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// BackupSource: min=0
	BackupSource []alloydbcluster.BackupSource `hcl:"backup_source,block" validate:"min=0"`
	// MigrationSource: min=0
	MigrationSource []alloydbcluster.MigrationSource `hcl:"migration_source,block" validate:"min=0"`
	// AutomatedBackupPolicy: optional
	AutomatedBackupPolicy *alloydbcluster.AutomatedBackupPolicy `hcl:"automated_backup_policy,block"`
	// InitialUser: optional
	InitialUser *alloydbcluster.InitialUser `hcl:"initial_user,block"`
	// Timeouts: optional
	Timeouts *alloydbcluster.Timeouts `hcl:"timeouts,block"`
}
type alloydbClusterAttributes struct {
	ref terra.Reference
}

// ClusterId returns a reference to field cluster_id of google_alloydb_cluster.
func (ac alloydbClusterAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("cluster_id"))
}

// DatabaseVersion returns a reference to field database_version of google_alloydb_cluster.
func (ac alloydbClusterAttributes) DatabaseVersion() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("database_version"))
}

// DisplayName returns a reference to field display_name of google_alloydb_cluster.
func (ac alloydbClusterAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("display_name"))
}

// Id returns a reference to field id of google_alloydb_cluster.
func (ac alloydbClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("id"))
}

// Labels returns a reference to field labels of google_alloydb_cluster.
func (ac alloydbClusterAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("labels"))
}

// Location returns a reference to field location of google_alloydb_cluster.
func (ac alloydbClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("location"))
}

// Name returns a reference to field name of google_alloydb_cluster.
func (ac alloydbClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("name"))
}

// Network returns a reference to field network of google_alloydb_cluster.
func (ac alloydbClusterAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("network"))
}

// Project returns a reference to field project of google_alloydb_cluster.
func (ac alloydbClusterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("project"))
}

// Uid returns a reference to field uid of google_alloydb_cluster.
func (ac alloydbClusterAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("uid"))
}

func (ac alloydbClusterAttributes) BackupSource() terra.ListValue[alloydbcluster.BackupSourceAttributes] {
	return terra.ReferenceAsList[alloydbcluster.BackupSourceAttributes](ac.ref.Append("backup_source"))
}

func (ac alloydbClusterAttributes) MigrationSource() terra.ListValue[alloydbcluster.MigrationSourceAttributes] {
	return terra.ReferenceAsList[alloydbcluster.MigrationSourceAttributes](ac.ref.Append("migration_source"))
}

func (ac alloydbClusterAttributes) AutomatedBackupPolicy() terra.ListValue[alloydbcluster.AutomatedBackupPolicyAttributes] {
	return terra.ReferenceAsList[alloydbcluster.AutomatedBackupPolicyAttributes](ac.ref.Append("automated_backup_policy"))
}

func (ac alloydbClusterAttributes) InitialUser() terra.ListValue[alloydbcluster.InitialUserAttributes] {
	return terra.ReferenceAsList[alloydbcluster.InitialUserAttributes](ac.ref.Append("initial_user"))
}

func (ac alloydbClusterAttributes) Timeouts() alloydbcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[alloydbcluster.TimeoutsAttributes](ac.ref.Append("timeouts"))
}

type alloydbClusterState struct {
	ClusterId             string                                      `json:"cluster_id"`
	DatabaseVersion       string                                      `json:"database_version"`
	DisplayName           string                                      `json:"display_name"`
	Id                    string                                      `json:"id"`
	Labels                map[string]string                           `json:"labels"`
	Location              string                                      `json:"location"`
	Name                  string                                      `json:"name"`
	Network               string                                      `json:"network"`
	Project               string                                      `json:"project"`
	Uid                   string                                      `json:"uid"`
	BackupSource          []alloydbcluster.BackupSourceState          `json:"backup_source"`
	MigrationSource       []alloydbcluster.MigrationSourceState       `json:"migration_source"`
	AutomatedBackupPolicy []alloydbcluster.AutomatedBackupPolicyState `json:"automated_backup_policy"`
	InitialUser           []alloydbcluster.InitialUserState           `json:"initial_user"`
	Timeouts              *alloydbcluster.TimeoutsState               `json:"timeouts"`
}
