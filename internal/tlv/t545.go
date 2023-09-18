package tlv

import "github.com/grayrail000/MiraiGo/binary"

func T545(qimei []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x545)
		w.WriteBytesShort(qimei)
	})
}
