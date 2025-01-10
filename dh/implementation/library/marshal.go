package library

import binary "encoding/binary"
import errors "errors"
//@ import . "dh-gobra/verification/bytes"
//@ import by "dh-gobra/verification/utilbytes"

const Msg2Tag uint32 = 0
const Msg3Tag uint32 = 1
const TransMsgTag uint32 = 2

type Msg1 struct {
	X []byte
}

type Msg2 struct {
	IdA uint32
	IdB uint32
	X []byte
	Y []byte
}

type Msg3 struct {
	IdA uint32
	IdB uint32
	X []byte
	Y []byte
}

/*@
pred (m *Msg1) Mem() {
	acc(m) && Mem(m.X)
}

pred (m *Msg2) Mem() {
	acc(m) && Mem(m.X) && Mem(m.Y)
}

pred (m *Msg3) Mem() {
	acc(m) && Mem(m.Y) && Mem(m.X)
}
@*/

//@ trusted
//@ preserves acc(msg1.Mem(), 1/16)
//@ ensures err == nil ==> Mem(res) && res != nil
//@ ensures err == nil ==> Abs(res) == (unfolding acc(msg1.Mem(), 1/16) in Abs(msg1.X))
func (l *LibState) MarshalMsg1(msg1 *Msg1) (res []byte, err error) {
	res = make([]byte, len(msg1.X))
	copy(res, msg1.X)
	return res, nil
}

//@ trusted
//@ preserves acc(Mem(data), 1/16)
//@ ensures err == nil ==> res.Mem()
//@ ensures err == nil ==> Abs(data) == (unfolding acc(res.Mem(), 1/16) in by.tuple5B(by.integer32B(Msg2Tag), by.integer32B(res.IdB), by.integer32B(res.IdA), Abs(res.X), Abs(res.Y)))
func (l *LibState) UnmarshalMsg2(data []byte) (res *Msg2, err error) {
	if len(data) < 2 * DHHalfKeyLength + 12 {
		return nil, errors.New("msg2 is too short")
	}

	tag := binary.BigEndian.Uint32(data[:4])
	// note that idB occurs before idA in the 2nd message:
	idB := binary.BigEndian.Uint32(data[4:8])
	idA := binary.BigEndian.Uint32(data[8:12])
	X := make([]byte, DHHalfKeyLength)
	copy(X, data[12:(DHHalfKeyLength + 12)])
	Y := make([]byte, DHHalfKeyLength)
	copy(Y, data[(DHHalfKeyLength + 12):(2 * DHHalfKeyLength + 12)])

	if tag != Msg2Tag {
		return nil, errors.New("unexpected message tag in msg2")
	}

	return &Msg2{IdA: idA, IdB: idB, X: X, Y: Y}, nil
}

//@ trusted
//@ preserves acc(msg3.Mem(), 1/16)
//@ ensures err == nil ==> Mem(res)
//@ ensures err == nil ==> Abs(res) == (unfolding acc(msg3.Mem(), 1/16) in by.tuple5B(by.integer32B(Msg3Tag), by.integer32B(msg3.IdA), by.integer32B(msg3.IdB), Abs(msg3.Y), Abs(msg3.X)))
func (l *LibState) MarshalMsg3(msg3 *Msg3) (res []byte, err error) {
	res = make([]byte, 12)
	binary.BigEndian.PutUint32(res[:4], Msg3Tag)
	binary.BigEndian.PutUint32(res[4:8], msg3.IdA)
	binary.BigEndian.PutUint32(res[8:12], msg3.IdB)
	// note that Y occurs before X in the 3rd message:
	res = append(res, msg3.Y...)
	return append(res, msg3.X...), nil
}

//@ trusted
//@ preserves acc(Mem(ciphertext), 1/16)
//@ ensures err == nil ==> Mem(res) && res != nil
//@ ensures err == nil ==> Abs(res) == by.tuple2B(by.integer32B(TransMsgTag), Abs(ciphertext))
func (l *LibState) MarshalTransportMsg(ciphertext []byte) (res []byte, err error) {
	res = make([]byte, 4)
	binary.BigEndian.PutUint32(res[:4], TransMsgTag)
	return append(res, ciphertext...), nil
}

//@ trusted
//@ preserves acc(Mem(data), 1/16)
//@ ensures err == nil ==> Mem(ciphertext) && Abs(data) == by.tuple2B(by.integer32B(TransMsgTag), Abs(ciphertext))
func (l *LibState) UnmarshalTransportMsg(data []byte) (ciphertext []byte, err error) {
	if len(data) < 4 {
		return nil, errors.New("transport message is too short")
	}

	tag := binary.BigEndian.Uint32(data[:4])
	ciphertext = make([]byte, len(data) - 4)
	copy(ciphertext, data[4:])

	if tag != TransMsgTag {
		return nil, errors.New("unexpected message tag in transport message")
	}

	return
}
