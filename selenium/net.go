package selenium

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"syscall"
	"time"
)

func isSysErrEACCES(port int) bool {
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		var ope = &net.OpError{}
		if errors.As(err, &ope) {
			var syse = &os.SyscallError{}
			if errors.As(ope.Err, &syse) {
				if syse.Err == syscall.EACCES {
					return true
				} else {
					panic(err)
				}
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	l.Close()
	return false
}

func RanPort() int {
	rand.Seed(time.Now().Unix())
	random := rand.Intn(65535+1024) - 1024
	if random < 0 {
		random = 1024
	}
	return random
}

func FindFreePort() (int, error) {
	port := RanPort()
	if port+5 >= 65535 {
		port = port - 5
	}
	for i := 0; i < 5; i++ {
		port = port + i
		if isSysErrEACCES(port) {
			continue
		} else {
			return port, nil
		}
	}
	return -1, errors.New("unable to find a free port")
}
