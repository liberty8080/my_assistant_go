// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

// Config is the golang structure for table config.
type Config struct {
	Id    int    `orm:"id,primary" json:"id"`    //
	Name  string `orm:"name"       json:"name"`  //
	Value string `orm:"value"      json:"value"` //
	Type  int    `orm:"type"       json:"type"`  //
}