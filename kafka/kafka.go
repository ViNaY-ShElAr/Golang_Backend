package kafka

import (
	configs "GO_PROJECT/config"
	"GO_PROJECT/logger"
	"os"
	"strings"
	"time"

	// "crypto/tls"
	// "crypto/x509"

	"github.com/segmentio/kafka-go"
	// "github.com/segmentio/kafka-go/sasl/scram"
)

var KafkaReader *kafka.Reader
var KafkaWriter *kafka.Writer

func InitialiseKafka() {
	// saslmech, err := scram.Mechanism(scram.SHA256, os.Getenv("KAFKA_USERNAME"), os.Getenv("KAFKA_PASSWORD"))
	// if err != nil {
	// 	panic(err)
	// }

	consumer := &kafka.Dialer{
		Timeout: 10 * time.Second,
		// SASLMechanism: saslmech,
		// TLS:           tlsConfig(),
	}

	// Initialize Kafka reader configuration
	readerConfig := kafka.ReaderConfig{
		Brokers:     strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
		GroupID:     configs.Config.KAFKA.CONSUMER_GROUP,
		Topic:       configs.Config.KAFKA.CONSUMER_TOPIC,
		MaxWait:     10000 * time.Millisecond,
		StartOffset: kafka.FirstOffset,
		Dialer:      consumer,
	}

	// Create a Kafka reader
	KafkaReader = kafka.NewReader(readerConfig)

	// kafka producer
	producer := &kafka.Dialer{
		Timeout: 10 * time.Second,
		// SASLMechanism: saslmech,
		// TLS:           tlsConfig(),

	}

	// Initialize Kafka writer configuration
	writerConfig := kafka.WriterConfig{
		Brokers:      strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
		Topic:        configs.Config.KAFKA.CONSUMER_TOPIC,
		Dialer:       producer,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    1,
		BatchTimeout: 10 * time.Millisecond,
		Async:        true,
	}

	// Create Kafka writer
	KafkaWriter = kafka.NewWriter(writerConfig)

	logger.Log.Info("Kafka Connected Sucessfully")
}

// func tlsConfig() *tls.Config {

// 	caCertPool := x509.NewCertPool()
// 	caCertPool.AppendCertsFromPEM([]byte(`Bag Attributes
//     friendlyName: rilrootca
//     2.16.840.1.113894.746875.1.1: <Unsupported tag 6>
// subject=CN = RILROOTCA

// issuer=CN = RILROOTCA

