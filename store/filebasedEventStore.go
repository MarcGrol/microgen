package store

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

type SimpleEventStore struct {
	filename string
	fio      *os.File
}

func NewSimpleEventStore() *SimpleEventStore {
	store := new(SimpleEventStore)
	return store
}

func (store *SimpleEventStore) Open(filename string) error {
	var err error
	store.filename = filename
	store.fio, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file %s (%v)", store.filename, err)
		return err
	}
	//log.Printf("Opened file %s", store.filename)
	return nil
}

func (store *SimpleEventStore) Append(blob []byte) error {
	_, err := store.fio.Seek(0, os.SEEK_END)
	if err != nil {
		log.Printf("Error going to end of file (%v)", err)
		return err
	}

	// write length to buffer
	err = binary.Write(store.fio, binary.BigEndian, int64(len(blob)))
	if err != nil {
		log.Printf("error writing size to file %s (%v)", store.filename, err)
		return err
	}
	//log.Printf("Written len %d", len(blob))

	// write json blob to buffer
	_, err = store.fio.Write(blob)
	if err != nil {
		log.Printf("error writing blob to file %s (%v)", store.filename, err)
		return err
	}
	//log.Printf("Write blob %d", written)

	// only return when file is on disk
	err = store.fio.Sync()
	if err != nil {
		log.Printf("error syncing file %s (%v)", store.filename, err)
		return err
	}

	return nil
}

type BlobHandlerFunc func(blob []byte)

func (store *SimpleEventStore) Iterate(handlerFunc BlobHandlerFunc) error {
	_, err := store.fio.Seek(0, os.SEEK_SET)
	if err != nil {
		log.Printf("Error going to start of file (%v)", err)
		return err
	}
	for {
		blob, err := store.readNextEvent()
		if err != nil {
			return err
		} else if blob == nil {
			break
		}
		handlerFunc(blob)
	}
	return nil
}

func (store *SimpleEventStore) readNextEvent() ([]byte, error) {

	// read length
	var jsonLength int64
	err := binary.Read(store.fio, binary.BigEndian, &jsonLength)
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}
		log.Printf("Error reading record length (%v)", err)
		return nil, err
	}
	//log.Printf("Read record length (%d)", jsonLength)

	// read blob
	blob := make([]byte, jsonLength)
	_, err = io.ReadFull(store.fio, blob)
	if err != nil {
		log.Printf("Error reading json blob (%v)", err)
		return nil, err
	}
	//log.Printf("Read blob with length %d (%s)", read, blob)

	return blob, nil
}

func (store *SimpleEventStore) Close() {
	if store.fio != nil {
		//log.Printf("Closed file %s", store.filename)
		store.fio.Close()
	}
}
