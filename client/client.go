// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type VerifyBankElementRequest struct {
	SceneId      *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	Mode         *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	BankCardNo   *string `json:"BankCardNo,omitempty" xml:"BankCardNo,omitempty"`
	IdNo         *string `json:"IdNo,omitempty" xml:"IdNo,omitempty"`
	BankCardUrl  *string `json:"BankCardUrl,omitempty" xml:"BankCardUrl,omitempty"`
	BankCardFile *string `json:"BankCardFile,omitempty" xml:"BankCardFile,omitempty"`
	IdName       *string `json:"IdName,omitempty" xml:"IdName,omitempty"`
	Mobile       *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s VerifyBankElementRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyBankElementRequest) GoString() string {
	return s.String()
}

func (s *VerifyBankElementRequest) SetSceneId(v int64) *VerifyBankElementRequest {
	s.SceneId = &v
	return s
}

func (s *VerifyBankElementRequest) SetOuterOrderNo(v string) *VerifyBankElementRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *VerifyBankElementRequest) SetMode(v string) *VerifyBankElementRequest {
	s.Mode = &v
	return s
}

func (s *VerifyBankElementRequest) SetBankCardNo(v string) *VerifyBankElementRequest {
	s.BankCardNo = &v
	return s
}

func (s *VerifyBankElementRequest) SetIdNo(v string) *VerifyBankElementRequest {
	s.IdNo = &v
	return s
}

func (s *VerifyBankElementRequest) SetBankCardUrl(v string) *VerifyBankElementRequest {
	s.BankCardUrl = &v
	return s
}

func (s *VerifyBankElementRequest) SetBankCardFile(v string) *VerifyBankElementRequest {
	s.BankCardFile = &v
	return s
}

func (s *VerifyBankElementRequest) SetIdName(v string) *VerifyBankElementRequest {
	s.IdName = &v
	return s
}

func (s *VerifyBankElementRequest) SetMobile(v string) *VerifyBankElementRequest {
	s.Mobile = &v
	return s
}

type VerifyBankElementResponse struct {
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message      *string                                `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResultObject *VerifyBankElementResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s VerifyBankElementResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyBankElementResponse) GoString() string {
	return s.String()
}

func (s *VerifyBankElementResponse) SetRequestId(v string) *VerifyBankElementResponse {
	s.RequestId = &v
	return s
}

func (s *VerifyBankElementResponse) SetMessage(v string) *VerifyBankElementResponse {
	s.Message = &v
	return s
}

func (s *VerifyBankElementResponse) SetCode(v string) *VerifyBankElementResponse {
	s.Code = &v
	return s
}

func (s *VerifyBankElementResponse) SetResultObject(v *VerifyBankElementResponseResultObject) *VerifyBankElementResponse {
	s.ResultObject = v
	return s
}

type VerifyBankElementResponseResultObject struct {
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty" require:"true"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty" require:"true"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty" require:"true"`
	CertifyId    *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty" require:"true"`
}

func (s VerifyBankElementResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s VerifyBankElementResponseResultObject) GoString() string {
	return s.String()
}

func (s *VerifyBankElementResponseResultObject) SetPassed(v string) *VerifyBankElementResponseResultObject {
	s.Passed = &v
	return s
}

func (s *VerifyBankElementResponseResultObject) SetSubCode(v string) *VerifyBankElementResponseResultObject {
	s.SubCode = &v
	return s
}

func (s *VerifyBankElementResponseResultObject) SetMaterialInfo(v string) *VerifyBankElementResponseResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *VerifyBankElementResponseResultObject) SetCertifyId(v string) *VerifyBankElementResponseResultObject {
	s.CertifyId = &v
	return s
}

type VerifyBankElementAdvanceRequest struct {
	BankCardFileObject io.Reader `json:"BankCardFileObject,omitempty" xml:"BankCardFileObject,omitempty" require:"true"`
	SceneId            *int64    `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	OuterOrderNo       *string   `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	Mode               *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	BankCardNo         *string   `json:"BankCardNo,omitempty" xml:"BankCardNo,omitempty"`
	IdNo               *string   `json:"IdNo,omitempty" xml:"IdNo,omitempty"`
	BankCardUrl        *string   `json:"BankCardUrl,omitempty" xml:"BankCardUrl,omitempty"`
	IdName             *string   `json:"IdName,omitempty" xml:"IdName,omitempty"`
	Mobile             *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s VerifyBankElementAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyBankElementAdvanceRequest) GoString() string {
	return s.String()
}

