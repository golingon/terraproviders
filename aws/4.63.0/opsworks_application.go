// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opsworksapplication "github.com/golingon/terraproviders/aws/4.63.0/opsworksapplication"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpsworksApplication creates a new instance of [OpsworksApplication].
func NewOpsworksApplication(name string, args OpsworksApplicationArgs) *OpsworksApplication {
	return &OpsworksApplication{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpsworksApplication)(nil)

// OpsworksApplication represents the Terraform resource aws_opsworks_application.
type OpsworksApplication struct {
	Name      string
	Args      OpsworksApplicationArgs
	state     *opsworksApplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpsworksApplication].
func (oa *OpsworksApplication) Type() string {
	return "aws_opsworks_application"
}

// LocalName returns the local name for [OpsworksApplication].
func (oa *OpsworksApplication) LocalName() string {
	return oa.Name
}

// Configuration returns the configuration (args) for [OpsworksApplication].
func (oa *OpsworksApplication) Configuration() interface{} {
	return oa.Args
}

// DependOn is used for other resources to depend on [OpsworksApplication].
func (oa *OpsworksApplication) DependOn() terra.Reference {
	return terra.ReferenceResource(oa)
}

// Dependencies returns the list of resources [OpsworksApplication] depends_on.
func (oa *OpsworksApplication) Dependencies() terra.Dependencies {
	return oa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpsworksApplication].
func (oa *OpsworksApplication) LifecycleManagement() *terra.Lifecycle {
	return oa.Lifecycle
}

// Attributes returns the attributes for [OpsworksApplication].
func (oa *OpsworksApplication) Attributes() opsworksApplicationAttributes {
	return opsworksApplicationAttributes{ref: terra.ReferenceResource(oa)}
}

// ImportState imports the given attribute values into [OpsworksApplication]'s state.
func (oa *OpsworksApplication) ImportState(av io.Reader) error {
	oa.state = &opsworksApplicationState{}
	if err := json.NewDecoder(av).Decode(oa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oa.Type(), oa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpsworksApplication] has state.
func (oa *OpsworksApplication) State() (*opsworksApplicationState, bool) {
	return oa.state, oa.state != nil
}

// StateMust returns the state for [OpsworksApplication]. Panics if the state is nil.
func (oa *OpsworksApplication) StateMust() *opsworksApplicationState {
	if oa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oa.Type(), oa.LocalName()))
	}
	return oa.state
}

