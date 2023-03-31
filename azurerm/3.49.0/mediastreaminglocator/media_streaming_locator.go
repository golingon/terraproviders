// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mediastreaminglocator

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ContentKey struct {
	// ContentKeyId: string, optional
	ContentKeyId terra.StringValue `hcl:"content_key_id,attr"`
	// LabelReferenceInStreamingPolicy: string, optional
	LabelReferenceInStreamingPolicy terra.StringValue `hcl:"label_reference_in_streaming_policy,attr"`
	// PolicyName: string, optional
	PolicyName terra.StringValue `hcl:"policy_name,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type ContentKeyAttributes struct {
	ref terra.Reference
}

func (ck ContentKeyAttributes) InternalRef() terra.Reference {
	return ck.ref
}

func (ck ContentKeyAttributes) InternalWithRef(ref terra.Reference) ContentKeyAttributes {
	return ContentKeyAttributes{ref: ref}
}

func (ck ContentKeyAttributes) InternalTokens() hclwrite.Tokens {
	return ck.ref.InternalTokens()
}

func (ck ContentKeyAttributes) ContentKeyId() terra.StringValue {
	return terra.ReferenceString(ck.ref.Append("content_key_id"))
}

func (ck ContentKeyAttributes) LabelReferenceInStreamingPolicy() terra.StringValue {
	return terra.ReferenceString(ck.ref.Append("label_reference_in_streaming_policy"))
}

func (ck ContentKeyAttributes) PolicyName() terra.StringValue {
	return terra.ReferenceString(ck.ref.Append("policy_name"))
}

func (ck ContentKeyAttributes) Type() terra.StringValue {
	return terra.ReferenceString(ck.ref.Append("type"))
}

func (ck ContentKeyAttributes) Value() terra.StringValue {
	return terra.ReferenceString(ck.ref.Append("value"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

type ContentKeyState struct {
	ContentKeyId                    string `json:"content_key_id"`
	LabelReferenceInStreamingPolicy string `json:"label_reference_in_streaming_policy"`
	PolicyName                      string `json:"policy_name"`
	Type                            string `json:"type"`
	Value                           string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
}
