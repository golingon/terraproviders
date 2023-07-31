// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computefirewall "github.com/golingon/terraproviders/googlebeta/4.75.1/computefirewall"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeFirewall creates a new instance of [ComputeFirewall].
func NewComputeFirewall(name string, args ComputeFirewallArgs) *ComputeFirewall {
	return &ComputeFirewall{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeFirewall)(nil)

// ComputeFirewall represents the Terraform resource google_compute_firewall.
type ComputeFirewall struct {
	Name      string
	Args      ComputeFirewallArgs
	state     *computeFirewallState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeFirewall].
func (cf *ComputeFirewall) Type() string {
	return "google_compute_firewall"
}

// LocalName returns the local name for [ComputeFirewall].
func (cf *ComputeFirewall) LocalName() string {
	return cf.Name
}

// Configuration returns the configuration (args) for [ComputeFirewall].
func (cf *ComputeFirewall) Configuration() interface{} {
	return cf.Args
}

// DependOn is used for other resources to depend on [ComputeFirewall].
func (cf *ComputeFirewall) DependOn() terra.Reference {
	return terra.ReferenceResource(cf)
}

// Dependencies returns the list of resources [ComputeFirewall] depends_on.
func (cf *ComputeFirewall) Dependencies() terra.Dependencies {
	return cf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeFirewall].
func (cf *ComputeFirewall) LifecycleManagement() *terra.Lifecycle {
	return cf.Lifecycle
}

// Attributes returns the attributes for [ComputeFirewall].
func (cf *ComputeFirewall) Attributes() computeFirewallAttributes {
	return computeFirewallAttributes{ref: terra.ReferenceResource(cf)}
}

// ImportState imports the given attribute values into [ComputeFirewall]'s state.
func (cf *ComputeFirewall) ImportState(av io.Reader) error {
	cf.state = &computeFirewallState{}
	if err := json.NewDecoder(av).Decode(cf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cf.Type(), cf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeFirewall] has state.
func (cf *ComputeFirewall) State() (*computeFirewallState, bool) {
	return cf.state, cf.state != nil
}

// StateMust returns the state for [ComputeFirewall]. Panics if the state is nil.
func (cf *ComputeFirewall) StateMust() *computeFirewallState {
	if cf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cf.Type(), cf.LocalName()))
	}
	return cf.state
}

// ComputeFirewallArgs contains the configurations for google_compute_firewall.
type ComputeFirewallArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DestinationRanges: set of string, optional
	DestinationRanges terra.SetValue[terra.StringValue] `hcl:"destination_ranges,attr"`
	// Direction: string, optional
	Direction terra.StringValue `hcl:"direction,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// EnableLogging: bool, optional
	EnableLogging terra.BoolValue `hcl:"enable_logging,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SourceRanges: set of string, optional
	SourceRanges terra.SetValue[terra.StringValue] `hcl:"source_ranges,attr"`
	// SourceServiceAccounts: set of string, optional
	SourceServiceAccounts terra.SetValue[terra.StringValue] `hcl:"source_service_accounts,attr"`
	// SourceTags: set of string, optional
	SourceTags terra.SetValue[terra.StringValue] `hcl:"source_tags,attr"`
	// TargetServiceAccounts: set of string, optional
	TargetServiceAccounts terra.SetValue[terra.StringValue] `hcl:"target_service_accounts,attr"`
	// TargetTags: set of string, optional
	TargetTags terra.SetValue[terra.StringValue] `hcl:"target_tags,attr"`
	// Allow: min=0
	Allow []computefirewall.Allow `hcl:"allow,block" validate:"min=0"`
	// Deny: min=0
	Deny []computefirewall.Deny `hcl:"deny,block" validate:"min=0"`
	// LogConfig: optional
	LogConfig *computefirewall.LogConfig `hcl:"log_config,block"`
	// Timeouts: optional
	Timeouts *computefirewall.Timeouts `hcl:"timeouts,block"`
}
type computeFirewallAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_firewall.
func (cf computeFirewallAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_firewall.
func (cf computeFirewallAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("description"))
}