func (s *VerifyBankElementAdvanceRequest) SetBankCardFileObject(v io.Reader) *VerifyBankElementAdvanceRequest {
	s.BankCardFileObject = v
	return s
}

func (s *VerifyBankElementAdvanceRequest) SetSceneId(v int64) *VerifyBankElementAdvanceRequest {
	s.SceneId = &v
	return s
}

func (s *VerifyBankElementAdvanceRequest) SetOuterOrderNo(v string) *VerifyBankElementAdvanceRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *VerifyBankElementAdvanceRequest) SetMode(v string) *VerifyBankElementAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *VerifyBankElementAdvanceRequest) SetBankCardNo(v string) *VerifyBankElementAdvanceRequest {
	s.BankCardNo = &v
	return s
}

func (s *VerifyBankElementAdvanceRequest) SetIdNo(v string) *VerifyBankElementAdvanceRequest {
	s.IdNo = &v
	return s
}

func (s *VerifyBankElementAdvanceRequest) SetBankCardUrl(v string) *VerifyBankElementAdvanceRequest {
	s.BankCardUrl = &v
	return s
}

func (s *VerifyBankElementAdvanceRequest) SetIdName(v string) *VerifyBankElementAdvanceRequest {
	s.IdName = &v
	return s
}

func (s *VerifyBankElementAdvanceRequest) SetMobile(v string) *VerifyBankElementAdvanceRequest {
	s.Mobile = &v
	return s
}

type ContrastSmartVerifyRequest struct {
	SceneId       *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	OuterOrderNo  *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	Mode          *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	CertType      *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	Mobile        *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo        *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	FacePicFile   *string `json:"FacePicFile,omitempty" xml:"FacePicFile,omitempty"`
	FacePicUrl    *string `json:"FacePicUrl,omitempty" xml:"FacePicUrl,omitempty"`
	FacePicString *string `json:"FacePicString,omitempty" xml:"FacePicString,omitempty"`
}

func (s ContrastSmartVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ContrastSmartVerifyRequest) GoString() string {
	return s.String()
}

func (s *ContrastSmartVerifyRequest) SetSceneId(v int64) *ContrastSmartVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *ContrastSmartVerifyRequest) SetOuterOrderNo(v string) *ContrastSmartVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *ContrastSmartVerifyRequest) SetMode(v string) *ContrastSmartVerifyRequest {
	s.Mode = &v
	return s
}

func (s *ContrastSmartVerifyRequest) SetCertType(v string) *ContrastSmartVerifyRequest {
	s.CertType = &v
	return s
}

func (s *ContrastSmartVerifyRequest) SetMobile(v string) *ContrastSmartVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *ContrastSmartVerifyRequest) SetIp(v string) *ContrastSmartVerifyRequest {
	s.Ip = &v
	return s
}

func (s *ContrastSmartVerifyRequest) SetUserId(v string) *ContrastSmartVerifyRequest {
	s.UserId = &v
	return s
}

func (s *ContrastSmartVerifyRequest) SetCertName(v string) *ContrastSmartVerifyRequest {
	s.CertName = &v
	return s
}

func (s *ContrastSmartVerifyRequest) SetCertNo(v string) *ContrastSmartVerifyRequest {
	s.CertNo = &v
	return s
}

func (s *ContrastSmartVerifyRequest) SetFacePicFile(v string) *ContrastSmartVerifyRequest {
	s.FacePicFile = &v
	return s
}

func (s *ContrastSmartVerifyRequest) SetFacePicUrl(v string) *ContrastSmartVerifyRequest {
	s.FacePicUrl = &v
	return s
}

func (s *ContrastSmartVerifyRequest) SetFacePicString(v string) *ContrastSmartVerifyRequest {
	s.FacePicString = &v
	return s
}

