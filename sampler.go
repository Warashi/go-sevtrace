package sevtrace

import (
	"strconv"

	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// SeverityKey is the attribute key used to attach the severity level to a span.
const severityKey = attribute.Key("severity")

// SeveritySampler is a custom sampler that checks the "severity" attribute
// on a span. If the attribute is not found, it defaults to Info (sevtrace.SeverityInfo).
// If the severity is greater than or equal to the specified threshold, the decision
// is delegated to the innerSampler.
type SeveritySampler struct {
	threshold    int64
	innerSampler sdktrace.Sampler
}

// NewSeveritySampler creates a new SeveritySampler.
//   - threshold: The severity threshold (e.g., sevtrace.SeverityWarn). If a span's severity
//     is greater than or equal to this threshold, the innerSampler is used.
//   - innerSampler: The sampling strategy to apply when the severity condition is met.
//     If nil, it defaults to AlwaysSample.
func NewSeveritySampler(threshold int64, innerSampler sdktrace.Sampler) sdktrace.Sampler {
	if innerSampler == nil {
		innerSampler = sdktrace.AlwaysSample()
	}
	return &SeveritySampler{
		threshold:    threshold,
		innerSampler: innerSampler,
	}
}

// ShouldSample implements the sdktrace.Sampler interface.
// It searches the span's attributes for the "severity" key and defaults to SeverityInfo (9)
// if not found. If the severity value is greater than or equal to the threshold,
// the decision is delegated to the innerSampler; otherwise, the span is dropped.
func (s *SeveritySampler) ShouldSample(params sdktrace.SamplingParameters) sdktrace.SamplingResult {
	// If the severity attribute is not found, default to Info (SeverityInfo = 9)
	sev := SeverityInfo
	for _, attr := range params.Attributes {
		if attr.Key == severityKey {
			if attr.Value.Type() == attribute.INT64 {
				sev = attr.Value.AsInt64()
			}
			break
		}
	}
	if sev >= s.threshold {
		return s.innerSampler.ShouldSample(params)
	}
	return sdktrace.SamplingResult{Decision: sdktrace.Drop}
}

// Description returns a description of the SeveritySampler.
func (s *SeveritySampler) Description() string {
	return "SeveritySampler(threshold=" + strconv.FormatInt(s.threshold, 10) +
		", inner=" + s.innerSampler.Description() + ")"
}
