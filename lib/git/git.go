package git

import (
	"io/ioutil"
	"os"
	"os/exec"
)

type Repo struct {
	url string
	dir string
}

func Clone(URL string) (*Repo, error) {
	// TODO: Create per-URL temp dir with fmt.Sprintf("supbot-%x", sha1.Sum([]byte(URL)))
	//       and then create (shallow) clones locally to speed things up.
	dir, err := ioutil.TempDir(os.TempDir(), "supbot")
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("git", "clone", URL)
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return &Repo{
		url: URL,
		dir: dir,
	}, nil
}

func (r *Repo) Checkout(branch string) error {
	cmd := exec.Command("git", "checkout", "-t", "origin/"+branch)
	cmd.Dir = r.dir
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (r *Repo) Dir() string {
	return r.dir
}
