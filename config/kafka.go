package config

type KafkaConfig struct {
	Brokers    []string
	Version    string
	Topics     []string
	Partitions int
}
