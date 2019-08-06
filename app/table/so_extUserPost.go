package table

import (
	"errors"
	fmt "fmt"
	"reflect"

	"github.com/coschain/contentos-go/common/encoding/kope"
	"github.com/coschain/contentos-go/iservices"
	prototype "github.com/coschain/contentos-go/prototype"
	proto "github.com/golang/protobuf/proto"
)

////////////// SECTION Prefix Mark ///////////////
var (
	ExtUserPostPostCreatedOrderTable uint32 = 555226009
	ExtUserPostPostIdUniTable        uint32 = 2411654352

	ExtUserPostPostIdRow uint32 = 3578023745
)

////////////// SECTION Wrap Define ///////////////
type SoExtUserPostWrap struct {
	dba       iservices.IDatabaseRW
	mainKey   *uint64
	mKeyFlag  int    //the flag of the main key exist state in db, -1:has not judged; 0:not exist; 1:already exist
	mKeyBuf   []byte //the buffer after the main key is encoded with prefix
	mBuf      []byte //the value after the main key is encoded
	mdFuncMap map[string]interface{}
}

func NewSoExtUserPostWrap(dba iservices.IDatabaseRW, key *uint64) *SoExtUserPostWrap {
	if dba == nil || key == nil {
		return nil
	}
	result := &SoExtUserPostWrap{dba, key, -1, nil, nil, nil}
	return result
}

func (s *SoExtUserPostWrap) CheckExist() bool {
	if s.dba == nil {
		return false
	}
	if s.mKeyFlag != -1 {
		//if you have already obtained the existence status of the primary key, use it directly
		if s.mKeyFlag == 0 {
			return false
		}
		return true
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}

	res, err := s.dba.Has(keyBuf)
	if err != nil {
		return false
	}
	if res == false {
		s.mKeyFlag = 0
	} else {
		s.mKeyFlag = 1
	}
	return res
}

func (s *SoExtUserPostWrap) Create(f func(tInfo *SoExtUserPost)) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if s.mainKey == nil {
		return errors.New("the main key is nil")
	}
	val := &SoExtUserPost{}
	f(val)
	if s.CheckExist() {
		return errors.New("the main key is already exist")
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return err

	}

	buf, err := proto.Marshal(val)
	if err != nil {
		return err
	}
	err = s.dba.Put(keyBuf, buf)
	if err != nil {
		return err
	}

	// update srt list keys
	if err = s.insertAllSortKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.dba.Delete(keyBuf)
		return err
	}

	//update unique list
	if sucNames, err := s.insertAllUniKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.delUniKeysWithNames(sucNames, val)
		s.dba.Delete(keyBuf)
		return err
	}

	s.mKeyFlag = 1
	return nil
}

func (s *SoExtUserPostWrap) getMainKeyBuf() ([]byte, error) {
	if s.mainKey == nil {
		return nil, errors.New("the main key is nil")
	}
	if s.mBuf == nil {
		var err error = nil
		s.mBuf, err = kope.Encode(s.mainKey)
		if err != nil {
			return nil, err
		}
	}
	return s.mBuf, nil
}

func (s *SoExtUserPostWrap) Modify(f func(tInfo *SoExtUserPost)) error {
	if !s.CheckExist() {
		return errors.New("the SoExtUserPost table does not exist. Please create a table first")
	}
	oriTable := s.getExtUserPost()
	if oriTable == nil {
		return errors.New("fail to get origin table SoExtUserPost")
	}
	curTable := *oriTable
	f(&curTable)

	//the main key is not support modify
	if !reflect.DeepEqual(curTable.PostId, oriTable.PostId) {
		return errors.New("primary key does not support modification")
	}

	fieldSli, err := s.getModifiedFields(oriTable, &curTable)
	if err != nil {
		return err
	}

	if fieldSli == nil || len(fieldSli) < 1 {
		return nil
	}

	//check whether modify sort and unique field to nil
	err = s.checkSortAndUniFieldValidity(&curTable, fieldSli)
	if err != nil {
		return err
	}

	//check unique
	err = s.handleFieldMd(FieldMdHandleTypeCheck, &curTable, fieldSli)
	if err != nil {
		return err
	}

	//delete sort and unique key
	err = s.handleFieldMd(FieldMdHandleTypeDel, oriTable, fieldSli)
	if err != nil {
		return err
	}

	//update table
	err = s.updateExtUserPost(&curTable)
	if err != nil {
		return err
	}

	//insert sort and unique key
	err = s.handleFieldMd(FieldMdHandleTypeInsert, &curTable, fieldSli)
	if err != nil {
		return err
	}

	return nil

}

