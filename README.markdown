# Date Manipulator

Simple application for generating [RFC3339 timestamps](https://datatracker.ietf.org/doc/html/rfc3339), written in
[Go](https://golang.org/).

Install the program like this:
```shell
go get -v github.com/mojotx/date-manipulator
```

To get the current UTC timestamp in RFC3339 format, just run the binary:
```shell
$ date-manipulator
2021-06-02T20:56:32Z
```

To get the UTC timestamp of 15 minutes ago, specify a valid Go [time.Duration](https://golang.org/pkg/time/#ParseDuration)
string, like this:
```shell
$ date-manipulator -15m
2021-06-02T20:43:38Z
```

To see 15 minutes ago, as well as current time, assuming [Bash](https://www.gnu.org/software/bash/) as your shell:
```shell
printf "Past %22s\n Now %22s\n" $(date-manipulator -15m) $(date-manipulator)
Past   2021-06-02T20:47:19Z
 Now   2021-06-02T21:02:19Z
```

Note that there are multiple [time.Duration](https://golang.org/pkg/time/#ParseDuration) units suffixes, including
"ns", "us" (or "µs"), "ms", "s", "m", and "h".
```shell
$ date-manipulator ; date-manipulator 12h
2021-06-02T21:05:45Z
2021-06-03T09:05:45Z
```

This is licensed under the [MIT license](https://opensource.org/licenses/MIT). If you find it useful, great!
