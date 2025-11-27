# glog

Logging, tracing, and monitoring in Go.

- [uber-zap](#zaplog)
- [zerolog](#zerolog)

## zerolog

- [Zerolog GitHub](https://github.com/rs/zerolog)

According to benchmark tests conducted by the creators of `zap` and others, `zerolog` demonstrates better performance than `uber-go/zap`.
This performance improvement is achieved by avoiding memory allocations and reflections, similar to `zap`, but with additional optimizations.

While `zerolog` is minimal and efficient, it lacks some flexibility and introduces verbosity to achieve features beyond standard logging expectations.
For example, features like stack traces and limited data type support for fields require specific handling. Additionally, its chaining API style and field type definitions can make it feel less intuitive.

## zaplog

- [uber-zap](https://github.com/uber-go/zap/tree/master)

**Zap Logger**: Structured, leveled, and fast logging framework for Go.

## Key features

**Separate Environments**:
- In development, include console output and a lower log level, mostly above debug.
- In production, use file output with a higher log level, typically error and above.

**Error Stack Traces**:
- Use `AddStacktrace` in the configuration to specify the log level for stack trace generation.

**Encoder Mode**:
- **Console Encoder** + **Debug level**: Good for development (human-readable output).
- **JSON Encoder** + **Error level**: Commonly used in production for structured logs.

**Sample Rate(optional)**:
- You can adjust the sample rate to control log verbosity, especially in production.
- This is useful for reducing log noise on frequent events (e.g., setting a rate limiter on info-level logs).

**Output Destinations**:
- Zap allows you to configure multiple output destinations.
- You can log to standard output, files, or external log aggregation services.
- Example:
   ```go
   zapConfig.OutputPaths = []string{"stdout", "/var/log/app.log"}
   ```

**Structured Logging**:
- Zap encourages structured logging by using key-value pairs; also child logger(clone) is possible.
- This can be done with `With` for context or directly in `Info`, `Error`, etc.:
   ```go
   logger.With(zap.String("userID", "123")).Info("User login successful")
   ```