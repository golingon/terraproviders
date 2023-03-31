// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dnsptrrecord "github.com/golingon/terraproviders/azurerm/3.49.0/dnsptrrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDnsPtrRecord(name string, args DnsPtrRecordArgs) *DnsPtrRecord {
	return &DnsPtrRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsPtrRecord)(nil)

type DnsPtrRecord struct {
	Name  string
	Args  DnsPtrRecordArgs
	state *dnsPtrRecordState
}

func (dpr *DnsPtrRecord) Type() string {
	return "azurerm_dns_ptr_record"
}

func (dpr *DnsPtrRecord) LocalName() string {
	return dpr.Name
}

func (dpr *DnsPtrRecord) Configuration() interface{} {
	return dpr.Args
}

func (dpr *DnsPtrRecord) Attributes() dnsPtrRecordAttributes {
	return dnsPtrRecordAttributes{ref: terra.ReferenceResource(dpr)}
}

func (dpr *DnsPtrRecord) ImportState(av io.Reader) error {
	dpr.state = &dnsPtrRecordState{}
	if err := json.NewDecoder(av).Decode(dpr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpr.Type(), dpr.LocalName(), err)
	}
	return nil
}

func (dpr *DnsPtrRecord) State() (*dnsPtrRecordState, bool) {
	return dpr.state, dpr.state != nil
}

func (dpr *DnsPtrRecord) StateMust() *dnsPtrRecordState {
	if dpr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpr.Type(), dpr.LocalName()))
	}
	return dpr.state
}

func (dpr *DnsPtrRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(dpr)
}

type DnsPtrRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Records: set of string, required
	Records terra.SetValue[terra.StringValue] `hcl:"records,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Ttl: number, required
	Ttl terra.NumberValue `hcl:"ttl,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dnsptrrecord.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that DnsPtrRecord depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dnsPtrRecordAttributes struct {
	ref terra.Reference
}

func (dpr dnsPtrRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceString(dpr.ref.Append("fqdn"))
}

func (dpr dnsPtrRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dpr.ref.Append("id"))
}

func (dpr dnsPtrRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dpr.ref.Append("name"))
}

func (dpr dnsPtrRecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](dpr.ref.Append("records"))
}

func (dpr dnsPtrRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(dpr.ref.Append("resource_group_name"))
}

func (dpr dnsPtrRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dpr.ref.Append("tags"))
}

func (dpr dnsPtrRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceNumber(dpr.ref.Append("ttl"))
}

func (dpr dnsPtrRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceString(dpr.ref.Append("zone_name"))
}

func (dpr dnsPtrRecordAttributes) Timeouts() dnsptrrecord.TimeoutsAttributes {
	return terra.ReferenceSingle[dnsptrrecord.TimeoutsAttributes](dpr.ref.Append("timeouts"))
}

type dnsPtrRecordState struct {
	Fqdn              string                      `json:"fqdn"`
	Id                string                      `json:"id"`
	Name              string                      `json:"name"`
	Records           []string                    `json:"records"`
	ResourceGroupName string                      `json:"resource_group_name"`
	Tags              map[string]string           `json:"tags"`
	Ttl               float64                     `json:"ttl"`
	ZoneName          string                      `json:"zone_name"`
	Timeouts          *dnsptrrecord.TimeoutsState `json:"timeouts"`
}
