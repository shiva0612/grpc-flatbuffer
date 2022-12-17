// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package models

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type UserRequestT struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	UserId string `json:"user_id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Dob string `json:"dob"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
}

func (t *UserRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	firstNameOffset := builder.CreateString(t.FirstName)
	lastNameOffset := builder.CreateString(t.LastName)
	middleNameOffset := builder.CreateString(t.MiddleName)
	userIdOffset := builder.CreateString(t.UserId)
	emailOffset := builder.CreateString(t.Email)
	phoneOffset := builder.CreateString(t.Phone)
	dobOffset := builder.CreateString(t.Dob)
	cityOffset := builder.CreateString(t.City)
	stateOffset := builder.CreateString(t.State)
	countryOffset := builder.CreateString(t.Country)
	UserRequestStart(builder)
	UserRequestAddFirstName(builder, firstNameOffset)
	UserRequestAddLastName(builder, lastNameOffset)
	UserRequestAddMiddleName(builder, middleNameOffset)
	UserRequestAddUserId(builder, userIdOffset)
	UserRequestAddEmail(builder, emailOffset)
	UserRequestAddPhone(builder, phoneOffset)
	UserRequestAddDob(builder, dobOffset)
	UserRequestAddCity(builder, cityOffset)
	UserRequestAddState(builder, stateOffset)
	UserRequestAddCountry(builder, countryOffset)
	return UserRequestEnd(builder)
}

func (rcv *UserRequest) UnPackTo(t *UserRequestT) {
	t.FirstName = string(rcv.FirstName())
	t.LastName = string(rcv.LastName())
	t.MiddleName = string(rcv.MiddleName())
	t.UserId = string(rcv.UserId())
	t.Email = string(rcv.Email())
	t.Phone = string(rcv.Phone())
	t.Dob = string(rcv.Dob())
	t.City = string(rcv.City())
	t.State = string(rcv.State())
	t.Country = string(rcv.Country())
}

func (rcv *UserRequest) UnPack() *UserRequestT {
	if rcv == nil { return nil }
	t := &UserRequestT{}
	rcv.UnPackTo(t)
	return t
}

type UserRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsUserRequest(buf []byte, offset flatbuffers.UOffsetT) *UserRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UserRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUserRequest(buf []byte, offset flatbuffers.UOffsetT) *UserRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &UserRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *UserRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UserRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *UserRequest) FirstName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *UserRequest) LastName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *UserRequest) MiddleName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *UserRequest) UserId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *UserRequest) Email() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *UserRequest) Phone() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *UserRequest) Dob() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *UserRequest) City() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *UserRequest) State() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *UserRequest) Country() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func UserRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}
func UserRequestAddFirstName(builder *flatbuffers.Builder, firstName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(firstName), 0)
}
func UserRequestAddLastName(builder *flatbuffers.Builder, lastName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(lastName), 0)
}
func UserRequestAddMiddleName(builder *flatbuffers.Builder, middleName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(middleName), 0)
}
func UserRequestAddUserId(builder *flatbuffers.Builder, userId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(userId), 0)
}
func UserRequestAddEmail(builder *flatbuffers.Builder, email flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(email), 0)
}
func UserRequestAddPhone(builder *flatbuffers.Builder, phone flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(phone), 0)
}
func UserRequestAddDob(builder *flatbuffers.Builder, dob flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(dob), 0)
}
func UserRequestAddCity(builder *flatbuffers.Builder, city flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(city), 0)
}
func UserRequestAddState(builder *flatbuffers.Builder, state flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(state), 0)
}
func UserRequestAddCountry(builder *flatbuffers.Builder, country flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(country), 0)
}
func UserRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
