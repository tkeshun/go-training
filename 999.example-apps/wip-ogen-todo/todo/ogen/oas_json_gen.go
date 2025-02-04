// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"math/bits"
	"strconv"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode encodes time.Time as json.
func (o OptDateTime) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptDateTime) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptDateTime to nil")
	}
	o.Set = true
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptDateTime) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.EncodeDateTime)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptDateTime) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.DecodeDateTime)
}

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt to nil")
	}
	o.Set = true
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes TodosFilterGetOKApplicationJSON as json.
func (s TodosFilterGetOKApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := []TodosFilterGetOKItem(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes TodosFilterGetOKApplicationJSON from json.
func (s *TodosFilterGetOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TodosFilterGetOKApplicationJSON to nil")
	}
	var unwrapped []TodosFilterGetOKItem
	if err := func() error {
		unwrapped = make([]TodosFilterGetOKItem, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem TodosFilterGetOKItem
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = TodosFilterGetOKApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s TodosFilterGetOKApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TodosFilterGetOKApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TodosFilterGetOKItem) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TodosFilterGetOKItem) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		if s.Status.Set {
			e.FieldStart("status")
			s.Status.Encode(e)
		}
	}
	{
		if s.Priority.Set {
			e.FieldStart("priority")
			s.Priority.Encode(e)
		}
	}
}

var jsonFieldsNameOfTodosFilterGetOKItem = [5]string{
	0: "id",
	1: "title",
	2: "description",
	3: "status",
	4: "priority",
}

// Decode decodes TodosFilterGetOKItem from json.
func (s *TodosFilterGetOKItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TodosFilterGetOKItem to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "status":
			if err := func() error {
				s.Status.Reset()
				if err := s.Status.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "priority":
			if err := func() error {
				s.Priority.Reset()
				if err := s.Priority.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"priority\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TodosFilterGetOKItem")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TodosFilterGetOKItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TodosFilterGetOKItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TodosGetOKItem) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TodosGetOKItem) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		if s.Status.Set {
			e.FieldStart("status")
			s.Status.Encode(e)
		}
	}
	{
		if s.Priority.Set {
			e.FieldStart("priority")
			s.Priority.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.UpdatedAt.Set {
			e.FieldStart("updated_at")
			s.UpdatedAt.Encode(e, json.EncodeDateTime)
		}
	}
}

var jsonFieldsNameOfTodosGetOKItem = [7]string{
	0: "id",
	1: "title",
	2: "description",
	3: "status",
	4: "priority",
	5: "created_at",
	6: "updated_at",
}

