// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCognitoUserPoolUiCustomization(name string, args CognitoUserPoolUiCustomizationArgs) *CognitoUserPoolUiCustomization {
	return &CognitoUserPoolUiCustomization{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitoUserPoolUiCustomization)(nil)

type CognitoUserPoolUiCustomization struct {
	Name  string
	Args  CognitoUserPoolUiCustomizationArgs
	state *cognitoUserPoolUiCustomizationState
}

func (cupuc *CognitoUserPoolUiCustomization) Type() string {
	return "aws_cognito_user_pool_ui_customization"
}

func (cupuc *CognitoUserPoolUiCustomization) LocalName() string {
	return cupuc.Name
}

func (cupuc *CognitoUserPoolUiCustomization) Configuration() interface{} {
	return cupuc.Args
}

func (cupuc *CognitoUserPoolUiCustomization) Attributes() cognitoUserPoolUiCustomizationAttributes {
	return cognitoUserPoolUiCustomizationAttributes{ref: terra.ReferenceResource(cupuc)}
}

func (cupuc *CognitoUserPoolUiCustomization) ImportState(av io.Reader) error {
	cupuc.state = &cognitoUserPoolUiCustomizationState{}
	if err := json.NewDecoder(av).Decode(cupuc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cupuc.Type(), cupuc.LocalName(), err)
	}
	return nil
}

func (cupuc *CognitoUserPoolUiCustomization) State() (*cognitoUserPoolUiCustomizationState, bool) {
	return cupuc.state, cupuc.state != nil
}

func (cupuc *CognitoUserPoolUiCustomization) StateMust() *cognitoUserPoolUiCustomizationState {
	if cupuc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cupuc.Type(), cupuc.LocalName()))
	}
	return cupuc.state
}

func (cupuc *CognitoUserPoolUiCustomization) DependOn() terra.Reference {
	return terra.ReferenceResource(cupuc)
}

type CognitoUserPoolUiCustomizationArgs struct {
	// ClientId: string, optional
	ClientId terra.StringValue `hcl:"client_id,attr"`
	// Css: string, optional
	Css terra.StringValue `hcl:"css,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageFile: string, optional
	ImageFile terra.StringValue `hcl:"image_file,attr"`
	// UserPoolId: string, required
	UserPoolId terra.StringValue `hcl:"user_pool_id,attr" validate:"required"`
	// DependsOn contains resources that CognitoUserPoolUiCustomization depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type cognitoUserPoolUiCustomizationAttributes struct {
	ref terra.Reference
}

func (cupuc cognitoUserPoolUiCustomizationAttributes) ClientId() terra.StringValue {
	return terra.ReferenceString(cupuc.ref.Append("client_id"))
}

func (cupuc cognitoUserPoolUiCustomizationAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceString(cupuc.ref.Append("creation_date"))
}

func (cupuc cognitoUserPoolUiCustomizationAttributes) Css() terra.StringValue {
	return terra.ReferenceString(cupuc.ref.Append("css"))
}

func (cupuc cognitoUserPoolUiCustomizationAttributes) CssVersion() terra.StringValue {
	return terra.ReferenceString(cupuc.ref.Append("css_version"))
}

func (cupuc cognitoUserPoolUiCustomizationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cupuc.ref.Append("id"))
}

func (cupuc cognitoUserPoolUiCustomizationAttributes) ImageFile() terra.StringValue {
	return terra.ReferenceString(cupuc.ref.Append("image_file"))
}

func (cupuc cognitoUserPoolUiCustomizationAttributes) ImageUrl() terra.StringValue {
	return terra.ReferenceString(cupuc.ref.Append("image_url"))
}

func (cupuc cognitoUserPoolUiCustomizationAttributes) LastModifiedDate() terra.StringValue {
	return terra.ReferenceString(cupuc.ref.Append("last_modified_date"))
}

func (cupuc cognitoUserPoolUiCustomizationAttributes) UserPoolId() terra.StringValue {
	return terra.ReferenceString(cupuc.ref.Append("user_pool_id"))
}

type cognitoUserPoolUiCustomizationState struct {
	ClientId         string `json:"client_id"`
	CreationDate     string `json:"creation_date"`
	Css              string `json:"css"`
	CssVersion       string `json:"css_version"`
	Id               string `json:"id"`
	ImageFile        string `json:"image_file"`
	ImageUrl         string `json:"image_url"`
	LastModifiedDate string `json:"last_modified_date"`
	UserPoolId       string `json:"user_pool_id"`
}
