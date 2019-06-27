package returncodes

import (
	"encoding/json"
	"github.com/pkg/errors"
	"testing"
)

func TestS(t *testing.T) {
	coder := ErrTimeout
	bytes, e := json.Marshal(coder)
	t.Logf("%s, %v", bytes, e)


	bytes, e = json.Marshal(Data("data"))
	t.Logf("%s, %v", bytes, e)
	bytes, e = json.Marshal(Mess("ok"))
	t.Logf("%s, %v", bytes, e)

	bytes, e = json.Marshal(Succ("aaa", "aaaa"))
	t.Logf("%s, %v", bytes, e)

	bytes, e = json.Marshal(Fail("a"))
	t.Logf("%s, %v", bytes, e)

	bytes, e = json.Marshal(Fail(errors.New("a")))
	t.Logf("%s, %v", bytes, e)

	bytes, e = json.Marshal(ErrArgument)
	t.Logf("%s, %v", bytes, e)
}