// -----BEGIN CERTIFICATE-----
// MIIFAzCCAuugAwIBAgIQbV1rQBPK0qBLNRPQwPe2ATANBgkqhkiG9w0BAQUFADAU
// MRIwEAYDVQQDEwlSSUxST09UQ0EwHhcNMDkxMDE0MjA1NzU3WhcNMjQxMDE0MjEw
// NzU2WjAUMRIwEAYDVQQDEwlSSUxST09UQ0EwggIiMA0GCSqGSIb3DQEBAQUAA4IC
// DwAwggIKAoICAQDBgLp5wP9pEshr90yJNaVPM+5jhCDkW8zHcZl9jFBB97IEEwpc
// Ejb3gKEU++p1iBR8N/RbYlJ6hXVpqIBoppuWbOVBDgnjl65abQRPs6eGcvVS5s56
// V3l9uHVAZo3Ug23SaTwxqliJ4zkqdLZ0UjX51wyjAc6SLODWG8INKmaZCB36nSq7
// g2pXPgYGBuKvDlMpNSd1/Kto8t4CT2ovGA6P+6HmVyeKpexNFz1f0nFhNbBmz4Z5
// 2YLBQ10X7VCRF8XYenVIx8SSmGYE0uPT3RBQmQ11hAVpi1XnkXsXiBlS/zxQIICu
// 5O7uBKH6f9233vCzTKHRRwr3vJoN2npw134eAJDIc1Vc1Z1KYY1LLDKw/8Q1I07s
// a71APjzCZQs+qsiwLeilknZ56GajyfTMcouyNSfl7/weU9DIoA2XPg5k1eDU3N/R
// ZCoTgnHlOVj8yGSZ2MJT1nyCEkgJDU1rYefUNIetGRlsZIToO6L0btmpS807q1Kw
// 7ObIJE6XwwIM9gL8wb4Hm2vM+t2oTYDTI68I9GpHjxunO3xywVIjPRvFt+q5n3+d
// UjV6gzsc6rYQgm5RvAmpTftB8thyzKXmMncYkMWYQWGNmaK4mPiTx+GIwmpUvGNq
// 7iCpwRCd04RbQKXMJtz0ZFvuZ74+7732v4kXnB1MbkYw5VxZv88T9SFCCQIDAQAB
// o1EwTzALBgNVHQ8EBAMCAYYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUNo1s
// aOPHEzOB/A+XIti3zDtm+Z8wEAYJKwYBBAGCNxUBBAMCAQAwDQYJKoZIhvcNAQEF
// BQADggIBAGqBrnJ1KucP+DG2dVSLtXG3z4YBDqeFpnsha3HzbKeOYfI1GtdYocEA
// KxK5TfBzC8yz2JkSp8nddIPiCepARjrKK/a+yEGDHRTGbbdCpKB5CnJrpsdduREB
// EeUyFmgsxOFObLV+o22UV0HomlVqehB2m5yxE+/3ZPKT39745sM6EQR4hOuX5Ses
// Fn1QCzbis6ta7OJO/j9zTEgwN1JV2hciVBfb5RMnAZaqq1serlBx5ND5dVpetAiB
// q1qd+KZoIFv702/2WG4bZreZBd3npf1fgxMSUvLNH4XYYbj3BK1tkK4aTEdH+vpg
// v8N/VS8tG45vzyP/EcENp1ps6YapDcc0W8mpYHJBohiRq3C3GkqpqFjpslQPbBHY
// +uBGE3FgeIxf2uSZXru72c1Gg3ecXw1fQ2VnaKZ8P6yqVIjNDeh5jMY0n3z0lWA3
// DypXM7HMbIF3WDJI0au6Rc7AJ7VsqNJpGz0X9hi91rN1s5b7wzMZ4/kjZecT5ba3
// jLsJ7Z6zKLK5VwjErnLbeoqFxJ6QErxFseoUpxVfogezYEj/kS+SdAwWOsMzbwMT
// z/vsrj1zFOhgFZN7EoZQ4ayM8203te9rRtwhgEPL2u7+K9y46Tlf
// -----END CERTIFICATE-----
// Bag Attributes
//     friendlyName: rilsubcaent
//     2.16.840.1.113894.746875.1.1: <Unsupported tag 6>
// subject=DC = com, DC = ril, DC = in, CN = RILSUBCAENT

// issuer=CN = RILROOTCA

