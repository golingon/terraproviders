// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	activedirectorydomain "github.com/golingon/terraproviders/google/4.76.0/activedirectorydomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewActiveDirectoryDomain creates a new instance of [ActiveDirectoryDomain].
func NewActiveDirectoryDomain(name string, args ActiveDirectoryDomainArgs) *ActiveDirectoryDomain {
	return &ActiveDirectoryDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ActiveDirectoryDomain)(nil)

// ActiveDirectoryDomain represents the Terraform resource google_active_directory_domain.
type ActiveDirectoryDomain struct {
	Name      string
	Args      ActiveDirectoryDomainArgs
	state     *activeDirectoryDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ActiveDirectoryDomain].
func (add *ActiveDirectoryDomain) Type() string {
	return "google_active_directory_domain"
}

// LocalName returns the local name for [ActiveDirectoryDomain].
func (add *ActiveDirectoryDomain) LocalName() string {
	return add.Name
}

// Configuration returns the configuration (args) for [ActiveDirectoryDomain].
func (add *ActiveDirectoryDomain) Configuration() interface{} {
	return add.Args
}

// DependOn is used for other resources to depend on [ActiveDirectoryDomain].
func (add *ActiveDirectoryDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(add)
}

// Dependencies returns the list of resources [ActiveDirectoryDomain] depends_on.
func (add *ActiveDirectoryDomain) Dependencies() terra.Dependencies {
	return add.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ActiveDirectoryDomain].
func (add *ActiveDirectoryDomain) LifecycleManagement() *terra.Lifecycle {
	return add.Lifecycle
}

// Attributes returns the attributes for [ActiveDirectoryDomain].
func (add *ActiveDirectoryDomain) Attributes() activeDirectoryDomainAttributes {
	return activeDirectoryDomainAttributes{ref: terra.ReferenceResource(add)}
}

// ImportState imports the given attribute values into [ActiveDirectoryDomain]'s state.
func (add *ActiveDirectoryDomain) ImportState(av io.Reader) error {
	add.state = &activeDirectoryDomainState{}
	if err := json.NewDecoder(av).Decode(add.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", add.Type(), add.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ActiveDirectoryDomain] has state.
func (add *ActiveDirectoryDomain) State() (*activeDirectoryDomainState, bool) {
	return add.state, add.state != nil
}

// StateMust returns the state for [ActiveDirectoryDomain]. Panics if the state is nil.
func (add *ActiveDirectoryDomain) StateMust() *activeDirectoryDomainState {
	if add.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", add.Type(), add.LocalName()))
	}
	return add.state
}

// ActiveDirectoryDomainArgs contains the configurations for google_active_directory_domain.
type ActiveDirectoryDomainArgs struct {
	// Admin: string, optional
	Admin terra.StringValue `hcl:"admin,attr"`
	// AuthorizedNetworks: set of string, optional
	AuthorizedNetworks terra.SetValue[terra.StringValue] `hcl:"authorized_networks,attr"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Locations: list of string, required
	Locations terra.ListValue[terra.StringValue] `hcl:"locations,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ReservedIpRange: string, required
	ReservedIpRange terra.StringValue `hcl:"reserved_ip_range,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *activedirectorydomain.Timeouts `hcl:"timeouts,block"`
}
type activeDirectoryDomainAttributes struct {
	ref terra.Reference
}

// Admin returns a reference to field admin of google_active_directory_domain.
func (add activeDirectoryDomainAttributes) Admin() terra.StringValue {
	return terra.ReferenceAsString(add.ref.Append("admin"))
}

// AuthorizedNetworks returns a reference to field authorized_networks of google_active_directory_domain.
func (add activeDirectoryDomainAttributes) AuthorizedNetworks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](add.ref.Append("authorized_networks"))
}

// DomainName returns a reference to field domain_name of google_active_directory_domain.
func (add activeDirectoryDomainAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(add.ref.Append("domain_name"))
}

// Fqdn returns a reference to field fqdn of google_active_directory_domain.
func (add activeDirectoryDomainAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(add.ref.Append("fqdn"))
}

// Id returns a reference to field id of google_active_directory_domain.
func (add activeDirectoryDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(add.ref.Append("id"))
}

// Labels returns a reference to field labels of google_active_directory_domain.
func (add activeDirectoryDomainAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](add.ref.Append("labels"))
}

// Locations returns a reference to field locations of google_active_directory_domain.
func (add activeDirectoryDomainAttributes) Locations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](add.ref.Append("locations"))
}

// Name returns a reference to field name of google_active_directory_domain.
func (add activeDirectoryDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(add.ref.Append("name"))
}

// Project returns a reference to field project of google_active_directory_domain.
func (add activeDirectoryDomainAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(add.ref.Append("project"))
}

// ReservedIpRange returns a reference to field reserved_ip_range of google_active_directory_domain.
func (add activeDirectoryDomainAttributes) ReservedIpRange() terra.StringValue {
	return terra.ReferenceAsString(add.ref.Append("reserved_ip_range"))
}

func (add activeDirectoryDomainAttributes) Timeouts() activedirectorydomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[activedirectorydomain.TimeoutsAttributes](add.ref.Append("timeouts"))
}

type activeDirectoryDomainState struct {
	Admin              string                               `json:"admin"`
	AuthorizedNetworks []string                             `json:"authorized_networks"`
	DomainName         string                               `json:"domain_name"`
	Fqdn               string                               `json:"fqdn"`
	Id                 string                               `json:"id"`
	Labels             map[string]string                    `json:"labels"`
	Locations          []string                             `json:"locations"`
	Name               string                               `json:"name"`
	Project            string                               `json:"project"`
	ReservedIpRange    string                               `json:"reserved_ip_range"`
	Timeouts           *activedirectorydomain.TimeoutsState `json:"timeouts"`
}
