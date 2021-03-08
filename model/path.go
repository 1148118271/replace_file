package model

import (
	"strings"
)

type Path struct {
	ProjectPath string
	ReplacePath string
	PackagePath []string
	FullPath map[string]string
}

// 创建配置文件路径对象
func New() *Path {
	return new(Path)
}

// 整合配置文件路径
func (p *Path) Integrate() {
	for _, v := range p.PackagePath {
		split := strings.Split(v, ":")
		if len(split) < 2 {
			panic("对应文件格式不正确,请检查格式!")
		}
		pp := joint(p.ProjectPath, split[0])
		rp := joint(p.ReplacePath, split[1])
		if p.FullPath == nil {
			p.FullPath = make(map[string]string)
		}
		p.FullPath[rp] = pp
	}
}

// 路径拼接
func joint(a string, b string) string {
	if a[(len(a) - 1)] != '/' && b[0] != '/' {
		return a + "/" + b
	}
	if a[(len(a) - 1)] == '/' && b[0] == '/' {
		return a + b[1:]
	}
	return a + b
}
