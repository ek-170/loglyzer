package filereader

import (
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/ssh"
)

type SshFileReader struct {
	path string
	sshKeyPath   string
	userName     string
	password     string
	host         string
	port         string
}

func (sfr *SshFileReader) ReadFile() (io.Reader, error) {
	sshConfig :=  &ssh.ClientConfig{}
	if sfr.userName!="" && sfr.sshKeyPath != "" {
		signer, err := readPrivateKeyFile(sfr.sshKeyPath)
		if err != nil {
			return nil, err
		}
		sshConfig = &ssh.ClientConfig{
			User: sfr.userName,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
		}
	} else if sfr.userName!="" && sfr.password!="" {
		sshConfig = &ssh.ClientConfig{
			User: sfr.userName,
			Auth: []ssh.AuthMethod{
				ssh.Password(sfr.password),
			},
		}
	}else {
		return nil, errors.New("Not sufficient credential(user Name is required, and either password or sshKeyPath is required)")
	}
	
	hostWithPort := sfr.host + ":" +  sfr.port
	client, err := ssh.Dial("tcp", hostWithPort, sshConfig)
	if err != nil {
		log.Println("Error establishing SSH connection:", err)
		return nil, errors.New("Error establishing SSH connection")
	}
	defer client.Close()

	// start SSH session
	session, err := client.NewSession()
	if err != nil {
		log.Println("Error creating SSH session:", err)
		return nil, errors.New("Error creating SSH session")
	}
	defer session.Close()

	remoteFile, err := session.Open(sfr.path)
	if err != nil {
		return nil, errors.New("Error establishing SSH connection")
	}
	return remoteFile, nil
}

func readPrivateKeyFile(path string) (ssh.Signer, error) {
	privateKey, err := os.ReadFile(path)
	if err != nil {
		log.Println("Error reading private key file:", err)
		return nil, errors.New("Error reading private key file")
	}

	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		log.Println("Error parsing private key:", err)
		return nil, errors.New("Error parsing private key")
	}

	return signer, nil
}

func NewSshFileReader(conf FileReaderConfig) SshFileReader{
	sfr := SshFileReader{}
	sfr.path = conf.Path
	sfr.sshKeyPath = conf.SshKeyPath
	sfr.userName = conf.UserName
	sfr.password = conf.Password
	sfr.host = conf.Host
	noPortValue := 0 
	if conf.Port != noPortValue {
		sfr.port = strconv.Itoa(conf.Port)
	}else {
		// SSH default Port
		sfr.port = "22"
	}
	return sfr
}