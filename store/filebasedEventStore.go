package store

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xebia/microgen/events"
	"io"
	"log"
	"os"
	"sync"
)

type SimpleEventStore struct {
	mutex              sync.RWMutex
	filename           string
	fio                *os.File
	lastSequenceNumber uint64
}

func NewSimpleEventStore() *SimpleEventStore {
	store := new(SimpleEventStore)
	return store
}

func (store *SimpleEventStore) Open(filename string) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	var err error
	store.filename = filename
	store.fio, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file %s (%v)", store.filename, err)
		return err
	}
	//log.Printf("Opened file %s", store.filename)
	store.lastSequenceNumber = store.getLastSequenceNumber()
	return nil
}

func (store *SimpleEventStore) Store(envelope *events.Envelope) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	return store.writeEvent(envelope)
}

func (store *SimpleEventStore) writeEvent(envelope *events.Envelope) error {
	_, err := store.fio.Seek(0, os.SEEK_END)
	if err != nil {
		log.Printf("Error going to end of file (%v)", err)
		return err
	}

    // assign incementing sequence number to determines order of events
    store.assignSequenceNumber(envelope)

    //log.Printf("write event: %v\n", envelope )

	// serialize event to json
	jsonBlob, err := json.Marshal(envelope)
	if err != nil {
		return errors.New(fmt.Sprintf("Error marshalling event (%v)", err))
	}
	//log.Printf("Marshalled envelope of type %d into %d bytes", envelope.Type, len(jsonBlob))

	// write length to buffer
	err = binary.Write(store.fio, binary.BigEndian, int64(len(jsonBlob)))
	if err != nil {
		log.Printf("error writing size to file %s (%v)", store.filename, err)
		return err
	}
	//log.Printf("Written len %d", len(jsonBlob))

	// write json blob to buffer
	_, err = store.fio.Write(jsonBlob)
	if err != nil {
		log.Printf("error writing blob to file %s (%v)", store.filename, err)
		return err
	}
	//log.Printf("Write blob %d", written)

	err = store.fio.Sync()
	if err != nil {
		log.Printf("error syncing file %s (%v)", store.filename, err)
		return err
	}

	return nil
}

func (store *SimpleEventStore) Iterate(handlerFunc events.StoredItemHandlerFunc) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

    return store.iterate( handlerFunc)
}

func (store *SimpleEventStore) iterate(handlerFunc events.StoredItemHandlerFunc) error {
	_, err := store.fio.Seek(0, os.SEEK_SET)
	if err != nil {
		log.Printf("Error going to start of file (%v)", err)
		return err
	}
	for {
		envelope, err := store.readNextEvent()
		if err != nil {
			return err
		} else if envelope == nil {
			break
		}
		done := handlerFunc(envelope)
		if done == true {
			break
		}
	}
	return nil
}

func (store *SimpleEventStore) readNextEvent() (*events.Envelope, error) {

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
	jsonDataBuffer := make([]byte, jsonLength)
	_, err = io.ReadFull(store.fio, jsonDataBuffer)
	if err != nil {
		log.Printf("Error reading json blob (%v)", err)
		return nil, err
	}
	//log.Printf("Read blob with length (%d)", read)

	// unserialize blob
	envelope := events.Envelope{Type: events.TypeUnknown}
	err = json.Unmarshal(jsonDataBuffer, &envelope)
	if err != nil {
		log.Printf("Error unmarshalling json blob (%v)", err)
		return nil, err
	}
	//log.Printf("Unmarshalled blob of type %d", envelope.Type)
    //log.Printf("read event: %v\n", envelope )

	return &envelope, nil
}

func (store *SimpleEventStore) assignSequenceNumber( envelope *events.Envelope) {
    store.lastSequenceNumber = store.lastSequenceNumber+1
    envelope.SequenceNumber = store.lastSequenceNumber
}

func (store *SimpleEventStore) getLastSequenceNumber() uint64 {
	var lastIndex uint64 = 0

	callback := func(envelope *events.Envelope) bool {
		lastIndex++
		return false
	}
	store.iterate(callback)

	return lastIndex
}

func (store *SimpleEventStore) Close() {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if store.fio != nil {
		//log.Printf("Closed file %s", store.filename)
		store.fio.Close()
	}
}
