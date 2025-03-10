package plasticcredit

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/EmpowerPlastic/empowerchain/testutil/sample"
)

type validateTest struct {
	msgUnderTest  sdk.Msg
	expectedError error
}

func TestMsgUpdateParams_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgUpdateParams{
				Authority: sample.AccAddress(),
				Params: Params{
					IssuerCreator: sample.AccAddress(),
				},
			},
			expectedError: nil,
		},
		"invalid authority": {
			msgUnderTest: &MsgUpdateParams{
				Authority: "invalid",
				Params: Params{
					IssuerCreator: sample.AccAddress(),
				},
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"invalid params": {
			msgUnderTest: &MsgUpdateParams{
				Authority: sample.AccAddress(),
				Params: Params{
					IssuerCreator: "invalid",
				},
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()

			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgCreateIssuer_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgCreateIssuer{
				Creator:     sample.AccAddress(),
				Name:        "Empower",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       sample.AccAddress(),
			},
			expectedError: nil,
		},
		"invalid creator": {
			msgUnderTest: &MsgCreateIssuer{
				Creator:     "invalid",
				Name:        "Empower",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       sample.AccAddress(),
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"invalid admin": {
			msgUnderTest: &MsgCreateIssuer{
				Creator:     sample.AccAddress(),
				Name:        "Empower",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       "invalid",
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"empty name not allowed": {
			msgUnderTest: &MsgCreateIssuer{
				Creator:     sample.AccAddress(),
				Name:        "",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       sample.AccAddress(),
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"empty description allowed": {
			msgUnderTest: &MsgCreateIssuer{
				Creator:     sample.AccAddress(),
				Name:        "Empower",
				Description: "",
				Admin:       sample.AccAddress(),
			},
			expectedError: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()

			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgUpdateIssuer_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgUpdateIssuer{
				Updater:     sample.AccAddress(),
				IssuerId:    1,
				Name:        "Empower",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       sample.AccAddress(),
			},
			expectedError: nil,
		},
		"invalid updater": {
			msgUnderTest: &MsgUpdateIssuer{
				Updater:     "invalid",
				IssuerId:    1,
				Name:        "Empower",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       sample.AccAddress(),
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"invalid issuer id": {
			msgUnderTest: &MsgUpdateIssuer{
				Updater:     sample.AccAddress(),
				IssuerId:    0,
				Name:        "Empower",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       sample.AccAddress(),
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"invalid admin": {
			msgUnderTest: &MsgUpdateIssuer{
				Updater:     sample.AccAddress(),
				IssuerId:    1,
				Name:        "Empower",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       "invalid",
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"empty name not allowed": {
			msgUnderTest: &MsgUpdateIssuer{
				Updater:     sample.AccAddress(),
				IssuerId:    1,
				Name:        "",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       sample.AccAddress(),
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"empty description allowed": {
			msgUnderTest: &MsgUpdateIssuer{
				Updater:     sample.AccAddress(),
				IssuerId:    1,
				Name:        "Empower",
				Description: "",
				Admin:       sample.AccAddress(),
			},
			expectedError: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()

			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgCreateApplicant_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgCreateApplicant{
				Name:        "Empower",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       sample.AccAddress(),
			},
			expectedError: nil,
		},
		"invalid admin": {
			msgUnderTest: &MsgCreateApplicant{
				Name:        "Empower",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       "invalid",
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"empty name not allowed": {
			msgUnderTest: &MsgCreateApplicant{
				Name:        "",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       sample.AccAddress(),
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"empty description allowed": {
			msgUnderTest: &MsgCreateApplicant{
				Name:        "Empower",
				Description: "",
				Admin:       sample.AccAddress(),
			},
			expectedError: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()

			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgUpdateApplicant_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgUpdateApplicant{
				ApplicantId: 1,
				Name:        "Empower",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       sample.AccAddress(),
				Updater:     sample.AccAddress(),
			},
			expectedError: nil,
		},
		"invalid admin": {
			msgUnderTest: &MsgUpdateApplicant{
				ApplicantId: 1,
				Name:        "Empower",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       "invalid",
				Updater:     sample.AccAddress(),
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"empty name not allowed": {
			msgUnderTest: &MsgUpdateApplicant{
				ApplicantId: 1,
				Name:        "",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       sample.AccAddress(),
				Updater:     sample.AccAddress(),
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"empty description allowed": {
			msgUnderTest: &MsgUpdateApplicant{
				ApplicantId: 1,
				Name:        "Empower",
				Description: "",
				Admin:       sample.AccAddress(),
				Updater:     sample.AccAddress(),
			},
			expectedError: nil,
		},
		"invalid applicant id": {
			msgUnderTest: &MsgUpdateApplicant{
				ApplicantId: 0,
				Name:        "Empower",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       sample.AccAddress(),
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"invalid updater": {
			msgUnderTest: &MsgUpdateApplicant{
				Updater:     "invalid",
				ApplicantId: 1,
				Name:        "Empower",
				Description: "Empower is the first and coolest plastic credit issuer!",
				Admin:       sample.AccAddress(),
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()

			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgCreateCreditClass_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgCreateCreditClass{
				Creator:      sample.AccAddress(),
				Abbreviation: "PCRD",
				IssuerId:     1,
				Name:         "Empower Plastic Credits",
			},
			expectedError: nil,
		},
		"invalid creator": {
			msgUnderTest: &MsgCreateCreditClass{
				Creator:      "hoppsasa",
				Abbreviation: "PCRD",
				IssuerId:     1,
				Name:         "Empower Plastic Credits",
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"empty abbreviation": {
			msgUnderTest: &MsgCreateCreditClass{
				Creator:      sample.AccAddress(),
				Abbreviation: "",
				IssuerId:     1,
				Name:         "Empower Plastic Credits",
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"empty issuer": {
			msgUnderTest: &MsgCreateCreditClass{
				Creator:      sample.AccAddress(),
				Abbreviation: "PCRD",
				IssuerId:     0,
				Name:         "Empower Plastic Credits",
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"empty name": {
			msgUnderTest: &MsgCreateCreditClass{
				Creator:      sample.AccAddress(),
				Abbreviation: "PCRD",
				IssuerId:     1,
				Name:         "",
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()
			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgUpdateCreditClass_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgUpdateCreditClass{
				Updater:      sample.AccAddress(),
				Abbreviation: "PCRD",
				Name:         "Empower Plastic Credits",
			},
			expectedError: nil,
		},
		"invalid creator": {
			msgUnderTest: &MsgUpdateCreditClass{
				Updater:      "hoppsasa",
				Abbreviation: "PCRD",
				Name:         "Empower Plastic Credits",
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"empty abbreviation": {
			msgUnderTest: &MsgUpdateCreditClass{
				Updater:      sample.AccAddress(),
				Abbreviation: "",
				Name:         "Empower Plastic Credits",
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"empty name": {
			msgUnderTest: &MsgUpdateCreditClass{
				Updater:      sample.AccAddress(),
				Abbreviation: "PCRD",
				Name:         "",
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()
			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgCreateProject_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgCreateProject{
				Creator:                 sample.AccAddress(),
				ApplicantId:             42,
				CreditClassAbbreviation: "PCRD",
				Name:                    "Project Name",
			},
			expectedError: nil,
		},
		"invalid creator": {
			msgUnderTest: &MsgCreateProject{
				Creator:                 "herpaderpa",
				ApplicantId:             42,
				CreditClassAbbreviation: "PCRD",
				Name:                    "Project Name",
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"invalid applicant id": {
			msgUnderTest: &MsgCreateProject{
				Creator:                 sample.AccAddress(),
				ApplicantId:             0,
				CreditClassAbbreviation: "PCRD",
				Name:                    "Project Name",
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"empty abbreviation": {
			msgUnderTest: &MsgCreateProject{
				Creator:                 sample.AccAddress(),
				ApplicantId:             42,
				CreditClassAbbreviation: "",
				Name:                    "Project Name",
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"empty name": {
			msgUnderTest: &MsgCreateProject{
				Creator:                 sample.AccAddress(),
				ApplicantId:             42,
				CreditClassAbbreviation: "PCRD",
				Name:                    "",
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()
			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgUpdateProject_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgUpdateProject{
				Updater:   sample.AccAddress(),
				ProjectId: 1,
				Name:      "Updated Project Name",
			},
			expectedError: nil,
		},
		"invalid updater": {
			msgUnderTest: &MsgUpdateProject{
				Updater:   "herpaderpa",
				ProjectId: 1,
				Name:      "Project Name",
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"invalid project id": {
			msgUnderTest: &MsgUpdateProject{
				Updater:   sample.AccAddress(),
				ProjectId: 0,
				Name:      "Project Name",
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"empty name": {
			msgUnderTest: &MsgUpdateProject{
				Updater:   sample.AccAddress(),
				ProjectId: 1,
				Name:      "",
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()
			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgApproveProject_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgApproveProject{
				Approver:  sample.AccAddress(),
				ProjectId: 1,
			},
			expectedError: nil,
		},
		"invalid approver": {
			msgUnderTest: &MsgApproveProject{
				Approver:  "invalid",
				ProjectId: 1,
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"invalid project id": {
			msgUnderTest: &MsgApproveProject{
				Approver:  sample.AccAddress(),
				ProjectId: 0,
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()
			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgRejectProject_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgRejectProject{
				Rejector:  sample.AccAddress(),
				ProjectId: 1,
			},
			expectedError: nil,
		},
		"invalid rejector": {
			msgUnderTest: &MsgRejectProject{
				Rejector:  "invalid",
				ProjectId: 1,
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"invalid project id": {
			msgUnderTest: &MsgRejectProject{
				Rejector:  sample.AccAddress(),
				ProjectId: 0,
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()
			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgSuspendProject_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgSuspendProject{
				Updater:   sample.AccAddress(),
				ProjectId: 1,
			},
			expectedError: nil,
		},
		"invalid updater": {
			msgUnderTest: &MsgSuspendProject{
				Updater:   "invalid",
				ProjectId: 1,
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"invalid project id": {
			msgUnderTest: &MsgSuspendProject{
				Updater:   sample.AccAddress(),
				ProjectId: 0,
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()
			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgIssueCredits_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgIssueCredits{
				Creator:      sample.AccAddress(),
				ProjectId:    1,
				SerialNumber: "123",
				CreditAmount: 10,
			},
			expectedError: nil,
		},
		"invalid creator": {
			msgUnderTest: &MsgIssueCredits{
				Creator:      "Empower",
				ProjectId:    1,
				SerialNumber: "123",
				CreditAmount: 10,
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"invalid project id": {
			msgUnderTest: &MsgIssueCredits{
				Creator:      sample.AccAddress(),
				SerialNumber: "123",
				CreditAmount: 10,
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"invalid serial number": {
			msgUnderTest: &MsgIssueCredits{
				Creator:      sample.AccAddress(),
				ProjectId:    1,
				CreditAmount: 10,
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"invalid credit amount": {
			msgUnderTest: &MsgIssueCredits{
				Creator:      sample.AccAddress(),
				ProjectId:    1,
				SerialNumber: "123",
				CreditAmount: 0,
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()
			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgTransferCredits(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgTransferCredits{
				From:   sample.AccAddress(),
				To:     sample.AccAddress(),
				Denom:  "EMP/123",
				Amount: 100,
				Retire: false,
			},
			expectedError: nil,
		},
		"invalid from": {
			msgUnderTest: &MsgTransferCredits{
				From:   "Empower",
				To:     sample.AccAddress(),
				Denom:  "EMP/123",
				Amount: 100,
				Retire: false,
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"invalid to": {
			msgUnderTest: &MsgTransferCredits{
				From:   sample.AccAddress(),
				To:     "Empower",
				Denom:  "EMP/123",
				Amount: 100,
				Retire: false,
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"invalid denom": {
			msgUnderTest: &MsgTransferCredits{
				From:   sample.AccAddress(),
				To:     sample.AccAddress(),
				Denom:  "",
				Amount: 100,
				Retire: false,
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"invalid amount": {
			msgUnderTest: &MsgTransferCredits{
				From:   sample.AccAddress(),
				To:     sample.AccAddress(),
				Denom:  "EMP/123",
				Amount: 0,
				Retire: false,
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()

			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgRetireCredits(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgRetireCredits{
				Owner:  sample.AccAddress(),
				Denom:  "EMP/123",
				Amount: 100,
			},
			expectedError: nil,
		},
		"invalid denom": {
			msgUnderTest: &MsgRetireCredits{
				Owner:  sample.AccAddress(),
				Denom:  "",
				Amount: 100,
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
		"invalid amount": {
			msgUnderTest: &MsgRetireCredits{
				Owner:  sample.AccAddress(),
				Denom:  "EMP/123",
				Amount: 0,
			},
			expectedError: sdkerrors.ErrInvalidRequest,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()

			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}
