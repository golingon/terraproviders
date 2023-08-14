// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	servicenetworkingpeereddnsdomain "github.com/golingon/terraproviders/googlebeta/4.77.0/servicenetworkingpeereddnsdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceNetworkingPeeredDnsDomain creates a new instance of [ServiceNetworkingPeeredDnsDomain].
func NewServiceNetworkingPeeredDnsDomain(name string, args ServiceNetworkingPeeredDnsDomainArgs) *ServiceNetworkingPeeredDnsDomain {
	return &ServiceNetworkingPeeredDnsDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceNetworkingPeeredDnsDomain)(nil)

// ServiceNetworkingPeeredDnsDomain represents the Terraform resource google_service_networking_peered_dns_domain.
type ServiceNetworkingPeeredDnsDomain struct {
	Name      string
	Args      ServiceNetworkingPeeredDnsDomainArgs
	state     *serviceNetworkingPeeredDnsDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceNetworkingPeeredDnsDomain].
func (snpdd *ServiceNetworkingPeeredDnsDomain) Type() string {
	return "google_service_networking_peered_dns_domain"
}

// LocalName returns the local name for [ServiceNetworkingPeeredDnsDomain].
func (snpdd *ServiceNetworkingPeeredDnsDomain) LocalName() string {
	return snpdd.Name
}

// Configuration returns the configuration (args) for [ServiceNetworkingPeeredDnsDomain].
func (snpdd *ServiceNetworkingPeeredDnsDomain) Configuration() interface{} {
	return snpdd.Args
}

// DependOn is used for other resources to depend on [ServiceNetworkingPeeredDnsDomain].
func (snpdd *ServiceNetworkingPeeredDnsDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(snpdd)
}

// Dependencies returns the list of resources [ServiceNetworkingPeeredDnsDomain] depends_on.
func (snpdd *ServiceNetworkingPeeredDnsDomain) Dependencies() terra.Dependencies {
	return snpdd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceNetworkingPeeredDnsDomain].
func (snpdd *ServiceNetworkingPeeredDnsDomain) LifecycleManagement() *terra.Lifecycle {
	return snpdd.Lifecycle
}

// Attributes returns the attributes for [ServiceNetworkingPeeredDnsDomain].
func (snpdd *ServiceNetworkingPeeredDnsDomain) Attributes() serviceNetworkingPeeredDnsDomainAttributes {
	return serviceNetworkingPeeredDnsDomainAttributes{ref: terra.ReferenceResource(snpdd)}
}

// ImportState imports the given attribute values into [ServiceNetworkingPeeredDnsDomain]'s state.
func (snpdd *ServiceNetworkingPeeredDnsDomain) ImportState(av io.Reader) error {
	snpdd.state = &serviceNetworkingPeeredDnsDomainState{}
	if err := json.NewDecoder(av).Decode(snpdd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", snpdd.Type(), snpdd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceNetworkingPeeredDnsDomain] has state.
func (snpdd *ServiceNetworkingPeeredDnsDomain) State() (*serviceNetworkingPeeredDnsDomainState, bool) {
	return snpdd.state, snpdd.state != nil
}

// StateMust returns the state for [ServiceNetworkingPeeredDnsDomain]. Panics if the state is nil.
func (snpdd *ServiceNetworkingPeeredDnsDomain) StateMust() *serviceNetworkingPeeredDnsDomainState {
	if snpdd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", snpdd.Type(), snpdd.LocalName()))
	}
	return snpdd.state
}

// ServiceNetworkingPeeredDnsDomainArgs contains the configurations for google_service_networking_peered_dns_domain.
type ServiceNetworkingPeeredDnsDomainArgs struct {
	// DnsSuffix: string, required
	DnsSuffix terra.StringValue `hcl:"dns_suffix,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Service: string, optional
	Service terra.StringValue `hcl:"service,attr"`
	// Timeouts: optional
	Timeouts *servicenetworkingpeereddnsdomain.Timeouts `hcl:"timeouts,block"`
}
type serviceNetworkingPeeredDnsDomainAttributes struct {
	ref terra.Reference
}

// DnsSuffix returns a reference to field dns_suffix of google_service_networking_peered_dns_domain.
func (snpdd serviceNetworkingPeeredDnsDomainAttributes) DnsSuffix() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("dns_suffix"))
}

// Id returns a reference to field id of google_service_networking_peered_dns_domain.
func (snpdd serviceNetworkingPeeredDnsDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("id"))
}

// Name returns a reference to field name of google_service_networking_peered_dns_domain.
func (snpdd serviceNetworkingPeeredDnsDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("name"))
}

// Network returns a reference to field network of google_service_networking_peered_dns_domain.
func (snpdd serviceNetworkingPeeredDnsDomainAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("network"))
}

// Parent returns a reference to field parent of google_service_networking_peered_dns_domain.
func (snpdd serviceNetworkingPeeredDnsDomainAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("parent"))
}

// Project returns a reference to field project of google_service_networking_peered_dns_domain.
func (snpdd serviceNetworkingPeeredDnsDomainAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("project"))
}

// Service returns a reference to field service of google_service_networking_peered_dns_domain.
func (snpdd serviceNetworkingPeeredDnsDomainAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("service"))
}

func (snpdd serviceNetworkingPeeredDnsDomainAttributes) Timeouts() servicenetworkingpeereddnsdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicenetworkingpeereddnsdomain.TimeoutsAttributes](snpdd.ref.Append("timeouts"))
}

type serviceNetworkingPeeredDnsDomainState struct {
	DnsSuffix string                                          `json:"dns_suffix"`
	Id        string                                          `json:"id"`
	Name      string                                          `json:"name"`
	Network   string                                          `json:"network"`
	Parent    string                                          `json:"parent"`
	Project   string                                          `json:"project"`
	Service   string                                          `json:"service"`
	Timeouts  *servicenetworkingpeereddnsdomain.TimeoutsState `json:"timeouts"`
}
