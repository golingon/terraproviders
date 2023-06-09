// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	securityscannerscanconfig "github.com/golingon/terraproviders/googlebeta/4.63.1/securityscannerscanconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecurityScannerScanConfig creates a new instance of [SecurityScannerScanConfig].
func NewSecurityScannerScanConfig(name string, args SecurityScannerScanConfigArgs) *SecurityScannerScanConfig {
	return &SecurityScannerScanConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityScannerScanConfig)(nil)

// SecurityScannerScanConfig represents the Terraform resource google_security_scanner_scan_config.
type SecurityScannerScanConfig struct {
	Name      string
	Args      SecurityScannerScanConfigArgs
	state     *securityScannerScanConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityScannerScanConfig].
func (sssc *SecurityScannerScanConfig) Type() string {
	return "google_security_scanner_scan_config"
}

// LocalName returns the local name for [SecurityScannerScanConfig].
func (sssc *SecurityScannerScanConfig) LocalName() string {
	return sssc.Name
}

// Configuration returns the configuration (args) for [SecurityScannerScanConfig].
func (sssc *SecurityScannerScanConfig) Configuration() interface{} {
	return sssc.Args
}

// DependOn is used for other resources to depend on [SecurityScannerScanConfig].
func (sssc *SecurityScannerScanConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(sssc)
}

// Dependencies returns the list of resources [SecurityScannerScanConfig] depends_on.
func (sssc *SecurityScannerScanConfig) Dependencies() terra.Dependencies {
	return sssc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityScannerScanConfig].
func (sssc *SecurityScannerScanConfig) LifecycleManagement() *terra.Lifecycle {
	return sssc.Lifecycle
}

// Attributes returns the attributes for [SecurityScannerScanConfig].
func (sssc *SecurityScannerScanConfig) Attributes() securityScannerScanConfigAttributes {
	return securityScannerScanConfigAttributes{ref: terra.ReferenceResource(sssc)}
}

// ImportState imports the given attribute values into [SecurityScannerScanConfig]'s state.
func (sssc *SecurityScannerScanConfig) ImportState(av io.Reader) error {
	sssc.state = &securityScannerScanConfigState{}
	if err := json.NewDecoder(av).Decode(sssc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sssc.Type(), sssc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityScannerScanConfig] has state.
func (sssc *SecurityScannerScanConfig) State() (*securityScannerScanConfigState, bool) {
	return sssc.state, sssc.state != nil
}

// StateMust returns the state for [SecurityScannerScanConfig]. Panics if the state is nil.
func (sssc *SecurityScannerScanConfig) StateMust() *securityScannerScanConfigState {
	if sssc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sssc.Type(), sssc.LocalName()))
	}
	return sssc.state
}

