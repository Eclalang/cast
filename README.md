# Cast library

[![Go Report Card](https://goreportcard.com/badge/github.com/Eclalang/cast)](https://goreportcard.com/report/github.com/Eclalang/cast)
[![codecov](https://codecov.io/gh/Eclalang/cast/graph/badge.svg?token=0ROO3NXG4P)](https://codecov.io/gh/Eclalang/cast)

## Candidate functions :

| Func Name  |                  Prototype                   |       Description        |                             Comments                              |
|:----------:|:--------------------------------------------:|:------------------------:|:-----------------------------------------------------------------:|
|    Atoi    |           Atoi(str string) int {}            |  Converts string to int  |                                N/A                                |
| FloatToInt |          FloatToInt(x float) int {}          |  Converts float to int   |                                N/A                                |
| IntToFloat |          IntToFloat(x int) float {}          |  Converts int to float   |                                N/A                                |
| ParseBool  |        ParseBool(str string) bool {}         | Converts string to bool  | Accepts : 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False |
| ParseFloat | ParseFloat(str string, bitSize int) float {} | Converts string to float |                                N/A                                |