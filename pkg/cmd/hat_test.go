package cmd

import (
	"net/url"
	"strings"
	"testing"
)

func TestBuildShopifyURL(t *testing.T) {
	tests := []struct {
		name        string
		addr        *shippingAddress
		wantCountry string
		wantState   string
	}{
		{
			name: "US address",
			addr: &shippingAddress{
				Email:     "user@example.com",
				FirstName: "Ada",
				LastName:  "Lovelace",
				Address1:  "1 Market St",
				City:      "San Francisco",
				State:     "CA",
				Zip:       "94105",
				Country:   "US",
			},
			wantCountry: "US",
			wantState:   "CA",
		},
		{
			name: "Canadian address",
			addr: &shippingAddress{
				Email:     "user@example.com",
				FirstName: "Grace",
				LastName:  "Hopper",
				Address1:  "100 Queen St W",
				City:      "Toronto",
				State:     "Ontario",
				Zip:       "M5H 2N2",
				Country:   "CA",
			},
			wantCountry: "CA",
			wantState:   "Ontario",
		},
		{
			name: "missing country defaults to US",
			addr: &shippingAddress{
				Email:     "user@example.com",
				FirstName: "Alan",
				LastName:  "Turing",
				Address1:  "1 Bletchley Park",
				City:      "Milton Keynes",
				State:     "NY",
				Zip:       "10001",
			},
			wantCountry: "US",
			wantState:   "NY",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := buildShopifyURL(tt.addr)
			if !strings.HasPrefix(raw, hatCartURL+"?") {
				t.Fatalf("URL missing cart prefix: %s", raw)
			}

			u, err := url.Parse(raw)
			if err != nil {
				t.Fatalf("parse URL: %v", err)
			}
			q := u.Query()

			if got := q.Get("checkout[shipping_address][country]"); got != tt.wantCountry {
				t.Errorf("country = %q, want %q", got, tt.wantCountry)
			}
			if got := q.Get("checkout[shipping_address][province]"); got != tt.wantState {
				t.Errorf("province = %q, want %q", got, tt.wantState)
			}
			if got := q.Get("checkout[email]"); got != tt.addr.Email {
				t.Errorf("email = %q, want %q", got, tt.addr.Email)
			}
			if got := q.Get("checkout[shipping_address][zip]"); got != tt.addr.Zip {
				t.Errorf("zip = %q, want %q", got, tt.addr.Zip)
			}
			if got := q.Get("utm_source"); got != "mercury-cli" {
				t.Errorf("utm_source = %q, want mercury-cli", got)
			}
		})
	}
}

func TestBuildShopifyURLOmitsEmptyAddress2(t *testing.T) {
	raw := buildShopifyURL(&shippingAddress{Country: "US"})
	u, err := url.Parse(raw)
	if err != nil {
		t.Fatalf("parse URL: %v", err)
	}
	if _, ok := u.Query()["checkout[shipping_address][address2]"]; ok {
		t.Error("address2 should be omitted when empty")
	}
}