// SecurityScannerScanConfigArgs contains the configurations for google_security_scanner_scan_config.
type SecurityScannerScanConfigArgs struct {
	// BlacklistPatterns: list of string, optional
	BlacklistPatterns terra.ListValue[terra.StringValue] `hcl:"blacklist_patterns,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// ExportToSecurityCommandCenter: string, optional
	ExportToSecurityCommandCenter terra.StringValue `hcl:"export_to_security_command_center,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxQps: number, optional
	MaxQps terra.NumberValue `hcl:"max_qps,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// StartingUrls: list of string, required
	StartingUrls terra.ListValue[terra.StringValue] `hcl:"starting_urls,attr" validate:"required"`
	// TargetPlatforms: list of string, optional
	TargetPlatforms terra.ListValue[terra.StringValue] `hcl:"target_platforms,attr"`
	// UserAgent: string, optional
	UserAgent terra.StringValue `hcl:"user_agent,attr"`
	// Authentication: optional
	Authentication *securityscannerscanconfig.Authentication `hcl:"authentication,block"`
	// Schedule: optional
	Schedule *securityscannerscanconfig.Schedule `hcl:"schedule,block"`
	// Timeouts: optional
	Timeouts *securityscannerscanconfig.Timeouts `hcl:"timeouts,block"`
}
type securityScannerScanConfigAttributes struct {
	ref terra.Reference
}

// BlacklistPatterns returns a reference to field blacklist_patterns of google_security_scanner_scan_config.
func (sssc securityScannerScanConfigAttributes) BlacklistPatterns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sssc.ref.Append("blacklist_patterns"))
}

// DisplayName returns a reference to field display_name of google_security_scanner_scan_config.
func (sssc securityScannerScanConfigAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sssc.ref.Append("display_name"))
}

// ExportToSecurityCommandCenter returns a reference to field export_to_security_command_center of google_security_scanner_scan_config.
func (sssc securityScannerScanConfigAttributes) ExportToSecurityCommandCenter() terra.StringValue {
	return terra.ReferenceAsString(sssc.ref.Append("export_to_security_command_center"))
}

// Id returns a reference to field id of google_security_scanner_scan_config.
func (sssc securityScannerScanConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sssc.ref.Append("id"))
}

// MaxQps returns a reference to field max_qps of google_security_scanner_scan_config.
func (sssc securityScannerScanConfigAttributes) MaxQps() terra.NumberValue {
	return terra.ReferenceAsNumber(sssc.ref.Append("max_qps"))
}

// Name returns a reference to field name of google_security_scanner_scan_config.
func (sssc securityScannerScanConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sssc.ref.Append("name"))
}

// Project returns a reference to field project of google_security_scanner_scan_config.
func (sssc securityScannerScanConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sssc.ref.Append("project"))
}

// StartingUrls returns a reference to field starting_urls of google_security_scanner_scan_config.
func (sssc securityScannerScanConfigAttributes) StartingUrls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sssc.ref.Append("starting_urls"))
}

// TargetPlatforms returns a reference to field target_platforms of google_security_scanner_scan_config.
func (sssc securityScannerScanConfigAttributes) TargetPlatforms() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sssc.ref.Append("target_platforms"))
}

// UserAgent returns a reference to field user_agent of google_security_scanner_scan_config.
func (sssc securityScannerScanConfigAttributes) UserAgent() terra.StringValue {
	return terra.ReferenceAsString(sssc.ref.Append("user_agent"))
}

func (sssc securityScannerScanConfigAttributes) Authentication() terra.ListValue[securityscannerscanconfig.AuthenticationAttributes] {
	return terra.ReferenceAsList[securityscannerscanconfig.AuthenticationAttributes](sssc.ref.Append("authentication"))
}

func (sssc securityScannerScanConfigAttributes) Schedule() terra.ListValue[securityscannerscanconfig.ScheduleAttributes] {
	return terra.ReferenceAsList[securityscannerscanconfig.ScheduleAttributes](sssc.ref.Append("schedule"))
}

func (sssc securityScannerScanConfigAttributes) Timeouts() securityscannerscanconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[securityscannerscanconfig.TimeoutsAttributes](sssc.ref.Append("timeouts"))
}

type securityScannerScanConfigState struct {
	BlacklistPatterns             []string                                        `json:"blacklist_patterns"`
	DisplayName                   string                                          `json:"display_name"`
	ExportToSecurityCommandCenter string                                          `json:"export_to_security_command_center"`
	Id                            string                                          `json:"id"`
	MaxQps                        float64                                         `json:"max_qps"`
	Name                          string                                          `json:"name"`
	Project                       string                                          `json:"project"`
	StartingUrls                  []string                                        `json:"starting_urls"`
	TargetPlatforms               []string                                        `json:"target_platforms"`
	UserAgent                     string                                          `json:"user_agent"`
	Authentication                []securityscannerscanconfig.AuthenticationState `json:"authentication"`
	Schedule                      []securityscannerscanconfig.ScheduleState       `json:"schedule"`
	Timeouts                      *securityscannerscanconfig.TimeoutsState        `json:"timeouts"`
}
