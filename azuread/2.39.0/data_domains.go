// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	datadomains "github.com/golingon/terraproviders/azuread/2.39.0/datadomains"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDomains creates a new instance of [DataDomains].
func NewDataDomains(name string, args DataDomainsArgs) *DataDomains {
	return &DataDomains{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDomains)(nil)

// DataDomains represents the Terraform data resource azuread_domains.
type DataDomains struct {
	Name string
	Args DataDomainsArgs
}

// DataSource returns the Terraform object type for [DataDomains].
func (d *DataDomains) DataSource() string {
	return "azuread_domains"
}

// LocalName returns the local name for [DataDomains].
func (d *DataDomains) LocalName() string {
	return d.Name
}

// Configuration returns the configuration (args) for [DataDomains].
func (d *DataDomains) Configuration() interface{} {
	return d.Args
}

// Attributes returns the attributes for [DataDomains].
func (d *DataDomains) Attributes() dataDomainsAttributes {
	return dataDomainsAttributes{ref: terra.ReferenceDataResource(d)}
}

// DataDomainsArgs contains the configurations for azuread_domains.
type DataDomainsArgs struct {
	// AdminManaged: bool, optional
	AdminManaged terra.BoolValue `hcl:"admin_managed,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludeUnverified: bool, optional
	IncludeUnverified terra.BoolValue `hcl:"include_unverified,attr"`
	// OnlyDefault: bool, optional
	OnlyDefault terra.BoolValue `hcl:"only_default,attr"`
	// OnlyInitial: bool, optional
	OnlyInitial terra.BoolValue `hcl:"only_initial,attr"`
	// OnlyRoot: bool, optional
	OnlyRoot terra.BoolValue `hcl:"only_root,attr"`
	// SupportsServices: list of string, optional
	SupportsServices terra.ListValue[terra.StringValue] `hcl:"supports_services,attr"`
	// Domains: min=0
	Domains []datadomains.Domains `hcl:"domains,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadomains.Timeouts `hcl:"timeouts,block"`
}
type dataDomainsAttributes struct {
	ref terra.Reference
}

// AdminManaged returns a reference to field admin_managed of azuread_domains.
func (d dataDomainsAttributes) AdminManaged() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("admin_managed"))
}

// Id returns a reference to field id of azuread_domains.
func (d dataDomainsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("id"))
}

// IncludeUnverified returns a reference to field include_unverified of azuread_domains.
func (d dataDomainsAttributes) IncludeUnverified() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("include_unverified"))
}

// OnlyDefault returns a reference to field only_default of azuread_domains.
func (d dataDomainsAttributes) OnlyDefault() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("only_default"))
}

// OnlyInitial returns a reference to field only_initial of azuread_domains.
func (d dataDomainsAttributes) OnlyInitial() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("only_initial"))
}

// OnlyRoot returns a reference to field only_root of azuread_domains.
func (d dataDomainsAttributes) OnlyRoot() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("only_root"))
}

// SupportsServices returns a reference to field supports_services of azuread_domains.
func (d dataDomainsAttributes) SupportsServices() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("supports_services"))
}

func (d dataDomainsAttributes) Domains() terra.ListValue[datadomains.DomainsAttributes] {
	return terra.ReferenceAsList[datadomains.DomainsAttributes](d.ref.Append("domains"))
}

func (d dataDomainsAttributes) Timeouts() datadomains.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadomains.TimeoutsAttributes](d.ref.Append("timeouts"))
}
