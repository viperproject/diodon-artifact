The taint analysis should succeed with output ending similarly to the following:

```
[INFO].dh Taint analysis took 3.7608 s
[INFO].dh 
[INFO].dh TARGET dh RESULT:
                No taint flows detected ✓
[INFO].dh TARGET dh ESCAPE ANALYSIS RESULT:
                Tainted data does not escape ✓
[INFO].dh ********************************************************************************
[INFO] Wrote final report in /gobra/dh/argot-proofs/argot-reports/overall-report-3602893344.json

real    0m32.541s
user    1m22.220s
sys     0m11.437s
```
