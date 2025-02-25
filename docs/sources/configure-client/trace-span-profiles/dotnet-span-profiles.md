---
title: Span profiles with Traces to profiles for .NET
menuTitle: Span profiles with Traces to profiles (.NET)
description: Learn about and configure Span profiles with Traces to profiles in Grafana for .NET applications.
weight: 103
---

# Span profiles with Traces to profiles for .NET

Span Profiles represents a major shift in profiling methodology, enabling deeper analysis of both tracing and profiling data.
Traditional continuous profiling provides an application-wide view over fixed intervals.
In contrast, Span Profiles delivers focused, dynamic analysis on specific execution scopes within applications, such as individual requests or specific trace spans.

This shift enables a more granular view of performance, enhancing the utility of profiles by linking them directly with traces for a comprehensive understanding of application behavior. As a result, engineering teams can more efficiently identify and address performance bottlenecks.

To learn more about Span Profiles, refer to [Combining tracing and profiling for enhanced observability: Introducing Span Profiles](/blog/2024/02/06/combining-tracing-and-profiling-for-enhanced-observability-introducing-span-profiles/).

![span-profiles screenshot](https://grafana.com/static/img/docs/tempo/profiles/tempo-profiles-Span-link-profile-data-source.png)

Pyroscope integrates with distributed tracing systems supporting the [**OpenTelemetry**](https://opentelemetry.io/docs/languages/net/getting-started/) standard.
This integration lets you link traces with the profiling data and find resource usage for specific lines of code for your trace spans.

{{< admonition type="note" >}}
* Only CPU profiling is supported at the moment.
* Because of how sampling profilers work, spans shorter than the sample interval may not be captured.
{{< /admonition >}}

To use Span Profiles, you need to:

* [Configure Pyroscope to send profiling data]({{< relref "../../configure-client" >}})
* Configure a client-side package to link traces and profiles: [.NET](https://github.com/grafana/pyroscope-dotnet/tree/main/Pyroscope/Pyroscope.OpenTelemetry)
* [Configure the Tempo data source in Grafana or Grafana Cloud to discover linked traces and profiles](/grafana-cloud/connect-externally-hosted/data-sources/tempo/configure-tempo-data-source/)

## Before you begin

Your applications must be instrumented for profiling and tracing before you can use span profiles.

* Profiling: Your application must be instrumented with Pyroscope's .NET instrumentation library. Refer to the [.NET]({{< relref "../language-sdks/dotnet" >}}) guide for instructions.
* Tracing: Your application must be instrumented with OpenTelemetry traces. Refer to the [OpenTelemetry](https://opentelemetry.io/docs/languages/net/getting-started/) guide for isntructions.

{{< admonition type="note" >}}
Span profiles in .NET are only supported using [OpenTelemetry manual instrumentation](https://opentelemetry.io/docs/languages/net/instrumentation/)
because Pyroscope's .NET profiler and OpenTelemetry's auto instrumentation are based on separate .NET CLR profilers.
{{< /admonition >}}

## Configure the `Pyroscope.OpenTelemetry` package

To start collecting Span Profiles for your .NET application, you need to include [Pyroscope.OpenTelemetry](https://github.com/grafana/pyroscope-dotnet/tree/main/Pyroscope/Pyroscope.OpenTelemetry) in your code.

This package provides a [`SpanProcessor`](https://github.com/open-telemetry/opentelemetry-dotnet/blob/main/src/OpenTelemetry/BaseProcessor.cs) implementation, which connects the two telemetry signals (traces and profiles) together.

```shell
dotnet add package Pyroscope.OpenTelemetry
```

Next, create and register the `PyroscopeSpanProcessor`:
```csharp
builder.Services.AddOpenTelemetry()
    .WithTracing(b =>
    {
        b
        .AddAspNetCoreInstrumentation()
        .AddConsoleExporter()
        .AddOtlpExporter()
        .AddProcessor(new Pyroscope.OpenTelemetry.PyroscopeSpanProcessor());
    });
```

With the span processor registered, spans created automatically (for example, HTTP handlers) and manually (`ActivitySource.StartActivity()`) have profiling data associated with them.

## View the span profiles in Grafana Tempo

To view the span profiles in Grafana Tempo, you need to have a Grafana instance running and a data source configured to link traces and profiles.

Refer to the [data source configuration documentation](/docs/grafana/datasources/tempo/configure-tempo-data-source) to see how to configure the visualization to link traces with profiles.

## Examples

Check out the [examples](https://github.com/grafana/pyroscope/tree/main/examples/tracing/tempo) directory for a complete demo application of span profiles in multiple languages.
