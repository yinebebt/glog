# zerolog

- [Zerolog GitHub](https://github.com/rs/zerolog)

According to benchmark tests conducted by the creators of `zap` and others, `zerolog` demonstrates better performance than `uber-go/zap`. 
This performance improvement is achieved by avoiding memory allocations and reflections, similar to `zap`, but with additional optimizations.

From the README, I found that while `zerolog` is minimal and efficient, it lacks some flexibility and introduces verbosity to achieve features beyond standard logging expectations. 
For example, features like stack traces and limited data type support for fields require specific handling. Additionally, its chaining API style and field type definitions can make it feel less intuitive.