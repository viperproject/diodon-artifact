package initiator

import (
	. "dh-gobra/library"
)

//@ import arb "dh-gobra/verification/arb"
//@ import by "dh-gobra/verification/bytes"
//@ import cl "dh-gobra/verification/claim"
//@ import ft "dh-gobra/verification/fact"
//@ import pl "dh-gobra/verification/place"
//@ import io "dh-gobra/verification/iospec"
//@ import tm "dh-gobra/verification/utilterm"
//@ import ay "dh-gobra/verification/utilbytes"
//@ import am "dh-gobra/verification/term"
//@ import p "dh-gobra/verification/pattern"
//@ import pub "dh-gobra/verification/pub"
//@ import fresh "dh-gobra/verification/fresh"

type Initiator struct {
	initiatorState InitiatorState
	l              *LibState
	idA            uint32
	idB            uint32
	skA            []byte
	pkB            []byte
	x              []byte
	X              []byte
	Y              []byte
	irKey          []byte
	riKey          []byte
	//@ ghost skAT tm.Term
	//@ ghost skBT tm.Term
	//@ ghost token pl.Place
	//@ ghost rid tm.Term
	//@ ghost absState mset[ft.Fact]
	//@ ghost xT tm.Term
	//@ ghost YT tm.Term
}

type InitiatorState int

const (
	Erroneous          InitiatorState = 0
	Initialized        InitiatorState = 1
	ProducedHsMsg1     InitiatorState = 2
	ProcessedHsMsg2    InitiatorState = 3
	HandshakeCompleted InitiatorState = 4
)

/*@
pred (i *Initiator) Inv() {
	acc(i) &&
	(i.initiatorState != Erroneous ==>
		i.l.Mem() && acc(Mem(i.skA), 1/4) && acc(Mem(i.pkB), 1/4) &&
		Abs(i.skA) == by.gamma(i.skAT) && Abs(i.pkB) == by.gamma(tm.pk(i.skBT)) &&
		pl.token(i.token) && io.P_Alice(i.token, i.rid, i.absState)) &&
	(i.initiatorState == Initialized ==>
		InitializedPred(i.rid, i.idA, i.idB, i.skAT, i.skBT, i.absState)) &&
	(i.initiatorState >= ProducedHsMsg1 ==>
		Mem(i.x) && Abs(i.x) == by.gamma(i.xT) &&
		Mem(i.X) && Abs(i.X) == by.gamma(tm.exp(tm.generator(), i.xT))) &&
	(i.initiatorState == ProducedHsMsg1 ==>
		ProducedHsMsg1Pred(i.rid, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.absState)) &&
	(i.initiatorState >= ProcessedHsMsg2 ==>
		Mem(i.Y) && Abs(i.Y) == by.gamma(i.YT) &&
		ProcessedHsMsg2Pred(i.rid, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.YT, i.absState)) &&
	(i.initiatorState == ProcessedHsMsg2 ==>
		HasHsMsg3OutFact(i.rid, i.idA, i.idB, i.YT, i.xT, i.skAT, i.absState)) &&
	(i.initiatorState >= HandshakeCompleted ==>
		HandshakeCompletedPred(i.irKey, i.riKey, i.xT, i.YT))
}

pred InitializedPred(rid tm.Term, idA, idB uint32, skAT, skBT tm.Term, s mset[ft.Fact]) {
	ft.Setup_Alice(rid, tm.integer32(idA), tm.integer32(idB), skAT, skBT) in s
}

pred ProducedHsMsg1Pred(rid tm.Term, idA, idB uint32, skAT, skBT, xT tm.Term, s mset[ft.Fact]) {
	ft.St_Alice_1(rid, tm.integer32(idA), tm.integer32(idB), skAT, skBT, xT) in s
}

pred ProcessedHsMsg2Pred(rid tm.Term, idA, idB uint32, skAT, skBT, xT, YT tm.Term, s mset[ft.Fact]) {
	ft.St_Alice_2(rid, tm.integer32(idA), tm.integer32(idB), skAT, skBT, xT, YT) in s
}

pred HasHsMsg3OutFact(rid tm.Term, idA, idB uint32, YT, xT, skAT tm.Term, s mset[ft.Fact]) {
	ft.OutFact_Alice(rid, tm.sign(tm.tuple5(tm.integer32(Msg3Tag), tm.integer32(idA), tm.integer32(idB), YT, tm.exp(tm.generator(), xT)), skAT)) in s
}

pred HandshakeCompletedPred(irKey, riKey []byte, xT, YT tm.Term) {
	Mem(irKey) && Abs(irKey) == by.gamma(tm.kdf1(tm.exp(YT, xT))) &&
	Mem(riKey) && Abs(riKey) == by.gamma(tm.kdf2(tm.exp(YT, xT)))
}

ghost
decreases
requires acc(i.Inv(), _)
pure
func (i *Initiator) getIdA() uint32 {
	return unfolding acc(i.Inv(), _) in i.idA
}

ghost
decreases
requires acc(i.Inv(), _)
pure
func (i *Initiator) getIdB() uint32 {
	return unfolding acc(i.Inv(), _) in i.idB
}
@*/

