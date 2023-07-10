// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	codegurureviewerrepositoryassociation "github.com/golingon/terraproviders/aws/5.7.0/codegurureviewerrepositoryassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodegurureviewerRepositoryAssociation creates a new instance of [CodegurureviewerRepositoryAssociation].
func NewCodegurureviewerRepositoryAssociation(name string, args CodegurureviewerRepositoryAssociationArgs) *CodegurureviewerRepositoryAssociation {
	return &CodegurureviewerRepositoryAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodegurureviewerRepositoryAssociation)(nil)

// CodegurureviewerRepositoryAssociation represents the Terraform resource aws_codegurureviewer_repository_association.
type CodegurureviewerRepositoryAssociation struct {
	Name      string
	Args      CodegurureviewerRepositoryAssociationArgs
	state     *codegurureviewerRepositoryAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodegurureviewerRepositoryAssociation].
func (cra *CodegurureviewerRepositoryAssociation) Type() string {
	return "aws_codegurureviewer_repository_association"
}

// LocalName returns the local name for [CodegurureviewerRepositoryAssociation].
func (cra *CodegurureviewerRepositoryAssociation) LocalName() string {
	return cra.Name
}

// Configuration returns the configuration (args) for [CodegurureviewerRepositoryAssociation].
func (cra *CodegurureviewerRepositoryAssociation) Configuration() interface{} {
	return cra.Args
}

// DependOn is used for other resources to depend on [CodegurureviewerRepositoryAssociation].
func (cra *CodegurureviewerRepositoryAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(cra)
}

// Dependencies returns the list of resources [CodegurureviewerRepositoryAssociation] depends_on.
func (cra *CodegurureviewerRepositoryAssociation) Dependencies() terra.Dependencies {
	return cra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodegurureviewerRepositoryAssociation].
func (cra *CodegurureviewerRepositoryAssociation) LifecycleManagement() *terra.Lifecycle {
	return cra.Lifecycle
}

// Attributes returns the attributes for [CodegurureviewerRepositoryAssociation].
func (cra *CodegurureviewerRepositoryAssociation) Attributes() codegurureviewerRepositoryAssociationAttributes {
	return codegurureviewerRepositoryAssociationAttributes{ref: terra.ReferenceResource(cra)}
}

// ImportState imports the given attribute values into [CodegurureviewerRepositoryAssociation]'s state.
func (cra *CodegurureviewerRepositoryAssociation) ImportState(av io.Reader) error {
	cra.state = &codegurureviewerRepositoryAssociationState{}
	if err := json.NewDecoder(av).Decode(cra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cra.Type(), cra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodegurureviewerRepositoryAssociation] has state.
func (cra *CodegurureviewerRepositoryAssociation) State() (*codegurureviewerRepositoryAssociationState, bool) {
	return cra.state, cra.state != nil
}

// StateMust returns the state for [CodegurureviewerRepositoryAssociation]. Panics if the state is nil.
func (cra *CodegurureviewerRepositoryAssociation) StateMust() *codegurureviewerRepositoryAssociationState {
	if cra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cra.Type(), cra.LocalName()))
	}
	return cra.state
}

// CodegurureviewerRepositoryAssociationArgs contains the configurations for aws_codegurureviewer_repository_association.
type CodegurureviewerRepositoryAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// S3RepositoryDetails: min=0
	S3RepositoryDetails []codegurureviewerrepositoryassociation.S3RepositoryDetails `hcl:"s3_repository_details,block" validate:"min=0"`
	// KmsKeyDetails: optional
	KmsKeyDetails *codegurureviewerrepositoryassociation.KmsKeyDetails `hcl:"kms_key_details,block"`
	// Repository: required
	Repository *codegurureviewerrepositoryassociation.Repository `hcl:"repository,block" validate:"required"`
	// Timeouts: optional
	Timeouts *codegurureviewerrepositoryassociation.Timeouts `hcl:"timeouts,block"`
}
type codegurureviewerRepositoryAssociationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_codegurureviewer_repository_association.
func (cra codegurureviewerRepositoryAssociationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("arn"))
}

