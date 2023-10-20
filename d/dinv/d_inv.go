package dinv

/*
Entity should depend on abstraction not on concretions
*/

type MyDB struct {
}

func (db MyDB) Query() []string {
	return []string{"name_1", "name_2", "name_3"}
}

type Post struct {
}

func (db Post) QueryPost() map[string]string {
	return map[string]string{
		"6b026a04-514f-4176-b8d0-79f7674cf864": "name_1",
		"d2f068c0-d94e-41d4-8b4a-07f4cdf8f178": "name_2",
		"b61d17f8-c2a5-4227-b20e-9784a088a206": "name_3",
	}
}

type UsersRepo struct {
	db MyDB
	//db Post
}

func (r UsersRepo) GetUsers() []string {
	responses := r.db.Query()
	//res := r.db.QueryPost()
	//var users []string
	//for _, u := range res {
	//	users = append(users, u)
	//}
	//return users
	return responses
}
