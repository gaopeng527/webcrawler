package middleware

// ID生成器的接口类型
type IdGenerator interface {
	GetUint32() uint32 // 获得一个uint32类型的ID
}
