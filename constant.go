package pcommon

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const DAY = 24 * time.Hour
const WEEK = 7 * DAY
const MONTH = 30 * DAY
const QUARTER = 90 * DAY

const MIN_TIME_FRAME = time.Second
const MAX_TIME_FRAME = QUARTER

const FUTURES_KEY = "_futures"
const SPOT_KEY = "_spot"

type env struct {
	ARCHIVES_DIR              string
	DATABASES_DIR             string
	MAX_SIMULTANEOUS_PARSING  int
	PARSER_SERVER_PORT        string
	INDEXER_SERVER_PORT       string
	MAX_SIMULTANEOUS_INDEXING int
}

var Env = env{
	ARCHIVES_DIR:              "archives",
	DATABASES_DIR:             "databases",
	MAX_SIMULTANEOUS_PARSING:  3,
	MAX_SIMULTANEOUS_INDEXING: 3,
	PARSER_SERVER_PORT:        "8889",
	INDEXER_SERVER_PORT:       "8890",
}

func (e env) Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Archives directory
	archiveDir := os.Getenv("ARCHIVES_DIR")
	if archiveDir != "" {
		if stat, err := os.Stat(archiveDir); os.IsNotExist(err) || !stat.IsDir() {
			log.Fatalf("archives directory not found or is not a directory")
		} else {
			Env.ARCHIVES_DIR = archiveDir
		}
	}

	// Databases directory
	dbDir := os.Getenv("DATABASES_DIR")
	if dbDir != "" {
		if stat, err := os.Stat(dbDir); os.IsNotExist(err) || !stat.IsDir() {
			log.Fatalf("databases directory not found or is not a directory")
		} else {
			Env.DATABASES_DIR = dbDir
		}
	}

	// Max simultaneous parsing workerss
	maxSimultaneousParsing := os.Getenv("MAX_SIMULTANEOUS_PARSING")
	if maxSimultaneousParsing != "" {
		max, err := strconv.Atoi(maxSimultaneousParsing)
		if err != nil {
			log.Fatal("Error parsing MAX_SIMULTANEOUS_PARSING")
		} else {
			Env.MAX_SIMULTANEOUS_PARSING = max
		}
	}

	// Max simultaneous indexing workers
	maxSimultaneousIndexing := os.Getenv("MAX_SIMULTANEOUS_INDEXING")
	if maxSimultaneousIndexing != "" {
		max, err := strconv.Atoi(maxSimultaneousIndexing)
		if err != nil {
			log.Fatal("Error parsing MAX_SIMULTANEOUS_INDEXING")
		} else {
			Env.MAX_SIMULTANEOUS_INDEXING = max
		}
	}

	// Parser server port
	parserServerPort := os.Getenv("PARSER_SERVER_PORT")
	if parserServerPort != "" {
		serverPortInt, err := strconv.Atoi(parserServerPort)
		if err != nil {
			log.Fatal("Error parsing PARSER_SERVER_PORT")
		} else {
			if serverPortInt < 0 || serverPortInt > 65535 {
				log.Fatal("Invalid port PARSER_SERVER_PORT")
			}
		}
		Env.PARSER_SERVER_PORT = strconv.Itoa(serverPortInt)
	}

	// Indexer server port
	indexerServerPort := os.Getenv("INDEXER_SERVER_PORT")
	if indexerServerPort != "" {
		serverPortInt, err := strconv.Atoi(indexerServerPort)
		if err != nil {
			log.Fatal("Error parsing INDEXER_SERVER_PORT")
		} else {
			if serverPortInt < 0 || serverPortInt > 65535 {
				log.Fatal("Invalid port INDEXER_SERVER_PORT")
			}
		}
		Env.INDEXER_SERVER_PORT = strconv.Itoa(serverPortInt)
	}
}
