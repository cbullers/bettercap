package network

var portServices = map[string]map[int]string{
	"tcp": map[int]string{
		6566:  "sane-port",
		1094:  "rootd",
		4557:  "fax",
		10000: "webmin",
		106:   "poppassd",
		5674:  "mrtd",
		98:    "linuxconf",
		371:   "clearcase",
		873:   "rsync",
		5269:  "xmpp-server",
		15345: "xpilot",
		143:   "imap2",
		345:   "pawserv",
		607:   "nqs",
		513:   "login",
		532:   "netnews",
		5002:  "rfe",
		20:    "ftp-data",
		70:    "gopher",
		10082: "amandaidx",
		13720: "bprd",
		871:   "supfilesrv",
		1:     "tcpmux",
		2135:  "gris",
		3260:  "iscsi-target",
		2607:  "ospfapi",
		370:   "codaauth2",
		546:   "dhcpv6-client",
		1813:  "radius-acct",
		775:   "moira-db",
		23:    "telnet",
		129:   "pwdgen",
		5050:  "mmcc",
		5190:  "aol",
		8140:  "puppet",
		22128: "gsidcap",
		2603:  "ripngd",
		9098:  "xinetd",
		1812:  "radius",
		3689:  "daap",
		24554: "binkp",
		2000:  "cisco-sccp",
		13:    "daytime",
		177:   "xdmcp",
		6005:  "x11-5",
		9101:  "bacula-dir",
		8081:  "tproxy",
		1080:  "socks",
		5061:  "sip-tls",
		7100:  "font-service",
		5051:  "enbd-cstatd",
		10081: "kamanda",
		178:   "nextstep",
		3050:  "gds-db",
		7002:  "afs3-prserver",
		11371: "hkp",
		115:   "sftp",
		5308:  "cfengine",
		547:   "dhcpv6-server",
		1959:  "remoteping",
		13782: "bpcd",
		1529:  "support",
		102:   "iso-tsap",
		163:   "cmip-man",
		2431:  "venus-se",
		22273: "wnn6",
		4559:  "hylafax",
		19:    "chargen",
		515:   "printer",
		2105:  "eklogin",
		8990:  "clc-build-daemon",
		2433:  "codasrv-se",
		4369:  "epmd",
		531:   "conference",
		5688:  "ggz",
		1236:  "rmtcfg",
		2600:  "zebrasrv",
		107:   "rtelnet",
		465:   "urd",
		11:    "systat",
		5052:  "enbd-sstatd",
		7003:  "afs3-vlserver",
		210:   "z3950",
		538:   "gdomap",
		13721: "bpdbm",
		161:   "snmp",
		2010:  "pipe-server",
		655:   "tinc",
		1863:  "msnp",
		2947:  "gpsd",
		777:   "moira-update",
		43:    "whois",
		389:   "ldap",
		10809: "nbd",
		808:   "omirr",
		2003:  "cfinger",
		60177: "tfido",
		135:   "loc-srv",
		139:   "netbios-ssn",
		2150:  "ninstall",
		20012: "vboxd",
		22:    "ssh",
		5432:  "postgresql",
		5222:  "xmpp-client",
		9:     "discard",
		347:   "fatserv",
		4224:  "xtell",
		6514:  "syslog-tls",
		27374: "asp",
		1127:  "supfiledbg",
		2606:  "ospf6d",
		2605:  "bgpd",
		4949:  "munin",
		57000: "dircproxy",
		444:   "snpp",
		9102:  "bacula-fd",
		49:    "tacacs",
		5672:  "amqp",
		7006:  "afs3-errors",
		8080:  "http-alt",
		2989:  "afmbackup",
		5680:  "canna",
		7:     "echo",
		17:    "qotd",
		995:   "pop3s",
		2601:  "zebra",
		9418:  "git",
		110:   "pop3",
		138:   "netbios-dgm",
		3493:  "nut",
		4031:  "suucp",
		4353:  "f5-iquery",
		7000:  "afs3-fileserver",
		15:    "netstat",
		213:   "ipx",
		990:   "ftps",
		1434:  "ms-sql-m",
		2053:  "knetd",
		101:   "hostnames",
		346:   "zserv",
		1701:  "l2f",
		87:    "link",
		1646:  "sa-msg-port",
		5353:  "mdns",
		6002:  "x11-2",
		6444:  "sge-qmaster",
		512:   "exec",
		1093:  "proofd",
		9673:  "zope",
		406:   "imsp",
		610:   "npmp-local",
		1241:  "nessus",
		95:    "supdup",
		530:   "courier",
		2792:  "f5-globalsite",
		6347:  "gnutella-rtr",
		5667:  "nsca",
		17004: "sgi-cad",
		206:   "at-zis",
		2628:  "dict",
		2111:  "kx",
		1314:  "xtelw",
		2608:  "isisd",
		18:    "msp",
		68:    "bootpc",
		21:    "ftp",
		174:   "mailq",
		4094:  "sysrqd",
		8088:  "omniorb",
		201:   "at-rtmp",
		612:   "hmmp-ind",
		751:   "kerberos-master",
		749:   "kerberos-adm",
		10080: "amanda",
		760:   "krbupdate",
		209:   "qmtp",
		992:   "telnets",
		587:   "submission",
		1352:  "lotusnote",
		1645:  "datametrics",
		2583:  "mon",
		1001:  "customs",
		6667:  "ircd",
		199:   "smux",
		514:   "shell",
		2811:  "gsiftp",
		6007:  "x11-7",
		204:   "at-echo",
		1958:  "log-server",
		123:   "ntp",
		202:   "at-nbp",
		1957:  "unix-status",
		2101:  "rtcm-sc104",
		2401:  "cvspserver",
		37:    "time",
		53:    "domain",
		194:   "irc",
		5556:  "freeciv",
		6003:  "x11-3",
		7009:  "afs3-rmtsys",
		11112: "dicom",
		25:    "smtp",
		42:    "nameserver",
		540:   "uucp",
		6445:  "sge-execd",
		7004:  "afs3-kaserver",
		10051: "zabbix-trapper",
		22125: "dcap",
		2604:  "ospfd",
		164:   "cmip-agent",
		427:   "svrloc",
		13722: "bpjava-msvc",
		4899:  "radmin-port",
		6004:  "x11-4",
		628:   "qmqp",
		556:   "remotefs",
		636:   "ldaps",
		1524:  "ingreslock",
		2432:  "codasrv",
		1109:  "kpop",
		372:   "ulistserv",
		443:   "https",
		2121:  "frox",
		5675:  "bgpsim",
		13783: "vopied",
		17500: "db-lsp",
		179:   "bgp",
		1433:  "ms-sql-s",
		2086:  "gnunet",
		9667:  "xmms2",
		1313:  "xtel",
		65:    "tacacs-ds",
		113:   "auth",
		1194:  "openvpn",
		7001:  "afs3-callback",
		1178:  "skkserv",
		104:   "acr-nema",
		487:   "saft",
		5151:  "pcrd",
		11201: "smsqp",
		989:   "ftps-data",
		4373:  "remctl",
		7005:  "afs3-volser",
		901:   "swat",
		783:   "spamd",
		5355:  "hostmon",
		3306:  "mysql",
		3632:  "distcc",
		2119:  "gsigatekeeper",
		2430:  "venus",
		3205:  "isns",
		79:    "finger",
		369:   "rpc2portmap",
		500:   "isakmp",
		1099:  "rmiregistry",
		30865: "csync2",
		60179: "fido",
		50:    "re-mail-ck",
		67:    "bootps",
		7008:  "afs3-update",
		5354:  "noclog",
		563:   "nntps",
		2049:  "nfs",
		6001:  "x11-1",
		6346:  "gnutella-svc",
		6446:  "mysql-proxy",
		750:   "kerberos4",
		754:   "krb-prop",
		1300:  "wipld",
		611:   "npmp-gui",
		765:   "webster",
		8021:  "zope-ftp",
		6000:  "x11",
		6697:  "ircs-u",
		10050: "zabbix-agent",
		20011: "isdnlog",
		445:   "microsoft-ds",
		5060:  "sip",
		1214:  "kazaa",
		6006:  "x11-6",
		543:   "klogin",
		549:   "idfp",
		706:   "silc",
		993:   "imaps",
		1649:  "kermit",
		1677:  "groupwise",
		3130:  "icpv2",
		7007:  "afs3-bos",
		80:    "http",
		162:   "snmp-trap",
		2602:  "ripd",
		10083: "amidxtape",
		13724: "vnetd",
		2988:  "afbackup",
		464:   "kpasswd",
		4569:  "iax",
		137:   "netbios-ns",
		554:   "rtsp",
		544:   "kshell",
		4190:  "sieve",
		9103:  "bacula-sd",
		5666:  "nrpe",
		111:   "sunrpc",
		119:   "nntp",
		631:   "ipp",
		526:   "tempo",
		548:   "afpovertcp",
		3690:  "svn",
		4691:  "mtn",
		5671:  "amqps",
		88:    "kerberos",
		105:   "csnet-ns",
		4600:  "distmp3",
	},
	"udp": map[int]string{
		2086:  "gnunet",
		3493:  "nut",
		4373:  "remctl",
		2988:  "afbackup",
		17003: "sgi-gcd",
		13:    "daytime",
		636:   "ldaps",
		1080:  "socks",
		370:   "codaauth2",
		1433:  "ms-sql-s",
		1645:  "datametrics",
		6004:  "x11-4",
		518:   "ntalk",
		533:   "netwall",
		164:   "cmip-agent",
		2792:  "f5-globalsite",
		107:   "rtelnet",
		610:   "npmp-local",
		517:   "talk",
		67:    "bootps",
		406:   "imsp",
		611:   "npmp-gui",
		2583:  "mon",
		22273: "wnn6",
		547:   "dhcpv6-server",
		1352:  "lotusnote",
		129:   "pwdgen",
		209:   "qmtp",
		369:   "rpc2portmap",
		5190:  "aol",
		7001:  "afs3-callback",
		135:   "loc-srv",
		7000:  "afs3-fileserver",
		7002:  "afs3-prserver",
		68:    "bootpc",
		194:   "irc",
		2135:  "gris",
		39:    "rlp",
		105:   "csnet-ns",
		751:   "kerberos-master",
		106:   "poppassd",
		623:   "asf-rmcp",
		512:   "biff",
		5060:  "sip",
		18:    "msp",
		1649:  "kermit",
		5269:  "xmpp-server",
		5556:  "freeciv",
		7009:  "afs3-rmtsys",
		750:   "kerberos4",
		779:   "moira-ureg",
		201:   "at-rtmp",
		1093:  "proofd",
		1214:  "kazaa",
		1210:  "predict",
		13722: "bpjava-msvc",
		15345: "xpilot",
		13720: "bprd",
		1001:  "customs",
		4031:  "suucp",
		1813:  "radius-acct",
		2401:  "cvspserver",
		7:     "echo",
		210:   "z3950",
		4094:  "sysrqd",
		7006:  "afs3-errors",
		2432:  "codasrv",
		3689:  "daap",
		3690:  "svn",
		808:   "omirr",
		11371: "hkp",
		372:   "ulistserv",
		1812:  "radius",
		9667:  "xmms2",
		6346:  "gnutella-svc",
		123:   "ntp",
		347:   "fatserv",
		5555:  "rplay",
		752:   "passwd-server",
		2989:  "afmbackup",
		464:   "kpasswd",
		554:   "rtsp",
		520:   "route",
		612:   "hmmp-ind",
		2430:  "venus",
		69:    "tftp",
		138:   "netbios-dgm",
		345:   "pawserv",
		2947:  "gpsd",
		27374: "asp",
		7100:  "font-service",
		10051: "zabbix-trapper",
		13721: "bpdbm",
		163:   "cmip-man",
		500:   "isakmp",
		204:   "at-echo",
		2104:  "zephyr-hm",
		20011: "isdnlog",
		1194:  "openvpn",
		2049:  "nfs",
		2101:  "rtcm-sc104",
		10081: "kamanda",
		427:   "svrloc",
		546:   "dhcpv6-client",
		4500:  "ipsec-nat-t",
		6003:  "x11-3",
		53:    "domain",
		174:   "mailq",
		5353:  "mdns",
		6006:  "x11-6",
		2103:  "zephyr-clt",
		444:   "snpp",
		2628:  "dict",
		1241:  "nessus",
		162:   "snmp-trap",
		202:   "at-nbp",
		513:   "who",
		525:   "timed",
		1094:  "rootd",
		4899:  "radmin-port",
		5432:  "postgresql",
		19:    "chargen",
		4569:  "iax",
		13783: "vopied",
		9102:  "bacula-fd",
		2150:  "ninstall",
		111:   "sunrpc",
		213:   "ipx",
		3632:  "distcc",
		6001:  "x11-1",
		37:    "time",
		20012: "vboxd",
		7003:  "afs3-vlserver",
		7008:  "afs3-update",
		10050: "zabbix-agent",
		13724: "vnetd",
		2811:  "gsiftp",
		538:   "gdomap",
		17001: "sgi-cmsd",
		21:    "fsp",
		5672:  "amqp",
		65:    "tacacs-ds",
		548:   "afpovertcp",
		5688:  "ggz",
		6444:  "sge-qmaster",
		9101:  "bacula-dir",
		2119:  "gsigatekeeper",
		6002:  "x11-2",
		10080: "amanda",
		9359:  "mandelspawn",
		2433:  "codasrv-se",
		11201: "smsqp",
		4353:  "f5-iquery",
		9103:  "bacula-sd",
		6445:  "sge-execd",
		88:    "kerberos",
		2431:  "venus-se",
		4369:  "epmd",
		3306:  "mysql",
		49:    "tacacs",
		161:   "snmp",
		2000:  "cisco-sccp",
		9:     "discard",
		3050:  "gds-db",
		6000:  "x11",
		514:   "syslog",
		1701:  "l2f",
		8080:  "http-alt",
		5354:  "noclog",
		178:   "nextstep",
		389:   "ldap",
		607:   "nqs",
		655:   "tinc",
		1099:  "rmiregistry",
		1863:  "msnp",
		346:   "zserv",
		6446:  "mysql-proxy",
		7007:  "afs3-bos",
		2102:  "zephyr-srv",
		5061:  "sip-tls",
		7004:  "afs3-kaserver",
		17002: "sgi-crsd",
		445:   "microsoft-ds",
		549:   "idfp",
		1524:  "ingreslock",
		371:   "clearcase",
		631:   "ipp",
		1434:  "ms-sql-m",
		5002:  "rfe",
		199:   "smux",
		5355:  "hostmon",
		6696:  "babel",
		104:   "acr-nema",
		487:   "saft",
		5050:  "mmcc",
		3130:  "icpv2",
		765:   "webster",
		1646:  "sa-msg-port",
		13782: "bpcd",
		139:   "netbios-ssn",
		3205:  "isns",
		5308:  "cfengine",
		6005:  "x11-5",
		6347:  "gnutella-rtr",
		706:   "silc",
		1677:  "groupwise",
		7005:  "afs3-volser",
		628:   "qmqp",
		4691:  "mtn",
		5222:  "xmpp-client",
		50:    "re-mail-ck",
		6007:  "x11-7",
		137:   "netbios-ns",
		8088:  "omniorb",
		177:   "xdmcp",
		206:   "at-zis",
	},
}

func GetServiceByPort(port int, proto string) string {
	if p, found := portServices[proto]; found {
		if svc, found := p[port]; found {
			return svc
		}
	}
	return ""
}
