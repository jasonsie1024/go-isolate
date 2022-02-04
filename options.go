package isolate

import "fmt"

type Options struct {
	TimeLimit, WallTimeLimit, ExtraTimeLimit        float64
	MemLimit, StackLimit, FileSizeLimit, QuotaLimit int
	ProcessLimit                                    int
	Metafile, Stdin, Stdout, Stderr                 string
	StderrToStdOut, ShareNet, InheritFds            bool
	CGMemLimit                                      int
	NoCGTiming                                      bool
}

func (o *Options) BuildArguments() []string {
	args := []string{}
	if o.TimeLimit != 0 {
		args = append(args, "-t", fmt.Sprint(o.TimeLimit))
	}
	if o.WallTimeLimit != 0 {
		args = append(args, "-w", fmt.Sprint(o.WallTimeLimit))
	}
	if o.ExtraTimeLimit != 0 {
		args = append(args, "-x", fmt.Sprint(o.ExtraTimeLimit))
	}

	if o.MemLimit != 0 {
		args = append(args, "-m", fmt.Sprint(o.MemLimit))
	}
	if o.StackLimit != 0 {
		args = append(args, "-k", fmt.Sprint(o.StackLimit))
	}
	if o.FileSizeLimit != 0 {
		args = append(args, "-f", fmt.Sprint(o.FileSizeLimit))
	}
	if o.QuotaLimit != 0 {
		args = append(args, "-q", fmt.Sprint(o.QuotaLimit))
	}

	if o.Metafile != "" {
		args = append(args, "-M", o.Metafile)
	}
	if o.Stdin != "" {
		args = append(args, "-i", o.Stdin)
	}
	if o.Stdout != "" {
		args = append(args, "-o", o.Stdout)
	}
	if o.Stderr != "" {
		args = append(args, "-r", o.Stderr)
	}

	if o.StderrToStdOut {
		args = append(args, "--stderr-to-stdout")
	}
	if o.ShareNet {
		args = append(args, "--share-net")
	}
	if o.InheritFds {
		args = append(args, "--inherit_fds")
	}

	if o.CGMemLimit != 0 {
		args = append(args, "--cg-mem", fmt.Sprint(o.CGMemLimit))
	}

	if o.NoCGTiming {
		args = append(args, "--no-cg-timing")
	}

	return args
}

type WithOption func(*Sandbox)

func Stdin(Stdin string) WithOption {
	return func(s *Sandbox) {
		s.Options.Stdin = Stdin
	}
}

func Stdout(Stdout string) WithOption {
	return func(s *Sandbox) {
		s.Options.Stdout = Stdout
	}
}

func Stderr(Stderr string) WithOption {
	return func(s *Sandbox) {
		s.Options.Stderr = Stderr
	}
}

func Metafile(Metafile string) WithOption {
	return func(s *Sandbox) {
		s.Options.Metafile = Metafile
	}
}

func ProcessLimit(ProcessLimit int) WithOption {
	return func(s *Sandbox) {
		s.Options.ProcessLimit = ProcessLimit
	}
}

func FileSizeLimit(FileSizeLimit int) WithOption {
	return func(s *Sandbox) {
		s.Options.FileSizeLimit = FileSizeLimit
	}
}

func QuotaLimit(QuotaLimit int) WithOption {
	return func(s *Sandbox) {
		s.Options.QuotaLimit = QuotaLimit
	}
}

func StackLimit(StackLimit int) WithOption {
	return func(s *Sandbox) {
		s.Options.StackLimit = StackLimit
	}
}

func TimeLimit(Time float64) WithOption {
	return func(s *Sandbox) {
		s.Options.TimeLimit = Time
	}
}

func WallTimeLimit(WallTimeLimit float64) WithOption {
	return func(s *Sandbox) {
		s.Options.WallTimeLimit = WallTimeLimit
	}
}

func ExtraTimeLimit(ExtraTimeLimit float64) WithOption {
	return func(s *Sandbox) {
		s.Options.ExtraTimeLimit = ExtraTimeLimit
	}
}

func MemLimit(Mem int) WithOption {
	return func(s *Sandbox) {
		s.Options.MemLimit = Mem
	}
}

func StderrToStdOut(StderrToStdOut bool) WithOption {
	return func(s *Sandbox) {
		s.Options.StderrToStdOut = StderrToStdOut
	}
}

func ShareNet(ShareNet bool) WithOption {
	return func(s *Sandbox) {
		s.Options.ShareNet = ShareNet
	}
}

func InheritFds(InheritFds bool) WithOption {
	return func(s *Sandbox) {
		s.Options.InheritFds = InheritFds
	}
}

func CGMemLimit(CGMemLimit int) WithOption {
	return func(s *Sandbox) {
		s.Options.CGMemLimit = CGMemLimit
	}
}

func NoCGTiming(NoCGTiming bool) WithOption {
	return func(s *Sandbox) {
		s.Options.NoCGTiming = NoCGTiming
	}
}
