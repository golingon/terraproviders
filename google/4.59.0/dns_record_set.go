// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dnsrecordset "github.com/golingon/terraproviders/google/4.59.0/dnsrecordset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDnsRecordSet(name string, args DnsRecordSetArgs) *DnsRecordSet {
	return &DnsRecordSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsRecordSet)(nil)

type DnsRecordSet struct {
	Name  string
	Args  DnsRecordSetArgs
	state *dnsRecordSetState
}

func (drs *DnsRecordSet) Type() string {
	return "google_dns_record_set"
}

func (drs *DnsRecordSet) LocalName() string {
	return drs.Name
}

func (drs *DnsRecordSet) Configuration() interface{} {
	return drs.Args
}

func (drs *DnsRecordSet) Attributes() dnsRecordSetAttributes {
	return dnsRecordSetAttributes{ref: terra.ReferenceResource(drs)}
}

func (drs *DnsRecordSet) ImportState(av io.Reader) error {
	drs.state = &dnsRecordSetState{}
	if err := json.NewDecoder(av).Decode(drs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", drs.Type(), drs.LocalName(), err)
	}
	return nil
}

func (drs *DnsRecordSet) State() (*dnsRecordSetState, bool) {
	return drs.state, drs.state != nil
}

func (drs *DnsRecordSet) StateMust() *dnsRecordSetState {
	if drs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", drs.Type(), drs.LocalName()))
	}
	return drs.state
}

func (drs *DnsRecordSet) DependOn() terra.Reference {
	return terra.ReferenceResource(drs)
}

type DnsRecordSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagedZone: string, required
	ManagedZone terra.StringValue `hcl:"managed_zone,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Rrdatas: list of string, optional
	Rrdatas terra.ListValue[terra.StringValue] `hcl:"rrdatas,attr"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// RoutingPolicy: optional
	RoutingPolicy *dnsrecordset.RoutingPolicy `hcl:"routing_policy,block"`
	// DependsOn contains resources that DnsRecordSet depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dnsRecordSetAttributes struct {
	ref terra.Reference
}

func (drs dnsRecordSetAttributes) Id() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("id"))
}

func (drs dnsRecordSetAttributes) ManagedZone() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("managed_zone"))
}

func (drs dnsRecordSetAttributes) Name() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("name"))
}

func (drs dnsRecordSetAttributes) Project() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("project"))
}

func (drs dnsRecordSetAttributes) Rrdatas() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](drs.ref.Append("rrdatas"))
}

func (drs dnsRecordSetAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceNumber(drs.ref.Append("ttl"))
}

func (drs dnsRecordSetAttributes) Type() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("type"))
}

func (drs dnsRecordSetAttributes) RoutingPolicy() terra.ListValue[dnsrecordset.RoutingPolicyAttributes] {
	return terra.ReferenceList[dnsrecordset.RoutingPolicyAttributes](drs.ref.Append("routing_policy"))
}

type dnsRecordSetState struct {
	Id            string                            `json:"id"`
	ManagedZone   string                            `json:"managed_zone"`
	Name          string                            `json:"name"`
	Project       string                            `json:"project"`
	Rrdatas       []string                          `json:"rrdatas"`
	Ttl           float64                           `json:"ttl"`
	Type          string                            `json:"type"`
	RoutingPolicy []dnsrecordset.RoutingPolicyState `json:"routing_policy"`
}
