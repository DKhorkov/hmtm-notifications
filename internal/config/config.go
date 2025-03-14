package config

import (
	"fmt"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/DKhorkov/libs/db"
	"github.com/DKhorkov/libs/loadenv"
	"github.com/DKhorkov/libs/logging"
	"github.com/DKhorkov/libs/tracing"
)

func New() Config {
	return Config{
		Environment: loadenv.GetEnv("ENVIRONMENT", "local"),
		Version:     loadenv.GetEnv("VERSION", "latest"),
		HTTP: HTTPConfig{
			Host: loadenv.GetEnv("HOST", "0.0.0.0"),
			Port: loadenv.GetEnvAsInt("PORT", 8040),
		},
		Database: db.Config{
			Host:         loadenv.GetEnv("POSTGRES_HOST", "0.0.0.0"),
			Port:         loadenv.GetEnvAsInt("POSTGRES_PORT", 5432),
			User:         loadenv.GetEnv("POSTGRES_USER", "postgres"),
			Password:     loadenv.GetEnv("POSTGRES_PASSWORD", "postgres"),
			DatabaseName: loadenv.GetEnv("POSTGRES_DB", "postgres"),
			SSLMode:      loadenv.GetEnv("POSTGRES_SSL_MODE", "disable"),
			Driver:       loadenv.GetEnv("POSTGRES_DRIVER", "postgres"),
			Pool: db.PoolConfig{
				MaxIdleConnections: loadenv.GetEnvAsInt("MAX_IDLE_CONNECTIONS", 1),
				MaxOpenConnections: loadenv.GetEnvAsInt("MAX_OPEN_CONNECTIONS", 1),
				MaxConnectionLifetime: time.Second * time.Duration(
					loadenv.GetEnvAsInt("MAX_CONNECTION_LIFETIME", 20),
				),
				MaxConnectionIdleTime: time.Second * time.Duration(
					loadenv.GetEnvAsInt("MAX_CONNECTION_IDLE_TIME", 10),
				),
			},
		},
		Logging: logging.Config{
			Level:       logging.Levels.DEBUG,
			LogFilePath: fmt.Sprintf("logs/%s.log", time.Now().UTC().Format("02-01-2006")),
		},
		Clients: ClientsConfig{
			SSO: ClientConfig{
				Host:         loadenv.GetEnv("SSO_CLIENT_HOST", "0.0.0.0"),
				Port:         loadenv.GetEnvAsInt("SSO_CLIENT_PORT", 8070),
				RetriesCount: loadenv.GetEnvAsInt("SSO_RETRIES_COUNT", 3),
				RetryTimeout: time.Second * time.Duration(
					loadenv.GetEnvAsInt("SSO_RETRIES_TIMEOUT", 1),
				),
			},
		},
		Tracing: TracingConfig{
			Server: tracing.Config{
				ServiceName:    loadenv.GetEnv("TRACING_SERVICE_NAME", "hmtm-notifications"),
				ServiceVersion: loadenv.GetEnv("VERSION", "latest"),
				JaegerURL: fmt.Sprintf(
					"http://%s:%d/api/traces",
					loadenv.GetEnv("TRACING_JAEGER_HOST", "0.0.0.0"),
					loadenv.GetEnvAsInt("TRACING_API_TRACES_PORT", 14268),
				),
			},
			Spans: SpansConfig{
				Root: tracing.SpanConfig{
					Opts: []trace.SpanStartOption{
						trace.WithAttributes(
							attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
						),
					},
					Events: tracing.SpanEventsConfig{
						Start: tracing.SpanEventConfig{
							Name: "Calling handler",
							Opts: []trace.EventOption{
								trace.WithAttributes(
									attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
								),
							},
						},
						End: tracing.SpanEventConfig{
							Name: "Received response from handler",
							Opts: []trace.EventOption{
								trace.WithAttributes(
									attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
								),
							},
						},
					},
				},
				Repositories: SpanRepositories{
					Emails: tracing.SpanConfig{
						Opts: []trace.SpanStartOption{
							trace.WithAttributes(
								attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
							),
						},
						Events: tracing.SpanEventsConfig{
							Start: tracing.SpanEventConfig{
								Name: "Calling database",
								Opts: []trace.EventOption{
									trace.WithAttributes(
										attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
									),
								},
							},
							End: tracing.SpanEventConfig{
								Name: "Received response from database",
								Opts: []trace.EventOption{
									trace.WithAttributes(
										attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
									),
								},
							},
						},
					},
				},
				Clients: SpanClients{
					SSO: tracing.SpanConfig{
						Opts: []trace.SpanStartOption{
							trace.WithAttributes(
								attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
							),
						},
						Events: tracing.SpanEventsConfig{
							Start: tracing.SpanEventConfig{
								Name: "Calling gRPC SSO client",
								Opts: []trace.EventOption{
									trace.WithAttributes(
										attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
									),
								},
							},
							End: tracing.SpanEventConfig{
								Name: "Received response from gRPC SSO client",
								Opts: []trace.EventOption{
									trace.WithAttributes(
										attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
									),
								},
							},
						},
					},
				},
				Handlers: SpanHandlers{
					VerifyEmail: tracing.SpanConfig{
						Opts: []trace.SpanStartOption{
							trace.WithAttributes(
								attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
							),
						},
						Events: tracing.SpanEventsConfig{
							Start: tracing.SpanEventConfig{
								Name: "Calling verify-email worker handler",
								Opts: []trace.EventOption{
									trace.WithAttributes(
										attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
									),
								},
							},
							End: tracing.SpanEventConfig{
								Name: "Received response from verify-email worker handler",
								Opts: []trace.EventOption{
									trace.WithAttributes(
										attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
									),
								},
							},
						},
					},
					ForgetPassword: tracing.SpanConfig{
						Opts: []trace.SpanStartOption{
							trace.WithAttributes(
								attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
							),
						},
						Events: tracing.SpanEventsConfig{
							Start: tracing.SpanEventConfig{
								Name: "Calling forget-password worker handler",
								Opts: []trace.EventOption{
									trace.WithAttributes(
										attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
									),
								},
							},
							End: tracing.SpanEventConfig{
								Name: "Received response from forget-password worker handler",
								Opts: []trace.EventOption{
									trace.WithAttributes(
										attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
									),
								},
							},
						},
					},
				},
				Senders: SpanSenders{
					Email: tracing.SpanConfig{
						Opts: []trace.SpanStartOption{
							trace.WithAttributes(
								attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
							),
						},
						Events: tracing.SpanEventsConfig{
							Start: tracing.SpanEventConfig{
								Name: "Sending email",
								Opts: []trace.EventOption{
									trace.WithAttributes(
										attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
									),
								},
							},
							End: tracing.SpanEventConfig{
								Name: "Sent email",
								Opts: []trace.EventOption{
									trace.WithAttributes(
										attribute.String("Environment", loadenv.GetEnv("ENVIRONMENT", "local")),
									),
								},
							},
						},
					},
				},
			},
		},
		NATS: NATSConfig{
			MessageChannelBufferSize: loadenv.GetEnvAsInt("NATS_MESSAGE_CHANNEL_BUFFER_SIZE", 1),
			GoroutinesPoolSize:       loadenv.GetEnvAsInt("NATS_GOROUTINES_POOL_SIZE", 1),
			ClientURL: fmt.Sprintf(
				"nats://%s:%d",
				loadenv.GetEnv("NATS_HOST", "0.0.0.0"),
				loadenv.GetEnvAsInt("NATS_CLIENT_PORT", 4222),
			),
			Subjects: NATSSubjects{
				VerifyEmail:    loadenv.GetEnv("NATS_VERIFY_EMAIL_SUBJECT", "verify-email"),
				ForgetPassword: loadenv.GetEnv("NATS_FORGET_PASSWORD_SUBJECT", "forget-password"),
			},
			Workers: NATSWorkers{
				VerifyEmail: NATSWorker{
					Name: loadenv.GetEnv("NATS_VERIFY_EMAIL_WORKER_NAME", "verify-email-worker"),
				},
				ForgetPassword: NATSWorker{
					Name: loadenv.GetEnv("NATS_FORGET_PASSWORD_WORKER_NAME", "forget-password-worker"),
				},
			},
		},
		Email: EmailConfig{
			SMTP: SMTPConfig{
				Host:     loadenv.GetEnv("EMAIL_SMTP_HOST", "smtp.freesmtpservers.com"),
				Port:     loadenv.GetEnvAsInt("EMAIL_SMTP_PORT", 25),
				Login:    loadenv.GetEnv("EMAIL_SMTP_LOGIN", "smtp"),
				Password: loadenv.GetEnv("EMAIL_SMTP_PASSWORD", "smtp"),
			},
			VerifyEmailURL:    loadenv.GetEnv("EMAIL_VERIFY_URL", "http://localhost:8090/sso/verify-email"),
			ForgetPasswordURL: loadenv.GetEnv("FORGET_PASSWORD_URL", "http://localhost:8090/sso/login"),
		},
	}
}

