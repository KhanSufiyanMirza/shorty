package config

import (
	"fmt"
	"time"
)

type KeyDBConfig struct {
	Host   string
	Passwd string
}
type ProductionConfig struct {
	ProdMode bool
	Passwd   string
}

// RDBConfig represents relational DB config
type RDBConfig struct {
	Host   string
	DBName string
	User   string
	Passwd string
}

type AppConfig struct {
	AppName             string
	AppPort             string
	JWTPubKeyCert       string
	JWTPrivKeyCert      string
	SSLCert             string
	SSLPrivKey          string
	APIQuota            int64
	DomainUrl           string
	HttpRouter          string
	APIQuotaTTL         time.Duration
	RDBConfig           RDBConfig
	KeyDBConfig         KeyDBConfig
	EventConsumerConfig EventConsumerConfig
	EventProducerConfig EventProducerConfig
}

func (a AppConfig) ValidateRDBConfig() error {
	if a.RDBConfig.Host == "" {
		return fmt.Errorf("invalid hostname for RDB")
	}
	if a.RDBConfig.DBName == "" {
		return fmt.Errorf("invalid hostname for RDB")
	}
	return nil
}

func (a AppConfig) ValidateKeyDBConfig() error {
	if a.KeyDBConfig.Host == "" {
		return fmt.Errorf("invalid hostname for Keydb")
	}
	return nil
}

func (a AppConfig) ValidateEventProducerConfig() error {
	if len(a.EventProducerConfig.KafkaProducerConfig.Brokers) <= 0 {
		return fmt.Errorf("invalid brokers for EventProducer")
	}
	return nil
}

func (a AppConfig) ValidateEventConsumerConfig() error {
	if len(a.EventConsumerConfig.KafkaConsumerConfig.Brokers) <= 0 {
		return fmt.Errorf("invalid brokers for EventConsumer")
	}
	return nil
}
