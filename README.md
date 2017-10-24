npidb
=====

Write weekly NPI updates to a SQL Server database.  Weekly updates are
downloaded from http://download.cms.gov/ (see
http://download.cms.gov/nppes/NPI_Files.html).  The zip file is written to a
temporary file, and the CSV file within is written to the NPI table.  Existing
records with the same NPI are updated; new records are inserted.

Usage
=====
```
npi-update <sqlserver://username:password@server:port?database=dbname>
```

The NPI and NPI_Update tables must exist in the database.  The schema is in
doc/schema.sql.

Installation
============
```
go install ./...
```
