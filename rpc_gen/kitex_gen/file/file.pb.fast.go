// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package file

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	annotations "github.com/doutokk/doutok/rpc_gen/kitex_gen/genproto/googleapis/api/annotations"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *UploadFileResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UploadFileResp[number], err)
}

func (x *UploadFileResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Key, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UploadFileResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Host, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UploadFileResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Policy, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UploadFileResp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SecurityToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UploadFileResp) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Signature, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UploadFileResp) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.XOssCredential, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UploadFileResp) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.XOssDate, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UploadFileResp) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.XOssSignatureVersion, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UploadFileReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UploadFileReq[number], err)
}

func (x *UploadFileReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *UploadFileReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FileName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FrontendUploadFileReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FrontendUploadFileReq[number], err)
}

func (x *FrontendUploadFileReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FileName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FrontendUploadFileResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FrontendUploadFileResp[number], err)
}

func (x *FrontendUploadFileResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Key, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FrontendUploadFileResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Host, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FrontendUploadFileResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Policy, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FrontendUploadFileResp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SecurityToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FrontendUploadFileResp) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Signature, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FrontendUploadFileResp) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.XOssCredential, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FrontendUploadFileResp) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.XOssDate, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FrontendUploadFileResp) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.XOssSignatureVersion, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UploadFileResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	return offset
}

func (x *UploadFileResp) fastWriteField1(buf []byte) (offset int) {
	if x.Key == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetKey())
	return offset
}

func (x *UploadFileResp) fastWriteField2(buf []byte) (offset int) {
	if x.Host == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetHost())
	return offset
}

func (x *UploadFileResp) fastWriteField3(buf []byte) (offset int) {
	if x.Policy == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPolicy())
	return offset
}

func (x *UploadFileResp) fastWriteField4(buf []byte) (offset int) {
	if x.SecurityToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetSecurityToken())
	return offset
}

func (x *UploadFileResp) fastWriteField5(buf []byte) (offset int) {
	if x.Signature == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetSignature())
	return offset
}

func (x *UploadFileResp) fastWriteField6(buf []byte) (offset int) {
	if x.XOssCredential == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetXOssCredential())
	return offset
}

func (x *UploadFileResp) fastWriteField7(buf []byte) (offset int) {
	if x.XOssDate == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetXOssDate())
	return offset
}

func (x *UploadFileResp) fastWriteField8(buf []byte) (offset int) {
	if x.XOssSignatureVersion == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.GetXOssSignatureVersion())
	return offset
}

func (x *UploadFileReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *UploadFileReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *UploadFileReq) fastWriteField2(buf []byte) (offset int) {
	if x.FileName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFileName())
	return offset
}

func (x *FrontendUploadFileReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *FrontendUploadFileReq) fastWriteField2(buf []byte) (offset int) {
	if x.FileName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFileName())
	return offset
}

func (x *FrontendUploadFileResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	return offset
}

func (x *FrontendUploadFileResp) fastWriteField1(buf []byte) (offset int) {
	if x.Key == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetKey())
	return offset
}

func (x *FrontendUploadFileResp) fastWriteField2(buf []byte) (offset int) {
	if x.Host == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetHost())
	return offset
}

func (x *FrontendUploadFileResp) fastWriteField3(buf []byte) (offset int) {
	if x.Policy == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPolicy())
	return offset
}

func (x *FrontendUploadFileResp) fastWriteField4(buf []byte) (offset int) {
	if x.SecurityToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetSecurityToken())
	return offset
}

func (x *FrontendUploadFileResp) fastWriteField5(buf []byte) (offset int) {
	if x.Signature == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetSignature())
	return offset
}

func (x *FrontendUploadFileResp) fastWriteField6(buf []byte) (offset int) {
	if x.XOssCredential == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetXOssCredential())
	return offset
}

