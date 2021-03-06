// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

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

type ContrastSmartVerifyResponseBody struct {
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *ContrastSmartVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s ContrastSmartVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContrastSmartVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *ContrastSmartVerifyResponseBody) SetCode(v string) *ContrastSmartVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *ContrastSmartVerifyResponseBody) SetMessage(v string) *ContrastSmartVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *ContrastSmartVerifyResponseBody) SetRequestId(v string) *ContrastSmartVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContrastSmartVerifyResponseBody) SetResultObject(v *ContrastSmartVerifyResponseBodyResultObject) *ContrastSmartVerifyResponseBody {
	s.ResultObject = v
	return s
}

type ContrastSmartVerifyResponseBodyResultObject struct {
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	CertifyId  *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	VerifyInfo *string `json:"VerifyInfo,omitempty" xml:"VerifyInfo,omitempty"`
	RiskInfo   *string `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty"`
	Passed     *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
}

func (s ContrastSmartVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s ContrastSmartVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *ContrastSmartVerifyResponseBodyResultObject) SetSubCode(v string) *ContrastSmartVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *ContrastSmartVerifyResponseBodyResultObject) SetCertifyId(v string) *ContrastSmartVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *ContrastSmartVerifyResponseBodyResultObject) SetVerifyInfo(v string) *ContrastSmartVerifyResponseBodyResultObject {
	s.VerifyInfo = &v
	return s
}

func (s *ContrastSmartVerifyResponseBodyResultObject) SetRiskInfo(v string) *ContrastSmartVerifyResponseBodyResultObject {
	s.RiskInfo = &v
	return s
}

func (s *ContrastSmartVerifyResponseBodyResultObject) SetPassed(v string) *ContrastSmartVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

type ContrastSmartVerifyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ContrastSmartVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ContrastSmartVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ContrastSmartVerifyResponse) GoString() string {
	return s.String()
}

func (s *ContrastSmartVerifyResponse) SetHeaders(v map[string]*string) *ContrastSmartVerifyResponse {
	s.Headers = v
	return s
}

func (s *ContrastSmartVerifyResponse) SetBody(v *ContrastSmartVerifyResponseBody) *ContrastSmartVerifyResponse {
	s.Body = v
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

type DescribeSmartVerifyResponseBody struct {
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DescribeSmartVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeSmartVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmartVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmartVerifyResponseBody) SetCode(v string) *DescribeSmartVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSmartVerifyResponseBody) SetMessage(v string) *DescribeSmartVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSmartVerifyResponseBody) SetRequestId(v string) *DescribeSmartVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmartVerifyResponseBody) SetResultObject(v *DescribeSmartVerifyResponseBodyResultObject) *DescribeSmartVerifyResponseBody {
	s.ResultObject = v
	return s
}

type DescribeSmartVerifyResponseBodyResultObject struct {
	SubCode      *string  `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	PassedScore  *float32 `json:"PassedScore,omitempty" xml:"PassedScore,omitempty"`
	MaterialInfo *string  `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	Passed       *string  `json:"Passed,omitempty" xml:"Passed,omitempty"`
}

func (s DescribeSmartVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmartVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSmartVerifyResponseBodyResultObject) SetSubCode(v string) *DescribeSmartVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *DescribeSmartVerifyResponseBodyResultObject) SetPassedScore(v float32) *DescribeSmartVerifyResponseBodyResultObject {
	s.PassedScore = &v
	return s
}

func (s *DescribeSmartVerifyResponseBodyResultObject) SetMaterialInfo(v string) *DescribeSmartVerifyResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *DescribeSmartVerifyResponseBodyResultObject) SetPassed(v string) *DescribeSmartVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

type DescribeSmartVerifyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSmartVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSmartVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmartVerifyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmartVerifyResponse) SetHeaders(v map[string]*string) *DescribeSmartVerifyResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmartVerifyResponse) SetBody(v *DescribeSmartVerifyResponseBody) *DescribeSmartVerifyResponse {
	s.Body = v
	return s
}

