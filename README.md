# go-sevtrace

go-sevtrace is a custom OpenTelemetry Go SDK package that provides severity-based trace sampling by annotating spans with severity levels defined in the OpenTelemetry Logs specification. It implements a custom sampler that uses a configurable severity threshold and delegates the final sampling decision to a user-defined inner sampler. In addition, the package offers helper functions to easily generate severity attributes.

## Features

- **Severity-based Sampling:** Annotate spans with severity levels.
- **Customizable Threshold:** Defaults to Info severity if no attribute is provided.
- **Inner Sampler Delegation:** Further sample spans using a user-defined inner sampler (e.g., random sampling).
- **Helper Functions:** Convenient helper functions for generating severity attributes:
  - `WithTrace()`
  - `WithDebug()`
  - `WithInfo()`
  - `WithWarn()`
  - `WithError()`
  - `WithFatal()`

## Installation

Install the package using Go modules:

```bash
go get github.com/Warashi/go-sevtrace
```

## Usage
### Setting Up the Sampler
Configure your tracer provider to use the severity-based sampler. In the example below, spans with severity Warn or above are further evaluated by an inner sampler that performs random sampling at a 10% rate.

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
	innerSampler := sdktrace.TraceIDRatioBased(0.1)
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
		otel.WithAttributes(sevtrace.WithWarn()),
	)
	defer span.End()

	// Your application logic here.
}
```

### Using Severity Helpers
The package provides helper functions to easily attach severity attributes when starting a span. For example:

```go
// Starting a span with Info severity
ctx, span := tracer.Start(ctx, "operation",
	otel.WithAttributes(sevtrace.WithInfo()),
)
defer span.End()
```

## License
This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
