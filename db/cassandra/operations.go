package cassandra

import (
	"GO_PROJECT/model"
)

func GetUser(email string) (map[string]interface{}, error) {
	query := "SELECT * FROM users WHERE email_id = ? LIMIT 1"
	iter := CassandraSession.Query(query, email).Iter()

	user := make(map[string]interface{})
	if !iter.MapScan(user) {
		return nil, iter.Close()
	}
	return user, nil
}

func AddUser(userData model.User) error {
	query := `INSERT INTO users (user_id, user_name, email_id, contact_no, gender, password, role, created_at, updated_at)
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Execute the query
	err := CassandraSession.Query(query,
		userData.UserId,
		userData.UserName,
		userData.EmailId,
		userData.ContactNo,
		userData.Gender,
		userData.Password,
		userData.Role,
		userData.CreatedAt,
		userData.UpdatedAt,
	).Exec()

	return err
}
