package ioc

import "sync"

type Container struct {
	instances map[string]interface{}
	mu        sync.Mutex
}

var container = &Container{
	instances: make(map[string]interface{}),
}

// Set 添加实例到容器
func Set(key string, instance interface{}) {
	container.mu.Lock()
	defer container.mu.Unlock()
	container.instances[key] = instance
}

// Get 从容器中获取实例
func Get(key string) interface{} {
	container.mu.Lock()
	defer container.mu.Unlock()
	return container.instances[key]
}
