// Package carddav provides a client and server CardDAV implementation.
//
// CardDAV is defined in RFC 6352.
package carddav

import (
	"time"

	"github.com/emersion/go-vcard"
)

// AddressDataType describes an address data type.
type AddressDataType struct {
	ContentType string
	Version     string
}

// AddressBook describes an address book.
type AddressBook struct {
	Path                 string
	Name                 string
	Description          string
	MaxResourceSize      int64
	SupportedAddressData []AddressDataType
	HomeSet              string
	PrincipalURL         string
	CurrentUserPrincipal string
}

// SupportsAddressData checks if the address book supports a given address data type.
func (ab *AddressBook) SupportsAddressData(contentType, version string) bool {
	if len(ab.SupportedAddressData) == 0 {
		return contentType == "text/vcard" && version == "3.0"
	}
	for _, t := range ab.SupportedAddressData {
		if t.ContentType == contentType && t.Version == version {
			return true
		}
	}
	return false
}

// AddressBookQuery describes a query on an address book.
type AddressBookQuery struct {
	DataRequest AddressDataRequest

	PropFilters []PropFilter
	FilterTest  FilterTest // defaults to FilterAnyOf

	Limit int // <= 0 means unlimited
}

// AddressDataRequest describes a request on address data.
type AddressDataRequest struct {
	Props   []string
	AllProp bool
}

// PropFilter describes a filter on props.
type PropFilter struct {
	Name string
	Test FilterTest // defaults to FilterAnyOf

	// if IsNotDefined is set, TextMatches and Params need to be unset
	IsNotDefined bool
	TextMatches  []TextMatch
	Params       []ParamFilter
}

// ParamFilter is used to describe PropFilter.
type ParamFilter struct {
	Name string

	// if IsNotDefined is set, TextMatch needs to be unset
	IsNotDefined bool
	TextMatch    *TextMatch
}

// TextMatch is used in to describe ParamFilter.
type TextMatch struct {
	Text            string
	NegateCondition bool
	MatchType       MatchType // defaults to MatchContains
}

// FilterTest is an alias used to describe PropFilter
type FilterTest string

// FilterTest constants
const (
	FilterAnyOf FilterTest = "anyof"
	FilterAllOf FilterTest = "allof"
)

// MatchType is used to describe TextMatch.
type MatchType string

// MatchType constants
const (
	MatchEquals     MatchType = "equals"
	MatchContains   MatchType = "contains"
	MatchStartsWith MatchType = "starts-with"
	MatchEndsWith   MatchType = "ends-with"
)

// AddressBookMultiGet describes a multi-get request.
type AddressBookMultiGet struct {
	Paths       []string
	DataRequest AddressDataRequest
}

// AddressObject describe an address object.
type AddressObject struct {
	Path    string
	ModTime time.Time
	ETag    string
	Card    vcard.Card
}

//SyncQuery is the query struct represents a sync-collection request
type SyncQuery struct {
	DataRequest AddressDataRequest
	SyncToken   string
	Limit       int // <= 0 means unlimited
}

//SyncResponse contains the returned sync-token for next time
type SyncResponse struct {
	SyncToken string
	Updated   []AddressObject
	Deleted   []string
}
