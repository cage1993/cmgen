// Code generated by cmgen.
// source:
// advanced.yaml

package base

import (
    "errors"
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
    "github.com/liamylian/jsontime"
    "time"
)

var (
    dbName = "test"
    ErrorInvalidObjectId = errors.New("invalid objectId")
    json = jsontime.ConfigWithCustomTimeFormat
    MongoSession *mgo.Session
)

func init() {
    //FIXME  Here to init MongoSession
    MongoSession = new(mgo.Session)
}

      
const (
    CollectionUser = "users"
)
    

func GetSessionAndCollection(collection string) (*mgo.Session, *mgo.Collection) {
	s := MongoSession.Copy()
	c := s.DB(dbName).C(collection)

	return s, c
}

type User struct {
    ID bson.ObjectId `bson:"_id" json:"id"`
    UserName string `bson:"user_name" json:"user_name,omitempty" valid:"required~first name is blank"`
    Email string `bson:"email" json:"email,omitempty" valid:"required,email"`
    Password string `bson:"password" json:"password,omitempty" valid:"required"`
    CreatedAt time.Time `bson:"created_at" json:"created_at" time_format:"2006-01-02 15:04:05"`
    UpdatedAt time.Time `bson:"updated_at" json:"updated_at" time_format:"2006-01-02 15:04:05"`
    Deleted int `bson:"deleted" json:"-"`
}

func NewUser() *User{
    return &User{}
}

func (user *User) Insert() error {
    s, c := GetSessionAndCollection(CollectionUser)
    defer s.Close()
    
    if err := c.EnsureIndex(mgo.Index{
        Key: []string{"user_name"},
        Unique:     true,
        DropDups:   false,
        Background: true,
    }); err != nil {
        return err
    }
    if err := c.EnsureIndex(mgo.Index{
        Key: []string{"email"},
        Unique:     true,
        DropDups:   false,
        Background: true,
    }); err != nil {
        return err
    }

    user.ID = bson.NewObjectId()
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    user.Deleted = 0

    return c.Insert(user)
}

func UpdateUserByID(id interface{}, user *User) error {
    s, c := GetSessionAndCollection(CollectionUser)
    defer s.Close()

    user.UpdatedAt = time.Now()

    switch id := id.(type) {
    case bson.ObjectId:
        return c.UpdateId(id, bson.M{
            "$set": user,
        })
    case string:
        if !bson.IsObjectIdHex(id) {
            return ErrorInvalidObjectId
        }
        return c.UpdateId(bson.ObjectIdHex(id), bson.M{
            "$set": user,
        })
    }

    return errors.New("no bson.ObjectId")
}

func UpdateUserByIDAndEntityMap(id interface{}, updateMap map[string]interface{}) error {
    s, c := GetSessionAndCollection(CollectionUser)
    defer s.Close()

    if updateMap == nil { return nil }
    updateMap["updated_at"] = time.Now()

    switch id := id.(type) {
    case bson.ObjectId:
        return c.UpdateId(id, bson.M{
            "$set": updateMap,
        })
    case string:
        if !bson.IsObjectIdHex(id) {
            return ErrorInvalidObjectId
        }
        return c.UpdateId(bson.ObjectIdHex(id), bson.M{
            "$set": updateMap,
        })
    }

    return errors.New("no bson.ObjectId")
}

// Update finds a single document matching the provided selector document
// and modifies it according to the update document.
// If the session is in safe mode (see SetSafe) a ErrNotFound error is
// returned if a document isn't found, or a value of type *LastError
// when some other error is detected.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Updating
//     http://www.mongodb.org/display/DOCS/Atomic+Operations
//
func UpdateUser(selector interface{}, user *User) error {
    s, c := GetSessionAndCollection(CollectionUser)
    defer s.Close()

    user.UpdatedAt = time.Now()

    return c.Update(selector, bson.M{
        "$set": user,
    })
}

// UpdateAll finds all documents matching the provided selector document
// and modifies them according to the update document.
// If the session is in safe mode (see SetSafe) details of the executed
// operation are returned in info or an error of type *LastError when
// some problem is detected. It is not an error for the update to not be
// applied on any documents because the selector doesn't match.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Updating
//     http://www.mongodb.org/display/DOCS/Atomic+Operations
//
func UpdateUserAll(selector interface{}, user *User) (*mgo.ChangeInfo, error) {
    s, c := GetSessionAndCollection(CollectionUser)
    defer s.Close()

    user.UpdatedAt = time.Now()

    return c.UpdateAll(selector, bson.M{
        "$set": user,
    })
}

func GetUserByID(id interface{}) (*User, error) {
    s, c := GetSessionAndCollection(CollectionUser)
    defer s.Close()

    user := new(User)
    var err error
    switch id := id.(type) {
    case bson.ObjectId:
    	err = c.FindId(id).One(user)
    case string:
        if !bson.IsObjectIdHex(id) {
            return nil, ErrorInvalidObjectId
        }
    	err = c.FindId(bson.ObjectIdHex(id)).One(user)
    }
    if err == mgo.ErrNotFound {
        return nil, nil
    }
    return user, nil
}

func GetOneUserByQuery(query map[string]interface{}) (*User, error) {
    s, c := GetSessionAndCollection(CollectionUser)
    defer s.Close()

    if query == nil { query = map[string]interface{}{} }
    query["deleted"] = 0

    user := new(User)

    err := c.Find(query).One(user)
    if err == mgo.ErrNotFound {
        return nil, nil
    }
    return user, nil
}

func ListAllUserByQuery(query map[string]interface{}) ([]*User, error) {
    s, c := GetSessionAndCollection(CollectionUser)
    defer s.Close()

    if query == nil { query = map[string]interface{}{} }
    query["deleted"] = 0

    user := make([]*User, 0)

    return user, c.Find(query).All(&user)
}

func ExistUserByID(id string) (bool, error) {
    if !bson.IsObjectIdHex(id) {
        return false, ErrorInvalidObjectId
    }

    s, c := GetSessionAndCollection(CollectionUser)
    defer s.Close()

    user := new(User)

    if err := c.FindId(bson.ObjectIdHex(id)).One(user); err != nil {
        if err == mgo.ErrNotFound {
            return false, nil
        }
        return false, err
    }

    return true, nil
}

func DeleteUserByID(id string) error {
    if !bson.IsObjectIdHex(id) {
        return ErrorInvalidObjectId
    }
    s, c := GetSessionAndCollection(CollectionUser)
    defer s.Close()

    return c.UpdateId(bson.ObjectIdHex(id), bson.M{
        "$set": bson.M{"deleted": 1},
    })
}