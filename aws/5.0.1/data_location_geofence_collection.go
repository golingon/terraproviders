// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataLocationGeofenceCollection creates a new instance of [DataLocationGeofenceCollection].
func NewDataLocationGeofenceCollection(name string, args DataLocationGeofenceCollectionArgs) *DataLocationGeofenceCollection {
	return &DataLocationGeofenceCollection{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLocationGeofenceCollection)(nil)

// DataLocationGeofenceCollection represents the Terraform data resource aws_location_geofence_collection.
type DataLocationGeofenceCollection struct {
	Name string
	Args DataLocationGeofenceCollectionArgs
}

// DataSource returns the Terraform object type for [DataLocationGeofenceCollection].
func (lgc *DataLocationGeofenceCollection) DataSource() string {
	return "aws_location_geofence_collection"
}

// LocalName returns the local name for [DataLocationGeofenceCollection].
func (lgc *DataLocationGeofenceCollection) LocalName() string {
	return lgc.Name
}

// Configuration returns the configuration (args) for [DataLocationGeofenceCollection].
func (lgc *DataLocationGeofenceCollection) Configuration() interface{} {
	return lgc.Args
}

// Attributes returns the attributes for [DataLocationGeofenceCollection].
func (lgc *DataLocationGeofenceCollection) Attributes() dataLocationGeofenceCollectionAttributes {
	return dataLocationGeofenceCollectionAttributes{ref: terra.ReferenceDataResource(lgc)}
}

// DataLocationGeofenceCollectionArgs contains the configurations for aws_location_geofence_collection.
type DataLocationGeofenceCollectionArgs struct {
	// CollectionName: string, required
	CollectionName terra.StringValue `hcl:"collection_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataLocationGeofenceCollectionAttributes struct {
	ref terra.Reference
}

// CollectionArn returns a reference to field collection_arn of aws_location_geofence_collection.
func (lgc dataLocationGeofenceCollectionAttributes) CollectionArn() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("collection_arn"))
}

// CollectionName returns a reference to field collection_name of aws_location_geofence_collection.
func (lgc dataLocationGeofenceCollectionAttributes) CollectionName() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("collection_name"))
}

// CreateTime returns a reference to field create_time of aws_location_geofence_collection.
func (lgc dataLocationGeofenceCollectionAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("create_time"))
}

// Description returns a reference to field description of aws_location_geofence_collection.
func (lgc dataLocationGeofenceCollectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("description"))
}

// Id returns a reference to field id of aws_location_geofence_collection.
func (lgc dataLocationGeofenceCollectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_location_geofence_collection.
func (lgc dataLocationGeofenceCollectionAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("kms_key_id"))
}

// Tags returns a reference to field tags of aws_location_geofence_collection.
func (lgc dataLocationGeofenceCollectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lgc.ref.Append("tags"))
}

// UpdateTime returns a reference to field update_time of aws_location_geofence_collection.
func (lgc dataLocationGeofenceCollectionAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("update_time"))
}