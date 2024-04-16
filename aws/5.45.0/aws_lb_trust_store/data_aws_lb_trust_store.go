// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_lb_trust_store

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_lb_trust_store.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (alts *DataSource) DataSource() string {
	return "aws_lb_trust_store"
}

// LocalName returns the local name for [DataSource].
func (alts *DataSource) LocalName() string {
	return alts.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (alts *DataSource) Configuration() interface{} {
	return alts.Args
}

// Attributes returns the attributes for [DataSource].
func (alts *DataSource) Attributes() dataAwsLbTrustStoreAttributes {
	return dataAwsLbTrustStoreAttributes{ref: terra.ReferenceDataSource(alts)}
}

// DataArgs contains the configurations for aws_lb_trust_store.
type DataArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}

type dataAwsLbTrustStoreAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lb_trust_store.
func (alts dataAwsLbTrustStoreAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(alts.ref.Append("arn"))
}

// Id returns a reference to field id of aws_lb_trust_store.
func (alts dataAwsLbTrustStoreAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(alts.ref.Append("id"))
}

// Name returns a reference to field name of aws_lb_trust_store.
func (alts dataAwsLbTrustStoreAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(alts.ref.Append("name"))
}