// Decode decodes TodosGetOKItem from json.
func (s *TodosGetOKItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TodosGetOKItem to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "status":
			if err := func() error {
				s.Status.Reset()
				if err := s.Status.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "priority":
			if err := func() error {
				s.Priority.Reset()
				if err := s.Priority.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"priority\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "updated_at":
			if err := func() error {
				s.UpdatedAt.Reset()
				if err := s.UpdatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TodosGetOKItem")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TodosGetOKItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TodosGetOKItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TodosIDGetOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TodosIDGetOK) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		if s.Status.Set {
			e.FieldStart("status")
			s.Status.Encode(e)
		}
	}
	{
		if s.Priority.Set {
			e.FieldStart("priority")
			s.Priority.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.UpdatedAt.Set {
			e.FieldStart("updated_at")
			s.UpdatedAt.Encode(e, json.EncodeDateTime)
		}
	}
}

var jsonFieldsNameOfTodosIDGetOK = [7]string{
	0: "id",
	1: "title",
	2: "description",
	3: "status",
	4: "priority",
	5: "created_at",
	6: "updated_at",
}

// Decode decodes TodosIDGetOK from json.
func (s *TodosIDGetOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TodosIDGetOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "status":
			if err := func() error {
				s.Status.Reset()
				if err := s.Status.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "priority":
			if err := func() error {
				s.Priority.Reset()
				if err := s.Priority.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"priority\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "updated_at":
			if err := func() error {
				s.UpdatedAt.Reset()
				if err := s.UpdatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TodosIDGetOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TodosIDGetOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TodosIDGetOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TodosIDPatchReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TodosIDPatchReq) encodeFields(e *jx.Encoder) {
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
}

var jsonFieldsNameOfTodosIDPatchReq = [2]string{
	0: "title",
	1: "description",
}

// Decode decodes TodosIDPatchReq from json.
func (s *TodosIDPatchReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TodosIDPatchReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		default:
			return errors.Errorf("unexpected field %q", k)
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TodosIDPatchReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TodosIDPatchReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TodosIDPatchReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TodosIDPriorityPatchReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TodosIDPriorityPatchReq) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("priority")
		s.Priority.Encode(e)
	}
}

var jsonFieldsNameOfTodosIDPriorityPatchReq = [1]string{
	0: "priority",
}

// Decode decodes TodosIDPriorityPatchReq from json.
func (s *TodosIDPriorityPatchReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TodosIDPriorityPatchReq to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "priority":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Priority.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"priority\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TodosIDPriorityPatchReq")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTodosIDPriorityPatchReq) {
					name = jsonFieldsNameOfTodosIDPriorityPatchReq[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TodosIDPriorityPatchReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TodosIDPriorityPatchReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes TodosIDPriorityPatchReqPriority as json.
func (s TodosIDPriorityPatchReqPriority) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes TodosIDPriorityPatchReqPriority from json.
func (s *TodosIDPriorityPatchReqPriority) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TodosIDPriorityPatchReqPriority to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch TodosIDPriorityPatchReqPriority(v) {
	case TodosIDPriorityPatchReqPriorityLOW:
		*s = TodosIDPriorityPatchReqPriorityLOW
	case TodosIDPriorityPatchReqPriorityMEDIUM:
		*s = TodosIDPriorityPatchReqPriorityMEDIUM
	case TodosIDPriorityPatchReqPriorityHIGH:
		*s = TodosIDPriorityPatchReqPriorityHIGH
	default:
		*s = TodosIDPriorityPatchReqPriority(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s TodosIDPriorityPatchReqPriority) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TodosIDPriorityPatchReqPriority) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TodosIDTagsPostReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TodosIDTagsPostReq) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("tagId")
		e.Int(s.TagId)
	}
}

var jsonFieldsNameOfTodosIDTagsPostReq = [1]string{
	0: "tagId",
}

// Decode decodes TodosIDTagsPostReq from json.
func (s *TodosIDTagsPostReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TodosIDTagsPostReq to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "tagId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.TagId = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"tagId\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TodosIDTagsPostReq")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTodosIDTagsPostReq) {
					name = jsonFieldsNameOfTodosIDTagsPostReq[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TodosIDTagsPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TodosIDTagsPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TodosPostReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TodosPostReq) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("title")
		e.Str(s.Title)
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
}

var jsonFieldsNameOfTodosPostReq = [2]string{
	0: "title",
	1: "description",
}

// Decode decodes TodosPostReq from json.
func (s *TodosPostReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TodosPostReq to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TodosPostReq")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTodosPostReq) {
					name = jsonFieldsNameOfTodosPostReq[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TodosPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TodosPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UsersIDGetOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UsersIDGetOK) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Name.Set {
			e.FieldStart("name")
			s.Name.Encode(e)
		}
	}
	{
		if s.Email.Set {
			e.FieldStart("email")
			s.Email.Encode(e)
		}
	}
	{
		if s.Role.Set {
			e.FieldStart("role")
			s.Role.Encode(e)
		}
	}
}

var jsonFieldsNameOfUsersIDGetOK = [4]string{
	0: "id",
	1: "name",
	2: "email",
	3: "role",
}

// Decode decodes UsersIDGetOK from json.
func (s *UsersIDGetOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UsersIDGetOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			if err := func() error {
				s.Name.Reset()
				if err := s.Name.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "email":
			if err := func() error {
				s.Email.Reset()
				if err := s.Email.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"email\"")
			}
		case "role":
			if err := func() error {
				s.Role.Reset()
				if err := s.Role.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"role\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UsersIDGetOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UsersIDGetOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UsersIDGetOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UsersPostReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UsersPostReq) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		e.FieldStart("email")
		e.Str(s.Email)
	}
	{
		e.FieldStart("password")
		e.Str(s.Password)
	}
}

var jsonFieldsNameOfUsersPostReq = [3]string{
	0: "name",
	1: "email",
	2: "password",
}

// Decode decodes UsersPostReq from json.
func (s *UsersPostReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UsersPostReq to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "name":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "email":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Email = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"email\"")
			}
		case "password":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Password = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"password\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UsersPostReq")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUsersPostReq) {
					name = jsonFieldsNameOfUsersPostReq[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UsersPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UsersPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