type ContrastSmartVerifyResponse struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message      *string                                  `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResultObject *ContrastSmartVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s ContrastSmartVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ContrastSmartVerifyResponse) GoString() string {
	return s.String()
}

func (s *ContrastSmartVerifyResponse) SetRequestId(v string) *ContrastSmartVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *ContrastSmartVerifyResponse) SetMessage(v string) *ContrastSmartVerifyResponse {
	s.Message = &v
	return s
}

func (s *ContrastSmartVerifyResponse) SetCode(v string) *ContrastSmartVerifyResponse {
	s.Code = &v
	return s
}

func (s *ContrastSmartVerifyResponse) SetResultObject(v *ContrastSmartVerifyResponseResultObject) *ContrastSmartVerifyResponse {
	s.ResultObject = v
	return s
}

type ContrastSmartVerifyResponseResultObject struct {
	CertifyId  *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty" require:"true"`
	Passed     *string `json:"Passed,omitempty" xml:"Passed,omitempty" require:"true"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty" require:"true"`
	VerifyInfo *string `json:"VerifyInfo,omitempty" xml:"VerifyInfo,omitempty" require:"true"`
	RiskInfo   *string `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty" require:"true"`
}

func (s ContrastSmartVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s ContrastSmartVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *ContrastSmartVerifyResponseResultObject) SetCertifyId(v string) *ContrastSmartVerifyResponseResultObject {
	s.CertifyId = &v
	return s
}

func (s *ContrastSmartVerifyResponseResultObject) SetPassed(v string) *ContrastSmartVerifyResponseResultObject {
	s.Passed = &v
	return s
}

func (s *ContrastSmartVerifyResponseResultObject) SetSubCode(v string) *ContrastSmartVerifyResponseResultObject {
	s.SubCode = &v
	return s
}

func (s *ContrastSmartVerifyResponseResultObject) SetVerifyInfo(v string) *ContrastSmartVerifyResponseResultObject {
	s.VerifyInfo = &v
	return s
}

func (s *ContrastSmartVerifyResponseResultObject) SetRiskInfo(v string) *ContrastSmartVerifyResponseResultObject {
	s.RiskInfo = &v
	return s
}

type ContrastSmartVerifyAdvanceRequest struct {
	FacePicFileObject io.Reader `json:"FacePicFileObject,omitempty" xml:"FacePicFileObject,omitempty" require:"true"`
	SceneId           *int64    `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	OuterOrderNo      *string   `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	Mode              *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	CertType          *string   `json:"CertType,omitempty" xml:"CertType,omitempty"`
	Mobile            *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Ip                *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	UserId            *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	CertName          *string   `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo            *string   `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	FacePicUrl        *string   `json:"FacePicUrl,omitempty" xml:"FacePicUrl,omitempty"`
	FacePicString     *string   `json:"FacePicString,omitempty" xml:"FacePicString,omitempty"`
}

func (s ContrastSmartVerifyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ContrastSmartVerifyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ContrastSmartVerifyAdvanceRequest) SetFacePicFileObject(v io.Reader) *ContrastSmartVerifyAdvanceRequest {
	s.FacePicFileObject = v
	return s
}

func (s *ContrastSmartVerifyAdvanceRequest) SetSceneId(v int64) *ContrastSmartVerifyAdvanceRequest {
	s.SceneId = &v
	return s
}

func (s *ContrastSmartVerifyAdvanceRequest) SetOuterOrderNo(v string) *ContrastSmartVerifyAdvanceRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *ContrastSmartVerifyAdvanceRequest) SetMode(v string) *ContrastSmartVerifyAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *ContrastSmartVerifyAdvanceRequest) SetCertType(v string) *ContrastSmartVerifyAdvanceRequest {
	s.CertType = &v
	return s
}

func (s *ContrastSmartVerifyAdvanceRequest) SetMobile(v string) *ContrastSmartVerifyAdvanceRequest {
	s.Mobile = &v
	return s
}

func (s *ContrastSmartVerifyAdvanceRequest) SetIp(v string) *ContrastSmartVerifyAdvanceRequest {
	s.Ip = &v
	return s
}

