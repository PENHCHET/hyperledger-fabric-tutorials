package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

const (
	PREFIX_ISS_AUTHINFO = "ISS_AUTHINFO"
	PREFIX_ISS_CERTINFO = "ISS_CERTINFO"
)

// IssAuthInfo asset
type IssAuthInfo struct {
	UUID 		string `json:"UUID"`
	CompanyIDNO string `json:"COMP_IDNO"`
	APIKey      string `json:"API_KEY"`
	BizCD       string `json:"BIZ_CD"`
	WebID       string `json:"WEB_ID"`
	WebPWD      string `json:"WEB_PW"`
}

// IssCertInfo asset
type IssCertInfo struct {
	UUID 		string `json:"UUID"`
	CompanyIDNO string `json:"COMP_IDNO"`
	APIKey      string `json:"API_KEY"`
	BizCD       string `json:"BIZ_CD"`
	CertOrg     string `json:"CERTORG"`
	CertName    string `json:"CERTNAME"`
	CertDate    string `json:"CERTDATE"`
	CertPWD     string `json:"CERTPWD"`
	CertPath    string `json:"CERTPATH"`
	CertData    string `json:"CERTDATA"`
	CertData1   string `json:"CERTDATA1"`
	CertData2   string `json:"CERTDATA2"`
	CertData3   string `json:"CERTDATA3"`
	CertData4   string `json:"CERTDATA4"`
	CertData5   string `json:"CERTDATA5"`
	CertData6   string `json:"CERTDATA6"`
	CertData7   string `json:"CERTDATA7"`
	CertData8   string `json:"CERTDATA8"`
	CertData9   string `json:"CERTDATA9"`
}


var logger = shim.NewLogger("main")

// SmartContract for building the Chaincode
type SmartContract struct {
}

/*
	IssAuthInfo
*/

/* -------------------------------------------------------------------------------------
addIssAuthInfo on mychannel ledger
CompanyIDNO 	string
APIKey      	string
BizCD      		string
WebID       	string
WebPWD      	string
-------------------------------------------------------------------------------------- */

