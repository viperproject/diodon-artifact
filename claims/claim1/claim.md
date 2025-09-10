# Claim 1 -- Protocol Model for SSM Agent
As described in Sec. 5.1.1, we model the security protocol for establishing a remote shell session between an SSM Agent running on an AWS EC2 instance and an AWS customer in Tamarin and prove (1) secrecy for the two symmetric session keys and (2) injective agreement on the SSM Agent's and AWS customer's identities and the session keys.

`run.sh` supports this claim by invoking Tamarin on the protocol model stored in `ssm-agent/model/protocol-model.spthy` and proving all lemmas stated therein.
The following lemmas are particularly noteworthy:
- `Agent_can_finish_wo_reveal` and `Client_can_finish_wo_reveal` prove that an SSM Agent and an AWS customer (called "client" in the model) can successfully execute the protocol without attacker intervention.
- `SharedSecret_is_secret` proves secrecy (1). Whenever a transition issues a `Secret(x)` action label for some term `x` then this lemma proves that `x` is unknown to the attacker (`not (Ex #j. K(x)@j)`) unless a signing key within the Key Management Service (KMS) or a long-term secret key is corrupted. The model states secrecy for the session keys `kdf1(g^x^y)` and `kdf2(g^x^y)` in the rules `Agent_SendHandshakeComplete` and `Client_RecvHandshakeComplete`.
- `injectiveagreement_agent` and `injectiveagreement_client` prove injective agreement (2) from the SSM Agent's and AWS customer's perspective, respectively, under the same corruption conditions.
