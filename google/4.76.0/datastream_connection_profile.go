// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	datastreamconnectionprofile "github.com/golingon/terraproviders/google/4.76.0/datastreamconnectionprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatastreamConnectionProfile creates a new instance of [DatastreamConnectionProfile].
func NewDatastreamConnectionProfile(name string, args DatastreamConnectionProfileArgs) *DatastreamConnectionProfile {
	return &DatastreamConnectionProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatastreamConnectionProfile)(nil)

// DatastreamConnectionProfile represents the Terraform resource google_datastream_connection_profile.
type DatastreamConnectionProfile struct {
	Name      string
	Args      DatastreamConnectionProfileArgs
	state     *datastreamConnectionProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatastreamConnectionProfile].
func (dcp *DatastreamConnectionProfile) Type() string {
	return "google_datastream_connection_profile"
}

// LocalName returns the local name for [DatastreamConnectionProfile].
func (dcp *DatastreamConnectionProfile) LocalName() string {
	return dcp.Name
}

// Configuration returns the configuration (args) for [DatastreamConnectionProfile].
func (dcp *DatastreamConnectionProfile) Configuration() interface{} {
	return dcp.Args
}

// DependOn is used for other resources to depend on [DatastreamConnectionProfile].
func (dcp *DatastreamConnectionProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(dcp)
}

// Dependencies returns the list of resources [DatastreamConnectionProfile] depends_on.
func (dcp *DatastreamConnectionProfile) Dependencies() terra.Dependencies {
	return dcp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatastreamConnectionProfile].
func (dcp *DatastreamConnectionProfile) LifecycleManagement() *terra.Lifecycle {
	return dcp.Lifecycle
}

// Attributes returns the attributes for [DatastreamConnectionProfile].
func (dcp *DatastreamConnectionProfile) Attributes() datastreamConnectionProfileAttributes {
	return datastreamConnectionProfileAttributes{ref: terra.ReferenceResource(dcp)}
}

// ImportState imports the given attribute values into [DatastreamConnectionProfile]'s state.
func (dcp *DatastreamConnectionProfile) ImportState(av io.Reader) error {
	dcp.state = &datastreamConnectionProfileState{}
	if err := json.NewDecoder(av).Decode(dcp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcp.Type(), dcp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatastreamConnectionProfile] has state.
func (dcp *DatastreamConnectionProfile) State() (*datastreamConnectionProfileState, bool) {
	return dcp.state, dcp.state != nil
}

// StateMust returns the state for [DatastreamConnectionProfile]. Panics if the state is nil.
func (dcp *DatastreamConnectionProfile) StateMust() *datastreamConnectionProfileState {
	if dcp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcp.Type(), dcp.LocalName()))
	}
	return dcp.state
}

