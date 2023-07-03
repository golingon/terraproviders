// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datasqltiers "github.com/golingon/terraproviders/googlebeta/4.71.0/datasqltiers"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSqlTiers creates a new instance of [DataSqlTiers].
func NewDataSqlTiers(name string, args DataSqlTiersArgs) *DataSqlTiers {
	return &DataSqlTiers{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSqlTiers)(nil)

// DataSqlTiers represents the Terraform data resource google_sql_tiers.
type DataSqlTiers struct {
	Name string
	Args DataSqlTiersArgs
}

// DataSource returns the Terraform object type for [DataSqlTiers].
func (st *DataSqlTiers) DataSource() string {
	return "google_sql_tiers"
}

// LocalName returns the local name for [DataSqlTiers].
func (st *DataSqlTiers) LocalName() string {
	return st.Name
}

// Configuration returns the configuration (args) for [DataSqlTiers].
func (st *DataSqlTiers) Configuration() interface{} {
	return st.Args
}

// Attributes returns the attributes for [DataSqlTiers].
func (st *DataSqlTiers) Attributes() dataSqlTiersAttributes {
	return dataSqlTiersAttributes{ref: terra.ReferenceDataResource(st)}
}

// DataSqlTiersArgs contains the configurations for google_sql_tiers.
type DataSqlTiersArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Tiers: min=0
	Tiers []datasqltiers.Tiers `hcl:"tiers,block" validate:"min=0"`
}
type dataSqlTiersAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_sql_tiers.
func (st dataSqlTiersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("id"))
}

// Project returns a reference to field project of google_sql_tiers.
func (st dataSqlTiersAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("project"))
}

func (st dataSqlTiersAttributes) Tiers() terra.ListValue[datasqltiers.TiersAttributes] {
	return terra.ReferenceAsList[datasqltiers.TiersAttributes](st.ref.Append("tiers"))
}
