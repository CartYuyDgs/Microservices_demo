package main

import "fmt"

var generatorMgr *GeneratorMgr = &GeneratorMgr{genMap: make(map[string]Generator)}

type GeneratorMgr struct {
	genMap map[string]Generator
}

func Register(name string, gen Generator) (err error) {
	_, ok := generatorMgr.genMap[name]
	if ok {
		err = fmt.Errorf("generator %s is exits", name)
		return
	}
	fmt.Println("name: ", name)
	generatorMgr.genMap[name] = gen
	return
}

func (g *GeneratorMgr) Run(opt *Option) (err error) {
	for _, gen := range g.genMap {

		err := gen.Run(opt)
		if err != nil {
			return err
		}
	}
	return
}