type DescribeSmsDetailRequest struct {
	Mobile       *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	SendDate     *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	SendStatus   *string `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	SignName     *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeSmsDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmsDetailRequest) SetMobile(v string) *DescribeSmsDetailRequest {
	s.Mobile = &v
	return s
}

func (s *DescribeSmsDetailRequest) SetSendDate(v string) *DescribeSmsDetailRequest {
	s.SendDate = &v
	return s
}

func (s *DescribeSmsDetailRequest) SetErrorCode(v string) *DescribeSmsDetailRequest {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSmsDetailRequest) SetOuterOrderNo(v string) *DescribeSmsDetailRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *DescribeSmsDetailRequest) SetSendStatus(v string) *DescribeSmsDetailRequest {
	s.SendStatus = &v
	return s
}

func (s *DescribeSmsDetailRequest) SetSignName(v string) *DescribeSmsDetailRequest {
	s.SignName = &v
	return s
}

func (s *DescribeSmsDetailRequest) SetTemplateCode(v string) *DescribeSmsDetailRequest {
	s.TemplateCode = &v
	return s
}

func (s *DescribeSmsDetailRequest) SetCurrentPage(v int32) *DescribeSmsDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSmsDetailRequest) SetPageSize(v int32) *DescribeSmsDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSmsDetailRequest) SetBizId(v string) *DescribeSmsDetailRequest {
	s.BizId = &v
	return s
}

type DescribeSmsDetailResponseBody struct {
	TotalItem   *int32                                `json:"TotalItem,omitempty" xml:"TotalItem,omitempty"`
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code        *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	TotalPage   *int32                                `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	PageSize    *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Items       []*DescribeSmsDetailResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeSmsDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmsDetailResponseBody) SetTotalItem(v int32) *DescribeSmsDetailResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeSmsDetailResponseBody) SetCurrentPage(v int32) *DescribeSmsDetailResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSmsDetailResponseBody) SetRequestId(v string) *DescribeSmsDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmsDetailResponseBody) SetCode(v string) *DescribeSmsDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSmsDetailResponseBody) SetMessage(v string) *DescribeSmsDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSmsDetailResponseBody) SetTotalPage(v int32) *DescribeSmsDetailResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSmsDetailResponseBody) SetPageSize(v int32) *DescribeSmsDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSmsDetailResponseBody) SetItems(v []*DescribeSmsDetailResponseBodyItems) *DescribeSmsDetailResponseBody {
	s.Items = v
	return s
}

type DescribeSmsDetailResponseBodyItems struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ReceiveDate  *string `json:"ReceiveDate,omitempty" xml:"ReceiveDate,omitempty"`
	SendDate     *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	TaskDate     *string `json:"TaskDate,omitempty" xml:"TaskDate,omitempty"`
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Mobile       *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	SmsSize      *int32  `json:"SmsSize,omitempty" xml:"SmsSize,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	SignName     *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SendStatus   *string `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
}

func (s DescribeSmsDetailResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsDetailResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSmsDetailResponseBodyItems) SetErrorMessage(v string) *DescribeSmsDetailResponseBodyItems {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSmsDetailResponseBodyItems) SetReceiveDate(v string) *DescribeSmsDetailResponseBodyItems {
	s.ReceiveDate = &v
	return s
}

func (s *DescribeSmsDetailResponseBodyItems) SetSendDate(v string) *DescribeSmsDetailResponseBodyItems {
	s.SendDate = &v
	return s
}

func (s *DescribeSmsDetailResponseBodyItems) SetBizId(v string) *DescribeSmsDetailResponseBodyItems {
	s.BizId = &v
	return s
}

func (s *DescribeSmsDetailResponseBodyItems) SetTaskDate(v string) *DescribeSmsDetailResponseBodyItems {
	s.TaskDate = &v
	return s
}

func (s *DescribeSmsDetailResponseBodyItems) SetTemplateCode(v string) *DescribeSmsDetailResponseBodyItems {
	s.TemplateCode = &v
	return s
}

