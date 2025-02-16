package sevtrace

// Severity constants based on the OpenTelemetry Logs specification.
// Constants with numeric suffixes are excluded.
const (
	SeverityUndefined int64 = 0

	// Trace severity.
	SeverityTrace int64 = 1

	// Debug severity.
	SeverityDebug int64 = 5

	// Info severity.
	SeverityInfo int64 = 9

	// Warn severity.
	SeverityWarn int64 = 13

	// Error severity.
	SeverityError int64 = 17

	// Fatal severity.
	SeverityFatal int64 = 21
)
