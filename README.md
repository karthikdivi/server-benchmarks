
## Commands

```sh
time go run go/measure.go
time node nodejs/measure.js
```

## Results

|language|command|result|time stats|
|---|---|---|---|
|Go| `go run go/measure.go` | `resultsLength: 1105478, resultsSize: 92383535, timeTaken: 743` | `1.04s user 0.55s system 116% cpu 1.362 total`|
|NodeJS| `go run go/measure.go` | `resultsLength: 1105478, resultsSize: 92383525, timeTaken: 939` | `1.01s user 0.14s system 112% cpu 1.023 total`|
