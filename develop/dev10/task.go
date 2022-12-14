package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type tlnCli struct {
	address    string
	connection net.Conn
	timeout    time.Duration
	in         io.ReadCloser
	out        io.Writer
	done       chan struct{}
}

func NewtlnCli(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) (tlnc *tlnCli) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan struct{})
	go func() {
		<-sigs
		close(done)
	}()
	return &tlnCli{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
		done:    done,
	}
}

func (tlnc *tlnCli) Connect() error {
	conn, err := net.DialTimeout("tcp", tlnc.address, tlnc.timeout)
	if err != nil {
		return err
	}
	tlnc.connection = conn
	return nil
}

func (tlnc *tlnCli) Send() error {
	buffer, err := bufio.NewReader(tlnc.in).ReadBytes(byte('\n'))
	switch {
	case err == io.EOF:
		select {
		case <-tlnc.done:
		default:
			close(tlnc.done)
		}
		return nil
	case err != nil:
		return err
	default:
	}
	_, err = tlnc.connection.Write(buffer)
	return err
}

func (tlnc *tlnCli) Receive() error {
	buffer, err := bufio.NewReader(tlnc.connection).ReadBytes(byte('\n'))
	switch {
	case err == io.EOF:
		select {
		case <-tlnc.done:
		default:
			close(tlnc.done)
		}
		return nil
	case err != nil:
		return err
	default:
	}
	_, err = tlnc.out.Write(buffer)
	return err
}

func (tlnc *tlnCli) Close() error {
	return tlnc.connection.Close()
}

func (tlnc *tlnCli) Done() <-chan struct{} {
	return tlnc.done
}

func main() {
	timeout := flag.String("timeout", "10s", "connection timeout")
	hostPort := flag.Args()[0]

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("%s -timeout 10s hostname:port\n", os.Args[0])
		return
	}

	tmotDur, err := time.ParseDuration(*timeout)
	if err != nil {
		fmt.Printf("%s -timeout 10s hostname:port\n", os.Args[0])
		panic(err.Error())
	}

	tlnetCl := NewtlnCli(hostPort, tmotDur, os.Stdin, os.Stdout)
	err = tlnSndRec(tlnetCl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	<-tlnetCl.Done()
	tlnetCl.Close()
}

func tlnSndRec(tc *tlnCli) error {
	if err := tc.Connect(); err != nil {
		return err
	}
	go func() {
		for {
			select {
			case <-tc.Done():
				return
			default:
			}
			tc.Receive()
		}
	}()

	go func() {
		for {
			select {
			case <-tc.Done():
				return
			default:
			}
			tc.Send()
		}
	}()

	return nil
}
