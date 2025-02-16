package sevtrace

import (
	"go.opentelemetry.io/otel/attribute"
)

// WithSeverity returns an attribute.KeyValue for the given severity level.
func WithSeverity(severity int64) attribute.KeyValue {
	return severityKey.Int64(severity)
}

// WithTrace returns an attribute.KeyValue for the Trace severity.
func WithTrace() attribute.KeyValue {
	return WithSeverity(SeverityTrace)
}

// WithDebug returns an attribute.KeyValue for the Debug severity.
func WithDebug() attribute.KeyValue {
	return WithSeverity(SeverityDebug)
}

// WithInfo returns an attribute.KeyValue for the Info severity.
func WithInfo() attribute.KeyValue {
	return WithSeverity(SeverityInfo)
}

// WithWarn returns an attribute.KeyValue for the Warn severity.
func WithWarn() attribute.KeyValue {
	return WithSeverity(SeverityWarn)
}

// WithError returns an attribute.KeyValue for the Error severity.
func WithError() attribute.KeyValue {
	return WithSeverity(SeverityError)
}

// WithFatal returns an attribute.KeyValue for the Fatal severity.
func WithFatal() attribute.KeyValue {
	return WithSeverity(SeverityFatal)
}

