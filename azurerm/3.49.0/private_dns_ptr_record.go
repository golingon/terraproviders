// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednsptrrecord "github.com/golingon/terraproviders/azurerm/3.49.0/privatednsptrrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewPrivateDnsPtrRecord(name string, args PrivateDnsPtrRecordArgs) *PrivateDnsPtrRecord {
	return &PrivateDnsPtrRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsPtrRecord)(nil)

type PrivateDnsPtrRecord struct {
	Name  string
	Args  PrivateDnsPtrRecordArgs
	state *privateDnsPtrRecordState
}

func (pdpr *PrivateDnsPtrRecord) Type() string {
	return "azurerm_private_dns_ptr_record"
}

func (pdpr *PrivateDnsPtrRecord) LocalName() string {
	return pdpr.Name
}

func (pdpr *PrivateDnsPtrRecord) Configuration() interface{} {
	return pdpr.Args
}

func (pdpr *PrivateDnsPtrRecord) Attributes() privateDnsPtrRecordAttributes {
	return privateDnsPtrRecordAttributes{ref: terra.ReferenceResource(pdpr)}
}

func (pdpr *PrivateDnsPtrRecord) ImportState(av io.Reader) error {
	pdpr.state = &privateDnsPtrRecordState{}
	if err := json.NewDecoder(av).Decode(pdpr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdpr.Type(), pdpr.LocalName(), err)
	}
	return nil
}

func (pdpr *PrivateDnsPtrRecord) State() (*privateDnsPtrRecordState, bool) {
	return pdpr.state, pdpr.state != nil
}

func (pdpr *PrivateDnsPtrRecord) StateMust() *privateDnsPtrRecordState {
	if pdpr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdpr.Type(), pdpr.LocalName()))
	}
	return pdpr.state
}

func (pdpr *PrivateDnsPtrRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(pdpr)
}

type PrivateDnsPtrRecordArgs struct {
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
	Timeouts *privatednsptrrecord.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that PrivateDnsPtrRecord depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type privateDnsPtrRecordAttributes struct {
	ref terra.Reference
}

func (pdpr privateDnsPtrRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceString(pdpr.ref.Append("fqdn"))
}

func (pdpr privateDnsPtrRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceString(pdpr.ref.Append("id"))
}

func (pdpr privateDnsPtrRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceString(pdpr.ref.Append("name"))
}

func (pdpr privateDnsPtrRecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](pdpr.ref.Append("records"))
}

func (pdpr privateDnsPtrRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(pdpr.ref.Append("resource_group_name"))
}

func (pdpr privateDnsPtrRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](pdpr.ref.Append("tags"))
}

func (pdpr privateDnsPtrRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceNumber(pdpr.ref.Append("ttl"))
}

func (pdpr privateDnsPtrRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceString(pdpr.ref.Append("zone_name"))
}

func (pdpr privateDnsPtrRecordAttributes) Timeouts() privatednsptrrecord.TimeoutsAttributes {
	return terra.ReferenceSingle[privatednsptrrecord.TimeoutsAttributes](pdpr.ref.Append("timeouts"))
}

type privateDnsPtrRecordState struct {
	Fqdn              string                             `json:"fqdn"`
	Id                string                             `json:"id"`
	Name              string                             `json:"name"`
	Records           []string                           `json:"records"`
	ResourceGroupName string                             `json:"resource_group_name"`
	Tags              map[string]string                  `json:"tags"`
	Ttl               float64                            `json:"ttl"`
	ZoneName          string                             `json:"zone_name"`
	Timeouts          *privatednsptrrecord.TimeoutsState `json:"timeouts"`
}
