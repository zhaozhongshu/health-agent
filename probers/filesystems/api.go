package filesystems

import (
	libprober "github.com/Symantec/health-agent/lib/prober"
	"github.com/Symantec/tricorder/go/tricorder"
	"io"
)

type prober struct {
	dir         *tricorder.DirectorySpec
	fileSystems map[string]*fileSystem // map[device]*fileSystem
}

type fileSystem struct {
	dir        *tricorder.DirectorySpec
	mountPoint string
	available  uint64
	device     string
	free       uint64
	options    string
	size       uint64
	writable   bool
	probed     bool
}

func Register(dir *tricorder.DirectorySpec) libprober.Prober {
	return register(dir)
}

func (p *prober) Probe() error {
	return p.probe()
}

func (p *prober) WriteHtml(writer io.Writer) {
	p.writeHtml(writer)
}
