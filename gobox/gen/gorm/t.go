package gorm

type Param struct {
	JsonAble     bool              `json:"json_able"`
	NullAble     bool              `json:"null_able"`
	GormAble     bool              `json:"gorm_able"`
	Package      string            `json:"package"`
	User         string            `json:"user"`
	Host         string            `json:"host"`
	Port         string            `json:"port"`
	Password     string            `json:"password"`
	Database     string            `json:"database"`
	Dst          string            `json:"dst"`
	Table2struct map[string]string `json:"table_2_struct"`
}