// DestinationRanges returns a reference to field destination_ranges of google_compute_firewall.
func (cf computeFirewallAttributes) DestinationRanges() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cf.ref.Append("destination_ranges"))
}

// Direction returns a reference to field direction of google_compute_firewall.
func (cf computeFirewallAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("direction"))
}

// Disabled returns a reference to field disabled of google_compute_firewall.
func (cf computeFirewallAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(cf.ref.Append("disabled"))
}

// EnableLogging returns a reference to field enable_logging of google_compute_firewall.
func (cf computeFirewallAttributes) EnableLogging() terra.BoolValue {
	return terra.ReferenceAsBool(cf.ref.Append("enable_logging"))
}

// Id returns a reference to field id of google_compute_firewall.
func (cf computeFirewallAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_firewall.
func (cf computeFirewallAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_firewall.
func (cf computeFirewallAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("network"))
}

// Priority returns a reference to field priority of google_compute_firewall.
func (cf computeFirewallAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(cf.ref.Append("priority"))
}

// Project returns a reference to field project of google_compute_firewall.
func (cf computeFirewallAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_firewall.
func (cf computeFirewallAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("self_link"))
}

// SourceRanges returns a reference to field source_ranges of google_compute_firewall.
func (cf computeFirewallAttributes) SourceRanges() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cf.ref.Append("source_ranges"))
}

// SourceServiceAccounts returns a reference to field source_service_accounts of google_compute_firewall.
func (cf computeFirewallAttributes) SourceServiceAccounts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cf.ref.Append("source_service_accounts"))
}

// SourceTags returns a reference to field source_tags of google_compute_firewall.
func (cf computeFirewallAttributes) SourceTags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cf.ref.Append("source_tags"))
}

// TargetServiceAccounts returns a reference to field target_service_accounts of google_compute_firewall.
func (cf computeFirewallAttributes) TargetServiceAccounts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cf.ref.Append("target_service_accounts"))
}

// TargetTags returns a reference to field target_tags of google_compute_firewall.
func (cf computeFirewallAttributes) TargetTags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cf.ref.Append("target_tags"))
}

func (cf computeFirewallAttributes) Allow() terra.SetValue[computefirewall.AllowAttributes] {
	return terra.ReferenceAsSet[computefirewall.AllowAttributes](cf.ref.Append("allow"))
}

func (cf computeFirewallAttributes) Deny() terra.SetValue[computefirewall.DenyAttributes] {
	return terra.ReferenceAsSet[computefirewall.DenyAttributes](cf.ref.Append("deny"))
}

func (cf computeFirewallAttributes) LogConfig() terra.ListValue[computefirewall.LogConfigAttributes] {
	return terra.ReferenceAsList[computefirewall.LogConfigAttributes](cf.ref.Append("log_config"))
}

func (cf computeFirewallAttributes) Timeouts() computefirewall.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computefirewall.TimeoutsAttributes](cf.ref.Append("timeouts"))
}

type computeFirewallState struct {
	CreationTimestamp     string                           `json:"creation_timestamp"`
	Description           string                           `json:"description"`
	DestinationRanges     []string                         `json:"destination_ranges"`
	Direction             string                           `json:"direction"`
	Disabled              bool                             `json:"disabled"`
	EnableLogging         bool                             `json:"enable_logging"`
	Id                    string                           `json:"id"`
	Name                  string                           `json:"name"`
	Network               string                           `json:"network"`
	Priority              float64                          `json:"priority"`
	Project               string                           `json:"project"`
	SelfLink              string                           `json:"self_link"`
	SourceRanges          []string                         `json:"source_ranges"`
	SourceServiceAccounts []string                         `json:"source_service_accounts"`
	SourceTags            []string                         `json:"source_tags"`
	TargetServiceAccounts []string                         `json:"target_service_accounts"`
	TargetTags            []string                         `json:"target_tags"`
	Allow                 []computefirewall.AllowState     `json:"allow"`
	Deny                  []computefirewall.DenyState      `json:"deny"`
	LogConfig             []computefirewall.LogConfigState `json:"log_config"`
	Timeouts              *computefirewall.TimeoutsState   `json:"timeouts"`
}
