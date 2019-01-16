package main

import (
	"fmt"
	"log"

	"golang.org/x/sys/windows"

	acl "github.com/hectane/go-acl"
	"github.com/taskcluster/generic-worker/process"
)

func (cot *ChainOfTrustTaskFeature) ensureTaskUserCantReadPrivateCotKey() error {
	accessToken := cot.task.PlatformData.CommandAccessToken
	signingKeyPaths := [2]string{
		config.OpenPGPSigningKeyLocation,
		config.Ed25519SigningKeyLocation,
	}
	for _, path := range signingKeyPaths {
		c, err := process.NewCommand([]string{"cmd.exe", "/c", "type", path}, cwd, nil, accessToken)
		if err != nil {
			panic(fmt.Errorf("SERIOUS BUG: Could not create command (not even trying to execute it yet) to cat private chain of trust key - %v", err))
		}
		r := c.Execute()
		if !r.Failed() {
			log.Print(r.String())
			return fmt.Errorf(ChainOfTrustKeyNotSecureMessage)
		}
	}
	return nil
}

// Ensure only administrators have access permissions for the chain of trust
// private signing key file, and grant them full control.
func secureSigningKey() (err error) {
	signingKeyPaths := [2]string{
		config.OpenPGPSigningKeyLocation,
		config.Ed25519SigningKeyLocation,
	}
	for _, path := range signingKeyPaths {
		err = acl.Apply(

			// Private signing key file
			path,
			// delete existing permissions (ACLs)
			true,
			// don't inherit permissions (ACLs)
			false,
			// grant Administrators group full control
			acl.GrantName(windows.GENERIC_ALL, "Administrators"),
		)
	}
	return
}
