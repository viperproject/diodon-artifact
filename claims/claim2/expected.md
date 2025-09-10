The taint analysis should succeed with output similar to the following:

```
Running taint analysis on SSM agent in directory /gobra/ssm-agent/implementation
[INFO] Argot taint tool - v0.4.8-alpha-DIODON
[INFO].ssm-session-worker Taint analysis of target "ssm-session-worker" = [agent/framework/processor/executer/outofproc/sessionworker/main.go]
[INFO].ssm-session-worker Loaded module github.com/aws/amazon-ssm-agent
[INFO].ssm-session-worker Loaded 21 annotations from program
[INFO].ssm-session-worker Gathering values and starting pointer analysis...
[INFO].ssm-session-worker Pointer analysis terminated (12.02 s)
[INFO].ssm-session-worker Gathering global variable declaration in the program...
[INFO].ssm-session-worker Computing information about types and functions for analysis...
[INFO].ssm-session-worker Global gathering terminated, added 3966 items (0.00 s)
[INFO].ssm-session-worker Pointer analysis state computed, added 2680 items (12.06 s)
[INFO].ssm-session-worker Starting intra-procedural analysis ...
[INFO].ssm-session-worker Intra-procedural pass done (1.27 s).
[INFO].ssm-session-worker.diodon-agent-io-independence Starting inter-procedural pass...
[INFO].ssm-session-worker.diodon-agent-io-independence Scanning for entry points ...
[INFO].ssm-session-worker.diodon-agent-io-independence ╭──────────────────────────╮
[INFO].ssm-session-worker.diodon-agent-io-independence │  1 analysis entry points │
[INFO].ssm-session-worker.diodon-agent-io-independence ╰──────────────────────────╯
[INFO].ssm-session-worker.diodon-agent-io-independence.#26127.5 //argot:ignore implicit flow to conditional statement at /gobra/ssm-agent/implementation/agent/session/plugins/sessionplugin/sessionplugin.go:78:9
[INFO].ssm-session-worker.diodon-agent-io-independence.#26127.5 //argot:ignore implicit flow to conditional statement at /gobra/ssm-agent/implementation/agent/session/datachannel/helper.go:414:8
[INFO].ssm-session-worker.diodon-agent-io-independence.#26127.5 //argot:ignore implicit flow to conditional statement at /gobra/ssm-agent/implementation/agent/session/datachannel/datachannel.go:330:8
[INFO].ssm-session-worker.diodon-agent-io-independence.#26127.5 //argot:ignore implicit flow to conditional statement at /gobra/ssm-agent/implementation/agent/session/datachannel/datachannel.go:36:8
[INFO].ssm-session-worker.diodon-agent-io-independence.#26127.5 //argot:ignore implicit flow to conditional statement at /gobra/ssm-agent/implementation/agent/session/datachannel/datachannel.go:354:8
[INFO].ssm-session-worker.diodon-agent-io-independence.#26127.5 //argot:ignore implicit flow to conditional statement at /gobra/ssm-agent/implementation/agent/session/datachannel/datachannel.go:378:8
[INFO].ssm-session-worker.diodon-agent-io-independence.#26127.5 //argot:ignore implicit flow to conditional statement at /gobra/ssm-agent/implementation/agent/session/datachannel/send.go:165:91
[INFO].ssm-session-worker.diodon-agent-io-independence inter-procedural pass done (8.34 s).
[INFO].ssm-session-worker.diodon-agent-io-independence Done analyzing diodon-agent-io-independence
[INFO].ssm-session-worker 
[INFO].ssm-session-worker Taint analysis took 9.6177 s
[INFO].ssm-session-worker 
[INFO].ssm-session-worker TARGET ssm-session-worker RESULT:
                No taint flows detected ✓
[INFO].ssm-session-worker ********************************************************************************
[INFO] Wrote final report in /gobra/ssm-agent/argot-proofs/argot-reports/overall-report-2817160620.json

real    1m26.127s
user    6m39.533s
sys     0m48.969s
```
