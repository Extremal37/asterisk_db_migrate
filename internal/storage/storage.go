package storage

import (
	"database/sql"
	"fmt"
	"github.com/Extremal37/asterisk_db_migrate/internal/logger"
	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	conn   *sql.DB
	log    *logger.Logger
	tables map[string]string
}

func New(conn *sql.DB, log *logger.Logger) *Storage {
	return &Storage{
		conn: conn,
		log:  log,
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
	s.log.Debug("Starting migrations func")

	for table, query := range s.tables {
		s.log.Debugf("Migrating table: %s", table)
		res, err := s.conn.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to migrate %s table: %w", table, err)
		}

		rows, err := res.RowsAffected()
		if err != nil {
			s.log.Warnf("Unable to fetch count of affected rows for table: %s", table)
		}

		s.log.Infof("Succesfully migrate %s table. Rows affected: %d ", table, rows)
	}

	return nil
}
