// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datamonitoringappengineservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Telemetry struct{}

type TelemetryAttributes struct {
	ref terra.Reference
}

func (t TelemetryAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TelemetryAttributes) InternalWithRef(ref terra.Reference) TelemetryAttributes {
	return TelemetryAttributes{ref: ref}
}

func (t TelemetryAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TelemetryAttributes) ResourceName() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("resource_name"))
}

type TelemetryState struct {
	ResourceName string `json:"resource_name"`
}