// @ ensures i.Inv()
func NewInitiator(privateKey [64]byte, peerPublicKey [32]byte) (i *Initiator, success bool) {
	// pick an arbitrary rid for this protocol session and inhale the IO specification for
	// the SSM agent and the chosen protocol session:
	//@ t0, rid, s0 := arb.GetArbPlace(), arb.GetArbTerm(), mset[ft.Fact]{}
	//@ inhale pl.token(t0) && io.P_Alice(t0, rid, s0)

	i = &Initiator{}
	var l *LibState
	l, err := NewLibState(0, 1, privateKey, peerPublicKey)
	if err != nil {
		//@ fold i.Inv()
		return
	}

	//@ unfold io.P_Alice(t0, rid, s0)
	//@ unfold io.phiRF_Alice_7(t0, rid, s0)
	//@ assert acc(io.e_Setup_Alice(t0, rid))
	//@ skAT := io.get_e_Setup_Alice_r3(t0, rid)
	//@ skBT := io.get_e_Setup_Alice_r4(t0, rid)

	var idA, idB uint32
	var skA, pkB []byte
	//@ ghost var t1 pl.Place
	idA, idB, skA, pkB, err /*@, t1 @*/ = l.Setup( /*@ t0, rid @*/ )
	//@ s1 := mset[ft.Fact]{ ft.Setup_Alice(rid, tm.integer32(idA), tm.integer32(idB), skAT, skBT) }
	//@ fold InitializedPred(rid, idA, idB, skAT, skBT, s1)
	if err != nil {
		//@ fold i.Inv()
		return
	}

	i.initiatorState = Initialized
	i.l = l
	i.idA = idA
	i.idB = idB
	i.skA = skA
	i.pkB = pkB
	//@ i.skAT = skAT
	//@ i.skBT = skBT
	//@ i.token = t1
	//@ i.rid = rid
	//@ i.absState = s1

	//@ fold i.Inv()
	success = true
	return
}

