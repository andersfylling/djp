package discordjsonparser

import "testing"

func TestMemberReplaceUserWithID(t *testing.T) {
	type entry struct {
		input  []byte
		output []byte
	}
	dataset := []entry{
		{
			[]byte(`{"members":[{"ok":true,"user":{"id":"12313"}},{"user":{"id":"2"},"ok":false},{"user":{"id":"3"}}]}`),
			[]byte(`{"members":[{"ok":true,"id":"12313"},{"id":"2","ok":false},{"id":"3"}]}`),
		},
		{
			[]byte(`{"members":[{"user":{"id":"1"}},{"user":{"id":"2"}},{"user":{"id":"3"}}]}`),
			[]byte(`{"members":[{"id":"1"},{"id":"2"},{"id":"3"}]}`),
		},
		{
			[]byte(`{"members":[{"user":{}},{"user":{"id":"2"}},{"user":{"id":"3"}}]}`),
			[]byte(`{"members":[{"id":""},{"id":"2"},{"id":"3"}]}`), // TODO: remove id
		},
	}

	for i := range dataset {
		input := dataset[i].input
		wants := string(dataset[i].output)
		got := string(MemberReplaceUserWithID(input, "members"))

		if wants != got {
			t.Errorf("results were not as expected \ngot\t\t%s\nwants\t%s", got, wants)
		}
	}
}
