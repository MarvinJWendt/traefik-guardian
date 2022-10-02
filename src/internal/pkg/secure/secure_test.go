package secure

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestHashes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		hashResolver HashResolver
		password     string
		hash         string
		shouldMatch  bool
	}{
		{hashResolver: &BcryptResolver{}, password: "Hello, World!", hash: "$2a$10$k8fBIpJInrE70BzYy5rO/OUSt1w2.IX0bWhiMdb2mJEhjheVHDhvK", shouldMatch: true},
		{hashResolver: &BcryptResolver{}, password: "Hello, World!", hash: "$2a$10$X8fBIpJonrE70BzYy5rO/OUSt1w2.IX0bWhiMdb2mJEhjheVHDhvK", shouldMatch: false},
		{hashResolver: &BcryptResolver{}, password: "Hello, World!", hash: "", shouldMatch: false},
		{hashResolver: &BcryptResolver{}, password: "", hash: "", shouldMatch: false},
		{hashResolver: &MD5Resolver{}, password: "Hello, World!", hash: "65a8e27d8879283831b664bd8b7f0ad4", shouldMatch: true},
		{hashResolver: &MD5Resolver{}, password: "Hello, World!", hash: "X5a8e27d8879283831b664bd8b7f0ad4", shouldMatch: false},
		{hashResolver: &MD5Resolver{}, password: "Hello, World!", hash: "", shouldMatch: false},
		{hashResolver: &MD5Resolver{}, password: "", hash: "", shouldMatch: false},
		{hashResolver: &PlaintextResolver{}, password: "Hello, World!", hash: "Hello, World!", shouldMatch: true},
		{hashResolver: &PlaintextResolver{}, password: "Hello, World!", hash: "Xello, World!", shouldMatch: false},
		{hashResolver: &PlaintextResolver{}, password: "Hello, World!", hash: "", shouldMatch: false},
		{hashResolver: &PlaintextResolver{}, password: "", hash: "", shouldMatch: true},
		{hashResolver: &SHA512Resolver{}, password: "Hello, World!", hash: "374d794a95cdcfd8b35993185fef9ba368f160d8daf432d08ba9f1ed1e5abe6cc69291e0fa2fe0006a52570ef18c19def4e617c33ce52ef0a6e5fbe318cb0387", shouldMatch: true},
		{hashResolver: &SHA512Resolver{}, password: "Hello, World!", hash: "X74d794a95cdcfd8b35993185fef9ba368f160d8daf432d08ba9f1ed1e5abe6cc69291e0fa2fe0006a52570ef18c19def4e617c33ce52ef0a6e5fbe318cb0387", shouldMatch: false},
		{hashResolver: &SHA512Resolver{}, password: "Hello, World!", hash: "", shouldMatch: false},
		{hashResolver: &SHA512Resolver{}, password: "", hash: "", shouldMatch: false},
	}
	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			t.Parallel()
			b := tt.hashResolver
			testza.AssertEqual(t, b.Check(tt.hash, tt.password), tt.shouldMatch, "Test: %#v", tt)
		})
	}
}
