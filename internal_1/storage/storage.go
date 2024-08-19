package storage

import "os"

func SaveToFile(record string) error {
	file, err := os.OpenFile("data/records.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(record)
	if err != nil {
		return err
	}

	return nil
}
