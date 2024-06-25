package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// url is the github url https://github.com/deepaksinghvi/cdule
	// directory is the local director where it should be cloned /tmp/cdule
	// token is the personal access token
	url, directory, token := os.Args[1], os.Args[2], os.Args[3]

	// Clone the given repository to the given directory
	log.Infof("git clone %s %s", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "abc123", // yes, this can be anything except an empty string
			Password: token,
		},
		URL:      url,
		Progress: os.Stdout,
	})
	if nil != err {
		log.Errorf("Failed to clone %s: %s", url, err)
	} else {
		log.Infof("Cloned %s", url)
		log.Infof("Repository %v", r)
	}
}
