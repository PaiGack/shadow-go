package shadow_go

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strings"
)

func gitStatusFiles(args []string, grep string) []string {
	gitCmd := exec.Command("git", args...)
	grpCmd := exec.Command("findstr", grep)

	outPutPip, _ := gitCmd.StdoutPipe()
	grepInPutPip, _ := grpCmd.StdinPipe()

	if err := gitCmd.Start(); err != nil {
		return nil
	}

	data, _ := ioutil.ReadAll(outPutPip)

	if err := gitCmd.Wait(); err != nil {
		return nil
	}

	var grepOutBuffer bytes.Buffer
	grpCmd.Stdout = &grepOutBuffer
	if err := grpCmd.Start(); err != nil {
		return nil
	}

	_, _ = grepInPutPip.Write(data)
	_ = grepInPutPip.Close()

	if err := grpCmd.Wait(); err != nil {
		return nil
	}

	var ret []string
	for _, str := range strings.Split(grepOutBuffer.String(), "\n") {
		tmps := strings.Split(str, " ")
		if len(tmps) > 1 {
			ret = append(ret, tmps[len(tmps)-1])
		}
	}
	return ret
}
