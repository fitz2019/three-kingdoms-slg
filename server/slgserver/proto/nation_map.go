package proto

type Conf struct {
	Type     int8   `json:"type"`
	Level    int8   `json:"level"`
	Name     string `json:"name"`
	Wood     int    `json:"Wood"`
	Iron     int    `json:"iron"`
	Stone    int    `json:"stone"`
	Grain    int    `json:"grain"`
	Durable  int    `json:"durable"`  //耐久
	Defender int    `json:"defender"` //防御等级
}

type ConfigReq struct {
}

type ConfigRsp struct {
	Confs []Conf
}

type MapRoleBuild struct {
	RId        int    `json:"rid"`
	RNick      string `json:"RNick"` //角色昵称
	Name       string `json:"name"`
	UnionId    int    `json:"union_id"`   //联盟id
	UnionName  string `json:"union_name"` //联盟名字
	ParentId   int    `json:"parent_id"`  //上级id
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Type       int8   `json:"type"`
	Level      int8   `json:"level"`
	OPLevel    int8   `json:"op_level"`
	CurDurable int    `json:"cur_durable"`
	MaxDurable int    `json:"max_durable"`
	Defender   int    `json:"defender"`
	OccupyTime int64  `json:"occupy_time"`
	EndTime    int64  `json:"end_time"`    //建造完的时间
	GiveUpTime int64  `json:"giveUp_time"` //领地到了这个时间会被放弃
}

type ScanReq struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ScanRsp struct {
	MRBuilds []MapRoleBuild `json:"mr_builds"` //角色建筑，包含被占领的基础建筑
	MCBuilds []MapRoleCity  `json:"mc_builds"` //角色城市
	Armys    []Army         `json:"armys"`     //军队
}

type ScanBlockReq struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Length int `json:"length"`
}

type GiveUpReq struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type GiveUpRsp struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type BuildReq struct {
	X    int  `json:"x"`
	Y    int  `json:"y"`
	Type int8 `json:"type"`
}

type BuildRsp struct {
	X    int  `json:"x"`
	Y    int  `json:"y"`
	Type int8 `json:"type"`
}

type UpBuildReq struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type UpBuildRsp struct {
	X     int          `json:"x"`
	Y     int          `json:"y"`
	Build MapRoleBuild `json:"build"`
}

type DelBuildReq struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type DelBuildRsp struct {
	X     int          `json:"x"`
	Y     int          `json:"y"`
	Build MapRoleBuild `json:"build"`
}
