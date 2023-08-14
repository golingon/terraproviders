// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	locationgeofencecollection "github.com/golingon/terraproviders/aws/5.12.0/locationgeofencecollection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLocationGeofenceCollection creates a new instance of [LocationGeofenceCollection].
func NewLocationGeofenceCollection(name string, args LocationGeofenceCollectionArgs) *LocationGeofenceCollection {
	return &LocationGeofenceCollection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LocationGeofenceCollection)(nil)

// LocationGeofenceCollection represents the Terraform resource aws_location_geofence_collection.
type LocationGeofenceCollection struct {
	Name      string
	Args      LocationGeofenceCollectionArgs
	state     *locationGeofenceCollectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LocationGeofenceCollection].
func (lgc *LocationGeofenceCollection) Type() string {
	return "aws_location_geofence_collection"
}

// LocalName returns the local name for [LocationGeofenceCollection].
func (lgc *LocationGeofenceCollection) LocalName() string {
	return lgc.Name
}

// Configuration returns the configuration (args) for [LocationGeofenceCollection].
func (lgc *LocationGeofenceCollection) Configuration() interface{} {
	return lgc.Args
}

// DependOn is used for other resources to depend on [LocationGeofenceCollection].
func (lgc *LocationGeofenceCollection) DependOn() terra.Reference {
	return terra.ReferenceResource(lgc)
}

// Dependencies returns the list of resources [LocationGeofenceCollection] depends_on.
func (lgc *LocationGeofenceCollection) Dependencies() terra.Dependencies {
	return lgc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LocationGeofenceCollection].
func (lgc *LocationGeofenceCollection) LifecycleManagement() *terra.Lifecycle {
	return lgc.Lifecycle
}

// Attributes returns the attributes for [LocationGeofenceCollection].
func (lgc *LocationGeofenceCollection) Attributes() locationGeofenceCollectionAttributes {
	return locationGeofenceCollectionAttributes{ref: terra.ReferenceResource(lgc)}
}

// ImportState imports the given attribute values into [LocationGeofenceCollection]'s state.
func (lgc *LocationGeofenceCollection) ImportState(av io.Reader) error {
	lgc.state = &locationGeofenceCollectionState{}
	if err := json.NewDecoder(av).Decode(lgc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lgc.Type(), lgc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LocationGeofenceCollection] has state.
func (lgc *LocationGeofenceCollection) State() (*locationGeofenceCollectionState, bool) {
	return lgc.state, lgc.state != nil
}

// StateMust returns the state for [LocationGeofenceCollection]. Panics if the state is nil.
func (lgc *LocationGeofenceCollection) StateMust() *locationGeofenceCollectionState {
	if lgc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lgc.Type(), lgc.LocalName()))
	}
	return lgc.state
}

// LocationGeofenceCollectionArgs contains the configurations for aws_location_geofence_collection.
type LocationGeofenceCollectionArgs struct {
	// CollectionName: string, required
	CollectionName terra.StringValue `hcl:"collection_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *locationgeofencecollection.Timeouts `hcl:"timeouts,block"`
}
type locationGeofenceCollectionAttributes struct {
	ref terra.Reference
}

// CollectionArn returns a reference to field collection_arn of aws_location_geofence_collection.
func (lgc locationGeofenceCollectionAttributes) CollectionArn() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("collection_arn"))
}

// CollectionName returns a reference to field collection_name of aws_location_geofence_collection.
func (lgc locationGeofenceCollectionAttributes) CollectionName() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("collection_name"))
}

// CreateTime returns a reference to field create_time of aws_location_geofence_collection.
func (lgc locationGeofenceCollectionAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("create_time"))
}

// Description returns a reference to field description of aws_location_geofence_collection.
func (lgc locationGeofenceCollectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("description"))
}

// Id returns a reference to field id of aws_location_geofence_collection.
func (lgc locationGeofenceCollectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_location_geofence_collection.
func (lgc locationGeofenceCollectionAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("kms_key_id"))
}

// Tags returns a reference to field tags of aws_location_geofence_collection.
func (lgc locationGeofenceCollectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lgc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_location_geofence_collection.
func (lgc locationGeofenceCollectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lgc.ref.Append("tags_all"))
}

// UpdateTime returns a reference to field update_time of aws_location_geofence_collection.
func (lgc locationGeofenceCollectionAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(lgc.ref.Append("update_time"))
}

func (lgc locationGeofenceCollectionAttributes) Timeouts() locationgeofencecollection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[locationgeofencecollection.TimeoutsAttributes](lgc.ref.Append("timeouts"))
}

type locationGeofenceCollectionState struct {
	CollectionArn  string                                    `json:"collection_arn"`
	CollectionName string                                    `json:"collection_name"`
	CreateTime     string                                    `json:"create_time"`
	Description    string                                    `json:"description"`
	Id             string                                    `json:"id"`
	KmsKeyId       string                                    `json:"kms_key_id"`
	Tags           map[string]string                         `json:"tags"`
	TagsAll        map[string]string                         `json:"tags_all"`
	UpdateTime     string                                    `json:"update_time"`
	Timeouts       *locationgeofencecollection.TimeoutsState `json:"timeouts"`
}