package site

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_site"
)

// GetRemoteSiteRemainingScore reads the latest remaining_score from the site's
// remote database (sys_sites.remaining_score) and returns it. Used to sync
// remote → local.
func GetRemoteSiteRemainingScore(ctx context.Context, site *manage_site.ManageSite) (float64, error) {
	if site.DbHost == "" || site.DbName == "" {
		return 0, fmt.Errorf("site db not configured (db_host or db_name empty)")
	}
	dsn := buildSiteDSN(site)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return 0, fmt.Errorf("open site db: %w", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		return 0, fmt.Errorf("ping site db: %w", err)
	}

	var remoteScore float64
	err = db.QueryRowContext(ctx, "SELECT remaining_score FROM sys_sites LIMIT 1").Scan(&remoteScore)
	if err != nil {
		return 0, fmt.Errorf("read sys_sites.remaining_score: %w", err)
	}
	return remoteScore, nil
}

// UpdateRemoteSiteRemainingScore connects to the site's database and updates
// sys_sites.remaining_score by delta (positive = add, negative = subtract).
func UpdateRemoteSiteRemainingScore(ctx context.Context, site *manage_site.ManageSite, delta float64) error {
	if site.DbHost == "" || site.DbName == "" {
		return fmt.Errorf("site db not configured (db_host or db_name empty)")
	}
	dsn := buildSiteDSN(site)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("open site db: %w", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("ping site db: %w", err)
	}

	_, err = db.ExecContext(ctx, "UPDATE sys_sites SET remaining_score = remaining_score + ?", delta)
	if err != nil {
		return fmt.Errorf("update sys_sites.remaining_score: %w", err)
	}
	return nil
}

func buildSiteDSN(site *manage_site.ManageSite) string {
	cfg := &mysql.Config{
		User:                 site.DbUsername,
		Passwd:               site.DbPassword,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", site.DbHost, site.DbPort),
		DBName:               site.DbName,
		ParseTime:            true,
		Loc:                  time.Local,
		Params:               map[string]string{"charset": "utf8mb4"},
		AllowNativePasswords: true,
	}
	return cfg.FormatDSN()
}
