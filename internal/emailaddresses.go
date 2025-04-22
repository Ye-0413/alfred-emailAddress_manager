package internal

// EmailAddress represents a stored email address with a name
type EmailAddress struct {
	Name  string
	Email string
}

// GetEmailAddresses returns a list of predefined email addresses
func GetEmailAddresses() []EmailAddress {
	return []EmailAddress{
		{Name: "PolyU Student Number", Email: "24052775r@connect.polyu.hk"},
		{Name: "PolyU Student Name", Email: "ye-aimmeng.jia@connect.polyu.hk"},
		{Name: "PolyU Staff", Email: "aimmeng-ye.jia@polyu.edu.hk"},
		{Name: "Gmail", Email: "aimmeng.life@gmail.com"},
		{Name: "QQ", Email: "1149042251@qq.com"},
		{Name: "Academic", Email: "academic-ye.jia@hotmail.com"},
		{Name: "anything", Email: "anything@jiaye.hk.cn"},
		{Name: "163", Email: "jy_academic@163.com"},
		// Add more email addresses as needed
	}
}
