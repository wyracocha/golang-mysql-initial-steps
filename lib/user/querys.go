package user

func Query () struct {save string; findById string; changeStatus string; changeFullname string}{
	return struct{
		save string
		findById string
		changeStatus string
		changeFullname string
	}{
		save: "insert into user(id, fullname, createdAt, birthdate, status) values(?, ?, ?, ?, ?)",
		findById: "select * from user where id = ?",
		changeStatus: "update user set status = ? where id = ?",
		changeFullname: "update user set fullname = ? where id = ?",
	}
}