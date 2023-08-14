// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apigeenataddress "github.com/golingon/terraproviders/googlebeta/4.77.0/apigeenataddress"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeNatAddress creates a new instance of [ApigeeNatAddress].
func NewApigeeNatAddress(name string, args ApigeeNatAddressArgs) *ApigeeNatAddress {
	return &ApigeeNatAddress{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeNatAddress)(nil)

// ApigeeNatAddress represents the Terraform resource google_apigee_nat_address.
type ApigeeNatAddress struct {
	Name      string
	Args      ApigeeNatAddressArgs
	state     *apigeeNatAddressState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeNatAddress].
func (ana *ApigeeNatAddress) Type() string {
	return "google_apigee_nat_address"
}

// LocalName returns the local name for [ApigeeNatAddress].
func (ana *ApigeeNatAddress) LocalName() string {
	return ana.Name
}

// Configuration returns the configuration (args) for [ApigeeNatAddress].
func (ana *ApigeeNatAddress) Configuration() interface{} {
	return ana.Args
}

// DependOn is used for other resources to depend on [ApigeeNatAddress].
func (ana *ApigeeNatAddress) DependOn() terra.Reference {
	return terra.ReferenceResource(ana)
}

// Dependencies returns the list of resources [ApigeeNatAddress] depends_on.
func (ana *ApigeeNatAddress) Dependencies() terra.Dependencies {
	return ana.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeNatAddress].
func (ana *ApigeeNatAddress) LifecycleManagement() *terra.Lifecycle {
	return ana.Lifecycle
}

// Attributes returns the attributes for [ApigeeNatAddress].
func (ana *ApigeeNatAddress) Attributes() apigeeNatAddressAttributes {
	return apigeeNatAddressAttributes{ref: terra.ReferenceResource(ana)}
}

// ImportState imports the given attribute values into [ApigeeNatAddress]'s state.
func (ana *ApigeeNatAddress) ImportState(av io.Reader) error {
	ana.state = &apigeeNatAddressState{}
	if err := json.NewDecoder(av).Decode(ana.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ana.Type(), ana.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeNatAddress] has state.
func (ana *ApigeeNatAddress) State() (*apigeeNatAddressState, bool) {
	return ana.state, ana.state != nil
}

// StateMust returns the state for [ApigeeNatAddress]. Panics if the state is nil.
func (ana *ApigeeNatAddress) StateMust() *apigeeNatAddressState {
	if ana.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ana.Type(), ana.LocalName()))
	}
	return ana.state
}

// ApigeeNatAddressArgs contains the configurations for google_apigee_nat_address.
type ApigeeNatAddressArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apigeenataddress.Timeouts `hcl:"timeouts,block"`
}
type apigeeNatAddressAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_apigee_nat_address.
func (ana apigeeNatAddressAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ana.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of google_apigee_nat_address.
func (ana apigeeNatAddressAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(ana.ref.Append("instance_id"))
}

// IpAddress returns a reference to field ip_address of google_apigee_nat_address.
func (ana apigeeNatAddressAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ana.ref.Append("ip_address"))
}

// Name returns a reference to field name of google_apigee_nat_address.
func (ana apigeeNatAddressAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ana.ref.Append("name"))
}

// State returns a reference to field state of google_apigee_nat_address.
func (ana apigeeNatAddressAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ana.ref.Append("state"))
}

func (ana apigeeNatAddressAttributes) Timeouts() apigeenataddress.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeenataddress.TimeoutsAttributes](ana.ref.Append("timeouts"))
}

type apigeeNatAddressState struct {
	Id         string                          `json:"id"`
	InstanceId string                          `json:"instance_id"`
	IpAddress  string                          `json:"ip_address"`
	Name       string                          `json:"name"`
	State      string                          `json:"state"`
	Timeouts   *apigeenataddress.TimeoutsState `json:"timeouts"`
}