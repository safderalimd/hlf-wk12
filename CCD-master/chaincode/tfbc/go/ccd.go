/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Trade Finance Use Case - WORK IN  PROGRESS
 */

 package main


 import (
	 "bytes"
	 "encoding/json"
	 "fmt"
	 "strconv"
	 "time"
 
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
 )
 
 // Define the Smart Contract structure
 type SmartContract struct {
 }
 
 
 // Define the letter of credit
 //type LetterOfCredit struct {
 type CurbCounterfiet struct {
	 LCId            string      `json:"lcId"`
	 ExpiryDate      string      `json:"expiryDate"`
	 Distributor    string   `json:"Distributor"`
	 Manufacturer        string      `json:"Manufacturer"`
	 Dealer      string      `json:"Dealer"`
	 Amount          int     `json:"amount,int"`
	 Status          string      `json:"status"`
	 DrugName        string      `json:"drugName"`
	 BatchNumber     string      `json:"batchNumber"`
	 MfgLicNo        string      `json:"mfgLicNo"`
	 ManufacturingDate       string      `json:"manufacturingDate"`
 }
 
 
 func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 // Retrieve the requested Smart Contract function and arguments
	 function, args := APIstub.GetFunctionAndParameters()
	 // Route to the appropriate handler function to interact with the ledger appropriately
	 if function == "requestDetails" {
		 return s.requestDetails(APIstub, args)
	 } else if function == "manufacturingDetails" {
		 return s.manufacturingDetails(APIstub, args)
	 } else if function == "acceptOrder" {
		 return s.acceptOrder(APIstub, args)
	 }else if function == "getDetails" {
		 return s.getDetails(APIstub, args)
	 }else if function == "getDetailsHistory" {
		 return s.getDetailsHistory(APIstub, args)
	 }
 
	 return shim.Error("Invalid Smart Contract function name.")
 }
 
 
 
 
 
 // This function is initiate by Distributor 
 func (s *SmartContract) requestDetails(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
 
	 CC := CurbCounterfiet{}
 
	 err  := json.Unmarshal([]byte(args[0]),&CC)
 if err != nil {
		 return shim.Error("Not able to parse args into CC")
	 }
	 LCBytes, err := json.Marshal(CC)
	 APIstub.PutState(CC.Distributor,LCBytes)
	 fmt.Println("Details Requested -> ", CC)
 
	 
 
	 return shim.Success(nil)
 }
 
 // This function is initiate by Manufacturer //dealer
 func (s *SmartContract) issueLC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 lcID := struct {
		 LcID  string `json:"lcID"`
	 }{}
	 err  := json.Unmarshal([]byte(args[0]),&lcID)
	 if err != nil {
		 return shim.Error("Not able to parse args into LCID")
	 }
	 
	 // if err != nil {
	 //  return shim.Error("No Amount")
	 // }
 
	 LCAsBytes, _ := APIstub.GetState(lcID.LcID)
 
	 var lc CurbCounterfiet
 
	 err = json.Unmarshal(LCAsBytes, &lc)
 
	 if err != nil {
		 return shim.Error("Issue with LC json unmarshaling")
	 }
 
 
	 //LC := CurbCounterfiet{LCId: lc.LCId, ExpiryDate: lc.ExpiryDate, Distributor: lc.Distributor, Manufacturer: lc.Manufacturer, Dealer: lc.Dealer, Amount: lc.Amount, Status: "Issued"}
	 LC := CurbCounterfiet{LCId: lc.LCId, ExpiryDate: lc.ExpiryDate, Manufacturer: lc.Manufacturer, Amount: lc.Amount, Status: "Issued",lc.DrugName ,lc.BatchNumber,lc.MfgLicNo,lc.ManufacturingDate}
	 LCBytes, err := json.Marshal(LC)
 
	 if err != nil {
		 return shim.Error("Issue with LC json marshaling")
	 }
 
	 APIstub.PutState(lc.LCId,LCBytes)
	 fmt.Println("LC Issued -> ", LC)
 
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) acceptOrder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 lcID := struct {
		 LcID  string `json:"lcID"`
	 }{}
	 err  := json.Unmarshal([]byte(args[0]),&lcID)
	 if err != nil {
		 return shim.Error("Not able to parse args into LC")
	 }
 
	 LCAsBytes, _ := APIstub.GetState(lcID.LcID)
 
	 var lc CurbCounterfiet
 
	 err = json.Unmarshal(LCAsBytes, &lc)
 
	 if err != nil {
		 return shim.Error("Issue with LC json unmarshaling")
	 }
 
 
	 //LC := CurbCounterfiet{LCId: lc.LCId, ExpiryDate: lc.ExpiryDate, Distributor: lc.Distributor, Manufacturer: lc.Manufacturer, Dealer: lc.Dealer, Amount: lc.Amount, Status: "Accepted"}
	 LC := CurbCounterfiet{LCId: lc.LCId, ExpiryDate: lc.ExpiryDate, Manufacturer: lc.Manufacturer, Amount: lc.Amount, Status: "DistributorAccepted",lc.DrugName ,lc.BatchNumber,lc.MfgLicNo,lc.ManufacturingDate}
	 LCBytes, err := json.Marshal(LC)
 
	 if err != nil {
		 return shim.Error("Issue with LC json marshaling")
	 }
 
	 APIstub.PutState(lc.LCId,LCBytes)
	 fmt.Println("LC Accepted -> ", LC)
 
 
	 
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) getDetails(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 lcId := args[0];
	 
	 // if err != nil {
	 //  return shim.Error("No Amount")
	 // }
 
	 LCAsBytes, _ := APIstub.GetState(lcId)
 
	 return shim.Success(LCAsBytes)
 }
 
 func (s *SmartContract) getDetailsHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 lcId := args[0];
	 
	 
 
	 resultsIterator, err := APIstub.GetHistoryForKey(lcId)
	 if err != nil {
		 return shim.Error("Error retrieving LC history.")
	 }
	 defer resultsIterator.Close()
 
	 // buffer is a JSON array containing historic values for the marble
	 var buffer bytes.Buffer
	 buffer.WriteString("[")
 
	 bArrayMemberAlreadyWritten := false
	 for resultsIterator.HasNext() {
		 response, err := resultsIterator.Next()
		 if err != nil {
			 return shim.Error("Error retrieving LC history.")
		 }
		 // Add a comma before array members, suppress it for the first array member
		 if bArrayMemberAlreadyWritten == true {
			 buffer.WriteString(",")
		 }
		 buffer.WriteString("{\"TxId\":")
		 buffer.WriteString("\"")
		 buffer.WriteString(response.TxId)
		 buffer.WriteString("\"")
 
		 buffer.WriteString(", \"Value\":")
		 // if it was a delete operation on given key, then we need to set the
		 //corresponding value null. Else, we will write the response.Value
		 //as-is (as the Value itself a JSON marble)
		 if response.IsDelete {
			 buffer.WriteString("null")
		 } else {
			 buffer.WriteString(string(response.Value))
		 }
 
		 buffer.WriteString(", \"Timestamp\":")
		 buffer.WriteString("\"")
		 buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		 buffer.WriteString("\"")
 
		 buffer.WriteString(", \"IsDelete\":")
		 buffer.WriteString("\"")
		 buffer.WriteString(strconv.FormatBool(response.IsDelete))
		 buffer.WriteString("\"")
 
		 buffer.WriteString("}")
		 bArrayMemberAlreadyWritten = true
	 }
	 buffer.WriteString("]")
 
	 fmt.Printf("- getDetailsHistory returning:\n%s\n", buffer.String())
 
	 
 
	 return shim.Success(buffer.Bytes())
 }
 
 // The main function is only relevant in unit test mode. Only included here for completeness.
 func main() {
 
	 // Create a new Smart Contract
	 err := shim.Start(new(SmartContract))
	 if err != nil {
		 fmt.Printf("Error creating new Smart Contract: %s", err)
	 }
 }