// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	networkmanagementconnectivitytest "github.com/golingon/terraproviders/google/4.59.0/networkmanagementconnectivitytest"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkManagementConnectivityTest creates a new instance of [NetworkManagementConnectivityTest].
func NewNetworkManagementConnectivityTest(name string, args NetworkManagementConnectivityTestArgs) *NetworkManagementConnectivityTest {
	return &NetworkManagementConnectivityTest{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManagementConnectivityTest)(nil)

// NetworkManagementConnectivityTest represents the Terraform resource google_network_management_connectivity_test.
type NetworkManagementConnectivityTest struct {
	Name      string
	Args      NetworkManagementConnectivityTestArgs
	state     *networkManagementConnectivityTestState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkManagementConnectivityTest].
func (nmct *NetworkManagementConnectivityTest) Type() string {
	return "google_network_management_connectivity_test"
}

// LocalName returns the local name for [NetworkManagementConnectivityTest].
func (nmct *NetworkManagementConnectivityTest) LocalName() string {
	return nmct.Name
}

// Configuration returns the configuration (args) for [NetworkManagementConnectivityTest].
func (nmct *NetworkManagementConnectivityTest) Configuration() interface{} {
	return nmct.Args
}

// DependOn is used for other resources to depend on [NetworkManagementConnectivityTest].
func (nmct *NetworkManagementConnectivityTest) DependOn() terra.Reference {
	return terra.ReferenceResource(nmct)
}

// Dependencies returns the list of resources [NetworkManagementConnectivityTest] depends_on.
func (nmct *NetworkManagementConnectivityTest) Dependencies() terra.Dependencies {
	return nmct.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkManagementConnectivityTest].
func (nmct *NetworkManagementConnectivityTest) LifecycleManagement() *terra.Lifecycle {
	return nmct.Lifecycle
}

// Attributes returns the attributes for [NetworkManagementConnectivityTest].
func (nmct *NetworkManagementConnectivityTest) Attributes() networkManagementConnectivityTestAttributes {
	return networkManagementConnectivityTestAttributes{ref: terra.ReferenceResource(nmct)}
}

// ImportState imports the given attribute values into [NetworkManagementConnectivityTest]'s state.
func (nmct *NetworkManagementConnectivityTest) ImportState(av io.Reader) error {
	nmct.state = &networkManagementConnectivityTestState{}
	if err := json.NewDecoder(av).Decode(nmct.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nmct.Type(), nmct.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkManagementConnectivityTest] has state.
func (nmct *NetworkManagementConnectivityTest) State() (*networkManagementConnectivityTestState, bool) {
	return nmct.state, nmct.state != nil
}

// StateMust returns the state for [NetworkManagementConnectivityTest]. Panics if the state is nil.
func (nmct *NetworkManagementConnectivityTest) StateMust() *networkManagementConnectivityTestState {
	if nmct.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nmct.Type(), nmct.LocalName()))
	}
	return nmct.state
}

// NetworkManagementConnectivityTestArgs contains the configurations for google_network_management_connectivity_test.
type NetworkManagementConnectivityTestArgs struct {
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
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// RelatedProjects: list of string, optional
	RelatedProjects terra.ListValue[terra.StringValue] `hcl:"related_projects,attr"`
	// Destination: required
	Destination *networkmanagementconnectivitytest.Destination `hcl:"destination,block" validate:"required"`
	// Source: required
	Source *networkmanagementconnectivitytest.Source `hcl:"source,block" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagementconnectivitytest.Timeouts `hcl:"timeouts,block"`
}
type networkManagementConnectivityTestAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_network_management_connectivity_test.
func (nmct networkManagementConnectivityTestAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nmct.ref.Append("description"))
}

// Id returns a reference to field id of google_network_management_connectivity_test.
func (nmct networkManagementConnectivityTestAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nmct.ref.Append("id"))
}

// Labels returns a reference to field labels of google_network_management_connectivity_test.
func (nmct networkManagementConnectivityTestAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nmct.ref.Append("labels"))
}

// Name returns a reference to field name of google_network_management_connectivity_test.
func (nmct networkManagementConnectivityTestAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nmct.ref.Append("name"))
}

// Project returns a reference to field project of google_network_management_connectivity_test.
func (nmct networkManagementConnectivityTestAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nmct.ref.Append("project"))
}

// Protocol returns a reference to field protocol of google_network_management_connectivity_test.
func (nmct networkManagementConnectivityTestAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(nmct.ref.Append("protocol"))
}

// RelatedProjects returns a reference to field related_projects of google_network_management_connectivity_test.
func (nmct networkManagementConnectivityTestAttributes) RelatedProjects() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nmct.ref.Append("related_projects"))
}

func (nmct networkManagementConnectivityTestAttributes) Destination() terra.ListValue[networkmanagementconnectivitytest.DestinationAttributes] {
	return terra.ReferenceAsList[networkmanagementconnectivitytest.DestinationAttributes](nmct.ref.Append("destination"))
}

func (nmct networkManagementConnectivityTestAttributes) Source() terra.ListValue[networkmanagementconnectivitytest.SourceAttributes] {
	return terra.ReferenceAsList[networkmanagementconnectivitytest.SourceAttributes](nmct.ref.Append("source"))
}

func (nmct networkManagementConnectivityTestAttributes) Timeouts() networkmanagementconnectivitytest.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagementconnectivitytest.TimeoutsAttributes](nmct.ref.Append("timeouts"))
}

type networkManagementConnectivityTestState struct {
	Description     string                                               `json:"description"`
	Id              string                                               `json:"id"`
	Labels          map[string]string                                    `json:"labels"`
	Name            string                                               `json:"name"`
	Project         string                                               `json:"project"`
	Protocol        string                                               `json:"protocol"`
	RelatedProjects []string                                             `json:"related_projects"`
	Destination     []networkmanagementconnectivitytest.DestinationState `json:"destination"`
	Source          []networkmanagementconnectivitytest.SourceState      `json:"source"`
	Timeouts        *networkmanagementconnectivitytest.TimeoutsState     `json:"timeouts"`
}
