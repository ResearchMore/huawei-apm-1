package common

// const for zipkin
const (
	APMKwySeparator = "|"

	// Endpoint to send Zipkin spans to.
	apmHTTPEndpoint = ""

	// Debug mode.
	zipkinDebug = false

	// same span
	sameSpan = true

	// make Tracer generate 128 bit traceID's for root spans.
	traceID128Bit = true
)

const (
	INTERMEDIATE = iota + 1
	FIRST_FOR_BACKEND
	FIRST_FOR_CLIENT
	FIRST_FOR_UNKNOWN
	FIRST_FOR_ENDPOINT
	ENDPOINT
	EXTERNAL
	FIRST_EXTERNAL
	ENDPOINT_EXTERNAL
	ISTIO
	ISTIO_EXTERNAL
)
