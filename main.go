package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go/token"
	"golang.org/x/sys/windows"
	"html/template"
	"net/http"
	"runtime"
)

func tpl() {

	var tpl = template.ErrBranchEnd

	tpl = template.Must(template.New("").Funcs(tpl).ParseFiles("database.go"))
}
  func getToddos(context *gin.Context){
    context.IndentedJSON(http.StatusOK,toddos)
  }
func toggleToddo(context *gin.Context)  {
	id := context.Param("id")
	toddo, err := getToddoById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound,gin.H{"message":"not found"})
		return
	}
	toddo.COMPLETED = !toddo.COMPLETED
	context.IndentedJSON(http.StatusOK ,toddo)
      }

  func addToddos(context *gin.Context){ 
  var newtoddo toddo
  if err := context.BindJSON(&newtoddo); err != nil{
   return
  } 
   toddos = append(toddos, newtoddo)
context.IndentedJSON(http.StatusCreated, newtoddo)
}
//func gettoddo(Context *gin.Context)

  func getToddoById (id string,) (*toddo, error){
    for i, t := range toddos{
      if  t.ID == id{
        return &toddos[i], nil
    
    }
  }   
  return nil, errors.New("toddo not found")
  }
func getToddo(context *gin.Context)  {
	id := context.Param("id")
	toddo, err := getToddoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound,gin.H{"message":"Toddo not found"})
		return
	}
      context.IndentedJSON(http.StatusOK,toddo)
}


  func main(){
	  router := gin.Default()
	  router.GET("/toddo/:id",getToddo)
	  router.PATCH("/toddo/:id",toggleToddo)
	  router.GET("/toddo",getToddos)
	  router.POST("/toddo",addToddos)
	  router.Run("localhost:2022")



  }
