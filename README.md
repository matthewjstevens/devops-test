# DevOps Test

Simple Go app with single HTTP path: `/hello` - which returns the contents of the `MESSAGE` env var.

Built using Golang 1.23. To run locally:

```
go build
MESSAGE="Hello World" ./devopsTest
```