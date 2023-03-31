// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednstxtrecord "github.com/golingon/terraproviders/azurerm/3.49.0/privatednstxtrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewPrivateDnsTxtRecord(name string, args PrivateDnsTxtRecordArgs) *PrivateDnsTxtRecord {
	return &PrivateDnsTxtRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsTxtRecord)(nil)

type PrivateDnsTxtRecord struct {
	Name  string
	Args  PrivateDnsTxtRecordArgs
	state *privateDnsTxtRecordState
}

func (pdtr *PrivateDnsTxtRecord) Type() string {
	return "azurerm_private_dns_txt_record"
}

func (pdtr *PrivateDnsTxtRecord) LocalName() string {
	return pdtr.Name
}

func (pdtr *PrivateDnsTxtRecord) Configuration() interface{} {
	return pdtr.Args
}

func (pdtr *PrivateDnsTxtRecord) Attributes() privateDnsTxtRecordAttributes {
	return privateDnsTxtRecordAttributes{ref: terra.ReferenceResource(pdtr)}
}

func (pdtr *PrivateDnsTxtRecord) ImportState(av io.Reader) error {
	pdtr.state = &privateDnsTxtRecordState{}
	if err := json.NewDecoder(av).Decode(pdtr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdtr.Type(), pdtr.LocalName(), err)
	}
	return nil
}

func (pdtr *PrivateDnsTxtRecord) State() (*privateDnsTxtRecordState, bool) {
	return pdtr.state, pdtr.state != nil
}

func (pdtr *PrivateDnsTxtRecord) StateMust() *privateDnsTxtRecordState {
	if pdtr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdtr.Type(), pdtr.LocalName()))
	}
	return pdtr.state
}

func (pdtr *PrivateDnsTxtRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(pdtr)
}

type PrivateDnsTxtRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Ttl: number, required
	Ttl terra.NumberValue `hcl:"ttl,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Record: min=1
	Record []privatednstxtrecord.Record `hcl:"record,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *privatednstxtrecord.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that PrivateDnsTxtRecord depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type privateDnsTxtRecordAttributes struct {
	ref terra.Reference
}

func (pdtr privateDnsTxtRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceString(pdtr.ref.Append("fqdn"))
}

func (pdtr privateDnsTxtRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceString(pdtr.ref.Append("id"))
}

func (pdtr privateDnsTxtRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceString(pdtr.ref.Append("name"))
}

func (pdtr privateDnsTxtRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(pdtr.ref.Append("resource_group_name"))
}

func (pdtr privateDnsTxtRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](pdtr.ref.Append("tags"))
}

func (pdtr privateDnsTxtRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceNumber(pdtr.ref.Append("ttl"))
}

func (pdtr privateDnsTxtRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceString(pdtr.ref.Append("zone_name"))
}

func (pdtr privateDnsTxtRecordAttributes) Record() terra.SetValue[privatednstxtrecord.RecordAttributes] {
	return terra.ReferenceSet[privatednstxtrecord.RecordAttributes](pdtr.ref.Append("record"))
}

func (pdtr privateDnsTxtRecordAttributes) Timeouts() privatednstxtrecord.TimeoutsAttributes {
	return terra.ReferenceSingle[privatednstxtrecord.TimeoutsAttributes](pdtr.ref.Append("timeouts"))
}

type privateDnsTxtRecordState struct {
	Fqdn              string                             `json:"fqdn"`
	Id                string                             `json:"id"`
	Name              string                             `json:"name"`
	ResourceGroupName string                             `json:"resource_group_name"`
	Tags              map[string]string                  `json:"tags"`
	Ttl               float64                            `json:"ttl"`
	ZoneName          string                             `json:"zone_name"`
	Record            []privatednstxtrecord.RecordState  `json:"record"`
	Timeouts          *privatednstxtrecord.TimeoutsState `json:"timeouts"`
}
