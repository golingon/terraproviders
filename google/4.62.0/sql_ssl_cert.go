// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	sqlsslcert "github.com/golingon/terraproviders/google/4.62.0/sqlsslcert"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlSslCert creates a new instance of [SqlSslCert].
func NewSqlSslCert(name string, args SqlSslCertArgs) *SqlSslCert {
	return &SqlSslCert{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlSslCert)(nil)

// SqlSslCert represents the Terraform resource google_sql_ssl_cert.
type SqlSslCert struct {
	Name      string
	Args      SqlSslCertArgs
	state     *sqlSslCertState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlSslCert].
func (ssc *SqlSslCert) Type() string {
	return "google_sql_ssl_cert"
}

// LocalName returns the local name for [SqlSslCert].
func (ssc *SqlSslCert) LocalName() string {
	return ssc.Name
}

// Configuration returns the configuration (args) for [SqlSslCert].
func (ssc *SqlSslCert) Configuration() interface{} {
	return ssc.Args
}

// DependOn is used for other resources to depend on [SqlSslCert].
func (ssc *SqlSslCert) DependOn() terra.Reference {
	return terra.ReferenceResource(ssc)
}

// Dependencies returns the list of resources [SqlSslCert] depends_on.
func (ssc *SqlSslCert) Dependencies() terra.Dependencies {
	return ssc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlSslCert].
func (ssc *SqlSslCert) LifecycleManagement() *terra.Lifecycle {
	return ssc.Lifecycle
}

// Attributes returns the attributes for [SqlSslCert].
func (ssc *SqlSslCert) Attributes() sqlSslCertAttributes {
	return sqlSslCertAttributes{ref: terra.ReferenceResource(ssc)}
}

// ImportState imports the given attribute values into [SqlSslCert]'s state.
func (ssc *SqlSslCert) ImportState(av io.Reader) error {
	ssc.state = &sqlSslCertState{}
	if err := json.NewDecoder(av).Decode(ssc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssc.Type(), ssc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlSslCert] has state.
func (ssc *SqlSslCert) State() (*sqlSslCertState, bool) {
	return ssc.state, ssc.state != nil
}

// StateMust returns the state for [SqlSslCert]. Panics if the state is nil.
func (ssc *SqlSslCert) StateMust() *sqlSslCertState {
	if ssc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssc.Type(), ssc.LocalName()))
	}
	return ssc.state
}

// SqlSslCertArgs contains the configurations for google_sql_ssl_cert.
type SqlSslCertArgs struct {
	// CommonName: string, required
	CommonName terra.StringValue `hcl:"common_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *sqlsslcert.Timeouts `hcl:"timeouts,block"`
}
type sqlSslCertAttributes struct {
	ref terra.Reference
}

// Cert returns a reference to field cert of google_sql_ssl_cert.
func (ssc sqlSslCertAttributes) Cert() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("cert"))
}

// CertSerialNumber returns a reference to field cert_serial_number of google_sql_ssl_cert.
func (ssc sqlSslCertAttributes) CertSerialNumber() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("cert_serial_number"))
}

// CommonName returns a reference to field common_name of google_sql_ssl_cert.
func (ssc sqlSslCertAttributes) CommonName() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("common_name"))
}

// CreateTime returns a reference to field create_time of google_sql_ssl_cert.
func (ssc sqlSslCertAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("create_time"))
}

// ExpirationTime returns a reference to field expiration_time of google_sql_ssl_cert.
func (ssc sqlSslCertAttributes) ExpirationTime() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("expiration_time"))
}

// Id returns a reference to field id of google_sql_ssl_cert.
func (ssc sqlSslCertAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("id"))
}

// Instance returns a reference to field instance of google_sql_ssl_cert.
func (ssc sqlSslCertAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("instance"))
}

// PrivateKey returns a reference to field private_key of google_sql_ssl_cert.
func (ssc sqlSslCertAttributes) PrivateKey() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("private_key"))
}

// Project returns a reference to field project of google_sql_ssl_cert.
func (ssc sqlSslCertAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("project"))
}

// ServerCaCert returns a reference to field server_ca_cert of google_sql_ssl_cert.
func (ssc sqlSslCertAttributes) ServerCaCert() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("server_ca_cert"))
}

// Sha1Fingerprint returns a reference to field sha1_fingerprint of google_sql_ssl_cert.
func (ssc sqlSslCertAttributes) Sha1Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(ssc.ref.Append("sha1_fingerprint"))
}

func (ssc sqlSslCertAttributes) Timeouts() sqlsslcert.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqlsslcert.TimeoutsAttributes](ssc.ref.Append("timeouts"))
}

type sqlSslCertState struct {
	Cert             string                    `json:"cert"`
	CertSerialNumber string                    `json:"cert_serial_number"`
	CommonName       string                    `json:"common_name"`
	CreateTime       string                    `json:"create_time"`
	ExpirationTime   string                    `json:"expiration_time"`
	Id               string                    `json:"id"`
	Instance         string                    `json:"instance"`
	PrivateKey       string                    `json:"private_key"`
	Project          string                    `json:"project"`
	ServerCaCert     string                    `json:"server_ca_cert"`
	Sha1Fingerprint  string                    `json:"sha1_fingerprint"`
	Timeouts         *sqlsslcert.TimeoutsState `json:"timeouts"`
}