func (s *DescribeSmsDetailResponseBodyItems) SetOuterOrderNo(v string) *DescribeSmsDetailResponseBodyItems {
	s.OuterOrderNo = &v
	return s
}

func (s *DescribeSmsDetailResponseBodyItems) SetErrorCode(v string) *DescribeSmsDetailResponseBodyItems {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSmsDetailResponseBodyItems) SetMobile(v string) *DescribeSmsDetailResponseBodyItems {
	s.Mobile = &v
	return s
}

func (s *DescribeSmsDetailResponseBodyItems) SetSmsSize(v int32) *DescribeSmsDetailResponseBodyItems {
	s.SmsSize = &v
	return s
}

func (s *DescribeSmsDetailResponseBodyItems) SetContent(v string) *DescribeSmsDetailResponseBodyItems {
	s.Content = &v
	return s
}

func (s *DescribeSmsDetailResponseBodyItems) SetSignName(v string) *DescribeSmsDetailResponseBodyItems {
	s.SignName = &v
	return s
}

func (s *DescribeSmsDetailResponseBodyItems) SetSendStatus(v string) *DescribeSmsDetailResponseBodyItems {
	s.SendStatus = &v
	return s
}

type DescribeSmsDetailResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSmsDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSmsDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmsDetailResponse) SetHeaders(v map[string]*string) *DescribeSmsDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmsDetailResponse) SetBody(v *DescribeSmsDetailResponseBody) *DescribeSmsDetailResponse {
	s.Body = v
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

type ElementSmartVerifyResponseBody struct {
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *ElementSmartVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s ElementSmartVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ElementSmartVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *ElementSmartVerifyResponseBody) SetCode(v string) *ElementSmartVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *ElementSmartVerifyResponseBody) SetMessage(v string) *ElementSmartVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *ElementSmartVerifyResponseBody) SetRequestId(v string) *ElementSmartVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ElementSmartVerifyResponseBody) SetResultObject(v *ElementSmartVerifyResponseBodyResultObject) *ElementSmartVerifyResponseBody {
	s.ResultObject = v
	return s
}

type ElementSmartVerifyResponseBodyResultObject struct {
	CertifyId    *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
}

func (s ElementSmartVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s ElementSmartVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *ElementSmartVerifyResponseBodyResultObject) SetCertifyId(v string) *ElementSmartVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *ElementSmartVerifyResponseBodyResultObject) SetSubCode(v string) *ElementSmartVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *ElementSmartVerifyResponseBodyResultObject) SetMaterialInfo(v string) *ElementSmartVerifyResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *ElementSmartVerifyResponseBodyResultObject) SetPassed(v string) *ElementSmartVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

type ElementSmartVerifyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ElementSmartVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ElementSmartVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ElementSmartVerifyResponse) GoString() string {
	return s.String()
}

func (s *ElementSmartVerifyResponse) SetHeaders(v map[string]*string) *ElementSmartVerifyResponse {
	s.Headers = v
	return s
}

func (s *ElementSmartVerifyResponse) SetBody(v *ElementSmartVerifyResponseBody) *ElementSmartVerifyResponse {
	s.Body = v
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

type InitSmartVerifyResponseBody struct {
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *InitSmartVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s InitSmartVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitSmartVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *InitSmartVerifyResponseBody) SetCode(v string) *InitSmartVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *InitSmartVerifyResponseBody) SetMessage(v string) *InitSmartVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *InitSmartVerifyResponseBody) SetRequestId(v string) *InitSmartVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitSmartVerifyResponseBody) SetResultObject(v *InitSmartVerifyResponseBodyResultObject) *InitSmartVerifyResponseBody {
	s.ResultObject = v
	return s
}

type InitSmartVerifyResponseBodyResultObject struct {
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
}

func (s InitSmartVerifyResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s InitSmartVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *InitSmartVerifyResponseBodyResultObject) SetCertifyId(v string) *InitSmartVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

type InitSmartVerifyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitSmartVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitSmartVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s InitSmartVerifyResponse) GoString() string {
	return s.String()
}

