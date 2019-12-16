package lib_gc_cache_helpers

// Cache Object metadata fields
const METADATA_VERSION = "VERSION"
const METADATA_STATUS = "STATUS"

const (
	CACHEOBJECTSTATUS_ENABLED  int16 = 0
	CACHEOBJECTSTATUS_DISABLED int16 = 1
	CACHEOBJECTSTATUS_DELETED  int16 = 2
)
