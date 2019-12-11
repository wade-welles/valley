# Valley

Valley is tool for generating plain Go validation code based on your Go code.

## Usage

...

## Built-In Constraints

* Equals
* Max
* MaxLength
* Min
* MinLength
* MutuallyExclusive
* NotEquals
* NotNil
* Required
* Valid

## Motivation

Previously I've implemented validation in Go using reflection, and while reflection isn't actually
as slow as you might expect it does come with other issues. By generating validation code instead of
resorting to reflection you regain the protection that the Go compiler gives you. Even if the output
of Valley is wrong (or you misconfigure the constraints) your application would fail to compile,
alerting you to the issue.

On the topic of performance, the code generated by Valley is still a lot faster than reflection.
This is for several reasons. One is that I've tried to be quite efficient doing things like building
up a path to fields (i.e. reusing memory where possible, and not adding to the path unless a
constraint violation occurs, or there's no choice not to). Another is that without using reflection,
the checks just become simple `if` statements and loops - these checks are extremely fast.

Another issue I found with reflection-based approaches is that you have to pass in references to
fields to validate as strings (i.e. the name of the field), rather than the fields themselves. This
is because you can't retrieve a field name as far as I can tell from a value passed in using
reflection. The configuration for Valley needs to be able to compile as Go code. If it's mis-used,
Valley will do it's best to tell you what's wrong, and where. References to fields should exist, and
your existing tooling, and Go toolchain will tell you if they don't - as well as Valley. On top of
that, the generated code also has to compile, further protecting you from runtime panics.

## TODO

* Add some benchmarks to the README?
* The ability to attach multiple constraints methods to a type, that generate different validate
functions (the `Valid` constraint would need an option to override which method is called).
* You might want to validate map keys too, so maybe a `Keys` method on `Field`?
* Allow overriding field names by using struct tags?
* The ability to define constraints in a separate file (in the same package).

## License

MIT

## Contributions

Feel free to open a [pull request][1], or file an [issue][2] on Github. I always welcome
contributions as long as they're for the benefit of all (potential) users of this project.

If you're unsure about anything, feel free to ask about it in an issue before you get your heart set
on fixing it yourself.

[1]: https://github.com/seeruk/valley/pulls
[2]: https://github.com/seeruk/valley/issues
