package ldDataBaseKit

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
)

var mMgoSession *mgo.Session

//InitMongoDBSession InitMongoDBSession
func InitMongoDBSession(dialInfo *mgo.DialInfo, mode mgo.Mode) error {
	var errDial error
	mMgoSession, errDial = mgo.DialWithInfo(dialInfo)
	if errDial != nil {

		return errors.New("hcLionDataBase.InitMongoDBSession:" + "Connect error " + errDial.Error())
	}
	mMgoSession.SetMode(mode, true)
	return nil
}

//GetMongoDBSession GetMongoDBSession
func GetMongoDBSession() (*mgo.Session, error) {
	if mMgoSession == nil {
		return nil, errors.New("hcLionDataBase.GetMongoDBSession:please InitMongoDBSession before use GetMongoDBSession")
	}
	return mMgoSession.Clone(), nil
}

//CloseMongo 关闭数据库链接
func CloseMongo() {
	mMgoSession.Close()
}
