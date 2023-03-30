// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCloudfrontKeyGroup(name string, args CloudfrontKeyGroupArgs) *CloudfrontKeyGroup {
	return &CloudfrontKeyGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfrontKeyGroup)(nil)

type CloudfrontKeyGroup struct {
	Name  string
	Args  CloudfrontKeyGroupArgs
	state *cloudfrontKeyGroupState
}

func (ckg *CloudfrontKeyGroup) Type() string {
	return "aws_cloudfront_key_group"
}

func (ckg *CloudfrontKeyGroup) LocalName() string {
	return ckg.Name
}

func (ckg *CloudfrontKeyGroup) Configuration() interface{} {
	return ckg.Args
}

func (ckg *CloudfrontKeyGroup) Attributes() cloudfrontKeyGroupAttributes {
	return cloudfrontKeyGroupAttributes{ref: terra.ReferenceResource(ckg)}
}

func (ckg *CloudfrontKeyGroup) ImportState(av io.Reader) error {
	ckg.state = &cloudfrontKeyGroupState{}
	if err := json.NewDecoder(av).Decode(ckg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ckg.Type(), ckg.LocalName(), err)
	}
	return nil
}

func (ckg *CloudfrontKeyGroup) State() (*cloudfrontKeyGroupState, bool) {
	return ckg.state, ckg.state != nil
}

func (ckg *CloudfrontKeyGroup) StateMust() *cloudfrontKeyGroupState {
	if ckg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ckg.Type(), ckg.LocalName()))
	}
	return ckg.state
}

func (ckg *CloudfrontKeyGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ckg)
}

type CloudfrontKeyGroupArgs struct {
	// Comment: string, optional
	Comment terra.StringValue `hcl:"comment,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Items: set of string, required
	Items terra.SetValue[terra.StringValue] `hcl:"items,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// DependsOn contains resources that CloudfrontKeyGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type cloudfrontKeyGroupAttributes struct {
	ref terra.Reference
}

func (ckg cloudfrontKeyGroupAttributes) Comment() terra.StringValue {
	return terra.ReferenceString(ckg.ref.Append("comment"))
}

func (ckg cloudfrontKeyGroupAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(ckg.ref.Append("etag"))
}

func (ckg cloudfrontKeyGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ckg.ref.Append("id"))
}

func (ckg cloudfrontKeyGroupAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](ckg.ref.Append("items"))
}

func (ckg cloudfrontKeyGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ckg.ref.Append("name"))
}

type cloudfrontKeyGroupState struct {
	Comment string   `json:"comment"`
	Etag    string   `json:"etag"`
	Id      string   `json:"id"`
	Items   []string `json:"items"`
	Name    string   `json:"name"`
}
