package dvalid

/*
Entity should depend on abstraction not on concretions
*/

type Mydb struct {
}

func (db Mydb) Query() interface{} {
	return []string{"name_1", "name_2", "name_3"}
}

type Post struct {
}

func (db Post) Query() interface{} {
	return map[string]string{
		"6b026a04-514f-4176-b8d0-79f7674cf864": "name_1",
		"d2f068c0-d94e-41d4-8b4a-07f4cdf8f178": "name_2",
		"b61d17f8-c2a5-4227-b20e-9784a088a206": "name_3",
	}
}

type DBConnection interface {
	Query() interface{}
}

type UserRepo struct {
	db DBConnection
}

func (r UserRepo) GetUsers() []string {
	var users []string
	response := r.db.Query()

	switch getTypes := response.(type) {
	case map[string]string:
		for _, u := range getTypes {
			users = append(users, u)
		}
		return users
	case []string:
		return response.([]string)
	}

	return []string{}
}
