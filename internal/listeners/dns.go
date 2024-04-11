package listeners

import (
	"github.com/miekg/dns"
	"go.uber.org/zap"
)

// DNS is the listener used for C2 communications over the DNS protocol.
type DNS struct {
	logger *zap.Logger
	server dns.Server
}

// NewDNS returns a new instance of the DNS listener.
func NewDNS(logger *zap.Logger) *DNS {
	return &DNS{
		logger: logger,
		server: dns.Server{},
	}
}

// Start runs the DNS listener.
func (s *DNS) Start() error {
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

// Stop ends the DNS listener.
func (s *DNS) Stop() error {
	err := s.server.Shutdown()
	if err != nil {
		return err
	}

	return nil
}