func (s *InitSmartVerifyResponse) SetHeaders(v map[string]*string) *InitSmartVerifyResponse {
	s.Headers = v
	return s
}

func (s *InitSmartVerifyResponse) SetBody(v *InitSmartVerifyResponseBody) *InitSmartVerifyResponse {
	s.Body = v
	return s
}

type SendSmsRequest struct {
	Mobile        *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	SignName      *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	TemplateCode  *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateParam *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
	OuterOrderNo  *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
}

func (s SendSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s SendSmsRequest) GoString() string {
	return s.String()
}

func (s *SendSmsRequest) SetMobile(v string) *SendSmsRequest {
	s.Mobile = &v
	return s
}

func (s *SendSmsRequest) SetSignName(v string) *SendSmsRequest {
	s.SignName = &v
	return s
}

func (s *SendSmsRequest) SetTemplateCode(v string) *SendSmsRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendSmsRequest) SetTemplateParam(v string) *SendSmsRequest {
	s.TemplateParam = &v
	return s
}

func (s *SendSmsRequest) SetOuterOrderNo(v string) *SendSmsRequest {
	s.OuterOrderNo = &v
	return s
}

type SendSmsResponseBody struct {
	Code         *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *SendSmsResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s SendSmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SendSmsResponseBody) SetCode(v string) *SendSmsResponseBody {
	s.Code = &v
	return s
}

func (s *SendSmsResponseBody) SetMessage(v string) *SendSmsResponseBody {
	s.Message = &v
	return s
}

func (s *SendSmsResponseBody) SetRequestId(v string) *SendSmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendSmsResponseBody) SetResultObject(v *SendSmsResponseBodyResultObject) *SendSmsResponseBody {
	s.ResultObject = v
	return s
}

type SendSmsResponseBodyResultObject struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s SendSmsResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s SendSmsResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *SendSmsResponseBodyResultObject) SetBizId(v string) *SendSmsResponseBodyResultObject {
	s.BizId = &v
	return s
}

type SendSmsResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendSmsResponse) String() string {
	return tea.Prettify(s)
}

func (s SendSmsResponse) GoString() string {
	return s.String()
}

func (s *SendSmsResponse) SetHeaders(v map[string]*string) *SendSmsResponse {
	s.Headers = v
	return s
}

func (s *SendSmsResponse) SetBody(v *SendSmsResponseBody) *SendSmsResponse {
	s.Body = v
	return s
}

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

type VerifyBankElementResponseBody struct {
	Code         *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *VerifyBankElementResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s VerifyBankElementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyBankElementResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyBankElementResponseBody) SetCode(v string) *VerifyBankElementResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyBankElementResponseBody) SetMessage(v string) *VerifyBankElementResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyBankElementResponseBody) SetRequestId(v string) *VerifyBankElementResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyBankElementResponseBody) SetResultObject(v *VerifyBankElementResponseBodyResultObject) *VerifyBankElementResponseBody {
	s.ResultObject = v
	return s
}

type VerifyBankElementResponseBodyResultObject struct {
	CertifyId    *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
}

func (s VerifyBankElementResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s VerifyBankElementResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *VerifyBankElementResponseBodyResultObject) SetCertifyId(v string) *VerifyBankElementResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *VerifyBankElementResponseBodyResultObject) SetSubCode(v string) *VerifyBankElementResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *VerifyBankElementResponseBodyResultObject) SetMaterialInfo(v string) *VerifyBankElementResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *VerifyBankElementResponseBodyResultObject) SetPassed(v string) *VerifyBankElementResponseBodyResultObject {
	s.Passed = &v
	return s
}

type VerifyBankElementResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyBankElementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyBankElementResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyBankElementResponse) GoString() string {
	return s.String()
}

func (s *VerifyBankElementResponse) SetHeaders(v map[string]*string) *VerifyBankElementResponse {
	s.Headers = v
	return s
}

func (s *VerifyBankElementResponse) SetBody(v *VerifyBankElementResponseBody) *VerifyBankElementResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
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

