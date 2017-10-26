package schema

import (
	"database/sql"
	"fmt"
)

func CreateTables(txn *sql.Tx, tableName string) (err error) {
	_, err = txn.Exec(npiTableSchema(tableName))
	if err != nil {
		return
	}
	_, err = txn.Exec(updateTableSchema(tableName))
	return
}

func npiTableSchema(tableName string) string {
	return fmt.Sprintf(`
		CREATE TABLE %s (
			"NPI" varchar(10),
			"Entity Type Code" varchar(1),
			"Replacement NPI" varchar(10),
			"Employer Identification Number (EIN)" varchar(9),
			"Provider Organization Name (Legal Business Name)" varchar(70),
			"Provider Last Name (Legal Name)" varchar(35),
			"Provider First Name" varchar(20),
			"Provider Middle Name" varchar(20),
			"Provider Name Prefix Text" varchar(5),
			"Provider Name Suffix Text" varchar(5),
			"Provider Credential Text" varchar(20),
			"Provider Other Organization Name" varchar(70),
			"Provider Other Organization Name Type Code" varchar(1),
			"Provider Other Last Name" varchar(35),
			"Provider Other First Name" varchar(20),
			"Provider Other Middle Name" varchar(20),
			"Provider Other Name Prefix Text" varchar(5),
			"Provider Other Name Suffix Text" varchar(5),
			"Provider Other Credential Text" varchar(20),
			"Provider Other Last Name Type Code" varchar(1),
			"Provider First Line Business Mailing Address" varchar(55),
			"Provider Second Line Business Mailing Address" varchar(55),
			"Provider Business Mailing Address City Name" varchar(40),
			"Provider Business Mailing Address State Name" varchar(40),
			"Provider Business Mailing Address Postal Code" varchar(20),
			"Provider Business Mailing Address Country Code (If outside U.S.)" varchar(2),
			"Provider Business Mailing Address Telephone Number" varchar(20),
			"Provider Business Mailing Address Fax Number" varchar(20),
			"Provider First Line Business Practice Location Address" varchar(55),
			"Provider Second Line Business Practice Location Address" varchar(55),
			"Provider Business Practice Location Address City Name" varchar(40),
			"Provider Business Practice Location Address State Name" varchar(40),
			"Provider Business Practice Location Address Postal Code" varchar(20),
			"Provider Business Practice Location Address Country Code (If outside U.S.)" varchar(20),
			"Provider Business Practice Location Address Telephone Number" varchar(20),
			"Provider Business Practice Location Address Fax Number" varchar(20),
			"Provider Enumeration Date" varchar(10),
			"Last Update Date" varchar(10),
			"NPI Deactivation Reason Code" varchar(2),
			"NPI Deactivation Date" varchar(10),
			"NPI Reactivation Date" varchar(10),
			"Provider Gender Code" varchar(1),
			"Authorized Official Last Name" varchar(35),
			"Authorized Official First Name" varchar(20),
			"Authorized Official Middle Name" varchar(20),
			"Authorized Official Title or Position" varchar(35),
			"Authorized Official Telephone Number" varchar(20),
			"Healthcare Provider Taxonomy Code_1" varchar(10),
			"Provider License Number_1" varchar(20),
			"Provider License Number State Code_1" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_1" varchar(1),
			"Healthcare Provider Taxonomy Code_2" varchar(10),
			"Provider License Number_2" varchar(20),
			"Provider License Number State Code_2" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_2" varchar(1),
			"Healthcare Provider Taxonomy Code_3" varchar(10),
			"Provider License Number_3" varchar(20),
			"Provider License Number State Code_3" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_3" varchar(1),
			"Healthcare Provider Taxonomy Code_4" varchar(10),
			"Provider License Number_4" varchar(20),
			"Provider License Number State Code_4" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_4" varchar(1),
			"Healthcare Provider Taxonomy Code_5" varchar(10),
			"Provider License Number_5" varchar(20),
			"Provider License Number State Code_5" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_5" varchar(1),
			"Healthcare Provider Taxonomy Code_6" varchar(10),
			"Provider License Number_6" varchar(20),
			"Provider License Number State Code_6" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_6" varchar(1),
			"Healthcare Provider Taxonomy Code_7" varchar(10),
			"Provider License Number_7" varchar(20),
			"Provider License Number State Code_7" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_7" varchar(1),
			"Healthcare Provider Taxonomy Code_8" varchar(10),
			"Provider License Number_8" varchar(20),
			"Provider License Number State Code_8" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_8" varchar(1),
			"Healthcare Provider Taxonomy Code_9" varchar(10),
			"Provider License Number_9" varchar(20),
			"Provider License Number State Code_9" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_9" varchar(1),
			"Healthcare Provider Taxonomy Code_10" varchar(10),
			"Provider License Number_10" varchar(20),
			"Provider License Number State Code_10" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_10" varchar(1),
			"Healthcare Provider Taxonomy Code_11" varchar(10),
			"Provider License Number_11" varchar(20),
			"Provider License Number State Code_11" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_11" varchar(1),
			"Healthcare Provider Taxonomy Code_12" varchar(10),
			"Provider License Number_12" varchar(20),
			"Provider License Number State Code_12" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_12" varchar(1),
			"Healthcare Provider Taxonomy Code_13" varchar(10),
			"Provider License Number_13" varchar(20),
			"Provider License Number State Code_13" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_13" varchar(1),
			"Healthcare Provider Taxonomy Code_14" varchar(10),
			"Provider License Number_14" varchar(20),
			"Provider License Number State Code_14" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_14" varchar(1),
			"Healthcare Provider Taxonomy Code_15" varchar(10),
			"Provider License Number_15" varchar(20),
			"Provider License Number State Code_15" varchar(2),
			"Healthcare Provider Primary Taxonomy Switch_15" varchar(1),
			"Other Provider Identifier_1" varchar(20),
			"Other Provider Identifier Type Code_1" varchar(2),
			"Other Provider Identifier State_1" varchar(2),
			"Other Provider Identifier Issuer_1" varchar(80),
			"Other Provider Identifier_2" varchar(20),
			"Other Provider Identifier Type Code_2" varchar(2),
			"Other Provider Identifier State_2" varchar(2),
			"Other Provider Identifier Issuer_2" varchar(80),
			"Other Provider Identifier_3" varchar(20),
			"Other Provider Identifier Type Code_3" varchar(2),
			"Other Provider Identifier State_3" varchar(2),
			"Other Provider Identifier Issuer_3" varchar(80),
			"Other Provider Identifier_4" varchar(20),
			"Other Provider Identifier Type Code_4" varchar(2),
			"Other Provider Identifier State_4" varchar(2),
			"Other Provider Identifier Issuer_4" varchar(80),
			"Other Provider Identifier_5" varchar(20),
			"Other Provider Identifier Type Code_5" varchar(2),
			"Other Provider Identifier State_5" varchar(2),
			"Other Provider Identifier Issuer_5" varchar(80),
			"Other Provider Identifier_6" varchar(20),
			"Other Provider Identifier Type Code_6" varchar(2),
			"Other Provider Identifier State_6" varchar(2),
			"Other Provider Identifier Issuer_6" varchar(80),
			"Other Provider Identifier_7" varchar(20),
			"Other Provider Identifier Type Code_7" varchar(2),
			"Other Provider Identifier State_7" varchar(2),
			"Other Provider Identifier Issuer_7" varchar(80),
			"Other Provider Identifier_8" varchar(20),
			"Other Provider Identifier Type Code_8" varchar(2),
			"Other Provider Identifier State_8" varchar(2),
			"Other Provider Identifier Issuer_8" varchar(80),
			"Other Provider Identifier_9" varchar(20),
			"Other Provider Identifier Type Code_9" varchar(2),
			"Other Provider Identifier State_9" varchar(2),
			"Other Provider Identifier Issuer_9" varchar(80),
			"Other Provider Identifier_10" varchar(20),
			"Other Provider Identifier Type Code_10" varchar(2),
			"Other Provider Identifier State_10" varchar(2),
			"Other Provider Identifier Issuer_10" varchar(80),
			"Other Provider Identifier_11" varchar(20),
			"Other Provider Identifier Type Code_11" varchar(2),
			"Other Provider Identifier State_11" varchar(2),
			"Other Provider Identifier Issuer_11" varchar(80),
			"Other Provider Identifier_12" varchar(20),
			"Other Provider Identifier Type Code_12" varchar(2),
			"Other Provider Identifier State_12" varchar(2),
			"Other Provider Identifier Issuer_12" varchar(80),
			"Other Provider Identifier_13" varchar(20),
			"Other Provider Identifier Type Code_13" varchar(2),
			"Other Provider Identifier State_13" varchar(2),
			"Other Provider Identifier Issuer_13" varchar(80),
			"Other Provider Identifier_14" varchar(20),
			"Other Provider Identifier Type Code_14" varchar(2),
			"Other Provider Identifier State_14" varchar(2),
			"Other Provider Identifier Issuer_14" varchar(80),
			"Other Provider Identifier_15" varchar(20),
			"Other Provider Identifier Type Code_15" varchar(2),
			"Other Provider Identifier State_15" varchar(2),
			"Other Provider Identifier Issuer_15" varchar(80),
			"Other Provider Identifier_16" varchar(20),
			"Other Provider Identifier Type Code_16" varchar(2),
			"Other Provider Identifier State_16" varchar(2),
			"Other Provider Identifier Issuer_16" varchar(80),
			"Other Provider Identifier_17" varchar(20),
			"Other Provider Identifier Type Code_17" varchar(2),
			"Other Provider Identifier State_17" varchar(2),
			"Other Provider Identifier Issuer_17" varchar(80),
			"Other Provider Identifier_18" varchar(20),
			"Other Provider Identifier Type Code_18" varchar(2),
			"Other Provider Identifier State_18" varchar(2),
			"Other Provider Identifier Issuer_18" varchar(80),
			"Other Provider Identifier_19" varchar(20),
			"Other Provider Identifier Type Code_19" varchar(2),
			"Other Provider Identifier State_19" varchar(2),
			"Other Provider Identifier Issuer_19" varchar(80),
			"Other Provider Identifier_20" varchar(20),
			"Other Provider Identifier Type Code_20" varchar(2),
			"Other Provider Identifier State_20" varchar(2),
			"Other Provider Identifier Issuer_20" varchar(80),
			"Other Provider Identifier_21" varchar(20),
			"Other Provider Identifier Type Code_21" varchar(2),
			"Other Provider Identifier State_21" varchar(2),
			"Other Provider Identifier Issuer_21" varchar(80),
			"Other Provider Identifier_22" varchar(20),
			"Other Provider Identifier Type Code_22" varchar(2),
			"Other Provider Identifier State_22" varchar(2),
			"Other Provider Identifier Issuer_22" varchar(80),
			"Other Provider Identifier_23" varchar(20),
			"Other Provider Identifier Type Code_23" varchar(2),
			"Other Provider Identifier State_23" varchar(2),
			"Other Provider Identifier Issuer_23" varchar(80),
			"Other Provider Identifier_24" varchar(20),
			"Other Provider Identifier Type Code_24" varchar(2),
			"Other Provider Identifier State_24" varchar(2),
			"Other Provider Identifier Issuer_24" varchar(80),
			"Other Provider Identifier_25" varchar(20),
			"Other Provider Identifier Type Code_25" varchar(2),
			"Other Provider Identifier State_25" varchar(2),
			"Other Provider Identifier Issuer_25" varchar(80),
			"Other Provider Identifier_26" varchar(20),
			"Other Provider Identifier Type Code_26" varchar(2),
			"Other Provider Identifier State_26" varchar(2),
			"Other Provider Identifier Issuer_26" varchar(80),
			"Other Provider Identifier_27" varchar(20),
			"Other Provider Identifier Type Code_27" varchar(2),
			"Other Provider Identifier State_27" varchar(2),
			"Other Provider Identifier Issuer_27" varchar(80),
			"Other Provider Identifier_28" varchar(20),
			"Other Provider Identifier Type Code_28" varchar(2),
			"Other Provider Identifier State_28" varchar(2),
			"Other Provider Identifier Issuer_28" varchar(80),
			"Other Provider Identifier_29" varchar(20),
			"Other Provider Identifier Type Code_29" varchar(2),
			"Other Provider Identifier State_29" varchar(2),
			"Other Provider Identifier Issuer_29" varchar(80),
			"Other Provider Identifier_30" varchar(20),
			"Other Provider Identifier Type Code_30" varchar(2),
			"Other Provider Identifier State_30" varchar(2),
			"Other Provider Identifier Issuer_30" varchar(80),
			"Other Provider Identifier_31" varchar(20),
			"Other Provider Identifier Type Code_31" varchar(2),
			"Other Provider Identifier State_31" varchar(2),
			"Other Provider Identifier Issuer_31" varchar(80),
			"Other Provider Identifier_32" varchar(20),
			"Other Provider Identifier Type Code_32" varchar(2),
			"Other Provider Identifier State_32" varchar(2),
			"Other Provider Identifier Issuer_32" varchar(80),
			"Other Provider Identifier_33" varchar(20),
			"Other Provider Identifier Type Code_33" varchar(2),
			"Other Provider Identifier State_33" varchar(2),
			"Other Provider Identifier Issuer_33" varchar(80),
			"Other Provider Identifier_34" varchar(20),
			"Other Provider Identifier Type Code_34" varchar(2),
			"Other Provider Identifier State_34" varchar(2),
			"Other Provider Identifier Issuer_34" varchar(80),
			"Other Provider Identifier_35" varchar(20),
			"Other Provider Identifier Type Code_35" varchar(2),
			"Other Provider Identifier State_35" varchar(2),
			"Other Provider Identifier Issuer_35" varchar(80),
			"Other Provider Identifier_36" varchar(20),
			"Other Provider Identifier Type Code_36" varchar(2),
			"Other Provider Identifier State_36" varchar(2),
			"Other Provider Identifier Issuer_36" varchar(80),
			"Other Provider Identifier_37" varchar(20),
			"Other Provider Identifier Type Code_37" varchar(2),
			"Other Provider Identifier State_37" varchar(2),
			"Other Provider Identifier Issuer_37" varchar(80),
			"Other Provider Identifier_38" varchar(20),
			"Other Provider Identifier Type Code_38" varchar(2),
			"Other Provider Identifier State_38" varchar(2),
			"Other Provider Identifier Issuer_38" varchar(80),
			"Other Provider Identifier_39" varchar(20),
			"Other Provider Identifier Type Code_39" varchar(2),
			"Other Provider Identifier State_39" varchar(2),
			"Other Provider Identifier Issuer_39" varchar(80),
			"Other Provider Identifier_40" varchar(20),
			"Other Provider Identifier Type Code_40" varchar(2),
			"Other Provider Identifier State_40" varchar(2),
			"Other Provider Identifier Issuer_40" varchar(80),
			"Other Provider Identifier_41" varchar(20),
			"Other Provider Identifier Type Code_41" varchar(2),
			"Other Provider Identifier State_41" varchar(2),
			"Other Provider Identifier Issuer_41" varchar(80),
			"Other Provider Identifier_42" varchar(20),
			"Other Provider Identifier Type Code_42" varchar(2),
			"Other Provider Identifier State_42" varchar(2),
			"Other Provider Identifier Issuer_42" varchar(80),
			"Other Provider Identifier_43" varchar(20),
			"Other Provider Identifier Type Code_43" varchar(2),
			"Other Provider Identifier State_43" varchar(2),
			"Other Provider Identifier Issuer_43" varchar(80),
			"Other Provider Identifier_44" varchar(20),
			"Other Provider Identifier Type Code_44" varchar(2),
			"Other Provider Identifier State_44" varchar(2),
			"Other Provider Identifier Issuer_44" varchar(80),
			"Other Provider Identifier_45" varchar(20),
			"Other Provider Identifier Type Code_45" varchar(2),
			"Other Provider Identifier State_45" varchar(2),
			"Other Provider Identifier Issuer_45" varchar(80),
			"Other Provider Identifier_46" varchar(20),
			"Other Provider Identifier Type Code_46" varchar(2),
			"Other Provider Identifier State_46" varchar(2),
			"Other Provider Identifier Issuer_46" varchar(80),
			"Other Provider Identifier_47" varchar(20),
			"Other Provider Identifier Type Code_47" varchar(2),
			"Other Provider Identifier State_47" varchar(2),
			"Other Provider Identifier Issuer_47" varchar(80),
			"Other Provider Identifier_48" varchar(20),
			"Other Provider Identifier Type Code_48" varchar(2),
			"Other Provider Identifier State_48" varchar(2),
			"Other Provider Identifier Issuer_48" varchar(80),
			"Other Provider Identifier_49" varchar(20),
			"Other Provider Identifier Type Code_49" varchar(2),
			"Other Provider Identifier State_49" varchar(2),
			"Other Provider Identifier Issuer_49" varchar(80),
			"Other Provider Identifier_50" varchar(20),
			"Other Provider Identifier Type Code_50" varchar(2),
			"Other Provider Identifier State_50" varchar(2),
			"Other Provider Identifier Issuer_50" varchar(80),
			"Is Sole Proprietor" varchar(1),
			"Is Organization Subpart" varchar(1),
			"Parent Organization LBN" varchar(70),
			"Parent Organization TIN" varchar(9),
			"Authorized Official Name Prefix Text" varchar(5),
			"Authorized Official Name Suffix Text" varchar(5),
			"Authorized Official Credential Text" varchar(20),
			"Healthcare Provider Taxonomy Group_1" varchar(70),
			"Healthcare Provider Taxonomy Group_2" varchar(70),
			"Healthcare Provider Taxonomy Group_3" varchar(70),
			"Healthcare Provider Taxonomy Group_4" varchar(70),
			"Healthcare Provider Taxonomy Group_5" varchar(70),
			"Healthcare Provider Taxonomy Group_6" varchar(70),
			"Healthcare Provider Taxonomy Group_7" varchar(70),
			"Healthcare Provider Taxonomy Group_8" varchar(70),
			"Healthcare Provider Taxonomy Group_9" varchar(70),
			"Healthcare Provider Taxonomy Group_10" varchar(70),
			"Healthcare Provider Taxonomy Group_11" varchar(70),
			"Healthcare Provider Taxonomy Group_12" varchar(70),
			"Healthcare Provider Taxonomy Group_13" varchar(70),
			"Healthcare Provider Taxonomy Group_14" varchar(70),
			"Healthcare Provider Taxonomy Group_15" varchar(70),
			CONSTRAINT %[1]s_PK PRIMARY KEY ("NPI")
		)`, tableName)
}

func updateTableSchema(prefix string) string {
	return fmt.Sprintf(`
		CREATE TABLE %s_Update (
			end_date date NOT NULL,
			CONSTRAINT %[1]s_Update_PK PRIMARY KEY (end_date)
		)`, prefix)
}
