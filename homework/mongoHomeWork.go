package homework

import (
	"DarProject-master/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var(
	collection *mongo.Collection
)
type howeworkrepo struct{
	dbcon *mongo.Database
}
func NewLessonRepository(config config.MongoConfig) (Repository,error){
	clientOptions:=options.Client().ApplyURI("mongodb://"+config.Host+":"+config.Port)
	client,err := mongo.Connect(context.TODO(),clientOptions)
	if err!=nil{
		return nil,err
	}
	err = client.Ping(context.TODO(),nil)
	if err!=nil{
		return nil,err
	}
	db:=client.Database(config.Database)
	collection=db.Collection("homework")
	return &howeworkrepo{dbcon:db},nil
}
func(hw *howeworkrepo) Get()([]*HomeWork,error){
	findOptions:=options.Find()
	var myHomeWorks []*HomeWork
	cur,err :=collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err!=nil{
		return nil,err
	}
	for cur.Next(context.TODO()){
		var homeWork HomeWork
		err:=cur.Decode(&homeWork)
		if err!=nil{
			return nil,err
		}
		myHomeWorks = append(myHomeWorks,&homeWork)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return myHomeWorks,nil
}
func(hw *howeworkrepo) Add(h *HomeWork) (*HomeWork,error){
	myhomeworks,err:=hw.Get()
	n:=len(myhomeworks)
	if n!=0{
		lastHomeWork:=myhomeworks[n-1]
		h.Id = lastHomeWork.Id+1
	}else{
		h.Id = 1
	}
	_,err=collection.InsertOne(context.TODO(),h)
	if err!=nil{
		return nil,err
	}
	return h,nil
}
func (hw *howeworkrepo) GetById(id int64) (*HomeWork,error){
	filter:=bson.D{{"id",id}}
	homework:=&HomeWork{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&homework)
	if err!=nil{
		return nil,err
	}
	return homework,nil
}
func (hw *howeworkrepo) Update(h *HomeWork)  (*HomeWork,error){
	filter:=bson.D{{"id",h.Id}}
	update:=bson.D{{"$set",bson.D{
		{"createdat",h.CreatedAt},
		{"updatedat",h.UpdatedAt},
		{"taskdescription",h.TaskDescription},
		{"taskmaterial",h.TaskMaterial},
		{"endtime",h.EndTime},
		{"lessonid",h.LessonId},
	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return h,nil
}
func (hw *howeworkrepo) Remove(h *HomeWork) error{
	filter:=bson.D{{"id",h.Id}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}
