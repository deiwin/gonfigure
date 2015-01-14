# GO-N-FIGURE

[![GoDoc](https://godoc.org/github.com/deiwin/gonfigure?status.svg)](https://godoc.org/github.com/deiwin/gonfigure)

Minimalistic configuration helper for your Go projects.

## Example

```go
var portProperty   = gonfigure.NewEnvProperty("PORT", "8080")
// If the $DOMAIN env variable is not set the configuration creation will fail with a fatal error
var domainProperty = gonfigure.NewRequiredEnvProperty("DOMAIN")

type Config struct {
  Port   string
  Domain string
}

func NewConfig() Config {
  return Config{
    Port:   portProperty.Value(),
    Domain: domainProperty.Value(),
  }
}
```
