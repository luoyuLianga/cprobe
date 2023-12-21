package probe

import (
	"github.com/cprobe/cprobe/types"

	_ "github.com/cprobe/cprobe/plugins/blackbox"
	_ "github.com/cprobe/cprobe/plugins/elasticsearch"
	_ "github.com/cprobe/cprobe/plugins/json"
	_ "github.com/cprobe/cprobe/plugins/kafka"
	_ "github.com/cprobe/cprobe/plugins/mysql"
	_ "github.com/cprobe/cprobe/plugins/oracledb"
	_ "github.com/cprobe/cprobe/plugins/prometheus"
	_ "github.com/cprobe/cprobe/plugins/redis"
	_ "github.com/cprobe/cprobe/plugins/whois"
)

func makeJobs() map[string]map[JobID]*JobGoroutine {
	return map[string]map[JobID]*JobGoroutine{
		types.PluginMySQL:         make(map[JobID]*JobGoroutine),
		types.PluginRedis:         make(map[JobID]*JobGoroutine),
		types.PluginMongoDB:       make(map[JobID]*JobGoroutine),
		types.PluginPostgreSQL:    make(map[JobID]*JobGoroutine),
		types.PluginElasticSearch: make(map[JobID]*JobGoroutine),
		types.PluginKafka:         make(map[JobID]*JobGoroutine),
		types.PluginBlackbox:      make(map[JobID]*JobGoroutine),
		types.PluginJson:          make(map[JobID]*JobGoroutine),
		types.PluginPrometheus:    make(map[JobID]*JobGoroutine),
		types.PluginOracleDB:      make(map[JobID]*JobGoroutine),
		types.PluginWhois:         make(map[JobID]*JobGoroutine),
	}
}
