// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sesconfigurationset "github.com/golingon/terraproviders/aws/5.9.0/sesconfigurationset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSesConfigurationSet creates a new instance of [SesConfigurationSet].
func NewSesConfigurationSet(name string, args SesConfigurationSetArgs) *SesConfigurationSet {
	return &SesConfigurationSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SesConfigurationSet)(nil)

// SesConfigurationSet represents the Terraform resource aws_ses_configuration_set.
type SesConfigurationSet struct {
	Name      string
	Args      SesConfigurationSetArgs
	state     *sesConfigurationSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SesConfigurationSet].
func (scs *SesConfigurationSet) Type() string {
	return "aws_ses_configuration_set"
}

// LocalName returns the local name for [SesConfigurationSet].
func (scs *SesConfigurationSet) LocalName() string {
	return scs.Name
}

// Configuration returns the configuration (args) for [SesConfigurationSet].
func (scs *SesConfigurationSet) Configuration() interface{} {
	return scs.Args
}

// DependOn is used for other resources to depend on [SesConfigurationSet].
func (scs *SesConfigurationSet) DependOn() terra.Reference {
	return terra.ReferenceResource(scs)
}

// Dependencies returns the list of resources [SesConfigurationSet] depends_on.
func (scs *SesConfigurationSet) Dependencies() terra.Dependencies {
	return scs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SesConfigurationSet].
func (scs *SesConfigurationSet) LifecycleManagement() *terra.Lifecycle {
	return scs.Lifecycle
}

// Attributes returns the attributes for [SesConfigurationSet].
func (scs *SesConfigurationSet) Attributes() sesConfigurationSetAttributes {
	return sesConfigurationSetAttributes{ref: terra.ReferenceResource(scs)}
}

// ImportState imports the given attribute values into [SesConfigurationSet]'s state.
func (scs *SesConfigurationSet) ImportState(av io.Reader) error {
	scs.state = &sesConfigurationSetState{}
	if err := json.NewDecoder(av).Decode(scs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scs.Type(), scs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SesConfigurationSet] has state.
func (scs *SesConfigurationSet) State() (*sesConfigurationSetState, bool) {
	return scs.state, scs.state != nil
}

// StateMust returns the state for [SesConfigurationSet]. Panics if the state is nil.
func (scs *SesConfigurationSet) StateMust() *sesConfigurationSetState {
	if scs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scs.Type(), scs.LocalName()))
	}
	return scs.state
}

// SesConfigurationSetArgs contains the configurations for aws_ses_configuration_set.
type SesConfigurationSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ReputationMetricsEnabled: bool, optional
	ReputationMetricsEnabled terra.BoolValue `hcl:"reputation_metrics_enabled,attr"`
	// SendingEnabled: bool, optional
	SendingEnabled terra.BoolValue `hcl:"sending_enabled,attr"`
	// DeliveryOptions: optional
	DeliveryOptions *sesconfigurationset.DeliveryOptions `hcl:"delivery_options,block"`
	// TrackingOptions: optional
	TrackingOptions *sesconfigurationset.TrackingOptions `hcl:"tracking_options,block"`
}
type sesConfigurationSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ses_configuration_set.
func (scs sesConfigurationSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ses_configuration_set.
func (scs sesConfigurationSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("id"))
}

// LastFreshStart returns a reference to field last_fresh_start of aws_ses_configuration_set.
func (scs sesConfigurationSetAttributes) LastFreshStart() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("last_fresh_start"))
}

// Name returns a reference to field name of aws_ses_configuration_set.
func (scs sesConfigurationSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("name"))
}

// ReputationMetricsEnabled returns a reference to field reputation_metrics_enabled of aws_ses_configuration_set.
func (scs sesConfigurationSetAttributes) ReputationMetricsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(scs.ref.Append("reputation_metrics_enabled"))
}

// SendingEnabled returns a reference to field sending_enabled of aws_ses_configuration_set.
func (scs sesConfigurationSetAttributes) SendingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(scs.ref.Append("sending_enabled"))
}

func (scs sesConfigurationSetAttributes) DeliveryOptions() terra.ListValue[sesconfigurationset.DeliveryOptionsAttributes] {
	return terra.ReferenceAsList[sesconfigurationset.DeliveryOptionsAttributes](scs.ref.Append("delivery_options"))
}

func (scs sesConfigurationSetAttributes) TrackingOptions() terra.ListValue[sesconfigurationset.TrackingOptionsAttributes] {
	return terra.ReferenceAsList[sesconfigurationset.TrackingOptionsAttributes](scs.ref.Append("tracking_options"))
}

type sesConfigurationSetState struct {
	Arn                      string                                     `json:"arn"`
	Id                       string                                     `json:"id"`
	LastFreshStart           string                                     `json:"last_fresh_start"`
	Name                     string                                     `json:"name"`
	ReputationMetricsEnabled bool                                       `json:"reputation_metrics_enabled"`
	SendingEnabled           bool                                       `json:"sending_enabled"`
	DeliveryOptions          []sesconfigurationset.DeliveryOptionsState `json:"delivery_options"`
	TrackingOptions          []sesconfigurationset.TrackingOptionsState `json:"tracking_options"`
}
