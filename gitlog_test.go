package gitrss

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGitlog(t *testing.T) {
	fp, err := os.Open("testdata/gitlog.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer fp.Close()
	r := parseGitlog(fp)
	lo := r[0]
	assert.Equal(t, "Songmu", lo.committer.name)
	assert.Equal(t, "y.songmu@gmail.com", lo.committer.email)
	assert.Equal(t, "2016-04-17 17:30:00 +0900 +0900", lo.committer.date.String())
	msg := `Checking in changes prior to tagging of version v0.9.8.

Changelog diff is:

diff --git a/Changes b/Changes
index b0c60be..9c690a3 100644
--- a/Changes
+++ b/Changes
@@ -2,6 +2,10 @@ Revision history for Perl extension Riji

 {{$NEXT}}

+v0.9.8 2016-04-17T08:29:51Z
+
+    - depends on MIME::Base32 1.301 or later and fix code around incompatible changes of it. fix #19
+
 v0.9.7 2015-04-12T12:55:49Z

     - support Front-matter
`
	assert.Equal(t, msg, lo.message)
}
