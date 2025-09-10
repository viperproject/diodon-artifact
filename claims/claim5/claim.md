# Claim 5 - Protocol Model for Diffie-Hellman
As described in Sec. 5.2, we model the security protocol for the signed Diffie-Hellman key exchange in Tamarin and prove (1) forward secrecy for the two symmetric session keys and (2) injective agreement on the identities of the two involved protocol role instances and the session keys.

`run.sh` supports this claim by invoking Tamarin on the protocol model stored in `dh/model/protocol-model.spthy` and proving all lemmas stated therein.
In particular, `i_agreement_init` and `i_agreement_resp` prove injective agreement from the perspective of the initiator and responder, respectively, and `forward_secrecy` proves forward secrecy for the symmetric session keys.
Cf. claim 1 for a more detailed explanation of lemmas in Tamarin.
