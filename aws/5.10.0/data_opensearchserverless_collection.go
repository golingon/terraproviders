// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataOpensearchserverlessCollection creates a new instance of [DataOpensearchserverlessCollection].
func NewDataOpensearchserverlessCollection(name string, args DataOpensearchserverlessCollectionArgs) *DataOpensearchserverlessCollection {
	return &DataOpensearchserverlessCollection{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOpensearchserverlessCollection)(nil)

// DataOpensearchserverlessCollection represents the Terraform data resource aws_opensearchserverless_collection.
type DataOpensearchserverlessCollection struct {
	Name string
	Args DataOpensearchserverlessCollectionArgs
}

// DataSource returns the Terraform object type for [DataOpensearchserverlessCollection].
func (oc *DataOpensearchserverlessCollection) DataSource() string {
	return "aws_opensearchserverless_collection"
}

// LocalName returns the local name for [DataOpensearchserverlessCollection].
func (oc *DataOpensearchserverlessCollection) LocalName() string {
	return oc.Name
}

// Configuration returns the configuration (args) for [DataOpensearchserverlessCollection].
func (oc *DataOpensearchserverlessCollection) Configuration() interface{} {
	return oc.Args
}

// Attributes returns the attributes for [DataOpensearchserverlessCollection].
func (oc *DataOpensearchserverlessCollection) Attributes() dataOpensearchserverlessCollectionAttributes {
	return dataOpensearchserverlessCollectionAttributes{ref: terra.ReferenceDataResource(oc)}
}

// DataOpensearchserverlessCollectionArgs contains the configurations for aws_opensearchserverless_collection.
type DataOpensearchserverlessCollectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}
type dataOpensearchserverlessCollectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_opensearchserverless_collection.
func (oc dataOpensearchserverlessCollectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("arn"))
}

// CollectionEndpoint returns a reference to field collection_endpoint of aws_opensearchserverless_collection.
func (oc dataOpensearchserverlessCollectionAttributes) CollectionEndpoint() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("collection_endpoint"))
}

// CreatedDate returns a reference to field created_date of aws_opensearchserverless_collection.
func (oc dataOpensearchserverlessCollectionAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("created_date"))
}

// DashboardEndpoint returns a reference to field dashboard_endpoint of aws_opensearchserverless_collection.
func (oc dataOpensearchserverlessCollectionAttributes) DashboardEndpoint() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("dashboard_endpoint"))
}

// Description returns a reference to field description of aws_opensearchserverless_collection.
func (oc dataOpensearchserverlessCollectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("description"))
}

// Id returns a reference to field id of aws_opensearchserverless_collection.
func (oc dataOpensearchserverlessCollectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_opensearchserverless_collection.
func (oc dataOpensearchserverlessCollectionAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("kms_key_arn"))
}

// LastModifiedDate returns a reference to field last_modified_date of aws_opensearchserverless_collection.
func (oc dataOpensearchserverlessCollectionAttributes) LastModifiedDate() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("last_modified_date"))
}

// Name returns a reference to field name of aws_opensearchserverless_collection.
func (oc dataOpensearchserverlessCollectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_opensearchserverless_collection.
func (oc dataOpensearchserverlessCollectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oc.ref.Append("tags"))
}

// Type returns a reference to field type of aws_opensearchserverless_collection.
func (oc dataOpensearchserverlessCollectionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("type"))
}