func addIssAuthInfo(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {

	fmt.Printf("addIssAuthInfo...")
	if len(args) != 6 {
		return shim.Error("addIssAuthInfo function needs the arguments: UUID, CompanyIDNO, APIKey, BizCD, WebID, WebPWD")
	}

	key, err := APIStub.CreateCompositeKey(PREFIX_ISS_AUTHINFO, []string{args[0]})

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("TxID ==> " + APIStub.GetTxID())

	issAuthInfo := IssAuthInfo{
		CompanyIDNO: args[1],
		APIKey:      args[2],
		BizCD:       args[3],
		WebID:       args[4],
		WebPWD:      args[5],
	}

	fmt.Printf("%#v", issAuthInfo)

	issAuthInfoAsBytes, err := json.Marshal(issAuthInfo)

	err = APIStub.PutState(key, issAuthInfoAsBytes)

	if err != nil {
		fmt.Printf("addIssAuthInfo APIStub.PutState() is error.")
	}

	fmt.Printf("IssAuthInfo has been added.")

	return shim.Success(nil)
}

/* -----------------------------------------------------------------------------------------
queryIssAuthInfo
CompanyIDNO 	string
APIKey      	string
BizCD      		string
WebID       	string
WebPWD      	string
------------------------------------------------------------------------------------------ */

func queryIssAuthInfo(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("queryIssAuthInfo...")
	if len(args) != 1 {
		return shim.Error("queryIssAuthInfo  function needs the arguments: UUID")
	}

	key, err := APIStub.CreateCompositeKey(PREFIX_ISS_AUTHINFO, []string{args[0]})

	if err != nil {
		return shim.Error(err.Error())
	}

	issAuthInfoAsBytes, err := APIStub.GetState(key)

	issAuthInfo := &IssAuthInfo{}

	json.Unmarshal(issAuthInfoAsBytes, &issAuthInfo)

	issAuthInfo.UUID = args[0]

	issAuthInfoAsBytes, err = json.Marshal(issAuthInfo)

	if err != nil {
		fmt.Printf("APIStub.GetState is error")
		return shim.Error(err.Error())
	}

	return shim.Success(issAuthInfoAsBytes)
}

/* -----------------------------------------------------------------------------------------
queryAllIssAuthInfo
------------------------------------------------------------------------------------------ */

func queryAllIssAuthInfo(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("queryAllIssAuthinfo...")

	resultsIterator, err := APIStub.GetStateByPartialCompositeKey(PREFIX_ISS_AUTHINFO, []string{})

	if err != nil {
		return shim.Error(err.Error())
	}

	defer resultsIterator.Close()

	results := []interface{}{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()

		if err != nil {
			return shim.Error(err.Error())
		}

		issAuthInfo := IssAuthInfo{}

		err = json.Unmarshal(kvResult.Value, &issAuthInfo)

		if err != nil {
			return shim.Error(err.Error())
		}

		prefix, keyParts, err := APIStub.SplitCompositeKey(kvResult.Key)

		if err != nil {
			return shim.Error(err.Error())
		}
		if len(keyParts) > 0 {
			issAuthInfo.UUID = keyParts[0]
		} else {
			issAuthInfo.UUID = prefix
		}

		results = append(results, issAuthInfo)
	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(resultsAsBytes)
}

/*
	IssCertInfo
*/
func addIssCertInfo(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {

	fmt.Printf("addIssCertInfo...")
	if len(args) != 9 {
		return shim.Error("addIssCertInfo function needs the arguments: CompanyIDNO, APIKey, BizCD, CertOrg, CertName, CertDate, CertPWD, CertPath, CertData")
	}

	key, err := APIStub.CreateCompositeKey(PREFIX_ISS_CERTINFO, []string{args[1]})

	if err != nil {
		return shim.Error(err.Error())
	}

	issCertInfo := IssCertInfo{
		CompanyIDNO: args[0],
		APIKey:      args[1],
		BizCD:       args[2],
		CertOrg:     args[3],
		CertName:    args[4],
		CertDate:    args[5],
		CertPWD:     args[6],
		CertPath:    args[7],
		CertData:    args[8],
		CertData1:   "CertData1",
		CertData2:   "CertData2",
		CertData3:   "CertData3",
		CertData4:   "CertData4",
		CertData5:   "CertData5",
		CertData6:   "CertData6",
		CertData7:   "CertData7",
		CertData8:   "CertData8",
		CertData9:   "CertData9",
	}

	fmt.Printf("%#v", issCertInfo)

	issCertInfoAsBytes, err := json.Marshal(issCertInfo)

	err = APIStub.PutState(key, issCertInfoAsBytes)

	if err != nil {
		fmt.Printf("addIssCertInfo APIStub.PutState() is error.")
	}

	fmt.Printf("IssCertInfo has been added.")

	return shim.Success(nil)
}

func queryIssCertInfo(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {

	fmt.Printf("queryIssCertInfo...")
	if len(args) != 1 {
		return shim.Error("queryIssCertInfo  function needs the arguments: APIKey")
	}

	key, err := APIStub.CreateCompositeKey(PREFIX_ISS_CERTINFO, []string{args[0]})

	if err != nil {
		return shim.Error(err.Error())
	}

	issCertInfoAsBytes, err := APIStub.GetState(key)

	if err != nil {
		fmt.Printf("APIStub.GetState is error")
		return shim.Error(err.Error())
	}

	return shim.Success(issCertInfoAsBytes)
}

func queryAllIssCertInfo(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("queryAllIssCertInfo...")

	resultsIterator, err := APIStub.GetStateByPartialCompositeKey(PREFIX_ISS_CERTINFO, []string{})

	if err != nil {
		return shim.Error(err.Error())
	}

	defer resultsIterator.Close()

	results := []interface{}{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()

		if err != nil {
			return shim.Error(err.Error())
		}

		issCertInfo := IssCertInfo{}

		err = json.Unmarshal(kvResult.Value, &issCertInfo)

		if err != nil {
			return shim.Error(err.Error())
		}

		results = append(results, issCertInfo)
	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(resultsAsBytes)
}


var blockchainFunctions = map[string]func(shim.ChaincodeStubInterface, []string) peer.Response{
	// Coocon Smart Contract Functions
	"addIssAuthInfo":      addIssAuthInfo,
	"queryIssAuthInfo":    queryIssAuthInfo,
	"queryAllIssAuthInfo": queryAllIssAuthInfo,

	// Webcash Smart Contract Functions
	"addIssCertInfo":      addIssCertInfo,
	"queryIssCertInfo":    queryIssCertInfo,
	"queryAllIssCertInfo": queryAllIssCertInfo,
}

/*------------------------------------------------------------------------------
SmartContract.Init Method is called when installing and instantiating the chaincode

*/
func (s *SmartContract) Init(APIStub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("========================================= Init Smart Contract =====================================")
	return shim.Success(nil)
}

/*---------------------------------------------------------------------------------
SmartContract.Invoke Method is called when invoking the chaincode
----------------------------------------------------------------------------------*/
func (s *SmartContract) Invoke(APIStub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("========================================= Invoke Smart Contract ======================================")

	function, args := APIStub.GetFunctionAndParameters()

	if function == "init" {
		return s.Init(APIStub)
	}

	// Invoking blockchainFunctions Map
	blockchainFunctions := blockchainFunctions[function]
	if blockchainFunctions == nil {
		return shim.Error("Invalid invoke function.")
	}
	return blockchainFunctions(APIStub, args)
}

/*---------------------------------------------------------------
main function is used for starting the SmartContract
----------------------------------------------------------------*/
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))

	// Check if have any errors
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
	fmt.Printf("SmartContract is starting successfully.")
}
