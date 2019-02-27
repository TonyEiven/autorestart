package communicator

import (
	"log"

	"github.com/autorestart/cmd"
	cons "github.com/autorestart/constant"
	"github.com/masterzen/winrm"
)

type connectionInfo struct {
	User     string
	Passwd   string
	Host     string
	Port     int
	HTTPS    bool
	Insecure bool
}

type communicator struct {
	connInfo *connectionInfo
	client   *winrm.Client
	endpoint *winrm.Endpoint
}

func NewCommunicator(addr string) (*communicator, error) {
	connInfo := &connectionInfo{
		User:     cons.DefaultUser,
		Passwd:   cons.DefaultPassword,
		Host:     addr,
		Port:     cons.DefaultPort,
		HTTPS:    false,
		Insecure: false,
	}
	endpoint := &winrm.Endpoint{
		Host:     connInfo.Host,
		Port:     connInfo.Port,
		HTTPS:    connInfo.HTTPS,
		Insecure: connInfo.Insecure,
	}
	comm := &communicator{
		connInfo: connInfo,
		endpoint: endpoint,
	}
	return comm, nil
}

func (c *communicator) Connect() error {
	if c.client != nil {
		return nil
	}
	params := winrm.DefaultParameters
	client, err := winrm.NewClientWithParameters(
		c.endpoint, c.connInfo.User, c.connInfo.Passwd, params)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] connecting to remote shell using WinRM")
	shell, err := client.CreateShell()
	if err != nil {
		log.Printf("[ERROR] error creating shell: %s", err)
		return err
	}
	err = shell.Close()
	if err != nil {
		log.Printf("[ERROR] error closing shell: %s", err)
		return err
	}
	c.client = client
	return nil
}

func (c *communicator) Disconnect() error {
	c.client = nil
	return nil
}

func (c *communicator) Start(rc *cmd.Cmd) error {
	rc.Init()
	log.Printf("[DEBUG] starting remote command: %s", rc.Command)

	if c.client == nil {
		log.Println("[WARN] winrm client not connected, attempting to connect")
		if err := c.Connect(); err != nil {
			return err
		}
	}
	status, err := c.client.Run(rc.Command, rc.Stdout, rc.Stderr)
	rc.SetExitStatus(status, err)
	return nil
}