// -----BEGIN CERTIFICATE-----
// MIIGBjCCA+6gAwIBAgIKGkbOuQAAAAAAEzANBgkqhkiG9w0BAQsFADAUMRIwEAYD
// VQQDEwlSSUxST09UQ0EwHhcNMTcwNTEwMTAxNTQwWhcNMjQxMDE0MjEwNzU2WjBU
// MRMwEQYKCZImiZPyLGQBGRYDY29tMRMwEQYKCZImiZPyLGQBGRYDcmlsMRIwEAYK
// CZImiZPyLGQBGRYCaW4xFDASBgNVBAMTC1JJTFNVQkNBRU5UMIICIjANBgkqhkiG
// 9w0BAQEFAAOCAg8AMIICCgKCAgEAt4LI0LV1jw72HXQjEQosydTahjEKIEnwgBd5
// lJyQXqct3XuWCe7saRmx/AF0K9ZHAjhGDD7v1+XlSavNxfC87kcwNRx4m06XjZkm
// AXlLr7ydm27ivS1hkJekSwdo47DIoD8VPpL1OOK5hF8kTh0gddoWQ0mu0RtMyAir
// f6NPbEe5uCWuBBdOrVhmsGMhPeZkDZWI2Oup7C+ITY2UaT5BudtWLRPYFso/AXbe
// RhA+HlOlyGWx2BpBa2o/ryuyW/RVFk7aE/Zvt3qWLcGTzU+1fPODEnjcA4laMmOt
// Akrl58rr8h82dDrgWsiwTNFY5LyrrDqr74WkcAdOXi9h62cRojwPaw7htmn+uTzS
// qYoq9agUGw70wZqdVXkiljbVeqwWrvSRalc1san1Tm/uWlhyVAJDboc3tsFEfSud
// NasZdSDnWHifevW9KE9jVE004IrzdfNf3XRzGhldcR46y33dcgowy7UfoHliV3xJ
// mCFQ8dIDx51NnNtBghDqgdIJargHj/3vTRIEGPXQ/MnkAKv2el7UqsAS+w5qRvbB
// u8H8UOjSZMb1PLMGnex3jVeLgPs69RuGHiSjPo3xnJzi9qsAM4LLWc1bVxUUm9vc
// t09gGNl/6yuvPrWRnZ87RjMCvf7rgPuqiRSF4sJ2EwZ204uMB92unNJ480sGMhx2
// YSc46CMCAwEAAaOCARgwggEUMBAGCSsGAQQBgjcVAQQDAgEAMB0GA1UdDgQWBBT3
// kP7IVBVEsIjD5FQ8bbT9bTORwzAZBgkrBgEEAYI3FAIEDB4KAFMAdQBiAEMAQTAL
// BgNVHQ8EBAMCAYYwDwYDVR0TAQH/BAUwAwEB/zAfBgNVHSMEGDAWgBQ2jWxo48cT
// M4H8D5ci2LfMO2b5nzA4BgNVHR8EMTAvMC2gK6AphidodHRwOi8vY3JsLnJpbC5j
// b20vcm9vdENBL1JJTHJvb3RDQS5jcmwwTQYIKwYBBQUHAQEEQTA/MD0GCCsGAQUF
// BzAChjFodHRwOi8vY3JsLnJpbC5jb20vUm9vdENBL1JJTFJPT1RDQV9SSUxST09U
// Q0EuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQA1p+w180yyAvwrqk5OpI263mDM/ZkR
// IudY0SIOAJfmJQyJRvkbsGxKy/NTJ+UrV+lREOwN8+XpPKJvJJFIEpIdWuT3Mki6
// N9yYgbmjP+bNhUVIt73jAegJ0qkUgedUXsNf4GnS7pFs0R2XGzgztFN5rA3icqDW
// 921SVygntucLpbKI/wQmvAgkfpGKB24gZNxochsle6RVnp8jvt40n7dvk81Iza3+
// HNcElj6xcFofRmu6tpzeF394nUF27d3IiAgZX9vFxsOcCUzsou33rj5V0UOuNvqv
// 8PZ460jqnvXWETTOwfECxevp9nmh2FmFsFAcgl5xK9v7jgY/EU9jAl5H7Cq1TPMX
// u7eSSnVuQuilH3VGt2XxNmgi/GiXZ5aWwQnc1KA9KUJU9xtZOS0lra1Cf82Vov3o
// kf5jWaOy8f9g8jBZD+5ZNJtvKaZZRfZUTtVcgeXU8BG2dcZeeEWsI2DRxhqguMEI
// 1CSWKD0lIIcDrVSET6pdZSoqifJcOy0bVfuTj0zXxFDNddPd3s2EDb5wWDcBt5Ov
// lJKKNza2t+2UIPxvf2kwa6LEWLaS3A9GLD3kSqeq3Q0hAoV4vZvve7wBbabI0wpZ
// H7GMRT9cKF4Tzg==
// -----END CERTIFICATE-----
// `))
// 	return &tls.Config{
// 		RootCAs:            caCertPool,
// 		InsecureSkipVerify: true,
// 	}
// }
