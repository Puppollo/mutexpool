package mutexpool

import "sync"

type MutexPool struct{ pool []*sync.Mutex }

func NewMutexPool(count int) *MutexPool {
	return &MutexPool{pool: make([]*sync.Mutex, count)}
}

func sum(i ...int) int {
	var r int
	for _, v := range i {
		r += v
	}
	return r
}

func (ms *MutexPool) mutex(id ...int) *sync.Mutex {
	i := sum(id...) % len(ms.pool)
	if ms.pool[i] == nil {
		ms.pool[i] = &sync.Mutex{}
	}
	return ms.pool[i]
}

func (ms *MutexPool) Lock(id ...int) {
	ms.mutex(id...).Lock()
}

func (ms *MutexPool) Unlock(id ...int) {
	ms.mutex(id...).Unlock()
}

func (ms *MutexPool) LockAll() {
	for i := range ms.pool {
		ms.pool[i].Lock()
	}
}

func (ms MutexPool) UnlockAll() {
	for i := range ms.pool {
		ms.pool[i].Unlock()
	}
}
