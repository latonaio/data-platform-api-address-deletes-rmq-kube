package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToAddress(rows *sql.Rows) (*Address, error) {
	defer rows.Close()
	address := Address{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&address.AddressID,
			&address.ValidityStartDate,
			&address.ValidityEndDate,
			&address.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &address, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &address, nil
	}

	return &address, nil
}
