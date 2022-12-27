package i18n

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var _i18nDict = make(map[string]map[string]string)

func Init(filepath string, msgdicts ...map[string]map[string]string) {
	//_i18nDict = make(map[string]map[string]string)
	for _, dict := range msgdicts {
		for k, v := range dict {
			_i18nDict[k] = v
		}
	}
	if jsondicts, err := getDictFromJsonFile(filepath); err == nil {
		for k, v := range jsondicts {
			_i18nDict[k] = v
		}
	}
}

//国际化消息文本
//输入:msgcode string:消息码字符串
//language string:zh-CN,en-US等
//msgdict map[string]map[string]string:字典
//输出:string: 消息文本
func I18n(msgcode, language string) string {
	ecode, ok := _i18nDict[msgcode]
	if !ok {
		return msgcode
	}
	msg, ok := ecode[language]
	if !ok {
		return msgcode
	}
	return msg
}

//将json字典转换为map字典
//func jsonDictToMap(js string) (dict map[string]map[string]string) {
//	json.Unmarshal([]byte(js), &dict)
//	return
//}

//从json文件获取字典数据
func getDictFromJsonFile(filepath string) (map[string]map[string]string, error) {
	if len(filepath) == 0 {
		return nil, fmt.Errorf("file path is empty")
	}
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	dict := make(map[string]map[string]string)
	if err := json.Unmarshal(data, &dict); err != nil {
		return nil, err
	}
	return dict, nil
}
