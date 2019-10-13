package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func normalEsc(b64 string) string {
	return "\x1B]52;;" + b64 + "\x1B\x5C"
}

func tmuxEsc(b64 string) string {
	return "\x1BPtmux;\x1B\x1B]52;;" + b64 + "\x1B\x1B\x5C\x5C\x1B\x5C"
}

func screenEsc(b64 string) string {
	out := []string{}
	for i := 0; ; i++ {
		begin, end := i*76, (i+1)*76
		if end > len(b64) {
			end = len(b64)
		}
		if begin == 0 {
			out = append(out, "\x1BP\x1B]52;;"+b64[begin:end])
		} else {
			out = append(out, "\x1B\x5C\x1BP"+b64[begin:end])
		}
		if end == len(b64) {
			break
		}
	}
	out = append(out, "\x07\x1B\x5C")
	return strings.Join(out, "")
}

func isTmuxCC(pid string) bool {
	out, err := exec.Command("ps", "-p", pid, "-o", "command=").Output()
	if err != nil {
		return false
	}
	out = bytes.TrimRight(out, "\n\r")
	for _, argv := range strings.Split(string(out), " ") {
		if argv == "-CC" {
			return true
		}
	}
	return false
}

func chooseEsc() func(string) string {
	if env := os.Getenv("TMUX"); env != "" {
		envs := strings.Split(env, ",")
		if len(envs) > 1 {
			pid := envs[1]
			if isTmuxCC(pid) {
				return normalEsc
			}
		}
		return tmuxEsc
	} else if env := os.Getenv("TERM"); strings.HasPrefix(env, "screen") {
		return screenEsc
	}
	return normalEsc
}

func run() error {
	var b []byte
	var err error
	if len(os.Args) == 1 {
		b, err = ioutil.ReadAll(os.Stdin)
	} else {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Print("Usage:\n  pbcopy FILE\n  some-command | pbcopy\n")
			os.Exit(1)
		}
		b, err = ioutil.ReadFile(os.Args[1])
	}
	if err != nil {
		return err
	}
	b = bytes.TrimRight(b, "\n\r")
	if len(b) == 0 {
		return nil
	}

	esc := chooseEsc()
	b64 := base64.RawStdEncoding.EncodeToString(b)
	fmt.Print(esc(b64))
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
