package contentful

type Includes struct {
	Entry	[]*Entry				`json:"Entry,omitempty"`
	Asset	[]*Asset				`json:"Asset,omitempty"`
}
