package sshtunnel

import (
	"fmt"
	"io"
	"net"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	boshssh "github.com/cloudfoundry/bosh-cli/ssh"
)

type SSHTunnel interface {
	Start(chan<- error, chan<- error)
}

type sshTunnel struct {
	client boshssh.Client

	localForwardPort  int
	remoteForwardPort int

	remoteListener net.Listener

	logTag string
	logger boshlog.Logger
}

func (s *sshTunnel) Start(readyErrCh chan<- error, errCh chan<- error) {
	err := s.client.Start()
	if err != nil {
		readyErrCh <- bosherr.WrapError(err, "Starting SSH tunnel")
		return
	}

	remoteListenAddr := fmt.Sprintf("127.0.0.1:%d", s.remoteForwardPort)
	s.logger.Debug(s.logTag, "Listening on remote server %s", remoteListenAddr)
	s.remoteListener, err = s.client.Listen("tcp", remoteListenAddr)
	if err != nil {
		readyErrCh <- bosherr.WrapError(err, "Listening on remote server")
		return
	}

	readyErrCh <- nil

	for {
		remoteConn, err := s.remoteListener.Accept()
		s.logger.Debug(s.logTag, "Received connection")
		if err != nil {
			s.logger.Warn(s.logTag, "Failed to accept connection: %s", err.Error())
			errCh <- bosherr.WrapError(err, "Accepting connection on remote server")
		}

		defer func() {
			if err = remoteConn.Close(); err != nil {
				s.logger.Warn(s.logTag, "Failed to close remote listener connection: %s", err.Error())
			}
		}()

		localDialAddr := fmt.Sprintf("127.0.0.1:%d", s.localForwardPort)
		s.logger.Debug(s.logTag, "Dialing local server: %s", localDialAddr)
		localConn, err := net.Dial("tcp", localDialAddr)
		if err != nil {
			s.logger.Warn(s.logTag, "Failed to dail local server, will return. Error: %s", err.Error())
			errCh <- bosherr.WrapError(err, "Dialing local server")
			return
		}

		go func() {
			bytesNum, err := io.Copy(remoteConn, localConn)

			defer func() {
				if err = localConn.Close(); err != nil {
					s.logger.Warn(s.logTag, "Failed to close local dial connection: %s", err.Error())
				}
			}()

			s.logger.Debug(s.logTag, "Copying bytes from local to remote %d", bytesNum)

			if err != nil {
				s.logger.Warn(s.logTag, "Failed to copy bytes from local to remote. Error: %s", err.Error())
				errCh <- bosherr.WrapError(err, "Copying bytes from local to remote")
			}
		}()

		go func() {
			bytesNum, err := io.Copy(localConn, remoteConn)

			defer func() {
				if err = localConn.Close(); err != nil {
					s.logger.Warn(s.logTag, "Failed to close local dial connection: %s", err.Error())
				}
			}()

			s.logger.Debug(s.logTag, "Copying bytes from remote to local %d", bytesNum)

			if err != nil {
				s.logger.Warn(s.logTag, "Failed to copy bytes from remote to local. Error: %s", err.Error())
				errCh <- bosherr.WrapError(err, "Copying bytes from remote to local")
			}
		}()
	}
}
