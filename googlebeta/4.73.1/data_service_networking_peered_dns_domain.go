// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataServiceNetworkingPeeredDnsDomain creates a new instance of [DataServiceNetworkingPeeredDnsDomain].
func NewDataServiceNetworkingPeeredDnsDomain(name string, args DataServiceNetworkingPeeredDnsDomainArgs) *DataServiceNetworkingPeeredDnsDomain {
	return &DataServiceNetworkingPeeredDnsDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServiceNetworkingPeeredDnsDomain)(nil)

// DataServiceNetworkingPeeredDnsDomain represents the Terraform data resource google_service_networking_peered_dns_domain.
type DataServiceNetworkingPeeredDnsDomain struct {
	Name string
	Args DataServiceNetworkingPeeredDnsDomainArgs
}

// DataSource returns the Terraform object type for [DataServiceNetworkingPeeredDnsDomain].
func (snpdd *DataServiceNetworkingPeeredDnsDomain) DataSource() string {
	return "google_service_networking_peered_dns_domain"
}

// LocalName returns the local name for [DataServiceNetworkingPeeredDnsDomain].
func (snpdd *DataServiceNetworkingPeeredDnsDomain) LocalName() string {
	return snpdd.Name
}

// Configuration returns the configuration (args) for [DataServiceNetworkingPeeredDnsDomain].
func (snpdd *DataServiceNetworkingPeeredDnsDomain) Configuration() interface{} {
	return snpdd.Args
}

// Attributes returns the attributes for [DataServiceNetworkingPeeredDnsDomain].
func (snpdd *DataServiceNetworkingPeeredDnsDomain) Attributes() dataServiceNetworkingPeeredDnsDomainAttributes {
	return dataServiceNetworkingPeeredDnsDomainAttributes{ref: terra.ReferenceDataResource(snpdd)}
}

// DataServiceNetworkingPeeredDnsDomainArgs contains the configurations for google_service_networking_peered_dns_domain.
type DataServiceNetworkingPeeredDnsDomainArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
}
type dataServiceNetworkingPeeredDnsDomainAttributes struct {
	ref terra.Reference
}

// DnsSuffix returns a reference to field dns_suffix of google_service_networking_peered_dns_domain.
func (snpdd dataServiceNetworkingPeeredDnsDomainAttributes) DnsSuffix() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("dns_suffix"))
}

// Id returns a reference to field id of google_service_networking_peered_dns_domain.
func (snpdd dataServiceNetworkingPeeredDnsDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("id"))
}

// Name returns a reference to field name of google_service_networking_peered_dns_domain.
func (snpdd dataServiceNetworkingPeeredDnsDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("name"))
}

// Network returns a reference to field network of google_service_networking_peered_dns_domain.
func (snpdd dataServiceNetworkingPeeredDnsDomainAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("network"))
}

// Parent returns a reference to field parent of google_service_networking_peered_dns_domain.
func (snpdd dataServiceNetworkingPeeredDnsDomainAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("parent"))
}

// Project returns a reference to field project of google_service_networking_peered_dns_domain.
func (snpdd dataServiceNetworkingPeeredDnsDomainAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("project"))
}

// Service returns a reference to field service of google_service_networking_peered_dns_domain.
func (snpdd dataServiceNetworkingPeeredDnsDomainAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(snpdd.ref.Append("service"))
}
