package ddl

import (
	"ti-pratice/labrpc"
	"time"
)

type Server struct {
	//服务器ID
	id int
	//配置服务器加载地址
	configServer *labrpc.ClientEnd
	//服务器状态
	state int
	//版本号
	version int
	lease   int
	//配置加载的时间间隔
	loadInterval int
	//配置最后更新的时间
	lastTime int64
}

type Config struct {
	//服务器状态
	state int
	lease int
	//配置加载的时间间隔
	loadInterval int
}

type LoadConfigArgs struct {
}

type LoadConfigReply struct {
}

func Make(serverId int, configServer *labrpc.ClientEnd, config *Config) *Server {
	now := time.Now().UnixNano() / 1e6
	server := Server{
		id:           serverId,
		configServer: configServer,
		state:        config.state,
		lease:        config.lease,
		loadInterval: config.loadInterval,
		lastTime:     now,
	}
	return &server
}

func (s *Server) sendLoadConfig(args *LoadConfigArgs, reply *LoadConfigReply) bool {
	ok := s.configServer.Call("Config.LoadConfig", args, reply)
	return ok
}

func (s *Server) loadConfig() {
}
