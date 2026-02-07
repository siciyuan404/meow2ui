package recovery

import "fmt"

func Restore(backupID string) error {
	if backupID == "" {
		return fmt.Errorf("backup id is required")
	}
	return nil
}

func ValidatePostRestore() error {
	return nil
}
