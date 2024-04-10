// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	route53domainsregistereddomain "github.com/golingon/terraproviders/aws/5.44.0/route53domainsregistereddomain"
	"io"
)

// NewRoute53DomainsRegisteredDomain creates a new instance of [Route53DomainsRegisteredDomain].
func NewRoute53DomainsRegisteredDomain(name string, args Route53DomainsRegisteredDomainArgs) *Route53DomainsRegisteredDomain {
	return &Route53DomainsRegisteredDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53DomainsRegisteredDomain)(nil)

// Route53DomainsRegisteredDomain represents the Terraform resource aws_route53domains_registered_domain.
type Route53DomainsRegisteredDomain struct {
	Name      string
	Args      Route53DomainsRegisteredDomainArgs
	state     *route53DomainsRegisteredDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53DomainsRegisteredDomain].
func (rrd *Route53DomainsRegisteredDomain) Type() string {
	return "aws_route53domains_registered_domain"
}

// LocalName returns the local name for [Route53DomainsRegisteredDomain].
func (rrd *Route53DomainsRegisteredDomain) LocalName() string {
	return rrd.Name
}

// Configuration returns the configuration (args) for [Route53DomainsRegisteredDomain].
func (rrd *Route53DomainsRegisteredDomain) Configuration() interface{} {
	return rrd.Args
}

// DependOn is used for other resources to depend on [Route53DomainsRegisteredDomain].
func (rrd *Route53DomainsRegisteredDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(rrd)
}

// Dependencies returns the list of resources [Route53DomainsRegisteredDomain] depends_on.
func (rrd *Route53DomainsRegisteredDomain) Dependencies() terra.Dependencies {
	return rrd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53DomainsRegisteredDomain].
func (rrd *Route53DomainsRegisteredDomain) LifecycleManagement() *terra.Lifecycle {
	return rrd.Lifecycle
}

// Attributes returns the attributes for [Route53DomainsRegisteredDomain].
func (rrd *Route53DomainsRegisteredDomain) Attributes() route53DomainsRegisteredDomainAttributes {
	return route53DomainsRegisteredDomainAttributes{ref: terra.ReferenceResource(rrd)}
}

// ImportState imports the given attribute values into [Route53DomainsRegisteredDomain]'s state.
func (rrd *Route53DomainsRegisteredDomain) ImportState(av io.Reader) error {
	rrd.state = &route53DomainsRegisteredDomainState{}
	if err := json.NewDecoder(av).Decode(rrd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rrd.Type(), rrd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53DomainsRegisteredDomain] has state.
func (rrd *Route53DomainsRegisteredDomain) State() (*route53DomainsRegisteredDomainState, bool) {
	return rrd.state, rrd.state != nil
}

// StateMust returns the state for [Route53DomainsRegisteredDomain]. Panics if the state is nil.
func (rrd *Route53DomainsRegisteredDomain) StateMust() *route53DomainsRegisteredDomainState {
	if rrd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rrd.Type(), rrd.LocalName()))
	}
	return rrd.state
}

// Route53DomainsRegisteredDomainArgs contains the configurations for aws_route53domains_registered_domain.
type Route53DomainsRegisteredDomainArgs struct {
	// AdminPrivacy: bool, optional
	AdminPrivacy terra.BoolValue `hcl:"admin_privacy,attr"`
	// AutoRenew: bool, optional
	AutoRenew terra.BoolValue `hcl:"auto_renew,attr"`
	// BillingPrivacy: bool, optional
	BillingPrivacy terra.BoolValue `hcl:"billing_privacy,attr"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RegistrantPrivacy: bool, optional
	RegistrantPrivacy terra.BoolValue `hcl:"registrant_privacy,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TechPrivacy: bool, optional
	TechPrivacy terra.BoolValue `hcl:"tech_privacy,attr"`
	// TransferLock: bool, optional
	TransferLock terra.BoolValue `hcl:"transfer_lock,attr"`
	// AdminContact: optional
	AdminContact *route53domainsregistereddomain.AdminContact `hcl:"admin_contact,block"`
	// BillingContact: optional
	BillingContact *route53domainsregistereddomain.BillingContact `hcl:"billing_contact,block"`
	// NameServer: min=0,max=6
	NameServer []route53domainsregistereddomain.NameServer `hcl:"name_server,block" validate:"min=0,max=6"`
	// RegistrantContact: optional
	RegistrantContact *route53domainsregistereddomain.RegistrantContact `hcl:"registrant_contact,block"`
	// TechContact: optional
	TechContact *route53domainsregistereddomain.TechContact `hcl:"tech_contact,block"`
	// Timeouts: optional
	Timeouts *route53domainsregistereddomain.Timeouts `hcl:"timeouts,block"`
}
type route53DomainsRegisteredDomainAttributes struct {
	ref terra.Reference
}

// AbuseContactEmail returns a reference to field abuse_contact_email of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) AbuseContactEmail() terra.StringValue {
	return terra.ReferenceAsString(rrd.ref.Append("abuse_contact_email"))
}

// AbuseContactPhone returns a reference to field abuse_contact_phone of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) AbuseContactPhone() terra.StringValue {
	return terra.ReferenceAsString(rrd.ref.Append("abuse_contact_phone"))
}

// AdminPrivacy returns a reference to field admin_privacy of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) AdminPrivacy() terra.BoolValue {
	return terra.ReferenceAsBool(rrd.ref.Append("admin_privacy"))
}

// AutoRenew returns a reference to field auto_renew of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) AutoRenew() terra.BoolValue {
	return terra.ReferenceAsBool(rrd.ref.Append("auto_renew"))
}

// BillingPrivacy returns a reference to field billing_privacy of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) BillingPrivacy() terra.BoolValue {
	return terra.ReferenceAsBool(rrd.ref.Append("billing_privacy"))
}

