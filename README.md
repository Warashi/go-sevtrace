# go-sevtrace

go-sevtrace is a custom OpenTelemetry Go SDK package that enables severity-based trace sampling by annotating spans with OTEL Logs severity levels and delegating the sampling decision to a user-defined inner sampler.

## Features

- **Severity Annotation:** Use severity levels defined in the OpenTelemetry Logs specification.
- **Custom Sampling:** Defaults to Info severity when no severity attribute is provided.
- **Inner Sampler Delegation:** Apply additional sampling strategies (e.g., random sampling) via a configurable inner sampler.
- **Helper Function:** Easily attach severity attributes when starting spans.

## Installation

Install the package using Go modules:

```bash
go get github.com/Warashi/go-sevtrace
```

```go
package main

import (
	"context"
	"github.com/Warashi/go-sevtrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() {
	// Set up an inner sampler (e.g., random sampling at 10%).
	innerSampler := sdktrace.TraceIDRatioBased(0.1)
	// Create a SeveritySampler that delegates if severity is at least SeverityWarn.
	sampler := sevtrace.NewSeveritySampler(sevtrace.SeverityWarn, innerSampler)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sampler),
		sdktrace.WithResource(resource.NewWithAttributes(
			attribute.String("service.name", "my-service"),
		)),
	)
	otel.SetTracerProvider(tp)
}

func main() {
	ctx := context.Background()
	tracer := otel.Tracer("example")

	// Start a span with a severity attribute using the helper function.
	ctx, span := tracer.Start(ctx, "sample-operation",
		otel.WithAttributes(sevtrace.WithSeverity(sevtrace.SeverityWarn)),
	)
	defer span.End()

	// Your application logic here.
}
```

## License
This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

