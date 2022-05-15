package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tanakornwry/mars-exploration-project/entities"
)

type input struct {
}

type Input interface {
	ReadCommand() entities.CommandConf
}

func NewInput() Input {
	return &input{}
}

func (i *input) ReadCommand() entities.CommandConf {
	// filepath := filepath.Join("$GOPATH", "src", "mars-exploration-project", "ftp", "command.txt")
	// Use a relative path to easiest to test and does not worried about $GOPATH
	filepath := "../../ftp/command.txt"

	if existed := fileExists(filepath); !existed {
		log.Fatal("Not found the command file")
	}

	f, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var rawCommand []string
	for scanner.Scan() {
		rawCommand = append(rawCommand, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	commandConf := refactorCommand(rawCommand)

	return commandConf
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func refactorCommand(rawCommand []string) entities.CommandConf {
	var cc entities.CommandConf
	for i, v := range rawCommand {
		if i == 0 {
			asInt, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("Can not defind the size map because:", err)
			}
			cc.SizeMap = asInt
		} else {
			for _, c := range strings.Split(v, " ") {
				subComm := strings.Split(c, "")
				if len(subComm) > 1 {
					asInt, _ := strconv.Atoi(subComm[1])
					if asInt >= 1 {
						if validateCommand(subComm[0]) {
							cc.Command = append(cc.Command, map[string]int{subComm[0]: asInt})
						}
					}

				} else {
					if validateCommand(subComm[0]) {
						cc.Command = append(cc.Command, map[string]int{subComm[0]: 1})
					}
				}
			}
		}
	}

	return cc
}

func validateCommand(c string) bool {
	if c == "L" || c == "R" || c == "F" || c == "B" {
		return true
	}

	return false
}
