// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectorawss3 "github.com/golingon/terraproviders/azurerm/3.65.0/sentineldataconnectorawss3"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelDataConnectorAwsS3 creates a new instance of [SentinelDataConnectorAwsS3].
func NewSentinelDataConnectorAwsS3(name string, args SentinelDataConnectorAwsS3Args) *SentinelDataConnectorAwsS3 {
	return &SentinelDataConnectorAwsS3{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorAwsS3)(nil)

// SentinelDataConnectorAwsS3 represents the Terraform resource azurerm_sentinel_data_connector_aws_s3.
type SentinelDataConnectorAwsS3 struct {
	Name      string
	Args      SentinelDataConnectorAwsS3Args
	state     *sentinelDataConnectorAwsS3State
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorAwsS3].
func (sdcas *SentinelDataConnectorAwsS3) Type() string {
	return "azurerm_sentinel_data_connector_aws_s3"
}

// LocalName returns the local name for [SentinelDataConnectorAwsS3].
func (sdcas *SentinelDataConnectorAwsS3) LocalName() string {
	return sdcas.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorAwsS3].
func (sdcas *SentinelDataConnectorAwsS3) Configuration() interface{} {
	return sdcas.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorAwsS3].
func (sdcas *SentinelDataConnectorAwsS3) DependOn() terra.Reference {
	return terra.ReferenceResource(sdcas)
}

// Dependencies returns the list of resources [SentinelDataConnectorAwsS3] depends_on.
func (sdcas *SentinelDataConnectorAwsS3) Dependencies() terra.Dependencies {
	return sdcas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorAwsS3].
func (sdcas *SentinelDataConnectorAwsS3) LifecycleManagement() *terra.Lifecycle {
	return sdcas.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorAwsS3].
func (sdcas *SentinelDataConnectorAwsS3) Attributes() sentinelDataConnectorAwsS3Attributes {
	return sentinelDataConnectorAwsS3Attributes{ref: terra.ReferenceResource(sdcas)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorAwsS3]'s state.
func (sdcas *SentinelDataConnectorAwsS3) ImportState(av io.Reader) error {
	sdcas.state = &sentinelDataConnectorAwsS3State{}
	if err := json.NewDecoder(av).Decode(sdcas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdcas.Type(), sdcas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorAwsS3] has state.
func (sdcas *SentinelDataConnectorAwsS3) State() (*sentinelDataConnectorAwsS3State, bool) {
	return sdcas.state, sdcas.state != nil
}

// StateMust returns the state for [SentinelDataConnectorAwsS3]. Panics if the state is nil.
func (sdcas *SentinelDataConnectorAwsS3) StateMust() *sentinelDataConnectorAwsS3State {
	if sdcas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdcas.Type(), sdcas.LocalName()))
	}
	return sdcas.state
}

// SentinelDataConnectorAwsS3Args contains the configurations for azurerm_sentinel_data_connector_aws_s3.
type SentinelDataConnectorAwsS3Args struct {
	// AwsRoleArn: string, required
	AwsRoleArn terra.StringValue `hcl:"aws_role_arn,attr" validate:"required"`
	// DestinationTable: string, required
	DestinationTable terra.StringValue `hcl:"destination_table,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SqsUrls: list of string, required
	SqsUrls terra.ListValue[terra.StringValue] `hcl:"sqs_urls,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *sentineldataconnectorawss3.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorAwsS3Attributes struct {
	ref terra.Reference
}

// AwsRoleArn returns a reference to field aws_role_arn of azurerm_sentinel_data_connector_aws_s3.
func (sdcas sentinelDataConnectorAwsS3Attributes) AwsRoleArn() terra.StringValue {
	return terra.ReferenceAsString(sdcas.ref.Append("aws_role_arn"))
}

// DestinationTable returns a reference to field destination_table of azurerm_sentinel_data_connector_aws_s3.
func (sdcas sentinelDataConnectorAwsS3Attributes) DestinationTable() terra.StringValue {
	return terra.ReferenceAsString(sdcas.ref.Append("destination_table"))
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_aws_s3.
func (sdcas sentinelDataConnectorAwsS3Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdcas.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_aws_s3.
func (sdcas sentinelDataConnectorAwsS3Attributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdcas.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_aws_s3.
func (sdcas sentinelDataConnectorAwsS3Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdcas.ref.Append("name"))
}

// SqsUrls returns a reference to field sqs_urls of azurerm_sentinel_data_connector_aws_s3.
func (sdcas sentinelDataConnectorAwsS3Attributes) SqsUrls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sdcas.ref.Append("sqs_urls"))
}

func (sdcas sentinelDataConnectorAwsS3Attributes) Timeouts() sentineldataconnectorawss3.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectorawss3.TimeoutsAttributes](sdcas.ref.Append("timeouts"))
}

type sentinelDataConnectorAwsS3State struct {
	AwsRoleArn              string                                    `json:"aws_role_arn"`
	DestinationTable        string                                    `json:"destination_table"`
	Id                      string                                    `json:"id"`
	LogAnalyticsWorkspaceId string                                    `json:"log_analytics_workspace_id"`
	Name                    string                                    `json:"name"`
	SqsUrls                 []string                                  `json:"sqs_urls"`
	Timeouts                *sentineldataconnectorawss3.TimeoutsState `json:"timeouts"`
}
