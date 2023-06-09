// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package elasticachecluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CacheNodes struct{}

type LogDeliveryConfiguration struct {
	// Destination: string, required
	Destination terra.StringValue `hcl:"destination,attr" validate:"required"`
	// DestinationType: string, required
	DestinationType terra.StringValue `hcl:"destination_type,attr" validate:"required"`
	// LogFormat: string, required
	LogFormat terra.StringValue `hcl:"log_format,attr" validate:"required"`
	// LogType: string, required
	LogType terra.StringValue `hcl:"log_type,attr" validate:"required"`
}

type CacheNodesAttributes struct {
	ref terra.Reference
}

func (cn CacheNodesAttributes) InternalRef() (terra.Reference, error) {
	return cn.ref, nil
}

func (cn CacheNodesAttributes) InternalWithRef(ref terra.Reference) CacheNodesAttributes {
	return CacheNodesAttributes{ref: ref}
}

func (cn CacheNodesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cn.ref.InternalTokens()
}

func (cn CacheNodesAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("address"))
}

func (cn CacheNodesAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("availability_zone"))
}

func (cn CacheNodesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("id"))
}

func (cn CacheNodesAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("outpost_arn"))
}

func (cn CacheNodesAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(cn.ref.Append("port"))
}

type LogDeliveryConfigurationAttributes struct {
	ref terra.Reference
}

func (ldc LogDeliveryConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ldc.ref, nil
}

func (ldc LogDeliveryConfigurationAttributes) InternalWithRef(ref terra.Reference) LogDeliveryConfigurationAttributes {
	return LogDeliveryConfigurationAttributes{ref: ref}
}

func (ldc LogDeliveryConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ldc.ref.InternalTokens()
}

func (ldc LogDeliveryConfigurationAttributes) Destination() terra.StringValue {
	return terra.ReferenceAsString(ldc.ref.Append("destination"))
}

func (ldc LogDeliveryConfigurationAttributes) DestinationType() terra.StringValue {
	return terra.ReferenceAsString(ldc.ref.Append("destination_type"))
}

func (ldc LogDeliveryConfigurationAttributes) LogFormat() terra.StringValue {
	return terra.ReferenceAsString(ldc.ref.Append("log_format"))
}

func (ldc LogDeliveryConfigurationAttributes) LogType() terra.StringValue {
	return terra.ReferenceAsString(ldc.ref.Append("log_type"))
}

type CacheNodesState struct {
	Address          string  `json:"address"`
	AvailabilityZone string  `json:"availability_zone"`
	Id               string  `json:"id"`
	OutpostArn       string  `json:"outpost_arn"`
	Port             float64 `json:"port"`
}

type LogDeliveryConfigurationState struct {
	Destination     string `json:"destination"`
	DestinationType string `json:"destination_type"`
	LogFormat       string `json:"log_format"`
	LogType         string `json:"log_type"`
}
