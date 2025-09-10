The output should end in the following lines, which state that all lemmas were successfully verified.

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

analyzed: /gobra/ssm-agent/model/protocol-model.spthy

  processing time: 585.38s
  
  agent_recv_sign_response_valid (all-traces): verified (8 steps)
  client_recv_sign_response_valid (all-traces): verified (7 steps)
  agent_recv_transport_valid (all-traces): verified (46 steps)
  kms_sign_response_unique (all-traces): verified (20 steps)
  loop_induction_client (all-traces): verified (12 steps)
  loop_induction_agent (all-traces): verified (12 steps)
  Agent_can_finish_wo_reveal (exists-trace): verified (27 steps)
  Client_can_finish_wo_reveal (exists-trace): verified (21 steps)
  x_is_secret (all-traces): verified (3 steps)
  y_is_secret (all-traces): verified (3 steps)
  AgentSendEncryptedSessionKey_is_unique (all-traces): verified (4 steps)
  SharedSecret_is_secret (all-traces): verified (1351 steps)
  injectiveagreement_agent (all-traces): verified (240 steps)
  injectiveagreement_client (all-traces): verified (2486 steps)

==============================================================================

real    9m45.899s
user    22m55.860s
sys     46m29.489s
```
