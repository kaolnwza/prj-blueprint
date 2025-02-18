package api

// require on calling external service
// func (c httpClient) NewTLSClientConfig(ctx context.Context, timeout string, m config.CertsConfig) (http.Client, error) {
// 	exp, err := time.ParseDuration(timeout)
// 	if err != nil {
// 		return http.Client{}, err
// 	}

// 	t := c.ts.Clone()
// 	if m.CertsRequire {
// 		if m.InsecureSkip {
// 			t.TLSClientConfig = &tls.Config{
// 				InsecureSkipVerify: true,
// 			}
// 		} else {
// 			certPool := x509.NewCertPool()
// 			certByte := []byte(m.Certs)
// 			if !certPool.AppendCertsFromPEM(certByte) {
// 				logger.GetByContext(ctx).Warn("Invalid certificate, using default transport settings")
// 				return http.Client{Transport: t, Timeout: exp}, nil
// 			}

// 			t.TLSClientConfig = &tls.Config{
// 				RootCAs: certPool,
// 			}
// 		}
// 	}

// 	return http.Client{Transport: t, Timeout: exp}, nil
// }

// func newTransport(httpConf config.HttpConfig) *http.Transport {
// 	t := http.DefaultTransport.(*http.Transport).Clone()
// 	t.MaxIdleConns = httpConf.MaxIdleConns
// 	t.MaxConnsPerHost = httpConf.MaxConnsPerHost
// 	t.MaxIdleConnsPerHost = httpConf.MaxIdleConnsPerHost
// 	return t
// }

// type prepareClient struct {
// 	cli      http.Client
// 	endpoint config.Endpoint
// }

// type httpgen interface {
// 	GetBaseUrl() string
// 	BuildHttpTransport(t *http.Transport)
// }

// func ToHttpClient[T any](ctx context.Context, httpConf config.HttpConfig, c config.BaseExtApiConf[T]) http.Client {
// 	defaultTimeout := time.Second * 30
// 	exp, err := time.ParseDuration(c.Timeout)
// 	if err != nil {
// 		logger.GetByContext(ctx).Errorf("Failed to set timeout, err = %v", err)
// 		exp = defaultTimeout
// 	}

// 	t := http.DefaultTransport.(*http.Transport).Clone()
// 	t.MaxIdleConns = httpConf.MaxIdleConns
// 	t.MaxConnsPerHost = httpConf.MaxConnsPerHost
// 	t.MaxIdleConnsPerHost = httpConf.MaxIdleConnsPerHost

// 	if !c.CertsRequire {
// 		if c.InsecureSkip {
// 			t.TLSClientConfig = &tls.Config{
// 				InsecureSkipVerify: true,
// 			}
// 		} else {
// 			certPool := x509.NewCertPool()
// 			certByte := []byte(c.Certs)
// 			certPool.AppendCertsFromPEM(certByte)

// 			t.TLSClientConfig = &tls.Config{
// 				RootCAs: certPool,
// 			}
// 		}
// 	}

// 	return http.Client{Transport: t, Timeout: exp}
// }

// func BuildHttpClient[T any](ctx context.Context, httpConf config.HttpConfig, c httpgen, ep config.EndpointConf) prepareClient {
// 	defaultTimeout := time.Second * 30
// 	exp, err := time.ParseDuration(c.Timeout)
// 	if err != nil {
// 		logger.GetByContext(ctx).Errorf("Failed to set timeout, err = %v", err)
// 		exp = defaultTimeout
// 	}

// 	t := http.DefaultTransport.(*http.Transport).Clone()
// 	t.MaxIdleConns = httpConf.MaxIdleConns
// 	t.MaxConnsPerHost = httpConf.MaxConnsPerHost
// 	t.MaxIdleConnsPerHost = httpConf.MaxIdleConnsPerHost

// }

// func (c prepareClient[T]) AddQuery(q ApiUrl) {
// 	c.endpoint
// }
