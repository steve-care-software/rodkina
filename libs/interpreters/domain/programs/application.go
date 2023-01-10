package programs

import (
	"github.com/steve-care-software/rodkina/libs/interpreters/domain/programs/modules"
)

type application struct {
	index       uint
	module      modules.Module
	attachments Attachments
}

func createApplication(
	index uint,
	module modules.Module,
) Application {
	return createApplicationInternally(index, module, nil)
}

func createApplicationWithAttachments(
	index uint,
	module modules.Module,
	attachments Attachments,
) Application {
	return createApplicationInternally(index, module, attachments)
}

func createApplicationInternally(
	index uint,
	module modules.Module,
	attachments Attachments,
) Application {
	out := application{
		index:       index,
		module:      module,
		attachments: attachments,
	}

	return &out
}

// Index returns the index
func (obj *application) Index() uint {
	return obj.index
}

// Module returns the module
func (obj *application) Module() modules.Module {
	return obj.module
}

// HasAttachments returns true if there is attachments, false otherwise
func (obj *application) HasAttachments() bool {
	return obj.attachments != nil
}

// Attachments returns the attachments, if any
func (obj *application) Attachments() Attachments {
	return obj.attachments
}
