module server

go 1.15

require local/logger v0.0.0-00010101000000-000000000000

replace local/logger => ./logger

require github.com/themakers/osinfo v0.0.0-20171110064631-efd22025c123 // indirect
