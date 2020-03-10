package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var Client, err = mongo.Connect(context.Background(), "mongodb+srv://nithyavardhan:hoaxapp@69@cluster0-xwnvc.mongodb.net/test?retryWrites=true&w=majority", nil)
