package tool

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path"
	"strings"
)

// ReadJson 读JSON并解析
func ReadJson[T any](path string, beforeParse func(content []byte) ([]byte, error)) (t T, err error) {
	// 检查文件是否存在
	if exist := IsFileExisted(path); !exist {
		return t, errors.New("file does not exist")
	}

	// 读取文件
	var content []byte
	if content, err = os.ReadFile(path); err != nil {
		return t, err
	}

	// 解析前处理
	if beforeParse != nil {
		if content, err = beforeParse(content); err != nil {
			return t, err
		}
	}

	// 初始化实例并解析JSON
	var inst T
	if err = json.Unmarshal(content, &inst); err != nil {
		return t, err
	}

	return inst, nil
}

// SaveJson 写JSON
func SaveJson[T any](t T, dstPath string, doIndent bool, afterParse func(content []byte) ([]byte, error)) error {
	// 规格化路径
	dstPath = strings.Replace(dstPath, "\\", "/", -1)

	// 检查文件是否存在
	dir := path.Dir(dstPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	JsonData, err := json.Marshal(t)
	if err != nil {
		return err
	}

	// json写入文件
	var prefix, indent string
	if doIndent {
		prefix, indent = "", "\t"
	} else {
		prefix, indent = "", ""
	}

	var str bytes.Buffer
	if err = json.Indent(&str, JsonData, prefix, indent); err != nil {
		return err
	}

	content := str.Bytes()

	// 写文件前处理
	if afterParse != nil {
		if content, err = afterParse(content); err != nil {
			return err
		}
	}

	// 写文件
	if err = os.WriteFile(dstPath, content, 0666); err != nil {
		return err
	}

	return nil
}
