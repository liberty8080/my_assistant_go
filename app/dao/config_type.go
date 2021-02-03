// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"my_assistant_go/app/dao/internal"
)

// configTypeDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type configTypeDao struct {
	*internal.ConfigTypeDao
}

var (
	// ConfigType is globally public accessible object for table config_type operations.
	ConfigType = &configTypeDao{
		internal.ConfigType,
	}
)

// Fill with you ideas below.