// @ preserves i != nil ==> i.Inv()
// @ ensures   success ==> msg != nil
// @ ensures   msg != nil ==> Mem(msg)
func (i *Initiator) ProduceHsMsg1() (msg []byte, success bool) {
	if i == nil { //argot:ignore diodon-dh-io-independence
		return
	}
	//@ unfold i.Inv()
	if i.initiatorState != Initialized {
		//@ fold i.Inv()
		return
	}
	//@ t0, ridT, s0 := i.token, i.rid, i.absState

	// create x
	//@ unfold io.P_Alice(t0, ridT, s0)
	//@ unfold io.phiRF_Alice_5(t0, ridT, s0)
	//@ assert acc(io.e_FrFact(t0, ridT))
	//@ i.xT = io.get_e_FrFact_r1(t0, ridT)
	//@ ghost var t1 pl.Place
	var err error
	i.x, err /*@, t1 @*/ = i.l.CreateNonce( /*@ t0, ridT @*/ )
	//@ s1 := s0 union mset[ft.Fact]{ ft.FrFact_Alice(ridT, i.xT) }
	if err != nil {
		//@ fold io.phiRF_Alice_5(t0, ridT, s0)
		//@ fold io.P_Alice(t0, ridT, s0)
		//@ fold i.Inv()
		return
	}

	//@ unfold InitializedPred(ridT, i.idA, i.idB, i.skAT, i.skBT, s0)
	// create X, i.e., X = g^x
	i.X, err = i.l.DhExp(i.x)
	//@ XT := tm.exp(tm.generator(), i.xT)
	if err != nil {
		//@ i.token = t1
		//@ i.absState = s1
		//@ fold InitializedPred(ridT, i.idA, i.idB, i.skAT, i.skBT, s1)
		//@ fold i.Inv()
		return
	}

	//@ unfold io.P_Alice(t1, ridT, s1)
	//@ unfold io.phiR_Alice_0(t1, ridT, s1)
	//@ idAT := tm.integer32(i.idA)
	//@ idBT := tm.integer32(i.idB)
	/*@
	l := mset[ft.Fact]{ ft.Setup_Alice(ridT, idAT, idBT, i.skAT, i.skBT),
		ft.FrFact_Alice(ridT, i.xT) }
	a := mset[cl.Claim]{}
	r := mset[ft.Fact]{ ft.St_Alice_1(ridT, idAT, idBT, i.skAT, i.skBT, i.xT),
		ft.OutFact_Alice(ridT, XT) }
	@*/
	//@ assert io.e_Alice_send(t1, ridT, idAT, idBT, i.skAT, i.skBT, i.xT, l, a, r)
	//@ t2 := io.internBIO_e_Alice_send(t1, ridT, idAT, idBT, i.skAT, i.skBT, i.xT, l, a, r)
	//@ s2 := ft.U(l, r, s1)

	msg1 := &Msg1{X: i.X}
	//@ fold acc(msg1.Mem(), 1/8)
	msg, err = i.l.MarshalMsg1(msg1)
	//@ unfold acc(msg1.Mem(), 1/8)
	if err != nil {
		i.initiatorState = Erroneous
		//@ i.token = t2
		//@ i.absState = s2
		//@ fold i.Inv()
		msg = nil
		return
	}

	//@ unfold io.P_Alice(t2, ridT, s2)
	//@ unfold io.phiRG_Alice_4(t2, ridT, s2)
	//@ assert io.e_OutFact(t2, ridT, XT)
	//@ ghost var t3 pl.Place
	msg /*@, t3 @*/ = PerformVirtualOutputOperation(msg /*@, t2, ridT, XT @*/)
	//@ s3 := s2 setminus mset[ft.Fact]{ ft.OutFact_Alice(ridT, XT) }
	//@ fold ProducedHsMsg1Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, s3)
	i.initiatorState = ProducedHsMsg1
	//@ i.token = t3
	//@ i.absState = s3
	//@ fold i.Inv()
	success = true
	return
}

