// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package cart

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	api "github.com/doutokk/doutok/rpc_gen/kitex_gen/cwgo/http/api"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *EditCartReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_EditCartReq[number], err)
}

func (x *EditCartReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *EditCartReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v CartItem
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Items = append(x.Items, &v)
	return offset, nil
}

func (x *EditCartResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *CartItem) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CartItem[number], err)
}

func (x *CartItem) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ProductId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CartItem) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Quantity, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *AddItemReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AddItemReq[number], err)
}

func (x *AddItemReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *AddItemReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v CartItem
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Item = &v
	return offset, nil
}

func (x *AddItemResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *EmptyCartReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_EmptyCartReq[number], err)
}

func (x *EmptyCartReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *GetCartReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetCartReq[number], err)
}

func (x *GetCartReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *GetCartResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetCartResp[number], err)
}

func (x *GetCartResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Cart
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Cart = &v
	return offset, nil
}

func (x *Cart) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Cart[number], err)
}

func (x *Cart) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *Cart) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v CartItem
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Items = append(x.Items, &v)
	return offset, nil
}

func (x *EmptyCartResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *EditCartReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *EditCartReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *EditCartReq) fastWriteField2(buf []byte) (offset int) {
	if x.Items == nil {
		return offset
	}
	for i := range x.GetItems() {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.GetItems()[i])
	}
	return offset
}

func (x *EditCartResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *CartItem) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CartItem) fastWriteField1(buf []byte) (offset int) {
	if x.ProductId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetProductId())
	return offset
}

func (x *CartItem) fastWriteField2(buf []byte) (offset int) {
	if x.Quantity == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetQuantity())
	return offset
}

func (x *AddItemReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *AddItemReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *AddItemReq) fastWriteField2(buf []byte) (offset int) {
	if x.Item == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetItem())
	return offset
}

func (x *AddItemResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *EmptyCartReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *EmptyCartReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *GetCartReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetCartReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *GetCartResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetCartResp) fastWriteField1(buf []byte) (offset int) {
	if x.Cart == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetCart())
	return offset
}

func (x *Cart) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *Cart) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *Cart) fastWriteField2(buf []byte) (offset int) {
	if x.Items == nil {
		return offset
	}
	for i := range x.GetItems() {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.GetItems()[i])
	}
	return offset
}

func (x *EmptyCartResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *EditCartReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *EditCartReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetUserId())
	return n
}

func (x *EditCartReq) sizeField2() (n int) {
	if x.Items == nil {
		return n
	}
	for i := range x.GetItems() {
		n += fastpb.SizeMessage(2, x.GetItems()[i])
	}
	return n
}

func (x *EditCartResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *CartItem) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CartItem) sizeField1() (n int) {
	if x.ProductId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetProductId())
	return n
}

func (x *CartItem) sizeField2() (n int) {
	if x.Quantity == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetQuantity())
	return n
}

func (x *AddItemReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *AddItemReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetUserId())
	return n
}

func (x *AddItemReq) sizeField2() (n int) {
	if x.Item == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetItem())
	return n
}

func (x *AddItemResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *EmptyCartReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *EmptyCartReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetUserId())
	return n
}

func (x *GetCartReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetCartReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetUserId())
	return n
}

func (x *GetCartResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetCartResp) sizeField1() (n int) {
	if x.Cart == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetCart())
	return n
}

func (x *Cart) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *Cart) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetUserId())
	return n
}

func (x *Cart) sizeField2() (n int) {
	if x.Items == nil {
		return n
	}
	for i := range x.GetItems() {
		n += fastpb.SizeMessage(2, x.GetItems()[i])
	}
	return n
}

func (x *EmptyCartResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

var fieldIDToName_EditCartReq = map[int32]string{
	1: "UserId",
	2: "Items",
}

var fieldIDToName_EditCartResp = map[int32]string{}

var fieldIDToName_CartItem = map[int32]string{
	1: "ProductId",
	2: "Quantity",
}

var fieldIDToName_AddItemReq = map[int32]string{
	1: "UserId",
	2: "Item",
}

var fieldIDToName_AddItemResp = map[int32]string{}

var fieldIDToName_EmptyCartReq = map[int32]string{
	1: "UserId",
}

var fieldIDToName_GetCartReq = map[int32]string{
	1: "UserId",
}

var fieldIDToName_GetCartResp = map[int32]string{
	1: "Cart",
}

var fieldIDToName_Cart = map[int32]string{
	1: "UserId",
	2: "Items",
}

var fieldIDToName_EmptyCartResp = map[int32]string{}

var _ = api.File_api_proto
