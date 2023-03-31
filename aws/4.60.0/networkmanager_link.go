// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagerlink "github.com/golingon/terraproviders/aws/4.60.0/networkmanagerlink"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerLink creates a new instance of [NetworkmanagerLink].
func NewNetworkmanagerLink(name string, args NetworkmanagerLinkArgs) *NetworkmanagerLink {
	return &NetworkmanagerLink{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerLink)(nil)

// NetworkmanagerLink represents the Terraform resource aws_networkmanager_link.
type NetworkmanagerLink struct {
	Name      string
	Args      NetworkmanagerLinkArgs
	state     *networkmanagerLinkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerLink].
func (nl *NetworkmanagerLink) Type() string {
	return "aws_networkmanager_link"
}

// LocalName returns the local name for [NetworkmanagerLink].
func (nl *NetworkmanagerLink) LocalName() string {
	return nl.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerLink].
func (nl *NetworkmanagerLink) Configuration() interface{} {
	return nl.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerLink].
func (nl *NetworkmanagerLink) DependOn() terra.Reference {
	return terra.ReferenceResource(nl)
}

// Dependencies returns the list of resources [NetworkmanagerLink] depends_on.
func (nl *NetworkmanagerLink) Dependencies() terra.Dependencies {
	return nl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerLink].
func (nl *NetworkmanagerLink) LifecycleManagement() *terra.Lifecycle {
	return nl.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerLink].
func (nl *NetworkmanagerLink) Attributes() networkmanagerLinkAttributes {
	return networkmanagerLinkAttributes{ref: terra.ReferenceResource(nl)}
}

// ImportState imports the given attribute values into [NetworkmanagerLink]'s state.
func (nl *NetworkmanagerLink) ImportState(av io.Reader) error {
	nl.state = &networkmanagerLinkState{}
	if err := json.NewDecoder(av).Decode(nl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nl.Type(), nl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerLink] has state.
func (nl *NetworkmanagerLink) State() (*networkmanagerLinkState, bool) {
	return nl.state, nl.state != nil
}

// StateMust returns the state for [NetworkmanagerLink]. Panics if the state is nil.
func (nl *NetworkmanagerLink) StateMust() *networkmanagerLinkState {
	if nl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nl.Type(), nl.LocalName()))
	}
	return nl.state
}

// NetworkmanagerLinkArgs contains the configurations for aws_networkmanager_link.
type NetworkmanagerLinkArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProviderName: string, optional
	ProviderName terra.StringValue `hcl:"provider_name,attr"`
	// SiteId: string, required
	SiteId terra.StringValue `hcl:"site_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Bandwidth: required
	Bandwidth *networkmanagerlink.Bandwidth `hcl:"bandwidth,block" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagerlink.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerLinkAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_link.
func (nl networkmanagerLinkAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("arn"))
}

// Description returns a reference to field description of aws_networkmanager_link.
func (nl networkmanagerLinkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("description"))
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_link.
func (nl networkmanagerLinkAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_link.
func (nl networkmanagerLinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("id"))
}

// ProviderName returns a reference to field provider_name of aws_networkmanager_link.
func (nl networkmanagerLinkAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("provider_name"))
}

// SiteId returns a reference to field site_id of aws_networkmanager_link.
func (nl networkmanagerLinkAttributes) SiteId() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("site_id"))
}

// Tags returns a reference to field tags of aws_networkmanager_link.
func (nl networkmanagerLinkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nl.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_networkmanager_link.
func (nl networkmanagerLinkAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nl.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_networkmanager_link.
func (nl networkmanagerLinkAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("type"))
}

func (nl networkmanagerLinkAttributes) Bandwidth() terra.ListValue[networkmanagerlink.BandwidthAttributes] {
	return terra.ReferenceAsList[networkmanagerlink.BandwidthAttributes](nl.ref.Append("bandwidth"))
}

func (nl networkmanagerLinkAttributes) Timeouts() networkmanagerlink.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagerlink.TimeoutsAttributes](nl.ref.Append("timeouts"))
}

type networkmanagerLinkState struct {
	Arn             string                              `json:"arn"`
	Description     string                              `json:"description"`
	GlobalNetworkId string                              `json:"global_network_id"`
	Id              string                              `json:"id"`
	ProviderName    string                              `json:"provider_name"`
	SiteId          string                              `json:"site_id"`
	Tags            map[string]string                   `json:"tags"`
	TagsAll         map[string]string                   `json:"tags_all"`
	Type            string                              `json:"type"`
	Bandwidth       []networkmanagerlink.BandwidthState `json:"bandwidth"`
	Timeouts        *networkmanagerlink.TimeoutsState   `json:"timeouts"`
}
