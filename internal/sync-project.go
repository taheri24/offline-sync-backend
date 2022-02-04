package internal

import (
	"bytes"

	"taheri24.ir/publish-server/database"
)

type SyncProject struct {
	data *database.Project
}

func NewSyncProject(data *database.Project) *SyncProject {
	result := new(SyncProject)
	result.data = data
	return result
}
func (project *SyncProject) GeneratePatch() (string, bytes.Buffer) {
	panic("sdf")
}
