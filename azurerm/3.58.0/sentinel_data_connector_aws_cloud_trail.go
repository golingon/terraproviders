// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectorawscloudtrail "github.com/golingon/terraproviders/azurerm/3.58.0/sentineldataconnectorawscloudtrail"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelDataConnectorAwsCloudTrail creates a new instance of [SentinelDataConnectorAwsCloudTrail].
func NewSentinelDataConnectorAwsCloudTrail(name string, args SentinelDataConnectorAwsCloudTrailArgs) *SentinelDataConnectorAwsCloudTrail {
	return &SentinelDataConnectorAwsCloudTrail{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorAwsCloudTrail)(nil)

// SentinelDataConnectorAwsCloudTrail represents the Terraform resource azurerm_sentinel_data_connector_aws_cloud_trail.
type SentinelDataConnectorAwsCloudTrail struct {
	Name      string
	Args      SentinelDataConnectorAwsCloudTrailArgs
	state     *sentinelDataConnectorAwsCloudTrailState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorAwsCloudTrail].
func (sdcact *SentinelDataConnectorAwsCloudTrail) Type() string {
	return "azurerm_sentinel_data_connector_aws_cloud_trail"
}

// LocalName returns the local name for [SentinelDataConnectorAwsCloudTrail].
func (sdcact *SentinelDataConnectorAwsCloudTrail) LocalName() string {
	return sdcact.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorAwsCloudTrail].
func (sdcact *SentinelDataConnectorAwsCloudTrail) Configuration() interface{} {
	return sdcact.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorAwsCloudTrail].
func (sdcact *SentinelDataConnectorAwsCloudTrail) DependOn() terra.Reference {
	return terra.ReferenceResource(sdcact)
}

// Dependencies returns the list of resources [SentinelDataConnectorAwsCloudTrail] depends_on.
func (sdcact *SentinelDataConnectorAwsCloudTrail) Dependencies() terra.Dependencies {
	return sdcact.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorAwsCloudTrail].
func (sdcact *SentinelDataConnectorAwsCloudTrail) LifecycleManagement() *terra.Lifecycle {
	return sdcact.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorAwsCloudTrail].
func (sdcact *SentinelDataConnectorAwsCloudTrail) Attributes() sentinelDataConnectorAwsCloudTrailAttributes {
	return sentinelDataConnectorAwsCloudTrailAttributes{ref: terra.ReferenceResource(sdcact)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorAwsCloudTrail]'s state.
func (sdcact *SentinelDataConnectorAwsCloudTrail) ImportState(av io.Reader) error {
	sdcact.state = &sentinelDataConnectorAwsCloudTrailState{}
	if err := json.NewDecoder(av).Decode(sdcact.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdcact.Type(), sdcact.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorAwsCloudTrail] has state.
func (sdcact *SentinelDataConnectorAwsCloudTrail) State() (*sentinelDataConnectorAwsCloudTrailState, bool) {
	return sdcact.state, sdcact.state != nil
}

// StateMust returns the state for [SentinelDataConnectorAwsCloudTrail]. Panics if the state is nil.
func (sdcact *SentinelDataConnectorAwsCloudTrail) StateMust() *sentinelDataConnectorAwsCloudTrailState {
	if sdcact.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdcact.Type(), sdcact.LocalName()))
	}
	return sdcact.state
}

// SentinelDataConnectorAwsCloudTrailArgs contains the configurations for azurerm_sentinel_data_connector_aws_cloud_trail.
type SentinelDataConnectorAwsCloudTrailArgs struct {
	// AwsRoleArn: string, required
	AwsRoleArn terra.StringValue `hcl:"aws_role_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *sentineldataconnectorawscloudtrail.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorAwsCloudTrailAttributes struct {
	ref terra.Reference
}

// AwsRoleArn returns a reference to field aws_role_arn of azurerm_sentinel_data_connector_aws_cloud_trail.
func (sdcact sentinelDataConnectorAwsCloudTrailAttributes) AwsRoleArn() terra.StringValue {
	return terra.ReferenceAsString(sdcact.ref.Append("aws_role_arn"))
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_aws_cloud_trail.
func (sdcact sentinelDataConnectorAwsCloudTrailAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdcact.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_aws_cloud_trail.
func (sdcact sentinelDataConnectorAwsCloudTrailAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdcact.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_aws_cloud_trail.
func (sdcact sentinelDataConnectorAwsCloudTrailAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdcact.ref.Append("name"))
}

func (sdcact sentinelDataConnectorAwsCloudTrailAttributes) Timeouts() sentineldataconnectorawscloudtrail.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectorawscloudtrail.TimeoutsAttributes](sdcact.ref.Append("timeouts"))
}

type sentinelDataConnectorAwsCloudTrailState struct {
	AwsRoleArn              string                                            `json:"aws_role_arn"`
	Id                      string                                            `json:"id"`
	LogAnalyticsWorkspaceId string                                            `json:"log_analytics_workspace_id"`
	Name                    string                                            `json:"name"`
	Timeouts                *sentineldataconnectorawscloudtrail.TimeoutsState `json:"timeouts"`
}
