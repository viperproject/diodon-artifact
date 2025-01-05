package library

import fmt "fmt"
//@ import . "dh-gobra/verification/bytes"
//@ import . "dh-gobra/verification/place"
//@ import . "dh-gobra/verification/iospec"
//@ import tm "dh-gobra/verification/utilterm"


//@ trusted
//@ preserves acc(l.Mem(), 1/16)
//@ requires acc(Mem(data), 1/16)
//@ requires token(t) && e_OutFact(t, rid, m) && gamma(m) == Abs(data)
//@ ensures  acc(Mem(data), 1/16)
//@ ensures  token(t1) && t1 == old(get_e_OutFact_placeDst(t, rid, m))
func (l *LibState) Declassify(data []byte /*@, ghost t Place, ghost rid tm.Term, ghost m tm.Term @*/) (/*@ ghost t1 Place @*/) {
	return
}

//@ trusted
//@ preserves acc(l.Mem(), 1/16)
//@ preserves acc(Mem(sharedSecret), 1/16)
func (l *LibState) PrintSharedSecret(sharedSecret []byte) {
	fmt.Printf("Initiator & responder agreed on shared secret %x\n", sharedSecret)
}

//@ trusted
//@ preserves acc(l.Mem(), 1/16)
//@ preserves acc(Mem(irKey), 1/16) && acc(Mem(riKey), 1/16)
func (l *LibState) PrintKeys(irKey, riKey []byte) {
	fmt.Printf("IR Key %x\nRI Key %x\n", irKey, riKey)
}
