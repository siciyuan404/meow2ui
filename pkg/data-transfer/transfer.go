package datatransfer

import "fmt"

type Bundle struct {
	Manifest string
	Payload  string
}

func ExportWorkspace(workspaceID string) (Bundle, error) {
	if workspaceID == "" {
		return Bundle{}, fmt.Errorf("workspace id is required")
	}
	return Bundle{Manifest: "manifest.json", Payload: "{}"}, nil
}

func ImportWorkspace(bundle Bundle) error {
	if bundle.Manifest == "" {
		return fmt.Errorf("manifest is required")
	}
	return nil
}
