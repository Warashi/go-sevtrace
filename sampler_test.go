// sampler_test.go
package sevtrace_test

import (
	"testing"

	"github.com/Warashi/go-sevtrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/attribute"
)

// TestSeveritySampler_DefaultSeverity verifies that when no severity attribute is provided,
// the sampler defaults to Info (sevtrace.SeverityInfo) and drops the span if the threshold is higher.
func TestSeveritySampler_DefaultSeverity(t *testing.T) {
	// Create a sampler with a threshold of SeverityWarn.
	sampler := sevtrace.NewSeveritySampler(sevtrace.SeverityWarn, sdktrace.AlwaysSample())
	params := sdktrace.SamplingParameters{
		Attributes: []attribute.KeyValue{}, // no severity attribute provided
	}
	result := sampler.ShouldSample(params)
	if result.Decision != sdktrace.Drop {
		t.Errorf("Expected decision to be Drop for default Info severity, got %v", result.Decision)
	}
}

// TestSeveritySampler_LowSeverity verifies that a span with a severity below the threshold is dropped.
func TestSeveritySampler_LowSeverity(t *testing.T) {
	sampler := sevtrace.NewSeveritySampler(sevtrace.SeverityWarn, sdktrace.AlwaysSample())
	// Using Debug severity (5), which is below the threshold (13).
	attr := sevtrace.SeverityKey.Int64(sevtrace.SeverityDebug)
	params := sdktrace.SamplingParameters{
		Attributes: []attribute.KeyValue{attr},
	}
	result := sampler.ShouldSample(params)
	if result.Decision != sdktrace.Drop {
		t.Errorf("Expected decision to be Drop for severity Debug, got %v", result.Decision)
	}
}

// TestSeveritySampler_EqualSeverity verifies that a span with severity equal to the threshold
// delegates the decision to the inner sampler.
func TestSeveritySampler_EqualSeverity(t *testing.T) {
	sampler := sevtrace.NewSeveritySampler(sevtrace.SeverityWarn, sdktrace.AlwaysSample())
	// Using Warn severity (13), which is equal to the threshold.
	attr := sevtrace.SeverityKey.Int64(sevtrace.SeverityWarn)
	params := sdktrace.SamplingParameters{
		Attributes: []attribute.KeyValue{attr},
	}
	result := sampler.ShouldSample(params)
	if result.Decision != sdktrace.RecordAndSample {
		t.Errorf("Expected decision to be RecordAndSample for severity Warn, got %v", result.Decision)
	}
}

// TestSeveritySampler_HighSeverity verifies that a span with severity above the threshold
// delegates the decision to the inner sampler.
func TestSeveritySampler_HighSeverity(t *testing.T) {
	sampler := sevtrace.NewSeveritySampler(sevtrace.SeverityWarn, sdktrace.AlwaysSample())
	// Using Error severity (17), which is above the threshold.
	attr := sevtrace.SeverityKey.Int64(sevtrace.SeverityError)
	params := sdktrace.SamplingParameters{
		Attributes: []attribute.KeyValue{attr},
	}
	result := sampler.ShouldSample(params)
	if result.Decision != sdktrace.RecordAndSample {
		t.Errorf("Expected decision to be RecordAndSample for severity Error, got %v", result.Decision)
	}
}

// TestAttributeHelpers verifies that the helper functions generate the correct severity attributes.
func TestAttributeHelpers(t *testing.T) {
	// Test WithTrace
	attr := sevtrace.WithTrace()
	if attr.Key != sevtrace.SeverityKey {
		t.Errorf("WithTrace: expected key %v, got %v", sevtrace.SeverityKey, attr.Key)
	}
	if attr.Value.AsInt64() != sevtrace.SeverityTrace {
		t.Errorf("WithTrace: expected value %d, got %d", sevtrace.SeverityTrace, attr.Value.AsInt64())
	}
	// Test WithDebug
	attr = sevtrace.WithDebug()
	if attr.Key != sevtrace.SeverityKey {
		t.Errorf("WithDebug: expected key %v, got %v", sevtrace.SeverityKey, attr.Key)
	}
	if attr.Value.AsInt64() != sevtrace.SeverityDebug {
		t.Errorf("WithDebug: expected value %d, got %d", sevtrace.SeverityDebug, attr.Value.AsInt64())
	}
	// Test WithInfo
	attr = sevtrace.WithInfo()
	if attr.Key != sevtrace.SeverityKey {
		t.Errorf("WithInfo: expected key %v, got %v", sevtrace.SeverityKey, attr.Key)
	}
	if attr.Value.AsInt64() != sevtrace.SeverityInfo {
		t.Errorf("WithInfo: expected value %d, got %d", sevtrace.SeverityInfo, attr.Value.AsInt64())
	}
	// Test WithWarn
	attr = sevtrace.WithWarn()
	if attr.Key != sevtrace.SeverityKey {
		t.Errorf("WithWarn: expected key %v, got %v", sevtrace.SeverityKey, attr.Key)
	}
	if attr.Value.AsInt64() != sevtrace.SeverityWarn {
		t.Errorf("WithWarn: expected value %d, got %d", sevtrace.SeverityWarn, attr.Value.AsInt64())
	}
	// Test WithError
	attr = sevtrace.WithError()
	if attr.Key != sevtrace.SeverityKey {
		t.Errorf("WithError: expected key %v, got %v", sevtrace.SeverityKey, attr.Key)
	}
	if attr.Value.AsInt64() != sevtrace.SeverityError {
		t.Errorf("WithError: expected value %d, got %d", sevtrace.SeverityError, attr.Value.AsInt64())
	}
	// Test WithFatal
	attr = sevtrace.WithFatal()
	if attr.Key != sevtrace.SeverityKey {
		t.Errorf("WithFatal: expected key %v, got %v", sevtrace.SeverityKey, attr.Key)
	}
	if attr.Value.AsInt64() != sevtrace.SeverityFatal {
		t.Errorf("WithFatal: expected value %d, got %d", sevtrace.SeverityFatal, attr.Value.AsInt64())
	}
}
