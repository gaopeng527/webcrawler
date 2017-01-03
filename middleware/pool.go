package middleware

type Entity interface {
	Id() uint32 // ID的获取方法
}

// 实体池的接口类型
type Pool interface {
	Take() (Entity, error)      // 取出实体
	Return(entity Entity) error // 归还实体
	Total() uint32              // 实体池的容量
	Used() uint32               // 实体池中已被使用的实体数量
}