package db

import (
	"context"
	"errors"
	"fmt"
	"log"

    "github.com/KonnorFrik/ChatServer/pkg/sql/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
    "github.com/ilyakaznacheev/cleanenv"
)

type DbConn struct {
    conn *pgx.Conn
    *models.Queries
}

type DbConfig struct {
    Host string `env:"DB_HOST" env-required`
    User string `env:"DB_USER" env-required`
    Password string `env:"DB_PASSWORD" env-required`
    DbName string `env:"DB_DBNAME" env-required`
    Port string `env:"DB_PORT" env-required env-default:"5432"`
    SSLMode string `env:"DB_SSLMODE" env-default:"disable"`
}

var (
    // ErrInvalidConfig - invalid configuration for connect to db
    ErrInvalidConfig = errors.New("invalid database configuration")
    // ErrDataBaseNotConnected - no connection to db
    ErrDataBaseNotConnected = errors.New("database is not connected")
    // ErrConstraintUniqueViolation - dublication in db
    ErrConstraintUniqueViolation = errors.New("unique constraint violation")
    // ErrConstraintForeignKeyViolation - db error
    ErrConstraintForeignKeyViolation = errors.New("foreign key violation")
    // ErrConstraintNotNullViolation - null value in db
    ErrConstraintNotNullViolation = errors.New("not null violation")
    // ErrConstraintCheckViolation - custom check not passed
    ErrConstraintCheckViolation = errors.New("check violation")
    // ErrUnknown - any other not documented error in db
    ErrUnknown = errors.New("unknown error")

    // Default configuration for connect to postgres container from this repo
	DefaultConfig DbConfig

    dbObj DbConn
)

func init() {
    if err := cleanenv.ReadEnv(&DefaultConfig); err != nil {
        log.Println("[db/LoadEnv]: Error:", err)
        return
    }

    err := dbObj.connect()

    if err != nil {
        log.Println("[db/Connect]: Error:", err)
        return 
    }

    log.Println("[db]: DataBase connected")
}

// DB - return a singletone object for interact with postgres db
func DB() *DbConn {
    return &dbObj
}

// WrapError - wraps an error from the db into one of the predefined ones
func (dc *DbConn) WrapError(err error) error {
    newErr := wrapError(err)

    if newErr == nil {
        return nil
    }

    if errors.Is(newErr, ErrDataBaseNotConnected) {
        dc.connect()
    }

    return newErr
}

func (dc *DbConn) connect() error {
    if !DefaultConfig.isValid() {
        log.Printf("[db]: Invalid config: %+v\n", DefaultConfig)
        return ErrInvalidConfig
    }

    var db *pgx.Conn
    var err error
    connectionInfo := DefaultConfig.String()
    db, err = pgx.Connect(context.Background(), connectionInfo)

    if err != nil {
        wrapped := wrapError(err)
        log.Printf("[db/.connect]: Error on connection: %q\n", wrapped)
        return wrapped
    }

    dc.conn = db
    dc.Queries = models.New(dc.conn)
    log.Println("[db/.connect]: Postgres Connected")
    return nil
}

func (dc *DbConfig) String() string {
    return fmt.Sprintf(
        "host=%s database=%s user=%s password=%s port=%s sslmode=%s",
        dc.Host,
        dc.DbName,
        dc.User,
        dc.Password,
        dc.Port,
        dc.SSLMode,
    )
}

func (dc *DbConfig) isValid() bool {
    return dc.Host != "" && dc.User != "" && dc.Password != "" && dc.Port != ""
}

func wrapError(err error) error {
    if err == nil {
        return nil
    } 

    var pgErr *pgconn.PgError

    if errors.As(err, &pgErr) {
        switch pgErr.Code {
        case "23505":
            return fmt.Errorf("%w: %w", ErrConstraintUniqueViolation, err)
        case "23503":
            return fmt.Errorf("%w: %w", ErrConstraintForeignKeyViolation, err)
        case "23502":
            return fmt.Errorf("%w: %w", ErrConstraintNotNullViolation, err)
        case "23514": 
            return fmt.Errorf("%w: %w", ErrConstraintCheckViolation, err)

        case "57P01":
            return fmt.Errorf("%w: %w", ErrDataBaseNotConnected, err)
        case "57P02":
            return fmt.Errorf("%w: %w", ErrDataBaseNotConnected, err)

        default:
            log.Printf("[db] Can't process code: %q\n", pgErr.Code)
        }
    }

    return fmt.Errorf("%w: %w", ErrUnknown, err)
}
