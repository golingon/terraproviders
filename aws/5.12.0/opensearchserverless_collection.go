// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opensearchserverlesscollection "github.com/golingon/terraproviders/aws/5.12.0/opensearchserverlesscollection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpensearchserverlessCollection creates a new instance of [OpensearchserverlessCollection].
func NewOpensearchserverlessCollection(name string, args OpensearchserverlessCollectionArgs) *OpensearchserverlessCollection {
	return &OpensearchserverlessCollection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpensearchserverlessCollection)(nil)

// OpensearchserverlessCollection represents the Terraform resource aws_opensearchserverless_collection.
type OpensearchserverlessCollection struct {
	Name      string
	Args      OpensearchserverlessCollectionArgs
	state     *opensearchserverlessCollectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpensearchserverlessCollection].
func (oc *OpensearchserverlessCollection) Type() string {
	return "aws_opensearchserverless_collection"
}

// LocalName returns the local name for [OpensearchserverlessCollection].
func (oc *OpensearchserverlessCollection) LocalName() string {
	return oc.Name
}

// Configuration returns the configuration (args) for [OpensearchserverlessCollection].
func (oc *OpensearchserverlessCollection) Configuration() interface{} {
	return oc.Args
}

// DependOn is used for other resources to depend on [OpensearchserverlessCollection].
func (oc *OpensearchserverlessCollection) DependOn() terra.Reference {
	return terra.ReferenceResource(oc)
}

// Dependencies returns the list of resources [OpensearchserverlessCollection] depends_on.
func (oc *OpensearchserverlessCollection) Dependencies() terra.Dependencies {
	return oc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpensearchserverlessCollection].
func (oc *OpensearchserverlessCollection) LifecycleManagement() *terra.Lifecycle {
	return oc.Lifecycle
}

// Attributes returns the attributes for [OpensearchserverlessCollection].
func (oc *OpensearchserverlessCollection) Attributes() opensearchserverlessCollectionAttributes {
	return opensearchserverlessCollectionAttributes{ref: terra.ReferenceResource(oc)}
}

// ImportState imports the given attribute values into [OpensearchserverlessCollection]'s state.
func (oc *OpensearchserverlessCollection) ImportState(av io.Reader) error {
	oc.state = &opensearchserverlessCollectionState{}
	if err := json.NewDecoder(av).Decode(oc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oc.Type(), oc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpensearchserverlessCollection] has state.
func (oc *OpensearchserverlessCollection) State() (*opensearchserverlessCollectionState, bool) {
	return oc.state, oc.state != nil
}

// StateMust returns the state for [OpensearchserverlessCollection]. Panics if the state is nil.
func (oc *OpensearchserverlessCollection) StateMust() *opensearchserverlessCollectionState {
	if oc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oc.Type(), oc.LocalName()))
	}
	return oc.state
}

// OpensearchserverlessCollectionArgs contains the configurations for aws_opensearchserverless_collection.
type OpensearchserverlessCollectionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Timeouts: optional
	Timeouts *opensearchserverlesscollection.Timeouts `hcl:"timeouts,block"`
}
type opensearchserverlessCollectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_opensearchserverless_collection.
func (oc opensearchserverlessCollectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("arn"))
}

// CollectionEndpoint returns a reference to field collection_endpoint of aws_opensearchserverless_collection.
func (oc opensearchserverlessCollectionAttributes) CollectionEndpoint() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("collection_endpoint"))
}

// DashboardEndpoint returns a reference to field dashboard_endpoint of aws_opensearchserverless_collection.
func (oc opensearchserverlessCollectionAttributes) DashboardEndpoint() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("dashboard_endpoint"))
}

// Description returns a reference to field description of aws_opensearchserverless_collection.
func (oc opensearchserverlessCollectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("description"))
}

// Id returns a reference to field id of aws_opensearchserverless_collection.
func (oc opensearchserverlessCollectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_opensearchserverless_collection.
func (oc opensearchserverlessCollectionAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("kms_key_arn"))
}

// Name returns a reference to field name of aws_opensearchserverless_collection.
func (oc opensearchserverlessCollectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_opensearchserverless_collection.
func (oc opensearchserverlessCollectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_opensearchserverless_collection.
func (oc opensearchserverlessCollectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oc.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_opensearchserverless_collection.
func (oc opensearchserverlessCollectionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("type"))
}

func (oc opensearchserverlessCollectionAttributes) Timeouts() opensearchserverlesscollection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[opensearchserverlesscollection.TimeoutsAttributes](oc.ref.Append("timeouts"))
}

type opensearchserverlessCollectionState struct {
	Arn                string                                        `json:"arn"`
	CollectionEndpoint string                                        `json:"collection_endpoint"`
	DashboardEndpoint  string                                        `json:"dashboard_endpoint"`
	Description        string                                        `json:"description"`
	Id                 string                                        `json:"id"`
	KmsKeyArn          string                                        `json:"kms_key_arn"`
	Name               string                                        `json:"name"`
	Tags               map[string]string                             `json:"tags"`
	TagsAll            map[string]string                             `json:"tags_all"`
	Type               string                                        `json:"type"`
	Timeouts           *opensearchserverlesscollection.TimeoutsState `json:"timeouts"`
}
