package git

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type Repo struct {
	url string
	dir string
}

// execCommand execs the given command on the command's working directory and
// prints debugging information.
func execCommand(cwd string, name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = cwd
	log.Printf("cwd: %q", cmd.Dir)

	var stdout, stderr bytes.Buffer

	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		log.Printf("err: %s", string(stderr.Bytes()))
		return stderr.Bytes(), err
	}

	log.Printf("out: %s", string(stderr.Bytes()))
	return stdout.Bytes(), nil
}

func Clone(uri string) (*Repo, error) {
	// TODO: Create per-URL temp dir with fmt.Sprintf("supbot-%x", sha1.Sum([]byte(URL)))
	//       and then create (shallow) clones locally to speed things up.
	dir, err := ioutil.TempDir(os.TempDir(), "supbot")
	if err != nil {
		return nil, err
	}

	_, err = execCommand(dir, "git", "clone", uri, ".")
	if err != nil {
		return nil, err
	}

	return &Repo{
		url: uri,
		dir: dir,
	}, nil
}

func (r *Repo) Checkout(branch string) error {
	if r.dir == "" {
		return errors.New("Missing dir.")
	}
	//_, err := execCommand(r.dir, "git", "checkout", "-t", "origin/"+branch) // says No error is expected but got exit status 128
	_, err := execCommand(r.dir, "git", "checkout", "origin/"+branch)
	return err
}

func (r *Repo) Dir() string {
	return r.dir
}