func (client *Client) ContrastSmartVerifyWithOptions(request *ContrastSmartVerifyRequest, runtime *util.RuntimeOptions) (_result *ContrastSmartVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ContrastSmartVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ContrastSmartVerify"), tea.String("2020-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ContrastSmartVerify(request *ContrastSmartVerifyRequest) (_result *ContrastSmartVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContrastSmartVerifyResponse{}
	_body, _err := client.ContrastSmartVerifyWithOptions(request, runtime)
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
	openapiutil.Convert(runtime, ossRuntime)
	contrastSmartVerifyReq := &ContrastSmartVerifyRequest{}
	openapiutil.Convert(request, contrastSmartVerifyReq)
	if !tea.BoolValue(util.IsUnset(request.FacePicFileObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	}

	contrastSmartVerifyResp, _err := client.ContrastSmartVerifyWithOptions(contrastSmartVerifyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = contrastSmartVerifyResp
	return _result, _err
}

func (client *Client) DescribeSmartVerifyWithOptions(request *DescribeSmartVerifyRequest, runtime *util.RuntimeOptions) (_result *DescribeSmartVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSmartVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSmartVerify"), tea.String("2020-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSmartVerify(request *DescribeSmartVerifyRequest) (_result *DescribeSmartVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSmartVerifyResponse{}
	_body, _err := client.DescribeSmartVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSmsDetailWithOptions(request *DescribeSmsDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeSmsDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSmsDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSmsDetail"), tea.String("2020-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSmsDetail(request *DescribeSmsDetailRequest) (_result *DescribeSmsDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSmsDetailResponse{}
	_body, _err := client.DescribeSmsDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ElementSmartVerifyWithOptions(request *ElementSmartVerifyRequest, runtime *util.RuntimeOptions) (_result *ElementSmartVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ElementSmartVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ElementSmartVerify"), tea.String("2020-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ElementSmartVerify(request *ElementSmartVerifyRequest) (_result *ElementSmartVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ElementSmartVerifyResponse{}
	_body, _err := client.ElementSmartVerifyWithOptions(request, runtime)
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
	openapiutil.Convert(runtime, ossRuntime)
	elementSmartVerifyReq := &ElementSmartVerifyRequest{}
	openapiutil.Convert(request, elementSmartVerifyReq)
	if !tea.BoolValue(util.IsUnset(request.CertFileObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	}

	elementSmartVerifyResp, _err := client.ElementSmartVerifyWithOptions(elementSmartVerifyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = elementSmartVerifyResp
	return _result, _err
}

func (client *Client) InitSmartVerifyWithOptions(request *InitSmartVerifyRequest, runtime *util.RuntimeOptions) (_result *InitSmartVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InitSmartVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InitSmartVerify"), tea.String("2020-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitSmartVerify(request *InitSmartVerifyRequest) (_result *InitSmartVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitSmartVerifyResponse{}
	_body, _err := client.InitSmartVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendSmsWithOptions(request *SendSmsRequest, runtime *util.RuntimeOptions) (_result *SendSmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendSmsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendSms"), tea.String("2020-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendSms(request *SendSmsRequest) (_result *SendSmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendSmsResponse{}
	_body, _err := client.SendSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyBankElementWithOptions(request *VerifyBankElementRequest, runtime *util.RuntimeOptions) (_result *VerifyBankElementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyBankElementResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyBankElement"), tea.String("2020-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyBankElement(request *VerifyBankElementRequest) (_result *VerifyBankElementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyBankElementResponse{}
	_body, _err := client.VerifyBankElementWithOptions(request, runtime)
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
	openapiutil.Convert(runtime, ossRuntime)
	verifyBankElementReq := &VerifyBankElementRequest{}
	openapiutil.Convert(request, verifyBankElementReq)
	if !tea.BoolValue(util.IsUnset(request.BankCardFileObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	}

	verifyBankElementResp, _err := client.VerifyBankElementWithOptions(verifyBankElementReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = verifyBankElementResp
	return _result, _err
}
