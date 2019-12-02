package Mappppper

import (
	"github.com/guregu/null"
)

type PoliciesDetail struct {
	Id      null.Int    `json:"id"`
	Product null.String `json:"product"`
	Fname   null.String `json:"fname"`
	Uptm    null.Time   `json:"uptm"`
	Val     null.String `json:"val"`
}
type Sysconf struct {
	Name null.String `json:"name"`
	Uptm null.Time   `json:"uptm"`
	Val  null.String `json:"val"`
}
type ClntsAssets struct {
	Id    null.Int    `json:"id"`
	CId   null.Int    `json:"c_id"`
	Cid   null.String `json:"cid"`
	Label null.String `json:"label"`
	Value null.String `json:"value"`
}
type ClntsTags struct {
	Id    null.Int    `json:"id"`
	CId   null.Int    `json:"c_id"`
	Cid   null.String `json:"cid"`
	Tagid null.Int    `json:"tagid"`
}
type Groups struct {
	Id     null.Int    `json:"id"`
	Name   null.String `json:"name"`
	Parent null.Int    `json:"parent"`
	Pol    null.Int    `json:"pol"`
	Rules  null.String `json:"rules"`
}
type LeakrepairKbsave struct {
	Id     null.Int    `json:"id"`
	Kbid   null.String `json:"kbid"`
	Sha1   null.String `json:"sha1"`
	Url    null.String `json:"url"`
	Size   null.Int    `json:"size"`
	Done   null.Int    `json:"done"`
	Donetm null.Time   `json:"donetm"`
}
type LogIntrusion struct {
	Id           null.Int    `json:"id"`
	Cid          null.String `json:"cid"`
	Upts         null.Time   `json:"upts"`
	Prodver      null.Int    `json:"prodver"`
	Dbver        null.Time   `json:"dbver"`
	Ts           null.Time   `json:"ts"`
	Name         null.String `json:"name"`
	Outbound     null.Int    `json:"outbound"`
	Blocked      null.Int    `json:"blocked"`
	Protocol     null.Int    `json:"protocol"`
	ProtocolData null.String `json:"protocol_data"`
	Laddr        null.Int    `json:"laddr"`
	Lport        null.Int    `json:"lport"`
	Raddr        null.Int    `json:"raddr"`
	Rport        null.Int    `json:"rport"`
	ParentCmd    null.String `json:"parent_cmd"`
	ProcPath     null.String `json:"proc_path"`
	ProcName     null.String `json:"proc_name"`
	Cmdline      null.String `json:"cmdline"`
	Origin       null.String `json:"origin"`
}
type LogUdiskmon struct {
	Id       null.Int    `json:"id"`
	CId      null.Int    `json:"c_id"`
	Cid      null.String `json:"cid"`
	Uptm     null.Time   `json:"uptm"`
	Prodtype null.String `json:"prodtype"`
	Prodver  null.Int    `json:"prodver"`
	Virdb    null.Time   `json:"virdb"`
	Tm       null.Time   `json:"tm"`
	Object   null.String `json:"object"`
	Objname  null.String `json:"objname"`
	Recid    null.String `json:"recid"`
	Recname  null.String `json:"recname"`
	Result   null.Int    `json:"result"`
	Origin   null.String `json:"origin"`
}
type Whitelist struct {
	Id     null.Int    `json:"id"`
	Type   null.String `json:"type"`
	Uptm   null.Time   `json:"uptm"`
	Val    null.String `json:"val"`
	Act    null.Int    `json:"act"`
	Enable null.Int    `json:"enable"`
	Memo   null.String `json:"memo"`
}
type Swinfo struct {
	Id        null.Int    `json:"id"`
	Name      null.String `json:"name"`
	Publisher null.String `json:"publisher"`
	Version   null.String `json:"version"`
}
type TagsDict struct {
	Id     null.Int    `json:"id"`
	Name   null.String `json:"name"`
	Group  null.String `json:"group"`
	Userid null.Int    `json:"userid"`
}
type ClntsOl struct {
	CId  null.Int    `json:"c_id"`
	Cid  null.String `json:"cid"`
	Mac  null.String `json:"mac"`
	Upip null.Int    `json:"upip"`
	Hbtm null.Time   `json:"hbtm"`
	Stat null.Int    `json:"stat"`
}
type LeakrepairKbinfo struct {
	Id      null.Int    `json:"id"`
	Kbid    null.String `json:"kbid"`
	Pubtm   null.Time   `json:"pubtm"`
	Level   null.Int    `json:"level"`
	Exclude null.Int    `json:"exclude"`
	Desc    null.String `json:"desc"`
}
type LogCron struct {
	Cid  null.String `json:"cid"`
	Date null.Time   `json:"date"`
	Wpos null.Int    `json:"wpos"`
	Rpos null.Int    `json:"rpos"`
}
type LogIpblacklist struct {
	Id        null.Int    `json:"id"`
	Cid       null.String `json:"cid"`
	Upts      null.Time   `json:"upts"`
	Prodver   null.Int    `json:"prodver"`
	Dbver     null.Time   `json:"dbver"`
	Ts        null.Time   `json:"ts"`
	Name      null.String `json:"name"`
	Outbound  null.Int    `json:"outbound"`
	Blocked   null.Int    `json:"blocked"`
	Laddr     null.Int    `json:"laddr"`
	Lport     null.Int    `json:"lport"`
	Raddr     null.Int    `json:"raddr"`
	Rport     null.Int    `json:"rport"`
	ParentCmd null.String `json:"parent_cmd"`
	ProcPath  null.String `json:"proc_path"`
	ProcName  null.String `json:"proc_name"`
	Cmdline   null.String `json:"cmdline"`
	Origin    null.String `json:"origin"`
}
type LogRemote struct {
	Id       null.Int    `json:"id"`
	Tsstart  null.Time   `json:"tsstart"`
	Tsend    null.Time   `json:"tsend"`
	Uid      null.Int    `json:"uid"`
	Ip       null.Int    `json:"ip"`
	Cid      null.String `json:"cid"`
	Viewonly null.Int    `json:"viewonly"`
	Reason   null.String `json:"reason"`
	Origin   null.String `json:"origin"`
}
type LogUpgrade struct {
	Id      null.Int    `json:"id"`
	Ts      null.Time   `json:"ts"`
	Product null.String `json:"product"`
	Orgver  null.String `json:"orgver"`
	Newver  null.String `json:"newver"`
	Result  null.Int    `json:"result"`
	Origin  null.String `json:"origin"`
}
type TaskInfo struct {
	Id        null.Int    `json:"id"`
	Userid    null.Int    `json:"userid"`
	Tmstart   null.Time   `json:"tmstart"`
	Tmexpire  null.Time   `json:"tmexpire"`
	Cls       null.String `json:"cls"`
	Param     null.String `json:"param"`
	ClntsAll  null.Int    `json:"clnts_all"`
	ClntsDone null.Int    `json:"clnts_done"`
}
type ClntsHwinfo struct {
	CId     null.Int    `json:"c_id"`
	Cid     null.String `json:"cid"`
	Uptm    null.Time   `json:"uptm"`
	Cpu     null.String `json:"cpu"`
	Mem     null.String `json:"mem"`
	Board   null.String `json:"board"`
	Hdd     null.String `json:"hdd"`
	Nic     null.String `json:"nic"`
	Video   null.String `json:"video"`
	InfoCrc null.Int    `json:"info_crc"`
	Info    null.String `json:"info"`
}
type LeakrepairClntkb struct {
	Id      null.Int    `json:"id"`
	Kbid    null.String `json:"kbid"`
	CId     null.Int    `json:"c_id"`
	Cid     null.String `json:"cid"`
	Exclude null.Int    `json:"exclude"`
}
type LogMailmon struct {
	Id       null.Int    `json:"id"`
	CId      null.Int    `json:"c_id"`
	Cid      null.String `json:"cid"`
	Uptm     null.Time   `json:"uptm"`
	Prodtype null.String `json:"prodtype"`
	Prodver  null.Int    `json:"prodver"`
	Virdb    null.Time   `json:"virdb"`
	Tm       null.Time   `json:"tm"`
	Object   null.String `json:"object"`
	Objname  null.String `json:"objname"`
	Recid    null.String `json:"recid"`
	Recname  null.String `json:"recname"`
	Result   null.Int    `json:"result"`
	Origin   null.String `json:"origin"`
}
type LogMalurl struct {
	Id       null.Int    `json:"id"`
	Cid      null.String `json:"cid"`
	Upts     null.Time   `json:"upts"`
	Prodver  null.Int    `json:"prodver"`
	Dbver    null.Time   `json:"dbver"`
	Ts       null.Time   `json:"ts"`
	Cls      null.String `json:"cls"`
	Domain   null.String `json:"domain"`
	Url      null.String `json:"url"`
	ProcPath null.String `json:"proc_path"`
	ProcName null.String `json:"proc_name"`
	Origin   null.String `json:"origin"`
}
type Clnts struct {
	Id           null.Int    `json:"id"`
	Cid          null.String `json:"cid"`
	Gid          null.Int    `json:"gid"`
	Pol          null.Int    `json:"pol"`
	Hostname     null.String `json:"hostname"`
	Aliasname    null.String `json:"aliasname"`
	Mac          null.String `json:"mac"`
	Ip           null.String `json:"ip"`
	Regtm        null.Time   `json:"regtm"`
	Uptm         null.Time   `json:"uptm"`
	Ostype       null.String `json:"ostype"`
	Osver        null.String `json:"osver"`
	Stat         null.Int    `json:"stat"`
	Prodtype     null.String `json:"prodtype"`
	Prodver      null.Int    `json:"prodver"`
	Virdb        null.Time   `json:"virdb"`
	SysprotFlags null.Int    `json:"sysprot_flags"`
	NetprotFlags null.Int    `json:"netprot_flags"`
	Lastuser     null.String `json:"lastuser"`
}
type ClntsSched struct {
	CId    null.Int    `json:"c_id"`
	Cid    null.String `json:"cid"`
	Taskid null.Int    `json:"taskid"`
}
type LogScan struct {
	Id       null.Int    `json:"id"`
	CId      null.Int    `json:"c_id"`
	Cid      null.String `json:"cid"`
	Uptm     null.Time   `json:"uptm"`
	Prodtype null.String `json:"prodtype"`
	Prodver  null.Int    `json:"prodver"`
	Virdb    null.Time   `json:"virdb"`
	Tm       null.Time   `json:"tm"`
	Object   null.String `json:"object"`
	Objname  null.String `json:"objname"`
	Recid    null.String `json:"recid"`
	Recname  null.String `json:"recname"`
	Result   null.Int    `json:"result"`
	Origin   null.String `json:"origin"`
}
type Distr struct {
	Id    null.Int    `json:"id"`
	Type  null.String `json:"type"`
	Name  null.String `json:"name"`
	Memo  null.String `json:"memo"`
	Size  null.Int    `json:"size"`
	Ctime null.Time   `json:"ctime"`
}
type LeakrepairClnts struct {
	CId     null.Int    `json:"c_id"`
	Cid     null.String `json:"cid"`
	Tm      null.Time   `json:"tm"`
	Exclude null.Int    `json:"exclude"`
}
type LogIpproto struct {
	Id           null.Int    `json:"id"`
	Cid          null.String `json:"cid"`
	Upts         null.Time   `json:"upts"`
	Prodver      null.Int    `json:"prodver"`
	Dbver        null.Time   `json:"dbver"`
	Ts           null.Time   `json:"ts"`
	Name         null.String `json:"name"`
	Outbound     null.Int    `json:"outbound"`
	Blocked      null.Int    `json:"blocked"`
	Protocol     null.Int    `json:"protocol"`
	ProtocolData null.String `json:"protocol_data"`
	Laddr        null.Int    `json:"laddr"`
	Lport        null.Int    `json:"lport"`
	Raddr        null.Int    `json:"raddr"`
	Rport        null.Int    `json:"rport"`
	ProcPath     null.String `json:"proc_path"`
	Cmdline      null.String `json:"cmdline"`
	Origin       null.String `json:"origin"`
}
type Quarantine struct {
	Id        null.Int    `json:"id"`
	Cid       null.String `json:"cid"`
	Uptm      null.Time   `json:"uptm"`
	Prodver   null.Int    `json:"prodver"`
	Dbver     null.Time   `json:"dbver"`
	Tm        null.Time   `json:"tm"`
	Fname     null.String `json:"fname"`
	Sha1      null.String `json:"sha1"`
	FilePath  null.String `json:"file_path"`
	FileName  null.String `json:"file_name"`
	VirusId   null.String `json:"virus_id"`
	VirusName null.String `json:"virus_name"`
}
type TaskLimit struct {
	CId      null.Int    `json:"c_id"`
	Cid      null.String `json:"cid"`
	Taskid   null.Int    `json:"taskid"`
	Tmexpire null.Time   `json:"tmexpire"`
}
type UdiskSec struct {
	Id       null.String `json:"id"`
	Label    null.String `json:"label"`
	Memo     null.String `json:"memo"`
	Vendor   null.String `json:"vendor"`
	Product  null.String `json:"product"`
	Serial   null.String `json:"serial"`
	Cap      null.Int    `json:"cap"`
	Regtm    null.Time   `json:"regtm"`
	Userid   null.Int    `json:"userid"`
	Password null.Int    `json:"password"`
	Outside  null.Int    `json:"outside"`
}
type Userlogin struct {
	Id  null.Int  `json:"id"`
	Cnt null.Int  `json:"cnt"`
	Tm  null.Time `json:"tm"`
}
type AssetsItems struct {
	Id      null.Int    `json:"id"`
	Label   null.String `json:"label"`
	Input   null.String `json:"input"`
	Empty   null.Int    `json:"empty"`
	Options null.String `json:"options"`
}
type HwinfoHistory struct {
	Id      null.Int    `json:"id"`
	CId     null.Int    `json:"c_id"`
	Cid     null.String `json:"cid"`
	Uptm    null.Time   `json:"uptm"`
	InfoCrc null.Int    `json:"info_crc"`
	Info    null.String `json:"info"`
}
type LogInstmon struct {
	Id         null.Int    `json:"id"`
	CId        null.Int    `json:"c_id"`
	Cid        null.String `json:"cid"`
	Uptm       null.Time   `json:"uptm"`
	Prodtype   null.String `json:"prodtype"`
	Prodver    null.Int    `json:"prodver"`
	Virdb      null.Time   `json:"virdb"`
	Tm         null.Time   `json:"tm"`
	PprocPath  null.String `json:"pproc_path"`
	PprocName  null.String `json:"pproc_name"`
	PprocCmdln null.String `json:"pproc_cmdln"`
	PprocSha1  null.String `json:"pproc_sha1"`
	ProcPath   null.String `json:"proc_path"`
	ProcName   null.String `json:"proc_name"`
	ProcCmdln  null.String `json:"proc_cmdln"`
	ProcSha1   null.String `json:"proc_sha1"`
	ResPath    null.String `json:"res_path"`
	ResName    null.String `json:"res_name"`
	ResCmdln   null.String `json:"res_cmdln"`
	ResSha1    null.String `json:"res_sha1"`
	Runcmd     null.String `json:"runcmd"`
	Recid      null.String `json:"recid"`
	Recname    null.String `json:"recname"`
	DispName   null.String `json:"disp_name"`
	Treatment  null.Int    `json:"treatment"`
	Origin     null.String `json:"origin"`
}
type LogIpattack struct {
	Id        null.Int    `json:"id"`
	Cid       null.String `json:"cid"`
	Upts      null.Time   `json:"upts"`
	Prodver   null.Int    `json:"prodver"`
	Dbver     null.Time   `json:"dbver"`
	Ts        null.Time   `json:"ts"`
	FloodType null.Int    `json:"flood_type"`
	Blocked   null.Int    `json:"blocked"`
	Raddr     null.Int    `json:"raddr"`
	ParentCmd null.String `json:"parent_cmd"`
	ProcPath  null.String `json:"proc_path"`
	ProcName  null.String `json:"proc_name"`
	Cmdline   null.String `json:"cmdline"`
	Origin    null.String `json:"origin"`
}
type LogRemoteprot struct {
	Id       null.Int    `json:"id"`
	Cid      null.String `json:"cid"`
	Upts     null.Time   `json:"upts"`
	Prodver  null.Int    `json:"prodver"`
	Dbver    null.Time   `json:"dbver"`
	Ts       null.Time   `json:"ts"`
	Raddr    null.Int    `json:"raddr"`
	LogonTs  null.Time   `json:"logon_ts"`
	Hostname null.String `json:"hostname"`
	Username null.String `json:"username"`
	Origin   null.String `json:"origin"`
}
type Userlog struct {
	Id          null.Int    `json:"id"`
	Uid         null.Int    `json:"uid"`
	Logts       null.Time   `json:"logts"`
	Upip        null.Int    `json:"upip"`
	Module      null.String `json:"module"`
	Description null.String `json:"description"`
	Method      null.String `json:"method"`
	Param       null.String `json:"param"`
	Origin      null.String `json:"origin"`
}
type SchedInfo struct {
	Id           null.Int    `json:"id"`
	Userid       null.Int    `json:"userid"`
	Ctime        null.Time   `json:"ctime"`
	Utime        null.Time   `json:"utime"`
	TriggerType  null.String `json:"trigger_type"`
	TriggerParam null.String `json:"trigger_param"`
	Cls          null.String `json:"cls"`
	Param        null.String `json:"param"`
}
type TaskStat struct {
	Taskid null.Int    `json:"taskid"`
	CId    null.Int    `json:"c_id"`
	Cid    null.String `json:"cid"`
	Uptm   null.Time   `json:"uptm"`
	State  null.Int    `json:"state"`
	Msg    null.String `json:"msg"`
}
type ClntsOsinfo struct {
	CId     null.Int    `json:"c_id"`
	Cid     null.String `json:"cid"`
	Osinfo  null.String `json:"osinfo"`
	Osusers null.String `json:"osusers"`
	Oslogon null.String `json:"oslogon"`
}
type ClntsStat struct {
	CId           null.Int    `json:"c_id"`
	Cid           null.String `json:"cid"`
	Tasktm        null.Time   `json:"tasktm"`
	SwinfoCrc     null.Int    `json:"swinfo_crc"`
	OsinfoCrc     null.Int    `json:"osinfo_crc"`
	HwinfoCrc     null.Int    `json:"hwinfo_crc"`
	OsuserinfoCrc null.Int    `json:"osuserinfo_crc"`
	KbinfoCrc     null.Int    `json:"kbinfo_crc"`
}
type ClntsSwinfo struct {
	CId  null.Int    `json:"c_id"`
	Cid  null.String `json:"cid"`
	Swid null.Int    `json:"swid"`
}
type LogBehav struct {
	Id       null.Int    `json:"id"`
	CId      null.Int    `json:"c_id"`
	Cid      null.String `json:"cid"`
	Uptm     null.Time   `json:"uptm"`
	Prodtype null.String `json:"prodtype"`
	Prodver  null.Int    `json:"prodver"`
	Virdb    null.Time   `json:"virdb"`
	Tm       null.Time   `json:"tm"`
	Object   null.String `json:"object"`
	Objname  null.String `json:"objname"`
	Recid    null.String `json:"recid"`
	Recname  null.String `json:"recname"`
	Result   null.Int    `json:"result"`
	Origin   null.String `json:"origin"`
}
type LogDlmon struct {
	Id       null.Int    `json:"id"`
	CId      null.Int    `json:"c_id"`
	Cid      null.String `json:"cid"`
	Uptm     null.Time   `json:"uptm"`
	Prodtype null.String `json:"prodtype"`
	Prodver  null.Int    `json:"prodver"`
	Virdb    null.Time   `json:"virdb"`
	Tm       null.Time   `json:"tm"`
	Object   null.String `json:"object"`
	Objname  null.String `json:"objname"`
	Recid    null.String `json:"recid"`
	Recname  null.String `json:"recname"`
	Result   null.Int    `json:"result"`
	Origin   null.String `json:"origin"`
}
type LogSysprot struct {
	Id        null.Int    `json:"id"`
	CId       null.Int    `json:"c_id"`
	Cid       null.String `json:"cid"`
	Uptm      null.Time   `json:"uptm"`
	Prodtype  null.String `json:"prodtype"`
	Prodver   null.Int    `json:"prodver"`
	Virdb     null.Time   `json:"virdb"`
	Tm        null.Time   `json:"tm"`
	Fname     null.String `json:"fname"`
	Recname   null.String `json:"recname"`
	Montype   null.Int    `json:"montype"`
	Acttype   null.Int    `json:"acttype"`
	ResPath   null.String `json:"res_path"`
	ResName   null.String `json:"res_name"`
	Treatment null.Int    `json:"treatment"`
	Origin    null.String `json:"origin"`
}
type Userinfo struct {
	Id        null.Int    `json:"id"`
	Group     null.String `json:"group"`
	Name      null.String `json:"name"`
	Phone     null.String `json:"phone"`
	Regtm     null.Time   `json:"regtm"`
	State     null.Int    `json:"state"`
	Password  null.String `json:"password"`
	Memo      null.String `json:"memo"`
	Privilege null.String `json:"privilege"`
	Config    null.String `json:"config"`
}
type LogLeakrepair struct {
	Id      null.Int    `json:"id"`
	Cid     null.String `json:"cid"`
	Upts    null.Time   `json:"upts"`
	Prodver null.Int    `json:"prodver"`
	Dbver   null.Time   `json:"dbver"`
	Ts      null.Time   `json:"ts"`
	Kbid    null.String `json:"kbid"`
	Level   null.Int    `json:"level"`
	State   null.Int    `json:"state"`
	Desc    null.String `json:"desc"`
	Origin  null.String `json:"origin"`
}
type LogStat struct {
	Ts   null.Time   `json:"ts"`
	Type null.Int    `json:"type"`
	Cid  null.String `json:"cid"`
	Cnt  null.Int    `json:"cnt"`
}
type Nodes struct {
	Id        null.Int    `json:"id"`
	Hostname  null.String `json:"hostname"`
	Aliasname null.String `json:"aliasname"`
	Ip        null.Int    `json:"ip"`
	Url       null.String `json:"url"`
	Hbtm      null.Time   `json:"hbtm"`
	Status    null.Int    `json:"status"`
	Data      null.String `json:"data"`
}
type Policies struct {
	Id   null.Int    `json:"id"`
	Name null.String `json:"name"`
	Uptm null.Time   `json:"uptm"`
}