func (s *SoExtUserPostWrap) MdPostCreatedOrder(p *prototype.UserPostCreateOrder) bool {
	err := s.Modify(func(r *SoExtUserPost) {
		r.PostCreatedOrder = p
	})
	return err == nil
}

func (s *SoExtUserPostWrap) checkSortAndUniFieldValidity(curTable *SoExtUserPost, fieldSli []string) error {
	if curTable != nil && fieldSli != nil && len(fieldSli) > 0 {
		for _, fName := range fieldSli {
			if len(fName) > 0 {

				if fName == "PostCreatedOrder" && curTable.PostCreatedOrder == nil {
					return errors.New("sort field PostCreatedOrder can't be modified to nil")
				}

			}
		}
	}
	return nil
}

//Get all the modified fields in the table
func (s *SoExtUserPostWrap) getModifiedFields(oriTable *SoExtUserPost, curTable *SoExtUserPost) ([]string, error) {
	if oriTable == nil {
		return nil, errors.New("table info is nil, can't get modified fields")
	}
	var list []string

	if !reflect.DeepEqual(oriTable.PostCreatedOrder, curTable.PostCreatedOrder) {
		list = append(list, "PostCreatedOrder")
	}

	return list, nil
}

func (s *SoExtUserPostWrap) handleFieldMd(t FieldMdHandleType, so *SoExtUserPost, fSli []string) error {
	if so == nil {
		return errors.New("fail to modify empty table")
	}

	//there is no field need to modify
	if fSli == nil || len(fSli) < 1 {
		return nil
	}

	errStr := ""
	for _, fName := range fSli {

		if fName == "PostCreatedOrder" {
			res := true
			if t == FieldMdHandleTypeCheck {
				res = s.mdFieldPostCreatedOrder(so.PostCreatedOrder, true, false, false, so)
				errStr = fmt.Sprintf("fail to modify exist value of %v", fName)
			} else if t == FieldMdHandleTypeDel {
				res = s.mdFieldPostCreatedOrder(so.PostCreatedOrder, false, true, false, so)
				errStr = fmt.Sprintf("fail to delete  sort or unique field  %v", fName)
			} else if t == FieldMdHandleTypeInsert {
				res = s.mdFieldPostCreatedOrder(so.PostCreatedOrder, false, false, true, so)
				errStr = fmt.Sprintf("fail to insert  sort or unique field  %v", fName)
			}
			if !res {
				return errors.New(errStr)
			}
		}

	}

	return nil
}

////////////// SECTION LKeys delete/insert ///////////////

func (s *SoExtUserPostWrap) delSortKeyPostCreatedOrder(sa *SoExtUserPost) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListExtUserPostByPostCreatedOrder{}
	if sa == nil {
		val.PostCreatedOrder = s.GetPostCreatedOrder()
		val.PostId = *s.mainKey
	} else {
		val.PostCreatedOrder = sa.PostCreatedOrder
		val.PostId = sa.PostId
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoExtUserPostWrap) insertSortKeyPostCreatedOrder(sa *SoExtUserPost) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListExtUserPostByPostCreatedOrder{}
	val.PostId = sa.PostId
	val.PostCreatedOrder = sa.PostCreatedOrder
	buf, err := proto.Marshal(&val)
	if err != nil {
		return false
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Put(subBuf, buf)
	return ordErr == nil
}

func (s *SoExtUserPostWrap) delAllSortKeys(br bool, val *SoExtUserPost) bool {
	if s.dba == nil {
		return false
	}
	res := true

	if !s.delSortKeyPostCreatedOrder(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoExtUserPostWrap) insertAllSortKeys(val *SoExtUserPost) error {
	if s.dba == nil {
		return errors.New("insert sort Field fail,the db is nil ")
	}
	if val == nil {
		return errors.New("insert sort Field fail,get the SoExtUserPost fail ")
	}

	if !s.insertSortKeyPostCreatedOrder(val) {
		return errors.New("insert sort Field PostCreatedOrder fail while insert table ")
	}

	return nil
}

////////////// SECTION LKeys delete/insert //////////////

func (s *SoExtUserPostWrap) RemoveExtUserPost() bool {
	if s.dba == nil {
		return false
	}
	//delete sort list key
	if res := s.delAllSortKeys(true, nil); !res {
		return false
	}

	//delete unique list
	if res := s.delAllUniKeys(true, nil); !res {
		return false
	}

	//delete table
	key, err := s.encodeMainKey()
	if err != nil {
		return false
	}
	err = s.dba.Delete(key)
	if err == nil {
		s.mKeyBuf = nil
		s.mKeyFlag = -1
		return true
	} else {
		return false
	}
}

////////////// SECTION Members Get/Modify ///////////////

func (s *SoExtUserPostWrap) GetPostCreatedOrder() *prototype.UserPostCreateOrder {
	res := true
	msg := &SoExtUserPost{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMainKey()
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.PostCreatedOrder
			}
		}
	}
	if !res {
		return nil

	}
	return msg.PostCreatedOrder
}

