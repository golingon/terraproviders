// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package flowlog

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DestinationOptions struct {
	// FileFormat: string, optional
	FileFormat terra.StringValue `hcl:"file_format,attr"`
	// HiveCompatiblePartitions: bool, optional
	HiveCompatiblePartitions terra.BoolValue `hcl:"hive_compatible_partitions,attr"`
	// PerHourPartition: bool, optional
	PerHourPartition terra.BoolValue `hcl:"per_hour_partition,attr"`
}

type DestinationOptionsAttributes struct {
	ref terra.Reference
}

func (do DestinationOptionsAttributes) InternalRef() terra.Reference {
	return do.ref
}

func (do DestinationOptionsAttributes) InternalWithRef(ref terra.Reference) DestinationOptionsAttributes {
	return DestinationOptionsAttributes{ref: ref}
}

func (do DestinationOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return do.ref.InternalTokens()
}

func (do DestinationOptionsAttributes) FileFormat() terra.StringValue {
	return terra.ReferenceAsString(do.ref.Append("file_format"))
}

func (do DestinationOptionsAttributes) HiveCompatiblePartitions() terra.BoolValue {
	return terra.ReferenceAsBool(do.ref.Append("hive_compatible_partitions"))
}

func (do DestinationOptionsAttributes) PerHourPartition() terra.BoolValue {
	return terra.ReferenceAsBool(do.ref.Append("per_hour_partition"))
}

type DestinationOptionsState struct {
	FileFormat               string `json:"file_format"`
	HiveCompatiblePartitions bool   `json:"hive_compatible_partitions"`
	PerHourPartition         bool   `json:"per_hour_partition"`
}
