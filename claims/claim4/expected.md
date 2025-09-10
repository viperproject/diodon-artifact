The immutability analysis succeeds, the escape analysis finds at most one escaping CORE instance and escaping logger objects, the pointer analysis checking non-aliasing of parameters (i.e., condition (C7)) succeeds, and the pass-through analysis fails due to false-positives.
This expected result is indicated by observing output that ends in the following:

```
real	3m38.713s
user	9m58.702s
sys	0m48.716s
Conditions were checked successfully modulo the stated exceptions.
```
