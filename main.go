package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	c "github.com/sebmaz93/rss_gogator/internal/cmds"
	"github.com/sebmaz93/rss_gogator/internal/config"
	"github.com/sebmaz93/rss_gogator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading the config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	dbQueries := database.New(db)

	appState := &c.State{
		Cfg: &cfg,
		DB:  dbQueries,
	}

	commands := c.Commands{
		RegisteredCommands: make(map[string]func(*c.State, c.Command) error),
	}
	commands.Register("login", c.CmdLogin)
	commands.Register("register", c.CmdRegister)
	commands.Register("reset", c.CmdReset)
	commands.Register("users", c.CmdListUsers)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = commands.Run(appState, c.Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
