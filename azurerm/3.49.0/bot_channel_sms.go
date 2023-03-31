// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	botchannelsms "github.com/golingon/terraproviders/azurerm/3.49.0/botchannelsms"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewBotChannelSms(name string, args BotChannelSmsArgs) *BotChannelSms {
	return &BotChannelSms{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BotChannelSms)(nil)

type BotChannelSms struct {
	Name  string
	Args  BotChannelSmsArgs
	state *botChannelSmsState
}

func (bcs *BotChannelSms) Type() string {
	return "azurerm_bot_channel_sms"
}

func (bcs *BotChannelSms) LocalName() string {
	return bcs.Name
}

func (bcs *BotChannelSms) Configuration() interface{} {
	return bcs.Args
}

func (bcs *BotChannelSms) Attributes() botChannelSmsAttributes {
	return botChannelSmsAttributes{ref: terra.ReferenceResource(bcs)}
}

func (bcs *BotChannelSms) ImportState(av io.Reader) error {
	bcs.state = &botChannelSmsState{}
	if err := json.NewDecoder(av).Decode(bcs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcs.Type(), bcs.LocalName(), err)
	}
	return nil
}

func (bcs *BotChannelSms) State() (*botChannelSmsState, bool) {
	return bcs.state, bcs.state != nil
}

func (bcs *BotChannelSms) StateMust() *botChannelSmsState {
	if bcs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcs.Type(), bcs.LocalName()))
	}
	return bcs.state
}

func (bcs *BotChannelSms) DependOn() terra.Reference {
	return terra.ReferenceResource(bcs)
}

type BotChannelSmsArgs struct {
	// BotName: string, required
	BotName terra.StringValue `hcl:"bot_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// PhoneNumber: string, required
	PhoneNumber terra.StringValue `hcl:"phone_number,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SmsChannelAccountSecurityId: string, required
	SmsChannelAccountSecurityId terra.StringValue `hcl:"sms_channel_account_security_id,attr" validate:"required"`
	// SmsChannelAuthToken: string, required
	SmsChannelAuthToken terra.StringValue `hcl:"sms_channel_auth_token,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *botchannelsms.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that BotChannelSms depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type botChannelSmsAttributes struct {
	ref terra.Reference
}

func (bcs botChannelSmsAttributes) BotName() terra.StringValue {
	return terra.ReferenceString(bcs.ref.Append("bot_name"))
}

func (bcs botChannelSmsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(bcs.ref.Append("id"))
}

func (bcs botChannelSmsAttributes) Location() terra.StringValue {
	return terra.ReferenceString(bcs.ref.Append("location"))
}

func (bcs botChannelSmsAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceString(bcs.ref.Append("phone_number"))
}

func (bcs botChannelSmsAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(bcs.ref.Append("resource_group_name"))
}

func (bcs botChannelSmsAttributes) SmsChannelAccountSecurityId() terra.StringValue {
	return terra.ReferenceString(bcs.ref.Append("sms_channel_account_security_id"))
}

func (bcs botChannelSmsAttributes) SmsChannelAuthToken() terra.StringValue {
	return terra.ReferenceString(bcs.ref.Append("sms_channel_auth_token"))
}

func (bcs botChannelSmsAttributes) Timeouts() botchannelsms.TimeoutsAttributes {
	return terra.ReferenceSingle[botchannelsms.TimeoutsAttributes](bcs.ref.Append("timeouts"))
}

type botChannelSmsState struct {
	BotName                     string                       `json:"bot_name"`
	Id                          string                       `json:"id"`
	Location                    string                       `json:"location"`
	PhoneNumber                 string                       `json:"phone_number"`
	ResourceGroupName           string                       `json:"resource_group_name"`
	SmsChannelAccountSecurityId string                       `json:"sms_channel_account_security_id"`
	SmsChannelAuthToken         string                       `json:"sms_channel_auth_token"`
	Timeouts                    *botchannelsms.TimeoutsState `json:"timeouts"`
}
