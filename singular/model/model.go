/*
Copyright 2016 - 2017 Huawei Technologies Co., Ltd. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package model

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Huawei/containerops/common"
	"github.com/Huawei/containerops/common/model"
)

var (
	DB *gorm.DB
)

// init()
func init() {

}

// OpenDatabase is
func OpenDatabase(dbconfig *common.DatabaseConfig) error {
	var err error

	driver, host, port, user, password, db := dbconfig.Driver, dbconfig.Host, dbconfig.Port, dbconfig.User, dbconfig.Password, dbconfig.Name

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local", user, password, host, port, db)
	if DB, err = gorm.Open(driver, dsn); err != nil {
		log.Fatal("Initialization database connection error.", err)
		return err
	} else {
		if db := DB.DB(); db == nil {
			return fmt.Errorf("database connection is invalid")
		}

		if err := DB.DB().Ping(); err != nil {
			return err
		}

		DB.DB().SetMaxIdleConns(10)
		DB.DB().SetMaxOpenConns(100)
		DB.SingularTable(true)
	}

	return nil
}

// Migrate is
func Migrate() error {
	// Singular Tables
	DB.AutoMigrate(&SingularV1{}, &DeploymentV1{}, &InfraV1{}, &ComponentV1{})

	// Label V1 Require Table
	DB.AutoMigrate(&model.LabelV1{})

	return nil
}
