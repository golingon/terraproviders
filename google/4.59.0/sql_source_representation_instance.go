// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	sqlsourcerepresentationinstance "github.com/golingon/terraproviders/google/4.59.0/sqlsourcerepresentationinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSqlSourceRepresentationInstance(name string, args SqlSourceRepresentationInstanceArgs) *SqlSourceRepresentationInstance {
	return &SqlSourceRepresentationInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlSourceRepresentationInstance)(nil)

type SqlSourceRepresentationInstance struct {
	Name  string
	Args  SqlSourceRepresentationInstanceArgs
	state *sqlSourceRepresentationInstanceState
}

func (ssri *SqlSourceRepresentationInstance) Type() string {
	return "google_sql_source_representation_instance"
}

func (ssri *SqlSourceRepresentationInstance) LocalName() string {
	return ssri.Name
}

func (ssri *SqlSourceRepresentationInstance) Configuration() interface{} {
	return ssri.Args
}

func (ssri *SqlSourceRepresentationInstance) Attributes() sqlSourceRepresentationInstanceAttributes {
	return sqlSourceRepresentationInstanceAttributes{ref: terra.ReferenceResource(ssri)}
}

func (ssri *SqlSourceRepresentationInstance) ImportState(av io.Reader) error {
	ssri.state = &sqlSourceRepresentationInstanceState{}
	if err := json.NewDecoder(av).Decode(ssri.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssri.Type(), ssri.LocalName(), err)
	}
	return nil
}

func (ssri *SqlSourceRepresentationInstance) State() (*sqlSourceRepresentationInstanceState, bool) {
	return ssri.state, ssri.state != nil
}

func (ssri *SqlSourceRepresentationInstance) StateMust() *sqlSourceRepresentationInstanceState {
	if ssri.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssri.Type(), ssri.LocalName()))
	}
	return ssri.state
}

func (ssri *SqlSourceRepresentationInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(ssri)
}

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
	// DependsOn contains resources that SqlSourceRepresentationInstance depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sqlSourceRepresentationInstanceAttributes struct {
	ref terra.Reference
}

func (ssri sqlSourceRepresentationInstanceAttributes) CaCertificate() terra.StringValue {
	return terra.ReferenceString(ssri.ref.Append("ca_certificate"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) ClientCertificate() terra.StringValue {
	return terra.ReferenceString(ssri.ref.Append("client_certificate"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) ClientKey() terra.StringValue {
	return terra.ReferenceString(ssri.ref.Append("client_key"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) DatabaseVersion() terra.StringValue {
	return terra.ReferenceString(ssri.ref.Append("database_version"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) DumpFilePath() terra.StringValue {
	return terra.ReferenceString(ssri.ref.Append("dump_file_path"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) Host() terra.StringValue {
	return terra.ReferenceString(ssri.ref.Append("host"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ssri.ref.Append("id"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ssri.ref.Append("name"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) Password() terra.StringValue {
	return terra.ReferenceString(ssri.ref.Append("password"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) Port() terra.NumberValue {
	return terra.ReferenceNumber(ssri.ref.Append("port"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceString(ssri.ref.Append("project"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) Region() terra.StringValue {
	return terra.ReferenceString(ssri.ref.Append("region"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) Username() terra.StringValue {
	return terra.ReferenceString(ssri.ref.Append("username"))
}

func (ssri sqlSourceRepresentationInstanceAttributes) Timeouts() sqlsourcerepresentationinstance.TimeoutsAttributes {
	return terra.ReferenceSingle[sqlsourcerepresentationinstance.TimeoutsAttributes](ssri.ref.Append("timeouts"))
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
