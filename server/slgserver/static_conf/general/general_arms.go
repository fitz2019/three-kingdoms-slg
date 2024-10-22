package general

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/llr104/slgserver/config"
	"github.com/llr104/slgserver/log"
	"go.uber.org/zap"
)

var GenArms Arms

type gArmsCondition struct {
	Level     int `json:"level"`
	StarLevel int `json:"star_lv"`
}

type gArmsCost struct {
	Gold int `json:"gold"`
}

type gArms struct {
	Id         int            `json:"id"`
	Name       string         `json:"name"`
	Condition  gArmsCondition `json:"condition"`
	ChangeCost gArmsCost      `json:"change_cost"`
	HarmRatio  []int          `json:"harm_ratio"`
}

type Arms struct {
	Title string  `json:"title"`
	Arms  []gArms `json:"arms"`
	AMap  map[int]gArms
}

func (this *Arms) Load() {
	jsonDir := config.File.MustValue("logic", "json_data", "../data/conf/")
	fileName := path.Join(jsonDir, "general", "general_arms.json")
	jdata, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.DefaultLog.Error("general load file error", zap.Error(err), zap.String("file", fileName))
		os.Exit(0)
	}

	json.Unmarshal(jdata, this)

	this.AMap = make(map[int]gArms)
	for _, v := range this.Arms {
		this.AMap[v.Id] = v
	}

	fmt.Println(this)
}

func (this *Arms) GetArm(id int) (gArms, error) {
	return this.AMap[id], nil
}

func (this *Arms) GetHarmRatio(attId, defId int) float64 {
	attArm, ok1 := this.AMap[attId]
	_, ok2 := this.AMap[defId]
	if ok1 && ok2 {
		return float64(attArm.HarmRatio[defId-1]) / 100.0
	} else {
		return 1.0
	}
}
