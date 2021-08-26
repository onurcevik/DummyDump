package internal

import (
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/util"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	pgDatabase           = "--dbname=postgres"
	pgFlagDatabase       = "-d"
	pgFlagFileName       = "-f"
	pgFlagCreateDatabase = "-C"
	pgFlagCreate         = "--create"
	pgFlagFormat         = "--format=c"
	pgVersion            = "--version"
	pgFlagPassword       = "-W"
	mysqlVersion         = "--version"

	mysqlDatabase       = "deneme"
	mysqlFlagUser       = "--user"
	mysqlFlagPassword   = "--password"
	mysqlFlagExecute    = "-e"
	mysqlFlagResultFile = "--result-file"

	host    = "--host="
	port    = "--port="
	noOwner = "-x"
)

// CreateCheckBinaryCommand ...
func CreateCheckBinaryCommand(sourceType config.SourceType) *exec.Cmd {
	return exec.Command(util.Which(), getCheckCommand(sourceType)...)
}

// CreateCheckBinaryPathCommand ...
func CreateCheckBinaryPathCommand(cfg config.Config) *exec.Cmd {
	return exec.Command(util.Which(), getCheckBinaryPathCommand(cfg)...)
}

// CreateImportBinaryCommand ...
func CreateImportBinaryCommand(sourceType config.SourceType) *exec.Cmd {
	return exec.Command(util.Which(), getImportCommand(sourceType)...)
}

// CreateVersionCommand ...
func CreateVersionCommand(binaryPath string, sourceType config.SourceType) *exec.Cmd {
	return exec.Command(binaryPath, getVersionCommandArg(sourceType)...)
}

// CreateExportBinaryCommand ...
func CreateExportBinaryCommand(sourceType config.SourceType) *exec.Cmd {
	return exec.Command(util.Which(), getExportCommand(sourceType)...)
}

// CreateExportCommand ...
func CreateExportCommand(cfg config.Config) *exec.Cmd {
	arg := getExportCommandArg(cfg)
	return exec.Command(cfg.BinaryPath, arg...)
}

// CreateImportCommand ...
func CreateImportCommand(cfg config.Config) *exec.Cmd {
	return exec.Command(cfg.BinaryPath, getImportCommandArg(cfg)...)
}

func getCheckBinaryPathCommand(cfg config.Config) (command []string) {
	switch runtime.GOOS {
	case "darwin":
		command = []string{cfg.BinaryPath}
	case "linux":
		command = []string{cfg.BinaryPath}
	case "windows":
		path, file := filepath.Split(cfg.BinaryPath)
		command = []string{"/r", path, file}
	}
	return command
}

// getVersionCommandArg ...
func getVersionCommandArg(sourceType config.SourceType) (arg []string) {
	switch sourceType {
	case config.PostgreSQL:
		arg = []string{pgVersion}
	case config.MySQL:
		arg = []string{mysqlVersion}
	}
	return arg
}

// getImportCommandArg ...
func getImportCommandArg(cfg config.Config) (arg []string) {
	host := fmt.Sprintf("%s%s", host, cfg.Host)
	port := fmt.Sprintf("%s%d", port, cfg.Port)
	switch cfg.Source {
	case config.PostgreSQL:
		arg = []string{"-d", fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB), filepath.Join(cfg.BackupFilePath, cfg.BackupName)}
	case config.MySQL:
		user := fmt.Sprintf("%s=%s", mysqlFlagUser, cfg.User)
		password := fmt.Sprintf("%s=%s", mysqlFlagPassword, cfg.Password)
		arg = []string{user, password, host, port, cfg.DB, mysqlFlagExecute, "source " + filepath.Join(cfg.BackupFilePath, cfg.BackupName)}
	}
	return arg
}

// getExportCommandArg ...
func getExportCommandArg(cfg config.Config) (arg []string) {
	host := fmt.Sprintf("%s%s", host, cfg.Host)
	port := fmt.Sprintf("%s%d", port, cfg.Port)
	switch cfg.Source {
	case config.PostgreSQL:
		filename := filepath.Join(cfg.BackupFilePath, cfg.BackupName)
		dns := fmt.Sprintf(`--dbname=postgresql://%s:%s@%s:%d/%s`, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
		arg = []string{dns, pgFlagFormat, noOwner, pgFlagFileName, filename}
	case config.MySQL:
		filename := filepath.Join(cfg.BackupFilePath, cfg.BackupName)
		user := fmt.Sprintf(`%s=%s`, mysqlFlagUser, cfg.User)
		password := fmt.Sprintf(`%s=%s`, mysqlFlagPassword, cfg.Password)
		resultFile := fmt.Sprintf(`%s=%s`, mysqlFlagResultFile, filename)
		arg = []string{user, password, host, port, cfg.DB, resultFile}
	}
	return arg
}

// getCheckCommand ...
func getCheckCommand(sourceType config.SourceType) (command []string) {
	switch sourceType {
	case config.PostgreSQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"psql"}
		case "linux":
			command = []string{"psql"}
		case "windows":
			p := strings.TrimSpace(postgresqlBinaryDirectories()[0])
			command = []string{"/r", p, "psql"}
		}
	case config.MySQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"mysql"}
		case "linux":
			command = []string{"mysql"}
		case "windows":
			p := strings.TrimSpace(mysqlBinaryDirectory()[0])
			command = []string{"/r", p, "mysql"}
		}
	}
	return command
}

// getImportCommand ...
func getImportCommand(sourceType config.SourceType) (command []string) {
	switch sourceType {
	case config.PostgreSQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"pg_restore"}
		case "linux":
			command = []string{"pg_restore"}
		case "windows":
			p := strings.TrimSpace(postgresqlBinaryDirectories()[0])
			command = []string{"/r", p, "pg_restore"}
		}
	case config.MySQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"mysql"}
		case "linux":
			command = []string{"mysql"}
		case "windows":
			p := strings.TrimSpace(mysqlBinaryDirectory()[0])
			command = []string{"/r", p, "mysql"}
		}
	}
	return command
}

// getExportCommand ...
func getExportCommand(sourceType config.SourceType) (command []string) {
	switch sourceType {
	case config.PostgreSQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"pg_dump"}
		case "linux":
			command = []string{"pg_dump"}
		case "windows":
			p := strings.TrimSpace(postgresqlBinaryDirectories()[0])
			command = []string{"/r", p, "pg_dump"}
		}
	case config.MySQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"mysqldump"}
		case "linux":
			command = []string{"mysqldump"}
		case "windows":
			p := strings.TrimSpace(mysqlBinaryDirectory()[0])
			command = []string{"/r", p, "mysqldump"}

		}
	}
	return command
}
