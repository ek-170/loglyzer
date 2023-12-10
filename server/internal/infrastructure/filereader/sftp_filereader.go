package filereader

import (
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	fr "github.com/ek-170/loglyzer/internal/domain/filereader"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SftpFileReader struct {
	path string
	sshKeyPath   string
	userName     string
	password     string
	host         string
	port         string
}

func (sfr *SftpFileReader) ReadFile() (io.Reader, error) {
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
		// case of insufficient credentials
		return nil, errors.New("insufficient credential(user Name is required, and either password or sshKeyPath is required)")
	}
	
	hostWithPort := sfr.host + ":" +  sfr.port
	sshClient, err := ssh.Dial("tcp", hostWithPort, sshConfig)
	if err != nil {
		log.Println("Error establishing SSH connection:", err)
		return nil, errors.New("establishing SSH connection error")
	}
	defer sshClient.Close()

	client, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Println("Error creating SFTP client:", err)
		return nil, errors.New("creating SFTP client error")
	}

	remoteFile, err := client.Open(sfr.path)
	if err != nil {
		return nil, errors.New("establishing SSH connection error")
	}
	return remoteFile, nil
}

func readPrivateKeyFile(path string) (ssh.Signer, error) {
	privateKey, err := os.ReadFile(path)
	if err != nil {
		log.Println("Error reading private key file:", err)
		return nil, errors.New("reading private key file error")
	}

	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		log.Println("Error parsing private key:", err)
		return nil, errors.New("parsing private key error")
	}

	return signer, nil
}

func NewSftpFileReader(conf fr.FileReaderConfig) *SftpFileReader{
	sfr := SftpFileReader{}
	sfr.path = conf.Path
	sfr.sshKeyPath = conf.SshKeyPath
	sfr.userName = conf.UserName
	sfr.password = conf.Password
	sfr.host = conf.Host
	if conf.Port != nil {
		sfr.port = strconv.Itoa(*conf.Port)
	}else {
		// SSH default Port
		sfr.port = "22"
	}
	return &sfr
}