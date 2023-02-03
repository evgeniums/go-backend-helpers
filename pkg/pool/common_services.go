package pool

import "github.com/evgeniums/go-backend-helpers/pkg/common"

const (
	TypePostgresDatabase string = "postgres"
	TypeRestApiServer    string = "rest_api"
	TypeRedisServer      string = "redis"
	TypeSqliteDatabase   string = "sqlite"
)

type PostgresServer struct {
	common.ObjectBase
	common.WithNameBase
	common.WithDescriptionBase
	common.WithActiveBase
	HOST         string `gorm:"index" json:"host"`
	PORT         uint16 `gorm:"index" json:"port"`
	DSN          string `gorm:"index" json:"dsn"`
	EXTRA_CONFIG string `json:"extra_config"`
}

type RestApiServer struct {
	common.ObjectBase
	common.WithNameBase
	common.WithDescriptionBase
	common.WithActiveBase
	URL         string `gorm:"index" json:"url"`
	CONTROL_URL string `gorm:"index" json:"control_url"`
}

type RedisServer struct {
	common.ObjectBase
	common.WithNameBase
	common.WithDescriptionBase
	common.WithActiveBase
	HOST     string `gorm:"index" json:"host"`
	PORT     uint16 `gorm:"index" json:"port"`
	DATABASE int    `gorm:"index" json:"database"`
}

type SqliteDatabase struct {
	common.ObjectBase
	common.WithNameBase
	common.WithDescriptionBase
	common.WithActiveBase
	FILE string `gorm:"index" json:"file"`
}
