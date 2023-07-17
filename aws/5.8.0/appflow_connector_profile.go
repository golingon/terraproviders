// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appflowconnectorprofile "github.com/golingon/terraproviders/aws/5.8.0/appflowconnectorprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppflowConnectorProfile creates a new instance of [AppflowConnectorProfile].
func NewAppflowConnectorProfile(name string, args AppflowConnectorProfileArgs) *AppflowConnectorProfile {
	return &AppflowConnectorProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppflowConnectorProfile)(nil)

// AppflowConnectorProfile represents the Terraform resource aws_appflow_connector_profile.
type AppflowConnectorProfile struct {
	Name      string
	Args      AppflowConnectorProfileArgs
	state     *appflowConnectorProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppflowConnectorProfile].
func (acp *AppflowConnectorProfile) Type() string {
	return "aws_appflow_connector_profile"
}

// LocalName returns the local name for [AppflowConnectorProfile].
func (acp *AppflowConnectorProfile) LocalName() string {
	return acp.Name
}

// Configuration returns the configuration (args) for [AppflowConnectorProfile].
func (acp *AppflowConnectorProfile) Configuration() interface{} {
	return acp.Args
}

// DependOn is used for other resources to depend on [AppflowConnectorProfile].
func (acp *AppflowConnectorProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(acp)
}

// Dependencies returns the list of resources [AppflowConnectorProfile] depends_on.
func (acp *AppflowConnectorProfile) Dependencies() terra.Dependencies {
	return acp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppflowConnectorProfile].
func (acp *AppflowConnectorProfile) LifecycleManagement() *terra.Lifecycle {
	return acp.Lifecycle
}

// Attributes returns the attributes for [AppflowConnectorProfile].
func (acp *AppflowConnectorProfile) Attributes() appflowConnectorProfileAttributes {
	return appflowConnectorProfileAttributes{ref: terra.ReferenceResource(acp)}
}

// ImportState imports the given attribute values into [AppflowConnectorProfile]'s state.
func (acp *AppflowConnectorProfile) ImportState(av io.Reader) error {
	acp.state = &appflowConnectorProfileState{}
	if err := json.NewDecoder(av).Decode(acp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acp.Type(), acp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppflowConnectorProfile] has state.
func (acp *AppflowConnectorProfile) State() (*appflowConnectorProfileState, bool) {
	return acp.state, acp.state != nil
}

// StateMust returns the state for [AppflowConnectorProfile]. Panics if the state is nil.
func (acp *AppflowConnectorProfile) StateMust() *appflowConnectorProfileState {
	if acp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acp.Type(), acp.LocalName()))
	}
	return acp.state
}

// AppflowConnectorProfileArgs contains the configurations for aws_appflow_connector_profile.
type AppflowConnectorProfileArgs struct {
	// ConnectionMode: string, required
	ConnectionMode terra.StringValue `hcl:"connection_mode,attr" validate:"required"`
	// ConnectorLabel: string, optional
	ConnectorLabel terra.StringValue `hcl:"connector_label,attr"`
	// ConnectorType: string, required
	ConnectorType terra.StringValue `hcl:"connector_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsArn: string, optional
	KmsArn terra.StringValue `hcl:"kms_arn,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ConnectorProfileConfig: required
	ConnectorProfileConfig *appflowconnectorprofile.ConnectorProfileConfig `hcl:"connector_profile_config,block" validate:"required"`
}
type appflowConnectorProfileAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appflow_connector_profile.
func (acp appflowConnectorProfileAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("arn"))
}

// ConnectionMode returns a reference to field connection_mode of aws_appflow_connector_profile.
func (acp appflowConnectorProfileAttributes) ConnectionMode() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("connection_mode"))
}

// ConnectorLabel returns a reference to field connector_label of aws_appflow_connector_profile.
func (acp appflowConnectorProfileAttributes) ConnectorLabel() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("connector_label"))
}

// ConnectorType returns a reference to field connector_type of aws_appflow_connector_profile.
func (acp appflowConnectorProfileAttributes) ConnectorType() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("connector_type"))
}

// CredentialsArn returns a reference to field credentials_arn of aws_appflow_connector_profile.
func (acp appflowConnectorProfileAttributes) CredentialsArn() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("credentials_arn"))
}

// Id returns a reference to field id of aws_appflow_connector_profile.
func (acp appflowConnectorProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("id"))
}

// KmsArn returns a reference to field kms_arn of aws_appflow_connector_profile.
func (acp appflowConnectorProfileAttributes) KmsArn() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("kms_arn"))
}

// Name returns a reference to field name of aws_appflow_connector_profile.
func (acp appflowConnectorProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("name"))
}

func (acp appflowConnectorProfileAttributes) ConnectorProfileConfig() terra.ListValue[appflowconnectorprofile.ConnectorProfileConfigAttributes] {
	return terra.ReferenceAsList[appflowconnectorprofile.ConnectorProfileConfigAttributes](acp.ref.Append("connector_profile_config"))
}

type appflowConnectorProfileState struct {
	Arn                    string                                                `json:"arn"`
	ConnectionMode         string                                                `json:"connection_mode"`
	ConnectorLabel         string                                                `json:"connector_label"`
	ConnectorType          string                                                `json:"connector_type"`
	CredentialsArn         string                                                `json:"credentials_arn"`
	Id                     string                                                `json:"id"`
	KmsArn                 string                                                `json:"kms_arn"`
	Name                   string                                                `json:"name"`
	ConnectorProfileConfig []appflowconnectorprofile.ConnectorProfileConfigState `json:"connector_profile_config"`
}
