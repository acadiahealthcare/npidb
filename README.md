npidb
=====

Write weekly NPI updates to a SQL Server database.  Weekly updates are
downloaded from http://download.cms.gov/ (see
http://download.cms.gov/nppes/NPI_Files.html).  The zip file is written to a
temporary file, and the CSV file within is written to the NPI table.  Existing
records with the same NPI are updated; new records are inserted.

Installation
============
```
go install ./...
```
Initialization
==============
Download one of the Full Replacement Monthly NPI files from http://download.cms.gov/nppes/NPI_Files.html, e.g. `NPPES_Data_Dissemination_October_2017.zip`.
Create the `NPI` and `NPI_Update` tables using `npi-init`:

```
npi-init [--table NPI] "sqlserver://username:password@server:port?database=dbname" /path/to/NPPES_Data_Dissemination_October_2017.zip
```

You may need to increase the default timeout in your DSN, e.g.  `sqlserver://username:password@server:port?database=dbname&connection+timeout=600`.

Usage
=====

```
npi-update [-table NPI] <sqlserver://username:password@server:port?database=dbname>
```