func (x *FrontendUploadFileResp) fastWriteField7(buf []byte) (offset int) {
	if x.XOssDate == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetXOssDate())
	return offset
}

func (x *FrontendUploadFileResp) fastWriteField8(buf []byte) (offset int) {
	if x.XOssSignatureVersion == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.GetXOssSignatureVersion())
	return offset
}

func (x *UploadFileResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	return n
}

func (x *UploadFileResp) sizeField1() (n int) {
	if x.Key == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetKey())
	return n
}

func (x *UploadFileResp) sizeField2() (n int) {
	if x.Host == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetHost())
	return n
}

func (x *UploadFileResp) sizeField3() (n int) {
	if x.Policy == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPolicy())
	return n
}

func (x *UploadFileResp) sizeField4() (n int) {
	if x.SecurityToken == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetSecurityToken())
	return n
}

func (x *UploadFileResp) sizeField5() (n int) {
	if x.Signature == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetSignature())
	return n
}

func (x *UploadFileResp) sizeField6() (n int) {
	if x.XOssCredential == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetXOssCredential())
	return n
}

func (x *UploadFileResp) sizeField7() (n int) {
	if x.XOssDate == "" {
		return n
	}
	n += fastpb.SizeString(7, x.GetXOssDate())
	return n
}

func (x *UploadFileResp) sizeField8() (n int) {
	if x.XOssSignatureVersion == "" {
		return n
	}
	n += fastpb.SizeString(8, x.GetXOssSignatureVersion())
	return n
}

func (x *UploadFileReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *UploadFileReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetUserId())
	return n
}

func (x *UploadFileReq) sizeField2() (n int) {
	if x.FileName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFileName())
	return n
}

func (x *FrontendUploadFileReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField2()
	return n
}

func (x *FrontendUploadFileReq) sizeField2() (n int) {
	if x.FileName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFileName())
	return n
}

func (x *FrontendUploadFileResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	return n
}

func (x *FrontendUploadFileResp) sizeField1() (n int) {
	if x.Key == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetKey())
	return n
}

func (x *FrontendUploadFileResp) sizeField2() (n int) {
	if x.Host == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetHost())
	return n
}

func (x *FrontendUploadFileResp) sizeField3() (n int) {
	if x.Policy == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPolicy())
	return n
}

func (x *FrontendUploadFileResp) sizeField4() (n int) {
	if x.SecurityToken == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetSecurityToken())
	return n
}

func (x *FrontendUploadFileResp) sizeField5() (n int) {
	if x.Signature == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetSignature())
	return n
}

func (x *FrontendUploadFileResp) sizeField6() (n int) {
	if x.XOssCredential == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetXOssCredential())
	return n
}

func (x *FrontendUploadFileResp) sizeField7() (n int) {
	if x.XOssDate == "" {
		return n
	}
	n += fastpb.SizeString(7, x.GetXOssDate())
	return n
}

func (x *FrontendUploadFileResp) sizeField8() (n int) {
	if x.XOssSignatureVersion == "" {
		return n
	}
	n += fastpb.SizeString(8, x.GetXOssSignatureVersion())
	return n
}

var fieldIDToName_UploadFileResp = map[int32]string{
	1: "Key",
	2: "Host",
	3: "Policy",
	4: "SecurityToken",
	5: "Signature",
	6: "XOssCredential",
	7: "XOssDate",
	8: "XOssSignatureVersion",
}

var fieldIDToName_UploadFileReq = map[int32]string{
	1: "UserId",
	2: "FileName",
}

var fieldIDToName_FrontendUploadFileReq = map[int32]string{
	2: "FileName",
}

var fieldIDToName_FrontendUploadFileResp = map[int32]string{
	1: "Key",
	2: "Host",
	3: "Policy",
	4: "SecurityToken",
	5: "Signature",
	6: "XOssCredential",
	7: "XOssDate",
	8: "XOssSignatureVersion",
}

var _ = annotations.File_google_api_annotations_proto
