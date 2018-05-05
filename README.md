mongodb-json: MongoDB extended JSON
===================================

`h12.io/mongodb-json` is a fork of gopkg.in/mgo.v2/internal/json
that preserve the order of the fields in an unmarshaled JSON document in bson.D
instead of bson.M.