// @ preserves i != nil ==> i.Inv()
// @ preserves msg != nil ==> Mem(msg)
func (i *Initiator) ProcessHsMsg2(msg []byte) (success bool) {
	if i == nil || msg == nil { //argot:ignore diodon-dh-io-independence
		return
	}
	//@ unfold i.Inv()
	if i.initiatorState != ProducedHsMsg1 {
		//@ fold i.Inv()
		return
	}
	//@ t0, ridT, s0 := i.token, i.rid, i.absState
	//@ unfold ProducedHsMsg1Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, s0)

	//@ unfold io.P_Alice(t0, ridT, s0)
	//@ unfold io.phiRF_Alice_6(t0, ridT, s0)
	//@ assert io.e_InFact(t0, ridT)
	/*@ t1, msgT := @*/
	PerformVirtualInputOperation(msg /*@, t0, ridT @*/)
	//@ s1 := s0 union mset[ft.Fact]{ ft.InFact_Alice(ridT, msgT) }

	msg2Data, err := i.l.Open(msg, i.pkB /*@, i.skBT @*/)
	if err != nil {
		//@ i.token = t1
		//@ i.absState = s1
		//@ fold ProducedHsMsg1Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, s1)
		//@ fold i.Inv()
		return
	}

	msg2, err := i.l.UnmarshalMsg2(msg2Data)
	if err != nil {
		//@ i.token = t1
		//@ i.absState = s1
		//@ fold ProducedHsMsg1Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, s1)
		//@ fold i.Inv()
		return
	}

	//@ unfold msg2.Mem()
	if msg2.IdA != i.idA || msg2.IdB != i.idB {
		//@ i.token = t1
		//@ i.absState = s1
		//@ fold ProducedHsMsg1Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, s1)
		//@ fold i.Inv()
		return
	}
	i.Y = msg2.Y

	// check receivedX
	if !Equals(i.X, msg2.X) {
		//@ i.token = t1
		//@ i.absState = s1
		//@ fold ProducedHsMsg1Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, s1)
		//@ fold i.Inv()
		return
	}

	//@ idAT := tm.integer32(i.idA)
	//@ idBT := tm.integer32(i.idB)
	//@ XT := tm.exp(tm.generator(), i.xT)
	//@ YT := p.patternRequirement2(ridT, idAT, idBT, i.skAT, i.skBT, i.xT, by.oneTerm(Abs(i.Y)), msgT, t1, s1)
	// the following 2 assert stmts are needed for triggering reasons:
	//@ assert ay.getMsgB(Abs(msg)) == Abs(msg2Data)
	//@ assert by.ex55B(Abs(msg2Data)) == Abs(i.Y)
	//@ assert Abs(i.Y) == by.gamma(YT)

	//@ unfold io.P_Alice(t1, ridT, s1)
	//@ unfold io.phiR_Alice_1(t1, ridT, s1)
	//@ msg3T := tm.sign(tm.tuple5(tm.integer32(Msg3Tag), idAT, idBT, YT, XT), i.skAT)
	/*@
		l := mset[ft.Fact]{ ft.St_Alice_1(ridT, idAT, idBT, i.skAT, i.skBT, i.xT),
			ft.InFact_Alice(ridT, msgT) }
		a := mset[cl.Claim]{
			cl.IN_ALICE(YT, tm.tuple5(tm.integer32(Msg2Tag), idBT, idAT, XT, YT)),
	        cl.Secret(idAT, idBT, tm.exp(YT, i.xT)),
	        cl.Running(tm.idR(), tm.idI(), tm.tuple3(idAT, idBT, tm.exp(YT, i.xT))),
	        cl.Commit(tm.idI(), tm.idR(), tm.tuple3(idAT, idBT, tm.exp(YT, i.xT))),
			cl.AliceHsDone(tm.exp(YT, i.xT)) }
		r := mset[ft.Fact]{ ft.St_Alice_2(ridT, idAT, idBT, i.skAT, i.skBT, i.xT, YT),
			ft.OutFact_Alice(ridT, msg3T) }
		@*/
	//@ assert io.e_Alice_recvAndSend(t1, ridT, idAT, idBT, i.skAT, i.skBT, i.xT, YT, l, a, r)
	//@ t2 := io.internBIO_e_Alice_recvAndSend(t1, ridT, idAT, idBT, i.skAT, i.skBT, i.xT, YT, l, a, r)
	//@ s2 := ft.U(l, r, s1)
	//@ fold ProcessedHsMsg2Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, YT, s2)
	//@ fold HasHsMsg3OutFact(ridT, i.idA, i.idB, YT, i.xT, i.skAT, s2)
	i.initiatorState = ProcessedHsMsg2
	//@ i.token = t2
	//@ i.absState = s2
	//@ i.YT = YT
	//@ fold i.Inv()
	success = true
	return
}