// AssociationId returns a reference to field association_id of aws_codegurureviewer_repository_association.
func (cra codegurureviewerRepositoryAssociationAttributes) AssociationId() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("association_id"))
}

// ConnectionArn returns a reference to field connection_arn of aws_codegurureviewer_repository_association.
func (cra codegurureviewerRepositoryAssociationAttributes) ConnectionArn() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("connection_arn"))
}

// Id returns a reference to field id of aws_codegurureviewer_repository_association.
func (cra codegurureviewerRepositoryAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("id"))
}

// Name returns a reference to field name of aws_codegurureviewer_repository_association.
func (cra codegurureviewerRepositoryAssociationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_codegurureviewer_repository_association.
func (cra codegurureviewerRepositoryAssociationAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("owner"))
}

// ProviderType returns a reference to field provider_type of aws_codegurureviewer_repository_association.
func (cra codegurureviewerRepositoryAssociationAttributes) ProviderType() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("provider_type"))
}

// State returns a reference to field state of aws_codegurureviewer_repository_association.
func (cra codegurureviewerRepositoryAssociationAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("state"))
}

// StateReason returns a reference to field state_reason of aws_codegurureviewer_repository_association.
func (cra codegurureviewerRepositoryAssociationAttributes) StateReason() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("state_reason"))
}

// Tags returns a reference to field tags of aws_codegurureviewer_repository_association.
func (cra codegurureviewerRepositoryAssociationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cra.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_codegurureviewer_repository_association.
func (cra codegurureviewerRepositoryAssociationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cra.ref.Append("tags_all"))
}

func (cra codegurureviewerRepositoryAssociationAttributes) S3RepositoryDetails() terra.ListValue[codegurureviewerrepositoryassociation.S3RepositoryDetailsAttributes] {
	return terra.ReferenceAsList[codegurureviewerrepositoryassociation.S3RepositoryDetailsAttributes](cra.ref.Append("s3_repository_details"))
}

func (cra codegurureviewerRepositoryAssociationAttributes) KmsKeyDetails() terra.ListValue[codegurureviewerrepositoryassociation.KmsKeyDetailsAttributes] {
	return terra.ReferenceAsList[codegurureviewerrepositoryassociation.KmsKeyDetailsAttributes](cra.ref.Append("kms_key_details"))
}

func (cra codegurureviewerRepositoryAssociationAttributes) Repository() terra.ListValue[codegurureviewerrepositoryassociation.RepositoryAttributes] {
	return terra.ReferenceAsList[codegurureviewerrepositoryassociation.RepositoryAttributes](cra.ref.Append("repository"))
}

func (cra codegurureviewerRepositoryAssociationAttributes) Timeouts() codegurureviewerrepositoryassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[codegurureviewerrepositoryassociation.TimeoutsAttributes](cra.ref.Append("timeouts"))
}

type codegurureviewerRepositoryAssociationState struct {
	Arn                 string                                                           `json:"arn"`
	AssociationId       string                                                           `json:"association_id"`
	ConnectionArn       string                                                           `json:"connection_arn"`
	Id                  string                                                           `json:"id"`
	Name                string                                                           `json:"name"`
	Owner               string                                                           `json:"owner"`
	ProviderType        string                                                           `json:"provider_type"`
	State               string                                                           `json:"state"`
	StateReason         string                                                           `json:"state_reason"`
	Tags                map[string]string                                                `json:"tags"`
	TagsAll             map[string]string                                                `json:"tags_all"`
	S3RepositoryDetails []codegurureviewerrepositoryassociation.S3RepositoryDetailsState `json:"s3_repository_details"`
	KmsKeyDetails       []codegurureviewerrepositoryassociation.KmsKeyDetailsState       `json:"kms_key_details"`
	Repository          []codegurureviewerrepositoryassociation.RepositoryState          `json:"repository"`
	Timeouts            *codegurureviewerrepositoryassociation.TimeoutsState             `json:"timeouts"`
}
