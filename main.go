package main

import (
	ss "github.com/zalando/go-keyring/secret_service"
)

func main() {
	err := unlock()
	if err != nil {
		panic(err)
	}
}

func unlock() error {
	svc, err := ss.NewSecretService()
	if err != nil {
		return err
	}

	session, err := svc.OpenSession()
	if err != nil {
		return err
	}
	defer svc.Close(session)

	return svc.Unlock(svc.GetLoginCollection().Path())
}
