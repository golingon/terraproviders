// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dbproxydefaulttargetgroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ConnectionPoolConfig struct {
	// ConnectionBorrowTimeout: number, optional
	ConnectionBorrowTimeout terra.NumberValue `hcl:"connection_borrow_timeout,attr"`
	// InitQuery: string, optional
	InitQuery terra.StringValue `hcl:"init_query,attr"`
	// MaxConnectionsPercent: number, optional
	MaxConnectionsPercent terra.NumberValue `hcl:"max_connections_percent,attr"`
	// MaxIdleConnectionsPercent: number, optional
	MaxIdleConnectionsPercent terra.NumberValue `hcl:"max_idle_connections_percent,attr"`
	// SessionPinningFilters: set of string, optional
	SessionPinningFilters terra.SetValue[terra.StringValue] `hcl:"session_pinning_filters,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ConnectionPoolConfigAttributes struct {
	ref terra.Reference
}

func (cpc ConnectionPoolConfigAttributes) InternalRef() (terra.Reference, error) {
	return cpc.ref, nil
}

func (cpc ConnectionPoolConfigAttributes) InternalWithRef(ref terra.Reference) ConnectionPoolConfigAttributes {
	return ConnectionPoolConfigAttributes{ref: ref}
}

func (cpc ConnectionPoolConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cpc.ref.InternalTokens()
}

func (cpc ConnectionPoolConfigAttributes) ConnectionBorrowTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(cpc.ref.Append("connection_borrow_timeout"))
}

func (cpc ConnectionPoolConfigAttributes) InitQuery() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("init_query"))
}

func (cpc ConnectionPoolConfigAttributes) MaxConnectionsPercent() terra.NumberValue {
	return terra.ReferenceAsNumber(cpc.ref.Append("max_connections_percent"))
}

func (cpc ConnectionPoolConfigAttributes) MaxIdleConnectionsPercent() terra.NumberValue {
	return terra.ReferenceAsNumber(cpc.ref.Append("max_idle_connections_percent"))
}

func (cpc ConnectionPoolConfigAttributes) SessionPinningFilters() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cpc.ref.Append("session_pinning_filters"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ConnectionPoolConfigState struct {
	ConnectionBorrowTimeout   float64  `json:"connection_borrow_timeout"`
	InitQuery                 string   `json:"init_query"`
	MaxConnectionsPercent     float64  `json:"max_connections_percent"`
	MaxIdleConnectionsPercent float64  `json:"max_idle_connections_percent"`
	SessionPinningFilters     []string `json:"session_pinning_filters"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Update string `json:"update"`
}