package cli

import (
	"encoding/json"
	"fmt"
	"github.com/joshinjohnson/authentication-service/internal/access"
	"github.com/joshinjohnson/authentication-service/pkg/api"
	"github.com/joshinjohnson/authentication-service/pkg/models"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	mode models.Mode
	modeStr    string
	typeStr 	string
	inputFilePath string
	cred models.UserCredential
	regDetails registrationDetails
	authAccess *access.AuthenticationAccess

	rootCmd    = &cobra.Command{
		Use:     "authentication-engine",
		Version: "0.0.1",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			b, err := ioutil.ReadFile(inputFilePath)
			if err != nil {
				fmt.Errorf(err.Error())
				os.Exit(1)
			}
			switch typeStr {
			case "authenticate":
				if err := json.Unmarshal(b, &cred); err != nil {
					fmt.Errorf(err.Error())
					os.Exit(1)
				}
			case "register":
				if err := json.Unmarshal(b, &regDetails); err != nil {
					fmt.Errorf(err.Error())
					os.Exit(1)
				}
			default:
				fmt.Errorf("invalid operation type selected")
				os.Exit(1)
			}
			mode = getMode(modeStr)
			authAccess = access.NewAuthenticationAccess()
		},
		Run: func(cmd *cobra.Command, args []string) {
			engine, err := api.New(models.Config{Mode: mode})
			if err != nil {
				fmt.Errorf(err.Error())
				os.Exit(1)
			}
			switch typeStr {
			case "authenticate":
				_, err := engine.Authenticate(cred)
				if err != nil {
					fmt.Errorf(err.Error())
					os.Exit(1)
				}
				fmt.Println("authenticated!")
			case "register":
				if err := engine.Register(regDetails.creds, regDetails.details); err != nil {
					fmt.Errorf(err.Error())
					os.Exit(1)
				}
				fmt.Println("registered!")
			}
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&modeStr, "mode", "m", "1", "modes available: 1 (emulation), 2 (operation)")
	rootCmd.Flags().StringVarP(&typeStr, "operation-type", "t", "authenticate", "operation type available: authenticate, register")
	rootCmd.Flags().StringVarP(&inputFilePath, "input-file-path", "f", "data/authenticate.json", "file path for operation type")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getMode(str string) models.Mode {
	if i, _ := strconv.Atoi(str); i == 2 {
		return models.Operation
	}
	return models.Emulation
}

type registrationDetails struct {
	creds models.UserCredential `json:"user-credentials"`
	details models.UserDetails `json:"user-details"`
}