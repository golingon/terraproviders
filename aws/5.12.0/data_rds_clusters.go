// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datardsclusters "github.com/golingon/terraproviders/aws/5.12.0/datardsclusters"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRdsClusters creates a new instance of [DataRdsClusters].
func NewDataRdsClusters(name string, args DataRdsClustersArgs) *DataRdsClusters {
	return &DataRdsClusters{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRdsClusters)(nil)

// DataRdsClusters represents the Terraform data resource aws_rds_clusters.
type DataRdsClusters struct {
	Name string
	Args DataRdsClustersArgs
}

// DataSource returns the Terraform object type for [DataRdsClusters].
func (rc *DataRdsClusters) DataSource() string {
	return "aws_rds_clusters"
}

// LocalName returns the local name for [DataRdsClusters].
func (rc *DataRdsClusters) LocalName() string {
	return rc.Name
}

// Configuration returns the configuration (args) for [DataRdsClusters].
func (rc *DataRdsClusters) Configuration() interface{} {
	return rc.Args
}

// Attributes returns the attributes for [DataRdsClusters].
func (rc *DataRdsClusters) Attributes() dataRdsClustersAttributes {
	return dataRdsClustersAttributes{ref: terra.ReferenceDataResource(rc)}
}

// DataRdsClustersArgs contains the configurations for aws_rds_clusters.
type DataRdsClustersArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Filter: min=0
	Filter []datardsclusters.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataRdsClustersAttributes struct {
	ref terra.Reference
}

// ClusterArns returns a reference to field cluster_arns of aws_rds_clusters.
func (rc dataRdsClustersAttributes) ClusterArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("cluster_arns"))
}

// ClusterIdentifiers returns a reference to field cluster_identifiers of aws_rds_clusters.
func (rc dataRdsClustersAttributes) ClusterIdentifiers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("cluster_identifiers"))
}

// Id returns a reference to field id of aws_rds_clusters.
func (rc dataRdsClustersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("id"))
}

func (rc dataRdsClustersAttributes) Filter() terra.SetValue[datardsclusters.FilterAttributes] {
	return terra.ReferenceAsSet[datardsclusters.FilterAttributes](rc.ref.Append("filter"))
}
