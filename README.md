## An implementation of Common check-digit mod11 algorithm

[![Go Report Card](https://goreportcard.com/badge/github.com/kolkov/esr)](https://goreportcard.com/report/github.com/kolkov/esr)
[![Coverage Status](https://coveralls.io/repos/github/kolkov/esr/badge.svg?branch=master)](https://coveralls.io/github/kolkov/esr?branch=master)
[![Build Status](https://travis-ci.org/kolkov/esr.svg?branch=master)](https://travis-ci.org/kolkov/esr)

Compatible with the ozzo-validation package.

### Usage ###

```
import "github.com/kolkov/esr"

err := esr.Validate("639400")

signed := esr.Generate("1")
```
