package sevtrace

import (
	"go.opentelemetry.io/otel/attribute"
)

// WithSeverity returns an attribute.KeyValue for the severity level.
// This helper can be used when starting a span, for example:
// tracer.Start(ctx, "operation", otel.WithAttributes(WithSeverity(sevtrace.SeverityWarn))).
func WithSeverity(severity int64) attribute.KeyValue {
	return severityKey.Int64(severity)
}
