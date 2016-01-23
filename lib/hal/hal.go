package hal

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/gophergala2016/supbot/lib/git"
	"github.com/gophergala2016/supbot/lib/sup"
)

var space = []byte(` `)

// Making sure this is a reader.
var _ = io.Writer(&Hal{})

var (
	msgMissingCommand = `Say what?!`
)

type Hal struct {
	out  io.Writer
	repo string // current working directory.
}

var (
	errMissingCommand    = errors.New(`Missing command.`)
	errIncompleteCommand = errors.New(`Incomplete command.`)
)

var db *bolt.DB

func init() {
	var err error
	db, err = bolt.Open("hal.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket([]byte("settings"))
		return nil
	})

}

func NewHal(out io.Writer) *Hal {
	h := &Hal{
		out: out,
	}

	h.restore()

	return h
}

func (h *Hal) reset() error {
	h.repo = ""
	return nil
}

func (h *Hal) save() error {
	// HAL remembers settings.
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("settings"))
		b.Put([]byte("repo"), []byte(h.repo))
		return nil
	})
	return err
}

func (h *Hal) restore() error {
	// HAL remembers settings.
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("settings"))
		v := b.Get([]byte("repo"))
		h.repo = string(v)
		return nil
	})
	return err
}

func (h *Hal) Write(cmd []byte) (n int, err error) {
	l := len(cmd)

	chunks := bytes.Split(cmd, space)

	if len(chunks) < 1 {
		h.out.Write([]byte(msgMissingCommand))
		return l, errMissingCommand
	}

	s := string(chunks[0])

	switch s {
	case "help":
		h.out.Write([]byte(`[repository]/[branch] [network] [target]`))
		return l, nil
	case "wipe":
		h.reset()
		h.out.Write([]byte(`Now I don't have a memory.`))
		return l, nil
	default:
		if len(chunks) > 0 {
			switch string(chunks[0]) {
			case "set-repo":
				if len(chunks) > 1 {
					repo := string(chunks[1])
					if repo != "" {
						// TODO: check this is an actual repo.
						h.repo = repo
						h.save()
						h.out.Write([]byte(fmt.Sprintf("You current repo is %q", h.repo)))
						return l, nil
					}
				}
				h.out.Write([]byte(fmt.Sprintf("Try `set-repo [repo-url]`")))
				return l, errMissingCommand
			}
			if h.repo != "" {

				h.out.Write([]byte(fmt.Sprintf("Hang in there, I'm cloning %q...", h.repo)))

				// TODO: grab branch name from URL, if any.
				repo, err := git.Clone(h.repo)
				if err != nil {
					return l, err
				}

				if err := repo.Checkout("master"); err != nil {
					return l, err
				}

				h.out.Write([]byte(fmt.Sprintf("Running sup...")))

				// TODO: insert sup magic here.
				var outbuf bytes.Buffer
				cmd := sup.NewSup(&outbuf).Setwd(repo.Dir())
				defer func() {
					log.Printf("Cleaning %v", repo.Dir())
					os.RemoveAll(repo.Dir())
				}()

				if len(chunks) > 0 {
					cmd.Network(string(chunks[0]))
				}
				if len(chunks) > 1 {
					cmd.Target(string(chunks[1]))
				}

				err = cmd.Exec()

				h.out.Write(outbuf.Bytes())
				return l, err
			} else {
				h.out.Write([]byte(fmt.Sprintf("Missing repo, try `set-repo [repo-url]`")))
			}
			return l, errMissingCommand
		}
	}

	return l, nil
}
