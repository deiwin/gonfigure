# GO-N-FIGURE

[![GoDoc](https://godoc.org/github.com/deiwin/gonfigure?status.svg)](https://godoc.org/github.com/deiwin/gonfigure)

Minimalistic configuration helper for your Go projects.

## Example

```go
var portProperty   = gonfigure.NewEnvProperty("PORT", "8080")
var domainProperty = gonfigure.NewRequiredEnvProperty("DOMAIN")

type Config struct {
  Port   string
  Domain string
}

func NewConfig() Config {
  return Config{
    Port:   portProperty.Value(),
    // If the $DOMAIN env variable is not set, this call will panic
    Domain: domainProperty.Value(),
  }
}
```
