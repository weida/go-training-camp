package data

import (
  "context"
  "github.com/google/wire"
  log "github.com/sirupsen/logrus"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"


)

var ProviderSet = wire.NewSet(NewData)

type Data struct {
  db *DB
}

func NewData(conf *conf.Data) (*Data , func(), error){
  client, err = gorm.Open(
	conf.Database.Driver,
	conf.Database.Source,
  )
  if err != nil {
     log.Error(" open db err %v", err)
  }
  d := &Data{
    db :  client
  }

  return d, func() {
     log.Info("closeing db")
     if err := d.db.Close(); err!= nil {
        log.Error(err)
    }
 
  }, nil
}
