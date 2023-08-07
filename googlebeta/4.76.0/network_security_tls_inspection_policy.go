// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	networksecuritytlsinspectionpolicy "github.com/golingon/terraproviders/googlebeta/4.76.0/networksecuritytlsinspectionpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkSecurityTlsInspectionPolicy creates a new instance of [NetworkSecurityTlsInspectionPolicy].
func NewNetworkSecurityTlsInspectionPolicy(name string, args NetworkSecurityTlsInspectionPolicyArgs) *NetworkSecurityTlsInspectionPolicy {
	return &NetworkSecurityTlsInspectionPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkSecurityTlsInspectionPolicy)(nil)

// NetworkSecurityTlsInspectionPolicy represents the Terraform resource google_network_security_tls_inspection_policy.
type NetworkSecurityTlsInspectionPolicy struct {
	Name      string
	Args      NetworkSecurityTlsInspectionPolicyArgs
	state     *networkSecurityTlsInspectionPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkSecurityTlsInspectionPolicy].
func (nstip *NetworkSecurityTlsInspectionPolicy) Type() string {
	return "google_network_security_tls_inspection_policy"
}

// LocalName returns the local name for [NetworkSecurityTlsInspectionPolicy].
func (nstip *NetworkSecurityTlsInspectionPolicy) LocalName() string {
	return nstip.Name
}

// Configuration returns the configuration (args) for [NetworkSecurityTlsInspectionPolicy].
func (nstip *NetworkSecurityTlsInspectionPolicy) Configuration() interface{} {
	return nstip.Args
}

// DependOn is used for other resources to depend on [NetworkSecurityTlsInspectionPolicy].
func (nstip *NetworkSecurityTlsInspectionPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(nstip)
}

// Dependencies returns the list of resources [NetworkSecurityTlsInspectionPolicy] depends_on.
func (nstip *NetworkSecurityTlsInspectionPolicy) Dependencies() terra.Dependencies {
	return nstip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkSecurityTlsInspectionPolicy].
func (nstip *NetworkSecurityTlsInspectionPolicy) LifecycleManagement() *terra.Lifecycle {
	return nstip.Lifecycle
}

// Attributes returns the attributes for [NetworkSecurityTlsInspectionPolicy].
func (nstip *NetworkSecurityTlsInspectionPolicy) Attributes() networkSecurityTlsInspectionPolicyAttributes {
	return networkSecurityTlsInspectionPolicyAttributes{ref: terra.ReferenceResource(nstip)}
}

// ImportState imports the given attribute values into [NetworkSecurityTlsInspectionPolicy]'s state.
func (nstip *NetworkSecurityTlsInspectionPolicy) ImportState(av io.Reader) error {
	nstip.state = &networkSecurityTlsInspectionPolicyState{}
	if err := json.NewDecoder(av).Decode(nstip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nstip.Type(), nstip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkSecurityTlsInspectionPolicy] has state.
func (nstip *NetworkSecurityTlsInspectionPolicy) State() (*networkSecurityTlsInspectionPolicyState, bool) {
	return nstip.state, nstip.state != nil
}

// StateMust returns the state for [NetworkSecurityTlsInspectionPolicy]. Panics if the state is nil.
func (nstip *NetworkSecurityTlsInspectionPolicy) StateMust() *networkSecurityTlsInspectionPolicyState {
	if nstip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nstip.Type(), nstip.LocalName()))
	}
	return nstip.state
}

// NetworkSecurityTlsInspectionPolicyArgs contains the configurations for google_network_security_tls_inspection_policy.
type NetworkSecurityTlsInspectionPolicyArgs struct {
	// CaPool: string, required
	CaPool terra.StringValue `hcl:"ca_pool,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ExcludePublicCaSet: bool, optional
	ExcludePublicCaSet terra.BoolValue `hcl:"exclude_public_ca_set,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *networksecuritytlsinspectionpolicy.Timeouts `hcl:"timeouts,block"`
}
type networkSecurityTlsInspectionPolicyAttributes struct {
	ref terra.Reference
}

// CaPool returns a reference to field ca_pool of google_network_security_tls_inspection_policy.
func (nstip networkSecurityTlsInspectionPolicyAttributes) CaPool() terra.StringValue {
	return terra.ReferenceAsString(nstip.ref.Append("ca_pool"))
}

// CreateTime returns a reference to field create_time of google_network_security_tls_inspection_policy.
func (nstip networkSecurityTlsInspectionPolicyAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nstip.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_security_tls_inspection_policy.
func (nstip networkSecurityTlsInspectionPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nstip.ref.Append("description"))
}

// ExcludePublicCaSet returns a reference to field exclude_public_ca_set of google_network_security_tls_inspection_policy.
func (nstip networkSecurityTlsInspectionPolicyAttributes) ExcludePublicCaSet() terra.BoolValue {
	return terra.ReferenceAsBool(nstip.ref.Append("exclude_public_ca_set"))
}

// Id returns a reference to field id of google_network_security_tls_inspection_policy.
func (nstip networkSecurityTlsInspectionPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nstip.ref.Append("id"))
}

// Location returns a reference to field location of google_network_security_tls_inspection_policy.
func (nstip networkSecurityTlsInspectionPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nstip.ref.Append("location"))
}

// Name returns a reference to field name of google_network_security_tls_inspection_policy.
func (nstip networkSecurityTlsInspectionPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nstip.ref.Append("name"))
}

// Project returns a reference to field project of google_network_security_tls_inspection_policy.
func (nstip networkSecurityTlsInspectionPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nstip.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_network_security_tls_inspection_policy.
func (nstip networkSecurityTlsInspectionPolicyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nstip.ref.Append("update_time"))
}

func (nstip networkSecurityTlsInspectionPolicyAttributes) Timeouts() networksecuritytlsinspectionpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networksecuritytlsinspectionpolicy.TimeoutsAttributes](nstip.ref.Append("timeouts"))
}

type networkSecurityTlsInspectionPolicyState struct {
	CaPool             string                                            `json:"ca_pool"`
	CreateTime         string                                            `json:"create_time"`
	Description        string                                            `json:"description"`
	ExcludePublicCaSet bool                                              `json:"exclude_public_ca_set"`
	Id                 string                                            `json:"id"`
	Location           string                                            `json:"location"`
	Name               string                                            `json:"name"`
	Project            string                                            `json:"project"`
	UpdateTime         string                                            `json:"update_time"`
	Timeouts           *networksecuritytlsinspectionpolicy.TimeoutsState `json:"timeouts"`
}
