// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	networksecurityurllists "github.com/golingon/terraproviders/googlebeta/4.76.0/networksecurityurllists"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkSecurityUrlLists creates a new instance of [NetworkSecurityUrlLists].
func NewNetworkSecurityUrlLists(name string, args NetworkSecurityUrlListsArgs) *NetworkSecurityUrlLists {
	return &NetworkSecurityUrlLists{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkSecurityUrlLists)(nil)

// NetworkSecurityUrlLists represents the Terraform resource google_network_security_url_lists.
type NetworkSecurityUrlLists struct {
	Name      string
	Args      NetworkSecurityUrlListsArgs
	state     *networkSecurityUrlListsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkSecurityUrlLists].
func (nsul *NetworkSecurityUrlLists) Type() string {
	return "google_network_security_url_lists"
}

// LocalName returns the local name for [NetworkSecurityUrlLists].
func (nsul *NetworkSecurityUrlLists) LocalName() string {
	return nsul.Name
}

// Configuration returns the configuration (args) for [NetworkSecurityUrlLists].
func (nsul *NetworkSecurityUrlLists) Configuration() interface{} {
	return nsul.Args
}

// DependOn is used for other resources to depend on [NetworkSecurityUrlLists].
func (nsul *NetworkSecurityUrlLists) DependOn() terra.Reference {
	return terra.ReferenceResource(nsul)
}

// Dependencies returns the list of resources [NetworkSecurityUrlLists] depends_on.
func (nsul *NetworkSecurityUrlLists) Dependencies() terra.Dependencies {
	return nsul.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkSecurityUrlLists].
func (nsul *NetworkSecurityUrlLists) LifecycleManagement() *terra.Lifecycle {
	return nsul.Lifecycle
}

// Attributes returns the attributes for [NetworkSecurityUrlLists].
func (nsul *NetworkSecurityUrlLists) Attributes() networkSecurityUrlListsAttributes {
	return networkSecurityUrlListsAttributes{ref: terra.ReferenceResource(nsul)}
}

// ImportState imports the given attribute values into [NetworkSecurityUrlLists]'s state.
func (nsul *NetworkSecurityUrlLists) ImportState(av io.Reader) error {
	nsul.state = &networkSecurityUrlListsState{}
	if err := json.NewDecoder(av).Decode(nsul.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nsul.Type(), nsul.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkSecurityUrlLists] has state.
func (nsul *NetworkSecurityUrlLists) State() (*networkSecurityUrlListsState, bool) {
	return nsul.state, nsul.state != nil
}

// StateMust returns the state for [NetworkSecurityUrlLists]. Panics if the state is nil.
func (nsul *NetworkSecurityUrlLists) StateMust() *networkSecurityUrlListsState {
	if nsul.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nsul.Type(), nsul.LocalName()))
	}
	return nsul.state
}

// NetworkSecurityUrlListsArgs contains the configurations for google_network_security_url_lists.
type NetworkSecurityUrlListsArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networksecurityurllists.Timeouts `hcl:"timeouts,block"`
}
type networkSecurityUrlListsAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_network_security_url_lists.
func (nsul networkSecurityUrlListsAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nsul.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_security_url_lists.
func (nsul networkSecurityUrlListsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nsul.ref.Append("description"))
}

// Id returns a reference to field id of google_network_security_url_lists.
func (nsul networkSecurityUrlListsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nsul.ref.Append("id"))
}

// Location returns a reference to field location of google_network_security_url_lists.
func (nsul networkSecurityUrlListsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nsul.ref.Append("location"))
}

// Name returns a reference to field name of google_network_security_url_lists.
func (nsul networkSecurityUrlListsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nsul.ref.Append("name"))
}

// Project returns a reference to field project of google_network_security_url_lists.
func (nsul networkSecurityUrlListsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nsul.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_network_security_url_lists.
func (nsul networkSecurityUrlListsAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nsul.ref.Append("update_time"))
}

// Values returns a reference to field values of google_network_security_url_lists.
func (nsul networkSecurityUrlListsAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nsul.ref.Append("values"))
}

func (nsul networkSecurityUrlListsAttributes) Timeouts() networksecurityurllists.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networksecurityurllists.TimeoutsAttributes](nsul.ref.Append("timeouts"))
}

type networkSecurityUrlListsState struct {
	CreateTime  string                                 `json:"create_time"`
	Description string                                 `json:"description"`
	Id          string                                 `json:"id"`
	Location    string                                 `json:"location"`
	Name        string                                 `json:"name"`
	Project     string                                 `json:"project"`
	UpdateTime  string                                 `json:"update_time"`
	Values      []string                               `json:"values"`
	Timeouts    *networksecurityurllists.TimeoutsState `json:"timeouts"`
}