// @ preserves i != nil ==> i.Inv()
// @ ensures   success ==> signedMsg3 != nil
// @ ensures   signedMsg3 != nil ==> Mem(signedMsg3)
func (i *Initiator) ProduceHsMsg3() (signedMsg3 []byte, success bool) {
	if i == nil { //argot:ignore diodon-dh-io-independence
		return
	}
	//@ unfold i.Inv()
	if i.initiatorState != ProcessedHsMsg2 {
		//@ fold i.Inv()
		return
	}
	//@ t0, ridT, s0 := i.token, i.rid, i.absState

	msg3 := &Msg3{IdA: i.idA, IdB: i.idB, X: i.X, Y: i.Y}
	//@ fold acc(msg3.Mem(), 1/8)
	msg3Data, err := i.l.MarshalMsg3(msg3)
	//@ unfold acc(msg3.Mem(), 1/8)
	if err != nil { //argot:ignore diodon-dh-io-independence
		//@ fold i.Inv()
		return
	}

	signedMsg3, err = i.l.Sign(msg3Data, i.skA)
	if err != nil { //argot:ignore diodon-dh-io-independence
		//@ fold i.Inv()
		signedMsg3 = nil
		return
	}

	//@ requires acc(i, 1/2) && acc(i.l.Mem(), 1/2)
	//@ requires Mem(signedMsg3) && signedMsg3 != nil && Abs(signedMsg3) == by.signB(ay.tuple5B(ay.integer32B(Msg3Tag), ay.integer32B(i.idA), ay.integer32B(i.idB), by.gamma(i.YT), by.expB(ay.generatorB(), by.gamma(i.xT))), by.gamma(i.skAT))
	//@ requires pl.token(t0) && io.P_Alice(t0, ridT, s0)
	//@ requires HasHsMsg3OutFact(ridT, i.idA, i.idB, i.YT, i.xT, i.skAT, s0)
	//@ requires ProcessedHsMsg2Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.YT, s0)
	//@ ensures  acc(i, 1/2) && acc(i.l.Mem(), 1/2)
	//@ ensures  Mem(signedMsg3) && signedMsg3 != nil && Abs(signedMsg3) == before(Abs(signedMsg3))
	//@ ensures  pl.token(t1) && io.P_Alice(t1, ridT, s1)
	//@ ensures  ProcessedHsMsg2Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.YT, s1)
	//@ outline(
	//@ unfold io.P_Alice(t0, ridT, s0)
	//@ unfold io.phiRG_Alice_4(t0, ridT, s0)
	//@ unfold ProcessedHsMsg2Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.YT, s0)
	//@ unfold HasHsMsg3OutFact(ridT, i.idA, i.idB, i.YT, i.xT, i.skAT, s0)
	//@ XT := tm.exp(tm.generator(), i.xT)
	//@ msgT := tm.sign(tm.tuple5(tm.integer32(Msg3Tag), tm.integer32(i.idA), tm.integer32(i.idB), i.YT, XT), i.skAT)
	//@ assert acc(io.e_OutFact(t0, ridT, msgT))
	//@ ghost var t1 pl.Place
	signedMsg3 /*@, t1 @*/ = PerformVirtualOutputOperation(signedMsg3 /*@, t0, ridT, msgT @*/)
	//@ s1 := s0 setminus mset[ft.Fact]{ ft.OutFact_Alice(ridT, msgT) }
	//@ fold ProcessedHsMsg2Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.YT, s1)
	//@ )

	//@ i.token = t1
	//@ i.absState = s1

	//@ preserves acc(i, 1/2) && acc(i.l.Mem(), 1/2)
	//@ preserves acc(Mem(i.x), 1/2) && acc(Mem(i.Y), 1/2)
	//@ preserves Abs(i.x) == by.gamma(i.xT) && Abs(i.Y) == by.gamma(i.YT)
	//@ preserves acc(&i.irKey, 1/2) && acc(&i.riKey, 1/2)
	//@ ensures   err == nil ==> HandshakeCompletedPred(i.irKey, i.riKey, i.xT, i.YT)
	//@ outline(
	var sharedSecret []byte
	//@ ghost var sharedSecretB by.Bytes
	sharedSecret, err /*@, sharedSecretB @*/ = i.l.DhSharedSecret(i.x, i.Y)
	if err == nil { //argot:ignore diodon-dh-io-independence
		i.irKey, i.riKey = NewBytes(32), NewBytes(32)
		err = KDF2Slice(i.irKey, i.riKey, sharedSecret)
		if err == nil {
			i.l.PrintKeys(i.irKey, i.riKey)
			//@ fold HandshakeCompletedPred(i.irKey, i.riKey, i.xT, i.YT)
		}
	}
	//@ )

	if err != nil {
		i.initiatorState = Erroneous
		//@ fold i.Inv()
		return
	}
	i.initiatorState = HandshakeCompleted
	//@ fold i.Inv()
	success = true
	return
}

