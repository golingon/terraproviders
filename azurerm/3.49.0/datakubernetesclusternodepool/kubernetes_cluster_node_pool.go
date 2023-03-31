// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datakubernetesclusternodepool

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type UpgradeSettings struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type UpgradeSettingsAttributes struct {
	ref terra.Reference
}

func (us UpgradeSettingsAttributes) InternalRef() terra.Reference {
	return us.ref
}

func (us UpgradeSettingsAttributes) InternalWithRef(ref terra.Reference) UpgradeSettingsAttributes {
	return UpgradeSettingsAttributes{ref: ref}
}

func (us UpgradeSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return us.ref.InternalTokens()
}

func (us UpgradeSettingsAttributes) MaxSurge() terra.StringValue {
	return terra.ReferenceString(us.ref.Append("max_surge"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

type UpgradeSettingsState struct {
	MaxSurge string `json:"max_surge"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
