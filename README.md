# Overview

Welcome to the documentation for the dummy-dump GitHub repository. This documentation provides an overview of the project, its purpose, features, and instructions for setting it up and contributing.

# Introduction

dummy-dump is a database backup program written in Go. It provides a simple and efficient solution for creating backups of your databases. This repository aims to simplify the backup process and ensure the safety and integrity of your data.

# Features

The key features of dummy-dump include:

- Support for multiple databases:
    - Currently supports MySQL, PostgreSQL, and SQLite databases.
    - Easily extendable to support other popular databases.
- Customizable backup options:
    - Specify the database connection details, including host, port, username, password, and database name.
    - Choose the backup destination directory.
    - Define backup retention policies, such as the number of backups to keep.
- Automated backup scheduling:
    - Schedule automated backups at regular intervals using cron expressions.
    - Configure the frequency and timing of backups based on your requirements.
    - Ensure regular and consistent backup operations.
- Backup encryption and compression:
    - Encrypt the backup files for enhanced security.
    - Compress the backup files to minimize storage space requirements.
    - Choose encryption and compression algorithms based on your security needs.

# Usage

### Requirements

* Go >= 1.15

## CLI Usage

### Scan
* How to use the program via scan

```
go run ./cmd -scan 
```

### Export

* How to run the program

```
go run ./cmd -source=<postgres/mysql/mssql> -export -user=<User Name> -db=<Database Name> -host=<Host> -port=<Port> -backupFilePath<Backup File Path> -backupName=<Backup Name> 
```
### Import

* How to run the program

```
go run ./cmd -source=<postgres/mysql/mssql> -import -user=<User Name> -db=<Database Name> -host=<Host> -port=<Port> -backupFilePath<Backup File Path> -backupName=<Backup Name> 
```

- *import*  
Flag: -import  
Type: bool  
Default: false  

- *export*  
Flag: -export  
Type: bool  
Default: false 

- *source*  
Flag: -source  
Type: string  
Default: null  
Values: **mysql**, **postgres**, **mssql** 

- *user*  
Flag: -user  
Type: string  
Default: null 

- *path*  
Flag: -path  
Type: string  
Default: null 

- *db*  
Flag: -db  
Type: string  
Default: null

- *host*  
  Flag: -host  
  Type: string  
  Default: null

- *port*  
  Flag: -port  
  Type: int  
  Default: null

- *binaryPath*  
Flag: -binaryPath  
Type: string  
Default: null  

- *backupName*  
  Flag: -backupName  
  Type: string  
  Default: null

- *backupFilePath*  
  Flag: -backupFilePath  
  Type: string  
  Default: null

## API Usage
* How to import database

```go

package main

import (
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	dd, err := dummydump.New(&config.Config{
		Source:         <postgres\\mysql\\mssql>,
		Import:         true,
		Export:         false,
		User:           <user>,
		Password:       <password>,
		DB:             <database name>,
		Host:           <host>,
		Port:           <port>,
		BackupFilePath: <path where to save or retrieve>,
		BackupName:     <backup name>,
		BinaryPath:     <binary path>,
	})

	// Check error
	if err != nil {
		panic(err)
	}

	// Call export import
	dd.Check().Import()
}
```
* How to export database
```go

package main

import (
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	dd, err := dummydump.New(&config.Config{
		Source:         <postgres\\mysql\\mssql>,
		Import:         false,
		Export:         true,
		User:           <user>,
		Password:       <password>,
		DB:             <database name>,
		Host:           <host>,
		Port:           <port>,
		BackupFilePath: <path where to save or retrieve>,
		BackupName:     <backup name>,
		BinaryPath:     <binary path>,
	})

	// Check error
	if err != nil {
		panic(err)
	}

	// Call export method
	dd.Check().Export()
}
```

* How to check is there any error while run
```go

package main

import (
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	dd, err := dummydump.New(&config.Config{
		Source:         <postgres\\mysql\\mssql>,
		Import:         false,
		Export:         true,
		User:           <user>,
		Password:       <password>,
		DB:             <database name>,
		Host:           <host>,
		Port:           <port>,
		BackupFilePath: <path where to save or retrieve>,
		BackupName:     <backup name>,
		BinaryPath:     <binary path>,
	})

	// Check error
	if err != nil {
		panic(err)
	}

	// Call export method with check error 
	if _, err = dd.Check().Export().Run(); err != nil {
		panic(err)
	}
}
```

Contributors names and contact info

ex. [@sadihakan](https://github.com/sadihakan/)    
ex. [@onurcevik](https://github.com/onurcevik/)



## License

This project is licensed under the sadihakan License - see the LICENSE.md file for details


