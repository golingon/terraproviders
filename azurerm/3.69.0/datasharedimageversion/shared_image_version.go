// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datasharedimageversion

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type TargetRegion struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type TargetRegionAttributes struct {
	ref terra.Reference
}

func (tr TargetRegionAttributes) InternalRef() (terra.Reference, error) {
	return tr.ref, nil
}

func (tr TargetRegionAttributes) InternalWithRef(ref terra.Reference) TargetRegionAttributes {
	return TargetRegionAttributes{ref: ref}
}

func (tr TargetRegionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tr.ref.InternalTokens()
}

func (tr TargetRegionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("name"))
}

func (tr TargetRegionAttributes) RegionalReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(tr.ref.Append("regional_replica_count"))
}

func (tr TargetRegionAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("storage_account_type"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type TargetRegionState struct {
	Name                 string  `json:"name"`
	RegionalReplicaCount float64 `json:"regional_replica_count"`
	StorageAccountType   string  `json:"storage_account_type"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}