// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	wafv2webaclloggingconfiguration "github.com/golingon/terraproviders/aws/5.44.0/wafv2webaclloggingconfiguration"
	"io"
)

// NewWafv2WebAclLoggingConfiguration creates a new instance of [Wafv2WebAclLoggingConfiguration].
func NewWafv2WebAclLoggingConfiguration(name string, args Wafv2WebAclLoggingConfigurationArgs) *Wafv2WebAclLoggingConfiguration {
	return &Wafv2WebAclLoggingConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Wafv2WebAclLoggingConfiguration)(nil)

// Wafv2WebAclLoggingConfiguration represents the Terraform resource aws_wafv2_web_acl_logging_configuration.
type Wafv2WebAclLoggingConfiguration struct {
	Name      string
	Args      Wafv2WebAclLoggingConfigurationArgs
	state     *wafv2WebAclLoggingConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Wafv2WebAclLoggingConfiguration].
func (wwalc *Wafv2WebAclLoggingConfiguration) Type() string {
	return "aws_wafv2_web_acl_logging_configuration"
}

// LocalName returns the local name for [Wafv2WebAclLoggingConfiguration].
func (wwalc *Wafv2WebAclLoggingConfiguration) LocalName() string {
	return wwalc.Name
}

// Configuration returns the configuration (args) for [Wafv2WebAclLoggingConfiguration].
func (wwalc *Wafv2WebAclLoggingConfiguration) Configuration() interface{} {
	return wwalc.Args
}

// DependOn is used for other resources to depend on [Wafv2WebAclLoggingConfiguration].
func (wwalc *Wafv2WebAclLoggingConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(wwalc)
}

// Dependencies returns the list of resources [Wafv2WebAclLoggingConfiguration] depends_on.
func (wwalc *Wafv2WebAclLoggingConfiguration) Dependencies() terra.Dependencies {
	return wwalc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Wafv2WebAclLoggingConfiguration].
func (wwalc *Wafv2WebAclLoggingConfiguration) LifecycleManagement() *terra.Lifecycle {
	return wwalc.Lifecycle
}

// Attributes returns the attributes for [Wafv2WebAclLoggingConfiguration].
func (wwalc *Wafv2WebAclLoggingConfiguration) Attributes() wafv2WebAclLoggingConfigurationAttributes {
	return wafv2WebAclLoggingConfigurationAttributes{ref: terra.ReferenceResource(wwalc)}
}

// ImportState imports the given attribute values into [Wafv2WebAclLoggingConfiguration]'s state.
func (wwalc *Wafv2WebAclLoggingConfiguration) ImportState(av io.Reader) error {
	wwalc.state = &wafv2WebAclLoggingConfigurationState{}
	if err := json.NewDecoder(av).Decode(wwalc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wwalc.Type(), wwalc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Wafv2WebAclLoggingConfiguration] has state.
func (wwalc *Wafv2WebAclLoggingConfiguration) State() (*wafv2WebAclLoggingConfigurationState, bool) {
	return wwalc.state, wwalc.state != nil
}

// StateMust returns the state for [Wafv2WebAclLoggingConfiguration]. Panics if the state is nil.
func (wwalc *Wafv2WebAclLoggingConfiguration) StateMust() *wafv2WebAclLoggingConfigurationState {
	if wwalc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wwalc.Type(), wwalc.LocalName()))
	}
	return wwalc.state
}

// Wafv2WebAclLoggingConfigurationArgs contains the configurations for aws_wafv2_web_acl_logging_configuration.
type Wafv2WebAclLoggingConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogDestinationConfigs: set of string, required
	LogDestinationConfigs terra.SetValue[terra.StringValue] `hcl:"log_destination_configs,attr" validate:"required"`
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
	// LoggingFilter: optional
	LoggingFilter *wafv2webaclloggingconfiguration.LoggingFilter `hcl:"logging_filter,block"`
	// RedactedFields: min=0,max=100
	RedactedFields []wafv2webaclloggingconfiguration.RedactedFields `hcl:"redacted_fields,block" validate:"min=0,max=100"`
}
type wafv2WebAclLoggingConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_wafv2_web_acl_logging_configuration.
func (wwalc wafv2WebAclLoggingConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwalc.ref.Append("id"))
}

// LogDestinationConfigs returns a reference to field log_destination_configs of aws_wafv2_web_acl_logging_configuration.
func (wwalc wafv2WebAclLoggingConfigurationAttributes) LogDestinationConfigs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wwalc.ref.Append("log_destination_configs"))
}

// ResourceArn returns a reference to field resource_arn of aws_wafv2_web_acl_logging_configuration.
func (wwalc wafv2WebAclLoggingConfigurationAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(wwalc.ref.Append("resource_arn"))
}

func (wwalc wafv2WebAclLoggingConfigurationAttributes) LoggingFilter() terra.ListValue[wafv2webaclloggingconfiguration.LoggingFilterAttributes] {
	return terra.ReferenceAsList[wafv2webaclloggingconfiguration.LoggingFilterAttributes](wwalc.ref.Append("logging_filter"))
}

func (wwalc wafv2WebAclLoggingConfigurationAttributes) RedactedFields() terra.ListValue[wafv2webaclloggingconfiguration.RedactedFieldsAttributes] {
	return terra.ReferenceAsList[wafv2webaclloggingconfiguration.RedactedFieldsAttributes](wwalc.ref.Append("redacted_fields"))
}

type wafv2WebAclLoggingConfigurationState struct {
	Id                    string                                                `json:"id"`
	LogDestinationConfigs []string                                              `json:"log_destination_configs"`
	ResourceArn           string                                                `json:"resource_arn"`
	LoggingFilter         []wafv2webaclloggingconfiguration.LoggingFilterState  `json:"logging_filter"`
	RedactedFields        []wafv2webaclloggingconfiguration.RedactedFieldsState `json:"redacted_fields"`
}
