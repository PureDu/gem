// Generated by bbc; DO NOT EDIT
package rt3

import (
	"bytes"
	"gem/encoding"
)

type FSIndex struct {
	Length     encoding.Int24
	StartBlock encoding.Int24
}

func (struc *FSIndex) Encode(buf *bytes.Buffer, flags interface{}) (err error) {
	err = struc.Length.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.StartBlock.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

func (struc *FSIndex) Decode(buf *bytes.Buffer, flags interface{}) (err error) {
	err = struc.Length.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.StartBlock.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

type FSBlock struct {
	FileID       encoding.Int16
	FilePosition encoding.Int16
	NextBlock    encoding.Int24
	Partition    encoding.Int8
	Data         [512]encoding.Int8
}

func (struc *FSBlock) Encode(buf *bytes.Buffer, flags interface{}) (err error) {
	err = struc.FileID.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.FilePosition.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.NextBlock.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Partition.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	for i := 0; i < 512; i++ {
		err = struc.Data[i].Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	return err
}

func (struc *FSBlock) Decode(buf *bytes.Buffer, flags interface{}) (err error) {
	err = struc.FileID.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.FilePosition.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.NextBlock.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Partition.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	for i := 0; i < 512; i++ {
		err = struc.Data[i].Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	return err
}
