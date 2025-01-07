package library

import fmt "fmt"

//@ import . "dh-gobra/verification/bytes"
//@ import . "dh-gobra/verification/place"
//@ import . "dh-gobra/verification/iospec"
//@ import tm "dh-gobra/verification/utilterm"

// calling this function results in a proof obligations that the necessary I/O permission
// for performing an output operation are present. Thus, we can treat this function as
// sanitizing `data`
// technically, fractional permissions for `data` is sufficient. However, we express the specification
// in a way that it does not matter for callers whether this function copies the slice of bytes or returns
// the same one.
// @ trusted
// @ requires Mem(data) && data != nil
// @ requires token(t) && e_OutFact(t, rid, m) && gamma(m) == Abs(data)
// @ ensures  token(t1) && t1 == old(get_e_OutFact_placeDst(t, rid, m))
// @ ensures  Mem(res) && Abs(res) == old(Abs(data)) && res != nil
func PerformVirtualOutputOperation(data []byte /*@, ghost t Place, ghost rid tm.Term, ghost m tm.Term @*/) (res []byte /*@, ghost t1 Place @*/) {
	// due to an imprecision in the data flow analysis, we have to copy the slice instead of directly returning `data`
	// otherwise, the data flow analysis considers the result tainted despite configuring this function as a sanitizer
	res = append([]byte(nil), data...)
	return
}

// calling this function results in treating `data` being treated as received from the environment / attacker (on the
// level of the abstract model). Thus, this function can be seen as performing a virtual input operation.
// However, since `data` is treated as coming from the attacker, `data` must not be tainted.
// @ trusted
// @ decreases
// @ requires acc(Mem(data), 1/16)
// @ requires token(t) && e_InFact(t, rid)
// @ ensures  acc(Mem(data), 1/16) && gamma(dataT) == Abs(data)
// @ ensures  token(t1) && t1 == old(get_e_InFact_placeDst(t, rid))
// @ ensures  dataT == old(get_e_InFact_r1(t, rid))
func PerformVirtualInputOperation(data []byte /*@, ghost t Place, ghost rid tm.Term @*/) /*@ (ghost t1 Place, dataT tm.Term) @*/ {
	return
}

// @ trusted
// @ preserves acc(l.Mem(), 1/16)
// @ preserves acc(Mem(sharedSecret), 1/16)
func (l *LibState) PrintSharedSecret(sharedSecret []byte) {
	fmt.Printf("Initiator & responder agreed on shared secret %x\n", sharedSecret)
}

// @ trusted
// @ preserves acc(l.Mem(), 1/16)
// @ preserves acc(Mem(irKey), 1/16) && acc(Mem(riKey), 1/16)
func (l *LibState) PrintKeys(irKey, riKey []byte) {
	fmt.Printf("IR Key %x\nRI Key %x\n", irKey, riKey)
}
