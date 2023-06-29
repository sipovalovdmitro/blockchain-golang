package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)

	if err != nil {
		log.Panic(err)
		return []byte{}
	}

	return result.Bytes()
}
