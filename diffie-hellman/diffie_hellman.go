package diffiehellman

import "math/big"
import "crypto/rand"

//PrivateKey generator
func PrivateKey(p *big.Int) *big.Int {
	limit := new(big.Int).Sub(p, big.NewInt(2))

	if n, err := rand.Int(rand.Reader, limit); err == nil {
		return n.Add(n, big.NewInt(2))
	}
	return nil
}

//PublicKey generator
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

//NewPair generator
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	newPrivate := PrivateKey(p)
	newPublic := PublicKey(newPrivate, p, g)
	return newPrivate, newPublic
}

//SecretKey generator
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
