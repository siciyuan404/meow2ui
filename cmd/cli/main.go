package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/example/a2ui-go-agent-platform/internal/infra/bootstrap"
	"github.com/example/a2ui-go-agent-platform/internal/infra/db"
	"github.com/example/a2ui-go-agent-platform/pkg/agent"
)

func main() {
	ctx := context.Background()

	args := os.Args
	if len(args) < 2 {
		usage()
		return
	}

	switch args[1] {
	case "db:create":
		cfg, err := db.LoadConfigFromEnv()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := db.CreateDatabaseIfNotExists(ctx, cfg); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("database ready: %s\n", cfg.Database)
	case "db:migrate":
		cfg, err := db.LoadConfigFromEnv()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := db.MigrateUp(ctx, cfg, "migrations"); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("migration complete")
	case "db:backup":
		fmt.Println("backup started")
	case "db:restore":
		if len(args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: db:restore <backup-id>")
			os.Exit(1)
		}
		fmt.Printf("restore requested for backup=%s\n", args[2])
	case "data:export":
		if len(args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: data:export <workspace-id>")
			os.Exit(1)
		}
		fmt.Printf("export requested workspace=%s\n", args[2])
	case "data:import":
		if len(args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: data:import <bundle-path>")
			os.Exit(1)
		}
		fmt.Printf("import requested bundle=%s\n", args[2])
	case "workspace:create", "session:create", "agent:run":
		app, err := bootstrap.New(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bootstrap failed: %v\n", err)
			os.Exit(1)
		}
		runAppCommand(ctx, app, args)
	default:
		usage()
	}
}

func runAppCommand(ctx context.Context, app *bootstrap.App, args []string) {
	switch args[1] {
	case "workspace:create":
		if len(args) < 4 {
			fmt.Fprintln(os.Stderr, "usage: workspace:create <name> <root>")
			os.Exit(1)
		}
		ws, err := app.Workspace.Create(ctx, args[2], args[3])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("workspace=%s\n", ws.ID)
	case "session:create":
		if len(args) < 4 {
			fmt.Fprintln(os.Stderr, "usage: session:create <workspace-id> <title>")
			os.Exit(1)
		}
		ssn, ver, err := app.Session.Create(ctx, args[2], args[3], "")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("session=%s version=%s\n", ssn.ID, ver.ID)
	case "agent:run":
		if len(args) < 4 {
			fmt.Fprintln(os.Stderr, "usage: agent:run <session-id> <prompt>")
			os.Exit(1)
		}
		prompt := strings.Join(args[3:], " ")
		out, err := app.Agent.Run(ctx, agent.RunInput{SessionID: args[2], Prompt: prompt})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("run=%s version=%s repaired=%v\n", out.RunID, out.VersionID, out.Repaired)
	default:
		usage()
	}
}

func usage() {
	fmt.Println("a2ui cli")
	fmt.Println("  db:create")
	fmt.Println("  db:migrate")
	fmt.Println("  db:backup")
	fmt.Println("  db:restore <backup-id>")
	fmt.Println("  data:export <workspace-id>")
	fmt.Println("  data:import <bundle-path>")
	fmt.Println("  workspace:create <name> <root>")
	fmt.Println("  session:create <workspace-id> <title>")
	fmt.Println("  agent:run <session-id> <prompt>")
}
