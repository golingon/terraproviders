// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSesReceiptFilter(name string, args SesReceiptFilterArgs) *SesReceiptFilter {
	return &SesReceiptFilter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SesReceiptFilter)(nil)

type SesReceiptFilter struct {
	Name  string
	Args  SesReceiptFilterArgs
	state *sesReceiptFilterState
}

func (srf *SesReceiptFilter) Type() string {
	return "aws_ses_receipt_filter"
}

func (srf *SesReceiptFilter) LocalName() string {
	return srf.Name
}

func (srf *SesReceiptFilter) Configuration() interface{} {
	return srf.Args
}

func (srf *SesReceiptFilter) Attributes() sesReceiptFilterAttributes {
	return sesReceiptFilterAttributes{ref: terra.ReferenceResource(srf)}
}

func (srf *SesReceiptFilter) ImportState(av io.Reader) error {
	srf.state = &sesReceiptFilterState{}
	if err := json.NewDecoder(av).Decode(srf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srf.Type(), srf.LocalName(), err)
	}
	return nil
}

func (srf *SesReceiptFilter) State() (*sesReceiptFilterState, bool) {
	return srf.state, srf.state != nil
}

func (srf *SesReceiptFilter) StateMust() *sesReceiptFilterState {
	if srf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srf.Type(), srf.LocalName()))
	}
	return srf.state
}

func (srf *SesReceiptFilter) DependOn() terra.Reference {
	return terra.ReferenceResource(srf)
}

type SesReceiptFilterArgs struct {
	// Cidr: string, required
	Cidr terra.StringValue `hcl:"cidr,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
	// DependsOn contains resources that SesReceiptFilter depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sesReceiptFilterAttributes struct {
	ref terra.Reference
}

func (srf sesReceiptFilterAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(srf.ref.Append("arn"))
}

func (srf sesReceiptFilterAttributes) Cidr() terra.StringValue {
	return terra.ReferenceString(srf.ref.Append("cidr"))
}

func (srf sesReceiptFilterAttributes) Id() terra.StringValue {
	return terra.ReferenceString(srf.ref.Append("id"))
}

func (srf sesReceiptFilterAttributes) Name() terra.StringValue {
	return terra.ReferenceString(srf.ref.Append("name"))
}

func (srf sesReceiptFilterAttributes) Policy() terra.StringValue {
	return terra.ReferenceString(srf.ref.Append("policy"))
}

type sesReceiptFilterState struct {
	Arn    string `json:"arn"`
	Cidr   string `json:"cidr"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Policy string `json:"policy"`
}
