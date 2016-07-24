package gitrss

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// git log --pretty=raw --no-color

type gitlog struct {
	commit    string
	tree      string
	parent    string
	author    *person
	committer *person
	message   string
}

type person struct {
	name  string
	email string
	date  time.Time
}

func parseGitlog(rdr io.Reader) (result []*gitlog) {
	var cur *gitlog
	var err error
	r := bufio.NewReader(rdr)
	for line := ""; err == nil; line, err = r.ReadString('\n') {
		line = strings.TrimSuffix(line, "\n")
		if strings.HasPrefix(line, "commit ") {
			if cur != nil {
				result = append(result, cur)
			}
			cur = &gitlog{commit: getHash(line)}
			continue
		}
		if strings.HasPrefix(line, "tree ") {
			cur.tree = getHash(line)
			continue
		}
		if strings.HasPrefix(line, "parent ") {
			cur.parent = getHash(line)
			continue
		}
		if strings.HasPrefix(line, "author ") {
			cur.author = parsePerson(line)
			continue
		}
		if strings.HasPrefix(line, "committer ") {
			cur.committer = parsePerson(line)
			continue
		}
		if strings.HasPrefix(line, "    ") || strings.HasPrefix(line, "\v") {
			if strings.HasPrefix(line, "    ") {
				line = strings.TrimPrefix(line, "    ")
			} else if strings.HasPrefix(line, "\v") {
				line = strings.TrimPrefix(line, "\v")
			}
			cur.message += line + "\n"
		}
	}
	if err != io.EOF {
		log.Fatal(err)
	}
	return
}

func getHash(line string) string {
	arr := strings.Fields(line)
	return arr[1]
}

// author Songmu <y.songmu@gmail.com> 1428772084 +0900
var personReg = regexp.MustCompile(`^([^<]+)\s<([^>]+)>\s([0-9]+)\s(.*)$`)

func parsePerson(line string) *person {
	arr := strings.SplitN(line, " ", 2)

	matches := personReg.FindStringSubmatch(arr[1])
	epoch := matches[3]
	tz := matches[4]
	_ = tz

	// lc, err := time.LoadLocation(tz)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	ep, err := strconv.ParseInt(epoch, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	t := time.Unix(ep, 0)
	// t.In(lc)
	return &person{
		name:  matches[1],
		email: matches[2],
		date:  t,
	}
}
