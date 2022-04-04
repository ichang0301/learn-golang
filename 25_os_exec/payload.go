package os_exec

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

type Payload struct {
	Message string `xml:"message"`
}

func GetData(data io.Reader) string {
	var payload Payload
	xml.NewDecoder(data).Decode(&payload)

	return strings.ToUpper(payload.Message)
}

func GetXMLFromCommand(name string, filePath string) io.Reader {
	cmd := exec.Command(name, filePath)
	out, _ := cmd.StdoutPipe() // these 3 can return errors but I'm ignoring for brevity

	cmd.Start()
	data, _ := ioutil.ReadAll(out)
	cmd.Wait()

	return bytes.NewReader(data)
}