func (s *SoExtUserPostWrap) mdFieldPostCreatedOrder(p *prototype.UserPostCreateOrder, isCheck bool, isDel bool, isInsert bool,
	so *SoExtUserPost) bool {
	if s.dba == nil {
		return false
	}

	if isCheck {
		res := s.checkPostCreatedOrderIsMetMdCondition(p)
		if !res {
			return false
		}
	}

	if isDel {
		res := s.delFieldPostCreatedOrder(so)
		if !res {
			return false
		}
	}

	if isInsert {
		res := s.insertFieldPostCreatedOrder(so)
		if !res {
			return false
		}
	}
	return true
}

func (s *SoExtUserPostWrap) delFieldPostCreatedOrder(so *SoExtUserPost) bool {
	if s.dba == nil {
		return false
	}

	if !s.delSortKeyPostCreatedOrder(so) {
		return false
	}

	return true
}

func (s *SoExtUserPostWrap) insertFieldPostCreatedOrder(so *SoExtUserPost) bool {
	if s.dba == nil {
		return false
	}

	if !s.insertSortKeyPostCreatedOrder(so) {
		return false
	}

	return true
}

func (s *SoExtUserPostWrap) checkPostCreatedOrderIsMetMdCondition(p *prototype.UserPostCreateOrder) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoExtUserPostWrap) GetPostId() uint64 {
	res := true
	msg := &SoExtUserPost{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMainKey()
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.PostId
			}
		}
	}
	if !res {
		var tmpValue uint64
		return tmpValue
	}
	return msg.PostId
}

////////////// SECTION List Keys ///////////////
type SExtUserPostPostCreatedOrderWrap struct {
	Dba iservices.IDatabaseRW
}

func NewExtUserPostPostCreatedOrderWrap(db iservices.IDatabaseRW) *SExtUserPostPostCreatedOrderWrap {
	if db == nil {
		return nil
	}
	wrap := SExtUserPostPostCreatedOrderWrap{Dba: db}
	return &wrap
}

func (s *SExtUserPostPostCreatedOrderWrap) GetMainVal(val []byte) *uint64 {
	res := &SoListExtUserPostByPostCreatedOrder{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}

	return &res.PostId

}

