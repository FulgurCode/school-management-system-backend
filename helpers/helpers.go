package helpers

import "os/exec"

// Generating unique UUID
func GetUniqueId() string {
	// Execute the uuidgen command
	id, err := exec.Command("uuidgen").Output()
  if err != nil {
    panic(err)
  }

	return string(id)
}
