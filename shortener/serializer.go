package shortener

type Serializer interface {
	Encode(redirect Redirect) ([]byte, error)
	Decode(raw []byte) (Redirect, error)
}
