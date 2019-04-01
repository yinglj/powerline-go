package main

import (
	"fmt"
	"strconv"
)

func getMeaningFromExitCode(exitCode int) string {
	switch exitCode {
	case 1:
		return "ERROR"
	case 2:
		return "USAGE"
	case 126:
		return "NOPERM"
	case 127:
		return "NOTFOUND"
	case 128 + 1:
		return "SIGHUP"
	case 128 + 2:
		return "SIGINT"
	case 128 + 3:
		return "SIGQUIT"
	case 128 + 4:
		return "SIGILL"
	case 128 + 5:
		return "SIGTRAP"
	case 128 + 6:
		return "SIGIOT"
	case 128 + 7:
		return "SIGBUS"
	case 128 + 8:
		return "SIGFPE"
	case 128 + 9:
		return "SIGKILL"
	case 128 + 10:
		return "SIGUSR1"
	case 128 + 11:
		return "SIGSEGV"
	case 128 + 12:
		return "SIGUSR2"
	case 128 + 13:
		return "SIGPIPE"
	case 128 + 14:
		return "SIGALRM"
	case 128 + 15:
		return "SIGTERM"
	case 128 + 16:
		return "SIGSTKFLT"
	case 128 + 17:
		return "SIGCHLD"
	case 128 + 18:
		return "SIGCONT"
	case 128 + 19:
		return "SIGSTOP"
	case 128 + 20:
		return "SIGTSTP"
	case 128 + 21:
		return "SIGTTIN"
	case 128 + 22:
		return "SIGTTOU"
	default:
		return fmt.Sprintf("%d", exitCode)
	}
}

func segmentExitCode(p *powerline) {
	if *p.args.ShowExitCode == true {
		var meaning string
		if *p.args.PrevError != 0 {
			if *p.args.NumericExitCodes {
				meaning = strconv.Itoa(*p.args.PrevError)
			} else {
				meaning = getMeaningFromExitCode(*p.args.PrevError)
			}
			p.appendSegment("exit", segment{
				content:    meaning,
				foreground: p.theme.CmdFailedFg,
				background: p.theme.CmdFailedBg,
			})
		}
	}
}
