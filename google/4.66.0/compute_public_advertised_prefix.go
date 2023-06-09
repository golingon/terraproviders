// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computepublicadvertisedprefix "github.com/golingon/terraproviders/google/4.66.0/computepublicadvertisedprefix"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputePublicAdvertisedPrefix creates a new instance of [ComputePublicAdvertisedPrefix].
func NewComputePublicAdvertisedPrefix(name string, args ComputePublicAdvertisedPrefixArgs) *ComputePublicAdvertisedPrefix {
	return &ComputePublicAdvertisedPrefix{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputePublicAdvertisedPrefix)(nil)

// ComputePublicAdvertisedPrefix represents the Terraform resource google_compute_public_advertised_prefix.
type ComputePublicAdvertisedPrefix struct {
	Name      string
	Args      ComputePublicAdvertisedPrefixArgs
	state     *computePublicAdvertisedPrefixState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputePublicAdvertisedPrefix].
func (cpap *ComputePublicAdvertisedPrefix) Type() string {
	return "google_compute_public_advertised_prefix"
}

// LocalName returns the local name for [ComputePublicAdvertisedPrefix].
func (cpap *ComputePublicAdvertisedPrefix) LocalName() string {
	return cpap.Name
}

// Configuration returns the configuration (args) for [ComputePublicAdvertisedPrefix].
func (cpap *ComputePublicAdvertisedPrefix) Configuration() interface{} {
	return cpap.Args
}

// DependOn is used for other resources to depend on [ComputePublicAdvertisedPrefix].
func (cpap *ComputePublicAdvertisedPrefix) DependOn() terra.Reference {
	return terra.ReferenceResource(cpap)
}

// Dependencies returns the list of resources [ComputePublicAdvertisedPrefix] depends_on.
func (cpap *ComputePublicAdvertisedPrefix) Dependencies() terra.Dependencies {
	return cpap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputePublicAdvertisedPrefix].
func (cpap *ComputePublicAdvertisedPrefix) LifecycleManagement() *terra.Lifecycle {
	return cpap.Lifecycle
}

// Attributes returns the attributes for [ComputePublicAdvertisedPrefix].
func (cpap *ComputePublicAdvertisedPrefix) Attributes() computePublicAdvertisedPrefixAttributes {
	return computePublicAdvertisedPrefixAttributes{ref: terra.ReferenceResource(cpap)}
}

// ImportState imports the given attribute values into [ComputePublicAdvertisedPrefix]'s state.
func (cpap *ComputePublicAdvertisedPrefix) ImportState(av io.Reader) error {
	cpap.state = &computePublicAdvertisedPrefixState{}
	if err := json.NewDecoder(av).Decode(cpap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cpap.Type(), cpap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputePublicAdvertisedPrefix] has state.
func (cpap *ComputePublicAdvertisedPrefix) State() (*computePublicAdvertisedPrefixState, bool) {
	return cpap.state, cpap.state != nil
}

// StateMust returns the state for [ComputePublicAdvertisedPrefix]. Panics if the state is nil.
func (cpap *ComputePublicAdvertisedPrefix) StateMust() *computePublicAdvertisedPrefixState {
	if cpap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cpap.Type(), cpap.LocalName()))
	}
	return cpap.state
}

// ComputePublicAdvertisedPrefixArgs contains the configurations for google_compute_public_advertised_prefix.
type ComputePublicAdvertisedPrefixArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DnsVerificationIp: string, required
	DnsVerificationIp terra.StringValue `hcl:"dns_verification_ip,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpCidrRange: string, required
	IpCidrRange terra.StringValue `hcl:"ip_cidr_range,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *computepublicadvertisedprefix.Timeouts `hcl:"timeouts,block"`
}
type computePublicAdvertisedPrefixAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_public_advertised_prefix.
func (cpap computePublicAdvertisedPrefixAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cpap.ref.Append("description"))
}

// DnsVerificationIp returns a reference to field dns_verification_ip of google_compute_public_advertised_prefix.
func (cpap computePublicAdvertisedPrefixAttributes) DnsVerificationIp() terra.StringValue {
	return terra.ReferenceAsString(cpap.ref.Append("dns_verification_ip"))
}

// Id returns a reference to field id of google_compute_public_advertised_prefix.
func (cpap computePublicAdvertisedPrefixAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cpap.ref.Append("id"))
}

// IpCidrRange returns a reference to field ip_cidr_range of google_compute_public_advertised_prefix.
func (cpap computePublicAdvertisedPrefixAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceAsString(cpap.ref.Append("ip_cidr_range"))
}

// Name returns a reference to field name of google_compute_public_advertised_prefix.
func (cpap computePublicAdvertisedPrefixAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cpap.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_public_advertised_prefix.
func (cpap computePublicAdvertisedPrefixAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cpap.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_public_advertised_prefix.
func (cpap computePublicAdvertisedPrefixAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cpap.ref.Append("self_link"))
}

func (cpap computePublicAdvertisedPrefixAttributes) Timeouts() computepublicadvertisedprefix.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computepublicadvertisedprefix.TimeoutsAttributes](cpap.ref.Append("timeouts"))
}

type computePublicAdvertisedPrefixState struct {
	Description       string                                       `json:"description"`
	DnsVerificationIp string                                       `json:"dns_verification_ip"`
	Id                string                                       `json:"id"`
	IpCidrRange       string                                       `json:"ip_cidr_range"`
	Name              string                                       `json:"name"`
	Project           string                                       `json:"project"`
	SelfLink          string                                       `json:"self_link"`
	Timeouts          *computepublicadvertisedprefix.TimeoutsState `json:"timeouts"`
}
