package internal

import "io"

func SendToServer(msg string, writer io.Writer) error {

	_, err := writer.Write([]byte(msg))
	if err != nil {
		return err
	}

	return nil
}
