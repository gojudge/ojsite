package judger

import (
	"github.com/duguying/judger/client"
	"github.com/gogather/com/log"
)

var J map[string]interface{}

func ConnectJudger(host string, port int, password string) error {
	J = make(map[string]interface{})

	cli, err := client.New(host, port, password)
	J["default"] = cli

	if err != nil {
		log.Warnln(err)
		return err
	}

	cli.SetDebug(true)

	return nil
}

func Get(name string) *client.JClient {
	jdg, ok := J[name].(*client.JClient)
	if !ok {
		log.Warnln("get failed")
		return nil
	}

	return jdg
}
