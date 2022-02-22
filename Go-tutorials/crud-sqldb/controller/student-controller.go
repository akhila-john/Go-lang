// package student_controller

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type Student struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Class string `json:"class"`
// 	Marks int    `json:"marks"`
// }

// var studentList []Student

// func getStudent(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	//return the list of rows selected by the query.
// 	result, err := db.Query("SELECT id,name,class,marks from student")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer result.Close()
// 	// Next prepares the next result row for reading with the Scan method
// 	// It returns true on success, or false if there is no next result row or an error
// 	// happened while preparing it.
// 	for result.Next() {
// 		var student Student
// 		//  the result columns are copied to the id,name,class,marks fields of the student struct .
// 		// tell Scan() to copy the retrieved data into the memory locations used by those variables,
// 		// and Scan() will return nil if this is successful, otherwise it returns an error.
// 		err := result.Scan(&student.ID, &student.Name, &student.Class, &student.Marks)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		studentList = append(studentList, student)
// 	}
// 	json.NewEncoder(w).Encode(studentList)
// }

// func createstudent(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")

// 	var keyVal Student
// 	// decodes the json body from the request and stores at keyval
// 	json.NewDecoder(r.Body).Decode(&keyVal)
// 	//prepare a statement that will create new post
// 	stmt, err := db.Prepare("INSERT INTO student(id, name, class, marks) VALUES (?, ?, ?, ?)")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer stmt.Close()
// 	//////////////// alternate way ////////////////////////////////////////////
// 	// read the request body and extract name etc.. value.
// 	// body, err := ioutil.ReadAll(r.Body)
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }

// 	//use “stmt.Exec()” and put in the  variables we created, check for errors and finally just sends a message
// 	// to the client that a new post was created.
// 	//keyVal := make(map[string]string)
// 	//json.Unmarshal(body, &keyVal)
// 	//id := keyVal["id"]
// 	// name := keyVal["name"]
// 	// class := keyVal["class"]
// 	// marks := keyVal["marks"]
// 	////////////////////////////////////////////////////////////////////////////////
// 	// need to pass the needed variables with the exec.
// 	_, err = stmt.Exec(keyVal.ID, keyVal.Name, keyVal.Class, keyVal.Marks)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	fmt.Fprintf(w, "new student added")

// }
// func getStudentById(w http.ResponseWriter, r *http.Request) {

// 	params := mux.Vars(r)
// 	// query that gets the student with the id from our url
// 	// extract params - params["id"]
// 	stmt, err := db.Query("SELECT id, name, class, marks FROM student WHERE id=?", params["id"])
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer stmt.Close()

// 	var student Student
// 	//	create a new Student object, iterate through the result (even if we know that there is only one post with that specific id)
// 	// and scan the values into our object.
// 	for stmt.Next() {

// 		err := stmt.Scan(&student.ID, &student.Name, &student.Class, &student.Marks)
// 		if err != nil {
// 			panic(err.Error())
// 		}

// 	}
// 	//encode our student to JSON and send it away.
// 	json.NewEncoder(w).Encode(student)
// }

// func deleteStudent(w http.ResponseWriter, r *http.Request) {
// 	// gets the variables from the url in params
// 	params := mux.Vars(r)

// 	stmt, err := db.Prepare("DELETE FROM student WHERE id = ?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	//extracts the id from the params
// 	_, err = stmt.Exec(params["id"])
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	fmt.Fprintf(w, "Student with ID = %s was deleted", params["id"])

// }