// OpsworksApplicationArgs contains the configurations for aws_opsworks_application.
type OpsworksApplicationArgs struct {
	// AutoBundleOnDeploy: string, optional
	AutoBundleOnDeploy terra.StringValue `hcl:"auto_bundle_on_deploy,attr"`
	// AwsFlowRubySettings: string, optional
	AwsFlowRubySettings terra.StringValue `hcl:"aws_flow_ruby_settings,attr"`
	// DataSourceArn: string, optional
	DataSourceArn terra.StringValue `hcl:"data_source_arn,attr"`
	// DataSourceDatabaseName: string, optional
	DataSourceDatabaseName terra.StringValue `hcl:"data_source_database_name,attr"`
	// DataSourceType: string, optional
	DataSourceType terra.StringValue `hcl:"data_source_type,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DocumentRoot: string, optional
	DocumentRoot terra.StringValue `hcl:"document_root,attr"`
	// Domains: list of string, optional
	Domains terra.ListValue[terra.StringValue] `hcl:"domains,attr"`
	// EnableSsl: bool, optional
	EnableSsl terra.BoolValue `hcl:"enable_ssl,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RailsEnv: string, optional
	RailsEnv terra.StringValue `hcl:"rails_env,attr"`
	// ShortName: string, optional
	ShortName terra.StringValue `hcl:"short_name,attr"`
	// StackId: string, required
	StackId terra.StringValue `hcl:"stack_id,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// AppSource: min=0
	AppSource []opsworksapplication.AppSource `hcl:"app_source,block" validate:"min=0"`
	// Environment: min=0
	Environment []opsworksapplication.Environment `hcl:"environment,block" validate:"min=0"`
	// SslConfiguration: min=0
	SslConfiguration []opsworksapplication.SslConfiguration `hcl:"ssl_configuration,block" validate:"min=0"`
}
type opsworksApplicationAttributes struct {
	ref terra.Reference
}

// AutoBundleOnDeploy returns a reference to field auto_bundle_on_deploy of aws_opsworks_application.
func (oa opsworksApplicationAttributes) AutoBundleOnDeploy() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("auto_bundle_on_deploy"))
}

// AwsFlowRubySettings returns a reference to field aws_flow_ruby_settings of aws_opsworks_application.
func (oa opsworksApplicationAttributes) AwsFlowRubySettings() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("aws_flow_ruby_settings"))
}

// DataSourceArn returns a reference to field data_source_arn of aws_opsworks_application.
func (oa opsworksApplicationAttributes) DataSourceArn() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("data_source_arn"))
}

// DataSourceDatabaseName returns a reference to field data_source_database_name of aws_opsworks_application.
func (oa opsworksApplicationAttributes) DataSourceDatabaseName() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("data_source_database_name"))
}

// DataSourceType returns a reference to field data_source_type of aws_opsworks_application.
func (oa opsworksApplicationAttributes) DataSourceType() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("data_source_type"))
}

// Description returns a reference to field description of aws_opsworks_application.
func (oa opsworksApplicationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("description"))
}

// DocumentRoot returns a reference to field document_root of aws_opsworks_application.
func (oa opsworksApplicationAttributes) DocumentRoot() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("document_root"))
}

// Domains returns a reference to field domains of aws_opsworks_application.
func (oa opsworksApplicationAttributes) Domains() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oa.ref.Append("domains"))
}

// EnableSsl returns a reference to field enable_ssl of aws_opsworks_application.
func (oa opsworksApplicationAttributes) EnableSsl() terra.BoolValue {
	return terra.ReferenceAsBool(oa.ref.Append("enable_ssl"))
}

// Id returns a reference to field id of aws_opsworks_application.
func (oa opsworksApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("id"))
}

// Name returns a reference to field name of aws_opsworks_application.
func (oa opsworksApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("name"))
}

// RailsEnv returns a reference to field rails_env of aws_opsworks_application.
func (oa opsworksApplicationAttributes) RailsEnv() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("rails_env"))
}

// ShortName returns a reference to field short_name of aws_opsworks_application.
func (oa opsworksApplicationAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("short_name"))
}

// StackId returns a reference to field stack_id of aws_opsworks_application.
func (oa opsworksApplicationAttributes) StackId() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("stack_id"))
}

// Type returns a reference to field type of aws_opsworks_application.
func (oa opsworksApplicationAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("type"))
}

func (oa opsworksApplicationAttributes) AppSource() terra.ListValue[opsworksapplication.AppSourceAttributes] {
	return terra.ReferenceAsList[opsworksapplication.AppSourceAttributes](oa.ref.Append("app_source"))
}

func (oa opsworksApplicationAttributes) Environment() terra.SetValue[opsworksapplication.EnvironmentAttributes] {
	return terra.ReferenceAsSet[opsworksapplication.EnvironmentAttributes](oa.ref.Append("environment"))
}

func (oa opsworksApplicationAttributes) SslConfiguration() terra.ListValue[opsworksapplication.SslConfigurationAttributes] {
	return terra.ReferenceAsList[opsworksapplication.SslConfigurationAttributes](oa.ref.Append("ssl_configuration"))
}

type opsworksApplicationState struct {
	AutoBundleOnDeploy     string                                      `json:"auto_bundle_on_deploy"`
	AwsFlowRubySettings    string                                      `json:"aws_flow_ruby_settings"`
	DataSourceArn          string                                      `json:"data_source_arn"`
	DataSourceDatabaseName string                                      `json:"data_source_database_name"`
	DataSourceType         string                                      `json:"data_source_type"`
	Description            string                                      `json:"description"`
	DocumentRoot           string                                      `json:"document_root"`
	Domains                []string                                    `json:"domains"`
	EnableSsl              bool                                        `json:"enable_ssl"`
	Id                     string                                      `json:"id"`
	Name                   string                                      `json:"name"`
	RailsEnv               string                                      `json:"rails_env"`
	ShortName              string                                      `json:"short_name"`
	StackId                string                                      `json:"stack_id"`
	Type                   string                                      `json:"type"`
	AppSource              []opsworksapplication.AppSourceState        `json:"app_source"`
	Environment            []opsworksapplication.EnvironmentState      `json:"environment"`
	SslConfiguration       []opsworksapplication.SslConfigurationState `json:"ssl_configuration"`
}
