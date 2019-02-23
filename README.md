## An implementation of Common check-digit mod11 algorithm

[![Go Report Card](https://goreportcard.com/badge/github.com/kolkov/esr)](https://goreportcard.com/report/github.com/kolkov/esr)

Compatible with the ozzo-validation package.

### Usage ###

```
import "github.com/kolkov/esr"

err := esr.Validate("639400")

signed := esr.Generate("1")
```
