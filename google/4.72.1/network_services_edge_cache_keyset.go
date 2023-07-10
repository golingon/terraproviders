// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	networkservicesedgecachekeyset "github.com/golingon/terraproviders/google/4.72.1/networkservicesedgecachekeyset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkServicesEdgeCacheKeyset creates a new instance of [NetworkServicesEdgeCacheKeyset].
func NewNetworkServicesEdgeCacheKeyset(name string, args NetworkServicesEdgeCacheKeysetArgs) *NetworkServicesEdgeCacheKeyset {
	return &NetworkServicesEdgeCacheKeyset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkServicesEdgeCacheKeyset)(nil)

// NetworkServicesEdgeCacheKeyset represents the Terraform resource google_network_services_edge_cache_keyset.
type NetworkServicesEdgeCacheKeyset struct {
	Name      string
	Args      NetworkServicesEdgeCacheKeysetArgs
	state     *networkServicesEdgeCacheKeysetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkServicesEdgeCacheKeyset].
func (nseck *NetworkServicesEdgeCacheKeyset) Type() string {
	return "google_network_services_edge_cache_keyset"
}

// LocalName returns the local name for [NetworkServicesEdgeCacheKeyset].
func (nseck *NetworkServicesEdgeCacheKeyset) LocalName() string {
	return nseck.Name
}

// Configuration returns the configuration (args) for [NetworkServicesEdgeCacheKeyset].
func (nseck *NetworkServicesEdgeCacheKeyset) Configuration() interface{} {
	return nseck.Args
}

// DependOn is used for other resources to depend on [NetworkServicesEdgeCacheKeyset].
func (nseck *NetworkServicesEdgeCacheKeyset) DependOn() terra.Reference {
	return terra.ReferenceResource(nseck)
}

// Dependencies returns the list of resources [NetworkServicesEdgeCacheKeyset] depends_on.
func (nseck *NetworkServicesEdgeCacheKeyset) Dependencies() terra.Dependencies {
	return nseck.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkServicesEdgeCacheKeyset].
func (nseck *NetworkServicesEdgeCacheKeyset) LifecycleManagement() *terra.Lifecycle {
	return nseck.Lifecycle
}

// Attributes returns the attributes for [NetworkServicesEdgeCacheKeyset].
func (nseck *NetworkServicesEdgeCacheKeyset) Attributes() networkServicesEdgeCacheKeysetAttributes {
	return networkServicesEdgeCacheKeysetAttributes{ref: terra.ReferenceResource(nseck)}
}

// ImportState imports the given attribute values into [NetworkServicesEdgeCacheKeyset]'s state.
func (nseck *NetworkServicesEdgeCacheKeyset) ImportState(av io.Reader) error {
	nseck.state = &networkServicesEdgeCacheKeysetState{}
	if err := json.NewDecoder(av).Decode(nseck.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nseck.Type(), nseck.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkServicesEdgeCacheKeyset] has state.
func (nseck *NetworkServicesEdgeCacheKeyset) State() (*networkServicesEdgeCacheKeysetState, bool) {
	return nseck.state, nseck.state != nil
}

// StateMust returns the state for [NetworkServicesEdgeCacheKeyset]. Panics if the state is nil.
func (nseck *NetworkServicesEdgeCacheKeyset) StateMust() *networkServicesEdgeCacheKeysetState {
	if nseck.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nseck.Type(), nseck.LocalName()))
	}
	return nseck.state
}

// NetworkServicesEdgeCacheKeysetArgs contains the configurations for google_network_services_edge_cache_keyset.
type NetworkServicesEdgeCacheKeysetArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// PublicKey: min=0,max=3
	PublicKey []networkservicesedgecachekeyset.PublicKey `hcl:"public_key,block" validate:"min=0,max=3"`
	// Timeouts: optional
	Timeouts *networkservicesedgecachekeyset.Timeouts `hcl:"timeouts,block"`
	// ValidationSharedKeys: min=0,max=3
	ValidationSharedKeys []networkservicesedgecachekeyset.ValidationSharedKeys `hcl:"validation_shared_keys,block" validate:"min=0,max=3"`
}
type networkServicesEdgeCacheKeysetAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_network_services_edge_cache_keyset.
func (nseck networkServicesEdgeCacheKeysetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nseck.ref.Append("description"))
}

// Id returns a reference to field id of google_network_services_edge_cache_keyset.
func (nseck networkServicesEdgeCacheKeysetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nseck.ref.Append("id"))
}

// Labels returns a reference to field labels of google_network_services_edge_cache_keyset.
func (nseck networkServicesEdgeCacheKeysetAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nseck.ref.Append("labels"))
}

// Name returns a reference to field name of google_network_services_edge_cache_keyset.
func (nseck networkServicesEdgeCacheKeysetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nseck.ref.Append("name"))
}

// Project returns a reference to field project of google_network_services_edge_cache_keyset.
func (nseck networkServicesEdgeCacheKeysetAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nseck.ref.Append("project"))
}

func (nseck networkServicesEdgeCacheKeysetAttributes) PublicKey() terra.ListValue[networkservicesedgecachekeyset.PublicKeyAttributes] {
	return terra.ReferenceAsList[networkservicesedgecachekeyset.PublicKeyAttributes](nseck.ref.Append("public_key"))
}

func (nseck networkServicesEdgeCacheKeysetAttributes) Timeouts() networkservicesedgecachekeyset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkservicesedgecachekeyset.TimeoutsAttributes](nseck.ref.Append("timeouts"))
}

func (nseck networkServicesEdgeCacheKeysetAttributes) ValidationSharedKeys() terra.ListValue[networkservicesedgecachekeyset.ValidationSharedKeysAttributes] {
	return terra.ReferenceAsList[networkservicesedgecachekeyset.ValidationSharedKeysAttributes](nseck.ref.Append("validation_shared_keys"))
}

type networkServicesEdgeCacheKeysetState struct {
	Description          string                                                     `json:"description"`
	Id                   string                                                     `json:"id"`
	Labels               map[string]string                                          `json:"labels"`
	Name                 string                                                     `json:"name"`
	Project              string                                                     `json:"project"`
	PublicKey            []networkservicesedgecachekeyset.PublicKeyState            `json:"public_key"`
	Timeouts             *networkservicesedgecachekeyset.TimeoutsState              `json:"timeouts"`
	ValidationSharedKeys []networkservicesedgecachekeyset.ValidationSharedKeysState `json:"validation_shared_keys"`
}
