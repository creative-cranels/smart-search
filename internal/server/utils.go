package server

type TargetDBType string

const (
	TargetDBTypePostgreSQL    TargetDBType = "postgresql"
	TargetDBTypeMongoDB       TargetDBType = "mongodb"
	TargetDBTypeElasticsearch TargetDBType = "elasticsearch"
	//TODO: add one vertical database
)

type SearchProtocolType string

const (
	SearchProtocolTypeREST SearchProtocolType = "rest"
	//TODO: add graphql, grpc
)

type ManagerProtocolType string

const (
	ManagerProtocolTypeREST ManagerProtocolType = "rest"
	//TODO: add graphql, kafka, rabbitmq, grpc
)

type ServerConfig struct {
	TargetDB        TargetDBType
	SearchProtocol  SearchProtocolType
	ManagerProtocol ManagerProtocolType
}
