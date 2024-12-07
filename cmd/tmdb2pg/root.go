package tmdb2pg

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/mgoodnight/tmdb2pg/db"
	"github.com/spf13/cobra"
)

var dbConn *pgx.Conn
var rootCmd = &cobra.Command{
	Use: `
tmdb2pg postgres://user:pass@localhost:5432/mydb
TMDB2PG_DSN=postgres://user:pass@localhost:5432/mydb tmdb2pg`,
	Short: "tmdb2pg - fetch latest TMDB daily export and load into a PostreSQL database",
	Long:  `tmdb2pg attempts to fetch the latest daily export file provided from TMDB and loads the data into a PostgreSQL database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running...")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		var dsn string
		if len(args) == 0 {
			envDsn := os.Getenv("TMDB2PG_DSN")
			if envDsn != "" {
				dsn = envDsn
			} else {
				return fmt.Errorf("missing connection string")
			}
		} else {
			dsn = args[0]
		}

		dbConn = db.OpenConn(dsn)

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "whoops. there was an error while executing tmdb2db '%s'\n", err)
		os.Exit(1)
	}

	if dbConn != nil {
		defer dbConn.Close(context.Background())
	}
}
