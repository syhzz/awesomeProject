package geecache

type ByteView struct {
	b []byte
}

func (B ByteView) Len() int {
	return len(B.b)
}
func (B ByteView) ByteSlice() []byte {
	return cloneBytes(B.b)
}

func (B ByteView) String() string {
	return string(B.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
