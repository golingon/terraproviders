// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package medialivemultiplex

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MultiplexSettings struct {
	// MaximumVideoBufferDelayMilliseconds: number, optional
	MaximumVideoBufferDelayMilliseconds terra.NumberValue `hcl:"maximum_video_buffer_delay_milliseconds,attr"`
	// TransportStreamBitrate: number, required
	TransportStreamBitrate terra.NumberValue `hcl:"transport_stream_bitrate,attr" validate:"required"`
	// TransportStreamId: number, required
	TransportStreamId terra.NumberValue `hcl:"transport_stream_id,attr" validate:"required"`
	// TransportStreamReservedBitrate: number, optional
	TransportStreamReservedBitrate terra.NumberValue `hcl:"transport_stream_reserved_bitrate,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type MultiplexSettingsAttributes struct {
	ref terra.Reference
}

func (ms MultiplexSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ms.ref, nil
}

func (ms MultiplexSettingsAttributes) InternalWithRef(ref terra.Reference) MultiplexSettingsAttributes {
	return MultiplexSettingsAttributes{ref: ref}
}

func (ms MultiplexSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ms.ref.InternalTokens()
}

func (ms MultiplexSettingsAttributes) MaximumVideoBufferDelayMilliseconds() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("maximum_video_buffer_delay_milliseconds"))
}

func (ms MultiplexSettingsAttributes) TransportStreamBitrate() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("transport_stream_bitrate"))
}

func (ms MultiplexSettingsAttributes) TransportStreamId() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("transport_stream_id"))
}

func (ms MultiplexSettingsAttributes) TransportStreamReservedBitrate() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("transport_stream_reserved_bitrate"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type MultiplexSettingsState struct {
	MaximumVideoBufferDelayMilliseconds float64 `json:"maximum_video_buffer_delay_milliseconds"`
	TransportStreamBitrate              float64 `json:"transport_stream_bitrate"`
	TransportStreamId                   float64 `json:"transport_stream_id"`
	TransportStreamReservedBitrate      float64 `json:"transport_stream_reserved_bitrate"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
