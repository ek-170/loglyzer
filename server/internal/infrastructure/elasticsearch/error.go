package elasticsearch

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type ElasticsearchError struct {
	Status int
	Failed string
	Reason string
}

func HandleElasticsearchError(err error) string {
	log.Println(err.Error())
	esErr, err := parseElasticsearchError(err.Error())
	if err != nil {
		return err.Error()
	}
	return convertMsg4EndUser(esErr)
}

func parseElasticsearchError(errorStr string) (ElasticsearchError, error) {
	var result ElasticsearchError

	re := regexp.MustCompile(`status: (\d+), failed: \[(.*)\], reason: (.*)`)
	matches := re.FindStringSubmatch(errorStr)
	if len(matches) != 4 {
    log.Println("failed to parse Elasticsearch error message")
		return result, errors.New("unknown error has occured when communicate Elasticsearch. please check docker log")
	}
	result.Status = parseInt(matches[1])
	result.Failed = matches[2]
	result.Reason = matches[3]

	return result, nil
}

func parseInt(s string) int {
	i := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			break
		}
		i = i*10 + int(c-'0')
	}
	return i
}

func convertMsg4EndUser(err ElasticsearchError) string {
	if err.Failed == ES_F00001 {
		return ES_EM00001
	}
	if err.Status == http.StatusNotFound {
		return ES_EM00003
	}
	return fmt.Sprintf("Elasticsearch Error has occured. failed: %s, reason: %s", err.Failed, err.Reason)
}