package mutexpool

import "sync"

type (
	// пул мьютексов для конкурентной работы на основе ввода
	// смысл этой конструкции в том, чтоб выбирать мьютекс на основе набора данных (int в данном случае)
	// блокируя операция не глобально для всех, а только для набора горутин с одинаковым индeксом в пуле
	MutexPool struct {
		pool []*sync.Mutex
		m    *sync.Mutex
	}
)

func New(count int) *MutexPool {
	if count <= 0 {
		panic("zero or less pool size")
	}
	p := &MutexPool{pool: make([]*sync.Mutex, count), m: &sync.Mutex{}}
	for i := range p.pool {
		p.pool[i] = &sync.Mutex{}
	}
	return p
}

func sum(i ...int) int {
	var r int
	for _, v := range i {
		r += v
	}
	return r
}

func (ms *MutexPool) mutex(id ...int) *sync.Mutex {
	ms.m.Lock()
	defer ms.m.Unlock()
	i := sum(id...) % len(ms.pool)
	return ms.pool[i]
}

func (ms *MutexPool) Lock(id ...int) {
	ms.mutex(id...).Lock()
}

func (ms *MutexPool) Unlock(id ...int) {
	ms.mutex(id...).Unlock()
}

func (ms *MutexPool) LockAll() {
	for _, m := range ms.pool {
		m.Lock()
	}
}

func (ms MutexPool) UnlockAll() {
	for _, m := range ms.pool {
		m.Unlock()
	}
}
