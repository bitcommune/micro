package types

type ProjectConfig struct {
	Name           string
	Version        string
	GoModule       string
	WithDB         bool
	DBType         string
	WithDocker     bool
	WithTests      bool
	WithMiddleware bool
}

type ProtoConfig struct {
	PackageName string
	GoPackage   string
	ServiceName string
	Messages    []Message
	RPCs        []RPC
}

type ServiceConfig struct {
	Name        string
	PackageName string
	ServiceName string
	RPCs        []RPC
}

type DatabaseConfig struct {
	Type   string
	Models []Model
}

type Message struct {
	Name   string
	Fields []Field
}

type Field struct {
	Type string
	Name string
	Tag  int
}

type RPC struct {
	Name     string
	Request  string
	Response string
}

type Model struct {
	Name   string
	Fields []Field
}

type GormField struct {
	Name    string
	Type    string
	GormTag string
}
