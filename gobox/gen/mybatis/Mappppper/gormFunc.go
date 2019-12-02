package Mappppper

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func init() {
	GormDB, e := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/hress2?charset=utf8mb4&collation=utf8mb4_bin&loc=Local&parseTime=true")
	if e != nil {
		panic(e)
	} else {
		DB = GormDB
	}
}

var (
	DB                 *gorm.DB
	LogMailmonDB       = LogMailmonOpe{DB: DB}
	LogMalurlDB        = LogMalurlOpe{DB: DB}
	ClntsHwinfoDB      = ClntsHwinfoOpe{DB: DB}
	LeakrepairClntkbDB = LeakrepairClntkbOpe{DB: DB}
	LogScanDB          = LogScanOpe{DB: DB}
	ClntsDB            = ClntsOpe{DB: DB}
	ClntsSchedDB       = ClntsSchedOpe{DB: DB}
	LogIpprotoDB       = LogIpprotoOpe{DB: DB}
	QuarantineDB       = QuarantineOpe{DB: DB}
	TaskLimitDB        = TaskLimitOpe{DB: DB}
	UdiskSecDB         = UdiskSecOpe{DB: DB}
	DistrDB            = DistrOpe{DB: DB}
	LeakrepairClntsDB  = LeakrepairClntsOpe{DB: DB}
	LogInstmonDB       = LogInstmonOpe{DB: DB}
	LogIpattackDB      = LogIpattackOpe{DB: DB}
	LogRemoteprotDB    = LogRemoteprotOpe{DB: DB}
	UserlogDB          = UserlogOpe{DB: DB}
	UserloginDB        = UserloginOpe{DB: DB}
	AssetsItemsDB      = AssetsItemsOpe{DB: DB}
	HwinfoHistoryDB    = HwinfoHistoryOpe{DB: DB}
	ClntsSwinfoDB      = ClntsSwinfoOpe{DB: DB}
	LogBehavDB         = LogBehavOpe{DB: DB}
	LogDlmonDB         = LogDlmonOpe{DB: DB}
	LogSysprotDB       = LogSysprotOpe{DB: DB}
	SchedInfoDB        = SchedInfoOpe{DB: DB}
	TaskStatDB         = TaskStatOpe{DB: DB}
	ClntsOsinfoDB      = ClntsOsinfoOpe{DB: DB}
	ClntsStatDB        = ClntsStatOpe{DB: DB}
	UserinfoDB         = UserinfoOpe{DB: DB}
	NodesDB            = NodesOpe{DB: DB}
	PoliciesDB         = PoliciesOpe{DB: DB}
	LogLeakrepairDB    = LogLeakrepairOpe{DB: DB}
	LogStatDB          = LogStatOpe{DB: DB}
	GroupsDB           = GroupsOpe{DB: DB}
	LeakrepairKbsaveDB = LeakrepairKbsaveOpe{DB: DB}
	LogIntrusionDB     = LogIntrusionOpe{DB: DB}
	LogUdiskmonDB      = LogUdiskmonOpe{DB: DB}
	PoliciesDetailDB   = PoliciesDetailOpe{DB: DB}
	SysconfDB          = SysconfOpe{DB: DB}
	ClntsAssetsDB      = ClntsAssetsOpe{DB: DB}
	ClntsTagsDB        = ClntsTagsOpe{DB: DB}
	WhitelistDB        = WhitelistOpe{DB: DB}
	LogCronDB          = LogCronOpe{DB: DB}
	LogIpblacklistDB   = LogIpblacklistOpe{DB: DB}
	LogRemoteDB        = LogRemoteOpe{DB: DB}
	LogUpgradeDB       = LogUpgradeOpe{DB: DB}
	SwinfoDB           = SwinfoOpe{DB: DB}
	TagsDictDB         = TagsDictOpe{DB: DB}
	ClntsOlDB          = ClntsOlOpe{DB: DB}
	LeakrepairKbinfoDB = LeakrepairKbinfoOpe{DB: DB}
	TaskInfoDB         = TaskInfoOpe{DB: DB}
)

type PoliciesOpe struct {
	DB *gorm.DB
}