func (s *ContrastSmartVerifyAdvanceRequest) SetUserId(v string) *ContrastSmartVerifyAdvanceRequest {
	s.UserId = &v
	return s
}

func (s *ContrastSmartVerifyAdvanceRequest) SetCertName(v string) *ContrastSmartVerifyAdvanceRequest {
	s.CertName = &v
	return s
}

func (s *ContrastSmartVerifyAdvanceRequest) SetCertNo(v string) *ContrastSmartVerifyAdvanceRequest {
	s.CertNo = &v
	return s
}

func (s *ContrastSmartVerifyAdvanceRequest) SetFacePicUrl(v string) *ContrastSmartVerifyAdvanceRequest {
	s.FacePicUrl = &v
	return s
}

func (s *ContrastSmartVerifyAdvanceRequest) SetFacePicString(v string) *ContrastSmartVerifyAdvanceRequest {
	s.FacePicString = &v
	return s
}

type ElementSmartVerifyRequest struct {
	SceneId               *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	OuterOrderNo          *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	Mode                  *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	CertType              *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertName              *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertUrl               *string `json:"CertUrl,omitempty" xml:"CertUrl,omitempty"`
	CertFile              *string `json:"CertFile,omitempty" xml:"CertFile,omitempty"`
	CertNationalEmblemUrl *string `json:"CertNationalEmblemUrl,omitempty" xml:"CertNationalEmblemUrl,omitempty"`
}

func (s ElementSmartVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ElementSmartVerifyRequest) GoString() string {
	return s.String()
}

func (s *ElementSmartVerifyRequest) SetSceneId(v int64) *ElementSmartVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *ElementSmartVerifyRequest) SetOuterOrderNo(v string) *ElementSmartVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *ElementSmartVerifyRequest) SetMode(v string) *ElementSmartVerifyRequest {
	s.Mode = &v
	return s
}

func (s *ElementSmartVerifyRequest) SetCertType(v string) *ElementSmartVerifyRequest {
	s.CertType = &v
	return s
}

func (s *ElementSmartVerifyRequest) SetCertName(v string) *ElementSmartVerifyRequest {
	s.CertName = &v
	return s
}

func (s *ElementSmartVerifyRequest) SetCertNo(v string) *ElementSmartVerifyRequest {
	s.CertNo = &v
	return s
}

func (s *ElementSmartVerifyRequest) SetCertUrl(v string) *ElementSmartVerifyRequest {
	s.CertUrl = &v
	return s
}

func (s *ElementSmartVerifyRequest) SetCertFile(v string) *ElementSmartVerifyRequest {
	s.CertFile = &v
	return s
}

func (s *ElementSmartVerifyRequest) SetCertNationalEmblemUrl(v string) *ElementSmartVerifyRequest {
	s.CertNationalEmblemUrl = &v
	return s
}

type ElementSmartVerifyResponse struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResultObject *ElementSmartVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s ElementSmartVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ElementSmartVerifyResponse) GoString() string {
	return s.String()
}

func (s *ElementSmartVerifyResponse) SetRequestId(v string) *ElementSmartVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *ElementSmartVerifyResponse) SetMessage(v string) *ElementSmartVerifyResponse {
	s.Message = &v
	return s
}

func (s *ElementSmartVerifyResponse) SetCode(v string) *ElementSmartVerifyResponse {
	s.Code = &v
	return s
}

func (s *ElementSmartVerifyResponse) SetResultObject(v *ElementSmartVerifyResponseResultObject) *ElementSmartVerifyResponse {
	s.ResultObject = v
	return s
}

type ElementSmartVerifyResponseResultObject struct {
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty" require:"true"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty" require:"true"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty" require:"true"`
	CertifyId    *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty" require:"true"`
}

func (s ElementSmartVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s ElementSmartVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *ElementSmartVerifyResponseResultObject) SetPassed(v string) *ElementSmartVerifyResponseResultObject {
	s.Passed = &v
	return s
}

func (s *ElementSmartVerifyResponseResultObject) SetSubCode(v string) *ElementSmartVerifyResponseResultObject {
	s.SubCode = &v
	return s
}

