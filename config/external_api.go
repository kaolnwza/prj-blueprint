package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/kaolnwza/proj-blueprint/libs/safetyper"
)

type BaseExtApiConf[T any] struct {
	BaseUrl        string      `mapstructure:"base_url"`
	Timeout        string      `mapstructure:"timeout"`
	CertsConf      CertsConfig `mapstructure:"certs"`
	httpPoolClient http.Client
	Endpoints      T
}

type CertsConfig struct {
	Type         string `mapstructure:"type"`
	CertsRequire bool   `mapstructure:"certs_require"`
	InsecureSkip bool   `mapstructure:"insecure_skip"`
	Certs        string `mapstructure:"certs"`
}

type ExternalApiConfig struct {
	UserCenterConf BaseExtApiConf[UserCenterEndpoints] `mapstructure:"user_centers"`
}

type (
	UserCenterEndpoints struct {
		Inq EndpointConf `mapstructure:"inq"`
	}
)

func (b *BaseExtApiConf[T]) setHttpClient(client http.Client) {
	b.httpPoolClient = client
}

func (m BaseExtApiConf[T]) GetHttpClient() http.Client {
	return m.httpPoolClient
}

const defaultTimeout = time.Second * 30
const (
	certsTypeSecret = "secret"
	certsTypeFile   = "file"
)

func (c CertsConfig) toCertPool() (*x509.CertPool, error) {
	certs := []byte(c.Certs)
	if c.Type == certsTypeFile {
		certBytes, err := os.ReadFile(c.Certs)
		if err != nil {
			return nil, fmt.Errorf("failed to read cert file: %w", err)
		}

		certs = certBytes
	}

	certPool := x509.NewCertPool()
	certByte := []byte(certs)
	if !certPool.AppendCertsFromPEM(certByte) {
		return nil, fmt.Errorf("failed to valid certs from AppendCertsFromPEM")
	}

	return certPool, nil
}

func setupHttpClient(fieldName string, c CertsConfig, timeout string, httpConf HttpConfig) http.Client {
	exp, err := time.ParseDuration(timeout)
	if err != nil {
		log.Println("[Config] failed to parse timeout on external api http client -> parsed default timeout, err = %w", err)
		exp = defaultTimeout
	}

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = httpConf.MaxIdleConns
	t.MaxConnsPerHost = httpConf.MaxConnsPerHost
	t.MaxIdleConnsPerHost = httpConf.MaxIdleConnsPerHost

	if c.CertsRequire {
		if c.InsecureSkip {
			t.TLSClientConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		} else {
			certPool, err := c.toCertPool()
			if err != nil {
				log.Printf("[Config][%s] %v\n", fieldName, err)
				return http.Client{Transport: t, Timeout: exp}
			}

			t.TLSClientConfig = &tls.Config{
				RootCAs: certPool,
			}
		}
	}

	return http.Client{Transport: t, Timeout: exp}
}

func (c *ExternalApiConfig) setupHttpPoolClients(httpConf HttpConfig) error {
	configValue := reflect.ValueOf(c).Elem()
	configType := configValue.Type()

	for i := 0; i < configType.NumField(); i++ {
		field := configValue.Field(i)
		if field.Kind() == reflect.Struct {
			certsConfigField := field.FieldByName("CertsConf")
			certsPtr := certsConfigField.Addr().Interface().(*CertsConfig)

			if safetyper.IsNilPtr(c) {
				return fmt.Errorf("certsPtr is null")
			}

			timeoutValue := field.FieldByName("Timeout").String()
			strSpl := strings.Split(field.Type().Name(), ".")
			fieldName := strSpl[len(strSpl)-1]
			client := setupHttpClient(fieldName, *certsPtr, timeoutValue, httpConf)

			// field.FieldByName("httpPoolClient").Set(reflect.ValueOf(client))

			if !field.CanAddr() {
				return fmt.Errorf("field is not addressable")
			}

			baseConfPtr := field.Addr().Interface()
			baseConf, ok := baseConfPtr.(interface{ setHttpClient(http.Client) })
			if !ok {
				return fmt.Errorf("field does not implement setHttpClient")
			}

			baseConf.setHttpClient(client)
		}

	}

	return nil
}