// CreationDate returns a reference to field creation_date of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(rrd.ref.Append("creation_date"))
}

// DomainName returns a reference to field domain_name of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(rrd.ref.Append("domain_name"))
}

// ExpirationDate returns a reference to field expiration_date of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(rrd.ref.Append("expiration_date"))
}

// Id returns a reference to field id of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rrd.ref.Append("id"))
}

// RegistrantPrivacy returns a reference to field registrant_privacy of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) RegistrantPrivacy() terra.BoolValue {
	return terra.ReferenceAsBool(rrd.ref.Append("registrant_privacy"))
}

// RegistrarName returns a reference to field registrar_name of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) RegistrarName() terra.StringValue {
	return terra.ReferenceAsString(rrd.ref.Append("registrar_name"))
}

// RegistrarUrl returns a reference to field registrar_url of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) RegistrarUrl() terra.StringValue {
	return terra.ReferenceAsString(rrd.ref.Append("registrar_url"))
}

// Reseller returns a reference to field reseller of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) Reseller() terra.StringValue {
	return terra.ReferenceAsString(rrd.ref.Append("reseller"))
}

// StatusList returns a reference to field status_list of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) StatusList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rrd.ref.Append("status_list"))
}

// Tags returns a reference to field tags of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rrd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rrd.ref.Append("tags_all"))
}

// TechPrivacy returns a reference to field tech_privacy of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) TechPrivacy() terra.BoolValue {
	return terra.ReferenceAsBool(rrd.ref.Append("tech_privacy"))
}

// TransferLock returns a reference to field transfer_lock of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) TransferLock() terra.BoolValue {
	return terra.ReferenceAsBool(rrd.ref.Append("transfer_lock"))
}

// UpdatedDate returns a reference to field updated_date of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) UpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(rrd.ref.Append("updated_date"))
}

// WhoisServer returns a reference to field whois_server of aws_route53domains_registered_domain.
func (rrd route53DomainsRegisteredDomainAttributes) WhoisServer() terra.StringValue {
	return terra.ReferenceAsString(rrd.ref.Append("whois_server"))
}

func (rrd route53DomainsRegisteredDomainAttributes) AdminContact() terra.ListValue[route53domainsregistereddomain.AdminContactAttributes] {
	return terra.ReferenceAsList[route53domainsregistereddomain.AdminContactAttributes](rrd.ref.Append("admin_contact"))
}

func (rrd route53DomainsRegisteredDomainAttributes) BillingContact() terra.ListValue[route53domainsregistereddomain.BillingContactAttributes] {
	return terra.ReferenceAsList[route53domainsregistereddomain.BillingContactAttributes](rrd.ref.Append("billing_contact"))
}

func (rrd route53DomainsRegisteredDomainAttributes) NameServer() terra.ListValue[route53domainsregistereddomain.NameServerAttributes] {
	return terra.ReferenceAsList[route53domainsregistereddomain.NameServerAttributes](rrd.ref.Append("name_server"))
}

func (rrd route53DomainsRegisteredDomainAttributes) RegistrantContact() terra.ListValue[route53domainsregistereddomain.RegistrantContactAttributes] {
	return terra.ReferenceAsList[route53domainsregistereddomain.RegistrantContactAttributes](rrd.ref.Append("registrant_contact"))
}

func (rrd route53DomainsRegisteredDomainAttributes) TechContact() terra.ListValue[route53domainsregistereddomain.TechContactAttributes] {
	return terra.ReferenceAsList[route53domainsregistereddomain.TechContactAttributes](rrd.ref.Append("tech_contact"))
}

func (rrd route53DomainsRegisteredDomainAttributes) Timeouts() route53domainsregistereddomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[route53domainsregistereddomain.TimeoutsAttributes](rrd.ref.Append("timeouts"))
}

type route53DomainsRegisteredDomainState struct {
	AbuseContactEmail string                                                  `json:"abuse_contact_email"`
	AbuseContactPhone string                                                  `json:"abuse_contact_phone"`
	AdminPrivacy      bool                                                    `json:"admin_privacy"`
	AutoRenew         bool                                                    `json:"auto_renew"`
	BillingPrivacy    bool                                                    `json:"billing_privacy"`
	CreationDate      string                                                  `json:"creation_date"`
	DomainName        string                                                  `json:"domain_name"`
	ExpirationDate    string                                                  `json:"expiration_date"`
	Id                string                                                  `json:"id"`
	RegistrantPrivacy bool                                                    `json:"registrant_privacy"`
	RegistrarName     string                                                  `json:"registrar_name"`
	RegistrarUrl      string                                                  `json:"registrar_url"`
	Reseller          string                                                  `json:"reseller"`
	StatusList        []string                                                `json:"status_list"`
	Tags              map[string]string                                       `json:"tags"`
	TagsAll           map[string]string                                       `json:"tags_all"`
	TechPrivacy       bool                                                    `json:"tech_privacy"`
	TransferLock      bool                                                    `json:"transfer_lock"`
	UpdatedDate       string                                                  `json:"updated_date"`
	WhoisServer       string                                                  `json:"whois_server"`
	AdminContact      []route53domainsregistereddomain.AdminContactState      `json:"admin_contact"`
	BillingContact    []route53domainsregistereddomain.BillingContactState    `json:"billing_contact"`
	NameServer        []route53domainsregistereddomain.NameServerState        `json:"name_server"`
	RegistrantContact []route53domainsregistereddomain.RegistrantContactState `json:"registrant_contact"`
	TechContact       []route53domainsregistereddomain.TechContactState       `json:"tech_contact"`
	Timeouts          *route53domainsregistereddomain.TimeoutsState           `json:"timeouts"`
}
