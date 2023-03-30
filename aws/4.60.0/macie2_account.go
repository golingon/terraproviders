// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewMacie2Account(name string, args Macie2AccountArgs) *Macie2Account {
	return &Macie2Account{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Macie2Account)(nil)

type Macie2Account struct {
	Name  string
	Args  Macie2AccountArgs
	state *macie2AccountState
}

func (ma *Macie2Account) Type() string {
	return "aws_macie2_account"
}

func (ma *Macie2Account) LocalName() string {
	return ma.Name
}

func (ma *Macie2Account) Configuration() interface{} {
	return ma.Args
}

func (ma *Macie2Account) Attributes() macie2AccountAttributes {
	return macie2AccountAttributes{ref: terra.ReferenceResource(ma)}
}

func (ma *Macie2Account) ImportState(av io.Reader) error {
	ma.state = &macie2AccountState{}
	if err := json.NewDecoder(av).Decode(ma.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ma.Type(), ma.LocalName(), err)
	}
	return nil
}

func (ma *Macie2Account) State() (*macie2AccountState, bool) {
	return ma.state, ma.state != nil
}

func (ma *Macie2Account) StateMust() *macie2AccountState {
	if ma.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ma.Type(), ma.LocalName()))
	}
	return ma.state
}

func (ma *Macie2Account) DependOn() terra.Reference {
	return terra.ReferenceResource(ma)
}

type Macie2AccountArgs struct {
	// FindingPublishingFrequency: string, optional
	FindingPublishingFrequency terra.StringValue `hcl:"finding_publishing_frequency,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// DependsOn contains resources that Macie2Account depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type macie2AccountAttributes struct {
	ref terra.Reference
}

func (ma macie2AccountAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceString(ma.ref.Append("created_at"))
}

func (ma macie2AccountAttributes) FindingPublishingFrequency() terra.StringValue {
	return terra.ReferenceString(ma.ref.Append("finding_publishing_frequency"))
}

func (ma macie2AccountAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ma.ref.Append("id"))
}

func (ma macie2AccountAttributes) ServiceRole() terra.StringValue {
	return terra.ReferenceString(ma.ref.Append("service_role"))
}

func (ma macie2AccountAttributes) Status() terra.StringValue {
	return terra.ReferenceString(ma.ref.Append("status"))
}

func (ma macie2AccountAttributes) UpdatedAt() terra.StringValue {
	return terra.ReferenceString(ma.ref.Append("updated_at"))
}

type macie2AccountState struct {
	CreatedAt                  string `json:"created_at"`
	FindingPublishingFrequency string `json:"finding_publishing_frequency"`
	Id                         string `json:"id"`
	ServiceRole                string `json:"service_role"`
	Status                     string `json:"status"`
	UpdatedAt                  string `json:"updated_at"`
}