// DatastreamConnectionProfileArgs contains the configurations for google_datastream_connection_profile.
type DatastreamConnectionProfileArgs struct {
	// ConnectionProfileId: string, required
	ConnectionProfileId terra.StringValue `hcl:"connection_profile_id,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// BigqueryProfile: optional
	BigqueryProfile *datastreamconnectionprofile.BigqueryProfile `hcl:"bigquery_profile,block"`
	// ForwardSshConnectivity: optional
	ForwardSshConnectivity *datastreamconnectionprofile.ForwardSshConnectivity `hcl:"forward_ssh_connectivity,block"`
	// GcsProfile: optional
	GcsProfile *datastreamconnectionprofile.GcsProfile `hcl:"gcs_profile,block"`
	// MysqlProfile: optional
	MysqlProfile *datastreamconnectionprofile.MysqlProfile `hcl:"mysql_profile,block"`
	// OracleProfile: optional
	OracleProfile *datastreamconnectionprofile.OracleProfile `hcl:"oracle_profile,block"`
	// PostgresqlProfile: optional
	PostgresqlProfile *datastreamconnectionprofile.PostgresqlProfile `hcl:"postgresql_profile,block"`
	// PrivateConnectivity: optional
	PrivateConnectivity *datastreamconnectionprofile.PrivateConnectivity `hcl:"private_connectivity,block"`
	// Timeouts: optional
	Timeouts *datastreamconnectionprofile.Timeouts `hcl:"timeouts,block"`
}
type datastreamConnectionProfileAttributes struct {
	ref terra.Reference
}

// ConnectionProfileId returns a reference to field connection_profile_id of google_datastream_connection_profile.
func (dcp datastreamConnectionProfileAttributes) ConnectionProfileId() terra.StringValue {
	return terra.ReferenceAsString(dcp.ref.Append("connection_profile_id"))
}

// DisplayName returns a reference to field display_name of google_datastream_connection_profile.
func (dcp datastreamConnectionProfileAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dcp.ref.Append("display_name"))
}

// Id returns a reference to field id of google_datastream_connection_profile.
func (dcp datastreamConnectionProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcp.ref.Append("id"))
}

// Labels returns a reference to field labels of google_datastream_connection_profile.
func (dcp datastreamConnectionProfileAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dcp.ref.Append("labels"))
}

// Location returns a reference to field location of google_datastream_connection_profile.
func (dcp datastreamConnectionProfileAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dcp.ref.Append("location"))
}

// Name returns a reference to field name of google_datastream_connection_profile.
func (dcp datastreamConnectionProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dcp.ref.Append("name"))
}

// Project returns a reference to field project of google_datastream_connection_profile.
func (dcp datastreamConnectionProfileAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dcp.ref.Append("project"))
}

func (dcp datastreamConnectionProfileAttributes) BigqueryProfile() terra.ListValue[datastreamconnectionprofile.BigqueryProfileAttributes] {
	return terra.ReferenceAsList[datastreamconnectionprofile.BigqueryProfileAttributes](dcp.ref.Append("bigquery_profile"))
}

func (dcp datastreamConnectionProfileAttributes) ForwardSshConnectivity() terra.ListValue[datastreamconnectionprofile.ForwardSshConnectivityAttributes] {
	return terra.ReferenceAsList[datastreamconnectionprofile.ForwardSshConnectivityAttributes](dcp.ref.Append("forward_ssh_connectivity"))
}

func (dcp datastreamConnectionProfileAttributes) GcsProfile() terra.ListValue[datastreamconnectionprofile.GcsProfileAttributes] {
	return terra.ReferenceAsList[datastreamconnectionprofile.GcsProfileAttributes](dcp.ref.Append("gcs_profile"))
}

func (dcp datastreamConnectionProfileAttributes) MysqlProfile() terra.ListValue[datastreamconnectionprofile.MysqlProfileAttributes] {
	return terra.ReferenceAsList[datastreamconnectionprofile.MysqlProfileAttributes](dcp.ref.Append("mysql_profile"))
}

func (dcp datastreamConnectionProfileAttributes) OracleProfile() terra.ListValue[datastreamconnectionprofile.OracleProfileAttributes] {
	return terra.ReferenceAsList[datastreamconnectionprofile.OracleProfileAttributes](dcp.ref.Append("oracle_profile"))
}

func (dcp datastreamConnectionProfileAttributes) PostgresqlProfile() terra.ListValue[datastreamconnectionprofile.PostgresqlProfileAttributes] {
	return terra.ReferenceAsList[datastreamconnectionprofile.PostgresqlProfileAttributes](dcp.ref.Append("postgresql_profile"))
}

func (dcp datastreamConnectionProfileAttributes) PrivateConnectivity() terra.ListValue[datastreamconnectionprofile.PrivateConnectivityAttributes] {
	return terra.ReferenceAsList[datastreamconnectionprofile.PrivateConnectivityAttributes](dcp.ref.Append("private_connectivity"))
}

func (dcp datastreamConnectionProfileAttributes) Timeouts() datastreamconnectionprofile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datastreamconnectionprofile.TimeoutsAttributes](dcp.ref.Append("timeouts"))
}

type datastreamConnectionProfileState struct {
	ConnectionProfileId    string                                                    `json:"connection_profile_id"`
	DisplayName            string                                                    `json:"display_name"`
	Id                     string                                                    `json:"id"`
	Labels                 map[string]string                                         `json:"labels"`
	Location               string                                                    `json:"location"`
	Name                   string                                                    `json:"name"`
	Project                string                                                    `json:"project"`
	BigqueryProfile        []datastreamconnectionprofile.BigqueryProfileState        `json:"bigquery_profile"`
	ForwardSshConnectivity []datastreamconnectionprofile.ForwardSshConnectivityState `json:"forward_ssh_connectivity"`
	GcsProfile             []datastreamconnectionprofile.GcsProfileState             `json:"gcs_profile"`
	MysqlProfile           []datastreamconnectionprofile.MysqlProfileState           `json:"mysql_profile"`
	OracleProfile          []datastreamconnectionprofile.OracleProfileState          `json:"oracle_profile"`
	PostgresqlProfile      []datastreamconnectionprofile.PostgresqlProfileState      `json:"postgresql_profile"`
	PrivateConnectivity    []datastreamconnectionprofile.PrivateConnectivityState    `json:"private_connectivity"`
	Timeouts               *datastreamconnectionprofile.TimeoutsState                `json:"timeouts"`
}