package config

type EventConsumerConfig struct {
	KafkaConsumerConfig KafkaConfig
}
type EventProducerConfig struct {
	KafkaProducerConfig KafkaConfig
}
