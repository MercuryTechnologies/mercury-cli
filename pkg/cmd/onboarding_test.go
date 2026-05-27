// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
)

func TestOnboardingSubmit1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"onboarding", "submit1",
			"--beneficial-owner", "{address1: address1, address2: address2, citizenshipStatus: USCitizen, city: city, country: country, dateOfBirth: '2016-07-22', email: email, firstName: firstName, identificationBlob: identificationBlob, identificationType: Passport, isPep: IsPep, jobTitle: ChiefExecutiveOfficer, lastName: lastName, otherJobTitle: otherJobTitle, percentOwnership: 0, phoneNumber: phoneNumber, postalCode: postalCode, region: region, socialProfileLinks: [string], state: state}",
			"--partner", "partner",
			"--about", "{countriesOfOperations: [string], countryOfOperation: countryOfOperation, description: description, industry: industry, legalBusinessName: legalBusinessName, website: website}",
			"--application-type", "PendingEINApplication",
			"--business-contact-details", "{address1: address1, address2: address2, city: city, country: country, phoneNumber: phoneNumber, postalCode: postalCode, state: state}",
			"--business-legal-address", "{address1: address1, address2: address2, city: city, country: country, postalCode: postalCode, region: region}",
			"--business-physical-address", "{address1: address1, address2: address2, city: city, country: country, postalCode: postalCode, region: region}",
			"--formation-details", "{federalEin: 12-3456789, formationDocumentFileBlob: 12-3456789, companyOriginCountry: companyOriginCountry, companyStructure: CCorp, einDocumentFileBlob: einDocumentFileBlob, eINDocumentFileBlob: eINDocumentFileBlob, foreignBusinessNumber: foreignBusinessNumber, formationDocumentType: ArticlesOfIncorporation}",
			"--invite-email", "inviteEmail",
			"--webhook-url", "webhookURL",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(onboardingSubmit1)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"onboarding", "submit1",
			"--beneficial-owner.address1", "address1",
			"--beneficial-owner.address2", "address2",
			"--beneficial-owner.citizenship-status", "USCitizen",
			"--beneficial-owner.city", "city",
			"--beneficial-owner.country", "country",
			"--beneficial-owner.date-of-birth", "2016-07-22",
			"--beneficial-owner.email", "email",
			"--beneficial-owner.first-name", "firstName",
			"--beneficial-owner.identification-blob", "identificationBlob",
			"--beneficial-owner.identification-type", "Passport",
			"--beneficial-owner.is-pep", "IsPep",
			"--beneficial-owner.job-title", "ChiefExecutiveOfficer",
			"--beneficial-owner.last-name", "lastName",
			"--beneficial-owner.other-job-title", "otherJobTitle",
			"--beneficial-owner.percent-ownership", "0",
			"--beneficial-owner.phone-number", "phoneNumber",
			"--beneficial-owner.postal-code", "postalCode",
			"--beneficial-owner.region", "region",
			"--beneficial-owner.social-profile-links", "[string]",
			"--beneficial-owner.state", "state",
			"--partner", "partner",
			"--about.countries-of-operations", "[string]",
			"--about.country-of-operation", "countryOfOperation",
			"--about.description", "description",
			"--about.industry", "industry",
			"--about.legal-business-name", "legalBusinessName",
			"--about.website", "website",
			"--application-type", "PendingEINApplication",
			"--business-contact-details.address1", "address1",
			"--business-contact-details.address2", "address2",
			"--business-contact-details.city", "city",
			"--business-contact-details.country", "country",
			"--business-contact-details.phone-number", "phoneNumber",
			"--business-contact-details.postal-code", "postalCode",
			"--business-contact-details.state", "state",
			"--business-legal-address.address1", "address1",
			"--business-legal-address.address2", "address2",
			"--business-legal-address.city", "city",
			"--business-legal-address.country", "country",
			"--business-legal-address.postal-code", "postalCode",
			"--business-legal-address.region", "region",
			"--business-physical-address.address1", "address1",
			"--business-physical-address.address2", "address2",
			"--business-physical-address.city", "city",
			"--business-physical-address.country", "country",
			"--business-physical-address.postal-code", "postalCode",
			"--business-physical-address.region", "region",
			"--formation-details.federal-ein", "12-3456789",
			"--formation-details.formation-document-file-blob", "12-3456789",
			"--formation-details.company-origin-country", "companyOriginCountry",
			"--formation-details.company-structure", "CCorp",
			"--formation-details.ein-document-file-blob", "einDocumentFileBlob",
			"--formation-details.e-in-document-file-blob", "eINDocumentFileBlob",
			"--formation-details.foreign-business-number", "foreignBusinessNumber",
			"--formation-details.formation-document-type", "ArticlesOfIncorporation",
			"--invite-email", "inviteEmail",
			"--webhook-url", "webhookURL",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"beneficialOwners:\n" +
			"  - address1: address1\n" +
			"    address2: address2\n" +
			"    citizenshipStatus: USCitizen\n" +
			"    city: city\n" +
			"    country: country\n" +
			"    dateOfBirth: '2016-07-22'\n" +
			"    email: email\n" +
			"    firstName: firstName\n" +
			"    identificationBlob: identificationBlob\n" +
			"    identificationType: Passport\n" +
			"    isPep: IsPep\n" +
			"    jobTitle: ChiefExecutiveOfficer\n" +
			"    lastName: lastName\n" +
			"    otherJobTitle: otherJobTitle\n" +
			"    percentOwnership: 0\n" +
			"    phoneNumber: phoneNumber\n" +
			"    postalCode: postalCode\n" +
			"    region: region\n" +
			"    socialProfileLinks:\n" +
			"      - string\n" +
			"    state: state\n" +
			"partner: partner\n" +
			"about:\n" +
			"  countriesOfOperations:\n" +
			"    - string\n" +
			"  countryOfOperation: countryOfOperation\n" +
			"  description: description\n" +
			"  industry: industry\n" +
			"  legalBusinessName: legalBusinessName\n" +
			"  website: website\n" +
			"applicationType: PendingEINApplication\n" +
			"businessContactDetails:\n" +
			"  address1: address1\n" +
			"  address2: address2\n" +
			"  city: city\n" +
			"  country: country\n" +
			"  phoneNumber: phoneNumber\n" +
			"  postalCode: postalCode\n" +
			"  state: state\n" +
			"businessLegalAddress:\n" +
			"  address1: address1\n" +
			"  address2: address2\n" +
			"  city: city\n" +
			"  country: country\n" +
			"  postalCode: postalCode\n" +
			"  region: region\n" +
			"businessPhysicalAddress:\n" +
			"  address1: address1\n" +
			"  address2: address2\n" +
			"  city: city\n" +
			"  country: country\n" +
			"  postalCode: postalCode\n" +
			"  region: region\n" +
			"formationDetails:\n" +
			"  federalEin: 12-3456789\n" +
			"  formationDocumentFileBlob: 12-3456789\n" +
			"  companyOriginCountry: companyOriginCountry\n" +
			"  companyStructure: CCorp\n" +
			"  einDocumentFileBlob: einDocumentFileBlob\n" +
			"  eINDocumentFileBlob: eINDocumentFileBlob\n" +
			"  foreignBusinessNumber: foreignBusinessNumber\n" +
			"  formationDocumentType: ArticlesOfIncorporation\n" +
			"inviteEmail: inviteEmail\n" +
			"webhookURL: webhookURL\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"onboarding", "submit1",
		)
	})
}
