package pidfinder

import (
	"os/exec"
	"strings"
)

type ps struct {
	pid string
	tty string
	time string
	cmd string
}
var pslist []ps
func cmdtostruct (pscommandline string) {
	s := &ps{
		pid:  "",
		tty:  "",
		time: "",
		cmd:  "",
	}
	piddone := false
	ttydone := false
	timedone := false
	atom := strings.Split(pscommandline," ")
	for i:=0;i<len(atom);i++ {
		if atom[i] != "" {
			if !timedone {
				if !ttydone {
					if !piddone {
						s.pid = atom[i]
						piddone = true
					}else {
						s.tty = atom[i]
						ttydone = true
					}
				} else {
					s.time = atom[i]
					timedone = true
				}
			}else {
				s.cmd = atom[i]
			}
		}
	}
	pslist = append(pslist,*s)
}


func NewPS() []ps {
	pslist = make([]ps, 0, 0)
	pscommand := exec.Command("bash", "-c", "ps -e")
	pscommandlist, err := pscommand.Output()
	if err != nil {
		panic(err)
	}
	pscommandline := strings.Split(string(pscommandlist), "\n")


	for i:=1;i<len(pscommandline);i++ {
		cmdtostruct(pscommandline[i])
	}
	return pslist
}

func SearchPID(list []ps, name string) string {
	for i:= 0;i<len(list);i++ {
		if list[i].cmd == name {
			return list[i].pid
		}
	}
	return "err"
}


