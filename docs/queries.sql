
INSERT INTO asterisk.cdr(calldate,clid,src,dst,dcontext,channel,dstchannel,lastapp,lastdata,duration,billsec,disposition,amaflags,accountcode,userfield,uniqueid) SELECT calldate,clid,src,dst,dcontext,channel,dstchannel,lastapp,lastdata,duration,billsec,disposition,amaflags,accountcode,userfield,uniqueid FROM cdr;

INSERT INTO asterisk.device_devices(id,mac,model_id) SELECT id,mac,model_id FROM device_devices;

INSERT INTO asterisk.device_lines(number,line_no,device_id) SELECT number,line_no,device_id FROM device_lines;

INSERT INTO asterisk.ivr_body(id,major,inc_number,check_time,active_time,active,ivr_comment,wait_exten,audio_id,holiday_id,setcid,disa) SELECT id,major,inc_number,check_time,active_time,active,ivr_comment,wait_exten,audio_id,holiday_id,setcid,disa FROM ivr_body;

INSERT INTO asterisk.ivr_events(id,ivr_id,`key`,action,parameters) SELECT id,ivr_id,`key`,action,parameters FROM ivr_events;

INSERT IGNORE INTO asterisk.queue_log(time,callid,queuename,agent,event,position,origposition,waittime,channel,logintime,holdtime,calltime,callerid,url,`key`,penalty,ringtime,exten,context) SELECT time,callid,queuename,agent,event,position,origposition,waittime,channel,logintime,holdtime,calltime,callerid,url,`key`,penalty,ringtime,exten,context FROM queue_log;

INSERT INTO asterisk.queue_member_table(uniqueid,membername,queue_name,interface,penalty,paused) SELECT uniqueid,membername,queue_name,interface,penalty,paused FROM queue_member_table

INSERT INTO asterisk.queue_table(name,musiconhold,announce,context,timeout,monitor_join,monitor_format,queue_youarenext,queue_thereare,queue_callswaiting,queue_holdtime,queue_minutes,queue_seconds,queue_lessthan,queue_thankyou,queue_reporthold,announce_frequency,announce_round_seconds,announce_holdtime,retry,wrapuptime,maxlen,servicelevel,strategy,joinempty,leavewhenempty,eventmemberstatus,eventwhencalled,reportholdtime,memberdelay,weight,timeoutrestart,periodic_announce,periodic_announce_frequency,ringinuse,description,`escalate-if-not-available`,wait_timeout,exit_exten) SELECT name,musiconhold,announce,context,timeout,monitor_join,monitor_format,queue_youarenext,queue_thereare,queue_callswaiting,queue_holdtime,queue_minutes,queue_seconds,queue_lessthan,queue_thankyou,queue_reporthold,announce_frequency,announce_round_seconds,announce_holdtime,retry,wrapuptime,maxlen,servicelevel,strategy,joinempty,leavewhenempty,eventmemberstatus,eventwhencalled,reportholdtime,memberdelay,weight,timeoutrestart,periodic_announce,periodic_announce_frequency,ringinuse,description,`escalate-if-not-available`,wait_timeout,exit_exten FROM queue_table;

INSERT INTO asterisk.sipfriends(name,type,defaultuser,fromuser,fromdomain,secret,md5secret,auth,mailbox,subscribemwi,vmexten,callerid,cid_number,callingpres,usereqphone,language,`call-limit`,context,dtmfmode,subscribecontext,amaflags,accountcode,musicclass,mohsuggest,allowtransfer,nat,callgroup,pickupgroup,autoframing,disallow,allow,maxcallbitrate,host,outboundproxy,ipaddr,defaultip,port,fullcontact,insecure,qualify,regseconds,regexten,regserver,rtptimeout,rtpholdtimeout,rtpkeepalive,lastms,setvar,useragent,t38pt_udptl,recs,afk,mobile) SELECT name,type,defaultuser,fromuser,fromdomain,secret,md5secret,auth,mailbox,subscribemwi,vmexten,callerid,cid_number,callingpres,usereqphone,language,`call-limit`,context,dtmfmode,subscribecontext,amaflags,accountcode,musicclass,mohsuggest,allowtransfer,nat,callgroup,pickupgroup,autoframing,disallow,allow,maxcallbitrate,host,outboundproxy,ipaddr,defaultip,port,fullcontact,insecure,qualify,regseconds,regexten,regserver,rtptimeout,rtpholdtimeout,rtpkeepalive,lastms,setvar,useragent,t38pt_udptl,recs,afk,mobile FROM sipfriends;

INSERT INTO asterisk.sounds(id,sound_comment,sound_file,sound_file_length) SELECT id,sound_comment,sound_file,sound_file_length FROM sounds;

INSERT INTO asterisk.subscribers(name,callerid,protocol,context,mobile,cfwd,recs,mailbox) SELECT name,callerid,protocol,context,mobile,cfwd,recs,mailbox FROM subscribers;

INSERT INTO asterisk.voicemail_users(uniqueid,customer_id,context,mailbox,password,attach,email,fullname,pager,tz,saycid,dialout,callback,review,operator,envelope,sayduration,saydurationm,sendvoicemail,`delete`,nextaftercmd,forcename,forcegreetings,hidefromdir,stamp) SELECT uniqueid,customer_id,context,mailbox,password,attach,email,fullname,pager,tz,saycid,dialout,callback,review,operator,envelope,sayduration,saydurationm,sendvoicemail,`delete`,nextaftercmd,forcename,forcegreetings,hidefromdir,stamp FROM voicemail_users;
