package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Note struct {
	// gorm fileds Model
	ID        uint           `gorm:"primaryKey"`
	Created   int64          `gorm:"autoCreateTime;->;<-:create"` // Использование формата unix в секундах при создании
	Updated   int64          `gorm:"autoUpdateTime;->;<-:create"` // Использование формата unix в секундах в качестве времени обновления
	DeletedAt gorm.DeletedAt `gorm:"index;->;<-:create"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Name        string `json:"name" gorm:"->"`
}

func (n *Note) GetNote(db *gorm.DB) *gorm.DB {
	return db.First(&n)
}

func (n *Note) UpdateNoteFields(db *gorm.DB) *gorm.DB {
	return db.Updates(&n)
}

func (n *Note) UpdateFullNote(db *gorm.DB) *gorm.DB {
	var find Note
	find = *n
	fmt.Println("asd: =", find)
	first := db.First(&find)
	fmt.Println("asd: =", find)
	fmt.Println("n: =", *n)

	if first.Error == nil {
		fmt.Println("first")
		find = *n
		// copyParams(&asd, n)
		return db.Updates(&find)
	} else {
		fmt.Println("first Error:", first.Error)
		return db.Create(&n)

	}

}

// type NoteInterface interface {
// 	Title_() string
// 	Description_() string
// 	Name_() string
// }

// func (note Note) Title_() string {
// 	return note.Title
// }
// func (note Note) Description_() string {
// 	return note.Description
// }

// func (note Note) Name_() string {
// 	return note.Name
// }

// func copyParams(left *Note, right *Note) {

// 	var noteRightInterface NoteInterface = Note(*right)
// 	var noteLeftInterface NoteInterface = Note(*left)

// 	// // var i NoteInt = Note{}

// 	// var note interface{} = Note(*right)

// 	fmt.Println("copyParams")

// 	fmt.Println("noteRightInterface1:  ", noteRightInterface)
// 	fmt.Println("noteLeftInterface1:  ", noteLeftInterface)

// 	// var n NoteInterface = Note{Title: "Title old"}
// 	// var n2 NoteInterface = Note{Title: "Title2"}

// 	r := reflect.ValueOf(&noteRightInterface).Elem()
// 	tmpRight := reflect.New(r.Elem().Type()).Elem()

// 	l := reflect.ValueOf(&noteLeftInterface).Elem()
// 	tmpLeft := reflect.New(r.Elem().Type()).Elem()

// 	// Copy the struct value contained in interface to
// 	// the temporary variable.
// 	tmpRight.Set(r.Elem())
// 	tmpLeft.Set(l.Elem())

// 	// Set the field.
// 	// tmpRight.FieldByName("Title").SetString("Hello")

// 	// Set the interface to the modified struct value.
// 	// r.Set(tmpRight)

// 	var noteRightInterface2 NoteInterface = Note{}

// 	var aa NoteInterface

// 	fields := reflect.VisibleFields(reflect.TypeOf(aa))

// 	fmt.Println("fields:", fields)

// 	t := reflect.TypeOf(noteRightInterface2)
// 	// v := reflect.ValueOf(noteRightInterface)
// 	// vLeft := reflect.ValueOf(noteLeftInterface)

// 	for i := 0; i < t.NumField(); i++ {
// 		fmt.Println("Field type ", t.Field(i))
// 	}

// 	for i := 0; i < tmpRight.NumField(); i++ {
// 		fmt.Println("<")
// 		fmt.Printf("%+v\n", tmpRight.Field(i))
// 		fmt.Println(tmpRight.Field(i))
// 		fmt.Println(tmpRight.Field(i).IsZero())
// 		fmt.Println("CanSet : ", tmpRight.Field(i).CanSet())
// 		fmt.Println(">")

// 		if !tmpRight.Field(i).IsZero() {
// 			tmpLeft.Field(i).Set(tmpRight.Field(i))
// 		}

// 		// vLeft.Field(i).Set(v.Field(i))
// 	}

// 	// fmt.Println("noteLeftInterface2:", noteLeftInterface)

// 	r.Set(tmpRight)
// 	l.Set(tmpLeft)

// 	fmt.Println("noteRightInterface2:  ", noteRightInterface)
// 	fmt.Println("noteLeftInterface2:  ", noteLeftInterface)
// }

func (n *Note) DeleteNote(db *gorm.DB) *gorm.DB {
	return db.Delete(&n)
}

func (n *Note) CreateNote(db *gorm.DB) *gorm.DB {
	return db.Create(&n)
}

func (n *Note) GetNotes(db *gorm.DB, notes *[]Note, limit, offset int) *gorm.DB {
	gormDb := db.Limit(limit).Offset(offset).Find(notes)
	return gormDb
}