func (s *ElementSmartVerifyResponseResultObject) SetMaterialInfo(v string) *ElementSmartVerifyResponseResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *ElementSmartVerifyResponseResultObject) SetCertifyId(v string) *ElementSmartVerifyResponseResultObject {
	s.CertifyId = &v
	return s
}

type ElementSmartVerifyAdvanceRequest struct {
	CertFileObject        io.Reader `json:"CertFileObject,omitempty" xml:"CertFileObject,omitempty" require:"true"`
	SceneId               *int64    `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	OuterOrderNo          *string   `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	Mode                  *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	CertType              *string   `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertName              *string   `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo                *string   `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	CertUrl               *string   `json:"CertUrl,omitempty" xml:"CertUrl,omitempty"`
	CertNationalEmblemUrl *string   `json:"CertNationalEmblemUrl,omitempty" xml:"CertNationalEmblemUrl,omitempty"`
}

func (s ElementSmartVerifyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ElementSmartVerifyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ElementSmartVerifyAdvanceRequest) SetCertFileObject(v io.Reader) *ElementSmartVerifyAdvanceRequest {
	s.CertFileObject = v
	return s
}

func (s *ElementSmartVerifyAdvanceRequest) SetSceneId(v int64) *ElementSmartVerifyAdvanceRequest {
	s.SceneId = &v
	return s
}

func (s *ElementSmartVerifyAdvanceRequest) SetOuterOrderNo(v string) *ElementSmartVerifyAdvanceRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *ElementSmartVerifyAdvanceRequest) SetMode(v string) *ElementSmartVerifyAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *ElementSmartVerifyAdvanceRequest) SetCertType(v string) *ElementSmartVerifyAdvanceRequest {
	s.CertType = &v
	return s
}

func (s *ElementSmartVerifyAdvanceRequest) SetCertName(v string) *ElementSmartVerifyAdvanceRequest {
	s.CertName = &v
	return s
}

func (s *ElementSmartVerifyAdvanceRequest) SetCertNo(v string) *ElementSmartVerifyAdvanceRequest {
	s.CertNo = &v
	return s
}

func (s *ElementSmartVerifyAdvanceRequest) SetCertUrl(v string) *ElementSmartVerifyAdvanceRequest {
	s.CertUrl = &v
	return s
}

func (s *ElementSmartVerifyAdvanceRequest) SetCertNationalEmblemUrl(v string) *ElementSmartVerifyAdvanceRequest {
	s.CertNationalEmblemUrl = &v
	return s
}

type InitSmartVerifyRequest struct {
	SceneId           *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	OuterOrderNo      *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	Mode              *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	CertType          *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	MetaInfo          *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	Mobile            *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Ip                *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	CertName          *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertNo            *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	Ocr               *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	CallbackUrl       *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	CallbackToken     *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	FacePictureUrl    *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	CertifyId         *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	OssBucketName     *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectName     *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	IdNo              *string `json:"IdNo,omitempty" xml:"IdNo,omitempty"`
	IdName            *string `json:"IdName,omitempty" xml:"IdName,omitempty"`
}

func (s InitSmartVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s InitSmartVerifyRequest) GoString() string {
	return s.String()
}

func (s *InitSmartVerifyRequest) SetSceneId(v int64) *InitSmartVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *InitSmartVerifyRequest) SetOuterOrderNo(v string) *InitSmartVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *InitSmartVerifyRequest) SetMode(v string) *InitSmartVerifyRequest {
	s.Mode = &v
	return s
}

func (s *InitSmartVerifyRequest) SetCertType(v string) *InitSmartVerifyRequest {
	s.CertType = &v
	return s
}

func (s *InitSmartVerifyRequest) SetMetaInfo(v string) *InitSmartVerifyRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitSmartVerifyRequest) SetMobile(v string) *InitSmartVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *InitSmartVerifyRequest) SetIp(v string) *InitSmartVerifyRequest {
	s.Ip = &v
	return s
}

func (s *InitSmartVerifyRequest) SetUserId(v string) *InitSmartVerifyRequest {
	s.UserId = &v
	return s
}

func (s *InitSmartVerifyRequest) SetCertName(v string) *InitSmartVerifyRequest {
	s.CertName = &v
	return s
}

func (s *InitSmartVerifyRequest) SetCertNo(v string) *InitSmartVerifyRequest {
	s.CertNo = &v
	return s
}

func (s *InitSmartVerifyRequest) SetOcr(v string) *InitSmartVerifyRequest {
	s.Ocr = &v
	return s
}

func (s *InitSmartVerifyRequest) SetCallbackUrl(v string) *InitSmartVerifyRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitSmartVerifyRequest) SetCallbackToken(v string) *InitSmartVerifyRequest {
	s.CallbackToken = &v
	return s
}

func (s *InitSmartVerifyRequest) SetFacePictureBase64(v string) *InitSmartVerifyRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *InitSmartVerifyRequest) SetFacePictureUrl(v string) *InitSmartVerifyRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *InitSmartVerifyRequest) SetCertifyId(v string) *InitSmartVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *InitSmartVerifyRequest) SetOssBucketName(v string) *InitSmartVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *InitSmartVerifyRequest) SetOssObjectName(v string) *InitSmartVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *InitSmartVerifyRequest) SetIdNo(v string) *InitSmartVerifyRequest {
	s.IdNo = &v
	return s
}

func (s *InitSmartVerifyRequest) SetIdName(v string) *InitSmartVerifyRequest {
	s.IdName = &v
	return s
}

type InitSmartVerifyResponse struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message      *string                              `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code         *string                              `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResultObject *InitSmartVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s InitSmartVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s InitSmartVerifyResponse) GoString() string {
	return s.String()
}

func (s *InitSmartVerifyResponse) SetRequestId(v string) *InitSmartVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *InitSmartVerifyResponse) SetMessage(v string) *InitSmartVerifyResponse {
	s.Message = &v
	return s
}

func (s *InitSmartVerifyResponse) SetCode(v string) *InitSmartVerifyResponse {
	s.Code = &v
	return s
}

func (s *InitSmartVerifyResponse) SetResultObject(v *InitSmartVerifyResponseResultObject) *InitSmartVerifyResponse {
	s.ResultObject = v
	return s
}

type InitSmartVerifyResponseResultObject struct {
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty" require:"true"`
}

