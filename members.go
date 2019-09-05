package djp

import "github.com/buger/jsonparser"

func MemberReplaceUserWithID(data []byte, keys ...string) []byte {
	type userKey struct {
		offset int
		length int
		id     []byte
	}

	var userKeys []userKey
	_, err := jsonparser.ArrayEach(data, func(d []byte, _ jsonparser.ValueType, offset int, _ error) {
		id, _, _, _ := jsonparser.Get(d, "user", "id")
		usrData, _, usrOffset, _ := jsonparser.Get(d, "user")
		length := len(usrData) + len(`"user":`)
		offset += usrOffset - len(`"user":`) - len(usrData)

		userKeys = append(userKeys, userKey{offset, length, id})
	}, keys...)
	if err != nil {
		return data
	}

	// replace with id
	const idPrefix = `"id":"`
	id := make([]byte, 0, len(idPrefix)+25)
	id = append(id, idPrefix...)
	for i := len(userKeys) - 1; i >= 0; i-- {
		k := userKeys[i]
		id = append(id, k.id...)
		id = append(id, '"')
		p := len(id) - 1
		for j := k.offset + k.length - 1; j >= k.offset && p >= 0; j-- {
			ii := id[p]
			data[j] = ii
			p--
		}
		userKeys[i].length -= len(id)
		id = id[:len(idPrefix)]
	}

	// shift
	for i := len(userKeys) - 1; i >= 0; i-- {
		k := userKeys[i]
		o := k.offset
		for j := o + k.length; j < len(data); j++ {
			data[o] = data[j]
			o++
		}
		data = data[:len(data)-k.length]
	}

	return data
}
