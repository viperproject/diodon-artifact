The script terminates successfully (exit code 0) and its output should end in the following lines, which state that all lemmas were successfully verified.

```
/* All wellformedness checks were successful. */

/*
Generated from:
Tamarin version 1.8.0
Maude version 2.7.1
Git revision: f172d7f00b1485446a1e7a42dc14623c2189cc42, branch: master
Compiled at: 2023-09-01 08:49:23.916032222 UTC
*/

end

==============================================================================
summary of summaries:

analyzed: /gobra/dh/model/protocol-model.spthy

  processing time: 4.58s
  
  sources (all-traces): verified (443 steps)
  exec (exists-trace): verified (10 steps)
  ni_agreement_init (all-traces): verified (7 steps)
  i_agreement_init (all-traces): verified (33 steps)
  ni_agreement_resp (all-traces): verified (7 steps)
  i_agreement_resp (all-traces): verified (33 steps)
  forward_secrecy (all-traces): verified (109 steps)

==============================================================================

real    0m4.749s
user    0m27.002s
sys     0m1.635s
```
