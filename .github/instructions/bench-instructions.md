# Benchmark Instructions

These instructions define standards for writing and maintaining benchmarks.

## Requirements
- Document the purpose of each benchmark.

## Best Practices
- Benchmark both typical and edge-case scenarios.
- Cover performance aspects including:
  - Startup and shutdown time
  - Scalability (e.g., increasing load, resource usage)
  - Throughput (operations per second, bandwidth)
  - Latency (response time, tail latency)
  - Recovery time (from failure or restart)
- Document any non-determinism or variability in results.

---

**For all Go-specific requirements and best practices (e.g., idioms, assertions, error handling), see `go.instructions.md`.**

---

**Use this file to guide all benchmark code generation and review.**
