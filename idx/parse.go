package idx

import (
	"encoding/binary"
	"github.com/pkg/errors"
	"os"
)

type IDXData struct {
	Labels         []int
	Data           []int
	DataDimensions []int
	file           *os.File
}

func (iD *IDXData) ParseLabels(filepath string) (err error) {
	iD.file, err = os.Open(filepath)
	if err != nil {
		return errors.Wrap(err, "Failed to open file, check filepath")
	}

	_, numDimensions, err := iD.parseMagic()
	_, err = iD.parseDimensions(numDimensions)
	iD.Labels, err = iD.parseDataBytes(1)

	// Clean up file pointer
	iD.file = nil

	return nil
}

func (iD *IDXData) ParseData(filepath string) (err error) {
	iD.file, err = os.Open(filepath)
	if err != nil {
		return errors.Wrap(err, "Failed to open file, check filepath")
	}

	_, numDimensions, err := iD.parseMagic()
	iD.DataDimensions, err = iD.parseDimensions(numDimensions)
	numDataBytes, err := iD.getNumBytes()
	iD.Data, err = iD.parseDataBytes(numDataBytes)

	// Clean up file pointer
	iD.file = nil

	return nil
}

func (iD *IDXData) parseMagic() (datatype, numDimensions int, err error) {
	magic := make([]byte, 4)
	iD.file.Read(magic)

	if len(magic) != 4 {
		err = errors.New("Invalid magic bytes")
	}
	if magic[0] != 0 || magic[1] != 0 {
		err = errors.New("Invalid magic bytes")
	}

	datatype = int(magic[2])
	numDimensions = int(magic[3])
	return
}

func (iD *IDXData) parseDimensions(numDimensions int) (dataDimensions []int, err error) {
	dataDimensions = make([]int, numDimensions)
	dimension := make([]byte, 4)
	for i:=0; i<numDimensions; i++ {
		iD.file.Read(dimension)
		dataDimensions[i] = int(binary.BigEndian.Uint32(dimension))
	}
	return
}

func (iD *IDXData) parseDataBytes(numBytes int) (data []int, err error) {
	data = make([]int, numBytes)
	values := make([]byte, numBytes)
	iD.file.Read(values)
	for i, value := range values {
		data[i] = int(value)
	}
	return
}

func (iD *IDXData) getNumBytes() (numBytes int, err error) {
	numBytes = 1
	for _, dim := range iD.DataDimensions {
		numBytes = numBytes * dim
	}
	return
}