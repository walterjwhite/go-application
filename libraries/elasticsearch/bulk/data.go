package bulk

import (
	"context"
	"fmt"
	"github.com/walterjwhite/go-application/libraries/elasticsearch"
	"github.com/walterjwhite/go-application/libraries/logging"
	"gopkg.in/olivere/elastic.v7"
	"time"
)

type Operation int

const (
	Index Operation = iota
	Update
	Delete
)

type MasterBatch struct {
	nodeConfiguration *elasticsearch.NodeConfiguration
	bulkProcessor     *elastic.BulkProcessor
}

// problem is that we don't have the granularity to report this now
/*
type RecordOperation struct {
	Operation Operation
	DocumentType string
	DocumentId string
}

type Batch struct {
	RecordOperations []RecordOperation
}
*/

func NewDefaultBatch(c *elasticsearch.NodeConfiguration) *MasterBatch {
	// TODO: make this dynamic based on the hardware
	return NewBatch(c, 10, 10, 1*time.Second, 2)
}

func NewBatch(c *elasticsearch.NodeConfiguration, actionSize int, dataSize int, interval time.Duration, workers int) *MasterBatch {
	masterBatch := MasterBatch{nodeConfiguration: c}

	// TODO: make the size configurable
	p, err := c.Client.BulkProcessor().
		Name(getProcessorName(c)).
		Workers(workers).
		BulkActions(actionSize).
		// default is 5 MB
		//BulkSize(2 << 20).
		FlushInterval(interval).
		Do(context.Background())

	logging.Panic(err)

	masterBatch.bulkProcessor = p

	return &masterBatch
}

func getProcessorName(c *elasticsearch.NodeConfiguration) string {
	return fmt.Sprintf("bulkProcessor.%v", c.IndexPrefix)
}

/*
func (b *MasterBatch) getDocumentTypeName(document interface{}) string {
	documentTypeName := typename.Get(document)
	b.nodeConfiguration.PrepareIndex(documentTypeName)

	return documentTypeName
}

func (b *MasterBatch) getIndexName(documentTypeName string) string {
	return b.nodeConfiguration.getIndexName(documentTypeName)
}
*/

func (b *MasterBatch) Flush() {
	logging.Panic(b.bulkProcessor.Flush())
}

/*
type BulkCommandFailed struct {
	RemainingCommands int
}

func (b *BulkCommandFailed) Error() string {
	return fmt.Sprintf("Bulk command has %v commands remaining, but should be 0\n", b.RemainingCommands)
}
*/
