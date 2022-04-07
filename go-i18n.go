package i18n

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var _i18nDict map[string]map[string]string

func Init(filepath string) (err error) {
	_i18nDict, err = GetDictFromJsonFile(filepath)
	return err
}

//国际化消息文本
//输入:msgcode string:消息码字符串
//language string:zh-CN,en-US等
//msgdict map[string]map[string]string:字典
//输出:string: 消息文本
func I18n(msgcode, language string, msgdicts ...map[string]map[string]string) string {
	for _, dict := range msgdicts {
		for k, v := range dict {
			_i18nDict[k] = v
		}
	}
	ecode, ok := _i18nDict[msgcode]
	if !ok {
		ecode = map[string]string{
			"zh-CN": "未定义的文本字典",
			"en-US": "Undefined text dictionary",
			"ru":    "неопределённый текстовый словарь", //俄语
			"es":    "Diccionario de texto indefinido",  //西班牙语
			"fr":    "Dictionnaire de texte non défini", //法语
		}
		//logs.Alert("[I18n]Undefined message code: %s", msgcode)
	}
	msg, ok := ecode[language]
	if !ok {
		msg = fmt.Sprintf("[I18n]Undefined languige type:%s", language)
		//logs.Alert(msg)
	}
	return msg
}

//将json字典转换为map字典
func JsonDictToMap(js string) (dict map[string]map[string]string) {
	json.Unmarshal([]byte(js), &dict)
	return
}

//从json文件获取字典数据
func GetDictFromJsonFile(filepath string) (map[string]map[string]string, error) {
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
