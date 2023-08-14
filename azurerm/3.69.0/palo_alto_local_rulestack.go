// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	paloaltolocalrulestack "github.com/golingon/terraproviders/azurerm/3.69.0/paloaltolocalrulestack"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPaloAltoLocalRulestack creates a new instance of [PaloAltoLocalRulestack].
func NewPaloAltoLocalRulestack(name string, args PaloAltoLocalRulestackArgs) *PaloAltoLocalRulestack {
	return &PaloAltoLocalRulestack{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PaloAltoLocalRulestack)(nil)

// PaloAltoLocalRulestack represents the Terraform resource azurerm_palo_alto_local_rulestack.
type PaloAltoLocalRulestack struct {
	Name      string
	Args      PaloAltoLocalRulestackArgs
	state     *paloAltoLocalRulestackState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PaloAltoLocalRulestack].
func (palr *PaloAltoLocalRulestack) Type() string {
	return "azurerm_palo_alto_local_rulestack"
}

// LocalName returns the local name for [PaloAltoLocalRulestack].
func (palr *PaloAltoLocalRulestack) LocalName() string {
	return palr.Name
}

// Configuration returns the configuration (args) for [PaloAltoLocalRulestack].
func (palr *PaloAltoLocalRulestack) Configuration() interface{} {
	return palr.Args
}

// DependOn is used for other resources to depend on [PaloAltoLocalRulestack].
func (palr *PaloAltoLocalRulestack) DependOn() terra.Reference {
	return terra.ReferenceResource(palr)
}

// Dependencies returns the list of resources [PaloAltoLocalRulestack] depends_on.
func (palr *PaloAltoLocalRulestack) Dependencies() terra.Dependencies {
	return palr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PaloAltoLocalRulestack].
func (palr *PaloAltoLocalRulestack) LifecycleManagement() *terra.Lifecycle {
	return palr.Lifecycle
}

// Attributes returns the attributes for [PaloAltoLocalRulestack].
func (palr *PaloAltoLocalRulestack) Attributes() paloAltoLocalRulestackAttributes {
	return paloAltoLocalRulestackAttributes{ref: terra.ReferenceResource(palr)}
}

// ImportState imports the given attribute values into [PaloAltoLocalRulestack]'s state.
func (palr *PaloAltoLocalRulestack) ImportState(av io.Reader) error {
	palr.state = &paloAltoLocalRulestackState{}
	if err := json.NewDecoder(av).Decode(palr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", palr.Type(), palr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PaloAltoLocalRulestack] has state.
func (palr *PaloAltoLocalRulestack) State() (*paloAltoLocalRulestackState, bool) {
	return palr.state, palr.state != nil
}

// StateMust returns the state for [PaloAltoLocalRulestack]. Panics if the state is nil.
func (palr *PaloAltoLocalRulestack) StateMust() *paloAltoLocalRulestackState {
	if palr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", palr.Type(), palr.LocalName()))
	}
	return palr.state
}

// PaloAltoLocalRulestackArgs contains the configurations for azurerm_palo_alto_local_rulestack.
type PaloAltoLocalRulestackArgs struct {
	// AntiSpywareProfile: string, optional
	AntiSpywareProfile terra.StringValue `hcl:"anti_spyware_profile,attr"`
	// AntiVirusProfile: string, optional
	AntiVirusProfile terra.StringValue `hcl:"anti_virus_profile,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DnsSubscription: string, optional
	DnsSubscription terra.StringValue `hcl:"dns_subscription,attr"`
	// FileBlockingProfile: string, optional
	FileBlockingProfile terra.StringValue `hcl:"file_blocking_profile,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// UrlFilteringProfile: string, optional
	UrlFilteringProfile terra.StringValue `hcl:"url_filtering_profile,attr"`
	// VulnerabilityProfile: string, optional
	VulnerabilityProfile terra.StringValue `hcl:"vulnerability_profile,attr"`
	// Timeouts: optional
	Timeouts *paloaltolocalrulestack.Timeouts `hcl:"timeouts,block"`
}
type paloAltoLocalRulestackAttributes struct {
	ref terra.Reference
}

// AntiSpywareProfile returns a reference to field anti_spyware_profile of azurerm_palo_alto_local_rulestack.
func (palr paloAltoLocalRulestackAttributes) AntiSpywareProfile() terra.StringValue {
	return terra.ReferenceAsString(palr.ref.Append("anti_spyware_profile"))
}

// AntiVirusProfile returns a reference to field anti_virus_profile of azurerm_palo_alto_local_rulestack.
func (palr paloAltoLocalRulestackAttributes) AntiVirusProfile() terra.StringValue {
	return terra.ReferenceAsString(palr.ref.Append("anti_virus_profile"))
}

// Description returns a reference to field description of azurerm_palo_alto_local_rulestack.
func (palr paloAltoLocalRulestackAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(palr.ref.Append("description"))
}

// DnsSubscription returns a reference to field dns_subscription of azurerm_palo_alto_local_rulestack.
func (palr paloAltoLocalRulestackAttributes) DnsSubscription() terra.StringValue {
	return terra.ReferenceAsString(palr.ref.Append("dns_subscription"))
}

// FileBlockingProfile returns a reference to field file_blocking_profile of azurerm_palo_alto_local_rulestack.
func (palr paloAltoLocalRulestackAttributes) FileBlockingProfile() terra.StringValue {
	return terra.ReferenceAsString(palr.ref.Append("file_blocking_profile"))
}

// Id returns a reference to field id of azurerm_palo_alto_local_rulestack.
func (palr paloAltoLocalRulestackAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(palr.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_palo_alto_local_rulestack.
func (palr paloAltoLocalRulestackAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(palr.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_palo_alto_local_rulestack.
func (palr paloAltoLocalRulestackAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(palr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_palo_alto_local_rulestack.
func (palr paloAltoLocalRulestackAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(palr.ref.Append("resource_group_name"))
}

// UrlFilteringProfile returns a reference to field url_filtering_profile of azurerm_palo_alto_local_rulestack.
func (palr paloAltoLocalRulestackAttributes) UrlFilteringProfile() terra.StringValue {
	return terra.ReferenceAsString(palr.ref.Append("url_filtering_profile"))
}

// VulnerabilityProfile returns a reference to field vulnerability_profile of azurerm_palo_alto_local_rulestack.
func (palr paloAltoLocalRulestackAttributes) VulnerabilityProfile() terra.StringValue {
	return terra.ReferenceAsString(palr.ref.Append("vulnerability_profile"))
}

func (palr paloAltoLocalRulestackAttributes) Timeouts() paloaltolocalrulestack.TimeoutsAttributes {
	return terra.ReferenceAsSingle[paloaltolocalrulestack.TimeoutsAttributes](palr.ref.Append("timeouts"))
}

type paloAltoLocalRulestackState struct {
	AntiSpywareProfile   string                                `json:"anti_spyware_profile"`
	AntiVirusProfile     string                                `json:"anti_virus_profile"`
	Description          string                                `json:"description"`
	DnsSubscription      string                                `json:"dns_subscription"`
	FileBlockingProfile  string                                `json:"file_blocking_profile"`
	Id                   string                                `json:"id"`
	Location             string                                `json:"location"`
	Name                 string                                `json:"name"`
	ResourceGroupName    string                                `json:"resource_group_name"`
	UrlFilteringProfile  string                                `json:"url_filtering_profile"`
	VulnerabilityProfile string                                `json:"vulnerability_profile"`
	Timeouts             *paloaltolocalrulestack.TimeoutsState `json:"timeouts"`
}
