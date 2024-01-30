// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     public/app/plugins/gen.go
// Using jennies:
//     PluginGoTypesJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package dataquery

// Defines values for SearchStreamingState.
const (
	SearchStreamingStateDone      SearchStreamingState = "done"
	SearchStreamingStateError     SearchStreamingState = "error"
	SearchStreamingStatePending   SearchStreamingState = "pending"
	SearchStreamingStateStreaming SearchStreamingState = "streaming"
)

// Defines values for SearchTableType.
const (
	SearchTableTypeSpans  SearchTableType = "spans"
	SearchTableTypeTraces SearchTableType = "traces"
)

// Defines values for TempoQueryType.
const (
	TempoQueryTypeClear         TempoQueryType = "clear"
	TempoQueryTypeNativeSearch  TempoQueryType = "nativeSearch"
	TempoQueryTypeSearch        TempoQueryType = "search"
	TempoQueryTypeServiceMap    TempoQueryType = "serviceMap"
	TempoQueryTypeTraceId       TempoQueryType = "traceId"
	TempoQueryTypeTraceql       TempoQueryType = "traceql"
	TempoQueryTypeTraceqlSearch TempoQueryType = "traceqlSearch"
	TempoQueryTypeUpload        TempoQueryType = "upload"
)

// Defines values for TraceqlSearchScope.
const (
	TraceqlSearchScopeIntrinsic TraceqlSearchScope = "intrinsic"
	TraceqlSearchScopeResource  TraceqlSearchScope = "resource"
	TraceqlSearchScopeSpan      TraceqlSearchScope = "span"
	TraceqlSearchScopeUnscoped  TraceqlSearchScope = "unscoped"
)

// These are the common properties available to all queries in all datasources.
// Specific implementations will *extend* this interface, adding the required
// properties for the given context.
type DataQuery struct {
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *any `json:"datasource,omitempty"`

	// Hide true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
}

// The state of the TraceQL streaming search query
type SearchStreamingState string

// The type of the table that is used to display the search results
type SearchTableType string

// TempoDataQuery defines model for TempoDataQuery.
type TempoDataQuery = map[string]any

// TempoQuery defines model for TempoQuery.
type TempoQuery struct {
	// DataQuery These are the common properties available to all queries in all datasources.
	// Specific implementations will *extend* this interface, adding the required
	// properties for the given context.
	DataQuery

	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *any            `json:"datasource,omitempty"`
	Filters    []TraceqlFilter `json:"filters"`

	// Filters that are used to query the metrics summary
	GroupBy []TraceqlFilter `json:"groupBy,omitempty"`

	// Hide true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`

	// Defines the maximum number of traces that are returned from Tempo
	Limit *int64 `json:"limit,omitempty"`

	// @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
	MaxDuration *string `json:"maxDuration,omitempty"`

	// @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
	MinDuration *string `json:"minDuration,omitempty"`

	// TraceQL query or trace ID
	Query *string `json:"query,omitempty"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`

	// @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
	Search *string `json:"search,omitempty"`

	// Use service.namespace in addition to service.name to uniquely identify a service.
	ServiceMapIncludeNamespace *bool `json:"serviceMapIncludeNamespace,omitempty"`

	// Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
	ServiceMapQuery *any `json:"serviceMapQuery,omitempty"`

	// @deprecated Query traces by service name
	ServiceName *string `json:"serviceName,omitempty"`

	// @deprecated Query traces by span name
	SpanName *string `json:"spanName,omitempty"`

	// Defines the maximum number of spans per spanset that are returned from Tempo
	Spss *int64 `json:"spss,omitempty"`

	// The type of the table that is used to display the search results
	TableType *SearchTableType `json:"tableType,omitempty"`
}

// TempoQueryType search = Loki search, nativeSearch = Tempo search for backwards compatibility
type TempoQueryType string

// TraceqlFilter defines model for TraceqlFilter.
type TraceqlFilter struct {
	// Uniquely identify the filter, will not be used in the query generation
	Id string `json:"id"`

	// The operator that connects the tag to the value, for example: =, >, !=, =~
	Operator *string `json:"operator,omitempty"`

	// Scope static fields are pre-set in the UI, dynamic fields are added by the user
	Scope *TraceqlSearchScope `json:"scope,omitempty"`

	// The tag for the search filter, for example: .http.status_code, .service.name, status
	Tag *string `json:"tag,omitempty"`

	// The value for the search filter
	Value *any `json:"value,omitempty"`

	// The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query
	ValueType *string `json:"valueType,omitempty"`
}

// TraceqlSearchScope static fields are pre-set in the UI, dynamic fields are added by the user
type TraceqlSearchScope string