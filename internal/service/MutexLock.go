package service

import "sync"

type MutexLock struct {
	ClientCount sync.Mutex
	BroadCast   sync.Mutex
}
