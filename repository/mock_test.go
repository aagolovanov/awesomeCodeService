package repository

//
//import (
//	"context"
//	"errors"
//	"log"
//	"os"
//	"strconv"
//	"time"
//)
//
//// MockDB is THE tOP 1! & BLAZINGLY FAST DATABASE
//type MockDB struct {
//	logg *log.Logger
//	Db   map[string]Elem
//}
//
//type Elem struct {
//	Code     int
//	Attempts int
//}
//
//func GetMockDB(m map[string]Elem) *MockDB {
//	logger := log.New(os.Stdout, "MOCKDB: ", log.Lmsgprefix)
//
//	mapp := make(map[string]Elem)
//
//	for k, v := range m {
//		mapp[k] = v
//	}
//
//	return &MockDB{
//		logg: logger,
//		Db:   mapp,
//	}
//}
//
//func (m MockDB) SetData(ctx context.Context, key string, fields map[string]string) error {
//	m.logg.Printf("SET %v: %v\n", key, fields)
//
//	code, _ := strconv.Atoi(fields["code"])
//	attempts, _ := strconv.Atoi(fields["attempts"])
//
//	m.Db[key] = Elem{
//		Code:     code,
//		Attempts: attempts,
//	}
//
//	m.logg.Printf("New map: %v\n\n", m.Db)
//
//	return nil
//}
//
//func (m MockDB) GetAllData(ctx context.Context, key string) (map[string]string, error) {
//	elem, ok := m.Db[key]
//	if !ok {
//		return nil, errors.New("no Elem with this key")
//	}
//
//	return map[string]string{
//		"code":     strconv.Itoa(elem.Code),
//		"attempts": strconv.Itoa(elem.Attempts),
//	}, nil
//}
//
//func (m MockDB) CheckExist(ctx context.Context, key string) bool {
//	_, ok := m.Db[key]
//	return ok
//}
//
//func (m MockDB) Delete(ctx context.Context, key string) error {
//	m.logg.Printf("Delete called for: %v\n", key)
//	if !m.CheckExist(ctx, key) {
//		return errors.New("no Elem with this key")
//	}
//	delete(m.Db, key)
//	m.logg.Printf("New map: %v\n\n", m.Db)
//	return nil
//}
//
//func (m MockDB) SetExpire(ctx context.Context, key string, duration time.Duration) error {
//	m.logg.Printf("expiration set for %v : %v", key, duration)
//	// noop
//	return nil
//}
//
//func (m MockDB) Increment(ctx context.Context, key, field string) error {
//	m.logg.Printf("increment called on : %v", key)
//	if !m.CheckExist(ctx, key) {
//		return errors.New("no Elem with this key")
//	}
//
//	code := m.Db[key].Code
//	attempts := m.Db[key].Attempts
//
//	m.Db[key] = Elem{
//		Code:     code + 1,
//		Attempts: attempts,
//	}
//	m.logg.Printf("New map: %v\n\n", m.Db)
//	return nil
//}
//
////var _ Storage = (*MockDB)(nil)
