package present

import (
	"github.com/concourse/concourse/v7/atc"
	"github.com/concourse/concourse/v7/atc/db"
)

func ResourceTypes(savedResourceTypes db.ResourceTypes) atc.ResourceTypes {
	return savedResourceTypes.Deserialize()
}
