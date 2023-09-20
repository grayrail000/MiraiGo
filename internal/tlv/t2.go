package tlv

import "github.com/grayrail000/MiraiGo/binary"

func T2(result string, Sign []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x02)
		pos := w.FillUInt16()
		w.WriteUInt16(0)
		w.WriteStringShort(result)
		w.WriteBytesShort(Sign)
		w.WriteUInt16At(pos, uint16(w.Len()-4))
	})
}