func (s InitSmartVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s InitSmartVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *InitSmartVerifyResponseResultObject) SetCertifyId(v string) *InitSmartVerifyResponseResultObject {
	s.CertifyId = &v
	return s
}

type DescribeSmartVerifyRequest struct {
	SceneId           *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	CertifyId         *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	PictureReturnType *string `json:"PictureReturnType,omitempty" xml:"PictureReturnType,omitempty"`
}

func (s DescribeSmartVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmartVerifyRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmartVerifyRequest) SetSceneId(v int64) *DescribeSmartVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeSmartVerifyRequest) SetCertifyId(v string) *DescribeSmartVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *DescribeSmartVerifyRequest) SetPictureReturnType(v string) *DescribeSmartVerifyRequest {
	s.PictureReturnType = &v
	return s
}

type DescribeSmartVerifyResponse struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message      *string                                  `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ResultObject *DescribeSmartVerifyResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s DescribeSmartVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmartVerifyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmartVerifyResponse) SetRequestId(v string) *DescribeSmartVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSmartVerifyResponse) SetMessage(v string) *DescribeSmartVerifyResponse {
	s.Message = &v
	return s
}

func (s *DescribeSmartVerifyResponse) SetCode(v string) *DescribeSmartVerifyResponse {
	s.Code = &v
	return s
}

func (s *DescribeSmartVerifyResponse) SetResultObject(v *DescribeSmartVerifyResponseResultObject) *DescribeSmartVerifyResponse {
	s.ResultObject = v
	return s
}

type DescribeSmartVerifyResponseResultObject struct {
	Passed       *string  `json:"Passed,omitempty" xml:"Passed,omitempty" require:"true"`
	SubCode      *string  `json:"SubCode,omitempty" xml:"SubCode,omitempty" require:"true"`
	MaterialInfo *string  `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty" require:"true"`
	PassedScore  *float32 `json:"PassedScore,omitempty" xml:"PassedScore,omitempty" require:"true"`
}

