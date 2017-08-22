package agregate

import mgo "gopkg.in/mgo.v2"

func NewRepo() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://localhost/local")
	if err != nil {

		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)

	return session.Copy(), nil
}
