package storage

import (
	"database/sql"
	"fmt"
	"github.com/Extremal37/asterisk_db_migrate/internal/logger"
	_ "github.com/go-sql-driver/mysql"
)

const (
	queryCDR              = "INSERT INTO asterisk.cdr (calldate,clid,src,dst,dcontext,channel,dstchannel,lastapp,lastdata,duration,billsec,disposition,amaflags,accountcode,userfield,uniqueid) SELECT * FROM asterisk_backup.cdr"
	queryExtensionsTable  = "INSERT INTO asterisk.extensions_table SELECT * FROM asterisk_backup.extensions_table"
	queryIVRBody          = "INSERT INTO asterisk.ivr_body SELECT * FROM asterisk_backup.ivr_body"
	queryIVREvents        = "INSERT INTO asterisk.ivr_events SELECT * FROM asterisk_backup.ivr_events"
	queryIVRHolidays      = "INSERT INTO asterisk.ivr_holidays SELECT * FROM asterisk_backup.ivr_holidays"
	queryMusicOnHold      = "INSERT INTO asterisk.musiconhold SELECT * FROM asterisk_backup.musiconhold"
	queryQueueLog         = "INSERT IGNORE INTO asterisk.queue_log(time,callid,queuename,agent,event,position,origposition,waittime,channel,logintime,holdtime,calltime,callerid,url,`key`,penalty,ringtime,exten,context)  SELECT time,callid,queuename,agent,event,position,origposition,waittime,channel,logintime,holdtime,calltime,callerid,url,`key`,penalty,ringtime,exten,context FROM asterisk_backup.queue_log"
	queryQueueMemberTable = "INSERT INTO asterisk.queue_member_table(uniqueid,membername,queue_name,interface,penalty,paused) SELECT uniqueid,membername,queue_name,interface,penalty,paused FROM asterisk_backup.queue_member_table"
	queryQueuePenalty     = "INSERT INTO asterisk.queue_penalty SELECT * FROM asterisk_backup.queue_penalty"
	queryQueueTable       = "INSERT INTO asterisk.queue_table(name,musiconhold,announce,context,timeout,monitor_join,monitor_format,queue_youarenext,queue_thereare,queue_callswaiting,queue_holdtime,queue_minutes,queue_seconds,queue_lessthan,queue_thankyou,queue_reporthold,announce_frequency,announce_round_seconds,announce_holdtime,retry,wrapuptime,maxlen,servicelevel,strategy,joinempty,leavewhenempty,eventmemberstatus,eventwhencalled,reportholdtime,memberdelay,weight,timeoutrestart,periodic_announce,periodic_announce_frequency,ringinuse,description,`escalate-if-not-available`,wait_timeout,exit_exten) SELECT name,musiconhold,announce,context,timeout,monitor_join,monitor_format,queue_youarenext,queue_thereare,queue_callswaiting,queue_holdtime,queue_minutes,queue_seconds,queue_lessthan,queue_thankyou,queue_reporthold,announce_frequency,announce_round_seconds,announce_holdtime,retry,wrapuptime,maxlen,servicelevel,strategy,joinempty,leavewhenempty,eventmemberstatus,eventwhencalled,reportholdtime,memberdelay,weight,timeoutrestart,periodic_announce,periodic_announce_frequency,ringinuse,description,`escalate-if-not-available`,wait_timeout,exit_exten FROM asterisk_backup.queue_table;"
	queryRecLevels        = "INSERT INTO asterisk.rec_levels SELECT * FROM asterisk_backup.rec_levels"
	querySipfriends       = "INSERT INTO asterisk.sipfriends(name,type,defaultuser,fromuser,fromdomain,secret,md5secret,auth,mailbox,subscribemwi,vmexten,callerid,cid_number,callingpres,usereqphone,language,`call-limit`,context,dtmfmode,subscribecontext,amaflags,accountcode,musicclass,mohsuggest,allowtransfer,nat,callgroup,pickupgroup,autoframing,disallow,allow,maxcallbitrate,host,outboundproxy,ipaddr,defaultip,port,fullcontact,insecure,qualify,regseconds,regexten,regserver,rtptimeout,rtpholdtimeout,rtpkeepalive,lastms,setvar,useragent,t38pt_udptl,recs,afk,mobile) SELECT name,type,defaultuser,fromuser,fromdomain,secret,md5secret,auth,mailbox,subscribemwi,vmexten,callerid,cid_number,callingpres,usereqphone,language,`call-limit`,context,dtmfmode,subscribecontext,amaflags,accountcode,musicclass,mohsuggest,allowtransfer,nat,callgroup,pickupgroup,autoframing,disallow,allow,maxcallbitrate,host,outboundproxy,ipaddr,defaultip,port,fullcontact,insecure,qualify,regseconds,regexten,regserver,rtptimeout,rtpholdtimeout,rtpkeepalive,lastms,setvar,useragent,t38pt_udptl,recs,afk,mobile FROM asterisk_backup.sipfriends"
	querySounds           = "INSERT INTO asterisk.sounds SELECT * FROM asterisk_backup.sounds"
	querySubscribers      = "INSERT INTO asterisk.subscribers(name,callerid,protocol,context,mobile,cfwd,recs,mailbox) SELECT name,callerid,protocol,context,mobile,cfwd,recs,mailbox FROM asterisk_backup.subscribers"
	queryVoicemailUsers   = "INSERT INTO asterisk.voicemail_users SELECT * FROM asterisk_backup.voicemail_users"
)

type Storage struct {
	conn   *sql.DB
	tables map[string]string
}

func New(conn *sql.DB) *Storage {
	return &Storage{
		conn: conn,
		tables: map[string]string{
			"cdr":                queryCDR,
			"extensions_table":   queryExtensionsTable,
			"ivr_body":           queryIVRBody,
			"ivr_events":         queryIVREvents,
			"ivr_holidays":       queryIVRHolidays,
			"musiconhold":        queryMusicOnHold,
			"queue_log":          queryQueueLog,
			"queue_member_table": queryQueueMemberTable,
			"queue_penalty":      queryQueuePenalty,
			"queue_table":        queryQueueTable,
			"rec_levels":         queryRecLevels,
			"sipfriends":         querySipfriends,
			"sounds":             querySounds,
			"subscribers":        querySubscribers,
			"voicemail_users":    queryVoicemailUsers,
		},
	}
}

func (s *Storage) Migrate() error {

	for table, query := range s.tables {
		res, err := s.conn.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to migrate %s table: %w", table, err)
		}

		rows, err := res.RowsAffected()
		if err != nil {
			logger.Warnf("Unable to fetch count of affected rows for table: %s", table)
		}

		logger.Infof("Succesfully migrate %s table. Rows affected: %d ", table, rows)
	}

	return nil
}