// @ preserves i != nil ==> i.Inv()
// @ preserves msgData != nil ==> Mem(msgData)
// @ ensures   success ==> payload != nil
// @ ensures   payload != nil ==> Mem(payload)
func (i *Initiator) ProcessTransportMsg(msgData []byte) (payload []byte, success bool) {
	if i == nil || msgData == nil { //argot:ignore diodon-dh-io-independence
		return
	}
	//@ unfold i.Inv()
	if i.initiatorState != HandshakeCompleted {
		//@ fold i.Inv()
		return
	}
	//@ t0, ridT, s0 := i.token, i.rid, i.absState
	//@ unfold ProcessedHsMsg2Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.YT, s0)

	//@ unfold io.P_Alice(t0, ridT, s0)
	//@ unfold io.phiRF_Alice_6(t0, ridT, s0)
	//@ assert io.e_InFact(t0, ridT)
	/*@ t1, msgDataT := @*/
	PerformVirtualInputOperation(msgData /*@, t0, ridT @*/)
	//@ s1 := s0 union mset[ft.Fact]{ ft.InFact_Alice(ridT, msgDataT) }

	ciphertext, err := i.l.UnmarshalTransportMsg(msgData)
	if err != nil {
		//@ i.token = t1
		//@ i.absState = s1
		//@ fold ProcessedHsMsg2Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.YT, s1)
		//@ fold i.Inv()
		return
	}

	//@ unfold HandshakeCompletedPred(i.irKey, i.riKey, i.xT, i.YT)
	payload, err = i.l.Decrypt(ciphertext, i.riKey /*@, tm.kdf2(tm.exp(i.YT, i.xT)) @*/)
	if err != nil {
		//@ i.token = t1
		//@ i.absState = s1
		//@ fold ProcessedHsMsg2Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.YT, s1)
		//@ fold HandshakeCompletedPred(i.irKey, i.riKey, i.xT, i.YT)
		//@ fold i.Inv()
		payload = nil
		return
	}

	//@ idAT := tm.integer32(i.idA)
	//@ idBT := tm.integer32(i.idB)
	//@ payloadT := p.patternRequirementTransMsg(ridT, idAT, idBT, i.skAT, i.skBT, i.xT, i.YT, by.oneTerm(Abs(payload)), msgDataT, t1, s1)
	// the following 2 assert stmts are needed for triggering reasons:
	//@ assert by.ex22B(Abs(msgData)) == Abs(ciphertext)
	//@ assert by.sdecB(Abs(ciphertext), Abs(i.riKey)) == Abs(payload)
	//@ assert Abs(payload) == by.gamma(payloadT)
	//@ fold HandshakeCompletedPred(i.irKey, i.riKey, i.xT, i.YT)

	//@ unfold io.P_Alice(t1, ridT, s1)
	//@ unfold io.phiR_Alice_3(t1, ridT, s1)
	/*@
	l := mset[ft.Fact]{ ft.St_Alice_2(ridT, idAT, idBT, i.skAT, i.skBT, i.xT, i.YT),
		ft.InFact_Alice(ridT, msgDataT) }
	a := mset[cl.Claim]{
		cl.AliceRecvLoop(tm.exp(i.YT, i.xT)),
		cl.AliceRecvTransMsg(payloadT, tm.kdf2(tm.exp(i.YT, i.xT))) }
	r := mset[ft.Fact]{ ft.St_Alice_2(ridT, idAT, idBT, i.skAT, i.skBT, i.xT, i.YT),
		ft.OutFact_Alice(ridT, payloadT) }
	@*/
	//@ assert io.e_Alice_recvMsg(t1, ridT, idAT, idBT, i.skAT, i.skBT, i.xT, i.YT, payloadT, l, a, r)
	//@ t2 := io.internBIO_e_Alice_recvMsg(t1, ridT, idAT, idBT, i.skAT, i.skBT, i.xT, i.YT, payloadT, l, a, r)
	//@ s2 := ft.U(l, r, s1)

	//@ unfold io.P_Alice(t2, ridT, s2)
	//@ unfold io.phiRG_Alice_4(t2, ridT, s2)
	//@ assert acc(io.e_OutFact(t2, ridT, payloadT))
	//@ ghost var t3 pl.Place
	payload /*@, t3 @*/ = PerformVirtualOutputOperation(payload /*@, t2, ridT, payloadT @*/)
	//@ s3 := s2 setminus mset[ft.Fact]{ ft.OutFact_Alice(ridT, payloadT) }

	//@ i.token = t3
	//@ i.absState = s3
	//@ fold ProcessedHsMsg2Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.YT, s3)
	//@ fold i.Inv()
	success = true
	return
}

