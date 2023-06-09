// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewQuicksightIngestion creates a new instance of [QuicksightIngestion].
func NewQuicksightIngestion(name string, args QuicksightIngestionArgs) *QuicksightIngestion {
	return &QuicksightIngestion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*QuicksightIngestion)(nil)

// QuicksightIngestion represents the Terraform resource aws_quicksight_ingestion.
type QuicksightIngestion struct {
	Name      string
	Args      QuicksightIngestionArgs
	state     *quicksightIngestionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [QuicksightIngestion].
func (qi *QuicksightIngestion) Type() string {
	return "aws_quicksight_ingestion"
}

// LocalName returns the local name for [QuicksightIngestion].
func (qi *QuicksightIngestion) LocalName() string {
	return qi.Name
}

// Configuration returns the configuration (args) for [QuicksightIngestion].
func (qi *QuicksightIngestion) Configuration() interface{} {
	return qi.Args
}

// DependOn is used for other resources to depend on [QuicksightIngestion].
func (qi *QuicksightIngestion) DependOn() terra.Reference {
	return terra.ReferenceResource(qi)
}

// Dependencies returns the list of resources [QuicksightIngestion] depends_on.
func (qi *QuicksightIngestion) Dependencies() terra.Dependencies {
	return qi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [QuicksightIngestion].
func (qi *QuicksightIngestion) LifecycleManagement() *terra.Lifecycle {
	return qi.Lifecycle
}

// Attributes returns the attributes for [QuicksightIngestion].
func (qi *QuicksightIngestion) Attributes() quicksightIngestionAttributes {
	return quicksightIngestionAttributes{ref: terra.ReferenceResource(qi)}
}

// ImportState imports the given attribute values into [QuicksightIngestion]'s state.
func (qi *QuicksightIngestion) ImportState(av io.Reader) error {
	qi.state = &quicksightIngestionState{}
	if err := json.NewDecoder(av).Decode(qi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", qi.Type(), qi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [QuicksightIngestion] has state.
func (qi *QuicksightIngestion) State() (*quicksightIngestionState, bool) {
	return qi.state, qi.state != nil
}

// StateMust returns the state for [QuicksightIngestion]. Panics if the state is nil.
func (qi *QuicksightIngestion) StateMust() *quicksightIngestionState {
	if qi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", qi.Type(), qi.LocalName()))
	}
	return qi.state
}

// QuicksightIngestionArgs contains the configurations for aws_quicksight_ingestion.
type QuicksightIngestionArgs struct {
	// AwsAccountId: string, optional
	AwsAccountId terra.StringValue `hcl:"aws_account_id,attr"`
	// DataSetId: string, required
	DataSetId terra.StringValue `hcl:"data_set_id,attr" validate:"required"`
	// IngestionId: string, required
	IngestionId terra.StringValue `hcl:"ingestion_id,attr" validate:"required"`
	// IngestionType: string, required
	IngestionType terra.StringValue `hcl:"ingestion_type,attr" validate:"required"`
}
type quicksightIngestionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_quicksight_ingestion.
func (qi quicksightIngestionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("arn"))
}

// AwsAccountId returns a reference to field aws_account_id of aws_quicksight_ingestion.
func (qi quicksightIngestionAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("aws_account_id"))
}

// DataSetId returns a reference to field data_set_id of aws_quicksight_ingestion.
func (qi quicksightIngestionAttributes) DataSetId() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("data_set_id"))
}

// Id returns a reference to field id of aws_quicksight_ingestion.
func (qi quicksightIngestionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("id"))
}

// IngestionId returns a reference to field ingestion_id of aws_quicksight_ingestion.
func (qi quicksightIngestionAttributes) IngestionId() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("ingestion_id"))
}

// IngestionStatus returns a reference to field ingestion_status of aws_quicksight_ingestion.
func (qi quicksightIngestionAttributes) IngestionStatus() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("ingestion_status"))
}

// IngestionType returns a reference to field ingestion_type of aws_quicksight_ingestion.
func (qi quicksightIngestionAttributes) IngestionType() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("ingestion_type"))
}

type quicksightIngestionState struct {
	Arn             string `json:"arn"`
	AwsAccountId    string `json:"aws_account_id"`
	DataSetId       string `json:"data_set_id"`
	Id              string `json:"id"`
	IngestionId     string `json:"ingestion_id"`
	IngestionStatus string `json:"ingestion_status"`
	IngestionType   string `json:"ingestion_type"`
}
