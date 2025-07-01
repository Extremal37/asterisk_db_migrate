package storage

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Extremal37/asterisk_db_migrate/internal/logger"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

const migrateTimeout = 30 * time.Minute

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

func (s *Storage) Migrate(ctx context.Context) error {
	s.log.Debug("Starting migrations func")
	c, cancel := context.WithTimeout(ctx, migrateTimeout)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(s.tables))

	successMigrate := true

	for table, query := range s.tables {
		go func() {
			defer wg.Done()
			success := s.migrateTable(c, table, query)
			if !success {
				successMigrate = false
			}
		}()
	}

	wg.Wait()

	if !successMigrate {
		return errors.New("failed to migrate")
	}
	return nil
}

func (s *Storage) Close() {
	err := s.conn.Close()
	if err != nil {
		s.log.Warnf("Failed to close DB connection: %v", err)
	} else {
		s.log.Debug("DB connection closed")
	}
}

func (s *Storage) migrateTable(ctx context.Context, table string, query string) bool {
	s.log.Debugf("Migrating table: %s", table)

	start := time.Now()

	res, err := s.conn.ExecContext(ctx, query)
	if err != nil {
		s.log.Errorf("failed to migrate %s table: %w", table, err)
		return false
	}

	rows, err := res.RowsAffected()
	if err != nil {
		s.log.Warnf("Unable to fetch count of affected rows for table: %s", table)
	}

	elapsed := time.Since(start).Round(time.Millisecond)
	s.log.Debugf("Succesfully migrate %s table. Rows affected: %d . Time elapsed: %v ", table, rows, elapsed)

	return true
}
