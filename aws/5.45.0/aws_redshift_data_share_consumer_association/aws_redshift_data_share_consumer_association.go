// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_redshift_data_share_consumer_association

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_redshift_data_share_consumer_association.
type Resource struct {
	Name      string
	Args      Args
	state     *awsRedshiftDataShareConsumerAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ardsca *Resource) Type() string {
	return "aws_redshift_data_share_consumer_association"
}

// LocalName returns the local name for [Resource].
func (ardsca *Resource) LocalName() string {
	return ardsca.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ardsca *Resource) Configuration() interface{} {
	return ardsca.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ardsca *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ardsca)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ardsca *Resource) Dependencies() terra.Dependencies {
	return ardsca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ardsca *Resource) LifecycleManagement() *terra.Lifecycle {
	return ardsca.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ardsca *Resource) Attributes() awsRedshiftDataShareConsumerAssociationAttributes {
	return awsRedshiftDataShareConsumerAssociationAttributes{ref: terra.ReferenceResource(ardsca)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ardsca *Resource) ImportState(state io.Reader) error {
	ardsca.state = &awsRedshiftDataShareConsumerAssociationState{}
	if err := json.NewDecoder(state).Decode(ardsca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ardsca.Type(), ardsca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ardsca *Resource) State() (*awsRedshiftDataShareConsumerAssociationState, bool) {
	return ardsca.state, ardsca.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ardsca *Resource) StateMust() *awsRedshiftDataShareConsumerAssociationState {
	if ardsca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ardsca.Type(), ardsca.LocalName()))
	}
	return ardsca.state
}

// Args contains the configurations for aws_redshift_data_share_consumer_association.
type Args struct {
	// AllowWrites: bool, optional
	AllowWrites terra.BoolValue `hcl:"allow_writes,attr"`
	// AssociateEntireAccount: bool, optional
	AssociateEntireAccount terra.BoolValue `hcl:"associate_entire_account,attr"`
	// ConsumerArn: string, optional
	ConsumerArn terra.StringValue `hcl:"consumer_arn,attr"`
	// ConsumerRegion: string, optional
	ConsumerRegion terra.StringValue `hcl:"consumer_region,attr"`
	// DataShareArn: string, required
	DataShareArn terra.StringValue `hcl:"data_share_arn,attr" validate:"required"`
}

type awsRedshiftDataShareConsumerAssociationAttributes struct {
	ref terra.Reference
}

// AllowWrites returns a reference to field allow_writes of aws_redshift_data_share_consumer_association.
func (ardsca awsRedshiftDataShareConsumerAssociationAttributes) AllowWrites() terra.BoolValue {
	return terra.ReferenceAsBool(ardsca.ref.Append("allow_writes"))
}

// AssociateEntireAccount returns a reference to field associate_entire_account of aws_redshift_data_share_consumer_association.
func (ardsca awsRedshiftDataShareConsumerAssociationAttributes) AssociateEntireAccount() terra.BoolValue {
	return terra.ReferenceAsBool(ardsca.ref.Append("associate_entire_account"))
}

// ConsumerArn returns a reference to field consumer_arn of aws_redshift_data_share_consumer_association.
func (ardsca awsRedshiftDataShareConsumerAssociationAttributes) ConsumerArn() terra.StringValue {
	return terra.ReferenceAsString(ardsca.ref.Append("consumer_arn"))
}

// ConsumerRegion returns a reference to field consumer_region of aws_redshift_data_share_consumer_association.
func (ardsca awsRedshiftDataShareConsumerAssociationAttributes) ConsumerRegion() terra.StringValue {
	return terra.ReferenceAsString(ardsca.ref.Append("consumer_region"))
}

// DataShareArn returns a reference to field data_share_arn of aws_redshift_data_share_consumer_association.
func (ardsca awsRedshiftDataShareConsumerAssociationAttributes) DataShareArn() terra.StringValue {
	return terra.ReferenceAsString(ardsca.ref.Append("data_share_arn"))
}

// Id returns a reference to field id of aws_redshift_data_share_consumer_association.
func (ardsca awsRedshiftDataShareConsumerAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ardsca.ref.Append("id"))
}

// ManagedBy returns a reference to field managed_by of aws_redshift_data_share_consumer_association.
func (ardsca awsRedshiftDataShareConsumerAssociationAttributes) ManagedBy() terra.StringValue {
	return terra.ReferenceAsString(ardsca.ref.Append("managed_by"))
}

// ProducerArn returns a reference to field producer_arn of aws_redshift_data_share_consumer_association.
func (ardsca awsRedshiftDataShareConsumerAssociationAttributes) ProducerArn() terra.StringValue {
	return terra.ReferenceAsString(ardsca.ref.Append("producer_arn"))
}

type awsRedshiftDataShareConsumerAssociationState struct {
	AllowWrites            bool   `json:"allow_writes"`
	AssociateEntireAccount bool   `json:"associate_entire_account"`
	ConsumerArn            string `json:"consumer_arn"`
	ConsumerRegion         string `json:"consumer_region"`
	DataShareArn           string `json:"data_share_arn"`
	Id                     string `json:"id"`
	ManagedBy              string `json:"managed_by"`
	ProducerArn            string `json:"producer_arn"`
}
