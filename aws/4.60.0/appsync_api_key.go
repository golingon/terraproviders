// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewAppsyncApiKey(name string, args AppsyncApiKeyArgs) *AppsyncApiKey {
	return &AppsyncApiKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppsyncApiKey)(nil)

type AppsyncApiKey struct {
	Name  string
	Args  AppsyncApiKeyArgs
	state *appsyncApiKeyState
}

func (aak *AppsyncApiKey) Type() string {
	return "aws_appsync_api_key"
}

func (aak *AppsyncApiKey) LocalName() string {
	return aak.Name
}

func (aak *AppsyncApiKey) Configuration() interface{} {
	return aak.Args
}

func (aak *AppsyncApiKey) Attributes() appsyncApiKeyAttributes {
	return appsyncApiKeyAttributes{ref: terra.ReferenceResource(aak)}
}

func (aak *AppsyncApiKey) ImportState(av io.Reader) error {
	aak.state = &appsyncApiKeyState{}
	if err := json.NewDecoder(av).Decode(aak.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aak.Type(), aak.LocalName(), err)
	}
	return nil
}

func (aak *AppsyncApiKey) State() (*appsyncApiKeyState, bool) {
	return aak.state, aak.state != nil
}

func (aak *AppsyncApiKey) StateMust() *appsyncApiKeyState {
	if aak.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aak.Type(), aak.LocalName()))
	}
	return aak.state
}

func (aak *AppsyncApiKey) DependOn() terra.Reference {
	return terra.ReferenceResource(aak)
}

type AppsyncApiKeyArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Expires: string, optional
	Expires terra.StringValue `hcl:"expires,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// DependsOn contains resources that AppsyncApiKey depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type appsyncApiKeyAttributes struct {
	ref terra.Reference
}

func (aak appsyncApiKeyAttributes) ApiId() terra.StringValue {
	return terra.ReferenceString(aak.ref.Append("api_id"))
}

func (aak appsyncApiKeyAttributes) Description() terra.StringValue {
	return terra.ReferenceString(aak.ref.Append("description"))
}

func (aak appsyncApiKeyAttributes) Expires() terra.StringValue {
	return terra.ReferenceString(aak.ref.Append("expires"))
}

func (aak appsyncApiKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(aak.ref.Append("id"))
}

func (aak appsyncApiKeyAttributes) Key() terra.StringValue {
	return terra.ReferenceString(aak.ref.Append("key"))
}

type appsyncApiKeyState struct {
	ApiId       string `json:"api_id"`
	Description string `json:"description"`
	Expires     string `json:"expires"`
	Id          string `json:"id"`
	Key         string `json:"key"`
}
