package store

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	DIRNAME  = "."
	FILENAME = "test.db"
)

func TestStore(t *testing.T) {

	os.Remove(DIRNAME + "/" + FILENAME)

	store := NewFileBlobStore(DIRNAME, FILENAME)

	{
		// write and close
		err := store.Open()
		assert.Nil(t, err)
		store.Append([]byte("first_record"))
		store.Close()
	}

	{
		// write and close
		err := store.Open()
		assert.Nil(t, err)
		err = store.Append([]byte("second_record"))
		assert.Nil(t, err)
		err = store.Append([]byte("third_record"))
		assert.Nil(t, err)
		err = store.Append([]byte("fourth_record"))
		assert.Nil(t, err)
		store.Close()
	}

	{
		// read all and close
		err := store.Open()
		assert.Nil(t, err)
		records := make([][]byte, 0, 10)
		cb := func(record []byte) {
			records = append(records, record)
		}
		store.Iterate(cb)
		assert.Equal(t, 4, len(records))

		assert.Equal(t, []byte("first_record"), records[0])
		assert.Equal(t, []byte("second_record"), records[1])
		assert.Equal(t, []byte("third_record"), records[2])
		assert.Equal(t, []byte("fourth_record"), records[3])

		store.Close()
	}

	store.Close()
	os.Remove(DIRNAME + "/" + FILENAME)
}

func BenchmarkWrite(b *testing.B) {
	os.Remove(DIRNAME + "/" + FILENAME)

	store := NewFileBlobStore(DIRNAME, FILENAME)
	store.Open()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		store.Append([]byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"))
		reader := func(record []byte) {
		}
		store.Iterate(reader)
	}

	store.Close()
	os.Remove(DIRNAME + "/" + FILENAME)
}