func (s DescribeSmartVerifyResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmartVerifyResponseResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSmartVerifyResponseResultObject) SetPassed(v string) *DescribeSmartVerifyResponseResultObject {
	s.Passed = &v
	return s
}

func (s *DescribeSmartVerifyResponseResultObject) SetSubCode(v string) *DescribeSmartVerifyResponseResultObject {
	s.SubCode = &v
	return s
}

func (s *DescribeSmartVerifyResponseResultObject) SetMaterialInfo(v string) *DescribeSmartVerifyResponseResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *DescribeSmartVerifyResponseResultObject) SetPassedScore(v float32) *DescribeSmartVerifyResponseResultObject {
	s.PassedScore = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudauth"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) VerifyBankElement(request *VerifyBankElementRequest, runtime *util.RuntimeOptions) (_result *VerifyBankElementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &VerifyBankElementResponse{}
	_body, _err := client.DoRequest(tea.String("VerifyBankElement"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyBankElementSimply(request *VerifyBankElementRequest) (_result *VerifyBankElementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyBankElementResponse{}
	_body, _err := client.VerifyBankElement(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyBankElementAdvance(request *VerifyBankElementAdvanceRequest, runtime *util.RuntimeOptions) (_result *VerifyBankElementResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("Cloudauth"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	verifyBankElementReq := &VerifyBankElementRequest{}
	rpcutil.Convert(request, verifyBankElementReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.BankCardFileObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	verifyBankElementReq.BankCardFile = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	verifyBankElementResp, _err := client.VerifyBankElement(verifyBankElementReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = verifyBankElementResp
	return _result, _err
}

func (client *Client) ContrastSmartVerify(request *ContrastSmartVerifyRequest, runtime *util.RuntimeOptions) (_result *ContrastSmartVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ContrastSmartVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("ContrastSmartVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ContrastSmartVerifySimply(request *ContrastSmartVerifyRequest) (_result *ContrastSmartVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContrastSmartVerifyResponse{}
	_body, _err := client.ContrastSmartVerify(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContrastSmartVerifyAdvance(request *ContrastSmartVerifyAdvanceRequest, runtime *util.RuntimeOptions) (_result *ContrastSmartVerifyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("Cloudauth"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	contrastSmartVerifyReq := &ContrastSmartVerifyRequest{}
	rpcutil.Convert(request, contrastSmartVerifyReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.FacePicFileObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	contrastSmartVerifyReq.FacePicFile = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	contrastSmartVerifyResp, _err := client.ContrastSmartVerify(contrastSmartVerifyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = contrastSmartVerifyResp
	return _result, _err
}

func (client *Client) ElementSmartVerify(request *ElementSmartVerifyRequest, runtime *util.RuntimeOptions) (_result *ElementSmartVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ElementSmartVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("ElementSmartVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ElementSmartVerifySimply(request *ElementSmartVerifyRequest) (_result *ElementSmartVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ElementSmartVerifyResponse{}
	_body, _err := client.ElementSmartVerify(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ElementSmartVerifyAdvance(request *ElementSmartVerifyAdvanceRequest, runtime *util.RuntimeOptions) (_result *ElementSmartVerifyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("Cloudauth"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	elementSmartVerifyReq := &ElementSmartVerifyRequest{}
	rpcutil.Convert(request, elementSmartVerifyReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.CertFileObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	elementSmartVerifyReq.CertFile = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	elementSmartVerifyResp, _err := client.ElementSmartVerify(elementSmartVerifyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = elementSmartVerifyResp
	return _result, _err
}

func (client *Client) InitSmartVerify(request *InitSmartVerifyRequest, runtime *util.RuntimeOptions) (_result *InitSmartVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InitSmartVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("InitSmartVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitSmartVerifySimply(request *InitSmartVerifyRequest) (_result *InitSmartVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitSmartVerifyResponse{}
	_body, _err := client.InitSmartVerify(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSmartVerify(request *DescribeSmartVerifyRequest, runtime *util.RuntimeOptions) (_result *DescribeSmartVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSmartVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSmartVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSmartVerifySimply(request *DescribeSmartVerifyRequest) (_result *DescribeSmartVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSmartVerifyResponse{}
	_body, _err := client.DescribeSmartVerify(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