func (s *SExtUserPostPostCreatedOrderWrap) GetSubVal(val []byte) *prototype.UserPostCreateOrder {
	res := &SoListExtUserPostByPostCreatedOrder{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.PostCreatedOrder

}

func (m *SoListExtUserPostByPostCreatedOrder) OpeEncode() ([]byte, error) {
	pre := ExtUserPostPostCreatedOrderTable
	sub := m.PostCreatedOrder
	if sub == nil {
		return nil, errors.New("the pro PostCreatedOrder is nil")
	}
	sub1 := m.PostId

	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query srt by order
//
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SExtUserPostPostCreatedOrderWrap) ForEachByOrder(start *prototype.UserPostCreateOrder, end *prototype.UserPostCreateOrder, lastMainKey *uint64,
	lastSubVal *prototype.UserPostCreateOrder, f func(mVal *uint64, sVal *prototype.UserPostCreateOrder, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := ExtUserPostPostCreatedOrderTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey, kope.MinimalKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey, kope.MinimalKey)
		}
		skeyList = append(skeyList, kope.MinimalKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	} else {
		eKeyList = append(eKeyList, kope.MaximumKey)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(sBuf, eBuf, false, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

//Query srt by reverse order
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SExtUserPostPostCreatedOrderWrap) ForEachByRevOrder(start *prototype.UserPostCreateOrder, end *prototype.UserPostCreateOrder, lastMainKey *uint64,
	lastSubVal *prototype.UserPostCreateOrder, f func(mVal *uint64, sVal *prototype.UserPostCreateOrder, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := ExtUserPostPostCreatedOrderTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey)
		}
		skeyList = append(skeyList, kope.MaximumKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(eBuf, sBuf, true, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

/////////////// SECTION Private function ////////////////

func (s *SoExtUserPostWrap) update(sa *SoExtUserPost) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	buf, err := proto.Marshal(sa)
	if err != nil {
		return false
	}

	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}

	return s.dba.Put(keyBuf, buf) == nil
}

func (s *SoExtUserPostWrap) getExtUserPost() *SoExtUserPost {
	if s.dba == nil {
		return nil
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return nil
	}
	resBuf, err := s.dba.Get(keyBuf)

	if err != nil {
		return nil
	}

	res := &SoExtUserPost{}
	if proto.Unmarshal(resBuf, res) != nil {
		return nil
	}
	return res
}

func (s *SoExtUserPostWrap) updateExtUserPost(so *SoExtUserPost) error {
	if s.dba == nil {
		return errors.New("update fail:the db is nil")
	}

	if so == nil {
		return errors.New("update fail: the SoExtUserPost is nil")
	}

	key, err := s.encodeMainKey()
	if err != nil {
		return nil
	}

	buf, err := proto.Marshal(so)
	if err != nil {
		return err
	}

	err = s.dba.Put(key, buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *SoExtUserPostWrap) encodeMainKey() ([]byte, error) {
	if s.mKeyBuf != nil {
		return s.mKeyBuf, nil
	}
	pre := ExtUserPostPostIdRow
	sub := s.mainKey
	if sub == nil {
		return nil, errors.New("the mainKey is nil")
	}
	preBuf, err := kope.Encode(pre)
	if err != nil {
		return nil, err
	}
	mBuf, err := s.getMainKeyBuf()
	if err != nil {
		return nil, err
	}
	list := make([][]byte, 2)
	list[0] = preBuf
	list[1] = mBuf
	s.mKeyBuf = kope.PackList(list)
	return s.mKeyBuf, nil
}

////////////// Unique Query delete/insert/query ///////////////

func (s *SoExtUserPostWrap) delAllUniKeys(br bool, val *SoExtUserPost) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delUniKeyPostId(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoExtUserPostWrap) delUniKeysWithNames(names map[string]string, val *SoExtUserPost) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if len(names["PostId"]) > 0 {
		if !s.delUniKeyPostId(val) {
			res = false
		}
	}

	return res
}

func (s *SoExtUserPostWrap) insertAllUniKeys(val *SoExtUserPost) (map[string]string, error) {
	if s.dba == nil {
		return nil, errors.New("insert uniuqe Field fail,the db is nil ")
	}
	if val == nil {
		return nil, errors.New("insert uniuqe Field fail,get the SoExtUserPost fail ")
	}
	sucFields := map[string]string{}
	if !s.insertUniKeyPostId(val) {
		return sucFields, errors.New("insert unique Field PostId fail while insert table ")
	}
	sucFields["PostId"] = "PostId"

	return sucFields, nil
}

func (s *SoExtUserPostWrap) delUniKeyPostId(sa *SoExtUserPost) bool {
	if s.dba == nil {
		return false
	}
	pre := ExtUserPostPostIdUniTable
	kList := []interface{}{pre}
	if sa != nil {

		sub := sa.PostId
		kList = append(kList, sub)
	} else {
		sub := s.GetPostId()

		kList = append(kList, sub)

	}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	return s.dba.Delete(kBuf) == nil
}

func (s *SoExtUserPostWrap) insertUniKeyPostId(sa *SoExtUserPost) bool {
	if s.dba == nil || sa == nil {
		return false
	}

	pre := ExtUserPostPostIdUniTable
	sub := sa.PostId
	kList := []interface{}{pre, sub}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	res, err := s.dba.Has(kBuf)
	if err == nil && res == true {
		//the unique key is already exist
		return false
	}
	val := SoUniqueExtUserPostByPostId{}
	val.PostId = sa.PostId

	buf, err := proto.Marshal(&val)

	if err != nil {
		return false
	}

	return s.dba.Put(kBuf, buf) == nil

}

type UniExtUserPostPostIdWrap struct {
	Dba iservices.IDatabaseRW
}

func NewUniExtUserPostPostIdWrap(db iservices.IDatabaseRW) *UniExtUserPostPostIdWrap {
	if db == nil {
		return nil
	}
	wrap := UniExtUserPostPostIdWrap{Dba: db}
	return &wrap
}

func (s *UniExtUserPostPostIdWrap) UniQueryPostId(start *uint64) *SoExtUserPostWrap {
	if start == nil || s.Dba == nil {
		return nil
	}
	pre := ExtUserPostPostIdUniTable
	kList := []interface{}{pre, start}
	bufStartkey, err := kope.EncodeSlice(kList)
	val, err := s.Dba.Get(bufStartkey)
	if err == nil {
		res := &SoUniqueExtUserPostByPostId{}
		rErr := proto.Unmarshal(val, res)
		if rErr == nil {
			wrap := NewSoExtUserPostWrap(s.Dba, &res.PostId)
			return wrap
		}
	}
	return nil
}
