package users

import "github.com/skycoin/skywire-utilities/pkg/cipher"

// KeyPair is a pair of public and secret keys
type KeyPair struct {
    PK cipher.PubKey
    SK cipher.SecKey
}

type UserAccount struct {
    UserName     string `json:"username"`
    UserPassword string `json:"password"`
    PublicKey    string `json:"pubkey"`
    PrivateKey   string `json:"privkey"`
}

type Message struct {
    MessageBody       				string `json:"message"`
    MessageTime        				string `json:"time"`
    MessageReceiveFromUsername 		string `json:"from_username"`
    MessageReceiveFromPublickey     string `json:"from_publickey"`
    MessageSendToUsername		    string `json:"to_username"`
    MessageSendToPublickey			string `json:"to_publickey"`
}

type Contact struct {
    UserName string `json:"UserName"`
    PublicKey string `json:"PublicKey"`
}

        from_username TEXT,
        from_publickey TEXT,
        to_username TEXT,
        to_publickey TEXT,
        message TEXT,
        time TEXT


// GenKeyPairs generates n random key pairs
func GenKeyPairs(n int) []KeyPair {
    pairs := make([]KeyPair, n)
    for i := range pairs {
        pk, sk, err := cipher.GenerateDeterministicKeyPair([]byte{byte(i)})
        if err != nil {
            panic(err)
        }
        pairs[i] = KeyPair{PK: pk, SK: sk}
    }
    return pairs
}
