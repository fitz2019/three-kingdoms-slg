package facility

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/llr104/slgserver/log"
	"go.uber.org/zap"
)

type conditions struct {
	Type  int `json:"type"`
	Level int `json:"level"`
}

type facility struct {
	Title      string       `json:"title"`
	Des        string       `json:"des"`
	Name       string       `json:"name"`
	Type       int8         `json:"type"`
	Additions  []int8       `json:"additions"`
	Conditions []conditions `json:"conditions"`
	Levels     []fLevel     `json:"levels"`
}

type NeedRes struct {
	Decree int `json:"decree"`
	Grain  int `json:"grain"`
	Wood   int `json:"wood"`
	Iron   int `json:"iron"`
	Stone  int `json:"stone"`
	Gold   int `json:"gold"`
}

type fLevel struct {
	Level  int     `json:"level"`
	Values []int   `json:"values"`
	Need   NeedRes `json:"need"`
	Time   int     `json:"time"` //升级需要的时间
}

func NewFacility(jsonName string) *facility {
	f := &facility{}
	f.load(jsonName)
	return f
}

func (this *facility) load(jsonName string) {

	jdata, err := ioutil.ReadFile(jsonName)
	if err != nil {
		log.DefaultLog.Error("facility load file error",
			zap.Error(err), zap.String("file", jsonName))
		os.Exit(0)
	}

	json.Unmarshal(jdata, this)

	fmt.Println(this)
}
