package www

import "strings"

var contentTypeMap map[string]string
var contentValMap map[string][]string

func ContentType(key string) string {
	s, ok := contentTypeMap[strings.ToLower(key)]
	if !ok {
		s = contentTypeMap[".*"]
	}
	return s
}
func ContentVal(val string) []string {
	return contentValMap[strings.ToLower(val)]
}
func init() {
	contentTypeMap = map[string]string{
		".*":     "application/octet-stream",
		".001":   "application/x-001",
		".301":   "application/x-301",
		".323":   "text/h323",
		".906":   "application/x-906",
		".907":   "drawing/907",
		".a11":   "application/x-a11",
		".acp":   "audio/x-mei-aac",
		".ai":    "application/postscript",
		".aif":   "audio/aiff",
		".aifc":  "audio/aiff",
		".aiff":  "audio/aiff",
		".anv":   "application/x-anv",
		".asa":   "text/asa",
		".asf":   "video/x-ms-asf",
		".asp":   "text/asp",
		".asx":   "video/x-ms-asf",
		".au":    "audio/basic",
		".avi":   "video/avi",
		".awf":   "application/vnd.adobe.workflow",
		".biz":   "text/xml",
		".bmp":   "application/x-bmp",
		".bot":   "application/x-bot",
		".c4t":   "application/x-c4t",
		".c90":   "application/x-c90",
		".cal":   "application/x-cals",
		".cat":   "application/vnd.ms-pki.seccat",
		".cdf":   "application/x-netcdf",
		".cdr":   "application/x-cdr",
		".cel":   "application/x-cel",
		".cer":   "application/x-x509-ca-cert",
		".cg4":   "application/x-g4",
		".cgm":   "application/x-cgm",
		".cit":   "application/x-cit",
		".class": "java/*",
		".cml":   "text/xml",
		".cmp":   "application/x-cmp",
		".cmx":   "application/x-cmx",
		".cot":   "application/x-cot",
		".crl":   "application/pkix-crl",
		".crt":   "application/x-x509-ca-cert",
		".csi":   "application/x-csi",
		".css":   "text/css",
		".cut":   "application/x-cut",
		".dbf":   "application/x-dbf",
		".dbm":   "application/x-dbm",
		".dbx":   "application/x-dbx",
		".dcd":   "text/xml",
		".dcx":   "application/x-dcx",
		".der":   "application/x-x509-ca-cert",
		".dgn":   "application/x-dgn",
		".dib":   "application/x-dib",
		".dll":   "application/x-msdownload",
		".doc":   "application/msword",
		".dot":   "application/msword",
		".drw":   "application/x-drw",
		".dtd":   "text/xml",
		".dwf":   "Model/vnd.dwf",
		//".dwf":"application/x-dwf",
		".dwg": "application/x-dwg",
		".dxb": "application/x-dxb",
		".dxf": "application/x-dxf",
		".edn": "application/vnd.adobe.edn",
		".emf": "application/x-emf",
		".eml": "message/rfc822",
		".ent": "text/xml",
		".epi": "application/x-epi",
		".eps": "application/x-ps",
		//".eps":"application/postscript",
		".etd":  "application/x-ebx",
		".exe":  "application/x-msdownload",
		".fax":  "image/fax",
		".fdf":  "application/vnd.fdf",
		".fif":  "application/fractals",
		".fo":   "text/xml",
		".frm":  "application/x-frm",
		".g4":   "application/x-g4",
		".gbr":  "application/x-gbr",
		".":     "application/x-",
		".gif":  "image/gif",
		".gl2":  "application/x-gl2",
		".gp4":  "application/x-gp4",
		".hgl":  "application/x-hgl",
		".hmr":  "application/x-hmr",
		".hpg":  "application/x-hpgl",
		".hpl":  "application/x-hpl",
		".hqx":  "application/mac-binhex40",
		".hrf":  "application/x-hrf",
		".hta":  "application/hta",
		".htc":  "text/x-component",
		".htm":  "text/html",
		".html": "text/html",
		".htt":  "text/webviewhtml",
		".htx":  "text/html",
		".icb":  "application/x-icb",
		".ico":  "image/x-icon",
		//".ico":"application/x-ico",
		".iff":  "application/x-iff",
		".ig4":  "application/x-g4",
		".igs":  "application/x-igs",
		".iii":  "application/x-iphone",
		".img":  "application/x-img",
		".ins":  "application/x-internet-signup",
		".isp":  "application/x-internet-signup",
		".IVF":  "video/x-ivf",
		".java": "java/*",
		".jfif": "image/jpeg",
		".jpe":  "image/jpeg",
		//".jpe":"application/x-jpe",
		".jpeg": "image/jpeg",
		".jpg":  "image/jpeg",
		//".jpg":"application/x-jpg",
		".js":    "application/x-javascript",
		".jsp":   "text/html",
		".la1":   "audio/x-liquid-file",
		".lar":   "application/x-laplayer-reg",
		".latex": "application/x-latex",
		".lavs":  "audio/x-liquid-secure",
		".lbm":   "application/x-lbm",
		".lmsff": "audio/x-la-lms",
		".ls":    "application/x-javascript",
		".ltr":   "application/x-ltr",
		".m1v":   "video/x-mpeg",
		".m2v":   "video/x-mpeg",
		".m3u":   "audio/mpegurl",
		".m4e":   "video/mpeg4",
		".mac":   "application/x-mac",
		".man":   "application/x-troff-man",
		".math":  "text/xml",
		".mdb":   "application/msaccess",
		//".mdb":"application/x-mdb",
		".mfp":   "application/x-shockwave-flash",
		".mht":   "message/rfc822",
		".mhtml": "message/rfc822",
		".mi":    "application/x-mi",
		".mid":   "audio/mid",
		".midi":  "audio/mid",
		".mil":   "application/x-mil",
		".mml":   "text/xml",
		".mnd":   "audio/x-musicnet-download",
		".mns":   "audio/x-musicnet-stream",
		".mocha": "application/x-javascript",
		".movie": "video/x-sgi-movie",
		".mp1":   "audio/mp1",
		".mp2":   "audio/mp2",
		".mp2v":  "video/mpeg",
		".mp3":   "audio/mp3",
		".mp4":   "video/mpeg4",
		".mpa":   "video/x-mpg",
		".mpd":   "application/vnd.ms-project",
		".mpe":   "video/x-mpeg",
		".mpeg":  "video/mpg",
		".mpg":   "video/mpg",
		".mpga":  "audio/rn-mpeg",
		".mpp":   "application/vnd.ms-project",
		".mps":   "video/x-mpeg",
		".mpt":   "application/vnd.ms-project",
		".mpv":   "video/mpg",
		".mpv2":  "video/mpeg",
		".mpw":   "application/vnd.ms-project",
		".mpx":   "application/vnd.ms-project",
		".mtx":   "text/xml",
		".mxp":   "application/x-mmxp",
		".net":   "image/pnetvue",
		".nrf":   "application/x-nrf",
		".nws":   "message/rfc822",
		".odc":   "text/x-ms-odc",
		".out":   "application/x-out",
		".p10":   "application/pkcs10",
		".p12":   "application/x-pkcs12",
		".p7b":   "application/x-pkcs7-certificates",
		".p7c":   "application/pkcs7-mime",
		".p7m":   "application/pkcs7-mime",
		".p7r":   "application/x-pkcs7-certreqresp",
		".p7s":   "application/pkcs7-signature",
		".pc5":   "application/x-pc5",
		".pci":   "application/x-pci",
		".pcl":   "application/x-pcl",
		".pcx":   "application/x-pcx",
		".pdf":   "application/pdf",
		".pdx":   "application/vnd.adobe.pdx",
		".pfx":   "application/x-pkcs12",
		".pgl":   "application/x-pgl",
		".pic":   "application/x-pic",
		".pko":   "application/vnd.ms-pki.pko",
		".pl":    "application/x-perl",
		".plg":   "text/html",
		".pls":   "audio/scpls",
		".plt":   "application/x-plt",
		".png":   "image/png",
		//".png":"application/x-png",
		".pot": "application/vnd.ms-powerpoint",
		".ppa": "application/vnd.ms-powerpoint",
		".ppm": "application/x-ppm",
		".pps": "application/vnd.ms-powerpoint",
		".ppt": "application/vnd.ms-powerpoint",
		//".ppt":"application/x-ppt",
		".pr":  "application/x-pr",
		".prf": "application/pics-rules",
		".prn": "application/x-prn",
		".prt": "application/x-prt",
		".ps":  "application/postscript",
		//".ps":"application/x-ps",
		".ptn":  "application/x-ptn",
		".pwz":  "application/vnd.ms-powerpoint",
		".r3t":  "text/vnd.rn-realtext3d",
		".ra":   "audio/vnd.rn-realaudio",
		".ram":  "audio/x-pn-realaudio",
		".ras":  "application/x-ras",
		".rat":  "application/rat-file",
		".rdf":  "text/xml",
		".rec":  "application/vnd.rn-recording",
		".red":  "application/x-red",
		".rgb":  "application/x-rgb",
		".rjs":  "application/vnd.rn-realsystem-rjs",
		".rjt":  "application/vnd.rn-realsystem-rjt",
		".rlc":  "application/x-rlc",
		".rle":  "application/x-rle",
		".rm":   "application/vnd.rn-realmedia",
		".rmf":  "application/vnd.adobe.rmf",
		".rmi":  "audio/mid",
		".rmj":  "application/vnd.rn-realsystem-rmj",
		".rmm":  "audio/x-pn-realaudio",
		".rmp":  "application/vnd.rn-rn_music_package",
		".rms":  "application/vnd.rn-realmedia-secure",
		".rmvb": "application/vnd.rn-realmedia-vbr",
		".rmx":  "application/vnd.rn-realsystem-rmx",
		".rnx":  "application/vnd.rn-realplayer",
		".rp":   "image/vnd.rn-realpix",
		".rpm":  "audio/x-pn-realaudio-plugin",
		".rsml": "application/vnd.rn-rsml",
		".rt":   "text/vnd.rn-realtext",
		".rtf":  "application/msword",
		//".rtf":"application/x-rtf",
		".rv":   "video/vnd.rn-realvideo",
		".sam":  "application/x-sam",
		".sat":  "application/x-sat",
		".sdp":  "application/sdp",
		".sdw":  "application/x-sdw",
		".sit":  "application/x-stuffit",
		".slb":  "application/x-slb",
		".sld":  "application/x-sld",
		".slk":  "drawing/x-slk",
		".smi":  "application/smil",
		".smil": "application/smil",
		".smk":  "application/x-smk",
		".snd":  "audio/basic",
		".sol":  "text/plain",
		".sor":  "text/plain",
		".spc":  "application/x-pkcs7-certificates",
		".spl":  "application/futuresplash",
		".spp":  "text/xml",
		".ssm":  "application/streamingmedia",
		".sst":  "application/vnd.ms-pki.certstore",
		".stl":  "application/vnd.ms-pki.stl",
		".stm":  "text/html",
		".sty":  "application/x-sty",
		".svg":  "text/xml",
		".swf":  "application/x-shockwave-flash",
		".tdf":  "application/x-tdf",
		".tg4":  "application/x-tg4",
		".tga":  "application/x-tga",
		".tif":  "image/tiff",
		//".tif":"application/x-tif",
		".tiff":    "image/tiff",
		".tld":     "text/xml",
		".top":     "drawing/x-top",
		".torrent": "application/x-bittorrent",
		".tsd":     "text/xml",
		".txt":     "text/plain",
		".uin":     "application/x-icq",
		".uls":     "text/iuls",
		".vcf":     "text/x-vcard",
		".vda":     "application/x-vda",
		".vdx":     "application/vnd.visio",
		".vml":     "text/xml",
		".vpg":     "application/x-vpeg005",
		".vsd":     "application/vnd.visio",
		//".vsd":"application/x-vsd",
		".vss": "application/vnd.visio",
		".vst": "application/vnd.visio",
		//".vst":"application/x-vst",
		".vsw":   "application/vnd.visio",
		".vsx":   "application/vnd.visio",
		".vtx":   "application/vnd.visio",
		".vxml":  "text/xml",
		".wav":   "audio/wav",
		".wax":   "audio/x-ms-wax",
		".wb1":   "application/x-wb1",
		".wb2":   "application/x-wb2",
		".wb3":   "application/x-wb3",
		".wbmp":  "image/vnd.wap.wbmp",
		".wiz":   "application/msword",
		".wk3":   "application/x-wk3",
		".wk4":   "application/x-wk4",
		".wkq":   "application/x-wkq",
		".wks":   "application/x-wks",
		".wm":    "video/x-ms-wm",
		".wma":   "audio/x-ms-wma",
		".wmd":   "application/x-ms-wmd",
		".wmf":   "application/x-wmf",
		".wml":   "text/vnd.wap.wml",
		".wmv":   "video/x-ms-wmv",
		".wmx":   "video/x-ms-wmx",
		".wmz":   "application/x-ms-wmz",
		".wp6":   "application/x-wp6",
		".wpd":   "application/x-wpd",
		".wpg":   "application/x-wpg",
		".wpl":   "application/vnd.ms-wpl",
		".wq1":   "application/x-wq1",
		".wr1":   "application/x-wr1",
		".wri":   "application/x-wri",
		".wrk":   "application/x-wrk",
		".ws":    "application/x-ws",
		".ws2":   "application/x-ws",
		".wsc":   "text/scriptlet",
		".wsdl":  "text/xml",
		".wvx":   "video/x-ms-wvx",
		".xdp":   "application/vnd.adobe.xdp",
		".xdr":   "text/xml",
		".xfd":   "application/vnd.adobe.xfd",
		".xfdf":  "application/vnd.adobe.xfdf",
		".xhtml": "text/html",
		".xls":   "application/vnd.ms-excel",
		//".xls":"application/x-xls",
		".xlw":    "application/x-xlw",
		".xml":    "text/xml",
		".xpl":    "audio/scpls",
		".xq":     "text/xml",
		".xql":    "text/xml",
		".xquery": "text/xml",
		".xsd":    "text/xml",
		".xsl":    "text/xml",
		".xslt":   "text/xml",
		".xwd":    "application/x-xwd",
		".x_b":    "application/x-x_b",
		".sis":    "application/vnd.symbian.install",
		".sisx":   "application/vnd.symbian.install",
		".x_t":    "application/x-x_t",
		".ipa":    "application/vnd.iphone",
		".apk":    "application/vnd.android.package-archive",
		".xap":    "application/x-silverlight-app",
	}
	contentValMap = map[string][]string{
		"application/octet-stream": {".*"},
		"application/x-001":        {".001"},
		"application/x-301":        {".301"},
		"text/h323":                {".323"},
		"application/x-906":        {".906"},
		"drawing/907":              {".907"},
		"application/x-a11":        {".a11"},
		"audio/x-mei-aac":          {".acp"},
		"application/postscript":   {".ai", ".eps", ".ps"},

		"audio/aiff":                              {".aif", ".aifc", ".aiff"},
		"application/x-anv":                       {".anv"},
		"text/asa":                                {".asa"},
		"text/asp":                                {".asp"},
		"video/x-ms-asf":                          {".asx", ".asf"},
		"audio/basic":                             {".au", ".snd"},
		"video/avi":                               {".avi"},
		"application/vnd.adobe.workflow":          {".awf"},
		"application/x-bmp":                       {".bmp"},
		"application/x-bot":                       {".bot"},
		"application/x-c4t":                       {".c4t"},
		"application/x-c90":                       {".c90"},
		"application/x-cals":                      {".cal"},
		"application/vnd.ms-pki.seccat":           {".cat"},
		"application/x-netcdf":                    {".cdf"},
		"application/x-cdr":                       {".cdr"},
		"application/x-cel":                       {".cel"},
		"application/x-x509-ca-cert":              {".cer", ".crt", ".der"},
		"application/x-g4":                        {".cg4", ".g4", ".ig4"},
		"application/x-cgm":                       {".cgm"},
		"application/x-cit":                       {".cit"},
		"application/x-cmp":                       {".cmp"},
		"application/x-cmx":                       {".cmx"},
		"application/x-cot":                       {".cot"},
		"application/pkix-crl":                    {".crl"},
		"application/x-csi":                       {".csi"},
		"text/css":                                {".css"},
		"application/x-cut":                       {".cut"},
		"application/x-dbf":                       {".dbf"},
		"application/x-dbm":                       {".dbm"},
		"application/x-dbx":                       {".dbx"},
		"application/x-dcx":                       {".dcx"},
		"java/*":                                  {".class", ".java"},
		"application/x-dgn":                       {".dgn"},
		"application/x-dib":                       {".dib"},
		"application/x-msdownload":                {".dll", ".exe"},
		"application/msword":                      {".doc", ".dot", ".rtf", ".wiz"},
		"application/x-drw":                       {".drw"},
		"Model/vnd.dwf":                           {".dwf"},
		"application/x-dwf":                       {".dwf"},
		"application/x-dwg":                       {".dwg"},
		"application/x-dxb":                       {".dxb"},
		"application/x-dxf":                       {".dxf"},
		"application/vnd.adobe.edn":               {".edn"},
		"application/x-emf":                       {".emf"},
		"message/rfc822":                          {".eml", ".mht", ".mhtml", ".nws"},
		"text/xml":                                {".cml", ".biz", ".dcd", ".dtd", ".ent", ".fo", ".math", ".mml", ".mtx", ".rdf", ".spp", ".svg", ".tld", ".tsd", ".vml", ".vxml", ".wsdl", ".xdr", ".xml", ".xq", ".xql", ".xquery", ".xsd", ".xsl", ".xslt"},
		"application/x-epi":                       {".epi"},
		"application/x-ps":                        {".eps", ".ps"},
		"application/x-ebx":                       {".etd"},
		"image/fax":                               {".fax"},
		"application/vnd.fdf":                     {".fdf"},
		"application/fractals":                    {".fif"},
		"application/x-frm":                       {".frm"},
		"application/x-gbr":                       {".gbr"},
		"application/x-":                          {"."},
		"image/gif":                               {".gif"},
		"application/x-gl2":                       {".gl2"},
		"application/x-gp4":                       {".gp4"},
		"application/x-hgl":                       {".hgl"},
		"application/x-hmr":                       {".hmr"},
		"application/x-hpgl":                      {".hpg"},
		"application/x-hpl":                       {".hpl"},
		"application/mac-binhex40":                {".hqx"},
		"application/x-hrf":                       {".hrf"},
		"application/hta":                         {".hta"},
		"text/x-component":                        {".htc"},
		"text/html":                               {".htm", ".html", ".htx", ".jsp", ".plg", ".stm", ".xhtml"},
		"text/webviewhtml":                        {".htt"},
		"application/x-icb":                       {".icb"},
		"image/x-icon":                            {".ico"},
		"application/x-ico":                       {".ico"},
		"application/x-iff":                       {".iff"},
		"application/x-igs":                       {".igs"},
		"application/x-iphone":                    {".iii"},
		"application/x-img":                       {".img"},
		"application/x-internet-signup":           {".ins", ".isp"},
		"video/x-ivf":                             {".IVF"},
		"image/jpeg":                              {".jfif", ".jpe", ".jpeg", ".jpg"},
		"application/x-jpe":                       {".jpe"},
		"application/x-jpg":                       {".jpg"},
		"application/x-javascript":                {".js", ".ls", ".mocha"},
		"audio/x-liquid-file":                     {".la1"},
		"application/x-laplayer-reg":              {".lar"},
		"application/x-latex":                     {".latex"},
		"audio/x-liquid-secure":                   {".lavs"},
		"application/x-lbm":                       {".lbm"},
		"audio/x-la-lms":                          {".lmsff"},
		"application/x-ltr":                       {".ltr"},
		"video/x-mpeg":                            {".m1v", ".m2v", ".mpe", ".mps"},
		"audio/mpegurl":                           {".m3u"},
		"video/mpeg4":                             {".mp4", ".m4e"},
		"application/x-mac":                       {".mac"},
		"application/x-troff-man":                 {".man"},
		"application/msaccess":                    {".mdb"},
		"application/x-mdb":                       {".mdb"},
		"application/x-shockwave-flash":           {".swf", ".mfp"},
		"application/x-mi":                        {".mi"},
		"audio/mid":                               {".mid", ".midi", ".rmi"},
		"application/x-mil":                       {".mil"},
		"audio/x-musicnet-download":               {".mnd"},
		"audio/x-musicnet-stream":                 {".mns"},
		"video/x-sgi-movie":                       {".movie"},
		"audio/mp1":                               {".mp1"},
		"audio/mp2":                               {".mp2"},
		"video/mpeg":                              {".mp2v", ".mpv2"},
		"audio/mp3":                               {".mp3"},
		"video/x-mpg":                             {".mpa"},
		"application/vnd.ms-project":              {".mpd", ".mpp", ".mpt", ".mpw", ".mpx"},
		"video/mpg":                               {".mpeg", ".mpg", ".mpv"},
		"audio/rn-mpeg":                           {".mpga"},
		"application/x-mmxp":                      {".mxp"},
		"image/pnetvue":                           {".net"},
		"application/x-nrf":                       {".nrf"},
		"text/x-ms-odc":                           {".odc"},
		"application/x-out":                       {".out"},
		"application/pkcs10":                      {".p10"},
		"application/x-pkcs12":                    {".p12", ".pfx"},
		"application/x-pkcs7-certificates":        {".p7b", ".spc"},
		"application/pkcs7-mime":                  {".p7c", ".p7m"},
		"application/x-pkcs7-certreqresp":         {".p7r"},
		"application/pkcs7-signature":             {".p7s"},
		"application/x-pc5":                       {".pc5"},
		"application/x-pci":                       {".pci"},
		"application/x-pcl":                       {".pcl"},
		"application/x-pcx":                       {".pcx"},
		"application/pdf":                         {".pdf"},
		"application/vnd.adobe.pdx":               {".pdx"},
		"application/x-pgl":                       {".pgl"},
		"application/x-pic":                       {".pic"},
		"application/vnd.ms-pki.pko":              {".pko"},
		"application/x-perl":                      {".pl"},
		"audio/scpls":                             {".pls", ".xpl"},
		"application/x-plt":                       {".plt"},
		"image/png":                               {".png"},
		"application/x-png":                       {".png"},
		"application/vnd.ms-powerpoint":           {".ppt", ".pot", ".ppa", ".pps", ".pwz"},
		"application/x-ppm":                       {".ppm"},
		"application/x-ppt":                       {".ppt"},
		"application/x-pr":                        {".pr"},
		"application/pics-rules":                  {".prf"},
		"application/x-prn":                       {".prn"},
		"application/x-prt":                       {".prt"},
		"application/x-ptn":                       {".ptn"},
		"text/vnd.rn-realtext3d":                  {".r3t"},
		"audio/vnd.rn-realaudio":                  {".ra"},
		"audio/x-pn-realaudio":                    {".ram", ".rmm"},
		"application/x-ras":                       {".ras"},
		"application/rat-file":                    {".rat"},
		"application/vnd.rn-recording":            {".rec"},
		"application/x-red":                       {".red"},
		"application/x-rgb":                       {".rgb"},
		"application/vnd.rn-realsystem-rjs":       {".rjs"},
		"application/vnd.rn-realsystem-rjt":       {".rjt"},
		"application/x-rlc":                       {".rlc"},
		"application/x-rle":                       {".rle"},
		"application/vnd.rn-realmedia":            {".rm"},
		"application/vnd.adobe.rmf":               {".rmf"},
		"application/vnd.rn-realsystem-rmj":       {".rmj"},
		"application/vnd.rn-rn_music_package":     {".rmp"},
		"application/vnd.rn-realmedia-secure":     {".rms"},
		"application/vnd.rn-realmedia-vbr":        {".rmvb"},
		"application/vnd.rn-realsystem-rmx":       {".rmx"},
		"application/vnd.rn-realplayer":           {".rnx"},
		"image/vnd.rn-realpix":                    {".rp"},
		"audio/x-pn-realaudio-plugin":             {".rpm"},
		"application/vnd.rn-rsml":                 {".rsml"},
		"text/vnd.rn-realtext":                    {".rt"},
		"application/x-rtf":                       {".rtf"},
		"video/vnd.rn-realvideo":                  {".rv"},
		"application/x-sam":                       {".sam"},
		"application/x-sat":                       {".sat"},
		"application/sdp":                         {".sdp"},
		"application/x-sdw":                       {".sdw"},
		"application/x-stuffit":                   {".sit"},
		"application/x-slb":                       {".slb"},
		"application/x-sld":                       {".sld"},
		"drawing/x-slk":                           {".slk"},
		"application/smil":                        {".smi", ".smil"},
		"application/x-smk":                       {".smk"},
		"text/plain":                              {".txt", ".sol", ".sor"},
		"application/futuresplash":                {".spl"},
		"application/streamingmedia":              {".ssm"},
		"application/vnd.ms-pki.certstore":        {".sst"},
		"application/vnd.ms-pki.stl":              {".stl"},
		"application/x-sty":                       {".sty"},
		"application/x-tdf":                       {".tdf"},
		"application/x-tg4":                       {".tg4"},
		"application/x-tga":                       {".tga"},
		"image/tiff":                              {".tif", ".tiff"},
		"application/x-tif":                       {".tif"},
		"drawing/x-top":                           {".top"},
		"application/x-bittorrent":                {".torrent"},
		"application/x-icq":                       {".uin"},
		"text/iuls":                               {".uls"},
		"text/x-vcard":                            {".vcf"},
		"application/x-vda":                       {".vda"},
		"application/vnd.visio":                   {".vdx", ".vsd", ".vss", ".vst", ".vsw", ".vsx", ".vtx"},
		"application/x-vpeg005":                   {".vpg"},
		"application/x-vsd":                       {".vsd"},
		"application/x-vst":                       {".vst"},
		"audio/wav":                               {".wav"},
		"audio/x-ms-wax":                          {".wax"},
		"application/x-wb1":                       {".wb1"},
		"application/x-wb2":                       {".wb2"},
		"application/x-wb3":                       {".wb3"},
		"image/vnd.wap.wbmp":                      {".wbmp"},
		"application/x-wk3":                       {".wk3"},
		"application/x-wk4":                       {".wk4"},
		"application/x-wkq":                       {".wkq"},
		"application/x-wks":                       {".wks"},
		"video/x-ms-wm":                           {".wm"},
		"audio/x-ms-wma":                          {".wma"},
		"application/x-ms-wmd":                    {".wmd"},
		"application/x-wmf":                       {".wmf"},
		"text/vnd.wap.wml":                        {".wml"},
		"video/x-ms-wmv":                          {".wmv"},
		"video/x-ms-wmx":                          {".wmx"},
		"application/x-ms-wmz":                    {".wmz"},
		"application/x-wp6":                       {".wp6"},
		"application/x-wpd":                       {".wpd"},
		"application/x-wpg":                       {".wpg"},
		"application/vnd.ms-wpl":                  {".wpl"},
		"application/x-wq1":                       {".wq1"},
		"application/x-wr1":                       {".wr1"},
		"application/x-wri":                       {".wri"},
		"application/x-wrk":                       {".wrk"},
		"application/x-ws":                        {".ws", ".ws2"},
		"text/scriptlet":                          {".wsc"},
		"video/x-ms-wvx":                          {".wvx"},
		"application/vnd.adobe.xdp":               {".xdp"},
		"application/vnd.adobe.xfd":               {".xfd"},
		"application/vnd.adobe.xfdf":              {".xfdf"},
		"application/vnd.ms-excel":                {".xls"},
		"application/x-xls":                       {".xls"},
		"application/x-xlw":                       {".xlw"},
		"application/x-xwd":                       {".xwd"},
		"application/x-x_b":                       {".x_b"},
		"application/vnd.symbian.install":         {".sis", ".sisx"},
		"application/x-x_t":                       {".x_t"},
		"application/vnd.iphone":                  {".ipa"},
		"application/vnd.android.package-archive": {".apk"},
		"application/x-silverlight-app":           {".xap"},
	}
}