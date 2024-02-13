package config

import "strconv"

type CacheConfig struct {
	Adders   map[string]string
	Password string
	DB       int
}

func NewCacheConfig(hosts []string, pass string, dbNum int) CacheConfig {
	adders := hostMap(hosts)

	return CacheConfig{
		Adders:   adders,
		Password: pass,
		DB:       dbNum,
	}
}

// hostMap to Converts list of hosts to map
func hostMap(hosts []string) map[string]string {
	hostMap := make(map[string]string)
	for i, host := range hosts {
		hostMap["shard"+strconv.Itoa(i)] = host
	}

	return hostMap
}