type ClientConfig struct {
	Host         string
	Port         int
	RetryTimeout time.Duration
	RetriesCount int
}

type ClientsConfig struct {
	SSO ClientConfig
}

type HTTPConfig struct {
	Host string
	Port int
}

type TracingConfig struct {
	Server tracing.Config
	Spans  SpansConfig
}

type SpansConfig struct {
	Root         tracing.SpanConfig
	Repositories SpanRepositories
	Clients      SpanClients
	Handlers     SpanHandlers
	Senders      SpanSenders
}

type SpanHandlers struct {
	VerifyEmail    tracing.SpanConfig
	ForgetPassword tracing.SpanConfig
}

type SpanSenders struct {
	Email tracing.SpanConfig
}

type SpanRepositories struct {
	Emails tracing.SpanConfig
}

type SpanClients struct {
	SSO tracing.SpanConfig
}

type NATSConfig struct {
	ClientURL                string
	MessageChannelBufferSize int
	GoroutinesPoolSize       int
	Subjects                 NATSSubjects
	Workers                  NATSWorkers
}

type NATSSubjects struct {
	VerifyEmail    string
	ForgetPassword string
}

type NATSWorkers struct {
	VerifyEmail    NATSWorker
	ForgetPassword NATSWorker
}

type NATSWorker struct {
	Name string
}

type EmailConfig struct {
	SMTP              SMTPConfig
	VerifyEmailURL    string
	ForgetPasswordURL string
}

type SMTPConfig struct {
	Host     string
	Port     int
	Login    string
	Password string
}

type Config struct {
	HTTP        HTTPConfig
	Database    db.Config
	Logging     logging.Config
	Clients     ClientsConfig
	Tracing     TracingConfig
	Environment string
	Version     string
	NATS        NATSConfig
	Email       EmailConfig
}
