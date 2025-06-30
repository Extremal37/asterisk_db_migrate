package storage

const (
	queryCDR = "INSERT IGNORE INTO " +
		"asterisk.cdr " +
		"(calldate,clid,src,dst,dcontext,channel,dstchannel,lastapp,lastdata,duration,billsec,disposition,amaflags,accountcode,userfield,uniqueid) " +
		"SELECT * " +
		"FROM asterisk_backup.cdr"
	queryExtensionsTable = "INSERT IGNORE INTO " +
		"asterisk.extensions_table " +
		"SELECT * " +
		"FROM asterisk_backup.extensions_table"
	queryIVRBody = "INSERT IGNORE INTO " +
		"asterisk.ivr_body " +
		"SELECT * " +
		"FROM asterisk_backup.ivr_body"
	queryIVREvents = "INSERT IGNORE INTO " +
		"asterisk.ivr_events " +
		"SELECT * " +
		"FROM asterisk_backup.ivr_events"
	queryIVRHolidays = "INSERT IGNORE INTO " +
		"asterisk.ivr_holidays " +
		"SELECT * " +
		"FROM asterisk_backup.ivr_holidays"
	queryMusicOnHold = "INSERT IGNORE INTO " +
		"asterisk.musiconhold " +
		"SELECT * " +
		"FROM asterisk_backup.musiconhold"
	queryQueueLog = "INSERT IGNORE INTO " +
		"asterisk.queue_log(time,callid,queuename,agent,event,position,origposition,waittime,channel,logintime,holdtime,calltime,callerid,url,`key`,penalty,ringtime,exten,context)  " +
		"SELECT time,callid,queuename,agent,event,position,origposition,waittime,channel,logintime,holdtime,calltime,callerid,url,`key`,penalty,ringtime,exten,context " +
		"FROM asterisk_backup.queue_log"
	queryQueueMemberTable = "INSERT IGNORE INTO " +
		"asterisk.queue_member_table(uniqueid,membername,queue_name,interface,penalty,paused) " +
		"SELECT uniqueid,membername,queue_name,interface,penalty,paused " +
		"FROM asterisk_backup.queue_member_table"
	queryQueuePenalty = "INSERT IGNORE INTO " +
		"asterisk.queue_penalty " +
		"SELECT * FROM " +
		"asterisk_backup.queue_penalty"
	queryQueueTable = "INSERT IGNORE INTO " +
		"asterisk.queue_table(name,musiconhold,announce,context,timeout,monitor_join,monitor_format,queue_youarenext,queue_thereare,queue_callswaiting,queue_holdtime,queue_minutes,queue_seconds,queue_lessthan,queue_thankyou,queue_reporthold,announce_frequency,announce_round_seconds,announce_holdtime,retry,wrapuptime,maxlen,servicelevel,strategy,joinempty,leavewhenempty,eventmemberstatus,eventwhencalled,reportholdtime,memberdelay,weight,timeoutrestart,periodic_announce,periodic_announce_frequency,ringinuse,description,`escalate-if-not-available`,wait_timeout,exit_exten) " +
		"SELECT name,musiconhold,announce,context,timeout,monitor_join,monitor_format,queue_youarenext,queue_thereare,queue_callswaiting,queue_holdtime,queue_minutes,queue_seconds,queue_lessthan,queue_thankyou,queue_reporthold,announce_frequency,announce_round_seconds,announce_holdtime,retry,wrapuptime,maxlen,servicelevel,strategy,joinempty,leavewhenempty,eventmemberstatus,eventwhencalled,reportholdtime,memberdelay,weight,timeoutrestart,periodic_announce,periodic_announce_frequency,ringinuse,description,`escalate-if-not-available`,wait_timeout,exit_exten " +
		"FROM asterisk_backup.queue_table"
	queryRecLevels = "INSERT IGNORE INTO " +
		"asterisk.rec_levels " +
		"SELECT * FROM " +
		"asterisk_backup.rec_levels"
	querySipfriends = "INSERT IGNORE INTO " +
		"asterisk.sipfriends(name,type,defaultuser,fromuser,fromdomain,secret,md5secret,auth,mailbox,subscribemwi,vmexten,callerid,cid_number,callingpres,usereqphone,language,`call-limit`,context,dtmfmode,subscribecontext,amaflags,accountcode,musicclass,mohsuggest,allowtransfer,nat,callgroup,pickupgroup,autoframing,disallow,allow,maxcallbitrate,host,outboundproxy,ipaddr,defaultip,port,fullcontact,insecure,qualify,regseconds,regexten,regserver,rtptimeout,rtpholdtimeout,rtpkeepalive,lastms,setvar,useragent,t38pt_udptl,recs,afk,mobile) " +
		"SELECT name,type,defaultuser,fromuser,fromdomain,secret,md5secret,auth,mailbox,subscribemwi,vmexten,callerid,cid_number,callingpres,usereqphone,language,`call-limit`,context,dtmfmode,subscribecontext,amaflags,accountcode,musicclass,mohsuggest,allowtransfer,nat,callgroup,pickupgroup,autoframing,disallow,allow,maxcallbitrate,host,outboundproxy,ipaddr,defaultip,port,fullcontact,insecure,qualify,regseconds,regexten,regserver,rtptimeout,rtpholdtimeout,rtpkeepalive,lastms,setvar,useragent,t38pt_udptl,recs,afk,mobile " +
		"FROM asterisk_backup.sipfriends"
	querySounds = "INSERT IGNORE INTO " +
		"asterisk.sounds " +
		"SELECT * FROM " +
		"asterisk_backup.sounds"
	querySubscribers = "INSERT IGNORE INTO " +
		"asterisk.subscribers(name,callerid,protocol,context,mobile,cfwd,recs,mailbox) " +
		"SELECT name,callerid,protocol,context,mobile,cfwd,recs,mailbox " +
		"FROM asterisk_backup.subscribers"
	queryVoicemailUsers = "INSERT IGNORE INTO " +
		"asterisk.voicemail_users " +
		"SELECT * FROM " +
		"asterisk_backup.voicemail_users"
)
