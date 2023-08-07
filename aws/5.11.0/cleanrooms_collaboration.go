// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cleanroomscollaboration "github.com/golingon/terraproviders/aws/5.11.0/cleanroomscollaboration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCleanroomsCollaboration creates a new instance of [CleanroomsCollaboration].
func NewCleanroomsCollaboration(name string, args CleanroomsCollaborationArgs) *CleanroomsCollaboration {
	return &CleanroomsCollaboration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CleanroomsCollaboration)(nil)

// CleanroomsCollaboration represents the Terraform resource aws_cleanrooms_collaboration.
type CleanroomsCollaboration struct {
	Name      string
	Args      CleanroomsCollaborationArgs
	state     *cleanroomsCollaborationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CleanroomsCollaboration].
func (cc *CleanroomsCollaboration) Type() string {
	return "aws_cleanrooms_collaboration"
}

// LocalName returns the local name for [CleanroomsCollaboration].
func (cc *CleanroomsCollaboration) LocalName() string {
	return cc.Name
}

// Configuration returns the configuration (args) for [CleanroomsCollaboration].
func (cc *CleanroomsCollaboration) Configuration() interface{} {
	return cc.Args
}

// DependOn is used for other resources to depend on [CleanroomsCollaboration].
func (cc *CleanroomsCollaboration) DependOn() terra.Reference {
	return terra.ReferenceResource(cc)
}

// Dependencies returns the list of resources [CleanroomsCollaboration] depends_on.
func (cc *CleanroomsCollaboration) Dependencies() terra.Dependencies {
	return cc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CleanroomsCollaboration].
func (cc *CleanroomsCollaboration) LifecycleManagement() *terra.Lifecycle {
	return cc.Lifecycle
}

// Attributes returns the attributes for [CleanroomsCollaboration].
func (cc *CleanroomsCollaboration) Attributes() cleanroomsCollaborationAttributes {
	return cleanroomsCollaborationAttributes{ref: terra.ReferenceResource(cc)}
}

// ImportState imports the given attribute values into [CleanroomsCollaboration]'s state.
func (cc *CleanroomsCollaboration) ImportState(av io.Reader) error {
	cc.state = &cleanroomsCollaborationState{}
	if err := json.NewDecoder(av).Decode(cc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cc.Type(), cc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CleanroomsCollaboration] has state.
func (cc *CleanroomsCollaboration) State() (*cleanroomsCollaborationState, bool) {
	return cc.state, cc.state != nil
}

// StateMust returns the state for [CleanroomsCollaboration]. Panics if the state is nil.
func (cc *CleanroomsCollaboration) StateMust() *cleanroomsCollaborationState {
	if cc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cc.Type(), cc.LocalName()))
	}
	return cc.state
}

// CleanroomsCollaborationArgs contains the configurations for aws_cleanrooms_collaboration.
type CleanroomsCollaborationArgs struct {
	// CreatorDisplayName: string, required
	CreatorDisplayName terra.StringValue `hcl:"creator_display_name,attr" validate:"required"`
	// CreatorMemberAbilities: list of string, required
	CreatorMemberAbilities terra.ListValue[terra.StringValue] `hcl:"creator_member_abilities,attr" validate:"required"`
	// Description: string, required
	Description terra.StringValue `hcl:"description,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// QueryLogStatus: string, required
	QueryLogStatus terra.StringValue `hcl:"query_log_status,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DataEncryptionMetadata: optional
	DataEncryptionMetadata *cleanroomscollaboration.DataEncryptionMetadata `hcl:"data_encryption_metadata,block"`
	// Member: min=0
	Member []cleanroomscollaboration.Member `hcl:"member,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *cleanroomscollaboration.Timeouts `hcl:"timeouts,block"`
}
type cleanroomsCollaborationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cleanrooms_collaboration.
func (cc cleanroomsCollaborationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("arn"))
}

// CreateTime returns a reference to field create_time of aws_cleanrooms_collaboration.
func (cc cleanroomsCollaborationAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("create_time"))
}

// CreatorDisplayName returns a reference to field creator_display_name of aws_cleanrooms_collaboration.
func (cc cleanroomsCollaborationAttributes) CreatorDisplayName() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("creator_display_name"))
}

// CreatorMemberAbilities returns a reference to field creator_member_abilities of aws_cleanrooms_collaboration.
func (cc cleanroomsCollaborationAttributes) CreatorMemberAbilities() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cc.ref.Append("creator_member_abilities"))
}

// Description returns a reference to field description of aws_cleanrooms_collaboration.
func (cc cleanroomsCollaborationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("description"))
}

// Id returns a reference to field id of aws_cleanrooms_collaboration.
func (cc cleanroomsCollaborationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("id"))
}

// Name returns a reference to field name of aws_cleanrooms_collaboration.
func (cc cleanroomsCollaborationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("name"))
}

// QueryLogStatus returns a reference to field query_log_status of aws_cleanrooms_collaboration.
func (cc cleanroomsCollaborationAttributes) QueryLogStatus() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("query_log_status"))
}

// Tags returns a reference to field tags of aws_cleanrooms_collaboration.
func (cc cleanroomsCollaborationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_cleanrooms_collaboration.
func (cc cleanroomsCollaborationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cc.ref.Append("tags_all"))
}

// UpdateTime returns a reference to field update_time of aws_cleanrooms_collaboration.
func (cc cleanroomsCollaborationAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("update_time"))
}

func (cc cleanroomsCollaborationAttributes) DataEncryptionMetadata() terra.ListValue[cleanroomscollaboration.DataEncryptionMetadataAttributes] {
	return terra.ReferenceAsList[cleanroomscollaboration.DataEncryptionMetadataAttributes](cc.ref.Append("data_encryption_metadata"))
}

func (cc cleanroomsCollaborationAttributes) Member() terra.SetValue[cleanroomscollaboration.MemberAttributes] {
	return terra.ReferenceAsSet[cleanroomscollaboration.MemberAttributes](cc.ref.Append("member"))
}

func (cc cleanroomsCollaborationAttributes) Timeouts() cleanroomscollaboration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cleanroomscollaboration.TimeoutsAttributes](cc.ref.Append("timeouts"))
}

type cleanroomsCollaborationState struct {
	Arn                    string                                                `json:"arn"`
	CreateTime             string                                                `json:"create_time"`
	CreatorDisplayName     string                                                `json:"creator_display_name"`
	CreatorMemberAbilities []string                                              `json:"creator_member_abilities"`
	Description            string                                                `json:"description"`
	Id                     string                                                `json:"id"`
	Name                   string                                                `json:"name"`
	QueryLogStatus         string                                                `json:"query_log_status"`
	Tags                   map[string]string                                     `json:"tags"`
	TagsAll                map[string]string                                     `json:"tags_all"`
	UpdateTime             string                                                `json:"update_time"`
	DataEncryptionMetadata []cleanroomscollaboration.DataEncryptionMetadataState `json:"data_encryption_metadata"`
	Member                 []cleanroomscollaboration.MemberState                 `json:"member"`
	Timeouts               *cleanroomscollaboration.TimeoutsState                `json:"timeouts"`
}
