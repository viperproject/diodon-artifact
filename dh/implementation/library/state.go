package library

//@ import by "dh-gobra/verification/bytes"


type LibState struct {
	idA uint32
	idB uint32
	skA [64]byte
	pkB [32]byte
}

/*@
pred (l *LibState) Mem()

pred Mem(data []byte)

ghost
decreases
requires acc(Mem(b), 1/32)
pure func Abs(b []byte) (res by.Bytes)
@*/

//@ trusted
//@ ensures Mem(res)
func NewBytes(length int) (res []byte) {
	return make([]byte, length)
}

//@ trusted
//@ ensures err == nil ==> l.Mem()
func NewLibState(idA, idB uint32, privateKey [64]byte, peerPublicKey [32]byte) (l *LibState, err error) {
	state := &LibState {
		idA: idA,
		idB: idB,
		skA: privateKey,
		pkB: peerPublicKey,
	}
	return state, nil
}
