package db

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Conn - global connection object. To be used though out the app
var Conn *MConn

// MConn - Mongo connection
type MConn struct {
	Session    *mgo.Session
	DbName     string
	ConnString string
}

// GetOne - Get one record
func (conn *MConn) GetOne(query bson.M, tableName string, result interface{}) error {
	session := conn.Session.Copy()
	defer session.Close()

	coll := session.DB(conn.DbName).C(tableName)
	err := coll.Find(query).One(result)

	return err
}

// Upsert - Insert or update record
func (conn *MConn) Upsert(query bson.M, tableName string, doc interface{}) error {
	session := conn.Session.Copy()
	defer session.Close()

	coll := session.DB(conn.DbName).C(tableName)
	_, err := coll.Upsert(query, doc)

	return err
}

// MakeConn - Create mongo connection
func MakeConn(connString string, dbname string) *MConn {
	dialInfo, err := mgo.ParseURL(connString)
	if err != nil {
		panic(err)
	}

	dialInfo.Timeout = 30 * time.Second
	dialInfo.Database = dbname
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return &MConn{Session: session, DbName: dbname, ConnString: connString}
}
