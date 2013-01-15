package sdb


// define simple database interface
type SimpleDatabase interface{
	GetTables()[]string
}

// define simple table interface
type SimepleTable interface{
	GetRows()int
	GetColnums()int
}

// define simple row interface
type SimpleRow interface{
	Next()bool
	GetString(column int)string
}


