package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Student1 struct {
	Stuid int    `bson:"stuid"`
	Name  string `bson:"name"`
}

func Insert(a interface{}, dbconn *mgo.Collection) {
	err := dbconn.Insert(a)
	if err != nil {
		panic("error in inserting")
	} else {
		fmt.Println("Inserted")
	}
}
func Retrive1(a *Student1, dbconn *mgo.Collection) {
	fmt.Println("\nChoose the option to retrive the value by using \n 01. student id \n 02. student Name \n 03. Retrive All ")
	var c int
	fmt.Scanf("%d", &c)
	switch c {
	///retrive using id
	case 1:
		fmt.Println("\n Enter the value of student id")
		var val int
		fmt.Scanf("%d", &val)
		err := dbconn.Find(bson.M{"stuid": val}).One(&a)
		if err != nil {
			panic("error in Retriving ")
		} else {
			fmt.Println("ok you Got it :-)")
		}
		fmt.Println(a.Name)
		///retrive using name
	case 2:
		fmt.Println("\n Enter the value of student Name")
		var val string
		fmt.Scanf("%s", &val)
		err := dbconn.Find(bson.M{"name": val}).One(&a)
		if err != nil {
			panic("error in Retriving ")
		} else {
			fmt.Println("ok you Got it :-)")
		}
		fmt.Println(a.Stuid)
	///retrive All
	case 3:
		iter := dbconn.Find(nil).Iter()
		fmt.Println("\n*************************Data Retrived***********************\n")
		for iter.Next(&a) {
			fmt.Println("\nStudent id is :", a.Stuid, "Student Name is :", a.Name, "\n")
		}
		fmt.Println(a)
	}
}
func update1(dbconn *mgo.Collection) {
	fmt.Println("\nChoose the option to update the value by using \n 01. update one \n 02. update all ")
	var c int
	fmt.Scanf("%d", &c)
	switch c {
	///retrive using id
	case 1:
		fmt.Println("\n replace by \n 01 .id \n 02 .name \n")
		var o int
		fmt.Scanf("%d", &o)
		switch o {
		case 1:
			fmt.Println("\n Enter the id value to replace \n Enter the name value to change ")
			var a int
			var b string
			fmt.Scanf("%d", &a)
			fmt.Scanf("%s", &b)
			err := dbconn.Update(bson.M{"stuid": a}, bson.M{"$set": bson.M{"name": b}})
			if err != nil {
				panic("Error in updating id")
			} else {
				fmt.Println("updated")
			}
		case 2:
			fmt.Println("\n Enter the student name value to replace \n  Enter the name value to change ")
			var a string
			var b string
			fmt.Scanf("%s", &a)
			fmt.Scanf("%s", &b)

			err := dbconn.Update(bson.M{"name": a}, bson.M{"$set": bson.M{"name": b}})
			if err != nil {
				panic("Error in updating name")
			} else {
				fmt.Println("updated")
			}
		}
	case 2:
		fmt.Println("\n replace by \n 01 .id \n 02 .name \n")
		var o int
		fmt.Scanf("%d", &o)
		switch o {
		case 1:
			fmt.Println("\n Enter the id value to replace \n  Enter the name value to change ")
			var a int
			var b string
			fmt.Scanf("%d", &a)
			fmt.Scanf("%s", &b)
			_, err := dbconn.UpdateAll(bson.M{"stuid": a}, bson.M{"$set": bson.M{"name": b}})
			if err != nil {
				panic("Error in updating id")
			} else {
				fmt.Println("updated")
			}
		case 2:
			fmt.Println("\n Enter the student name value to replace \n  Enter the name value to change ")
			var a string
			var b string
			fmt.Scanf("%s", &a)
			fmt.Scanf("%s", &b)
			_, err := dbconn.UpdateAll(bson.M{"name": a}, bson.M{"$set": bson.M{"name": b}})
			if err != nil {
				panic("Error in updating name")
			} else {
				fmt.Println("updated")
			}
		}

	}
}
func Remove1(dbconn *mgo.Collection) {
	fmt.Println("\nChoose the option to delete record \n 01. Delete one record \n 02. delete all record ")
	var o int
	fmt.Scanf("%d", &o)
	switch o {
	///retrive using id
	case 1:
		var a int
		fmt.Println("Enter the student id to remove ")
		fmt.Scanf("%d", &a)
		err := dbconn.Remove(bson.M{"stuid": a})
		if err != nil {
			fmt.Printf("remove fail %v\n", err)
			// os.Exit(1)
		} else {
			fmt.Println("removed")
		}
	case 2:
		fmt.Println("\nChoose the option to delete record \n 01. Delete all record using id \n 02. delete all record ")
		var cc int
		fmt.Scanf("%d", &cc)
		switch cc {
		case 1:
			var a int
			fmt.Println("Enter the student id to remove ")
			fmt.Scanf("%d", &a)
			_, err := dbconn.RemoveAll(bson.M{"stuid": a})
			if err != nil {
				panic("Error in Deleting documents")
			} else {
				fmt.Println("Deleted")
			}
		case 2:
			_, err := dbconn.RemoveAll(nil)
			if err != nil {
				panic("Error in Deleting documents")
			} else {
				fmt.Println("Deleted")
			}
		}
	}
}
func main() {
	s := &Student1{}
	session, err := mgo.Dial("127.0.0.1:27017")
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		panic("----------:-(--Error in connecting databases------------------:-(--------")
	} else {
		fmt.Println("------------------------>Database is connected<----------------------")
	}
	dbconn := session.DB("school1").C("student2")
	var c int
	///Author
	//Patern for N LETTER
	for kk := 1; kk <= 7; kk++ {
		for ll := 1; ll <= 10; ll++ {
			kkk := kk + 1
			lll := ll

			if ll <= 3 || ll >= 8 && ll <= 10 || kk == ll || kkk == lll {
				fmt.Print("-")
				time.Sleep(100 * time.Millisecond)

			} else {
				fmt.Print(" ")
				time.Sleep(100 * time.Millisecond)
			}
		}
		fmt.Println("")
	}
	fmt.Println("\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x50\x72\x6f\x67\x72\x61\x6d\x6d\x65\x64\x20\x42\x79\x20\x4e\x41\x47\x41\x52\x41\x4a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20")
	for {
		fmt.Println("\nchoose the option \n 1.insert \n 2.retrive \n 3.update \n 4.remove \n 0.Exit")
		fmt.Scanf("%d", &c)
		if c == 0 {
			os.Exit(1)
		}
		switch c {
		case 1:
			fmt.Println("\n--------------INSERTION-------------------------------\n")
			fmt.Println("enter the value to store in student Collection \n enter the student id \n")
			fmt.Scanf("%d", &s.Stuid)
			fmt.Println("\n enter the student Name \n")
			fmt.Scanf("%q\n", &s.Name)
			Insert(s, dbconn)
		case 2:
			fmt.Println("\n--------------RETRIVE-------------------------------\n")
			Retrive1(s, dbconn)
		case 3:
			fmt.Println("\n--------------UPDATING-------------------------------\n")
			update1(dbconn)
		case 4:
			Remove1(dbconn)
		}
	}
}
