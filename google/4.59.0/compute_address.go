// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeaddress "github.com/golingon/terraproviders/google/4.59.0/computeaddress"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeAddress(name string, args ComputeAddressArgs) *ComputeAddress {
	return &ComputeAddress{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeAddress)(nil)

type ComputeAddress struct {
	Name  string
	Args  ComputeAddressArgs
	state *computeAddressState
}

func (ca *ComputeAddress) Type() string {
	return "google_compute_address"
}

func (ca *ComputeAddress) LocalName() string {
	return ca.Name
}

func (ca *ComputeAddress) Configuration() interface{} {
	return ca.Args
}

func (ca *ComputeAddress) Attributes() computeAddressAttributes {
	return computeAddressAttributes{ref: terra.ReferenceResource(ca)}
}

func (ca *ComputeAddress) ImportState(av io.Reader) error {
	ca.state = &computeAddressState{}
	if err := json.NewDecoder(av).Decode(ca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ca.Type(), ca.LocalName(), err)
	}
	return nil
}

func (ca *ComputeAddress) State() (*computeAddressState, bool) {
	return ca.state, ca.state != nil
}

func (ca *ComputeAddress) StateMust() *computeAddressState {
	if ca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ca.Type(), ca.LocalName()))
	}
	return ca.state
}

func (ca *ComputeAddress) DependOn() terra.Reference {
	return terra.ReferenceResource(ca)
}

type ComputeAddressArgs struct {
	// Address: string, optional
	Address terra.StringValue `hcl:"address,attr"`
	// AddressType: string, optional
	AddressType terra.StringValue `hcl:"address_type,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// NetworkTier: string, optional
	NetworkTier terra.StringValue `hcl:"network_tier,attr"`
	// PrefixLength: number, optional
	PrefixLength terra.NumberValue `hcl:"prefix_length,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Purpose: string, optional
	Purpose terra.StringValue `hcl:"purpose,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Subnetwork: string, optional
	Subnetwork terra.StringValue `hcl:"subnetwork,attr"`
	// Timeouts: optional
	Timeouts *computeaddress.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ComputeAddress depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeAddressAttributes struct {
	ref terra.Reference
}

func (ca computeAddressAttributes) Address() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("address"))
}

func (ca computeAddressAttributes) AddressType() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("address_type"))
}

func (ca computeAddressAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("creation_timestamp"))
}

func (ca computeAddressAttributes) Description() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("description"))
}

func (ca computeAddressAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("id"))
}

func (ca computeAddressAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("name"))
}

func (ca computeAddressAttributes) Network() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("network"))
}

func (ca computeAddressAttributes) NetworkTier() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("network_tier"))
}

func (ca computeAddressAttributes) PrefixLength() terra.NumberValue {
	return terra.ReferenceNumber(ca.ref.Append("prefix_length"))
}

func (ca computeAddressAttributes) Project() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("project"))
}

func (ca computeAddressAttributes) Purpose() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("purpose"))
}

func (ca computeAddressAttributes) Region() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("region"))
}

func (ca computeAddressAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("self_link"))
}

func (ca computeAddressAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("subnetwork"))
}

func (ca computeAddressAttributes) Users() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ca.ref.Append("users"))
}

func (ca computeAddressAttributes) Timeouts() computeaddress.TimeoutsAttributes {
	return terra.ReferenceSingle[computeaddress.TimeoutsAttributes](ca.ref.Append("timeouts"))
}

type computeAddressState struct {
	Address           string                        `json:"address"`
	AddressType       string                        `json:"address_type"`
	CreationTimestamp string                        `json:"creation_timestamp"`
	Description       string                        `json:"description"`
	Id                string                        `json:"id"`
	Name              string                        `json:"name"`
	Network           string                        `json:"network"`
	NetworkTier       string                        `json:"network_tier"`
	PrefixLength      float64                       `json:"prefix_length"`
	Project           string                        `json:"project"`
	Purpose           string                        `json:"purpose"`
	Region            string                        `json:"region"`
	SelfLink          string                        `json:"self_link"`
	Subnetwork        string                        `json:"subnetwork"`
	Users             []string                      `json:"users"`
	Timeouts          *computeaddress.TimeoutsState `json:"timeouts"`
}
