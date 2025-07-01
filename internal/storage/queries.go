package storage

const (
	fieldsCDR = " calldate,clid,src,dst,dcontext,channel,dstchannel,lastapp,lastdata,duration,billsec,disposition,amaflags,accountcode,userfield,uniqueid "
	queryCDR  = "INSERT IGNORE INTO " +
		"asterisk.cdr " +
		"(" + fieldsCDR + ") " +
		"SELECT " + fieldsCDR +
		"FROM asterisk_backup.cdr"

	fieldsExtensionsTable = " id,context,exten,app,appdata "
	queryExtensionsTable  = "INSERT IGNORE INTO " +
		"asterisk.extensions_table " +
		"(" + fieldsExtensionsTable + ") " +
		"SELECT " + fieldsExtensionsTable +
		"FROM asterisk_backup.extensions_table"

	fieldsIVRBody = " id,major,inc_number,check_time,active_time,active,ivr_comment,wait_exten,audio_id,holiday_id,setcid,disa "
	queryIVRBody  = "INSERT IGNORE INTO " +
		"asterisk.ivr_body " +
		"(" + fieldsIVRBody + ") " +
		"SELECT " + fieldsIVRBody +
		"FROM asterisk_backup.ivr_body"

	fieldsIVREvents = " id,ivr_id,`key`,action,parameters "
	queryIVREvents  = "INSERT IGNORE INTO " +
		"asterisk.ivr_events " +
		"(" + fieldsIVREvents + ") " +
		"SELECT " + fieldsIVREvents +
		"FROM asterisk_backup.ivr_events"

	fieldsIVRHolidays = " id,comment "
	queryIVRHolidays  = "INSERT IGNORE INTO " +
		"asterisk.ivr_holidays " +
		"(" + fieldsIVRHolidays + ") " +
		"SELECT " + fieldsIVRHolidays +
		"FROM asterisk_backup.ivr_holidays"

	fieldsMusicOnHold = " name,directory,application,mode,digit,sort,format,description,sda "
	queryMusicOnHold  = "INSERT IGNORE INTO " +
		"asterisk.musiconhold " +
		"(" + fieldsMusicOnHold + ") " +
		"SELECT " + fieldsMusicOnHold +
		"FROM asterisk_backup.musiconhold"

	fieldsQueueLog = " time,callid,queuename,agent,event,position,origposition,waittime,channel,logintime,holdtime,calltime,callerid,url,`key`,penalty,ringtime,exten,context "
	queryQueueLog  = "INSERT IGNORE INTO " +
		"asterisk.queue_log " +
		"(" + fieldsQueueLog + ") " +
		"SELECT " + fieldsQueueLog +
		"FROM asterisk_backup.queue_log"

	fieldsQueueMemberTable = " uniqueid,membername,queue_name,interface,penalty,paused "
	queryQueueMemberTable  = "INSERT IGNORE INTO " +
		"asterisk.queue_member_table " +
		"(" + fieldsQueueMemberTable + ") " +
		"SELECT " + fieldsQueueMemberTable +
		"FROM asterisk_backup.queue_member_table"

	fieldsQueuePenalty = " interface,queue,penalty "
	queryQueuePenalty  = "INSERT IGNORE INTO " +
		"asterisk.queue_penalty " +
		"(" + fieldsQueuePenalty + ") " +
		"SELECT " + fieldsQueuePenalty +
		"FROM asterisk_backup.queue_penalty"

	fieldsQueueTable = " name,musiconhold,announce,context,timeout,monitor_join,monitor_format,queue_youarenext,queue_thereare,queue_callswaiting,queue_holdtime,queue_minutes,queue_seconds,queue_lessthan,queue_thankyou,queue_reporthold,announce_frequency,announce_round_seconds,announce_holdtime,retry,wrapuptime,maxlen,servicelevel,strategy,joinempty,leavewhenempty,eventmemberstatus,eventwhencalled,reportholdtime,memberdelay,weight,timeoutrestart,periodic_announce,periodic_announce_frequency,ringinuse,description,`escalate-if-not-available`,wait_timeout,exit_exten "
	queryQueueTable  = "INSERT IGNORE INTO " +
		"asterisk.queue_table " +
		"(" + fieldsQueueTable + ") " +
		"SELECT " + fieldsQueueTable +
		"FROM asterisk_backup.queue_table"

	fieldsRecLevels = " user_id,subscriber_name "
	queryRecLevels  = "INSERT IGNORE INTO " +
		"asterisk.rec_levels " +
		"(" + fieldsRecLevels + ") " +
		"SELECT " + fieldsRecLevels +
		"FROM asterisk_backup.rec_levels"

	fieldsSipfriends = " name,type,defaultuser,fromuser,fromdomain,secret,md5secret,auth,mailbox,subscribemwi,vmexten,callerid,cid_number,callingpres,usereqphone,language,`call-limit`,context,dtmfmode,subscribecontext,amaflags,accountcode,musicclass,mohsuggest,allowtransfer,nat,callgroup,pickupgroup,autoframing,disallow,allow,maxcallbitrate,host,outboundproxy,ipaddr,defaultip,port,fullcontact,insecure,qualify,regseconds,regexten,regserver,rtptimeout,rtpholdtimeout,rtpkeepalive,lastms,setvar,useragent,t38pt_udptl,recs,afk,mobile "
	querySipfriends  = "INSERT IGNORE INTO " +
		"asterisk.sipfriends " +
		"(" + fieldsSipfriends + ") " +
		"SELECT " + fieldsSipfriends +
		"FROM asterisk_backup.sipfriends"

	fieldsSounds = " id,sound_comment,sound_file,sound_file_length "
	querySounds  = "INSERT IGNORE INTO " +
		"asterisk.sounds " +
		"(" + fieldsSounds + ") " +
		"SELECT " + fieldsSounds +
		"FROM asterisk_backup.sounds"

	fieldsSubscribers = " name,callerid,protocol,context,mobile,cfwd,recs,mailbox "
	querySubscribers  = "INSERT IGNORE INTO " +
		"asterisk.subscribers " +
		"(" + fieldsSubscribers + ") " +
		"SELECT " + fieldsSubscribers +
		"FROM asterisk_backup.subscribers"

	fieldsVoicemailUsers = " uniqueid, customer_id, context, mailbox, password, attach, email, fullname, pager, tz, saycid, dialout, callback, review, operator, envelope, sayduration, saydurationm, sendvoicemail, delete, nextaftercmd, forcename, forcegreetings, hidefromdir, stamp "
	queryVoicemailUsers  = "INSERT IGNORE INTO " +
		"asterisk.voicemail_users " +
		"(" + fieldsVoicemailUsers + ") " +
		"SELECT " + fieldsVoicemailUsers +
		"FROM asterisk_backup.voicemail_users"
)
