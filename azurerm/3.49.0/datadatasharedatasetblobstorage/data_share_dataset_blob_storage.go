// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datadatasharedatasetblobstorage

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type StorageAccount struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type StorageAccountAttributes struct {
	ref terra.Reference
}

func (sa StorageAccountAttributes) InternalRef() terra.Reference {
	return sa.ref
}

func (sa StorageAccountAttributes) InternalWithRef(ref terra.Reference) StorageAccountAttributes {
	return StorageAccountAttributes{ref: ref}
}

func (sa StorageAccountAttributes) InternalTokens() hclwrite.Tokens {
	return sa.ref.InternalTokens()
}

func (sa StorageAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("name"))
}

func (sa StorageAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("resource_group_name"))
}

func (sa StorageAccountAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("subscription_id"))
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

type StorageAccountState struct {
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	SubscriptionId    string `json:"subscription_id"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
