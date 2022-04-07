package i18n

import "testing"

func TestI18n(t *testing.T) {
	textMap, err := GetDictFromJsonFile("testdictfile.json")
	if err != nil {
		t.Error(err)
	} else {
		cn := I18n("user_unlogin", "zh-CN", textMap)
		en := I18n("user_unlogin", "en-US", textMap)
		t.Log(cn)
		t.Log(en)
	}
}
