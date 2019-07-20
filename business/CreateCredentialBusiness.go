package business

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	strfmt "github.com/go-openapi/strfmt"

	uid "github.com/segmentio/ksuid"

	"github.com/ccamaleon5/SaludChainClient/models"
	"encoding/hex"
)

//CreateCredential saving the hash into blockchain
func CreateCredential(subject *models.CredentialSubject, issuer string) (*models.Credential, string, error) {
	credential := new(models.Credential)
	id := generateID()
	credential.ID = &id
	types := make([]string, 0, 4)
	types = append(types, "VerifiableCredential")
	types = append(types, subject.Type)
	credential.Type = types
	credential.CredentialSubject = subject.Content
	credential.Evidence = subject.Evidence
	credential.Issuer = issuer
	credential.Proof = getProof("SmartContract", "saludchain")

	rawCredential, err := json.Marshal(credential)
	if err != nil {
		return nil, "",errors.New("Credential isn't Json format")
	}

	fmt.Println(string(rawCredential))
	fmt.Println("date:", subject.ExpirationDate.String())

	credentialHash, _ := generateHexaHash(sha256.Sum256(rawCredential))

	return credential, credentialHash, nil
}

func getProof(typeProof string, verificationMethod string) *models.Proof {
	var proof = new(models.Proof)
	proof.Type = typeProof
	proof.VerificationMethod = verificationMethod
	proof.Created = strfmt.DateTime(time.Now())

	return proof
}

func generateID() string {
	id := uid.New()
	return id.String()
}

func generateHexaHash(hashCredential [32]byte) (string, error) {
	var hash = make([]byte, 32, 64)

	for i, j := range hashCredential {
		hash[i] = j
	}

	fmt.Println("credential hex:", hex.EncodeToString(hash))
	return hex.EncodeToString(hash), nil
}
