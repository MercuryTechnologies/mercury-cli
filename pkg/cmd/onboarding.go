// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/MercuryTechnologies/mercury-cli/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	"github.com/MercuryTechnologies/mercury-go"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var onboardingSubmit = requestflag.WithInnerFlags(cli.Command{
	Name:    "submit",
	Usage:   "Submit onboarding data for applicants to pre-fill their Mercury application",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "beneficial-owner",
			Required: true,
			BodyPath: "beneficialOwners",
		},
		&requestflag.Flag[string]{
			Name:     "partner",
			Required: true,
			BodyPath: "partner",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "about",
			BodyPath: "about",
		},
		&requestflag.Flag[*string]{
			Name:     "application-type",
			Usage:    `Allowed values: "PendingEINApplication", "DefaultApplication".`,
			BodyPath: "applicationType",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "business-contact-details",
			BodyPath: "businessContactDetails",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "business-legal-address",
			BodyPath: "businessLegalAddress",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "business-physical-address",
			BodyPath: "businessPhysicalAddress",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "formation-details",
			BodyPath: "formationDetails",
		},
		&requestflag.Flag[*string]{
			Name:     "invite-email",
			BodyPath: "inviteEmail",
		},
		&requestflag.Flag[*string]{
			Name:     "webhook-url",
			BodyPath: "webhookURL",
		},
	},
	Action:          handleOnboardingSubmit,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"beneficial-owner": {
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.address1",
			Usage:      " Address line 1 of Beneficial Owner's address",
			InnerField: "address1",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.address2",
			Usage:      " Address line 2 of Beneficial Owner's address",
			InnerField: "address2",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.citizenship-status",
			Usage:      " Beneficial Owner's Citizenship Status",
			InnerField: "citizenshipStatus",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.city",
			Usage:      " City of Beneficial Owner's address",
			InnerField: "city",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.country",
			Usage:      " Country of Beneficial Owner's address",
			InnerField: "country",
		},
		&requestflag.InnerFlag[any]{
			Name:       "beneficial-owner.date-of-birth",
			Usage:      " Beneficial Owner's Date of Birth",
			InnerField: "dateOfBirth",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.email",
			Usage:      " Beneficial Owner's Email Address",
			InnerField: "email",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.first-name",
			Usage:      " Beneficial Owner's First Name",
			InnerField: "firstName",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.identification-blob",
			Usage:      " Beneficial Owner's Identification File",
			InnerField: "identificationBlob",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.identification-type",
			Usage:      " Beneficial Owner's Identification File Type",
			InnerField: "identificationType",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.is-pep",
			Usage:      " Beneficial Owner's pep status",
			InnerField: "isPep",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.job-title",
			Usage:      " Beneficial Owner's Job Title",
			InnerField: "jobTitle",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.last-name",
			Usage:      " Beneficial Owner's Last Name",
			InnerField: "lastName",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.other-job-title",
			Usage:      " Beneficial Owner's Alternate Job Title",
			InnerField: "otherJobTitle",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "beneficial-owner.percent-ownership",
			Usage:      " Beneficial Owner's Ownership Percentage",
			InnerField: "percentOwnership",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.phone-number",
			Usage:      " Beneficial Owner's Phone Number",
			InnerField: "phoneNumber",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.postal-code",
			Usage:      " Postal Code of Beneficial Owner's address",
			InnerField: "postalCode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.region",
			Usage:      " Region or State of Beneficial Owner's address",
			InnerField: "region",
		},
		&requestflag.InnerFlag[any]{
			Name:       "beneficial-owner.social-profile-links",
			Usage:      " Beneficial Owner's Social Profile Websites",
			InnerField: "socialProfileLinks",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "beneficial-owner.state",
			Usage:      " State or Region of Beneficial Owner's address (Deprecated)",
			InnerField: "state",
		},
	},
	"about": {
		&requestflag.InnerFlag[any]{
			Name:       "about.countries-of-operations",
			Usage:      " The countries where the company operates.",
			InnerField: "countriesOfOperations",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "about.country-of-operation",
			InnerField: "countryOfOperation",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "about.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "about.industry",
			InnerField: "industry",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "about.legal-business-name",
			InnerField: "legalBusinessName",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "about.website",
			InnerField: "website",
		},
	},
	"business-contact-details": {
		&requestflag.InnerFlag[*string]{
			Name:       "business-contact-details.address1",
			InnerField: "address1",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-contact-details.address2",
			InnerField: "address2",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-contact-details.city",
			InnerField: "city",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-contact-details.country",
			InnerField: "country",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-contact-details.phone-number",
			InnerField: "phoneNumber",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-contact-details.postal-code",
			InnerField: "postalCode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-contact-details.state",
			InnerField: "state",
		},
	},
	"business-legal-address": {
		&requestflag.InnerFlag[*string]{
			Name:       "business-legal-address.address1",
			InnerField: "address1",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-legal-address.address2",
			InnerField: "address2",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-legal-address.city",
			InnerField: "city",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-legal-address.country",
			InnerField: "country",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-legal-address.postal-code",
			InnerField: "postalCode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-legal-address.region",
			InnerField: "region",
		},
	},
	"business-physical-address": {
		&requestflag.InnerFlag[*string]{
			Name:       "business-physical-address.address1",
			InnerField: "address1",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-physical-address.address2",
			InnerField: "address2",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-physical-address.city",
			InnerField: "city",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-physical-address.country",
			InnerField: "country",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-physical-address.postal-code",
			InnerField: "postalCode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-physical-address.region",
			InnerField: "region",
		},
	},
	"formation-details": {
		&requestflag.InnerFlag[*string]{
			Name:       "formation-details.federal-ein",
			Usage:      "Field should be null (no value), 'Pending' (value will be provided at a later date), or a valid value",
			InnerField: "federalEin",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "formation-details.formation-document-file-blob",
			Usage:      "Field should be null (no value), 'Pending' (value will be provided at a later date), or a valid value",
			InnerField: "formationDocumentFileBlob",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "formation-details.company-origin-country",
			InnerField: "companyOriginCountry",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "formation-details.company-structure",
			Usage:      `Allowed values: "CCorp", "LLC", "LLP", "NonProfit", "Partnership", "ProfessionalAssociation", "ProfessionalCorporation", "SCorp", "GeneralPartnership", "LimitedPartnership", "JointVenture", "LLCTaxedAsSoleProprietorship", "SoleProprietorship", "ExemptedCompany", "Limited".`,
			InnerField: "companyStructure",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "formation-details.ein-document-file-blob",
			InnerField: "einDocumentFileBlob",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "formation-details.e-in-document-file-blob",
			InnerField: "eINDocumentFileBlob",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "formation-details.foreign-business-number",
			InnerField: "foreignBusinessNumber",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "formation-details.formation-document-type",
			Usage:      `Allowed values: "ArticlesOfIncorporation", "ArticlesOfOrganization", "CertificateOfFormation", "PartnershipAgreement", "SecretaryOfStateRegistrationPage".`,
			InnerField: "formationDocumentType",
		},
	},
})

func handleOnboardingSubmit(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := mercury.OnboardingSubmitParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Onboarding.Submit(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "onboarding submit",
		Transform:      transform,
	})
}
