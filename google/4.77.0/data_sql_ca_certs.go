// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datasqlcacerts "github.com/golingon/terraproviders/google/4.77.0/datasqlcacerts"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSqlCaCerts creates a new instance of [DataSqlCaCerts].
func NewDataSqlCaCerts(name string, args DataSqlCaCertsArgs) *DataSqlCaCerts {
	return &DataSqlCaCerts{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSqlCaCerts)(nil)

// DataSqlCaCerts represents the Terraform data resource google_sql_ca_certs.
type DataSqlCaCerts struct {
	Name string
	Args DataSqlCaCertsArgs
}

// DataSource returns the Terraform object type for [DataSqlCaCerts].
func (scc *DataSqlCaCerts) DataSource() string {
	return "google_sql_ca_certs"
}

// LocalName returns the local name for [DataSqlCaCerts].
func (scc *DataSqlCaCerts) LocalName() string {
	return scc.Name
}

// Configuration returns the configuration (args) for [DataSqlCaCerts].
func (scc *DataSqlCaCerts) Configuration() interface{} {
	return scc.Args
}

// Attributes returns the attributes for [DataSqlCaCerts].
func (scc *DataSqlCaCerts) Attributes() dataSqlCaCertsAttributes {
	return dataSqlCaCertsAttributes{ref: terra.ReferenceDataResource(scc)}
}

// DataSqlCaCertsArgs contains the configurations for google_sql_ca_certs.
type DataSqlCaCertsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Certs: min=0
	Certs []datasqlcacerts.Certs `hcl:"certs,block" validate:"min=0"`
}
type dataSqlCaCertsAttributes struct {
	ref terra.Reference
}

// ActiveVersion returns a reference to field active_version of google_sql_ca_certs.
func (scc dataSqlCaCertsAttributes) ActiveVersion() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("active_version"))
}

// Id returns a reference to field id of google_sql_ca_certs.
func (scc dataSqlCaCertsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("id"))
}

// Instance returns a reference to field instance of google_sql_ca_certs.
func (scc dataSqlCaCertsAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("instance"))
}

// Project returns a reference to field project of google_sql_ca_certs.
func (scc dataSqlCaCertsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("project"))
}

func (scc dataSqlCaCertsAttributes) Certs() terra.ListValue[datasqlcacerts.CertsAttributes] {
	return terra.ReferenceAsList[datasqlcacerts.CertsAttributes](scc.ref.Append("certs"))
}
