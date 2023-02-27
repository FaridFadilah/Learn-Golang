package packageinit

var connection string

func init(){
	connection = "mysql"
}

func Database() string {
	return connection
}