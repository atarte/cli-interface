package utils

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type WindSize struct {
	Row int
	Col int
}

func GetWindSize() WindSize {

	cmd := exec.Command("stty", "size")

	cmd.Stdin = os.Stdin

	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Console size data could not be accuired, check if stty is install on your device, err :" + err.Error())
	}

	dataSplit := strings.Split(string(out), " ")
	dataSplit[1] = dataSplit[1][:2]

	row, _ := strconv.ParseInt(dataSplit[0], 10, 32)
	col, _ := strconv.ParseInt(dataSplit[1], 10, 32)

	return WindSize{
		Row: int(row),
		Col: int(col),
	}
}

func (w *WindSize) RefeshWindowsSize() {

}

func (w *WindSize) SetWindSize(col, row int) {

}

func IsSTTYAvailable() bool {
	return true
}
