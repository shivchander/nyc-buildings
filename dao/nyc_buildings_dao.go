package dao

import (
	"log"
	"fmt"
	"os"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/csv"
	"strconv"
)

type Buildings struct {
	ID          	bson.ObjectId 	`bson:"_id" json:"id"`
	BIN         	int 			`bson:"BIN" json:"BIN"`
	Name        	string        	`bson:"Name" json:"Name"`
	Construct_Yr 	int       		`bson:"Construct_Yr" json:"Construct_Yr"`
	Shape_Area 		float64        	`bson:"Shape_Area" json:"Shape_Area"`
	Shape_Len 		float64        	`bson:"Shape_Len" json:"Shape_Len"`
	lststatype		string			`bson:"lststatype" json:"lststatype"`
	GEOMSOURCE		string			`bson:"GEOMSOURCE" json:"GEOMSOURCE"`
}


type BuildingsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database


const (
	COLLECTION = "buildings"
)

// Establish a connection to database
func (m *BuildingsDAO) Connect() {
	fmt.Println("Connecting DB...")
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
	fmt.Println("Connection Established")

}

func (m *BuildingsDAO) Insert(buildings Buildings) error {
	err := db.C(COLLECTION).Insert(&buildings)
	return err
}

func(m *BuildingsDAO) PushAll() {

	var Buildings []Buildings
	 db.C(COLLECTION).Find(bson.M{}).All(&Buildings)
	//check if data already pushed
	if len(Buildings) <=0 {
	//insert only if table is empty
	fmt.Println("Push 1000 records from excel Invoked")
		buildings := extract()
		//fmt.Println( len(buildings))
	
	for i, b := range buildings {
		if i <= 1000  && i > 0 {
		b.ID = bson.NewObjectId()
		db.C(COLLECTION).Insert(*b) 
		}
	}		 

	fmt.Println("Push 1000 records into mongodb Finished")
	}
}

func extract() []*Buildings {
	fmt.Println("Extraction Beginning")
	result := []*Buildings{}
	f, _ := os.Open("./building.csv")
	defer f.Close()
	r := csv.NewReader(f)

	for record, err := r.Read(); err == nil; record, err = r.Read() {
		building := new(Buildings)
		building.BIN, _ = strconv.Atoi(record[1])
		building.Name = record[3]
		building.Construct_Yr, _ = strconv.Atoi(record[2])
		building.Shape_Area, _ = strconv.ParseFloat(record[10], 64)
		building.Shape_Len, _ = strconv.ParseFloat(record[11], 64)
		building.lststatype = record[5]
		building.GEOMSOURCE = record[14]
		
		result = append(result, building)
	}
	return result
}


// Find list of Buildings
func (m *BuildingsDAO) FindAll() ([]Buildings, error) {
	var Buildings []Buildings
	err := db.C(COLLECTION).Find(bson.M{}).All(&Buildings)
	return Buildings, err
}

// Find a Buildings by its id
func (m *BuildingsDAO) FindById(bin int) (Buildings, error) {
	var Buildings Buildings
	err := db.C(COLLECTION).Find(bson.D{{"BIN", bin}}).One(&Buildings)
	fmt.Println(Buildings.Shape_Area)
	return Buildings, err
}


func (m *BuildingsDAO) FindByYr(start int, end int) ([]Buildings, error) {
	var Buildings []Buildings
	err := db.C(COLLECTION).Find(
		bson.M{
			"Construct_Yr": bson.M{
				"$gte": start,
				"$lte": end,
			},
		}).All(&Buildings)
	return Buildings, err
}


