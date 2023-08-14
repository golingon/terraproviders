// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	sqlsourcerepresentationinstance "github.com/golingon/terraproviders/google/4.77.0/sqlsourcerepresentationinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlSourceRepresentationInstance creates a new instance of [SqlSourceRepresentationInstance].
func NewSqlSourceRepresentationInstance(name string, args SqlSourceRepresentationInstanceArgs) *SqlSourceRepresentationInstance {
	return &SqlSourceRepresentationInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlSourceRepresentationInstance)(nil)

// SqlSourceRepresentationInstance represents the Terraform resource google_sql_source_representation_instance.
type SqlSourceRepresentationInstance struct {
	Name      string
	Args      SqlSourceRepresentationInstanceArgs
	state     *sqlSourceRepresentationInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlSourceRepresentationInstance].
func (ssri *SqlSourceRepresentationInstance) Type() string {
	return "google_sql_source_representation_instance"
}

// LocalName returns the local name for [SqlSourceRepresentationInstance].
func (ssri *SqlSourceRepresentationInstance) LocalName() string {
	return ssri.Name
}

// Configuration returns the configuration (args) for [SqlSourceRepresentationInstance].
func (ssri *SqlSourceRepresentationInstance) Configuration() interface{} {
	return ssri.Args
}

// DependOn is used for other resources to depend on [SqlSourceRepresentationInstance].
func (ssri *SqlSourceRepresentationInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(ssri)
}

// Dependencies returns the list of resources [SqlSourceRepresentationInstance] depends_on.
func (ssri *SqlSourceRepresentationInstance) Dependencies() terra.Dependencies {
	return ssri.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlSourceRepresentationInstance].
func (ssri *SqlSourceRepresentationInstance) LifecycleManagement() *terra.Lifecycle {
	return ssri.Lifecycle
}

// Attributes returns the attributes for [SqlSourceRepresentationInstance].
func (ssri *SqlSourceRepresentationInstance) Attributes() sqlSourceRepresentationInstanceAttributes {
	return sqlSourceRepresentationInstanceAttributes{ref: terra.ReferenceResource(ssri)}
}

// ImportState imports the given attribute values into [SqlSourceRepresentationInstance]'s state.
func (ssri *SqlSourceRepresentationInstance) ImportState(av io.Reader) error {
	ssri.state = &sqlSourceRepresentationInstanceState{}
	if err := json.NewDecoder(av).Decode(ssri.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssri.Type(), ssri.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlSourceRepresentationInstance] has state.
func (ssri *SqlSourceRepresentationInstance) State() (*sqlSourceRepresentationInstanceState, bool) {
	return ssri.state, ssri.state != nil
}

// StateMust returns the state for [SqlSourceRepresentationInstance]. Panics if the state is nil.
func (ssri *SqlSourceRepresentationInstance) StateMust() *sqlSourceRepresentationInstanceState {
	if ssri.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssri.Type(), ssri.LocalName()))
	}
	return ssri.state
}

// SqlSourceRepresentationInstanceArgs contains the configurations for google_sql_source_representation_instance.
type SqlSourceRepresentationInstanceArgs struct {
	// CaCertificate: string, optional
	CaCertificate terra.StringValue `hcl:"ca_certificate,attr"`
	// ClientCertificate: string, optional
	ClientCertificate terra.StringValue `hcl:"client_certificate,attr"`
	// ClientKey: string, optional
	ClientKey terra.StringValue `hcl:"client_key,attr"`
	// DatabaseVersion: string, required
	DatabaseVersion terra.StringValue `hcl:"database_version,attr" validate:"required"`
	// DumpFilePath: string, optional
	DumpFilePath terra.StringValue `hcl:"dump_file_path,attr"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Username: string, optional
	Username terra.StringValue `hcl:"username,attr"`
	// Timeouts: optional
	Timeouts *sqlsourcerepresentationinstance.Timeouts `hcl:"timeouts,block"`
}
type sqlSourceRepresentationInstanceAttributes struct {
	ref terra.Reference
}

// CaCertificate returns a reference to field ca_certificate of google_sql_source_representation_instance.
func (ssri sqlSourceRepresentationInstanceAttributes) CaCertificate() terra.StringValue {
	return terra.ReferenceAsString(ssri.ref.Append("ca_certificate"))
}

// ClientCertificate returns a reference to field client_certificate of google_sql_source_representation_instance.
func (ssri sqlSourceRepresentationInstanceAttributes) ClientCertificate() terra.StringValue {
	return terra.ReferenceAsString(ssri.ref.Append("client_certificate"))
}

// ClientKey returns a reference to field client_key of google_sql_source_representation_instance.
func (ssri sqlSourceRepresentationInstanceAttributes) ClientKey() terra.StringValue {
	return terra.ReferenceAsString(ssri.ref.Append("client_key"))
}

// DatabaseVersion returns a reference to field database_version of google_sql_source_representation_instance.
func (ssri sqlSourceRepresentationInstanceAttributes) DatabaseVersion() terra.StringValue {
	return terra.ReferenceAsString(ssri.ref.Append("database_version"))
}

// DumpFilePath returns a reference to field dump_file_path of google_sql_source_representation_instance.
func (ssri sqlSourceRepresentationInstanceAttributes) DumpFilePath() terra.StringValue {
	return terra.ReferenceAsString(ssri.ref.Append("dump_file_path"))
}

// Host returns a reference to field host of google_sql_source_representation_instance.
func (ssri sqlSourceRepresentationInstanceAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(ssri.ref.Append("host"))
}

// Id returns a reference to field id of google_sql_source_representation_instance.
func (ssri sqlSourceRepresentationInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssri.ref.Append("id"))
}

// Name returns a reference to field name of google_sql_source_representation_instance.
func (ssri sqlSourceRepresentationInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ssri.ref.Append("name"))
}

// Password returns a reference to field password of google_sql_source_representation_instance.
func (ssri sqlSourceRepresentationInstanceAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(ssri.ref.Append("password"))
}

// Port returns a reference to field port of google_sql_source_representation_instance.
func (ssri sqlSourceRepresentationInstanceAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ssri.ref.Append("port"))
}

// Project returns a reference to field project of google_sql_source_representation_instance.
func (ssri sqlSourceRepresentationInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ssri.ref.Append("project"))
}

// Region returns a reference to field region of google_sql_source_representation_instance.
func (ssri sqlSourceRepresentationInstanceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ssri.ref.Append("region"))
}

// Username returns a reference to field username of google_sql_source_representation_instance.
func (ssri sqlSourceRepresentationInstanceAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(ssri.ref.Append("username"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) Timeouts() sqlsourcerepresentationinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqlsourcerepresentationinstance.TimeoutsAttributes](ssri.ref.Append("timeouts"))
}

type sqlSourceRepresentationInstanceState struct {
	CaCertificate     string                                         `json:"ca_certificate"`
	ClientCertificate string                                         `json:"client_certificate"`
	ClientKey         string                                         `json:"client_key"`
	DatabaseVersion   string                                         `json:"database_version"`
	DumpFilePath      string                                         `json:"dump_file_path"`
	Host              string                                         `json:"host"`
	Id                string                                         `json:"id"`
	Name              string                                         `json:"name"`
	Password          string                                         `json:"password"`
	Port              float64                                        `json:"port"`
	Project           string                                         `json:"project"`
	Region            string                                         `json:"region"`
	Username          string                                         `json:"username"`
	Timeouts          *sqlsourcerepresentationinstance.TimeoutsState `json:"timeouts"`
}
