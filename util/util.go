package util

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"replace_file/model"
	"strings"
)

// 项目路径配置文件名称
const ProjectPath string = "project-path"
// 替换项目文件的文件夹名称
const ReplacePath string = "replace-path"
// 对应的文件替换路径
const PackagePath string = "package-path"

// 获得配置文件内信息
func GetConfig(configPath string) (model.Path, error) {
	nilPathModel := model.Path{}
	open, err := os.Open(configPath)
	if err != nil {
		return nilPathModel, err
	}
	defer open.Close()
	path := model.New()
	reader := bufio.NewReader(open)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nilPathModel, err
			}
		}
		if len(line) <= 0 {
			continue
		}
		if isComments(line) {
			continue
		}
		err = handleConfig(path, line)
		if err != nil {
			return nilPathModel, err
		}
	}
	path.Integrate()
	return *path, nil
}

// 配置文件筛选注释
func isComments(line []uint8) bool {
	for _, v := range line {
		if v != 32 && v != 9 {
			if v == 35 {
				return true
			}
		}
	}
	return false
}

// 处理配置文件信息到 Path 结构体
// 方法私有,不可被外部调用
func handleConfig(path *model.Path, line []byte) error {
	lineStr := string(line)
	split := strings.Split(lineStr, "=")
	if len(split) < 2 {
		return errors.New("配置文件不正确, 请检查配置文件参数: " + lineStr)
	}
	switch split[0] {
	case ProjectPath:
		path.ProjectPath = split[1]
		break
	case ReplacePath:
		path.ReplacePath = split[1]
		break
	case PackagePath:
		path.PackagePath = strings.Split(split[1], ",")
		break
	}
	return nil
}

// 替换文件内容
func CopyFile(path model.Path) error {
	// k -> 替换文件
	// v -> 项目中需要替换的文件路径
	for k, v := range path.FullPath {
		log.Println(k, "开始替换!")
		kFile, err := os.Open(k)
		if err != nil {
			return err
		}
		kContent, err := ioutil.ReadAll(kFile)
		if err != nil {
			return err
		}
		vFile, err := os.OpenFile(v, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		_, err = vFile.Write(kContent)
		if err != nil {
			return err
		}
		if vFile != nil {
			defer vFile.Close()
		}
		if kFile != nil {
			defer kFile.Close()
		}
		log.Println(k, "替换成功!")
	}
	return nil
}
