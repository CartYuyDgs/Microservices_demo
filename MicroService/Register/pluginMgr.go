package Register

import (
	"context"
	"log"
	"sync"
)

var (
	pluginMgr = &PluginMgr{
		plugins: make(map[string]Registry),
	}
)

//插件管理
type PluginMgr struct {
	plugins map[string]Registry
	lock    sync.Mutex
}

func (p *PluginMgr) registerPlugin(plugin Registry) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	_, ok := p.plugins[plugin.Name()]
	if ok {
		log.Fatalf("Error %v", ok)
		return
	}

	p.plugins[plugin.Name()] = plugin
	return
}

func (p *PluginMgr) initRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	plugin, ok := p.plugins[name]
	if ok {
		log.Fatalf("Error %v", ok)
		return
	}

	registry = plugin
	err = plugin.Init(ctx, opts...)
	return

}

func RegisterPlugin(registry Registry) (err error) {
	return pluginMgr.registerPlugin(registry)
}

func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return pluginMgr.initRegistry(ctx, name)
}