// @ preserves i != nil ==> i.Inv()
// @ preserves payload != nil ==> Mem(payload)
// @ ensures   success ==> msgData != nil
// @ ensures   msgData != nil ==> Mem(msgData)
func (i *Initiator) ProduceTransportMsg(payload []byte) (msgData []byte, success bool) {
	if i == nil || payload == nil { //argot:ignore diodon-dh-io-independence
		return
	}
	//@ unfold i.Inv()
	if i.initiatorState != HandshakeCompleted {
		//@ fold i.Inv()
		return
	}
	//@ t0, ridT, s0 := i.token, i.rid, i.absState
	//@ unfold ProcessedHsMsg2Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.YT, s0)

	//@ unfold io.P_Alice(t0, ridT, s0)
	//@ unfold io.phiRF_Alice_6(t0, ridT, s0)
	//@ assert io.e_InFact(t0, ridT)
	/*@ t1, payloadT := @*/
	PerformVirtualInputOperation(payload /*@, t0, ridT @*/)
	//@ s1 := s0 union mset[ft.Fact]{ ft.InFact_Alice(ridT, payloadT) }

	//@ unfold HandshakeCompletedPred(i.irKey, i.riKey, i.xT, i.YT)
	ciphertext, err := i.l.Encrypt(payload, i.irKey /*@, tm.kdf1(tm.exp(i.YT, i.xT)) @*/)
	//@ fold HandshakeCompletedPred(i.irKey, i.riKey, i.xT, i.YT)
	if err != nil {
		//@ i.token = t1
		//@ i.absState = s1
		//@ fold ProcessedHsMsg2Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.YT, s1)
		//@ fold i.Inv()
		return
	}

	tmpMsgData, err := i.l.MarshalTransportMsg(ciphertext)
	if err != nil {
		//@ i.token = t1
		//@ i.absState = s1
		//@ fold ProcessedHsMsg2Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.YT, s1)
		//@ fold i.Inv()
		return
	}
	//@ msgDataT := tm.tuple2(tm.integer32(TransMsgTag), tm.senc(payloadT, tm.kdf1(tm.exp(i.YT, i.xT))))
	//@ assert Abs(tmpMsgData) == by.gamma(msgDataT)

	//@ idAT := tm.integer32(i.idA)
	//@ idBT := tm.integer32(i.idB)
	//@ unfold io.P_Alice(t1, ridT, s1)
	//@ unfold io.phiR_Alice_2(t1, ridT, s1)
	/*@
	l := mset[ft.Fact]{ ft.St_Alice_2(ridT, idAT, idBT, i.skAT, i.skBT, i.xT, i.YT),
		ft.InFact_Alice(ridT, payloadT) }
	a := mset[cl.Claim]{
		cl.AliceSendLoop(tm.exp(i.YT, i.xT)),
		cl.AliceSendTransMsg(payloadT, tm.kdf1(tm.exp(i.YT, i.xT))) }
	r := mset[ft.Fact]{ ft.St_Alice_2(ridT, idAT, idBT, i.skAT, i.skBT, i.xT, i.YT),
		ft.OutFact_Alice(ridT, msgDataT) }
	@*/
	//@ assert io.e_Alice_sendMsg(t1, ridT, idAT, idBT, i.skAT, i.skBT, i.xT, i.YT, payloadT, l, a, r)
	//@ t2 := io.internBIO_e_Alice_sendMsg(t1, ridT, idAT, idBT, i.skAT, i.skBT, i.xT, i.YT, payloadT, l, a, r)
	//@ s2 := ft.U(l, r, s1)

	//@ unfold io.P_Alice(t2, ridT, s2)
	//@ unfold io.phiRG_Alice_4(t2, ridT, s2)
	//@ assert io.e_OutFact(t2, ridT, msgDataT)
	//@ ghost var t3 pl.Place
	msgData /*@, t3 @*/ = PerformVirtualOutputOperation(tmpMsgData /*@, t2, ridT, msgDataT @*/)
	//@ s3 := s2 setminus mset[ft.Fact]{ ft.OutFact_Alice(ridT, msgDataT) }

	//@ i.token = t3
	//@ i.absState = s3
	//@ fold ProcessedHsMsg2Pred(ridT, i.idA, i.idB, i.skAT, i.skBT, i.xT, i.YT, s3)
	//@ fold i.Inv()
	success = true
	return
}
