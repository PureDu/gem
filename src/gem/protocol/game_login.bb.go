// Generated by bbc; DO NOT EDIT
package protocol

import (
	"gem/encoding"
	"io"
)

type ClientSecureLoginBlock struct {
	Magic           encoding.Int8
	ClientISAACSeed [2]encoding.Int32
	ClientUID       encoding.Int32
	Username        encoding.JString
	Password        encoding.JString
}

func (struc *ClientSecureLoginBlock) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.Magic.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	for i := 0; i < 2; i++ {
		err = struc.ClientISAACSeed[i].Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	err = struc.ClientUID.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Username.Encode(buf, 0)
	if err != nil {
		return err
	}

	err = struc.Password.Encode(buf, 0)
	if err != nil {
		return err
	}

	return err
}

func (struc *ClientSecureLoginBlock) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.Magic.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	for i := 0; i < 2; i++ {
		err = struc.ClientISAACSeed[i].Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	err = struc.ClientUID.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Username.Decode(buf, 0)
	if err != nil {
		return err
	}

	err = struc.Password.Decode(buf, 0)
	if err != nil {
		return err
	}

	return err
}

type ServerLoginResponse struct {
	Response encoding.Int8
	Rights   encoding.Int8
	Flagged  encoding.Int8
}

func (struc *ServerLoginResponse) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.Response.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Rights.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Flagged.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

func (struc *ServerLoginResponse) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.Response.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Rights.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Flagged.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

type ClientLoginBlock struct {
	LoginType   encoding.Int8
	LoginLen    encoding.Int8
	Magic       encoding.Int8
	Revision    encoding.Int16
	MemType     encoding.Int8
	ArchiveCRCs [9]encoding.Int32
}

func (struc *ClientLoginBlock) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.LoginType.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.LoginLen.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Magic.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Revision.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.MemType.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	for i := 0; i < 9; i++ {
		err = struc.ArchiveCRCs[i].Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	return err
}

func (struc *ClientLoginBlock) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.LoginType.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.LoginLen.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Magic.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Revision.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.MemType.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	for i := 0; i < 9; i++ {
		err = struc.ArchiveCRCs[i].Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	return err
}