func (a PoliciesOpe) GetByID(id int) (Policies, error) {
	obj := Policies{}
	err := a.DB.Model(&Policies{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a PoliciesOpe) ListByLimit(o, l int) ([]Policies, error) {
	var list []Policies
	err := a.DB.Model(&Policies{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a PoliciesOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Policies{}).Find(&[]Policies{}).Count(&count).Error
	return count, err
}
func (a PoliciesOpe) UpdateById(up Policies) error {
	return a.DB.Model(&Policies{}).Update(up).Error
}
func (a PoliciesOpe) DeleteById(id int) error {
	return a.DB.Model(&Policies{}).Where("`id`=?", id).Delete(nil).Error
}
func (a PoliciesOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Policies{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------policies------------------------//
type LogLeakrepairOpe struct {
	DB *gorm.DB
}

func (a LogLeakrepairOpe) GetByID(id int) (LogLeakrepair, error) {
	obj := LogLeakrepair{}
	err := a.DB.Model(&LogLeakrepair{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogLeakrepairOpe) ListByLimit(o, l int) ([]LogLeakrepair, error) {
	var list []LogLeakrepair
	err := a.DB.Model(&LogLeakrepair{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogLeakrepairOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogLeakrepair{}).Find(&[]LogLeakrepair{}).Count(&count).Error
	return count, err
}
func (a LogLeakrepairOpe) UpdateById(up LogLeakrepair) error {
	return a.DB.Model(&LogLeakrepair{}).Update(up).Error
}
func (a LogLeakrepairOpe) DeleteById(id int) error {
	return a.DB.Model(&LogLeakrepair{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogLeakrepairOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogLeakrepair{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_leakrepair------------------------//
type LogStatOpe struct {
	DB *gorm.DB
}

func (a LogStatOpe) GetByID(id int) (LogStat, error) {
	obj := LogStat{}
	err := a.DB.Model(&LogStat{}).Where("`ts`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogStatOpe) ListByLimit(o, l int) ([]LogStat, error) {
	var list []LogStat
	err := a.DB.Model(&LogStat{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogStatOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogStat{}).Find(&[]LogStat{}).Count(&count).Error
	return count, err
}
func (a LogStatOpe) UpdateById(up LogStat) error {
	return a.DB.Model(&LogStat{}).Update(up).Error
}
func (a LogStatOpe) DeleteById(id int) error {
	return a.DB.Model(&LogStat{}).Where("`ts`=?", id).Delete(nil).Error
}
func (a LogStatOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogStat{}).Where("`ts` in (?)", id).Delete(nil).Error
}

//------------------------log_stat------------------------//
type NodesOpe struct {
	DB *gorm.DB
}

func (a NodesOpe) GetByID(id int) (Nodes, error) {
	obj := Nodes{}
	err := a.DB.Model(&Nodes{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a NodesOpe) ListByLimit(o, l int) ([]Nodes, error) {
	var list []Nodes
	err := a.DB.Model(&Nodes{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a NodesOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Nodes{}).Find(&[]Nodes{}).Count(&count).Error
	return count, err
}
func (a NodesOpe) UpdateById(up Nodes) error {
	return a.DB.Model(&Nodes{}).Update(up).Error
}
func (a NodesOpe) DeleteById(id int) error {
	return a.DB.Model(&Nodes{}).Where("`id`=?", id).Delete(nil).Error
}
func (a NodesOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Nodes{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------nodes------------------------//
type LeakrepairKbsaveOpe struct {
	DB *gorm.DB
}

func (a LeakrepairKbsaveOpe) GetByID(id int) (LeakrepairKbsave, error) {
	obj := LeakrepairKbsave{}
	err := a.DB.Model(&LeakrepairKbsave{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LeakrepairKbsaveOpe) ListByLimit(o, l int) ([]LeakrepairKbsave, error) {
	var list []LeakrepairKbsave
	err := a.DB.Model(&LeakrepairKbsave{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LeakrepairKbsaveOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LeakrepairKbsave{}).Find(&[]LeakrepairKbsave{}).Count(&count).Error
	return count, err
}
func (a LeakrepairKbsaveOpe) UpdateById(up LeakrepairKbsave) error {
	return a.DB.Model(&LeakrepairKbsave{}).Update(up).Error
}
func (a LeakrepairKbsaveOpe) DeleteById(id int) error {
	return a.DB.Model(&LeakrepairKbsave{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LeakrepairKbsaveOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LeakrepairKbsave{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------leakrepair_kbsave------------------------//
type LogIntrusionOpe struct {
	DB *gorm.DB
}

func (a LogIntrusionOpe) GetByID(id int) (LogIntrusion, error) {
	obj := LogIntrusion{}
	err := a.DB.Model(&LogIntrusion{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogIntrusionOpe) ListByLimit(o, l int) ([]LogIntrusion, error) {
	var list []LogIntrusion
	err := a.DB.Model(&LogIntrusion{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogIntrusionOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogIntrusion{}).Find(&[]LogIntrusion{}).Count(&count).Error
	return count, err
}
func (a LogIntrusionOpe) UpdateById(up LogIntrusion) error {
	return a.DB.Model(&LogIntrusion{}).Update(up).Error
}
func (a LogIntrusionOpe) DeleteById(id int) error {
	return a.DB.Model(&LogIntrusion{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogIntrusionOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogIntrusion{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_intrusion------------------------//
type LogUdiskmonOpe struct {
	DB *gorm.DB
}

func (a LogUdiskmonOpe) GetByID(id int) (LogUdiskmon, error) {
	obj := LogUdiskmon{}
	err := a.DB.Model(&LogUdiskmon{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogUdiskmonOpe) ListByLimit(o, l int) ([]LogUdiskmon, error) {
	var list []LogUdiskmon
	err := a.DB.Model(&LogUdiskmon{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogUdiskmonOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogUdiskmon{}).Find(&[]LogUdiskmon{}).Count(&count).Error
	return count, err
}
func (a LogUdiskmonOpe) UpdateById(up LogUdiskmon) error {
	return a.DB.Model(&LogUdiskmon{}).Update(up).Error
}
func (a LogUdiskmonOpe) DeleteById(id int) error {
	return a.DB.Model(&LogUdiskmon{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogUdiskmonOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogUdiskmon{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_udiskmon------------------------//
type PoliciesDetailOpe struct {
	DB *gorm.DB
}

func (a PoliciesDetailOpe) GetByID(id int) (PoliciesDetail, error) {
	obj := PoliciesDetail{}
	err := a.DB.Model(&PoliciesDetail{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a PoliciesDetailOpe) ListByLimit(o, l int) ([]PoliciesDetail, error) {
	var list []PoliciesDetail
	err := a.DB.Model(&PoliciesDetail{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a PoliciesDetailOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&PoliciesDetail{}).Find(&[]PoliciesDetail{}).Count(&count).Error
	return count, err
}
func (a PoliciesDetailOpe) UpdateById(up PoliciesDetail) error {
	return a.DB.Model(&PoliciesDetail{}).Update(up).Error
}
func (a PoliciesDetailOpe) DeleteById(id int) error {
	return a.DB.Model(&PoliciesDetail{}).Where("`id`=?", id).Delete(nil).Error
}
func (a PoliciesDetailOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&PoliciesDetail{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------policies_detail------------------------//
type SysconfOpe struct {
	DB *gorm.DB
}

func (a SysconfOpe) GetByID(id int) (Sysconf, error) {
	obj := Sysconf{}
	err := a.DB.Model(&Sysconf{}).Where("`name`=?", id).Find(&obj).Error
	return obj, err
}

func (a SysconfOpe) ListByLimit(o, l int) ([]Sysconf, error) {
	var list []Sysconf
	err := a.DB.Model(&Sysconf{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a SysconfOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Sysconf{}).Find(&[]Sysconf{}).Count(&count).Error
	return count, err
}
func (a SysconfOpe) UpdateById(up Sysconf) error {
	return a.DB.Model(&Sysconf{}).Update(up).Error
}
func (a SysconfOpe) DeleteById(id int) error {
	return a.DB.Model(&Sysconf{}).Where("`name`=?", id).Delete(nil).Error
}
func (a SysconfOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Sysconf{}).Where("`name` in (?)", id).Delete(nil).Error
}

//------------------------sysconf------------------------//
type ClntsAssetsOpe struct {
	DB *gorm.DB
}

func (a ClntsAssetsOpe) GetByID(id int) (ClntsAssets, error) {
	obj := ClntsAssets{}
	err := a.DB.Model(&ClntsAssets{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a ClntsAssetsOpe) ListByLimit(o, l int) ([]ClntsAssets, error) {
	var list []ClntsAssets
	err := a.DB.Model(&ClntsAssets{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a ClntsAssetsOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&ClntsAssets{}).Find(&[]ClntsAssets{}).Count(&count).Error
	return count, err
}
func (a ClntsAssetsOpe) UpdateById(up ClntsAssets) error {
	return a.DB.Model(&ClntsAssets{}).Update(up).Error
}
func (a ClntsAssetsOpe) DeleteById(id int) error {
	return a.DB.Model(&ClntsAssets{}).Where("`id`=?", id).Delete(nil).Error
}
func (a ClntsAssetsOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&ClntsAssets{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------clnts_assets------------------------//
type ClntsTagsOpe struct {
	DB *gorm.DB
}

func (a ClntsTagsOpe) GetByID(id int) (ClntsTags, error) {
	obj := ClntsTags{}
	err := a.DB.Model(&ClntsTags{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a ClntsTagsOpe) ListByLimit(o, l int) ([]ClntsTags, error) {
	var list []ClntsTags
	err := a.DB.Model(&ClntsTags{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a ClntsTagsOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&ClntsTags{}).Find(&[]ClntsTags{}).Count(&count).Error
	return count, err
}
func (a ClntsTagsOpe) UpdateById(up ClntsTags) error {
	return a.DB.Model(&ClntsTags{}).Update(up).Error
}
func (a ClntsTagsOpe) DeleteById(id int) error {
	return a.DB.Model(&ClntsTags{}).Where("`id`=?", id).Delete(nil).Error
}
func (a ClntsTagsOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&ClntsTags{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------clnts_tags------------------------//
type GroupsOpe struct {
	DB *gorm.DB
}

func (a GroupsOpe) GetByID(id int) (Groups, error) {
	obj := Groups{}
	err := a.DB.Model(&Groups{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a GroupsOpe) ListByLimit(o, l int) ([]Groups, error) {
	var list []Groups
	err := a.DB.Model(&Groups{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a GroupsOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Groups{}).Find(&[]Groups{}).Count(&count).Error
	return count, err
}
func (a GroupsOpe) UpdateById(up Groups) error {
	return a.DB.Model(&Groups{}).Update(up).Error
}
func (a GroupsOpe) DeleteById(id int) error {
	return a.DB.Model(&Groups{}).Where("`id`=?", id).Delete(nil).Error
}
func (a GroupsOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Groups{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------groups------------------------//
type WhitelistOpe struct {
	DB *gorm.DB
}

func (a WhitelistOpe) GetByID(id int) (Whitelist, error) {
	obj := Whitelist{}
	err := a.DB.Model(&Whitelist{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a WhitelistOpe) ListByLimit(o, l int) ([]Whitelist, error) {
	var list []Whitelist
	err := a.DB.Model(&Whitelist{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a WhitelistOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Whitelist{}).Find(&[]Whitelist{}).Count(&count).Error
	return count, err
}
func (a WhitelistOpe) UpdateById(up Whitelist) error {
	return a.DB.Model(&Whitelist{}).Update(up).Error
}
func (a WhitelistOpe) DeleteById(id int) error {
	return a.DB.Model(&Whitelist{}).Where("`id`=?", id).Delete(nil).Error
}
func (a WhitelistOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Whitelist{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------whitelist------------------------//
type LogIpblacklistOpe struct {
	DB *gorm.DB
}

func (a LogIpblacklistOpe) GetByID(id int) (LogIpblacklist, error) {
	obj := LogIpblacklist{}
	err := a.DB.Model(&LogIpblacklist{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogIpblacklistOpe) ListByLimit(o, l int) ([]LogIpblacklist, error) {
	var list []LogIpblacklist
	err := a.DB.Model(&LogIpblacklist{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogIpblacklistOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogIpblacklist{}).Find(&[]LogIpblacklist{}).Count(&count).Error
	return count, err
}
func (a LogIpblacklistOpe) UpdateById(up LogIpblacklist) error {
	return a.DB.Model(&LogIpblacklist{}).Update(up).Error
}
func (a LogIpblacklistOpe) DeleteById(id int) error {
	return a.DB.Model(&LogIpblacklist{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogIpblacklistOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogIpblacklist{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_ipblacklist------------------------//
type LogRemoteOpe struct {
	DB *gorm.DB
}

func (a LogRemoteOpe) GetByID(id int) (LogRemote, error) {
	obj := LogRemote{}
	err := a.DB.Model(&LogRemote{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogRemoteOpe) ListByLimit(o, l int) ([]LogRemote, error) {
	var list []LogRemote
	err := a.DB.Model(&LogRemote{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogRemoteOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogRemote{}).Find(&[]LogRemote{}).Count(&count).Error
	return count, err
}
func (a LogRemoteOpe) UpdateById(up LogRemote) error {
	return a.DB.Model(&LogRemote{}).Update(up).Error
}
func (a LogRemoteOpe) DeleteById(id int) error {
	return a.DB.Model(&LogRemote{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogRemoteOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogRemote{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_remote------------------------//
type LogUpgradeOpe struct {
	DB *gorm.DB
}

func (a LogUpgradeOpe) GetByID(id int) (LogUpgrade, error) {
	obj := LogUpgrade{}
	err := a.DB.Model(&LogUpgrade{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogUpgradeOpe) ListByLimit(o, l int) ([]LogUpgrade, error) {
	var list []LogUpgrade
	err := a.DB.Model(&LogUpgrade{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogUpgradeOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogUpgrade{}).Find(&[]LogUpgrade{}).Count(&count).Error
	return count, err
}
func (a LogUpgradeOpe) UpdateById(up LogUpgrade) error {
	return a.DB.Model(&LogUpgrade{}).Update(up).Error
}
func (a LogUpgradeOpe) DeleteById(id int) error {
	return a.DB.Model(&LogUpgrade{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogUpgradeOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogUpgrade{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_upgrade------------------------//
type SwinfoOpe struct {
	DB *gorm.DB
}

func (a SwinfoOpe) GetByID(id int) (Swinfo, error) {
	obj := Swinfo{}
	err := a.DB.Model(&Swinfo{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a SwinfoOpe) ListByLimit(o, l int) ([]Swinfo, error) {
	var list []Swinfo
	err := a.DB.Model(&Swinfo{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a SwinfoOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Swinfo{}).Find(&[]Swinfo{}).Count(&count).Error
	return count, err
}
func (a SwinfoOpe) UpdateById(up Swinfo) error {
	return a.DB.Model(&Swinfo{}).Update(up).Error
}
func (a SwinfoOpe) DeleteById(id int) error {
	return a.DB.Model(&Swinfo{}).Where("`id`=?", id).Delete(nil).Error
}
func (a SwinfoOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Swinfo{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------swinfo------------------------//
type TagsDictOpe struct {
	DB *gorm.DB
}

func (a TagsDictOpe) GetByID(id int) (TagsDict, error) {
	obj := TagsDict{}
	err := a.DB.Model(&TagsDict{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a TagsDictOpe) ListByLimit(o, l int) ([]TagsDict, error) {
	var list []TagsDict
	err := a.DB.Model(&TagsDict{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a TagsDictOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&TagsDict{}).Find(&[]TagsDict{}).Count(&count).Error
	return count, err
}
func (a TagsDictOpe) UpdateById(up TagsDict) error {
	return a.DB.Model(&TagsDict{}).Update(up).Error
}
func (a TagsDictOpe) DeleteById(id int) error {
	return a.DB.Model(&TagsDict{}).Where("`id`=?", id).Delete(nil).Error
}
func (a TagsDictOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&TagsDict{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------tags_dict------------------------//
type ClntsOlOpe struct {
	DB *gorm.DB
}

func (a ClntsOlOpe) GetByID(id int) (ClntsOl, error) {
	obj := ClntsOl{}
	err := a.DB.Model(&ClntsOl{}).Where("`c_id`=?", id).Find(&obj).Error
	return obj, err
}

func (a ClntsOlOpe) ListByLimit(o, l int) ([]ClntsOl, error) {
	var list []ClntsOl
	err := a.DB.Model(&ClntsOl{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a ClntsOlOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&ClntsOl{}).Find(&[]ClntsOl{}).Count(&count).Error
	return count, err
}
func (a ClntsOlOpe) UpdateById(up ClntsOl) error {
	return a.DB.Model(&ClntsOl{}).Update(up).Error
}
func (a ClntsOlOpe) DeleteById(id int) error {
	return a.DB.Model(&ClntsOl{}).Where("`c_id`=?", id).Delete(nil).Error
}
func (a ClntsOlOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&ClntsOl{}).Where("`c_id` in (?)", id).Delete(nil).Error
}

//------------------------clnts_ol------------------------//
type LeakrepairKbinfoOpe struct {
	DB *gorm.DB
}

func (a LeakrepairKbinfoOpe) GetByID(id int) (LeakrepairKbinfo, error) {
	obj := LeakrepairKbinfo{}
	err := a.DB.Model(&LeakrepairKbinfo{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LeakrepairKbinfoOpe) ListByLimit(o, l int) ([]LeakrepairKbinfo, error) {
	var list []LeakrepairKbinfo
	err := a.DB.Model(&LeakrepairKbinfo{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LeakrepairKbinfoOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LeakrepairKbinfo{}).Find(&[]LeakrepairKbinfo{}).Count(&count).Error
	return count, err
}
func (a LeakrepairKbinfoOpe) UpdateById(up LeakrepairKbinfo) error {
	return a.DB.Model(&LeakrepairKbinfo{}).Update(up).Error
}
func (a LeakrepairKbinfoOpe) DeleteById(id int) error {
	return a.DB.Model(&LeakrepairKbinfo{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LeakrepairKbinfoOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LeakrepairKbinfo{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------leakrepair_kbinfo------------------------//
type LogCronOpe struct {
	DB *gorm.DB
}

func (a LogCronOpe) GetByID(id int) (LogCron, error) {
	obj := LogCron{}
	err := a.DB.Model(&LogCron{}).Where("`cid`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogCronOpe) ListByLimit(o, l int) ([]LogCron, error) {
	var list []LogCron
	err := a.DB.Model(&LogCron{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogCronOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogCron{}).Find(&[]LogCron{}).Count(&count).Error
	return count, err
}
func (a LogCronOpe) UpdateById(up LogCron) error {
	return a.DB.Model(&LogCron{}).Update(up).Error
}
func (a LogCronOpe) DeleteById(id int) error {
	return a.DB.Model(&LogCron{}).Where("`cid`=?", id).Delete(nil).Error
}
func (a LogCronOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogCron{}).Where("`cid` in (?)", id).Delete(nil).Error
}

//------------------------log_cron------------------------//
type TaskInfoOpe struct {
	DB *gorm.DB
}

func (a TaskInfoOpe) GetByID(id int) (TaskInfo, error) {
	obj := TaskInfo{}
	err := a.DB.Model(&TaskInfo{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a TaskInfoOpe) ListByLimit(o, l int) ([]TaskInfo, error) {
	var list []TaskInfo
	err := a.DB.Model(&TaskInfo{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a TaskInfoOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&TaskInfo{}).Find(&[]TaskInfo{}).Count(&count).Error
	return count, err
}
func (a TaskInfoOpe) UpdateById(up TaskInfo) error {
	return a.DB.Model(&TaskInfo{}).Update(up).Error
}
func (a TaskInfoOpe) DeleteById(id int) error {
	return a.DB.Model(&TaskInfo{}).Where("`id`=?", id).Delete(nil).Error
}
func (a TaskInfoOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&TaskInfo{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------task_info------------------------//
type LogMalurlOpe struct {
	DB *gorm.DB
}

func (a LogMalurlOpe) GetByID(id int) (LogMalurl, error) {
	obj := LogMalurl{}
	err := a.DB.Model(&LogMalurl{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogMalurlOpe) ListByLimit(o, l int) ([]LogMalurl, error) {
	var list []LogMalurl
	err := a.DB.Model(&LogMalurl{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogMalurlOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogMalurl{}).Find(&[]LogMalurl{}).Count(&count).Error
	return count, err
}
func (a LogMalurlOpe) UpdateById(up LogMalurl) error {
	return a.DB.Model(&LogMalurl{}).Update(up).Error
}
func (a LogMalurlOpe) DeleteById(id int) error {
	return a.DB.Model(&LogMalurl{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogMalurlOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogMalurl{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_malurl------------------------//
type ClntsHwinfoOpe struct {
	DB *gorm.DB
}

func (a ClntsHwinfoOpe) GetByID(id int) (ClntsHwinfo, error) {
	obj := ClntsHwinfo{}
	err := a.DB.Model(&ClntsHwinfo{}).Where("`c_id`=?", id).Find(&obj).Error
	return obj, err
}

func (a ClntsHwinfoOpe) ListByLimit(o, l int) ([]ClntsHwinfo, error) {
	var list []ClntsHwinfo
	err := a.DB.Model(&ClntsHwinfo{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a ClntsHwinfoOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&ClntsHwinfo{}).Find(&[]ClntsHwinfo{}).Count(&count).Error
	return count, err
}
func (a ClntsHwinfoOpe) UpdateById(up ClntsHwinfo) error {
	return a.DB.Model(&ClntsHwinfo{}).Update(up).Error
}
func (a ClntsHwinfoOpe) DeleteById(id int) error {
	return a.DB.Model(&ClntsHwinfo{}).Where("`c_id`=?", id).Delete(nil).Error
}
func (a ClntsHwinfoOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&ClntsHwinfo{}).Where("`c_id` in (?)", id).Delete(nil).Error
}

//------------------------clnts_hwinfo------------------------//
type LeakrepairClntkbOpe struct {
	DB *gorm.DB
}

func (a LeakrepairClntkbOpe) GetByID(id int) (LeakrepairClntkb, error) {
	obj := LeakrepairClntkb{}
	err := a.DB.Model(&LeakrepairClntkb{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LeakrepairClntkbOpe) ListByLimit(o, l int) ([]LeakrepairClntkb, error) {
	var list []LeakrepairClntkb
	err := a.DB.Model(&LeakrepairClntkb{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LeakrepairClntkbOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LeakrepairClntkb{}).Find(&[]LeakrepairClntkb{}).Count(&count).Error
	return count, err
}
func (a LeakrepairClntkbOpe) UpdateById(up LeakrepairClntkb) error {
	return a.DB.Model(&LeakrepairClntkb{}).Update(up).Error
}
func (a LeakrepairClntkbOpe) DeleteById(id int) error {
	return a.DB.Model(&LeakrepairClntkb{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LeakrepairClntkbOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LeakrepairClntkb{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------leakrepair_clntkb------------------------//
type LogMailmonOpe struct {
	DB *gorm.DB
}

func (a LogMailmonOpe) GetByID(id int) (LogMailmon, error) {
	obj := LogMailmon{}
	err := a.DB.Model(&LogMailmon{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogMailmonOpe) ListByLimit(o, l int) ([]LogMailmon, error) {
	var list []LogMailmon
	err := a.DB.Model(&LogMailmon{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogMailmonOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogMailmon{}).Find(&[]LogMailmon{}).Count(&count).Error
	return count, err
}
func (a LogMailmonOpe) UpdateById(up LogMailmon) error {
	return a.DB.Model(&LogMailmon{}).Update(up).Error
}
func (a LogMailmonOpe) DeleteById(id int) error {
	return a.DB.Model(&LogMailmon{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogMailmonOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogMailmon{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_mailmon------------------------//
type ClntsOpe struct {
	DB *gorm.DB
}

func (a ClntsOpe) GetByID(id int) (Clnts, error) {
	obj := Clnts{}
	err := a.DB.Model(&Clnts{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a ClntsOpe) ListByLimit(o, l int) ([]Clnts, error) {
	var list []Clnts
	err := a.DB.Model(&Clnts{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a ClntsOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Clnts{}).Find(&[]Clnts{}).Count(&count).Error
	return count, err
}
func (a ClntsOpe) UpdateById(up Clnts) error {
	return a.DB.Model(&Clnts{}).Update(up).Error
}
func (a ClntsOpe) DeleteById(id int) error {
	return a.DB.Model(&Clnts{}).Where("`id`=?", id).Delete(nil).Error
}
func (a ClntsOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Clnts{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------clnts------------------------//
type ClntsSchedOpe struct {
	DB *gorm.DB
}

func (a ClntsSchedOpe) GetByID(id int) (ClntsSched, error) {
	obj := ClntsSched{}
	err := a.DB.Model(&ClntsSched{}).Where("`c_id`=?", id).Find(&obj).Error
	return obj, err
}

func (a ClntsSchedOpe) ListByLimit(o, l int) ([]ClntsSched, error) {
	var list []ClntsSched
	err := a.DB.Model(&ClntsSched{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a ClntsSchedOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&ClntsSched{}).Find(&[]ClntsSched{}).Count(&count).Error
	return count, err
}
func (a ClntsSchedOpe) UpdateById(up ClntsSched) error {
	return a.DB.Model(&ClntsSched{}).Update(up).Error
}
func (a ClntsSchedOpe) DeleteById(id int) error {
	return a.DB.Model(&ClntsSched{}).Where("`c_id`=?", id).Delete(nil).Error
}
func (a ClntsSchedOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&ClntsSched{}).Where("`c_id` in (?)", id).Delete(nil).Error
}

//------------------------clnts_sched------------------------//
type LogScanOpe struct {
	DB *gorm.DB
}

func (a LogScanOpe) GetByID(id int) (LogScan, error) {
	obj := LogScan{}
	err := a.DB.Model(&LogScan{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogScanOpe) ListByLimit(o, l int) ([]LogScan, error) {
	var list []LogScan
	err := a.DB.Model(&LogScan{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogScanOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogScan{}).Find(&[]LogScan{}).Count(&count).Error
	return count, err
}
func (a LogScanOpe) UpdateById(up LogScan) error {
	return a.DB.Model(&LogScan{}).Update(up).Error
}
func (a LogScanOpe) DeleteById(id int) error {
	return a.DB.Model(&LogScan{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogScanOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogScan{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_scan------------------------//
type QuarantineOpe struct {
	DB *gorm.DB
}

func (a QuarantineOpe) GetByID(id int) (Quarantine, error) {
	obj := Quarantine{}
	err := a.DB.Model(&Quarantine{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a QuarantineOpe) ListByLimit(o, l int) ([]Quarantine, error) {
	var list []Quarantine
	err := a.DB.Model(&Quarantine{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a QuarantineOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Quarantine{}).Find(&[]Quarantine{}).Count(&count).Error
	return count, err
}
func (a QuarantineOpe) UpdateById(up Quarantine) error {
	return a.DB.Model(&Quarantine{}).Update(up).Error
}
func (a QuarantineOpe) DeleteById(id int) error {
	return a.DB.Model(&Quarantine{}).Where("`id`=?", id).Delete(nil).Error
}
func (a QuarantineOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Quarantine{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------quarantine------------------------//
type TaskLimitOpe struct {
	DB *gorm.DB
}

func (a TaskLimitOpe) GetByID(id int) (TaskLimit, error) {
	obj := TaskLimit{}
	err := a.DB.Model(&TaskLimit{}).Where("`c_id`=?", id).Find(&obj).Error
	return obj, err
}

func (a TaskLimitOpe) ListByLimit(o, l int) ([]TaskLimit, error) {
	var list []TaskLimit
	err := a.DB.Model(&TaskLimit{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a TaskLimitOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&TaskLimit{}).Find(&[]TaskLimit{}).Count(&count).Error
	return count, err
}
func (a TaskLimitOpe) UpdateById(up TaskLimit) error {
	return a.DB.Model(&TaskLimit{}).Update(up).Error
}
func (a TaskLimitOpe) DeleteById(id int) error {
	return a.DB.Model(&TaskLimit{}).Where("`c_id`=?", id).Delete(nil).Error
}
func (a TaskLimitOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&TaskLimit{}).Where("`c_id` in (?)", id).Delete(nil).Error
}

//------------------------task_limit------------------------//
type UdiskSecOpe struct {
	DB *gorm.DB
}

func (a UdiskSecOpe) GetByID(id int) (UdiskSec, error) {
	obj := UdiskSec{}
	err := a.DB.Model(&UdiskSec{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a UdiskSecOpe) ListByLimit(o, l int) ([]UdiskSec, error) {
	var list []UdiskSec
	err := a.DB.Model(&UdiskSec{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a UdiskSecOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&UdiskSec{}).Find(&[]UdiskSec{}).Count(&count).Error
	return count, err
}
func (a UdiskSecOpe) UpdateById(up UdiskSec) error {
	return a.DB.Model(&UdiskSec{}).Update(up).Error
}
func (a UdiskSecOpe) DeleteById(id int) error {
	return a.DB.Model(&UdiskSec{}).Where("`id`=?", id).Delete(nil).Error
}
func (a UdiskSecOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&UdiskSec{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------udisk_sec------------------------//
type DistrOpe struct {
	DB *gorm.DB
}

func (a DistrOpe) GetByID(id int) (Distr, error) {
	obj := Distr{}
	err := a.DB.Model(&Distr{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a DistrOpe) ListByLimit(o, l int) ([]Distr, error) {
	var list []Distr
	err := a.DB.Model(&Distr{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a DistrOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Distr{}).Find(&[]Distr{}).Count(&count).Error
	return count, err
}
func (a DistrOpe) UpdateById(up Distr) error {
	return a.DB.Model(&Distr{}).Update(up).Error
}
func (a DistrOpe) DeleteById(id int) error {
	return a.DB.Model(&Distr{}).Where("`id`=?", id).Delete(nil).Error
}
func (a DistrOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Distr{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------distr------------------------//
type LeakrepairClntsOpe struct {
	DB *gorm.DB
}

func (a LeakrepairClntsOpe) GetByID(id int) (LeakrepairClnts, error) {
	obj := LeakrepairClnts{}
	err := a.DB.Model(&LeakrepairClnts{}).Where("`c_id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LeakrepairClntsOpe) ListByLimit(o, l int) ([]LeakrepairClnts, error) {
	var list []LeakrepairClnts
	err := a.DB.Model(&LeakrepairClnts{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LeakrepairClntsOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LeakrepairClnts{}).Find(&[]LeakrepairClnts{}).Count(&count).Error
	return count, err
}
func (a LeakrepairClntsOpe) UpdateById(up LeakrepairClnts) error {
	return a.DB.Model(&LeakrepairClnts{}).Update(up).Error
}
func (a LeakrepairClntsOpe) DeleteById(id int) error {
	return a.DB.Model(&LeakrepairClnts{}).Where("`c_id`=?", id).Delete(nil).Error
}
func (a LeakrepairClntsOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LeakrepairClnts{}).Where("`c_id` in (?)", id).Delete(nil).Error
}

//------------------------leakrepair_clnts------------------------//
type LogIpprotoOpe struct {
	DB *gorm.DB
}

func (a LogIpprotoOpe) GetByID(id int) (LogIpproto, error) {
	obj := LogIpproto{}
	err := a.DB.Model(&LogIpproto{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogIpprotoOpe) ListByLimit(o, l int) ([]LogIpproto, error) {
	var list []LogIpproto
	err := a.DB.Model(&LogIpproto{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogIpprotoOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogIpproto{}).Find(&[]LogIpproto{}).Count(&count).Error
	return count, err
}
func (a LogIpprotoOpe) UpdateById(up LogIpproto) error {
	return a.DB.Model(&LogIpproto{}).Update(up).Error
}
func (a LogIpprotoOpe) DeleteById(id int) error {
	return a.DB.Model(&LogIpproto{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogIpprotoOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogIpproto{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_ipproto------------------------//
type LogIpattackOpe struct {
	DB *gorm.DB
}

func (a LogIpattackOpe) GetByID(id int) (LogIpattack, error) {
	obj := LogIpattack{}
	err := a.DB.Model(&LogIpattack{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogIpattackOpe) ListByLimit(o, l int) ([]LogIpattack, error) {
	var list []LogIpattack
	err := a.DB.Model(&LogIpattack{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogIpattackOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogIpattack{}).Find(&[]LogIpattack{}).Count(&count).Error
	return count, err
}
func (a LogIpattackOpe) UpdateById(up LogIpattack) error {
	return a.DB.Model(&LogIpattack{}).Update(up).Error
}
func (a LogIpattackOpe) DeleteById(id int) error {
	return a.DB.Model(&LogIpattack{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogIpattackOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogIpattack{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_ipattack------------------------//
type LogRemoteprotOpe struct {
	DB *gorm.DB
}

func (a LogRemoteprotOpe) GetByID(id int) (LogRemoteprot, error) {
	obj := LogRemoteprot{}
	err := a.DB.Model(&LogRemoteprot{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogRemoteprotOpe) ListByLimit(o, l int) ([]LogRemoteprot, error) {
	var list []LogRemoteprot
	err := a.DB.Model(&LogRemoteprot{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogRemoteprotOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogRemoteprot{}).Find(&[]LogRemoteprot{}).Count(&count).Error
	return count, err
}
func (a LogRemoteprotOpe) UpdateById(up LogRemoteprot) error {
	return a.DB.Model(&LogRemoteprot{}).Update(up).Error
}
func (a LogRemoteprotOpe) DeleteById(id int) error {
	return a.DB.Model(&LogRemoteprot{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogRemoteprotOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogRemoteprot{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_remoteprot------------------------//
type UserlogOpe struct {
	DB *gorm.DB
}

func (a UserlogOpe) GetByID(id int) (Userlog, error) {
	obj := Userlog{}
	err := a.DB.Model(&Userlog{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a UserlogOpe) ListByLimit(o, l int) ([]Userlog, error) {
	var list []Userlog
	err := a.DB.Model(&Userlog{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a UserlogOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Userlog{}).Find(&[]Userlog{}).Count(&count).Error
	return count, err
}
func (a UserlogOpe) UpdateById(up Userlog) error {
	return a.DB.Model(&Userlog{}).Update(up).Error
}
func (a UserlogOpe) DeleteById(id int) error {
	return a.DB.Model(&Userlog{}).Where("`id`=?", id).Delete(nil).Error
}
func (a UserlogOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Userlog{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------userlog------------------------//
type UserloginOpe struct {
	DB *gorm.DB
}

func (a UserloginOpe) GetByID(id int) (Userlogin, error) {
	obj := Userlogin{}
	err := a.DB.Model(&Userlogin{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a UserloginOpe) ListByLimit(o, l int) ([]Userlogin, error) {
	var list []Userlogin
	err := a.DB.Model(&Userlogin{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a UserloginOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Userlogin{}).Find(&[]Userlogin{}).Count(&count).Error
	return count, err
}
func (a UserloginOpe) UpdateById(up Userlogin) error {
	return a.DB.Model(&Userlogin{}).Update(up).Error
}
func (a UserloginOpe) DeleteById(id int) error {
	return a.DB.Model(&Userlogin{}).Where("`id`=?", id).Delete(nil).Error
}
func (a UserloginOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Userlogin{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------userlogin------------------------//
type AssetsItemsOpe struct {
	DB *gorm.DB
}

func (a AssetsItemsOpe) GetByID(id int) (AssetsItems, error) {
	obj := AssetsItems{}
	err := a.DB.Model(&AssetsItems{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a AssetsItemsOpe) ListByLimit(o, l int) ([]AssetsItems, error) {
	var list []AssetsItems
	err := a.DB.Model(&AssetsItems{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a AssetsItemsOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&AssetsItems{}).Find(&[]AssetsItems{}).Count(&count).Error
	return count, err
}
func (a AssetsItemsOpe) UpdateById(up AssetsItems) error {
	return a.DB.Model(&AssetsItems{}).Update(up).Error
}
func (a AssetsItemsOpe) DeleteById(id int) error {
	return a.DB.Model(&AssetsItems{}).Where("`id`=?", id).Delete(nil).Error
}
func (a AssetsItemsOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&AssetsItems{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------assets_items------------------------//
type HwinfoHistoryOpe struct {
	DB *gorm.DB
}

func (a HwinfoHistoryOpe) GetByID(id int) (HwinfoHistory, error) {
	obj := HwinfoHistory{}
	err := a.DB.Model(&HwinfoHistory{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a HwinfoHistoryOpe) ListByLimit(o, l int) ([]HwinfoHistory, error) {
	var list []HwinfoHistory
	err := a.DB.Model(&HwinfoHistory{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a HwinfoHistoryOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&HwinfoHistory{}).Find(&[]HwinfoHistory{}).Count(&count).Error
	return count, err
}
func (a HwinfoHistoryOpe) UpdateById(up HwinfoHistory) error {
	return a.DB.Model(&HwinfoHistory{}).Update(up).Error
}
func (a HwinfoHistoryOpe) DeleteById(id int) error {
	return a.DB.Model(&HwinfoHistory{}).Where("`id`=?", id).Delete(nil).Error
}
func (a HwinfoHistoryOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&HwinfoHistory{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------hwinfo_history------------------------//
type LogInstmonOpe struct {
	DB *gorm.DB
}

func (a LogInstmonOpe) GetByID(id int) (LogInstmon, error) {
	obj := LogInstmon{}
	err := a.DB.Model(&LogInstmon{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogInstmonOpe) ListByLimit(o, l int) ([]LogInstmon, error) {
	var list []LogInstmon
	err := a.DB.Model(&LogInstmon{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogInstmonOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogInstmon{}).Find(&[]LogInstmon{}).Count(&count).Error
	return count, err
}
func (a LogInstmonOpe) UpdateById(up LogInstmon) error {
	return a.DB.Model(&LogInstmon{}).Update(up).Error
}
func (a LogInstmonOpe) DeleteById(id int) error {
	return a.DB.Model(&LogInstmon{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogInstmonOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogInstmon{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_instmon------------------------//
type LogBehavOpe struct {
	DB *gorm.DB
}

func (a LogBehavOpe) GetByID(id int) (LogBehav, error) {
	obj := LogBehav{}
	err := a.DB.Model(&LogBehav{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogBehavOpe) ListByLimit(o, l int) ([]LogBehav, error) {
	var list []LogBehav
	err := a.DB.Model(&LogBehav{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogBehavOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogBehav{}).Find(&[]LogBehav{}).Count(&count).Error
	return count, err
}
func (a LogBehavOpe) UpdateById(up LogBehav) error {
	return a.DB.Model(&LogBehav{}).Update(up).Error
}
func (a LogBehavOpe) DeleteById(id int) error {
	return a.DB.Model(&LogBehav{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogBehavOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogBehav{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_behav------------------------//
type LogDlmonOpe struct {
	DB *gorm.DB
}

func (a LogDlmonOpe) GetByID(id int) (LogDlmon, error) {
	obj := LogDlmon{}
	err := a.DB.Model(&LogDlmon{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogDlmonOpe) ListByLimit(o, l int) ([]LogDlmon, error) {
	var list []LogDlmon
	err := a.DB.Model(&LogDlmon{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogDlmonOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogDlmon{}).Find(&[]LogDlmon{}).Count(&count).Error
	return count, err
}
func (a LogDlmonOpe) UpdateById(up LogDlmon) error {
	return a.DB.Model(&LogDlmon{}).Update(up).Error
}
func (a LogDlmonOpe) DeleteById(id int) error {
	return a.DB.Model(&LogDlmon{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogDlmonOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogDlmon{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_dlmon------------------------//
type LogSysprotOpe struct {
	DB *gorm.DB
}

func (a LogSysprotOpe) GetByID(id int) (LogSysprot, error) {
	obj := LogSysprot{}
	err := a.DB.Model(&LogSysprot{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a LogSysprotOpe) ListByLimit(o, l int) ([]LogSysprot, error) {
	var list []LogSysprot
	err := a.DB.Model(&LogSysprot{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a LogSysprotOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&LogSysprot{}).Find(&[]LogSysprot{}).Count(&count).Error
	return count, err
}
func (a LogSysprotOpe) UpdateById(up LogSysprot) error {
	return a.DB.Model(&LogSysprot{}).Update(up).Error
}
func (a LogSysprotOpe) DeleteById(id int) error {
	return a.DB.Model(&LogSysprot{}).Where("`id`=?", id).Delete(nil).Error
}
func (a LogSysprotOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&LogSysprot{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------log_sysprot------------------------//
type SchedInfoOpe struct {
	DB *gorm.DB
}

func (a SchedInfoOpe) GetByID(id int) (SchedInfo, error) {
	obj := SchedInfo{}
	err := a.DB.Model(&SchedInfo{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a SchedInfoOpe) ListByLimit(o, l int) ([]SchedInfo, error) {
	var list []SchedInfo
	err := a.DB.Model(&SchedInfo{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a SchedInfoOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&SchedInfo{}).Find(&[]SchedInfo{}).Count(&count).Error
	return count, err
}
func (a SchedInfoOpe) UpdateById(up SchedInfo) error {
	return a.DB.Model(&SchedInfo{}).Update(up).Error
}
func (a SchedInfoOpe) DeleteById(id int) error {
	return a.DB.Model(&SchedInfo{}).Where("`id`=?", id).Delete(nil).Error
}
func (a SchedInfoOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&SchedInfo{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------sched_info------------------------//
type TaskStatOpe struct {
	DB *gorm.DB
}

func (a TaskStatOpe) GetByID(id int) (TaskStat, error) {
	obj := TaskStat{}
	err := a.DB.Model(&TaskStat{}).Where("`taskid`=?", id).Find(&obj).Error
	return obj, err
}

func (a TaskStatOpe) ListByLimit(o, l int) ([]TaskStat, error) {
	var list []TaskStat
	err := a.DB.Model(&TaskStat{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a TaskStatOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&TaskStat{}).Find(&[]TaskStat{}).Count(&count).Error
	return count, err
}
func (a TaskStatOpe) UpdateById(up TaskStat) error {
	return a.DB.Model(&TaskStat{}).Update(up).Error
}
func (a TaskStatOpe) DeleteById(id int) error {
	return a.DB.Model(&TaskStat{}).Where("`taskid`=?", id).Delete(nil).Error
}
func (a TaskStatOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&TaskStat{}).Where("`taskid` in (?)", id).Delete(nil).Error
}

//------------------------task_stat------------------------//
type ClntsOsinfoOpe struct {
	DB *gorm.DB
}

func (a ClntsOsinfoOpe) GetByID(id int) (ClntsOsinfo, error) {
	obj := ClntsOsinfo{}
	err := a.DB.Model(&ClntsOsinfo{}).Where("`c_id`=?", id).Find(&obj).Error
	return obj, err
}

func (a ClntsOsinfoOpe) ListByLimit(o, l int) ([]ClntsOsinfo, error) {
	var list []ClntsOsinfo
	err := a.DB.Model(&ClntsOsinfo{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a ClntsOsinfoOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&ClntsOsinfo{}).Find(&[]ClntsOsinfo{}).Count(&count).Error
	return count, err
}
func (a ClntsOsinfoOpe) UpdateById(up ClntsOsinfo) error {
	return a.DB.Model(&ClntsOsinfo{}).Update(up).Error
}
func (a ClntsOsinfoOpe) DeleteById(id int) error {
	return a.DB.Model(&ClntsOsinfo{}).Where("`c_id`=?", id).Delete(nil).Error
}
func (a ClntsOsinfoOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&ClntsOsinfo{}).Where("`c_id` in (?)", id).Delete(nil).Error
}

//------------------------clnts_osinfo------------------------//
type ClntsStatOpe struct {
	DB *gorm.DB
}

func (a ClntsStatOpe) GetByID(id int) (ClntsStat, error) {
	obj := ClntsStat{}
	err := a.DB.Model(&ClntsStat{}).Where("`c_id`=?", id).Find(&obj).Error
	return obj, err
}

func (a ClntsStatOpe) ListByLimit(o, l int) ([]ClntsStat, error) {
	var list []ClntsStat
	err := a.DB.Model(&ClntsStat{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a ClntsStatOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&ClntsStat{}).Find(&[]ClntsStat{}).Count(&count).Error
	return count, err
}
func (a ClntsStatOpe) UpdateById(up ClntsStat) error {
	return a.DB.Model(&ClntsStat{}).Update(up).Error
}
func (a ClntsStatOpe) DeleteById(id int) error {
	return a.DB.Model(&ClntsStat{}).Where("`c_id`=?", id).Delete(nil).Error
}
func (a ClntsStatOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&ClntsStat{}).Where("`c_id` in (?)", id).Delete(nil).Error
}

//------------------------clnts_stat------------------------//
type ClntsSwinfoOpe struct {
	DB *gorm.DB
}

func (a ClntsSwinfoOpe) GetByID(id int) (ClntsSwinfo, error) {
	obj := ClntsSwinfo{}
	err := a.DB.Model(&ClntsSwinfo{}).Where("`c_id`=?", id).Find(&obj).Error
	return obj, err
}

func (a ClntsSwinfoOpe) ListByLimit(o, l int) ([]ClntsSwinfo, error) {
	var list []ClntsSwinfo
	err := a.DB.Model(&ClntsSwinfo{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a ClntsSwinfoOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&ClntsSwinfo{}).Find(&[]ClntsSwinfo{}).Count(&count).Error
	return count, err
}
func (a ClntsSwinfoOpe) UpdateById(up ClntsSwinfo) error {
	return a.DB.Model(&ClntsSwinfo{}).Update(up).Error
}
func (a ClntsSwinfoOpe) DeleteById(id int) error {
	return a.DB.Model(&ClntsSwinfo{}).Where("`c_id`=?", id).Delete(nil).Error
}
func (a ClntsSwinfoOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&ClntsSwinfo{}).Where("`c_id` in (?)", id).Delete(nil).Error
}

//------------------------clnts_swinfo------------------------//
type UserinfoOpe struct {
	DB *gorm.DB
}

func (a UserinfoOpe) GetByID(id int) (Userinfo, error) {
	obj := Userinfo{}
	err := a.DB.Model(&Userinfo{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a UserinfoOpe) ListByLimit(o, l int) ([]Userinfo, error) {
	var list []Userinfo
	err := a.DB.Model(&Userinfo{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a UserinfoOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Userinfo{}).Find(&[]Userinfo{}).Count(&count).Error
	return count, err
}
func (a UserinfoOpe) UpdateById(up Userinfo) error {
	return a.DB.Model(&Userinfo{}).Update(up).Error
}
func (a UserinfoOpe) DeleteById(id int) error {
	return a.DB.Model(&Userinfo{}).Where("`id`=?", id).Delete(nil).Error
}
func (a UserinfoOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Userinfo{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------userinfo------------------------//
