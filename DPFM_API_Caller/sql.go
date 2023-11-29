package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-address-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-address-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) AddressRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Address {
	where := fmt.Sprintf("WHERE address.AddressID = \"%s\"", input.Address.AddressID)
	where := fmt.Sprintf("WHERE address.ValidityStartDate = \"%s\"", input.Address.ValidityStartDate)
	where := fmt.Sprintf("WHERE address.ValidityEndDate = \"%s\"", input.Address.ValidityEndDate)
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_address_address_data as address 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
