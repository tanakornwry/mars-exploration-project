package input

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/tanakornwry/mars-exploration-project/config"
	"github.com/tanakornwry/mars-exploration-project/entities"
)

type input struct {
}

type Input interface {
	ReadCommand() (entities.CommandConf, string)
}

func NewInput() Input {
	return &input{}
}

func (i *input) ReadCommand() (entities.CommandConf, string) {
	conf := config.LoadConfiguration("./config/config.json")

	// filepath := filepath.Join("$GOPATH", "src", "mars-exploration-project", "ftp", "command.txt")
	// Use a relative path to easiest to test and does not worried about $GOPATH
	filepath := filepath.Join(conf.CommandFile.Path, conf.CommandFile.Filename)
	if existed := fileExists(filepath); !existed {
		if existed := fileExists("../../ftp/command_testcase.txt"); existed {
			// Set this to allow pass the test cases
			filepath = "../../ftp/command_testcase.txt"
		}
	}

	f, err := os.Open(filepath)

	if err != nil {
		log.Println("Error: Not found the command file.")
		return entities.CommandConf{}, "Error: Not found the command file."
	}
	defer f.Close()
	log.Println("Reading the command file from", filepath)

	scanner := bufio.NewScanner(f)

	var rawCommand []string
	for scanner.Scan() {
		rawCommand = append(rawCommand, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	commandConf := refactorCommand(rawCommand)

	return commandConf, ""
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
		v = strings.Trim(v, " ")
		if i == 0 { // The first line is always size map
			asInt, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("Can not defind the size map because:", err)
			}
			cc.SizeMap = asInt
		} else {
			for _, c := range strings.Split(v, " ") {
				newCC := analyzeInstruction(c)
				if len(newCC.Command) > 0 {
					cc.Command = append(cc.Command, newCC.Command...)
				}
			}
		}
	}

	return cc
}

func analyzeInstruction(command string) entities.CommandConf {
	var cc entities.CommandConf

	cSet := []rune(command)
	if len(cSet) > 1 {
		var inst string
		var unit string
		if cSet[0] >= 65 && cSet[0] <= 90 {
			// First charecter
			inst = inst + string(cSet[0])

			// Second charecter
			if cSet[1] >= 65 && cSet[1] <= 90 {
				inst = inst + string(cSet[1])
			} else {
				unit = unit + string(cSet[1])
			}

			if validInstruction(inst) {

				// The rest charecter
				for i := 2; i < len(cSet); i++ {
					unit = unit + string(cSet[i])
				}

				if unit == "" {
					cc.Command = append(cc.Command, map[string]int{inst: 1})
				} else {
					unit, err := strconv.Atoi(unit)
					if err == nil && unit >= 1 && (inst == "F" || inst == "B") {
						for i := 0; i < unit; i++ {
							cc.Command = append(cc.Command, map[string]int{inst: 1})
						}
					} else {
						log.Println("Checkpoint 3 - Found incorrect instruction:", command)
					}
				}

			} else {
				log.Println("Checkpoint 2 - Found incorrect instruction:", command)
			}

		} else {
			log.Println("Checkpoint 1 - Found incorrect instruction:", command)
		}

	} else {
		if validInstruction(string(cSet[0])) {
			cc.Command = append(cc.Command, map[string]int{string(cSet[0]): 1})
		} else {
			log.Println("Found incorrect instruction:", command)
		}
	}

	return cc
}

func validInstruction(str string) bool {
	for _, v := range entities.Instruction.AllowInstruction {
		if v == str {
			return true
		}
	}

	return false
}
